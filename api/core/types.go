package core

import (
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/types"
)

// ConditionStatus defines conditions of resources
type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ObjectReference struct {
	// +optional
	Kind string `json:"kind"`
	// +optional
	Namespace string `json:"namespace"`
	// +optional
	Name string `json:"name"`
	// +optional
	UID types.UID `json:"uid"`
	// +optional
	APIVersion string `json:"apiVersion"`
	// +optional
	ResourceVersion string `json:"resourceVersion"`

	FieldPath string `json:"fieldPath"`
}

// ResourceName is the name identifying various resources in a ResourceList.
type ResourceName string

const (
	// CPU, in cores. (500m = .5 cores)
	ResourceCPU ResourceName = "cpu"
	// Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	ResourceMemory ResourceName = "memory"
	// Volume size, in bytes (e,g. 5Gi = 5GiB = 5 * 1024 * 1024 * 1024)
	ResourceStorage ResourceName = "storage"
	// Local ephemeral storage, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)
	// The resource name for ResourceEphemeralStorage is alpha and it can change across releases.
	ResourceEphemeralStorage ResourceName = "ephemeral-storage"
)

type ResourceList map[ResourceName]resource.Quantity

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
