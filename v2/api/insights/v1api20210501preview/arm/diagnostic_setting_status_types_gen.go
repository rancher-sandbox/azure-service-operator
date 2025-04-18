// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type DiagnosticSetting_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of a Diagnostic Settings Resource.
	Properties *DiagnosticSettings_STATUS `json:"properties,omitempty"`

	// SystemData: The system metadata related to this resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The diagnostic settings.
type DiagnosticSettings_STATUS struct {
	// EventHubAuthorizationRuleId: The resource Id for the event hub authorization rule.
	EventHubAuthorizationRuleId *string `json:"eventHubAuthorizationRuleId,omitempty"`

	// EventHubName: The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName *string `json:"eventHubName,omitempty"`

	// LogAnalyticsDestinationType: A string indicating whether the export to Log Analytics should use the default destination
	// type, i.e. AzureDiagnostics, or use a destination type constructed as follows: <normalized service identity>_<normalized
	// category name>. Possible values are: Dedicated and null (null is default.)
	LogAnalyticsDestinationType *string `json:"logAnalyticsDestinationType,omitempty"`

	// Logs: The list of logs settings.
	Logs []LogSettings_STATUS `json:"logs,omitempty"`

	// MarketplacePartnerId: The full ARM resource ID of the Marketplace resource to which you would like to send Diagnostic
	// Logs.
	MarketplacePartnerId *string `json:"marketplacePartnerId,omitempty"`

	// Metrics: The list of metric settings.
	Metrics []MetricSettings_STATUS `json:"metrics,omitempty"`

	// ServiceBusRuleId: The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleId *string `json:"serviceBusRuleId,omitempty"`

	// StorageAccountId: The resource ID of the storage account to which you would like to send Diagnostic Logs.
	StorageAccountId *string `json:"storageAccountId,omitempty"`

	// WorkspaceId: The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs.
	// Example:
	// /subscriptions/4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/resourceGroups/insights-integration/providers/Microsoft.OperationalInsights/workspaces/viruela2
	WorkspaceId *string `json:"workspaceId,omitempty"`
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

// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
type LogSettings_STATUS struct {
	// Category: Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of
	// Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
	Category *string `json:"category,omitempty"`

	// CategoryGroup: Name of a Diagnostic Log category group for a resource type this setting is applied to. To obtain the
	// list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
	CategoryGroup *string `json:"categoryGroup,omitempty"`

	// Enabled: a value indicating whether this log is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// RetentionPolicy: the retention policy for this log.
	RetentionPolicy *RetentionPolicy_STATUS `json:"retentionPolicy,omitempty"`
}

// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.
type MetricSettings_STATUS struct {
	// Category: Name of a Diagnostic Metric category for a resource type this setting is applied to. To obtain the list of
	// Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation.
	Category *string `json:"category,omitempty"`

	// Enabled: a value indicating whether this category is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// RetentionPolicy: the retention policy for this category.
	RetentionPolicy *RetentionPolicy_STATUS `json:"retentionPolicy,omitempty"`

	// TimeGrain: the timegrain of the metric in ISO8601 format.
	TimeGrain *string `json:"timeGrain,omitempty"`
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

// Specifies the retention policy for the log.
type RetentionPolicy_STATUS struct {
	// Days: the number of days for the retention in days. A value of 0 will retain the events indefinitely.
	Days *int `json:"days,omitempty"`

	// Enabled: a value indicating whether the retention policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}
