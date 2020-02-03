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
func (in *Deployment) GetNotificationARNs() []string {
	notifcations := []string{}
	for _, notifarn := range in.Spec.NotificationARNs {
		notifcations = append(notifcations, *notifarn)
	}
	return notifcations
}

// GetTemplate will return the JSON version of the CFN to use.
func (in *Deployment) GetTemplate(client dynamic.Interface) (string, error) {
	if client == nil {
		return "", fmt.Errorf("k8s client not loaded for template")
	}
	template := cloudformation.NewTemplate()

	template.Description = "AWS Controller - apigateway.Deployment (ac-{TODO})"

	template.Outputs = map[string]interface{}{
		"ResourceRef": map[string]interface{}{
			"Value": cloudformation.Ref("Deployment"),
			"Export": map[string]interface{}{
				"Name": in.Name + "Ref",
			},
		},
	}

	apigatewayDeployment := &apigateway.Deployment{}

	if !reflect.DeepEqual(in.Spec.DeploymentCanarySettings, Deployment_DeploymentCanarySettings{}) {
		apigatewayDeploymentDeploymentCanarySettings := apigateway.Deployment_DeploymentCanarySettings{}

		if float64(in.Spec.DeploymentCanarySettings.PercentTraffic) != apigatewayDeploymentDeploymentCanarySettings.PercentTraffic {
			apigatewayDeploymentDeploymentCanarySettings.PercentTraffic = float64(in.Spec.DeploymentCanarySettings.PercentTraffic)
		}

		if !reflect.DeepEqual(in.Spec.DeploymentCanarySettings.StageVariableOverrides, map[string]string{}) {
			apigatewayDeploymentDeploymentCanarySettings.StageVariableOverrides = in.Spec.DeploymentCanarySettings.StageVariableOverrides
		}

		if in.Spec.DeploymentCanarySettings.UseStageCache || !in.Spec.DeploymentCanarySettings.UseStageCache {
			apigatewayDeploymentDeploymentCanarySettings.UseStageCache = in.Spec.DeploymentCanarySettings.UseStageCache
		}

		apigatewayDeployment.DeploymentCanarySettings = &apigatewayDeploymentDeploymentCanarySettings
	}

	if in.Spec.Description != "" {
		apigatewayDeployment.Description = in.Spec.Description
	}

	// TODO(christopherhein) move these to a defaulter
	apigatewayDeploymentRestApiRefItem := in.Spec.RestApiRef.DeepCopy()

	if apigatewayDeploymentRestApiRefItem.ObjectRef.Namespace == "" {
		apigatewayDeploymentRestApiRefItem.ObjectRef.Namespace = in.Namespace
	}

	in.Spec.RestApiRef = *apigatewayDeploymentRestApiRefItem
	restApiId, err := in.Spec.RestApiRef.String(client)
	if err != nil {
		return "", err
	}

	if restApiId != "" {
		apigatewayDeployment.RestApiId = restApiId
	}

	if !reflect.DeepEqual(in.Spec.StageDescription, Deployment_StageDescription{}) {
		apigatewayDeploymentStageDescription := apigateway.Deployment_StageDescription{}

		if !reflect.DeepEqual(in.Spec.StageDescription.AccessLogSetting, Deployment_AccessLogSetting{}) {
			apigatewayDeploymentStageDescriptionAccessLogSetting := apigateway.Deployment_AccessLogSetting{}

			// TODO(christopherhein) move these to a defaulter
			apigatewayDeploymentStageDescriptionAccessLogSettingDestinationRefItem := in.Spec.StageDescription.AccessLogSetting.DestinationRef.DeepCopy()

			if apigatewayDeploymentStageDescriptionAccessLogSettingDestinationRefItem.ObjectRef.Namespace == "" {
				apigatewayDeploymentStageDescriptionAccessLogSettingDestinationRefItem.ObjectRef.Namespace = in.Namespace
			}

			in.Spec.StageDescription.AccessLogSetting.DestinationRef = *apigatewayDeploymentStageDescriptionAccessLogSettingDestinationRefItem
			destinationArn, err := in.Spec.StageDescription.AccessLogSetting.DestinationRef.String(client)
			if err != nil {
				return "", err
			}

			if destinationArn != "" {
				apigatewayDeploymentStageDescriptionAccessLogSetting.DestinationArn = destinationArn
			}

			if in.Spec.StageDescription.AccessLogSetting.Format != "" {
				apigatewayDeploymentStageDescriptionAccessLogSetting.Format = in.Spec.StageDescription.AccessLogSetting.Format
			}

			apigatewayDeploymentStageDescription.AccessLogSetting = &apigatewayDeploymentStageDescriptionAccessLogSetting
		}

		if in.Spec.StageDescription.CacheClusterEnabled || !in.Spec.StageDescription.CacheClusterEnabled {
			apigatewayDeploymentStageDescription.CacheClusterEnabled = in.Spec.StageDescription.CacheClusterEnabled
		}

		if in.Spec.StageDescription.CacheClusterSize != "" {
			apigatewayDeploymentStageDescription.CacheClusterSize = in.Spec.StageDescription.CacheClusterSize
		}

		if in.Spec.StageDescription.CacheDataEncrypted || !in.Spec.StageDescription.CacheDataEncrypted {
			apigatewayDeploymentStageDescription.CacheDataEncrypted = in.Spec.StageDescription.CacheDataEncrypted
		}

		if in.Spec.StageDescription.CacheTtlInSeconds != apigatewayDeploymentStageDescription.CacheTtlInSeconds {
			apigatewayDeploymentStageDescription.CacheTtlInSeconds = in.Spec.StageDescription.CacheTtlInSeconds
		}

		if in.Spec.StageDescription.CachingEnabled || !in.Spec.StageDescription.CachingEnabled {
			apigatewayDeploymentStageDescription.CachingEnabled = in.Spec.StageDescription.CachingEnabled
		}

		if !reflect.DeepEqual(in.Spec.StageDescription.CanarySetting, Deployment_CanarySetting{}) {
			apigatewayDeploymentStageDescriptionCanarySetting := apigateway.Deployment_CanarySetting{}

			if float64(in.Spec.StageDescription.CanarySetting.PercentTraffic) != apigatewayDeploymentStageDescriptionCanarySetting.PercentTraffic {
				apigatewayDeploymentStageDescriptionCanarySetting.PercentTraffic = float64(in.Spec.StageDescription.CanarySetting.PercentTraffic)
			}

			if !reflect.DeepEqual(in.Spec.StageDescription.CanarySetting.StageVariableOverrides, map[string]string{}) {
				apigatewayDeploymentStageDescriptionCanarySetting.StageVariableOverrides = in.Spec.StageDescription.CanarySetting.StageVariableOverrides
			}

			if in.Spec.StageDescription.CanarySetting.UseStageCache || !in.Spec.StageDescription.CanarySetting.UseStageCache {
				apigatewayDeploymentStageDescriptionCanarySetting.UseStageCache = in.Spec.StageDescription.CanarySetting.UseStageCache
			}

			apigatewayDeploymentStageDescription.CanarySetting = &apigatewayDeploymentStageDescriptionCanarySetting
		}

		// TODO(christopherhein) move these to a defaulter
		apigatewayDeploymentStageDescriptionClientCertificateRefItem := in.Spec.StageDescription.ClientCertificateRef.DeepCopy()

		if apigatewayDeploymentStageDescriptionClientCertificateRefItem.ObjectRef.Namespace == "" {
			apigatewayDeploymentStageDescriptionClientCertificateRefItem.ObjectRef.Namespace = in.Namespace
		}

		in.Spec.StageDescription.ClientCertificateRef = *apigatewayDeploymentStageDescriptionClientCertificateRefItem
		clientCertificateId, err := in.Spec.StageDescription.ClientCertificateRef.String(client)
		if err != nil {
			return "", err
		}

		if clientCertificateId != "" {
			apigatewayDeploymentStageDescription.ClientCertificateId = clientCertificateId
		}

		if in.Spec.StageDescription.DataTraceEnabled || !in.Spec.StageDescription.DataTraceEnabled {
			apigatewayDeploymentStageDescription.DataTraceEnabled = in.Spec.StageDescription.DataTraceEnabled
		}

		if in.Spec.StageDescription.Description != "" {
			apigatewayDeploymentStageDescription.Description = in.Spec.StageDescription.Description
		}

		if in.Spec.StageDescription.DocumentationVersion != "" {
			apigatewayDeploymentStageDescription.DocumentationVersion = in.Spec.StageDescription.DocumentationVersion
		}

		if in.Spec.StageDescription.LoggingLevel != "" {
			apigatewayDeploymentStageDescription.LoggingLevel = in.Spec.StageDescription.LoggingLevel
		}

		apigatewayDeploymentStageDescriptionMethodSettings := []apigateway.Deployment_MethodSetting{}

		for _, item := range in.Spec.StageDescription.MethodSettings {
			apigatewayDeploymentStageDescriptionMethodSetting := apigateway.Deployment_MethodSetting{}

			if item.CacheDataEncrypted || !item.CacheDataEncrypted {
				apigatewayDeploymentStageDescriptionMethodSetting.CacheDataEncrypted = item.CacheDataEncrypted
			}

			if item.CacheTtlInSeconds != apigatewayDeploymentStageDescriptionMethodSetting.CacheTtlInSeconds {
				apigatewayDeploymentStageDescriptionMethodSetting.CacheTtlInSeconds = item.CacheTtlInSeconds
			}

			if item.CachingEnabled || !item.CachingEnabled {
				apigatewayDeploymentStageDescriptionMethodSetting.CachingEnabled = item.CachingEnabled
			}

			if item.DataTraceEnabled || !item.DataTraceEnabled {
				apigatewayDeploymentStageDescriptionMethodSetting.DataTraceEnabled = item.DataTraceEnabled
			}

			if item.HttpMethod != "" {
				apigatewayDeploymentStageDescriptionMethodSetting.HttpMethod = item.HttpMethod
			}

			if item.LoggingLevel != "" {
				apigatewayDeploymentStageDescriptionMethodSetting.LoggingLevel = item.LoggingLevel
			}

			if item.MetricsEnabled || !item.MetricsEnabled {
				apigatewayDeploymentStageDescriptionMethodSetting.MetricsEnabled = item.MetricsEnabled
			}

			if item.ResourcePath != "" {
				apigatewayDeploymentStageDescriptionMethodSetting.ResourcePath = item.ResourcePath
			}

			if item.ThrottlingBurstLimit != apigatewayDeploymentStageDescriptionMethodSetting.ThrottlingBurstLimit {
				apigatewayDeploymentStageDescriptionMethodSetting.ThrottlingBurstLimit = item.ThrottlingBurstLimit
			}

			if float64(item.ThrottlingRateLimit) != apigatewayDeploymentStageDescriptionMethodSetting.ThrottlingRateLimit {
				apigatewayDeploymentStageDescriptionMethodSetting.ThrottlingRateLimit = float64(item.ThrottlingRateLimit)
			}

		}

		if len(apigatewayDeploymentStageDescriptionMethodSettings) > 0 {
			apigatewayDeploymentStageDescription.MethodSettings = apigatewayDeploymentStageDescriptionMethodSettings
		}
		if in.Spec.StageDescription.MetricsEnabled || !in.Spec.StageDescription.MetricsEnabled {
			apigatewayDeploymentStageDescription.MetricsEnabled = in.Spec.StageDescription.MetricsEnabled
		}

		// TODO(christopherhein): implement tags this could be easy now that I have the mechanims of nested objects
		if in.Spec.StageDescription.ThrottlingBurstLimit != apigatewayDeploymentStageDescription.ThrottlingBurstLimit {
			apigatewayDeploymentStageDescription.ThrottlingBurstLimit = in.Spec.StageDescription.ThrottlingBurstLimit
		}

		if float64(in.Spec.StageDescription.ThrottlingRateLimit) != apigatewayDeploymentStageDescription.ThrottlingRateLimit {
			apigatewayDeploymentStageDescription.ThrottlingRateLimit = float64(in.Spec.StageDescription.ThrottlingRateLimit)
		}

		if in.Spec.StageDescription.TracingEnabled || !in.Spec.StageDescription.TracingEnabled {
			apigatewayDeploymentStageDescription.TracingEnabled = in.Spec.StageDescription.TracingEnabled
		}

		if !reflect.DeepEqual(in.Spec.StageDescription.Variables, map[string]string{}) {
			apigatewayDeploymentStageDescription.Variables = in.Spec.StageDescription.Variables
		}

		apigatewayDeployment.StageDescription = &apigatewayDeploymentStageDescription
	}

	if in.Spec.StageName != "" {
		apigatewayDeployment.StageName = in.Spec.StageName
	}

	template.Resources = map[string]cloudformation.Resource{
		"Deployment": apigatewayDeployment,
	}

	// json, err := template.JSONWithOptions(&intrinsics.ProcessorOptions{NoEvaluateConditions: true})
	json, err := template.JSON()
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// GetStackID will return stackID
func (in *Deployment) GetStackID() string {
	return in.Status.StackID
}

// GenerateStackName will generate a StackName
func (in *Deployment) GenerateStackName() string {
	return strings.Join([]string{"apigateway", "deployment", in.GetName(), in.GetNamespace()}, "-")
}

// GetStackName will return stackName
func (in *Deployment) GetStackName() string {
	return in.Spec.StackName
}

// GetTemplateVersionLabel will return the stack template version
func (in *Deployment) GetTemplateVersionLabel() (value string, ok bool) {
	value, ok = in.Labels[controllerutils.StackTemplateVersionLabel]
	return
}

// GetParameters will return CFN Parameters
func (in *Deployment) GetParameters() map[string]string {
	params := map[string]string{}
	cfnencoder.MarshalTypes(params, in.Spec, "Parameter")
	return params
}

// GetCloudFormationMeta will return CFN meta object
func (in *Deployment) GetCloudFormationMeta() metav1alpha1.CloudFormationMeta {
	return in.Spec.CloudFormationMeta
}

// GetStatus will return the CFN Status
func (in *Deployment) GetStatus() metav1alpha1.ConditionStatus {
	return in.Status.Status
}

// SetStackID will put a stackID
func (in *Deployment) SetStackID(input string) {
	in.Status.StackID = input
	return
}

// SetStackName will return stackName
func (in *Deployment) SetStackName(input string) {
	in.Spec.StackName = input
	return
}

// SetTemplateVersionLabel will set the template version label
func (in *Deployment) SetTemplateVersionLabel() {
	if len(in.Labels) == 0 {
		in.Labels = map[string]string{}
	}

	in.Labels[controllerutils.StackTemplateVersionLabel] = controllerutils.ComputeHash(in.Spec)
}

// TemplateVersionChanged will return bool if template has changed
func (in *Deployment) TemplateVersionChanged() bool {
	// Ignore bool since it will still record changed
	label, _ := in.GetTemplateVersionLabel()
	return label != controllerutils.ComputeHash(in.Spec)
}

// SetStatus will set status for object
func (in *Deployment) SetStatus(status *metav1alpha1.StatusMeta) {
	in.Status.StatusMeta = *status
}
