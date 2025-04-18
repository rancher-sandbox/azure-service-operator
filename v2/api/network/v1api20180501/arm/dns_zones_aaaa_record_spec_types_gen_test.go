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

func Test_ARecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ARecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForARecord, ARecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForARecord runs a test to see if a specific instance of ARecord round trips to JSON and back losslessly
func RunJSONSerializationTestForARecord(subject ARecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ARecord
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

// Generator of ARecord instances for property testing - lazily instantiated by ARecordGenerator()
var aRecordGenerator gopter.Gen

// ARecordGenerator returns a generator of ARecord instances for property testing.
func ARecordGenerator() gopter.Gen {
	if aRecordGenerator != nil {
		return aRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForARecord(generators)
	aRecordGenerator = gen.Struct(reflect.TypeOf(ARecord{}), generators)

	return aRecordGenerator
}

// AddIndependentPropertyGeneratorsForARecord is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForARecord(gens map[string]gopter.Gen) {
	gens["Ipv4Address"] = gen.PtrOf(gen.AlphaString())
}

func Test_AaaaRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AaaaRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAaaaRecord, AaaaRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAaaaRecord runs a test to see if a specific instance of AaaaRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForAaaaRecord(subject AaaaRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AaaaRecord
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

// Generator of AaaaRecord instances for property testing - lazily instantiated by AaaaRecordGenerator()
var aaaaRecordGenerator gopter.Gen

// AaaaRecordGenerator returns a generator of AaaaRecord instances for property testing.
func AaaaRecordGenerator() gopter.Gen {
	if aaaaRecordGenerator != nil {
		return aaaaRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAaaaRecord(generators)
	aaaaRecordGenerator = gen.Struct(reflect.TypeOf(AaaaRecord{}), generators)

	return aaaaRecordGenerator
}

// AddIndependentPropertyGeneratorsForAaaaRecord is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAaaaRecord(gens map[string]gopter.Gen) {
	gens["Ipv6Address"] = gen.PtrOf(gen.AlphaString())
}

func Test_CaaRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CaaRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCaaRecord, CaaRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCaaRecord runs a test to see if a specific instance of CaaRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForCaaRecord(subject CaaRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CaaRecord
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

// Generator of CaaRecord instances for property testing - lazily instantiated by CaaRecordGenerator()
var caaRecordGenerator gopter.Gen

// CaaRecordGenerator returns a generator of CaaRecord instances for property testing.
func CaaRecordGenerator() gopter.Gen {
	if caaRecordGenerator != nil {
		return caaRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaaRecord(generators)
	caaRecordGenerator = gen.Struct(reflect.TypeOf(CaaRecord{}), generators)

	return caaRecordGenerator
}

// AddIndependentPropertyGeneratorsForCaaRecord is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCaaRecord(gens map[string]gopter.Gen) {
	gens["Flags"] = gen.PtrOf(gen.Int())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

func Test_CnameRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CnameRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCnameRecord, CnameRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCnameRecord runs a test to see if a specific instance of CnameRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForCnameRecord(subject CnameRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CnameRecord
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

// Generator of CnameRecord instances for property testing - lazily instantiated by CnameRecordGenerator()
var cnameRecordGenerator gopter.Gen

// CnameRecordGenerator returns a generator of CnameRecord instances for property testing.
func CnameRecordGenerator() gopter.Gen {
	if cnameRecordGenerator != nil {
		return cnameRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCnameRecord(generators)
	cnameRecordGenerator = gen.Struct(reflect.TypeOf(CnameRecord{}), generators)

	return cnameRecordGenerator
}

// AddIndependentPropertyGeneratorsForCnameRecord is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCnameRecord(gens map[string]gopter.Gen) {
	gens["Cname"] = gen.PtrOf(gen.AlphaString())
}

func Test_DnsZonesAAAARecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesAAAARecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesAAAARecord_Spec, DnsZonesAAAARecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesAAAARecord_Spec runs a test to see if a specific instance of DnsZonesAAAARecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesAAAARecord_Spec(subject DnsZonesAAAARecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesAAAARecord_Spec
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

// Generator of DnsZonesAAAARecord_Spec instances for property testing - lazily instantiated by
// DnsZonesAAAARecord_SpecGenerator()
var dnsZonesAAAARecord_SpecGenerator gopter.Gen

// DnsZonesAAAARecord_SpecGenerator returns a generator of DnsZonesAAAARecord_Spec instances for property testing.
// We first initialize dnsZonesAAAARecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesAAAARecord_SpecGenerator() gopter.Gen {
	if dnsZonesAAAARecord_SpecGenerator != nil {
		return dnsZonesAAAARecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesAAAARecord_Spec(generators)
	dnsZonesAAAARecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesAAAARecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesAAAARecord_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZonesAAAARecord_Spec(generators)
	dnsZonesAAAARecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesAAAARecord_Spec{}), generators)

	return dnsZonesAAAARecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesAAAARecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesAAAARecord_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForDnsZonesAAAARecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesAAAARecord_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetPropertiesGenerator())
}

func Test_MxRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MxRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMxRecord, MxRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMxRecord runs a test to see if a specific instance of MxRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForMxRecord(subject MxRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MxRecord
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

// Generator of MxRecord instances for property testing - lazily instantiated by MxRecordGenerator()
var mxRecordGenerator gopter.Gen

// MxRecordGenerator returns a generator of MxRecord instances for property testing.
func MxRecordGenerator() gopter.Gen {
	if mxRecordGenerator != nil {
		return mxRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMxRecord(generators)
	mxRecordGenerator = gen.Struct(reflect.TypeOf(MxRecord{}), generators)

	return mxRecordGenerator
}

// AddIndependentPropertyGeneratorsForMxRecord is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMxRecord(gens map[string]gopter.Gen) {
	gens["Exchange"] = gen.PtrOf(gen.AlphaString())
	gens["Preference"] = gen.PtrOf(gen.Int())
}

func Test_NsRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NsRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNsRecord, NsRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNsRecord runs a test to see if a specific instance of NsRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForNsRecord(subject NsRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NsRecord
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

// Generator of NsRecord instances for property testing - lazily instantiated by NsRecordGenerator()
var nsRecordGenerator gopter.Gen

// NsRecordGenerator returns a generator of NsRecord instances for property testing.
func NsRecordGenerator() gopter.Gen {
	if nsRecordGenerator != nil {
		return nsRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNsRecord(generators)
	nsRecordGenerator = gen.Struct(reflect.TypeOf(NsRecord{}), generators)

	return nsRecordGenerator
}

// AddIndependentPropertyGeneratorsForNsRecord is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNsRecord(gens map[string]gopter.Gen) {
	gens["Nsdname"] = gen.PtrOf(gen.AlphaString())
}

func Test_PtrRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PtrRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPtrRecord, PtrRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPtrRecord runs a test to see if a specific instance of PtrRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForPtrRecord(subject PtrRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PtrRecord
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

// Generator of PtrRecord instances for property testing - lazily instantiated by PtrRecordGenerator()
var ptrRecordGenerator gopter.Gen

// PtrRecordGenerator returns a generator of PtrRecord instances for property testing.
func PtrRecordGenerator() gopter.Gen {
	if ptrRecordGenerator != nil {
		return ptrRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPtrRecord(generators)
	ptrRecordGenerator = gen.Struct(reflect.TypeOf(PtrRecord{}), generators)

	return ptrRecordGenerator
}

// AddIndependentPropertyGeneratorsForPtrRecord is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPtrRecord(gens map[string]gopter.Gen) {
	gens["Ptrdname"] = gen.PtrOf(gen.AlphaString())
}

func Test_RecordSetProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RecordSetProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRecordSetProperties, RecordSetPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRecordSetProperties runs a test to see if a specific instance of RecordSetProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForRecordSetProperties(subject RecordSetProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RecordSetProperties
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

// Generator of RecordSetProperties instances for property testing - lazily instantiated by
// RecordSetPropertiesGenerator()
var recordSetPropertiesGenerator gopter.Gen

// RecordSetPropertiesGenerator returns a generator of RecordSetProperties instances for property testing.
// We first initialize recordSetPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RecordSetPropertiesGenerator() gopter.Gen {
	if recordSetPropertiesGenerator != nil {
		return recordSetPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRecordSetProperties(generators)
	recordSetPropertiesGenerator = gen.Struct(reflect.TypeOf(RecordSetProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRecordSetProperties(generators)
	AddRelatedPropertyGeneratorsForRecordSetProperties(generators)
	recordSetPropertiesGenerator = gen.Struct(reflect.TypeOf(RecordSetProperties{}), generators)

	return recordSetPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForRecordSetProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRecordSetProperties(gens map[string]gopter.Gen) {
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["TTL"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForRecordSetProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRecordSetProperties(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecordGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecordGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecordGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResourceGenerator())
}

func Test_SoaRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SoaRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSoaRecord, SoaRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSoaRecord runs a test to see if a specific instance of SoaRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForSoaRecord(subject SoaRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SoaRecord
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

// Generator of SoaRecord instances for property testing - lazily instantiated by SoaRecordGenerator()
var soaRecordGenerator gopter.Gen

// SoaRecordGenerator returns a generator of SoaRecord instances for property testing.
func SoaRecordGenerator() gopter.Gen {
	if soaRecordGenerator != nil {
		return soaRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSoaRecord(generators)
	soaRecordGenerator = gen.Struct(reflect.TypeOf(SoaRecord{}), generators)

	return soaRecordGenerator
}

// AddIndependentPropertyGeneratorsForSoaRecord is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSoaRecord(gens map[string]gopter.Gen) {
	gens["Email"] = gen.PtrOf(gen.AlphaString())
	gens["ExpireTime"] = gen.PtrOf(gen.Int())
	gens["Host"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTTL"] = gen.PtrOf(gen.Int())
	gens["RefreshTime"] = gen.PtrOf(gen.Int())
	gens["RetryTime"] = gen.PtrOf(gen.Int())
	gens["SerialNumber"] = gen.PtrOf(gen.Int())
}

func Test_SrvRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SrvRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSrvRecord, SrvRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSrvRecord runs a test to see if a specific instance of SrvRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForSrvRecord(subject SrvRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SrvRecord
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

// Generator of SrvRecord instances for property testing - lazily instantiated by SrvRecordGenerator()
var srvRecordGenerator gopter.Gen

// SrvRecordGenerator returns a generator of SrvRecord instances for property testing.
func SrvRecordGenerator() gopter.Gen {
	if srvRecordGenerator != nil {
		return srvRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSrvRecord(generators)
	srvRecordGenerator = gen.Struct(reflect.TypeOf(SrvRecord{}), generators)

	return srvRecordGenerator
}

// AddIndependentPropertyGeneratorsForSrvRecord is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSrvRecord(gens map[string]gopter.Gen) {
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Weight"] = gen.PtrOf(gen.Int())
}

func Test_TxtRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TxtRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTxtRecord, TxtRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTxtRecord runs a test to see if a specific instance of TxtRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForTxtRecord(subject TxtRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TxtRecord
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

// Generator of TxtRecord instances for property testing - lazily instantiated by TxtRecordGenerator()
var txtRecordGenerator gopter.Gen

// TxtRecordGenerator returns a generator of TxtRecord instances for property testing.
func TxtRecordGenerator() gopter.Gen {
	if txtRecordGenerator != nil {
		return txtRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTxtRecord(generators)
	txtRecordGenerator = gen.Struct(reflect.TypeOf(TxtRecord{}), generators)

	return txtRecordGenerator
}

// AddIndependentPropertyGeneratorsForTxtRecord is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTxtRecord(gens map[string]gopter.Gen) {
	gens["Value"] = gen.SliceOf(gen.AlphaString())
}
