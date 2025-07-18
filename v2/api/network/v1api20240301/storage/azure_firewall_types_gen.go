// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=azurefirewalls,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={azurefirewalls/status,azurefirewalls/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20240301.AzureFirewall
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2024-03-01/azureFirewall.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}
type AzureFirewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzureFirewall_Spec   `json:"spec,omitempty"`
	Status            AzureFirewall_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &AzureFirewall{}

// GetConditions returns the conditions of the resource
func (firewall *AzureFirewall) GetConditions() conditions.Conditions {
	return firewall.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (firewall *AzureFirewall) SetConditions(conditions conditions.Conditions) {
	firewall.Status.Conditions = conditions
}

var _ configmaps.Exporter = &AzureFirewall{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (firewall *AzureFirewall) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if firewall.Spec.OperatorSpec == nil {
		return nil
	}
	return firewall.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &AzureFirewall{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (firewall *AzureFirewall) SecretDestinationExpressions() []*core.DestinationExpression {
	if firewall.Spec.OperatorSpec == nil {
		return nil
	}
	return firewall.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &AzureFirewall{}

// AzureName returns the Azure name of the resource
func (firewall *AzureFirewall) AzureName() string {
	return firewall.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-03-01"
func (firewall AzureFirewall) GetAPIVersion() string {
	return "2024-03-01"
}

// GetResourceScope returns the scope of the resource
func (firewall *AzureFirewall) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (firewall *AzureFirewall) GetSpec() genruntime.ConvertibleSpec {
	return &firewall.Spec
}

// GetStatus returns the status of this resource
func (firewall *AzureFirewall) GetStatus() genruntime.ConvertibleStatus {
	return &firewall.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (firewall *AzureFirewall) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/azureFirewalls"
func (firewall *AzureFirewall) GetType() string {
	return "Microsoft.Network/azureFirewalls"
}

// NewEmptyStatus returns a new empty (blank) status
func (firewall *AzureFirewall) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &AzureFirewall_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (firewall *AzureFirewall) Owner() *genruntime.ResourceReference {
	if firewall.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(firewall.Spec)
	return firewall.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (firewall *AzureFirewall) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*AzureFirewall_STATUS); ok {
		firewall.Status = *st
		return nil
	}

	// Convert status to required version
	var st AzureFirewall_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	firewall.Status = st
	return nil
}

// Hub marks that this AzureFirewall is the hub type for conversion
func (firewall *AzureFirewall) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (firewall *AzureFirewall) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: firewall.Spec.OriginalVersion,
		Kind:    "AzureFirewall",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20240301.AzureFirewall
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2024-03-01/azureFirewall.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}
type AzureFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureFirewall `json:"items"`
}

// Storage version of v1api20240301.APIVersion
// +kubebuilder:validation:Enum={"2024-03-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2024-03-01")

// Storage version of v1api20240301.AzureFirewall_Spec
type AzureFirewall_Spec struct {
	AdditionalProperties       map[string]string                        `json:"additionalProperties,omitempty"`
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollection `json:"applicationRuleCollections,omitempty"`
	AutoscaleConfiguration     *AzureFirewallAutoscaleConfiguration     `json:"autoscaleConfiguration,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                 string                               `json:"azureName,omitempty"`
	FirewallPolicy            *SubResource                         `json:"firewallPolicy,omitempty"`
	HubIPAddresses            *HubIPAddresses                      `json:"hubIPAddresses,omitempty"`
	IpConfigurations          []AzureFirewallIPConfiguration       `json:"ipConfigurations,omitempty"`
	Location                  *string                              `json:"location,omitempty"`
	ManagementIpConfiguration *AzureFirewallIPConfiguration        `json:"managementIpConfiguration,omitempty"`
	NatRuleCollections        []AzureFirewallNatRuleCollection     `json:"natRuleCollections,omitempty"`
	NetworkRuleCollections    []AzureFirewallNetworkRuleCollection `json:"networkRuleCollections,omitempty"`
	OperatorSpec              *AzureFirewallOperatorSpec           `json:"operatorSpec,omitempty"`
	OriginalVersion           string                               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner           *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku             *AzureFirewallSku                  `json:"sku,omitempty"`
	Tags            map[string]string                  `json:"tags,omitempty"`
	ThreatIntelMode *string                            `json:"threatIntelMode,omitempty"`
	VirtualHub      *SubResource                       `json:"virtualHub,omitempty"`
	Zones           []string                           `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &AzureFirewall_Spec{}

// ConvertSpecFrom populates our AzureFirewall_Spec from the provided source
func (firewall *AzureFirewall_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == firewall {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(firewall)
}

// ConvertSpecTo populates the provided destination from our AzureFirewall_Spec
func (firewall *AzureFirewall_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == firewall {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(firewall)
}

// Storage version of v1api20240301.AzureFirewall_STATUS
// Azure Firewall resource.
type AzureFirewall_STATUS struct {
	AdditionalProperties       map[string]string                               `json:"additionalProperties,omitempty"`
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollection_STATUS `json:"applicationRuleCollections,omitempty"`
	AutoscaleConfiguration     *AzureFirewallAutoscaleConfiguration_STATUS     `json:"autoscaleConfiguration,omitempty"`
	Conditions                 []conditions.Condition                          `json:"conditions,omitempty"`
	Etag                       *string                                         `json:"etag,omitempty"`
	FirewallPolicy             *SubResource_STATUS                             `json:"firewallPolicy,omitempty"`
	HubIPAddresses             *HubIPAddresses_STATUS                          `json:"hubIPAddresses,omitempty"`
	Id                         *string                                         `json:"id,omitempty"`
	IpConfigurations           []AzureFirewallIPConfiguration_STATUS           `json:"ipConfigurations,omitempty"`
	IpGroups                   []AzureFirewallIpGroups_STATUS                  `json:"ipGroups,omitempty"`
	Location                   *string                                         `json:"location,omitempty"`
	ManagementIpConfiguration  *AzureFirewallIPConfiguration_STATUS            `json:"managementIpConfiguration,omitempty"`
	Name                       *string                                         `json:"name,omitempty"`
	NatRuleCollections         []AzureFirewallNatRuleCollection_STATUS         `json:"natRuleCollections,omitempty"`
	NetworkRuleCollections     []AzureFirewallNetworkRuleCollection_STATUS     `json:"networkRuleCollections,omitempty"`
	PropertyBag                genruntime.PropertyBag                          `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                         `json:"provisioningState,omitempty"`
	Sku                        *AzureFirewallSku_STATUS                        `json:"sku,omitempty"`
	Tags                       map[string]string                               `json:"tags,omitempty"`
	ThreatIntelMode            *string                                         `json:"threatIntelMode,omitempty"`
	Type                       *string                                         `json:"type,omitempty"`
	VirtualHub                 *SubResource_STATUS                             `json:"virtualHub,omitempty"`
	Zones                      []string                                        `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &AzureFirewall_STATUS{}

// ConvertStatusFrom populates our AzureFirewall_STATUS from the provided source
func (firewall *AzureFirewall_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == firewall {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(firewall)
}

// ConvertStatusTo populates the provided destination from our AzureFirewall_STATUS
func (firewall *AzureFirewall_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == firewall {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(firewall)
}

// Storage version of v1api20240301.AzureFirewallApplicationRuleCollection
// Application rule collection resource.
type AzureFirewallApplicationRuleCollection struct {
	Action      *AzureFirewallRCAction         `json:"action,omitempty"`
	Name        *string                        `json:"name,omitempty"`
	Priority    *int                           `json:"priority,omitempty"`
	PropertyBag genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	Rules       []AzureFirewallApplicationRule `json:"rules,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallApplicationRuleCollection_STATUS
// Application rule collection resource.
type AzureFirewallApplicationRuleCollection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallAutoscaleConfiguration
// Azure Firewall Autoscale Configuration parameters.
type AzureFirewallAutoscaleConfiguration struct {
	MaxCapacity *int                   `json:"maxCapacity,omitempty"`
	MinCapacity *int                   `json:"minCapacity,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallAutoscaleConfiguration_STATUS
// Azure Firewall Autoscale Configuration parameters.
type AzureFirewallAutoscaleConfiguration_STATUS struct {
	MaxCapacity *int                   `json:"maxCapacity,omitempty"`
	MinCapacity *int                   `json:"minCapacity,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallIPConfiguration
// IP configuration of an Azure Firewall.
type AzureFirewallIPConfiguration struct {
	Name            *string                `json:"name,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicIPAddress *SubResource           `json:"publicIPAddress,omitempty"`
	Subnet          *SubResource           `json:"subnet,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallIPConfiguration_STATUS
// IP configuration of an Azure Firewall.
type AzureFirewallIPConfiguration_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallIpGroups_STATUS
// IpGroups associated with azure firewall.
type AzureFirewallIpGroups_STATUS struct {
	ChangeNumber *string                `json:"changeNumber,omitempty"`
	Id           *string                `json:"id,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallNatRuleCollection
// NAT rule collection resource.
type AzureFirewallNatRuleCollection struct {
	Action      *AzureFirewallNatRCAction `json:"action,omitempty"`
	Name        *string                   `json:"name,omitempty"`
	Priority    *int                      `json:"priority,omitempty"`
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Rules       []AzureFirewallNatRule    `json:"rules,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallNatRuleCollection_STATUS
// NAT rule collection resource.
type AzureFirewallNatRuleCollection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallNetworkRuleCollection
// Network rule collection resource.
type AzureFirewallNetworkRuleCollection struct {
	Action      *AzureFirewallRCAction     `json:"action,omitempty"`
	Name        *string                    `json:"name,omitempty"`
	Priority    *int                       `json:"priority,omitempty"`
	PropertyBag genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Rules       []AzureFirewallNetworkRule `json:"rules,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallNetworkRuleCollection_STATUS
// Network rule collection resource.
type AzureFirewallNetworkRuleCollection_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type AzureFirewallOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallSku
// SKU of an Azure Firewall.
type AzureFirewallSku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallSku_STATUS
// SKU of an Azure Firewall.
type AzureFirewallSku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20240301.HubIPAddresses
// IP addresses associated with azure firewall.
type HubIPAddresses struct {
	PrivateIPAddress *string                `json:"privateIPAddress,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicIPs        *HubPublicIPAddresses  `json:"publicIPs,omitempty"`
}

// Storage version of v1api20240301.HubIPAddresses_STATUS
// IP addresses associated with azure firewall.
type HubIPAddresses_STATUS struct {
	PrivateIPAddress *string                      `json:"privateIPAddress,omitempty"`
	PropertyBag      genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	PublicIPs        *HubPublicIPAddresses_STATUS `json:"publicIPs,omitempty"`
}

// Storage version of v1api20240301.SubResource
// Reference to another subresource.
type SubResource struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// AssignProperties_From_SubResource populates our SubResource from the provided source SubResource
func (resource *SubResource) AssignProperties_From_SubResource(source *storage.SubResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Reference
	if source.Reference != nil {
		reference := source.Reference.Copy()
		resource.Reference = &reference
	} else {
		resource.Reference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SubResource populates the provided destination SubResource from our SubResource
func (resource *SubResource) AssignProperties_To_SubResource(destination *storage.SubResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Reference
	if resource.Reference != nil {
		reference := resource.Reference.Copy()
		destination.Reference = &reference
	} else {
		destination.Reference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20240301.SubResource_STATUS
// Reference to another subresource.
type SubResource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_SubResource_STATUS populates our SubResource_STATUS from the provided source SubResource_STATUS
func (resource *SubResource_STATUS) AssignProperties_From_SubResource_STATUS(source *storage.SubResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource_STATUS); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SubResource_STATUS populates the provided destination SubResource_STATUS from our SubResource_STATUS
func (resource *SubResource_STATUS) AssignProperties_To_SubResource_STATUS(destination *storage.SubResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource_STATUS); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSubResource interface {
	AssignPropertiesFrom(src *storage.SubResource) error
	AssignPropertiesTo(dst *storage.SubResource) error
}

type augmentConversionForSubResource_STATUS interface {
	AssignPropertiesFrom(src *storage.SubResource_STATUS) error
	AssignPropertiesTo(dst *storage.SubResource_STATUS) error
}

// Storage version of v1api20240301.AzureFirewallApplicationRule
// Properties of an application rule.
type AzureFirewallApplicationRule struct {
	Description     *string                                `json:"description,omitempty"`
	FqdnTags        []string                               `json:"fqdnTags,omitempty"`
	Name            *string                                `json:"name,omitempty"`
	PropertyBag     genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	Protocols       []AzureFirewallApplicationRuleProtocol `json:"protocols,omitempty"`
	SourceAddresses []string                               `json:"sourceAddresses,omitempty"`
	SourceIpGroups  []string                               `json:"sourceIpGroups,omitempty"`
	TargetFqdns     []string                               `json:"targetFqdns,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallNatRCAction
// AzureFirewall NAT Rule Collection Action.
type AzureFirewallNatRCAction struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallNatRule
// Properties of a NAT rule.
type AzureFirewallNatRule struct {
	Description          *string                `json:"description,omitempty"`
	DestinationAddresses []string               `json:"destinationAddresses,omitempty"`
	DestinationPorts     []string               `json:"destinationPorts,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocols            []string               `json:"protocols,omitempty"`
	SourceAddresses      []string               `json:"sourceAddresses,omitempty"`
	SourceIpGroups       []string               `json:"sourceIpGroups,omitempty"`
	TranslatedAddress    *string                `json:"translatedAddress,omitempty"`
	TranslatedFqdn       *string                `json:"translatedFqdn,omitempty"`
	TranslatedPort       *string                `json:"translatedPort,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallNetworkRule
// Properties of the network rule.
type AzureFirewallNetworkRule struct {
	Description          *string                `json:"description,omitempty"`
	DestinationAddresses []string               `json:"destinationAddresses,omitempty"`
	DestinationFqdns     []string               `json:"destinationFqdns,omitempty"`
	DestinationIpGroups  []string               `json:"destinationIpGroups,omitempty"`
	DestinationPorts     []string               `json:"destinationPorts,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocols            []string               `json:"protocols,omitempty"`
	SourceAddresses      []string               `json:"sourceAddresses,omitempty"`
	SourceIpGroups       []string               `json:"sourceIpGroups,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallRCAction
// Properties of the AzureFirewallRCAction.
type AzureFirewallRCAction struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20240301.HubPublicIPAddresses
// Public IP addresses associated with azure firewall.
type HubPublicIPAddresses struct {
	Addresses   []AzureFirewallPublicIPAddress `json:"addresses,omitempty"`
	Count       *int                           `json:"count,omitempty"`
	PropertyBag genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.HubPublicIPAddresses_STATUS
// Public IP addresses associated with azure firewall.
type HubPublicIPAddresses_STATUS struct {
	Addresses   []AzureFirewallPublicIPAddress_STATUS `json:"addresses,omitempty"`
	Count       *int                                  `json:"count,omitempty"`
	PropertyBag genruntime.PropertyBag                `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallApplicationRuleProtocol
// Properties of the application rule protocol.
type AzureFirewallApplicationRuleProtocol struct {
	Port         *int                   `json:"port,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProtocolType *string                `json:"protocolType,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallPublicIPAddress
// Public IP Address associated with azure firewall.
type AzureFirewallPublicIPAddress struct {
	Address     *string                `json:"address,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20240301.AzureFirewallPublicIPAddress_STATUS
// Public IP Address associated with azure firewall.
type AzureFirewallPublicIPAddress_STATUS struct {
	Address     *string                `json:"address,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&AzureFirewall{}, &AzureFirewallList{})
}
