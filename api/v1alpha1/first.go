package v1alpha1

import (
	"fmt"
	"k8s.io/apimachinery/pkg/types"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (instance *First) GetDeploymentName() types.NamespacedName {
	return types.NamespacedName{
		Name:      instance.GetName() + "-deploy",
		Namespace: instance.GetNamespace(),
	}
}

func (instance *First) GetLabelMap() map[string]string {
	labels := make(map[string]string)
	labels["app"] = "deployment"
	labels["first-deploy"] = "my-deployment"

	return labels
}

func CreateDeployment(instance *First) *appsv1.Deployment {
	fmt.Println("Started created deployment")
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.GetDeploymentName().Name,
			Namespace: instance.GetDeploymentName().Namespace,
			Labels:    instance.GetLabelMap(),
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: instance.GetLabelMap(),
			},
			Replicas: instance.Spec.Member,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   instance.GetName(),
					Labels: instance.GetLabelMap(),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            instance.Spec.Name,
							Image:           instance.Spec.Container.Image,
							ImagePullPolicy: corev1.PullIfNotPresent,
							Command:         []string{"sleep", "20"},
						},
					},
					RestartPolicy: corev1.RestartPolicyAlways,
				},
			},
		},
	}
}
