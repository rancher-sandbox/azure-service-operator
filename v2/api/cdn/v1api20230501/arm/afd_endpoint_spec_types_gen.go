// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type AfdEndpoint_Spec struct {
	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties required to create an endpoint.
	Properties *AFDEndpointProperties `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &AfdEndpoint_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (endpoint AfdEndpoint_Spec) GetAPIVersion() string {
	return "2023-05-01"
}

// GetName returns the Name of the resource
func (endpoint *AfdEndpoint_Spec) GetName() string {
	return endpoint.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/afdEndpoints"
func (endpoint *AfdEndpoint_Spec) GetType() string {
	return "Microsoft.Cdn/profiles/afdEndpoints"
}

// The JSON object that contains the properties required to create an endpoint.
type AFDEndpointProperties struct {
	// AutoGeneratedDomainNameLabelScope: Indicates the endpoint name reuse scope. The default value is TenantReuse.
	AutoGeneratedDomainNameLabelScope *AutoGeneratedDomainNameLabelScope `json:"autoGeneratedDomainNameLabelScope,omitempty"`

	// EnabledState: Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
	EnabledState *AFDEndpointProperties_EnabledState `json:"enabledState,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type AFDEndpointProperties_EnabledState string

const (
	AFDEndpointProperties_EnabledState_Disabled = AFDEndpointProperties_EnabledState("Disabled")
	AFDEndpointProperties_EnabledState_Enabled  = AFDEndpointProperties_EnabledState("Enabled")
)

// Mapping from string to AFDEndpointProperties_EnabledState
var aFDEndpointProperties_EnabledState_Values = map[string]AFDEndpointProperties_EnabledState{
	"disabled": AFDEndpointProperties_EnabledState_Disabled,
	"enabled":  AFDEndpointProperties_EnabledState_Enabled,
}

// Indicates the endpoint name reuse scope. The default value is TenantReuse.
// +kubebuilder:validation:Enum={"NoReuse","ResourceGroupReuse","SubscriptionReuse","TenantReuse"}
type AutoGeneratedDomainNameLabelScope string

const (
	AutoGeneratedDomainNameLabelScope_NoReuse            = AutoGeneratedDomainNameLabelScope("NoReuse")
	AutoGeneratedDomainNameLabelScope_ResourceGroupReuse = AutoGeneratedDomainNameLabelScope("ResourceGroupReuse")
	AutoGeneratedDomainNameLabelScope_SubscriptionReuse  = AutoGeneratedDomainNameLabelScope("SubscriptionReuse")
	AutoGeneratedDomainNameLabelScope_TenantReuse        = AutoGeneratedDomainNameLabelScope("TenantReuse")
)

// Mapping from string to AutoGeneratedDomainNameLabelScope
var autoGeneratedDomainNameLabelScope_Values = map[string]AutoGeneratedDomainNameLabelScope{
	"noreuse":            AutoGeneratedDomainNameLabelScope_NoReuse,
	"resourcegroupreuse": AutoGeneratedDomainNameLabelScope_ResourceGroupReuse,
	"subscriptionreuse":  AutoGeneratedDomainNameLabelScope_SubscriptionReuse,
	"tenantreuse":        AutoGeneratedDomainNameLabelScope_TenantReuse,
}
