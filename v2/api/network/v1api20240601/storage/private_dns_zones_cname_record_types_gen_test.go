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

func Test_PrivateDnsZonesCNAMERecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesCNAMERecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesCNAMERecord, PrivateDnsZonesCNAMERecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesCNAMERecord runs a test to see if a specific instance of PrivateDnsZonesCNAMERecord round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesCNAMERecord(subject PrivateDnsZonesCNAMERecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesCNAMERecord
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

// Generator of PrivateDnsZonesCNAMERecord instances for property testing - lazily instantiated by
// PrivateDnsZonesCNAMERecordGenerator()
var privateDnsZonesCNAMERecordGenerator gopter.Gen

// PrivateDnsZonesCNAMERecordGenerator returns a generator of PrivateDnsZonesCNAMERecord instances for property testing.
func PrivateDnsZonesCNAMERecordGenerator() gopter.Gen {
	if privateDnsZonesCNAMERecordGenerator != nil {
		return privateDnsZonesCNAMERecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord(generators)
	privateDnsZonesCNAMERecordGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecord{}), generators)

	return privateDnsZonesCNAMERecordGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord(gens map[string]gopter.Gen) {
	gens["Spec"] = PrivateDnsZonesCNAMERecord_SpecGenerator()
	gens["Status"] = PrivateDnsZonesCNAMERecord_STATUSGenerator()
}

func Test_PrivateDnsZonesCNAMERecordOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesCNAMERecordOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesCNAMERecordOperatorSpec, PrivateDnsZonesCNAMERecordOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesCNAMERecordOperatorSpec runs a test to see if a specific instance of PrivateDnsZonesCNAMERecordOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesCNAMERecordOperatorSpec(subject PrivateDnsZonesCNAMERecordOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesCNAMERecordOperatorSpec
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

// Generator of PrivateDnsZonesCNAMERecordOperatorSpec instances for property testing - lazily instantiated by
// PrivateDnsZonesCNAMERecordOperatorSpecGenerator()
var privateDnsZonesCNAMERecordOperatorSpecGenerator gopter.Gen

// PrivateDnsZonesCNAMERecordOperatorSpecGenerator returns a generator of PrivateDnsZonesCNAMERecordOperatorSpec instances for property testing.
func PrivateDnsZonesCNAMERecordOperatorSpecGenerator() gopter.Gen {
	if privateDnsZonesCNAMERecordOperatorSpecGenerator != nil {
		return privateDnsZonesCNAMERecordOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	privateDnsZonesCNAMERecordOperatorSpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecordOperatorSpec{}), generators)

	return privateDnsZonesCNAMERecordOperatorSpecGenerator
}

func Test_PrivateDnsZonesCNAMERecord_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesCNAMERecord_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_STATUS, PrivateDnsZonesCNAMERecord_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_STATUS runs a test to see if a specific instance of PrivateDnsZonesCNAMERecord_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_STATUS(subject PrivateDnsZonesCNAMERecord_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesCNAMERecord_STATUS
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

// Generator of PrivateDnsZonesCNAMERecord_STATUS instances for property testing - lazily instantiated by
// PrivateDnsZonesCNAMERecord_STATUSGenerator()
var privateDnsZonesCNAMERecord_STATUSGenerator gopter.Gen

// PrivateDnsZonesCNAMERecord_STATUSGenerator returns a generator of PrivateDnsZonesCNAMERecord_STATUS instances for property testing.
// We first initialize privateDnsZonesCNAMERecord_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesCNAMERecord_STATUSGenerator() gopter.Gen {
	if privateDnsZonesCNAMERecord_STATUSGenerator != nil {
		return privateDnsZonesCNAMERecord_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_STATUS(generators)
	privateDnsZonesCNAMERecord_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecord_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_STATUS(generators)
	privateDnsZonesCNAMERecord_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecord_STATUS{}), generators)

	return privateDnsZonesCNAMERecord_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsAutoRegistered"] = gen.PtrOf(gen.Bool())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Ttl"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_STATUS(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
}

func Test_PrivateDnsZonesCNAMERecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesCNAMERecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_Spec, PrivateDnsZonesCNAMERecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_Spec runs a test to see if a specific instance of PrivateDnsZonesCNAMERecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_Spec(subject PrivateDnsZonesCNAMERecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesCNAMERecord_Spec
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

// Generator of PrivateDnsZonesCNAMERecord_Spec instances for property testing - lazily instantiated by
// PrivateDnsZonesCNAMERecord_SpecGenerator()
var privateDnsZonesCNAMERecord_SpecGenerator gopter.Gen

// PrivateDnsZonesCNAMERecord_SpecGenerator returns a generator of PrivateDnsZonesCNAMERecord_Spec instances for property testing.
// We first initialize privateDnsZonesCNAMERecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesCNAMERecord_SpecGenerator() gopter.Gen {
	if privateDnsZonesCNAMERecord_SpecGenerator != nil {
		return privateDnsZonesCNAMERecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec(generators)
	privateDnsZonesCNAMERecord_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec(generators)
	privateDnsZonesCNAMERecord_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecord_Spec{}), generators)

	return privateDnsZonesCNAMERecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Ttl"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec(gens map[string]gopter.Gen) {
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["AaaaRecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["CnameRecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["MxRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["OperatorSpec"] = gen.PtrOf(PrivateDnsZonesCNAMERecordOperatorSpecGenerator())
	gens["PtrRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SoaRecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SrvRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TxtRecords"] = gen.SliceOf(TxtRecordGenerator())
}
