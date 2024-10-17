// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

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

func Test_PrivateDnsZonesCNAMERecord_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesCNAMERecord_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_Spec_ARM, PrivateDnsZonesCNAMERecord_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_Spec_ARM runs a test to see if a specific instance of PrivateDnsZonesCNAMERecord_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_Spec_ARM(subject PrivateDnsZonesCNAMERecord_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesCNAMERecord_Spec_ARM
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

// Generator of PrivateDnsZonesCNAMERecord_Spec_ARM instances for property testing - lazily instantiated by
// PrivateDnsZonesCNAMERecord_Spec_ARMGenerator()
var privateDnsZonesCNAMERecord_Spec_ARMGenerator gopter.Gen

// PrivateDnsZonesCNAMERecord_Spec_ARMGenerator returns a generator of PrivateDnsZonesCNAMERecord_Spec_ARM instances for property testing.
// We first initialize privateDnsZonesCNAMERecord_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesCNAMERecord_Spec_ARMGenerator() gopter.Gen {
	if privateDnsZonesCNAMERecord_Spec_ARMGenerator != nil {
		return privateDnsZonesCNAMERecord_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec_ARM(generators)
	privateDnsZonesCNAMERecord_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecord_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec_ARM(generators)
	privateDnsZonesCNAMERecord_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecord_Spec_ARM{}), generators)

	return privateDnsZonesCNAMERecord_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetProperties_ARMGenerator())
}