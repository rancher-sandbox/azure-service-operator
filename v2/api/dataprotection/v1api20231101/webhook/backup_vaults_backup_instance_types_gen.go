// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20231101 "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type BackupVaultsBackupInstance struct {
}

// +kubebuilder:webhook:path=/mutate-dataprotection-azure-com-v1api20231101-backupvaultsbackupinstance,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dataprotection.azure.com,resources=backupvaultsbackupinstances,verbs=create;update,versions=v1api20231101,name=default.v1api20231101.backupvaultsbackupinstances.dataprotection.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &BackupVaultsBackupInstance{}

// Default applies defaults to the BackupVaultsBackupInstance resource
func (instance *BackupVaultsBackupInstance) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20231101.BackupVaultsBackupInstance)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101/BackupVaultsBackupInstance, but got %T", obj)
	}
	err := instance.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = instance
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (instance *BackupVaultsBackupInstance) defaultAzureName(ctx context.Context, obj *v20231101.BackupVaultsBackupInstance) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the BackupVaultsBackupInstance resource
func (instance *BackupVaultsBackupInstance) defaultImpl(ctx context.Context, obj *v20231101.BackupVaultsBackupInstance) error {
	err := instance.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-dataprotection-azure-com-v1api20231101-backupvaultsbackupinstance,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dataprotection.azure.com,resources=backupvaultsbackupinstances,verbs=create;update,versions=v1api20231101,name=validate.v1api20231101.backupvaultsbackupinstances.dataprotection.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &BackupVaultsBackupInstance{}

// ValidateCreate validates the creation of the resource
func (instance *BackupVaultsBackupInstance) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20231101.BackupVaultsBackupInstance)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101/BackupVaultsBackupInstance, but got %T", resource)
	}
	validations := instance.createValidations()
	var temp any = instance
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20231101.BackupVaultsBackupInstance]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (instance *BackupVaultsBackupInstance) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20231101.BackupVaultsBackupInstance)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101/BackupVaultsBackupInstance, but got %T", resource)
	}
	validations := instance.deleteValidations()
	var temp any = instance
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20231101.BackupVaultsBackupInstance]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (instance *BackupVaultsBackupInstance) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20231101.BackupVaultsBackupInstance)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101/BackupVaultsBackupInstance, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20231101.BackupVaultsBackupInstance)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101/BackupVaultsBackupInstance, but got %T", oldResource)
	}
	validations := instance.updateValidations()
	var temp any = instance
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20231101.BackupVaultsBackupInstance]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (instance *BackupVaultsBackupInstance) createValidations() []func(ctx context.Context, obj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error){
		instance.validateResourceReferences,
		instance.validateOwnerReference,
		instance.validateSecretDestinations,
		instance.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (instance *BackupVaultsBackupInstance) deleteValidations() []func(ctx context.Context, obj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (instance *BackupVaultsBackupInstance) updateValidations() []func(ctx context.Context, oldObj *v20231101.BackupVaultsBackupInstance, newObj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20231101.BackupVaultsBackupInstance, newObj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20231101.BackupVaultsBackupInstance, newObj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
			return instance.validateResourceReferences(ctx, newObj)
		},
		instance.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20231101.BackupVaultsBackupInstance, newObj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
			return instance.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20231101.BackupVaultsBackupInstance, newObj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
			return instance.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20231101.BackupVaultsBackupInstance, newObj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
			return instance.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (instance *BackupVaultsBackupInstance) validateConfigMapDestinations(ctx context.Context, obj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (instance *BackupVaultsBackupInstance) validateOwnerReference(ctx context.Context, obj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (instance *BackupVaultsBackupInstance) validateResourceReferences(ctx context.Context, obj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (instance *BackupVaultsBackupInstance) validateSecretDestinations(ctx context.Context, obj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (instance *BackupVaultsBackupInstance) validateWriteOnceProperties(ctx context.Context, oldObj *v20231101.BackupVaultsBackupInstance, newObj *v20231101.BackupVaultsBackupInstance) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
