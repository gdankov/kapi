package v1alpha1

import (
	"encoding/json"

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
	Status StagingStatus `json:"status,omitempty"`
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
	State string `json:"state,omitempty"`
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
	GUID                    string                      `json:"guid"`
	Version                 string                      `json:"version"`
	ProcessGUID             string                      `json:"process_guid"`
	Ports                   []int32                     `json:"ports"`
	Routes                  map[string]*json.RawMessage `json:"routes"`
	Environment             map[string]string           `json:"environment"`
	NumInstances            int                         `json:"instances"`
	LastUpdated             string                      `json:"last_updated"`
	HealthCheckType         string                      `json:"health_check_type"`
	HealthCheckHTTPEndpoint string                      `json:"health_check_http_endpoint"`
	HealthCheckTimeoutMs    uint                        `json:"health_check_timeout_ms"`
	MemoryMB                int64                       `json:"memory_mb"`
	CPUWeight               uint8                       `json:"cpu_weight"`
	VolumeMounts            []VolumeMount               `json:"volume_mounts"`
	Lifecycle               Lifecycle                   `json:"lifecycle"`
	DropletHash             string                      `json:"droplet_hash"`
	DropletGUID             string                      `json:"droplet_guid"`
	StartCommand            string                      `json:"start_command"`
	State                   string                      `json:"state"`
}

type Lifecycle struct {
	DockerLifecycle    *DockerLifecycle    `json:"docker_lifecycle"`
	BuildpackLifecycle *BuildpackLifecycle `json:"buildpack_lifecycle"`
}

type DockerLifecycle struct {
	Image   string   `json:"image"`
	Command []string `json:"command"`
}

type BuildpackLifecycle struct {
	DropletHash  string `json:"droplet_hash"`
	DropletGUID  string `json:"droplet_guid"`
	StartCommand string `json:"start_command"`
}

type VolumeMount struct {
	VolumeID string `json:"volume_id"`
	MountDir string `json:"mount_dir"`
}

type LRPStatus struct {
	AvailableReplicas int32  `json:"availableReplicas"`
	State             string `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LRPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LRP `json:"items"`
}
