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

	"sigs.k8s.io/controller-runtime/pkg/api/config"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this GenericControllerConfiguration to the Hub version (v1).
func (src *GenericControllerConfiguration) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*config.ControllerRuntimeConfiguration)

	dst.TypeMeta = src.TypeMeta
	dst.SyncPeriod = src.SyncPeriod
	dst.LeaderElection = src.LeaderElection
	dst.Namespace = src.Namespace
	dst.MetricsBindAddress = src.MetricsBindAddress
	if !reflect.DeepEqual(src.Health, ControllerHealth{}) {
		dst.Health.HealthProbeBindAddress = src.Health.HealthProbeBindAddress
		dst.Health.ReadinessEndpointName = src.Health.ReadinessEndpointName
		dst.Health.LivenessEndpointName = src.Health.LivenessEndpointName
	}
	dst.Port = src.Port
	dst.Host = src.Host
	dst.CertDir = src.CertDir

	return nil
}

// ConvertFrom converts from the Hub version (internal type) to this version.
func (dst *GenericControllerConfiguration) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*config.ControllerRuntimeConfiguration)

	dst.TypeMeta = src.TypeMeta
	dst.SyncPeriod = src.SyncPeriod
	dst.LeaderElection = src.LeaderElection
	dst.Namespace = src.Namespace
	dst.MetricsBindAddress = src.MetricsBindAddress
	if !reflect.DeepEqual(src.Health, ControllerHealth{}) {
		dst.Health.HealthProbeBindAddress = src.Health.HealthProbeBindAddress
		dst.Health.ReadinessEndpointName = src.Health.ReadinessEndpointName
		dst.Health.LivenessEndpointName = src.Health.LivenessEndpointName
	}
	dst.Port = src.Port
	dst.Host = src.Host
	dst.CertDir = src.CertDir

	return nil
}
