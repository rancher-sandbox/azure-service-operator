/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"path/filepath"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/reporting"
)

// ReportResourceStructureStageID is the unique identifier for this stage
const ReportResourceStructureStageID = "reportResourceStructure"

// ReportResourceStructure creates a pipeline stage that reports the structure of resources in each package
func ReportResourceStructure(configuration *config.Configuration) *Stage {
	return NewStage(
		ReportResourceStructureStageID,
		"Reports the structure of resources in each package",
		func(ctx context.Context, state *State) (*State, error) {
			report := NewResourceStructureReport(state.Definitions())
			err := report.SaveReports(configuration.FullTypesOutputPath())
			return state, err
		})
}

type ResourceStructureReport struct {
	lists map[astmodel.InternalPackageReference]astmodel.TypeDefinitionSet // A separate list of resources for each package
}

func NewResourceStructureReport(defs astmodel.TypeDefinitionSet) *ResourceStructureReport {
	result := &ResourceStructureReport{
		lists: make(map[astmodel.InternalPackageReference]astmodel.TypeDefinitionSet),
	}

	result.summarize(defs)
	return result
}

// SaveReports writes the reports to the specified files
func (report *ResourceStructureReport) SaveReports(baseFolder string) error {
	for pkg, defs := range report.lists {
		folder := filepath.Join(baseFolder, pkg.FolderPath(), "structure.txt")
		err := report.saveReport(folder, defs)
		if err != nil {
			return err
		}
	}

	return nil
}

// summarize collates a list of all definitions, grouped by package
func (report *ResourceStructureReport) summarize(definitions astmodel.TypeDefinitionSet) {
	for name, def := range definitions {
		pkg := name.InternalPackageReference()

		if _, ok := report.lists[pkg]; !ok {
			report.lists[pkg] = make(astmodel.TypeDefinitionSet)
		}

		report.lists[pkg].Add(def)
	}
}

func (report *ResourceStructureReport) saveReport(filePath string, defs astmodel.TypeDefinitionSet) error {
	rpt := reporting.NewTypeCatalogReport(defs, reporting.InlineTypes)
	rpt.AddHeader(astmodel.CodeGenerationComments...)
	err := rpt.SaveTo(filePath)
	return eris.Wrapf(err, "unable to save type catalog report to %q", filePath)
}
