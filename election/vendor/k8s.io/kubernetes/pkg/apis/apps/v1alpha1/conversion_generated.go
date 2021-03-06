// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	api "k8s.io/kubernetes/pkg/api"
	apps "k8s.io/kubernetes/pkg/apis/apps"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_PetSet_To_apps_PetSet,
		Convert_apps_PetSet_To_v1alpha1_PetSet,
		Convert_v1alpha1_PetSetList_To_apps_PetSetList,
		Convert_apps_PetSetList_To_v1alpha1_PetSetList,
		Convert_v1alpha1_PetSetSpec_To_apps_PetSetSpec,
		Convert_apps_PetSetSpec_To_v1alpha1_PetSetSpec,
		Convert_v1alpha1_PetSetStatus_To_apps_PetSetStatus,
		Convert_apps_PetSetStatus_To_v1alpha1_PetSetStatus,
	); err != nil {
		// if one of the conversion functions is malformed, detect it immediately.
		panic(err)
	}
}

func autoConvert_v1alpha1_PetSet_To_apps_PetSet(in *PetSet, out *apps.PetSet, s conversion.Scope) error {
	SetDefaults_PetSet(in)
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PetSetSpec_To_apps_PetSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PetSetStatus_To_apps_PetSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_PetSet_To_apps_PetSet(in *PetSet, out *apps.PetSet, s conversion.Scope) error {
	return autoConvert_v1alpha1_PetSet_To_apps_PetSet(in, out, s)
}

func autoConvert_apps_PetSet_To_v1alpha1_PetSet(in *apps.PetSet, out *PetSet, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_apps_PetSetSpec_To_v1alpha1_PetSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_apps_PetSetStatus_To_v1alpha1_PetSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_apps_PetSet_To_v1alpha1_PetSet(in *apps.PetSet, out *PetSet, s conversion.Scope) error {
	return autoConvert_apps_PetSet_To_v1alpha1_PetSet(in, out, s)
}

func autoConvert_v1alpha1_PetSetList_To_apps_PetSetList(in *PetSetList, out *apps.PetSetList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]apps.PetSet, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_PetSet_To_apps_PetSet(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1alpha1_PetSetList_To_apps_PetSetList(in *PetSetList, out *apps.PetSetList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PetSetList_To_apps_PetSetList(in, out, s)
}

func autoConvert_apps_PetSetList_To_v1alpha1_PetSetList(in *apps.PetSetList, out *PetSetList, s conversion.Scope) error {
	if err := api.Convert_unversioned_TypeMeta_To_unversioned_TypeMeta(&in.TypeMeta, &out.TypeMeta, s); err != nil {
		return err
	}
	if err := api.Convert_unversioned_ListMeta_To_unversioned_ListMeta(&in.ListMeta, &out.ListMeta, s); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PetSet, len(*in))
		for i := range *in {
			if err := Convert_apps_PetSet_To_v1alpha1_PetSet(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_apps_PetSetList_To_v1alpha1_PetSetList(in *apps.PetSetList, out *PetSetList, s conversion.Scope) error {
	return autoConvert_apps_PetSetList_To_v1alpha1_PetSetList(in, out, s)
}

func autoConvert_v1alpha1_PetSetStatus_To_apps_PetSetStatus(in *PetSetStatus, out *apps.PetSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = int(in.Replicas)
	return nil
}

func Convert_v1alpha1_PetSetStatus_To_apps_PetSetStatus(in *PetSetStatus, out *apps.PetSetStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_PetSetStatus_To_apps_PetSetStatus(in, out, s)
}

func autoConvert_apps_PetSetStatus_To_v1alpha1_PetSetStatus(in *apps.PetSetStatus, out *PetSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = int32(in.Replicas)
	return nil
}

func Convert_apps_PetSetStatus_To_v1alpha1_PetSetStatus(in *apps.PetSetStatus, out *PetSetStatus, s conversion.Scope) error {
	return autoConvert_apps_PetSetStatus_To_v1alpha1_PetSetStatus(in, out, s)
}
