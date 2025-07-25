// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20220801 "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type ApiVersionSet struct {
}

// +kubebuilder:webhook:path=/mutate-apimanagement-azure-com-v1api20220801-apiversionset,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=apiversionsets,verbs=create;update,versions=v1api20220801,name=default.v1api20220801.apiversionsets.apimanagement.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &ApiVersionSet{}

// Default applies defaults to the ApiVersionSet resource
func (versionSet *ApiVersionSet) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20220801.ApiVersionSet)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/ApiVersionSet, but got %T", obj)
	}
	err := versionSet.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = versionSet
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (versionSet *ApiVersionSet) defaultAzureName(ctx context.Context, obj *v20220801.ApiVersionSet) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the ApiVersionSet resource
func (versionSet *ApiVersionSet) defaultImpl(ctx context.Context, obj *v20220801.ApiVersionSet) error {
	err := versionSet.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-apimanagement-azure-com-v1api20220801-apiversionset,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=apiversionsets,verbs=create;update,versions=v1api20220801,name=validate.v1api20220801.apiversionsets.apimanagement.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &ApiVersionSet{}

// ValidateCreate validates the creation of the resource
func (versionSet *ApiVersionSet) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220801.ApiVersionSet)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/ApiVersionSet, but got %T", resource)
	}
	validations := versionSet.createValidations()
	var temp any = versionSet
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220801.ApiVersionSet]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (versionSet *ApiVersionSet) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220801.ApiVersionSet)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/ApiVersionSet, but got %T", resource)
	}
	validations := versionSet.deleteValidations()
	var temp any = versionSet
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220801.ApiVersionSet]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (versionSet *ApiVersionSet) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20220801.ApiVersionSet)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/ApiVersionSet, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20220801.ApiVersionSet)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/ApiVersionSet, but got %T", oldResource)
	}
	validations := versionSet.updateValidations()
	var temp any = versionSet
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220801.ApiVersionSet]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (versionSet *ApiVersionSet) createValidations() []func(ctx context.Context, obj *v20220801.ApiVersionSet) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20220801.ApiVersionSet) (admission.Warnings, error){
		versionSet.validateResourceReferences,
		versionSet.validateOwnerReference,
		versionSet.validateSecretDestinations,
		versionSet.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (versionSet *ApiVersionSet) deleteValidations() []func(ctx context.Context, obj *v20220801.ApiVersionSet) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (versionSet *ApiVersionSet) updateValidations() []func(ctx context.Context, oldObj *v20220801.ApiVersionSet, newObj *v20220801.ApiVersionSet) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20220801.ApiVersionSet, newObj *v20220801.ApiVersionSet) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20220801.ApiVersionSet, newObj *v20220801.ApiVersionSet) (admission.Warnings, error) {
			return versionSet.validateResourceReferences(ctx, newObj)
		},
		versionSet.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20220801.ApiVersionSet, newObj *v20220801.ApiVersionSet) (admission.Warnings, error) {
			return versionSet.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220801.ApiVersionSet, newObj *v20220801.ApiVersionSet) (admission.Warnings, error) {
			return versionSet.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220801.ApiVersionSet, newObj *v20220801.ApiVersionSet) (admission.Warnings, error) {
			return versionSet.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (versionSet *ApiVersionSet) validateConfigMapDestinations(ctx context.Context, obj *v20220801.ApiVersionSet) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (versionSet *ApiVersionSet) validateOwnerReference(ctx context.Context, obj *v20220801.ApiVersionSet) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (versionSet *ApiVersionSet) validateResourceReferences(ctx context.Context, obj *v20220801.ApiVersionSet) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (versionSet *ApiVersionSet) validateSecretDestinations(ctx context.Context, obj *v20220801.ApiVersionSet) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (versionSet *ApiVersionSet) validateWriteOnceProperties(ctx context.Context, oldObj *v20220801.ApiVersionSet, newObj *v20220801.ApiVersionSet) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
