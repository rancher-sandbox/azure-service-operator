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

func Test_DnsZone_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZone_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZone_Spec, DnsZone_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZone_Spec runs a test to see if a specific instance of DnsZone_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZone_Spec(subject DnsZone_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZone_Spec
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

// Generator of DnsZone_Spec instances for property testing - lazily instantiated by DnsZone_SpecGenerator()
var dnsZone_SpecGenerator gopter.Gen

// DnsZone_SpecGenerator returns a generator of DnsZone_Spec instances for property testing.
// We first initialize dnsZone_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZone_SpecGenerator() gopter.Gen {
	if dnsZone_SpecGenerator != nil {
		return dnsZone_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZone_Spec(generators)
	dnsZone_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZone_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZone_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZone_Spec(generators)
	dnsZone_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZone_Spec{}), generators)

	return dnsZone_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZone_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZone_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsZone_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZone_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ZonePropertiesGenerator())
}

func Test_SubResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource, SubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource runs a test to see if a specific instance of SubResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource(subject SubResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource
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

// Generator of SubResource instances for property testing - lazily instantiated by SubResourceGenerator()
var subResourceGenerator gopter.Gen

// SubResourceGenerator returns a generator of SubResource instances for property testing.
func SubResourceGenerator() gopter.Gen {
	if subResourceGenerator != nil {
		return subResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource(generators)
	subResourceGenerator = gen.Struct(reflect.TypeOf(SubResource{}), generators)

	return subResourceGenerator
}

// AddIndependentPropertyGeneratorsForSubResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_ZoneProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ZoneProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForZoneProperties, ZonePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForZoneProperties runs a test to see if a specific instance of ZoneProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForZoneProperties(subject ZoneProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ZoneProperties
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

// Generator of ZoneProperties instances for property testing - lazily instantiated by ZonePropertiesGenerator()
var zonePropertiesGenerator gopter.Gen

// ZonePropertiesGenerator returns a generator of ZoneProperties instances for property testing.
// We first initialize zonePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ZonePropertiesGenerator() gopter.Gen {
	if zonePropertiesGenerator != nil {
		return zonePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForZoneProperties(generators)
	zonePropertiesGenerator = gen.Struct(reflect.TypeOf(ZoneProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForZoneProperties(generators)
	AddRelatedPropertyGeneratorsForZoneProperties(generators)
	zonePropertiesGenerator = gen.Struct(reflect.TypeOf(ZoneProperties{}), generators)

	return zonePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForZoneProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForZoneProperties(gens map[string]gopter.Gen) {
	gens["ZoneType"] = gen.PtrOf(gen.OneConstOf(ZoneProperties_ZoneType_Private, ZoneProperties_ZoneType_Public))
}

// AddRelatedPropertyGeneratorsForZoneProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForZoneProperties(gens map[string]gopter.Gen) {
	gens["RegistrationVirtualNetworks"] = gen.SliceOf(SubResourceGenerator())
	gens["ResolutionVirtualNetworks"] = gen.SliceOf(SubResourceGenerator())
}
