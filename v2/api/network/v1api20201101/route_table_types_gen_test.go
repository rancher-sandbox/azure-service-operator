// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101

import (
	"encoding/json"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101/storage"
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

func Test_RouteTable_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTable to hub returns original",
		prop.ForAll(RunResourceConversionTestForRouteTable, RouteTableGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRouteTable tests if a specific instance of RouteTable round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRouteTable(subject RouteTable) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20201101s.RouteTable
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RouteTable
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

func Test_RouteTable_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTable to RouteTable via AssignProperties_To_RouteTable & AssignProperties_From_RouteTable returns original",
		prop.ForAll(RunPropertyAssignmentTestForRouteTable, RouteTableGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRouteTable tests if a specific instance of RouteTable can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRouteTable(subject RouteTable) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.RouteTable
	err := copied.AssignProperties_To_RouteTable(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RouteTable
	err = actual.AssignProperties_From_RouteTable(&other)
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

func Test_RouteTable_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable, RouteTableGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable runs a test to see if a specific instance of RouteTable round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable(subject RouteTable) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable
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

// Generator of RouteTable instances for property testing - lazily instantiated by RouteTableGenerator()
var routeTableGenerator gopter.Gen

// RouteTableGenerator returns a generator of RouteTable instances for property testing.
func RouteTableGenerator() gopter.Gen {
	if routeTableGenerator != nil {
		return routeTableGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRouteTable(generators)
	routeTableGenerator = gen.Struct(reflect.TypeOf(RouteTable{}), generators)

	return routeTableGenerator
}

// AddRelatedPropertyGeneratorsForRouteTable is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTable(gens map[string]gopter.Gen) {
	gens["Spec"] = RouteTable_SpecGenerator()
	gens["Status"] = RouteTable_STATUSGenerator()
}

func Test_RouteTable_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTable_Spec to RouteTable_Spec via AssignProperties_To_RouteTable_Spec & AssignProperties_From_RouteTable_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRouteTable_Spec, RouteTable_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRouteTable_Spec tests if a specific instance of RouteTable_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRouteTable_Spec(subject RouteTable_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.RouteTable_Spec
	err := copied.AssignProperties_To_RouteTable_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RouteTable_Spec
	err = actual.AssignProperties_From_RouteTable_Spec(&other)
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

func Test_RouteTable_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable_Spec, RouteTable_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable_Spec runs a test to see if a specific instance of RouteTable_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable_Spec(subject RouteTable_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable_Spec
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

// Generator of RouteTable_Spec instances for property testing - lazily instantiated by RouteTable_SpecGenerator()
var routeTable_SpecGenerator gopter.Gen

// RouteTable_SpecGenerator returns a generator of RouteTable_Spec instances for property testing.
func RouteTable_SpecGenerator() gopter.Gen {
	if routeTable_SpecGenerator != nil {
		return routeTable_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_Spec(generators)
	routeTable_SpecGenerator = gen.Struct(reflect.TypeOf(RouteTable_Spec{}), generators)

	return routeTable_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRouteTable_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTable_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["DisableBgpRoutePropagation"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

func Test_RouteTable_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RouteTable_STATUS to RouteTable_STATUS via AssignProperties_To_RouteTable_STATUS & AssignProperties_From_RouteTable_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRouteTable_STATUS, RouteTable_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRouteTable_STATUS tests if a specific instance of RouteTable_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRouteTable_STATUS(subject RouteTable_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20201101s.RouteTable_STATUS
	err := copied.AssignProperties_To_RouteTable_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RouteTable_STATUS
	err = actual.AssignProperties_From_RouteTable_STATUS(&other)
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

func Test_RouteTable_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTable_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTable_STATUS, RouteTable_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTable_STATUS runs a test to see if a specific instance of RouteTable_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTable_STATUS(subject RouteTable_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTable_STATUS
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

// Generator of RouteTable_STATUS instances for property testing - lazily instantiated by RouteTable_STATUSGenerator()
var routeTable_STATUSGenerator gopter.Gen

// RouteTable_STATUSGenerator returns a generator of RouteTable_STATUS instances for property testing.
func RouteTable_STATUSGenerator() gopter.Gen {
	if routeTable_STATUSGenerator != nil {
		return routeTable_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTable_STATUS(generators)
	routeTable_STATUSGenerator = gen.Struct(reflect.TypeOf(RouteTable_STATUS{}), generators)

	return routeTable_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRouteTable_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTable_STATUS(gens map[string]gopter.Gen) {
	gens["DisableBgpRoutePropagation"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
