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

func Test_ContainerProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ContainerProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForContainerProperties_STATUS, ContainerProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForContainerProperties_STATUS runs a test to see if a specific instance of ContainerProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForContainerProperties_STATUS(subject ContainerProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ContainerProperties_STATUS
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

// Generator of ContainerProperties_STATUS instances for property testing - lazily instantiated by
// ContainerProperties_STATUSGenerator()
var containerProperties_STATUSGenerator gopter.Gen

// ContainerProperties_STATUSGenerator returns a generator of ContainerProperties_STATUS instances for property testing.
// We first initialize containerProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ContainerProperties_STATUSGenerator() gopter.Gen {
	if containerProperties_STATUSGenerator != nil {
		return containerProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerProperties_STATUS(generators)
	containerProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ContainerProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForContainerProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForContainerProperties_STATUS(generators)
	containerProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ContainerProperties_STATUS{}), generators)

	return containerProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForContainerProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForContainerProperties_STATUS(gens map[string]gopter.Gen) {
	gens["DefaultEncryptionScope"] = gen.PtrOf(gen.AlphaString())
	gens["Deleted"] = gen.PtrOf(gen.Bool())
	gens["DeletedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DenyEncryptionScopeOverride"] = gen.PtrOf(gen.Bool())
	gens["EnableNfsV3AllSquash"] = gen.PtrOf(gen.Bool())
	gens["EnableNfsV3RootSquash"] = gen.PtrOf(gen.Bool())
	gens["HasImmutabilityPolicy"] = gen.PtrOf(gen.Bool())
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
	gens["LastModifiedTime"] = gen.PtrOf(gen.AlphaString())
	gens["LeaseDuration"] = gen.PtrOf(gen.OneConstOf(ContainerProperties_LeaseDuration_STATUS_Fixed, ContainerProperties_LeaseDuration_STATUS_Infinite))
	gens["LeaseState"] = gen.PtrOf(gen.OneConstOf(
		ContainerProperties_LeaseState_STATUS_Available,
		ContainerProperties_LeaseState_STATUS_Breaking,
		ContainerProperties_LeaseState_STATUS_Broken,
		ContainerProperties_LeaseState_STATUS_Expired,
		ContainerProperties_LeaseState_STATUS_Leased))
	gens["LeaseStatus"] = gen.PtrOf(gen.OneConstOf(ContainerProperties_LeaseStatus_STATUS_Locked, ContainerProperties_LeaseStatus_STATUS_Unlocked))
	gens["Metadata"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["PublicAccess"] = gen.PtrOf(gen.OneConstOf(ContainerProperties_PublicAccess_STATUS_Blob, ContainerProperties_PublicAccess_STATUS_Container, ContainerProperties_PublicAccess_STATUS_None))
	gens["RemainingRetentionDays"] = gen.PtrOf(gen.Int())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForContainerProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForContainerProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ImmutabilityPolicy"] = gen.PtrOf(ImmutabilityPolicyProperties_STATUSGenerator())
	gens["ImmutableStorageWithVersioning"] = gen.PtrOf(ImmutableStorageWithVersioning_STATUSGenerator())
	gens["LegalHold"] = gen.PtrOf(LegalHoldProperties_STATUSGenerator())
}

func Test_ImmutabilityPolicyProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyProperties_STATUS, ImmutabilityPolicyProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyProperties_STATUS runs a test to see if a specific instance of ImmutabilityPolicyProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyProperties_STATUS(subject ImmutabilityPolicyProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperties_STATUS
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

// Generator of ImmutabilityPolicyProperties_STATUS instances for property testing - lazily instantiated by
// ImmutabilityPolicyProperties_STATUSGenerator()
var immutabilityPolicyProperties_STATUSGenerator gopter.Gen

// ImmutabilityPolicyProperties_STATUSGenerator returns a generator of ImmutabilityPolicyProperties_STATUS instances for property testing.
// We first initialize immutabilityPolicyProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ImmutabilityPolicyProperties_STATUSGenerator() gopter.Gen {
	if immutabilityPolicyProperties_STATUSGenerator != nil {
		return immutabilityPolicyProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUS(generators)
	immutabilityPolicyProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_STATUS(generators)
	immutabilityPolicyProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperties_STATUS{}), generators)

	return immutabilityPolicyProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForImmutabilityPolicyProperties_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ImmutabilityPolicyProperty_STATUSGenerator())
	gens["UpdateHistory"] = gen.SliceOf(UpdateHistoryProperty_STATUSGenerator())
}

func Test_ImmutabilityPolicyProperty_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilityPolicyProperty_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilityPolicyProperty_STATUS, ImmutabilityPolicyProperty_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilityPolicyProperty_STATUS runs a test to see if a specific instance of ImmutabilityPolicyProperty_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilityPolicyProperty_STATUS(subject ImmutabilityPolicyProperty_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilityPolicyProperty_STATUS
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

// Generator of ImmutabilityPolicyProperty_STATUS instances for property testing - lazily instantiated by
// ImmutabilityPolicyProperty_STATUSGenerator()
var immutabilityPolicyProperty_STATUSGenerator gopter.Gen

// ImmutabilityPolicyProperty_STATUSGenerator returns a generator of ImmutabilityPolicyProperty_STATUS instances for property testing.
func ImmutabilityPolicyProperty_STATUSGenerator() gopter.Gen {
	if immutabilityPolicyProperty_STATUSGenerator != nil {
		return immutabilityPolicyProperty_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilityPolicyProperty_STATUS(generators)
	immutabilityPolicyProperty_STATUSGenerator = gen.Struct(reflect.TypeOf(ImmutabilityPolicyProperty_STATUS{}), generators)

	return immutabilityPolicyProperty_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilityPolicyProperty_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilityPolicyProperty_STATUS(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWrites"] = gen.PtrOf(gen.Bool())
	gens["AllowProtectedAppendWritesAll"] = gen.PtrOf(gen.Bool())
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["State"] = gen.PtrOf(gen.OneConstOf(ImmutabilityPolicyProperty_State_STATUS_Locked, ImmutabilityPolicyProperty_State_STATUS_Unlocked))
}

func Test_ImmutableStorageWithVersioning_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutableStorageWithVersioning_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutableStorageWithVersioning_STATUS, ImmutableStorageWithVersioning_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutableStorageWithVersioning_STATUS runs a test to see if a specific instance of ImmutableStorageWithVersioning_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutableStorageWithVersioning_STATUS(subject ImmutableStorageWithVersioning_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutableStorageWithVersioning_STATUS
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

// Generator of ImmutableStorageWithVersioning_STATUS instances for property testing - lazily instantiated by
// ImmutableStorageWithVersioning_STATUSGenerator()
var immutableStorageWithVersioning_STATUSGenerator gopter.Gen

// ImmutableStorageWithVersioning_STATUSGenerator returns a generator of ImmutableStorageWithVersioning_STATUS instances for property testing.
func ImmutableStorageWithVersioning_STATUSGenerator() gopter.Gen {
	if immutableStorageWithVersioning_STATUSGenerator != nil {
		return immutableStorageWithVersioning_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_STATUS(generators)
	immutableStorageWithVersioning_STATUSGenerator = gen.Struct(reflect.TypeOf(ImmutableStorageWithVersioning_STATUS{}), generators)

	return immutableStorageWithVersioning_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutableStorageWithVersioning_STATUS(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["MigrationState"] = gen.PtrOf(gen.OneConstOf(ImmutableStorageWithVersioning_MigrationState_STATUS_Completed, ImmutableStorageWithVersioning_MigrationState_STATUS_InProgress))
	gens["TimeStamp"] = gen.PtrOf(gen.AlphaString())
}

func Test_LegalHoldProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LegalHoldProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLegalHoldProperties_STATUS, LegalHoldProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLegalHoldProperties_STATUS runs a test to see if a specific instance of LegalHoldProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForLegalHoldProperties_STATUS(subject LegalHoldProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LegalHoldProperties_STATUS
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

// Generator of LegalHoldProperties_STATUS instances for property testing - lazily instantiated by
// LegalHoldProperties_STATUSGenerator()
var legalHoldProperties_STATUSGenerator gopter.Gen

// LegalHoldProperties_STATUSGenerator returns a generator of LegalHoldProperties_STATUS instances for property testing.
// We first initialize legalHoldProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LegalHoldProperties_STATUSGenerator() gopter.Gen {
	if legalHoldProperties_STATUSGenerator != nil {
		return legalHoldProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUS(generators)
	legalHoldProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForLegalHoldProperties_STATUS(generators)
	legalHoldProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(LegalHoldProperties_STATUS{}), generators)

	return legalHoldProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLegalHoldProperties_STATUS(gens map[string]gopter.Gen) {
	gens["HasLegalHold"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForLegalHoldProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLegalHoldProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ProtectedAppendWritesHistory"] = gen.PtrOf(ProtectedAppendWritesHistory_STATUSGenerator())
	gens["Tags"] = gen.SliceOf(TagProperty_STATUSGenerator())
}

func Test_ProtectedAppendWritesHistory_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProtectedAppendWritesHistory_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProtectedAppendWritesHistory_STATUS, ProtectedAppendWritesHistory_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProtectedAppendWritesHistory_STATUS runs a test to see if a specific instance of ProtectedAppendWritesHistory_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProtectedAppendWritesHistory_STATUS(subject ProtectedAppendWritesHistory_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProtectedAppendWritesHistory_STATUS
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

// Generator of ProtectedAppendWritesHistory_STATUS instances for property testing - lazily instantiated by
// ProtectedAppendWritesHistory_STATUSGenerator()
var protectedAppendWritesHistory_STATUSGenerator gopter.Gen

// ProtectedAppendWritesHistory_STATUSGenerator returns a generator of ProtectedAppendWritesHistory_STATUS instances for property testing.
func ProtectedAppendWritesHistory_STATUSGenerator() gopter.Gen {
	if protectedAppendWritesHistory_STATUSGenerator != nil {
		return protectedAppendWritesHistory_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProtectedAppendWritesHistory_STATUS(generators)
	protectedAppendWritesHistory_STATUSGenerator = gen.Struct(reflect.TypeOf(ProtectedAppendWritesHistory_STATUS{}), generators)

	return protectedAppendWritesHistory_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForProtectedAppendWritesHistory_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProtectedAppendWritesHistory_STATUS(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWritesAll"] = gen.PtrOf(gen.Bool())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
}

func Test_StorageAccountsBlobServicesContainer_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsBlobServicesContainer_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsBlobServicesContainer_STATUS, StorageAccountsBlobServicesContainer_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsBlobServicesContainer_STATUS runs a test to see if a specific instance of StorageAccountsBlobServicesContainer_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsBlobServicesContainer_STATUS(subject StorageAccountsBlobServicesContainer_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsBlobServicesContainer_STATUS
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

// Generator of StorageAccountsBlobServicesContainer_STATUS instances for property testing - lazily instantiated by
// StorageAccountsBlobServicesContainer_STATUSGenerator()
var storageAccountsBlobServicesContainer_STATUSGenerator gopter.Gen

// StorageAccountsBlobServicesContainer_STATUSGenerator returns a generator of StorageAccountsBlobServicesContainer_STATUS instances for property testing.
// We first initialize storageAccountsBlobServicesContainer_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsBlobServicesContainer_STATUSGenerator() gopter.Gen {
	if storageAccountsBlobServicesContainer_STATUSGenerator != nil {
		return storageAccountsBlobServicesContainer_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainer_STATUS(generators)
	storageAccountsBlobServicesContainer_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainer_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainer_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer_STATUS(generators)
	storageAccountsBlobServicesContainer_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsBlobServicesContainer_STATUS{}), generators)

	return storageAccountsBlobServicesContainer_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainer_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsBlobServicesContainer_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsBlobServicesContainer_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ContainerProperties_STATUSGenerator())
}

func Test_TagProperty_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TagProperty_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTagProperty_STATUS, TagProperty_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTagProperty_STATUS runs a test to see if a specific instance of TagProperty_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForTagProperty_STATUS(subject TagProperty_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TagProperty_STATUS
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

// Generator of TagProperty_STATUS instances for property testing - lazily instantiated by TagProperty_STATUSGenerator()
var tagProperty_STATUSGenerator gopter.Gen

// TagProperty_STATUSGenerator returns a generator of TagProperty_STATUS instances for property testing.
func TagProperty_STATUSGenerator() gopter.Gen {
	if tagProperty_STATUSGenerator != nil {
		return tagProperty_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTagProperty_STATUS(generators)
	tagProperty_STATUSGenerator = gen.Struct(reflect.TypeOf(TagProperty_STATUS{}), generators)

	return tagProperty_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForTagProperty_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTagProperty_STATUS(gens map[string]gopter.Gen) {
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}

func Test_UpdateHistoryProperty_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UpdateHistoryProperty_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUpdateHistoryProperty_STATUS, UpdateHistoryProperty_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUpdateHistoryProperty_STATUS runs a test to see if a specific instance of UpdateHistoryProperty_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForUpdateHistoryProperty_STATUS(subject UpdateHistoryProperty_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UpdateHistoryProperty_STATUS
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

// Generator of UpdateHistoryProperty_STATUS instances for property testing - lazily instantiated by
// UpdateHistoryProperty_STATUSGenerator()
var updateHistoryProperty_STATUSGenerator gopter.Gen

// UpdateHistoryProperty_STATUSGenerator returns a generator of UpdateHistoryProperty_STATUS instances for property testing.
func UpdateHistoryProperty_STATUSGenerator() gopter.Gen {
	if updateHistoryProperty_STATUSGenerator != nil {
		return updateHistoryProperty_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUpdateHistoryProperty_STATUS(generators)
	updateHistoryProperty_STATUSGenerator = gen.Struct(reflect.TypeOf(UpdateHistoryProperty_STATUS{}), generators)

	return updateHistoryProperty_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForUpdateHistoryProperty_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUpdateHistoryProperty_STATUS(gens map[string]gopter.Gen) {
	gens["AllowProtectedAppendWrites"] = gen.PtrOf(gen.Bool())
	gens["AllowProtectedAppendWritesAll"] = gen.PtrOf(gen.Bool())
	gens["ImmutabilityPeriodSinceCreationInDays"] = gen.PtrOf(gen.Int())
	gens["ObjectIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Timestamp"] = gen.PtrOf(gen.AlphaString())
	gens["Update"] = gen.PtrOf(gen.OneConstOf(UpdateHistoryProperty_Update_STATUS_Extend, UpdateHistoryProperty_Update_STATUS_Lock, UpdateHistoryProperty_Update_STATUS_Put))
	gens["Upn"] = gen.PtrOf(gen.AlphaString())
}
