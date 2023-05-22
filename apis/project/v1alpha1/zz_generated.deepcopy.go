//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApprovalSettingsObservation) DeepCopyInto(out *ApprovalSettingsObservation) {
	*out = *in
	if in.CanApplyDeclinedChanges != nil {
		in, out := &in.CanApplyDeclinedChanges, &out.CanApplyDeclinedChanges
		*out = new(bool)
		**out = **in
	}
	if in.CanReviewOwnRequest != nil {
		in, out := &in.CanReviewOwnRequest, &out.CanReviewOwnRequest
		*out = new(bool)
		**out = **in
	}
	if in.MinNumApprovals != nil {
		in, out := &in.MinNumApprovals, &out.MinNumApprovals
		*out = new(float64)
		**out = **in
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(bool)
		**out = **in
	}
	if in.RequiredApprovalTags != nil {
		in, out := &in.RequiredApprovalTags, &out.RequiredApprovalTags
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApprovalSettingsObservation.
func (in *ApprovalSettingsObservation) DeepCopy() *ApprovalSettingsObservation {
	if in == nil {
		return nil
	}
	out := new(ApprovalSettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApprovalSettingsParameters) DeepCopyInto(out *ApprovalSettingsParameters) {
	*out = *in
	if in.CanApplyDeclinedChanges != nil {
		in, out := &in.CanApplyDeclinedChanges, &out.CanApplyDeclinedChanges
		*out = new(bool)
		**out = **in
	}
	if in.CanReviewOwnRequest != nil {
		in, out := &in.CanReviewOwnRequest, &out.CanReviewOwnRequest
		*out = new(bool)
		**out = **in
	}
	if in.MinNumApprovals != nil {
		in, out := &in.MinNumApprovals, &out.MinNumApprovals
		*out = new(float64)
		**out = **in
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = new(bool)
		**out = **in
	}
	if in.RequiredApprovalTags != nil {
		in, out := &in.RequiredApprovalTags, &out.RequiredApprovalTags
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApprovalSettingsParameters.
func (in *ApprovalSettingsParameters) DeepCopy() *ApprovalSettingsParameters {
	if in == nil {
		return nil
	}
	out := new(ApprovalSettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultClientSideAvailabilityObservation) DeepCopyInto(out *DefaultClientSideAvailabilityObservation) {
	*out = *in
	if in.UsingEnvironmentID != nil {
		in, out := &in.UsingEnvironmentID, &out.UsingEnvironmentID
		*out = new(bool)
		**out = **in
	}
	if in.UsingMobileKey != nil {
		in, out := &in.UsingMobileKey, &out.UsingMobileKey
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultClientSideAvailabilityObservation.
func (in *DefaultClientSideAvailabilityObservation) DeepCopy() *DefaultClientSideAvailabilityObservation {
	if in == nil {
		return nil
	}
	out := new(DefaultClientSideAvailabilityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultClientSideAvailabilityParameters) DeepCopyInto(out *DefaultClientSideAvailabilityParameters) {
	*out = *in
	if in.UsingEnvironmentID != nil {
		in, out := &in.UsingEnvironmentID, &out.UsingEnvironmentID
		*out = new(bool)
		**out = **in
	}
	if in.UsingMobileKey != nil {
		in, out := &in.UsingMobileKey, &out.UsingMobileKey
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultClientSideAvailabilityParameters.
func (in *DefaultClientSideAvailabilityParameters) DeepCopy() *DefaultClientSideAvailabilityParameters {
	if in == nil {
		return nil
	}
	out := new(DefaultClientSideAvailabilityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentsObservation) DeepCopyInto(out *EnvironmentsObservation) {
	*out = *in
	if in.ApprovalSettings != nil {
		in, out := &in.ApprovalSettings, &out.ApprovalSettings
		*out = make([]ApprovalSettingsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Color != nil {
		in, out := &in.Color, &out.Color
		*out = new(string)
		**out = **in
	}
	if in.ConfirmChanges != nil {
		in, out := &in.ConfirmChanges, &out.ConfirmChanges
		*out = new(bool)
		**out = **in
	}
	if in.DefaultTTL != nil {
		in, out := &in.DefaultTTL, &out.DefaultTTL
		*out = new(float64)
		**out = **in
	}
	if in.DefaultTrackEvents != nil {
		in, out := &in.DefaultTrackEvents, &out.DefaultTrackEvents
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RequireComments != nil {
		in, out := &in.RequireComments, &out.RequireComments
		*out = new(bool)
		**out = **in
	}
	if in.SecureMode != nil {
		in, out := &in.SecureMode, &out.SecureMode
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentsObservation.
func (in *EnvironmentsObservation) DeepCopy() *EnvironmentsObservation {
	if in == nil {
		return nil
	}
	out := new(EnvironmentsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentsParameters) DeepCopyInto(out *EnvironmentsParameters) {
	*out = *in
	if in.ApprovalSettings != nil {
		in, out := &in.ApprovalSettings, &out.ApprovalSettings
		*out = make([]ApprovalSettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Color != nil {
		in, out := &in.Color, &out.Color
		*out = new(string)
		**out = **in
	}
	if in.ConfirmChanges != nil {
		in, out := &in.ConfirmChanges, &out.ConfirmChanges
		*out = new(bool)
		**out = **in
	}
	if in.DefaultTTL != nil {
		in, out := &in.DefaultTTL, &out.DefaultTTL
		*out = new(float64)
		**out = **in
	}
	if in.DefaultTrackEvents != nil {
		in, out := &in.DefaultTrackEvents, &out.DefaultTrackEvents
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RequireComments != nil {
		in, out := &in.RequireComments, &out.RequireComments
		*out = new(bool)
		**out = **in
	}
	if in.SecureMode != nil {
		in, out := &in.SecureMode, &out.SecureMode
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentsParameters.
func (in *EnvironmentsParameters) DeepCopy() *EnvironmentsParameters {
	if in == nil {
		return nil
	}
	out := new(EnvironmentsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectObservation) DeepCopyInto(out *ProjectObservation) {
	*out = *in
	if in.DefaultClientSideAvailability != nil {
		in, out := &in.DefaultClientSideAvailability, &out.DefaultClientSideAvailability
		*out = make([]DefaultClientSideAvailabilityObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Environments != nil {
		in, out := &in.Environments, &out.Environments
		*out = make([]EnvironmentsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IncludeInSnippet != nil {
		in, out := &in.IncludeInSnippet, &out.IncludeInSnippet
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectObservation.
func (in *ProjectObservation) DeepCopy() *ProjectObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectParameters) DeepCopyInto(out *ProjectParameters) {
	*out = *in
	if in.DefaultClientSideAvailability != nil {
		in, out := &in.DefaultClientSideAvailability, &out.DefaultClientSideAvailability
		*out = make([]DefaultClientSideAvailabilityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Environments != nil {
		in, out := &in.Environments, &out.Environments
		*out = make([]EnvironmentsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludeInSnippet != nil {
		in, out := &in.IncludeInSnippet, &out.IncludeInSnippet
		*out = new(bool)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectParameters.
func (in *ProjectParameters) DeepCopy() *ProjectParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}