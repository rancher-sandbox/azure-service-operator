// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20250601 "github.com/Azure/azure-service-operator/v2/api/cognitiveservices/v1api20250601"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type Account struct {
}

// +kubebuilder:webhook:path=/mutate-cognitiveservices-azure-com-v1api20250601-account,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cognitiveservices.azure.com,resources=accounts,verbs=create;update,versions=v1api20250601,name=default.v1api20250601.accounts.cognitiveservices.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &Account{}

// Default applies defaults to the Account resource
func (account *Account) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20250601.Account)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cognitiveservices/v1api20250601/Account, but got %T", obj)
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
func (account *Account) defaultAzureName(ctx context.Context, obj *v20250601.Account) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the Account resource
func (account *Account) defaultImpl(ctx context.Context, obj *v20250601.Account) error {
	err := account.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-cognitiveservices-azure-com-v1api20250601-account,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cognitiveservices.azure.com,resources=accounts,verbs=create;update,versions=v1api20250601,name=validate.v1api20250601.accounts.cognitiveservices.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &Account{}

// ValidateCreate validates the creation of the resource
func (account *Account) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20250601.Account)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cognitiveservices/v1api20250601/Account, but got %T", resource)
	}
	validations := account.createValidations()
	var temp any = account
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20250601.Account]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (account *Account) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20250601.Account)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cognitiveservices/v1api20250601/Account, but got %T", resource)
	}
	validations := account.deleteValidations()
	var temp any = account
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20250601.Account]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (account *Account) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20250601.Account)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cognitiveservices/v1api20250601/Account, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20250601.Account)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cognitiveservices/v1api20250601/Account, but got %T", oldResource)
	}
	validations := account.updateValidations()
	var temp any = account
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20250601.Account]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (account *Account) createValidations() []func(ctx context.Context, obj *v20250601.Account) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20250601.Account) (admission.Warnings, error){
		account.validateResourceReferences,
		account.validateOwnerReference,
		account.validateSecretDestinations,
		account.validateConfigMapDestinations,
		account.validateOptionalConfigMapReferences,
	}
}

// deleteValidations validates the deletion of the resource
func (account *Account) deleteValidations() []func(ctx context.Context, obj *v20250601.Account) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (account *Account) updateValidations() []func(ctx context.Context, oldObj *v20250601.Account, newObj *v20250601.Account) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20250601.Account, newObj *v20250601.Account) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20250601.Account, newObj *v20250601.Account) (admission.Warnings, error) {
			return account.validateResourceReferences(ctx, newObj)
		},
		account.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20250601.Account, newObj *v20250601.Account) (admission.Warnings, error) {
			return account.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20250601.Account, newObj *v20250601.Account) (admission.Warnings, error) {
			return account.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20250601.Account, newObj *v20250601.Account) (admission.Warnings, error) {
			return account.validateConfigMapDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20250601.Account, newObj *v20250601.Account) (admission.Warnings, error) {
			return account.validateOptionalConfigMapReferences(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (account *Account) validateConfigMapDestinations(ctx context.Context, obj *v20250601.Account) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOptionalConfigMapReferences validates all optional configmap reference pairs to ensure that at most 1 is set
func (account *Account) validateOptionalConfigMapReferences(ctx context.Context, obj *v20250601.Account) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return configmaps.ValidateOptionalReferences(refs)
}

// validateOwnerReference validates the owner field
func (account *Account) validateOwnerReference(ctx context.Context, obj *v20250601.Account) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (account *Account) validateResourceReferences(ctx context.Context, obj *v20250601.Account) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (account *Account) validateSecretDestinations(ctx context.Context, obj *v20250601.Account) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.SecretDestination
	if obj.Spec.OperatorSpec.Secrets != nil {
		toValidate = []*genruntime.SecretDestination{
			obj.Spec.OperatorSpec.Secrets.Key1,
			obj.Spec.OperatorSpec.Secrets.Key2,
		}
	}
	return secrets.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (account *Account) validateWriteOnceProperties(ctx context.Context, oldObj *v20250601.Account, newObj *v20250601.Account) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
