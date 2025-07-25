// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type Topic struct {
}

// +kubebuilder:webhook:path=/mutate-eventgrid-azure-com-v1api20200601-topic,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventgrid.azure.com,resources=topics,verbs=create;update,versions=v1api20200601,name=default.v1api20200601.topics.eventgrid.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &Topic{}

// Default applies defaults to the Topic resource
func (topic *Topic) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20200601.Topic)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601/Topic, but got %T", obj)
	}
	err := topic.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = topic
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (topic *Topic) defaultAzureName(ctx context.Context, obj *v20200601.Topic) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the Topic resource
func (topic *Topic) defaultImpl(ctx context.Context, obj *v20200601.Topic) error {
	err := topic.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-eventgrid-azure-com-v1api20200601-topic,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventgrid.azure.com,resources=topics,verbs=create;update,versions=v1api20200601,name=validate.v1api20200601.topics.eventgrid.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &Topic{}

// ValidateCreate validates the creation of the resource
func (topic *Topic) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20200601.Topic)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601/Topic, but got %T", resource)
	}
	validations := topic.createValidations()
	var temp any = topic
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20200601.Topic]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (topic *Topic) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20200601.Topic)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601/Topic, but got %T", resource)
	}
	validations := topic.deleteValidations()
	var temp any = topic
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20200601.Topic]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (topic *Topic) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20200601.Topic)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601/Topic, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20200601.Topic)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601/Topic, but got %T", oldResource)
	}
	validations := topic.updateValidations()
	var temp any = topic
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20200601.Topic]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (topic *Topic) createValidations() []func(ctx context.Context, obj *v20200601.Topic) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20200601.Topic) (admission.Warnings, error){
		topic.validateResourceReferences,
		topic.validateOwnerReference,
		topic.validateSecretDestinations,
		topic.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (topic *Topic) deleteValidations() []func(ctx context.Context, obj *v20200601.Topic) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (topic *Topic) updateValidations() []func(ctx context.Context, oldObj *v20200601.Topic, newObj *v20200601.Topic) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20200601.Topic, newObj *v20200601.Topic) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20200601.Topic, newObj *v20200601.Topic) (admission.Warnings, error) {
			return topic.validateResourceReferences(ctx, newObj)
		},
		topic.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20200601.Topic, newObj *v20200601.Topic) (admission.Warnings, error) {
			return topic.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20200601.Topic, newObj *v20200601.Topic) (admission.Warnings, error) {
			return topic.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20200601.Topic, newObj *v20200601.Topic) (admission.Warnings, error) {
			return topic.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (topic *Topic) validateConfigMapDestinations(ctx context.Context, obj *v20200601.Topic) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.ConfigMapDestination
	if obj.Spec.OperatorSpec.ConfigMaps != nil {
		toValidate = []*genruntime.ConfigMapDestination{
			obj.Spec.OperatorSpec.ConfigMaps.Endpoint,
		}
	}
	return configmaps.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (topic *Topic) validateOwnerReference(ctx context.Context, obj *v20200601.Topic) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (topic *Topic) validateResourceReferences(ctx context.Context, obj *v20200601.Topic) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (topic *Topic) validateSecretDestinations(ctx context.Context, obj *v20200601.Topic) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.SecretDestination
	if obj.Spec.OperatorSpec.Secrets != nil {
		toValidate = []*genruntime.SecretDestination{
			obj.Spec.OperatorSpec.Secrets.Key1,
			obj.Spec.OperatorSpec.Secrets.Key2,
		}
	}
	return secrets.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (topic *Topic) validateWriteOnceProperties(ctx context.Context, oldObj *v20200601.Topic, newObj *v20200601.Topic) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
