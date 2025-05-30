// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601/storage"
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

func Test_DomainsTopic_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopic to hub returns original",
		prop.ForAll(RunResourceConversionTestForDomainsTopic, DomainsTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDomainsTopic tests if a specific instance of DomainsTopic round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDomainsTopic(subject DomainsTopic) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.DomainsTopic
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual DomainsTopic
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

func Test_DomainsTopic_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopic to DomainsTopic via AssignProperties_To_DomainsTopic & AssignProperties_From_DomainsTopic returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomainsTopic, DomainsTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomainsTopic tests if a specific instance of DomainsTopic can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDomainsTopic(subject DomainsTopic) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DomainsTopic
	err := copied.AssignProperties_To_DomainsTopic(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DomainsTopic
	err = actual.AssignProperties_From_DomainsTopic(&other)
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

func Test_DomainsTopic_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopic via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic, DomainsTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic runs a test to see if a specific instance of DomainsTopic round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic(subject DomainsTopic) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic
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

// Generator of DomainsTopic instances for property testing - lazily instantiated by DomainsTopicGenerator()
var domainsTopicGenerator gopter.Gen

// DomainsTopicGenerator returns a generator of DomainsTopic instances for property testing.
func DomainsTopicGenerator() gopter.Gen {
	if domainsTopicGenerator != nil {
		return domainsTopicGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDomainsTopic(generators)
	domainsTopicGenerator = gen.Struct(reflect.TypeOf(DomainsTopic{}), generators)

	return domainsTopicGenerator
}

// AddRelatedPropertyGeneratorsForDomainsTopic is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsTopic(gens map[string]gopter.Gen) {
	gens["Spec"] = DomainsTopic_SpecGenerator()
	gens["Status"] = DomainsTopic_STATUSGenerator()
}

func Test_DomainsTopicOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopicOperatorSpec to DomainsTopicOperatorSpec via AssignProperties_To_DomainsTopicOperatorSpec & AssignProperties_From_DomainsTopicOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomainsTopicOperatorSpec, DomainsTopicOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomainsTopicOperatorSpec tests if a specific instance of DomainsTopicOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDomainsTopicOperatorSpec(subject DomainsTopicOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DomainsTopicOperatorSpec
	err := copied.AssignProperties_To_DomainsTopicOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DomainsTopicOperatorSpec
	err = actual.AssignProperties_From_DomainsTopicOperatorSpec(&other)
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

func Test_DomainsTopicOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopicOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopicOperatorSpec, DomainsTopicOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopicOperatorSpec runs a test to see if a specific instance of DomainsTopicOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopicOperatorSpec(subject DomainsTopicOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopicOperatorSpec
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

// Generator of DomainsTopicOperatorSpec instances for property testing - lazily instantiated by
// DomainsTopicOperatorSpecGenerator()
var domainsTopicOperatorSpecGenerator gopter.Gen

// DomainsTopicOperatorSpecGenerator returns a generator of DomainsTopicOperatorSpec instances for property testing.
func DomainsTopicOperatorSpecGenerator() gopter.Gen {
	if domainsTopicOperatorSpecGenerator != nil {
		return domainsTopicOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	domainsTopicOperatorSpecGenerator = gen.Struct(reflect.TypeOf(DomainsTopicOperatorSpec{}), generators)

	return domainsTopicOperatorSpecGenerator
}

func Test_DomainsTopic_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopic_STATUS to DomainsTopic_STATUS via AssignProperties_To_DomainsTopic_STATUS & AssignProperties_From_DomainsTopic_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomainsTopic_STATUS, DomainsTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomainsTopic_STATUS tests if a specific instance of DomainsTopic_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDomainsTopic_STATUS(subject DomainsTopic_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DomainsTopic_STATUS
	err := copied.AssignProperties_To_DomainsTopic_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DomainsTopic_STATUS
	err = actual.AssignProperties_From_DomainsTopic_STATUS(&other)
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

func Test_DomainsTopic_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic_STATUS, DomainsTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic_STATUS runs a test to see if a specific instance of DomainsTopic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic_STATUS(subject DomainsTopic_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic_STATUS
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

// Generator of DomainsTopic_STATUS instances for property testing - lazily instantiated by
// DomainsTopic_STATUSGenerator()
var domainsTopic_STATUSGenerator gopter.Gen

// DomainsTopic_STATUSGenerator returns a generator of DomainsTopic_STATUS instances for property testing.
// We first initialize domainsTopic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainsTopic_STATUSGenerator() gopter.Gen {
	if domainsTopic_STATUSGenerator != nil {
		return domainsTopic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_STATUS(generators)
	domainsTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_STATUS(generators)
	AddRelatedPropertyGeneratorsForDomainsTopic_STATUS(generators)
	domainsTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_STATUS{}), generators)

	return domainsTopic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDomainsTopic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsTopic_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DomainTopicProperties_ProvisioningState_STATUS_Canceled,
		DomainTopicProperties_ProvisioningState_STATUS_Creating,
		DomainTopicProperties_ProvisioningState_STATUS_Deleting,
		DomainTopicProperties_ProvisioningState_STATUS_Failed,
		DomainTopicProperties_ProvisioningState_STATUS_Succeeded,
		DomainTopicProperties_ProvisioningState_STATUS_Updating))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDomainsTopic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsTopic_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_DomainsTopic_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DomainsTopic_Spec to DomainsTopic_Spec via AssignProperties_To_DomainsTopic_Spec & AssignProperties_From_DomainsTopic_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDomainsTopic_Spec, DomainsTopic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDomainsTopic_Spec tests if a specific instance of DomainsTopic_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDomainsTopic_Spec(subject DomainsTopic_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DomainsTopic_Spec
	err := copied.AssignProperties_To_DomainsTopic_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DomainsTopic_Spec
	err = actual.AssignProperties_From_DomainsTopic_Spec(&other)
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

func Test_DomainsTopic_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DomainsTopic_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDomainsTopic_Spec, DomainsTopic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDomainsTopic_Spec runs a test to see if a specific instance of DomainsTopic_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDomainsTopic_Spec(subject DomainsTopic_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DomainsTopic_Spec
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

// Generator of DomainsTopic_Spec instances for property testing - lazily instantiated by DomainsTopic_SpecGenerator()
var domainsTopic_SpecGenerator gopter.Gen

// DomainsTopic_SpecGenerator returns a generator of DomainsTopic_Spec instances for property testing.
// We first initialize domainsTopic_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DomainsTopic_SpecGenerator() gopter.Gen {
	if domainsTopic_SpecGenerator != nil {
		return domainsTopic_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_Spec(generators)
	domainsTopic_SpecGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDomainsTopic_Spec(generators)
	AddRelatedPropertyGeneratorsForDomainsTopic_Spec(generators)
	domainsTopic_SpecGenerator = gen.Struct(reflect.TypeOf(DomainsTopic_Spec{}), generators)

	return domainsTopic_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDomainsTopic_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDomainsTopic_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForDomainsTopic_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDomainsTopic_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(DomainsTopicOperatorSpecGenerator())
}
