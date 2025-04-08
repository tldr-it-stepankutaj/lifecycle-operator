package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OperatorUpgradeRequestSpec struct {
	OperatorName string `json:"operatorName"`
	Namespace    string `json:"namespace"`
	TargetCSV    string `json:"targetCSV"`
}

type OperatorUpgradeRequestStatus struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type OperatorUpgradeRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OperatorUpgradeRequestSpec   `json:"spec,omitempty"`
	Status OperatorUpgradeRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type OperatorUpgradeRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OperatorUpgradeRequest `json:"items"`
}
