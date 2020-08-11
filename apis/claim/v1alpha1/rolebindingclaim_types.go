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
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RoleBindingClaimStatus defines the observed state of RoleBindingClaim
type RoleBindingClaimStatus struct {
	Message            string      `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
	Reason             string      `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`

	// +kubebuilder:validation:Enum=Awaiting;Success;Reject;Error
	Status string `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=rbc
// +kubebuilder:rbac:groups="",resources=rolebindings,verbs=*

// RoleBindingClaim is the Schema for the rolebindingclaims API
type RoleBindingClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ResourceName string                 `json:"resourceName"`
	RoleRef      rbacv1.RoleRef         `json:"roleRef" protobuf:"bytes,3,opt,name=roleRef"`
	Subjects     []rbacv1.Subject       `json:"subjects" protobuf:"bytes,2,rep,name=subjects"`
	Status       RoleBindingClaimStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleBindingClaimList contains a list of RoleBindingClaim
type RoleBindingClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleBindingClaim `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RoleBindingClaim{}, &RoleBindingClaimList{})
}
