// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240302

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/compute/v1api20240302/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/compute/v1api20240302/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/diskAccess.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}
type DiskAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskAccess_Spec   `json:"spec,omitempty"`
	Status            DiskAccess_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DiskAccess{}

// GetConditions returns the conditions of the resource
func (access *DiskAccess) GetConditions() conditions.Conditions {
	return access.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (access *DiskAccess) SetConditions(conditions conditions.Conditions) {
	access.Status.Conditions = conditions
}

var _ conversion.Convertible = &DiskAccess{}

// ConvertFrom populates our DiskAccess from the provided hub DiskAccess
func (access *DiskAccess) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.DiskAccess)
	if !ok {
		return fmt.Errorf("expected compute/v1api20240302/storage/DiskAccess but received %T instead", hub)
	}

	return access.AssignProperties_From_DiskAccess(source)
}

// ConvertTo populates the provided hub DiskAccess from our DiskAccess
func (access *DiskAccess) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.DiskAccess)
	if !ok {
		return fmt.Errorf("expected compute/v1api20240302/storage/DiskAccess but received %T instead", hub)
	}

	return access.AssignProperties_To_DiskAccess(destination)
}

var _ configmaps.Exporter = &DiskAccess{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (access *DiskAccess) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if access.Spec.OperatorSpec == nil {
		return nil
	}
	return access.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &DiskAccess{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (access *DiskAccess) SecretDestinationExpressions() []*core.DestinationExpression {
	if access.Spec.OperatorSpec == nil {
		return nil
	}
	return access.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &DiskAccess{}

// InitializeSpec initializes the spec for this resource from the given status
func (access *DiskAccess) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*DiskAccess_STATUS); ok {
		return access.Spec.Initialize_From_DiskAccess_STATUS(s)
	}

	return fmt.Errorf("expected Status of type DiskAccess_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &DiskAccess{}

// AzureName returns the Azure name of the resource
func (access *DiskAccess) AzureName() string {
	return access.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-03-02"
func (access DiskAccess) GetAPIVersion() string {
	return "2024-03-02"
}

// GetResourceScope returns the scope of the resource
func (access *DiskAccess) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (access *DiskAccess) GetSpec() genruntime.ConvertibleSpec {
	return &access.Spec
}

// GetStatus returns the status of this resource
func (access *DiskAccess) GetStatus() genruntime.ConvertibleStatus {
	return &access.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (access *DiskAccess) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/diskAccesses"
func (access *DiskAccess) GetType() string {
	return "Microsoft.Compute/diskAccesses"
}

// NewEmptyStatus returns a new empty (blank) status
func (access *DiskAccess) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DiskAccess_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (access *DiskAccess) Owner() *genruntime.ResourceReference {
	if access.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(access.Spec)
	return access.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (access *DiskAccess) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DiskAccess_STATUS); ok {
		access.Status = *st
		return nil
	}

	// Convert status to required version
	var st DiskAccess_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	access.Status = st
	return nil
}

// AssignProperties_From_DiskAccess populates our DiskAccess from the provided source DiskAccess
func (access *DiskAccess) AssignProperties_From_DiskAccess(source *storage.DiskAccess) error {

	// ObjectMeta
	access.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DiskAccess_Spec
	err := spec.AssignProperties_From_DiskAccess_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_DiskAccess_Spec() to populate field Spec")
	}
	access.Spec = spec

	// Status
	var status DiskAccess_STATUS
	err = status.AssignProperties_From_DiskAccess_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_DiskAccess_STATUS() to populate field Status")
	}
	access.Status = status

	// No error
	return nil
}

// AssignProperties_To_DiskAccess populates the provided destination DiskAccess from our DiskAccess
func (access *DiskAccess) AssignProperties_To_DiskAccess(destination *storage.DiskAccess) error {

	// ObjectMeta
	destination.ObjectMeta = *access.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.DiskAccess_Spec
	err := access.Spec.AssignProperties_To_DiskAccess_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_DiskAccess_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.DiskAccess_STATUS
	err = access.Status.AssignProperties_To_DiskAccess_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_DiskAccess_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (access *DiskAccess) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: access.Spec.OriginalVersion(),
		Kind:    "DiskAccess",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/diskAccess.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskAccesses/{diskAccessName}
type DiskAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskAccess `json:"items"`
}

type DiskAccess_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// ExtendedLocation: The extended location where the disk access will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation `json:"extendedLocation,omitempty"`

	// +kubebuilder:validation:Required
	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *DiskAccessOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DiskAccess_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (access *DiskAccess_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if access == nil {
		return nil, nil
	}
	result := &arm.DiskAccess_Spec{}

	// Set property "ExtendedLocation":
	if access.ExtendedLocation != nil {
		extendedLocation_ARM, err := (*access.ExtendedLocation).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		extendedLocation := *extendedLocation_ARM.(*arm.ExtendedLocation)
		result.ExtendedLocation = &extendedLocation
	}

	// Set property "Location":
	if access.Location != nil {
		location := *access.Location
		result.Location = &location
	}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Tags":
	if access.Tags != nil {
		result.Tags = make(map[string]string, len(access.Tags))
		for key, value := range access.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (access *DiskAccess_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.DiskAccess_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (access *DiskAccess_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.DiskAccess_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.DiskAccess_Spec, got %T", armInput)
	}

	// Set property "AzureName":
	access.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "ExtendedLocation":
	if typedInput.ExtendedLocation != nil {
		var extendedLocation1 ExtendedLocation
		err := extendedLocation1.PopulateFromARM(owner, *typedInput.ExtendedLocation)
		if err != nil {
			return err
		}
		extendedLocation := extendedLocation1
		access.ExtendedLocation = &extendedLocation
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		access.Location = &location
	}

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	access.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		access.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			access.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DiskAccess_Spec{}

// ConvertSpecFrom populates our DiskAccess_Spec from the provided source
func (access *DiskAccess_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.DiskAccess_Spec)
	if ok {
		// Populate our instance from source
		return access.AssignProperties_From_DiskAccess_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.DiskAccess_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = access.AssignProperties_From_DiskAccess_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DiskAccess_Spec
func (access *DiskAccess_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.DiskAccess_Spec)
	if ok {
		// Populate destination from our instance
		return access.AssignProperties_To_DiskAccess_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.DiskAccess_Spec{}
	err := access.AssignProperties_To_DiskAccess_Spec(dst)
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

// AssignProperties_From_DiskAccess_Spec populates our DiskAccess_Spec from the provided source DiskAccess_Spec
func (access *DiskAccess_Spec) AssignProperties_From_DiskAccess_Spec(source *storage.DiskAccess_Spec) error {

	// AzureName
	access.AzureName = source.AzureName

	// ExtendedLocation
	if source.ExtendedLocation != nil {
		var extendedLocation ExtendedLocation
		err := extendedLocation.AssignProperties_From_ExtendedLocation(source.ExtendedLocation)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_ExtendedLocation() to populate field ExtendedLocation")
		}
		access.ExtendedLocation = &extendedLocation
	} else {
		access.ExtendedLocation = nil
	}

	// Location
	access.Location = genruntime.ClonePointerToString(source.Location)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec DiskAccessOperatorSpec
		err := operatorSpec.AssignProperties_From_DiskAccessOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_DiskAccessOperatorSpec() to populate field OperatorSpec")
		}
		access.OperatorSpec = &operatorSpec
	} else {
		access.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		access.Owner = &owner
	} else {
		access.Owner = nil
	}

	// Tags
	access.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_DiskAccess_Spec populates the provided destination DiskAccess_Spec from our DiskAccess_Spec
func (access *DiskAccess_Spec) AssignProperties_To_DiskAccess_Spec(destination *storage.DiskAccess_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = access.AzureName

	// ExtendedLocation
	if access.ExtendedLocation != nil {
		var extendedLocation storage.ExtendedLocation
		err := access.ExtendedLocation.AssignProperties_To_ExtendedLocation(&extendedLocation)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_ExtendedLocation() to populate field ExtendedLocation")
		}
		destination.ExtendedLocation = &extendedLocation
	} else {
		destination.ExtendedLocation = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(access.Location)

	// OperatorSpec
	if access.OperatorSpec != nil {
		var operatorSpec storage.DiskAccessOperatorSpec
		err := access.OperatorSpec.AssignProperties_To_DiskAccessOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_DiskAccessOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = access.OriginalVersion()

	// Owner
	if access.Owner != nil {
		owner := access.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(access.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_DiskAccess_STATUS populates our DiskAccess_Spec from the provided source DiskAccess_STATUS
func (access *DiskAccess_Spec) Initialize_From_DiskAccess_STATUS(source *DiskAccess_STATUS) error {

	// ExtendedLocation
	if source.ExtendedLocation != nil {
		var extendedLocation ExtendedLocation
		err := extendedLocation.Initialize_From_ExtendedLocation_STATUS(source.ExtendedLocation)
		if err != nil {
			return eris.Wrap(err, "calling Initialize_From_ExtendedLocation_STATUS() to populate field ExtendedLocation")
		}
		access.ExtendedLocation = &extendedLocation
	} else {
		access.ExtendedLocation = nil
	}

	// Location
	access.Location = genruntime.ClonePointerToString(source.Location)

	// Tags
	access.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (access *DiskAccess_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (access *DiskAccess_Spec) SetAzureName(azureName string) { access.AzureName = azureName }

// disk access resource.
type DiskAccess_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// ExtendedLocation: The extended location where the disk access will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation_STATUS `json:"extendedLocation,omitempty"`

	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// PrivateEndpointConnections: A readonly collection of private endpoint connections created on the disk. Currently only
	// one endpoint connection is supported.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: The disk access resource provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// TimeCreated: The time when the disk access was created.
	TimeCreated *string `json:"timeCreated,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DiskAccess_STATUS{}

// ConvertStatusFrom populates our DiskAccess_STATUS from the provided source
func (access *DiskAccess_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.DiskAccess_STATUS)
	if ok {
		// Populate our instance from source
		return access.AssignProperties_From_DiskAccess_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.DiskAccess_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = access.AssignProperties_From_DiskAccess_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DiskAccess_STATUS
func (access *DiskAccess_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.DiskAccess_STATUS)
	if ok {
		// Populate destination from our instance
		return access.AssignProperties_To_DiskAccess_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.DiskAccess_STATUS{}
	err := access.AssignProperties_To_DiskAccess_STATUS(dst)
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

var _ genruntime.FromARMConverter = &DiskAccess_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (access *DiskAccess_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.DiskAccess_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (access *DiskAccess_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.DiskAccess_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.DiskAccess_STATUS, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "ExtendedLocation":
	if typedInput.ExtendedLocation != nil {
		var extendedLocation1 ExtendedLocation_STATUS
		err := extendedLocation1.PopulateFromARM(owner, *typedInput.ExtendedLocation)
		if err != nil {
			return err
		}
		extendedLocation := extendedLocation1
		access.ExtendedLocation = &extendedLocation
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		access.Id = &id
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		access.Location = &location
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		access.Name = &name
	}

	// Set property "PrivateEndpointConnections":
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.PrivateEndpointConnections {
			var item1 PrivateEndpointConnection_STATUS
			err := item1.PopulateFromARM(owner, item)
			if err != nil {
				return err
			}
			access.PrivateEndpointConnections = append(access.PrivateEndpointConnections, item1)
		}
	}

	// Set property "ProvisioningState":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			access.ProvisioningState = &provisioningState
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		access.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			access.Tags[key] = value
		}
	}

	// Set property "TimeCreated":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.TimeCreated != nil {
			timeCreated := *typedInput.Properties.TimeCreated
			access.TimeCreated = &timeCreated
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		access.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_DiskAccess_STATUS populates our DiskAccess_STATUS from the provided source DiskAccess_STATUS
func (access *DiskAccess_STATUS) AssignProperties_From_DiskAccess_STATUS(source *storage.DiskAccess_STATUS) error {

	// Conditions
	access.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// ExtendedLocation
	if source.ExtendedLocation != nil {
		var extendedLocation ExtendedLocation_STATUS
		err := extendedLocation.AssignProperties_From_ExtendedLocation_STATUS(source.ExtendedLocation)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_ExtendedLocation_STATUS() to populate field ExtendedLocation")
		}
		access.ExtendedLocation = &extendedLocation
	} else {
		access.ExtendedLocation = nil
	}

	// Id
	access.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	access.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	access.Name = genruntime.ClonePointerToString(source.Name)

	// PrivateEndpointConnections
	if source.PrivateEndpointConnections != nil {
		privateEndpointConnectionList := make([]PrivateEndpointConnection_STATUS, len(source.PrivateEndpointConnections))
		for privateEndpointConnectionIndex, privateEndpointConnectionItem := range source.PrivateEndpointConnections {
			// Shadow the loop variable to avoid aliasing
			privateEndpointConnectionItem := privateEndpointConnectionItem
			var privateEndpointConnection PrivateEndpointConnection_STATUS
			err := privateEndpointConnection.AssignProperties_From_PrivateEndpointConnection_STATUS(&privateEndpointConnectionItem)
			if err != nil {
				return eris.Wrap(err, "calling AssignProperties_From_PrivateEndpointConnection_STATUS() to populate field PrivateEndpointConnections")
			}
			privateEndpointConnectionList[privateEndpointConnectionIndex] = privateEndpointConnection
		}
		access.PrivateEndpointConnections = privateEndpointConnectionList
	} else {
		access.PrivateEndpointConnections = nil
	}

	// ProvisioningState
	access.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// Tags
	access.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// TimeCreated
	access.TimeCreated = genruntime.ClonePointerToString(source.TimeCreated)

	// Type
	access.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_DiskAccess_STATUS populates the provided destination DiskAccess_STATUS from our DiskAccess_STATUS
func (access *DiskAccess_STATUS) AssignProperties_To_DiskAccess_STATUS(destination *storage.DiskAccess_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(access.Conditions)

	// ExtendedLocation
	if access.ExtendedLocation != nil {
		var extendedLocation storage.ExtendedLocation_STATUS
		err := access.ExtendedLocation.AssignProperties_To_ExtendedLocation_STATUS(&extendedLocation)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_ExtendedLocation_STATUS() to populate field ExtendedLocation")
		}
		destination.ExtendedLocation = &extendedLocation
	} else {
		destination.ExtendedLocation = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(access.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(access.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(access.Name)

	// PrivateEndpointConnections
	if access.PrivateEndpointConnections != nil {
		privateEndpointConnectionList := make([]storage.PrivateEndpointConnection_STATUS, len(access.PrivateEndpointConnections))
		for privateEndpointConnectionIndex, privateEndpointConnectionItem := range access.PrivateEndpointConnections {
			// Shadow the loop variable to avoid aliasing
			privateEndpointConnectionItem := privateEndpointConnectionItem
			var privateEndpointConnection storage.PrivateEndpointConnection_STATUS
			err := privateEndpointConnectionItem.AssignProperties_To_PrivateEndpointConnection_STATUS(&privateEndpointConnection)
			if err != nil {
				return eris.Wrap(err, "calling AssignProperties_To_PrivateEndpointConnection_STATUS() to populate field PrivateEndpointConnections")
			}
			privateEndpointConnectionList[privateEndpointConnectionIndex] = privateEndpointConnection
		}
		destination.PrivateEndpointConnections = privateEndpointConnectionList
	} else {
		destination.PrivateEndpointConnections = nil
	}

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(access.ProvisioningState)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(access.Tags)

	// TimeCreated
	destination.TimeCreated = genruntime.ClonePointerToString(access.TimeCreated)

	// Type
	destination.Type = genruntime.ClonePointerToString(access.Type)

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
type DiskAccessOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_DiskAccessOperatorSpec populates our DiskAccessOperatorSpec from the provided source DiskAccessOperatorSpec
func (operator *DiskAccessOperatorSpec) AssignProperties_From_DiskAccessOperatorSpec(source *storage.DiskAccessOperatorSpec) error {

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

// AssignProperties_To_DiskAccessOperatorSpec populates the provided destination DiskAccessOperatorSpec from our DiskAccessOperatorSpec
func (operator *DiskAccessOperatorSpec) AssignProperties_To_DiskAccessOperatorSpec(destination *storage.DiskAccessOperatorSpec) error {
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

// The Private Endpoint Connection resource.
type PrivateEndpointConnection_STATUS struct {
	// Id: private endpoint connection Id
	Id *string `json:"id,omitempty"`
}

var _ genruntime.FromARMConverter = &PrivateEndpointConnection_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (connection *PrivateEndpointConnection_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.PrivateEndpointConnection_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (connection *PrivateEndpointConnection_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.PrivateEndpointConnection_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.PrivateEndpointConnection_STATUS, got %T", armInput)
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		connection.Id = &id
	}

	// No error
	return nil
}

// AssignProperties_From_PrivateEndpointConnection_STATUS populates our PrivateEndpointConnection_STATUS from the provided source PrivateEndpointConnection_STATUS
func (connection *PrivateEndpointConnection_STATUS) AssignProperties_From_PrivateEndpointConnection_STATUS(source *storage.PrivateEndpointConnection_STATUS) error {

	// Id
	connection.Id = genruntime.ClonePointerToString(source.Id)

	// No error
	return nil
}

// AssignProperties_To_PrivateEndpointConnection_STATUS populates the provided destination PrivateEndpointConnection_STATUS from our PrivateEndpointConnection_STATUS
func (connection *PrivateEndpointConnection_STATUS) AssignProperties_To_PrivateEndpointConnection_STATUS(destination *storage.PrivateEndpointConnection_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Id
	destination.Id = genruntime.ClonePointerToString(connection.Id)

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
	SchemeBuilder.Register(&DiskAccess{}, &DiskAccessList{})
}
