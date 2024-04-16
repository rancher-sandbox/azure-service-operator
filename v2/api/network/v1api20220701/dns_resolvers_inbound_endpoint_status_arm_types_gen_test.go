// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

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

func Test_DnsResolvers_InboundEndpoint_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolvers_InboundEndpoint_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolvers_InboundEndpoint_STATUS_ARM, DnsResolvers_InboundEndpoint_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolvers_InboundEndpoint_STATUS_ARM runs a test to see if a specific instance of DnsResolvers_InboundEndpoint_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolvers_InboundEndpoint_STATUS_ARM(subject DnsResolvers_InboundEndpoint_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolvers_InboundEndpoint_STATUS_ARM
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

// Generator of DnsResolvers_InboundEndpoint_STATUS_ARM instances for property testing - lazily instantiated by
// DnsResolvers_InboundEndpoint_STATUS_ARMGenerator()
var dnsResolvers_InboundEndpoint_STATUS_ARMGenerator gopter.Gen

// DnsResolvers_InboundEndpoint_STATUS_ARMGenerator returns a generator of DnsResolvers_InboundEndpoint_STATUS_ARM instances for property testing.
// We first initialize dnsResolvers_InboundEndpoint_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolvers_InboundEndpoint_STATUS_ARMGenerator() gopter.Gen {
	if dnsResolvers_InboundEndpoint_STATUS_ARMGenerator != nil {
		return dnsResolvers_InboundEndpoint_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolvers_InboundEndpoint_STATUS_ARM(generators)
	dnsResolvers_InboundEndpoint_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsResolvers_InboundEndpoint_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolvers_InboundEndpoint_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsResolvers_InboundEndpoint_STATUS_ARM(generators)
	dnsResolvers_InboundEndpoint_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsResolvers_InboundEndpoint_STATUS_ARM{}), generators)

	return dnsResolvers_InboundEndpoint_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolvers_InboundEndpoint_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolvers_InboundEndpoint_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolvers_InboundEndpoint_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolvers_InboundEndpoint_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(InboundEndpointProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_InboundEndpointProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundEndpointProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundEndpointProperties_STATUS_ARM, InboundEndpointProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundEndpointProperties_STATUS_ARM runs a test to see if a specific instance of InboundEndpointProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundEndpointProperties_STATUS_ARM(subject InboundEndpointProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundEndpointProperties_STATUS_ARM
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

// Generator of InboundEndpointProperties_STATUS_ARM instances for property testing - lazily instantiated by
// InboundEndpointProperties_STATUS_ARMGenerator()
var inboundEndpointProperties_STATUS_ARMGenerator gopter.Gen

// InboundEndpointProperties_STATUS_ARMGenerator returns a generator of InboundEndpointProperties_STATUS_ARM instances for property testing.
// We first initialize inboundEndpointProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func InboundEndpointProperties_STATUS_ARMGenerator() gopter.Gen {
	if inboundEndpointProperties_STATUS_ARMGenerator != nil {
		return inboundEndpointProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundEndpointProperties_STATUS_ARM(generators)
	inboundEndpointProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(InboundEndpointProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundEndpointProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForInboundEndpointProperties_STATUS_ARM(generators)
	inboundEndpointProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(InboundEndpointProperties_STATUS_ARM{}), generators)

	return inboundEndpointProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForInboundEndpointProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundEndpointProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DnsresolverProvisioningState_STATUS_Canceled,
		DnsresolverProvisioningState_STATUS_Creating,
		DnsresolverProvisioningState_STATUS_Deleting,
		DnsresolverProvisioningState_STATUS_Failed,
		DnsresolverProvisioningState_STATUS_Succeeded,
		DnsresolverProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForInboundEndpointProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForInboundEndpointProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["IpConfigurations"] = gen.SliceOf(IpConfiguration_STATUS_ARMGenerator())
}

func Test_IpConfiguration_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpConfiguration_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpConfiguration_STATUS_ARM, IpConfiguration_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpConfiguration_STATUS_ARM runs a test to see if a specific instance of IpConfiguration_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForIpConfiguration_STATUS_ARM(subject IpConfiguration_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpConfiguration_STATUS_ARM
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

// Generator of IpConfiguration_STATUS_ARM instances for property testing - lazily instantiated by
// IpConfiguration_STATUS_ARMGenerator()
var ipConfiguration_STATUS_ARMGenerator gopter.Gen

// IpConfiguration_STATUS_ARMGenerator returns a generator of IpConfiguration_STATUS_ARM instances for property testing.
// We first initialize ipConfiguration_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IpConfiguration_STATUS_ARMGenerator() gopter.Gen {
	if ipConfiguration_STATUS_ARMGenerator != nil {
		return ipConfiguration_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpConfiguration_STATUS_ARM(generators)
	ipConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(IpConfiguration_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpConfiguration_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForIpConfiguration_STATUS_ARM(generators)
	ipConfiguration_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(IpConfiguration_STATUS_ARM{}), generators)

	return ipConfiguration_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForIpConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrivateIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIpAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IpConfiguration_PrivateIpAllocationMethod_STATUS_Dynamic, IpConfiguration_PrivateIpAllocationMethod_STATUS_Static))
}

// AddRelatedPropertyGeneratorsForIpConfiguration_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIpConfiguration_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(DnsresolverSubResource_STATUS_ARMGenerator())
}
