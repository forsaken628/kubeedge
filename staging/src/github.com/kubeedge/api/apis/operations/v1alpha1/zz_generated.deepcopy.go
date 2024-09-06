//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KubeEdge Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePrePullJob) DeepCopyInto(out *ImagePrePullJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePrePullJob.
func (in *ImagePrePullJob) DeepCopy() *ImagePrePullJob {
	if in == nil {
		return nil
	}
	out := new(ImagePrePullJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagePrePullJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePrePullJobList) DeepCopyInto(out *ImagePrePullJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImagePrePullJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePrePullJobList.
func (in *ImagePrePullJobList) DeepCopy() *ImagePrePullJobList {
	if in == nil {
		return nil
	}
	out := new(ImagePrePullJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImagePrePullJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePrePullJobSpec) DeepCopyInto(out *ImagePrePullJobSpec) {
	*out = *in
	in.ImagePrePullTemplate.DeepCopyInto(&out.ImagePrePullTemplate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePrePullJobSpec.
func (in *ImagePrePullJobSpec) DeepCopy() *ImagePrePullJobSpec {
	if in == nil {
		return nil
	}
	out := new(ImagePrePullJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePrePullJobStatus) DeepCopyInto(out *ImagePrePullJobStatus) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make([]ImagePrePullStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePrePullJobStatus.
func (in *ImagePrePullJobStatus) DeepCopy() *ImagePrePullJobStatus {
	if in == nil {
		return nil
	}
	out := new(ImagePrePullJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePrePullStatus) DeepCopyInto(out *ImagePrePullStatus) {
	*out = *in
	if in.TaskStatus != nil {
		in, out := &in.TaskStatus, &out.TaskStatus
		*out = new(TaskStatus)
		**out = **in
	}
	if in.ImageStatus != nil {
		in, out := &in.ImageStatus, &out.ImageStatus
		*out = make([]ImageStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePrePullStatus.
func (in *ImagePrePullStatus) DeepCopy() *ImagePrePullStatus {
	if in == nil {
		return nil
	}
	out := new(ImagePrePullStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePrePullTemplate) DeepCopyInto(out *ImagePrePullTemplate) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.CheckItems != nil {
		in, out := &in.CheckItems, &out.CheckItems
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(uint32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePrePullTemplate.
func (in *ImagePrePullTemplate) DeepCopy() *ImagePrePullTemplate {
	if in == nil {
		return nil
	}
	out := new(ImagePrePullTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatus) DeepCopyInto(out *ImageStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatus.
func (in *ImageStatus) DeepCopy() *ImageStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeUpgradeJob) DeepCopyInto(out *NodeUpgradeJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeUpgradeJob.
func (in *NodeUpgradeJob) DeepCopy() *NodeUpgradeJob {
	if in == nil {
		return nil
	}
	out := new(NodeUpgradeJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeUpgradeJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeUpgradeJobList) DeepCopyInto(out *NodeUpgradeJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeUpgradeJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeUpgradeJobList.
func (in *NodeUpgradeJobList) DeepCopy() *NodeUpgradeJobList {
	if in == nil {
		return nil
	}
	out := new(NodeUpgradeJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeUpgradeJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeUpgradeJobSpec) DeepCopyInto(out *NodeUpgradeJobSpec) {
	*out = *in
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(uint32)
		**out = **in
	}
	if in.NodeNames != nil {
		in, out := &in.NodeNames, &out.NodeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.CheckItems != nil {
		in, out := &in.CheckItems, &out.CheckItems
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeUpgradeJobSpec.
func (in *NodeUpgradeJobSpec) DeepCopy() *NodeUpgradeJobSpec {
	if in == nil {
		return nil
	}
	out := new(NodeUpgradeJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeUpgradeJobStatus) DeepCopyInto(out *NodeUpgradeJobStatus) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make([]TaskStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeUpgradeJobStatus.
func (in *NodeUpgradeJobStatus) DeepCopy() *NodeUpgradeJobStatus {
	if in == nil {
		return nil
	}
	out := new(NodeUpgradeJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStatus) DeepCopyInto(out *TaskStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStatus.
func (in *TaskStatus) DeepCopy() *TaskStatus {
	if in == nil {
		return nil
	}
	out := new(TaskStatus)
	in.DeepCopyInto(out)
	return out
}