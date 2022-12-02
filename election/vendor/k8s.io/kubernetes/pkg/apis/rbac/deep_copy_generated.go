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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package rbac

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_rbac_ClusterRole,
		DeepCopy_rbac_ClusterRoleBinding,
		DeepCopy_rbac_ClusterRoleBindingList,
		DeepCopy_rbac_ClusterRoleList,
		DeepCopy_rbac_PolicyRule,
		DeepCopy_rbac_Role,
		DeepCopy_rbac_RoleBinding,
		DeepCopy_rbac_RoleBindingList,
		DeepCopy_rbac_RoleList,
		DeepCopy_rbac_Subject,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_rbac_ClusterRole(in ClusterRole, out *ClusterRole, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api.DeepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if in.Rules != nil {
		in, out := in.Rules, &out.Rules
		*out = make([]PolicyRule, len(in))
		for i := range in {
			if err := DeepCopy_rbac_PolicyRule(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func DeepCopy_rbac_ClusterRoleBinding(in ClusterRoleBinding, out *ClusterRoleBinding, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api.DeepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if in.Subjects != nil {
		in, out := in.Subjects, &out.Subjects
		*out = make([]Subject, len(in))
		for i := range in {
			if err := DeepCopy_rbac_Subject(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Subjects = nil
	}
	if err := api.DeepCopy_api_ObjectReference(in.RoleRef, &out.RoleRef, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_rbac_ClusterRoleBindingList(in ClusterRoleBindingList, out *ClusterRoleBindingList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]ClusterRoleBinding, len(in))
		for i := range in {
			if err := DeepCopy_rbac_ClusterRoleBinding(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_rbac_ClusterRoleList(in ClusterRoleList, out *ClusterRoleList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]ClusterRole, len(in))
		for i := range in {
			if err := DeepCopy_rbac_ClusterRole(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_rbac_PolicyRule(in PolicyRule, out *PolicyRule, c *conversion.Cloner) error {
	if in.Verbs != nil {
		in, out := in.Verbs, &out.Verbs
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Verbs = nil
	}
	if in.AttributeRestrictions == nil {
		out.AttributeRestrictions = nil
	} else if newVal, err := c.DeepCopy(in.AttributeRestrictions); err != nil {
		return err
	} else {
		out.AttributeRestrictions = newVal.(runtime.Object)
	}
	if in.APIGroups != nil {
		in, out := in.APIGroups, &out.APIGroups
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.APIGroups = nil
	}
	if in.Resources != nil {
		in, out := in.Resources, &out.Resources
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Resources = nil
	}
	if in.ResourceNames != nil {
		in, out := in.ResourceNames, &out.ResourceNames
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.ResourceNames = nil
	}
	if in.NonResourceURLs != nil {
		in, out := in.NonResourceURLs, &out.NonResourceURLs
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.NonResourceURLs = nil
	}
	return nil
}

func DeepCopy_rbac_Role(in Role, out *Role, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api.DeepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if in.Rules != nil {
		in, out := in.Rules, &out.Rules
		*out = make([]PolicyRule, len(in))
		for i := range in {
			if err := DeepCopy_rbac_PolicyRule(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func DeepCopy_rbac_RoleBinding(in RoleBinding, out *RoleBinding, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := api.DeepCopy_api_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if in.Subjects != nil {
		in, out := in.Subjects, &out.Subjects
		*out = make([]Subject, len(in))
		for i := range in {
			if err := DeepCopy_rbac_Subject(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Subjects = nil
	}
	if err := api.DeepCopy_api_ObjectReference(in.RoleRef, &out.RoleRef, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_rbac_RoleBindingList(in RoleBindingList, out *RoleBindingList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]RoleBinding, len(in))
		for i := range in {
			if err := DeepCopy_rbac_RoleBinding(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_rbac_RoleList(in RoleList, out *RoleList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]Role, len(in))
		for i := range in {
			if err := DeepCopy_rbac_Role(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_rbac_Subject(in Subject, out *Subject, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.APIVersion = in.APIVersion
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}
