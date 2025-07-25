// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20220901 "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type StorageAccountsBlobServicesContainer struct {
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1api20220901-storageaccountsblobservicescontainer,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsblobservicescontainers,verbs=create;update,versions=v1api20220901,name=default.v1api20220901.storageaccountsblobservicescontainers.storage.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &StorageAccountsBlobServicesContainer{}

// Default applies defaults to the StorageAccountsBlobServicesContainer resource
func (container *StorageAccountsBlobServicesContainer) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20220901.StorageAccountsBlobServicesContainer)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/StorageAccountsBlobServicesContainer, but got %T", obj)
	}
	err := container.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = container
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (container *StorageAccountsBlobServicesContainer) defaultAzureName(ctx context.Context, obj *v20220901.StorageAccountsBlobServicesContainer) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the StorageAccountsBlobServicesContainer resource
func (container *StorageAccountsBlobServicesContainer) defaultImpl(ctx context.Context, obj *v20220901.StorageAccountsBlobServicesContainer) error {
	err := container.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1api20220901-storageaccountsblobservicescontainer,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsblobservicescontainers,verbs=create;update,versions=v1api20220901,name=validate.v1api20220901.storageaccountsblobservicescontainers.storage.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &StorageAccountsBlobServicesContainer{}

// ValidateCreate validates the creation of the resource
func (container *StorageAccountsBlobServicesContainer) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220901.StorageAccountsBlobServicesContainer)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/StorageAccountsBlobServicesContainer, but got %T", resource)
	}
	validations := container.createValidations()
	var temp any = container
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220901.StorageAccountsBlobServicesContainer]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (container *StorageAccountsBlobServicesContainer) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220901.StorageAccountsBlobServicesContainer)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/StorageAccountsBlobServicesContainer, but got %T", resource)
	}
	validations := container.deleteValidations()
	var temp any = container
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220901.StorageAccountsBlobServicesContainer]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (container *StorageAccountsBlobServicesContainer) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20220901.StorageAccountsBlobServicesContainer)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/StorageAccountsBlobServicesContainer, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20220901.StorageAccountsBlobServicesContainer)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/StorageAccountsBlobServicesContainer, but got %T", oldResource)
	}
	validations := container.updateValidations()
	var temp any = container
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220901.StorageAccountsBlobServicesContainer]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (container *StorageAccountsBlobServicesContainer) createValidations() []func(ctx context.Context, obj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error){
		container.validateResourceReferences,
		container.validateOwnerReference,
		container.validateSecretDestinations,
		container.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (container *StorageAccountsBlobServicesContainer) deleteValidations() []func(ctx context.Context, obj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (container *StorageAccountsBlobServicesContainer) updateValidations() []func(ctx context.Context, oldObj *v20220901.StorageAccountsBlobServicesContainer, newObj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20220901.StorageAccountsBlobServicesContainer, newObj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20220901.StorageAccountsBlobServicesContainer, newObj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
			return container.validateResourceReferences(ctx, newObj)
		},
		container.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20220901.StorageAccountsBlobServicesContainer, newObj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
			return container.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220901.StorageAccountsBlobServicesContainer, newObj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
			return container.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220901.StorageAccountsBlobServicesContainer, newObj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
			return container.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (container *StorageAccountsBlobServicesContainer) validateConfigMapDestinations(ctx context.Context, obj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (container *StorageAccountsBlobServicesContainer) validateOwnerReference(ctx context.Context, obj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (container *StorageAccountsBlobServicesContainer) validateResourceReferences(ctx context.Context, obj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (container *StorageAccountsBlobServicesContainer) validateSecretDestinations(ctx context.Context, obj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (container *StorageAccountsBlobServicesContainer) validateWriteOnceProperties(ctx context.Context, oldObj *v20220901.StorageAccountsBlobServicesContainer, newObj *v20220901.StorageAccountsBlobServicesContainer) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
