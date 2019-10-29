// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Buildpack) DeepCopyInto(out *Buildpack) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Buildpack.
func (in *Buildpack) DeepCopy() *Buildpack {
	if in == nil {
		return nil
	}
	out := new(Buildpack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildpackLifecycle) DeepCopyInto(out *BuildpackLifecycle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildpackLifecycle.
func (in *BuildpackLifecycle) DeepCopy() *BuildpackLifecycle {
	if in == nil {
		return nil
	}
	out := new(BuildpackLifecycle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerLifecycle) DeepCopyInto(out *DockerLifecycle) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerLifecycle.
func (in *DockerLifecycle) DeepCopy() *DockerLifecycle {
	if in == nil {
		return nil
	}
	out := new(DockerLifecycle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LRP) DeepCopyInto(out *LRP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LRP.
func (in *LRP) DeepCopy() *LRP {
	if in == nil {
		return nil
	}
	out := new(LRP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LRP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LRPList) DeepCopyInto(out *LRPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LRP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LRPList.
func (in *LRPList) DeepCopy() *LRPList {
	if in == nil {
		return nil
	}
	out := new(LRPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LRPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LRPSpec) DeepCopyInto(out *LRPSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make(map[string]*json.RawMessage, len(*in))
		for key, val := range *in {
			var outVal *json.RawMessage
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(json.RawMessage)
				if **in != nil {
					in, out := *in, *out
					*out = make([]byte, len(*in))
					copy(*out, *in)
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]VolumeMount, len(*in))
		copy(*out, *in)
	}
	in.Lifecycle.DeepCopyInto(&out.Lifecycle)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LRPSpec.
func (in *LRPSpec) DeepCopy() *LRPSpec {
	if in == nil {
		return nil
	}
	out := new(LRPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LRPStatus) DeepCopyInto(out *LRPStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LRPStatus.
func (in *LRPStatus) DeepCopy() *LRPStatus {
	if in == nil {
		return nil
	}
	out := new(LRPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lifecycle) DeepCopyInto(out *Lifecycle) {
	*out = *in
	if in.DockerLifecycle != nil {
		in, out := &in.DockerLifecycle, &out.DockerLifecycle
		*out = new(DockerLifecycle)
		(*in).DeepCopyInto(*out)
	}
	if in.BuildpackLifecycle != nil {
		in, out := &in.BuildpackLifecycle, &out.BuildpackLifecycle
		*out = new(BuildpackLifecycle)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lifecycle.
func (in *Lifecycle) DeepCopy() *Lifecycle {
	if in == nil {
		return nil
	}
	out := new(Lifecycle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleData) DeepCopyInto(out *LifecycleData) {
	*out = *in
	if in.Buildpacks != nil {
		in, out := &in.Buildpacks, &out.Buildpacks
		*out = make([]Buildpack, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleData.
func (in *LifecycleData) DeepCopy() *LifecycleData {
	if in == nil {
		return nil
	}
	out := new(LifecycleData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Staging) DeepCopyInto(out *Staging) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Staging.
func (in *Staging) DeepCopy() *Staging {
	if in == nil {
		return nil
	}
	out := new(Staging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Staging) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StagingList) DeepCopyInto(out *StagingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Staging, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StagingList.
func (in *StagingList) DeepCopy() *StagingList {
	if in == nil {
		return nil
	}
	out := new(StagingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StagingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StagingSpec) DeepCopyInto(out *StagingSpec) {
	*out = *in
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
	in.LifecycleData.DeepCopyInto(&out.LifecycleData)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StagingSpec.
func (in *StagingSpec) DeepCopy() *StagingSpec {
	if in == nil {
		return nil
	}
	out := new(StagingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StagingStatus) DeepCopyInto(out *StagingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StagingStatus.
func (in *StagingStatus) DeepCopy() *StagingStatus {
	if in == nil {
		return nil
	}
	out := new(StagingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeMount) DeepCopyInto(out *VolumeMount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeMount.
func (in *VolumeMount) DeepCopy() *VolumeMount {
	if in == nil {
		return nil
	}
	out := new(VolumeMount)
	in.DeepCopyInto(out)
	return out
}