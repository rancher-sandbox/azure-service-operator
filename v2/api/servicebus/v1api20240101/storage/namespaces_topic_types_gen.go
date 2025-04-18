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

// +kubebuilder:rbac:groups=servicebus.azure.com,resources=namespacestopics,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=servicebus.azure.com,resources={namespacestopics/status,namespacestopics/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240101.NamespacesTopic
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/stable/2024-01-01/topics.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}
type NamespacesTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesTopic_Spec   `json:"spec,omitempty"`
	Status            NamespacesTopic_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesTopic{}

// GetConditions returns the conditions of the resource
func (topic *NamespacesTopic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *NamespacesTopic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ configmaps.Exporter = &NamespacesTopic{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (topic *NamespacesTopic) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if topic.Spec.OperatorSpec == nil {
		return nil
	}
	return topic.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &NamespacesTopic{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (topic *NamespacesTopic) SecretDestinationExpressions() []*core.DestinationExpression {
	if topic.Spec.OperatorSpec == nil {
		return nil
	}
	return topic.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &NamespacesTopic{}

// AzureName returns the Azure name of the resource
func (topic *NamespacesTopic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-01-01"
func (topic NamespacesTopic) GetAPIVersion() string {
	return "2024-01-01"
}

// GetResourceScope returns the scope of the resource
func (topic *NamespacesTopic) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (topic *NamespacesTopic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *NamespacesTopic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (topic *NamespacesTopic) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics"
func (topic *NamespacesTopic) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *NamespacesTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NamespacesTopic_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (topic *NamespacesTopic) Owner() *genruntime.ResourceReference {
	if topic.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return topic.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (topic *NamespacesTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NamespacesTopic_STATUS); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st NamespacesTopic_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// Hub marks that this NamespacesTopic is the hub type for conversion
func (topic *NamespacesTopic) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *NamespacesTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion,
		Kind:    "NamespacesTopic",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240101.NamespacesTopic
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/stable/2024-01-01/topics.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}
type NamespacesTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesTopic `json:"items"`
}

// Storage version of v1api20240101.NamespacesTopic_Spec
type NamespacesTopic_Spec struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                           string                       `json:"azureName,omitempty"`
	DefaultMessageTimeToLive            *string                      `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                      `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                        `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                        `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                        `json:"enablePartitioning,omitempty"`
	MaxMessageSizeInKilobytes           *int                         `json:"maxMessageSizeInKilobytes,omitempty"`
	MaxSizeInMegabytes                  *int                         `json:"maxSizeInMegabytes,omitempty"`
	OperatorSpec                        *NamespacesTopicOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion                     string                       `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a servicebus.azure.com/Namespace resource
	Owner                      *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection *bool                              `json:"requiresDuplicateDetection,omitempty"`
	SupportOrdering            *bool                              `json:"supportOrdering,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesTopic_Spec{}

// ConvertSpecFrom populates our NamespacesTopic_Spec from the provided source
func (topic *NamespacesTopic_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == topic {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(topic)
}

// ConvertSpecTo populates the provided destination from our NamespacesTopic_Spec
func (topic *NamespacesTopic_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == topic {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(topic)
}

// Storage version of v1api20240101.NamespacesTopic_STATUS
type NamespacesTopic_STATUS struct {
	AccessedAt                          *string                     `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                     `json:"autoDeleteOnIdle,omitempty"`
	Conditions                          []conditions.Condition      `json:"conditions,omitempty"`
	CountDetails                        *MessageCountDetails_STATUS `json:"countDetails,omitempty"`
	CreatedAt                           *string                     `json:"createdAt,omitempty"`
	DefaultMessageTimeToLive            *string                     `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                     `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                       `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                       `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                       `json:"enablePartitioning,omitempty"`
	Id                                  *string                     `json:"id,omitempty"`
	Location                            *string                     `json:"location,omitempty"`
	MaxMessageSizeInKilobytes           *int                        `json:"maxMessageSizeInKilobytes,omitempty"`
	MaxSizeInMegabytes                  *int                        `json:"maxSizeInMegabytes,omitempty"`
	Name                                *string                     `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection          *bool                       `json:"requiresDuplicateDetection,omitempty"`
	SizeInBytes                         *int                        `json:"sizeInBytes,omitempty"`
	Status                              *string                     `json:"status,omitempty"`
	SubscriptionCount                   *int                        `json:"subscriptionCount,omitempty"`
	SupportOrdering                     *bool                       `json:"supportOrdering,omitempty"`
	SystemData                          *SystemData_STATUS          `json:"systemData,omitempty"`
	Type                                *string                     `json:"type,omitempty"`
	UpdatedAt                           *string                     `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NamespacesTopic_STATUS{}

// ConvertStatusFrom populates our NamespacesTopic_STATUS from the provided source
func (topic *NamespacesTopic_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == topic {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(topic)
}

// ConvertStatusTo populates the provided destination from our NamespacesTopic_STATUS
func (topic *NamespacesTopic_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == topic {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(topic)
}

// Storage version of v1api20240101.NamespacesTopicOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type NamespacesTopicOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NamespacesTopic{}, &NamespacesTopicList{})
}
