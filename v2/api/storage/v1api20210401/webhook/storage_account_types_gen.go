// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20210401 "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type StorageAccount struct {
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1api20210401-storageaccount,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccounts,verbs=create;update,versions=v1api20210401,name=default.v1api20210401.storageaccounts.storage.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &StorageAccount{}

// Default applies defaults to the StorageAccount resource
func (account *StorageAccount) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20210401.StorageAccount)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401/StorageAccount, but got %T", obj)
	}
	err := account.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = account
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (account *StorageAccount) defaultAzureName(ctx context.Context, obj *v20210401.StorageAccount) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the StorageAccount resource
func (account *StorageAccount) defaultImpl(ctx context.Context, obj *v20210401.StorageAccount) error {
	err := account.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1api20210401-storageaccount,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccounts,verbs=create;update,versions=v1api20210401,name=validate.v1api20210401.storageaccounts.storage.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &StorageAccount{}

// ValidateCreate validates the creation of the resource
func (account *StorageAccount) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210401.StorageAccount)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401/StorageAccount, but got %T", resource)
	}
	validations := account.createValidations()
	var temp any = account
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210401.StorageAccount]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (account *StorageAccount) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210401.StorageAccount)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401/StorageAccount, but got %T", resource)
	}
	validations := account.deleteValidations()
	var temp any = account
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210401.StorageAccount]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (account *StorageAccount) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20210401.StorageAccount)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401/StorageAccount, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20210401.StorageAccount)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401/StorageAccount, but got %T", oldResource)
	}
	validations := account.updateValidations()
	var temp any = account
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210401.StorageAccount]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (account *StorageAccount) createValidations() []func(ctx context.Context, obj *v20210401.StorageAccount) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20210401.StorageAccount) (admission.Warnings, error){
		account.validateResourceReferences,
		account.validateOwnerReference,
		account.validateSecretDestinations,
		account.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (account *StorageAccount) deleteValidations() []func(ctx context.Context, obj *v20210401.StorageAccount) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (account *StorageAccount) updateValidations() []func(ctx context.Context, oldObj *v20210401.StorageAccount, newObj *v20210401.StorageAccount) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20210401.StorageAccount, newObj *v20210401.StorageAccount) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20210401.StorageAccount, newObj *v20210401.StorageAccount) (admission.Warnings, error) {
			return account.validateResourceReferences(ctx, newObj)
		},
		account.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20210401.StorageAccount, newObj *v20210401.StorageAccount) (admission.Warnings, error) {
			return account.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210401.StorageAccount, newObj *v20210401.StorageAccount) (admission.Warnings, error) {
			return account.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210401.StorageAccount, newObj *v20210401.StorageAccount) (admission.Warnings, error) {
			return account.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (account *StorageAccount) validateConfigMapDestinations(ctx context.Context, obj *v20210401.StorageAccount) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.ConfigMapDestination
	if obj.Spec.OperatorSpec.ConfigMaps != nil {
		toValidate = []*genruntime.ConfigMapDestination{
			obj.Spec.OperatorSpec.ConfigMaps.BlobEndpoint,
			obj.Spec.OperatorSpec.ConfigMaps.DfsEndpoint,
			obj.Spec.OperatorSpec.ConfigMaps.FileEndpoint,
			obj.Spec.OperatorSpec.ConfigMaps.QueueEndpoint,
			obj.Spec.OperatorSpec.ConfigMaps.TableEndpoint,
			obj.Spec.OperatorSpec.ConfigMaps.WebEndpoint,
		}
	}
	return configmaps.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (account *StorageAccount) validateOwnerReference(ctx context.Context, obj *v20210401.StorageAccount) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (account *StorageAccount) validateResourceReferences(ctx context.Context, obj *v20210401.StorageAccount) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (account *StorageAccount) validateSecretDestinations(ctx context.Context, obj *v20210401.StorageAccount) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.SecretDestination
	if obj.Spec.OperatorSpec.Secrets != nil {
		toValidate = []*genruntime.SecretDestination{
			obj.Spec.OperatorSpec.Secrets.BlobEndpoint,
			obj.Spec.OperatorSpec.Secrets.DfsEndpoint,
			obj.Spec.OperatorSpec.Secrets.FileEndpoint,
			obj.Spec.OperatorSpec.Secrets.Key1,
			obj.Spec.OperatorSpec.Secrets.Key2,
			obj.Spec.OperatorSpec.Secrets.QueueEndpoint,
			obj.Spec.OperatorSpec.Secrets.TableEndpoint,
			obj.Spec.OperatorSpec.Secrets.WebEndpoint,
		}
	}
	return secrets.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (account *StorageAccount) validateWriteOnceProperties(ctx context.Context, oldObj *v20210401.StorageAccount, newObj *v20210401.StorageAccount) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
