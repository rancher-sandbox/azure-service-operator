// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20240301 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type AzureFirewall struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240301-azurefirewall,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=azurefirewalls,verbs=create;update,versions=v1api20240301,name=default.v1api20240301.azurefirewalls.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &AzureFirewall{}

// Default applies defaults to the AzureFirewall resource
func (firewall *AzureFirewall) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240301.AzureFirewall)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/AzureFirewall, but got %T", obj)
	}
	err := firewall.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = firewall
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (firewall *AzureFirewall) defaultAzureName(ctx context.Context, obj *v20240301.AzureFirewall) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the AzureFirewall resource
func (firewall *AzureFirewall) defaultImpl(ctx context.Context, obj *v20240301.AzureFirewall) error {
	err := firewall.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240301-azurefirewall,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=azurefirewalls,verbs=create;update,versions=v1api20240301,name=validate.v1api20240301.azurefirewalls.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &AzureFirewall{}

// ValidateCreate validates the creation of the resource
func (firewall *AzureFirewall) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.AzureFirewall)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/AzureFirewall, but got %T", resource)
	}
	validations := firewall.createValidations()
	var temp any = firewall
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.AzureFirewall]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (firewall *AzureFirewall) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.AzureFirewall)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/AzureFirewall, but got %T", resource)
	}
	validations := firewall.deleteValidations()
	var temp any = firewall
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.AzureFirewall]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (firewall *AzureFirewall) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240301.AzureFirewall)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/AzureFirewall, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240301.AzureFirewall)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/AzureFirewall, but got %T", oldResource)
	}
	validations := firewall.updateValidations()
	var temp any = firewall
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.AzureFirewall]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (firewall *AzureFirewall) createValidations() []func(ctx context.Context, obj *v20240301.AzureFirewall) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240301.AzureFirewall) (admission.Warnings, error){
		firewall.validateResourceReferences,
		firewall.validateOwnerReference,
		firewall.validateSecretDestinations,
		firewall.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (firewall *AzureFirewall) deleteValidations() []func(ctx context.Context, obj *v20240301.AzureFirewall) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (firewall *AzureFirewall) updateValidations() []func(ctx context.Context, oldObj *v20240301.AzureFirewall, newObj *v20240301.AzureFirewall) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240301.AzureFirewall, newObj *v20240301.AzureFirewall) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240301.AzureFirewall, newObj *v20240301.AzureFirewall) (admission.Warnings, error) {
			return firewall.validateResourceReferences(ctx, newObj)
		},
		firewall.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240301.AzureFirewall, newObj *v20240301.AzureFirewall) (admission.Warnings, error) {
			return firewall.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.AzureFirewall, newObj *v20240301.AzureFirewall) (admission.Warnings, error) {
			return firewall.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.AzureFirewall, newObj *v20240301.AzureFirewall) (admission.Warnings, error) {
			return firewall.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (firewall *AzureFirewall) validateConfigMapDestinations(ctx context.Context, obj *v20240301.AzureFirewall) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (firewall *AzureFirewall) validateOwnerReference(ctx context.Context, obj *v20240301.AzureFirewall) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (firewall *AzureFirewall) validateResourceReferences(ctx context.Context, obj *v20240301.AzureFirewall) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (firewall *AzureFirewall) validateSecretDestinations(ctx context.Context, obj *v20240301.AzureFirewall) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (firewall *AzureFirewall) validateWriteOnceProperties(ctx context.Context, oldObj *v20240301.AzureFirewall, newObj *v20240301.AzureFirewall) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
