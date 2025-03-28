// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20211101 "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type ServersDatabase struct {
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversdatabase,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabases,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversdatabases.sql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &ServersDatabase{}

// Default applies defaults to the ServersDatabase resource
func (database *ServersDatabase) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20211101.ServersDatabase)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabase, but got %T", obj)
	}
	err := database.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = database
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (database *ServersDatabase) defaultAzureName(ctx context.Context, obj *v20211101.ServersDatabase) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the ServersDatabase resource
func (database *ServersDatabase) defaultImpl(ctx context.Context, obj *v20211101.ServersDatabase) error {
	err := database.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversdatabase,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabases,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversdatabases.sql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &ServersDatabase{}

// ValidateCreate validates the creation of the resource
func (database *ServersDatabase) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.ServersDatabase)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabase, but got %T", resource)
	}
	validations := database.createValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersDatabase]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (database *ServersDatabase) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.ServersDatabase)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabase, but got %T", resource)
	}
	validations := database.deleteValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersDatabase]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (database *ServersDatabase) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20211101.ServersDatabase)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabase, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20211101.ServersDatabase)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabase, but got %T", oldResource)
	}
	validations := database.updateValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersDatabase]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (database *ServersDatabase) createValidations() []func(ctx context.Context, obj *v20211101.ServersDatabase) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20211101.ServersDatabase) (admission.Warnings, error){database.validateResourceReferences, database.validateOwnerReference, database.validateSecretDestinations, database.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (database *ServersDatabase) deleteValidations() []func(ctx context.Context, obj *v20211101.ServersDatabase) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (database *ServersDatabase) updateValidations() []func(ctx context.Context, oldObj *v20211101.ServersDatabase, newObj *v20211101.ServersDatabase) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20211101.ServersDatabase, newObj *v20211101.ServersDatabase) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20211101.ServersDatabase, newObj *v20211101.ServersDatabase) (admission.Warnings, error) {
			return database.validateResourceReferences(ctx, newObj)
		},
		database.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20211101.ServersDatabase, newObj *v20211101.ServersDatabase) (admission.Warnings, error) {
			return database.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.ServersDatabase, newObj *v20211101.ServersDatabase) (admission.Warnings, error) {
			return database.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.ServersDatabase, newObj *v20211101.ServersDatabase) (admission.Warnings, error) {
			return database.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (database *ServersDatabase) validateConfigMapDestinations(ctx context.Context, obj *v20211101.ServersDatabase) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (database *ServersDatabase) validateOwnerReference(ctx context.Context, obj *v20211101.ServersDatabase) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (database *ServersDatabase) validateResourceReferences(ctx context.Context, obj *v20211101.ServersDatabase) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (database *ServersDatabase) validateSecretDestinations(ctx context.Context, obj *v20211101.ServersDatabase) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (database *ServersDatabase) validateWriteOnceProperties(ctx context.Context, oldObj *v20211101.ServersDatabase, newObj *v20211101.ServersDatabase) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
