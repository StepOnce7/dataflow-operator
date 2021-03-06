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

package v1

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DataflowEngineSpec defines the desired state of DataflowEngine
type DataflowEngineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image string `json:"image,omitempty"`

	Master ServerMasterSpec `json:"master,omitempty"`

	Executor ServerExecutorSpec `json:"executor,omitempty"`

	FrameStandalone FrameStandaloneSpec `json:"frameStandalone,omitempty"`

	UserStandalone UserStandaloneSpec `json:"userStandalone,omitempty"`
}

type FrameStandaloneSpec struct {
	// mysql-standalone-sample
	Name string `json:"name"`

	Image string `json:"image"`

	Size *int32 `json:"size,omitempty"`

	Platform string `json:"platform,omitempty"`

	Port int32 `json:"port"`

	ClusterTag bool `json:"clusterTag"`

	BackupCommand []string `json:"backupCommand,omitempty"`
}

type UserStandaloneSpec struct {
	// etcd-standalone-sample
	Name string `json:"name"`

	Size *int32 `json:"size,omitempty"`

	Image string `json:"image"`

	Command []string `json:"command,omitempty"`

	Ports []int32 `json:"ports"`

	ClusterTag bool `json:"clusterTag"`
}

type ServerMasterSpec struct {
	Name string `json:"name"`

	Size *int32 `json:"size,omitempty"`

	Command []string `json:"command,omitempty"`

	Ports int32 `json:"ports"`

	ClusterTag bool `json:"clusterTag"`
}

type ServerExecutorSpec struct {
	Name string `json:"name"`

	Size *int32 `json:"size,omitempty"`

	Command []string `json:"command,omitempty"`

	Ports int32 `json:"ports"`

	ClusterTag bool `json:"clusterTag"`
}

// DataflowEngineStatus defines the observed state of DataflowEngine
type DataflowEngineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	MasterNames []string `json:"masterNames"`

	ExecutorNames []string `json:"executorNames"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DataflowEngine is the Schema for the dataflowengines API
type DataflowEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataflowEngineSpec   `json:"spec,omitempty"`
	Status DataflowEngineStatus `json:"status,omitempty"`
}

func (in *DataflowEngine) String() string {

	var res []string

	res = append(res, fmt.Sprintf("Namespance [%s], FrameStandaloneName [%s], FrameStandaloneImage [%s], FrameStandalonePort [%d]",
		in.Namespace,
		in.Spec.FrameStandalone.Name,
		in.Spec.FrameStandalone.Image,
		in.Spec.FrameStandalone.Port))

	res = append(res, fmt.Sprintf("Namespance [%s], UserStandaloneSpecSize [%d], UserStandaloneSpecName [%s], UserStandaloneSpecImage [%s], UserStandaloneSpecPorts [%v]",
		in.Namespace,
		in.Spec.UserStandalone.Size,
		in.Spec.UserStandalone.Name,
		in.Spec.UserStandalone.Image,
		in.Spec.UserStandalone.Ports))

	str := ""
	for _, s := range res {
		str += s
		str += " "
	}
	return str

}

//+kubebuilder:object:root=true

// DataflowEngineList contains a list of DataflowEngine
type DataflowEngineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataflowEngine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataflowEngine{}, &DataflowEngineList{})
}
