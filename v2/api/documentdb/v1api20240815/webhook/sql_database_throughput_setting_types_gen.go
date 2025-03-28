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

type SqlDatabaseThroughputSetting struct {
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1api20240815-sqldatabasethroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasethroughputsettings,verbs=create;update,versions=v1api20240815,name=default.v1api20240815.sqldatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &SqlDatabaseThroughputSetting{}

// Default applies defaults to the SqlDatabaseThroughputSetting resource
func (setting *SqlDatabaseThroughputSetting) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240815.SqlDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlDatabaseThroughputSetting, but got %T", obj)
	}
	err := setting.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = setting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultImpl applies the code generated defaults to the SqlDatabaseThroughputSetting resource
func (setting *SqlDatabaseThroughputSetting) defaultImpl(ctx context.Context, obj *v20240815.SqlDatabaseThroughputSetting) error {
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1api20240815-sqldatabasethroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasethroughputsettings,verbs=create;update,versions=v1api20240815,name=validate.v1api20240815.sqldatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &SqlDatabaseThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (setting *SqlDatabaseThroughputSetting) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240815.SqlDatabaseThroughputSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlDatabaseThroughputSetting, but got %T", resource)
	}
	validations := setting.createValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240815.SqlDatabaseThroughputSetting]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (setting *SqlDatabaseThroughputSetting) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240815.SqlDatabaseThroughputSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlDatabaseThroughputSetting, but got %T", resource)
	}
	validations := setting.deleteValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240815.SqlDatabaseThroughputSetting]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (setting *SqlDatabaseThroughputSetting) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240815.SqlDatabaseThroughputSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlDatabaseThroughputSetting, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240815.SqlDatabaseThroughputSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlDatabaseThroughputSetting, but got %T", oldResource)
	}
	validations := setting.updateValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240815.SqlDatabaseThroughputSetting]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (setting *SqlDatabaseThroughputSetting) createValidations() []func(ctx context.Context, obj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error){setting.validateResourceReferences, setting.validateOwnerReference, setting.validateSecretDestinations, setting.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (setting *SqlDatabaseThroughputSetting) deleteValidations() []func(ctx context.Context, obj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (setting *SqlDatabaseThroughputSetting) updateValidations() []func(ctx context.Context, oldObj *v20240815.SqlDatabaseThroughputSetting, newObj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240815.SqlDatabaseThroughputSetting, newObj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240815.SqlDatabaseThroughputSetting, newObj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
			return setting.validateResourceReferences(ctx, newObj)
		},
		setting.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240815.SqlDatabaseThroughputSetting, newObj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
			return setting.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240815.SqlDatabaseThroughputSetting, newObj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
			return setting.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240815.SqlDatabaseThroughputSetting, newObj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
			return setting.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (setting *SqlDatabaseThroughputSetting) validateConfigMapDestinations(ctx context.Context, obj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (setting *SqlDatabaseThroughputSetting) validateOwnerReference(ctx context.Context, obj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (setting *SqlDatabaseThroughputSetting) validateResourceReferences(ctx context.Context, obj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (setting *SqlDatabaseThroughputSetting) validateSecretDestinations(ctx context.Context, obj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (setting *SqlDatabaseThroughputSetting) validateWriteOnceProperties(ctx context.Context, oldObj *v20240815.SqlDatabaseThroughputSetting, newObj *v20240815.SqlDatabaseThroughputSetting) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
