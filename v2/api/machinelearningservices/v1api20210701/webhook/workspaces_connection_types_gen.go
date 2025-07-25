// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20210701 "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20210701"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type WorkspacesConnection struct {
}

// +kubebuilder:webhook:path=/mutate-machinelearningservices-azure-com-v1api20210701-workspacesconnection,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=machinelearningservices.azure.com,resources=workspacesconnections,verbs=create;update,versions=v1api20210701,name=default.v1api20210701.workspacesconnections.machinelearningservices.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &WorkspacesConnection{}

// Default applies defaults to the WorkspacesConnection resource
func (connection *WorkspacesConnection) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20210701.WorkspacesConnection)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20210701/WorkspacesConnection, but got %T", obj)
	}
	err := connection.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = connection
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (connection *WorkspacesConnection) defaultAzureName(ctx context.Context, obj *v20210701.WorkspacesConnection) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the WorkspacesConnection resource
func (connection *WorkspacesConnection) defaultImpl(ctx context.Context, obj *v20210701.WorkspacesConnection) error {
	err := connection.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-machinelearningservices-azure-com-v1api20210701-workspacesconnection,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=machinelearningservices.azure.com,resources=workspacesconnections,verbs=create;update,versions=v1api20210701,name=validate.v1api20210701.workspacesconnections.machinelearningservices.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &WorkspacesConnection{}

// ValidateCreate validates the creation of the resource
func (connection *WorkspacesConnection) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210701.WorkspacesConnection)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20210701/WorkspacesConnection, but got %T", resource)
	}
	validations := connection.createValidations()
	var temp any = connection
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210701.WorkspacesConnection]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (connection *WorkspacesConnection) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210701.WorkspacesConnection)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20210701/WorkspacesConnection, but got %T", resource)
	}
	validations := connection.deleteValidations()
	var temp any = connection
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210701.WorkspacesConnection]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (connection *WorkspacesConnection) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20210701.WorkspacesConnection)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20210701/WorkspacesConnection, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20210701.WorkspacesConnection)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1api20210701/WorkspacesConnection, but got %T", oldResource)
	}
	validations := connection.updateValidations()
	var temp any = connection
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210701.WorkspacesConnection]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (connection *WorkspacesConnection) createValidations() []func(ctx context.Context, obj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20210701.WorkspacesConnection) (admission.Warnings, error){
		connection.validateResourceReferences,
		connection.validateOwnerReference,
		connection.validateSecretDestinations,
		connection.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (connection *WorkspacesConnection) deleteValidations() []func(ctx context.Context, obj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (connection *WorkspacesConnection) updateValidations() []func(ctx context.Context, oldObj *v20210701.WorkspacesConnection, newObj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20210701.WorkspacesConnection, newObj *v20210701.WorkspacesConnection) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20210701.WorkspacesConnection, newObj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
			return connection.validateResourceReferences(ctx, newObj)
		},
		connection.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20210701.WorkspacesConnection, newObj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
			return connection.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210701.WorkspacesConnection, newObj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
			return connection.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210701.WorkspacesConnection, newObj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
			return connection.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (connection *WorkspacesConnection) validateConfigMapDestinations(ctx context.Context, obj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (connection *WorkspacesConnection) validateOwnerReference(ctx context.Context, obj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (connection *WorkspacesConnection) validateResourceReferences(ctx context.Context, obj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (connection *WorkspacesConnection) validateSecretDestinations(ctx context.Context, obj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (connection *WorkspacesConnection) validateWriteOnceProperties(ctx context.Context, oldObj *v20210701.WorkspacesConnection, newObj *v20210701.WorkspacesConnection) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
