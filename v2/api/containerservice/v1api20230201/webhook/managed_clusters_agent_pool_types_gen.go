// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20230201 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type ManagedClustersAgentPool struct {
}

// +kubebuilder:webhook:path=/mutate-containerservice-azure-com-v1api20230201-managedclustersagentpool,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=containerservice.azure.com,resources=managedclustersagentpools,verbs=create;update,versions=v1api20230201,name=default.v1api20230201.managedclustersagentpools.containerservice.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &ManagedClustersAgentPool{}

// Default applies defaults to the ManagedClustersAgentPool resource
func (pool *ManagedClustersAgentPool) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20230201.ManagedClustersAgentPool)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201/ManagedClustersAgentPool, but got %T", obj)
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
func (pool *ManagedClustersAgentPool) defaultAzureName(ctx context.Context, obj *v20230201.ManagedClustersAgentPool) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the ManagedClustersAgentPool resource
func (pool *ManagedClustersAgentPool) defaultImpl(ctx context.Context, obj *v20230201.ManagedClustersAgentPool) error {
	err := pool.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-containerservice-azure-com-v1api20230201-managedclustersagentpool,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=containerservice.azure.com,resources=managedclustersagentpools,verbs=create;update,versions=v1api20230201,name=validate.v1api20230201.managedclustersagentpools.containerservice.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &ManagedClustersAgentPool{}

// ValidateCreate validates the creation of the resource
func (pool *ManagedClustersAgentPool) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230201.ManagedClustersAgentPool)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201/ManagedClustersAgentPool, but got %T", resource)
	}
	validations := pool.createValidations()
	var temp any = pool
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230201.ManagedClustersAgentPool]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (pool *ManagedClustersAgentPool) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20230201.ManagedClustersAgentPool)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201/ManagedClustersAgentPool, but got %T", resource)
	}
	validations := pool.deleteValidations()
	var temp any = pool
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230201.ManagedClustersAgentPool]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (pool *ManagedClustersAgentPool) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20230201.ManagedClustersAgentPool)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201/ManagedClustersAgentPool, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20230201.ManagedClustersAgentPool)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201/ManagedClustersAgentPool, but got %T", oldResource)
	}
	validations := pool.updateValidations()
	var temp any = pool
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20230201.ManagedClustersAgentPool]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (pool *ManagedClustersAgentPool) createValidations() []func(ctx context.Context, obj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error){pool.validateResourceReferences, pool.validateOwnerReference, pool.validateSecretDestinations, pool.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (pool *ManagedClustersAgentPool) deleteValidations() []func(ctx context.Context, obj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (pool *ManagedClustersAgentPool) updateValidations() []func(ctx context.Context, oldObj *v20230201.ManagedClustersAgentPool, newObj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20230201.ManagedClustersAgentPool, newObj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20230201.ManagedClustersAgentPool, newObj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
			return pool.validateResourceReferences(ctx, newObj)
		},
		pool.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20230201.ManagedClustersAgentPool, newObj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
			return pool.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230201.ManagedClustersAgentPool, newObj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
			return pool.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20230201.ManagedClustersAgentPool, newObj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
			return pool.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (pool *ManagedClustersAgentPool) validateConfigMapDestinations(ctx context.Context, obj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (pool *ManagedClustersAgentPool) validateOwnerReference(ctx context.Context, obj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (pool *ManagedClustersAgentPool) validateResourceReferences(ctx context.Context, obj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (pool *ManagedClustersAgentPool) validateSecretDestinations(ctx context.Context, obj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (pool *ManagedClustersAgentPool) validateWriteOnceProperties(ctx context.Context, oldObj *v20230201.ManagedClustersAgentPool, newObj *v20230201.ManagedClustersAgentPool) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
