// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20240301 "github.com/Azure/azure-service-operator/v2/api/app/v1api20240301"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type ManagedEnvironment struct {
}

// +kubebuilder:webhook:path=/mutate-app-azure-com-v1api20240301-managedenvironment,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=app.azure.com,resources=managedenvironments,verbs=create;update,versions=v1api20240301,name=default.v1api20240301.managedenvironments.app.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &ManagedEnvironment{}

// Default applies defaults to the ManagedEnvironment resource
func (environment *ManagedEnvironment) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240301.ManagedEnvironment)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/app/v1api20240301/ManagedEnvironment, but got %T", obj)
	}
	err := environment.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = environment
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (environment *ManagedEnvironment) defaultAzureName(ctx context.Context, obj *v20240301.ManagedEnvironment) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the ManagedEnvironment resource
func (environment *ManagedEnvironment) defaultImpl(ctx context.Context, obj *v20240301.ManagedEnvironment) error {
	err := environment.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-app-azure-com-v1api20240301-managedenvironment,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=app.azure.com,resources=managedenvironments,verbs=create;update,versions=v1api20240301,name=validate.v1api20240301.managedenvironments.app.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &ManagedEnvironment{}

// ValidateCreate validates the creation of the resource
func (environment *ManagedEnvironment) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.ManagedEnvironment)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/app/v1api20240301/ManagedEnvironment, but got %T", resource)
	}
	validations := environment.createValidations()
	var temp any = environment
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.ManagedEnvironment]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (environment *ManagedEnvironment) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.ManagedEnvironment)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/app/v1api20240301/ManagedEnvironment, but got %T", resource)
	}
	validations := environment.deleteValidations()
	var temp any = environment
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.ManagedEnvironment]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (environment *ManagedEnvironment) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240301.ManagedEnvironment)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/app/v1api20240301/ManagedEnvironment, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240301.ManagedEnvironment)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/app/v1api20240301/ManagedEnvironment, but got %T", oldResource)
	}
	validations := environment.updateValidations()
	var temp any = environment
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.ManagedEnvironment]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (environment *ManagedEnvironment) createValidations() []func(ctx context.Context, obj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240301.ManagedEnvironment) (admission.Warnings, error){environment.validateResourceReferences, environment.validateOwnerReference, environment.validateSecretDestinations, environment.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (environment *ManagedEnvironment) deleteValidations() []func(ctx context.Context, obj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (environment *ManagedEnvironment) updateValidations() []func(ctx context.Context, oldObj *v20240301.ManagedEnvironment, newObj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240301.ManagedEnvironment, newObj *v20240301.ManagedEnvironment) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240301.ManagedEnvironment, newObj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
			return environment.validateResourceReferences(ctx, newObj)
		},
		environment.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240301.ManagedEnvironment, newObj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
			return environment.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.ManagedEnvironment, newObj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
			return environment.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.ManagedEnvironment, newObj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
			return environment.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (environment *ManagedEnvironment) validateConfigMapDestinations(ctx context.Context, obj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (environment *ManagedEnvironment) validateOwnerReference(ctx context.Context, obj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (environment *ManagedEnvironment) validateResourceReferences(ctx context.Context, obj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (environment *ManagedEnvironment) validateSecretDestinations(ctx context.Context, obj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (environment *ManagedEnvironment) validateWriteOnceProperties(ctx context.Context, oldObj *v20240301.ManagedEnvironment, newObj *v20240301.ManagedEnvironment) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
