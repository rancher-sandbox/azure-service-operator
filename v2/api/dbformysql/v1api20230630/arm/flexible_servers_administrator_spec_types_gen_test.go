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

func Test_AdministratorProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AdministratorProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAdministratorProperties, AdministratorPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAdministratorProperties runs a test to see if a specific instance of AdministratorProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForAdministratorProperties(subject AdministratorProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AdministratorProperties
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

// Generator of AdministratorProperties instances for property testing - lazily instantiated by
// AdministratorPropertiesGenerator()
var administratorPropertiesGenerator gopter.Gen

// AdministratorPropertiesGenerator returns a generator of AdministratorProperties instances for property testing.
func AdministratorPropertiesGenerator() gopter.Gen {
	if administratorPropertiesGenerator != nil {
		return administratorPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAdministratorProperties(generators)
	administratorPropertiesGenerator = gen.Struct(reflect.TypeOf(AdministratorProperties{}), generators)

	return administratorPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForAdministratorProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAdministratorProperties(gens map[string]gopter.Gen) {
	gens["AdministratorType"] = gen.PtrOf(gen.OneConstOf(AdministratorProperties_AdministratorType_ActiveDirectory))
	gens["IdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Login"] = gen.PtrOf(gen.AlphaString())
	gens["Sid"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_FlexibleServersAdministrator_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersAdministrator_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersAdministrator_Spec, FlexibleServersAdministrator_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersAdministrator_Spec runs a test to see if a specific instance of FlexibleServersAdministrator_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersAdministrator_Spec(subject FlexibleServersAdministrator_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersAdministrator_Spec
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

// Generator of FlexibleServersAdministrator_Spec instances for property testing - lazily instantiated by
// FlexibleServersAdministrator_SpecGenerator()
var flexibleServersAdministrator_SpecGenerator gopter.Gen

// FlexibleServersAdministrator_SpecGenerator returns a generator of FlexibleServersAdministrator_Spec instances for property testing.
// We first initialize flexibleServersAdministrator_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersAdministrator_SpecGenerator() gopter.Gen {
	if flexibleServersAdministrator_SpecGenerator != nil {
		return flexibleServersAdministrator_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_Spec(generators)
	flexibleServersAdministrator_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministrator_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_Spec(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_Spec(generators)
	flexibleServersAdministrator_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersAdministrator_Spec{}), generators)

	return flexibleServersAdministrator_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersAdministrator_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersAdministrator_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AdministratorPropertiesGenerator())
}
