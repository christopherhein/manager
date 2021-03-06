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

// Package generic implements the basic CFN backed controller functions
package generic

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"go.awsctrl.io/manager/meta"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"
	controllerutils "go.awsctrl.io/manager/controllers/utils"
)

// Generic implements the functions for generic CloudFormation controllers
type Generic interface {
	Reconcile(context.Context, logr.Logger, meta.StackObject) (time.Duration, error)
}

type generic struct {
	client.Client
	dynamic.Interface
	Scheme *runtime.Scheme
}

// New configures a new controller
func New(c client.Client, d dynamic.Interface, scheme *runtime.Scheme) (Generic, error) {
	return &generic{
		Client:    c,
		Interface: d,
		Scheme:    scheme,
	}, nil
}

// Reconcile will take in a awsctrl.io/meta.StackObject and create or update a stack
func (r *generic) Reconcile(ctx context.Context, log logr.Logger, instance meta.StackObject) (time.Duration, error) {
	requeueZero := time.Duration(0 * time.Second)
	requeueLater := time.Duration(2 * time.Second)

	// TODO(christopherhein) Add finalizer if termination protection is enabled
	// Need to add a new func which implements reading the field from a code generated function
	// defining the interface in the ./meta/ package and then adding and updating the record

	if instance.GetStackName() == "" {
		log.Info("Create CloudFormation Stack")
		if err := r.addStack(ctx, instance); err != nil {
			return requeueZero, err
		}
		return requeueLater, r.updateTemplateVersion(ctx, instance)
	}

	if controllerutils.IsStatusComplete(instance.GetStatus()) &&
		instance.TemplateVersionChanged() {
		log.Info("Update CloudFormation Stack")
		if err := r.updateStack(ctx, instance); err != nil {
			return requeueZero, err
		}
		return requeueLater, r.updateTemplateVersion(ctx, instance)
	}

	log.Info("Update Status")
	return requeueZero, r.updateStatus(ctx, instance)
}

func (r *generic) addStack(ctx context.Context, instance meta.StackObject) error {
	stack, err := r.newStack(instance)
	if err != nil {
		return err
	}

	if err := r.Create(ctx, stack); err != nil && !apierrs.IsAlreadyExists(err) {
		return err
	}

	instanceCopy := instance.DeepCopyObject().(meta.StackObject)
	instanceCopy.SetStackName(stack.Name)

	return r.Update(ctx, instanceCopy)
}

func (r *generic) updateStack(ctx context.Context, instance meta.StackObject) error {
	var stack cloudformationv1alpha1.Stack
	if err := r.Get(ctx, types.NamespacedName{Name: instance.GetStackName(), Namespace: instance.GetNamespace()}, &stack); err != nil {
		return err
	}

	temp, err := instance.GetTemplate(r.Interface)
	if err != nil {
		return err
	}

	stackCopy := stack.DeepCopy()
	// stackCopy.Spec.Parameters = instance.GetParameters()
	stackCopy.Spec.TemplateBody = temp
	stackCopy.Spec.CloudFormationMeta = instance.GetCloudFormationMeta()

	return r.Update(ctx, stackCopy)
}

func (r *generic) updateStatus(ctx context.Context, instance meta.StackObject) error {
	var stack cloudformationv1alpha1.Stack
	if err := r.Get(ctx, types.NamespacedName{Name: instance.GetStackName(), Namespace: instance.GetNamespace()}, &stack); err != nil {
		return err
	}

	instance.SetStatus(stack.Status.StatusMeta.DeepCopy())
	return r.Status().Update(ctx, instance)
}

func (r *generic) updateTemplateVersion(ctx context.Context, instance meta.StackObject) error {
	nsn := types.NamespacedName{Namespace: instance.GetNamespace(), Name: instance.GetName()}

	var newInstance = instance.DeepCopyObject()
	if err := r.Get(ctx, nsn, newInstance); err != nil {
		return err
	}

	instanceCopy := newInstance.DeepCopyObject().(meta.StackObject)
	instanceCopy.SetTemplateVersionLabel()

	return r.Update(ctx, instanceCopy)
}

func (r *generic) newStack(instance meta.StackObject) (*cloudformationv1alpha1.Stack, error) {
	stack := &cloudformationv1alpha1.Stack{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.GenerateStackName(),
			Namespace: instance.GetNamespace(),
		},
		Spec: cloudformationv1alpha1.StackSpec{
			CloudFormationMeta: instance.GetCloudFormationMeta(),
			Parameters:         map[string]string{},
			// Parameters:         instance.GetParameters(),
		},
	}

	temp, err := instance.GetTemplate(r.Interface)
	if err != nil {
		return stack, err
	}
	stack.Spec.TemplateBody = temp

	if err := ctrl.SetControllerReference(instance, stack, r.Scheme); err != nil {
		return nil, err
	}
	return stack, nil
}
