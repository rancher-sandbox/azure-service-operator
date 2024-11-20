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

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccountsblobservicescontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccountsblobservicescontainers/status,storageaccountsblobservicescontainers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230101.StorageAccountsBlobServicesContainer
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2023-01-01/blob.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}
type StorageAccountsBlobServicesContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsBlobServicesContainer_Spec   `json:"spec,omitempty"`
	Status            StorageAccountsBlobServicesContainer_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsBlobServicesContainer{}

// GetConditions returns the conditions of the resource
func (container *StorageAccountsBlobServicesContainer) GetConditions() conditions.Conditions {
	return container.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (container *StorageAccountsBlobServicesContainer) SetConditions(conditions conditions.Conditions) {
	container.Status.Conditions = conditions
}

var _ configmaps.Exporter = &StorageAccountsBlobServicesContainer{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (container *StorageAccountsBlobServicesContainer) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if container.Spec.OperatorSpec == nil {
		return nil
	}
	return container.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &StorageAccountsBlobServicesContainer{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (container *StorageAccountsBlobServicesContainer) SecretDestinationExpressions() []*core.DestinationExpression {
	if container.Spec.OperatorSpec == nil {
		return nil
	}
	return container.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &StorageAccountsBlobServicesContainer{}

// AzureName returns the Azure name of the resource
func (container *StorageAccountsBlobServicesContainer) AzureName() string {
	return container.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-01"
func (container StorageAccountsBlobServicesContainer) GetAPIVersion() string {
	return "2023-01-01"
}

// GetResourceScope returns the scope of the resource
func (container *StorageAccountsBlobServicesContainer) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (container *StorageAccountsBlobServicesContainer) GetSpec() genruntime.ConvertibleSpec {
	return &container.Spec
}

// GetStatus returns the status of this resource
func (container *StorageAccountsBlobServicesContainer) GetStatus() genruntime.ConvertibleStatus {
	return &container.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (container *StorageAccountsBlobServicesContainer) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices/containers"
func (container *StorageAccountsBlobServicesContainer) GetType() string {
	return "Microsoft.Storage/storageAccounts/blobServices/containers"
}

// NewEmptyStatus returns a new empty (blank) status
func (container *StorageAccountsBlobServicesContainer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccountsBlobServicesContainer_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (container *StorageAccountsBlobServicesContainer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(container.Spec)
	return container.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (container *StorageAccountsBlobServicesContainer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccountsBlobServicesContainer_STATUS); ok {
		container.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccountsBlobServicesContainer_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	container.Status = st
	return nil
}

// Hub marks that this StorageAccountsBlobServicesContainer is the hub type for conversion
func (container *StorageAccountsBlobServicesContainer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (container *StorageAccountsBlobServicesContainer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: container.Spec.OriginalVersion,
		Kind:    "StorageAccountsBlobServicesContainer",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230101.StorageAccountsBlobServicesContainer
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2023-01-01/blob.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}
type StorageAccountsBlobServicesContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsBlobServicesContainer `json:"items"`
}

// Storage version of v1api20230101.StorageAccountsBlobServicesContainer_Spec
type StorageAccountsBlobServicesContainer_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                      string                                            `json:"azureName,omitempty"`
	DefaultEncryptionScope         *string                                           `json:"defaultEncryptionScope,omitempty"`
	DenyEncryptionScopeOverride    *bool                                             `json:"denyEncryptionScopeOverride,omitempty"`
	EnableNfsV3AllSquash           *bool                                             `json:"enableNfsV3AllSquash,omitempty"`
	EnableNfsV3RootSquash          *bool                                             `json:"enableNfsV3RootSquash,omitempty"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning                   `json:"immutableStorageWithVersioning,omitempty"`
	Metadata                       map[string]string                                 `json:"metadata,omitempty"`
	OperatorSpec                   *StorageAccountsBlobServicesContainerOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion                string                                            `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccountsBlobService resource
	Owner        *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccountsBlobService"`
	PropertyBag  genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicAccess *string                            `json:"publicAccess,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsBlobServicesContainer_Spec{}

// ConvertSpecFrom populates our StorageAccountsBlobServicesContainer_Spec from the provided source
func (container *StorageAccountsBlobServicesContainer_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == container {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(container)
}

// ConvertSpecTo populates the provided destination from our StorageAccountsBlobServicesContainer_Spec
func (container *StorageAccountsBlobServicesContainer_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == container {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(container)
}

// Storage version of v1api20230101.StorageAccountsBlobServicesContainer_STATUS
type StorageAccountsBlobServicesContainer_STATUS struct {
	Conditions                     []conditions.Condition                 `json:"conditions,omitempty"`
	DefaultEncryptionScope         *string                                `json:"defaultEncryptionScope,omitempty"`
	Deleted                        *bool                                  `json:"deleted,omitempty"`
	DeletedTime                    *string                                `json:"deletedTime,omitempty"`
	DenyEncryptionScopeOverride    *bool                                  `json:"denyEncryptionScopeOverride,omitempty"`
	EnableNfsV3AllSquash           *bool                                  `json:"enableNfsV3AllSquash,omitempty"`
	EnableNfsV3RootSquash          *bool                                  `json:"enableNfsV3RootSquash,omitempty"`
	Etag                           *string                                `json:"etag,omitempty"`
	HasImmutabilityPolicy          *bool                                  `json:"hasImmutabilityPolicy,omitempty"`
	HasLegalHold                   *bool                                  `json:"hasLegalHold,omitempty"`
	Id                             *string                                `json:"id,omitempty"`
	ImmutabilityPolicy             *ImmutabilityPolicyProperties_STATUS   `json:"immutabilityPolicy,omitempty"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning_STATUS `json:"immutableStorageWithVersioning,omitempty"`
	LastModifiedTime               *string                                `json:"lastModifiedTime,omitempty"`
	LeaseDuration                  *string                                `json:"leaseDuration,omitempty"`
	LeaseState                     *string                                `json:"leaseState,omitempty"`
	LeaseStatus                    *string                                `json:"leaseStatus,omitempty"`
	LegalHold                      *LegalHoldProperties_STATUS            `json:"legalHold,omitempty"`
	Metadata                       map[string]string                      `json:"metadata,omitempty"`
	Name                           *string                                `json:"name,omitempty"`
	PropertyBag                    genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	PublicAccess                   *string                                `json:"publicAccess,omitempty"`
	RemainingRetentionDays         *int                                   `json:"remainingRetentionDays,omitempty"`
	Type                           *string                                `json:"type,omitempty"`
	Version                        *string                                `json:"version,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccountsBlobServicesContainer_STATUS{}

// ConvertStatusFrom populates our StorageAccountsBlobServicesContainer_STATUS from the provided source
func (container *StorageAccountsBlobServicesContainer_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == container {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(container)
}

// ConvertStatusTo populates the provided destination from our StorageAccountsBlobServicesContainer_STATUS
func (container *StorageAccountsBlobServicesContainer_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == container {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(container)
}

// Storage version of v1api20230101.ImmutabilityPolicyProperties_STATUS
// The properties of an ImmutabilityPolicy of a blob container.
type ImmutabilityPolicyProperties_STATUS struct {
	AllowProtectedAppendWrites            *bool                          `json:"allowProtectedAppendWrites,omitempty"`
	AllowProtectedAppendWritesAll         *bool                          `json:"allowProtectedAppendWritesAll,omitempty"`
	Etag                                  *string                        `json:"etag,omitempty"`
	ImmutabilityPeriodSinceCreationInDays *int                           `json:"immutabilityPeriodSinceCreationInDays,omitempty"`
	PropertyBag                           genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	State                                 *string                        `json:"state,omitempty"`
	UpdateHistory                         []UpdateHistoryProperty_STATUS `json:"updateHistory,omitempty"`
}

// Storage version of v1api20230101.ImmutableStorageWithVersioning
// Object level immutability properties of the container.
type ImmutableStorageWithVersioning struct {
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.ImmutableStorageWithVersioning_STATUS
// Object level immutability properties of the container.
type ImmutableStorageWithVersioning_STATUS struct {
	Enabled        *bool                  `json:"enabled,omitempty"`
	MigrationState *string                `json:"migrationState,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TimeStamp      *string                `json:"timeStamp,omitempty"`
}

// Storage version of v1api20230101.LegalHoldProperties_STATUS
// The LegalHold property of a blob container.
type LegalHoldProperties_STATUS struct {
	HasLegalHold                 *bool                                `json:"hasLegalHold,omitempty"`
	PropertyBag                  genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	ProtectedAppendWritesHistory *ProtectedAppendWritesHistory_STATUS `json:"protectedAppendWritesHistory,omitempty"`
	Tags                         []TagProperty_STATUS                 `json:"tags,omitempty"`
}

// Storage version of v1api20230101.StorageAccountsBlobServicesContainerOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type StorageAccountsBlobServicesContainerOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20230101.ProtectedAppendWritesHistory_STATUS
// Protected append writes history setting for the blob container with Legal holds.
type ProtectedAppendWritesHistory_STATUS struct {
	AllowProtectedAppendWritesAll *bool                  `json:"allowProtectedAppendWritesAll,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Timestamp                     *string                `json:"timestamp,omitempty"`
}

// Storage version of v1api20230101.TagProperty_STATUS
// A tag of the LegalHold of a blob container.
type TagProperty_STATUS struct {
	ObjectIdentifier *string                `json:"objectIdentifier,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tag              *string                `json:"tag,omitempty"`
	TenantId         *string                `json:"tenantId,omitempty"`
	Timestamp        *string                `json:"timestamp,omitempty"`
	Upn              *string                `json:"upn,omitempty"`
}

// Storage version of v1api20230101.UpdateHistoryProperty_STATUS
// An update history of the ImmutabilityPolicy of a blob container.
type UpdateHistoryProperty_STATUS struct {
	AllowProtectedAppendWrites            *bool                  `json:"allowProtectedAppendWrites,omitempty"`
	AllowProtectedAppendWritesAll         *bool                  `json:"allowProtectedAppendWritesAll,omitempty"`
	ImmutabilityPeriodSinceCreationInDays *int                   `json:"immutabilityPeriodSinceCreationInDays,omitempty"`
	ObjectIdentifier                      *string                `json:"objectIdentifier,omitempty"`
	PropertyBag                           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TenantId                              *string                `json:"tenantId,omitempty"`
	Timestamp                             *string                `json:"timestamp,omitempty"`
	Update                                *string                `json:"update,omitempty"`
	Upn                                   *string                `json:"upn,omitempty"`
}

func init() {
	SchemeBuilder.Register(&StorageAccountsBlobServicesContainer{}, &StorageAccountsBlobServicesContainerList{})
}
