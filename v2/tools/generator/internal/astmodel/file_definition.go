/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"
	"sort"
	"strings"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

// FileDefinition is the content of a file we're generating
type FileDefinition struct {
	// the package this file is in
	packageReference InternalPackageReference
	// definitions to include in this file
	definitions []TypeDefinition

	// other generated packages whose references may be needed for code generation
	generatedPackages map[InternalPackageReference]*PackageDefinition
}

var _ GoSourceFile = &FileDefinition{}

// NewFileDefinition creates a file definition containing specified definitions.
// packageRef is the package to which this file belongs.
// definitions are the type definitions to include in this specific file.
// generatedPackages is a map of all other packages being generated (to allow for cross-package references).
func NewFileDefinition(
	packageRef InternalPackageReference,
	definitions []TypeDefinition,
	generatedPackages map[InternalPackageReference]*PackageDefinition,
) *FileDefinition {
	// Topological sort of the definitions, putting them in order of reference
	ranks := calcRanks(definitions)
	sort.Slice(definitions, func(i, j int) bool {
		iRank := ranks[definitions[i].Name()]
		jRank := ranks[definitions[j].Name()]
		if iRank != jRank {
			return iRank < jRank
		}

		// Case-insensitive sort
		iName := definitions[i].Name().Name()
		jName := definitions[j].Name().Name()

		return strings.ToLower(iName) < strings.ToLower(jName)
	})

	// TODO: check that all definitions are from same package
	return &FileDefinition{
		packageReference:  packageRef,
		definitions:       definitions,
		generatedPackages: generatedPackages,
	}
}

// Calculate the ranks for each type
// Can't use a recursive algorithm, so have to use an iterative one
func calcRanks(definitions []TypeDefinition) map[TypeName]int {
	ranks := make(map[TypeName]int)

	// First need a way to identify all the root type definers
	// These are the ones not referenced by any other in this file
	nonRoots := make(map[TypeName]bool)
	for _, d := range definitions {
		for ref := range d.References() {
			nonRoots[ref] = true
		}
	}

	// Create a queue of all the definitions we need to process
	queue := make([]TypeDefinition, 0, len(definitions))
	for _, d := range definitions {
		if _, ok := d.Type().(*ResourceType); ok {
			// Resources have rank 0
			ranks[d.Name()] = 0
		} else if _, ok := nonRoots[d.Name()]; !ok {
			// Roots have rank 0
			ranks[d.Name()] = 0
		}
		queue = append(queue, d)
	}

	lastLength := len(queue)
	for len(queue) > 0 {
		queue = assignRanks(queue, ranks)
		if len(queue) == lastLength {
			// No progress made - give everything remaining a fallback rank
			for _, d := range queue {
				ranks[d.Name()] = 10000
			}

			queue = nil
		}

		lastLength = len(queue)
	}

	return ranks
}

// assignRanks allocates ranks to any type definers whose types have known rank,
// returning a slice containing the remaining type definers for later processing
func assignRanks(definers []TypeDefinition, ranks map[TypeName]int) []TypeDefinition {
	var assignable []TypeDefinition
	var remaining []TypeDefinition

	// Partition type definers into ones we can allocate, and ones to defer. We do this before
	// making any changes to ranks to avoid iteration ordering having any impact on the ranks
	// assigned.
	for _, d := range definers {
		if _, ok := ranks[d.Name()]; ok {
			assignable = append(assignable, d)
		} else {
			remaining = append(remaining, d)
		}
	}

	// Assign ranks
	for _, d := range assignable {
		rank := ranks[d.Name()]
		for ref := range d.References() {
			if _, ok := ranks[ref]; !ok {
				// Never overwrite an existing rank
				ranks[ref] = rank + 1
			}
		}
	}

	return remaining
}

// generateImports products the definitive set of imports for use in this file
func (file *FileDefinition) generateImports() *PackageImportSet {
	allReferences := NewPackageReferenceSet()
	for _, s := range file.definitions {
		allReferences.Merge(s.RequiredPackageReferences())
	}

	// Don't need to import the current package
	allReferences.Remove(file.packageReference)

	// Create the set of imports
	requiredImports := NewPackageImportSet()
	requiredImports.AddImportsForPackageReferenceSet(allReferences)

	// TODO: Make this configurable
	requiredImports.ApplyName(MetaV1Reference, "metav1")
	requiredImports.ApplyName(APIMachineryErrorsReference, "kerrors")
	requiredImports.ApplyName(MakeSubPackageReference(ARMPackageName, file.packageReference), "arm")

	return requiredImports
}

// AsAst generates an AST node representing this file
func (file *FileDefinition) AsAst() (result *dst.File, err error) {
	// Create context from imports
	codeGenContext := NewCodeGenerationContext(
		file.packageReference,
		file.generateImports(),
		file.generatedPackages)

	// Create all definitions:
	var declarations []dst.Decl
	var errs []error
	for _, s := range file.definitions {
		decls, err := s.AsDeclarations(codeGenContext)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		declarations = append(declarations, decls...)
	}

	// If we had any errors, return them
	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	var decls []dst.Decl

	// Create import header if needed
	usedImports := codeGenContext.UsedPackageImports()
	if usedImports.Length() > 0 {
		decls = append(decls, &dst.GenDecl{
			Decs: dst.GenDeclDecorations{
				NodeDecs: dst.NodeDecs{
					After: dst.EmptyLine,
				},
			},
			Tok:   token.IMPORT,
			Specs: usedImports.AsImportSpecs(),
		})
	}

	// Add generated definitions
	decls = append(decls, declarations...)

	// Emit registration for each resource:
	var exprs []dst.Expr
	for _, defn := range file.definitions {
		if resource, ok := defn.Type().(*ResourceType); ok {
			for _, t := range resource.SchemeTypes(defn.Name()) {
				tExpr, err := t.AsTypeExpr(codeGenContext)
				if err != nil {
					return nil, eris.Wrapf(err, "creating type expression for %s", t.Name())
				}

				literal := astbuilder.NewCompositeLiteralBuilder(tExpr)
				exprs = append(exprs, astbuilder.AddrOf(literal.Build()))
			}
		}
	}

	if len(exprs) > 0 {
		decls = append(decls,
			&dst.FuncDecl{
				Type: &dst.FuncType{Params: &dst.FieldList{}},
				Name: dst.NewIdent("init"),
				Body: astbuilder.StatementBlock(
					&dst.ExprStmt{
						Decs: dst.ExprStmtDecorations{
							NodeDecs: dst.NodeDecs{
								Before: dst.NewLine,
							},
						},
						X: &dst.CallExpr{
							Fun:  dst.NewIdent("SchemeBuilder.Register"), // HACK
							Args: exprs,
						},
					},
				),
			})
	}

	var header []string
	header = append(header, CodeGenerationComments...)
	header = append(header,
		"// Copyright (c) Microsoft Corporation.",
		"// Licensed under the MIT license.")

	packageName := file.packageReference.PackageName()

	result = &dst.File{
		Decs: dst.FileDecorations{
			NodeDecs: dst.NodeDecs{
				Start: header,
				After: dst.EmptyLine,
			},
		},
		Name:  dst.NewIdent(packageName),
		Decls: decls,
	}

	return
}
