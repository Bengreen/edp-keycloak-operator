//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package common

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceRef) DeepCopyInto(out *SourceRef) {
	*out = *in
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(ConfigMapKeySelector)
		**out = **in
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceRef.
func (in *SourceRef) DeepCopy() *SourceRef {
	if in == nil {
		return nil
	}
	out := new(SourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenSettings) DeepCopyInto(out *TokenSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenSettings.
func (in *TokenSettings) DeepCopy() *TokenSettings {
	if in == nil {
		return nil
	}
	out := new(TokenSettings)
	in.DeepCopyInto(out)
	return out
}
