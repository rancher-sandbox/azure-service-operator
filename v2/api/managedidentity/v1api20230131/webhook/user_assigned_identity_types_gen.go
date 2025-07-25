// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20230131 "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type UserAssignedIdentity struct {
}

// +kubebuilder:webhook:path=/mutate-managedidentity-azure-com-v1api20230131-userassignedidentity,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=managedidentity.azure.com,resources=userassignedidentities,verbs=create;update,versions=v1api20230131,name=default.v1api20230131.userassignedidentities.managedidentity.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &UserAssignedIdentity{}

// Default applies defaults to the UserAssignedIdentity resource
func (identity *UserAssignedIdentity) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20230131.UserAssignedIdentity)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131/UserAssignedIdentity, but got %T", obj)
	}
	err := identity.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = identity
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (identity *UserAssignedIdentity) defaultAzureName(ctx context.Context, obj *v20230131.UserAssignedIdentity) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the UserAssignedIdentity resource
func (identity *UserAssignedIdentity) defaultImpl(ctx context.Context, obj *v20230131.UserAssignedIdentity) error {
	err := identity.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-managedidentity-azure-com-v1api20230131-userassignedidentity,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=managedidentity.azure.com,resources=userassignedidentities,verbs=create;update,versions=v1api20230131,name=validate.v1api20230131.userassignedidentities.managedidentity.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &UserAssignedIdentity{}

// ValidateCreate validates the creation of the resource
func (identity *UserAssignedIdentity) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230131.UserAssignedIdentity)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131/UserAssignedIdentity, but got %T", resource)
	}
	validations := identity.createValidations()
	var temp any = identity
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230131.UserAssignedIdentity]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (identity *UserAssignedIdentity) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230131.UserAssignedIdentity)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131/UserAssignedIdentity, but got %T", resource)
	}
	validations := identity.deleteValidations()
	var temp any = identity
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230131.UserAssignedIdentity]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (identity *UserAssignedIdentity) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20230131.UserAssignedIdentity)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131/UserAssignedIdentity, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20230131.UserAssignedIdentity)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131/UserAssignedIdentity, but got %T", oldResource)
	}
	validations := identity.updateValidations()
	var temp any = identity
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230131.UserAssignedIdentity]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (identity *UserAssignedIdentity) createValidations() []func(ctx context.Context, obj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20230131.UserAssignedIdentity) (admission.Warnings, error){
		identity.validateResourceReferences,
		identity.validateOwnerReference,
		identity.validateSecretDestinations,
		identity.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (identity *UserAssignedIdentity) deleteValidations() []func(ctx context.Context, obj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (identity *UserAssignedIdentity) updateValidations() []func(ctx context.Context, oldObj *v20230131.UserAssignedIdentity, newObj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20230131.UserAssignedIdentity, newObj *v20230131.UserAssignedIdentity) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20230131.UserAssignedIdentity, newObj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
			return identity.validateResourceReferences(ctx, newObj)
		},
		identity.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20230131.UserAssignedIdentity, newObj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
			return identity.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230131.UserAssignedIdentity, newObj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
			return identity.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230131.UserAssignedIdentity, newObj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
			return identity.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (identity *UserAssignedIdentity) validateConfigMapDestinations(ctx context.Context, obj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.ConfigMapDestination
	if obj.Spec.OperatorSpec.ConfigMaps != nil {
		toValidate = []*genruntime.ConfigMapDestination{
			obj.Spec.OperatorSpec.ConfigMaps.ClientId,
			obj.Spec.OperatorSpec.ConfigMaps.PrincipalId,
			obj.Spec.OperatorSpec.ConfigMaps.TenantId,
		}
	}
	return configmaps.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (identity *UserAssignedIdentity) validateOwnerReference(ctx context.Context, obj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (identity *UserAssignedIdentity) validateResourceReferences(ctx context.Context, obj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (identity *UserAssignedIdentity) validateSecretDestinations(ctx context.Context, obj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.SecretDestination
	if obj.Spec.OperatorSpec.Secrets != nil {
		toValidate = []*genruntime.SecretDestination{
			obj.Spec.OperatorSpec.Secrets.ClientId,
			obj.Spec.OperatorSpec.Secrets.PrincipalId,
			obj.Spec.OperatorSpec.Secrets.TenantId,
		}
	}
	return secrets.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (identity *UserAssignedIdentity) validateWriteOnceProperties(ctx context.Context, oldObj *v20230131.UserAssignedIdentity, newObj *v20230131.UserAssignedIdentity) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
