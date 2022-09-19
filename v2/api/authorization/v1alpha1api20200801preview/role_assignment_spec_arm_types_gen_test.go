// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

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

func Test_RoleAssignment_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignment_Spec_ARM, RoleAssignment_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignment_Spec_ARM runs a test to see if a specific instance of RoleAssignment_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignment_Spec_ARM(subject RoleAssignment_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment_Spec_ARM
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

// Generator of RoleAssignment_Spec_ARM instances for property testing - lazily instantiated by
// RoleAssignment_Spec_ARMGenerator()
var roleAssignment_Spec_ARMGenerator gopter.Gen

// RoleAssignment_Spec_ARMGenerator returns a generator of RoleAssignment_Spec_ARM instances for property testing.
// We first initialize roleAssignment_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RoleAssignment_Spec_ARMGenerator() gopter.Gen {
	if roleAssignment_Spec_ARMGenerator != nil {
		return roleAssignment_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_Spec_ARM(generators)
	roleAssignment_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignment_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForRoleAssignment_Spec_ARM(generators)
	roleAssignment_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_Spec_ARM{}), generators)

	return roleAssignment_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignment_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignment_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRoleAssignment_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleAssignment_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoleAssignmentProperties_ARMGenerator())
}

func Test_RoleAssignmentProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignmentProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentProperties_ARM, RoleAssignmentProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentProperties_ARM runs a test to see if a specific instance of RoleAssignmentProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentProperties_ARM(subject RoleAssignmentProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignmentProperties_ARM
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

// Generator of RoleAssignmentProperties_ARM instances for property testing - lazily instantiated by
// RoleAssignmentProperties_ARMGenerator()
var roleAssignmentProperties_ARMGenerator gopter.Gen

// RoleAssignmentProperties_ARMGenerator returns a generator of RoleAssignmentProperties_ARM instances for property testing.
func RoleAssignmentProperties_ARMGenerator() gopter.Gen {
	if roleAssignmentProperties_ARMGenerator != nil {
		return roleAssignmentProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentProperties_ARM(generators)
	roleAssignmentProperties_ARMGenerator = gen.Struct(reflect.TypeOf(RoleAssignmentProperties_ARM{}), generators)

	return roleAssignmentProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentProperties_ARM(gens map[string]gopter.Gen) {
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.OneConstOf(
		RoleAssignmentProperties_PrincipalType_ForeignGroup,
		RoleAssignmentProperties_PrincipalType_Group,
		RoleAssignmentProperties_PrincipalType_ServicePrincipal,
		RoleAssignmentProperties_PrincipalType_User))
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
}