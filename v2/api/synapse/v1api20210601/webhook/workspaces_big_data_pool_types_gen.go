// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20210601 "github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type WorkspacesBigDataPool struct {
}

// +kubebuilder:webhook:path=/mutate-synapse-azure-com-v1api20210601-workspacesbigdatapool,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=synapse.azure.com,resources=workspacesbigdatapools,verbs=create;update,versions=v1api20210601,name=default.v1api20210601.workspacesbigdatapools.synapse.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &WorkspacesBigDataPool{}

// Default applies defaults to the WorkspacesBigDataPool resource
func (pool *WorkspacesBigDataPool) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20210601.WorkspacesBigDataPool)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601/WorkspacesBigDataPool, but got %T", obj)
	}
	err := pool.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = pool
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (pool *WorkspacesBigDataPool) defaultAzureName(ctx context.Context, obj *v20210601.WorkspacesBigDataPool) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the WorkspacesBigDataPool resource
func (pool *WorkspacesBigDataPool) defaultImpl(ctx context.Context, obj *v20210601.WorkspacesBigDataPool) error {
	err := pool.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-synapse-azure-com-v1api20210601-workspacesbigdatapool,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=synapse.azure.com,resources=workspacesbigdatapools,verbs=create;update,versions=v1api20210601,name=validate.v1api20210601.workspacesbigdatapools.synapse.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &WorkspacesBigDataPool{}

// ValidateCreate validates the creation of the resource
func (pool *WorkspacesBigDataPool) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210601.WorkspacesBigDataPool)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601/WorkspacesBigDataPool, but got %T", resource)
	}
	validations := pool.createValidations()
	var temp any = pool
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210601.WorkspacesBigDataPool]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (pool *WorkspacesBigDataPool) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210601.WorkspacesBigDataPool)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601/WorkspacesBigDataPool, but got %T", resource)
	}
	validations := pool.deleteValidations()
	var temp any = pool
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210601.WorkspacesBigDataPool]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (pool *WorkspacesBigDataPool) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20210601.WorkspacesBigDataPool)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601/WorkspacesBigDataPool, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20210601.WorkspacesBigDataPool)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/synapse/v1api20210601/WorkspacesBigDataPool, but got %T", oldResource)
	}
	validations := pool.updateValidations()
	var temp any = pool
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210601.WorkspacesBigDataPool]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (pool *WorkspacesBigDataPool) createValidations() []func(ctx context.Context, obj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error){
		pool.validateResourceReferences,
		pool.validateOwnerReference,
		pool.validateSecretDestinations,
		pool.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (pool *WorkspacesBigDataPool) deleteValidations() []func(ctx context.Context, obj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (pool *WorkspacesBigDataPool) updateValidations() []func(ctx context.Context, oldObj *v20210601.WorkspacesBigDataPool, newObj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20210601.WorkspacesBigDataPool, newObj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20210601.WorkspacesBigDataPool, newObj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
			return pool.validateResourceReferences(ctx, newObj)
		},
		pool.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20210601.WorkspacesBigDataPool, newObj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
			return pool.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210601.WorkspacesBigDataPool, newObj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
			return pool.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210601.WorkspacesBigDataPool, newObj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
			return pool.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (pool *WorkspacesBigDataPool) validateConfigMapDestinations(ctx context.Context, obj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (pool *WorkspacesBigDataPool) validateOwnerReference(ctx context.Context, obj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (pool *WorkspacesBigDataPool) validateResourceReferences(ctx context.Context, obj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (pool *WorkspacesBigDataPool) validateSecretDestinations(ctx context.Context, obj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (pool *WorkspacesBigDataPool) validateWriteOnceProperties(ctx context.Context, oldObj *v20210601.WorkspacesBigDataPool, newObj *v20210601.WorkspacesBigDataPool) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
