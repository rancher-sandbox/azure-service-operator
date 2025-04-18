// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
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

func Test_CustomDomainProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CustomDomainProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCustomDomainProperties, CustomDomainPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCustomDomainProperties runs a test to see if a specific instance of CustomDomainProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForCustomDomainProperties(subject CustomDomainProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CustomDomainProperties
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

// Generator of CustomDomainProperties instances for property testing - lazily instantiated by
// CustomDomainPropertiesGenerator()
var customDomainPropertiesGenerator gopter.Gen

// CustomDomainPropertiesGenerator returns a generator of CustomDomainProperties instances for property testing.
// We first initialize customDomainPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CustomDomainPropertiesGenerator() gopter.Gen {
	if customDomainPropertiesGenerator != nil {
		return customDomainPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomDomainProperties(generators)
	customDomainPropertiesGenerator = gen.Struct(reflect.TypeOf(CustomDomainProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomDomainProperties(generators)
	AddRelatedPropertyGeneratorsForCustomDomainProperties(generators)
	customDomainPropertiesGenerator = gen.Struct(reflect.TypeOf(CustomDomainProperties{}), generators)

	return customDomainPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForCustomDomainProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCustomDomainProperties(gens map[string]gopter.Gen) {
	gens["DomainName"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForCustomDomainProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCustomDomainProperties(gens map[string]gopter.Gen) {
	gens["CustomCertificate"] = gen.PtrOf(ResourceReferenceGenerator())
}

func Test_CustomDomain_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CustomDomain_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCustomDomain_Spec, CustomDomain_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCustomDomain_Spec runs a test to see if a specific instance of CustomDomain_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForCustomDomain_Spec(subject CustomDomain_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CustomDomain_Spec
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

// Generator of CustomDomain_Spec instances for property testing - lazily instantiated by CustomDomain_SpecGenerator()
var customDomain_SpecGenerator gopter.Gen

// CustomDomain_SpecGenerator returns a generator of CustomDomain_Spec instances for property testing.
// We first initialize customDomain_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CustomDomain_SpecGenerator() gopter.Gen {
	if customDomain_SpecGenerator != nil {
		return customDomain_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomDomain_Spec(generators)
	customDomain_SpecGenerator = gen.Struct(reflect.TypeOf(CustomDomain_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomDomain_Spec(generators)
	AddRelatedPropertyGeneratorsForCustomDomain_Spec(generators)
	customDomain_SpecGenerator = gen.Struct(reflect.TypeOf(CustomDomain_Spec{}), generators)

	return customDomain_SpecGenerator
}

// AddIndependentPropertyGeneratorsForCustomDomain_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCustomDomain_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForCustomDomain_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCustomDomain_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(CustomDomainPropertiesGenerator())
}

func Test_ResourceReference_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceReference via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceReference, ResourceReferenceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceReference runs a test to see if a specific instance of ResourceReference round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceReference(subject ResourceReference) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceReference
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

// Generator of ResourceReference instances for property testing - lazily instantiated by ResourceReferenceGenerator()
var resourceReferenceGenerator gopter.Gen

// ResourceReferenceGenerator returns a generator of ResourceReference instances for property testing.
func ResourceReferenceGenerator() gopter.Gen {
	if resourceReferenceGenerator != nil {
		return resourceReferenceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceReference(generators)
	resourceReferenceGenerator = gen.Struct(reflect.TypeOf(ResourceReference{}), generators)

	return resourceReferenceGenerator
}

// AddIndependentPropertyGeneratorsForResourceReference is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceReference(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
