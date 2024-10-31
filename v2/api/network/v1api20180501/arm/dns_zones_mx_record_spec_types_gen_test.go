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

func Test_DnsZonesMXRecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesMXRecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesMXRecord_Spec, DnsZonesMXRecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesMXRecord_Spec runs a test to see if a specific instance of DnsZonesMXRecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesMXRecord_Spec(subject DnsZonesMXRecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesMXRecord_Spec
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

// Generator of DnsZonesMXRecord_Spec instances for property testing - lazily instantiated by
// DnsZonesMXRecord_SpecGenerator()
var dnsZonesMXRecord_SpecGenerator gopter.Gen

// DnsZonesMXRecord_SpecGenerator returns a generator of DnsZonesMXRecord_Spec instances for property testing.
// We first initialize dnsZonesMXRecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesMXRecord_SpecGenerator() gopter.Gen {
	if dnsZonesMXRecord_SpecGenerator != nil {
		return dnsZonesMXRecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesMXRecord_Spec(generators)
	dnsZonesMXRecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesMXRecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesMXRecord_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZonesMXRecord_Spec(generators)
	dnsZonesMXRecord_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZonesMXRecord_Spec{}), generators)

	return dnsZonesMXRecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesMXRecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesMXRecord_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForDnsZonesMXRecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesMXRecord_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetPropertiesGenerator())
}