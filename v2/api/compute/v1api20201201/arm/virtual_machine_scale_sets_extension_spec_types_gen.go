// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type VirtualMachineScaleSets_Extension_Spec struct {
	// Name: The name of the extension.
	Name string `json:"name,omitempty"`

	// Properties: Describes the properties of a Virtual Machine Scale Set Extension.
	Properties *VirtualMachineScaleSetExtensionProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualMachineScaleSets_Extension_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (extension VirtualMachineScaleSets_Extension_Spec) GetAPIVersion() string {
	return "2020-12-01"
}

// GetName returns the Name of the resource
func (extension *VirtualMachineScaleSets_Extension_Spec) GetName() string {
	return extension.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/virtualMachineScaleSets/extensions"
func (extension *VirtualMachineScaleSets_Extension_Spec) GetType() string {
	return "Microsoft.Compute/virtualMachineScaleSets/extensions"
}

// Describes the properties of a Virtual Machine Scale Set Extension.
type VirtualMachineScaleSetExtensionProperties struct {
	// AutoUpgradeMinorVersion: Indicates whether the extension should use a newer minor version if one is available at
	// deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this
	// property set to true.
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty"`

	// EnableAutomaticUpgrade: Indicates whether the extension should be automatically upgraded by the platform if there is a
	// newer version of the extension available.
	EnableAutomaticUpgrade *bool `json:"enableAutomaticUpgrade,omitempty"`

	// ForceUpdateTag: If a value is provided and is different from the previous value, the extension handler will be forced to
	// update even if the extension configuration has not changed.
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty"`

	// ProtectedSettings: The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected
	// settings at all.
	ProtectedSettings map[string]string `json:"protectedSettings,omitempty"`

	// ProvisionAfterExtensions: Collection of extension names after which this extension needs to be provisioned.
	ProvisionAfterExtensions []string `json:"provisionAfterExtensions,omitempty"`

	// Publisher: The name of the extension handler publisher.
	Publisher *string `json:"publisher,omitempty"`

	// Settings: Json formatted public settings for the extension.
	Settings map[string]v1.JSON `json:"settings,omitempty"`

	// Type: Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `json:"type,omitempty"`

	// TypeHandlerVersion: Specifies the version of the script handler.
	TypeHandlerVersion *string `json:"typeHandlerVersion,omitempty"`
}