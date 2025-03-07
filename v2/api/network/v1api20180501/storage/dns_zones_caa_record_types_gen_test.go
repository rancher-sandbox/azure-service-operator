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

func Test_DnsZonesCAARecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesCAARecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesCAARecord, DnsZonesCAARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesCAARecord runs a test to see if a specific instance of DnsZonesCAARecord round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesCAARecord(subject DnsZonesCAARecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesCAARecord
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

// Generator of DnsZonesCAARecord instances for property testing - lazily instantiated by DnsZonesCAARecordGenerator()
var dnsZonesCAARecordGenerator gopter.Gen

// DnsZonesCAARecordGenerator returns a generator of DnsZonesCAARecord instances for property testing.
func DnsZonesCAARecordGenerator() gopter.Gen {
	if dnsZonesCAARecordGenerator != nil {
		return dnsZonesCAARecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsZonesCAARecord(generators)
	dnsZonesCAARecordGenerator = gen.Struct(reflect.TypeOf(DnsZonesCAARecord{}), generators)

	return dnsZonesCAARecordGenerator
}

// AddRelatedPropertyGeneratorsForDnsZonesCAARecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesCAARecord(gens map[string]gopter.Gen) {
	gens["Spec"] = DnsZonesCAARecord_SpecGenerator()
	gens["Status"] = DnsZonesCAARecord_STATUSGenerator()
}

func Test_DnsZonesCAARecordOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesCAARecordOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesCAARecordOperatorSpec, DnsZonesCAARecordOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesCAARecordOperatorSpec runs a test to see if a specific instance of DnsZonesCAARecordOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesCAARecordOperatorSpec(subject DnsZonesCAARecordOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesCAARecordOperatorSpec
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

// Generator of DnsZonesCAARecordOperatorSpec instances for property testing - lazily instantiated by
// DnsZonesCAARecordOperatorSpecGenerator()
var dnsZonesCAARecordOperatorSpecGenerator gopter.Gen

// DnsZonesCAARecordOperatorSpecGenerator returns a generator of DnsZonesCAARecordOperatorSpec instances for property testing.
func DnsZonesCAARecordOperatorSpecGenerator() gopter.Gen {
	if dnsZonesCAARecordOperatorSpecGenerator != nil {
		return dnsZonesCAARecordOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	dnsZonesCAARecordOperatorSpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesCAARecordOperatorSpec{}), generators)

	return dnsZonesCAARecordOperatorSpecGenerator
}

func Test_DnsZonesCAARecord_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesCAARecord_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesCAARecord_STATUS, DnsZonesCAARecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesCAARecord_STATUS runs a test to see if a specific instance of DnsZonesCAARecord_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesCAARecord_STATUS(subject DnsZonesCAARecord_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesCAARecord_STATUS
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

// Generator of DnsZonesCAARecord_STATUS instances for property testing - lazily instantiated by
// DnsZonesCAARecord_STATUSGenerator()
var dnsZonesCAARecord_STATUSGenerator gopter.Gen

// DnsZonesCAARecord_STATUSGenerator returns a generator of DnsZonesCAARecord_STATUS instances for property testing.
// We first initialize dnsZonesCAARecord_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesCAARecord_STATUSGenerator() gopter.Gen {
	if dnsZonesCAARecord_STATUSGenerator != nil {
		return dnsZonesCAARecord_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesCAARecord_STATUS(generators)
	dnsZonesCAARecord_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZonesCAARecord_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesCAARecord_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsZonesCAARecord_STATUS(generators)
	dnsZonesCAARecord_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZonesCAARecord_STATUS{}), generators)

	return dnsZonesCAARecord_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesCAARecord_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesCAARecord_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["TTL"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsZonesCAARecord_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesCAARecord_STATUS(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecord_STATUSGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecord_STATUSGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResource_STATUSGenerator())
}

func Test_DnsZonesCAARecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesCAARecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesCAARecord_Spec, DnsZonesCAARecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesCAARecord_Spec runs a test to see if a specific instance of DnsZonesCAARecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesCAARecord_Spec(subject DnsZonesCAARecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesCAARecord_Spec
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

// Generator of DnsZonesCAARecord_Spec instances for property testing - lazily instantiated by
// DnsZonesCAARecord_SpecGenerator()
var dnsZonesCAARecord_SpecGenerator gopter.Gen

// DnsZonesCAARecord_SpecGenerator returns a generator of DnsZonesCAARecord_Spec instances for property testing.
// We first initialize dnsZonesCAARecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesCAARecord_SpecGenerator() gopter.Gen {
	if dnsZonesCAARecord_SpecGenerator != nil {
		return dnsZonesCAARecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesCAARecord_Spec(generators)
	dnsZonesCAARecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesCAARecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesCAARecord_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZonesCAARecord_Spec(generators)
	dnsZonesCAARecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesCAARecord_Spec{}), generators)

	return dnsZonesCAARecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesCAARecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesCAARecord_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["TTL"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDnsZonesCAARecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesCAARecord_Spec(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecordGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecordGenerator())
	gens["OperatorSpec"] = gen.PtrOf(DnsZonesCAARecordOperatorSpecGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecordGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResourceGenerator())
}
