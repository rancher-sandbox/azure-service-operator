// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20221001p "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type NamespacesTopicsSubscriptionsRule struct {
}

// +kubebuilder:webhook:path=/mutate-servicebus-azure-com-v1api20221001preview-namespacestopicssubscriptionsrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=servicebus.azure.com,resources=namespacestopicssubscriptionsrules,verbs=create;update,versions=v1api20221001preview,name=default.v1api20221001preview.namespacestopicssubscriptionsrules.servicebus.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &NamespacesTopicsSubscriptionsRule{}

// Default applies defaults to the NamespacesTopicsSubscriptionsRule resource
func (rule *NamespacesTopicsSubscriptionsRule) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20221001p.NamespacesTopicsSubscriptionsRule)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview/NamespacesTopicsSubscriptionsRule, but got %T", obj)
	}
	err := rule.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *NamespacesTopicsSubscriptionsRule) defaultAzureName(ctx context.Context, obj *v20221001p.NamespacesTopicsSubscriptionsRule) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the NamespacesTopicsSubscriptionsRule resource
func (rule *NamespacesTopicsSubscriptionsRule) defaultImpl(ctx context.Context, obj *v20221001p.NamespacesTopicsSubscriptionsRule) error {
	err := rule.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-servicebus-azure-com-v1api20221001preview-namespacestopicssubscriptionsrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=servicebus.azure.com,resources=namespacestopicssubscriptionsrules,verbs=create;update,versions=v1api20221001preview,name=validate.v1api20221001preview.namespacestopicssubscriptionsrules.servicebus.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &NamespacesTopicsSubscriptionsRule{}

// ValidateCreate validates the creation of the resource
func (rule *NamespacesTopicsSubscriptionsRule) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20221001p.NamespacesTopicsSubscriptionsRule)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview/NamespacesTopicsSubscriptionsRule, but got %T", resource)
	}
	validations := rule.createValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20221001p.NamespacesTopicsSubscriptionsRule]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (rule *NamespacesTopicsSubscriptionsRule) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20221001p.NamespacesTopicsSubscriptionsRule)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview/NamespacesTopicsSubscriptionsRule, but got %T", resource)
	}
	validations := rule.deleteValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20221001p.NamespacesTopicsSubscriptionsRule]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (rule *NamespacesTopicsSubscriptionsRule) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20221001p.NamespacesTopicsSubscriptionsRule)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview/NamespacesTopicsSubscriptionsRule, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20221001p.NamespacesTopicsSubscriptionsRule)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview/NamespacesTopicsSubscriptionsRule, but got %T", oldResource)
	}
	validations := rule.updateValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20221001p.NamespacesTopicsSubscriptionsRule]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (rule *NamespacesTopicsSubscriptionsRule) createValidations() []func(ctx context.Context, obj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error){rule.validateResourceReferences, rule.validateOwnerReference, rule.validateSecretDestinations, rule.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (rule *NamespacesTopicsSubscriptionsRule) deleteValidations() []func(ctx context.Context, obj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (rule *NamespacesTopicsSubscriptionsRule) updateValidations() []func(ctx context.Context, oldObj *v20221001p.NamespacesTopicsSubscriptionsRule, newObj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20221001p.NamespacesTopicsSubscriptionsRule, newObj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20221001p.NamespacesTopicsSubscriptionsRule, newObj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
			return rule.validateResourceReferences(ctx, newObj)
		},
		rule.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20221001p.NamespacesTopicsSubscriptionsRule, newObj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
			return rule.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20221001p.NamespacesTopicsSubscriptionsRule, newObj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
			return rule.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20221001p.NamespacesTopicsSubscriptionsRule, newObj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
			return rule.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (rule *NamespacesTopicsSubscriptionsRule) validateConfigMapDestinations(ctx context.Context, obj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (rule *NamespacesTopicsSubscriptionsRule) validateOwnerReference(ctx context.Context, obj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (rule *NamespacesTopicsSubscriptionsRule) validateResourceReferences(ctx context.Context, obj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (rule *NamespacesTopicsSubscriptionsRule) validateSecretDestinations(ctx context.Context, obj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *NamespacesTopicsSubscriptionsRule) validateWriteOnceProperties(ctx context.Context, oldObj *v20221001p.NamespacesTopicsSubscriptionsRule, newObj *v20221001p.NamespacesTopicsSubscriptionsRule) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
