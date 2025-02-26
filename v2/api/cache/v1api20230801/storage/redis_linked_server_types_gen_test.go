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

func Test_RedisLinkedServer_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer, RedisLinkedServerGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer runs a test to see if a specific instance of RedisLinkedServer round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer(subject RedisLinkedServer) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer
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

// Generator of RedisLinkedServer instances for property testing - lazily instantiated by RedisLinkedServerGenerator()
var redisLinkedServerGenerator gopter.Gen

// RedisLinkedServerGenerator returns a generator of RedisLinkedServer instances for property testing.
func RedisLinkedServerGenerator() gopter.Gen {
	if redisLinkedServerGenerator != nil {
		return redisLinkedServerGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisLinkedServer(generators)
	redisLinkedServerGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer{}), generators)

	return redisLinkedServerGenerator
}

// AddRelatedPropertyGeneratorsForRedisLinkedServer is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisLinkedServer(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisLinkedServer_SpecGenerator()
	gens["Status"] = Redis_LinkedServer_STATUSGenerator()
}

func Test_RedisLinkedServerOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServerOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerOperatorSpec, RedisLinkedServerOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerOperatorSpec runs a test to see if a specific instance of RedisLinkedServerOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerOperatorSpec(subject RedisLinkedServerOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServerOperatorSpec
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

// Generator of RedisLinkedServerOperatorSpec instances for property testing - lazily instantiated by
// RedisLinkedServerOperatorSpecGenerator()
var redisLinkedServerOperatorSpecGenerator gopter.Gen

// RedisLinkedServerOperatorSpecGenerator returns a generator of RedisLinkedServerOperatorSpec instances for property testing.
func RedisLinkedServerOperatorSpecGenerator() gopter.Gen {
	if redisLinkedServerOperatorSpecGenerator != nil {
		return redisLinkedServerOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	redisLinkedServerOperatorSpecGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerOperatorSpec{}), generators)

	return redisLinkedServerOperatorSpecGenerator
}

func Test_RedisLinkedServer_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer_Spec, RedisLinkedServer_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer_Spec runs a test to see if a specific instance of RedisLinkedServer_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer_Spec(subject RedisLinkedServer_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_Spec
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

// Generator of RedisLinkedServer_Spec instances for property testing - lazily instantiated by
// RedisLinkedServer_SpecGenerator()
var redisLinkedServer_SpecGenerator gopter.Gen

// RedisLinkedServer_SpecGenerator returns a generator of RedisLinkedServer_Spec instances for property testing.
// We first initialize redisLinkedServer_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisLinkedServer_SpecGenerator() gopter.Gen {
	if redisLinkedServer_SpecGenerator != nil {
		return redisLinkedServer_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec(generators)
	redisLinkedServer_SpecGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec(generators)
	AddRelatedPropertyGeneratorsForRedisLinkedServer_Spec(generators)
	redisLinkedServer_SpecGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_Spec{}), generators)

	return redisLinkedServer_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["ServerRole"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisLinkedServer_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisLinkedServer_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(RedisLinkedServerOperatorSpecGenerator())
}

func Test_Redis_LinkedServer_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_LinkedServer_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis_LinkedServer_STATUS, Redis_LinkedServer_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis_LinkedServer_STATUS runs a test to see if a specific instance of Redis_LinkedServer_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis_LinkedServer_STATUS(subject Redis_LinkedServer_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_LinkedServer_STATUS
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

// Generator of Redis_LinkedServer_STATUS instances for property testing - lazily instantiated by
// Redis_LinkedServer_STATUSGenerator()
var redis_LinkedServer_STATUSGenerator gopter.Gen

// Redis_LinkedServer_STATUSGenerator returns a generator of Redis_LinkedServer_STATUS instances for property testing.
func Redis_LinkedServer_STATUSGenerator() gopter.Gen {
	if redis_LinkedServer_STATUSGenerator != nil {
		return redis_LinkedServer_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedis_LinkedServer_STATUS(generators)
	redis_LinkedServer_STATUSGenerator = gen.Struct(reflect.TypeOf(Redis_LinkedServer_STATUS{}), generators)

	return redis_LinkedServer_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedis_LinkedServer_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedis_LinkedServer_STATUS(gens map[string]gopter.Gen) {
	gens["GeoReplicatedPrimaryHostName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheId"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryHostName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ServerRole"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
