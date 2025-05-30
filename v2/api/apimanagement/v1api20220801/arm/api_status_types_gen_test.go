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

func Test_ApiContactInformation_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiContactInformation_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiContactInformation_STATUS, ApiContactInformation_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiContactInformation_STATUS runs a test to see if a specific instance of ApiContactInformation_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForApiContactInformation_STATUS(subject ApiContactInformation_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiContactInformation_STATUS
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

// Generator of ApiContactInformation_STATUS instances for property testing - lazily instantiated by
// ApiContactInformation_STATUSGenerator()
var apiContactInformation_STATUSGenerator gopter.Gen

// ApiContactInformation_STATUSGenerator returns a generator of ApiContactInformation_STATUS instances for property testing.
func ApiContactInformation_STATUSGenerator() gopter.Gen {
	if apiContactInformation_STATUSGenerator != nil {
		return apiContactInformation_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiContactInformation_STATUS(generators)
	apiContactInformation_STATUSGenerator = gen.Struct(reflect.TypeOf(ApiContactInformation_STATUS{}), generators)

	return apiContactInformation_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForApiContactInformation_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiContactInformation_STATUS(gens map[string]gopter.Gen) {
	gens["Email"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

func Test_ApiContractProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiContractProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiContractProperties_STATUS, ApiContractProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiContractProperties_STATUS runs a test to see if a specific instance of ApiContractProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForApiContractProperties_STATUS(subject ApiContractProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiContractProperties_STATUS
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

// Generator of ApiContractProperties_STATUS instances for property testing - lazily instantiated by
// ApiContractProperties_STATUSGenerator()
var apiContractProperties_STATUSGenerator gopter.Gen

// ApiContractProperties_STATUSGenerator returns a generator of ApiContractProperties_STATUS instances for property testing.
// We first initialize apiContractProperties_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ApiContractProperties_STATUSGenerator() gopter.Gen {
	if apiContractProperties_STATUSGenerator != nil {
		return apiContractProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiContractProperties_STATUS(generators)
	apiContractProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ApiContractProperties_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiContractProperties_STATUS(generators)
	AddRelatedPropertyGeneratorsForApiContractProperties_STATUS(generators)
	apiContractProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ApiContractProperties_STATUS{}), generators)

	return apiContractProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForApiContractProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiContractProperties_STATUS(gens map[string]gopter.Gen) {
	gens["APIVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ApiRevision"] = gen.PtrOf(gen.AlphaString())
	gens["ApiRevisionDescription"] = gen.PtrOf(gen.AlphaString())
	gens["ApiVersionDescription"] = gen.PtrOf(gen.AlphaString())
	gens["ApiVersionSetId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["IsCurrent"] = gen.PtrOf(gen.Bool())
	gens["IsOnline"] = gen.PtrOf(gen.Bool())
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Protocols"] = gen.SliceOf(gen.OneConstOf(
		ApiContractProperties_Protocols_STATUS_Http,
		ApiContractProperties_Protocols_STATUS_Https,
		ApiContractProperties_Protocols_STATUS_Ws,
		ApiContractProperties_Protocols_STATUS_Wss))
	gens["ServiceUrl"] = gen.PtrOf(gen.AlphaString())
	gens["SourceApiId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionRequired"] = gen.PtrOf(gen.Bool())
	gens["TermsOfServiceUrl"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		ApiContractProperties_Type_STATUS_Graphql,
		ApiContractProperties_Type_STATUS_Http,
		ApiContractProperties_Type_STATUS_Soap,
		ApiContractProperties_Type_STATUS_Websocket))
}

// AddRelatedPropertyGeneratorsForApiContractProperties_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApiContractProperties_STATUS(gens map[string]gopter.Gen) {
	gens["ApiVersionSet"] = gen.PtrOf(ApiVersionSetContractDetails_STATUSGenerator())
	gens["AuthenticationSettings"] = gen.PtrOf(AuthenticationSettingsContract_STATUSGenerator())
	gens["Contact"] = gen.PtrOf(ApiContactInformation_STATUSGenerator())
	gens["License"] = gen.PtrOf(ApiLicenseInformation_STATUSGenerator())
	gens["SubscriptionKeyParameterNames"] = gen.PtrOf(SubscriptionKeyParameterNamesContract_STATUSGenerator())
}

func Test_ApiLicenseInformation_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiLicenseInformation_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiLicenseInformation_STATUS, ApiLicenseInformation_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiLicenseInformation_STATUS runs a test to see if a specific instance of ApiLicenseInformation_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForApiLicenseInformation_STATUS(subject ApiLicenseInformation_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiLicenseInformation_STATUS
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

// Generator of ApiLicenseInformation_STATUS instances for property testing - lazily instantiated by
// ApiLicenseInformation_STATUSGenerator()
var apiLicenseInformation_STATUSGenerator gopter.Gen

// ApiLicenseInformation_STATUSGenerator returns a generator of ApiLicenseInformation_STATUS instances for property testing.
func ApiLicenseInformation_STATUSGenerator() gopter.Gen {
	if apiLicenseInformation_STATUSGenerator != nil {
		return apiLicenseInformation_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiLicenseInformation_STATUS(generators)
	apiLicenseInformation_STATUSGenerator = gen.Struct(reflect.TypeOf(ApiLicenseInformation_STATUS{}), generators)

	return apiLicenseInformation_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForApiLicenseInformation_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiLicenseInformation_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

func Test_ApiVersionSetContractDetails_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiVersionSetContractDetails_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiVersionSetContractDetails_STATUS, ApiVersionSetContractDetails_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiVersionSetContractDetails_STATUS runs a test to see if a specific instance of ApiVersionSetContractDetails_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForApiVersionSetContractDetails_STATUS(subject ApiVersionSetContractDetails_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiVersionSetContractDetails_STATUS
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

// Generator of ApiVersionSetContractDetails_STATUS instances for property testing - lazily instantiated by
// ApiVersionSetContractDetails_STATUSGenerator()
var apiVersionSetContractDetails_STATUSGenerator gopter.Gen

// ApiVersionSetContractDetails_STATUSGenerator returns a generator of ApiVersionSetContractDetails_STATUS instances for property testing.
func ApiVersionSetContractDetails_STATUSGenerator() gopter.Gen {
	if apiVersionSetContractDetails_STATUSGenerator != nil {
		return apiVersionSetContractDetails_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiVersionSetContractDetails_STATUS(generators)
	apiVersionSetContractDetails_STATUSGenerator = gen.Struct(reflect.TypeOf(ApiVersionSetContractDetails_STATUS{}), generators)

	return apiVersionSetContractDetails_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForApiVersionSetContractDetails_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiVersionSetContractDetails_STATUS(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["VersionHeaderName"] = gen.PtrOf(gen.AlphaString())
	gens["VersionQueryName"] = gen.PtrOf(gen.AlphaString())
	gens["VersioningScheme"] = gen.PtrOf(gen.OneConstOf(ApiVersionSetContractDetails_VersioningScheme_STATUS_Header, ApiVersionSetContractDetails_VersioningScheme_STATUS_Query, ApiVersionSetContractDetails_VersioningScheme_STATUS_Segment))
}

func Test_Api_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Api_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApi_STATUS, Api_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApi_STATUS runs a test to see if a specific instance of Api_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForApi_STATUS(subject Api_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Api_STATUS
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

// Generator of Api_STATUS instances for property testing - lazily instantiated by Api_STATUSGenerator()
var api_STATUSGenerator gopter.Gen

// Api_STATUSGenerator returns a generator of Api_STATUS instances for property testing.
// We first initialize api_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Api_STATUSGenerator() gopter.Gen {
	if api_STATUSGenerator != nil {
		return api_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApi_STATUS(generators)
	api_STATUSGenerator = gen.Struct(reflect.TypeOf(Api_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApi_STATUS(generators)
	AddRelatedPropertyGeneratorsForApi_STATUS(generators)
	api_STATUSGenerator = gen.Struct(reflect.TypeOf(Api_STATUS{}), generators)

	return api_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForApi_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApi_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForApi_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApi_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ApiContractProperties_STATUSGenerator())
}

func Test_AuthenticationSettingsContract_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthenticationSettingsContract_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthenticationSettingsContract_STATUS, AuthenticationSettingsContract_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthenticationSettingsContract_STATUS runs a test to see if a specific instance of AuthenticationSettingsContract_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthenticationSettingsContract_STATUS(subject AuthenticationSettingsContract_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthenticationSettingsContract_STATUS
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

// Generator of AuthenticationSettingsContract_STATUS instances for property testing - lazily instantiated by
// AuthenticationSettingsContract_STATUSGenerator()
var authenticationSettingsContract_STATUSGenerator gopter.Gen

// AuthenticationSettingsContract_STATUSGenerator returns a generator of AuthenticationSettingsContract_STATUS instances for property testing.
func AuthenticationSettingsContract_STATUSGenerator() gopter.Gen {
	if authenticationSettingsContract_STATUSGenerator != nil {
		return authenticationSettingsContract_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAuthenticationSettingsContract_STATUS(generators)
	authenticationSettingsContract_STATUSGenerator = gen.Struct(reflect.TypeOf(AuthenticationSettingsContract_STATUS{}), generators)

	return authenticationSettingsContract_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForAuthenticationSettingsContract_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthenticationSettingsContract_STATUS(gens map[string]gopter.Gen) {
	gens["OAuth2"] = gen.PtrOf(OAuth2AuthenticationSettingsContract_STATUSGenerator())
	gens["OAuth2AuthenticationSettings"] = gen.SliceOf(OAuth2AuthenticationSettingsContract_STATUSGenerator())
	gens["Openid"] = gen.PtrOf(OpenIdAuthenticationSettingsContract_STATUSGenerator())
	gens["OpenidAuthenticationSettings"] = gen.SliceOf(OpenIdAuthenticationSettingsContract_STATUSGenerator())
}

func Test_OAuth2AuthenticationSettingsContract_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OAuth2AuthenticationSettingsContract_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOAuth2AuthenticationSettingsContract_STATUS, OAuth2AuthenticationSettingsContract_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOAuth2AuthenticationSettingsContract_STATUS runs a test to see if a specific instance of OAuth2AuthenticationSettingsContract_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForOAuth2AuthenticationSettingsContract_STATUS(subject OAuth2AuthenticationSettingsContract_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OAuth2AuthenticationSettingsContract_STATUS
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

// Generator of OAuth2AuthenticationSettingsContract_STATUS instances for property testing - lazily instantiated by
// OAuth2AuthenticationSettingsContract_STATUSGenerator()
var oAuth2AuthenticationSettingsContract_STATUSGenerator gopter.Gen

// OAuth2AuthenticationSettingsContract_STATUSGenerator returns a generator of OAuth2AuthenticationSettingsContract_STATUS instances for property testing.
func OAuth2AuthenticationSettingsContract_STATUSGenerator() gopter.Gen {
	if oAuth2AuthenticationSettingsContract_STATUSGenerator != nil {
		return oAuth2AuthenticationSettingsContract_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOAuth2AuthenticationSettingsContract_STATUS(generators)
	oAuth2AuthenticationSettingsContract_STATUSGenerator = gen.Struct(reflect.TypeOf(OAuth2AuthenticationSettingsContract_STATUS{}), generators)

	return oAuth2AuthenticationSettingsContract_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForOAuth2AuthenticationSettingsContract_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOAuth2AuthenticationSettingsContract_STATUS(gens map[string]gopter.Gen) {
	gens["AuthorizationServerId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
}

func Test_OpenIdAuthenticationSettingsContract_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OpenIdAuthenticationSettingsContract_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOpenIdAuthenticationSettingsContract_STATUS, OpenIdAuthenticationSettingsContract_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOpenIdAuthenticationSettingsContract_STATUS runs a test to see if a specific instance of OpenIdAuthenticationSettingsContract_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForOpenIdAuthenticationSettingsContract_STATUS(subject OpenIdAuthenticationSettingsContract_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OpenIdAuthenticationSettingsContract_STATUS
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

// Generator of OpenIdAuthenticationSettingsContract_STATUS instances for property testing - lazily instantiated by
// OpenIdAuthenticationSettingsContract_STATUSGenerator()
var openIdAuthenticationSettingsContract_STATUSGenerator gopter.Gen

// OpenIdAuthenticationSettingsContract_STATUSGenerator returns a generator of OpenIdAuthenticationSettingsContract_STATUS instances for property testing.
func OpenIdAuthenticationSettingsContract_STATUSGenerator() gopter.Gen {
	if openIdAuthenticationSettingsContract_STATUSGenerator != nil {
		return openIdAuthenticationSettingsContract_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOpenIdAuthenticationSettingsContract_STATUS(generators)
	openIdAuthenticationSettingsContract_STATUSGenerator = gen.Struct(reflect.TypeOf(OpenIdAuthenticationSettingsContract_STATUS{}), generators)

	return openIdAuthenticationSettingsContract_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForOpenIdAuthenticationSettingsContract_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOpenIdAuthenticationSettingsContract_STATUS(gens map[string]gopter.Gen) {
	gens["BearerTokenSendingMethods"] = gen.SliceOf(gen.OneConstOf(BearerTokenSendingMethodsContract_STATUS_AuthorizationHeader, BearerTokenSendingMethodsContract_STATUS_Query))
	gens["OpenidProviderId"] = gen.PtrOf(gen.AlphaString())
}

func Test_SubscriptionKeyParameterNamesContract_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionKeyParameterNamesContract_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionKeyParameterNamesContract_STATUS, SubscriptionKeyParameterNamesContract_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionKeyParameterNamesContract_STATUS runs a test to see if a specific instance of SubscriptionKeyParameterNamesContract_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionKeyParameterNamesContract_STATUS(subject SubscriptionKeyParameterNamesContract_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionKeyParameterNamesContract_STATUS
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

// Generator of SubscriptionKeyParameterNamesContract_STATUS instances for property testing - lazily instantiated by
// SubscriptionKeyParameterNamesContract_STATUSGenerator()
var subscriptionKeyParameterNamesContract_STATUSGenerator gopter.Gen

// SubscriptionKeyParameterNamesContract_STATUSGenerator returns a generator of SubscriptionKeyParameterNamesContract_STATUS instances for property testing.
func SubscriptionKeyParameterNamesContract_STATUSGenerator() gopter.Gen {
	if subscriptionKeyParameterNamesContract_STATUSGenerator != nil {
		return subscriptionKeyParameterNamesContract_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionKeyParameterNamesContract_STATUS(generators)
	subscriptionKeyParameterNamesContract_STATUSGenerator = gen.Struct(reflect.TypeOf(SubscriptionKeyParameterNamesContract_STATUS{}), generators)

	return subscriptionKeyParameterNamesContract_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionKeyParameterNamesContract_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionKeyParameterNamesContract_STATUS(gens map[string]gopter.Gen) {
	gens["Header"] = gen.PtrOf(gen.AlphaString())
	gens["Query"] = gen.PtrOf(gen.AlphaString())
}
