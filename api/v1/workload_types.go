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
	Type WorkloadType `json:"type"`

	// k8s raw definitions if type is k8s-object
	K8SObject []runtime.RawExtension `json:"k8sObject"`

	DestAffinity *Affinity `json:"destAffinity"`

	SchedulerName string `json:"schedulerName"`

	CheckFilters []K8SObject `json:"checkFilters"`
}

// WorkloadStatus defines the observed state of Workload
type WorkloadStatus struct {
	Phase              WorkloadPhase       `json:"phase"`
	ObservedGeneration int64               `json:"observedGeneration"`
	Conditions         []WorkloadCondition `json:"conditions"`
	K8SObjectStatus    []K8SObjectStatus   `json:"k8sObjectStatus"`

	AllocatedResources core.ResourceList `json:"allocatedResources"`
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
	WorkloadAffinity     *WorkloadAffinity     `json:"workloadAffinity"`
	WorkloadAntiAffinity *WorkloadAntiAffinity `json:"workloadAntiAffinity"`
}

type WorkloadAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []WorkloadAffinityTerm         `json:"requiredDuringSchedulingIgnoredDuringExecution"`
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedWorkloadAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution"`
}

type WorkloadAntiAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []WorkloadAffinityTerm         `json:"requiredDuringSchedulingIgnoredDuringExecution"`
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedWorkloadAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution"`
}

type WorkloadAffinityTerm struct {
	LabelSelector     *metav1.LabelSelector `json:"labelSelector"`
	Namespaces        []string              `json:"namespaces"`
	TopologyKey       string                `json:"topologyKey"`
	NamespaceSelector *metav1.LabelSelector `json:"namespaceSelector"`
}

type WeightedWorkloadAffinityTerm struct {
	// range 1-100
	Weight               int32                `json:"weight"`
	WorkloadAffinityTerm WorkloadAffinityTerm `json:"workloadAffinityTerm"`
}

type K8SObject struct {
	Group   string `json:"group" protobuf:"bytes,1,opt,name=group"`
	Version string `json:"version" protobuf:"bytes,2,opt,name=version"`
	Kind    string `json:"kind" protobuf:"bytes,3,opt,name=kind"`
	Name    string `json:"name"`
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
	Type WorkloadConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status core.ConditionStatus `json:"status"`
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message"`
}

type K8SObjectStatus struct {
	Reference core.ObjectReference `json:"reference"`

	// Scheduled specifies whether the resource is scheduled.
	Scheduled bool `json:"scheduled"`

	// Ready specifies whether the resource is ready.
	Ready bool `json:"ready"`

	// Deleted
	Deleted bool `json:"deleted"`

	// Replicas
	Replicas int32 `json:"replicas"`

	// ReadyReplicas
	ReadyReplicas int32 `json:"readyReplicas"`
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
