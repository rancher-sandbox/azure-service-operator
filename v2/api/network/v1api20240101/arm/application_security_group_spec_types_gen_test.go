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

func Test_ApplicationSecurityGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroup_Spec, ApplicationSecurityGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroup_Spec runs a test to see if a specific instance of ApplicationSecurityGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroup_Spec(subject ApplicationSecurityGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_Spec
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

// Generator of ApplicationSecurityGroup_Spec instances for property testing - lazily instantiated by
// ApplicationSecurityGroup_SpecGenerator()
var applicationSecurityGroup_SpecGenerator gopter.Gen

// ApplicationSecurityGroup_SpecGenerator returns a generator of ApplicationSecurityGroup_Spec instances for property testing.
func ApplicationSecurityGroup_SpecGenerator() gopter.Gen {
	if applicationSecurityGroup_SpecGenerator != nil {
		return applicationSecurityGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroup_Spec(generators)
	applicationSecurityGroup_SpecGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_Spec{}), generators)

	return applicationSecurityGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroup_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}
