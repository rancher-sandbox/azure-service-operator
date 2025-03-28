// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20230501p "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type Service struct {
}

// +kubebuilder:webhook:path=/mutate-apimanagement-azure-com-v1api20230501preview-service,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=services,verbs=create;update,versions=v1api20230501preview,name=default.v1api20230501preview.services.apimanagement.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &Service{}

// Default applies defaults to the Service resource
func (service *Service) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20230501p.Service)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/Service, but got %T", obj)
	}
	err := service.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = service
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (service *Service) defaultAzureName(ctx context.Context, obj *v20230501p.Service) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the Service resource
func (service *Service) defaultImpl(ctx context.Context, obj *v20230501p.Service) error {
	err := service.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-apimanagement-azure-com-v1api20230501preview-service,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=services,verbs=create;update,versions=v1api20230501preview,name=validate.v1api20230501preview.services.apimanagement.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &Service{}

// ValidateCreate validates the creation of the resource
func (service *Service) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230501p.Service)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/Service, but got %T", resource)
	}
	validations := service.createValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501p.Service]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (service *Service) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230501p.Service)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/Service, but got %T", resource)
	}
	validations := service.deleteValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501p.Service]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (service *Service) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20230501p.Service)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/Service, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20230501p.Service)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/Service, but got %T", oldResource)
	}
	validations := service.updateValidations()
	var temp any = service
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230501p.Service]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (service *Service) createValidations() []func(ctx context.Context, obj *v20230501p.Service) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20230501p.Service) (admission.Warnings, error){service.validateResourceReferences, service.validateOwnerReference, service.validateSecretDestinations, service.validateConfigMapDestinations, service.validateOptionalConfigMapReferences}
}

// deleteValidations validates the deletion of the resource
func (service *Service) deleteValidations() []func(ctx context.Context, obj *v20230501p.Service) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (service *Service) updateValidations() []func(ctx context.Context, oldObj *v20230501p.Service, newObj *v20230501p.Service) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20230501p.Service, newObj *v20230501p.Service) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20230501p.Service, newObj *v20230501p.Service) (admission.Warnings, error) {
			return service.validateResourceReferences(ctx, newObj)
		},
		service.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20230501p.Service, newObj *v20230501p.Service) (admission.Warnings, error) {
			return service.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230501p.Service, newObj *v20230501p.Service) (admission.Warnings, error) {
			return service.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230501p.Service, newObj *v20230501p.Service) (admission.Warnings, error) {
			return service.validateConfigMapDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230501p.Service, newObj *v20230501p.Service) (admission.Warnings, error) {
			return service.validateOptionalConfigMapReferences(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (service *Service) validateConfigMapDestinations(ctx context.Context, obj *v20230501p.Service) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOptionalConfigMapReferences validates all optional configmap reference pairs to ensure that at most 1 is set
func (service *Service) validateOptionalConfigMapReferences(ctx context.Context, obj *v20230501p.Service) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return configmaps.ValidateOptionalReferences(refs)
}

// validateOwnerReference validates the owner field
func (service *Service) validateOwnerReference(ctx context.Context, obj *v20230501p.Service) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (service *Service) validateResourceReferences(ctx context.Context, obj *v20230501p.Service) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (service *Service) validateSecretDestinations(ctx context.Context, obj *v20230501p.Service) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (service *Service) validateWriteOnceProperties(ctx context.Context, oldObj *v20230501p.Service, newObj *v20230501p.Service) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
