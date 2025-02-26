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

func Test_UserAssignedIdentity_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity_Spec, UserAssignedIdentity_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity_Spec runs a test to see if a specific instance of UserAssignedIdentity_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity_Spec(subject UserAssignedIdentity_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity_Spec
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

// Generator of UserAssignedIdentity_Spec instances for property testing - lazily instantiated by
// UserAssignedIdentity_SpecGenerator()
var userAssignedIdentity_SpecGenerator gopter.Gen

// UserAssignedIdentity_SpecGenerator returns a generator of UserAssignedIdentity_Spec instances for property testing.
func UserAssignedIdentity_SpecGenerator() gopter.Gen {
	if userAssignedIdentity_SpecGenerator != nil {
		return userAssignedIdentity_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec(generators)
	userAssignedIdentity_SpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity_Spec{}), generators)

	return userAssignedIdentity_SpecGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentity_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}
