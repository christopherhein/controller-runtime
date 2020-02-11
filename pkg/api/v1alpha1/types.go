/*
Copyright 2020 The Kubernetes Authors.

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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DefaultControllerConfigurationSpec defines the desired state of DefaultControllerConfiguration
type DefaultControllerConfigurationSpec struct {
	SyncPeriod *time.Duration `json:"syncPeriod,omitempty"`

	LeaderElection DefaultControllerConfigurationLeaderElection `json:"leaderElection,omitempty"`

	LeaseDuration *time.Duration `json:"leaseDuration,omitempty"`
	RenewDeadline *time.Duration `json:"renewDeadline,omitempty"`
	RetryPeriod   *time.Duration `json:"retryPeriod,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	MetricsBindAddress string `json:"metricsBindAddress,omitempty"`

	Health DefaultControllerConfigurationHealth `json:"health,omitempty"`

	Port *int   `json:"port,omitempty"`
	Host string `json:"host,omitempty"`

	CertDir string `json:"certDir,omitempty"`
}

// DefaultControllerConfigurationLeaderElection defines the leaderelection config
type DefaultControllerConfigurationLeaderElection struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	ID        string `json:"id,omitempty"`
}

// DefaultControllerConfigurationHealth defines the health configs
type DefaultControllerConfigurationHealth struct {
	HealthProbeBindAddress string `json:"healthProbeBindAddress,omitempty"`

	ReadinessEndpointName string `json:"readinessEndpointName,omitempty"`
	LivenessEndpointName  string `json:"livenessEndpointName,omitempty"`
}

// DefaultControllerConfigurationStatus defines the observed state of DefaultControllerConfiguration
type DefaultControllerConfigurationStatus struct {
}

// DefaultControllerConfiguration is the Schema for the DefaultControllerConfigurations API
type DefaultControllerConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	Spec   DefaultControllerConfigurationSpec   `json:"spec,omitempty"`
	Status DefaultControllerConfigurationStatus `json:"status,omitempty"`
}
