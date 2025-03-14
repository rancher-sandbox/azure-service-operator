// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesEventhubsConsumerGroup_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Single item in List or Get Consumer group operation
	Properties *Namespaces_Eventhubs_Consumergroup_Properties_Spec `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesEventhubsConsumerGroup_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-01-01"
func (group NamespacesEventhubsConsumerGroup_Spec) GetAPIVersion() string {
	return "2024-01-01"
}

// GetName returns the Name of the resource
func (group *NamespacesEventhubsConsumerGroup_Spec) GetName() string {
	return group.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/consumergroups"
func (group *NamespacesEventhubsConsumerGroup_Spec) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/consumergroups"
}

type Namespaces_Eventhubs_Consumergroup_Properties_Spec struct {
	// UserMetadata: User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be
	// used to store descriptive data, such as list of teams and their contact information also user-defined configuration
	// settings can be stored.
	UserMetadata *string `json:"userMetadata,omitempty"`
}
