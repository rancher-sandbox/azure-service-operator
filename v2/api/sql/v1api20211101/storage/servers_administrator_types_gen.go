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

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversadministrators,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversadministrators/status,serversadministrators/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.ServersAdministrator
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAzureADAdministrators.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}
type ServersAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServersAdministrator_Spec   `json:"spec,omitempty"`
	Status            ServersAdministrator_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersAdministrator{}

// GetConditions returns the conditions of the resource
func (administrator *ServersAdministrator) GetConditions() conditions.Conditions {
	return administrator.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (administrator *ServersAdministrator) SetConditions(conditions conditions.Conditions) {
	administrator.Status.Conditions = conditions
}

var _ configmaps.Exporter = &ServersAdministrator{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (administrator *ServersAdministrator) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if administrator.Spec.OperatorSpec == nil {
		return nil
	}
	return administrator.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ServersAdministrator{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (administrator *ServersAdministrator) SecretDestinationExpressions() []*core.DestinationExpression {
	if administrator.Spec.OperatorSpec == nil {
		return nil
	}
	return administrator.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &ServersAdministrator{}

// AzureName returns the Azure name of the resource (always "ActiveDirectory")
func (administrator *ServersAdministrator) AzureName() string {
	return "ActiveDirectory"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (administrator ServersAdministrator) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceScope returns the scope of the resource
func (administrator *ServersAdministrator) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (administrator *ServersAdministrator) GetSpec() genruntime.ConvertibleSpec {
	return &administrator.Spec
}

// GetStatus returns the status of this resource
func (administrator *ServersAdministrator) GetStatus() genruntime.ConvertibleStatus {
	return &administrator.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (administrator *ServersAdministrator) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/administrators"
func (administrator *ServersAdministrator) GetType() string {
	return "Microsoft.Sql/servers/administrators"
}

// NewEmptyStatus returns a new empty (blank) status
func (administrator *ServersAdministrator) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ServersAdministrator_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (administrator *ServersAdministrator) Owner() *genruntime.ResourceReference {
	if administrator.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(administrator.Spec)
	return administrator.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (administrator *ServersAdministrator) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ServersAdministrator_STATUS); ok {
		administrator.Status = *st
		return nil
	}

	// Convert status to required version
	var st ServersAdministrator_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	administrator.Status = st
	return nil
}

// Hub marks that this ServersAdministrator is the hub type for conversion
func (administrator *ServersAdministrator) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (administrator *ServersAdministrator) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: administrator.Spec.OriginalVersion,
		Kind:    "ServersAdministrator",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.ServersAdministrator
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAzureADAdministrators.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}
type ServersAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersAdministrator `json:"items"`
}

// Storage version of v1api20211101.ServersAdministrator_Spec
type ServersAdministrator_Spec struct {
	AdministratorType *string                           `json:"administratorType,omitempty"`
	Login             *string                           `json:"login,omitempty"`
	OperatorSpec      *ServersAdministratorOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion   string                            `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner              *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`
	PropertyBag        genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sid                *string                            `json:"sid,omitempty" optionalConfigMapPair:"Sid"`
	SidFromConfig      *genruntime.ConfigMapReference     `json:"sidFromConfig,omitempty" optionalConfigMapPair:"Sid"`
	TenantId           *string                            `json:"tenantId,omitempty" optionalConfigMapPair:"TenantId"`
	TenantIdFromConfig *genruntime.ConfigMapReference     `json:"tenantIdFromConfig,omitempty" optionalConfigMapPair:"TenantId"`
}

var _ genruntime.ConvertibleSpec = &ServersAdministrator_Spec{}

// ConvertSpecFrom populates our ServersAdministrator_Spec from the provided source
func (administrator *ServersAdministrator_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == administrator {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(administrator)
}

// ConvertSpecTo populates the provided destination from our ServersAdministrator_Spec
func (administrator *ServersAdministrator_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == administrator {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(administrator)
}

// Storage version of v1api20211101.ServersAdministrator_STATUS
type ServersAdministrator_STATUS struct {
	AdministratorType         *string                `json:"administratorType,omitempty"`
	AzureADOnlyAuthentication *bool                  `json:"azureADOnlyAuthentication,omitempty"`
	Conditions                []conditions.Condition `json:"conditions,omitempty"`
	Id                        *string                `json:"id,omitempty"`
	Login                     *string                `json:"login,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Sid                       *string                `json:"sid,omitempty"`
	TenantId                  *string                `json:"tenantId,omitempty"`
	Type                      *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ServersAdministrator_STATUS{}

// ConvertStatusFrom populates our ServersAdministrator_STATUS from the provided source
func (administrator *ServersAdministrator_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == administrator {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(administrator)
}

// ConvertStatusTo populates the provided destination from our ServersAdministrator_STATUS
func (administrator *ServersAdministrator_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == administrator {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(administrator)
}

// Storage version of v1api20211101.ServersAdministratorOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ServersAdministratorOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ServersAdministrator{}, &ServersAdministratorList{})
}
