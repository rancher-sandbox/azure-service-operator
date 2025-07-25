// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20230101 "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type StorageAccountsFileService struct {
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1api20230101-storageaccountsfileservice,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsfileservices,verbs=create;update,versions=v1api20230101,name=default.v1api20230101.storageaccountsfileservices.storage.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &StorageAccountsFileService{}

// Default applies defaults to the StorageAccountsFileService resource
func (service *StorageAccountsFileService) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20230101.StorageAccountsFileService)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/StorageAccountsFileService, but got %T", obj)
	}
	err := service.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = service
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultImpl applies the code generated defaults to the StorageAccountsFileService resource
func (service *StorageAccountsFileService) defaultImpl(ctx context.Context, obj *v20230101.StorageAccountsFileService) error {
	return nil
}

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1api20230101-storageaccountsfileservice,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsfileservices,verbs=create;update,versions=v1api20230101,name=validate.v1api20230101.storageaccountsfileservices.storage.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &StorageAccountsFileService{}

// ValidateCreate validates the creation of the resource
func (service *StorageAccountsFileService) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230101.StorageAccountsFileService)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/StorageAccountsFileService, but got %T", resource)
	}
	validations := service.createValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230101.StorageAccountsFileService]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (service *StorageAccountsFileService) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230101.StorageAccountsFileService)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/StorageAccountsFileService, but got %T", resource)
	}
	validations := service.deleteValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230101.StorageAccountsFileService]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (service *StorageAccountsFileService) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20230101.StorageAccountsFileService)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/StorageAccountsFileService, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20230101.StorageAccountsFileService)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/StorageAccountsFileService, but got %T", oldResource)
	}
	validations := service.updateValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230101.StorageAccountsFileService]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (service *StorageAccountsFileService) createValidations() []func(ctx context.Context, obj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20230101.StorageAccountsFileService) (admission.Warnings, error){
		service.validateResourceReferences,
		service.validateOwnerReference,
		service.validateSecretDestinations,
		service.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (service *StorageAccountsFileService) deleteValidations() []func(ctx context.Context, obj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (service *StorageAccountsFileService) updateValidations() []func(ctx context.Context, oldObj *v20230101.StorageAccountsFileService, newObj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20230101.StorageAccountsFileService, newObj *v20230101.StorageAccountsFileService) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20230101.StorageAccountsFileService, newObj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
			return service.validateResourceReferences(ctx, newObj)
		},
		service.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20230101.StorageAccountsFileService, newObj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
			return service.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230101.StorageAccountsFileService, newObj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
			return service.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230101.StorageAccountsFileService, newObj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
			return service.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (service *StorageAccountsFileService) validateConfigMapDestinations(ctx context.Context, obj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (service *StorageAccountsFileService) validateOwnerReference(ctx context.Context, obj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (service *StorageAccountsFileService) validateResourceReferences(ctx context.Context, obj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (service *StorageAccountsFileService) validateSecretDestinations(ctx context.Context, obj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (service *StorageAccountsFileService) validateWriteOnceProperties(ctx context.Context, oldObj *v20230101.StorageAccountsFileService, newObj *v20230101.StorageAccountsFileService) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
