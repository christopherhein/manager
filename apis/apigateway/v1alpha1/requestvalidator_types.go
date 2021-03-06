/*
Copyright © 2019 AWS Controller authors

Licensed under the Apache License, Version 2.0 (the &#34;License&#34;);
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an &#34;AS IS&#34; BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RequestValidatorSpec defines the desired state of RequestValidator
type RequestValidatorSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// RestApiRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-restapiid
	RestApiRef metav1alpha1.ObjectReference `json:"restApiRef,omitempty" cloudformation:"RestApiId,Parameter"`

	// ValidateRequestBody http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-validaterequestbody
	ValidateRequestBody bool `json:"validateRequestBody,omitempty" cloudformation:"ValidateRequestBody,Parameter"`

	// ValidateRequestParameters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html#cfn-apigateway-requestvalidator-validaterequestparameters
	ValidateRequestParameters bool `json:"validateRequestParameters,omitempty" cloudformation:"ValidateRequestParameters,Parameter"`
}

// RequestValidatorStatus defines the observed state of RequestValidator
type RequestValidatorStatus struct {
	metav1alpha1.StatusMeta `json:",inline"`
}

// RequestValidatorOutput defines the stack outputs
type RequestValidatorOutput struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html
	Ref string `json:"ref,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=aws;apigateway
// +kubebuilder:printcolumn:JSONPath=.status.status,description="status of the stack",name=Status,priority=0,type=string
// +kubebuilder:printcolumn:JSONPath=.status.message,description="reason for the stack status",name=Message,priority=1,type=string
// +kubebuilder:printcolumn:JSONPath=.status.stackID,description="CloudFormation Stack ID",name=StackID,priority=2,type=string

// RequestValidator is the Schema for the apigateway RequestValidator API
type RequestValidator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RequestValidatorSpec   `json:"spec,omitempty"`
	Status RequestValidatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RequestValidatorList contains a list of Account
type RequestValidatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []RequestValidator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RequestValidator{}, &RequestValidatorList{})
}
