/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FirstSpec defines the desired state of First
type FirstSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name
	Name string `json:"name,omitempty"`

	// Member describe the number of values
	Member *int32 `json:"member,omitempty"`

	// Container Image
	Container Container `json:"container,omitempty"`
}

type Container struct {
	Image string `json:"image,omitempty"`
}

const (
	PhasePending = "PENDING"
	PhaseRunning = "RUNNING"
)

// FirstStatus defines the observed state of First
type FirstStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="InstanceName",type="string",JSONPath=".metadata.name"
// +kubebuilder:printcolumn:name="Replica",type="string",JSONPath=".spec.member"

// First is the Schema for the firsts API
type First struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirstSpec   `json:"spec,omitempty"`
	Status FirstStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirstList contains a list of First
type FirstList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []First `json:"items"`
}

func init() {
	SchemeBuilder.Register(&First{}, &FirstList{})
}
