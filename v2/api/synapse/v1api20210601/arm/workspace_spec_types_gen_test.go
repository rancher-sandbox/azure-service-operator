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

func Test_CspWorkspaceAdminProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CspWorkspaceAdminProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCspWorkspaceAdminProperties, CspWorkspaceAdminPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCspWorkspaceAdminProperties runs a test to see if a specific instance of CspWorkspaceAdminProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForCspWorkspaceAdminProperties(subject CspWorkspaceAdminProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CspWorkspaceAdminProperties
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

// Generator of CspWorkspaceAdminProperties instances for property testing - lazily instantiated by
// CspWorkspaceAdminPropertiesGenerator()
var cspWorkspaceAdminPropertiesGenerator gopter.Gen

// CspWorkspaceAdminPropertiesGenerator returns a generator of CspWorkspaceAdminProperties instances for property testing.
func CspWorkspaceAdminPropertiesGenerator() gopter.Gen {
	if cspWorkspaceAdminPropertiesGenerator != nil {
		return cspWorkspaceAdminPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCspWorkspaceAdminProperties(generators)
	cspWorkspaceAdminPropertiesGenerator = gen.Struct(reflect.TypeOf(CspWorkspaceAdminProperties{}), generators)

	return cspWorkspaceAdminPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForCspWorkspaceAdminProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCspWorkspaceAdminProperties(gens map[string]gopter.Gen) {
	gens["InitialWorkspaceAdminObjectId"] = gen.PtrOf(gen.AlphaString())
}

func Test_CustomerManagedKeyDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CustomerManagedKeyDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCustomerManagedKeyDetails, CustomerManagedKeyDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCustomerManagedKeyDetails runs a test to see if a specific instance of CustomerManagedKeyDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForCustomerManagedKeyDetails(subject CustomerManagedKeyDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CustomerManagedKeyDetails
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

// Generator of CustomerManagedKeyDetails instances for property testing - lazily instantiated by
// CustomerManagedKeyDetailsGenerator()
var customerManagedKeyDetailsGenerator gopter.Gen

// CustomerManagedKeyDetailsGenerator returns a generator of CustomerManagedKeyDetails instances for property testing.
func CustomerManagedKeyDetailsGenerator() gopter.Gen {
	if customerManagedKeyDetailsGenerator != nil {
		return customerManagedKeyDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForCustomerManagedKeyDetails(generators)
	customerManagedKeyDetailsGenerator = gen.Struct(reflect.TypeOf(CustomerManagedKeyDetails{}), generators)

	return customerManagedKeyDetailsGenerator
}

// AddRelatedPropertyGeneratorsForCustomerManagedKeyDetails is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCustomerManagedKeyDetails(gens map[string]gopter.Gen) {
	gens["KekIdentity"] = gen.PtrOf(KekIdentityPropertiesGenerator())
	gens["Key"] = gen.PtrOf(WorkspaceKeyDetailsGenerator())
}

func Test_DataLakeStorageAccountDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DataLakeStorageAccountDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDataLakeStorageAccountDetails, DataLakeStorageAccountDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDataLakeStorageAccountDetails runs a test to see if a specific instance of DataLakeStorageAccountDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForDataLakeStorageAccountDetails(subject DataLakeStorageAccountDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DataLakeStorageAccountDetails
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

// Generator of DataLakeStorageAccountDetails instances for property testing - lazily instantiated by
// DataLakeStorageAccountDetailsGenerator()
var dataLakeStorageAccountDetailsGenerator gopter.Gen

// DataLakeStorageAccountDetailsGenerator returns a generator of DataLakeStorageAccountDetails instances for property testing.
func DataLakeStorageAccountDetailsGenerator() gopter.Gen {
	if dataLakeStorageAccountDetailsGenerator != nil {
		return dataLakeStorageAccountDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDataLakeStorageAccountDetails(generators)
	dataLakeStorageAccountDetailsGenerator = gen.Struct(reflect.TypeOf(DataLakeStorageAccountDetails{}), generators)

	return dataLakeStorageAccountDetailsGenerator
}

// AddIndependentPropertyGeneratorsForDataLakeStorageAccountDetails is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDataLakeStorageAccountDetails(gens map[string]gopter.Gen) {
	gens["AccountUrl"] = gen.PtrOf(gen.AlphaString())
	gens["CreateManagedPrivateEndpoint"] = gen.PtrOf(gen.Bool())
	gens["Filesystem"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_EncryptionDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionDetails, EncryptionDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionDetails runs a test to see if a specific instance of EncryptionDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionDetails(subject EncryptionDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionDetails
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

// Generator of EncryptionDetails instances for property testing - lazily instantiated by EncryptionDetailsGenerator()
var encryptionDetailsGenerator gopter.Gen

// EncryptionDetailsGenerator returns a generator of EncryptionDetails instances for property testing.
func EncryptionDetailsGenerator() gopter.Gen {
	if encryptionDetailsGenerator != nil {
		return encryptionDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForEncryptionDetails(generators)
	encryptionDetailsGenerator = gen.Struct(reflect.TypeOf(EncryptionDetails{}), generators)

	return encryptionDetailsGenerator
}

// AddRelatedPropertyGeneratorsForEncryptionDetails is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionDetails(gens map[string]gopter.Gen) {
	gens["Cmk"] = gen.PtrOf(CustomerManagedKeyDetailsGenerator())
}

func Test_KekIdentityProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KekIdentityProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKekIdentityProperties, KekIdentityPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKekIdentityProperties runs a test to see if a specific instance of KekIdentityProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForKekIdentityProperties(subject KekIdentityProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KekIdentityProperties
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

// Generator of KekIdentityProperties instances for property testing - lazily instantiated by
// KekIdentityPropertiesGenerator()
var kekIdentityPropertiesGenerator gopter.Gen

// KekIdentityPropertiesGenerator returns a generator of KekIdentityProperties instances for property testing.
func KekIdentityPropertiesGenerator() gopter.Gen {
	if kekIdentityPropertiesGenerator != nil {
		return kekIdentityPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKekIdentityProperties(generators)
	kekIdentityPropertiesGenerator = gen.Struct(reflect.TypeOf(KekIdentityProperties{}), generators)

	return kekIdentityPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForKekIdentityProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKekIdentityProperties(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}

func Test_ManagedIdentity_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedIdentity via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedIdentity, ManagedIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedIdentity runs a test to see if a specific instance of ManagedIdentity round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedIdentity(subject ManagedIdentity) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedIdentity
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

// Generator of ManagedIdentity instances for property testing - lazily instantiated by ManagedIdentityGenerator()
var managedIdentityGenerator gopter.Gen

// ManagedIdentityGenerator returns a generator of ManagedIdentity instances for property testing.
// We first initialize managedIdentityGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedIdentityGenerator() gopter.Gen {
	if managedIdentityGenerator != nil {
		return managedIdentityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedIdentity(generators)
	managedIdentityGenerator = gen.Struct(reflect.TypeOf(ManagedIdentity{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedIdentity(generators)
	AddRelatedPropertyGeneratorsForManagedIdentity(generators)
	managedIdentityGenerator = gen.Struct(reflect.TypeOf(ManagedIdentity{}), generators)

	return managedIdentityGenerator
}

// AddIndependentPropertyGeneratorsForManagedIdentity is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedIdentity(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(ManagedIdentity_Type_None, ManagedIdentity_Type_SystemAssigned, ManagedIdentity_Type_SystemAssignedUserAssigned))
}

// AddRelatedPropertyGeneratorsForManagedIdentity is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedIdentity(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserAssignedIdentityDetailsGenerator())
}

func Test_ManagedVirtualNetworkSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedVirtualNetworkSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedVirtualNetworkSettings, ManagedVirtualNetworkSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedVirtualNetworkSettings runs a test to see if a specific instance of ManagedVirtualNetworkSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedVirtualNetworkSettings(subject ManagedVirtualNetworkSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedVirtualNetworkSettings
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

// Generator of ManagedVirtualNetworkSettings instances for property testing - lazily instantiated by
// ManagedVirtualNetworkSettingsGenerator()
var managedVirtualNetworkSettingsGenerator gopter.Gen

// ManagedVirtualNetworkSettingsGenerator returns a generator of ManagedVirtualNetworkSettings instances for property testing.
func ManagedVirtualNetworkSettingsGenerator() gopter.Gen {
	if managedVirtualNetworkSettingsGenerator != nil {
		return managedVirtualNetworkSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedVirtualNetworkSettings(generators)
	managedVirtualNetworkSettingsGenerator = gen.Struct(reflect.TypeOf(ManagedVirtualNetworkSettings{}), generators)

	return managedVirtualNetworkSettingsGenerator
}

// AddIndependentPropertyGeneratorsForManagedVirtualNetworkSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedVirtualNetworkSettings(gens map[string]gopter.Gen) {
	gens["AllowedAadTenantIdsForLinking"] = gen.SliceOf(gen.AlphaString())
	gens["LinkedAccessCheckOnTargetResource"] = gen.PtrOf(gen.Bool())
	gens["PreventDataExfiltration"] = gen.PtrOf(gen.Bool())
}

func Test_PurviewConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PurviewConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPurviewConfiguration, PurviewConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPurviewConfiguration runs a test to see if a specific instance of PurviewConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForPurviewConfiguration(subject PurviewConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PurviewConfiguration
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

// Generator of PurviewConfiguration instances for property testing - lazily instantiated by
// PurviewConfigurationGenerator()
var purviewConfigurationGenerator gopter.Gen

// PurviewConfigurationGenerator returns a generator of PurviewConfiguration instances for property testing.
func PurviewConfigurationGenerator() gopter.Gen {
	if purviewConfigurationGenerator != nil {
		return purviewConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPurviewConfiguration(generators)
	purviewConfigurationGenerator = gen.Struct(reflect.TypeOf(PurviewConfiguration{}), generators)

	return purviewConfigurationGenerator
}

// AddIndependentPropertyGeneratorsForPurviewConfiguration is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPurviewConfiguration(gens map[string]gopter.Gen) {
	gens["PurviewResourceId"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentityDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityDetails, UserAssignedIdentityDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityDetails runs a test to see if a specific instance of UserAssignedIdentityDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityDetails(subject UserAssignedIdentityDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityDetails
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

// Generator of UserAssignedIdentityDetails instances for property testing - lazily instantiated by
// UserAssignedIdentityDetailsGenerator()
var userAssignedIdentityDetailsGenerator gopter.Gen

// UserAssignedIdentityDetailsGenerator returns a generator of UserAssignedIdentityDetails instances for property testing.
func UserAssignedIdentityDetailsGenerator() gopter.Gen {
	if userAssignedIdentityDetailsGenerator != nil {
		return userAssignedIdentityDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	userAssignedIdentityDetailsGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityDetails{}), generators)

	return userAssignedIdentityDetailsGenerator
}

func Test_VirtualNetworkProfile_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkProfile via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkProfile, VirtualNetworkProfileGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkProfile runs a test to see if a specific instance of VirtualNetworkProfile round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkProfile(subject VirtualNetworkProfile) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkProfile
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

// Generator of VirtualNetworkProfile instances for property testing - lazily instantiated by
// VirtualNetworkProfileGenerator()
var virtualNetworkProfileGenerator gopter.Gen

// VirtualNetworkProfileGenerator returns a generator of VirtualNetworkProfile instances for property testing.
func VirtualNetworkProfileGenerator() gopter.Gen {
	if virtualNetworkProfileGenerator != nil {
		return virtualNetworkProfileGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkProfile(generators)
	virtualNetworkProfileGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkProfile{}), generators)

	return virtualNetworkProfileGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkProfile is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkProfile(gens map[string]gopter.Gen) {
	gens["ComputeSubnetId"] = gen.PtrOf(gen.AlphaString())
}

func Test_WorkspaceKeyDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceKeyDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceKeyDetails, WorkspaceKeyDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceKeyDetails runs a test to see if a specific instance of WorkspaceKeyDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceKeyDetails(subject WorkspaceKeyDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceKeyDetails
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

// Generator of WorkspaceKeyDetails instances for property testing - lazily instantiated by
// WorkspaceKeyDetailsGenerator()
var workspaceKeyDetailsGenerator gopter.Gen

// WorkspaceKeyDetailsGenerator returns a generator of WorkspaceKeyDetails instances for property testing.
func WorkspaceKeyDetailsGenerator() gopter.Gen {
	if workspaceKeyDetailsGenerator != nil {
		return workspaceKeyDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceKeyDetails(generators)
	workspaceKeyDetailsGenerator = gen.Struct(reflect.TypeOf(WorkspaceKeyDetails{}), generators)

	return workspaceKeyDetailsGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceKeyDetails is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceKeyDetails(gens map[string]gopter.Gen) {
	gens["KeyVaultUrl"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_WorkspaceProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceProperties, WorkspacePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceProperties runs a test to see if a specific instance of WorkspaceProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceProperties(subject WorkspaceProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceProperties
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

// Generator of WorkspaceProperties instances for property testing - lazily instantiated by
// WorkspacePropertiesGenerator()
var workspacePropertiesGenerator gopter.Gen

// WorkspacePropertiesGenerator returns a generator of WorkspaceProperties instances for property testing.
// We first initialize workspacePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspacePropertiesGenerator() gopter.Gen {
	if workspacePropertiesGenerator != nil {
		return workspacePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties(generators)
	workspacePropertiesGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties(generators)
	AddRelatedPropertyGeneratorsForWorkspaceProperties(generators)
	workspacePropertiesGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties{}), generators)

	return workspacePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceProperties(gens map[string]gopter.Gen) {
	gens["AzureADOnlyAuthentication"] = gen.PtrOf(gen.Bool())
	gens["ManagedResourceGroupName"] = gen.PtrOf(gen.AlphaString())
	gens["ManagedVirtualNetwork"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(WorkspaceProperties_PublicNetworkAccess_Disabled, WorkspaceProperties_PublicNetworkAccess_Enabled))
	gens["SqlAdministratorLogin"] = gen.PtrOf(gen.AlphaString())
	gens["SqlAdministratorLoginPassword"] = gen.PtrOf(gen.AlphaString())
	gens["TrustedServiceBypassEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForWorkspaceProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceProperties(gens map[string]gopter.Gen) {
	gens["CspWorkspaceAdminProperties"] = gen.PtrOf(CspWorkspaceAdminPropertiesGenerator())
	gens["DefaultDataLakeStorage"] = gen.PtrOf(DataLakeStorageAccountDetailsGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionDetailsGenerator())
	gens["ManagedVirtualNetworkSettings"] = gen.PtrOf(ManagedVirtualNetworkSettingsGenerator())
	gens["PurviewConfiguration"] = gen.PtrOf(PurviewConfigurationGenerator())
	gens["VirtualNetworkProfile"] = gen.PtrOf(VirtualNetworkProfileGenerator())
	gens["WorkspaceRepositoryConfiguration"] = gen.PtrOf(WorkspaceRepositoryConfigurationGenerator())
}

func Test_WorkspaceRepositoryConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceRepositoryConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceRepositoryConfiguration, WorkspaceRepositoryConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceRepositoryConfiguration runs a test to see if a specific instance of WorkspaceRepositoryConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceRepositoryConfiguration(subject WorkspaceRepositoryConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceRepositoryConfiguration
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

// Generator of WorkspaceRepositoryConfiguration instances for property testing - lazily instantiated by
// WorkspaceRepositoryConfigurationGenerator()
var workspaceRepositoryConfigurationGenerator gopter.Gen

// WorkspaceRepositoryConfigurationGenerator returns a generator of WorkspaceRepositoryConfiguration instances for property testing.
func WorkspaceRepositoryConfigurationGenerator() gopter.Gen {
	if workspaceRepositoryConfigurationGenerator != nil {
		return workspaceRepositoryConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceRepositoryConfiguration(generators)
	workspaceRepositoryConfigurationGenerator = gen.Struct(reflect.TypeOf(WorkspaceRepositoryConfiguration{}), generators)

	return workspaceRepositoryConfigurationGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceRepositoryConfiguration is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceRepositoryConfiguration(gens map[string]gopter.Gen) {
	gens["AccountName"] = gen.PtrOf(gen.AlphaString())
	gens["CollaborationBranch"] = gen.PtrOf(gen.AlphaString())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["LastCommitId"] = gen.PtrOf(gen.AlphaString())
	gens["ProjectName"] = gen.PtrOf(gen.AlphaString())
	gens["RepositoryName"] = gen.PtrOf(gen.AlphaString())
	gens["RootFolder"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_Workspace_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspace_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspace_Spec, Workspace_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspace_Spec runs a test to see if a specific instance of Workspace_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspace_Spec(subject Workspace_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspace_Spec
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

// Generator of Workspace_Spec instances for property testing - lazily instantiated by Workspace_SpecGenerator()
var workspace_SpecGenerator gopter.Gen

// Workspace_SpecGenerator returns a generator of Workspace_Spec instances for property testing.
// We first initialize workspace_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Workspace_SpecGenerator() gopter.Gen {
	if workspace_SpecGenerator != nil {
		return workspace_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_Spec(generators)
	workspace_SpecGenerator = gen.Struct(reflect.TypeOf(Workspace_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_Spec(generators)
	AddRelatedPropertyGeneratorsForWorkspace_Spec(generators)
	workspace_SpecGenerator = gen.Struct(reflect.TypeOf(Workspace_Spec{}), generators)

	return workspace_SpecGenerator
}

// AddIndependentPropertyGeneratorsForWorkspace_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspace_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspace_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspace_Spec(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(ManagedIdentityGenerator())
	gens["Properties"] = gen.PtrOf(WorkspacePropertiesGenerator())
}
