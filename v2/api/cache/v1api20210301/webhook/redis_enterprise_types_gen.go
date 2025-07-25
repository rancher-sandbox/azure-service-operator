// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20210301 "github.com/Azure/azure-service-operator/v2/api/cache/v1api20210301"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type RedisEnterprise struct {
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1api20210301-redisenterprise,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redisenterprises,verbs=create;update,versions=v1api20210301,name=default.v1api20210301.redisenterprises.cache.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &RedisEnterprise{}

// Default applies defaults to the RedisEnterprise resource
func (enterprise *RedisEnterprise) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20210301.RedisEnterprise)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cache/v1api20210301/RedisEnterprise, but got %T", obj)
	}
	err := enterprise.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = enterprise
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (enterprise *RedisEnterprise) defaultAzureName(ctx context.Context, obj *v20210301.RedisEnterprise) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the RedisEnterprise resource
func (enterprise *RedisEnterprise) defaultImpl(ctx context.Context, obj *v20210301.RedisEnterprise) error {
	err := enterprise.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1api20210301-redisenterprise,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redisenterprises,verbs=create;update,versions=v1api20210301,name=validate.v1api20210301.redisenterprises.cache.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &RedisEnterprise{}

// ValidateCreate validates the creation of the resource
func (enterprise *RedisEnterprise) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210301.RedisEnterprise)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cache/v1api20210301/RedisEnterprise, but got %T", resource)
	}
	validations := enterprise.createValidations()
	var temp any = enterprise
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210301.RedisEnterprise]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (enterprise *RedisEnterprise) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210301.RedisEnterprise)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cache/v1api20210301/RedisEnterprise, but got %T", resource)
	}
	validations := enterprise.deleteValidations()
	var temp any = enterprise
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210301.RedisEnterprise]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (enterprise *RedisEnterprise) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20210301.RedisEnterprise)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cache/v1api20210301/RedisEnterprise, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20210301.RedisEnterprise)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/cache/v1api20210301/RedisEnterprise, but got %T", oldResource)
	}
	validations := enterprise.updateValidations()
	var temp any = enterprise
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210301.RedisEnterprise]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (enterprise *RedisEnterprise) createValidations() []func(ctx context.Context, obj *v20210301.RedisEnterprise) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20210301.RedisEnterprise) (admission.Warnings, error){
		enterprise.validateResourceReferences,
		enterprise.validateOwnerReference,
		enterprise.validateSecretDestinations,
		enterprise.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (enterprise *RedisEnterprise) deleteValidations() []func(ctx context.Context, obj *v20210301.RedisEnterprise) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (enterprise *RedisEnterprise) updateValidations() []func(ctx context.Context, oldObj *v20210301.RedisEnterprise, newObj *v20210301.RedisEnterprise) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20210301.RedisEnterprise, newObj *v20210301.RedisEnterprise) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20210301.RedisEnterprise, newObj *v20210301.RedisEnterprise) (admission.Warnings, error) {
			return enterprise.validateResourceReferences(ctx, newObj)
		},
		enterprise.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20210301.RedisEnterprise, newObj *v20210301.RedisEnterprise) (admission.Warnings, error) {
			return enterprise.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210301.RedisEnterprise, newObj *v20210301.RedisEnterprise) (admission.Warnings, error) {
			return enterprise.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210301.RedisEnterprise, newObj *v20210301.RedisEnterprise) (admission.Warnings, error) {
			return enterprise.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (enterprise *RedisEnterprise) validateConfigMapDestinations(ctx context.Context, obj *v20210301.RedisEnterprise) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (enterprise *RedisEnterprise) validateOwnerReference(ctx context.Context, obj *v20210301.RedisEnterprise) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (enterprise *RedisEnterprise) validateResourceReferences(ctx context.Context, obj *v20210301.RedisEnterprise) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (enterprise *RedisEnterprise) validateSecretDestinations(ctx context.Context, obj *v20210301.RedisEnterprise) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (enterprise *RedisEnterprise) validateWriteOnceProperties(ctx context.Context, oldObj *v20210301.RedisEnterprise, newObj *v20210301.RedisEnterprise) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
