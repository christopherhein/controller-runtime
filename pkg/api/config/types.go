/*
The Kubernetes Authors.

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

package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

// ControllerConfiguration defines the desired state of GenericControllerConfiguration
type ControllerConfiguration struct {
	// SyncPeriod returns the SyncPeriod
	// +optional
	SyncPeriod *metav1.Duration `json:"syncPeriod,omitempty"`

	// LeaderElection returns the LeaderElection config
	// +optional
	LeaderElection configv1alpha1.LeaderElectionConfiguration `json:"leaderElection,omitempty"`

	// Namespace returns the namespace for the controller
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// MetricsBindAddress returns the bind address for the metrics server
	// +optional
	MetricsBindAddress string `json:"metricsBindAddress,omitempty"`

	// MetricsBindAddress returns the bind address for the metrics server
	// +optional
	Health ControllerHealth `json:"health,omitempty"`

	// Port returns the Port for the server
	// +optional
	Port *int `json:"port,omitempty"`

	// Host returns the Host for the server
	// +optional
	Host string `json:"host,omitempty"`

	// CertDir returns the CertDir
	// +optional
	CertDir string `json:"certDir,omitempty"`
}

// ControllerHealth defines the health configs
type ControllerHealth struct {
	// HealthProbeBindAddress returns the bind address for the health probe
	// +optional
	HealthProbeBindAddress string `json:"healthProbeBindAddress,omitempty"`

	// ReadinessEndpointName returns the readiness endpoint name
	// +optional
	ReadinessEndpointName string `json:"readinessEndpointName,omitempty"`

	// LivenessEndpointName returns the liveness endpoint name
	// +optional
	LivenessEndpointName string `json:"livenessEndpointName,omitempty"`
}

// +kubebuilder:object:root=true

// ControllerRuntimeConfiguration is the Schema for the GenericControllerConfigurations API
type ControllerRuntimeConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// ControllerConfiguration returns the contfigurations for controllers
	ControllerConfiguration `json:",inline"`
}
