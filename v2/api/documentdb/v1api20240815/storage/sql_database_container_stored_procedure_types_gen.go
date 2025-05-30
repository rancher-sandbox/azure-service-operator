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

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabasecontainerstoredprocedures,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabasecontainerstoredprocedures/status,sqldatabasecontainerstoredprocedures/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240815.SqlDatabaseContainerStoredProcedure
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures/{storedProcedureName}
type SqlDatabaseContainerStoredProcedure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseContainerStoredProcedure_Spec   `json:"spec,omitempty"`
	Status            SqlDatabaseContainerStoredProcedure_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerStoredProcedure{}

// GetConditions returns the conditions of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetConditions() conditions.Conditions {
	return procedure.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (procedure *SqlDatabaseContainerStoredProcedure) SetConditions(conditions conditions.Conditions) {
	procedure.Status.Conditions = conditions
}

var _ configmaps.Exporter = &SqlDatabaseContainerStoredProcedure{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (procedure *SqlDatabaseContainerStoredProcedure) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if procedure.Spec.OperatorSpec == nil {
		return nil
	}
	return procedure.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &SqlDatabaseContainerStoredProcedure{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (procedure *SqlDatabaseContainerStoredProcedure) SecretDestinationExpressions() []*core.DestinationExpression {
	if procedure.Spec.OperatorSpec == nil {
		return nil
	}
	return procedure.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerStoredProcedure{}

// AzureName returns the Azure name of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) AzureName() string {
	return procedure.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-08-15"
func (procedure SqlDatabaseContainerStoredProcedure) GetAPIVersion() string {
	return "2024-08-15"
}

// GetResourceScope returns the scope of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetSpec() genruntime.ConvertibleSpec {
	return &procedure.Spec
}

// GetStatus returns the status of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetStatus() genruntime.ConvertibleStatus {
	return &procedure.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
func (procedure *SqlDatabaseContainerStoredProcedure) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
}

// NewEmptyStatus returns a new empty (blank) status
func (procedure *SqlDatabaseContainerStoredProcedure) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlDatabaseContainerStoredProcedure_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (procedure *SqlDatabaseContainerStoredProcedure) Owner() *genruntime.ResourceReference {
	if procedure.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(procedure.Spec)
	return procedure.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlDatabaseContainerStoredProcedure_STATUS); ok {
		procedure.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlDatabaseContainerStoredProcedure_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	procedure.Status = st
	return nil
}

// Hub marks that this SqlDatabaseContainerStoredProcedure is the hub type for conversion
func (procedure *SqlDatabaseContainerStoredProcedure) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (procedure *SqlDatabaseContainerStoredProcedure) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: procedure.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerStoredProcedure",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240815.SqlDatabaseContainerStoredProcedure
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures/{storedProcedureName}
type SqlDatabaseContainerStoredProcedureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerStoredProcedure `json:"items"`
}

// Storage version of v1api20240815.SqlDatabaseContainerStoredProcedure_Spec
type SqlDatabaseContainerStoredProcedure_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                           `json:"azureName,omitempty"`
	Location        *string                                          `json:"location,omitempty"`
	OperatorSpec    *SqlDatabaseContainerStoredProcedureOperatorSpec `json:"operatorSpec,omitempty"`
	Options         *CreateUpdateOptions                             `json:"options,omitempty"`
	OriginalVersion string                                           `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *SqlStoredProcedureResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &SqlDatabaseContainerStoredProcedure_Spec{}

// ConvertSpecFrom populates our SqlDatabaseContainerStoredProcedure_Spec from the provided source
func (procedure *SqlDatabaseContainerStoredProcedure_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == procedure {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(procedure)
}

// ConvertSpecTo populates the provided destination from our SqlDatabaseContainerStoredProcedure_Spec
func (procedure *SqlDatabaseContainerStoredProcedure_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == procedure {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(procedure)
}

// Storage version of v1api20240815.SqlDatabaseContainerStoredProcedure_STATUS
type SqlDatabaseContainerStoredProcedure_STATUS struct {
	Conditions  []conditions.Condition                           `json:"conditions,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	Location    *string                                          `json:"location,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	Resource    *SqlStoredProcedureGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                                `json:"tags,omitempty"`
	Type        *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlDatabaseContainerStoredProcedure_STATUS{}

// ConvertStatusFrom populates our SqlDatabaseContainerStoredProcedure_STATUS from the provided source
func (procedure *SqlDatabaseContainerStoredProcedure_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == procedure {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(procedure)
}

// ConvertStatusTo populates the provided destination from our SqlDatabaseContainerStoredProcedure_STATUS
func (procedure *SqlDatabaseContainerStoredProcedure_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == procedure {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(procedure)
}

// Storage version of v1api20240815.SqlDatabaseContainerStoredProcedureOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type SqlDatabaseContainerStoredProcedureOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20240815.SqlStoredProcedureGetProperties_Resource_STATUS
type SqlStoredProcedureGetProperties_Resource_STATUS struct {
	Body        *string                `json:"body,omitempty"`
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
}

// Storage version of v1api20240815.SqlStoredProcedureResource
// Cosmos DB SQL storedProcedure resource object
type SqlStoredProcedureResource struct {
	Body        *string                `json:"body,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerStoredProcedure{}, &SqlDatabaseContainerStoredProcedureList{})
}
