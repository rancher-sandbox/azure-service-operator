// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type SitesSourcecontrol_Spec struct {
	// Kind: Kind of resource.
	Kind *string `json:"kind,omitempty"`
	Name string  `json:"name,omitempty"`

	// Properties: SiteSourceControl resource specific properties
	Properties *Sites_Sourcecontrol_Properties_Spec `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &SitesSourcecontrol_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-03-01"
func (sourcecontrol SitesSourcecontrol_Spec) GetAPIVersion() string {
	return "2022-03-01"
}

// GetName returns the Name of the resource
func (sourcecontrol *SitesSourcecontrol_Spec) GetName() string {
	return sourcecontrol.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Web/sites/sourcecontrols"
func (sourcecontrol *SitesSourcecontrol_Spec) GetType() string {
	return "Microsoft.Web/sites/sourcecontrols"
}

type Sites_Sourcecontrol_Properties_Spec struct {
	// Branch: Name of branch to use for deployment.
	Branch *string `json:"branch,omitempty"`

	// DeploymentRollbackEnabled: <code>true</code> to enable deployment rollback; otherwise, <code>false</code>.
	DeploymentRollbackEnabled *bool `json:"deploymentRollbackEnabled,omitempty"`

	// GitHubActionConfiguration: If GitHub Action is selected, than the associated configuration.
	GitHubActionConfiguration *GitHubActionConfiguration `json:"gitHubActionConfiguration,omitempty"`

	// IsGitHubAction: <code>true</code> if this is deployed via GitHub action.
	IsGitHubAction *bool `json:"isGitHubAction,omitempty"`

	// IsManualIntegration: <code>true</code> to limit to manual integration; <code>false</code> to enable continuous
	// integration (which configures webhooks into online repos like GitHub).
	IsManualIntegration *bool `json:"isManualIntegration,omitempty"`

	// IsMercurial: <code>true</code> for a Mercurial repository; <code>false</code> for a Git repository.
	IsMercurial *bool `json:"isMercurial,omitempty"`

	// RepoUrl: Repository or source control URL.
	RepoUrl *string `json:"repoUrl,omitempty"`
}

// The GitHub action configuration.
type GitHubActionConfiguration struct {
	// CodeConfiguration: GitHub Action code configuration.
	CodeConfiguration *GitHubActionCodeConfiguration `json:"codeConfiguration,omitempty"`

	// ContainerConfiguration: GitHub Action container configuration.
	ContainerConfiguration *GitHubActionContainerConfiguration `json:"containerConfiguration,omitempty"`

	// GenerateWorkflowFile: Workflow option to determine whether the workflow file should be generated and written to the
	// repository.
	GenerateWorkflowFile *bool `json:"generateWorkflowFile,omitempty"`

	// IsLinux: This will help determine the workflow configuration to select.
	IsLinux *bool `json:"isLinux,omitempty"`
}

// The GitHub action code configuration.
type GitHubActionCodeConfiguration struct {
	// RuntimeStack: Runtime stack is used to determine the workflow file content for code base apps.
	RuntimeStack *string `json:"runtimeStack,omitempty"`

	// RuntimeVersion: Runtime version is used to determine what build version to set in the workflow file.
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`
}

// The GitHub action container configuration.
type GitHubActionContainerConfiguration struct {
	// ImageName: The image name for the build.
	ImageName *string `json:"imageName,omitempty"`

	// Password: The password used to upload the image to the container registry.
	Password *string `json:"password,omitempty"`

	// ServerUrl: The server URL for the container registry where the build will be hosted.
	ServerUrl *string `json:"serverUrl,omitempty"`

	// Username: The username used to upload the image to the container registry.
	Username *string `json:"username,omitempty"`
}
