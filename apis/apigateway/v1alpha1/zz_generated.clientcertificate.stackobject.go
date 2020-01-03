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
	"fmt"
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/apigateway"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *ClientCertificate) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *ClientCertificate) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - apigateway.ClientCertificate (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("ClientCertificate"),
		},
	}

	apigatewayClientCertificate := &apigateway.ClientCertificate{}

	if in.Spec.Description != "" {
		apigatewayClientCertificate.Description = in.Spec.Description
	}

	// TODO(christopherhein): implement tags this could be easy now that I have the mechanims of nested objects

	template.Resources = map[string]cloudformation.Resource{
		"ClientCertificate": apigatewayClientCertificate,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *ClientCertificate) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *ClientCertificate) GenerateStackName() string {
	return strings.Join([]string{"apigateway", "clientcertificate", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *ClientCertificate) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *ClientCertificate) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *ClientCertificate) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *ClientCertificate) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *ClientCertificate) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *ClientCertificate) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *ClientCertificate) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *ClientCertificate) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *ClientCertificate) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *ClientCertificate) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}