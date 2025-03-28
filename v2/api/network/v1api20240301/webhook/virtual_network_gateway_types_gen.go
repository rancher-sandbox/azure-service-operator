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

type VirtualNetworkGateway struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240301-virtualnetworkgateway,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=virtualnetworkgateways,verbs=create;update,versions=v1api20240301,name=default.v1api20240301.virtualnetworkgateways.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &VirtualNetworkGateway{}

// Default applies defaults to the VirtualNetworkGateway resource
func (gateway *VirtualNetworkGateway) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240301.VirtualNetworkGateway)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/VirtualNetworkGateway, but got %T", obj)
	}
	err := gateway.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = gateway
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (gateway *VirtualNetworkGateway) defaultAzureName(ctx context.Context, obj *v20240301.VirtualNetworkGateway) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the VirtualNetworkGateway resource
func (gateway *VirtualNetworkGateway) defaultImpl(ctx context.Context, obj *v20240301.VirtualNetworkGateway) error {
	err := gateway.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240301-virtualnetworkgateway,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=virtualnetworkgateways,verbs=create;update,versions=v1api20240301,name=validate.v1api20240301.virtualnetworkgateways.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &VirtualNetworkGateway{}

// ValidateCreate validates the creation of the resource
func (gateway *VirtualNetworkGateway) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.VirtualNetworkGateway)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/VirtualNetworkGateway, but got %T", resource)
	}
	validations := gateway.createValidations()
	var temp any = gateway
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.VirtualNetworkGateway]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (gateway *VirtualNetworkGateway) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.VirtualNetworkGateway)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/VirtualNetworkGateway, but got %T", resource)
	}
	validations := gateway.deleteValidations()
	var temp any = gateway
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.VirtualNetworkGateway]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (gateway *VirtualNetworkGateway) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240301.VirtualNetworkGateway)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/VirtualNetworkGateway, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240301.VirtualNetworkGateway)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/VirtualNetworkGateway, but got %T", oldResource)
	}
	validations := gateway.updateValidations()
	var temp any = gateway
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.VirtualNetworkGateway]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (gateway *VirtualNetworkGateway) createValidations() []func(ctx context.Context, obj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240301.VirtualNetworkGateway) (admission.Warnings, error){gateway.validateResourceReferences, gateway.validateOwnerReference, gateway.validateSecretDestinations, gateway.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (gateway *VirtualNetworkGateway) deleteValidations() []func(ctx context.Context, obj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (gateway *VirtualNetworkGateway) updateValidations() []func(ctx context.Context, oldObj *v20240301.VirtualNetworkGateway, newObj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240301.VirtualNetworkGateway, newObj *v20240301.VirtualNetworkGateway) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240301.VirtualNetworkGateway, newObj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
			return gateway.validateResourceReferences(ctx, newObj)
		},
		gateway.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240301.VirtualNetworkGateway, newObj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
			return gateway.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.VirtualNetworkGateway, newObj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
			return gateway.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.VirtualNetworkGateway, newObj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
			return gateway.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (gateway *VirtualNetworkGateway) validateConfigMapDestinations(ctx context.Context, obj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (gateway *VirtualNetworkGateway) validateOwnerReference(ctx context.Context, obj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (gateway *VirtualNetworkGateway) validateResourceReferences(ctx context.Context, obj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (gateway *VirtualNetworkGateway) validateSecretDestinations(ctx context.Context, obj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (gateway *VirtualNetworkGateway) validateWriteOnceProperties(ctx context.Context, oldObj *v20240301.VirtualNetworkGateway, newObj *v20240301.VirtualNetworkGateway) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
