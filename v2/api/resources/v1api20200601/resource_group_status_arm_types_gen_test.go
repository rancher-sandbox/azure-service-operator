// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

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

func Test_ResourceGroup_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceGroup_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceGroup_STATUS_ARM, ResourceGroup_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceGroup_STATUS_ARM runs a test to see if a specific instance of ResourceGroup_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceGroup_STATUS_ARM(subject ResourceGroup_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceGroup_STATUS_ARM
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

// Generator of ResourceGroup_STATUS_ARM instances for property testing - lazily instantiated by
// ResourceGroup_STATUS_ARMGenerator()
var resourceGroup_STATUS_ARMGenerator gopter.Gen

// ResourceGroup_STATUS_ARMGenerator returns a generator of ResourceGroup_STATUS_ARM instances for property testing.
// We first initialize resourceGroup_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ResourceGroup_STATUS_ARMGenerator() gopter.Gen {
	if resourceGroup_STATUS_ARMGenerator != nil {
		return resourceGroup_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceGroup_STATUS_ARM(generators)
	resourceGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ResourceGroup_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceGroup_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForResourceGroup_STATUS_ARM(generators)
	resourceGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ResourceGroup_STATUS_ARM{}), generators)

	return resourceGroup_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForResourceGroup_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ManagedBy"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForResourceGroup_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForResourceGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ResourceGroupProperties_STATUS_ARMGenerator())
}

func Test_ResourceGroupProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ResourceGroupProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForResourceGroupProperties_STATUS_ARM, ResourceGroupProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForResourceGroupProperties_STATUS_ARM runs a test to see if a specific instance of ResourceGroupProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForResourceGroupProperties_STATUS_ARM(subject ResourceGroupProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ResourceGroupProperties_STATUS_ARM
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

// Generator of ResourceGroupProperties_STATUS_ARM instances for property testing - lazily instantiated by
// ResourceGroupProperties_STATUS_ARMGenerator()
var resourceGroupProperties_STATUS_ARMGenerator gopter.Gen

// ResourceGroupProperties_STATUS_ARMGenerator returns a generator of ResourceGroupProperties_STATUS_ARM instances for property testing.
func ResourceGroupProperties_STATUS_ARMGenerator() gopter.Gen {
	if resourceGroupProperties_STATUS_ARMGenerator != nil {
		return resourceGroupProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForResourceGroupProperties_STATUS_ARM(generators)
	resourceGroupProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ResourceGroupProperties_STATUS_ARM{}), generators)

	return resourceGroupProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForResourceGroupProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForResourceGroupProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
}
