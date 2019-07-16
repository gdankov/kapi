package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	NotStartedState = "NotStarted"
	StartedState    = "Started"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Staging struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StagingSpec   `json:"spec"`
	Status StagingStatus `json:"status"`
}

type StagingSpec struct {
	AppGUID            string        `json:"app_guid"`
	CompletionCallback string        `json:"completion_callback"`
	Environment        []EnvVar      `json:"environment"`
	LifecycleData      LifecycleData `json:"lifecycle_data"`
	State              string        `json:"state"`
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type LifecycleData struct {
	AppBitsDownloadURI string      `json:"app_bits_download_uri"`
	DropletUploadURI   string      `json:"droplet_upload_uri"`
	Buildpacks         []Buildpack `json:"buildpacks"`
}

type Buildpack struct {
	Name       string `json:"name"`
	Key        string `json:"key"`
	URL        string `json:"url"`
	SkipDetect bool   `json:"skip_detect"`
}

type StagingStatus struct {
	State string `json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StagingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Staging `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LRP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LRPSpec   `json:"spec"`
	Status LRPStatus `json:"status"`
}

type LRPSpec struct {
	AppGUID string `json:"app_guid"`
}

type LRPStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LRPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LRP `json:"items"`
}
