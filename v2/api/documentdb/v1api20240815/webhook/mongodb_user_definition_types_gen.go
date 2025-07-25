// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20240815 "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type MongodbUserDefinition struct {
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1api20240815-mongodbuserdefinition,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=mongodbuserdefinitions,verbs=create;update,versions=v1api20240815,name=default.v1api20240815.mongodbuserdefinitions.documentdb.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &MongodbUserDefinition{}

// Default applies defaults to the MongodbUserDefinition resource
func (definition *MongodbUserDefinition) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240815.MongodbUserDefinition)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/MongodbUserDefinition, but got %T", obj)
	}
	err := definition.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = definition
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (definition *MongodbUserDefinition) defaultAzureName(ctx context.Context, obj *v20240815.MongodbUserDefinition) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the MongodbUserDefinition resource
func (definition *MongodbUserDefinition) defaultImpl(ctx context.Context, obj *v20240815.MongodbUserDefinition) error {
	err := definition.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1api20240815-mongodbuserdefinition,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=mongodbuserdefinitions,verbs=create;update,versions=v1api20240815,name=validate.v1api20240815.mongodbuserdefinitions.documentdb.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &MongodbUserDefinition{}

// ValidateCreate validates the creation of the resource
func (definition *MongodbUserDefinition) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240815.MongodbUserDefinition)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/MongodbUserDefinition, but got %T", resource)
	}
	validations := definition.createValidations()
	var temp any = definition
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240815.MongodbUserDefinition]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (definition *MongodbUserDefinition) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240815.MongodbUserDefinition)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/MongodbUserDefinition, but got %T", resource)
	}
	validations := definition.deleteValidations()
	var temp any = definition
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240815.MongodbUserDefinition]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (definition *MongodbUserDefinition) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240815.MongodbUserDefinition)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/MongodbUserDefinition, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240815.MongodbUserDefinition)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/MongodbUserDefinition, but got %T", oldResource)
	}
	validations := definition.updateValidations()
	var temp any = definition
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240815.MongodbUserDefinition]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (definition *MongodbUserDefinition) createValidations() []func(ctx context.Context, obj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240815.MongodbUserDefinition) (admission.Warnings, error){
		definition.validateResourceReferences,
		definition.validateOwnerReference,
		definition.validateSecretDestinations,
		definition.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (definition *MongodbUserDefinition) deleteValidations() []func(ctx context.Context, obj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (definition *MongodbUserDefinition) updateValidations() []func(ctx context.Context, oldObj *v20240815.MongodbUserDefinition, newObj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240815.MongodbUserDefinition, newObj *v20240815.MongodbUserDefinition) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240815.MongodbUserDefinition, newObj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
			return definition.validateResourceReferences(ctx, newObj)
		},
		definition.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240815.MongodbUserDefinition, newObj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
			return definition.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240815.MongodbUserDefinition, newObj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
			return definition.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240815.MongodbUserDefinition, newObj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
			return definition.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (definition *MongodbUserDefinition) validateConfigMapDestinations(ctx context.Context, obj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (definition *MongodbUserDefinition) validateOwnerReference(ctx context.Context, obj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (definition *MongodbUserDefinition) validateResourceReferences(ctx context.Context, obj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (definition *MongodbUserDefinition) validateSecretDestinations(ctx context.Context, obj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (definition *MongodbUserDefinition) validateWriteOnceProperties(ctx context.Context, oldObj *v20240815.MongodbUserDefinition, newObj *v20240815.MongodbUserDefinition) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
