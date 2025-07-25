//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EffectiveReplicationInitParameters) DeepCopyInto(out *EffectiveReplicationInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EffectiveReplicationInitParameters.
func (in *EffectiveReplicationInitParameters) DeepCopy() *EffectiveReplicationInitParameters {
	if in == nil {
		return nil
	}
	out := new(EffectiveReplicationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EffectiveReplicationObservation) DeepCopyInto(out *EffectiveReplicationObservation) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]ReplicasObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EffectiveReplicationObservation.
func (in *EffectiveReplicationObservation) DeepCopy() *EffectiveReplicationObservation {
	if in == nil {
		return nil
	}
	out := new(EffectiveReplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EffectiveReplicationParameters) DeepCopyInto(out *EffectiveReplicationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EffectiveReplicationParameters.
func (in *EffectiveReplicationParameters) DeepCopy() *EffectiveReplicationParameters {
	if in == nil {
		return nil
	}
	out := new(EffectiveReplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSharesInitParameters) DeepCopyInto(out *FileSharesInitParameters) {
	*out = *in
	if in.CapacityGb != nil {
		in, out := &in.CapacityGb, &out.CapacityGb
		*out = new(float64)
		**out = **in
	}
	if in.NFSExportOptions != nil {
		in, out := &in.NFSExportOptions, &out.NFSExportOptions
		*out = make([]NFSExportOptionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SourceBackup != nil {
		in, out := &in.SourceBackup, &out.SourceBackup
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSharesInitParameters.
func (in *FileSharesInitParameters) DeepCopy() *FileSharesInitParameters {
	if in == nil {
		return nil
	}
	out := new(FileSharesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSharesObservation) DeepCopyInto(out *FileSharesObservation) {
	*out = *in
	if in.CapacityGb != nil {
		in, out := &in.CapacityGb, &out.CapacityGb
		*out = new(float64)
		**out = **in
	}
	if in.NFSExportOptions != nil {
		in, out := &in.NFSExportOptions, &out.NFSExportOptions
		*out = make([]NFSExportOptionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SourceBackup != nil {
		in, out := &in.SourceBackup, &out.SourceBackup
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSharesObservation.
func (in *FileSharesObservation) DeepCopy() *FileSharesObservation {
	if in == nil {
		return nil
	}
	out := new(FileSharesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSharesParameters) DeepCopyInto(out *FileSharesParameters) {
	*out = *in
	if in.CapacityGb != nil {
		in, out := &in.CapacityGb, &out.CapacityGb
		*out = new(float64)
		**out = **in
	}
	if in.NFSExportOptions != nil {
		in, out := &in.NFSExportOptions, &out.NFSExportOptions
		*out = make([]NFSExportOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SourceBackup != nil {
		in, out := &in.SourceBackup, &out.SourceBackup
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSharesParameters.
func (in *FileSharesParameters) DeepCopy() *FileSharesParameters {
	if in == nil {
		return nil
	}
	out := new(FileSharesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedIopsInitParameters) DeepCopyInto(out *FixedIopsInitParameters) {
	*out = *in
	if in.MaxIops != nil {
		in, out := &in.MaxIops, &out.MaxIops
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedIopsInitParameters.
func (in *FixedIopsInitParameters) DeepCopy() *FixedIopsInitParameters {
	if in == nil {
		return nil
	}
	out := new(FixedIopsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedIopsObservation) DeepCopyInto(out *FixedIopsObservation) {
	*out = *in
	if in.MaxIops != nil {
		in, out := &in.MaxIops, &out.MaxIops
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedIopsObservation.
func (in *FixedIopsObservation) DeepCopy() *FixedIopsObservation {
	if in == nil {
		return nil
	}
	out := new(FixedIopsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedIopsParameters) DeepCopyInto(out *FixedIopsParameters) {
	*out = *in
	if in.MaxIops != nil {
		in, out := &in.MaxIops, &out.MaxIops
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedIopsParameters.
func (in *FixedIopsParameters) DeepCopy() *FixedIopsParameters {
	if in == nil {
		return nil
	}
	out := new(FixedIopsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialReplicationInitParameters) DeepCopyInto(out *InitialReplicationInitParameters) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]InitialReplicationReplicasInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialReplicationInitParameters.
func (in *InitialReplicationInitParameters) DeepCopy() *InitialReplicationInitParameters {
	if in == nil {
		return nil
	}
	out := new(InitialReplicationInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialReplicationObservation) DeepCopyInto(out *InitialReplicationObservation) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]InitialReplicationReplicasObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialReplicationObservation.
func (in *InitialReplicationObservation) DeepCopy() *InitialReplicationObservation {
	if in == nil {
		return nil
	}
	out := new(InitialReplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialReplicationParameters) DeepCopyInto(out *InitialReplicationParameters) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]InitialReplicationReplicasParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialReplicationParameters.
func (in *InitialReplicationParameters) DeepCopy() *InitialReplicationParameters {
	if in == nil {
		return nil
	}
	out := new(InitialReplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialReplicationReplicasInitParameters) DeepCopyInto(out *InitialReplicationReplicasInitParameters) {
	*out = *in
	if in.PeerInstance != nil {
		in, out := &in.PeerInstance, &out.PeerInstance
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialReplicationReplicasInitParameters.
func (in *InitialReplicationReplicasInitParameters) DeepCopy() *InitialReplicationReplicasInitParameters {
	if in == nil {
		return nil
	}
	out := new(InitialReplicationReplicasInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialReplicationReplicasObservation) DeepCopyInto(out *InitialReplicationReplicasObservation) {
	*out = *in
	if in.PeerInstance != nil {
		in, out := &in.PeerInstance, &out.PeerInstance
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialReplicationReplicasObservation.
func (in *InitialReplicationReplicasObservation) DeepCopy() *InitialReplicationReplicasObservation {
	if in == nil {
		return nil
	}
	out := new(InitialReplicationReplicasObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialReplicationReplicasParameters) DeepCopyInto(out *InitialReplicationReplicasParameters) {
	*out = *in
	if in.PeerInstance != nil {
		in, out := &in.PeerInstance, &out.PeerInstance
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialReplicationReplicasParameters.
func (in *InitialReplicationReplicasParameters) DeepCopy() *InitialReplicationReplicasParameters {
	if in == nil {
		return nil
	}
	out := new(InitialReplicationReplicasParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceInitParameters) DeepCopyInto(out *InstanceInitParameters) {
	*out = *in
	if in.DeletionProtectionEnabled != nil {
		in, out := &in.DeletionProtectionEnabled, &out.DeletionProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DeletionProtectionReason != nil {
		in, out := &in.DeletionProtectionReason, &out.DeletionProtectionReason
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FileShares != nil {
		in, out := &in.FileShares, &out.FileShares
		*out = new(FileSharesInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.InitialReplication != nil {
		in, out := &in.InitialReplication, &out.InitialReplication
		*out = new(InitialReplicationInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyName != nil {
		in, out := &in.KMSKeyName, &out.KMSKeyName
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyNameRef != nil {
		in, out := &in.KMSKeyNameRef, &out.KMSKeyNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyNameSelector != nil {
		in, out := &in.KMSKeyNameSelector, &out.KMSKeyNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworksInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PerformanceConfig != nil {
		in, out := &in.PerformanceConfig, &out.PerformanceConfig
		*out = new(PerformanceConfigInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceInitParameters.
func (in *InstanceInitParameters) DeepCopy() *InstanceInitParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceObservation) DeepCopyInto(out *InstanceObservation) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.DeletionProtectionEnabled != nil {
		in, out := &in.DeletionProtectionEnabled, &out.DeletionProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DeletionProtectionReason != nil {
		in, out := &in.DeletionProtectionReason, &out.DeletionProtectionReason
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EffectiveLabels != nil {
		in, out := &in.EffectiveLabels, &out.EffectiveLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.EffectiveReplication != nil {
		in, out := &in.EffectiveReplication, &out.EffectiveReplication
		*out = make([]EffectiveReplicationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.FileShares != nil {
		in, out := &in.FileShares, &out.FileShares
		*out = new(FileSharesObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InitialReplication != nil {
		in, out := &in.InitialReplication, &out.InitialReplication
		*out = new(InitialReplicationObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyName != nil {
		in, out := &in.KMSKeyName, &out.KMSKeyName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworksObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PerformanceConfig != nil {
		in, out := &in.PerformanceConfig, &out.PerformanceConfig
		*out = new(PerformanceConfigObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TerraformLabels != nil {
		in, out := &in.TerraformLabels, &out.TerraformLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceObservation.
func (in *InstanceObservation) DeepCopy() *InstanceObservation {
	if in == nil {
		return nil
	}
	out := new(InstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceParameters) DeepCopyInto(out *InstanceParameters) {
	*out = *in
	if in.DeletionProtectionEnabled != nil {
		in, out := &in.DeletionProtectionEnabled, &out.DeletionProtectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DeletionProtectionReason != nil {
		in, out := &in.DeletionProtectionReason, &out.DeletionProtectionReason
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FileShares != nil {
		in, out := &in.FileShares, &out.FileShares
		*out = new(FileSharesParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.InitialReplication != nil {
		in, out := &in.InitialReplication, &out.InitialReplication
		*out = new(InitialReplicationParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyName != nil {
		in, out := &in.KMSKeyName, &out.KMSKeyName
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyNameRef != nil {
		in, out := &in.KMSKeyNameRef, &out.KMSKeyNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyNameSelector != nil {
		in, out := &in.KMSKeyNameSelector, &out.KMSKeyNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworksParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PerformanceConfig != nil {
		in, out := &in.PerformanceConfig, &out.PerformanceConfig
		*out = new(PerformanceConfigParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceParameters.
func (in *InstanceParameters) DeepCopy() *InstanceParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IopsPerTbInitParameters) DeepCopyInto(out *IopsPerTbInitParameters) {
	*out = *in
	if in.MaxIopsPerTb != nil {
		in, out := &in.MaxIopsPerTb, &out.MaxIopsPerTb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IopsPerTbInitParameters.
func (in *IopsPerTbInitParameters) DeepCopy() *IopsPerTbInitParameters {
	if in == nil {
		return nil
	}
	out := new(IopsPerTbInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IopsPerTbObservation) DeepCopyInto(out *IopsPerTbObservation) {
	*out = *in
	if in.MaxIopsPerTb != nil {
		in, out := &in.MaxIopsPerTb, &out.MaxIopsPerTb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IopsPerTbObservation.
func (in *IopsPerTbObservation) DeepCopy() *IopsPerTbObservation {
	if in == nil {
		return nil
	}
	out := new(IopsPerTbObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IopsPerTbParameters) DeepCopyInto(out *IopsPerTbParameters) {
	*out = *in
	if in.MaxIopsPerTb != nil {
		in, out := &in.MaxIopsPerTb, &out.MaxIopsPerTb
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IopsPerTbParameters.
func (in *IopsPerTbParameters) DeepCopy() *IopsPerTbParameters {
	if in == nil {
		return nil
	}
	out := new(IopsPerTbParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSExportOptionsInitParameters) DeepCopyInto(out *NFSExportOptionsInitParameters) {
	*out = *in
	if in.AccessMode != nil {
		in, out := &in.AccessMode, &out.AccessMode
		*out = new(string)
		**out = **in
	}
	if in.AnonGID != nil {
		in, out := &in.AnonGID, &out.AnonGID
		*out = new(float64)
		**out = **in
	}
	if in.AnonUID != nil {
		in, out := &in.AnonUID, &out.AnonUID
		*out = new(float64)
		**out = **in
	}
	if in.IPRanges != nil {
		in, out := &in.IPRanges, &out.IPRanges
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SquashMode != nil {
		in, out := &in.SquashMode, &out.SquashMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSExportOptionsInitParameters.
func (in *NFSExportOptionsInitParameters) DeepCopy() *NFSExportOptionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(NFSExportOptionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSExportOptionsObservation) DeepCopyInto(out *NFSExportOptionsObservation) {
	*out = *in
	if in.AccessMode != nil {
		in, out := &in.AccessMode, &out.AccessMode
		*out = new(string)
		**out = **in
	}
	if in.AnonGID != nil {
		in, out := &in.AnonGID, &out.AnonGID
		*out = new(float64)
		**out = **in
	}
	if in.AnonUID != nil {
		in, out := &in.AnonUID, &out.AnonUID
		*out = new(float64)
		**out = **in
	}
	if in.IPRanges != nil {
		in, out := &in.IPRanges, &out.IPRanges
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SquashMode != nil {
		in, out := &in.SquashMode, &out.SquashMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSExportOptionsObservation.
func (in *NFSExportOptionsObservation) DeepCopy() *NFSExportOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(NFSExportOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSExportOptionsParameters) DeepCopyInto(out *NFSExportOptionsParameters) {
	*out = *in
	if in.AccessMode != nil {
		in, out := &in.AccessMode, &out.AccessMode
		*out = new(string)
		**out = **in
	}
	if in.AnonGID != nil {
		in, out := &in.AnonGID, &out.AnonGID
		*out = new(float64)
		**out = **in
	}
	if in.AnonUID != nil {
		in, out := &in.AnonUID, &out.AnonUID
		*out = new(float64)
		**out = **in
	}
	if in.IPRanges != nil {
		in, out := &in.IPRanges, &out.IPRanges
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SquashMode != nil {
		in, out := &in.SquashMode, &out.SquashMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSExportOptionsParameters.
func (in *NFSExportOptionsParameters) DeepCopy() *NFSExportOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(NFSExportOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworksInitParameters) DeepCopyInto(out *NetworksInitParameters) {
	*out = *in
	if in.ConnectMode != nil {
		in, out := &in.ConnectMode, &out.ConnectMode
		*out = new(string)
		**out = **in
	}
	if in.Modes != nil {
		in, out := &in.Modes, &out.Modes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.ReservedIPRange != nil {
		in, out := &in.ReservedIPRange, &out.ReservedIPRange
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworksInitParameters.
func (in *NetworksInitParameters) DeepCopy() *NetworksInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworksInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworksObservation) DeepCopyInto(out *NetworksObservation) {
	*out = *in
	if in.ConnectMode != nil {
		in, out := &in.ConnectMode, &out.ConnectMode
		*out = new(string)
		**out = **in
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Modes != nil {
		in, out := &in.Modes, &out.Modes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.ReservedIPRange != nil {
		in, out := &in.ReservedIPRange, &out.ReservedIPRange
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworksObservation.
func (in *NetworksObservation) DeepCopy() *NetworksObservation {
	if in == nil {
		return nil
	}
	out := new(NetworksObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworksParameters) DeepCopyInto(out *NetworksParameters) {
	*out = *in
	if in.ConnectMode != nil {
		in, out := &in.ConnectMode, &out.ConnectMode
		*out = new(string)
		**out = **in
	}
	if in.Modes != nil {
		in, out := &in.Modes, &out.Modes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.ReservedIPRange != nil {
		in, out := &in.ReservedIPRange, &out.ReservedIPRange
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworksParameters.
func (in *NetworksParameters) DeepCopy() *NetworksParameters {
	if in == nil {
		return nil
	}
	out := new(NetworksParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerformanceConfigInitParameters) DeepCopyInto(out *PerformanceConfigInitParameters) {
	*out = *in
	if in.FixedIops != nil {
		in, out := &in.FixedIops, &out.FixedIops
		*out = new(FixedIopsInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.IopsPerTb != nil {
		in, out := &in.IopsPerTb, &out.IopsPerTb
		*out = new(IopsPerTbInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerformanceConfigInitParameters.
func (in *PerformanceConfigInitParameters) DeepCopy() *PerformanceConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(PerformanceConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerformanceConfigObservation) DeepCopyInto(out *PerformanceConfigObservation) {
	*out = *in
	if in.FixedIops != nil {
		in, out := &in.FixedIops, &out.FixedIops
		*out = new(FixedIopsObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.IopsPerTb != nil {
		in, out := &in.IopsPerTb, &out.IopsPerTb
		*out = new(IopsPerTbObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerformanceConfigObservation.
func (in *PerformanceConfigObservation) DeepCopy() *PerformanceConfigObservation {
	if in == nil {
		return nil
	}
	out := new(PerformanceConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerformanceConfigParameters) DeepCopyInto(out *PerformanceConfigParameters) {
	*out = *in
	if in.FixedIops != nil {
		in, out := &in.FixedIops, &out.FixedIops
		*out = new(FixedIopsParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.IopsPerTb != nil {
		in, out := &in.IopsPerTb, &out.IopsPerTb
		*out = new(IopsPerTbParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerformanceConfigParameters.
func (in *PerformanceConfigParameters) DeepCopy() *PerformanceConfigParameters {
	if in == nil {
		return nil
	}
	out := new(PerformanceConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicasInitParameters) DeepCopyInto(out *ReplicasInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicasInitParameters.
func (in *ReplicasInitParameters) DeepCopy() *ReplicasInitParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicasInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicasObservation) DeepCopyInto(out *ReplicasObservation) {
	*out = *in
	if in.LastActiveSyncTime != nil {
		in, out := &in.LastActiveSyncTime, &out.LastActiveSyncTime
		*out = new(string)
		**out = **in
	}
	if in.PeerInstance != nil {
		in, out := &in.PeerInstance, &out.PeerInstance
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateReasons != nil {
		in, out := &in.StateReasons, &out.StateReasons
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicasObservation.
func (in *ReplicasObservation) DeepCopy() *ReplicasObservation {
	if in == nil {
		return nil
	}
	out := new(ReplicasObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicasParameters) DeepCopyInto(out *ReplicasParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicasParameters.
func (in *ReplicasParameters) DeepCopy() *ReplicasParameters {
	if in == nil {
		return nil
	}
	out := new(ReplicasParameters)
	in.DeepCopyInto(out)
	return out
}
