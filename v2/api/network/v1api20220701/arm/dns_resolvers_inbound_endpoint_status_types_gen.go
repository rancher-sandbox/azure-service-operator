// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type DnsResolversInboundEndpoint_STATUS struct {
	// Etag: ETag of the inbound endpoint.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the inbound endpoint.
	Properties *InboundEndpointProperties_STATUS `json:"properties,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Represents the properties of an inbound endpoint for a DNS resolver.
type InboundEndpointProperties_STATUS struct {
	// IpConfigurations: IP configurations for the inbound endpoint.
	IpConfigurations []IpConfiguration_STATUS `json:"ipConfigurations,omitempty"`

	// ProvisioningState: The current provisioning state of the inbound endpoint. This is a read-only property and any attempt
	// to set this value will be ignored.
	ProvisioningState *DnsresolverProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceGuid: The resourceGuid property of the inbound endpoint resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`
}

// IP configuration.
type IpConfiguration_STATUS struct {
	// PrivateIpAddress: Private IP address of the IP configuration.
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`

	// PrivateIpAllocationMethod: Private IP address allocation method.
	PrivateIpAllocationMethod *IpConfiguration_PrivateIpAllocationMethod_STATUS `json:"privateIpAllocationMethod,omitempty"`

	// Subnet: The reference to the subnet bound to the IP configuration.
	Subnet *SubResource_STATUS `json:"subnet,omitempty"`
}

type IpConfiguration_PrivateIpAllocationMethod_STATUS string

const (
	IpConfiguration_PrivateIpAllocationMethod_STATUS_Dynamic = IpConfiguration_PrivateIpAllocationMethod_STATUS("Dynamic")
	IpConfiguration_PrivateIpAllocationMethod_STATUS_Static  = IpConfiguration_PrivateIpAllocationMethod_STATUS("Static")
)

// Mapping from string to IpConfiguration_PrivateIpAllocationMethod_STATUS
var ipConfiguration_PrivateIpAllocationMethod_STATUS_Values = map[string]IpConfiguration_PrivateIpAllocationMethod_STATUS{
	"dynamic": IpConfiguration_PrivateIpAllocationMethod_STATUS_Dynamic,
	"static":  IpConfiguration_PrivateIpAllocationMethod_STATUS_Static,
}
