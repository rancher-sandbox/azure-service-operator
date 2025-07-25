// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20211101 "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type ServersOutboundFirewallRule struct {
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversoutboundfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversoutboundfirewallrules,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversoutboundfirewallrules.sql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &ServersOutboundFirewallRule{}

// Default applies defaults to the ServersOutboundFirewallRule resource
func (rule *ServersOutboundFirewallRule) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20211101.ServersOutboundFirewallRule)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersOutboundFirewallRule, but got %T", obj)
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
func (rule *ServersOutboundFirewallRule) defaultAzureName(ctx context.Context, obj *v20211101.ServersOutboundFirewallRule) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the ServersOutboundFirewallRule resource
func (rule *ServersOutboundFirewallRule) defaultImpl(ctx context.Context, obj *v20211101.ServersOutboundFirewallRule) error {
	err := rule.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversoutboundfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversoutboundfirewallrules,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversoutboundfirewallrules.sql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &ServersOutboundFirewallRule{}

// ValidateCreate validates the creation of the resource
func (rule *ServersOutboundFirewallRule) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.ServersOutboundFirewallRule)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersOutboundFirewallRule, but got %T", resource)
	}
	validations := rule.createValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersOutboundFirewallRule]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (rule *ServersOutboundFirewallRule) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.ServersOutboundFirewallRule)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersOutboundFirewallRule, but got %T", resource)
	}
	validations := rule.deleteValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersOutboundFirewallRule]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (rule *ServersOutboundFirewallRule) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20211101.ServersOutboundFirewallRule)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersOutboundFirewallRule, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20211101.ServersOutboundFirewallRule)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersOutboundFirewallRule, but got %T", oldResource)
	}
	validations := rule.updateValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersOutboundFirewallRule]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (rule *ServersOutboundFirewallRule) createValidations() []func(ctx context.Context, obj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error){
		rule.validateResourceReferences,
		rule.validateOwnerReference,
		rule.validateSecretDestinations,
		rule.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (rule *ServersOutboundFirewallRule) deleteValidations() []func(ctx context.Context, obj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (rule *ServersOutboundFirewallRule) updateValidations() []func(ctx context.Context, oldObj *v20211101.ServersOutboundFirewallRule, newObj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20211101.ServersOutboundFirewallRule, newObj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20211101.ServersOutboundFirewallRule, newObj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
			return rule.validateResourceReferences(ctx, newObj)
		},
		rule.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20211101.ServersOutboundFirewallRule, newObj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
			return rule.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.ServersOutboundFirewallRule, newObj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
			return rule.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.ServersOutboundFirewallRule, newObj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
			return rule.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (rule *ServersOutboundFirewallRule) validateConfigMapDestinations(ctx context.Context, obj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (rule *ServersOutboundFirewallRule) validateOwnerReference(ctx context.Context, obj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (rule *ServersOutboundFirewallRule) validateResourceReferences(ctx context.Context, obj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (rule *ServersOutboundFirewallRule) validateSecretDestinations(ctx context.Context, obj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *ServersOutboundFirewallRule) validateWriteOnceProperties(ctx context.Context, oldObj *v20211101.ServersOutboundFirewallRule, newObj *v20211101.ServersOutboundFirewallRule) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
