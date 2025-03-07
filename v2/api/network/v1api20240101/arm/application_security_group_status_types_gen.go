// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// An application security group in a resource group.
type ApplicationSecurityGroup_STATUS struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the application security group.
	Properties *ApplicationSecurityGroupPropertiesFormat_STATUS `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Application security group properties.
type ApplicationSecurityGroupPropertiesFormat_STATUS struct {
	// ProvisioningState: The provisioning state of the application security group resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceGuid: The resource GUID property of the application security group resource. It uniquely identifies a resource,
	// even if the user changes its name or migrate the resource across subscriptions or resource groups.
	ResourceGuid *string `json:"resourceGuid,omitempty"`
}

// The current provisioning state.
type ProvisioningState_STATUS string

const (
	ProvisioningState_STATUS_Deleting  = ProvisioningState_STATUS("Deleting")
	ProvisioningState_STATUS_Failed    = ProvisioningState_STATUS("Failed")
	ProvisioningState_STATUS_Succeeded = ProvisioningState_STATUS("Succeeded")
	ProvisioningState_STATUS_Updating  = ProvisioningState_STATUS("Updating")
)

// Mapping from string to ProvisioningState_STATUS
var provisioningState_STATUS_Values = map[string]ProvisioningState_STATUS{
	"deleting":  ProvisioningState_STATUS_Deleting,
	"failed":    ProvisioningState_STATUS_Failed,
	"succeeded": ProvisioningState_STATUS_Succeeded,
	"updating":  ProvisioningState_STATUS_Updating,
}
