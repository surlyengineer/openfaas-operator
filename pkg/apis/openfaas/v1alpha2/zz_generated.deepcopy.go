// +build !ignore_autogenerated

/*
Copyright 2018 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha2

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Function).DeepCopyInto(out.(*Function))
			return nil
		}, InType: reflect.TypeOf(&Function{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FunctionList).DeepCopyInto(out.(*FunctionList))
			return nil
		}, InType: reflect.TypeOf(&FunctionList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FunctionResources).DeepCopyInto(out.(*FunctionResources))
			return nil
		}, InType: reflect.TypeOf(&FunctionResources{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FunctionSpec).DeepCopyInto(out.(*FunctionSpec))
			return nil
		}, InType: reflect.TypeOf(&FunctionSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*FunctionStatus).DeepCopyInto(out.(*FunctionStatus))
			return nil
		}, InType: reflect.TypeOf(&FunctionStatus{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionResources) DeepCopyInto(out *FunctionResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionResources.
func (in *FunctionResources) DeepCopy() *FunctionResources {
	if in == nil {
		return nil
	}
	out := new(FunctionResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		if *in == nil {
			*out = nil
		} else {
			*out = new(map[string]string)
			if **in != nil {
				in, out := *in, *out
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		if *in == nil {
			*out = nil
		} else {
			*out = new(map[string]string)
			if **in != nil {
				in, out := *in, *out
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.Constraints != nil {
		in, out := &in.Constraints, &out.Constraints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		if *in == nil {
			*out = nil
		} else {
			*out = new(FunctionResources)
			**out = **in
		}
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		if *in == nil {
			*out = nil
		} else {
			*out = new(FunctionResources)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}