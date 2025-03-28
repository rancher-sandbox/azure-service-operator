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

type AfdEndpoint struct {
}

// +kubebuilder:webhook:path=/mutate-cdn-azure-com-v1api20230501-afdendpoint,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cdn.azure.com,resources=afdendpoints,verbs=create;update,versions=v1api20230501,name=default.v1api20230501.afdendpoints.cdn.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &AfdEndpoint{}

// Default applies defaults to the AfdEndpoint resource
func (endpoint *AfdEndpoint) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20230501.AfdEndpoint)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/AfdEndpoint, but got %T", obj)
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
func (endpoint *AfdEndpoint) defaultAzureName(ctx context.Context, obj *v20230501.AfdEndpoint) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the AfdEndpoint resource
func (endpoint *AfdEndpoint) defaultImpl(ctx context.Context, obj *v20230501.AfdEndpoint) error {
	err := endpoint.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-cdn-azure-com-v1api20230501-afdendpoint,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cdn.azure.com,resources=afdendpoints,verbs=create;update,versions=v1api20230501,name=validate.v1api20230501.afdendpoints.cdn.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &AfdEndpoint{}

// ValidateCreate validates the creation of the resource
func (endpoint *AfdEndpoint) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230501.AfdEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/AfdEndpoint, but got %T", resource)
	}
	validations := endpoint.createValidations()
	var temp any = endpoint
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501.AfdEndpoint]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (endpoint *AfdEndpoint) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230501.AfdEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/AfdEndpoint, but got %T", resource)
	}
	validations := endpoint.deleteValidations()
	var temp any = endpoint
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501.AfdEndpoint]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (endpoint *AfdEndpoint) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20230501.AfdEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/AfdEndpoint, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20230501.AfdEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501/AfdEndpoint, but got %T", oldResource)
	}
	validations := endpoint.updateValidations()
	var temp any = endpoint
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501.AfdEndpoint]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (endpoint *AfdEndpoint) createValidations() []func(ctx context.Context, obj *v20230501.AfdEndpoint) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20230501.AfdEndpoint) (admission.Warnings, error){endpoint.validateResourceReferences, endpoint.validateOwnerReference, endpoint.validateSecretDestinations, endpoint.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (endpoint *AfdEndpoint) deleteValidations() []func(ctx context.Context, obj *v20230501.AfdEndpoint) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (endpoint *AfdEndpoint) updateValidations() []func(ctx context.Context, oldObj *v20230501.AfdEndpoint, newObj *v20230501.AfdEndpoint) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20230501.AfdEndpoint, newObj *v20230501.AfdEndpoint) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20230501.AfdEndpoint, newObj *v20230501.AfdEndpoint) (admission.Warnings, error) {
			return endpoint.validateResourceReferences(ctx, newObj)
		},
		endpoint.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20230501.AfdEndpoint, newObj *v20230501.AfdEndpoint) (admission.Warnings, error) {
			return endpoint.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230501.AfdEndpoint, newObj *v20230501.AfdEndpoint) (admission.Warnings, error) {
			return endpoint.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230501.AfdEndpoint, newObj *v20230501.AfdEndpoint) (admission.Warnings, error) {
			return endpoint.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (endpoint *AfdEndpoint) validateConfigMapDestinations(ctx context.Context, obj *v20230501.AfdEndpoint) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (endpoint *AfdEndpoint) validateOwnerReference(ctx context.Context, obj *v20230501.AfdEndpoint) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (endpoint *AfdEndpoint) validateResourceReferences(ctx context.Context, obj *v20230501.AfdEndpoint) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (endpoint *AfdEndpoint) validateSecretDestinations(ctx context.Context, obj *v20230501.AfdEndpoint) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (endpoint *AfdEndpoint) validateWriteOnceProperties(ctx context.Context, oldObj *v20230501.AfdEndpoint, newObj *v20230501.AfdEndpoint) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
