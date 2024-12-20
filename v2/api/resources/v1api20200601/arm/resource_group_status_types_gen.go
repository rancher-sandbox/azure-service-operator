// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// Resource group information.
type ResourceGroup_STATUS struct {
	// Id: The ID of the resource group.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group. It cannot be changed after the resource group has been created. It must be
	// one of the supported Azure locations.
	Location *string `json:"location,omitempty"`

	// ManagedBy: The ID of the resource that manages this resource group.
	ManagedBy *string `json:"managedBy,omitempty"`

	// Name: The name of the resource group.
	Name *string `json:"name,omitempty"`

	// Properties: The resource group properties.
	Properties *ResourceGroupProperties_STATUS `json:"properties,omitempty"`

	// Tags: The tags attached to the resource group.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource group.
	Type *string `json:"type,omitempty"`
}

// The resource group properties.
type ResourceGroupProperties_STATUS struct {
	// ProvisioningState: The provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}
