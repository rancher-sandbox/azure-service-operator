// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20230501 "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type AfdCustomDomain struct {
}

// +kubebuilder:webhook:path=/mutate-cdn-azure-com-v1api20230501-afdcustomdomain,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cdn.azure.com,resources=afdcustomdomains,verbs=create;update,versions=v1api20230501,name=default.v1api20230501.afdcustomdomains.cdn.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &AfdCustomDomain{}

// Default applies defaults to the AfdCustomDomain resource
func (domain *AfdCustomDomain) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20230501.AfdCustomDomain)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/AfdCustomDomain, but got %T", obj)
	}
	err := domain.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = domain
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (domain *AfdCustomDomain) defaultAzureName(ctx context.Context, obj *v20230501.AfdCustomDomain) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the AfdCustomDomain resource
func (domain *AfdCustomDomain) defaultImpl(ctx context.Context, obj *v20230501.AfdCustomDomain) error {
	err := domain.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-cdn-azure-com-v1api20230501-afdcustomdomain,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cdn.azure.com,resources=afdcustomdomains,verbs=create;update,versions=v1api20230501,name=validate.v1api20230501.afdcustomdomains.cdn.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &AfdCustomDomain{}

// ValidateCreate validates the creation of the resource
func (domain *AfdCustomDomain) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230501.AfdCustomDomain)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/AfdCustomDomain, but got %T", resource)
	}
	validations := domain.createValidations()
	var temp any = domain
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501.AfdCustomDomain]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (domain *AfdCustomDomain) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230501.AfdCustomDomain)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/AfdCustomDomain, but got %T", resource)
	}
	validations := domain.deleteValidations()
	var temp any = domain
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501.AfdCustomDomain]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (domain *AfdCustomDomain) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20230501.AfdCustomDomain)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/AfdCustomDomain, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20230501.AfdCustomDomain)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/AfdCustomDomain, but got %T", oldResource)
	}
	validations := domain.updateValidations()
	var temp any = domain
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501.AfdCustomDomain]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (domain *AfdCustomDomain) createValidations() []func(ctx context.Context, obj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20230501.AfdCustomDomain) (admission.Warnings, error){
		domain.validateResourceReferences,
		domain.validateOwnerReference,
		domain.validateSecretDestinations,
		domain.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (domain *AfdCustomDomain) deleteValidations() []func(ctx context.Context, obj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (domain *AfdCustomDomain) updateValidations() []func(ctx context.Context, oldObj *v20230501.AfdCustomDomain, newObj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20230501.AfdCustomDomain, newObj *v20230501.AfdCustomDomain) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20230501.AfdCustomDomain, newObj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
			return domain.validateResourceReferences(ctx, newObj)
		},
		domain.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20230501.AfdCustomDomain, newObj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
			return domain.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230501.AfdCustomDomain, newObj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
			return domain.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230501.AfdCustomDomain, newObj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
			return domain.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (domain *AfdCustomDomain) validateConfigMapDestinations(ctx context.Context, obj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (domain *AfdCustomDomain) validateOwnerReference(ctx context.Context, obj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (domain *AfdCustomDomain) validateResourceReferences(ctx context.Context, obj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (domain *AfdCustomDomain) validateSecretDestinations(ctx context.Context, obj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (domain *AfdCustomDomain) validateWriteOnceProperties(ctx context.Context, oldObj *v20230501.AfdCustomDomain, newObj *v20230501.AfdCustomDomain) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
