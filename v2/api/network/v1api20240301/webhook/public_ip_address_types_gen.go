// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20240301 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type PublicIPAddress struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240301-publicipaddress,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=publicipaddresses,verbs=create;update,versions=v1api20240301,name=default.v1api20240301.publicipaddresses.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &PublicIPAddress{}

// Default applies defaults to the PublicIPAddress resource
func (address *PublicIPAddress) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240301.PublicIPAddress)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/PublicIPAddress, but got %T", obj)
	}
	err := address.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = address
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (address *PublicIPAddress) defaultAzureName(ctx context.Context, obj *v20240301.PublicIPAddress) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the PublicIPAddress resource
func (address *PublicIPAddress) defaultImpl(ctx context.Context, obj *v20240301.PublicIPAddress) error {
	err := address.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240301-publicipaddress,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=publicipaddresses,verbs=create;update,versions=v1api20240301,name=validate.v1api20240301.publicipaddresses.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &PublicIPAddress{}

// ValidateCreate validates the creation of the resource
func (address *PublicIPAddress) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.PublicIPAddress)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/PublicIPAddress, but got %T", resource)
	}
	validations := address.createValidations()
	var temp any = address
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.PublicIPAddress]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (address *PublicIPAddress) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.PublicIPAddress)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/PublicIPAddress, but got %T", resource)
	}
	validations := address.deleteValidations()
	var temp any = address
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.PublicIPAddress]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (address *PublicIPAddress) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240301.PublicIPAddress)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/PublicIPAddress, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240301.PublicIPAddress)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/PublicIPAddress, but got %T", oldResource)
	}
	validations := address.updateValidations()
	var temp any = address
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.PublicIPAddress]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (address *PublicIPAddress) createValidations() []func(ctx context.Context, obj *v20240301.PublicIPAddress) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240301.PublicIPAddress) (admission.Warnings, error){address.validateResourceReferences, address.validateOwnerReference, address.validateSecretDestinations, address.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (address *PublicIPAddress) deleteValidations() []func(ctx context.Context, obj *v20240301.PublicIPAddress) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (address *PublicIPAddress) updateValidations() []func(ctx context.Context, oldObj *v20240301.PublicIPAddress, newObj *v20240301.PublicIPAddress) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240301.PublicIPAddress, newObj *v20240301.PublicIPAddress) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240301.PublicIPAddress, newObj *v20240301.PublicIPAddress) (admission.Warnings, error) {
			return address.validateResourceReferences(ctx, newObj)
		},
		address.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240301.PublicIPAddress, newObj *v20240301.PublicIPAddress) (admission.Warnings, error) {
			return address.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.PublicIPAddress, newObj *v20240301.PublicIPAddress) (admission.Warnings, error) {
			return address.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.PublicIPAddress, newObj *v20240301.PublicIPAddress) (admission.Warnings, error) {
			return address.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (address *PublicIPAddress) validateConfigMapDestinations(ctx context.Context, obj *v20240301.PublicIPAddress) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (address *PublicIPAddress) validateOwnerReference(ctx context.Context, obj *v20240301.PublicIPAddress) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (address *PublicIPAddress) validateResourceReferences(ctx context.Context, obj *v20240301.PublicIPAddress) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (address *PublicIPAddress) validateSecretDestinations(ctx context.Context, obj *v20240301.PublicIPAddress) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (address *PublicIPAddress) validateWriteOnceProperties(ctx context.Context, oldObj *v20240301.PublicIPAddress, newObj *v20240301.PublicIPAddress) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
