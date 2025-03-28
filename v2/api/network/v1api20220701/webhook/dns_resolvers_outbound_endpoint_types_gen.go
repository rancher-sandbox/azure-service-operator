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

type DnsResolversOutboundEndpoint struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20220701-dnsresolversoutboundendpoint,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=dnsresolversoutboundendpoints,verbs=create;update,versions=v1api20220701,name=default.v1api20220701.dnsresolversoutboundendpoints.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &DnsResolversOutboundEndpoint{}

// Default applies defaults to the DnsResolversOutboundEndpoint resource
func (endpoint *DnsResolversOutboundEndpoint) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20220701.DnsResolversOutboundEndpoint)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/DnsResolversOutboundEndpoint, but got %T", obj)
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
func (endpoint *DnsResolversOutboundEndpoint) defaultAzureName(ctx context.Context, obj *v20220701.DnsResolversOutboundEndpoint) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the DnsResolversOutboundEndpoint resource
func (endpoint *DnsResolversOutboundEndpoint) defaultImpl(ctx context.Context, obj *v20220701.DnsResolversOutboundEndpoint) error {
	err := endpoint.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20220701-dnsresolversoutboundendpoint,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=dnsresolversoutboundendpoints,verbs=create;update,versions=v1api20220701,name=validate.v1api20220701.dnsresolversoutboundendpoints.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &DnsResolversOutboundEndpoint{}

// ValidateCreate validates the creation of the resource
func (endpoint *DnsResolversOutboundEndpoint) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220701.DnsResolversOutboundEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/DnsResolversOutboundEndpoint, but got %T", resource)
	}
	validations := endpoint.createValidations()
	var temp any = endpoint
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220701.DnsResolversOutboundEndpoint]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (endpoint *DnsResolversOutboundEndpoint) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20220701.DnsResolversOutboundEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/DnsResolversOutboundEndpoint, but got %T", resource)
	}
	validations := endpoint.deleteValidations()
	var temp any = endpoint
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220701.DnsResolversOutboundEndpoint]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (endpoint *DnsResolversOutboundEndpoint) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20220701.DnsResolversOutboundEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/DnsResolversOutboundEndpoint, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20220701.DnsResolversOutboundEndpoint)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20220701/DnsResolversOutboundEndpoint, but got %T", oldResource)
	}
	validations := endpoint.updateValidations()
	var temp any = endpoint
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20220701.DnsResolversOutboundEndpoint]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (endpoint *DnsResolversOutboundEndpoint) createValidations() []func(ctx context.Context, obj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error){endpoint.validateResourceReferences, endpoint.validateOwnerReference, endpoint.validateSecretDestinations, endpoint.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (endpoint *DnsResolversOutboundEndpoint) deleteValidations() []func(ctx context.Context, obj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (endpoint *DnsResolversOutboundEndpoint) updateValidations() []func(ctx context.Context, oldObj *v20220701.DnsResolversOutboundEndpoint, newObj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20220701.DnsResolversOutboundEndpoint, newObj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20220701.DnsResolversOutboundEndpoint, newObj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
			return endpoint.validateResourceReferences(ctx, newObj)
		},
		endpoint.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20220701.DnsResolversOutboundEndpoint, newObj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
			return endpoint.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220701.DnsResolversOutboundEndpoint, newObj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
			return endpoint.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20220701.DnsResolversOutboundEndpoint, newObj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
			return endpoint.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (endpoint *DnsResolversOutboundEndpoint) validateConfigMapDestinations(ctx context.Context, obj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (endpoint *DnsResolversOutboundEndpoint) validateOwnerReference(ctx context.Context, obj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (endpoint *DnsResolversOutboundEndpoint) validateResourceReferences(ctx context.Context, obj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (endpoint *DnsResolversOutboundEndpoint) validateSecretDestinations(ctx context.Context, obj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (endpoint *DnsResolversOutboundEndpoint) validateWriteOnceProperties(ctx context.Context, oldObj *v20220701.DnsResolversOutboundEndpoint, newObj *v20220701.DnsResolversOutboundEndpoint) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
