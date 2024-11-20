// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230801

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230801/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2023-08-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/firewallRules/{ruleName}
type RedisFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisFirewallRule_Spec   `json:"spec,omitempty"`
	Status            RedisFirewallRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisFirewallRule{}

// GetConditions returns the conditions of the resource
func (rule *RedisFirewallRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *RedisFirewallRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisFirewallRule{}

// ConvertFrom populates our RedisFirewallRule from the provided hub RedisFirewallRule
func (rule *RedisFirewallRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.RedisFirewallRule)
	if !ok {
		return fmt.Errorf("expected cache/v1api20230801/storage/RedisFirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_RedisFirewallRule(source)
}

// ConvertTo populates the provided hub RedisFirewallRule from our RedisFirewallRule
func (rule *RedisFirewallRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.RedisFirewallRule)
	if !ok {
		return fmt.Errorf("expected cache/v1api20230801/storage/RedisFirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_RedisFirewallRule(destination)
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1api20230801-redisfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redisfirewallrules,verbs=create;update,versions=v1api20230801,name=default.v1api20230801.redisfirewallrules.cache.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &RedisFirewallRule{}

// Default applies defaults to the RedisFirewallRule resource
func (rule *RedisFirewallRule) Default() {
	rule.defaultImpl()
	var temp any = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *RedisFirewallRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the RedisFirewallRule resource
func (rule *RedisFirewallRule) defaultImpl() { rule.defaultAzureName() }

var _ configmaps.Exporter = &RedisFirewallRule{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (rule *RedisFirewallRule) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &RedisFirewallRule{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (rule *RedisFirewallRule) SecretDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &RedisFirewallRule{}

// InitializeSpec initializes the spec for this resource from the given status
func (rule *RedisFirewallRule) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*RedisFirewallRule_STATUS); ok {
		return rule.Spec.Initialize_From_RedisFirewallRule_STATUS(s)
	}

	return fmt.Errorf("expected Status of type RedisFirewallRule_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &RedisFirewallRule{}

// AzureName returns the Azure name of the resource
func (rule *RedisFirewallRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-08-01"
func (rule RedisFirewallRule) GetAPIVersion() string {
	return "2023-08-01"
}

// GetResourceScope returns the scope of the resource
func (rule *RedisFirewallRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *RedisFirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *RedisFirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (rule *RedisFirewallRule) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/firewallRules"
func (rule *RedisFirewallRule) GetType() string {
	return "Microsoft.Cache/redis/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *RedisFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisFirewallRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *RedisFirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return rule.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (rule *RedisFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisFirewallRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisFirewallRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1api20230801-redisfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redisfirewallrules,verbs=create;update,versions=v1api20230801,name=validate.v1api20230801.redisfirewallrules.cache.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &RedisFirewallRule{}

// ValidateCreate validates the creation of the resource
func (rule *RedisFirewallRule) ValidateCreate() (admission.Warnings, error) {
	validations := rule.createValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (rule *RedisFirewallRule) ValidateDelete() (admission.Warnings, error) {
	validations := rule.deleteValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (rule *RedisFirewallRule) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := rule.updateValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (rule *RedisFirewallRule) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){rule.validateResourceReferences, rule.validateOwnerReference, rule.validateSecretDestinations, rule.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (rule *RedisFirewallRule) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (rule *RedisFirewallRule) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateSecretDestinations()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateConfigMapDestinations()
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (rule *RedisFirewallRule) validateConfigMapDestinations() (admission.Warnings, error) {
	if rule.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(rule, nil, rule.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (rule *RedisFirewallRule) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(rule)
}

// validateResourceReferences validates all resource references
func (rule *RedisFirewallRule) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (rule *RedisFirewallRule) validateSecretDestinations() (admission.Warnings, error) {
	if rule.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(rule, nil, rule.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *RedisFirewallRule) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*RedisFirewallRule)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignProperties_From_RedisFirewallRule populates our RedisFirewallRule from the provided source RedisFirewallRule
func (rule *RedisFirewallRule) AssignProperties_From_RedisFirewallRule(source *storage.RedisFirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisFirewallRule_Spec
	err := spec.AssignProperties_From_RedisFirewallRule_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_RedisFirewallRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status RedisFirewallRule_STATUS
	err = status.AssignProperties_From_RedisFirewallRule_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_RedisFirewallRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignProperties_To_RedisFirewallRule populates the provided destination RedisFirewallRule from our RedisFirewallRule
func (rule *RedisFirewallRule) AssignProperties_To_RedisFirewallRule(destination *storage.RedisFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.RedisFirewallRule_Spec
	err := rule.Spec.AssignProperties_To_RedisFirewallRule_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_RedisFirewallRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.RedisFirewallRule_STATUS
	err = rule.Status.AssignProperties_To_RedisFirewallRule_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_RedisFirewallRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *RedisFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "RedisFirewallRule",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2023-08-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/firewallRules/{ruleName}
type RedisFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisFirewallRule `json:"items"`
}

type RedisFirewallRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// EndIP: highest IP address included in the range
	EndIP *string `json:"endIP,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *RedisFirewallRuleOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`

	// +kubebuilder:validation:Required
	// StartIP: lowest IP address included in the range
	StartIP *string `json:"startIP,omitempty"`
}

var _ genruntime.ARMTransformer = &RedisFirewallRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *RedisFirewallRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &arm.RedisFirewallRule_Spec{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if rule.EndIP != nil || rule.StartIP != nil {
		result.Properties = &arm.RedisFirewallRuleProperties{}
	}
	if rule.EndIP != nil {
		endIP := *rule.EndIP
		result.Properties.EndIP = &endIP
	}
	if rule.StartIP != nil {
		startIP := *rule.StartIP
		result.Properties.StartIP = &startIP
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *RedisFirewallRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.RedisFirewallRule_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *RedisFirewallRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.RedisFirewallRule_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.RedisFirewallRule_Spec, got %T", armInput)
	}

	// Set property "AzureName":
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "EndIP":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIP != nil {
			endIP := *typedInput.Properties.EndIP
			rule.EndIP = &endIP
		}
	}

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	rule.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "StartIP":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIP != nil {
			startIP := *typedInput.Properties.StartIP
			rule.StartIP = &startIP
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RedisFirewallRule_Spec{}

// ConvertSpecFrom populates our RedisFirewallRule_Spec from the provided source
func (rule *RedisFirewallRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.RedisFirewallRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_RedisFirewallRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.RedisFirewallRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_RedisFirewallRule_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisFirewallRule_Spec
func (rule *RedisFirewallRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.RedisFirewallRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_RedisFirewallRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.RedisFirewallRule_Spec{}
	err := rule.AssignProperties_To_RedisFirewallRule_Spec(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_RedisFirewallRule_Spec populates our RedisFirewallRule_Spec from the provided source RedisFirewallRule_Spec
func (rule *RedisFirewallRule_Spec) AssignProperties_From_RedisFirewallRule_Spec(source *storage.RedisFirewallRule_Spec) error {

	// AzureName
	rule.AzureName = source.AzureName

	// EndIP
	rule.EndIP = genruntime.ClonePointerToString(source.EndIP)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec RedisFirewallRuleOperatorSpec
		err := operatorSpec.AssignProperties_From_RedisFirewallRuleOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_RedisFirewallRuleOperatorSpec() to populate field OperatorSpec")
		}
		rule.OperatorSpec = &operatorSpec
	} else {
		rule.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// StartIP
	rule.StartIP = genruntime.ClonePointerToString(source.StartIP)

	// No error
	return nil
}

// AssignProperties_To_RedisFirewallRule_Spec populates the provided destination RedisFirewallRule_Spec from our RedisFirewallRule_Spec
func (rule *RedisFirewallRule_Spec) AssignProperties_To_RedisFirewallRule_Spec(destination *storage.RedisFirewallRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

	// EndIP
	destination.EndIP = genruntime.ClonePointerToString(rule.EndIP)

	// OperatorSpec
	if rule.OperatorSpec != nil {
		var operatorSpec storage.RedisFirewallRuleOperatorSpec
		err := rule.OperatorSpec.AssignProperties_To_RedisFirewallRuleOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_RedisFirewallRuleOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = rule.OriginalVersion()

	// Owner
	if rule.Owner != nil {
		owner := rule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// StartIP
	destination.StartIP = genruntime.ClonePointerToString(rule.StartIP)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_RedisFirewallRule_STATUS populates our RedisFirewallRule_Spec from the provided source RedisFirewallRule_STATUS
func (rule *RedisFirewallRule_Spec) Initialize_From_RedisFirewallRule_STATUS(source *RedisFirewallRule_STATUS) error {

	// EndIP
	rule.EndIP = genruntime.ClonePointerToString(source.EndIP)

	// StartIP
	rule.StartIP = genruntime.ClonePointerToString(source.StartIP)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (rule *RedisFirewallRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *RedisFirewallRule_Spec) SetAzureName(azureName string) { rule.AzureName = azureName }

type RedisFirewallRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// EndIP: highest IP address included in the range
	EndIP *string `json:"endIP,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// StartIP: lowest IP address included in the range
	StartIP *string `json:"startIP,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisFirewallRule_STATUS{}

// ConvertStatusFrom populates our RedisFirewallRule_STATUS from the provided source
func (rule *RedisFirewallRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.RedisFirewallRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_RedisFirewallRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.RedisFirewallRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_RedisFirewallRule_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisFirewallRule_STATUS
func (rule *RedisFirewallRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.RedisFirewallRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_RedisFirewallRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.RedisFirewallRule_STATUS{}
	err := rule.AssignProperties_To_RedisFirewallRule_STATUS(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &RedisFirewallRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *RedisFirewallRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.RedisFirewallRule_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *RedisFirewallRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.RedisFirewallRule_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.RedisFirewallRule_STATUS, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "EndIP":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIP != nil {
			endIP := *typedInput.Properties.EndIP
			rule.EndIP = &endIP
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		rule.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		rule.Name = &name
	}

	// Set property "StartIP":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIP != nil {
			startIP := *typedInput.Properties.StartIP
			rule.StartIP = &startIP
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		rule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_RedisFirewallRule_STATUS populates our RedisFirewallRule_STATUS from the provided source RedisFirewallRule_STATUS
func (rule *RedisFirewallRule_STATUS) AssignProperties_From_RedisFirewallRule_STATUS(source *storage.RedisFirewallRule_STATUS) error {

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// EndIP
	rule.EndIP = genruntime.ClonePointerToString(source.EndIP)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// StartIP
	rule.StartIP = genruntime.ClonePointerToString(source.StartIP)

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_RedisFirewallRule_STATUS populates the provided destination RedisFirewallRule_STATUS from our RedisFirewallRule_STATUS
func (rule *RedisFirewallRule_STATUS) AssignProperties_To_RedisFirewallRule_STATUS(destination *storage.RedisFirewallRule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// EndIP
	destination.EndIP = genruntime.ClonePointerToString(rule.EndIP)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// StartIP
	destination.StartIP = genruntime.ClonePointerToString(rule.StartIP)

	// Type
	destination.Type = genruntime.ClonePointerToString(rule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type RedisFirewallRuleOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_RedisFirewallRuleOperatorSpec populates our RedisFirewallRuleOperatorSpec from the provided source RedisFirewallRuleOperatorSpec
func (operator *RedisFirewallRuleOperatorSpec) AssignProperties_From_RedisFirewallRuleOperatorSpec(source *storage.RedisFirewallRuleOperatorSpec) error {

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// No error
	return nil
}

// AssignProperties_To_RedisFirewallRuleOperatorSpec populates the provided destination RedisFirewallRuleOperatorSpec from our RedisFirewallRuleOperatorSpec
func (operator *RedisFirewallRuleOperatorSpec) AssignProperties_To_RedisFirewallRuleOperatorSpec(destination *storage.RedisFirewallRuleOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&RedisFirewallRule{}, &RedisFirewallRuleList{})
}
