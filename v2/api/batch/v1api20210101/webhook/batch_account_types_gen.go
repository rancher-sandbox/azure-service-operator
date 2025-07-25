// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20210101 "github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type BatchAccount struct {
}

// +kubebuilder:webhook:path=/mutate-batch-azure-com-v1api20210101-batchaccount,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=batch.azure.com,resources=batchaccounts,verbs=create;update,versions=v1api20210101,name=default.v1api20210101.batchaccounts.batch.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &BatchAccount{}

// Default applies defaults to the BatchAccount resource
func (account *BatchAccount) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20210101.BatchAccount)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101/BatchAccount, but got %T", obj)
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
func (account *BatchAccount) defaultAzureName(ctx context.Context, obj *v20210101.BatchAccount) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the BatchAccount resource
func (account *BatchAccount) defaultImpl(ctx context.Context, obj *v20210101.BatchAccount) error {
	err := account.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-batch-azure-com-v1api20210101-batchaccount,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=batch.azure.com,resources=batchaccounts,verbs=create;update,versions=v1api20210101,name=validate.v1api20210101.batchaccounts.batch.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &BatchAccount{}

// ValidateCreate validates the creation of the resource
func (account *BatchAccount) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210101.BatchAccount)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101/BatchAccount, but got %T", resource)
	}
	validations := account.createValidations()
	var temp any = account
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210101.BatchAccount]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (account *BatchAccount) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210101.BatchAccount)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101/BatchAccount, but got %T", resource)
	}
	validations := account.deleteValidations()
	var temp any = account
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210101.BatchAccount]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (account *BatchAccount) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20210101.BatchAccount)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101/BatchAccount, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20210101.BatchAccount)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/batch/v1api20210101/BatchAccount, but got %T", oldResource)
	}
	validations := account.updateValidations()
	var temp any = account
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210101.BatchAccount]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (account *BatchAccount) createValidations() []func(ctx context.Context, obj *v20210101.BatchAccount) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20210101.BatchAccount) (admission.Warnings, error){
		account.validateResourceReferences,
		account.validateOwnerReference,
		account.validateSecretDestinations,
		account.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (account *BatchAccount) deleteValidations() []func(ctx context.Context, obj *v20210101.BatchAccount) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (account *BatchAccount) updateValidations() []func(ctx context.Context, oldObj *v20210101.BatchAccount, newObj *v20210101.BatchAccount) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20210101.BatchAccount, newObj *v20210101.BatchAccount) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20210101.BatchAccount, newObj *v20210101.BatchAccount) (admission.Warnings, error) {
			return account.validateResourceReferences(ctx, newObj)
		},
		account.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20210101.BatchAccount, newObj *v20210101.BatchAccount) (admission.Warnings, error) {
			return account.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210101.BatchAccount, newObj *v20210101.BatchAccount) (admission.Warnings, error) {
			return account.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210101.BatchAccount, newObj *v20210101.BatchAccount) (admission.Warnings, error) {
			return account.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (account *BatchAccount) validateConfigMapDestinations(ctx context.Context, obj *v20210101.BatchAccount) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (account *BatchAccount) validateOwnerReference(ctx context.Context, obj *v20210101.BatchAccount) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (account *BatchAccount) validateResourceReferences(ctx context.Context, obj *v20210101.BatchAccount) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (account *BatchAccount) validateSecretDestinations(ctx context.Context, obj *v20210101.BatchAccount) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (account *BatchAccount) validateWriteOnceProperties(ctx context.Context, oldObj *v20210101.BatchAccount, newObj *v20210101.BatchAccount) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
