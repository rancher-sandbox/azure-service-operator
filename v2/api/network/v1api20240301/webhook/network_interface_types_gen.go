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

type NetworkInterface struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240301-networkinterface,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=networkinterfaces,verbs=create;update,versions=v1api20240301,name=default.v1api20240301.networkinterfaces.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &NetworkInterface{}

// Default applies defaults to the NetworkInterface resource
func (networkInterface *NetworkInterface) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240301.NetworkInterface)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/NetworkInterface, but got %T", obj)
	}
	err := networkInterface.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = networkInterface
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (networkInterface *NetworkInterface) defaultAzureName(ctx context.Context, obj *v20240301.NetworkInterface) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the NetworkInterface resource
func (networkInterface *NetworkInterface) defaultImpl(ctx context.Context, obj *v20240301.NetworkInterface) error {
	err := networkInterface.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240301-networkinterface,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=networkinterfaces,verbs=create;update,versions=v1api20240301,name=validate.v1api20240301.networkinterfaces.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &NetworkInterface{}

// ValidateCreate validates the creation of the resource
func (networkInterface *NetworkInterface) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.NetworkInterface)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/NetworkInterface, but got %T", resource)
	}
	validations := networkInterface.createValidations()
	var temp any = networkInterface
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.NetworkInterface]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (networkInterface *NetworkInterface) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.NetworkInterface)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/NetworkInterface, but got %T", resource)
	}
	validations := networkInterface.deleteValidations()
	var temp any = networkInterface
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.NetworkInterface]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (networkInterface *NetworkInterface) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240301.NetworkInterface)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/NetworkInterface, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240301.NetworkInterface)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/NetworkInterface, but got %T", oldResource)
	}
	validations := networkInterface.updateValidations()
	var temp any = networkInterface
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.NetworkInterface]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (networkInterface *NetworkInterface) createValidations() []func(ctx context.Context, obj *v20240301.NetworkInterface) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240301.NetworkInterface) (admission.Warnings, error){
		networkInterface.validateResourceReferences,
		networkInterface.validateOwnerReference,
		networkInterface.validateSecretDestinations,
		networkInterface.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (networkInterface *NetworkInterface) deleteValidations() []func(ctx context.Context, obj *v20240301.NetworkInterface) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (networkInterface *NetworkInterface) updateValidations() []func(ctx context.Context, oldObj *v20240301.NetworkInterface, newObj *v20240301.NetworkInterface) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240301.NetworkInterface, newObj *v20240301.NetworkInterface) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240301.NetworkInterface, newObj *v20240301.NetworkInterface) (admission.Warnings, error) {
			return networkInterface.validateResourceReferences(ctx, newObj)
		},
		networkInterface.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240301.NetworkInterface, newObj *v20240301.NetworkInterface) (admission.Warnings, error) {
			return networkInterface.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.NetworkInterface, newObj *v20240301.NetworkInterface) (admission.Warnings, error) {
			return networkInterface.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.NetworkInterface, newObj *v20240301.NetworkInterface) (admission.Warnings, error) {
			return networkInterface.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (networkInterface *NetworkInterface) validateConfigMapDestinations(ctx context.Context, obj *v20240301.NetworkInterface) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (networkInterface *NetworkInterface) validateOwnerReference(ctx context.Context, obj *v20240301.NetworkInterface) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (networkInterface *NetworkInterface) validateResourceReferences(ctx context.Context, obj *v20240301.NetworkInterface) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (networkInterface *NetworkInterface) validateSecretDestinations(ctx context.Context, obj *v20240301.NetworkInterface) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (networkInterface *NetworkInterface) validateWriteOnceProperties(ctx context.Context, oldObj *v20240301.NetworkInterface, newObj *v20240301.NetworkInterface) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
