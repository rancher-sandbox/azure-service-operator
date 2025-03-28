// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20210601 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type FlexibleServersDatabase struct {
}

// +kubebuilder:webhook:path=/mutate-dbforpostgresql-azure-com-v1api20210601-flexibleserversdatabase,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversdatabases,verbs=create;update,versions=v1api20210601,name=default.v1api20210601.flexibleserversdatabases.dbforpostgresql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &FlexibleServersDatabase{}

// Default applies defaults to the FlexibleServersDatabase resource
func (database *FlexibleServersDatabase) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20210601.FlexibleServersDatabase)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601/FlexibleServersDatabase, but got %T", obj)
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
func (database *FlexibleServersDatabase) defaultAzureName(ctx context.Context, obj *v20210601.FlexibleServersDatabase) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the FlexibleServersDatabase resource
func (database *FlexibleServersDatabase) defaultImpl(ctx context.Context, obj *v20210601.FlexibleServersDatabase) error {
	err := database.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-dbforpostgresql-azure-com-v1api20210601-flexibleserversdatabase,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversdatabases,verbs=create;update,versions=v1api20210601,name=validate.v1api20210601.flexibleserversdatabases.dbforpostgresql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &FlexibleServersDatabase{}

// ValidateCreate validates the creation of the resource
func (database *FlexibleServersDatabase) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210601.FlexibleServersDatabase)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601/FlexibleServersDatabase, but got %T", resource)
	}
	validations := database.createValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210601.FlexibleServersDatabase]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (database *FlexibleServersDatabase) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20210601.FlexibleServersDatabase)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601/FlexibleServersDatabase, but got %T", resource)
	}
	validations := database.deleteValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210601.FlexibleServersDatabase]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (database *FlexibleServersDatabase) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20210601.FlexibleServersDatabase)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601/FlexibleServersDatabase, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20210601.FlexibleServersDatabase)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601/FlexibleServersDatabase, but got %T", oldResource)
	}
	validations := database.updateValidations()
	var temp any = database
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20210601.FlexibleServersDatabase]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (database *FlexibleServersDatabase) createValidations() []func(ctx context.Context, obj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20210601.FlexibleServersDatabase) (admission.Warnings, error){database.validateResourceReferences, database.validateOwnerReference, database.validateSecretDestinations, database.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (database *FlexibleServersDatabase) deleteValidations() []func(ctx context.Context, obj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (database *FlexibleServersDatabase) updateValidations() []func(ctx context.Context, oldObj *v20210601.FlexibleServersDatabase, newObj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20210601.FlexibleServersDatabase, newObj *v20210601.FlexibleServersDatabase) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20210601.FlexibleServersDatabase, newObj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
			return database.validateResourceReferences(ctx, newObj)
		},
		database.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20210601.FlexibleServersDatabase, newObj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
			return database.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210601.FlexibleServersDatabase, newObj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
			return database.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20210601.FlexibleServersDatabase, newObj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
			return database.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (database *FlexibleServersDatabase) validateConfigMapDestinations(ctx context.Context, obj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (database *FlexibleServersDatabase) validateOwnerReference(ctx context.Context, obj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (database *FlexibleServersDatabase) validateResourceReferences(ctx context.Context, obj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (database *FlexibleServersDatabase) validateSecretDestinations(ctx context.Context, obj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (database *FlexibleServersDatabase) validateWriteOnceProperties(ctx context.Context, oldObj *v20210601.FlexibleServersDatabase, newObj *v20210601.FlexibleServersDatabase) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
