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

type PrivateEndpoint struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240301-privateendpoint,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=privateendpoints,verbs=create;update,versions=v1api20240301,name=default.v1api20240301.privateendpoints.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &PrivateEndpoint{}

// Default applies defaults to the PrivateEndpoint resource
func (endpoint *PrivateEndpoint) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240301.PrivateEndpoint)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/PrivateEndpoint, but got %T", obj)
	}
	err := endpoint.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = endpoint
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (endpoint *PrivateEndpoint) defaultAzureName(ctx context.Context, obj *v20240301.PrivateEndpoint) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the PrivateEndpoint resource
func (endpoint *PrivateEndpoint) defaultImpl(ctx context.Context, obj *v20240301.PrivateEndpoint) error {
	err := endpoint.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240301-privateendpoint,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=privateendpoints,verbs=create;update,versions=v1api20240301,name=validate.v1api20240301.privateendpoints.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &PrivateEndpoint{}

// ValidateCreate validates the creation of the resource
func (endpoint *PrivateEndpoint) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.PrivateEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/PrivateEndpoint, but got %T", resource)
	}
	validations := endpoint.createValidations()
	var temp any = endpoint
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.PrivateEndpoint]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (endpoint *PrivateEndpoint) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240301.PrivateEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/PrivateEndpoint, but got %T", resource)
	}
	validations := endpoint.deleteValidations()
	var temp any = endpoint
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.PrivateEndpoint]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (endpoint *PrivateEndpoint) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240301.PrivateEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/PrivateEndpoint, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240301.PrivateEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/PrivateEndpoint, but got %T", oldResource)
	}
	validations := endpoint.updateValidations()
	var temp any = endpoint
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240301.PrivateEndpoint]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (endpoint *PrivateEndpoint) createValidations() []func(ctx context.Context, obj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240301.PrivateEndpoint) (admission.Warnings, error){
		endpoint.validateResourceReferences,
		endpoint.validateOwnerReference,
		endpoint.validateSecretDestinations,
		endpoint.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (endpoint *PrivateEndpoint) deleteValidations() []func(ctx context.Context, obj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (endpoint *PrivateEndpoint) updateValidations() []func(ctx context.Context, oldObj *v20240301.PrivateEndpoint, newObj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240301.PrivateEndpoint, newObj *v20240301.PrivateEndpoint) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240301.PrivateEndpoint, newObj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
			return endpoint.validateResourceReferences(ctx, newObj)
		},
		endpoint.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240301.PrivateEndpoint, newObj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
			return endpoint.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.PrivateEndpoint, newObj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
			return endpoint.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240301.PrivateEndpoint, newObj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
			return endpoint.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (endpoint *PrivateEndpoint) validateConfigMapDestinations(ctx context.Context, obj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.ConfigMapDestination
	if obj.Spec.OperatorSpec.ConfigMaps != nil {
		toValidate = []*genruntime.ConfigMapDestination{
			obj.Spec.OperatorSpec.ConfigMaps.PrimaryNicPrivateIpAddress,
		}
	}
	return configmaps.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (endpoint *PrivateEndpoint) validateOwnerReference(ctx context.Context, obj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (endpoint *PrivateEndpoint) validateResourceReferences(ctx context.Context, obj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (endpoint *PrivateEndpoint) validateSecretDestinations(ctx context.Context, obj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (endpoint *PrivateEndpoint) validateWriteOnceProperties(ctx context.Context, oldObj *v20240301.PrivateEndpoint, newObj *v20240301.PrivateEndpoint) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
