package e2e_test

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	cloudformationv1alpha1 "go.awsctrl.io/manager/apis/cloudformation/v1alpha1"
	lambdav1alpha1 "go.awsctrl.io/manager/apis/lambda/v1alpha1"

	metav1alpha1 "go.awsctrl.io/manager/apis/meta/v1alpha1"
)

// RunLayerVersionSpecs allows all instance E2E tests to run
var _ = Describe("Run lambda LayerVersion Controller", func() {

	Context("Without LayerVersion{} existing", func() {

		It("Should create lambda.LayerVersion{}", func() {
			Skip("Skip lambda.LayerVersion{}")

			var stackID string
			var stackName string
			var stack *cloudformationv1alpha1.Stack
			k8sclient := k8smanager.GetClient()
			Expect(k8sclient).ToNot(BeNil())

			instance := &lambdav1alpha1.LayerVersion{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: "sample-layerversion-",
					Namespace:    podnamespace,
				},
				Spec: lambdav1alpha1.LayerVersionSpec{},
			}
			By("Creating new lambda LayerVersion")
			Expect(k8sclient.Create(context.Background(), instance)).Should(Succeed())

			key := types.NamespacedName{
				Name:      instance.GetName(),
				Namespace: podnamespace,
			}

			By("Expecting CreateComplete")
			Eventually(func() bool {
				By("Getting latest lambda LayerVersion")
				instance = &lambdav1alpha1.LayerVersion{}
				err := k8sclient.Get(context.Background(), key, instance)
				if err != nil {
					return false
				}

				stackID = instance.GetStackID()
				stackName = instance.GetStackName()

				return instance.Status.Status == metav1alpha1.CreateCompleteStatus ||
					(os.Getenv("USE_AWS_CLIENT") != "true" && instance.Status.Status != "")
			}, timeout, interval).Should(BeTrue())

			By("Checking object OwnerShip")
			Eventually(func() bool {
				stackkey := types.NamespacedName{
					Name:      stackName,
					Namespace: key.Namespace,
				}

				stack = &cloudformationv1alpha1.Stack{}
				err := k8sclient.Get(context.Background(), stackkey, stack)
				if err != nil {
					return false
				}

				expectedOwnerReference := v1.OwnerReference{
					Kind:       instance.Kind,
					APIVersion: instance.APIVersion,
					UID:        instance.UID,
					Name:       instance.Name,
				}

				ownerrefs := stack.GetOwnerReferences()
				Expect(len(ownerrefs)).To(Equal(1))

				return ownerrefs[0].Name == expectedOwnerReference.Name
			}, timeout, interval).Should(BeTrue())

			By("Deleting lambda LayerVersion")
			Expect(k8sclient.Delete(context.Background(), instance)).Should(Succeed())

			By("Deleting LayerVersion Stack")
			Expect(k8sclient.Delete(context.Background(), stack)).Should(Succeed())

			By("Expecting metav1alpha1.DeleteCompleteStatus")
			Eventually(func() bool {
				if os.Getenv("USE_AWS_CLIENT") != "true" {
					return true
				}

				output, err := awsclient.GetClient("us-west-2").DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackID)})
				Expect(err).To(BeNil())
				stackoutput := output.Stacks[0].StackStatus
				return *stackoutput == "DELETE_COMPLETE"
			}, timeout, interval).Should(BeTrue())
		})
	})
})
