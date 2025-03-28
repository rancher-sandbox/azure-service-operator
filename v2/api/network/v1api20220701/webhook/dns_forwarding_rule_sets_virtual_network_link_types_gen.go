// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20220701 "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type DnsForwardingRuleSetsVirtualNetworkLink struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20220701-dnsforwardingrulesetsvirtualnetworklink,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=dnsforwardingrulesetsvirtualnetworklinks,verbs=create;update,versions=v1api20220701,name=default.v1api20220701.dnsforwardingrulesetsvirtualnetworklinks.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &DnsForwardingRuleSetsVirtualNetworkLink{}

// Default applies defaults to the DnsForwardingRuleSetsVirtualNetworkLink resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20220701.DnsForwardingRuleSetsVirtualNetworkLink)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/DnsForwardingRuleSetsVirtualNetworkLink, but got %T", obj)
	}
	err := link.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = link
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (link *DnsForwardingRuleSetsVirtualNetworkLink) defaultAzureName(ctx context.Context, obj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the DnsForwardingRuleSetsVirtualNetworkLink resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) defaultImpl(ctx context.Context, obj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) error {
	err := link.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20220701-dnsforwardingrulesetsvirtualnetworklink,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=dnsforwardingrulesetsvirtualnetworklinks,verbs=create;update,versions=v1api20220701,name=validate.v1api20220701.dnsforwardingrulesetsvirtualnetworklinks.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &DnsForwardingRuleSetsVirtualNetworkLink{}

// ValidateCreate validates the creation of the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220701.DnsForwardingRuleSetsVirtualNetworkLink)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/DnsForwardingRuleSetsVirtualNetworkLink, but got %T", resource)
	}
	validations := link.createValidations()
	var temp any = link
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220701.DnsForwardingRuleSetsVirtualNetworkLink]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220701.DnsForwardingRuleSetsVirtualNetworkLink)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/DnsForwardingRuleSetsVirtualNetworkLink, but got %T", resource)
	}
	validations := link.deleteValidations()
	var temp any = link
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220701.DnsForwardingRuleSetsVirtualNetworkLink]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20220701.DnsForwardingRuleSetsVirtualNetworkLink)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/DnsForwardingRuleSetsVirtualNetworkLink, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20220701.DnsForwardingRuleSetsVirtualNetworkLink)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/DnsForwardingRuleSetsVirtualNetworkLink, but got %T", oldResource)
	}
	validations := link.updateValidations()
	var temp any = link
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220701.DnsForwardingRuleSetsVirtualNetworkLink]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) createValidations() []func(ctx context.Context, obj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error){link.validateResourceReferences, link.validateOwnerReference, link.validateSecretDestinations, link.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) deleteValidations() []func(ctx context.Context, obj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) updateValidations() []func(ctx context.Context, oldObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink, newObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink, newObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink, newObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
			return link.validateResourceReferences(ctx, newObj)
		},
		link.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink, newObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
			return link.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink, newObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
			return link.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink, newObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
			return link.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (link *DnsForwardingRuleSetsVirtualNetworkLink) validateConfigMapDestinations(ctx context.Context, obj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (link *DnsForwardingRuleSetsVirtualNetworkLink) validateOwnerReference(ctx context.Context, obj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (link *DnsForwardingRuleSetsVirtualNetworkLink) validateResourceReferences(ctx context.Context, obj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (link *DnsForwardingRuleSetsVirtualNetworkLink) validateSecretDestinations(ctx context.Context, obj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (link *DnsForwardingRuleSetsVirtualNetworkLink) validateWriteOnceProperties(ctx context.Context, oldObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink, newObj *v20220701.DnsForwardingRuleSetsVirtualNetworkLink) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
