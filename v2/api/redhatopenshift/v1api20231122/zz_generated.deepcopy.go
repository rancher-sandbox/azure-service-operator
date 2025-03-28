//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20231122

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServerProfile) DeepCopyInto(out *APIServerProfile) {
	*out = *in
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(Visibility)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServerProfile.
func (in *APIServerProfile) DeepCopy() *APIServerProfile {
	if in == nil {
		return nil
	}
	out := new(APIServerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServerProfile_STATUS) DeepCopyInto(out *APIServerProfile_STATUS) {
	*out = *in
	if in.Ip != nil {
		in, out := &in.Ip, &out.Ip
		*out = new(string)
		**out = **in
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(Visibility_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServerProfile_STATUS.
func (in *APIServerProfile_STATUS) DeepCopy() *APIServerProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(APIServerProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProfile) DeepCopyInto(out *ClusterProfile) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.FipsValidatedModules != nil {
		in, out := &in.FipsValidatedModules, &out.FipsValidatedModules
		*out = new(FipsValidatedModules)
		**out = **in
	}
	if in.PullSecret != nil {
		in, out := &in.PullSecret, &out.PullSecret
		*out = new(genruntime.SecretReference)
		**out = **in
	}
	if in.ResourceGroupId != nil {
		in, out := &in.ResourceGroupId, &out.ResourceGroupId
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProfile.
func (in *ClusterProfile) DeepCopy() *ClusterProfile {
	if in == nil {
		return nil
	}
	out := new(ClusterProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProfile_STATUS) DeepCopyInto(out *ClusterProfile_STATUS) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.FipsValidatedModules != nil {
		in, out := &in.FipsValidatedModules, &out.FipsValidatedModules
		*out = new(FipsValidatedModules_STATUS)
		**out = **in
	}
	if in.ResourceGroupId != nil {
		in, out := &in.ResourceGroupId, &out.ResourceGroupId
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProfile_STATUS.
func (in *ClusterProfile_STATUS) DeepCopy() *ClusterProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(ClusterProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleProfile_STATUS) DeepCopyInto(out *ConsoleProfile_STATUS) {
	*out = *in
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleProfile_STATUS.
func (in *ConsoleProfile_STATUS) DeepCopy() *ConsoleProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(ConsoleProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EffectiveOutboundIP_STATUS) DeepCopyInto(out *EffectiveOutboundIP_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EffectiveOutboundIP_STATUS.
func (in *EffectiveOutboundIP_STATUS) DeepCopy() *EffectiveOutboundIP_STATUS {
	if in == nil {
		return nil
	}
	out := new(EffectiveOutboundIP_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressProfile) DeepCopyInto(out *IngressProfile) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(Visibility)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressProfile.
func (in *IngressProfile) DeepCopy() *IngressProfile {
	if in == nil {
		return nil
	}
	out := new(IngressProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressProfile_STATUS) DeepCopyInto(out *IngressProfile_STATUS) {
	*out = *in
	if in.Ip != nil {
		in, out := &in.Ip, &out.Ip
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(Visibility_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressProfile_STATUS.
func (in *IngressProfile_STATUS) DeepCopy() *IngressProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(IngressProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerProfile) DeepCopyInto(out *LoadBalancerProfile) {
	*out = *in
	if in.ManagedOutboundIps != nil {
		in, out := &in.ManagedOutboundIps, &out.ManagedOutboundIps
		*out = new(ManagedOutboundIPs)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerProfile.
func (in *LoadBalancerProfile) DeepCopy() *LoadBalancerProfile {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerProfile_STATUS) DeepCopyInto(out *LoadBalancerProfile_STATUS) {
	*out = *in
	if in.EffectiveOutboundIps != nil {
		in, out := &in.EffectiveOutboundIps, &out.EffectiveOutboundIps
		*out = make([]EffectiveOutboundIP_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagedOutboundIps != nil {
		in, out := &in.ManagedOutboundIps, &out.ManagedOutboundIps
		*out = new(ManagedOutboundIPs_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerProfile_STATUS.
func (in *LoadBalancerProfile_STATUS) DeepCopy() *LoadBalancerProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOutboundIPs) DeepCopyInto(out *ManagedOutboundIPs) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOutboundIPs.
func (in *ManagedOutboundIPs) DeepCopy() *ManagedOutboundIPs {
	if in == nil {
		return nil
	}
	out := new(ManagedOutboundIPs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOutboundIPs_STATUS) DeepCopyInto(out *ManagedOutboundIPs_STATUS) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOutboundIPs_STATUS.
func (in *ManagedOutboundIPs_STATUS) DeepCopy() *ManagedOutboundIPs_STATUS {
	if in == nil {
		return nil
	}
	out := new(ManagedOutboundIPs_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterProfile) DeepCopyInto(out *MasterProfile) {
	*out = *in
	if in.DiskEncryptionSetReference != nil {
		in, out := &in.DiskEncryptionSetReference, &out.DiskEncryptionSetReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.EncryptionAtHost != nil {
		in, out := &in.EncryptionAtHost, &out.EncryptionAtHost
		*out = new(EncryptionAtHost)
		**out = **in
	}
	if in.SubnetReference != nil {
		in, out := &in.SubnetReference, &out.SubnetReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.VmSize != nil {
		in, out := &in.VmSize, &out.VmSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterProfile.
func (in *MasterProfile) DeepCopy() *MasterProfile {
	if in == nil {
		return nil
	}
	out := new(MasterProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterProfile_STATUS) DeepCopyInto(out *MasterProfile_STATUS) {
	*out = *in
	if in.DiskEncryptionSetId != nil {
		in, out := &in.DiskEncryptionSetId, &out.DiskEncryptionSetId
		*out = new(string)
		**out = **in
	}
	if in.EncryptionAtHost != nil {
		in, out := &in.EncryptionAtHost, &out.EncryptionAtHost
		*out = new(EncryptionAtHost_STATUS)
		**out = **in
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
		*out = new(string)
		**out = **in
	}
	if in.VmSize != nil {
		in, out := &in.VmSize, &out.VmSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterProfile_STATUS.
func (in *MasterProfile_STATUS) DeepCopy() *MasterProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(MasterProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkProfile) DeepCopyInto(out *NetworkProfile) {
	*out = *in
	if in.LoadBalancerProfile != nil {
		in, out := &in.LoadBalancerProfile, &out.LoadBalancerProfile
		*out = new(LoadBalancerProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.OutboundType != nil {
		in, out := &in.OutboundType, &out.OutboundType
		*out = new(OutboundType)
		**out = **in
	}
	if in.PodCidr != nil {
		in, out := &in.PodCidr, &out.PodCidr
		*out = new(string)
		**out = **in
	}
	if in.PreconfiguredNSG != nil {
		in, out := &in.PreconfiguredNSG, &out.PreconfiguredNSG
		*out = new(PreconfiguredNSG)
		**out = **in
	}
	if in.ServiceCidr != nil {
		in, out := &in.ServiceCidr, &out.ServiceCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkProfile.
func (in *NetworkProfile) DeepCopy() *NetworkProfile {
	if in == nil {
		return nil
	}
	out := new(NetworkProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkProfile_STATUS) DeepCopyInto(out *NetworkProfile_STATUS) {
	*out = *in
	if in.LoadBalancerProfile != nil {
		in, out := &in.LoadBalancerProfile, &out.LoadBalancerProfile
		*out = new(LoadBalancerProfile_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.OutboundType != nil {
		in, out := &in.OutboundType, &out.OutboundType
		*out = new(OutboundType_STATUS)
		**out = **in
	}
	if in.PodCidr != nil {
		in, out := &in.PodCidr, &out.PodCidr
		*out = new(string)
		**out = **in
	}
	if in.PreconfiguredNSG != nil {
		in, out := &in.PreconfiguredNSG, &out.PreconfiguredNSG
		*out = new(PreconfiguredNSG_STATUS)
		**out = **in
	}
	if in.ServiceCidr != nil {
		in, out := &in.ServiceCidr, &out.ServiceCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkProfile_STATUS.
func (in *NetworkProfile_STATUS) DeepCopy() *NetworkProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(NetworkProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftCluster) DeepCopyInto(out *OpenShiftCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftCluster.
func (in *OpenShiftCluster) DeepCopy() *OpenShiftCluster {
	if in == nil {
		return nil
	}
	out := new(OpenShiftCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenShiftCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftClusterList) DeepCopyInto(out *OpenShiftClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenShiftCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftClusterList.
func (in *OpenShiftClusterList) DeepCopy() *OpenShiftClusterList {
	if in == nil {
		return nil
	}
	out := new(OpenShiftClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenShiftClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftClusterOperatorSpec) DeepCopyInto(out *OpenShiftClusterOperatorSpec) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftClusterOperatorSpec.
func (in *OpenShiftClusterOperatorSpec) DeepCopy() *OpenShiftClusterOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(OpenShiftClusterOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftCluster_STATUS) DeepCopyInto(out *OpenShiftCluster_STATUS) {
	*out = *in
	if in.ApiserverProfile != nil {
		in, out := &in.ApiserverProfile, &out.ApiserverProfile
		*out = new(APIServerProfile_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterProfile != nil {
		in, out := &in.ClusterProfile, &out.ClusterProfile
		*out = new(ClusterProfile_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConsoleProfile != nil {
		in, out := &in.ConsoleProfile, &out.ConsoleProfile
		*out = new(ConsoleProfile_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.IngressProfiles != nil {
		in, out := &in.IngressProfiles, &out.IngressProfiles
		*out = make([]IngressProfile_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MasterProfile != nil {
		in, out := &in.MasterProfile, &out.MasterProfile
		*out = new(MasterProfile_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkProfile != nil {
		in, out := &in.NetworkProfile, &out.NetworkProfile
		*out = new(NetworkProfile_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(ProvisioningState_STATUS)
		**out = **in
	}
	if in.ServicePrincipalProfile != nil {
		in, out := &in.ServicePrincipalProfile, &out.ServicePrincipalProfile
		*out = new(ServicePrincipalProfile_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.WorkerProfiles != nil {
		in, out := &in.WorkerProfiles, &out.WorkerProfiles
		*out = make([]WorkerProfile_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkerProfilesStatus != nil {
		in, out := &in.WorkerProfilesStatus, &out.WorkerProfilesStatus
		*out = make([]WorkerProfile_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftCluster_STATUS.
func (in *OpenShiftCluster_STATUS) DeepCopy() *OpenShiftCluster_STATUS {
	if in == nil {
		return nil
	}
	out := new(OpenShiftCluster_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftCluster_Spec) DeepCopyInto(out *OpenShiftCluster_Spec) {
	*out = *in
	if in.ApiserverProfile != nil {
		in, out := &in.ApiserverProfile, &out.ApiserverProfile
		*out = new(APIServerProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterProfile != nil {
		in, out := &in.ClusterProfile, &out.ClusterProfile
		*out = new(ClusterProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.IngressProfiles != nil {
		in, out := &in.IngressProfiles, &out.IngressProfiles
		*out = make([]IngressProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MasterProfile != nil {
		in, out := &in.MasterProfile, &out.MasterProfile
		*out = new(MasterProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkProfile != nil {
		in, out := &in.NetworkProfile, &out.NetworkProfile
		*out = new(NetworkProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(OpenShiftClusterOperatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.ServicePrincipalProfile != nil {
		in, out := &in.ServicePrincipalProfile, &out.ServicePrincipalProfile
		*out = new(ServicePrincipalProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.WorkerProfiles != nil {
		in, out := &in.WorkerProfiles, &out.WorkerProfiles
		*out = make([]WorkerProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftCluster_Spec.
func (in *OpenShiftCluster_Spec) DeepCopy() *OpenShiftCluster_Spec {
	if in == nil {
		return nil
	}
	out := new(OpenShiftCluster_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePrincipalProfile) DeepCopyInto(out *ServicePrincipalProfile) {
	*out = *in
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.ClientIdFromConfig != nil {
		in, out := &in.ClientIdFromConfig, &out.ClientIdFromConfig
		*out = new(genruntime.ConfigMapReference)
		**out = **in
	}
	if in.ClientSecret != nil {
		in, out := &in.ClientSecret, &out.ClientSecret
		*out = new(genruntime.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePrincipalProfile.
func (in *ServicePrincipalProfile) DeepCopy() *ServicePrincipalProfile {
	if in == nil {
		return nil
	}
	out := new(ServicePrincipalProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePrincipalProfile_STATUS) DeepCopyInto(out *ServicePrincipalProfile_STATUS) {
	*out = *in
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePrincipalProfile_STATUS.
func (in *ServicePrincipalProfile_STATUS) DeepCopy() *ServicePrincipalProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(ServicePrincipalProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_STATUS) DeepCopyInto(out *SystemData_STATUS) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedByType != nil {
		in, out := &in.CreatedByType, &out.CreatedByType
		*out = new(SystemData_CreatedByType_STATUS)
		**out = **in
	}
	if in.LastModifiedAt != nil {
		in, out := &in.LastModifiedAt, &out.LastModifiedAt
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedByType != nil {
		in, out := &in.LastModifiedByType, &out.LastModifiedByType
		*out = new(SystemData_LastModifiedByType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS.
func (in *SystemData_STATUS) DeepCopy() *SystemData_STATUS {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerProfile) DeepCopyInto(out *WorkerProfile) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
	if in.DiskEncryptionSetReference != nil {
		in, out := &in.DiskEncryptionSetReference, &out.DiskEncryptionSetReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.EncryptionAtHost != nil {
		in, out := &in.EncryptionAtHost, &out.EncryptionAtHost
		*out = new(EncryptionAtHost)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SubnetReference != nil {
		in, out := &in.SubnetReference, &out.SubnetReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.VmSize != nil {
		in, out := &in.VmSize, &out.VmSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerProfile.
func (in *WorkerProfile) DeepCopy() *WorkerProfile {
	if in == nil {
		return nil
	}
	out := new(WorkerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerProfile_STATUS) DeepCopyInto(out *WorkerProfile_STATUS) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int)
		**out = **in
	}
	if in.DiskEncryptionSetId != nil {
		in, out := &in.DiskEncryptionSetId, &out.DiskEncryptionSetId
		*out = new(string)
		**out = **in
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.EncryptionAtHost != nil {
		in, out := &in.EncryptionAtHost, &out.EncryptionAtHost
		*out = new(EncryptionAtHost_STATUS)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SubnetId != nil {
		in, out := &in.SubnetId, &out.SubnetId
		*out = new(string)
		**out = **in
	}
	if in.VmSize != nil {
		in, out := &in.VmSize, &out.VmSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerProfile_STATUS.
func (in *WorkerProfile_STATUS) DeepCopy() *WorkerProfile_STATUS {
	if in == nil {
		return nil
	}
	out := new(WorkerProfile_STATUS)
	in.DeepCopyInto(out)
	return out
}
