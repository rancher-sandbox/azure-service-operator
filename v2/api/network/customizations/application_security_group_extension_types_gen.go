// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20240101 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240101"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20240101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ApplicationSecurityGroupExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ApplicationSecurityGroupExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20240101.ApplicationSecurityGroup{},
		&storage.ApplicationSecurityGroup{}}
}