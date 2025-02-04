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

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=mongodbdatabasecollections,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={mongodbdatabasecollections/status,mongodbdatabasecollections/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240815.MongodbDatabaseCollection
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}
type MongodbDatabaseCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongodbDatabaseCollection_Spec   `json:"spec,omitempty"`
	Status            MongodbDatabaseCollection_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseCollection{}

// GetConditions returns the conditions of the resource
func (collection *MongodbDatabaseCollection) GetConditions() conditions.Conditions {
	return collection.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (collection *MongodbDatabaseCollection) SetConditions(conditions conditions.Conditions) {
	collection.Status.Conditions = conditions
}

var _ configmaps.Exporter = &MongodbDatabaseCollection{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (collection *MongodbDatabaseCollection) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if collection.Spec.OperatorSpec == nil {
		return nil
	}
	return collection.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &MongodbDatabaseCollection{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (collection *MongodbDatabaseCollection) SecretDestinationExpressions() []*core.DestinationExpression {
	if collection.Spec.OperatorSpec == nil {
		return nil
	}
	return collection.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &MongodbDatabaseCollection{}

// AzureName returns the Azure name of the resource
func (collection *MongodbDatabaseCollection) AzureName() string {
	return collection.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-08-15"
func (collection MongodbDatabaseCollection) GetAPIVersion() string {
	return "2024-08-15"
}

// GetResourceScope returns the scope of the resource
func (collection *MongodbDatabaseCollection) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (collection *MongodbDatabaseCollection) GetSpec() genruntime.ConvertibleSpec {
	return &collection.Spec
}

// GetStatus returns the status of this resource
func (collection *MongodbDatabaseCollection) GetStatus() genruntime.ConvertibleStatus {
	return &collection.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (collection *MongodbDatabaseCollection) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
func (collection *MongodbDatabaseCollection) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
}

// NewEmptyStatus returns a new empty (blank) status
func (collection *MongodbDatabaseCollection) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &MongodbDatabaseCollection_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (collection *MongodbDatabaseCollection) Owner() *genruntime.ResourceReference {
	if collection.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(collection.Spec)
	return collection.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (collection *MongodbDatabaseCollection) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*MongodbDatabaseCollection_STATUS); ok {
		collection.Status = *st
		return nil
	}

	// Convert status to required version
	var st MongodbDatabaseCollection_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	collection.Status = st
	return nil
}

// Hub marks that this MongodbDatabaseCollection is the hub type for conversion
func (collection *MongodbDatabaseCollection) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (collection *MongodbDatabaseCollection) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: collection.Spec.OriginalVersion,
		Kind:    "MongodbDatabaseCollection",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240815.MongodbDatabaseCollection
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}
type MongodbDatabaseCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseCollection `json:"items"`
}

// Storage version of v1api20240815.MongodbDatabaseCollection_Spec
type MongodbDatabaseCollection_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                 `json:"azureName,omitempty"`
	Location        *string                                `json:"location,omitempty"`
	OperatorSpec    *MongodbDatabaseCollectionOperatorSpec `json:"operatorSpec,omitempty"`
	Options         *CreateUpdateOptions                   `json:"options,omitempty"`
	OriginalVersion string                                 `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/MongodbDatabase resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"MongodbDatabase"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *MongoDBCollectionResource         `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &MongodbDatabaseCollection_Spec{}

// ConvertSpecFrom populates our MongodbDatabaseCollection_Spec from the provided source
func (collection *MongodbDatabaseCollection_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == collection {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(collection)
}

// ConvertSpecTo populates the provided destination from our MongodbDatabaseCollection_Spec
func (collection *MongodbDatabaseCollection_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == collection {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(collection)
}

// Storage version of v1api20240815.MongodbDatabaseCollection_STATUS
type MongodbDatabaseCollection_STATUS struct {
	Conditions  []conditions.Condition                          `json:"conditions,omitempty"`
	Id          *string                                         `json:"id,omitempty"`
	Location    *string                                         `json:"location,omitempty"`
	Name        *string                                         `json:"name,omitempty"`
	Options     *OptionsResource_STATUS                         `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                          `json:"$propertyBag,omitempty"`
	Resource    *MongoDBCollectionGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                               `json:"tags,omitempty"`
	Type        *string                                         `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &MongodbDatabaseCollection_STATUS{}

// ConvertStatusFrom populates our MongodbDatabaseCollection_STATUS from the provided source
func (collection *MongodbDatabaseCollection_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == collection {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(collection)
}

// ConvertStatusTo populates the provided destination from our MongodbDatabaseCollection_STATUS
func (collection *MongodbDatabaseCollection_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == collection {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(collection)
}

// Storage version of v1api20240815.MongoDBCollectionGetProperties_Resource_STATUS
type MongoDBCollectionGetProperties_Resource_STATUS struct {
	AnalyticalStorageTtl *int                          `json:"analyticalStorageTtl,omitempty"`
	CreateMode           *string                       `json:"createMode,omitempty"`
	Etag                 *string                       `json:"_etag,omitempty"`
	Id                   *string                       `json:"id,omitempty"`
	Indexes              []MongoIndex_STATUS           `json:"indexes,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RestoreParameters    *RestoreParametersBase_STATUS `json:"restoreParameters,omitempty"`
	Rid                  *string                       `json:"_rid,omitempty"`
	ShardKey             map[string]string             `json:"shardKey,omitempty"`
	Ts                   *float64                      `json:"_ts,omitempty"`
}

// Storage version of v1api20240815.MongoDBCollectionResource
// Cosmos DB MongoDB collection resource object
type MongoDBCollectionResource struct {
	AnalyticalStorageTtl *int                   `json:"analyticalStorageTtl,omitempty"`
	CreateMode           *string                `json:"createMode,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Indexes              []MongoIndex           `json:"indexes,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RestoreParameters    *RestoreParametersBase `json:"restoreParameters,omitempty"`
	ShardKey             map[string]string      `json:"shardKey,omitempty"`
}

// Storage version of v1api20240815.MongodbDatabaseCollectionOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type MongodbDatabaseCollectionOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20240815.MongoIndex
// Cosmos DB MongoDB collection index key
type MongoIndex struct {
	Key         *MongoIndexKeys        `json:"key,omitempty"`
	Options     *MongoIndexOptions     `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240815.MongoIndex_STATUS
// Cosmos DB MongoDB collection index key
type MongoIndex_STATUS struct {
	Key         *MongoIndexKeys_STATUS    `json:"key,omitempty"`
	Options     *MongoIndexOptions_STATUS `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240815.MongoIndexKeys
// Cosmos DB MongoDB collection resource object
type MongoIndexKeys struct {
	Keys        []string               `json:"keys,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240815.MongoIndexKeys_STATUS
// Cosmos DB MongoDB collection resource object
type MongoIndexKeys_STATUS struct {
	Keys        []string               `json:"keys,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240815.MongoIndexOptions
// Cosmos DB MongoDB collection index options
type MongoIndexOptions struct {
	ExpireAfterSeconds *int                   `json:"expireAfterSeconds,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Unique             *bool                  `json:"unique,omitempty"`
}

// Storage version of v1api20240815.MongoIndexOptions_STATUS
// Cosmos DB MongoDB collection index options
type MongoIndexOptions_STATUS struct {
	ExpireAfterSeconds *int                   `json:"expireAfterSeconds,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Unique             *bool                  `json:"unique,omitempty"`
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseCollection{}, &MongodbDatabaseCollectionList{})
}
