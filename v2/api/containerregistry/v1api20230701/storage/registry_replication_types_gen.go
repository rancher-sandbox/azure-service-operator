// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=containerregistry.azure.com,resources=registryreplications,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=containerregistry.azure.com,resources={registryreplications/status,registryreplications/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230701.RegistryReplication
// Generator information:
// - Generated from: /containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/containerregistry.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}
type RegistryReplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegistryReplication_Spec   `json:"spec,omitempty"`
	Status            RegistryReplication_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RegistryReplication{}

// GetConditions returns the conditions of the resource
func (replication *RegistryReplication) GetConditions() conditions.Conditions {
	return replication.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (replication *RegistryReplication) SetConditions(conditions conditions.Conditions) {
	replication.Status.Conditions = conditions
}

var _ configmaps.Exporter = &RegistryReplication{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (replication *RegistryReplication) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if replication.Spec.OperatorSpec == nil {
		return nil
	}
	return replication.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &RegistryReplication{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (replication *RegistryReplication) SecretDestinationExpressions() []*core.DestinationExpression {
	if replication.Spec.OperatorSpec == nil {
		return nil
	}
	return replication.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &RegistryReplication{}

// AzureName returns the Azure name of the resource
func (replication *RegistryReplication) AzureName() string {
	return replication.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-07-01"
func (replication RegistryReplication) GetAPIVersion() string {
	return "2023-07-01"
}

// GetResourceScope returns the scope of the resource
func (replication *RegistryReplication) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (replication *RegistryReplication) GetSpec() genruntime.ConvertibleSpec {
	return &replication.Spec
}

// GetStatus returns the status of this resource
func (replication *RegistryReplication) GetStatus() genruntime.ConvertibleStatus {
	return &replication.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (replication *RegistryReplication) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ContainerRegistry/registries/replications"
func (replication *RegistryReplication) GetType() string {
	return "Microsoft.ContainerRegistry/registries/replications"
}

// NewEmptyStatus returns a new empty (blank) status
func (replication *RegistryReplication) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RegistryReplication_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (replication *RegistryReplication) Owner() *genruntime.ResourceReference {
	if replication.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(replication.Spec)
	return replication.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (replication *RegistryReplication) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RegistryReplication_STATUS); ok {
		replication.Status = *st
		return nil
	}

	// Convert status to required version
	var st RegistryReplication_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	replication.Status = st
	return nil
}

// Hub marks that this RegistryReplication is the hub type for conversion
func (replication *RegistryReplication) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (replication *RegistryReplication) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: replication.Spec.OriginalVersion,
		Kind:    "RegistryReplication",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230701.RegistryReplication
// Generator information:
// - Generated from: /containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/containerregistry.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}
type RegistryReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryReplication `json:"items"`
}

// Storage version of v1api20230701.RegistryReplication_Spec
type RegistryReplication_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                           `json:"azureName,omitempty"`
	Location        *string                          `json:"location,omitempty"`
	OperatorSpec    *RegistryReplicationOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                           `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a containerregistry.azure.com/Registry resource
	Owner                 *genruntime.KnownResourceReference `group:"containerregistry.azure.com" json:"owner,omitempty" kind:"Registry"`
	PropertyBag           genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RegionEndpointEnabled *bool                              `json:"regionEndpointEnabled,omitempty"`
	Tags                  map[string]string                  `json:"tags,omitempty"`
	ZoneRedundancy        *string                            `json:"zoneRedundancy,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RegistryReplication_Spec{}

// ConvertSpecFrom populates our RegistryReplication_Spec from the provided source
func (replication *RegistryReplication_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == replication {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(replication)
}

// ConvertSpecTo populates the provided destination from our RegistryReplication_Spec
func (replication *RegistryReplication_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == replication {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(replication)
}

// Storage version of v1api20230701.RegistryReplication_STATUS
type RegistryReplication_STATUS struct {
	Conditions            []conditions.Condition `json:"conditions,omitempty"`
	Id                    *string                `json:"id,omitempty"`
	Location              *string                `json:"location,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState     *string                `json:"provisioningState,omitempty"`
	RegionEndpointEnabled *bool                  `json:"regionEndpointEnabled,omitempty"`
	Status                *Status_STATUS         `json:"status,omitempty"`
	SystemData            *SystemData_STATUS     `json:"systemData,omitempty"`
	Tags                  map[string]string      `json:"tags,omitempty"`
	Type                  *string                `json:"type,omitempty"`
	ZoneRedundancy        *string                `json:"zoneRedundancy,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RegistryReplication_STATUS{}

// ConvertStatusFrom populates our RegistryReplication_STATUS from the provided source
func (replication *RegistryReplication_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == replication {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(replication)
}

// ConvertStatusTo populates the provided destination from our RegistryReplication_STATUS
func (replication *RegistryReplication_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == replication {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(replication)
}

// Storage version of v1api20230701.RegistryReplicationOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type RegistryReplicationOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&RegistryReplication{}, &RegistryReplicationList{})
}
