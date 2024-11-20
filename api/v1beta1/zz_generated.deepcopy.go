//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"github.com/openstack-k8s-operators/lib-common/modules/common/service"
	"github.com/openstack-k8s-operators/lib-common/modules/storage"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIOverrideSpec) DeepCopyInto(out *APIOverrideSpec) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = make(map[service.Endpoint]service.RoutedOverrideSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIOverrideSpec.
func (in *APIOverrideSpec) DeepCopy() *APIOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(APIOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cinder) DeepCopyInto(out *Cinder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cinder.
func (in *Cinder) DeepCopy() *Cinder {
	if in == nil {
		return nil
	}
	out := new(Cinder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cinder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderAPI) DeepCopyInto(out *CinderAPI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderAPI.
func (in *CinderAPI) DeepCopy() *CinderAPI {
	if in == nil {
		return nil
	}
	out := new(CinderAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CinderAPI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderAPIList) DeepCopyInto(out *CinderAPIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CinderAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderAPIList.
func (in *CinderAPIList) DeepCopy() *CinderAPIList {
	if in == nil {
		return nil
	}
	out := new(CinderAPIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CinderAPIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderAPISpec) DeepCopyInto(out *CinderAPISpec) {
	*out = *in
	out.CinderTemplate = in.CinderTemplate
	in.CinderAPITemplate.DeepCopyInto(&out.CinderAPITemplate)
	if in.ExtraMounts != nil {
		in, out := &in.ExtraMounts, &out.ExtraMounts
		*out = make([]CinderExtraVolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderAPISpec.
func (in *CinderAPISpec) DeepCopy() *CinderAPISpec {
	if in == nil {
		return nil
	}
	out := new(CinderAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderAPIStatus) DeepCopyInto(out *CinderAPIStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceIDs != nil {
		in, out := &in.ServiceIDs, &out.ServiceIDs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderAPIStatus.
func (in *CinderAPIStatus) DeepCopy() *CinderAPIStatus {
	if in == nil {
		return nil
	}
	out := new(CinderAPIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderAPITemplate) DeepCopyInto(out *CinderAPITemplate) {
	*out = *in
	in.CinderAPITemplateCore.DeepCopyInto(&out.CinderAPITemplateCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderAPITemplate.
func (in *CinderAPITemplate) DeepCopy() *CinderAPITemplate {
	if in == nil {
		return nil
	}
	out := new(CinderAPITemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderAPITemplateCore) DeepCopyInto(out *CinderAPITemplateCore) {
	*out = *in
	in.CinderServiceTemplate.DeepCopyInto(&out.CinderServiceTemplate)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Override.DeepCopyInto(&out.Override)
	in.TLS.DeepCopyInto(&out.TLS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderAPITemplateCore.
func (in *CinderAPITemplateCore) DeepCopy() *CinderAPITemplateCore {
	if in == nil {
		return nil
	}
	out := new(CinderAPITemplateCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderBackup) DeepCopyInto(out *CinderBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderBackup.
func (in *CinderBackup) DeepCopy() *CinderBackup {
	if in == nil {
		return nil
	}
	out := new(CinderBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CinderBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderBackupList) DeepCopyInto(out *CinderBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CinderBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderBackupList.
func (in *CinderBackupList) DeepCopy() *CinderBackupList {
	if in == nil {
		return nil
	}
	out := new(CinderBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CinderBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderBackupSpec) DeepCopyInto(out *CinderBackupSpec) {
	*out = *in
	out.CinderTemplate = in.CinderTemplate
	in.CinderBackupTemplate.DeepCopyInto(&out.CinderBackupTemplate)
	if in.ExtraMounts != nil {
		in, out := &in.ExtraMounts, &out.ExtraMounts
		*out = make([]CinderExtraVolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderBackupSpec.
func (in *CinderBackupSpec) DeepCopy() *CinderBackupSpec {
	if in == nil {
		return nil
	}
	out := new(CinderBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderBackupStatus) DeepCopyInto(out *CinderBackupStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderBackupStatus.
func (in *CinderBackupStatus) DeepCopy() *CinderBackupStatus {
	if in == nil {
		return nil
	}
	out := new(CinderBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderBackupTemplate) DeepCopyInto(out *CinderBackupTemplate) {
	*out = *in
	in.CinderBackupTemplateCore.DeepCopyInto(&out.CinderBackupTemplateCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderBackupTemplate.
func (in *CinderBackupTemplate) DeepCopy() *CinderBackupTemplate {
	if in == nil {
		return nil
	}
	out := new(CinderBackupTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderBackupTemplateCore) DeepCopyInto(out *CinderBackupTemplateCore) {
	*out = *in
	in.CinderServiceTemplate.DeepCopyInto(&out.CinderServiceTemplate)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderBackupTemplateCore.
func (in *CinderBackupTemplateCore) DeepCopy() *CinderBackupTemplateCore {
	if in == nil {
		return nil
	}
	out := new(CinderBackupTemplateCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderDefaults) DeepCopyInto(out *CinderDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderDefaults.
func (in *CinderDefaults) DeepCopy() *CinderDefaults {
	if in == nil {
		return nil
	}
	out := new(CinderDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderExtraVolMounts) DeepCopyInto(out *CinderExtraVolMounts) {
	*out = *in
	if in.VolMounts != nil {
		in, out := &in.VolMounts, &out.VolMounts
		*out = make([]storage.VolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderExtraVolMounts.
func (in *CinderExtraVolMounts) DeepCopy() *CinderExtraVolMounts {
	if in == nil {
		return nil
	}
	out := new(CinderExtraVolMounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderList) DeepCopyInto(out *CinderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cinder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderList.
func (in *CinderList) DeepCopy() *CinderList {
	if in == nil {
		return nil
	}
	out := new(CinderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CinderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderScheduler) DeepCopyInto(out *CinderScheduler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderScheduler.
func (in *CinderScheduler) DeepCopy() *CinderScheduler {
	if in == nil {
		return nil
	}
	out := new(CinderScheduler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CinderScheduler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSchedulerList) DeepCopyInto(out *CinderSchedulerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CinderScheduler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSchedulerList.
func (in *CinderSchedulerList) DeepCopy() *CinderSchedulerList {
	if in == nil {
		return nil
	}
	out := new(CinderSchedulerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CinderSchedulerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSchedulerSpec) DeepCopyInto(out *CinderSchedulerSpec) {
	*out = *in
	out.CinderTemplate = in.CinderTemplate
	in.CinderSchedulerTemplate.DeepCopyInto(&out.CinderSchedulerTemplate)
	if in.ExtraMounts != nil {
		in, out := &in.ExtraMounts, &out.ExtraMounts
		*out = make([]CinderExtraVolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSchedulerSpec.
func (in *CinderSchedulerSpec) DeepCopy() *CinderSchedulerSpec {
	if in == nil {
		return nil
	}
	out := new(CinderSchedulerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSchedulerStatus) DeepCopyInto(out *CinderSchedulerStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSchedulerStatus.
func (in *CinderSchedulerStatus) DeepCopy() *CinderSchedulerStatus {
	if in == nil {
		return nil
	}
	out := new(CinderSchedulerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSchedulerTemplate) DeepCopyInto(out *CinderSchedulerTemplate) {
	*out = *in
	in.CinderSchedulerTemplateCore.DeepCopyInto(&out.CinderSchedulerTemplateCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSchedulerTemplate.
func (in *CinderSchedulerTemplate) DeepCopy() *CinderSchedulerTemplate {
	if in == nil {
		return nil
	}
	out := new(CinderSchedulerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSchedulerTemplateCore) DeepCopyInto(out *CinderSchedulerTemplateCore) {
	*out = *in
	in.CinderServiceTemplate.DeepCopyInto(&out.CinderServiceTemplate)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSchedulerTemplateCore.
func (in *CinderSchedulerTemplateCore) DeepCopy() *CinderSchedulerTemplateCore {
	if in == nil {
		return nil
	}
	out := new(CinderSchedulerTemplateCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderServiceTemplate) DeepCopyInto(out *CinderServiceTemplate) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.CustomServiceConfigSecrets != nil {
		in, out := &in.CustomServiceConfigSecrets, &out.CustomServiceConfigSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderServiceTemplate.
func (in *CinderServiceTemplate) DeepCopy() *CinderServiceTemplate {
	if in == nil {
		return nil
	}
	out := new(CinderServiceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSpec) DeepCopyInto(out *CinderSpec) {
	*out = *in
	in.CinderSpecBase.DeepCopyInto(&out.CinderSpecBase)
	in.CinderAPI.DeepCopyInto(&out.CinderAPI)
	in.CinderScheduler.DeepCopyInto(&out.CinderScheduler)
	in.CinderBackup.DeepCopyInto(&out.CinderBackup)
	if in.CinderVolumes != nil {
		in, out := &in.CinderVolumes, &out.CinderVolumes
		*out = make(map[string]CinderVolumeTemplate, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSpec.
func (in *CinderSpec) DeepCopy() *CinderSpec {
	if in == nil {
		return nil
	}
	out := new(CinderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSpecBase) DeepCopyInto(out *CinderSpecBase) {
	*out = *in
	out.CinderTemplate = in.CinderTemplate
	if in.ExtraMounts != nil {
		in, out := &in.ExtraMounts, &out.ExtraMounts
		*out = make([]CinderExtraVolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	out.DBPurge = in.DBPurge
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSpecBase.
func (in *CinderSpecBase) DeepCopy() *CinderSpecBase {
	if in == nil {
		return nil
	}
	out := new(CinderSpecBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSpecCore) DeepCopyInto(out *CinderSpecCore) {
	*out = *in
	in.CinderSpecBase.DeepCopyInto(&out.CinderSpecBase)
	in.CinderAPI.DeepCopyInto(&out.CinderAPI)
	in.CinderScheduler.DeepCopyInto(&out.CinderScheduler)
	in.CinderBackup.DeepCopyInto(&out.CinderBackup)
	if in.CinderVolumes != nil {
		in, out := &in.CinderVolumes, &out.CinderVolumes
		*out = make(map[string]CinderVolumeTemplateCore, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSpecCore.
func (in *CinderSpecCore) DeepCopy() *CinderSpecCore {
	if in == nil {
		return nil
	}
	out := new(CinderSpecCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderStatus) DeepCopyInto(out *CinderStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.ServiceIDs != nil {
		in, out := &in.ServiceIDs, &out.ServiceIDs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CinderVolumesReadyCounts != nil {
		in, out := &in.CinderVolumesReadyCounts, &out.CinderVolumesReadyCounts
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderStatus.
func (in *CinderStatus) DeepCopy() *CinderStatus {
	if in == nil {
		return nil
	}
	out := new(CinderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderTemplate) DeepCopyInto(out *CinderTemplate) {
	*out = *in
	out.PasswordSelectors = in.PasswordSelectors
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderTemplate.
func (in *CinderTemplate) DeepCopy() *CinderTemplate {
	if in == nil {
		return nil
	}
	out := new(CinderTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderVolume) DeepCopyInto(out *CinderVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderVolume.
func (in *CinderVolume) DeepCopy() *CinderVolume {
	if in == nil {
		return nil
	}
	out := new(CinderVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CinderVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderVolumeList) DeepCopyInto(out *CinderVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CinderVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderVolumeList.
func (in *CinderVolumeList) DeepCopy() *CinderVolumeList {
	if in == nil {
		return nil
	}
	out := new(CinderVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CinderVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderVolumeSpec) DeepCopyInto(out *CinderVolumeSpec) {
	*out = *in
	out.CinderTemplate = in.CinderTemplate
	in.CinderVolumeTemplate.DeepCopyInto(&out.CinderVolumeTemplate)
	if in.ExtraMounts != nil {
		in, out := &in.ExtraMounts, &out.ExtraMounts
		*out = make([]CinderExtraVolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderVolumeSpec.
func (in *CinderVolumeSpec) DeepCopy() *CinderVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(CinderVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderVolumeStatus) DeepCopyInto(out *CinderVolumeStatus) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderVolumeStatus.
func (in *CinderVolumeStatus) DeepCopy() *CinderVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(CinderVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderVolumeTemplate) DeepCopyInto(out *CinderVolumeTemplate) {
	*out = *in
	in.CinderVolumeTemplateCore.DeepCopyInto(&out.CinderVolumeTemplateCore)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderVolumeTemplate.
func (in *CinderVolumeTemplate) DeepCopy() *CinderVolumeTemplate {
	if in == nil {
		return nil
	}
	out := new(CinderVolumeTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderVolumeTemplateCore) DeepCopyInto(out *CinderVolumeTemplateCore) {
	*out = *in
	in.CinderServiceTemplate.DeepCopyInto(&out.CinderServiceTemplate)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderVolumeTemplateCore.
func (in *CinderVolumeTemplateCore) DeepCopy() *CinderVolumeTemplateCore {
	if in == nil {
		return nil
	}
	out := new(CinderVolumeTemplateCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBPurge) DeepCopyInto(out *DBPurge) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBPurge.
func (in *DBPurge) DeepCopy() *DBPurge {
	if in == nil {
		return nil
	}
	out := new(DBPurge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordSelector) DeepCopyInto(out *PasswordSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSelector.
func (in *PasswordSelector) DeepCopy() *PasswordSelector {
	if in == nil {
		return nil
	}
	out := new(PasswordSelector)
	in.DeepCopyInto(out)
	return out
}
