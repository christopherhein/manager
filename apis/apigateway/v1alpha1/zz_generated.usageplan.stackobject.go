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
	"reflect"
	"strings"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
	cfnencoder "go.awsctrl.io/manager/encoding/cloudformation"

	"github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/apigateway"
	"k8s.io/client-go/dynamic"
)

// GetNotificationARNs is an autogenerated deepcopy function, will return notifications for stack
func (in *UsagePlan) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *UsagePlan) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - apigateway.UsagePlan (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("UsagePlan"),
		},
	}

	apigatewayUsagePlan := &apigateway.UsagePlan{}

	if !reflect.DeepEqual(in.Spec.Throttle, UsagePlan_ThrottleSettings{}) {
		apigatewayUsagePlanThrottleSettings := apigateway.UsagePlan_ThrottleSettings{}

		if in.Spec.Throttle.BurstLimit != apigatewayUsagePlanThrottleSettings.BurstLimit {
			apigatewayUsagePlanThrottleSettings.BurstLimit = in.Spec.Throttle.BurstLimit
		}

		if float64(in.Spec.Throttle.RateLimit) != apigatewayUsagePlanThrottleSettings.RateLimit {
			apigatewayUsagePlanThrottleSettings.RateLimit = float64(in.Spec.Throttle.RateLimit)
		}

		apigatewayUsagePlan.Throttle = &apigatewayUsagePlanThrottleSettings
	}

	if in.Spec.UsagePlanName != "" {
		apigatewayUsagePlan.UsagePlanName = in.Spec.UsagePlanName
	}

	apigatewayUsagePlanApiStages := []apigateway.UsagePlan_ApiStage{}

	for _, item := range in.Spec.ApiStages {
		apigatewayUsagePlanApiStage := apigateway.UsagePlan_ApiStage{}

		if !reflect.DeepEqual(item.Throttle, map[string]UsagePlan_ThrottleSettings{}) {
			for key, prop := range item.Throttle {
				apigatewayUsagePlanApiStageThrottleSettings := apigateway.UsagePlan_ThrottleSettings{}
				if prop.BurstLimit != apigatewayUsagePlanApiStageThrottleSettings.BurstLimit {
					apigatewayUsagePlanApiStageThrottleSettings.BurstLimit = prop.BurstLimit
				}

				if float64(prop.RateLimit) != apigatewayUsagePlanApiStageThrottleSettings.RateLimit {
					apigatewayUsagePlanApiStageThrottleSettings.RateLimit = float64(prop.RateLimit)
				}

				apigatewayUsagePlanApiStage.Throttle[key] = apigatewayUsagePlanApiStageThrottleSettings
			}
		}

		// TODO(christopherhein) move these to a defaulter
		apigatewayUsagePlanApiStageApiItem := item.Api.DeepCopy()

		if apigatewayUsagePlanApiStageApiItem.ObjectRef.Kind == "" {
			apigatewayUsagePlanApiStageApiItem.ObjectRef.Kind = "Deployment"
		}

		if apigatewayUsagePlanApiStageApiItem.ObjectRef.APIVersion == "" {
			apigatewayUsagePlanApiStageApiItem.ObjectRef.APIVersion = "apigateway.awsctrl.io/v1alpha1"
		}

		if apigatewayUsagePlanApiStageApiItem.ObjectRef.Namespace == "" {
			apigatewayUsagePlanApiStageApiItem.ObjectRef.Namespace = in.Namespace
		}

		item.Api = *apigatewayUsagePlanApiStageApiItem
		apiId, err := item.Api.String(client)
		if err != nil {
			return "", err
		}

		if apiId != "" {
			apigatewayUsagePlanApiStage.ApiId = apiId
		}

		if item.Stage != "" {
			apigatewayUsagePlanApiStage.Stage = item.Stage
		}

	}

	if len(apigatewayUsagePlanApiStages) > 0 {
		apigatewayUsagePlan.ApiStages = apigatewayUsagePlanApiStages
	}
	if in.Spec.Description != "" {
		apigatewayUsagePlan.Description = in.Spec.Description
	}

	if !reflect.DeepEqual(in.Spec.Quota, UsagePlan_QuotaSettings{}) {
		apigatewayUsagePlanQuotaSettings := apigateway.UsagePlan_QuotaSettings{}

		if in.Spec.Quota.Limit != apigatewayUsagePlanQuotaSettings.Limit {
			apigatewayUsagePlanQuotaSettings.Limit = in.Spec.Quota.Limit
		}

		if in.Spec.Quota.Offset != apigatewayUsagePlanQuotaSettings.Offset {
			apigatewayUsagePlanQuotaSettings.Offset = in.Spec.Quota.Offset
		}

		if in.Spec.Quota.Period != "" {
			apigatewayUsagePlanQuotaSettings.Period = in.Spec.Quota.Period
		}

		apigatewayUsagePlan.Quota = &apigatewayUsagePlanQuotaSettings
	}

	// TODO(christopherhein): implement tags this could be easy now that I have the mechanims of nested objects

	template.Resources = map[string]cloudformation.Resource{
		"UsagePlan": apigatewayUsagePlan,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *UsagePlan) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *UsagePlan) GenerateStackName() string {
	return strings.Join([]string{"apigateway", "usageplan", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *UsagePlan) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *UsagePlan) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *UsagePlan) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *UsagePlan) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *UsagePlan) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *UsagePlan) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *UsagePlan) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *UsagePlan) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *UsagePlan) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *UsagePlan) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
