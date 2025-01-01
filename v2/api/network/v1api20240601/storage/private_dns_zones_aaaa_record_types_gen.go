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

// +kubebuilder:rbac:groups=network.azure.com,resources=privatednszonesaaaarecords,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={privatednszonesaaaarecords/status,privatednszonesaaaarecords/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240601.PrivateDnsZonesAAAARecord
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/AAAA/{relativeRecordSetName}
type PrivateDnsZonesAAAARecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDnsZonesAAAARecord_Spec   `json:"spec,omitempty"`
	Status            PrivateDnsZonesAAAARecord_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PrivateDnsZonesAAAARecord{}

// GetConditions returns the conditions of the resource
func (record *PrivateDnsZonesAAAARecord) GetConditions() conditions.Conditions {
	return record.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (record *PrivateDnsZonesAAAARecord) SetConditions(conditions conditions.Conditions) {
	record.Status.Conditions = conditions
}

var _ configmaps.Exporter = &PrivateDnsZonesAAAARecord{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (record *PrivateDnsZonesAAAARecord) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if record.Spec.OperatorSpec == nil {
		return nil
	}
	return record.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &PrivateDnsZonesAAAARecord{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (record *PrivateDnsZonesAAAARecord) SecretDestinationExpressions() []*core.DestinationExpression {
	if record.Spec.OperatorSpec == nil {
		return nil
	}
	return record.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &PrivateDnsZonesAAAARecord{}

// AzureName returns the Azure name of the resource
func (record *PrivateDnsZonesAAAARecord) AzureName() string {
	return record.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-06-01"
func (record PrivateDnsZonesAAAARecord) GetAPIVersion() string {
	return "2024-06-01"
}

// GetResourceScope returns the scope of the resource
func (record *PrivateDnsZonesAAAARecord) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (record *PrivateDnsZonesAAAARecord) GetSpec() genruntime.ConvertibleSpec {
	return &record.Spec
}

// GetStatus returns the status of this resource
func (record *PrivateDnsZonesAAAARecord) GetStatus() genruntime.ConvertibleStatus {
	return &record.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (record *PrivateDnsZonesAAAARecord) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateDnsZones/AAAA"
func (record *PrivateDnsZonesAAAARecord) GetType() string {
	return "Microsoft.Network/privateDnsZones/AAAA"
}

// NewEmptyStatus returns a new empty (blank) status
func (record *PrivateDnsZonesAAAARecord) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PrivateDnsZonesAAAARecord_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (record *PrivateDnsZonesAAAARecord) Owner() *genruntime.ResourceReference {
	if record.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(record.Spec)
	return record.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (record *PrivateDnsZonesAAAARecord) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PrivateDnsZonesAAAARecord_STATUS); ok {
		record.Status = *st
		return nil
	}

	// Convert status to required version
	var st PrivateDnsZonesAAAARecord_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	record.Status = st
	return nil
}

// Hub marks that this PrivateDnsZonesAAAARecord is the hub type for conversion
func (record *PrivateDnsZonesAAAARecord) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (record *PrivateDnsZonesAAAARecord) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: record.Spec.OriginalVersion,
		Kind:    "PrivateDnsZonesAAAARecord",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240601.PrivateDnsZonesAAAARecord
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/AAAA/{relativeRecordSetName}
type PrivateDnsZonesAAAARecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDnsZonesAAAARecord `json:"items"`
}

// Storage version of v1api20240601.PrivateDnsZonesAAAARecord_Spec
type PrivateDnsZonesAAAARecord_Spec struct {
	ARecords    []ARecord    `json:"aRecords,omitempty"`
	AaaaRecords []AaaaRecord `json:"aaaaRecords,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                 `json:"azureName,omitempty"`
	CnameRecord     *CnameRecord                           `json:"cnameRecord,omitempty"`
	Etag            *string                                `json:"etag,omitempty"`
	Metadata        map[string]string                      `json:"metadata,omitempty"`
	MxRecords       []MxRecord                             `json:"mxRecords,omitempty"`
	OperatorSpec    *PrivateDnsZonesAAAARecordOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                 `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/PrivateDnsZone resource
	Owner       *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"PrivateDnsZone"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PtrRecords  []PtrRecord                        `json:"ptrRecords,omitempty"`
	SoaRecord   *SoaRecord                         `json:"soaRecord,omitempty"`
	SrvRecords  []SrvRecord                        `json:"srvRecords,omitempty"`
	Ttl         *int                               `json:"ttl,omitempty"`
	TxtRecords  []TxtRecord                        `json:"txtRecords,omitempty"`
}

var _ genruntime.ConvertibleSpec = &PrivateDnsZonesAAAARecord_Spec{}

// ConvertSpecFrom populates our PrivateDnsZonesAAAARecord_Spec from the provided source
func (record *PrivateDnsZonesAAAARecord_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == record {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(record)
}

// ConvertSpecTo populates the provided destination from our PrivateDnsZonesAAAARecord_Spec
func (record *PrivateDnsZonesAAAARecord_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == record {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(record)
}

// Storage version of v1api20240601.PrivateDnsZonesAAAARecord_STATUS
type PrivateDnsZonesAAAARecord_STATUS struct {
	ARecords         []ARecord_STATUS       `json:"aRecords,omitempty"`
	AaaaRecords      []AaaaRecord_STATUS    `json:"aaaaRecords,omitempty"`
	CnameRecord      *CnameRecord_STATUS    `json:"cnameRecord,omitempty"`
	Conditions       []conditions.Condition `json:"conditions,omitempty"`
	Etag             *string                `json:"etag,omitempty"`
	Fqdn             *string                `json:"fqdn,omitempty"`
	Id               *string                `json:"id,omitempty"`
	IsAutoRegistered *bool                  `json:"isAutoRegistered,omitempty"`
	Metadata         map[string]string      `json:"metadata,omitempty"`
	MxRecords        []MxRecord_STATUS      `json:"mxRecords,omitempty"`
	Name             *string                `json:"name,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PtrRecords       []PtrRecord_STATUS     `json:"ptrRecords,omitempty"`
	SoaRecord        *SoaRecord_STATUS      `json:"soaRecord,omitempty"`
	SrvRecords       []SrvRecord_STATUS     `json:"srvRecords,omitempty"`
	Ttl              *int                   `json:"ttl,omitempty"`
	TxtRecords       []TxtRecord_STATUS     `json:"txtRecords,omitempty"`
	Type             *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &PrivateDnsZonesAAAARecord_STATUS{}

// ConvertStatusFrom populates our PrivateDnsZonesAAAARecord_STATUS from the provided source
func (record *PrivateDnsZonesAAAARecord_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == record {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(record)
}

// ConvertStatusTo populates the provided destination from our PrivateDnsZonesAAAARecord_STATUS
func (record *PrivateDnsZonesAAAARecord_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == record {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(record)
}

// Storage version of v1api20240601.AaaaRecord
// An AAAA record.
type AaaaRecord struct {
	Ipv6Address *string                `json:"ipv6Address,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240601.AaaaRecord_STATUS
// An AAAA record.
type AaaaRecord_STATUS struct {
	Ipv6Address *string                `json:"ipv6Address,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240601.ARecord
// An A record.
type ARecord struct {
	Ipv4Address *string                `json:"ipv4Address,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240601.ARecord_STATUS
// An A record.
type ARecord_STATUS struct {
	Ipv4Address *string                `json:"ipv4Address,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240601.CnameRecord
// A CNAME record.
type CnameRecord struct {
	Cname       *string                `json:"cname,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240601.CnameRecord_STATUS
// A CNAME record.
type CnameRecord_STATUS struct {
	Cname       *string                `json:"cname,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240601.MxRecord
// An MX record.
type MxRecord struct {
	Exchange    *string                `json:"exchange,omitempty"`
	Preference  *int                   `json:"preference,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240601.MxRecord_STATUS
// An MX record.
type MxRecord_STATUS struct {
	Exchange    *string                `json:"exchange,omitempty"`
	Preference  *int                   `json:"preference,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240601.PrivateDnsZonesAAAARecordOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type PrivateDnsZonesAAAARecordOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20240601.PtrRecord
// A PTR record.
type PtrRecord struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Ptrdname    *string                `json:"ptrdname,omitempty"`
}

// Storage version of v1api20240601.PtrRecord_STATUS
// A PTR record.
type PtrRecord_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Ptrdname    *string                `json:"ptrdname,omitempty"`
}

// Storage version of v1api20240601.SoaRecord
// An SOA record.
type SoaRecord struct {
	Email        *string                `json:"email,omitempty"`
	ExpireTime   *int                   `json:"expireTime,omitempty"`
	Host         *string                `json:"host,omitempty"`
	MinimumTtl   *int                   `json:"minimumTtl,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RefreshTime  *int                   `json:"refreshTime,omitempty"`
	RetryTime    *int                   `json:"retryTime,omitempty"`
	SerialNumber *int                   `json:"serialNumber,omitempty"`
}

// Storage version of v1api20240601.SoaRecord_STATUS
// An SOA record.
type SoaRecord_STATUS struct {
	Email        *string                `json:"email,omitempty"`
	ExpireTime   *int                   `json:"expireTime,omitempty"`
	Host         *string                `json:"host,omitempty"`
	MinimumTtl   *int                   `json:"minimumTtl,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RefreshTime  *int                   `json:"refreshTime,omitempty"`
	RetryTime    *int                   `json:"retryTime,omitempty"`
	SerialNumber *int                   `json:"serialNumber,omitempty"`
}

// Storage version of v1api20240601.SrvRecord
// An SRV record.
type SrvRecord struct {
	Port        *int                   `json:"port,omitempty"`
	Priority    *int                   `json:"priority,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Target      *string                `json:"target,omitempty"`
	Weight      *int                   `json:"weight,omitempty"`
}

// Storage version of v1api20240601.SrvRecord_STATUS
// An SRV record.
type SrvRecord_STATUS struct {
	Port        *int                   `json:"port,omitempty"`
	Priority    *int                   `json:"priority,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Target      *string                `json:"target,omitempty"`
	Weight      *int                   `json:"weight,omitempty"`
}

// Storage version of v1api20240601.TxtRecord
// A TXT record.
type TxtRecord struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       []string               `json:"value,omitempty"`
}

// Storage version of v1api20240601.TxtRecord_STATUS
// A TXT record.
type TxtRecord_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       []string               `json:"value,omitempty"`
}

func init() {
	SchemeBuilder.Register(&PrivateDnsZonesAAAARecord{}, &PrivateDnsZonesAAAARecordList{})
}