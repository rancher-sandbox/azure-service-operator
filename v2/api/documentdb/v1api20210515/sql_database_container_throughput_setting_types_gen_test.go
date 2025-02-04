// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

import (
	"encoding/json"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515/storage"
	v20240815s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/storage"
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

func Test_SqlDatabaseContainerThroughputSetting_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerThroughputSetting to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseContainerThroughputSetting, SqlDatabaseContainerThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseContainerThroughputSetting tests if a specific instance of SqlDatabaseContainerThroughputSetting round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseContainerThroughputSetting(subject SqlDatabaseContainerThroughputSetting) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20240815s.SqlDatabaseContainerThroughputSetting
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseContainerThroughputSetting
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseContainerThroughputSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerThroughputSetting to SqlDatabaseContainerThroughputSetting via AssignProperties_To_SqlDatabaseContainerThroughputSetting & AssignProperties_From_SqlDatabaseContainerThroughputSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting, SqlDatabaseContainerThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting tests if a specific instance of SqlDatabaseContainerThroughputSetting can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting(subject SqlDatabaseContainerThroughputSetting) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseContainerThroughputSetting
	err := copied.AssignProperties_To_SqlDatabaseContainerThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerThroughputSetting
	err = actual.AssignProperties_From_SqlDatabaseContainerThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseContainerThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting, SqlDatabaseContainerThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting runs a test to see if a specific instance of SqlDatabaseContainerThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting(subject SqlDatabaseContainerThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerThroughputSetting
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

// Generator of SqlDatabaseContainerThroughputSetting instances for property testing - lazily instantiated by
// SqlDatabaseContainerThroughputSettingGenerator()
var sqlDatabaseContainerThroughputSettingGenerator gopter.Gen

// SqlDatabaseContainerThroughputSettingGenerator returns a generator of SqlDatabaseContainerThroughputSetting instances for property testing.
func SqlDatabaseContainerThroughputSettingGenerator() gopter.Gen {
	if sqlDatabaseContainerThroughputSettingGenerator != nil {
		return sqlDatabaseContainerThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting(generators)
	sqlDatabaseContainerThroughputSettingGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSetting{}), generators)

	return sqlDatabaseContainerThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = SqlDatabaseContainerThroughputSetting_SpecGenerator()
	gens["Status"] = SqlDatabaseContainerThroughputSetting_STATUSGenerator()
}

func Test_SqlDatabaseContainerThroughputSettingOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerThroughputSettingOperatorSpec to SqlDatabaseContainerThroughputSettingOperatorSpec via AssignProperties_To_SqlDatabaseContainerThroughputSettingOperatorSpec & AssignProperties_From_SqlDatabaseContainerThroughputSettingOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSettingOperatorSpec, SqlDatabaseContainerThroughputSettingOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSettingOperatorSpec tests if a specific instance of SqlDatabaseContainerThroughputSettingOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSettingOperatorSpec(subject SqlDatabaseContainerThroughputSettingOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseContainerThroughputSettingOperatorSpec
	err := copied.AssignProperties_To_SqlDatabaseContainerThroughputSettingOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerThroughputSettingOperatorSpec
	err = actual.AssignProperties_From_SqlDatabaseContainerThroughputSettingOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseContainerThroughputSettingOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerThroughputSettingOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerThroughputSettingOperatorSpec, SqlDatabaseContainerThroughputSettingOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerThroughputSettingOperatorSpec runs a test to see if a specific instance of SqlDatabaseContainerThroughputSettingOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerThroughputSettingOperatorSpec(subject SqlDatabaseContainerThroughputSettingOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerThroughputSettingOperatorSpec
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

// Generator of SqlDatabaseContainerThroughputSettingOperatorSpec instances for property testing - lazily instantiated
// by SqlDatabaseContainerThroughputSettingOperatorSpecGenerator()
var sqlDatabaseContainerThroughputSettingOperatorSpecGenerator gopter.Gen

// SqlDatabaseContainerThroughputSettingOperatorSpecGenerator returns a generator of SqlDatabaseContainerThroughputSettingOperatorSpec instances for property testing.
func SqlDatabaseContainerThroughputSettingOperatorSpecGenerator() gopter.Gen {
	if sqlDatabaseContainerThroughputSettingOperatorSpecGenerator != nil {
		return sqlDatabaseContainerThroughputSettingOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	sqlDatabaseContainerThroughputSettingOperatorSpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSettingOperatorSpec{}), generators)

	return sqlDatabaseContainerThroughputSettingOperatorSpecGenerator
}

func Test_SqlDatabaseContainerThroughputSetting_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerThroughputSetting_STATUS to SqlDatabaseContainerThroughputSetting_STATUS via AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS & AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting_STATUS, SqlDatabaseContainerThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting_STATUS tests if a specific instance of SqlDatabaseContainerThroughputSetting_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting_STATUS(subject SqlDatabaseContainerThroughputSetting_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseContainerThroughputSetting_STATUS
	err := copied.AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerThroughputSetting_STATUS
	err = actual.AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseContainerThroughputSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerThroughputSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting_STATUS, SqlDatabaseContainerThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting_STATUS runs a test to see if a specific instance of SqlDatabaseContainerThroughputSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting_STATUS(subject SqlDatabaseContainerThroughputSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerThroughputSetting_STATUS
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

// Generator of SqlDatabaseContainerThroughputSetting_STATUS instances for property testing - lazily instantiated by
// SqlDatabaseContainerThroughputSetting_STATUSGenerator()
var sqlDatabaseContainerThroughputSetting_STATUSGenerator gopter.Gen

// SqlDatabaseContainerThroughputSetting_STATUSGenerator returns a generator of SqlDatabaseContainerThroughputSetting_STATUS instances for property testing.
// We first initialize sqlDatabaseContainerThroughputSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseContainerThroughputSetting_STATUSGenerator() gopter.Gen {
	if sqlDatabaseContainerThroughputSetting_STATUSGenerator != nil {
		return sqlDatabaseContainerThroughputSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS(generators)
	sqlDatabaseContainerThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS(generators)
	sqlDatabaseContainerThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSetting_STATUS{}), generators)

	return sqlDatabaseContainerThroughputSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsGetProperties_Resource_STATUSGenerator())
}

func Test_SqlDatabaseContainerThroughputSetting_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerThroughputSetting_Spec to SqlDatabaseContainerThroughputSetting_Spec via AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec & AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting_Spec, SqlDatabaseContainerThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting_Spec tests if a specific instance of SqlDatabaseContainerThroughputSetting_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerThroughputSetting_Spec(subject SqlDatabaseContainerThroughputSetting_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210515s.SqlDatabaseContainerThroughputSetting_Spec
	err := copied.AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerThroughputSetting_Spec
	err = actual.AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseContainerThroughputSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerThroughputSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting_Spec, SqlDatabaseContainerThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting_Spec runs a test to see if a specific instance of SqlDatabaseContainerThroughputSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting_Spec(subject SqlDatabaseContainerThroughputSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerThroughputSetting_Spec
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

// Generator of SqlDatabaseContainerThroughputSetting_Spec instances for property testing - lazily instantiated by
// SqlDatabaseContainerThroughputSetting_SpecGenerator()
var sqlDatabaseContainerThroughputSetting_SpecGenerator gopter.Gen

// SqlDatabaseContainerThroughputSetting_SpecGenerator returns a generator of SqlDatabaseContainerThroughputSetting_Spec instances for property testing.
// We first initialize sqlDatabaseContainerThroughputSetting_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseContainerThroughputSetting_SpecGenerator() gopter.Gen {
	if sqlDatabaseContainerThroughputSetting_SpecGenerator != nil {
		return sqlDatabaseContainerThroughputSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_Spec(generators)
	sqlDatabaseContainerThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSetting_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_Spec(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_Spec(generators)
	sqlDatabaseContainerThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSetting_Spec{}), generators)

	return sqlDatabaseContainerThroughputSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(SqlDatabaseContainerThroughputSettingOperatorSpecGenerator())
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}
