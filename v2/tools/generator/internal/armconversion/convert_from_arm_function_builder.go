/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

type convertFromARMBuilder struct {
	conversionBuilder
	typedInputIdent       string
	inputIdent            string
	typeConversionBuilder *astmodel.ConversionFunctionBuilder
	locals                *astmodel.KnownLocalsSet
}

func newConvertFromARMFunctionBuilder(
	c *ARMConversionFunction,
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.InternalTypeName,
	methodName string,
) (*convertFromARMBuilder, error) {
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", receiver)
	}

	result := &convertFromARMBuilder{
		// Note: If we have a property with these names we will have a compilation issue in the generated
		// code. Right now that doesn't seem to be the case anywhere but if it does happen we may need
		// to harden this logic some to choose an unused name.
		typedInputIdent: "typedInput",
		inputIdent:      "armInput",

		conversionBuilder: conversionBuilder{
			methodName:            methodName,
			sourceType:            c.armType,
			sourceTypeName:        c.armTypeName,
			destinationType:       getReceiverObjectType(codeGenerationContext, receiver),
			destinationTypeName:   c.kubeTypeName,
			receiverIdent:         c.idFactory.CreateReceiver(receiver.Name()),
			receiverTypeExpr:      receiverExpr,
			idFactory:             c.idFactory,
			typeKind:              c.typeKind,
			codeGenerationContext: codeGenerationContext,
		},
		typeConversionBuilder: astmodel.NewConversionFunctionBuilder(c.idFactory, codeGenerationContext),
		locals:                astmodel.NewKnownLocalsSet(c.idFactory),
	}
	// Add the receiver ident into the known locals
	result.locals.Add(result.receiverIdent)

	// It's a bit awkward that there are two levels of "handler" here, but they serve different purposes:
	// The top level propertyConversionHandlers is about determining which properties are involved: given a property on
	// the destination type it determines which property (if any) on the source type will be converted to the
	// destination.
	// The "inner" handler (typeConversionBuilder) is about determining how to convert between two types: given a
	// source type and a destination type, figure out how to make the assignment work. It has no knowledge of broader
	// object structure or other properties.
	result.typeConversionBuilder.AddConversionHandlers(result.convertComplexTypeNameProperty)

	result.propertyConversionHandlers = []propertyConversionHandler{
		// Handlers for specific properties come first
		skipPropertiesFlaggedWithNoARMConversion,
		result.namePropertyHandler,
		result.ownerPropertyHandler,
		result.conditionsPropertyHandler,
		result.operatorSpecPropertyHandler,
		result.userAssignedIdentitiesPropertyHandler,
		// Generic handlers come second
		result.referencePropertyHandler,
		result.secretPropertyHandler,
		result.configMapPropertyHandler,
		result.flattenedPropertyHandler,
		result.propertiesByNameHandler,
	}

	return result, nil
}

func (builder *convertFromARMBuilder) functionDeclaration() (*dst.FuncDecl, error) {
	body, err := builder.functionBodyStatements()
	if err != nil {
		return nil, err
	}

	fn := &astbuilder.FuncDetails{
		Name:          builder.methodName,
		ReceiverIdent: builder.receiverIdent,
		ReceiverType:  astbuilder.PointerTo(builder.receiverTypeExpr),
		Body:          body,
	}

	fn.AddComments("populates a Kubernetes CRD object from an Azure ARM object")
	ownerReferenceExpr, err := astmodel.ArbitraryOwnerReference.AsTypeExpr(builder.codeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating owner reference type expression")
	}

	fn.AddParameter(
		builder.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported),
		ownerReferenceExpr)

	fn.AddParameter(builder.inputIdent, dst.NewIdent("interface{}"))
	fn.AddReturns("error")
	return fn.DefineFunc(), nil
}

func (builder *convertFromARMBuilder) functionBodyStatements() ([]dst.Stmt, error) {
	// Find all (any!) properties where their values need to be promoted from the nested ARM type
	promotions := builder.findPromotions(builder.destinationType, builder.sourceType)

	conversionStmts, err := generateTypeConversionAssignments(
		builder.sourceType,
		builder.destinationType,
		builder.propertyConversionHandler(promotions))
	if err != nil {
		return nil, eris.Wrapf(err, "unable to generate conversion statements for %s", builder.methodName)
	}

	// We remove empty statements here as they may have been used to store comments or other
	// notes about properties which were not transformed. We want to keep these statements in the
	// set of statements, but they don't count as a conversion for the purposes of actually
	// using the typedInput variable
	hasConversions := len(removeEmptyStatements(conversionStmts)) > 0

	assertStmts := builder.assertInputTypeIsARM(hasConversions)

	return astbuilder.Statements(
		assertStmts,
		conversionStmts,
		astbuilder.ReturnNoError()), nil
}

func (builder *convertFromARMBuilder) propertyConversionHandler(
	promotions map[string][]propertyPair,
) func(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) ([]dst.Stmt, error) {
	return func(
		toProp *astmodel.PropertyDefinition,
		fromType *astmodel.ObjectType,
	) ([]dst.Stmt, error) {
		result, err := builder.conversionBuilder.propertyConversionHandler(toProp, fromType)
		if err != nil {
			return nil, eris.Wrapf(err, "unable to convert property %s", toProp.PropertyName())
		}

		if len(result) == 0 {
			return nil, nil
		}

		if toProp.HasTagValue(ConversionTag, PushToOneOfLeaf) || len(promotions) == 0 {
			return result, nil
		}

		// Promote any properties from child ARM objects to the parent ARM object
		// We know the types are going to be identical, so we can just generate
		// simple assignments, but only when the property exists
		if tn, ok := astmodel.AsInternalTypeName(toProp.PropertyType()); ok {
			if availablePromotions, ok := promotions[tn.Name()]; ok {
				var promotionStmts []dst.Stmt
				for _, promotion := range availablePromotions {
					assign := astbuilder.SimpleAssignment(
						astbuilder.Selector(
							dst.NewIdent(builder.receiverIdent),
							promotion.crdProperty),
						astbuilder.Selector(
							dst.NewIdent(builder.typedInputIdent),
							string(toProp.PropertyName()),
							promotion.armProperty))
					promotionStmts = append(promotionStmts, assign)
				}

				// Heuristic/hack: If the only statement is an IfStmt, we want to append the promotions
				// within the body of the if (because it will be a guard clause avoiding nils).
				// Otherwise we just add them to the end.
				if guard, ok := result[0].(*dst.IfStmt); ok && len(result) == 1 {
					guard.Body.List = append(guard.Body.List, promotionStmts...)
				} else {
					result = append(result, promotionStmts...)
				}
			}
		}

		return result, err
	}
}

func (builder *convertFromARMBuilder) assertInputTypeIsARM(needsResult bool) []dst.Stmt {
	fmtPackage := builder.codeGenerationContext.MustGetImportedPackageName(astmodel.FmtReference)

	dest := builder.typedInputIdent
	if !needsResult {
		dest = "_" // drop result
	}

	// perform a type assert
	// <dest>, ok := <inputIdent>.(<inputIdent>)
	typeAssert := astbuilder.TypeAssert(
		dst.NewIdent(dest),
		dst.NewIdent(builder.inputIdent),
		builder.sourceTypeIdent())

	// Check the result of the type assert
	// if !ok {
	//     return fmt.Errorf("unexpected type supplied ...", <inputIdent>)
	// }
	returnIfNotOk := astbuilder.ReturnIfNotOk(
		astbuilder.FormatError(
			fmtPackage,
			fmt.Sprintf("unexpected type supplied for %s() function. Expected %s, got %%T",
				builder.methodName,
				builder.sourceTypeString()),
			dst.NewIdent(builder.inputIdent)))

	return astbuilder.Statements(typeAssert, returnIfNotOk)
}

//////////////////////
// Conversion handlers
//////////////////////

func (builder *convertFromARMBuilder) namePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {
	if builder.typeKind != TypeKindSpec || !toProp.HasName(astmodel.AzureNameProperty) {
		return notHandled, nil
	}

	// Check to make sure that the ARM object has a "Name" property (which matches our "AzureName")
	fromProp, ok := fromType.Property("Name")
	if !ok {
		return notHandled, eris.New("ARM resource missing property 'Name'")
	}

	// Invoke SetAzureName(ExtractKubernetesResourceNameFromARMName(this.Name)):
	setAzureName := astbuilder.CallQualifiedFuncAsStmt(
		builder.receiverIdent,
		"SetAzureName",
		astbuilder.CallQualifiedFunc(
			astmodel.GenRuntimeReference.PackageName(),
			"ExtractKubernetesResourceNameFromARMName",
			astbuilder.Selector(dst.NewIdent(builder.typedInputIdent), string(fromProp.PropertyName()))))

	return handleWith(setAzureName), nil
}

func (builder *convertFromARMBuilder) userAssignedIdentitiesPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {
	if _, ok := astmodel.IsUserAssignedIdentityProperty(toProp); !ok {
		return notHandled, nil
	}

	// TODO: For now we are not assigning these, as we don't know how to rebuild the reference
	return handledWithNoOp, nil
}

func (builder *convertFromARMBuilder) referencePropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {
	if !astmodel.IsTypeResourceReference(toProp.PropertyType()) {
		return notHandled, nil
	}

	// TODO: For now, we are NOT assigning to these. _Status types don't have them, and it's unclear what
	// TODO: the fromARM functions do for us on Spec types. We may need them for diffing though. If so we will
	// TODO: need to revisit this and actually assign something
	return handledWithNoOp, nil
}

func (builder *convertFromARMBuilder) secretPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {
	if !astmodel.IsTypeSecretReference(toProp.PropertyType()) {
		return notHandled, nil
	}

	// TODO: For now, we are NOT assigning to these. _Status types don't have them, and it's unclear what
	// TODO: the fromARM functions do for us on Spec types. We may need them for diffing though. If so we will
	// TODO: need to revisit this and actually assign something
	return handledWithNoOp, nil
}

func (builder *convertFromARMBuilder) configMapPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {
	isConfigMapReference := astmodel.TypeEquals(toProp.PropertyType(), astmodel.ConfigMapReferenceType)
	isConfigMapReferencePtr := astmodel.TypeEquals(toProp.PropertyType(), astmodel.OptionalConfigMapReferenceType)

	if !isConfigMapReference && !isConfigMapReferencePtr {
		return notHandled, nil
	}

	// TODO: For now, we are NOT assigning to these. _Status types don't have them, and it's unclear what
	// TODO: the fromARM functions do for us on Spec types. We may need them for diffing though. If so we will
	// TODO: need to revisit this and actually assign something
	return handledWithNoOp, nil
}

func (builder *convertFromARMBuilder) ownerPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {
	ownerParameter := builder.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported)
	ownerProp := builder.idFactory.CreatePropertyName(astmodel.OwnerProperty, astmodel.Exported)
	if toProp.PropertyName() != ownerProp || builder.typeKind != TypeKindSpec {
		return notHandled, nil
	}

	// Confirm that the destination type is the type we expect
	ownerNameType, ok := astmodel.AsTypeName(toProp.PropertyType())
	if !ok {
		var kubeDescription strings.Builder
		builder.destinationType.WriteDebugDescription(&kubeDescription, nil)

		var armDescription strings.Builder
		builder.sourceType.WriteDebugDescription(&armDescription, nil)

		return notHandled,
			eris.Errorf(
				"owner property was not of type TypeName. Kube: %s, ARM: %s",
				kubeDescription.String(),
				armDescription.String())

	}

	var convertedOwner dst.Expr
	switch ownerNameType {
	case astmodel.KnownResourceReferenceType:
		knownResourceReferenceExpr, err := astmodel.KnownResourceReferenceType.AsTypeExpr(builder.codeGenerationContext)
		if err != nil {
			return notHandled,
				eris.Wrapf(err, "creating known resource reference type expression for %s", ownerProp)
		}

		compositeLit := astbuilder.NewCompositeLiteralBuilder(knownResourceReferenceExpr)
		compositeLit.AddField("Name", astbuilder.Selector(dst.NewIdent(ownerParameter), "Name"))
		compositeLit.AddField("ARMID", astbuilder.Selector(dst.NewIdent(ownerParameter), "ARMID"))
		convertedOwner = astbuilder.AddrOf(compositeLit.Build())
	case astmodel.ArbitraryOwnerReference:
		convertedOwner = astbuilder.AddrOf(dst.NewIdent(ownerParameter))
	default:
		return notHandled,
			eris.Errorf(
				"found Owner property on spec with unexpected TypeName %s",
				ownerNameType.String())
	}

	setOwner := astbuilder.QualifiedAssignment(
		dst.NewIdent(builder.receiverIdent),
		string(toProp.PropertyName()),
		token.ASSIGN,
		convertedOwner)

	return handleWith(setOwner), nil
}

// conditionsPropertyHandler generates conversions for the "Conditions" status property. This property is set by the controller
// after each reconcile and so does not need to be preserved.
func (builder *convertFromARMBuilder) conditionsPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {
	isPropConditions := toProp.PropertyName() == builder.idFactory.CreatePropertyName(astmodel.ConditionsProperty, astmodel.Exported)
	if !isPropConditions || builder.typeKind != TypeKindStatus {
		return notHandled, nil
	}

	return handledWithNoOp, nil
}

// operatorSpecPropertyHandler generates conversions for the "OperatorSpec" property.
// TODO: This property should be copied from the "previous" spec, it can't be sourced from ARM. We'll need to come up
// TODO: with some paradigm for that if/when we start doing diffing, but for now we don't actually use FromARM with Spec types
// TODO: so just skip this
func (builder *convertFromARMBuilder) operatorSpecPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	_ *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {
	if toProp.PropertyName() != astmodel.OperatorSpecProperty || builder.typeKind != TypeKindSpec {
		return notHandled, nil
	}

	return handledWithNoOp, nil
}

// flattenedPropertyHandler generates conversions for properties that
// were flattened out from inside other properties. The code it generates will
// look something like:
//
// If 'X' is a property that was flattened:
//
//	k8sObj.Y1 = armObj.X.Y1;
//	k8sObj.Y2 = armObj.X.Y2;
//
// in reality each assignment is likely to be another conversion that is specific
// to the type being converted.
func (builder *convertFromARMBuilder) flattenedPropertyHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {
	if !toProp.WasFlattened() {
		return notHandled, nil
	}

	for _, fromProp := range fromType.Properties().Copy() {
		if toProp.WasFlattenedFrom(fromProp.PropertyName()) {
			result, err := builder.buildFlattenedAssignment(toProp, fromProp)
			if err != nil {
				return notHandled,
					eris.Wrapf(err,
						"failed to build flattened assignment for property %s",
						toProp.PropertyName())
			}

			return result, nil
		}
	}

	return notHandled,
		eris.Errorf(
			"couldn’t find source ARM property %q that k8s property %q was flattened from",
			toProp.FlattenedFrom()[0],
			toProp.PropertyName())
}

func (builder *convertFromARMBuilder) buildFlattenedAssignment(
	toProp *astmodel.PropertyDefinition,
	fromProp *astmodel.PropertyDefinition,
) (propertyConversionHandlerResult, error) {
	if len(toProp.FlattenedFrom()) > 2 {
		// this doesn't appear to happen anywhere in the JSON schemas currently

		var props []string
		for _, ff := range toProp.FlattenedFrom() {
			props = append(props, string(ff))
		}

		return notHandled,
			eris.Errorf(
				"need to implement multiple levels of flattening: property %q on %s was flattened from %q",
				toProp.PropertyName(),
				builder.receiverIdent,
				strings.Join(props, "."))

	}

	allDefs := builder.codeGenerationContext.GetAllReachableDefinitions()

	// the 'from' shape here must be:
	// 1. maybe a typename, pointing to…
	// 2. maybe optional, wrapping …
	// 3. maybe a typename, pointing to…
	// 4. an object type

	// (1.) resolve any outer typename
	fromPropType, err := allDefs.FullyResolve(fromProp.PropertyType())
	if err != nil {
		return notHandled,
			eris.Wrapf(
				err,
				"failed to resolve type for property %s",
				fromProp.PropertyName())
	}

	var fromPropObjType *astmodel.ObjectType
	var objOk bool
	// (2.) resolve any optional type
	generateNilCheck := false
	if fromPropOptType, ok := fromPropType.(*astmodel.OptionalType); ok {
		generateNilCheck = true
		// (3.) resolve any inner typename
		var elementType astmodel.Type
		elementType, err = allDefs.FullyResolve(fromPropOptType.Element())
		if err != nil {
			return notHandled,
				eris.Wrapf(
					err,
					"failed to resolve type for property %s",
					fromProp.PropertyName())
		}

		// (4.) resolve the inner object type
		fromPropObjType, objOk = elementType.(*astmodel.ObjectType)
	} else {
		// (4.) resolve the inner object type
		fromPropObjType, objOk = fromPropType.(*astmodel.ObjectType)
	}

	if !objOk {
		// see pipeline_flatten_properties.go:flattenPropType which will only flatten from (optional) object types
		return notHandled,
			eris.Errorf(
				"property %q marked as flattened from non-object type %T, which shouldn’t be possible",
				toProp.PropertyName(),
				fromPropType)
	}

	// *** Now generate the code! ***
	toPropFlattenedFrom := toProp.FlattenedFrom()
	originalPropName := toPropFlattenedFrom[len(toPropFlattenedFrom)-1]
	nestedProp, ok := fromPropObjType.Property(originalPropName)
	if !ok {
		return notHandled,
			eris.Errorf(
				"couldn't find source of flattened property %q on %s",
				toProp.PropertyName(),
				builder.receiverIdent)
	}

	// need to make a clone of builder.locals if we are going to nest in an if statement
	locals := builder.locals
	if generateNilCheck {
		locals = locals.Clone()
	}

	stmts, err := builder.typeConversionBuilder.BuildConversion(
		astmodel.ConversionParameters{
			Source:              astbuilder.Selector(dst.NewIdent(builder.typedInputIdent), string(fromProp.PropertyName()), string(originalPropName)),
			SourceType:          nestedProp.PropertyType(),
			Destination:         astbuilder.Selector(dst.NewIdent(builder.receiverIdent), string(toProp.PropertyName())),
			DestinationType:     toProp.PropertyType(),
			NameHint:            string(toProp.PropertyName()),
			ConversionContext:   nil,
			AssignmentHandler:   nil,
			Locals:              locals,
			SourceProperty:      fromProp,
			DestinationProperty: toProp,
		})
	if err != nil {
		return notHandled,
			eris.Wrapf(
				err,
				"failed to generate conversion for flattened property %s",
				toProp.PropertyName())
	}

	// we were unable to generate an inner conversion, so we cannot generate the overall conversion
	if len(stmts) == 0 {
		return notHandled, nil
	}

	if generateNilCheck {
		propToCheck := astbuilder.Selector(dst.NewIdent(builder.typedInputIdent), string(fromProp.PropertyName()))
		stmts = astbuilder.Statements(
			astbuilder.IfNotNil(propToCheck, stmts...))
	}

	astbuilder.AddComment(
		&stmts[0].Decorations().Start,
		"// copying flattened property:")

	return handleWith(stmts), nil
}

func (builder *convertFromARMBuilder) propertiesByNameHandler(
	toProp *astmodel.PropertyDefinition,
	fromType *astmodel.ObjectType,
) (propertyConversionHandlerResult, error) {
	fromProp, found := fromType.Property(toProp.PropertyName())
	if !found {
		// No direct match by name, look for renames
		if originalName, ok := toProp.Renamed(); ok {
			// toProp has been renamed, so we need to look for an ARM property with the original name
			fromProp, found = fromType.Property(originalName)
		}
	}

	if !found {
		return notHandled, nil
	}

	conversion, err := builder.typeConversionBuilder.BuildConversion(
		astmodel.ConversionParameters{
			Source:              astbuilder.Selector(dst.NewIdent(builder.typedInputIdent), string(fromProp.PropertyName())),
			SourceType:          fromProp.PropertyType(),
			Destination:         astbuilder.Selector(dst.NewIdent(builder.receiverIdent), string(toProp.PropertyName())),
			DestinationType:     toProp.PropertyType(),
			NameHint:            string(toProp.PropertyName()),
			ConversionContext:   nil,
			AssignmentHandler:   nil,
			Locals:              builder.locals,
			SourceProperty:      fromProp,
			DestinationProperty: toProp,
		})
	if err != nil {
		return notHandled,
			eris.Wrapf(
				err,
				"failed to generate conversion for property %s",
				toProp.PropertyName())
	}

	return handleWith(conversion), nil
}

//////////////////////////////////////////////////////////////////////////////////
// Complex property conversion (for when properties aren't simple primitive types)
//////////////////////////////////////////////////////////////////////////////////

// convertComplexTypeNameProperty handles conversion of complex TypeName properties.
// This function generates code that looks like this:
//
//	<nameHint>Converted := <destinationType>{}
//	err = <nameHint>Converted.FromARM(owner, <source>)
//	if err != nil {
//		return err
//	}
//	<destination> = <nameHint>
func (builder *convertFromARMBuilder) convertComplexTypeNameProperty(
	_ *astmodel.ConversionFunctionBuilder,
	params astmodel.ConversionParameters,
) ([]dst.Stmt, error) {
	destinationType, ok := params.DestinationType.(astmodel.InternalTypeName)
	if !ok {
		return nil, nil
	}

	sourceType, ok := params.SourceType.(astmodel.TypeName)
	if !ok {
		return nil, nil
	}

	// This is for handling type names that aren't equal
	if astmodel.TypeEquals(sourceType, destinationType) {
		return nil, nil
	}

	propertyLocalVar := builder.typeConversionBuilder.CreateLocal(params.Locals, "", params.NameHint)
	ownerName := builder.idFactory.CreateIdentifier(astmodel.OwnerProperty, astmodel.NotExported)

	newVariable := astbuilder.NewVariable(propertyLocalVar, destinationType.Name())
	if !destinationType.InternalPackageReference().Equals(builder.codeGenerationContext.CurrentPackage()) {
		// struct name has to be qualified
		packageName := builder.codeGenerationContext.MustGetImportedPackageName(destinationType.InternalPackageReference())
		newVariable = astbuilder.NewVariableQualified(
			propertyLocalVar,
			packageName,
			destinationType.Name())
	}

	tok := token.ASSIGN
	if !params.Locals.HasName("err") {
		tok = token.DEFINE
		params.Locals.Add("err")
	}

	var results []dst.Stmt
	results = append(results, newVariable)
	results = append(
		results,
		astbuilder.AssignmentStatement(
			dst.NewIdent("err"),
			tok,
			astbuilder.CallQualifiedFunc(
				propertyLocalVar, builder.methodName, dst.NewIdent(ownerName), params.GetSource())))
	results = append(results, astbuilder.CheckErrorAndReturn())
	if params.AssignmentHandler == nil {
		results = append(
			results,
			astbuilder.SimpleAssignment(
				params.GetDestination(),
				dst.NewIdent(propertyLocalVar)))
	} else {
		results = append(
			results,
			params.AssignmentHandler(params.GetDestination(), dst.NewIdent(propertyLocalVar)))
	}

	return results, nil
}
