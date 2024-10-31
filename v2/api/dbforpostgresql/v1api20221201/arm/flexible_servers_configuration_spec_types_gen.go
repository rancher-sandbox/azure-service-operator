// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersConfiguration_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a configuration.
	Properties *ConfigurationProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersConfiguration_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-12-01"
func (configuration FlexibleServersConfiguration_Spec) GetAPIVersion() string {
	return "2022-12-01"
}

// GetName returns the Name of the resource
func (configuration *FlexibleServersConfiguration_Spec) GetName() string {
	return configuration.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (configuration *FlexibleServersConfiguration_Spec) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
}

// The properties of a configuration.
type ConfigurationProperties struct {
	// Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}