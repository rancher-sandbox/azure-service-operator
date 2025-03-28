// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20231001 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type TrustedAccessRoleBinding struct {
}

// +kubebuilder:webhook:path=/mutate-containerservice-azure-com-v1api20231001-trustedaccessrolebinding,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=containerservice.azure.com,resources=trustedaccessrolebindings,verbs=create;update,versions=v1api20231001,name=default.v1api20231001.trustedaccessrolebindings.containerservice.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &TrustedAccessRoleBinding{}

// Default applies defaults to the TrustedAccessRoleBinding resource
func (binding *TrustedAccessRoleBinding) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20231001.TrustedAccessRoleBinding)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/TrustedAccessRoleBinding, but got %T", obj)
	}
	err := binding.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = binding
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (binding *TrustedAccessRoleBinding) defaultAzureName(ctx context.Context, obj *v20231001.TrustedAccessRoleBinding) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the TrustedAccessRoleBinding resource
func (binding *TrustedAccessRoleBinding) defaultImpl(ctx context.Context, obj *v20231001.TrustedAccessRoleBinding) error {
	err := binding.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-containerservice-azure-com-v1api20231001-trustedaccessrolebinding,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=containerservice.azure.com,resources=trustedaccessrolebindings,verbs=create;update,versions=v1api20231001,name=validate.v1api20231001.trustedaccessrolebindings.containerservice.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &TrustedAccessRoleBinding{}

// ValidateCreate validates the creation of the resource
func (binding *TrustedAccessRoleBinding) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20231001.TrustedAccessRoleBinding)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/TrustedAccessRoleBinding, but got %T", resource)
	}
	validations := binding.createValidations()
	var temp any = binding
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20231001.TrustedAccessRoleBinding]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (binding *TrustedAccessRoleBinding) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20231001.TrustedAccessRoleBinding)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/TrustedAccessRoleBinding, but got %T", resource)
	}
	validations := binding.deleteValidations()
	var temp any = binding
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20231001.TrustedAccessRoleBinding]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (binding *TrustedAccessRoleBinding) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20231001.TrustedAccessRoleBinding)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/TrustedAccessRoleBinding, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20231001.TrustedAccessRoleBinding)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/TrustedAccessRoleBinding, but got %T", oldResource)
	}
	validations := binding.updateValidations()
	var temp any = binding
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20231001.TrustedAccessRoleBinding]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (binding *TrustedAccessRoleBinding) createValidations() []func(ctx context.Context, obj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error){binding.validateResourceReferences, binding.validateOwnerReference, binding.validateSecretDestinations, binding.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (binding *TrustedAccessRoleBinding) deleteValidations() []func(ctx context.Context, obj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (binding *TrustedAccessRoleBinding) updateValidations() []func(ctx context.Context, oldObj *v20231001.TrustedAccessRoleBinding, newObj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20231001.TrustedAccessRoleBinding, newObj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20231001.TrustedAccessRoleBinding, newObj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
			return binding.validateResourceReferences(ctx, newObj)
		},
		binding.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20231001.TrustedAccessRoleBinding, newObj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
			return binding.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20231001.TrustedAccessRoleBinding, newObj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
			return binding.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20231001.TrustedAccessRoleBinding, newObj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
			return binding.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (binding *TrustedAccessRoleBinding) validateConfigMapDestinations(ctx context.Context, obj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (binding *TrustedAccessRoleBinding) validateOwnerReference(ctx context.Context, obj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (binding *TrustedAccessRoleBinding) validateResourceReferences(ctx context.Context, obj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (binding *TrustedAccessRoleBinding) validateSecretDestinations(ctx context.Context, obj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (binding *TrustedAccessRoleBinding) validateWriteOnceProperties(ctx context.Context, oldObj *v20231001.TrustedAccessRoleBinding, newObj *v20231001.TrustedAccessRoleBinding) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
