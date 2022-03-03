/*
Copyright 2022.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:validation:Enum=create;update;delete;run
type PipelineAction string

const (
	PipelineCreate PipelineAction = "create"
	PipelineUpdate PipelineAction = "update"
	PipelineDelete PipelineAction = "delete"
	PipelineRun    PipelineAction = "run"
)

// PipelineSpec defines the desired state of Pipeline
type PipelineSpec struct {
	// 流水线操作
	Action PipelineAction `json:"action"`
	// 仅测试传入 jenkinsfile 的流水线
	Pipeline *NoScmPipeline `json:"pipeline,omitempty" description:"no scm pipeline structs"`
}

// PipelineStatus defines the observed state of Pipeline
type PipelineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	LastBuild *Build `json:"last_build_number,omitempty"`
}

type Build struct {
	Number string `json:"build_number"`
	Result string `json:"result"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Pipeline is the Schema for the pipelines API
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineSpec   `json:"spec,omitempty"`
	Status PipelineStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PipelineList contains a list of Pipeline
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items"`
}

type NoScmPipeline struct {
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Discarder   *DiscarderProperty `json:"discarder,omitempty"`
	Jenkinsfile string             `json:"jenkinsfile,omitempty"`
}

type DiscarderProperty struct {
	DaysToKeep string `json:"days_to_keep,omitempty"`
	NumToKeep  string `json:"num_to_keep,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Pipeline{}, &PipelineList{})
}
