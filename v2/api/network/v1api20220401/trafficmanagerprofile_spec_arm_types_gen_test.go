// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220401

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

func Test_Trafficmanagerprofile_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Trafficmanagerprofile_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTrafficmanagerprofile_Spec_ARM, Trafficmanagerprofile_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTrafficmanagerprofile_Spec_ARM runs a test to see if a specific instance of Trafficmanagerprofile_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTrafficmanagerprofile_Spec_ARM(subject Trafficmanagerprofile_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Trafficmanagerprofile_Spec_ARM
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

// Generator of Trafficmanagerprofile_Spec_ARM instances for property testing - lazily instantiated by
// Trafficmanagerprofile_Spec_ARMGenerator()
var trafficmanagerprofile_Spec_ARMGenerator gopter.Gen

// Trafficmanagerprofile_Spec_ARMGenerator returns a generator of Trafficmanagerprofile_Spec_ARM instances for property testing.
// We first initialize trafficmanagerprofile_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Trafficmanagerprofile_Spec_ARMGenerator() gopter.Gen {
	if trafficmanagerprofile_Spec_ARMGenerator != nil {
		return trafficmanagerprofile_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficmanagerprofile_Spec_ARM(generators)
	trafficmanagerprofile_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Trafficmanagerprofile_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTrafficmanagerprofile_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForTrafficmanagerprofile_Spec_ARM(generators)
	trafficmanagerprofile_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Trafficmanagerprofile_Spec_ARM{}), generators)

	return trafficmanagerprofile_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForTrafficmanagerprofile_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTrafficmanagerprofile_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTrafficmanagerprofile_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTrafficmanagerprofile_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ProfileProperties_ARMGenerator())
}

func Test_ProfileProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProfileProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfileProperties_ARM, ProfileProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfileProperties_ARM runs a test to see if a specific instance of ProfileProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfileProperties_ARM(subject ProfileProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProfileProperties_ARM
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

// Generator of ProfileProperties_ARM instances for property testing - lazily instantiated by
// ProfileProperties_ARMGenerator()
var profileProperties_ARMGenerator gopter.Gen

// ProfileProperties_ARMGenerator returns a generator of ProfileProperties_ARM instances for property testing.
// We first initialize profileProperties_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ProfileProperties_ARMGenerator() gopter.Gen {
	if profileProperties_ARMGenerator != nil {
		return profileProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileProperties_ARM(generators)
	profileProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ProfileProperties_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfileProperties_ARM(generators)
	AddRelatedPropertyGeneratorsForProfileProperties_ARM(generators)
	profileProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ProfileProperties_ARM{}), generators)

	return profileProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfileProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfileProperties_ARM(gens map[string]gopter.Gen) {
	gens["AllowedEndpointRecordTypes"] = gen.SliceOf(gen.OneConstOf(
		AllowedEndpointRecordType_Any,
		AllowedEndpointRecordType_DomainName,
		AllowedEndpointRecordType_IPv4Address,
		AllowedEndpointRecordType_IPv6Address))
	gens["MaxReturn"] = gen.PtrOf(gen.Int())
	gens["ProfileStatus"] = gen.PtrOf(gen.OneConstOf(ProfileProperties_ProfileStatus_Disabled, ProfileProperties_ProfileStatus_Enabled))
	gens["TrafficRoutingMethod"] = gen.PtrOf(gen.OneConstOf(
		ProfileProperties_TrafficRoutingMethod_Geographic,
		ProfileProperties_TrafficRoutingMethod_MultiValue,
		ProfileProperties_TrafficRoutingMethod_Performance,
		ProfileProperties_TrafficRoutingMethod_Priority,
		ProfileProperties_TrafficRoutingMethod_Subnet,
		ProfileProperties_TrafficRoutingMethod_Weighted))
	gens["TrafficViewEnrollmentStatus"] = gen.PtrOf(gen.OneConstOf(ProfileProperties_TrafficViewEnrollmentStatus_Disabled, ProfileProperties_TrafficViewEnrollmentStatus_Enabled))
}

// AddRelatedPropertyGeneratorsForProfileProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfileProperties_ARM(gens map[string]gopter.Gen) {
	gens["DnsConfig"] = gen.PtrOf(DnsConfig_ARMGenerator())
	gens["MonitorConfig"] = gen.PtrOf(MonitorConfig_ARMGenerator())
}

func Test_DnsConfig_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsConfig_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsConfig_ARM, DnsConfig_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsConfig_ARM runs a test to see if a specific instance of DnsConfig_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsConfig_ARM(subject DnsConfig_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsConfig_ARM
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

// Generator of DnsConfig_ARM instances for property testing - lazily instantiated by DnsConfig_ARMGenerator()
var dnsConfig_ARMGenerator gopter.Gen

// DnsConfig_ARMGenerator returns a generator of DnsConfig_ARM instances for property testing.
func DnsConfig_ARMGenerator() gopter.Gen {
	if dnsConfig_ARMGenerator != nil {
		return dnsConfig_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsConfig_ARM(generators)
	dnsConfig_ARMGenerator = gen.Struct(reflect.TypeOf(DnsConfig_ARM{}), generators)

	return dnsConfig_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsConfig_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsConfig_ARM(gens map[string]gopter.Gen) {
	gens["RelativeName"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

func Test_MonitorConfig_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MonitorConfig_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMonitorConfig_ARM, MonitorConfig_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMonitorConfig_ARM runs a test to see if a specific instance of MonitorConfig_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMonitorConfig_ARM(subject MonitorConfig_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MonitorConfig_ARM
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

// Generator of MonitorConfig_ARM instances for property testing - lazily instantiated by MonitorConfig_ARMGenerator()
var monitorConfig_ARMGenerator gopter.Gen

// MonitorConfig_ARMGenerator returns a generator of MonitorConfig_ARM instances for property testing.
// We first initialize monitorConfig_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MonitorConfig_ARMGenerator() gopter.Gen {
	if monitorConfig_ARMGenerator != nil {
		return monitorConfig_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMonitorConfig_ARM(generators)
	monitorConfig_ARMGenerator = gen.Struct(reflect.TypeOf(MonitorConfig_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMonitorConfig_ARM(generators)
	AddRelatedPropertyGeneratorsForMonitorConfig_ARM(generators)
	monitorConfig_ARMGenerator = gen.Struct(reflect.TypeOf(MonitorConfig_ARM{}), generators)

	return monitorConfig_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMonitorConfig_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMonitorConfig_ARM(gens map[string]gopter.Gen) {
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProfileMonitorStatus"] = gen.PtrOf(gen.OneConstOf(
		MonitorConfig_ProfileMonitorStatus_CheckingEndpoints,
		MonitorConfig_ProfileMonitorStatus_Degraded,
		MonitorConfig_ProfileMonitorStatus_Disabled,
		MonitorConfig_ProfileMonitorStatus_Inactive,
		MonitorConfig_ProfileMonitorStatus_Online))
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(MonitorConfig_Protocol_HTTP, MonitorConfig_Protocol_HTTPS, MonitorConfig_Protocol_TCP))
	gens["TimeoutInSeconds"] = gen.PtrOf(gen.Int())
	gens["ToleratedNumberOfFailures"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForMonitorConfig_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMonitorConfig_ARM(gens map[string]gopter.Gen) {
	gens["CustomHeaders"] = gen.SliceOf(MonitorConfig_CustomHeaders_ARMGenerator())
	gens["ExpectedStatusCodeRanges"] = gen.SliceOf(MonitorConfig_ExpectedStatusCodeRanges_ARMGenerator())
}

func Test_MonitorConfig_CustomHeaders_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MonitorConfig_CustomHeaders_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMonitorConfig_CustomHeaders_ARM, MonitorConfig_CustomHeaders_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMonitorConfig_CustomHeaders_ARM runs a test to see if a specific instance of MonitorConfig_CustomHeaders_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMonitorConfig_CustomHeaders_ARM(subject MonitorConfig_CustomHeaders_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MonitorConfig_CustomHeaders_ARM
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

// Generator of MonitorConfig_CustomHeaders_ARM instances for property testing - lazily instantiated by
// MonitorConfig_CustomHeaders_ARMGenerator()
var monitorConfig_CustomHeaders_ARMGenerator gopter.Gen

// MonitorConfig_CustomHeaders_ARMGenerator returns a generator of MonitorConfig_CustomHeaders_ARM instances for property testing.
func MonitorConfig_CustomHeaders_ARMGenerator() gopter.Gen {
	if monitorConfig_CustomHeaders_ARMGenerator != nil {
		return monitorConfig_CustomHeaders_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMonitorConfig_CustomHeaders_ARM(generators)
	monitorConfig_CustomHeaders_ARMGenerator = gen.Struct(reflect.TypeOf(MonitorConfig_CustomHeaders_ARM{}), generators)

	return monitorConfig_CustomHeaders_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMonitorConfig_CustomHeaders_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMonitorConfig_CustomHeaders_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_MonitorConfig_ExpectedStatusCodeRanges_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MonitorConfig_ExpectedStatusCodeRanges_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMonitorConfig_ExpectedStatusCodeRanges_ARM, MonitorConfig_ExpectedStatusCodeRanges_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMonitorConfig_ExpectedStatusCodeRanges_ARM runs a test to see if a specific instance of MonitorConfig_ExpectedStatusCodeRanges_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMonitorConfig_ExpectedStatusCodeRanges_ARM(subject MonitorConfig_ExpectedStatusCodeRanges_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MonitorConfig_ExpectedStatusCodeRanges_ARM
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

// Generator of MonitorConfig_ExpectedStatusCodeRanges_ARM instances for property testing - lazily instantiated by
// MonitorConfig_ExpectedStatusCodeRanges_ARMGenerator()
var monitorConfig_ExpectedStatusCodeRanges_ARMGenerator gopter.Gen

// MonitorConfig_ExpectedStatusCodeRanges_ARMGenerator returns a generator of MonitorConfig_ExpectedStatusCodeRanges_ARM instances for property testing.
func MonitorConfig_ExpectedStatusCodeRanges_ARMGenerator() gopter.Gen {
	if monitorConfig_ExpectedStatusCodeRanges_ARMGenerator != nil {
		return monitorConfig_ExpectedStatusCodeRanges_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMonitorConfig_ExpectedStatusCodeRanges_ARM(generators)
	monitorConfig_ExpectedStatusCodeRanges_ARMGenerator = gen.Struct(reflect.TypeOf(MonitorConfig_ExpectedStatusCodeRanges_ARM{}), generators)

	return monitorConfig_ExpectedStatusCodeRanges_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMonitorConfig_ExpectedStatusCodeRanges_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMonitorConfig_ExpectedStatusCodeRanges_ARM(gens map[string]gopter.Gen) {
	gens["Max"] = gen.PtrOf(gen.Int())
	gens["Min"] = gen.PtrOf(gen.Int())
}
