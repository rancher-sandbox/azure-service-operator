//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20220131preview

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedIdentityCredential) DeepCopyInto(out *FederatedIdentityCredential) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedIdentityCredential.
func (in *FederatedIdentityCredential) DeepCopy() *FederatedIdentityCredential {
	if in == nil {
		return nil
	}
	out := new(FederatedIdentityCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedIdentityCredential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedIdentityCredentialList) DeepCopyInto(out *FederatedIdentityCredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedIdentityCredential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedIdentityCredentialList.
func (in *FederatedIdentityCredentialList) DeepCopy() *FederatedIdentityCredentialList {
	if in == nil {
		return nil
	}
	out := new(FederatedIdentityCredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedIdentityCredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedIdentityCredentialOperatorSpec) DeepCopyInto(out *FederatedIdentityCredentialOperatorSpec) {
	*out = *in
	if in.ConfigMapExpressions != nil {
		in, out := &in.ConfigMapExpressions, &out.ConfigMapExpressions
		*out = make([]*core.DestinationExpression, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(core.DestinationExpression)
				**out = **in
			}
		}
	}
	if in.SecretExpressions != nil {
		in, out := &in.SecretExpressions, &out.SecretExpressions
		*out = make([]*core.DestinationExpression, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(core.DestinationExpression)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedIdentityCredentialOperatorSpec.
func (in *FederatedIdentityCredentialOperatorSpec) DeepCopy() *FederatedIdentityCredentialOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedIdentityCredentialOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedIdentityCredential_STATUS) DeepCopyInto(out *FederatedIdentityCredential_STATUS) {
	*out = *in
	if in.Audiences != nil {
		in, out := &in.Audiences, &out.Audiences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedIdentityCredential_STATUS.
func (in *FederatedIdentityCredential_STATUS) DeepCopy() *FederatedIdentityCredential_STATUS {
	if in == nil {
		return nil
	}
	out := new(FederatedIdentityCredential_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedIdentityCredential_Spec) DeepCopyInto(out *FederatedIdentityCredential_Spec) {
	*out = *in
	if in.Audiences != nil {
		in, out := &in.Audiences, &out.Audiences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
	if in.IssuerFromConfig != nil {
		in, out := &in.IssuerFromConfig, &out.IssuerFromConfig
		*out = new(genruntime.ConfigMapReference)
		**out = **in
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(FederatedIdentityCredentialOperatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(string)
		**out = **in
	}
	if in.SubjectFromConfig != nil {
		in, out := &in.SubjectFromConfig, &out.SubjectFromConfig
		*out = new(genruntime.ConfigMapReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedIdentityCredential_Spec.
func (in *FederatedIdentityCredential_Spec) DeepCopy() *FederatedIdentityCredential_Spec {
	if in == nil {
		return nil
	}
	out := new(FederatedIdentityCredential_Spec)
	in.DeepCopyInto(out)
	return out
}
