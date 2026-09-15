package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type AppDeploymentSpec struct {
	Image    string `json:"image"`
	Replicas int32  `json:"replicas"`
	Port     int32  `json:"port"`
}

type AppDeploymentStatus struct {
}

type AppDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppDeploymentSpec   `json:"spec,omitempty"`
	Status AppDeploymentStatus `json:"status,omitempty"`
}

type AppDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AppDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&AppDeployment{},
		&AppDeploymentList{},
	)
}
