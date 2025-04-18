/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

// ConversionParameters are parameters for converting between a source type and a destination type.
type ConversionParameters struct {
	Source            dst.Expr
	Destination       dst.Expr
	SourceType        Type
	DestinationType   Type
	NameHint          string
	ConversionContext []Type
	AssignmentHandler func(destination, source dst.Expr) dst.Stmt
	Locals            *KnownLocalsSet

	// SourceProperty is the source property for this conversion. Note that for recursive conversions this still refers
	// to the source property that the conversion chain will ultimately be sourced from
	SourceProperty *PropertyDefinition
	// DestinationProperty is the destination property for this conversion. Note that for recursive conversions this still refers
	// to the destination property that the conversion chain will ultimately be assigned to
	DestinationProperty *PropertyDefinition
}

// GetSource gets the Source field.
func (params ConversionParameters) GetSource() dst.Expr {
	return dst.Clone(params.Source).(dst.Expr)
}

// GetDestination gets the Destination field.
func (params ConversionParameters) GetDestination() dst.Expr {
	return dst.Clone(params.Destination).(dst.Expr)
}

// WithSource returns a new ConversionParameters with the updated Source.
func (params ConversionParameters) WithSource(source dst.Expr) ConversionParameters {
	result := params.copy()
	result.Source = source

	return result
}

// WithSourceType returns a new ConversionParameters with the updated SourceType.
func (params ConversionParameters) WithSourceType(t Type) ConversionParameters {
	result := params.copy()
	result.SourceType = t

	return result
}

// WithDestination returns a new ConversionParameters with the updated Destination.
func (params ConversionParameters) WithDestination(destination dst.Expr) ConversionParameters {
	result := params.copy()
	result.Destination = destination

	return result
}

// WithDestinationType returns a new ConversionParameters with the updated DestinationType.
func (params ConversionParameters) WithDestinationType(t Type) ConversionParameters {
	result := params.copy()
	result.DestinationType = t

	return result
}

// WithAssignmentHandler returns a new ConversionParameters with the updated AssignmentHandler.
func (params ConversionParameters) WithAssignmentHandler(
	assignmentHandler func(result dst.Expr, destination dst.Expr) dst.Stmt,
) ConversionParameters {
	result := params.copy()
	result.AssignmentHandler = assignmentHandler

	return result
}

// AssignmentHandlerOrDefault returns the AssignmentHandler or a default assignment handler if AssignmentHandler was nil.
func (params ConversionParameters) AssignmentHandlerOrDefault() func(destination, source dst.Expr) dst.Stmt {
	if params.AssignmentHandler == nil {
		return AssignmentHandlerAssign
	}
	return params.AssignmentHandler
}

// CountArraysAndMapsInConversionContext returns the number of arrays/maps which are in the conversion context.
// This is to aid in situations where there are deeply nested conversions (i.e. array of map of maps). In these contexts,
// temporary variables need to be declared to store intermediate conversion results.
func (params ConversionParameters) CountArraysAndMapsInConversionContext() int {
	result := 0
	for _, t := range params.ConversionContext {
		switch t.(type) {
		case *MapType:
			result += 1
		case *ArrayType:
			result += 1
		}
	}

	return result
}

func (params ConversionParameters) copy() ConversionParameters {
	result := params
	result.ConversionContext = append([]Type(nil), params.ConversionContext...)
	result.Locals = result.Locals.Clone()

	return result
}

type ConversionHandler func(builder *ConversionFunctionBuilder, params ConversionParameters) ([]dst.Stmt, error)

// TODO: There feels like overlap between this and the Storage Conversion Factories? Need further thinking to combine them?

// ConversionFunctionBuilder is used to build a function converting between two similar types.
// It has a set of built-in conversions and can be configured with additional conversions.
type ConversionFunctionBuilder struct {
	// TODO: Better way to let you fuss with this? How can you pick out what I've already put in here to overwrite it?
	conversions []ConversionHandler

	IDFactory             IdentifierFactory
	CodeGenerationContext *CodeGenerationContext

	// SupportExplicitEmptyCollectionsSerializationMode signals that collections can be initialized to a size 0 collection, rather than an empty collection
	SupportExplicitEmptyCollectionsSerializationMode bool
}

// NewConversionFunctionBuilder creates a new ConversionFunctionBuilder with the default conversions already added.
func NewConversionFunctionBuilder(
	idFactory IdentifierFactory,
	codeGenerationContext *CodeGenerationContext,
) *ConversionFunctionBuilder {
	return &ConversionFunctionBuilder{
		IDFactory:             idFactory,
		CodeGenerationContext: codeGenerationContext,
		conversions: []ConversionHandler{
			// Complex wrapper types checked first
			IdentityConvertComplexOptionalProperty,
			IdentityConvertComplexArrayProperty,
			IdentityConvertComplexMapProperty,

			// TODO: a flip function of some kind would be kinda nice (for source vs dest)
			IdentityAssignValidatedTypeDestination,
			IdentityAssignValidatedTypeSource,
			IdentityAssignPrimitiveType,
			AssignToOptional,
			AssignFromOptional,
			AssignToAliasOfPrimitive,
			AssignFromAliasOfPrimitive,
			AssignToAliasOfCollection,
			AssignFromAliasOfCollection,
			AssignToEnum,
			AssignFromEnum,
			IdentityDeepCopyJSON,
			IdentityAssignTypeName,
		},
	}
}

func (builder *ConversionFunctionBuilder) WithSupportExplicitEmptyCollectionsSerializationMode() *ConversionFunctionBuilder {
	builder.SupportExplicitEmptyCollectionsSerializationMode = true
	return builder
}

func (builder *ConversionFunctionBuilder) ShouldInitializeCollectionToEmpty(prop *PropertyDefinition) bool {
	return prop.HasTagValue(SerializationType, SerializationTypeExplicitEmptyCollection)
}

// AddConversionHandlers adds the specified conversion handlers to the end of the conversion list.
func (builder *ConversionFunctionBuilder) AddConversionHandlers(conversionHandlers ...ConversionHandler) {
	builder.conversions = append(builder.conversions, conversionHandlers...)
}

// PrependConversionHandlers adds the specified conversion handlers to the beginning of the conversion list.
func (builder *ConversionFunctionBuilder) PrependConversionHandlers(conversionHandlers ...ConversionHandler) {
	builder.conversions = append(conversionHandlers, builder.conversions...)
}

// BuildConversion creates a conversion between the source and destination defined by params.
func (builder *ConversionFunctionBuilder) BuildConversion(params ConversionParameters) ([]dst.Stmt, error) {
	for _, conversion := range builder.conversions {
		result, err := conversion(builder, params)
		if err != nil {
			return nil, err
		}

		if len(result) > 0 {
			return result, nil
		}
	}

	msg := fmt.Sprintf(
		"don't know how to perform conversion for %s -> %s",
		DebugDescription(params.SourceType),
		DebugDescription(params.DestinationType))
	panic(msg)
}

// IdentityConvertComplexOptionalProperty handles conversion for optional properties with complex elements
// This function generates code that looks like this:
//
//	if <source> != nil {
//		<code for producing result from destinationType.Element()>
//		<destination> = &<result>
//	}
func IdentityConvertComplexOptionalProperty(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	destinationType, ok := params.DestinationType.(*OptionalType)
	if !ok {
		return nil, nil
	}

	sourceType, ok := params.SourceType.(*OptionalType)
	if !ok {
		return nil, nil
	}

	locals := params.Locals.Clone()
	tempVarIdent := locals.CreateLocal(params.NameHint)

	conversion, err := builder.BuildConversion(
		ConversionParameters{
			Source:              astbuilder.Dereference(params.GetSource()),
			SourceType:          sourceType.Element(),
			Destination:         dst.NewIdent(tempVarIdent),
			DestinationType:     destinationType.Element(),
			NameHint:            params.NameHint,
			ConversionContext:   append(params.ConversionContext, destinationType),
			AssignmentHandler:   AssignmentHandlerDefine,
			Locals:              locals,
			SourceProperty:      params.SourceProperty,
			DestinationProperty: params.DestinationProperty,
		})
	if err != nil {
		return nil, eris.Wrap(err, "unable to build conversion for optional element")
	}

	// Tack on the final assignment
	conversion = append(
		conversion,
		astbuilder.SimpleAssignment(
			params.GetDestination(),
			astbuilder.AddrOf(dst.NewIdent(tempVarIdent))))

	result := astbuilder.IfNotNil(
		params.GetSource(),
		conversion...)

	return astbuilder.Statements(result), nil
}

// IdentityConvertComplexArrayProperty handles conversion for array properties with complex elements
// This function generates code that looks like this:
//
//	for _, item := range <source> {
//		<code for producing result from destinationType.Element()>
//		<destination> = append(<destination>, <result>)
//	}
func IdentityConvertComplexArrayProperty(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	destinationType, ok := params.DestinationType.(*ArrayType)
	if !ok {
		return nil, nil
	}

	sourceType, ok := params.SourceType.(*ArrayType)
	if !ok {
		return nil, nil
	}

	var results []dst.Stmt

	locals := params.Locals.Clone() // Loop variables are scoped inside the loop
	itemIdent := builder.CreateLocal(locals, "item", params.NameHint)

	depth := params.CountArraysAndMapsInConversionContext()
	destination := params.GetDestination()

	// Check what depth we're at to determine if we need to define an intermediate variable to hold the result
	// or if we'll be able to use the final destination directly.
	if depth > 0 {
		// TODO: The suffix here should maybe be configurable on the function builder?
		innerDestinationIdent := locals.CreateLocal(params.NameHint, "Temp")
		destination = dst.NewIdent(innerDestinationIdent)
		destinationTypeExpr, err := destinationType.AsTypeExpr(builder.CodeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating destination type expression")
		}

		results = append(
			results,
			astbuilder.LocalVariableDeclaration(
				innerDestinationIdent,
				destinationTypeExpr,
				""))
	}

	conversion, err := builder.BuildConversion(
		ConversionParameters{
			Source:              dst.NewIdent(itemIdent),
			SourceType:          sourceType.Element(),
			Destination:         dst.Clone(destination).(dst.Expr),
			DestinationType:     destinationType.Element(),
			NameHint:            itemIdent,
			ConversionContext:   append(params.ConversionContext, destinationType),
			AssignmentHandler:   astbuilder.AppendItemToSlice,
			Locals:              locals,
			SourceProperty:      params.SourceProperty,
			DestinationProperty: params.DestinationProperty,
		})
	if err != nil {
		return nil, eris.Wrap(err, "unable to build conversion for array element")
	}

	result := &dst.RangeStmt{
		Key:   dst.NewIdent("_"),
		Value: dst.NewIdent(itemIdent),
		X:     params.GetSource(),
		Tok:   token.DEFINE,
		Body:  astbuilder.StatementBlock(conversion...),
	}
	results = append(results, result)

	// If we have an assignment handler, we need to make sure to call it. This only happens in the case of nested
	// maps/arrays, where we need to make sure we generate the map assignment/array append before returning (otherwise
	// the "actual" assignment will just end up being to an empty array/map).
	if params.AssignmentHandler != nil {
		results = append(results, params.AssignmentHandler(params.GetDestination(), dst.Clone(destination).(dst.Expr)))
	}

	// If we must forcibly construct empty collections, check if the destination is nil and if so, construct an empty collection
	// This only applies for top-level collections (we don't forcibly construct nested collections)
	if depth == 0 && builder.ShouldInitializeCollectionToEmpty(params.SourceProperty) {
		destinationTypeExpr, err := destinationType.Element().AsTypeExpr(builder.CodeGenerationContext)
		if err != nil {
			return nil, eris.Wrap(err, "creating destination type expression")
		}

		emptySlice := astbuilder.SliceLiteral(destinationTypeExpr)
		assignEmpty := astbuilder.SimpleAssignment(params.GetDestination(), emptySlice)
		astbuilder.AddComment(
			&assignEmpty.Decs.Start,
			"// Set property to empty map, as this resource is set to serialize all collections explicitly")
		assignEmpty.Decs.Before = dst.NewLine

		ifNil := astbuilder.IfNil(
			params.GetDestination(),
			assignEmpty)
		results = append(results, ifNil)
	}

	return results, nil
}

// IdentityConvertComplexMapProperty handles conversion for map properties with complex values.
// This function panics if the map keys are not primitive types.
// This function generates code that looks like this:
//
//	if <source> != nil {
//		<destination> = make(map[<destinationType.KeyType()]<destinationType.ValueType()>, len(<source>))
//		for key, value := range <source> {
//			<code for producing result from destinationType.ValueType()>
//			<destination>[key] = <result>
//		}
//	}
func IdentityConvertComplexMapProperty(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	destinationType, ok := params.DestinationType.(*MapType)
	if !ok {
		return nil, nil
	}

	sourceType, ok := params.SourceType.(*MapType)
	if !ok {
		return nil, nil
	}

	if _, ok := destinationType.KeyType().(*PrimitiveType); !ok {
		msg := fmt.Sprintf(
			"map had non-primitive key type: %s",
			DebugDescription(destinationType.KeyType()))
		panic(msg)
	}

	depth := params.CountArraysAndMapsInConversionContext()

	locals := params.Locals.Clone() // Loop variables are scoped inside the loop

	keyIdent := builder.CreateLocal(locals, "key", params.NameHint)
	valueIdent := builder.CreateLocal(locals, "value", params.NameHint)

	nameHint := valueIdent
	destination := params.GetDestination()
	makeMapToken := token.ASSIGN

	// Check what depth we're at to determine if we need to define an intermediate variable to hold the result
	// or if we'll be able to use the final destination directly.
	if depth > 0 {
		innerDestinationIdent := locals.CreateLocal(params.NameHint, "Temp")
		destination = dst.NewIdent(innerDestinationIdent)
		makeMapToken = token.DEFINE
	}

	handler := func(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
		return astbuilder.InsertMap(lhs, dst.NewIdent(keyIdent), rhs)
	}

	keyTypeExpr, err := destinationType.KeyType().AsTypeExpr(builder.CodeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating key type expression")
	}

	valueTypeExpr, err := destinationType.ValueType().AsTypeExpr(builder.CodeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating value type expression")
	}

	makeMapStatement := astbuilder.AssignmentStatement(
		destination,
		makeMapToken,
		astbuilder.MakeMapWithCapacity(keyTypeExpr, valueTypeExpr,
			astbuilder.CallFunc("len", params.GetSource())))

	conversion, err := builder.BuildConversion(
		ConversionParameters{
			Source:              dst.NewIdent(valueIdent),
			SourceType:          sourceType.ValueType(),
			Destination:         dst.Clone(destination).(dst.Expr),
			DestinationType:     destinationType.ValueType(),
			NameHint:            nameHint,
			ConversionContext:   append(params.ConversionContext, destinationType),
			AssignmentHandler:   handler,
			Locals:              locals,
			SourceProperty:      params.SourceProperty,
			DestinationProperty: params.DestinationProperty,
		})
	if err != nil {
		return nil, eris.Wrap(err, "unable to build conversion for map value")
	}

	rangeStatement := &dst.RangeStmt{
		Key:   dst.NewIdent(keyIdent),
		Value: dst.NewIdent(valueIdent),
		X:     params.GetSource(),
		Tok:   token.DEFINE,
		Body:  astbuilder.StatementBlock(conversion...),
	}

	// If we must forcibly construct empty collections, check if the destination is nil and if so, construct an empty collection
	// This only applies for top-level collections (we don't forcibly construct nested collections)
	var result *dst.IfStmt
	if depth == 0 && builder.ShouldInitializeCollectionToEmpty(params.SourceProperty) {
		emptyMap := astbuilder.MakeMap(keyTypeExpr, valueTypeExpr)

		assignEmpty := astbuilder.SimpleAssignment(params.GetDestination(), emptyMap)
		astbuilder.AddComment(
			&assignEmpty.Decs.Start,
			"// Set property to empty map, as this resource is set to serialize all collections explicitly")
		assignEmpty.Decs.Before = dst.NewLine

		result = astbuilder.SimpleIfElse(
			astbuilder.NotNil(params.GetSource()),
			astbuilder.Statements(
				makeMapStatement,
				rangeStatement,
			),
			astbuilder.Statements(
				assignEmpty,
			),
		)
	} else {
		result = astbuilder.IfNotNil(
			params.GetSource(),
			makeMapStatement,
			rangeStatement)
	}

	// If we have an assignment handler, we need to make sure to call it. This only happens in the case of nested
	// maps/arrays, where we need to make sure we generate the map assignment/array append before returning (otherwise
	// the "actual" assignment will just end up being to an empty array/map).
	if params.AssignmentHandler != nil {
		result.Body.List = append(result.Body.List, params.AssignmentHandler(params.GetDestination(), dst.Clone(destination).(dst.Expr)))
	}

	return astbuilder.Statements(result), nil
}

// IdentityAssignTypeName handles conversion for TypeName's that are the same
// Note that because this handler is dealing with TypeName's and not Optional<TypeName>, it is safe to
// perform a simple assignment rather than a copy.
// This function generates code that looks like this:
//
//	<destination> <assignmentHandler> <source>
func IdentityAssignTypeName(
	_ *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	destinationType, ok := params.DestinationType.(TypeName)
	if !ok {
		return nil, nil
	}

	sourceType, ok := params.SourceType.(TypeName)
	if !ok {
		return nil, nil
	}

	// Can only apply basic assignment for typeNames that are the same
	if !TypeEquals(sourceType, destinationType) {
		return nil, nil
	}

	return astbuilder.Statements(
		params.AssignmentHandlerOrDefault()(
			params.GetDestination(),
			params.GetSource())), nil
}

// IdentityAssignPrimitiveType just assigns source to destination directly, no conversion needed.
// This function generates code that looks like this:
// <destination> <assignmentHandler> <source>
func IdentityAssignPrimitiveType(
	_ *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	if _, ok := params.DestinationType.(*PrimitiveType); !ok {
		return nil, nil
	}

	if _, ok := params.SourceType.(*PrimitiveType); !ok {
		return nil, nil
	}

	return astbuilder.Statements(
		params.AssignmentHandlerOrDefault()(
			params.GetDestination(),
			params.GetSource())), nil
}

// AssignToOptional assigns address of source to destination.
// This function generates code that looks like this, for simple conversions:
// <destination> <assignmentHandler> &<source>
//
// or:
// <destination>Temp := convert(<source>)
// <destination> <assignmentHandler> &<destination>Temp
func AssignToOptional(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	optDest, ok := params.DestinationType.(*OptionalType)
	if !ok {
		return nil, nil
	}

	if TypeEquals(optDest.Element(), params.SourceType) {
		return astbuilder.Statements(
			params.AssignmentHandlerOrDefault()(
				params.GetDestination(),
				astbuilder.AddrOf(params.GetSource()))), nil
	}

	// a more complex conversion is needed
	dstType := optDest.Element()
	tmpLocal := builder.CreateLocal(params.Locals, "temp", params.NameHint)

	conversion, err := builder.BuildConversion(
		ConversionParameters{
			Source:              params.Source,
			SourceType:          params.SourceType,
			Destination:         dst.NewIdent(tmpLocal),
			DestinationType:     dstType,
			NameHint:            tmpLocal,
			ConversionContext:   nil,
			AssignmentHandler:   nil,
			Locals:              params.Locals,
			SourceProperty:      params.SourceProperty,
			DestinationProperty: params.DestinationProperty,
		})
	if err != nil {
		return nil, eris.Wrap(err, "unable to build inner conversion to optional")
	}

	if len(conversion) == 0 {
		return nil, nil // unable to build inner conversion
	}

	destinationTypeExpr, err := dstType.AsTypeExpr(builder.CodeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating destination type expression")
	}

	return astbuilder.Statements(
		astbuilder.LocalVariableDeclaration(tmpLocal, destinationTypeExpr, ""),
		conversion,
		params.AssignmentHandlerOrDefault()(
			params.GetDestination(),
			astbuilder.AddrOf(dst.NewIdent(tmpLocal)))), nil
}

// AssignFromOptional assigns address of source to destination.
// This function generates code that looks like this, for simple conversions:
//
//	if (<source> != nil) {
//	    <destination> <assignmentHandler> *<source>
//	}
//
// or:
//
//	if (<source> != nil) {
//	    <destination>Temp := convert(*<source>)
//	    <destination> <assignmentHandler> <destination>Temp
//	}
func AssignFromOptional(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	optSrc, ok := params.SourceType.(*OptionalType)
	if !ok {
		return nil, nil
	}

	if TypeEquals(optSrc.Element(), params.DestinationType) {
		return astbuilder.Statements(
			astbuilder.IfNotNil(params.GetSource(),
				params.AssignmentHandlerOrDefault()(params.GetDestination(), astbuilder.Dereference(params.GetSource())),
			)), nil
	}

	// a more complex conversion is needed
	srcType := optSrc.Element()
	locals := params.Locals.Clone()
	tmpLocal := builder.CreateLocal(locals, "temp", params.NameHint)

	conversion, err := builder.BuildConversion(
		ConversionParameters{
			Source:              astbuilder.Dereference(params.GetSource()),
			SourceType:          srcType,
			Destination:         dst.NewIdent(tmpLocal),
			DestinationType:     params.DestinationType,
			NameHint:            tmpLocal,
			ConversionContext:   nil,
			AssignmentHandler:   nil,
			Locals:              locals,
			SourceProperty:      params.SourceProperty,
			DestinationProperty: params.DestinationProperty,
		})
	if err != nil {
		return nil, eris.Wrap(err, "unable to build inner conversion from optional")
	}

	if len(conversion) == 0 {
		return nil, nil // unable to build inner conversion
	}

	var result []dst.Stmt
	destinationTypeExpr, err := params.DestinationType.AsTypeExpr(builder.CodeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating destination type expression")
	}

	result = append(result, astbuilder.LocalVariableDeclaration(tmpLocal, destinationTypeExpr, ""))
	result = append(result, conversion...)
	result = append(result, params.AssignmentHandlerOrDefault()(params.GetDestination(), dst.NewIdent(tmpLocal)))

	return astbuilder.Statements(
		astbuilder.IfNotNil(
			params.GetSource(),
			result...)), nil
}

// AssignToAliasOfPrimitive assigns a primitive value to an alias of that same type.
// This function generates code that looks like this:
//
// <destination> <assignmentHandler> <destinationType>(<source>)
func AssignToAliasOfPrimitive(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	// Source must be a primitive type
	srcPrim, ok := params.SourceType.(*PrimitiveType)
	if !ok {
		return nil, nil
	}

	// Destination must be an internal type name
	dstName, ok := params.DestinationType.(InternalTypeName)
	if !ok {
		return nil, nil
	}

	// ... who's definition we know ...
	dstDef, err := builder.CodeGenerationContext.GetDefinition(dstName)
	if err != nil {
		//nolint:nilerr // err is not nil, we defer to a different conversion
		return nil, nil
	}

	// ... and it's not optional (hedging against oddities) ...
	if _, ok = AsOptionalType(dstDef.Type()); ok {
		return nil, nil
	}

	// ... and it is a primitive type ...
	dstPrim, ok := AsPrimitiveType(dstDef.Type())
	if !ok {
		return nil, nil
	}

	// ... and is an alias of the source type ...
	if !TypeEquals(srcPrim, dstPrim) {
		return nil, nil
	}

	return astbuilder.Statements(
			params.AssignmentHandlerOrDefault()(
				params.GetDestination(),
				astbuilder.CallFunc(
					dstName.Name(),
					params.GetSource()))),
		nil
}

// AssignFromAliasOfPrimitive assigns a primitive value from an alias of that same type.
// This function generates code that looks like this:
//
// <destination> <assignmentHandler> <destinationType>(<source>)
func AssignFromAliasOfPrimitive(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	// Source must be an internal type name
	srcName, ok := params.SourceType.(InternalTypeName)
	if !ok {
		return nil, nil
	}

	// ... who's definition we know ...
	srcDef, err := builder.CodeGenerationContext.GetDefinition(srcName)
	if err != nil {
		//nolint:nilerr // err is not nil, we defer to a different conversion
		return nil, nil
	}

	// ... and it's not optional (hedging against oddities) ...
	if _, ok = AsOptionalType(srcDef.Type()); ok {
		return nil, nil
	}

	// ... and it is a primitive type ...
	srcPrim, ok := AsPrimitiveType(srcDef.Type())
	if !ok {
		return nil, nil
	}

	// Destination must be a primitive type
	dstPrim, ok := params.DestinationType.(*PrimitiveType)
	if !ok {
		return nil, nil
	}

	// ... and is an alias of the source type ...
	if !TypeEquals(srcPrim, dstPrim) {
		return nil, nil
	}

	return astbuilder.Statements(
			params.AssignmentHandlerOrDefault()(
				params.GetDestination(),
				astbuilder.CallFunc(
					dstPrim.String(),
					params.GetSource()))),
		nil
}

// AssignToAliasOfCollection assigns an array of values to an alias of that same type.
// This function generates code that looks like this:
//
// <destination> <assignmentHandler> <destinationType>(<source>)
func AssignToAliasOfCollection(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	// Source must be an array type or a map type
	_, srcIsArray := AsArrayType(params.SourceType)
	_, srcIsMap := AsMapType(params.SourceType)
	if !srcIsArray && !srcIsMap {
		return nil, nil
	}

	// Destination must be an internal type name
	dstName, ok := params.DestinationType.(InternalTypeName)
	if !ok {
		return nil, nil
	}

	// ... who's definition we know ...
	dstDef, err := builder.CodeGenerationContext.GetDefinition(dstName)
	if err != nil {
		//nolint:nilerr // err is not nil, we defer to a different conversion
		return nil, nil
	}

	// ... and it's not optional (hedging against oddities) ...
	if _, ok = AsOptionalType(dstDef.Type()); ok {
		return nil, nil
	}

	// ... and it must also be an array type or a map type (and the same kind as source)
	dstArray, dstIsArray := AsArrayType(dstDef.Type())
	dstMap, dstIsMap := AsMapType(dstDef.Type())
	if dstIsArray != srcIsArray || dstIsMap != srcIsMap {
		return nil, nil
	}

	var dstType Type
	if dstIsArray {
		dstType = dstArray
	} else if dstIsMap {
		dstType = dstMap
	}

	// ... and we can convert between those array types ...
	conversion, err := builder.BuildConversion(
		ConversionParameters{
			Source:            params.Source,
			SourceType:        params.SourceType,
			Destination:       params.Destination,
			DestinationType:   dstDef.Type(),
			NameHint:          params.NameHint,
			ConversionContext: append(params.ConversionContext, dstType),
			AssignmentHandler: func(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
				// Use the existing assignment handler, but make sure we cast the rhs to the right type
				return params.AssignmentHandlerOrDefault()(
					lhs,
					astbuilder.CallFunc(
						dstName.Name(),
						rhs))
			},
			Locals:              params.Locals,
			SourceProperty:      params.SourceProperty,
			DestinationProperty: params.DestinationProperty,
		})
	if err != nil {
		return nil, eris.Wrap(err, "unable to build conversion for array element")
	}

	return astbuilder.Statements(
			conversion,
		),
		nil
}

// AssignFromAliasOfCollection assigns an array of values from an alias of that same type.
// This function generates code that looks like this:
//
// <destination> <assignmentHandler> <destinationType>(<source>)
func AssignFromAliasOfCollection(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	// Source must be an internal type name
	srcName, ok := params.SourceType.(InternalTypeName)
	if !ok {
		return nil, nil
	}

	// ... who's definition we know ...
	srcDef, err := builder.CodeGenerationContext.GetDefinition(srcName)
	if err != nil {
		//nolint:nilerr // err is not nil, we defer to a different conversion
		return nil, nil
	}

	// ... and it's not optional (hedging against oddities) ...
	if _, ok = AsOptionalType(srcDef.Type()); ok {
		return nil, nil
	}

	// ... and it is an array type or a map type ...
	_, srcIsArray := AsArrayType(srcDef.Type())
	_, srcIsMap := AsMapType(srcDef.Type())
	if !srcIsArray && !srcIsMap {
		return nil, nil
	}

	// Destination must also be an array type or a map type (and the same kind as source)
	_, dstIsArray := AsArrayType(params.DestinationType)
	_, dstIsMap := AsMapType(params.DestinationType)
	if dstIsArray != srcIsArray || dstIsMap != srcIsMap {
		return nil, nil
	}

	// ... and we can convert between those collection types ...
	conversion, err := builder.BuildConversion(
		ConversionParameters{
			Source:            params.Source,
			SourceType:        srcDef.Type(),
			Destination:       params.Destination,
			DestinationType:   params.DestinationType,
			NameHint:          params.NameHint,
			ConversionContext: append(params.ConversionContext, params.DestinationType),
			AssignmentHandler: func(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
				// Use the existing assignment handler if there is one, but make sure we always assign
				return params.AssignmentHandlerOrDefault()(
					lhs,
					rhs)
			},
			Locals:              params.Locals,
			SourceProperty:      params.SourceProperty,
			DestinationProperty: params.DestinationProperty,
		})
	if err != nil {
		return nil, eris.Wrap(err, "unable to build conversion for array element")
	}

	return astbuilder.Statements(
			conversion,
		),
		nil
}

// AssignToEnum stores a value into an enum type.
// This function generates code that looks like this:
// <destination> = <enumType>(<source>)
func AssignToEnum(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	// Is the destination a typename of an enum?
	itn, ok := params.DestinationType.(InternalTypeName)
	if !ok {
		// Not a typename
		return nil, nil
	}

	def, err := builder.CodeGenerationContext.GetDefinition(itn)
	if err != nil {
		// Couldn't the definition, this handler isn't the one we want
		// (not actually an error)
		return nil, nil //nolint:nilerr // err is not nil, we defer to a different conversion
	}

	et, ok := AsEnumType(def.Type())
	if !ok {
		// Definition isn't for an enum type,
		return nil, nil
	}

	if TypeEquals(et.BaseType(), params.SourceType) {
		// We can directly cast the source to the destination
		var cast dst.Expr
		if builder.CodeGenerationContext.CurrentPackage() == itn.PackageReference() {
			cast = astbuilder.CallFunc(itn.Name(), params.GetSource())
		} else {
			alias := builder.CodeGenerationContext.MustGetImportedPackageName(itn.PackageReference())
			cast = astbuilder.CallQualifiedFunc(alias, itn.Name(), params.GetSource())
		}

		return astbuilder.Statements(
			params.AssignmentHandlerOrDefault()(
				params.GetDestination(),
				cast)), nil
	}

	// a more complex conversion is needed
	dstType := et.BaseType()
	tmpLocal := builder.CreateLocal(params.Locals, "temp", params.NameHint)

	conversion, err := builder.BuildConversion(
		ConversionParameters{
			Source:              params.Source,
			SourceType:          params.SourceType,
			Destination:         dst.NewIdent(tmpLocal),
			DestinationType:     dstType,
			NameHint:            tmpLocal,
			ConversionContext:   nil,
			AssignmentHandler:   nil,
			Locals:              params.Locals,
			SourceProperty:      params.SourceProperty,
			DestinationProperty: params.DestinationProperty,
		})
	if err != nil {
		return nil, eris.Wrapf(err, "unable to build inner conversion to enum %s", itn.name)
	}

	if len(conversion) == 0 {
		// unable to build inner conversion
		return nil, nil
	}

	destinationTypeExpr, err := dstType.AsTypeExpr(builder.CodeGenerationContext)
	if err != nil {
		return nil, eris.Wrap(err, "creating destination type expression")
	}

	var cast dst.Expr
	if builder.CodeGenerationContext.CurrentPackage() == itn.PackageReference() {
		cast = astbuilder.CallFunc(itn.Name(), dst.NewIdent(tmpLocal))
	} else {
		alias := builder.CodeGenerationContext.MustGetImportedPackageName(itn.PackageReference())
		cast = astbuilder.CallQualifiedFunc(alias, itn.Name(), dst.NewIdent(tmpLocal))
	}

	return astbuilder.Statements(
		astbuilder.LocalVariableDeclaration(tmpLocal, destinationTypeExpr, ""),
		conversion,
		params.AssignmentHandlerOrDefault()(
			params.GetDestination(),
			cast)), nil
}

// AssignFromEnum reads a value from an enum type.
// This function generates code that looks like this:
// <destination> = <primitiveType>(<source>)
func AssignFromEnum(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	// Is the source a typename of an enum?
	itn, ok := params.SourceType.(InternalTypeName)
	if !ok {
		// Not a typename
		return nil, nil
	}

	def, err := builder.CodeGenerationContext.GetDefinition(itn)
	if err != nil {
		return nil, nil //nolint:nilerr // err is not nil, we defer to a different conversion
	}

	et, ok := AsEnumType(def.Type())
	if !ok {
		// Definition isn't for an enum type,
		return nil, nil
	}

	srcType := et.BaseType()
	cast := astbuilder.CallFunc(srcType.Name(), params.GetSource())

	conversion, err := builder.BuildConversion(
		ConversionParameters{
			Source:              cast,
			SourceType:          srcType,
			Destination:         params.Destination,
			DestinationType:     params.DestinationType,
			NameHint:            "",
			ConversionContext:   nil,
			AssignmentHandler:   nil,
			Locals:              params.Locals,
			SourceProperty:      params.SourceProperty,
			DestinationProperty: params.DestinationProperty,
		})
	if err != nil {
		return nil, eris.Wrapf(err, "unable to build inner conversion to enum %s", itn.name)
	}

	if len(conversion) == 0 {
		// unable to build inner conversion
		return nil, nil
	}

	return conversion, nil
}

// IdentityAssignValidatedTypeDestination generates an assignment to the underlying validated type Element
func IdentityAssignValidatedTypeDestination(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	validatedType, ok := params.DestinationType.(*ValidatedType)
	if !ok {
		return nil, nil
	}

	// pass through to underlying type
	params = params.WithDestinationType(validatedType.ElementType())
	return builder.BuildConversion(params)
}

// IdentityAssignValidatedTypeSource generates an assignment to the underlying validated type Element
func IdentityAssignValidatedTypeSource(
	builder *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	validatedType, ok := params.SourceType.(*ValidatedType)
	if !ok {
		return nil, nil
	}

	// pass through to underlying type
	params = params.WithSourceType(validatedType.ElementType())
	return builder.BuildConversion(params)
}

// IdentityDeepCopyJSON special cases copying JSON-type fields to call the DeepCopy method.
// It generates code that looks like:
//
//	<destination> = *<source>.DeepCopy()
func IdentityDeepCopyJSON(
	_ *ConversionFunctionBuilder,
	params ConversionParameters,
) ([]dst.Stmt, error) {
	if !TypeEquals(params.DestinationType, JSONType) {
		return nil, nil
	}

	newSource := astbuilder.Dereference(
		&dst.CallExpr{
			Fun:  astbuilder.Selector(params.GetSource(), "DeepCopy"),
			Args: []dst.Expr{},
		})

	return astbuilder.Statements(
		params.AssignmentHandlerOrDefault()(params.GetDestination(), newSource)), nil
}

// AssignmentHandlerDefine is an assignment handler for definitions, using :=
func AssignmentHandlerDefine(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
	return astbuilder.AssignmentStatement(lhs, token.DEFINE, rhs)
}

// AssignmentHandlerAssign is an assignment handler for standard assignments to existing variables, using =
func AssignmentHandlerAssign(lhs dst.Expr, rhs dst.Expr) dst.Stmt {
	return astbuilder.SimpleAssignment(lhs, rhs)
}

// CreateLocal creates an unused local variable name.
// Names are chosen according to the following rules:
//  1. If there is no local variable with the <suffix> name, use that.
//  2. If there is a local variable with the <suffix> name, create a variable name <nameHint><suffix>.
//
// In the case that <nameHint><suffix> is also taken append numbers to the end in standard KnownLocalsSet fashion.
// Note that this function trims numbers on the right hand side of nameHint, so a nameHint of "item1" will get a local
// variable named item<suffix>.
func (builder *ConversionFunctionBuilder) CreateLocal(locals *KnownLocalsSet, suffix string, nameHint string) string {
	ident := suffix
	if locals.HasName(ident) || ident == "" {
		// Trim any trailing numbers so that we don't end up with ident1111
		// Note that this can end up trimming trailing digits on fields that have them (i.e. "loop1" might become "loop").
		// This is ok as it doesn't hurt anything and helps avoid "loop12" (for loop1 number 2).
		trimmedNameHint := strings.TrimRight(nameHint, "0123456789")
		ident = locals.CreateLocal(trimmedNameHint, suffix)
	} else {
		locals.Add(ident)
	}

	return ident
}
