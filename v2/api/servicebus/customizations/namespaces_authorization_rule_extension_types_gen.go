// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210101p "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20210101preview"
	v20210101ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20210101preview/storage"
	v20211101 "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/storage"
	v20221001p "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview"
	v20221001ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview/storage"
	v20240101 "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20240101"
	v20240101s "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20240101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type NamespacesAuthorizationRuleExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *NamespacesAuthorizationRuleExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210101p.NamespacesAuthorizationRule{},
		&v20210101ps.NamespacesAuthorizationRule{},
		&v20211101.NamespacesAuthorizationRule{},
		&v20211101s.NamespacesAuthorizationRule{},
		&v20221001p.NamespacesAuthorizationRule{},
		&v20221001ps.NamespacesAuthorizationRule{},
		&v20240101.NamespacesAuthorizationRule{},
		&v20240101s.NamespacesAuthorizationRule{}}
}
