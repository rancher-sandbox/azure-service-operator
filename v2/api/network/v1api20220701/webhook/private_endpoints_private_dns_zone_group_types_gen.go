// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20220701 "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type PrivateEndpointsPrivateDnsZoneGroup struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20220701-privateendpointsprivatednszonegroup,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=privateendpointsprivatednszonegroups,verbs=create;update,versions=v1api20220701,name=default.v1api20220701.privateendpointsprivatednszonegroups.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &PrivateEndpointsPrivateDnsZoneGroup{}

// Default applies defaults to the PrivateEndpointsPrivateDnsZoneGroup resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20220701.PrivateEndpointsPrivateDnsZoneGroup)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/PrivateEndpointsPrivateDnsZoneGroup, but got %T", obj)
	}
	err := group.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = group
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (group *PrivateEndpointsPrivateDnsZoneGroup) defaultAzureName(ctx context.Context, obj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the PrivateEndpointsPrivateDnsZoneGroup resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) defaultImpl(ctx context.Context, obj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) error {
	err := group.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20220701-privateendpointsprivatednszonegroup,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=privateendpointsprivatednszonegroups,verbs=create;update,versions=v1api20220701,name=validate.v1api20220701.privateendpointsprivatednszonegroups.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &PrivateEndpointsPrivateDnsZoneGroup{}

// ValidateCreate validates the creation of the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220701.PrivateEndpointsPrivateDnsZoneGroup)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/PrivateEndpointsPrivateDnsZoneGroup, but got %T", resource)
	}
	validations := group.createValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220701.PrivateEndpointsPrivateDnsZoneGroup]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220701.PrivateEndpointsPrivateDnsZoneGroup)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/PrivateEndpointsPrivateDnsZoneGroup, but got %T", resource)
	}
	validations := group.deleteValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220701.PrivateEndpointsPrivateDnsZoneGroup]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20220701.PrivateEndpointsPrivateDnsZoneGroup)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/PrivateEndpointsPrivateDnsZoneGroup, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20220701.PrivateEndpointsPrivateDnsZoneGroup)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/PrivateEndpointsPrivateDnsZoneGroup, but got %T", oldResource)
	}
	validations := group.updateValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220701.PrivateEndpointsPrivateDnsZoneGroup]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) createValidations() []func(ctx context.Context, obj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error){group.validateResourceReferences, group.validateOwnerReference, group.validateSecretDestinations, group.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) deleteValidations() []func(ctx context.Context, obj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) updateValidations() []func(ctx context.Context, oldObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup, newObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup, newObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup, newObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
			return group.validateResourceReferences(ctx, newObj)
		},
		group.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup, newObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
			return group.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup, newObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
			return group.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup, newObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
			return group.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (group *PrivateEndpointsPrivateDnsZoneGroup) validateConfigMapDestinations(ctx context.Context, obj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (group *PrivateEndpointsPrivateDnsZoneGroup) validateOwnerReference(ctx context.Context, obj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (group *PrivateEndpointsPrivateDnsZoneGroup) validateResourceReferences(ctx context.Context, obj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (group *PrivateEndpointsPrivateDnsZoneGroup) validateSecretDestinations(ctx context.Context, obj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (group *PrivateEndpointsPrivateDnsZoneGroup) validateWriteOnceProperties(ctx context.Context, oldObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup, newObj *v20220701.PrivateEndpointsPrivateDnsZoneGroup) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
