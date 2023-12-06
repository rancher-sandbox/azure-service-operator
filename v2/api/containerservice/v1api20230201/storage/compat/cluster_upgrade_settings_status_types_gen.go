// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	v20231001s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/pkg/errors"
)

// Storage version of v1api20230202preview.ClusterUpgradeSettings_STATUS
// Settings for upgrading a cluster.
type ClusterUpgradeSettings_STATUS struct {
	OverrideSettings *UpgradeOverrideSettings_STATUS `json:"overrideSettings,omitempty"`
	PropertyBag      genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_ClusterUpgradeSettings_STATUS populates our ClusterUpgradeSettings_STATUS from the provided source ClusterUpgradeSettings_STATUS
func (settings *ClusterUpgradeSettings_STATUS) AssignProperties_From_ClusterUpgradeSettings_STATUS(source *v20231001s.ClusterUpgradeSettings_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// OverrideSettings
	if source.OverrideSettings != nil {
		var overrideSetting UpgradeOverrideSettings_STATUS
		err := overrideSetting.AssignProperties_From_UpgradeOverrideSettings_STATUS(source.OverrideSettings)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_UpgradeOverrideSettings_STATUS() to populate field OverrideSettings")
		}
		settings.OverrideSettings = &overrideSetting
	} else {
		settings.OverrideSettings = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		settings.PropertyBag = propertyBag
	} else {
		settings.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_ClusterUpgradeSettings_STATUS populates the provided destination ClusterUpgradeSettings_STATUS from our ClusterUpgradeSettings_STATUS
func (settings *ClusterUpgradeSettings_STATUS) AssignProperties_To_ClusterUpgradeSettings_STATUS(destination *v20231001s.ClusterUpgradeSettings_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(settings.PropertyBag)

	// OverrideSettings
	if settings.OverrideSettings != nil {
		var overrideSetting v20231001s.UpgradeOverrideSettings_STATUS
		err := settings.OverrideSettings.AssignProperties_To_UpgradeOverrideSettings_STATUS(&overrideSetting)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_UpgradeOverrideSettings_STATUS() to populate field OverrideSettings")
		}
		destination.OverrideSettings = &overrideSetting
	} else {
		destination.OverrideSettings = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}