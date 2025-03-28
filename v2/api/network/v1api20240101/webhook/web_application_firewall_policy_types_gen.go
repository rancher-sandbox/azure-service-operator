// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20240101 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240101"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type WebApplicationFirewallPolicy struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240101-webapplicationfirewallpolicy,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=webapplicationfirewallpolicies,verbs=create;update,versions=v1api20240101,name=default.v1api20240101.webapplicationfirewallpolicies.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &WebApplicationFirewallPolicy{}

// Default applies defaults to the WebApplicationFirewallPolicy resource
func (policy *WebApplicationFirewallPolicy) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240101.WebApplicationFirewallPolicy)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240101/WebApplicationFirewallPolicy, but got %T", obj)
	}
	err := policy.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = policy
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (policy *WebApplicationFirewallPolicy) defaultAzureName(ctx context.Context, obj *v20240101.WebApplicationFirewallPolicy) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the WebApplicationFirewallPolicy resource
func (policy *WebApplicationFirewallPolicy) defaultImpl(ctx context.Context, obj *v20240101.WebApplicationFirewallPolicy) error {
	err := policy.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240101-webapplicationfirewallpolicy,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=webapplicationfirewallpolicies,verbs=create;update,versions=v1api20240101,name=validate.v1api20240101.webapplicationfirewallpolicies.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &WebApplicationFirewallPolicy{}

// ValidateCreate validates the creation of the resource
func (policy *WebApplicationFirewallPolicy) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240101.WebApplicationFirewallPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240101/WebApplicationFirewallPolicy, but got %T", resource)
	}
	validations := policy.createValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240101.WebApplicationFirewallPolicy]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (policy *WebApplicationFirewallPolicy) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240101.WebApplicationFirewallPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240101/WebApplicationFirewallPolicy, but got %T", resource)
	}
	validations := policy.deleteValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240101.WebApplicationFirewallPolicy]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (policy *WebApplicationFirewallPolicy) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240101.WebApplicationFirewallPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240101/WebApplicationFirewallPolicy, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240101.WebApplicationFirewallPolicy)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240101/WebApplicationFirewallPolicy, but got %T", oldResource)
	}
	validations := policy.updateValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240101.WebApplicationFirewallPolicy]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (policy *WebApplicationFirewallPolicy) createValidations() []func(ctx context.Context, obj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error){policy.validateResourceReferences, policy.validateOwnerReference, policy.validateSecretDestinations, policy.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (policy *WebApplicationFirewallPolicy) deleteValidations() []func(ctx context.Context, obj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (policy *WebApplicationFirewallPolicy) updateValidations() []func(ctx context.Context, oldObj *v20240101.WebApplicationFirewallPolicy, newObj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240101.WebApplicationFirewallPolicy, newObj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240101.WebApplicationFirewallPolicy, newObj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
			return policy.validateResourceReferences(ctx, newObj)
		},
		policy.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240101.WebApplicationFirewallPolicy, newObj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
			return policy.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240101.WebApplicationFirewallPolicy, newObj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
			return policy.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240101.WebApplicationFirewallPolicy, newObj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
			return policy.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (policy *WebApplicationFirewallPolicy) validateConfigMapDestinations(ctx context.Context, obj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (policy *WebApplicationFirewallPolicy) validateOwnerReference(ctx context.Context, obj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (policy *WebApplicationFirewallPolicy) validateResourceReferences(ctx context.Context, obj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (policy *WebApplicationFirewallPolicy) validateSecretDestinations(ctx context.Context, obj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (policy *WebApplicationFirewallPolicy) validateWriteOnceProperties(ctx context.Context, oldObj *v20240101.WebApplicationFirewallPolicy, newObj *v20240101.WebApplicationFirewallPolicy) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
