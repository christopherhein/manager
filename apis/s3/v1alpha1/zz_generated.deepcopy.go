// +build !ignore_autogenerated

/*
Copyright © 2019 AWS Controller authors

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPoint) DeepCopyInto(out *AccessPoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPoint.
func (in *AccessPoint) DeepCopy() *AccessPoint {
	if in == nil {
		return nil
	}
	out := new(AccessPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessPoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointList) DeepCopyInto(out *AccessPointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessPoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointList.
func (in *AccessPointList) DeepCopy() *AccessPointList {
	if in == nil {
		return nil
	}
	out := new(AccessPointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessPointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointOutput) DeepCopyInto(out *AccessPointOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointOutput.
func (in *AccessPointOutput) DeepCopy() *AccessPointOutput {
	if in == nil {
		return nil
	}
	out := new(AccessPointOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointSpec) DeepCopyInto(out *AccessPointSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.PublicAccessBlockConfiguration = in.PublicAccessBlockConfiguration
	out.VpcConfiguration = in.VpcConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointSpec.
func (in *AccessPointSpec) DeepCopy() *AccessPointSpec {
	if in == nil {
		return nil
	}
	out := new(AccessPointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointStatus) DeepCopyInto(out *AccessPointStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointStatus.
func (in *AccessPointStatus) DeepCopy() *AccessPointStatus {
	if in == nil {
		return nil
	}
	out := new(AccessPointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPoint_PublicAccessBlockConfiguration) DeepCopyInto(out *AccessPoint_PublicAccessBlockConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPoint_PublicAccessBlockConfiguration.
func (in *AccessPoint_PublicAccessBlockConfiguration) DeepCopy() *AccessPoint_PublicAccessBlockConfiguration {
	if in == nil {
		return nil
	}
	out := new(AccessPoint_PublicAccessBlockConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPoint_VpcConfiguration) DeepCopyInto(out *AccessPoint_VpcConfiguration) {
	*out = *in
	out.VpcRef = in.VpcRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPoint_VpcConfiguration.
func (in *AccessPoint_VpcConfiguration) DeepCopy() *AccessPoint_VpcConfiguration {
	if in == nil {
		return nil
	}
	out := new(AccessPoint_VpcConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketList) DeepCopyInto(out *BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketList.
func (in *BucketList) DeepCopy() *BucketList {
	if in == nil {
		return nil
	}
	out := new(BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketOutput) DeepCopyInto(out *BucketOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketOutput.
func (in *BucketOutput) DeepCopy() *BucketOutput {
	if in == nil {
		return nil
	}
	out := new(BucketOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicy) DeepCopyInto(out *BucketPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicy.
func (in *BucketPolicy) DeepCopy() *BucketPolicy {
	if in == nil {
		return nil
	}
	out := new(BucketPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyList) DeepCopyInto(out *BucketPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BucketPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyList.
func (in *BucketPolicyList) DeepCopy() *BucketPolicyList {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyOutput) DeepCopyInto(out *BucketPolicyOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyOutput.
func (in *BucketPolicyOutput) DeepCopy() *BucketPolicyOutput {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicySpec) DeepCopyInto(out *BucketPolicySpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicySpec.
func (in *BucketPolicySpec) DeepCopy() *BucketPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BucketPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketPolicyStatus) DeepCopyInto(out *BucketPolicyStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketPolicyStatus.
func (in *BucketPolicyStatus) DeepCopy() *BucketPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(BucketPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketSpec) DeepCopyInto(out *BucketSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.AccelerateConfiguration = in.AccelerateConfiguration
	if in.AnalyticsConfigurations != nil {
		in, out := &in.AnalyticsConfigurations, &out.AnalyticsConfigurations
		*out = make([]Bucket_AnalyticsConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.BucketEncryption.DeepCopyInto(&out.BucketEncryption)
	in.CorsConfiguration.DeepCopyInto(&out.CorsConfiguration)
	if in.InventoryConfigurations != nil {
		in, out := &in.InventoryConfigurations, &out.InventoryConfigurations
		*out = make([]Bucket_InventoryConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LifecycleConfiguration.DeepCopyInto(&out.LifecycleConfiguration)
	out.LoggingConfiguration = in.LoggingConfiguration
	if in.MetricsConfigurations != nil {
		in, out := &in.MetricsConfigurations, &out.MetricsConfigurations
		*out = make([]Bucket_MetricsConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.NotificationConfiguration.DeepCopyInto(&out.NotificationConfiguration)
	out.ObjectLockConfiguration = in.ObjectLockConfiguration
	out.PublicAccessBlockConfiguration = in.PublicAccessBlockConfiguration
	in.ReplicationConfiguration.DeepCopyInto(&out.ReplicationConfiguration)
	out.VersioningConfiguration = in.VersioningConfiguration
	in.WebsiteConfiguration.DeepCopyInto(&out.WebsiteConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketSpec.
func (in *BucketSpec) DeepCopy() *BucketSpec {
	if in == nil {
		return nil
	}
	out := new(BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketStatus) DeepCopyInto(out *BucketStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketStatus.
func (in *BucketStatus) DeepCopy() *BucketStatus {
	if in == nil {
		return nil
	}
	out := new(BucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_AbortIncompleteMultipartUpload) DeepCopyInto(out *Bucket_AbortIncompleteMultipartUpload) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_AbortIncompleteMultipartUpload.
func (in *Bucket_AbortIncompleteMultipartUpload) DeepCopy() *Bucket_AbortIncompleteMultipartUpload {
	if in == nil {
		return nil
	}
	out := new(Bucket_AbortIncompleteMultipartUpload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_AccelerateConfiguration) DeepCopyInto(out *Bucket_AccelerateConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_AccelerateConfiguration.
func (in *Bucket_AccelerateConfiguration) DeepCopy() *Bucket_AccelerateConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_AccelerateConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_AccessControlTranslation) DeepCopyInto(out *Bucket_AccessControlTranslation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_AccessControlTranslation.
func (in *Bucket_AccessControlTranslation) DeepCopy() *Bucket_AccessControlTranslation {
	if in == nil {
		return nil
	}
	out := new(Bucket_AccessControlTranslation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_AnalyticsConfiguration) DeepCopyInto(out *Bucket_AnalyticsConfiguration) {
	*out = *in
	out.Ref = in.Ref
	out.StorageClassAnalysis = in.StorageClassAnalysis
	if in.TagFilters != nil {
		in, out := &in.TagFilters, &out.TagFilters
		*out = make([]Bucket_TagFilter, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_AnalyticsConfiguration.
func (in *Bucket_AnalyticsConfiguration) DeepCopy() *Bucket_AnalyticsConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_AnalyticsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_BucketEncryption) DeepCopyInto(out *Bucket_BucketEncryption) {
	*out = *in
	if in.ServerSideEncryptionConfiguration != nil {
		in, out := &in.ServerSideEncryptionConfiguration, &out.ServerSideEncryptionConfiguration
		*out = make([]Bucket_ServerSideEncryptionRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_BucketEncryption.
func (in *Bucket_BucketEncryption) DeepCopy() *Bucket_BucketEncryption {
	if in == nil {
		return nil
	}
	out := new(Bucket_BucketEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_CorsConfiguration) DeepCopyInto(out *Bucket_CorsConfiguration) {
	*out = *in
	if in.CorsRules != nil {
		in, out := &in.CorsRules, &out.CorsRules
		*out = make([]Bucket_CorsRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_CorsConfiguration.
func (in *Bucket_CorsConfiguration) DeepCopy() *Bucket_CorsConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_CorsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_CorsRule) DeepCopyInto(out *Bucket_CorsRule) {
	*out = *in
	if in.AllowedHeaders != nil {
		in, out := &in.AllowedHeaders, &out.AllowedHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedMethods != nil {
		in, out := &in.AllowedMethods, &out.AllowedMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedOrigins != nil {
		in, out := &in.AllowedOrigins, &out.AllowedOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExposedHeaders != nil {
		in, out := &in.ExposedHeaders, &out.ExposedHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Ref = in.Ref
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_CorsRule.
func (in *Bucket_CorsRule) DeepCopy() *Bucket_CorsRule {
	if in == nil {
		return nil
	}
	out := new(Bucket_CorsRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_DataExport) DeepCopyInto(out *Bucket_DataExport) {
	*out = *in
	out.Destination = in.Destination
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_DataExport.
func (in *Bucket_DataExport) DeepCopy() *Bucket_DataExport {
	if in == nil {
		return nil
	}
	out := new(Bucket_DataExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_DefaultRetention) DeepCopyInto(out *Bucket_DefaultRetention) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_DefaultRetention.
func (in *Bucket_DefaultRetention) DeepCopy() *Bucket_DefaultRetention {
	if in == nil {
		return nil
	}
	out := new(Bucket_DefaultRetention)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_Destination) DeepCopyInto(out *Bucket_Destination) {
	*out = *in
	out.BucketAccountRef = in.BucketAccountRef
	out.BucketRef = in.BucketRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_Destination.
func (in *Bucket_Destination) DeepCopy() *Bucket_Destination {
	if in == nil {
		return nil
	}
	out := new(Bucket_Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_EncryptionConfiguration) DeepCopyInto(out *Bucket_EncryptionConfiguration) {
	*out = *in
	out.ReplicaKmsKeyRef = in.ReplicaKmsKeyRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_EncryptionConfiguration.
func (in *Bucket_EncryptionConfiguration) DeepCopy() *Bucket_EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_EncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_FilterRule) DeepCopyInto(out *Bucket_FilterRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_FilterRule.
func (in *Bucket_FilterRule) DeepCopy() *Bucket_FilterRule {
	if in == nil {
		return nil
	}
	out := new(Bucket_FilterRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_InventoryConfiguration) DeepCopyInto(out *Bucket_InventoryConfiguration) {
	*out = *in
	out.Destination = in.Destination
	out.Ref = in.Ref
	if in.OptionalFields != nil {
		in, out := &in.OptionalFields, &out.OptionalFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_InventoryConfiguration.
func (in *Bucket_InventoryConfiguration) DeepCopy() *Bucket_InventoryConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_InventoryConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_LambdaConfiguration) DeepCopyInto(out *Bucket_LambdaConfiguration) {
	*out = *in
	in.Filter.DeepCopyInto(&out.Filter)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_LambdaConfiguration.
func (in *Bucket_LambdaConfiguration) DeepCopy() *Bucket_LambdaConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_LambdaConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_LifecycleConfiguration) DeepCopyInto(out *Bucket_LifecycleConfiguration) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Bucket_Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_LifecycleConfiguration.
func (in *Bucket_LifecycleConfiguration) DeepCopy() *Bucket_LifecycleConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_LifecycleConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_LoggingConfiguration) DeepCopyInto(out *Bucket_LoggingConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_LoggingConfiguration.
func (in *Bucket_LoggingConfiguration) DeepCopy() *Bucket_LoggingConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_LoggingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_MetricsConfiguration) DeepCopyInto(out *Bucket_MetricsConfiguration) {
	*out = *in
	out.Ref = in.Ref
	if in.TagFilters != nil {
		in, out := &in.TagFilters, &out.TagFilters
		*out = make([]Bucket_TagFilter, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_MetricsConfiguration.
func (in *Bucket_MetricsConfiguration) DeepCopy() *Bucket_MetricsConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_MetricsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_NoncurrentVersionTransition) DeepCopyInto(out *Bucket_NoncurrentVersionTransition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_NoncurrentVersionTransition.
func (in *Bucket_NoncurrentVersionTransition) DeepCopy() *Bucket_NoncurrentVersionTransition {
	if in == nil {
		return nil
	}
	out := new(Bucket_NoncurrentVersionTransition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_NotificationConfiguration) DeepCopyInto(out *Bucket_NotificationConfiguration) {
	*out = *in
	if in.LambdaConfigurations != nil {
		in, out := &in.LambdaConfigurations, &out.LambdaConfigurations
		*out = make([]Bucket_LambdaConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.QueueConfigurations != nil {
		in, out := &in.QueueConfigurations, &out.QueueConfigurations
		*out = make([]Bucket_QueueConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopicConfigurations != nil {
		in, out := &in.TopicConfigurations, &out.TopicConfigurations
		*out = make([]Bucket_TopicConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_NotificationConfiguration.
func (in *Bucket_NotificationConfiguration) DeepCopy() *Bucket_NotificationConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_NotificationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_NotificationFilter) DeepCopyInto(out *Bucket_NotificationFilter) {
	*out = *in
	in.S3Key.DeepCopyInto(&out.S3Key)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_NotificationFilter.
func (in *Bucket_NotificationFilter) DeepCopy() *Bucket_NotificationFilter {
	if in == nil {
		return nil
	}
	out := new(Bucket_NotificationFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_ObjectLockConfiguration) DeepCopyInto(out *Bucket_ObjectLockConfiguration) {
	*out = *in
	out.Rule = in.Rule
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_ObjectLockConfiguration.
func (in *Bucket_ObjectLockConfiguration) DeepCopy() *Bucket_ObjectLockConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_ObjectLockConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_ObjectLockRule) DeepCopyInto(out *Bucket_ObjectLockRule) {
	*out = *in
	out.DefaultRetention = in.DefaultRetention
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_ObjectLockRule.
func (in *Bucket_ObjectLockRule) DeepCopy() *Bucket_ObjectLockRule {
	if in == nil {
		return nil
	}
	out := new(Bucket_ObjectLockRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_PublicAccessBlockConfiguration) DeepCopyInto(out *Bucket_PublicAccessBlockConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_PublicAccessBlockConfiguration.
func (in *Bucket_PublicAccessBlockConfiguration) DeepCopy() *Bucket_PublicAccessBlockConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_PublicAccessBlockConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_QueueConfiguration) DeepCopyInto(out *Bucket_QueueConfiguration) {
	*out = *in
	in.Filter.DeepCopyInto(&out.Filter)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_QueueConfiguration.
func (in *Bucket_QueueConfiguration) DeepCopy() *Bucket_QueueConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_QueueConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_RedirectAllRequestsTo) DeepCopyInto(out *Bucket_RedirectAllRequestsTo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_RedirectAllRequestsTo.
func (in *Bucket_RedirectAllRequestsTo) DeepCopy() *Bucket_RedirectAllRequestsTo {
	if in == nil {
		return nil
	}
	out := new(Bucket_RedirectAllRequestsTo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_RedirectRule) DeepCopyInto(out *Bucket_RedirectRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_RedirectRule.
func (in *Bucket_RedirectRule) DeepCopy() *Bucket_RedirectRule {
	if in == nil {
		return nil
	}
	out := new(Bucket_RedirectRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_ReplicationConfiguration) DeepCopyInto(out *Bucket_ReplicationConfiguration) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Bucket_ReplicationRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_ReplicationConfiguration.
func (in *Bucket_ReplicationConfiguration) DeepCopy() *Bucket_ReplicationConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_ReplicationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_ReplicationDestination) DeepCopyInto(out *Bucket_ReplicationDestination) {
	*out = *in
	out.AccessControlTranslation = in.AccessControlTranslation
	out.EncryptionConfiguration = in.EncryptionConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_ReplicationDestination.
func (in *Bucket_ReplicationDestination) DeepCopy() *Bucket_ReplicationDestination {
	if in == nil {
		return nil
	}
	out := new(Bucket_ReplicationDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_ReplicationRule) DeepCopyInto(out *Bucket_ReplicationRule) {
	*out = *in
	out.Destination = in.Destination
	out.Ref = in.Ref
	out.SourceSelectionCriteria = in.SourceSelectionCriteria
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_ReplicationRule.
func (in *Bucket_ReplicationRule) DeepCopy() *Bucket_ReplicationRule {
	if in == nil {
		return nil
	}
	out := new(Bucket_ReplicationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_RoutingRule) DeepCopyInto(out *Bucket_RoutingRule) {
	*out = *in
	out.RedirectRule = in.RedirectRule
	out.RoutingRuleCondition = in.RoutingRuleCondition
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_RoutingRule.
func (in *Bucket_RoutingRule) DeepCopy() *Bucket_RoutingRule {
	if in == nil {
		return nil
	}
	out := new(Bucket_RoutingRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_RoutingRuleCondition) DeepCopyInto(out *Bucket_RoutingRuleCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_RoutingRuleCondition.
func (in *Bucket_RoutingRuleCondition) DeepCopy() *Bucket_RoutingRuleCondition {
	if in == nil {
		return nil
	}
	out := new(Bucket_RoutingRuleCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_Rule) DeepCopyInto(out *Bucket_Rule) {
	*out = *in
	out.AbortIncompleteMultipartUpload = in.AbortIncompleteMultipartUpload
	out.Ref = in.Ref
	out.NoncurrentVersionTransition = in.NoncurrentVersionTransition
	if in.NoncurrentVersionTransitions != nil {
		in, out := &in.NoncurrentVersionTransitions, &out.NoncurrentVersionTransitions
		*out = make([]Bucket_NoncurrentVersionTransition, len(*in))
		copy(*out, *in)
	}
	if in.TagFilters != nil {
		in, out := &in.TagFilters, &out.TagFilters
		*out = make([]Bucket_TagFilter, len(*in))
		copy(*out, *in)
	}
	out.Transition = in.Transition
	if in.Transitions != nil {
		in, out := &in.Transitions, &out.Transitions
		*out = make([]Bucket_Transition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_Rule.
func (in *Bucket_Rule) DeepCopy() *Bucket_Rule {
	if in == nil {
		return nil
	}
	out := new(Bucket_Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_S3KeyFilter) DeepCopyInto(out *Bucket_S3KeyFilter) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Bucket_FilterRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_S3KeyFilter.
func (in *Bucket_S3KeyFilter) DeepCopy() *Bucket_S3KeyFilter {
	if in == nil {
		return nil
	}
	out := new(Bucket_S3KeyFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_ServerSideEncryptionByDefault) DeepCopyInto(out *Bucket_ServerSideEncryptionByDefault) {
	*out = *in
	out.KMSMasterKeyRef = in.KMSMasterKeyRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_ServerSideEncryptionByDefault.
func (in *Bucket_ServerSideEncryptionByDefault) DeepCopy() *Bucket_ServerSideEncryptionByDefault {
	if in == nil {
		return nil
	}
	out := new(Bucket_ServerSideEncryptionByDefault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_ServerSideEncryptionRule) DeepCopyInto(out *Bucket_ServerSideEncryptionRule) {
	*out = *in
	out.ServerSideEncryptionByDefault = in.ServerSideEncryptionByDefault
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_ServerSideEncryptionRule.
func (in *Bucket_ServerSideEncryptionRule) DeepCopy() *Bucket_ServerSideEncryptionRule {
	if in == nil {
		return nil
	}
	out := new(Bucket_ServerSideEncryptionRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_SourceSelectionCriteria) DeepCopyInto(out *Bucket_SourceSelectionCriteria) {
	*out = *in
	out.SseKmsEncryptedObjects = in.SseKmsEncryptedObjects
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_SourceSelectionCriteria.
func (in *Bucket_SourceSelectionCriteria) DeepCopy() *Bucket_SourceSelectionCriteria {
	if in == nil {
		return nil
	}
	out := new(Bucket_SourceSelectionCriteria)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_SseKmsEncryptedObjects) DeepCopyInto(out *Bucket_SseKmsEncryptedObjects) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_SseKmsEncryptedObjects.
func (in *Bucket_SseKmsEncryptedObjects) DeepCopy() *Bucket_SseKmsEncryptedObjects {
	if in == nil {
		return nil
	}
	out := new(Bucket_SseKmsEncryptedObjects)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_StorageClassAnalysis) DeepCopyInto(out *Bucket_StorageClassAnalysis) {
	*out = *in
	out.DataExport = in.DataExport
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_StorageClassAnalysis.
func (in *Bucket_StorageClassAnalysis) DeepCopy() *Bucket_StorageClassAnalysis {
	if in == nil {
		return nil
	}
	out := new(Bucket_StorageClassAnalysis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_TagFilter) DeepCopyInto(out *Bucket_TagFilter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_TagFilter.
func (in *Bucket_TagFilter) DeepCopy() *Bucket_TagFilter {
	if in == nil {
		return nil
	}
	out := new(Bucket_TagFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_TopicConfiguration) DeepCopyInto(out *Bucket_TopicConfiguration) {
	*out = *in
	in.Filter.DeepCopyInto(&out.Filter)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_TopicConfiguration.
func (in *Bucket_TopicConfiguration) DeepCopy() *Bucket_TopicConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_TopicConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_Transition) DeepCopyInto(out *Bucket_Transition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_Transition.
func (in *Bucket_Transition) DeepCopy() *Bucket_Transition {
	if in == nil {
		return nil
	}
	out := new(Bucket_Transition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_VersioningConfiguration) DeepCopyInto(out *Bucket_VersioningConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_VersioningConfiguration.
func (in *Bucket_VersioningConfiguration) DeepCopy() *Bucket_VersioningConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_VersioningConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket_WebsiteConfiguration) DeepCopyInto(out *Bucket_WebsiteConfiguration) {
	*out = *in
	out.RedirectAllRequestsTo = in.RedirectAllRequestsTo
	if in.RoutingRules != nil {
		in, out := &in.RoutingRules, &out.RoutingRules
		*out = make([]Bucket_RoutingRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket_WebsiteConfiguration.
func (in *Bucket_WebsiteConfiguration) DeepCopy() *Bucket_WebsiteConfiguration {
	if in == nil {
		return nil
	}
	out := new(Bucket_WebsiteConfiguration)
	in.DeepCopyInto(out)
	return out
}
