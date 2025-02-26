// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// Virtual Network resource.
type VirtualNetwork_STATUS struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// ExtendedLocation: The extended location of the virtual network.
	ExtendedLocation *ExtendedLocation_STATUS `json:"extendedLocation,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the virtual network.
	Properties *VirtualNetworkPropertiesFormat_STATUS `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of the virtual network.
type VirtualNetworkPropertiesFormat_STATUS struct {
	// AddressSpace: The AddressSpace that contains an array of IP address ranges that can be used by subnets.
	AddressSpace *AddressSpace_STATUS `json:"addressSpace,omitempty"`

	// BgpCommunities: Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
	BgpCommunities *VirtualNetworkBgpCommunities_STATUS `json:"bgpCommunities,omitempty"`

	// DdosProtectionPlan: The DDoS protection plan associated with the virtual network.
	DdosProtectionPlan *SubResource_STATUS `json:"ddosProtectionPlan,omitempty"`

	// DhcpOptions: The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
	DhcpOptions *DhcpOptions_STATUS `json:"dhcpOptions,omitempty"`

	// EnableDdosProtection: Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It
	// requires a DDoS protection plan associated with the resource.
	EnableDdosProtection *bool `json:"enableDdosProtection,omitempty"`

	// EnableVmProtection: Indicates if VM protection is enabled for all the subnets in the virtual network.
	EnableVmProtection *bool `json:"enableVmProtection,omitempty"`

	// IpAllocations: Array of IpAllocation which reference this VNET.
	IpAllocations []SubResource_STATUS `json:"ipAllocations,omitempty"`

	// ProvisioningState: The provisioning state of the virtual network resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceGuid: The resourceGuid property of the Virtual Network resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`
}

// DhcpOptions contains an array of DNS servers available to VMs deployed in the virtual network. Standard DHCP option for
// a subnet overrides VNET DHCP options.
type DhcpOptions_STATUS struct {
	// DnsServers: The list of DNS servers IP addresses.
	DnsServers []string `json:"dnsServers,omitempty"`
}

// Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
type VirtualNetworkBgpCommunities_STATUS struct {
	// RegionalCommunity: The BGP community associated with the region of the virtual network.
	RegionalCommunity *string `json:"regionalCommunity,omitempty"`

	// VirtualNetworkCommunity: The BGP community associated with the virtual network.
	VirtualNetworkCommunity *string `json:"virtualNetworkCommunity,omitempty"`
}
