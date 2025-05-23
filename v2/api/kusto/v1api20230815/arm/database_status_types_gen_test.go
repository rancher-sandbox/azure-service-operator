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

func Test_DatabaseStatistics_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseStatistics_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseStatistics_STATUS, DatabaseStatistics_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseStatistics_STATUS runs a test to see if a specific instance of DatabaseStatistics_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseStatistics_STATUS(subject DatabaseStatistics_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseStatistics_STATUS
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

// Generator of DatabaseStatistics_STATUS instances for property testing - lazily instantiated by
// DatabaseStatistics_STATUSGenerator()
var databaseStatistics_STATUSGenerator gopter.Gen

// DatabaseStatistics_STATUSGenerator returns a generator of DatabaseStatistics_STATUS instances for property testing.
func DatabaseStatistics_STATUSGenerator() gopter.Gen {
	if databaseStatistics_STATUSGenerator != nil {
		return databaseStatistics_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseStatistics_STATUS(generators)
	databaseStatistics_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseStatistics_STATUS{}), generators)

	return databaseStatistics_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseStatistics_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseStatistics_STATUS(gens map[string]gopter.Gen) {
	gens["Size"] = gen.PtrOf(gen.Float64())
}

func Test_Database_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase_STATUS, Database_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase_STATUS runs a test to see if a specific instance of Database_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase_STATUS(subject Database_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_STATUS
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

// Generator of Database_STATUS instances for property testing - lazily instantiated by Database_STATUSGenerator()
var database_STATUSGenerator gopter.Gen

// Database_STATUSGenerator returns a generator of Database_STATUS instances for property testing.
func Database_STATUSGenerator() gopter.Gen {
	if database_STATUSGenerator != nil {
		return database_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDatabase_STATUS(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		props := map[string]gopter.Gen{propName: propGen}
		gens = append(gens, gen.Struct(reflect.TypeOf(Database_STATUS{}), props))
	}
	database_STATUSGenerator = gen.OneGenOf(gens...)

	return database_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["ReadWrite"] = ReadWriteDatabase_STATUSGenerator().Map(func(it ReadWriteDatabase_STATUS) *ReadWriteDatabase_STATUS {
		return &it
	}) // generate one case for OneOf type
}

func Test_ReadWriteDatabaseProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadWriteDatabaseProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadWriteDatabaseProperties_STATUS, ReadWriteDatabaseProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadWriteDatabaseProperties_STATUS runs a test to see if a specific instance of ReadWriteDatabaseProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForReadWriteDatabaseProperties_STATUS(subject ReadWriteDatabaseProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadWriteDatabaseProperties_STATUS
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

// Generator of ReadWriteDatabaseProperties_STATUS instances for property testing - lazily instantiated by
// ReadWriteDatabaseProperties_STATUSGenerator()
var readWriteDatabaseProperties_STATUSGenerator gopter.Gen

// ReadWriteDatabaseProperties_STATUSGenerator returns a generator of ReadWriteDatabaseProperties_STATUS instances for property testing.
// We first initialize readWriteDatabaseProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadWriteDatabaseProperties_STATUSGenerator() gopter.Gen {
	if readWriteDatabaseProperties_STATUSGenerator != nil {
		return readWriteDatabaseProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS(generators)
	readWriteDatabaseProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabaseProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties_STATUS(generators)
	readWriteDatabaseProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabaseProperties_STATUS{}), generators)

	return readWriteDatabaseProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadWriteDatabaseProperties_STATUS(gens map[string]gopter.Gen) {
	gens["HotCachePeriod"] = gen.PtrOf(gen.AlphaString())
	gens["IsFollowed"] = gen.PtrOf(gen.Bool())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Canceled,
		ProvisioningState_STATUS_Creating,
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Moving,
		ProvisioningState_STATUS_Running,
		ProvisioningState_STATUS_Succeeded))
	gens["SoftDeletePeriod"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadWriteDatabaseProperties_STATUS(gens map[string]gopter.Gen) {
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultProperties_STATUSGenerator())
	gens["Statistics"] = gen.PtrOf(DatabaseStatistics_STATUSGenerator())
	gens["SuspensionDetails"] = gen.PtrOf(SuspensionDetails_STATUSGenerator())
}

func Test_ReadWriteDatabase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReadWriteDatabase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReadWriteDatabase_STATUS, ReadWriteDatabase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReadWriteDatabase_STATUS runs a test to see if a specific instance of ReadWriteDatabase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForReadWriteDatabase_STATUS(subject ReadWriteDatabase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReadWriteDatabase_STATUS
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

// Generator of ReadWriteDatabase_STATUS instances for property testing - lazily instantiated by
// ReadWriteDatabase_STATUSGenerator()
var readWriteDatabase_STATUSGenerator gopter.Gen

// ReadWriteDatabase_STATUSGenerator returns a generator of ReadWriteDatabase_STATUS instances for property testing.
// We first initialize readWriteDatabase_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ReadWriteDatabase_STATUSGenerator() gopter.Gen {
	if readWriteDatabase_STATUSGenerator != nil {
		return readWriteDatabase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS(generators)
	readWriteDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabase_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForReadWriteDatabase_STATUS(generators)
	readWriteDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(ReadWriteDatabase_STATUS{}), generators)

	return readWriteDatabase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReadWriteDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Kind"] = gen.OneConstOf(ReadWriteDatabase_Kind_STATUS_ReadWrite)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForReadWriteDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForReadWriteDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ReadWriteDatabaseProperties_STATUSGenerator())
}

func Test_SuspensionDetails_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SuspensionDetails_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSuspensionDetails_STATUS, SuspensionDetails_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSuspensionDetails_STATUS runs a test to see if a specific instance of SuspensionDetails_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSuspensionDetails_STATUS(subject SuspensionDetails_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SuspensionDetails_STATUS
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

// Generator of SuspensionDetails_STATUS instances for property testing - lazily instantiated by
// SuspensionDetails_STATUSGenerator()
var suspensionDetails_STATUSGenerator gopter.Gen

// SuspensionDetails_STATUSGenerator returns a generator of SuspensionDetails_STATUS instances for property testing.
func SuspensionDetails_STATUSGenerator() gopter.Gen {
	if suspensionDetails_STATUSGenerator != nil {
		return suspensionDetails_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSuspensionDetails_STATUS(generators)
	suspensionDetails_STATUSGenerator = gen.Struct(reflect.TypeOf(SuspensionDetails_STATUS{}), generators)

	return suspensionDetails_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSuspensionDetails_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSuspensionDetails_STATUS(gens map[string]gopter.Gen) {
	gens["SuspensionStartDate"] = gen.PtrOf(gen.AlphaString())
}
