package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WatchedOperatorSpec struct {
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	Channel     string `json:"channel"`
	AutoUpgrade bool   `json:"autoUpgrade"`
	WebhookURL  string `json:"webhookURL,omitempty"`
}

type WatchedOperatorStatus struct {
	CurrentVersion string `json:"currentVersion,omitempty"`
	LatestVersion  string `json:"latestVersion,omitempty"`
	UpgradeReady   bool   `json:"upgradeReady,omitempty"`
}

type WatchedOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WatchedOperatorSpec   `json:"spec,omitempty"`
	Status WatchedOperatorStatus `json:"status,omitempty"`
}

type WatchedOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WatchedOperator `json:"items"`
}
