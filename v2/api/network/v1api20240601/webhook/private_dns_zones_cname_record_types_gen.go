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

type PrivateDnsZonesCNAMERecord struct {
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240601-privatednszonescnamerecord,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=privatednszonescnamerecords,verbs=create;update,versions=v1api20240601,name=default.v1api20240601.privatednszonescnamerecords.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &PrivateDnsZonesCNAMERecord{}

// Default applies defaults to the PrivateDnsZonesCNAMERecord resource
func (record *PrivateDnsZonesCNAMERecord) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240601.PrivateDnsZonesCNAMERecord)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/PrivateDnsZonesCNAMERecord, but got %T", obj)
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
func (record *PrivateDnsZonesCNAMERecord) defaultAzureName(ctx context.Context, obj *v20240601.PrivateDnsZonesCNAMERecord) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the PrivateDnsZonesCNAMERecord resource
func (record *PrivateDnsZonesCNAMERecord) defaultImpl(ctx context.Context, obj *v20240601.PrivateDnsZonesCNAMERecord) error {
	err := record.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240601-privatednszonescnamerecord,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=privatednszonescnamerecords,verbs=create;update,versions=v1api20240601,name=validate.v1api20240601.privatednszonescnamerecords.network.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &PrivateDnsZonesCNAMERecord{}

// ValidateCreate validates the creation of the resource
func (record *PrivateDnsZonesCNAMERecord) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240601.PrivateDnsZonesCNAMERecord)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/PrivateDnsZonesCNAMERecord, but got %T", resource)
	}
	validations := record.createValidations()
	var temp any = record
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240601.PrivateDnsZonesCNAMERecord]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (record *PrivateDnsZonesCNAMERecord) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240601.PrivateDnsZonesCNAMERecord)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/PrivateDnsZonesCNAMERecord, but got %T", resource)
	}
	validations := record.deleteValidations()
	var temp any = record
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240601.PrivateDnsZonesCNAMERecord]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (record *PrivateDnsZonesCNAMERecord) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240601.PrivateDnsZonesCNAMERecord)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/PrivateDnsZonesCNAMERecord, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240601.PrivateDnsZonesCNAMERecord)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/network/v1api20240601/PrivateDnsZonesCNAMERecord, but got %T", oldResource)
	}
	validations := record.updateValidations()
	var temp any = record
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240601.PrivateDnsZonesCNAMERecord]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (record *PrivateDnsZonesCNAMERecord) createValidations() []func(ctx context.Context, obj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error){
		record.validateResourceReferences,
		record.validateOwnerReference,
		record.validateSecretDestinations,
		record.validateConfigMapDestinations,
	}
}

// deleteValidations validates the deletion of the resource
func (record *PrivateDnsZonesCNAMERecord) deleteValidations() []func(ctx context.Context, obj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (record *PrivateDnsZonesCNAMERecord) updateValidations() []func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesCNAMERecord, newObj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesCNAMERecord, newObj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesCNAMERecord, newObj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
			return record.validateResourceReferences(ctx, newObj)
		},
		record.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesCNAMERecord, newObj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
			return record.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesCNAMERecord, newObj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
			return record.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240601.PrivateDnsZonesCNAMERecord, newObj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
			return record.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (record *PrivateDnsZonesCNAMERecord) validateConfigMapDestinations(ctx context.Context, obj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (record *PrivateDnsZonesCNAMERecord) validateOwnerReference(ctx context.Context, obj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (record *PrivateDnsZonesCNAMERecord) validateResourceReferences(ctx context.Context, obj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (record *PrivateDnsZonesCNAMERecord) validateSecretDestinations(ctx context.Context, obj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (record *PrivateDnsZonesCNAMERecord) validateWriteOnceProperties(ctx context.Context, oldObj *v20240601.PrivateDnsZonesCNAMERecord, newObj *v20240601.PrivateDnsZonesCNAMERecord) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
