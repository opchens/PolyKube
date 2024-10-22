/*
Copyright 2024.

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

package v1

import (
	"cloudkube/polykube/api/core"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// WorkloadSpec defines the desired state of Workload
type WorkloadSpec struct {
	Type WorkloadType

	// k8s raw definitions if type is k8s-object
	K8SObject []runtime.RawExtension

	DestAffinity *Affinity

	SchedulerName string

	CheckFilters []K8SObject
}

// WorkloadStatus defines the observed state of Workload
type WorkloadStatus struct {
	Phase              WorkloadPhase
	ObservedGeneration int64
	Conditions         []WorkloadCondition
	K8SObjectStatus    []K8SObjectStatus

	AllocatedResources core.ResourceList
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Workload is the Schema for the workloads API
type Workload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkloadSpec   `json:"spec,omitempty"`
	Status WorkloadStatus `json:"status,omitempty"`
}

type WorkloadType string

const (
	K8SOBJECT WorkloadType = "k8s-object"
)

type WorkloadPhase string

const (
	WorkloadPending  WorkloadPhase = "Pending"
	WorkloadCreating WorkloadPhase = "Creating"
	WorkloadRunning  WorkloadPhase = "Running"
	WorkloadUnknown  WorkloadPhase = "Unknown"
)

// Affinity is a group of affinity scheduling rules.
type Affinity struct {
	WorkloadAffinity     *WorkloadAffinity
	WorkloadAntiAffinity *WorkloadAntiAffinity
}

type WorkloadAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []WorkloadAffinityTerm
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedWorkloadAffinityTerm
}

type WorkloadAntiAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []WorkloadAffinityTerm
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedWorkloadAffinityTerm
}

type WorkloadAffinityTerm struct {
	LabelSelector     *metav1.LabelSelector
	Namespaces        []string
	TopologyKey       string
	NamespaceSelector *metav1.LabelSelector
}

type WeightedWorkloadAffinityTerm struct {
	// range 1-100
	Weight               int32
	WorkloadAffinityTerm WorkloadAffinityTerm
}

type K8SObject struct {
	metav1.GroupVersionKind
	Name string
}

type WorkloadConditionType string

// These are valid conditions of a workload.
const (
	WorkloadScheduled WorkloadConditionType = "WorkloadScheduled"
	WorkloadReady     WorkloadConditionType = "Ready"
)

// WorkloadCondition describes the state of a workload at a certain point.
type WorkloadCondition struct {
	// Type of replica set condition.
	Type WorkloadConditionType
	// Status of the condition, one of True, False, Unknown.
	Status core.ConditionStatus
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time
	// The reason for the condition's last transition.
	// +optional
	Reason string
	// A human readable message indicating details about the transition.
	// +optional
	Message string
}

type K8SObjectStatus struct {
	Reference core.ObjectReference

	// Scheduled specifies whether the resource is scheduled.
	Scheduled bool

	// Ready specifies whether the resource is ready.
	Ready bool

	// Deleted
	Deleted bool

	// Replicas
	Replicas int32

	// ReadyReplicas
	ReadyReplicas int32
}

//+kubebuilder:object:root=true

// WorkloadList contains a list of Workload
type WorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workload `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Workload{}, &WorkloadList{})
}
