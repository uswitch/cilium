// +build !ignore_autogenerated

// Copyright 2017-2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package ctmap

import (
	bpf "github.com/cilium/cilium/pkg/bpf"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CtEntry) DeepCopyInto(out *CtEntry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CtEntry.
func (in *CtEntry) DeepCopy() *CtEntry {
	if in == nil {
		return nil
	}
	out := new(CtEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (in *CtEntry) DeepCopyMapValue() bpf.MapValue {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CtKey4) DeepCopyInto(out *CtKey4) {
	*out = *in
	in.TupleKey4.DeepCopyInto(&out.TupleKey4)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CtKey4.
func (in *CtKey4) DeepCopy() *CtKey4 {
	if in == nil {
		return nil
	}
	out := new(CtKey4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *CtKey4) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CtKey4Global) DeepCopyInto(out *CtKey4Global) {
	*out = *in
	in.TupleKey4Global.DeepCopyInto(&out.TupleKey4Global)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CtKey4Global.
func (in *CtKey4Global) DeepCopy() *CtKey4Global {
	if in == nil {
		return nil
	}
	out := new(CtKey4Global)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *CtKey4Global) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CtKey6) DeepCopyInto(out *CtKey6) {
	*out = *in
	in.TupleKey6.DeepCopyInto(&out.TupleKey6)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CtKey6.
func (in *CtKey6) DeepCopy() *CtKey6 {
	if in == nil {
		return nil
	}
	out := new(CtKey6)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *CtKey6) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CtKey6Global) DeepCopyInto(out *CtKey6Global) {
	*out = *in
	in.TupleKey6Global.DeepCopyInto(&out.TupleKey6Global)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CtKey6Global.
func (in *CtKey6Global) DeepCopy() *CtKey6Global {
	if in == nil {
		return nil
	}
	out := new(CtKey6Global)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *CtKey6Global) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
