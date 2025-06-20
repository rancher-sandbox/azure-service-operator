// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersAdvancedThreatProtectionSettings_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Advanced threat protection properties.
	Properties *ServerThreatProtectionProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersAdvancedThreatProtectionSettings_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-08-01"
func (settings FlexibleServersAdvancedThreatProtectionSettings_Spec) GetAPIVersion() string {
	return "2024-08-01"
}

// GetName returns the Name of the resource
func (settings *FlexibleServersAdvancedThreatProtectionSettings_Spec) GetName() string {
	return settings.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/advancedThreatProtectionSettings"
func (settings *FlexibleServersAdvancedThreatProtectionSettings_Spec) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/advancedThreatProtectionSettings"
}

// Properties of advanced threat protection state for a flexible server.
type ServerThreatProtectionProperties struct {
	// State: Specifies the state of the advanced threat protection, whether it is enabled, disabled, or a state has not been
	// applied yet on the flexible server.
	State *ServerThreatProtectionProperties_State `json:"state,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerThreatProtectionProperties_State string

const (
	ServerThreatProtectionProperties_State_Disabled = ServerThreatProtectionProperties_State("Disabled")
	ServerThreatProtectionProperties_State_Enabled  = ServerThreatProtectionProperties_State("Enabled")
)

// Mapping from string to ServerThreatProtectionProperties_State
var serverThreatProtectionProperties_State_Values = map[string]ServerThreatProtectionProperties_State{
	"disabled": ServerThreatProtectionProperties_State_Disabled,
	"enabled":  ServerThreatProtectionProperties_State_Enabled,
}
