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

type SqlRoleAssignment struct {
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1api20240815-sqlroleassignment,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqlroleassignments,verbs=create;update,versions=v1api20240815,name=default.v1api20240815.sqlroleassignments.documentdb.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &SqlRoleAssignment{}

// Default applies defaults to the SqlRoleAssignment resource
func (assignment *SqlRoleAssignment) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20240815.SqlRoleAssignment)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlRoleAssignment, but got %T", obj)
	}
	err := assignment.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = assignment
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultImpl applies the code generated defaults to the SqlRoleAssignment resource
func (assignment *SqlRoleAssignment) defaultImpl(ctx context.Context, obj *v20240815.SqlRoleAssignment) error {
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1api20240815-sqlroleassignment,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqlroleassignments,verbs=create;update,versions=v1api20240815,name=validate.v1api20240815.sqlroleassignments.documentdb.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &SqlRoleAssignment{}

// ValidateCreate validates the creation of the resource
func (assignment *SqlRoleAssignment) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240815.SqlRoleAssignment)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlRoleAssignment, but got %T", resource)
	}
	validations := assignment.createValidations()
	var temp any = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240815.SqlRoleAssignment]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (assignment *SqlRoleAssignment) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20240815.SqlRoleAssignment)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlRoleAssignment, but got %T", resource)
	}
	validations := assignment.deleteValidations()
	var temp any = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240815.SqlRoleAssignment]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (assignment *SqlRoleAssignment) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20240815.SqlRoleAssignment)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlRoleAssignment, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20240815.SqlRoleAssignment)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlRoleAssignment, but got %T", oldResource)
	}
	validations := assignment.updateValidations()
	var temp any = assignment
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20240815.SqlRoleAssignment]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (assignment *SqlRoleAssignment) createValidations() []func(ctx context.Context, obj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20240815.SqlRoleAssignment) (admission.Warnings, error){assignment.validateResourceReferences, assignment.validateOwnerReference, assignment.validateSecretDestinations, assignment.validateConfigMapDestinations, assignment.validateOptionalConfigMapReferences}
}

// deleteValidations validates the deletion of the resource
func (assignment *SqlRoleAssignment) deleteValidations() []func(ctx context.Context, obj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (assignment *SqlRoleAssignment) updateValidations() []func(ctx context.Context, oldObj *v20240815.SqlRoleAssignment, newObj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20240815.SqlRoleAssignment, newObj *v20240815.SqlRoleAssignment) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20240815.SqlRoleAssignment, newObj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
			return assignment.validateResourceReferences(ctx, newObj)
		},
		assignment.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20240815.SqlRoleAssignment, newObj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
			return assignment.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240815.SqlRoleAssignment, newObj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
			return assignment.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240815.SqlRoleAssignment, newObj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
			return assignment.validateConfigMapDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20240815.SqlRoleAssignment, newObj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
			return assignment.validateOptionalConfigMapReferences(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (assignment *SqlRoleAssignment) validateConfigMapDestinations(ctx context.Context, obj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOptionalConfigMapReferences validates all optional configmap reference pairs to ensure that at most 1 is set
func (assignment *SqlRoleAssignment) validateOptionalConfigMapReferences(ctx context.Context, obj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return configmaps.ValidateOptionalReferences(refs)
}

// validateOwnerReference validates the owner field
func (assignment *SqlRoleAssignment) validateOwnerReference(ctx context.Context, obj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (assignment *SqlRoleAssignment) validateResourceReferences(ctx context.Context, obj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (assignment *SqlRoleAssignment) validateSecretDestinations(ctx context.Context, obj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (assignment *SqlRoleAssignment) validateWriteOnceProperties(ctx context.Context, oldObj *v20240815.SqlRoleAssignment, newObj *v20240815.SqlRoleAssignment) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
