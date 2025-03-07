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

func Test_StorageAccountsTableServicesTable_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsTableServicesTable_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsTableServicesTable_Spec, StorageAccountsTableServicesTable_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsTableServicesTable_Spec runs a test to see if a specific instance of StorageAccountsTableServicesTable_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsTableServicesTable_Spec(subject StorageAccountsTableServicesTable_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsTableServicesTable_Spec
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

// Generator of StorageAccountsTableServicesTable_Spec instances for property testing - lazily instantiated by
// StorageAccountsTableServicesTable_SpecGenerator()
var storageAccountsTableServicesTable_SpecGenerator gopter.Gen

// StorageAccountsTableServicesTable_SpecGenerator returns a generator of StorageAccountsTableServicesTable_Spec instances for property testing.
// We first initialize storageAccountsTableServicesTable_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsTableServicesTable_SpecGenerator() gopter.Gen {
	if storageAccountsTableServicesTable_SpecGenerator != nil {
		return storageAccountsTableServicesTable_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_Spec(generators)
	storageAccountsTableServicesTable_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_Spec(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_Spec(generators)
	storageAccountsTableServicesTable_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableServicesTable_Spec{}), generators)

	return storageAccountsTableServicesTable_SpecGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsTableServicesTable_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsTableServicesTable_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(TablePropertiesGenerator())
}

func Test_TableAccessPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableAccessPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableAccessPolicy, TableAccessPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableAccessPolicy runs a test to see if a specific instance of TableAccessPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForTableAccessPolicy(subject TableAccessPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableAccessPolicy
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

// Generator of TableAccessPolicy instances for property testing - lazily instantiated by TableAccessPolicyGenerator()
var tableAccessPolicyGenerator gopter.Gen

// TableAccessPolicyGenerator returns a generator of TableAccessPolicy instances for property testing.
func TableAccessPolicyGenerator() gopter.Gen {
	if tableAccessPolicyGenerator != nil {
		return tableAccessPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableAccessPolicy(generators)
	tableAccessPolicyGenerator = gen.Struct(reflect.TypeOf(TableAccessPolicy{}), generators)

	return tableAccessPolicyGenerator
}

// AddIndependentPropertyGeneratorsForTableAccessPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableAccessPolicy(gens map[string]gopter.Gen) {
	gens["ExpiryTime"] = gen.PtrOf(gen.AlphaString())
	gens["Permission"] = gen.PtrOf(gen.AlphaString())
	gens["StartTime"] = gen.PtrOf(gen.AlphaString())
}

func Test_TableProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableProperties, TablePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableProperties runs a test to see if a specific instance of TableProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForTableProperties(subject TableProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableProperties
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

// Generator of TableProperties instances for property testing - lazily instantiated by TablePropertiesGenerator()
var tablePropertiesGenerator gopter.Gen

// TablePropertiesGenerator returns a generator of TableProperties instances for property testing.
func TablePropertiesGenerator() gopter.Gen {
	if tablePropertiesGenerator != nil {
		return tablePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForTableProperties(generators)
	tablePropertiesGenerator = gen.Struct(reflect.TypeOf(TableProperties{}), generators)

	return tablePropertiesGenerator
}

// AddRelatedPropertyGeneratorsForTableProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTableProperties(gens map[string]gopter.Gen) {
	gens["SignedIdentifiers"] = gen.SliceOf(TableSignedIdentifierGenerator())
}

func Test_TableSignedIdentifier_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TableSignedIdentifier via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTableSignedIdentifier, TableSignedIdentifierGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTableSignedIdentifier runs a test to see if a specific instance of TableSignedIdentifier round trips to JSON and back losslessly
func RunJSONSerializationTestForTableSignedIdentifier(subject TableSignedIdentifier) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TableSignedIdentifier
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

// Generator of TableSignedIdentifier instances for property testing - lazily instantiated by
// TableSignedIdentifierGenerator()
var tableSignedIdentifierGenerator gopter.Gen

// TableSignedIdentifierGenerator returns a generator of TableSignedIdentifier instances for property testing.
// We first initialize tableSignedIdentifierGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TableSignedIdentifierGenerator() gopter.Gen {
	if tableSignedIdentifierGenerator != nil {
		return tableSignedIdentifierGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableSignedIdentifier(generators)
	tableSignedIdentifierGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTableSignedIdentifier(generators)
	AddRelatedPropertyGeneratorsForTableSignedIdentifier(generators)
	tableSignedIdentifierGenerator = gen.Struct(reflect.TypeOf(TableSignedIdentifier{}), generators)

	return tableSignedIdentifierGenerator
}

// AddIndependentPropertyGeneratorsForTableSignedIdentifier is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTableSignedIdentifier(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTableSignedIdentifier is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTableSignedIdentifier(gens map[string]gopter.Gen) {
	gens["AccessPolicy"] = gen.PtrOf(TableAccessPolicyGenerator())
}
