// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20211101 "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type NamespacesQueue struct {
}

// +kubebuilder:webhook:path=/mutate-servicebus-azure-com-v1api20211101-namespacesqueue,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=servicebus.azure.com,resources=namespacesqueues,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.namespacesqueues.servicebus.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &NamespacesQueue{}

// Default applies defaults to the NamespacesQueue resource
func (queue *NamespacesQueue) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20211101.NamespacesQueue)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/NamespacesQueue, but got %T", obj)
	}
	err := queue.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = queue
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (queue *NamespacesQueue) defaultAzureName(ctx context.Context, obj *v20211101.NamespacesQueue) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the NamespacesQueue resource
func (queue *NamespacesQueue) defaultImpl(ctx context.Context, obj *v20211101.NamespacesQueue) error {
	err := queue.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-servicebus-azure-com-v1api20211101-namespacesqueue,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=servicebus.azure.com,resources=namespacesqueues,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.namespacesqueues.servicebus.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &NamespacesQueue{}

// ValidateCreate validates the creation of the resource
func (queue *NamespacesQueue) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.NamespacesQueue)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/NamespacesQueue, but got %T", resource)
	}
	validations := queue.createValidations()
	var temp any = queue
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.NamespacesQueue]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (queue *NamespacesQueue) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.NamespacesQueue)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/NamespacesQueue, but got %T", resource)
	}
	validations := queue.deleteValidations()
	var temp any = queue
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.NamespacesQueue]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (queue *NamespacesQueue) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20211101.NamespacesQueue)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/NamespacesQueue, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20211101.NamespacesQueue)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/NamespacesQueue, but got %T", oldResource)
	}
	validations := queue.updateValidations()
	var temp any = queue
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.NamespacesQueue]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (queue *NamespacesQueue) createValidations() []func(ctx context.Context, obj *v20211101.NamespacesQueue) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20211101.NamespacesQueue) (admission.Warnings, error){
		queue.validateResourceReferences,
		queue.validateOwnerReference,
		queue.validateSecretDestinations,
		queue.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (queue *NamespacesQueue) deleteValidations() []func(ctx context.Context, obj *v20211101.NamespacesQueue) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (queue *NamespacesQueue) updateValidations() []func(ctx context.Context, oldObj *v20211101.NamespacesQueue, newObj *v20211101.NamespacesQueue) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20211101.NamespacesQueue, newObj *v20211101.NamespacesQueue) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20211101.NamespacesQueue, newObj *v20211101.NamespacesQueue) (admission.Warnings, error) {
			return queue.validateResourceReferences(ctx, newObj)
		},
		queue.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20211101.NamespacesQueue, newObj *v20211101.NamespacesQueue) (admission.Warnings, error) {
			return queue.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.NamespacesQueue, newObj *v20211101.NamespacesQueue) (admission.Warnings, error) {
			return queue.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.NamespacesQueue, newObj *v20211101.NamespacesQueue) (admission.Warnings, error) {
			return queue.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (queue *NamespacesQueue) validateConfigMapDestinations(ctx context.Context, obj *v20211101.NamespacesQueue) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (queue *NamespacesQueue) validateOwnerReference(ctx context.Context, obj *v20211101.NamespacesQueue) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (queue *NamespacesQueue) validateResourceReferences(ctx context.Context, obj *v20211101.NamespacesQueue) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (queue *NamespacesQueue) validateSecretDestinations(ctx context.Context, obj *v20211101.NamespacesQueue) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (queue *NamespacesQueue) validateWriteOnceProperties(ctx context.Context, oldObj *v20211101.NamespacesQueue, newObj *v20211101.NamespacesQueue) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
