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

func Test_PrivateDnsZone_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZone_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZone_STATUS, PrivateDnsZone_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZone_STATUS runs a test to see if a specific instance of PrivateDnsZone_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZone_STATUS(subject PrivateDnsZone_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZone_STATUS
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

// Generator of PrivateDnsZone_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZone_STATUSGenerator()
var privateDnsZone_STATUSGenerator gopter.Gen

// PrivateDnsZone_STATUSGenerator returns a generator of PrivateDnsZone_STATUS instances for property testing.
// We first initialize privateDnsZone_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZone_STATUSGenerator() gopter.Gen {
	if privateDnsZone_STATUSGenerator != nil {
		return privateDnsZone_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZone_STATUS(generators)
	privateDnsZone_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZone_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZone_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZone_STATUS(generators)
	privateDnsZone_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZone_STATUS{}), generators)

	return privateDnsZone_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZone_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZone_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZone_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZone_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateZoneProperties_STATUSGenerator())
}

func Test_PrivateZoneProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateZoneProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateZoneProperties_STATUS, PrivateZoneProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateZoneProperties_STATUS runs a test to see if a specific instance of PrivateZoneProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateZoneProperties_STATUS(subject PrivateZoneProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateZoneProperties_STATUS
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

// Generator of PrivateZoneProperties_STATUS instances for property testing - lazily instantiated by
// PrivateZoneProperties_STATUSGenerator()
var privateZoneProperties_STATUSGenerator gopter.Gen

// PrivateZoneProperties_STATUSGenerator returns a generator of PrivateZoneProperties_STATUS instances for property testing.
func PrivateZoneProperties_STATUSGenerator() gopter.Gen {
	if privateZoneProperties_STATUSGenerator != nil {
		return privateZoneProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateZoneProperties_STATUS(generators)
	privateZoneProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateZoneProperties_STATUS{}), generators)

	return privateZoneProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateZoneProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateZoneProperties_STATUS(gens map[string]gopter.Gen) {
	gens["MaxNumberOfRecordSets"] = gen.PtrOf(gen.Int())
	gens["MaxNumberOfVirtualNetworkLinks"] = gen.PtrOf(gen.Int())
	gens["MaxNumberOfVirtualNetworkLinksWithRegistration"] = gen.PtrOf(gen.Int())
	gens["NumberOfRecordSets"] = gen.PtrOf(gen.Int())
	gens["NumberOfVirtualNetworkLinks"] = gen.PtrOf(gen.Int())
	gens["NumberOfVirtualNetworkLinksWithRegistration"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		PrivateZoneProperties_ProvisioningState_STATUS_Canceled,
		PrivateZoneProperties_ProvisioningState_STATUS_Creating,
		PrivateZoneProperties_ProvisioningState_STATUS_Deleting,
		PrivateZoneProperties_ProvisioningState_STATUS_Failed,
		PrivateZoneProperties_ProvisioningState_STATUS_Succeeded,
		PrivateZoneProperties_ProvisioningState_STATUS_Updating))
}
