// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Api_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: API entity create of update properties.
	Properties *ApiCreateOrUpdateProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Api_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (api Api_Spec_ARM) GetAPIVersion() string {
	return "2023-05-01-preview"
}

// GetName returns the Name of the resource
func (api *Api_Spec_ARM) GetName() string {
	return api.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/apis"
func (api *Api_Spec_ARM) GetType() string {
	return "Microsoft.ApiManagement/service/apis"
}

// API Create or Update Properties.
type ApiCreateOrUpdateProperties_ARM struct {
	// APIVersion: Indicates the version identifier of the API if the API is versioned
	APIVersion *string `json:"apiVersion,omitempty"`

	// ApiRevision: Describes the revision of the API. If no value is provided, default revision 1 is created
	ApiRevision *string `json:"apiRevision,omitempty"`

	// ApiRevisionDescription: Description of the API Revision.
	ApiRevisionDescription *string `json:"apiRevisionDescription,omitempty"`

	// ApiType: Type of API to create.
	// * `http` creates a REST API
	// * `soap` creates a SOAP pass-through API
	// * `websocket` creates websocket API
	// * `graphql` creates GraphQL API.
	// New types can be added in the future.
	ApiType *ApiCreateOrUpdateProperties_ApiType_ARM `json:"apiType,omitempty"`

	// ApiVersionDescription: Description of the API Version.
	ApiVersionDescription *string `json:"apiVersionDescription,omitempty"`

	// ApiVersionSet: Version set details
	ApiVersionSet   *ApiVersionSetContractDetails_ARM `json:"apiVersionSet,omitempty"`
	ApiVersionSetId *string                           `json:"apiVersionSetId,omitempty"`

	// AuthenticationSettings: Collection of authentication settings included into this API.
	AuthenticationSettings *AuthenticationSettingsContract_ARM `json:"authenticationSettings,omitempty"`

	// Contact: Contact information for the API.
	Contact *ApiContactInformation_ARM `json:"contact,omitempty"`

	// Description: Description of the API. May include HTML formatting tags.
	Description *string `json:"description,omitempty"`

	// DisplayName: API name. Must be 1 to 300 characters long.
	DisplayName *string `json:"displayName,omitempty"`

	// Format: Format of the Content in which the API is getting imported. New formats can be added in the future
	Format *ApiCreateOrUpdateProperties_Format_ARM `json:"format,omitempty"`

	// IsCurrent: Indicates if API revision is current api revision.
	IsCurrent *bool `json:"isCurrent,omitempty"`

	// License: License information for the API.
	License *ApiLicenseInformation_ARM `json:"license,omitempty"`

	// Path: Relative URL uniquely identifying this API and all of its resource paths within the API Management service
	// instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public
	// URL for this API.
	Path *string `json:"path,omitempty"`

	// Protocols: Describes on which protocols the operations in this API can be invoked.
	Protocols []ApiCreateOrUpdateProperties_Protocols_ARM `json:"protocols,omitempty"`

	// ServiceUrl: Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl  *string `json:"serviceUrl,omitempty"`
	SourceApiId *string `json:"sourceApiId,omitempty"`

	// SubscriptionKeyParameterNames: Protocols over which API is made available.
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContract_ARM `json:"subscriptionKeyParameterNames,omitempty"`

	// SubscriptionRequired: Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty"`

	// TermsOfServiceUrl:  A URL to the Terms of Service for the API. MUST be in the format of a URL.
	TermsOfServiceUrl *string `json:"termsOfServiceUrl,omitempty"`

	// TranslateRequiredQueryParameters: Strategy of translating required query parameters to template ones. By default has
	// value 'template'. Possible values: 'template', 'query'
	TranslateRequiredQueryParameters *ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM `json:"translateRequiredQueryParameters,omitempty"`

	// Type: Type of API.
	Type *ApiCreateOrUpdateProperties_Type_ARM `json:"type,omitempty"`

	// Value: Content value when Importing an API.
	Value *string `json:"value,omitempty"`

	// WsdlSelector: Criteria to limit import of WSDL to a subset of the document.
	WsdlSelector *ApiCreateOrUpdateProperties_WsdlSelector_ARM `json:"wsdlSelector,omitempty"`
}

// API contact information
type ApiContactInformation_ARM struct {
	// Email: The email address of the contact person/organization. MUST be in the format of an email address
	Email *string `json:"email,omitempty"`

	// Name: The identifying name of the contact person/organization
	Name *string `json:"name,omitempty"`

	// Url: The URL pointing to the contact information. MUST be in the format of a URL
	Url *string `json:"url,omitempty"`
}

// +kubebuilder:validation:Enum={"graphql","grpc","http","odata","soap","websocket"}
type ApiCreateOrUpdateProperties_ApiType_ARM string

const (
	ApiCreateOrUpdateProperties_ApiType_ARM_Graphql   = ApiCreateOrUpdateProperties_ApiType_ARM("graphql")
	ApiCreateOrUpdateProperties_ApiType_ARM_Grpc      = ApiCreateOrUpdateProperties_ApiType_ARM("grpc")
	ApiCreateOrUpdateProperties_ApiType_ARM_Http      = ApiCreateOrUpdateProperties_ApiType_ARM("http")
	ApiCreateOrUpdateProperties_ApiType_ARM_Odata     = ApiCreateOrUpdateProperties_ApiType_ARM("odata")
	ApiCreateOrUpdateProperties_ApiType_ARM_Soap      = ApiCreateOrUpdateProperties_ApiType_ARM("soap")
	ApiCreateOrUpdateProperties_ApiType_ARM_Websocket = ApiCreateOrUpdateProperties_ApiType_ARM("websocket")
)

// Mapping from string to ApiCreateOrUpdateProperties_ApiType_ARM
var apiCreateOrUpdateProperties_ApiType_ARM_Values = map[string]ApiCreateOrUpdateProperties_ApiType_ARM{
	"graphql":   ApiCreateOrUpdateProperties_ApiType_ARM_Graphql,
	"grpc":      ApiCreateOrUpdateProperties_ApiType_ARM_Grpc,
	"http":      ApiCreateOrUpdateProperties_ApiType_ARM_Http,
	"odata":     ApiCreateOrUpdateProperties_ApiType_ARM_Odata,
	"soap":      ApiCreateOrUpdateProperties_ApiType_ARM_Soap,
	"websocket": ApiCreateOrUpdateProperties_ApiType_ARM_Websocket,
}

// +kubebuilder:validation:Enum={"graphql-link","grpc","grpc-link","odata","odata-link","openapi","openapi+json","openapi+json-link","openapi-link","swagger-json","swagger-link-json","wadl-link-json","wadl-xml","wsdl","wsdl-link"}
type ApiCreateOrUpdateProperties_Format_ARM string

const (
	ApiCreateOrUpdateProperties_Format_ARM_GraphqlLink     = ApiCreateOrUpdateProperties_Format_ARM("graphql-link")
	ApiCreateOrUpdateProperties_Format_ARM_Grpc            = ApiCreateOrUpdateProperties_Format_ARM("grpc")
	ApiCreateOrUpdateProperties_Format_ARM_GrpcLink        = ApiCreateOrUpdateProperties_Format_ARM("grpc-link")
	ApiCreateOrUpdateProperties_Format_ARM_Odata           = ApiCreateOrUpdateProperties_Format_ARM("odata")
	ApiCreateOrUpdateProperties_Format_ARM_OdataLink       = ApiCreateOrUpdateProperties_Format_ARM("odata-link")
	ApiCreateOrUpdateProperties_Format_ARM_Openapi         = ApiCreateOrUpdateProperties_Format_ARM("openapi")
	ApiCreateOrUpdateProperties_Format_ARM_OpenapiJson     = ApiCreateOrUpdateProperties_Format_ARM("openapi+json")
	ApiCreateOrUpdateProperties_Format_ARM_OpenapiJsonLink = ApiCreateOrUpdateProperties_Format_ARM("openapi+json-link")
	ApiCreateOrUpdateProperties_Format_ARM_OpenapiLink     = ApiCreateOrUpdateProperties_Format_ARM("openapi-link")
	ApiCreateOrUpdateProperties_Format_ARM_SwaggerJson     = ApiCreateOrUpdateProperties_Format_ARM("swagger-json")
	ApiCreateOrUpdateProperties_Format_ARM_SwaggerLinkJson = ApiCreateOrUpdateProperties_Format_ARM("swagger-link-json")
	ApiCreateOrUpdateProperties_Format_ARM_WadlLinkJson    = ApiCreateOrUpdateProperties_Format_ARM("wadl-link-json")
	ApiCreateOrUpdateProperties_Format_ARM_WadlXml         = ApiCreateOrUpdateProperties_Format_ARM("wadl-xml")
	ApiCreateOrUpdateProperties_Format_ARM_Wsdl            = ApiCreateOrUpdateProperties_Format_ARM("wsdl")
	ApiCreateOrUpdateProperties_Format_ARM_WsdlLink        = ApiCreateOrUpdateProperties_Format_ARM("wsdl-link")
)

// Mapping from string to ApiCreateOrUpdateProperties_Format_ARM
var apiCreateOrUpdateProperties_Format_ARM_Values = map[string]ApiCreateOrUpdateProperties_Format_ARM{
	"graphql-link":      ApiCreateOrUpdateProperties_Format_ARM_GraphqlLink,
	"grpc":              ApiCreateOrUpdateProperties_Format_ARM_Grpc,
	"grpc-link":         ApiCreateOrUpdateProperties_Format_ARM_GrpcLink,
	"odata":             ApiCreateOrUpdateProperties_Format_ARM_Odata,
	"odata-link":        ApiCreateOrUpdateProperties_Format_ARM_OdataLink,
	"openapi":           ApiCreateOrUpdateProperties_Format_ARM_Openapi,
	"openapi+json":      ApiCreateOrUpdateProperties_Format_ARM_OpenapiJson,
	"openapi+json-link": ApiCreateOrUpdateProperties_Format_ARM_OpenapiJsonLink,
	"openapi-link":      ApiCreateOrUpdateProperties_Format_ARM_OpenapiLink,
	"swagger-json":      ApiCreateOrUpdateProperties_Format_ARM_SwaggerJson,
	"swagger-link-json": ApiCreateOrUpdateProperties_Format_ARM_SwaggerLinkJson,
	"wadl-link-json":    ApiCreateOrUpdateProperties_Format_ARM_WadlLinkJson,
	"wadl-xml":          ApiCreateOrUpdateProperties_Format_ARM_WadlXml,
	"wsdl":              ApiCreateOrUpdateProperties_Format_ARM_Wsdl,
	"wsdl-link":         ApiCreateOrUpdateProperties_Format_ARM_WsdlLink,
}

// +kubebuilder:validation:Enum={"http","https","ws","wss"}
type ApiCreateOrUpdateProperties_Protocols_ARM string

const (
	ApiCreateOrUpdateProperties_Protocols_ARM_Http  = ApiCreateOrUpdateProperties_Protocols_ARM("http")
	ApiCreateOrUpdateProperties_Protocols_ARM_Https = ApiCreateOrUpdateProperties_Protocols_ARM("https")
	ApiCreateOrUpdateProperties_Protocols_ARM_Ws    = ApiCreateOrUpdateProperties_Protocols_ARM("ws")
	ApiCreateOrUpdateProperties_Protocols_ARM_Wss   = ApiCreateOrUpdateProperties_Protocols_ARM("wss")
)

// Mapping from string to ApiCreateOrUpdateProperties_Protocols_ARM
var apiCreateOrUpdateProperties_Protocols_ARM_Values = map[string]ApiCreateOrUpdateProperties_Protocols_ARM{
	"http":  ApiCreateOrUpdateProperties_Protocols_ARM_Http,
	"https": ApiCreateOrUpdateProperties_Protocols_ARM_Https,
	"ws":    ApiCreateOrUpdateProperties_Protocols_ARM_Ws,
	"wss":   ApiCreateOrUpdateProperties_Protocols_ARM_Wss,
}

// +kubebuilder:validation:Enum={"query","template"}
type ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM string

const (
	ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM_Query    = ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM("query")
	ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM_Template = ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM("template")
)

// Mapping from string to ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM
var apiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM_Values = map[string]ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM{
	"query":    ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM_Query,
	"template": ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters_ARM_Template,
}

// +kubebuilder:validation:Enum={"graphql","grpc","http","odata","soap","websocket"}
type ApiCreateOrUpdateProperties_Type_ARM string

const (
	ApiCreateOrUpdateProperties_Type_ARM_Graphql   = ApiCreateOrUpdateProperties_Type_ARM("graphql")
	ApiCreateOrUpdateProperties_Type_ARM_Grpc      = ApiCreateOrUpdateProperties_Type_ARM("grpc")
	ApiCreateOrUpdateProperties_Type_ARM_Http      = ApiCreateOrUpdateProperties_Type_ARM("http")
	ApiCreateOrUpdateProperties_Type_ARM_Odata     = ApiCreateOrUpdateProperties_Type_ARM("odata")
	ApiCreateOrUpdateProperties_Type_ARM_Soap      = ApiCreateOrUpdateProperties_Type_ARM("soap")
	ApiCreateOrUpdateProperties_Type_ARM_Websocket = ApiCreateOrUpdateProperties_Type_ARM("websocket")
)

// Mapping from string to ApiCreateOrUpdateProperties_Type_ARM
var apiCreateOrUpdateProperties_Type_ARM_Values = map[string]ApiCreateOrUpdateProperties_Type_ARM{
	"graphql":   ApiCreateOrUpdateProperties_Type_ARM_Graphql,
	"grpc":      ApiCreateOrUpdateProperties_Type_ARM_Grpc,
	"http":      ApiCreateOrUpdateProperties_Type_ARM_Http,
	"odata":     ApiCreateOrUpdateProperties_Type_ARM_Odata,
	"soap":      ApiCreateOrUpdateProperties_Type_ARM_Soap,
	"websocket": ApiCreateOrUpdateProperties_Type_ARM_Websocket,
}

type ApiCreateOrUpdateProperties_WsdlSelector_ARM struct {
	// WsdlEndpointName: Name of endpoint(port) to import from WSDL
	WsdlEndpointName *string `json:"wsdlEndpointName,omitempty"`

	// WsdlServiceName: Name of service to import from WSDL
	WsdlServiceName *string `json:"wsdlServiceName,omitempty"`
}

// API license information
type ApiLicenseInformation_ARM struct {
	// Name: The license name used for the API
	Name *string `json:"name,omitempty"`

	// Url: A URL to the license used for the API. MUST be in the format of a URL
	Url *string `json:"url,omitempty"`
}

// An API Version Set contains the common configuration for a set of API Versions relating
type ApiVersionSetContractDetails_ARM struct {
	// Description: Description of API Version Set.
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`

	// Name: The display Name of the API Version Set.
	Name *string `json:"name,omitempty"`

	// VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
	VersionHeaderName *string `json:"versionHeaderName,omitempty"`

	// VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
	VersionQueryName *string `json:"versionQueryName,omitempty"`

	// VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.
	VersioningScheme *ApiVersionSetContractDetails_VersioningScheme_ARM `json:"versioningScheme,omitempty"`
}

// API Authentication Settings.
type AuthenticationSettingsContract_ARM struct {
	// OAuth2: OAuth2 Authentication settings
	OAuth2 *OAuth2AuthenticationSettingsContract_ARM `json:"oAuth2,omitempty"`

	// OAuth2AuthenticationSettings: Collection of OAuth2 authentication settings included into this API.
	OAuth2AuthenticationSettings []OAuth2AuthenticationSettingsContract_ARM `json:"oAuth2AuthenticationSettings,omitempty"`

	// Openid: OpenID Connect Authentication Settings
	Openid *OpenIdAuthenticationSettingsContract_ARM `json:"openid,omitempty"`

	// OpenidAuthenticationSettings: Collection of Open ID Connect authentication settings included into this API.
	OpenidAuthenticationSettings []OpenIdAuthenticationSettingsContract_ARM `json:"openidAuthenticationSettings,omitempty"`
}

// Subscription key parameter names details.
type SubscriptionKeyParameterNamesContract_ARM struct {
	// Header: Subscription key header name.
	Header *string `json:"header,omitempty"`

	// Query: Subscription key query string parameter name.
	Query *string `json:"query,omitempty"`
}

// +kubebuilder:validation:Enum={"Header","Query","Segment"}
type ApiVersionSetContractDetails_VersioningScheme_ARM string

const (
	ApiVersionSetContractDetails_VersioningScheme_ARM_Header  = ApiVersionSetContractDetails_VersioningScheme_ARM("Header")
	ApiVersionSetContractDetails_VersioningScheme_ARM_Query   = ApiVersionSetContractDetails_VersioningScheme_ARM("Query")
	ApiVersionSetContractDetails_VersioningScheme_ARM_Segment = ApiVersionSetContractDetails_VersioningScheme_ARM("Segment")
)

// Mapping from string to ApiVersionSetContractDetails_VersioningScheme_ARM
var apiVersionSetContractDetails_VersioningScheme_ARM_Values = map[string]ApiVersionSetContractDetails_VersioningScheme_ARM{
	"header":  ApiVersionSetContractDetails_VersioningScheme_ARM_Header,
	"query":   ApiVersionSetContractDetails_VersioningScheme_ARM_Query,
	"segment": ApiVersionSetContractDetails_VersioningScheme_ARM_Segment,
}

// API OAuth2 Authentication settings details.
type OAuth2AuthenticationSettingsContract_ARM struct {
	// AuthorizationServerId: OAuth authorization server identifier.
	AuthorizationServerId *string `json:"authorizationServerId,omitempty"`

	// Scope: operations scope.
	Scope *string `json:"scope,omitempty"`
}

// API OAuth2 Authentication settings details.
type OpenIdAuthenticationSettingsContract_ARM struct {
	// BearerTokenSendingMethods: How to send token to the server.
	BearerTokenSendingMethods []BearerTokenSendingMethodsContract_ARM `json:"bearerTokenSendingMethods,omitempty"`

	// OpenidProviderId: OAuth authorization server identifier.
	OpenidProviderId *string `json:"openidProviderId,omitempty"`
}

// Form of an authorization grant, which the client uses to request the access token.
// +kubebuilder:validation:Enum={"authorizationHeader","query"}
type BearerTokenSendingMethodsContract_ARM string

const (
	BearerTokenSendingMethodsContract_ARM_AuthorizationHeader = BearerTokenSendingMethodsContract_ARM("authorizationHeader")
	BearerTokenSendingMethodsContract_ARM_Query               = BearerTokenSendingMethodsContract_ARM("query")
)

// Mapping from string to BearerTokenSendingMethodsContract_ARM
var bearerTokenSendingMethodsContract_ARM_Values = map[string]BearerTokenSendingMethodsContract_ARM{
	"authorizationheader": BearerTokenSendingMethodsContract_ARM_AuthorizationHeader,
	"query":               BearerTokenSendingMethodsContract_ARM_Query,
}