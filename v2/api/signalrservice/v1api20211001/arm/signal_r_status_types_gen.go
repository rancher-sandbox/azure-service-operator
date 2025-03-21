// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type SignalR_STATUS struct {
	// Id: Fully qualified resource Id for the resource.
	Id *string `json:"id,omitempty"`

	// Identity: A class represent managed identities used for request and response
	Identity *ManagedIdentity_STATUS `json:"identity,omitempty"`

	// Kind: The kind of the service, it can be SignalR or RawWebSockets
	Kind *ServiceKind_STATUS `json:"kind,omitempty"`

	// Location: The GEO location of the resource. e.g. West US | East US | North Central US | South Central US.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource.
	Name *string `json:"name,omitempty"`

	// Properties: A class that describes the properties of the resource
	Properties *SignalRProperties_STATUS `json:"properties,omitempty"`

	// Sku: The billing information of the resource.
	Sku *ResourceSku_STATUS `json:"sku,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type *string `json:"type,omitempty"`
}

// A class represent managed identities used for request and response
type ManagedIdentity_STATUS struct {
	// PrincipalId: Get the principal id for the system assigned identity.
	// Only be used in response.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: Get the tenant id for the system assigned identity.
	// Only be used in response
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Represents the identity type: systemAssigned, userAssigned, None
	Type *ManagedIdentityType_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: Get or set the user assigned identities
	UserAssignedIdentities map[string]UserAssignedIdentityProperty_STATUS `json:"userAssignedIdentities,omitempty"`
}

// The billing information of the resource.
type ResourceSku_STATUS struct {
	// Capacity: Optional, integer. The unit count of the resource. 1 by default.
	// If present, following values are allowed:
	// Free: 1
	// Standard: 1,2,5,10,20,50,100
	Capacity *int `json:"capacity,omitempty"`

	// Family: Not used. Retained for future use.
	Family *string `json:"family,omitempty"`

	// Name: The name of the SKU. Required.
	// Allowed values: Standard_S1, Free_F1
	Name *string `json:"name,omitempty"`

	// Size: Not used. Retained for future use.
	Size *string `json:"size,omitempty"`

	// Tier: Optional tier of this particular SKU. 'Standard' or 'Free'.
	// `Basic` is deprecated, use `Standard` instead.
	Tier *SignalRSkuTier_STATUS `json:"tier,omitempty"`
}

// The kind of the service, it can be SignalR or RawWebSockets
type ServiceKind_STATUS string

const (
	ServiceKind_STATUS_RawWebSockets = ServiceKind_STATUS("RawWebSockets")
	ServiceKind_STATUS_SignalR       = ServiceKind_STATUS("SignalR")
)

// Mapping from string to ServiceKind_STATUS
var serviceKind_STATUS_Values = map[string]ServiceKind_STATUS{
	"rawwebsockets": ServiceKind_STATUS_RawWebSockets,
	"signalr":       ServiceKind_STATUS_SignalR,
}

// A class that describes the properties of the resource
type SignalRProperties_STATUS struct {
	// Cors: Cross-Origin Resource Sharing (CORS) settings.
	Cors *SignalRCorsSettings_STATUS `json:"cors,omitempty"`

	// DisableAadAuth: DisableLocalAuth
	// Enable or disable aad auth
	// When set as true, connection with AuthType=aad won't work.
	DisableAadAuth *bool `json:"disableAadAuth,omitempty"`

	// DisableLocalAuth: DisableLocalAuth
	// Enable or disable local auth with AccessKey
	// When set as true, connection with AccessKey=xxx won't work.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// ExternalIP: The publicly accessible IP of the resource.
	ExternalIP *string `json:"externalIP,omitempty"`

	// Features: List of the featureFlags.
	// FeatureFlags that are not included in the parameters for the update operation will not be modified.
	// And the response will only include featureFlags that are explicitly set.
	// When a featureFlag is not explicitly set, its globally default value will be used
	// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
	Features []SignalRFeature_STATUS `json:"features,omitempty"`

	// HostName: FQDN of the service instance.
	HostName *string `json:"hostName,omitempty"`

	// HostNamePrefix: Deprecated.
	HostNamePrefix *string `json:"hostNamePrefix,omitempty"`

	// NetworkACLs: Network ACLs for the resource
	NetworkACLs *SignalRNetworkACLs_STATUS `json:"networkACLs,omitempty"`

	// PrivateEndpointConnections: Private endpoint connections to the resource.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Provisioning state of the resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: Enable or disable public network access. Default to "Enabled".
	// When it's Enabled, network ACLs still apply.
	// When it's Disabled, public network access is always disabled no matter what you set in network ACLs.
	PublicNetworkAccess *string `json:"publicNetworkAccess,omitempty"`

	// PublicPort: The publicly accessible port of the resource which is designed for browser/client side usage.
	PublicPort *int `json:"publicPort,omitempty"`

	// ResourceLogConfiguration: Resource log configuration of a Microsoft.SignalRService resource.
	ResourceLogConfiguration *ResourceLogConfiguration_STATUS `json:"resourceLogConfiguration,omitempty"`

	// ServerPort: The publicly accessible port of the resource which is designed for customer server side usage.
	ServerPort *int `json:"serverPort,omitempty"`

	// SharedPrivateLinkResources: The list of shared private link resources.
	SharedPrivateLinkResources []SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded `json:"sharedPrivateLinkResources,omitempty"`

	// Tls: TLS settings for the resource
	Tls *SignalRTlsSettings_STATUS `json:"tls,omitempty"`

	// Upstream: The settings for the Upstream when the service is in server-less mode.
	Upstream *ServerlessUpstreamSettings_STATUS `json:"upstream,omitempty"`

	// Version: Version of the resource. Probably you need the same or higher version of client SDKs.
	Version *string `json:"version,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// Represents the identity type: systemAssigned, userAssigned, None
type ManagedIdentityType_STATUS string

const (
	ManagedIdentityType_STATUS_None           = ManagedIdentityType_STATUS("None")
	ManagedIdentityType_STATUS_SystemAssigned = ManagedIdentityType_STATUS("SystemAssigned")
	ManagedIdentityType_STATUS_UserAssigned   = ManagedIdentityType_STATUS("UserAssigned")
)

// Mapping from string to ManagedIdentityType_STATUS
var managedIdentityType_STATUS_Values = map[string]ManagedIdentityType_STATUS{
	"none":           ManagedIdentityType_STATUS_None,
	"systemassigned": ManagedIdentityType_STATUS_SystemAssigned,
	"userassigned":   ManagedIdentityType_STATUS_UserAssigned,
}

// A private endpoint connection to an azure resource
type PrivateEndpointConnection_STATUS_SignalR_SubResourceEmbedded struct {
	// Id: Fully qualified resource Id for the resource.
	Id *string `json:"id,omitempty"`
}

// Provisioning state of the resource.
type ProvisioningState_STATUS string

const (
	ProvisioningState_STATUS_Canceled  = ProvisioningState_STATUS("Canceled")
	ProvisioningState_STATUS_Creating  = ProvisioningState_STATUS("Creating")
	ProvisioningState_STATUS_Deleting  = ProvisioningState_STATUS("Deleting")
	ProvisioningState_STATUS_Failed    = ProvisioningState_STATUS("Failed")
	ProvisioningState_STATUS_Moving    = ProvisioningState_STATUS("Moving")
	ProvisioningState_STATUS_Running   = ProvisioningState_STATUS("Running")
	ProvisioningState_STATUS_Succeeded = ProvisioningState_STATUS("Succeeded")
	ProvisioningState_STATUS_Unknown   = ProvisioningState_STATUS("Unknown")
	ProvisioningState_STATUS_Updating  = ProvisioningState_STATUS("Updating")
)

// Mapping from string to ProvisioningState_STATUS
var provisioningState_STATUS_Values = map[string]ProvisioningState_STATUS{
	"canceled":  ProvisioningState_STATUS_Canceled,
	"creating":  ProvisioningState_STATUS_Creating,
	"deleting":  ProvisioningState_STATUS_Deleting,
	"failed":    ProvisioningState_STATUS_Failed,
	"moving":    ProvisioningState_STATUS_Moving,
	"running":   ProvisioningState_STATUS_Running,
	"succeeded": ProvisioningState_STATUS_Succeeded,
	"unknown":   ProvisioningState_STATUS_Unknown,
	"updating":  ProvisioningState_STATUS_Updating,
}

// Resource log configuration of a Microsoft.SignalRService resource.
type ResourceLogConfiguration_STATUS struct {
	// Categories: Gets or sets the list of category configurations.
	Categories []ResourceLogCategory_STATUS `json:"categories,omitempty"`
}

// The settings for the Upstream when the service is in server-less mode.
type ServerlessUpstreamSettings_STATUS struct {
	// Templates: Gets or sets the list of Upstream URL templates. Order matters, and the first matching template takes effects.
	Templates []UpstreamTemplate_STATUS `json:"templates,omitempty"`
}

// Describes a Shared Private Link Resource
type SharedPrivateLinkResource_STATUS_SignalR_SubResourceEmbedded struct {
	// Id: Fully qualified resource Id for the resource.
	Id *string `json:"id,omitempty"`
}

// Cross-Origin Resource Sharing (CORS) settings.
type SignalRCorsSettings_STATUS struct {
	// AllowedOrigins: Gets or sets the list of origins that should be allowed to make cross-origin calls (for example:
	// http://example.com:12345). Use "*" to allow all. If omitted, allow all by default.
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
}

// Feature of a resource, which controls the runtime behavior.
type SignalRFeature_STATUS struct {
	// Flag: FeatureFlags is the supported features of Azure SignalR service.
	// - ServiceMode: Flag for backend server for SignalR  service. Values allowed: "Default": have your own backend server;
	// "Serverless": your application doesn't have a backend  server; "Classic": for backward compatibility. Support both
	// Default and Serverless mode but not recommended;  "PredefinedOnly": for future use.
	// - EnableConnectivityLogs: "true"/"false", to enable/disable the connectivity log  category respectively.
	// - EnableMessagingLogs: "true"/"false", to enable/disable the connectivity log category  respectively.
	// - EnableLiveTrace: Live Trace allows you to know what's happening inside Azure SignalR service, it will  give you live
	// traces in real time, it will be helpful when you developing your own Azure SignalR based web application  or
	// self-troubleshooting some issues. Please note that live traces are counted as outbound messages that will be charged.
	// Values allowed: "true"/"false", to enable/disable live trace feature.
	Flag *FeatureFlags_STATUS `json:"flag,omitempty"`

	// Properties: Optional properties related to this feature.
	Properties map[string]string `json:"properties,omitempty"`

	// Value: Value of the feature flag. See Azure SignalR service document https://docs.microsoft.com/azure/azure-signalr/ for
	// allowed values.
	Value *string `json:"value,omitempty"`
}

// Network ACLs for the resource
type SignalRNetworkACLs_STATUS struct {
	// DefaultAction: Azure Networking ACL Action.
	DefaultAction *ACLAction_STATUS `json:"defaultAction,omitempty"`

	// PrivateEndpoints: ACLs for requests from private endpoints
	PrivateEndpoints []PrivateEndpointACL_STATUS `json:"privateEndpoints,omitempty"`

	// PublicNetwork: Network ACL
	PublicNetwork *NetworkACL_STATUS `json:"publicNetwork,omitempty"`
}

// Optional tier of this particular SKU. 'Standard' or 'Free'.
// `Basic` is deprecated, use `Standard` instead.
type SignalRSkuTier_STATUS string

const (
	SignalRSkuTier_STATUS_Basic    = SignalRSkuTier_STATUS("Basic")
	SignalRSkuTier_STATUS_Free     = SignalRSkuTier_STATUS("Free")
	SignalRSkuTier_STATUS_Premium  = SignalRSkuTier_STATUS("Premium")
	SignalRSkuTier_STATUS_Standard = SignalRSkuTier_STATUS("Standard")
)

// Mapping from string to SignalRSkuTier_STATUS
var signalRSkuTier_STATUS_Values = map[string]SignalRSkuTier_STATUS{
	"basic":    SignalRSkuTier_STATUS_Basic,
	"free":     SignalRSkuTier_STATUS_Free,
	"premium":  SignalRSkuTier_STATUS_Premium,
	"standard": SignalRSkuTier_STATUS_Standard,
}

// TLS settings for the resource
type SignalRTlsSettings_STATUS struct {
	// ClientCertEnabled: Request client certificate during TLS handshake if enabled
	ClientCertEnabled *bool `json:"clientCertEnabled,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}

// Properties of user assigned identity.
type UserAssignedIdentityProperty_STATUS struct {
	// ClientId: Get the client id for the user assigned identity
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: Get the principal id for the user assigned identity
	PrincipalId *string `json:"principalId,omitempty"`
}

// Azure Networking ACL Action.
type ACLAction_STATUS string

const (
	ACLAction_STATUS_Allow = ACLAction_STATUS("Allow")
	ACLAction_STATUS_Deny  = ACLAction_STATUS("Deny")
)

// Mapping from string to ACLAction_STATUS
var aCLAction_STATUS_Values = map[string]ACLAction_STATUS{
	"allow": ACLAction_STATUS_Allow,
	"deny":  ACLAction_STATUS_Deny,
}

// FeatureFlags is the supported features of Azure SignalR service.
// - ServiceMode: Flag for backend server for SignalR
// service. Values allowed: "Default": have your own backend server; "Serverless": your application doesn't have a backend
// server; "Classic": for backward compatibility. Support both Default and Serverless mode but not recommended;
// "PredefinedOnly": for future use.
// - EnableConnectivityLogs: "true"/"false", to enable/disable the connectivity log
// category respectively.
// - EnableMessagingLogs: "true"/"false", to enable/disable the connectivity log category
// respectively.
// - EnableLiveTrace: Live Trace allows you to know what's happening inside Azure SignalR service, it will
// give you live traces in real time, it will be helpful when you developing your own Azure SignalR based web application
// or self-troubleshooting some issues. Please note that live traces are counted as outbound messages that will be charged.
// Values allowed: "true"/"false", to enable/disable live trace feature.
type FeatureFlags_STATUS string

const (
	FeatureFlags_STATUS_EnableConnectivityLogs = FeatureFlags_STATUS("EnableConnectivityLogs")
	FeatureFlags_STATUS_EnableLiveTrace        = FeatureFlags_STATUS("EnableLiveTrace")
	FeatureFlags_STATUS_EnableMessagingLogs    = FeatureFlags_STATUS("EnableMessagingLogs")
	FeatureFlags_STATUS_ServiceMode            = FeatureFlags_STATUS("ServiceMode")
)

// Mapping from string to FeatureFlags_STATUS
var featureFlags_STATUS_Values = map[string]FeatureFlags_STATUS{
	"enableconnectivitylogs": FeatureFlags_STATUS_EnableConnectivityLogs,
	"enablelivetrace":        FeatureFlags_STATUS_EnableLiveTrace,
	"enablemessaginglogs":    FeatureFlags_STATUS_EnableMessagingLogs,
	"servicemode":            FeatureFlags_STATUS_ServiceMode,
}

// Network ACL
type NetworkACL_STATUS struct {
	// Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Allow []SignalRRequestType_STATUS `json:"allow,omitempty"`

	// Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Deny []SignalRRequestType_STATUS `json:"deny,omitempty"`
}

// ACL for a private endpoint
type PrivateEndpointACL_STATUS struct {
	// Allow: Allowed request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Allow []SignalRRequestType_STATUS `json:"allow,omitempty"`

	// Deny: Denied request types. The value can be one or more of: ClientConnection, ServerConnection, RESTAPI.
	Deny []SignalRRequestType_STATUS `json:"deny,omitempty"`

	// Name: Name of the private endpoint connection
	Name *string `json:"name,omitempty"`
}

// Resource log category configuration of a Microsoft.SignalRService resource.
type ResourceLogCategory_STATUS struct {
	// Enabled: Indicates whether or the resource log category is enabled.
	// Available values: true, false.
	// Case insensitive.
	Enabled *string `json:"enabled,omitempty"`

	// Name: Gets or sets the resource log category's name.
	// Available values: ConnectivityLogs, MessagingLogs.
	// Case insensitive.
	Name *string `json:"name,omitempty"`
}

// Upstream template item settings. It defines the Upstream URL of the incoming requests.
// The template defines the pattern
// of the event, the hub or the category of the incoming request that matches current URL template.
type UpstreamTemplate_STATUS struct {
	// Auth: Upstream auth settings. If not set, no auth is used for upstream messages.
	Auth *UpstreamAuthSettings_STATUS `json:"auth,omitempty"`

	// CategoryPattern: Gets or sets the matching pattern for category names. If not set, it matches any category.
	// There are 3 kind of patterns supported:
	// 1. "*", it to matches any category name
	// 2. Combine multiple categories with ",", for example "connections,messages", it matches category "connections" and
	// "messages"
	// 3. The single category name, for example, "connections", it matches the category "connections"
	CategoryPattern *string `json:"categoryPattern,omitempty"`

	// EventPattern: Gets or sets the matching pattern for event names. If not set, it matches any event.
	// There are 3 kind of patterns supported:
	// 1. "*", it to matches any event name
	// 2. Combine multiple events with ",", for example "connect,disconnect", it matches event "connect" and "disconnect"
	// 3. The single event name, for example, "connect", it matches "connect"
	EventPattern *string `json:"eventPattern,omitempty"`

	// HubPattern: Gets or sets the matching pattern for hub names. If not set, it matches any hub.
	// There are 3 kind of patterns supported:
	// 1. "*", it to matches any hub name
	// 2. Combine multiple hubs with ",", for example "hub1,hub2", it matches "hub1" and "hub2"
	// 3. The single hub name, for example, "hub1", it matches "hub1"
	HubPattern *string `json:"hubPattern,omitempty"`

	// UrlTemplate: Gets or sets the Upstream URL template. You can use 3 predefined parameters {hub}, {category} {event}
	// inside the template, the value of the Upstream URL is dynamically calculated when the client request comes in.
	// For example, if the urlTemplate is `http://example.com/{hub}/api/{event}`, with a client request from hub `chat`
	// connects, it will first POST to this URL: `http://example.com/chat/api/connect`.
	UrlTemplate *string `json:"urlTemplate,omitempty"`
}

// The incoming request type to the service
type SignalRRequestType_STATUS string

const (
	SignalRRequestType_STATUS_ClientConnection = SignalRRequestType_STATUS("ClientConnection")
	SignalRRequestType_STATUS_RESTAPI          = SignalRRequestType_STATUS("RESTAPI")
	SignalRRequestType_STATUS_ServerConnection = SignalRRequestType_STATUS("ServerConnection")
	SignalRRequestType_STATUS_Trace            = SignalRRequestType_STATUS("Trace")
)

// Mapping from string to SignalRRequestType_STATUS
var signalRRequestType_STATUS_Values = map[string]SignalRRequestType_STATUS{
	"clientconnection": SignalRRequestType_STATUS_ClientConnection,
	"restapi":          SignalRRequestType_STATUS_RESTAPI,
	"serverconnection": SignalRRequestType_STATUS_ServerConnection,
	"trace":            SignalRRequestType_STATUS_Trace,
}

// Upstream auth settings. If not set, no auth is used for upstream messages.
type UpstreamAuthSettings_STATUS struct {
	// ManagedIdentity: Managed identity settings for upstream.
	ManagedIdentity *ManagedIdentitySettings_STATUS `json:"managedIdentity,omitempty"`

	// Type: Upstream auth type enum.
	Type *UpstreamAuthType_STATUS `json:"type,omitempty"`
}

// Managed identity settings for upstream.
type ManagedIdentitySettings_STATUS struct {
	// Resource: The Resource indicating the App ID URI of the target resource.
	// It also appears in the aud (audience) claim of the issued token.
	Resource *string `json:"resource,omitempty"`
}

// Upstream auth type enum.
type UpstreamAuthType_STATUS string

const (
	UpstreamAuthType_STATUS_ManagedIdentity = UpstreamAuthType_STATUS("ManagedIdentity")
	UpstreamAuthType_STATUS_None            = UpstreamAuthType_STATUS("None")
)

// Mapping from string to UpstreamAuthType_STATUS
var upstreamAuthType_STATUS_Values = map[string]UpstreamAuthType_STATUS{
	"managedidentity": UpstreamAuthType_STATUS_ManagedIdentity,
	"none":            UpstreamAuthType_STATUS_None,
}
