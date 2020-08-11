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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NamespaceClaimSpec defines the desired state of NamespaceClaim
type NamespaceClaimSpec struct {
	Hard v1.ResourceList `json:"hard,omitempty" protobuf:"bytes,1,rep,name=hard,casttype=ResourceList,castkey=ResourceName"`
}

// +kubebuilder:validation:Enum=Awaiting;Success;Reject;Error

// NamespaceClaimStatusType defines namespaceclaim status type
type NamespaceClaimStatusType string

const (
	NamespaceClaimStatusTypeAwaiting NamespaceClaimStatusType = "Awaiting"
	NamespaceClaimStatusTypeSuccess  NamespaceClaimStatusType = "Success"
	NamespaceClaimStatusTypeReject   NamespaceClaimStatusType = "Reject"
	NamespaceClaimStatueTypeError    NamespaceClaimStatusType = "Error"
)

// NamespaceClaimStatus defines the observed state of NamespaceClaim
type NamespaceClaimStatus struct {
	Message            string                   `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
	Reason             string                   `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
	Status             NamespaceClaimStatusType `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
	LastTransitionTime metav1.Time              `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=nsc
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=*

// kubebuilder:subresource:status might be use
// NamespaceClaim is the Schema for the namespaceclaims API
type NamespaceClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceClaimSpec   `json:"spec,omitempty"`
	Status NamespaceClaimStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamespaceClaimList contains a list of NamespaceClaim
type NamespaceClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceClaim `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NamespaceClaim{}, &NamespaceClaimList{})
}
