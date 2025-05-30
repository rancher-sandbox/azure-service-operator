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

func Test_CopyCompletionError_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CopyCompletionError via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCopyCompletionError, CopyCompletionErrorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCopyCompletionError runs a test to see if a specific instance of CopyCompletionError round trips to JSON and back losslessly
func RunJSONSerializationTestForCopyCompletionError(subject CopyCompletionError) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CopyCompletionError
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

// Generator of CopyCompletionError instances for property testing - lazily instantiated by
// CopyCompletionErrorGenerator()
var copyCompletionErrorGenerator gopter.Gen

// CopyCompletionErrorGenerator returns a generator of CopyCompletionError instances for property testing.
func CopyCompletionErrorGenerator() gopter.Gen {
	if copyCompletionErrorGenerator != nil {
		return copyCompletionErrorGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCopyCompletionError(generators)
	copyCompletionErrorGenerator = gen.Struct(reflect.TypeOf(CopyCompletionError{}), generators)

	return copyCompletionErrorGenerator
}

// AddIndependentPropertyGeneratorsForCopyCompletionError is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCopyCompletionError(gens map[string]gopter.Gen) {
	gens["ErrorCode"] = gen.PtrOf(gen.AlphaString())
	gens["ErrorMessage"] = gen.PtrOf(gen.AlphaString())
}

func Test_CopyCompletionError_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CopyCompletionError_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCopyCompletionError_STATUS, CopyCompletionError_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCopyCompletionError_STATUS runs a test to see if a specific instance of CopyCompletionError_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCopyCompletionError_STATUS(subject CopyCompletionError_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CopyCompletionError_STATUS
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

// Generator of CopyCompletionError_STATUS instances for property testing - lazily instantiated by
// CopyCompletionError_STATUSGenerator()
var copyCompletionError_STATUSGenerator gopter.Gen

// CopyCompletionError_STATUSGenerator returns a generator of CopyCompletionError_STATUS instances for property testing.
func CopyCompletionError_STATUSGenerator() gopter.Gen {
	if copyCompletionError_STATUSGenerator != nil {
		return copyCompletionError_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCopyCompletionError_STATUS(generators)
	copyCompletionError_STATUSGenerator = gen.Struct(reflect.TypeOf(CopyCompletionError_STATUS{}), generators)

	return copyCompletionError_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCopyCompletionError_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCopyCompletionError_STATUS(gens map[string]gopter.Gen) {
	gens["ErrorCode"] = gen.PtrOf(gen.AlphaString())
	gens["ErrorMessage"] = gen.PtrOf(gen.AlphaString())
}

func Test_Snapshot_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshot, SnapshotGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshot runs a test to see if a specific instance of Snapshot round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshot(subject Snapshot) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot
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

// Generator of Snapshot instances for property testing - lazily instantiated by SnapshotGenerator()
var snapshotGenerator gopter.Gen

// SnapshotGenerator returns a generator of Snapshot instances for property testing.
func SnapshotGenerator() gopter.Gen {
	if snapshotGenerator != nil {
		return snapshotGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSnapshot(generators)
	snapshotGenerator = gen.Struct(reflect.TypeOf(Snapshot{}), generators)

	return snapshotGenerator
}

// AddRelatedPropertyGeneratorsForSnapshot is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshot(gens map[string]gopter.Gen) {
	gens["Spec"] = Snapshot_SpecGenerator()
	gens["Status"] = Snapshot_STATUSGenerator()
}

func Test_SnapshotOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotOperatorSpec, SnapshotOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotOperatorSpec runs a test to see if a specific instance of SnapshotOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotOperatorSpec(subject SnapshotOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotOperatorSpec
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

// Generator of SnapshotOperatorSpec instances for property testing - lazily instantiated by
// SnapshotOperatorSpecGenerator()
var snapshotOperatorSpecGenerator gopter.Gen

// SnapshotOperatorSpecGenerator returns a generator of SnapshotOperatorSpec instances for property testing.
func SnapshotOperatorSpecGenerator() gopter.Gen {
	if snapshotOperatorSpecGenerator != nil {
		return snapshotOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	snapshotOperatorSpecGenerator = gen.Struct(reflect.TypeOf(SnapshotOperatorSpec{}), generators)

	return snapshotOperatorSpecGenerator
}

func Test_SnapshotSku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSku, SnapshotSkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSku runs a test to see if a specific instance of SnapshotSku round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSku(subject SnapshotSku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku
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

// Generator of SnapshotSku instances for property testing - lazily instantiated by SnapshotSkuGenerator()
var snapshotSkuGenerator gopter.Gen

// SnapshotSkuGenerator returns a generator of SnapshotSku instances for property testing.
func SnapshotSkuGenerator() gopter.Gen {
	if snapshotSkuGenerator != nil {
		return snapshotSkuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSku(generators)
	snapshotSkuGenerator = gen.Struct(reflect.TypeOf(SnapshotSku{}), generators)

	return snapshotSkuGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_SnapshotSku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSku_STATUS, SnapshotSku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSku_STATUS runs a test to see if a specific instance of SnapshotSku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSku_STATUS(subject SnapshotSku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku_STATUS
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

// Generator of SnapshotSku_STATUS instances for property testing - lazily instantiated by SnapshotSku_STATUSGenerator()
var snapshotSku_STATUSGenerator gopter.Gen

// SnapshotSku_STATUSGenerator returns a generator of SnapshotSku_STATUS instances for property testing.
func SnapshotSku_STATUSGenerator() gopter.Gen {
	if snapshotSku_STATUSGenerator != nil {
		return snapshotSku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSku_STATUS(generators)
	snapshotSku_STATUSGenerator = gen.Struct(reflect.TypeOf(SnapshotSku_STATUS{}), generators)

	return snapshotSku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSku_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_Snapshot_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshot_STATUS, Snapshot_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshot_STATUS runs a test to see if a specific instance of Snapshot_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshot_STATUS(subject Snapshot_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot_STATUS
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

// Generator of Snapshot_STATUS instances for property testing - lazily instantiated by Snapshot_STATUSGenerator()
var snapshot_STATUSGenerator gopter.Gen

// Snapshot_STATUSGenerator returns a generator of Snapshot_STATUS instances for property testing.
// We first initialize snapshot_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Snapshot_STATUSGenerator() gopter.Gen {
	if snapshot_STATUSGenerator != nil {
		return snapshot_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_STATUS(generators)
	snapshot_STATUSGenerator = gen.Struct(reflect.TypeOf(Snapshot_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_STATUS(generators)
	AddRelatedPropertyGeneratorsForSnapshot_STATUS(generators)
	snapshot_STATUSGenerator = gen.Struct(reflect.TypeOf(Snapshot_STATUS{}), generators)

	return snapshot_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSnapshot_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshot_STATUS(gens map[string]gopter.Gen) {
	gens["CompletionPercent"] = gen.PtrOf(gen.Float64())
	gens["DataAccessAuthMode"] = gen.PtrOf(gen.AlphaString())
	gens["DiskAccessId"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeBytes"] = gen.PtrOf(gen.Int())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.AlphaString())
	gens["HyperVGeneration"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["IncrementalSnapshotFamilyId"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ManagedBy"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["OsType"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["SupportsHibernation"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["TimeCreated"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UniqueId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshot_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshot_STATUS(gens map[string]gopter.Gen) {
	gens["CopyCompletionError"] = gen.PtrOf(CopyCompletionError_STATUSGenerator())
	gens["CreationData"] = gen.PtrOf(CreationData_STATUSGenerator())
	gens["Encryption"] = gen.PtrOf(Encryption_STATUSGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollection_STATUSGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlan_STATUSGenerator())
	gens["SecurityProfile"] = gen.PtrOf(DiskSecurityProfile_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSku_STATUSGenerator())
	gens["SupportedCapabilities"] = gen.PtrOf(SupportedCapabilities_STATUSGenerator())
}

func Test_Snapshot_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshot_Spec, Snapshot_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshot_Spec runs a test to see if a specific instance of Snapshot_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshot_Spec(subject Snapshot_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot_Spec
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

// Generator of Snapshot_Spec instances for property testing - lazily instantiated by Snapshot_SpecGenerator()
var snapshot_SpecGenerator gopter.Gen

// Snapshot_SpecGenerator returns a generator of Snapshot_Spec instances for property testing.
// We first initialize snapshot_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Snapshot_SpecGenerator() gopter.Gen {
	if snapshot_SpecGenerator != nil {
		return snapshot_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_Spec(generators)
	snapshot_SpecGenerator = gen.Struct(reflect.TypeOf(Snapshot_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_Spec(generators)
	AddRelatedPropertyGeneratorsForSnapshot_Spec(generators)
	snapshot_SpecGenerator = gen.Struct(reflect.TypeOf(Snapshot_Spec{}), generators)

	return snapshot_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSnapshot_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshot_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["CompletionPercent"] = gen.PtrOf(gen.Float64())
	gens["DataAccessAuthMode"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.AlphaString())
	gens["HyperVGeneration"] = gen.PtrOf(gen.AlphaString())
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["OsType"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["SupportsHibernation"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshot_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshot_Spec(gens map[string]gopter.Gen) {
	gens["CopyCompletionError"] = gen.PtrOf(CopyCompletionErrorGenerator())
	gens["CreationData"] = gen.PtrOf(CreationDataGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollectionGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
	gens["OperatorSpec"] = gen.PtrOf(SnapshotOperatorSpecGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlanGenerator())
	gens["SecurityProfile"] = gen.PtrOf(DiskSecurityProfileGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSkuGenerator())
	gens["SupportedCapabilities"] = gen.PtrOf(SupportedCapabilitiesGenerator())
}
