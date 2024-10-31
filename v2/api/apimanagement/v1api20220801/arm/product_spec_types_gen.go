// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Product_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Product entity contract properties.
	Properties *ProductContractProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Product_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (product Product_Spec) GetAPIVersion() string {
	return "2022-08-01"
}

// GetName returns the Name of the resource
func (product *Product_Spec) GetName() string {
	return product.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/products"
func (product *Product_Spec) GetType() string {
	return "Microsoft.ApiManagement/service/products"
}

// Product profile.
type ProductContractProperties struct {
	// ApprovalRequired: whether subscription approval is required. If false, new subscriptions will be approved automatically
	// enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually
	// approve the subscription before the developer can any of the product’s APIs. Can be present only if
	// subscriptionRequired property is present and has a value of false.
	ApprovalRequired *bool `json:"approvalRequired,omitempty"`

	// Description: Product description. May include HTML formatting tags.
	Description *string `json:"description,omitempty"`

	// DisplayName: Product name.
	DisplayName *string `json:"displayName,omitempty"`

	// State: whether product is published or not. Published products are discoverable by users of developer portal. Non
	// published products are visible only to administrators. Default state of Product is notPublished.
	State *ProductContractProperties_State `json:"state,omitempty"`

	// SubscriptionRequired: Whether a product subscription is required for accessing APIs included in this product. If true,
	// the product is referred to as "protected" and a valid subscription key is required for a request to an API included in
	// the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product
	// can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be
	// true.
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty"`

	// SubscriptionsLimit: Whether the number of subscriptions a user can have to this product at the same time. Set to null or
	// omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has
	// a value of false.
	SubscriptionsLimit *int `json:"subscriptionsLimit,omitempty"`

	// Terms: Product terms of use. Developers trying to subscribe to the product will be presented and required to accept
	// these terms before they can complete the subscription process.
	Terms *string `json:"terms,omitempty"`
}

// +kubebuilder:validation:Enum={"notPublished","published"}
type ProductContractProperties_State string

const (
	ProductContractProperties_State_NotPublished = ProductContractProperties_State("notPublished")
	ProductContractProperties_State_Published    = ProductContractProperties_State("published")
)

// Mapping from string to ProductContractProperties_State
var productContractProperties_State_Values = map[string]ProductContractProperties_State{
	"notpublished": ProductContractProperties_State_NotPublished,
	"published":    ProductContractProperties_State_Published,
}