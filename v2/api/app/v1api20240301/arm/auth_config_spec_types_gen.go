// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type AuthConfig_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: AuthConfig resource specific properties
	Properties *ContainerApps_AuthConfig_Properties_Spec `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &AuthConfig_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-03-01"
func (config AuthConfig_Spec) GetAPIVersion() string {
	return "2024-03-01"
}

// GetName returns the Name of the resource
func (config *AuthConfig_Spec) GetName() string {
	return config.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.App/containerApps/authConfigs"
func (config *AuthConfig_Spec) GetType() string {
	return "Microsoft.App/containerApps/authConfigs"
}

type ContainerApps_AuthConfig_Properties_Spec struct {
	// EncryptionSettings: The configuration settings of the secrets references of encryption key and signing key for
	// ContainerApp Service Authentication/Authorization.
	EncryptionSettings *EncryptionSettings `json:"encryptionSettings,omitempty"`

	// GlobalValidation: The configuration settings that determines the validation flow of users using  Service
	// Authentication/Authorization.
	GlobalValidation *GlobalValidation `json:"globalValidation,omitempty"`

	// HttpSettings: The configuration settings of the HTTP requests for authentication and authorization requests made against
	// ContainerApp Service Authentication/Authorization.
	HttpSettings *HttpSettings `json:"httpSettings,omitempty"`

	// IdentityProviders: The configuration settings of each of the identity providers used to configure ContainerApp Service
	// Authentication/Authorization.
	IdentityProviders *IdentityProviders `json:"identityProviders,omitempty"`

	// Login: The configuration settings of the login flow of users using ContainerApp Service Authentication/Authorization.
	Login *Login `json:"login,omitempty"`

	// Platform: The configuration settings of the platform of ContainerApp Service Authentication/Authorization.
	Platform *AuthPlatform `json:"platform,omitempty"`
}

// The configuration settings of the platform of ContainerApp Service Authentication/Authorization.
type AuthPlatform struct {
	// Enabled: <code>true</code> if the Authentication / Authorization feature is enabled for the current app; otherwise,
	// <code>false</code>.
	Enabled *bool `json:"enabled,omitempty"`

	// RuntimeVersion: The RuntimeVersion of the Authentication / Authorization feature in use for the current app.
	// The setting in this value can control the behavior of certain features in the Authentication / Authorization module.
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`
}

// The configuration settings of the secrets references of encryption key and signing key for ContainerApp Service
// Authentication/Authorization.
type EncryptionSettings struct {
	// ContainerAppAuthEncryptionSecretName: The secret name which is referenced for EncryptionKey.
	ContainerAppAuthEncryptionSecretName *string `json:"containerAppAuthEncryptionSecretName,omitempty"`

	// ContainerAppAuthSigningSecretName: The secret name which is referenced for SigningKey.
	ContainerAppAuthSigningSecretName *string `json:"containerAppAuthSigningSecretName,omitempty"`
}

// The configuration settings that determines the validation flow of users using ContainerApp Service
// Authentication/Authorization.
type GlobalValidation struct {
	// ExcludedPaths: The paths for which unauthenticated flow would not be redirected to the login page.
	ExcludedPaths []string `json:"excludedPaths,omitempty"`

	// RedirectToProvider: The default authentication provider to use when multiple providers are configured.
	// This setting is only needed if multiple providers are configured and the unauthenticated client
	// action is set to "RedirectToLoginPage".
	RedirectToProvider *string `json:"redirectToProvider,omitempty"`

	// UnauthenticatedClientAction: The action to take when an unauthenticated client attempts to access the app.
	UnauthenticatedClientAction *GlobalValidation_UnauthenticatedClientAction `json:"unauthenticatedClientAction,omitempty"`
}

// The configuration settings of the HTTP requests for authentication and authorization requests made against ContainerApp
// Service Authentication/Authorization.
type HttpSettings struct {
	// ForwardProxy: The configuration settings of a forward proxy used to make the requests.
	ForwardProxy *ForwardProxy `json:"forwardProxy,omitempty"`

	// RequireHttps: <code>false</code> if the authentication/authorization responses not having the HTTPS scheme are
	// permissible; otherwise, <code>true</code>.
	RequireHttps *bool `json:"requireHttps,omitempty"`

	// Routes: The configuration settings of the paths HTTP requests.
	Routes *HttpSettingsRoutes `json:"routes,omitempty"`
}

// The configuration settings of each of the identity providers used to configure ContainerApp Service
// Authentication/Authorization.
type IdentityProviders struct {
	// Apple: The configuration settings of the Apple provider.
	Apple *Apple `json:"apple,omitempty"`

	// AzureActiveDirectory: The configuration settings of the Azure Active directory provider.
	AzureActiveDirectory *AzureActiveDirectory `json:"azureActiveDirectory,omitempty"`

	// AzureStaticWebApps: The configuration settings of the Azure Static Web Apps provider.
	AzureStaticWebApps *AzureStaticWebApps `json:"azureStaticWebApps,omitempty"`

	// CustomOpenIdConnectProviders: The map of the name of the alias of each custom Open ID Connect provider to the
	// configuration settings of the custom Open ID Connect provider.
	CustomOpenIdConnectProviders map[string]CustomOpenIdConnectProvider `json:"customOpenIdConnectProviders,omitempty"`

	// Facebook: The configuration settings of the Facebook provider.
	Facebook *Facebook `json:"facebook,omitempty"`

	// GitHub: The configuration settings of the GitHub provider.
	GitHub *GitHub `json:"gitHub,omitempty"`

	// Google: The configuration settings of the Google provider.
	Google *Google `json:"google,omitempty"`

	// Twitter: The configuration settings of the Twitter provider.
	Twitter *Twitter `json:"twitter,omitempty"`
}

// The configuration settings of the login flow of users using ContainerApp Service Authentication/Authorization.
type Login struct {
	// AllowedExternalRedirectUrls: External URLs that can be redirected to as part of logging in or logging out of the app.
	// Note that the query string part of the URL is ignored.
	// This is an advanced setting typically only needed by Windows Store application backends.
	// Note that URLs within the current domain are always implicitly allowed.
	AllowedExternalRedirectUrls []string `json:"allowedExternalRedirectUrls,omitempty"`

	// CookieExpiration: The configuration settings of the session cookie's expiration.
	CookieExpiration *CookieExpiration `json:"cookieExpiration,omitempty"`

	// Nonce: The configuration settings of the nonce used in the login flow.
	Nonce *Nonce `json:"nonce,omitempty"`

	// PreserveUrlFragmentsForLogins: <code>true</code> if the fragments from the request are preserved after the login request
	// is made; otherwise, <code>false</code>.
	PreserveUrlFragmentsForLogins *bool `json:"preserveUrlFragmentsForLogins,omitempty"`

	// Routes: The routes that specify the endpoints used for login and logout requests.
	Routes *LoginRoutes `json:"routes,omitempty"`

	// TokenStore: The configuration settings of the token store.
	TokenStore *TokenStore `json:"tokenStore,omitempty"`
}

// The configuration settings of the Apple provider.
type Apple struct {
	// Enabled: <code>false</code> if the Apple provider should not be enabled despite the set registration; otherwise,
	// <code>true</code>.
	Enabled *bool `json:"enabled,omitempty"`

	// Login: The configuration settings of the login flow.
	Login *LoginScopes `json:"login,omitempty"`

	// Registration: The configuration settings of the Apple registration.
	Registration *AppleRegistration `json:"registration,omitempty"`
}

// The configuration settings of the Azure Active directory provider.
type AzureActiveDirectory struct {
	// Enabled: <code>false</code> if the Azure Active Directory provider should not be enabled despite the set registration;
	// otherwise, <code>true</code>.
	Enabled *bool `json:"enabled,omitempty"`

	// IsAutoProvisioned: Gets a value indicating whether the Azure AD configuration was auto-provisioned using 1st party
	// tooling.
	// This is an internal flag primarily intended to support the Azure Management Portal. Users should not
	// read or write to this property.
	IsAutoProvisioned *bool `json:"isAutoProvisioned,omitempty"`

	// Login: The configuration settings of the Azure Active Directory login flow.
	Login *AzureActiveDirectoryLogin `json:"login,omitempty"`

	// Registration: The configuration settings of the Azure Active Directory app registration.
	Registration *AzureActiveDirectoryRegistration `json:"registration,omitempty"`

	// Validation: The configuration settings of the Azure Active Directory token validation flow.
	Validation *AzureActiveDirectoryValidation `json:"validation,omitempty"`
}

// The configuration settings of the Azure Static Web Apps provider.
type AzureStaticWebApps struct {
	// Enabled: <code>false</code> if the Azure Static Web Apps provider should not be enabled despite the set registration;
	// otherwise, <code>true</code>.
	Enabled *bool `json:"enabled,omitempty"`

	// Registration: The configuration settings of the Azure Static Web Apps registration.
	Registration *AzureStaticWebAppsRegistration `json:"registration,omitempty"`
}

// The configuration settings of the session cookie's expiration.
type CookieExpiration struct {
	// Convention: The convention used when determining the session cookie's expiration.
	Convention *CookieExpiration_Convention `json:"convention,omitempty"`

	// TimeToExpiration: The time after the request is made when the session cookie should expire.
	TimeToExpiration *string `json:"timeToExpiration,omitempty"`
}

// The configuration settings of the custom Open ID Connect provider.
type CustomOpenIdConnectProvider struct {
	// Enabled: <code>false</code> if the custom Open ID provider provider should not be enabled; otherwise, <code>true</code>.
	Enabled *bool `json:"enabled,omitempty"`

	// Login: The configuration settings of the login flow of the custom Open ID Connect provider.
	Login *OpenIdConnectLogin `json:"login,omitempty"`

	// Registration: The configuration settings of the app registration for the custom Open ID Connect provider.
	Registration *OpenIdConnectRegistration `json:"registration,omitempty"`
}

// The configuration settings of the Facebook provider.
type Facebook struct {
	// Enabled: <code>false</code> if the Facebook provider should not be enabled despite the set registration; otherwise,
	// <code>true</code>.
	Enabled *bool `json:"enabled,omitempty"`

	// GraphApiVersion: The version of the Facebook api to be used while logging in.
	GraphApiVersion *string `json:"graphApiVersion,omitempty"`

	// Login: The configuration settings of the login flow.
	Login *LoginScopes `json:"login,omitempty"`

	// Registration: The configuration settings of the app registration for the Facebook provider.
	Registration *AppRegistration `json:"registration,omitempty"`
}

// The configuration settings of a forward proxy used to make the requests.
type ForwardProxy struct {
	// Convention: The convention used to determine the url of the request made.
	Convention *ForwardProxy_Convention `json:"convention,omitempty"`

	// CustomHostHeaderName: The name of the header containing the host of the request.
	CustomHostHeaderName *string `json:"customHostHeaderName,omitempty"`

	// CustomProtoHeaderName: The name of the header containing the scheme of the request.
	CustomProtoHeaderName *string `json:"customProtoHeaderName,omitempty"`
}

// The configuration settings of the GitHub provider.
type GitHub struct {
	// Enabled: <code>false</code> if the GitHub provider should not be enabled despite the set registration; otherwise,
	// <code>true</code>.
	Enabled *bool `json:"enabled,omitempty"`

	// Login: The configuration settings of the login flow.
	Login *LoginScopes `json:"login,omitempty"`

	// Registration: The configuration settings of the app registration for the GitHub provider.
	Registration *ClientRegistration `json:"registration,omitempty"`
}

// +kubebuilder:validation:Enum={"AllowAnonymous","RedirectToLoginPage","Return401","Return403"}
type GlobalValidation_UnauthenticatedClientAction string

const (
	GlobalValidation_UnauthenticatedClientAction_AllowAnonymous      = GlobalValidation_UnauthenticatedClientAction("AllowAnonymous")
	GlobalValidation_UnauthenticatedClientAction_RedirectToLoginPage = GlobalValidation_UnauthenticatedClientAction("RedirectToLoginPage")
	GlobalValidation_UnauthenticatedClientAction_Return401           = GlobalValidation_UnauthenticatedClientAction("Return401")
	GlobalValidation_UnauthenticatedClientAction_Return403           = GlobalValidation_UnauthenticatedClientAction("Return403")
)

// Mapping from string to GlobalValidation_UnauthenticatedClientAction
var globalValidation_UnauthenticatedClientAction_Values = map[string]GlobalValidation_UnauthenticatedClientAction{
	"allowanonymous":      GlobalValidation_UnauthenticatedClientAction_AllowAnonymous,
	"redirecttologinpage": GlobalValidation_UnauthenticatedClientAction_RedirectToLoginPage,
	"return401":           GlobalValidation_UnauthenticatedClientAction_Return401,
	"return403":           GlobalValidation_UnauthenticatedClientAction_Return403,
}

// The configuration settings of the Google provider.
type Google struct {
	// Enabled: <code>false</code> if the Google provider should not be enabled despite the set registration; otherwise,
	// <code>true</code>.
	Enabled *bool `json:"enabled,omitempty"`

	// Login: The configuration settings of the login flow.
	Login *LoginScopes `json:"login,omitempty"`

	// Registration: The configuration settings of the app registration for the Google provider.
	Registration *ClientRegistration `json:"registration,omitempty"`

	// Validation: The configuration settings of the Azure Active Directory token validation flow.
	Validation *AllowedAudiencesValidation `json:"validation,omitempty"`
}

// The configuration settings of the paths HTTP requests.
type HttpSettingsRoutes struct {
	// ApiPrefix: The prefix that should precede all the authentication/authorization paths.
	ApiPrefix *string `json:"apiPrefix,omitempty"`
}

// The routes that specify the endpoints used for login and logout requests.
type LoginRoutes struct {
	// LogoutEndpoint: The endpoint at which a logout request should be made.
	LogoutEndpoint *string `json:"logoutEndpoint,omitempty"`
}

// The configuration settings of the nonce used in the login flow.
type Nonce struct {
	// NonceExpirationInterval: The time after the request is made when the nonce should expire.
	NonceExpirationInterval *string `json:"nonceExpirationInterval,omitempty"`

	// ValidateNonce: <code>false</code> if the nonce should not be validated while completing the login flow; otherwise,
	// <code>true</code>.
	ValidateNonce *bool `json:"validateNonce,omitempty"`
}

// The configuration settings of the token store.
type TokenStore struct {
	// AzureBlobStorage: The configuration settings of the storage of the tokens if blob storage is used.
	AzureBlobStorage *BlobStorageTokenStore `json:"azureBlobStorage,omitempty"`

	// Enabled: <code>true</code> to durably store platform-specific security tokens that are obtained during login flows;
	// otherwise, <code>false</code>.
	// The default is <code>false</code>.
	Enabled *bool `json:"enabled,omitempty"`

	// TokenRefreshExtensionHours: The number of hours after session token expiration that a session token can be used to
	// call the token refresh API. The default is 72 hours.
	TokenRefreshExtensionHours *float64 `json:"tokenRefreshExtensionHours,omitempty"`
}

// The configuration settings of the Twitter provider.
type Twitter struct {
	// Enabled: <code>false</code> if the Twitter provider should not be enabled despite the set registration; otherwise,
	// <code>true</code>.
	Enabled *bool `json:"enabled,omitempty"`

	// Registration: The configuration settings of the app registration for the Twitter provider.
	Registration *TwitterRegistration `json:"registration,omitempty"`
}

// The configuration settings of the Allowed Audiences validation flow.
type AllowedAudiencesValidation struct {
	// AllowedAudiences: The configuration settings of the allowed list of audiences from which to validate the JWT token.
	AllowedAudiences []string `json:"allowedAudiences,omitempty"`
}

// The configuration settings of the registration for the Apple provider
type AppleRegistration struct {
	// ClientId: The Client ID of the app used for login.
	ClientId *string `json:"clientId,omitempty"`

	// ClientSecretSettingName: The app setting name that contains the client secret.
	ClientSecretSettingName *string `json:"clientSecretSettingName,omitempty"`
}

// The configuration settings of the app registration for providers that have app ids and app secrets
type AppRegistration struct {
	// AppId: The App ID of the app used for login.
	AppId *string `json:"appId,omitempty"`

	// AppSecretSettingName: The app setting name that contains the app secret.
	AppSecretSettingName *string `json:"appSecretSettingName,omitempty"`
}

// The configuration settings of the Azure Active Directory login flow.
type AzureActiveDirectoryLogin struct {
	// DisableWWWAuthenticate: <code>true</code> if the www-authenticate provider should be omitted from the request;
	// otherwise, <code>false</code>.
	DisableWWWAuthenticate *bool `json:"disableWWWAuthenticate,omitempty"`

	// LoginParameters: Login parameters to send to the OpenID Connect authorization endpoint when
	// a user logs in. Each parameter must be in the form "key=value".
	LoginParameters []string `json:"loginParameters,omitempty"`
}

// The configuration settings of the Azure Active Directory app registration.
type AzureActiveDirectoryRegistration struct {
	// ClientId: The Client ID of this relying party application, known as the client_id.
	// This setting is required for enabling OpenID Connection authentication with Azure Active Directory or
	// other 3rd party OpenID Connect providers.
	// More information on OpenID Connect: http://openid.net/specs/openid-connect-core-1_0.html
	ClientId *string `json:"clientId,omitempty"`

	// ClientSecretCertificateIssuer: An alternative to the client secret thumbprint, that is the issuer of a certificate used
	// for signing purposes. This property acts as
	// a replacement for the Client Secret Certificate Thumbprint. It is also optional.
	ClientSecretCertificateIssuer *string `json:"clientSecretCertificateIssuer,omitempty"`

	// ClientSecretCertificateSubjectAlternativeName: An alternative to the client secret thumbprint, that is the subject
	// alternative name of a certificate used for signing purposes. This property acts as
	// a replacement for the Client Secret Certificate Thumbprint. It is also optional.
	ClientSecretCertificateSubjectAlternativeName *string `json:"clientSecretCertificateSubjectAlternativeName,omitempty"`

	// ClientSecretCertificateThumbprint: An alternative to the client secret, that is the thumbprint of a certificate used for
	// signing purposes. This property acts as
	// a replacement for the Client Secret. It is also optional.
	ClientSecretCertificateThumbprint *string `json:"clientSecretCertificateThumbprint,omitempty"`

	// ClientSecretSettingName: The app setting name that contains the client secret of the relying party application.
	ClientSecretSettingName *string `json:"clientSecretSettingName,omitempty"`

	// OpenIdIssuer: The OpenID Connect Issuer URI that represents the entity which issues access tokens for this application.
	// When using Azure Active Directory, this value is the URI of the directory tenant, e.g.
	// `https://login.microsoftonline.com/v2.0/{tenant-guid}/`.
	// This URI is a case-sensitive identifier for the token issuer.
	// More information on OpenID Connect Discovery: http://openid.net/specs/openid-connect-discovery-1_0.html
	OpenIdIssuer *string `json:"openIdIssuer,omitempty"`
}

// The configuration settings of the Azure Active Directory token validation flow.
type AzureActiveDirectoryValidation struct {
	// AllowedAudiences: The list of audiences that can make successful authentication/authorization requests.
	AllowedAudiences []string `json:"allowedAudiences,omitempty"`

	// DefaultAuthorizationPolicy: The configuration settings of the default authorization policy.
	DefaultAuthorizationPolicy *DefaultAuthorizationPolicy `json:"defaultAuthorizationPolicy,omitempty"`

	// JwtClaimChecks: The configuration settings of the checks that should be made while validating the JWT Claims.
	JwtClaimChecks *JwtClaimChecks `json:"jwtClaimChecks,omitempty"`
}

// The configuration settings of the registration for the Azure Static Web Apps provider
type AzureStaticWebAppsRegistration struct {
	// ClientId: The Client ID of the app used for login.
	ClientId *string `json:"clientId,omitempty"`
}

// The configuration settings of the storage of the tokens if blob storage is used.
type BlobStorageTokenStore struct {
	// SasUrlSettingName: The name of the app secrets containing the SAS URL of the blob storage containing the tokens.
	SasUrlSettingName *string `json:"sasUrlSettingName,omitempty"`
}

// The configuration settings of the app registration for providers that have client ids and client secrets
type ClientRegistration struct {
	// ClientId: The Client ID of the app used for login.
	ClientId *string `json:"clientId,omitempty"`

	// ClientSecretSettingName: The app setting name that contains the client secret.
	ClientSecretSettingName *string `json:"clientSecretSettingName,omitempty"`
}

// +kubebuilder:validation:Enum={"FixedTime","IdentityProviderDerived"}
type CookieExpiration_Convention string

const (
	CookieExpiration_Convention_FixedTime               = CookieExpiration_Convention("FixedTime")
	CookieExpiration_Convention_IdentityProviderDerived = CookieExpiration_Convention("IdentityProviderDerived")
)

// Mapping from string to CookieExpiration_Convention
var cookieExpiration_Convention_Values = map[string]CookieExpiration_Convention{
	"fixedtime":               CookieExpiration_Convention_FixedTime,
	"identityproviderderived": CookieExpiration_Convention_IdentityProviderDerived,
}

// +kubebuilder:validation:Enum={"Custom","NoProxy","Standard"}
type ForwardProxy_Convention string

const (
	ForwardProxy_Convention_Custom   = ForwardProxy_Convention("Custom")
	ForwardProxy_Convention_NoProxy  = ForwardProxy_Convention("NoProxy")
	ForwardProxy_Convention_Standard = ForwardProxy_Convention("Standard")
)

// Mapping from string to ForwardProxy_Convention
var forwardProxy_Convention_Values = map[string]ForwardProxy_Convention{
	"custom":   ForwardProxy_Convention_Custom,
	"noproxy":  ForwardProxy_Convention_NoProxy,
	"standard": ForwardProxy_Convention_Standard,
}

// The configuration settings of the login flow, including the scopes that should be requested.
type LoginScopes struct {
	// Scopes: A list of the scopes that should be requested while authenticating.
	Scopes []string `json:"scopes,omitempty"`
}

// The configuration settings of the login flow of the custom Open ID Connect provider.
type OpenIdConnectLogin struct {
	// NameClaimType: The name of the claim that contains the users name.
	NameClaimType *string `json:"nameClaimType,omitempty"`

	// Scopes: A list of the scopes that should be requested while authenticating.
	Scopes []string `json:"scopes,omitempty"`
}

// The configuration settings of the app registration for the custom Open ID Connect provider.
type OpenIdConnectRegistration struct {
	// ClientCredential: The authentication credentials of the custom Open ID Connect provider.
	ClientCredential *OpenIdConnectClientCredential `json:"clientCredential,omitempty"`

	// ClientId: The client id of the custom Open ID Connect provider.
	ClientId *string `json:"clientId,omitempty"`

	// OpenIdConnectConfiguration: The configuration settings of the endpoints used for the custom Open ID Connect provider.
	OpenIdConnectConfiguration *OpenIdConnectConfig `json:"openIdConnectConfiguration,omitempty"`
}

// The configuration settings of the app registration for the Twitter provider.
type TwitterRegistration struct {
	// ConsumerKey: The OAuth 1.0a consumer key of the Twitter application used for sign-in.
	// This setting is required for enabling Twitter Sign-In.
	// Twitter Sign-In documentation: https://dev.twitter.com/web/sign-in
	ConsumerKey *string `json:"consumerKey,omitempty"`

	// ConsumerSecretSettingName: The app setting name that contains the OAuth 1.0a consumer secret of the Twitter
	// application used for sign-in.
	ConsumerSecretSettingName *string `json:"consumerSecretSettingName,omitempty"`
}

// The configuration settings of the Azure Active Directory default authorization policy.
type DefaultAuthorizationPolicy struct {
	// AllowedApplications: The configuration settings of the Azure Active Directory allowed applications.
	AllowedApplications []string `json:"allowedApplications,omitempty"`

	// AllowedPrincipals: The configuration settings of the Azure Active Directory allowed principals.
	AllowedPrincipals *AllowedPrincipals `json:"allowedPrincipals,omitempty"`
}

// The configuration settings of the checks that should be made while validating the JWT Claims.
type JwtClaimChecks struct {
	// AllowedClientApplications: The list of the allowed client applications.
	AllowedClientApplications []string `json:"allowedClientApplications,omitempty"`

	// AllowedGroups: The list of the allowed groups.
	AllowedGroups []string `json:"allowedGroups,omitempty"`
}

// The authentication client credentials of the custom Open ID Connect provider.
type OpenIdConnectClientCredential struct {
	// ClientSecretSettingName: The app setting that contains the client secret for the custom Open ID Connect provider.
	ClientSecretSettingName *string `json:"clientSecretSettingName,omitempty"`

	// Method: The method that should be used to authenticate the user.
	Method *OpenIdConnectClientCredential_Method `json:"method,omitempty"`
}

// The configuration settings of the endpoints used for the custom Open ID Connect provider.
type OpenIdConnectConfig struct {
	// AuthorizationEndpoint: The endpoint to be used to make an authorization request.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty"`

	// CertificationUri: The endpoint that provides the keys necessary to validate the token.
	CertificationUri *string `json:"certificationUri,omitempty"`

	// Issuer: The endpoint that issues the token.
	Issuer *string `json:"issuer,omitempty"`

	// TokenEndpoint: The endpoint to be used to request a token.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`

	// WellKnownOpenIdConfiguration: The endpoint that contains all the configuration endpoints for the provider.
	WellKnownOpenIdConfiguration *string `json:"wellKnownOpenIdConfiguration,omitempty"`
}

// The configuration settings of the Azure Active Directory allowed principals.
type AllowedPrincipals struct {
	// Groups: The list of the allowed groups.
	Groups []string `json:"groups,omitempty"`

	// Identities: The list of the allowed identities.
	Identities []string `json:"identities,omitempty"`
}

// +kubebuilder:validation:Enum={"ClientSecretPost"}
type OpenIdConnectClientCredential_Method string

const OpenIdConnectClientCredential_Method_ClientSecretPost = OpenIdConnectClientCredential_Method("ClientSecretPost")

// Mapping from string to OpenIdConnectClientCredential_Method
var openIdConnectClientCredential_Method_Values = map[string]OpenIdConnectClientCredential_Method{
	"clientsecretpost": OpenIdConnectClientCredential_Method_ClientSecretPost,
}
