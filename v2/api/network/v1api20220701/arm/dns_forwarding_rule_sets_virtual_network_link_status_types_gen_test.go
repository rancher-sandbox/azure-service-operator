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

func Test_DnsForwardingRuleSetsVirtualNetworkLink_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsForwardingRuleSetsVirtualNetworkLink_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsForwardingRuleSetsVirtualNetworkLink_STATUS, DnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsForwardingRuleSetsVirtualNetworkLink_STATUS runs a test to see if a specific instance of DnsForwardingRuleSetsVirtualNetworkLink_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsForwardingRuleSetsVirtualNetworkLink_STATUS(subject DnsForwardingRuleSetsVirtualNetworkLink_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsForwardingRuleSetsVirtualNetworkLink_STATUS
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

// Generator of DnsForwardingRuleSetsVirtualNetworkLink_STATUS instances for property testing - lazily instantiated by
// DnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator()
var dnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator gopter.Gen

// DnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator returns a generator of DnsForwardingRuleSetsVirtualNetworkLink_STATUS instances for property testing.
// We first initialize dnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator() gopter.Gen {
	if dnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator != nil {
		return dnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_STATUS(generators)
	dnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleSetsVirtualNetworkLink_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_STATUS(generators)
	dnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsForwardingRuleSetsVirtualNetworkLink_STATUS{}), generators)

	return dnsForwardingRuleSetsVirtualNetworkLink_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsForwardingRuleSetsVirtualNetworkLink_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(VirtualNetworkLinkProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
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
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DnsresolverProvisioningState_STATUS_Canceled,
		DnsresolverProvisioningState_STATUS_Creating,
		DnsresolverProvisioningState_STATUS_Deleting,
		DnsresolverProvisioningState_STATUS_Failed,
		DnsresolverProvisioningState_STATUS_Succeeded,
		DnsresolverProvisioningState_STATUS_Updating))
}

// AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkLinkProperties_STATUS(gens map[string]gopter.Gen) {
	gens["VirtualNetwork"] = gen.PtrOf(SubResource_STATUSGenerator())
}
