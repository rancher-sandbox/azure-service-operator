// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_RuleSet_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RuleSet via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRuleSet, RuleSetGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRuleSet runs a test to see if a specific instance of RuleSet round trips to JSON and back losslessly
func RunJSONSerializationTestForRuleSet(subject RuleSet) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RuleSet
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

// Generator of RuleSet instances for property testing - lazily instantiated by RuleSetGenerator()
var ruleSetGenerator gopter.Gen

// RuleSetGenerator returns a generator of RuleSet instances for property testing.
func RuleSetGenerator() gopter.Gen {
	if ruleSetGenerator != nil {
		return ruleSetGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRuleSet(generators)
	ruleSetGenerator = gen.Struct(reflect.TypeOf(RuleSet{}), generators)

	return ruleSetGenerator
}

// AddRelatedPropertyGeneratorsForRuleSet is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRuleSet(gens map[string]gopter.Gen) {
	gens["Spec"] = RuleSet_SpecGenerator()
	gens["Status"] = RuleSet_STATUSGenerator()
}

func Test_RuleSetOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RuleSetOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRuleSetOperatorSpec, RuleSetOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRuleSetOperatorSpec runs a test to see if a specific instance of RuleSetOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForRuleSetOperatorSpec(subject RuleSetOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RuleSetOperatorSpec
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

// Generator of RuleSetOperatorSpec instances for property testing - lazily instantiated by
// RuleSetOperatorSpecGenerator()
var ruleSetOperatorSpecGenerator gopter.Gen

// RuleSetOperatorSpecGenerator returns a generator of RuleSetOperatorSpec instances for property testing.
func RuleSetOperatorSpecGenerator() gopter.Gen {
	if ruleSetOperatorSpecGenerator != nil {
		return ruleSetOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	ruleSetOperatorSpecGenerator = gen.Struct(reflect.TypeOf(RuleSetOperatorSpec{}), generators)

	return ruleSetOperatorSpecGenerator
}

func Test_RuleSet_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RuleSet_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRuleSet_STATUS, RuleSet_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRuleSet_STATUS runs a test to see if a specific instance of RuleSet_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRuleSet_STATUS(subject RuleSet_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RuleSet_STATUS
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

// Generator of RuleSet_STATUS instances for property testing - lazily instantiated by RuleSet_STATUSGenerator()
var ruleSet_STATUSGenerator gopter.Gen

// RuleSet_STATUSGenerator returns a generator of RuleSet_STATUS instances for property testing.
// We first initialize ruleSet_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RuleSet_STATUSGenerator() gopter.Gen {
	if ruleSet_STATUSGenerator != nil {
		return ruleSet_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleSet_STATUS(generators)
	ruleSet_STATUSGenerator = gen.Struct(reflect.TypeOf(RuleSet_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleSet_STATUS(generators)
	AddRelatedPropertyGeneratorsForRuleSet_STATUS(generators)
	ruleSet_STATUSGenerator = gen.Struct(reflect.TypeOf(RuleSet_STATUS{}), generators)

	return ruleSet_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRuleSet_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRuleSet_STATUS(gens map[string]gopter.Gen) {
	gens["DeploymentStatus"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProfileName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRuleSet_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRuleSet_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_RuleSet_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RuleSet_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRuleSet_Spec, RuleSet_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRuleSet_Spec runs a test to see if a specific instance of RuleSet_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRuleSet_Spec(subject RuleSet_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RuleSet_Spec
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

// Generator of RuleSet_Spec instances for property testing - lazily instantiated by RuleSet_SpecGenerator()
var ruleSet_SpecGenerator gopter.Gen

// RuleSet_SpecGenerator returns a generator of RuleSet_Spec instances for property testing.
// We first initialize ruleSet_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RuleSet_SpecGenerator() gopter.Gen {
	if ruleSet_SpecGenerator != nil {
		return ruleSet_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleSet_Spec(generators)
	ruleSet_SpecGenerator = gen.Struct(reflect.TypeOf(RuleSet_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleSet_Spec(generators)
	AddRelatedPropertyGeneratorsForRuleSet_Spec(generators)
	ruleSet_SpecGenerator = gen.Struct(reflect.TypeOf(RuleSet_Spec{}), generators)

	return ruleSet_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRuleSet_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRuleSet_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRuleSet_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRuleSet_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(RuleSetOperatorSpecGenerator())
}
