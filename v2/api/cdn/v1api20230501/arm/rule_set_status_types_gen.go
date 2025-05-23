// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type RuleSet_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties of the Rule Set to create.
	Properties *RuleSetProperties_STATUS `json:"properties,omitempty"`

	// SystemData: Read only system data
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The JSON object that contains the properties of the Rule Set to create.
type RuleSetProperties_STATUS struct {
	DeploymentStatus *RuleSetProperties_DeploymentStatus_STATUS `json:"deploymentStatus,omitempty"`

	// ProfileName: The name of the profile which holds the rule set.
	ProfileName *string `json:"profileName,omitempty"`

	// ProvisioningState: Provisioning status
	ProvisioningState *RuleSetProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

type RuleSetProperties_DeploymentStatus_STATUS string

const (
	RuleSetProperties_DeploymentStatus_STATUS_Failed     = RuleSetProperties_DeploymentStatus_STATUS("Failed")
	RuleSetProperties_DeploymentStatus_STATUS_InProgress = RuleSetProperties_DeploymentStatus_STATUS("InProgress")
	RuleSetProperties_DeploymentStatus_STATUS_NotStarted = RuleSetProperties_DeploymentStatus_STATUS("NotStarted")
	RuleSetProperties_DeploymentStatus_STATUS_Succeeded  = RuleSetProperties_DeploymentStatus_STATUS("Succeeded")
)

// Mapping from string to RuleSetProperties_DeploymentStatus_STATUS
var ruleSetProperties_DeploymentStatus_STATUS_Values = map[string]RuleSetProperties_DeploymentStatus_STATUS{
	"failed":     RuleSetProperties_DeploymentStatus_STATUS_Failed,
	"inprogress": RuleSetProperties_DeploymentStatus_STATUS_InProgress,
	"notstarted": RuleSetProperties_DeploymentStatus_STATUS_NotStarted,
	"succeeded":  RuleSetProperties_DeploymentStatus_STATUS_Succeeded,
}

type RuleSetProperties_ProvisioningState_STATUS string

const (
	RuleSetProperties_ProvisioningState_STATUS_Creating  = RuleSetProperties_ProvisioningState_STATUS("Creating")
	RuleSetProperties_ProvisioningState_STATUS_Deleting  = RuleSetProperties_ProvisioningState_STATUS("Deleting")
	RuleSetProperties_ProvisioningState_STATUS_Failed    = RuleSetProperties_ProvisioningState_STATUS("Failed")
	RuleSetProperties_ProvisioningState_STATUS_Succeeded = RuleSetProperties_ProvisioningState_STATUS("Succeeded")
	RuleSetProperties_ProvisioningState_STATUS_Updating  = RuleSetProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to RuleSetProperties_ProvisioningState_STATUS
var ruleSetProperties_ProvisioningState_STATUS_Values = map[string]RuleSetProperties_ProvisioningState_STATUS{
	"creating":  RuleSetProperties_ProvisioningState_STATUS_Creating,
	"deleting":  RuleSetProperties_ProvisioningState_STATUS_Deleting,
	"failed":    RuleSetProperties_ProvisioningState_STATUS_Failed,
	"succeeded": RuleSetProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  RuleSetProperties_ProvisioningState_STATUS_Updating,
}
