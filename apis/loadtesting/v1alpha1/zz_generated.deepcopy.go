//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Agent) DeepCopyInto(out *Agent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Agent.
func (in *Agent) DeepCopy() *Agent {
	if in == nil {
		return nil
	}
	out := new(Agent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Agent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentInitParameters) DeepCopyInto(out *AgentInitParameters) {
	*out = *in
	if in.ComputeInstance != nil {
		in, out := &in.ComputeInstance, &out.ComputeInstance
		*out = make([]ComputeInstanceInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentInitParameters.
func (in *AgentInitParameters) DeepCopy() *AgentInitParameters {
	if in == nil {
		return nil
	}
	out := new(AgentInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentList) DeepCopyInto(out *AgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Agent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentList.
func (in *AgentList) DeepCopy() *AgentList {
	if in == nil {
		return nil
	}
	out := new(AgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentObservation) DeepCopyInto(out *AgentObservation) {
	*out = *in
	if in.ComputeInstance != nil {
		in, out := &in.ComputeInstance, &out.ComputeInstance
		*out = make([]ComputeInstanceObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ComputeInstanceID != nil {
		in, out := &in.ComputeInstanceID, &out.ComputeInstanceID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentObservation.
func (in *AgentObservation) DeepCopy() *AgentObservation {
	if in == nil {
		return nil
	}
	out := new(AgentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentParameters) DeepCopyInto(out *AgentParameters) {
	*out = *in
	if in.ComputeInstance != nil {
		in, out := &in.ComputeInstance, &out.ComputeInstance
		*out = make([]ComputeInstanceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.FolderIDRef != nil {
		in, out := &in.FolderIDRef, &out.FolderIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FolderIDSelector != nil {
		in, out := &in.FolderIDSelector, &out.FolderIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentParameters.
func (in *AgentParameters) DeepCopy() *AgentParameters {
	if in == nil {
		return nil
	}
	out := new(AgentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentSpec) DeepCopyInto(out *AgentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSpec.
func (in *AgentSpec) DeepCopy() *AgentSpec {
	if in == nil {
		return nil
	}
	out := new(AgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentStatus) DeepCopyInto(out *AgentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentStatus.
func (in *AgentStatus) DeepCopy() *AgentStatus {
	if in == nil {
		return nil
	}
	out := new(AgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootDiskInitParameters) DeepCopyInto(out *BootDiskInitParameters) {
	*out = *in
	if in.AutoDelete != nil {
		in, out := &in.AutoDelete, &out.AutoDelete
		*out = new(bool)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.InitializeParams != nil {
		in, out := &in.InitializeParams, &out.InitializeParams
		*out = make([]InitializeParamsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootDiskInitParameters.
func (in *BootDiskInitParameters) DeepCopy() *BootDiskInitParameters {
	if in == nil {
		return nil
	}
	out := new(BootDiskInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootDiskObservation) DeepCopyInto(out *BootDiskObservation) {
	*out = *in
	if in.AutoDelete != nil {
		in, out := &in.AutoDelete, &out.AutoDelete
		*out = new(bool)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(string)
		**out = **in
	}
	if in.InitializeParams != nil {
		in, out := &in.InitializeParams, &out.InitializeParams
		*out = make([]InitializeParamsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootDiskObservation.
func (in *BootDiskObservation) DeepCopy() *BootDiskObservation {
	if in == nil {
		return nil
	}
	out := new(BootDiskObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootDiskParameters) DeepCopyInto(out *BootDiskParameters) {
	*out = *in
	if in.AutoDelete != nil {
		in, out := &in.AutoDelete, &out.AutoDelete
		*out = new(bool)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.InitializeParams != nil {
		in, out := &in.InitializeParams, &out.InitializeParams
		*out = make([]InitializeParamsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootDiskParameters.
func (in *BootDiskParameters) DeepCopy() *BootDiskParameters {
	if in == nil {
		return nil
	}
	out := new(BootDiskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeInstanceInitParameters) DeepCopyInto(out *ComputeInstanceInitParameters) {
	*out = *in
	if in.BootDisk != nil {
		in, out := &in.BootDisk, &out.BootDisk
		*out = make([]BootDiskInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
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
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]NetworkInterfaceInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourcesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeInstanceInitParameters.
func (in *ComputeInstanceInitParameters) DeepCopy() *ComputeInstanceInitParameters {
	if in == nil {
		return nil
	}
	out := new(ComputeInstanceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeInstanceObservation) DeepCopyInto(out *ComputeInstanceObservation) {
	*out = *in
	if in.BootDisk != nil {
		in, out := &in.BootDisk, &out.BootDisk
		*out = make([]BootDiskObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ComputedMetadata != nil {
		in, out := &in.ComputedMetadata, &out.ComputedMetadata
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
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
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
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]NetworkInterfaceObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourcesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeInstanceObservation.
func (in *ComputeInstanceObservation) DeepCopy() *ComputeInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(ComputeInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeInstanceParameters) DeepCopyInto(out *ComputeInstanceParameters) {
	*out = *in
	if in.BootDisk != nil {
		in, out := &in.BootDisk, &out.BootDisk
		*out = make([]BootDiskParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
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
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]NetworkInterfaceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourcesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountIDRef != nil {
		in, out := &in.ServiceAccountIDRef, &out.ServiceAccountIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccountIDSelector != nil {
		in, out := &in.ServiceAccountIDSelector, &out.ServiceAccountIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeInstanceParameters.
func (in *ComputeInstanceParameters) DeepCopy() *ComputeInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(ComputeInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializeParamsInitParameters) DeepCopyInto(out *InitializeParamsInitParameters) {
	*out = *in
	if in.BlockSize != nil {
		in, out := &in.BlockSize, &out.BlockSize
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializeParamsInitParameters.
func (in *InitializeParamsInitParameters) DeepCopy() *InitializeParamsInitParameters {
	if in == nil {
		return nil
	}
	out := new(InitializeParamsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializeParamsObservation) DeepCopyInto(out *InitializeParamsObservation) {
	*out = *in
	if in.BlockSize != nil {
		in, out := &in.BlockSize, &out.BlockSize
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializeParamsObservation.
func (in *InitializeParamsObservation) DeepCopy() *InitializeParamsObservation {
	if in == nil {
		return nil
	}
	out := new(InitializeParamsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializeParamsParameters) DeepCopyInto(out *InitializeParamsParameters) {
	*out = *in
	if in.BlockSize != nil {
		in, out := &in.BlockSize, &out.BlockSize
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializeParamsParameters.
func (in *InitializeParamsParameters) DeepCopy() *InitializeParamsParameters {
	if in == nil {
		return nil
	}
	out := new(InitializeParamsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceInitParameters) DeepCopyInto(out *NetworkInterfaceInitParameters) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(bool)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.NAT != nil {
		in, out := &in.NAT, &out.NAT
		*out = new(bool)
		**out = **in
	}
	if in.NATIPAddress != nil {
		in, out := &in.NATIPAddress, &out.NATIPAddress
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceInitParameters.
func (in *NetworkInterfaceInitParameters) DeepCopy() *NetworkInterfaceInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceObservation) DeepCopyInto(out *NetworkInterfaceObservation) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(bool)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.Index != nil {
		in, out := &in.Index, &out.Index
		*out = new(float64)
		**out = **in
	}
	if in.MacAddress != nil {
		in, out := &in.MacAddress, &out.MacAddress
		*out = new(string)
		**out = **in
	}
	if in.NAT != nil {
		in, out := &in.NAT, &out.NAT
		*out = new(bool)
		**out = **in
	}
	if in.NATIPAddress != nil {
		in, out := &in.NATIPAddress, &out.NATIPAddress
		*out = new(string)
		**out = **in
	}
	if in.NATIPVersion != nil {
		in, out := &in.NATIPVersion, &out.NATIPVersion
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceObservation.
func (in *NetworkInterfaceObservation) DeepCopy() *NetworkInterfaceObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInterfaceParameters) DeepCopyInto(out *NetworkInterfaceParameters) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(bool)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.NAT != nil {
		in, out := &in.NAT, &out.NAT
		*out = new(bool)
		**out = **in
	}
	if in.NATIPAddress != nil {
		in, out := &in.NATIPAddress, &out.NATIPAddress
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInterfaceParameters.
func (in *NetworkInterfaceParameters) DeepCopy() *NetworkInterfaceParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkInterfaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesInitParameters) DeepCopyInto(out *ResourcesInitParameters) {
	*out = *in
	if in.CoreFraction != nil {
		in, out := &in.CoreFraction, &out.CoreFraction
		*out = new(float64)
		**out = **in
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesInitParameters.
func (in *ResourcesInitParameters) DeepCopy() *ResourcesInitParameters {
	if in == nil {
		return nil
	}
	out := new(ResourcesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesObservation) DeepCopyInto(out *ResourcesObservation) {
	*out = *in
	if in.CoreFraction != nil {
		in, out := &in.CoreFraction, &out.CoreFraction
		*out = new(float64)
		**out = **in
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesObservation.
func (in *ResourcesObservation) DeepCopy() *ResourcesObservation {
	if in == nil {
		return nil
	}
	out := new(ResourcesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesParameters) DeepCopyInto(out *ResourcesParameters) {
	*out = *in
	if in.CoreFraction != nil {
		in, out := &in.CoreFraction, &out.CoreFraction
		*out = new(float64)
		**out = **in
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesParameters.
func (in *ResourcesParameters) DeepCopy() *ResourcesParameters {
	if in == nil {
		return nil
	}
	out := new(ResourcesParameters)
	in.DeepCopyInto(out)
	return out
}
