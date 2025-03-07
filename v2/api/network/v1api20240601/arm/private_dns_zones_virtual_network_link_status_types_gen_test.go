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

func Test_PrivateDnsZonesVirtualNetworkLink_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesVirtualNetworkLink_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_STATUS, PrivateDnsZonesVirtualNetworkLink_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_STATUS runs a test to see if a specific instance of PrivateDnsZonesVirtualNetworkLink_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesVirtualNetworkLink_STATUS(subject PrivateDnsZonesVirtualNetworkLink_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesVirtualNetworkLink_STATUS
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

// Generator of PrivateDnsZonesVirtualNetworkLink_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZonesVirtualNetworkLink_STATUSGenerator()
var privateDnsZonesVirtualNetworkLink_STATUSGenerator gopter.Gen

// PrivateDnsZonesVirtualNetworkLink_STATUSGenerator returns a generator of PrivateDnsZonesVirtualNetworkLink_STATUS instances for property testing.
// We first initialize privateDnsZonesVirtualNetworkLink_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesVirtualNetworkLink_STATUSGenerator() gopter.Gen {
	if privateDnsZonesVirtualNetworkLink_STATUSGenerator != nil {
		return privateDnsZonesVirtualNetworkLink_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS(generators)
	privateDnsZonesVirtualNetworkLink_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLink_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS(generators)
	privateDnsZonesVirtualNetworkLink_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesVirtualNetworkLink_STATUS{}), generators)

	return privateDnsZonesVirtualNetworkLink_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesVirtualNetworkLink_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkLinkProperties_STATUSGenerator())
}

func Test_SubResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_STATUS, SubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_STATUS runs a test to see if a specific instance of SubResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_STATUS(subject SubResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_STATUS
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

// Generator of SubResource_STATUS instances for property testing - lazily instantiated by SubResource_STATUSGenerator()
var subResource_STATUSGenerator gopter.Gen

// SubResource_STATUSGenerator returns a generator of SubResource_STATUS instances for property testing.
func SubResource_STATUSGenerator() gopter.Gen {
	if subResource_STATUSGenerator != nil {
		return subResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_STATUS(generators)
	subResource_STATUSGenerator = gen.Struct(reflect.TypeOf(SubResource_STATUS{}), generators)

	return subResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_VirtualNetworkLinkProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkLinkProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkLinkProperties_STATUS, VirtualNetworkLinkProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkLinkProperties_STATUS runs a test to see if a specific instance of VirtualNetworkLinkProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkLinkProperties_STATUS(subject VirtualNetworkLinkProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkLinkProperties_STATUS
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

// Generator of VirtualNetworkLinkProperties_STATUS instances for property testing - lazily instantiated by
// VirtualNetworkLinkProperties_STATUSGenerator()
var virtualNetworkLinkProperties_STATUSGenerator gopter.Gen

// VirtualNetworkLinkProperties_STATUSGenerator returns a generator of VirtualNetworkLinkProperties_STATUS instances for property testing.
// We first initialize virtualNetworkLinkProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkLinkProperties_STATUSGenerator() gopter.Gen {
	if virtualNetworkLinkProperties_STATUSGenerator != nil {
		return virtualNetworkLinkProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties_STATUS(generators)
	virtualNetworkLinkProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkLinkProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties_STATUS(generators)
	virtualNetworkLinkProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkLinkProperties_STATUS{}), generators)

	return virtualNetworkLinkProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkLinkProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Canceled,
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Creating,
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Deleting,
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Failed,
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Succeeded,
		VirtualNetworkLinkProperties_ProvisioningState_STATUS_Updating))
	gens["RegistrationEnabled"] = gen.PtrOf(gen.Bool())
	gens["ResolutionPolicy"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkLinkProperties_ResolutionPolicy_STATUS_Default, VirtualNetworkLinkProperties_ResolutionPolicy_STATUS_NxDomainRedirect))
	gens["VirtualNetworkLinkState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_Completed, VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_InProgress))
}

// AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties_STATUS(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(SubResource_STATUSGenerator())
}
