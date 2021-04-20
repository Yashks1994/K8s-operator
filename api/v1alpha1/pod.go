package v1alpha1

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "k8s.io/api/core/v1"
)

func NewPod(instance *First) *corev1.Pod {
	fmt.Println("Started created Pod")
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Name,
			Namespace: instance.Namespace,
			Labels:    instance.Labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  instance.Name,
					Image: instance.Spec.Container.Image,
				},
			},
			RestartPolicy: corev1.RestartPolicyOnFailure,
		},
	}
}
