// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

import (
	"encoding/json"
	v20220801s "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	v20230501ps "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/storage"
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

func Test_KeyVaultContractCreateProperties_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from KeyVaultContractCreateProperties to KeyVaultContractCreateProperties via AssignProperties_To_KeyVaultContractCreateProperties & AssignProperties_From_KeyVaultContractCreateProperties returns original",
		prop.ForAll(RunPropertyAssignmentTestForKeyVaultContractCreateProperties, KeyVaultContractCreatePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForKeyVaultContractCreateProperties tests if a specific instance of KeyVaultContractCreateProperties can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForKeyVaultContractCreateProperties(subject KeyVaultContractCreateProperties) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.KeyVaultContractCreateProperties
	err := copied.AssignProperties_To_KeyVaultContractCreateProperties(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual KeyVaultContractCreateProperties
	err = actual.AssignProperties_From_KeyVaultContractCreateProperties(&other)
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

func Test_KeyVaultContractCreateProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultContractCreateProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultContractCreateProperties, KeyVaultContractCreatePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultContractCreateProperties runs a test to see if a specific instance of KeyVaultContractCreateProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultContractCreateProperties(subject KeyVaultContractCreateProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultContractCreateProperties
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

// Generator of KeyVaultContractCreateProperties instances for property testing - lazily instantiated by
// KeyVaultContractCreatePropertiesGenerator()
var keyVaultContractCreatePropertiesGenerator gopter.Gen

// KeyVaultContractCreatePropertiesGenerator returns a generator of KeyVaultContractCreateProperties instances for property testing.
func KeyVaultContractCreatePropertiesGenerator() gopter.Gen {
	if keyVaultContractCreatePropertiesGenerator != nil {
		return keyVaultContractCreatePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultContractCreateProperties(generators)
	keyVaultContractCreatePropertiesGenerator = gen.Struct(reflect.TypeOf(KeyVaultContractCreateProperties{}), generators)

	return keyVaultContractCreatePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultContractCreateProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultContractCreateProperties(gens map[string]gopter.Gen) {
	gens["IdentityClientId"] = gen.PtrOf(gen.AlphaString())
	gens["SecretIdentifier"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultContractProperties_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from KeyVaultContractProperties_STATUS to KeyVaultContractProperties_STATUS via AssignProperties_To_KeyVaultContractProperties_STATUS & AssignProperties_From_KeyVaultContractProperties_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForKeyVaultContractProperties_STATUS, KeyVaultContractProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForKeyVaultContractProperties_STATUS tests if a specific instance of KeyVaultContractProperties_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForKeyVaultContractProperties_STATUS(subject KeyVaultContractProperties_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.KeyVaultContractProperties_STATUS
	err := copied.AssignProperties_To_KeyVaultContractProperties_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual KeyVaultContractProperties_STATUS
	err = actual.AssignProperties_From_KeyVaultContractProperties_STATUS(&other)
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

func Test_KeyVaultContractProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultContractProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultContractProperties_STATUS, KeyVaultContractProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultContractProperties_STATUS runs a test to see if a specific instance of KeyVaultContractProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultContractProperties_STATUS(subject KeyVaultContractProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultContractProperties_STATUS
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

// Generator of KeyVaultContractProperties_STATUS instances for property testing - lazily instantiated by
// KeyVaultContractProperties_STATUSGenerator()
var keyVaultContractProperties_STATUSGenerator gopter.Gen

// KeyVaultContractProperties_STATUSGenerator returns a generator of KeyVaultContractProperties_STATUS instances for property testing.
// We first initialize keyVaultContractProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyVaultContractProperties_STATUSGenerator() gopter.Gen {
	if keyVaultContractProperties_STATUSGenerator != nil {
		return keyVaultContractProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultContractProperties_STATUS(generators)
	keyVaultContractProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(KeyVaultContractProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultContractProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForKeyVaultContractProperties_STATUS(generators)
	keyVaultContractProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(KeyVaultContractProperties_STATUS{}), generators)

	return keyVaultContractProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultContractProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultContractProperties_STATUS(gens map[string]gopter.Gen) {
	gens["IdentityClientId"] = gen.PtrOf(gen.AlphaString())
	gens["SecretIdentifier"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyVaultContractProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyVaultContractProperties_STATUS(gens map[string]gopter.Gen) {
	gens["LastStatus"] = gen.PtrOf(KeyVaultLastAccessStatusContractProperties_STATUSGenerator())
}

func Test_KeyVaultLastAccessStatusContractProperties_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from KeyVaultLastAccessStatusContractProperties_STATUS to KeyVaultLastAccessStatusContractProperties_STATUS via AssignProperties_To_KeyVaultLastAccessStatusContractProperties_STATUS & AssignProperties_From_KeyVaultLastAccessStatusContractProperties_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForKeyVaultLastAccessStatusContractProperties_STATUS, KeyVaultLastAccessStatusContractProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForKeyVaultLastAccessStatusContractProperties_STATUS tests if a specific instance of KeyVaultLastAccessStatusContractProperties_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForKeyVaultLastAccessStatusContractProperties_STATUS(subject KeyVaultLastAccessStatusContractProperties_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.KeyVaultLastAccessStatusContractProperties_STATUS
	err := copied.AssignProperties_To_KeyVaultLastAccessStatusContractProperties_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual KeyVaultLastAccessStatusContractProperties_STATUS
	err = actual.AssignProperties_From_KeyVaultLastAccessStatusContractProperties_STATUS(&other)
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

func Test_KeyVaultLastAccessStatusContractProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultLastAccessStatusContractProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultLastAccessStatusContractProperties_STATUS, KeyVaultLastAccessStatusContractProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultLastAccessStatusContractProperties_STATUS runs a test to see if a specific instance of KeyVaultLastAccessStatusContractProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultLastAccessStatusContractProperties_STATUS(subject KeyVaultLastAccessStatusContractProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultLastAccessStatusContractProperties_STATUS
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

// Generator of KeyVaultLastAccessStatusContractProperties_STATUS instances for property testing - lazily instantiated
// by KeyVaultLastAccessStatusContractProperties_STATUSGenerator()
var keyVaultLastAccessStatusContractProperties_STATUSGenerator gopter.Gen

// KeyVaultLastAccessStatusContractProperties_STATUSGenerator returns a generator of KeyVaultLastAccessStatusContractProperties_STATUS instances for property testing.
func KeyVaultLastAccessStatusContractProperties_STATUSGenerator() gopter.Gen {
	if keyVaultLastAccessStatusContractProperties_STATUSGenerator != nil {
		return keyVaultLastAccessStatusContractProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultLastAccessStatusContractProperties_STATUS(generators)
	keyVaultLastAccessStatusContractProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(KeyVaultLastAccessStatusContractProperties_STATUS{}), generators)

	return keyVaultLastAccessStatusContractProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultLastAccessStatusContractProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultLastAccessStatusContractProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["TimeStampUtc"] = gen.PtrOf(gen.AlphaString())
}

func Test_NamedValue_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamedValue to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamedValue, NamedValueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamedValue tests if a specific instance of NamedValue round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamedValue(subject NamedValue) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220801s.NamedValue
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamedValue
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

func Test_NamedValue_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamedValue to NamedValue via AssignProperties_To_NamedValue & AssignProperties_From_NamedValue returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamedValue, NamedValueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamedValue tests if a specific instance of NamedValue can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamedValue(subject NamedValue) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.NamedValue
	err := copied.AssignProperties_To_NamedValue(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamedValue
	err = actual.AssignProperties_From_NamedValue(&other)
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

func Test_NamedValue_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamedValue via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamedValue, NamedValueGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamedValue runs a test to see if a specific instance of NamedValue round trips to JSON and back losslessly
func RunJSONSerializationTestForNamedValue(subject NamedValue) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamedValue
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

// Generator of NamedValue instances for property testing - lazily instantiated by NamedValueGenerator()
var namedValueGenerator gopter.Gen

// NamedValueGenerator returns a generator of NamedValue instances for property testing.
func NamedValueGenerator() gopter.Gen {
	if namedValueGenerator != nil {
		return namedValueGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamedValue(generators)
	namedValueGenerator = gen.Struct(reflect.TypeOf(NamedValue{}), generators)

	return namedValueGenerator
}

// AddRelatedPropertyGeneratorsForNamedValue is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamedValue(gens map[string]gopter.Gen) {
	gens["Spec"] = NamedValue_SpecGenerator()
	gens["Status"] = NamedValue_STATUSGenerator()
}

func Test_NamedValueOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamedValueOperatorSpec to NamedValueOperatorSpec via AssignProperties_To_NamedValueOperatorSpec & AssignProperties_From_NamedValueOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamedValueOperatorSpec, NamedValueOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamedValueOperatorSpec tests if a specific instance of NamedValueOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamedValueOperatorSpec(subject NamedValueOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.NamedValueOperatorSpec
	err := copied.AssignProperties_To_NamedValueOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamedValueOperatorSpec
	err = actual.AssignProperties_From_NamedValueOperatorSpec(&other)
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

func Test_NamedValueOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamedValueOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamedValueOperatorSpec, NamedValueOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamedValueOperatorSpec runs a test to see if a specific instance of NamedValueOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamedValueOperatorSpec(subject NamedValueOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamedValueOperatorSpec
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

// Generator of NamedValueOperatorSpec instances for property testing - lazily instantiated by
// NamedValueOperatorSpecGenerator()
var namedValueOperatorSpecGenerator gopter.Gen

// NamedValueOperatorSpecGenerator returns a generator of NamedValueOperatorSpec instances for property testing.
func NamedValueOperatorSpecGenerator() gopter.Gen {
	if namedValueOperatorSpecGenerator != nil {
		return namedValueOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	namedValueOperatorSpecGenerator = gen.Struct(reflect.TypeOf(NamedValueOperatorSpec{}), generators)

	return namedValueOperatorSpecGenerator
}

func Test_NamedValue_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamedValue_STATUS to NamedValue_STATUS via AssignProperties_To_NamedValue_STATUS & AssignProperties_From_NamedValue_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamedValue_STATUS, NamedValue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamedValue_STATUS tests if a specific instance of NamedValue_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamedValue_STATUS(subject NamedValue_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.NamedValue_STATUS
	err := copied.AssignProperties_To_NamedValue_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamedValue_STATUS
	err = actual.AssignProperties_From_NamedValue_STATUS(&other)
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

func Test_NamedValue_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamedValue_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamedValue_STATUS, NamedValue_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamedValue_STATUS runs a test to see if a specific instance of NamedValue_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamedValue_STATUS(subject NamedValue_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamedValue_STATUS
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

// Generator of NamedValue_STATUS instances for property testing - lazily instantiated by NamedValue_STATUSGenerator()
var namedValue_STATUSGenerator gopter.Gen

// NamedValue_STATUSGenerator returns a generator of NamedValue_STATUS instances for property testing.
// We first initialize namedValue_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamedValue_STATUSGenerator() gopter.Gen {
	if namedValue_STATUSGenerator != nil {
		return namedValue_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamedValue_STATUS(generators)
	namedValue_STATUSGenerator = gen.Struct(reflect.TypeOf(NamedValue_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamedValue_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamedValue_STATUS(generators)
	namedValue_STATUSGenerator = gen.Struct(reflect.TypeOf(NamedValue_STATUS{}), generators)

	return namedValue_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamedValue_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamedValue_STATUS(gens map[string]gopter.Gen) {
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Secret"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamedValue_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamedValue_STATUS(gens map[string]gopter.Gen) {
	gens["KeyVault"] = gen.PtrOf(KeyVaultContractProperties_STATUSGenerator())
}

func Test_NamedValue_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamedValue_Spec to NamedValue_Spec via AssignProperties_To_NamedValue_Spec & AssignProperties_From_NamedValue_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamedValue_Spec, NamedValue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamedValue_Spec tests if a specific instance of NamedValue_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamedValue_Spec(subject NamedValue_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.NamedValue_Spec
	err := copied.AssignProperties_To_NamedValue_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamedValue_Spec
	err = actual.AssignProperties_From_NamedValue_Spec(&other)
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

func Test_NamedValue_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamedValue_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamedValue_Spec, NamedValue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamedValue_Spec runs a test to see if a specific instance of NamedValue_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamedValue_Spec(subject NamedValue_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamedValue_Spec
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

// Generator of NamedValue_Spec instances for property testing - lazily instantiated by NamedValue_SpecGenerator()
var namedValue_SpecGenerator gopter.Gen

// NamedValue_SpecGenerator returns a generator of NamedValue_Spec instances for property testing.
// We first initialize namedValue_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamedValue_SpecGenerator() gopter.Gen {
	if namedValue_SpecGenerator != nil {
		return namedValue_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamedValue_Spec(generators)
	namedValue_SpecGenerator = gen.Struct(reflect.TypeOf(NamedValue_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamedValue_Spec(generators)
	AddRelatedPropertyGeneratorsForNamedValue_Spec(generators)
	namedValue_SpecGenerator = gen.Struct(reflect.TypeOf(NamedValue_Spec{}), generators)

	return namedValue_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamedValue_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamedValue_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["Secret"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.SliceOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamedValue_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamedValue_Spec(gens map[string]gopter.Gen) {
	gens["KeyVault"] = gen.PtrOf(KeyVaultContractCreatePropertiesGenerator())
	gens["OperatorSpec"] = gen.PtrOf(NamedValueOperatorSpecGenerator())
}
