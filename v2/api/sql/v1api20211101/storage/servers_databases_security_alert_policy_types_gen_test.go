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

func Test_ServersDatabasesSecurityAlertPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesSecurityAlertPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy, ServersDatabasesSecurityAlertPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy runs a test to see if a specific instance of ServersDatabasesSecurityAlertPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy(subject ServersDatabasesSecurityAlertPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesSecurityAlertPolicy
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

// Generator of ServersDatabasesSecurityAlertPolicy instances for property testing - lazily instantiated by
// ServersDatabasesSecurityAlertPolicyGenerator()
var serversDatabasesSecurityAlertPolicyGenerator gopter.Gen

// ServersDatabasesSecurityAlertPolicyGenerator returns a generator of ServersDatabasesSecurityAlertPolicy instances for property testing.
func ServersDatabasesSecurityAlertPolicyGenerator() gopter.Gen {
	if serversDatabasesSecurityAlertPolicyGenerator != nil {
		return serversDatabasesSecurityAlertPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy(generators)
	serversDatabasesSecurityAlertPolicyGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicy{}), generators)

	return serversDatabasesSecurityAlertPolicyGenerator
}

// AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy(gens map[string]gopter.Gen) {
	gens["Spec"] = ServersDatabasesSecurityAlertPolicy_SpecGenerator()
	gens["Status"] = ServersDatabasesSecurityAlertPolicy_STATUSGenerator()
}

func Test_ServersDatabasesSecurityAlertPolicyOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesSecurityAlertPolicyOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesSecurityAlertPolicyOperatorSpec, ServersDatabasesSecurityAlertPolicyOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesSecurityAlertPolicyOperatorSpec runs a test to see if a specific instance of ServersDatabasesSecurityAlertPolicyOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesSecurityAlertPolicyOperatorSpec(subject ServersDatabasesSecurityAlertPolicyOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesSecurityAlertPolicyOperatorSpec
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

// Generator of ServersDatabasesSecurityAlertPolicyOperatorSpec instances for property testing - lazily instantiated by
// ServersDatabasesSecurityAlertPolicyOperatorSpecGenerator()
var serversDatabasesSecurityAlertPolicyOperatorSpecGenerator gopter.Gen

// ServersDatabasesSecurityAlertPolicyOperatorSpecGenerator returns a generator of ServersDatabasesSecurityAlertPolicyOperatorSpec instances for property testing.
func ServersDatabasesSecurityAlertPolicyOperatorSpecGenerator() gopter.Gen {
	if serversDatabasesSecurityAlertPolicyOperatorSpecGenerator != nil {
		return serversDatabasesSecurityAlertPolicyOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	serversDatabasesSecurityAlertPolicyOperatorSpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicyOperatorSpec{}), generators)

	return serversDatabasesSecurityAlertPolicyOperatorSpecGenerator
}

func Test_ServersDatabasesSecurityAlertPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesSecurityAlertPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_STATUS, ServersDatabasesSecurityAlertPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_STATUS runs a test to see if a specific instance of ServersDatabasesSecurityAlertPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_STATUS(subject ServersDatabasesSecurityAlertPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesSecurityAlertPolicy_STATUS
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

// Generator of ServersDatabasesSecurityAlertPolicy_STATUS instances for property testing - lazily instantiated by
// ServersDatabasesSecurityAlertPolicy_STATUSGenerator()
var serversDatabasesSecurityAlertPolicy_STATUSGenerator gopter.Gen

// ServersDatabasesSecurityAlertPolicy_STATUSGenerator returns a generator of ServersDatabasesSecurityAlertPolicy_STATUS instances for property testing.
// We first initialize serversDatabasesSecurityAlertPolicy_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersDatabasesSecurityAlertPolicy_STATUSGenerator() gopter.Gen {
	if serversDatabasesSecurityAlertPolicy_STATUSGenerator != nil {
		return serversDatabasesSecurityAlertPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS(generators)
	serversDatabasesSecurityAlertPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicy_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS(generators)
	AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS(generators)
	serversDatabasesSecurityAlertPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicy_STATUS{}), generators)

	return serversDatabasesSecurityAlertPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["CreationTime"] = gen.PtrOf(gen.AlphaString())
	gens["DisabledAlerts"] = gen.SliceOf(gen.AlphaString())
	gens["EmailAccountAdmins"] = gen.PtrOf(gen.Bool())
	gens["EmailAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_ServersDatabasesSecurityAlertPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServersDatabasesSecurityAlertPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_Spec, ServersDatabasesSecurityAlertPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_Spec runs a test to see if a specific instance of ServersDatabasesSecurityAlertPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForServersDatabasesSecurityAlertPolicy_Spec(subject ServersDatabasesSecurityAlertPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServersDatabasesSecurityAlertPolicy_Spec
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

// Generator of ServersDatabasesSecurityAlertPolicy_Spec instances for property testing - lazily instantiated by
// ServersDatabasesSecurityAlertPolicy_SpecGenerator()
var serversDatabasesSecurityAlertPolicy_SpecGenerator gopter.Gen

// ServersDatabasesSecurityAlertPolicy_SpecGenerator returns a generator of ServersDatabasesSecurityAlertPolicy_Spec instances for property testing.
// We first initialize serversDatabasesSecurityAlertPolicy_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ServersDatabasesSecurityAlertPolicy_SpecGenerator() gopter.Gen {
	if serversDatabasesSecurityAlertPolicy_SpecGenerator != nil {
		return serversDatabasesSecurityAlertPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec(generators)
	serversDatabasesSecurityAlertPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicy_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec(generators)
	AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec(generators)
	serversDatabasesSecurityAlertPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ServersDatabasesSecurityAlertPolicy_Spec{}), generators)

	return serversDatabasesSecurityAlertPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec(gens map[string]gopter.Gen) {
	gens["DisabledAlerts"] = gen.SliceOf(gen.AlphaString())
	gens["EmailAccountAdmins"] = gen.PtrOf(gen.Bool())
	gens["EmailAddresses"] = gen.SliceOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["RetentionDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.AlphaString())
	gens["StorageEndpoint"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServersDatabasesSecurityAlertPolicy_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(ServersDatabasesSecurityAlertPolicyOperatorSpecGenerator())
}
