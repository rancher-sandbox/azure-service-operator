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

type ServersAuditingSetting struct {
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversauditingsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversauditingsettings,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversauditingsettings.sql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &ServersAuditingSetting{}

// Default applies defaults to the ServersAuditingSetting resource
func (setting *ServersAuditingSetting) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20211101.ServersAuditingSetting)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersAuditingSetting, but got %T", obj)
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

// defaultImpl applies the code generated defaults to the ServersAuditingSetting resource
func (setting *ServersAuditingSetting) defaultImpl(ctx context.Context, obj *v20211101.ServersAuditingSetting) error {
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversauditingsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversauditingsettings,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversauditingsettings.sql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &ServersAuditingSetting{}

// ValidateCreate validates the creation of the resource
func (setting *ServersAuditingSetting) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.ServersAuditingSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersAuditingSetting, but got %T", resource)
	}
	validations := setting.createValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersAuditingSetting]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (setting *ServersAuditingSetting) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.ServersAuditingSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersAuditingSetting, but got %T", resource)
	}
	validations := setting.deleteValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersAuditingSetting]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (setting *ServersAuditingSetting) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20211101.ServersAuditingSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersAuditingSetting, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20211101.ServersAuditingSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersAuditingSetting, but got %T", oldResource)
	}
	validations := setting.updateValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersAuditingSetting]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (setting *ServersAuditingSetting) createValidations() []func(ctx context.Context, obj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20211101.ServersAuditingSetting) (admission.Warnings, error){setting.validateResourceReferences, setting.validateOwnerReference, setting.validateSecretDestinations, setting.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (setting *ServersAuditingSetting) deleteValidations() []func(ctx context.Context, obj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (setting *ServersAuditingSetting) updateValidations() []func(ctx context.Context, oldObj *v20211101.ServersAuditingSetting, newObj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20211101.ServersAuditingSetting, newObj *v20211101.ServersAuditingSetting) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20211101.ServersAuditingSetting, newObj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
			return setting.validateResourceReferences(ctx, newObj)
		},
		setting.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20211101.ServersAuditingSetting, newObj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
			return setting.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.ServersAuditingSetting, newObj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
			return setting.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.ServersAuditingSetting, newObj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
			return setting.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (setting *ServersAuditingSetting) validateConfigMapDestinations(ctx context.Context, obj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (setting *ServersAuditingSetting) validateOwnerReference(ctx context.Context, obj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (setting *ServersAuditingSetting) validateResourceReferences(ctx context.Context, obj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (setting *ServersAuditingSetting) validateSecretDestinations(ctx context.Context, obj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (setting *ServersAuditingSetting) validateWriteOnceProperties(ctx context.Context, oldObj *v20211101.ServersAuditingSetting, newObj *v20211101.ServersAuditingSetting) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
