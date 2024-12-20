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

func Test_PolicyContractProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PolicyContractProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicyContractProperties_STATUS, PolicyContractProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicyContractProperties_STATUS runs a test to see if a specific instance of PolicyContractProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicyContractProperties_STATUS(subject PolicyContractProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PolicyContractProperties_STATUS
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

// Generator of PolicyContractProperties_STATUS instances for property testing - lazily instantiated by
// PolicyContractProperties_STATUSGenerator()
var policyContractProperties_STATUSGenerator gopter.Gen

// PolicyContractProperties_STATUSGenerator returns a generator of PolicyContractProperties_STATUS instances for property testing.
func PolicyContractProperties_STATUSGenerator() gopter.Gen {
	if policyContractProperties_STATUSGenerator != nil {
		return policyContractProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPolicyContractProperties_STATUS(generators)
	policyContractProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(PolicyContractProperties_STATUS{}), generators)

	return policyContractProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPolicyContractProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPolicyContractProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Format"] = gen.PtrOf(gen.OneConstOf(
		PolicyContractProperties_Format_STATUS_Rawxml,
		PolicyContractProperties_Format_STATUS_RawxmlLink,
		PolicyContractProperties_Format_STATUS_Xml,
		PolicyContractProperties_Format_STATUS_XmlLink))
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_Policy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Policy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPolicy_STATUS, Policy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPolicy_STATUS runs a test to see if a specific instance of Policy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPolicy_STATUS(subject Policy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Policy_STATUS
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

// Generator of Policy_STATUS instances for property testing - lazily instantiated by Policy_STATUSGenerator()
var policy_STATUSGenerator gopter.Gen

// Policy_STATUSGenerator returns a generator of Policy_STATUS instances for property testing.
// We first initialize policy_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Policy_STATUSGenerator() gopter.Gen {
	if policy_STATUSGenerator != nil {
		return policy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPolicy_STATUS(generators)
	policy_STATUSGenerator = gen.Struct(reflect.TypeOf(Policy_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPolicy_STATUS(generators)
	AddRelatedPropertyGeneratorsForPolicy_STATUS(generators)
	policy_STATUSGenerator = gen.Struct(reflect.TypeOf(Policy_STATUS{}), generators)

	return policy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPolicy_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PolicyContractProperties_STATUSGenerator())
}
