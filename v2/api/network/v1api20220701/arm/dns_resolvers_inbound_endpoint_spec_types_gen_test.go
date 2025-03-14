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

func Test_DnsResolversInboundEndpoint_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolversInboundEndpoint_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolversInboundEndpoint_Spec, DnsResolversInboundEndpoint_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolversInboundEndpoint_Spec runs a test to see if a specific instance of DnsResolversInboundEndpoint_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolversInboundEndpoint_Spec(subject DnsResolversInboundEndpoint_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolversInboundEndpoint_Spec
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

// Generator of DnsResolversInboundEndpoint_Spec instances for property testing - lazily instantiated by
// DnsResolversInboundEndpoint_SpecGenerator()
var dnsResolversInboundEndpoint_SpecGenerator gopter.Gen

// DnsResolversInboundEndpoint_SpecGenerator returns a generator of DnsResolversInboundEndpoint_Spec instances for property testing.
// We first initialize dnsResolversInboundEndpoint_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolversInboundEndpoint_SpecGenerator() gopter.Gen {
	if dnsResolversInboundEndpoint_SpecGenerator != nil {
		return dnsResolversInboundEndpoint_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolversInboundEndpoint_Spec(generators)
	dnsResolversInboundEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(DnsResolversInboundEndpoint_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolversInboundEndpoint_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsResolversInboundEndpoint_Spec(generators)
	dnsResolversInboundEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(DnsResolversInboundEndpoint_Spec{}), generators)

	return dnsResolversInboundEndpoint_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolversInboundEndpoint_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolversInboundEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolversInboundEndpoint_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolversInboundEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(InboundEndpointPropertiesGenerator())
}

func Test_InboundEndpointProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundEndpointProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundEndpointProperties, InboundEndpointPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundEndpointProperties runs a test to see if a specific instance of InboundEndpointProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundEndpointProperties(subject InboundEndpointProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundEndpointProperties
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

// Generator of InboundEndpointProperties instances for property testing - lazily instantiated by
// InboundEndpointPropertiesGenerator()
var inboundEndpointPropertiesGenerator gopter.Gen

// InboundEndpointPropertiesGenerator returns a generator of InboundEndpointProperties instances for property testing.
func InboundEndpointPropertiesGenerator() gopter.Gen {
	if inboundEndpointPropertiesGenerator != nil {
		return inboundEndpointPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForInboundEndpointProperties(generators)
	inboundEndpointPropertiesGenerator = gen.Struct(reflect.TypeOf(InboundEndpointProperties{}), generators)

	return inboundEndpointPropertiesGenerator
}

// AddRelatedPropertyGeneratorsForInboundEndpointProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForInboundEndpointProperties(gens map[string]gopter.Gen) {
	gens["IpConfigurations"] = gen.SliceOf(IpConfigurationGenerator())
}

func Test_IpConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IpConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIpConfiguration, IpConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIpConfiguration runs a test to see if a specific instance of IpConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForIpConfiguration(subject IpConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IpConfiguration
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

// Generator of IpConfiguration instances for property testing - lazily instantiated by IpConfigurationGenerator()
var ipConfigurationGenerator gopter.Gen

// IpConfigurationGenerator returns a generator of IpConfiguration instances for property testing.
// We first initialize ipConfigurationGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IpConfigurationGenerator() gopter.Gen {
	if ipConfigurationGenerator != nil {
		return ipConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpConfiguration(generators)
	ipConfigurationGenerator = gen.Struct(reflect.TypeOf(IpConfiguration{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIpConfiguration(generators)
	AddRelatedPropertyGeneratorsForIpConfiguration(generators)
	ipConfigurationGenerator = gen.Struct(reflect.TypeOf(IpConfiguration{}), generators)

	return ipConfigurationGenerator
}

// AddIndependentPropertyGeneratorsForIpConfiguration is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIpConfiguration(gens map[string]gopter.Gen) {
	gens["PrivateIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateIpAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IpConfiguration_PrivateIpAllocationMethod_Dynamic, IpConfiguration_PrivateIpAllocationMethod_Static))
}

// AddRelatedPropertyGeneratorsForIpConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIpConfiguration(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(SubResourceGenerator())
}
