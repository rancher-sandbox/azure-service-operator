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

type ServersDatabasesAdvancedThreatProtectionSetting struct {
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversdatabasesadvancedthreatprotectionsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasesadvancedthreatprotectionsettings,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversdatabasesadvancedthreatprotectionsettings.sql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &ServersDatabasesAdvancedThreatProtectionSetting{}

// Default applies defaults to the ServersDatabasesAdvancedThreatProtectionSetting resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20211101.ServersDatabasesAdvancedThreatProtectionSetting)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabasesAdvancedThreatProtectionSetting, but got %T", obj)
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

// defaultImpl applies the code generated defaults to the ServersDatabasesAdvancedThreatProtectionSetting resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) defaultImpl(ctx context.Context, obj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) error {
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversdatabasesadvancedthreatprotectionsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasesadvancedthreatprotectionsettings,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversdatabasesadvancedthreatprotectionsettings.sql.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &ServersDatabasesAdvancedThreatProtectionSetting{}

// ValidateCreate validates the creation of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.ServersDatabasesAdvancedThreatProtectionSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabasesAdvancedThreatProtectionSetting, but got %T", resource)
	}
	validations := setting.createValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersDatabasesAdvancedThreatProtectionSetting]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20211101.ServersDatabasesAdvancedThreatProtectionSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabasesAdvancedThreatProtectionSetting, but got %T", resource)
	}
	validations := setting.deleteValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersDatabasesAdvancedThreatProtectionSetting]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20211101.ServersDatabasesAdvancedThreatProtectionSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabasesAdvancedThreatProtectionSetting, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20211101.ServersDatabasesAdvancedThreatProtectionSetting)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/ServersDatabasesAdvancedThreatProtectionSetting, but got %T", oldResource)
	}
	validations := setting.updateValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20211101.ServersDatabasesAdvancedThreatProtectionSetting]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) createValidations() []func(ctx context.Context, obj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error){setting.validateResourceReferences, setting.validateOwnerReference, setting.validateSecretDestinations, setting.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) deleteValidations() []func(ctx context.Context, obj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) updateValidations() []func(ctx context.Context, oldObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting, newObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting, newObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting, newObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
			return setting.validateResourceReferences(ctx, newObj)
		},
		setting.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting, newObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
			return setting.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting, newObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
			return setting.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting, newObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
			return setting.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) validateConfigMapDestinations(ctx context.Context, obj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) validateOwnerReference(ctx context.Context, obj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) validateResourceReferences(ctx context.Context, obj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) validateSecretDestinations(ctx context.Context, obj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (setting *ServersDatabasesAdvancedThreatProtectionSetting) validateWriteOnceProperties(ctx context.Context, oldObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting, newObj *v20211101.ServersDatabasesAdvancedThreatProtectionSetting) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
