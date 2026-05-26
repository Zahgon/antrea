/*
Copyright 2016 The Kubernetes Authors.

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
/*
// Copyright 2025 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

Original file https://raw.githubusercontent.com/kubernetes/kubernetes/refs/tags/v1.34.2/pkg/api/v1/service/util.go

Modifies:

- Rename package "service" to "util".

*/

package util

import (
	v1 "k8s.io/api/core/v1"
	utilnet "k8s.io/utils/net"
)

const (
	defaultLoadBalancerSourceRanges = "0.0.0.0/0"
)

// IsAllowAll checks whether the utilnet.IPNet allows traffic from 0.0.0.0/0
func IsAllowAll(ipnets utilnet.IPNetSet) bool { _ = "STUB: not implemented"; return false }

// GetLoadBalancerSourceRanges first try to parse and verify LoadBalancerSourceRanges field from a service.
// If the field is not specified, turn to parse and verify the AnnotationLoadBalancerSourceRangesKey annotation from a service,
// extracting the source ranges to allow, and if not present returns a default (allow-all) value.
func GetLoadBalancerSourceRanges(service *v1.Service) (utilnet.IPNetSet, error) {
	_ = "STUB: not implemented"
	return *new(utilnet.IPNetSet), nil
}

// if SourceRange field is specified, ignore sourceRange annotation

// ExternallyAccessible checks if service is externally accessible.
func ExternallyAccessible(service *v1.Service) bool { _ = "STUB: not implemented"; return false }

// ExternalPolicyLocal checks if service is externally accessible and has ETP = Local.
func ExternalPolicyLocal(service *v1.Service) bool { _ = "STUB: not implemented"; return false }

// InternalPolicyLocal checks if service has ITP = Local.
func InternalPolicyLocal(service *v1.Service) bool { _ = "STUB: not implemented"; return false }

// NeedsHealthCheck checks if service needs health check.
func NeedsHealthCheck(service *v1.Service) bool { _ = "STUB: not implemented"; return false }
