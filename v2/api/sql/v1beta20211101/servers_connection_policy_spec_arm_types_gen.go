// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Servers_ConnectionPolicy_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *ServerConnectionPolicyProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_ConnectionPolicy_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy Servers_ConnectionPolicy_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (policy *Servers_ConnectionPolicy_Spec_ARM) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/connectionPolicies"
func (policy *Servers_ConnectionPolicy_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers/connectionPolicies"
}

// The properties of a server connection policy.
type ServerConnectionPolicyProperties_ARM struct {
	// ConnectionType: The server connection type.
	ConnectionType *ServerConnectionPolicyProperties_ConnectionType `json:"connectionType,omitempty"`
}