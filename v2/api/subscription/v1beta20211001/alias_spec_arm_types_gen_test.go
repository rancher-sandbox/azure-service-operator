// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

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

func Test_Alias_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Alias_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAlias_Spec_ARM, Alias_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAlias_Spec_ARM runs a test to see if a specific instance of Alias_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAlias_Spec_ARM(subject Alias_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Alias_Spec_ARM
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

// Generator of Alias_Spec_ARM instances for property testing - lazily instantiated by Alias_Spec_ARMGenerator()
var alias_Spec_ARMGenerator gopter.Gen

// Alias_Spec_ARMGenerator returns a generator of Alias_Spec_ARM instances for property testing.
// We first initialize alias_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Alias_Spec_ARMGenerator() gopter.Gen {
	if alias_Spec_ARMGenerator != nil {
		return alias_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAlias_Spec_ARM(generators)
	alias_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Alias_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAlias_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForAlias_Spec_ARM(generators)
	alias_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Alias_Spec_ARM{}), generators)

	return alias_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAlias_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAlias_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForAlias_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAlias_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PutAliasRequestProperties_ARMGenerator())
}

func Test_PutAliasRequestProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PutAliasRequestProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPutAliasRequestProperties_ARM, PutAliasRequestProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPutAliasRequestProperties_ARM runs a test to see if a specific instance of PutAliasRequestProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPutAliasRequestProperties_ARM(subject PutAliasRequestProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PutAliasRequestProperties_ARM
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

// Generator of PutAliasRequestProperties_ARM instances for property testing - lazily instantiated by
// PutAliasRequestProperties_ARMGenerator()
var putAliasRequestProperties_ARMGenerator gopter.Gen

// PutAliasRequestProperties_ARMGenerator returns a generator of PutAliasRequestProperties_ARM instances for property testing.
// We first initialize putAliasRequestProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PutAliasRequestProperties_ARMGenerator() gopter.Gen {
	if putAliasRequestProperties_ARMGenerator != nil {
		return putAliasRequestProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestProperties_ARM(generators)
	putAliasRequestProperties_ARMGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForPutAliasRequestProperties_ARM(generators)
	putAliasRequestProperties_ARMGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestProperties_ARM{}), generators)

	return putAliasRequestProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPutAliasRequestProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPutAliasRequestProperties_ARM(gens map[string]gopter.Gen) {
	gens["BillingScope"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["ResellerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Workload"] = gen.PtrOf(gen.OneConstOf(PutAliasRequestProperties_Workload_DevTest, PutAliasRequestProperties_Workload_Production))
}

// AddRelatedPropertyGeneratorsForPutAliasRequestProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPutAliasRequestProperties_ARM(gens map[string]gopter.Gen) {
	gens["AdditionalProperties"] = gen.PtrOf(PutAliasRequestAdditionalProperties_ARMGenerator())
}

func Test_PutAliasRequestAdditionalProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PutAliasRequestAdditionalProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPutAliasRequestAdditionalProperties_ARM, PutAliasRequestAdditionalProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPutAliasRequestAdditionalProperties_ARM runs a test to see if a specific instance of PutAliasRequestAdditionalProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPutAliasRequestAdditionalProperties_ARM(subject PutAliasRequestAdditionalProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PutAliasRequestAdditionalProperties_ARM
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

// Generator of PutAliasRequestAdditionalProperties_ARM instances for property testing - lazily instantiated by
// PutAliasRequestAdditionalProperties_ARMGenerator()
var putAliasRequestAdditionalProperties_ARMGenerator gopter.Gen

// PutAliasRequestAdditionalProperties_ARMGenerator returns a generator of PutAliasRequestAdditionalProperties_ARM instances for property testing.
func PutAliasRequestAdditionalProperties_ARMGenerator() gopter.Gen {
	if putAliasRequestAdditionalProperties_ARMGenerator != nil {
		return putAliasRequestAdditionalProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties_ARM(generators)
	putAliasRequestAdditionalProperties_ARMGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestAdditionalProperties_ARM{}), generators)

	return putAliasRequestAdditionalProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties_ARM(gens map[string]gopter.Gen) {
	gens["ManagementGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionOwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionTenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}