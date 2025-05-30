// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

import (
	"encoding/json"
	v20220801s "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	v20230501ps "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ProductPolicy_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ProductPolicy to hub returns original",
		prop.ForAll(RunResourceConversionTestForProductPolicy, ProductPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForProductPolicy tests if a specific instance of ProductPolicy round trips to the hub storage version and back losslessly
func RunResourceConversionTestForProductPolicy(subject ProductPolicy) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220801s.ProductPolicy
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual ProductPolicy
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ProductPolicy_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ProductPolicy to ProductPolicy via AssignProperties_To_ProductPolicy & AssignProperties_From_ProductPolicy returns original",
		prop.ForAll(RunPropertyAssignmentTestForProductPolicy, ProductPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProductPolicy tests if a specific instance of ProductPolicy can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProductPolicy(subject ProductPolicy) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.ProductPolicy
	err := copied.AssignProperties_To_ProductPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ProductPolicy
	err = actual.AssignProperties_From_ProductPolicy(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ProductPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductPolicy, ProductPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductPolicy runs a test to see if a specific instance of ProductPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForProductPolicy(subject ProductPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductPolicy
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ProductPolicy instances for property testing - lazily instantiated by ProductPolicyGenerator()
var productPolicyGenerator gopter.Gen

// ProductPolicyGenerator returns a generator of ProductPolicy instances for property testing.
func ProductPolicyGenerator() gopter.Gen {
	if productPolicyGenerator != nil {
		return productPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForProductPolicy(generators)
	productPolicyGenerator = gen.Struct(reflect.TypeOf(ProductPolicy{}), generators)

	return productPolicyGenerator
}

// AddRelatedPropertyGeneratorsForProductPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProductPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = ProductPolicy_SpecGenerator()
	gens["Status"] = ProductPolicy_STATUSGenerator()
}

func Test_ProductPolicyOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ProductPolicyOperatorSpec to ProductPolicyOperatorSpec via AssignProperties_To_ProductPolicyOperatorSpec & AssignProperties_From_ProductPolicyOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForProductPolicyOperatorSpec, ProductPolicyOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProductPolicyOperatorSpec tests if a specific instance of ProductPolicyOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProductPolicyOperatorSpec(subject ProductPolicyOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.ProductPolicyOperatorSpec
	err := copied.AssignProperties_To_ProductPolicyOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ProductPolicyOperatorSpec
	err = actual.AssignProperties_From_ProductPolicyOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ProductPolicyOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductPolicyOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductPolicyOperatorSpec, ProductPolicyOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductPolicyOperatorSpec runs a test to see if a specific instance of ProductPolicyOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForProductPolicyOperatorSpec(subject ProductPolicyOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductPolicyOperatorSpec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ProductPolicyOperatorSpec instances for property testing - lazily instantiated by
// ProductPolicyOperatorSpecGenerator()
var productPolicyOperatorSpecGenerator gopter.Gen

// ProductPolicyOperatorSpecGenerator returns a generator of ProductPolicyOperatorSpec instances for property testing.
func ProductPolicyOperatorSpecGenerator() gopter.Gen {
	if productPolicyOperatorSpecGenerator != nil {
		return productPolicyOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	productPolicyOperatorSpecGenerator = gen.Struct(reflect.TypeOf(ProductPolicyOperatorSpec{}), generators)

	return productPolicyOperatorSpecGenerator
}

func Test_ProductPolicy_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ProductPolicy_STATUS to ProductPolicy_STATUS via AssignProperties_To_ProductPolicy_STATUS & AssignProperties_From_ProductPolicy_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForProductPolicy_STATUS, ProductPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProductPolicy_STATUS tests if a specific instance of ProductPolicy_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProductPolicy_STATUS(subject ProductPolicy_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.ProductPolicy_STATUS
	err := copied.AssignProperties_To_ProductPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ProductPolicy_STATUS
	err = actual.AssignProperties_From_ProductPolicy_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ProductPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductPolicy_STATUS, ProductPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductPolicy_STATUS runs a test to see if a specific instance of ProductPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProductPolicy_STATUS(subject ProductPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductPolicy_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ProductPolicy_STATUS instances for property testing - lazily instantiated by
// ProductPolicy_STATUSGenerator()
var productPolicy_STATUSGenerator gopter.Gen

// ProductPolicy_STATUSGenerator returns a generator of ProductPolicy_STATUS instances for property testing.
func ProductPolicy_STATUSGenerator() gopter.Gen {
	if productPolicy_STATUSGenerator != nil {
		return productPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductPolicy_STATUS(generators)
	productPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(ProductPolicy_STATUS{}), generators)

	return productPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForProductPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProductPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Format"] = gen.PtrOf(gen.OneConstOf(
		PolicyContractProperties_Format_STATUS_Rawxml,
		PolicyContractProperties_Format_STATUS_RawxmlLink,
		PolicyContractProperties_Format_STATUS_Xml,
		PolicyContractProperties_Format_STATUS_XmlLink))
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_ProductPolicy_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ProductPolicy_Spec to ProductPolicy_Spec via AssignProperties_To_ProductPolicy_Spec & AssignProperties_From_ProductPolicy_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForProductPolicy_Spec, ProductPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProductPolicy_Spec tests if a specific instance of ProductPolicy_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProductPolicy_Spec(subject ProductPolicy_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.ProductPolicy_Spec
	err := copied.AssignProperties_To_ProductPolicy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ProductPolicy_Spec
	err = actual.AssignProperties_From_ProductPolicy_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ProductPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductPolicy_Spec, ProductPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductPolicy_Spec runs a test to see if a specific instance of ProductPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForProductPolicy_Spec(subject ProductPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductPolicy_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ProductPolicy_Spec instances for property testing - lazily instantiated by ProductPolicy_SpecGenerator()
var productPolicy_SpecGenerator gopter.Gen

// ProductPolicy_SpecGenerator returns a generator of ProductPolicy_Spec instances for property testing.
// We first initialize productPolicy_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ProductPolicy_SpecGenerator() gopter.Gen {
	if productPolicy_SpecGenerator != nil {
		return productPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductPolicy_Spec(generators)
	productPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ProductPolicy_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductPolicy_Spec(generators)
	AddRelatedPropertyGeneratorsForProductPolicy_Spec(generators)
	productPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ProductPolicy_Spec{}), generators)

	return productPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForProductPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProductPolicy_Spec(gens map[string]gopter.Gen) {
	gens["Format"] = gen.PtrOf(gen.OneConstOf(
		PolicyContractProperties_Format_Rawxml,
		PolicyContractProperties_Format_RawxmlLink,
		PolicyContractProperties_Format_Xml,
		PolicyContractProperties_Format_XmlLink))
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProductPolicy_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProductPolicy_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(ProductPolicyOperatorSpecGenerator())
}
