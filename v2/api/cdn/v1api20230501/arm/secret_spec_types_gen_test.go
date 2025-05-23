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

func Test_AzureFirstPartyManagedCertificateParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AzureFirstPartyManagedCertificateParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAzureFirstPartyManagedCertificateParameters, AzureFirstPartyManagedCertificateParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAzureFirstPartyManagedCertificateParameters runs a test to see if a specific instance of AzureFirstPartyManagedCertificateParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForAzureFirstPartyManagedCertificateParameters(subject AzureFirstPartyManagedCertificateParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AzureFirstPartyManagedCertificateParameters
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

// Generator of AzureFirstPartyManagedCertificateParameters instances for property testing - lazily instantiated by
// AzureFirstPartyManagedCertificateParametersGenerator()
var azureFirstPartyManagedCertificateParametersGenerator gopter.Gen

// AzureFirstPartyManagedCertificateParametersGenerator returns a generator of AzureFirstPartyManagedCertificateParameters instances for property testing.
func AzureFirstPartyManagedCertificateParametersGenerator() gopter.Gen {
	if azureFirstPartyManagedCertificateParametersGenerator != nil {
		return azureFirstPartyManagedCertificateParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters(generators)
	azureFirstPartyManagedCertificateParametersGenerator = gen.Struct(reflect.TypeOf(AzureFirstPartyManagedCertificateParameters{}), generators)

	return azureFirstPartyManagedCertificateParametersGenerator
}

// AddIndependentPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAzureFirstPartyManagedCertificateParameters(gens map[string]gopter.Gen) {
	gens["SubjectAlternativeNames"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(AzureFirstPartyManagedCertificateParameters_Type_AzureFirstPartyManagedCertificate)
}

func Test_CustomerCertificateParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CustomerCertificateParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCustomerCertificateParameters, CustomerCertificateParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCustomerCertificateParameters runs a test to see if a specific instance of CustomerCertificateParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForCustomerCertificateParameters(subject CustomerCertificateParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CustomerCertificateParameters
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

// Generator of CustomerCertificateParameters instances for property testing - lazily instantiated by
// CustomerCertificateParametersGenerator()
var customerCertificateParametersGenerator gopter.Gen

// CustomerCertificateParametersGenerator returns a generator of CustomerCertificateParameters instances for property testing.
// We first initialize customerCertificateParametersGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CustomerCertificateParametersGenerator() gopter.Gen {
	if customerCertificateParametersGenerator != nil {
		return customerCertificateParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomerCertificateParameters(generators)
	customerCertificateParametersGenerator = gen.Struct(reflect.TypeOf(CustomerCertificateParameters{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCustomerCertificateParameters(generators)
	AddRelatedPropertyGeneratorsForCustomerCertificateParameters(generators)
	customerCertificateParametersGenerator = gen.Struct(reflect.TypeOf(CustomerCertificateParameters{}), generators)

	return customerCertificateParametersGenerator
}

// AddIndependentPropertyGeneratorsForCustomerCertificateParameters is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCustomerCertificateParameters(gens map[string]gopter.Gen) {
	gens["SecretVersion"] = gen.PtrOf(gen.AlphaString())
	gens["SubjectAlternativeNames"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(CustomerCertificateParameters_Type_CustomerCertificate)
	gens["UseLatestVersion"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCustomerCertificateParameters is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCustomerCertificateParameters(gens map[string]gopter.Gen) {
	gens["SecretSource"] = gen.PtrOf(ResourceReferenceGenerator())
}

func Test_ManagedCertificateParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedCertificateParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedCertificateParameters, ManagedCertificateParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedCertificateParameters runs a test to see if a specific instance of ManagedCertificateParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedCertificateParameters(subject ManagedCertificateParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedCertificateParameters
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

// Generator of ManagedCertificateParameters instances for property testing - lazily instantiated by
// ManagedCertificateParametersGenerator()
var managedCertificateParametersGenerator gopter.Gen

// ManagedCertificateParametersGenerator returns a generator of ManagedCertificateParameters instances for property testing.
func ManagedCertificateParametersGenerator() gopter.Gen {
	if managedCertificateParametersGenerator != nil {
		return managedCertificateParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedCertificateParameters(generators)
	managedCertificateParametersGenerator = gen.Struct(reflect.TypeOf(ManagedCertificateParameters{}), generators)

	return managedCertificateParametersGenerator
}

// AddIndependentPropertyGeneratorsForManagedCertificateParameters is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedCertificateParameters(gens map[string]gopter.Gen) {
	gens["Type"] = gen.OneConstOf(ManagedCertificateParameters_Type_ManagedCertificate)
}

func Test_SecretParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecretParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecretParameters, SecretParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecretParameters runs a test to see if a specific instance of SecretParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForSecretParameters(subject SecretParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecretParameters
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

// Generator of SecretParameters instances for property testing - lazily instantiated by SecretParametersGenerator()
var secretParametersGenerator gopter.Gen

// SecretParametersGenerator returns a generator of SecretParameters instances for property testing.
func SecretParametersGenerator() gopter.Gen {
	if secretParametersGenerator != nil {
		return secretParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecretParameters(generators)

	// handle OneOf by choosing only one field to instantiate
	var gens []gopter.Gen
	for propName, propGen := range generators {
		props := map[string]gopter.Gen{propName: propGen}
		gens = append(gens, gen.Struct(reflect.TypeOf(SecretParameters{}), props))
	}
	secretParametersGenerator = gen.OneGenOf(gens...)

	return secretParametersGenerator
}

// AddRelatedPropertyGeneratorsForSecretParameters is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecretParameters(gens map[string]gopter.Gen) {
	gens["AzureFirstPartyManagedCertificate"] = AzureFirstPartyManagedCertificateParametersGenerator().Map(func(it AzureFirstPartyManagedCertificateParameters) *AzureFirstPartyManagedCertificateParameters {
		return &it
	}) // generate one case for OneOf type
	gens["CustomerCertificate"] = CustomerCertificateParametersGenerator().Map(func(it CustomerCertificateParameters) *CustomerCertificateParameters {
		return &it
	}) // generate one case for OneOf type
	gens["ManagedCertificate"] = ManagedCertificateParametersGenerator().Map(func(it ManagedCertificateParameters) *ManagedCertificateParameters {
		return &it
	}) // generate one case for OneOf type
	gens["UrlSigningKey"] = UrlSigningKeyParametersGenerator().Map(func(it UrlSigningKeyParameters) *UrlSigningKeyParameters {
		return &it
	}) // generate one case for OneOf type
}

func Test_SecretProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecretProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecretProperties, SecretPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecretProperties runs a test to see if a specific instance of SecretProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForSecretProperties(subject SecretProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecretProperties
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

// Generator of SecretProperties instances for property testing - lazily instantiated by SecretPropertiesGenerator()
var secretPropertiesGenerator gopter.Gen

// SecretPropertiesGenerator returns a generator of SecretProperties instances for property testing.
func SecretPropertiesGenerator() gopter.Gen {
	if secretPropertiesGenerator != nil {
		return secretPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecretProperties(generators)
	secretPropertiesGenerator = gen.Struct(reflect.TypeOf(SecretProperties{}), generators)

	return secretPropertiesGenerator
}

// AddRelatedPropertyGeneratorsForSecretProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecretProperties(gens map[string]gopter.Gen) {
	gens["Parameters"] = gen.PtrOf(SecretParametersGenerator())
}

func Test_Secret_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Secret_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecret_Spec, Secret_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecret_Spec runs a test to see if a specific instance of Secret_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSecret_Spec(subject Secret_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Secret_Spec
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

// Generator of Secret_Spec instances for property testing - lazily instantiated by Secret_SpecGenerator()
var secret_SpecGenerator gopter.Gen

// Secret_SpecGenerator returns a generator of Secret_Spec instances for property testing.
// We first initialize secret_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Secret_SpecGenerator() gopter.Gen {
	if secret_SpecGenerator != nil {
		return secret_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecret_Spec(generators)
	secret_SpecGenerator = gen.Struct(reflect.TypeOf(Secret_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecret_Spec(generators)
	AddRelatedPropertyGeneratorsForSecret_Spec(generators)
	secret_SpecGenerator = gen.Struct(reflect.TypeOf(Secret_Spec{}), generators)

	return secret_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSecret_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecret_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForSecret_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecret_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SecretPropertiesGenerator())
}

func Test_UrlSigningKeyParameters_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UrlSigningKeyParameters via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUrlSigningKeyParameters, UrlSigningKeyParametersGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUrlSigningKeyParameters runs a test to see if a specific instance of UrlSigningKeyParameters round trips to JSON and back losslessly
func RunJSONSerializationTestForUrlSigningKeyParameters(subject UrlSigningKeyParameters) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UrlSigningKeyParameters
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

// Generator of UrlSigningKeyParameters instances for property testing - lazily instantiated by
// UrlSigningKeyParametersGenerator()
var urlSigningKeyParametersGenerator gopter.Gen

// UrlSigningKeyParametersGenerator returns a generator of UrlSigningKeyParameters instances for property testing.
// We first initialize urlSigningKeyParametersGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func UrlSigningKeyParametersGenerator() gopter.Gen {
	if urlSigningKeyParametersGenerator != nil {
		return urlSigningKeyParametersGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUrlSigningKeyParameters(generators)
	urlSigningKeyParametersGenerator = gen.Struct(reflect.TypeOf(UrlSigningKeyParameters{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUrlSigningKeyParameters(generators)
	AddRelatedPropertyGeneratorsForUrlSigningKeyParameters(generators)
	urlSigningKeyParametersGenerator = gen.Struct(reflect.TypeOf(UrlSigningKeyParameters{}), generators)

	return urlSigningKeyParametersGenerator
}

// AddIndependentPropertyGeneratorsForUrlSigningKeyParameters is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUrlSigningKeyParameters(gens map[string]gopter.Gen) {
	gens["KeyId"] = gen.PtrOf(gen.AlphaString())
	gens["SecretVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.OneConstOf(UrlSigningKeyParameters_Type_UrlSigningKey)
}

// AddRelatedPropertyGeneratorsForUrlSigningKeyParameters is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUrlSigningKeyParameters(gens map[string]gopter.Gen) {
	gens["SecretSource"] = gen.PtrOf(ResourceReferenceGenerator())
}
