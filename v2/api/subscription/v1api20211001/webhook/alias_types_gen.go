// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20211001 "github.com/Azure/azure-service-operator/v2/api/subscription/v1api20211001"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type Alias struct {
}

// +kubebuilder:webhook:path=/mutate-subscription-azure-com-v1api20211001-alias,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=subscription.azure.com,resources=aliases,verbs=create;update,versions=v1api20211001,name=default.v1api20211001.aliases.subscription.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &Alias{}

// Default applies defaults to the Alias resource
func (alias *Alias) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20211001.Alias)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/subscription/v1api20211001/Alias, but got %T", obj)
	}
	err := alias.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = alias
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (alias *Alias) defaultAzureName(ctx context.Context, obj *v20211001.Alias) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the Alias resource
func (alias *Alias) defaultImpl(ctx context.Context, obj *v20211001.Alias) error {
	err := alias.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-subscription-azure-com-v1api20211001-alias,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=subscription.azure.com,resources=aliases,verbs=create;update,versions=v1api20211001,name=validate.v1api20211001.aliases.subscription.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &Alias{}

// ValidateCreate validates the creation of the resource
func (alias *Alias) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211001.Alias)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/subscription/v1api20211001/Alias, but got %T", resource)
	}
	validations := alias.createValidations()
	var temp any = alias
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211001.Alias]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (alias *Alias) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211001.Alias)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/subscription/v1api20211001/Alias, but got %T", resource)
	}
	validations := alias.deleteValidations()
	var temp any = alias
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211001.Alias]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (alias *Alias) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20211001.Alias)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/subscription/v1api20211001/Alias, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20211001.Alias)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/subscription/v1api20211001/Alias, but got %T", oldResource)
	}
	validations := alias.updateValidations()
	var temp any = alias
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211001.Alias]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (alias *Alias) createValidations() []func(ctx context.Context, obj *v20211001.Alias) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20211001.Alias) (admission.Warnings, error){alias.validateResourceReferences, alias.validateSecretDestinations, alias.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (alias *Alias) deleteValidations() []func(ctx context.Context, obj *v20211001.Alias) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (alias *Alias) updateValidations() []func(ctx context.Context, oldObj *v20211001.Alias, newObj *v20211001.Alias) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20211001.Alias, newObj *v20211001.Alias) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20211001.Alias, newObj *v20211001.Alias) (admission.Warnings, error) {
			return alias.validateResourceReferences(ctx, newObj)
		},
		alias.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20211001.Alias, newObj *v20211001.Alias) (admission.Warnings, error) {
			return alias.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211001.Alias, newObj *v20211001.Alias) (admission.Warnings, error) {
			return alias.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (alias *Alias) validateConfigMapDestinations(ctx context.Context, obj *v20211001.Alias) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateResourceReferences validates all resource references
func (alias *Alias) validateResourceReferences(ctx context.Context, obj *v20211001.Alias) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (alias *Alias) validateSecretDestinations(ctx context.Context, obj *v20211001.Alias) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (alias *Alias) validateWriteOnceProperties(ctx context.Context, oldObj *v20211001.Alias, newObj *v20211001.Alias) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
