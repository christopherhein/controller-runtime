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
)

func (in *DefaultControllerConfiguration) GetSyncPeriod() *time.Duration {
	return in.Spec.SyncPeriod
}

func (in *DefaultControllerConfiguration) GetLeaderElection() bool {
	return in.Spec.LeaderElection.Enabled
}
func (in *DefaultControllerConfiguration) GetLeaderElectionNamespace() string {
	return in.Spec.LeaderElection.Namespace
}
func (in *DefaultControllerConfiguration) GetLeaderElectionID() string {
	return in.Spec.LeaderElection.ID
}

func (in *DefaultControllerConfiguration) GetLeaseDuration() *time.Duration {
	return in.Spec.LeaseDuration
}
func (in *DefaultControllerConfiguration) GetRenewDeadline() *time.Duration {
	return in.Spec.RenewDeadline
}
func (in *DefaultControllerConfiguration) GetRetryPeriod() *time.Duration {
	return in.Spec.RetryPeriod
}

func (in *DefaultControllerConfiguration) GetNamespace() string {
	return in.Spec.Namespace
}
func (in *DefaultControllerConfiguration) GetMetricsBindAddress() string {
	return in.Spec.MetricsBindAddress
}
func (in *DefaultControllerConfiguration) GetHealthProbeBindAddress() string {
	return in.Spec.Health.HealthProbeBindAddress
}

func (in *DefaultControllerConfiguration) GetReadinessEndpointName() string {
	return in.Spec.Health.ReadinessEndpointName
}
func (in *DefaultControllerConfiguration) GetLivenessEndpointName() string {
	return in.Spec.Health.LivenessEndpointName
}

func (in *DefaultControllerConfiguration) GetPort() *int {
	return in.Spec.Port
}
func (in *DefaultControllerConfiguration) GetHost() string {
	return in.Spec.Host
}

func (in *DefaultControllerConfiguration) GetCertDir() string {
	return in.Spec.CertDir
}
