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

package v1alpha1

import (
	"reflect"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

// Getters

// GetSyncPeriod returns the sync period in time.Duration
func (in *ControllerConfiguration) GetSyncPeriod() *time.Duration {
	if in.SyncPeriod != nil {
		return &in.SyncPeriod.Duration
	}
	return nil
}

// GetLeaderElection returns the LeaderElection
func (in *ControllerConfiguration) GetLeaderElection() *bool {
	return in.LeaderElection.LeaderElect
}

// GetLeaderElectionNamespace returns the LeaderElectionNamespace
func (in *ControllerConfiguration) GetLeaderElectionNamespace() string {
	return in.LeaderElection.ResourceNamespace
}

// GetLeaderElectionID returns the LeaderElectionID
func (in *ControllerConfiguration) GetLeaderElectionID() string {
	return in.LeaderElection.ResourceName
}

// GetLeaseDuration returns the LeaseDuration
func (in *ControllerConfiguration) GetLeaseDuration() *time.Duration {
	return &in.LeaderElection.LeaseDuration.Duration
}

// GetRenewDeadline returns the RenewDeadline
func (in *ControllerConfiguration) GetRenewDeadline() *time.Duration {
	return &in.LeaderElection.RenewDeadline.Duration
}

// GetRetryPeriod returns the RetryPeriod
func (in *ControllerConfiguration) GetRetryPeriod() *time.Duration {
	return &in.LeaderElection.RetryPeriod.Duration
}

// GetNamespace returns the Namespace
func (in *ControllerConfiguration) GetNamespace() string {
	return in.Namespace
}

// GetMetricsBindAddress returns the MetricsBindAddress
func (in *ControllerConfiguration) GetMetricsBindAddress() string {
	return in.MetricsBindAddress
}

// GetHealthProbeBindAddress returns the HealthProbeBindAddress
func (in *ControllerConfiguration) GetHealthProbeBindAddress() string {
	return in.Health.HealthProbeBindAddress
}

// GetReadinessEndpointName returns the ReadinessEndpointName
func (in *ControllerConfiguration) GetReadinessEndpointName() string {
	return in.Health.ReadinessEndpointName
}

// GetLivenessEndpointName returns the LivenessEndpointName
func (in *ControllerConfiguration) GetLivenessEndpointName() string {
	return in.Health.LivenessEndpointName
}

// GetPort returns the Port
func (in *ControllerConfiguration) GetPort() *int {
	return in.Port
}

// GetHost returns the Host
func (in *ControllerConfiguration) GetHost() string {
	return in.Host
}

// GetCertDir returns the CertDir
func (in *ControllerConfiguration) GetCertDir() string {
	return in.CertDir
}

// Setters

// SetSyncPeriod sets the sync period in time.Duration
func (in *ControllerConfiguration) SetSyncPeriod(syncPeriod *metav1.Duration) {
	in.SyncPeriod = syncPeriod
}

// SetLeaderElectionConfiguration sets the leader election configuration
func (in *ControllerConfiguration) SetLeaderElectionConfiguration(leaderElection configv1alpha1.LeaderElectionConfiguration) {
	in.LeaderElection = leaderElection
}

// SetLeaderElection sets the LeaderElection config
func (in *ControllerConfiguration) SetLeaderElection(leaderElection bool) {
	if reflect.DeepEqual(in.LeaderElection, configv1alpha1.LeaderElectionConfiguration{}) {
		in.SetLeaderElectionConfiguration(configv1alpha1.LeaderElectionConfiguration{})
	}
	in.LeaderElection.LeaderElect = &leaderElection
}

// SetLeaderElectionNamespace sets the LeaderElectionNamespace
func (in *ControllerConfiguration) SetLeaderElectionNamespace(resourceNamespace string) {
	if reflect.DeepEqual(in.LeaderElection, configv1alpha1.LeaderElectionConfiguration{}) {
		in.SetLeaderElectionConfiguration(configv1alpha1.LeaderElectionConfiguration{})
	}
	in.LeaderElection.ResourceNamespace = resourceNamespace
}

// SetLeaderElectionID sets the LeaderElectionID
func (in *ControllerConfiguration) SetLeaderElectionID(resourceName string) {
	if reflect.DeepEqual(in.LeaderElection, configv1alpha1.LeaderElectionConfiguration{}) {
		in.SetLeaderElectionConfiguration(configv1alpha1.LeaderElectionConfiguration{})
	}
	in.LeaderElection.ResourceName = resourceName
}

// SetLeaseDuration sets the LeaseDuration
func (in *ControllerConfiguration) SetLeaseDuration(leaseDuration metav1.Duration) {
	if reflect.DeepEqual(in.LeaderElection, configv1alpha1.LeaderElectionConfiguration{}) {
		in.SetLeaderElectionConfiguration(configv1alpha1.LeaderElectionConfiguration{})
	}
	in.LeaderElection.LeaseDuration = leaseDuration
}

// SetRenewDeadline sets the RenewDeadline
func (in *ControllerConfiguration) SetRenewDeadline(renewDeadline metav1.Duration) {
	if reflect.DeepEqual(in.LeaderElection, configv1alpha1.LeaderElectionConfiguration{}) {
		in.SetLeaderElectionConfiguration(configv1alpha1.LeaderElectionConfiguration{})
	}
	in.LeaderElection.RenewDeadline = renewDeadline
}

// SetRetryPeriod sets the RetryPeriod
func (in *ControllerConfiguration) SetRetryPeriod(retryPeriod metav1.Duration) {
	if reflect.DeepEqual(in.LeaderElection, configv1alpha1.LeaderElectionConfiguration{}) {
		in.SetLeaderElectionConfiguration(configv1alpha1.LeaderElectionConfiguration{})
	}
	in.LeaderElection.RetryPeriod = retryPeriod
}

// SetNamespace sets the Namespace
func (in *ControllerConfiguration) SetNamespace(namespace string) {
	in.Namespace = namespace
}

// SetMetricsBindAddress sets the MetricsBindAddress
func (in *ControllerConfiguration) SetMetricsBindAddress(metricsBindAddress string) {
	in.MetricsBindAddress = metricsBindAddress
}

// SetHealthProbeBindAddress sets the HealthProbeBindAddress
func (in *ControllerConfiguration) SetHealthProbeBindAddress(healthProbeBindAddress string) {
	in.Health.HealthProbeBindAddress = healthProbeBindAddress
}

// SetReadinessEndpointName sets the ReadinessEndpointName
func (in *ControllerConfiguration) SetReadinessEndpointName(readinessEndpointName string) {
	in.Health.ReadinessEndpointName = readinessEndpointName
}

// SetLivenessEndpointName sets the LivenessEndpointName
func (in *ControllerConfiguration) SetLivenessEndpointName(livenessEndpointName string) {
	in.Health.LivenessEndpointName = livenessEndpointName
}

// SetPort sets the Port
func (in *ControllerConfiguration) SetPort(port *int) {
	in.Port = port
}

// SetHost sets the Host
func (in *ControllerConfiguration) SetHost(host string) {
	in.Host = host
}

// SetCertDir sets the CertDir
func (in *ControllerConfiguration) SetCertDir(certDir string) {
	in.CertDir = certDir
}
