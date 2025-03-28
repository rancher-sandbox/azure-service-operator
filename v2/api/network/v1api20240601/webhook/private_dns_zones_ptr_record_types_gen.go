// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20240601 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240601"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type PrivateDnsZonesPTRRecord struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240601-privatednszonesptrrecord,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=privatednszonesptrrecords,verbs=create;update,versions=v1api20240601,name=default.v1api20240601.privatednszonesptrrecords.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &PrivateDnsZonesPTRRecord{}

// Default applies defaults to the PrivateDnsZonesPTRRecord resource
func (record *PrivateDnsZonesPTRRecord) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240601.PrivateDnsZonesPTRRecord)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/PrivateDnsZonesPTRRecord, but got %T", obj)
	}
	err := record.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = record
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (record *PrivateDnsZonesPTRRecord) defaultAzureName(ctx context.Context, obj *v20240601.PrivateDnsZonesPTRRecord) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the PrivateDnsZonesPTRRecord resource
func (record *PrivateDnsZonesPTRRecord) defaultImpl(ctx context.Context, obj *v20240601.PrivateDnsZonesPTRRecord) error {
	err := record.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240601-privatednszonesptrrecord,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=privatednszonesptrrecords,verbs=create;update,versions=v1api20240601,name=validate.v1api20240601.privatednszonesptrrecords.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &PrivateDnsZonesPTRRecord{}

// ValidateCreate validates the creation of the resource
func (record *PrivateDnsZonesPTRRecord) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240601.PrivateDnsZonesPTRRecord)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/PrivateDnsZonesPTRRecord, but got %T", resource)
	}
	validations := record.createValidations()
	var temp any = record
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240601.PrivateDnsZonesPTRRecord]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (record *PrivateDnsZonesPTRRecord) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240601.PrivateDnsZonesPTRRecord)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/PrivateDnsZonesPTRRecord, but got %T", resource)
	}
	validations := record.deleteValidations()
	var temp any = record
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240601.PrivateDnsZonesPTRRecord]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (record *PrivateDnsZonesPTRRecord) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240601.PrivateDnsZonesPTRRecord)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/PrivateDnsZonesPTRRecord, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240601.PrivateDnsZonesPTRRecord)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/PrivateDnsZonesPTRRecord, but got %T", oldResource)
	}
	validations := record.updateValidations()
	var temp any = record
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240601.PrivateDnsZonesPTRRecord]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (record *PrivateDnsZonesPTRRecord) createValidations() []func(ctx context.Context, obj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error){record.validateResourceReferences, record.validateOwnerReference, record.validateSecretDestinations, record.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (record *PrivateDnsZonesPTRRecord) deleteValidations() []func(ctx context.Context, obj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (record *PrivateDnsZonesPTRRecord) updateValidations() []func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesPTRRecord, newObj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesPTRRecord, newObj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesPTRRecord, newObj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
			return record.validateResourceReferences(ctx, newObj)
		},
		record.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesPTRRecord, newObj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
			return record.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesPTRRecord, newObj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
			return record.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesPTRRecord, newObj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
			return record.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (record *PrivateDnsZonesPTRRecord) validateConfigMapDestinations(ctx context.Context, obj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (record *PrivateDnsZonesPTRRecord) validateOwnerReference(ctx context.Context, obj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (record *PrivateDnsZonesPTRRecord) validateResourceReferences(ctx context.Context, obj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (record *PrivateDnsZonesPTRRecord) validateSecretDestinations(ctx context.Context, obj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (record *PrivateDnsZonesPTRRecord) validateWriteOnceProperties(ctx context.Context, oldObj *v20240601.PrivateDnsZonesPTRRecord, newObj *v20240601.PrivateDnsZonesPTRRecord) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
