// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220615 "github.com/Azure/azure-service-operator/v2/api/insights/v1api20220615"
	v20220615s "github.com/Azure/azure-service-operator/v2/api/insights/v1api20220615/storage"
	v20240101p "github.com/Azure/azure-service-operator/v2/api/insights/v1api20240101preview"
	v20240101ps "github.com/Azure/azure-service-operator/v2/api/insights/v1api20240101preview/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ScheduledQueryRuleExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ScheduledQueryRuleExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20220615.ScheduledQueryRule{},
		&v20220615s.ScheduledQueryRule{},
		&v20240101p.ScheduledQueryRule{},
		&v20240101ps.ScheduledQueryRule{}}
}
