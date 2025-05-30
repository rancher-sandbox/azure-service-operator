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

// +kubebuilder:rbac:groups=kusto.azure.com,resources=dataconnections,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kusto.azure.com,resources={dataconnections/status,dataconnections/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230815.DataConnection
// Generator information:
// - Generated from: /azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/kusto.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}
type DataConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataConnection_Spec   `json:"spec,omitempty"`
	Status            DataConnection_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DataConnection{}

// GetConditions returns the conditions of the resource
func (connection *DataConnection) GetConditions() conditions.Conditions {
	return connection.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (connection *DataConnection) SetConditions(conditions conditions.Conditions) {
	connection.Status.Conditions = conditions
}

var _ configmaps.Exporter = &DataConnection{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (connection *DataConnection) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if connection.Spec.OperatorSpec == nil {
		return nil
	}
	return connection.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &DataConnection{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (connection *DataConnection) SecretDestinationExpressions() []*core.DestinationExpression {
	if connection.Spec.OperatorSpec == nil {
		return nil
	}
	return connection.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &DataConnection{}

// AzureName returns the Azure name of the resource
func (connection *DataConnection) AzureName() string {
	return connection.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-08-15"
func (connection DataConnection) GetAPIVersion() string {
	return "2023-08-15"
}

// GetResourceScope returns the scope of the resource
func (connection *DataConnection) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (connection *DataConnection) GetSpec() genruntime.ConvertibleSpec {
	return &connection.Spec
}

// GetStatus returns the status of this resource
func (connection *DataConnection) GetStatus() genruntime.ConvertibleStatus {
	return &connection.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (connection *DataConnection) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Kusto/clusters/databases/dataConnections"
func (connection *DataConnection) GetType() string {
	return "Microsoft.Kusto/clusters/databases/dataConnections"
}

// NewEmptyStatus returns a new empty (blank) status
func (connection *DataConnection) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DataConnection_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (connection *DataConnection) Owner() *genruntime.ResourceReference {
	if connection.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(connection.Spec)
	return connection.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (connection *DataConnection) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DataConnection_STATUS); ok {
		connection.Status = *st
		return nil
	}

	// Convert status to required version
	var st DataConnection_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	connection.Status = st
	return nil
}

// Hub marks that this DataConnection is the hub type for conversion
func (connection *DataConnection) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (connection *DataConnection) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: connection.Spec.OriginalVersion,
		Kind:    "DataConnection",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230815.DataConnection
// Generator information:
// - Generated from: /azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/kusto.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}
type DataConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataConnection `json:"items"`
}

// Storage version of v1api20230815.DataConnection_Spec
type DataConnection_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                      `json:"azureName,omitempty"`
	CosmosDb        *CosmosDbDataConnection     `json:"cosmosDbDataConnection,omitempty"`
	EventGrid       *EventGridDataConnection    `json:"eventGridDataConnection,omitempty"`
	EventHub        *EventHubDataConnection     `json:"eventHubDataConnection,omitempty"`
	IotHub          *IotHubDataConnection       `json:"iotHubDataConnection,omitempty"`
	OperatorSpec    *DataConnectionOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                      `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a kusto.azure.com/Database resource
	Owner       *genruntime.KnownResourceReference `group:"kusto.azure.com" json:"owner,omitempty" kind:"Database"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DataConnection_Spec{}

// ConvertSpecFrom populates our DataConnection_Spec from the provided source
func (connection *DataConnection_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == connection {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(connection)
}

// ConvertSpecTo populates the provided destination from our DataConnection_Spec
func (connection *DataConnection_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == connection {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(connection)
}

// Storage version of v1api20230815.DataConnection_STATUS
type DataConnection_STATUS struct {
	Conditions  []conditions.Condition          `json:"conditions,omitempty"`
	CosmosDb    *CosmosDbDataConnection_STATUS  `json:"cosmosDb,omitempty"`
	EventGrid   *EventGridDataConnection_STATUS `json:"eventGrid,omitempty"`
	EventHub    *EventHubDataConnection_STATUS  `json:"eventHub,omitempty"`
	IotHub      *IotHubDataConnection_STATUS    `json:"iotHub,omitempty"`
	Name        *string                         `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DataConnection_STATUS{}

// ConvertStatusFrom populates our DataConnection_STATUS from the provided source
func (connection *DataConnection_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == connection {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(connection)
}

// ConvertStatusTo populates the provided destination from our DataConnection_STATUS
func (connection *DataConnection_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == connection {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(connection)
}

// Storage version of v1api20230815.CosmosDbDataConnection
type CosmosDbDataConnection struct {
	// +kubebuilder:validation:Required
	// CosmosDbAccountResourceReference: The resource ID of the Cosmos DB account used to create the data connection.
	CosmosDbAccountResourceReference *genruntime.ResourceReference `armReference:"CosmosDbAccountResourceId" json:"cosmosDbAccountResourceReference,omitempty"`
	CosmosDbContainer                *string                       `json:"cosmosDbContainer,omitempty"`
	CosmosDbDatabase                 *string                       `json:"cosmosDbDatabase,omitempty"`
	Kind                             *string                       `json:"kind,omitempty"`
	Location                         *string                       `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// ManagedIdentityResourceReference: The resource ID of a managed system or user-assigned identity. The identity is used to
	// authenticate with Cosmos DB.
	ManagedIdentityResourceReference *genruntime.ResourceReference `armReference:"ManagedIdentityResourceId" json:"managedIdentityResourceReference,omitempty"`
	MappingRuleName                  *string                       `json:"mappingRuleName,omitempty"`
	PropertyBag                      genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RetrievalStartDate               *string                       `json:"retrievalStartDate,omitempty"`
	TableName                        *string                       `json:"tableName,omitempty"`
}

// Storage version of v1api20230815.CosmosDbDataConnection_STATUS
type CosmosDbDataConnection_STATUS struct {
	CosmosDbAccountResourceId *string                `json:"cosmosDbAccountResourceId,omitempty"`
	CosmosDbContainer         *string                `json:"cosmosDbContainer,omitempty"`
	CosmosDbDatabase          *string                `json:"cosmosDbDatabase,omitempty"`
	Id                        *string                `json:"id,omitempty"`
	Kind                      *string                `json:"kind,omitempty"`
	Location                  *string                `json:"location,omitempty"`
	ManagedIdentityObjectId   *string                `json:"managedIdentityObjectId,omitempty"`
	ManagedIdentityResourceId *string                `json:"managedIdentityResourceId,omitempty"`
	MappingRuleName           *string                `json:"mappingRuleName,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                `json:"provisioningState,omitempty"`
	RetrievalStartDate        *string                `json:"retrievalStartDate,omitempty"`
	TableName                 *string                `json:"tableName,omitempty"`
	Type                      *string                `json:"type,omitempty"`
}

// Storage version of v1api20230815.DataConnectionOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type DataConnectionOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20230815.EventGridDataConnection
type EventGridDataConnection struct {
	BlobStorageEventType *string `json:"blobStorageEventType,omitempty"`
	ConsumerGroup        *string `json:"consumerGroup,omitempty"`
	DataFormat           *string `json:"dataFormat,omitempty"`
	DatabaseRouting      *string `json:"databaseRouting,omitempty"`

	// EventGridResourceReference: The resource ID of the event grid that is subscribed to the storage account events.
	EventGridResourceReference *genruntime.ResourceReference `armReference:"EventGridResourceId" json:"eventGridResourceReference,omitempty"`

	// +kubebuilder:validation:Required
	// EventHubResourceReference: The resource ID where the event grid is configured to send events.
	EventHubResourceReference *genruntime.ResourceReference `armReference:"EventHubResourceId" json:"eventHubResourceReference,omitempty"`
	IgnoreFirstRecord         *bool                         `json:"ignoreFirstRecord,omitempty"`
	Kind                      *string                       `json:"kind,omitempty"`
	Location                  *string                       `json:"location,omitempty"`

	// ManagedIdentityResourceReference: The resource ID of a managed identity (system or user assigned) to be used to
	// authenticate with event hub and storage account.
	ManagedIdentityResourceReference *genruntime.ResourceReference `armReference:"ManagedIdentityResourceId" json:"managedIdentityResourceReference,omitempty"`
	MappingRuleName                  *string                       `json:"mappingRuleName,omitempty"`
	PropertyBag                      genruntime.PropertyBag        `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// StorageAccountResourceReference: The resource ID of the storage account where the data resides.
	StorageAccountResourceReference *genruntime.ResourceReference `armReference:"StorageAccountResourceId" json:"storageAccountResourceReference,omitempty"`
	TableName                       *string                       `json:"tableName,omitempty"`
}

// Storage version of v1api20230815.EventGridDataConnection_STATUS
type EventGridDataConnection_STATUS struct {
	BlobStorageEventType      *string                `json:"blobStorageEventType,omitempty"`
	ConsumerGroup             *string                `json:"consumerGroup,omitempty"`
	DataFormat                *string                `json:"dataFormat,omitempty"`
	DatabaseRouting           *string                `json:"databaseRouting,omitempty"`
	EventGridResourceId       *string                `json:"eventGridResourceId,omitempty"`
	EventHubResourceId        *string                `json:"eventHubResourceId,omitempty"`
	Id                        *string                `json:"id,omitempty"`
	IgnoreFirstRecord         *bool                  `json:"ignoreFirstRecord,omitempty"`
	Kind                      *string                `json:"kind,omitempty"`
	Location                  *string                `json:"location,omitempty"`
	ManagedIdentityObjectId   *string                `json:"managedIdentityObjectId,omitempty"`
	ManagedIdentityResourceId *string                `json:"managedIdentityResourceId,omitempty"`
	MappingRuleName           *string                `json:"mappingRuleName,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                `json:"provisioningState,omitempty"`
	StorageAccountResourceId  *string                `json:"storageAccountResourceId,omitempty"`
	TableName                 *string                `json:"tableName,omitempty"`
	Type                      *string                `json:"type,omitempty"`
}

// Storage version of v1api20230815.EventHubDataConnection
type EventHubDataConnection struct {
	Compression     *string `json:"compression,omitempty"`
	ConsumerGroup   *string `json:"consumerGroup,omitempty"`
	DataFormat      *string `json:"dataFormat,omitempty"`
	DatabaseRouting *string `json:"databaseRouting,omitempty"`

	// +kubebuilder:validation:Required
	// EventHubResourceReference: The resource ID of the event hub to be used to create a data connection.
	EventHubResourceReference *genruntime.ResourceReference `armReference:"EventHubResourceId" json:"eventHubResourceReference,omitempty"`
	EventSystemProperties     []string                      `json:"eventSystemProperties,omitempty"`
	Kind                      *string                       `json:"kind,omitempty"`
	Location                  *string                       `json:"location,omitempty"`

	// ManagedIdentityResourceReference: The resource ID of a managed identity (system or user assigned) to be used to
	// authenticate with event hub.
	ManagedIdentityResourceReference *genruntime.ResourceReference `armReference:"ManagedIdentityResourceId" json:"managedIdentityResourceReference,omitempty"`
	MappingRuleName                  *string                       `json:"mappingRuleName,omitempty"`
	PropertyBag                      genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RetrievalStartDate               *string                       `json:"retrievalStartDate,omitempty"`
	TableName                        *string                       `json:"tableName,omitempty"`
}

// Storage version of v1api20230815.EventHubDataConnection_STATUS
type EventHubDataConnection_STATUS struct {
	Compression               *string                `json:"compression,omitempty"`
	ConsumerGroup             *string                `json:"consumerGroup,omitempty"`
	DataFormat                *string                `json:"dataFormat,omitempty"`
	DatabaseRouting           *string                `json:"databaseRouting,omitempty"`
	EventHubResourceId        *string                `json:"eventHubResourceId,omitempty"`
	EventSystemProperties     []string               `json:"eventSystemProperties,omitempty"`
	Id                        *string                `json:"id,omitempty"`
	Kind                      *string                `json:"kind,omitempty"`
	Location                  *string                `json:"location,omitempty"`
	ManagedIdentityObjectId   *string                `json:"managedIdentityObjectId,omitempty"`
	ManagedIdentityResourceId *string                `json:"managedIdentityResourceId,omitempty"`
	MappingRuleName           *string                `json:"mappingRuleName,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                `json:"provisioningState,omitempty"`
	RetrievalStartDate        *string                `json:"retrievalStartDate,omitempty"`
	TableName                 *string                `json:"tableName,omitempty"`
	Type                      *string                `json:"type,omitempty"`
}

// Storage version of v1api20230815.IotHubDataConnection
type IotHubDataConnection struct {
	ConsumerGroup         *string  `json:"consumerGroup,omitempty"`
	DataFormat            *string  `json:"dataFormat,omitempty"`
	DatabaseRouting       *string  `json:"databaseRouting,omitempty"`
	EventSystemProperties []string `json:"eventSystemProperties,omitempty"`

	// +kubebuilder:validation:Required
	// IotHubResourceReference: The resource ID of the Iot hub to be used to create a data connection.
	IotHubResourceReference *genruntime.ResourceReference `armReference:"IotHubResourceId" json:"iotHubResourceReference,omitempty"`
	Kind                    *string                       `json:"kind,omitempty"`
	Location                *string                       `json:"location,omitempty"`
	MappingRuleName         *string                       `json:"mappingRuleName,omitempty"`
	PropertyBag             genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	RetrievalStartDate      *string                       `json:"retrievalStartDate,omitempty"`
	SharedAccessPolicyName  *string                       `json:"sharedAccessPolicyName,omitempty"`
	TableName               *string                       `json:"tableName,omitempty"`
}

// Storage version of v1api20230815.IotHubDataConnection_STATUS
type IotHubDataConnection_STATUS struct {
	ConsumerGroup          *string                `json:"consumerGroup,omitempty"`
	DataFormat             *string                `json:"dataFormat,omitempty"`
	DatabaseRouting        *string                `json:"databaseRouting,omitempty"`
	EventSystemProperties  []string               `json:"eventSystemProperties,omitempty"`
	Id                     *string                `json:"id,omitempty"`
	IotHubResourceId       *string                `json:"iotHubResourceId,omitempty"`
	Kind                   *string                `json:"kind,omitempty"`
	Location               *string                `json:"location,omitempty"`
	MappingRuleName        *string                `json:"mappingRuleName,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState      *string                `json:"provisioningState,omitempty"`
	RetrievalStartDate     *string                `json:"retrievalStartDate,omitempty"`
	SharedAccessPolicyName *string                `json:"sharedAccessPolicyName,omitempty"`
	TableName              *string                `json:"tableName,omitempty"`
	Type                   *string                `json:"type,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DataConnection{}, &DataConnectionList{})
}
