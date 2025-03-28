// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package webhook

import (
	"context"
	"fmt"
	v20180601 "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1api20180601"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type Server struct {
}

// +kubebuilder:webhook:path=/mutate-dbformariadb-azure-com-v1api20180601-server,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformariadb.azure.com,resources=servers,verbs=create;update,versions=v1api20180601,name=default.v1api20180601.servers.dbformariadb.azure.com,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &Server{}

// Default applies defaults to the Server resource
func (server *Server) Default(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v20180601.Server)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1api20180601/Server, but got %T", obj)
	}
	err := server.defaultImpl(ctx, resource)
	if err != nil {
		return err
	}
	var temp any = server
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		err = runtimeDefaulter.CustomDefault(ctx, resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (server *Server) defaultAzureName(ctx context.Context, obj *v20180601.Server) error {
	if obj.Spec.AzureName == "" {
		obj.Spec.AzureName = obj.Name
	}
	return nil
}

// defaultImpl applies the code generated defaults to the Server resource
func (server *Server) defaultImpl(ctx context.Context, obj *v20180601.Server) error {
	err := server.defaultAzureName(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

// +kubebuilder:webhook:path=/validate-dbformariadb-azure-com-v1api20180601-server,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformariadb.azure.com,resources=servers,verbs=create;update,versions=v1api20180601,name=validate.v1api20180601.servers.dbformariadb.azure.com,admissionReviewVersions=v1

var _ webhook.CustomValidator = &Server{}

// ValidateCreate validates the creation of the resource
func (server *Server) ValidateCreate(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20180601.Server)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1api20180601/Server, but got %T", resource)
	}
	validations := server.createValidations()
	var temp any = server
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20180601.Server]); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(ctx, obj, validations)
}

// ValidateDelete validates the deletion of the resource
func (server *Server) ValidateDelete(ctx context.Context, resource runtime.Object) (admission.Warnings, error) {
	obj, ok := resource.(*v20180601.Server)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1api20180601/Server, but got %T", resource)
	}
	validations := server.deleteValidations()
	var temp any = server
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20180601.Server]); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(ctx, obj, validations)
}

// ValidateUpdate validates an update of the resource
func (server *Server) ValidateUpdate(ctx context.Context, oldResource runtime.Object, newResource runtime.Object) (admission.Warnings, error) {
	newObj, ok := newResource.(*v20180601.Server)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1api20180601/Server, but got %T", newResource)
	}
	oldObj, ok := oldResource.(*v20180601.Server)
	if !ok {
		return nil, fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1api20180601/Server, but got %T", oldResource)
	}
	validations := server.updateValidations()
	var temp any = server
	if runtimeValidator, ok := temp.(genruntime.Validator[*v20180601.Server]); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(
		ctx,
		oldObj,
		newObj,
		validations)
}

// createValidations validates the creation of the resource
func (server *Server) createValidations() []func(ctx context.Context, obj *v20180601.Server) (admission.Warnings, error) {
	return []func(ctx context.Context, obj *v20180601.Server) (admission.Warnings, error){server.validateResourceReferences, server.validateOwnerReference, server.validateSecretDestinations, server.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (server *Server) deleteValidations() []func(ctx context.Context, obj *v20180601.Server) (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (server *Server) updateValidations() []func(ctx context.Context, oldObj *v20180601.Server, newObj *v20180601.Server) (admission.Warnings, error) {
	return []func(ctx context.Context, oldObj *v20180601.Server, newObj *v20180601.Server) (admission.Warnings, error){
		func(ctx context.Context, oldObj *v20180601.Server, newObj *v20180601.Server) (admission.Warnings, error) {
			return server.validateResourceReferences(ctx, newObj)
		},
		server.validateWriteOnceProperties,
		func(ctx context.Context, oldObj *v20180601.Server, newObj *v20180601.Server) (admission.Warnings, error) {
			return server.validateOwnerReference(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20180601.Server, newObj *v20180601.Server) (admission.Warnings, error) {
			return server.validateSecretDestinations(ctx, newObj)
		},
		func(ctx context.Context, oldObj *v20180601.Server, newObj *v20180601.Server) (admission.Warnings, error) {
			return server.validateConfigMapDestinations(ctx, newObj)
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (server *Server) validateConfigMapDestinations(ctx context.Context, obj *v20180601.Server) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(obj, nil, obj.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (server *Server) validateOwnerReference(ctx context.Context, obj *v20180601.Server) (admission.Warnings, error) {
	return genruntime.ValidateOwner(obj)
}

// validateResourceReferences validates all resource references
func (server *Server) validateResourceReferences(ctx context.Context, obj *v20180601.Server) (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&obj.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (server *Server) validateSecretDestinations(ctx context.Context, obj *v20180601.Server) (admission.Warnings, error) {
	if obj.Spec.OperatorSpec == nil {
		return nil, nil
	}
	var toValidate []*genruntime.SecretDestination
	if obj.Spec.OperatorSpec.Secrets != nil {
		toValidate = []*genruntime.SecretDestination{
			obj.Spec.OperatorSpec.Secrets.FullyQualifiedDomainName,
		}
	}
	return secrets.ValidateDestinations(obj, toValidate, obj.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (server *Server) validateWriteOnceProperties(ctx context.Context, oldObj *v20180601.Server, newObj *v20180601.Server) (admission.Warnings, error) {
	return genruntime.ValidateWriteOnceProperties(oldObj, newObj)
}
