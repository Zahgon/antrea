// Copyright 2026 Antrea Authors
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

package k8s

import (
	"iter"

	corev1 "k8s.io/api/core/v1"

	"antrea.io/antrea/v2/pkg/agent/nodeportlocal/types"
)

// ipFamilies implements a set of IP families using a bitmask.
// It is an efficient implementation, without any memory allocation.
type ipFamilies int

const (
	ipv4Offset = 4
	ipv6Offset = 6
)

func (ipf ipFamilies) add(ipFamily corev1.IPFamily) ipFamilies {
	_ = "STUB: not implemented"
	return *new(ipFamilies)
}

func (ipf ipFamilies) union(other ipFamilies) ipFamilies {
	_ = "STUB: not implemented"
	return *new(ipFamilies)
}

func (ipf ipFamilies) values() iter.Seq[corev1.IPFamily] { _ = "STUB: not implemented"; return nil }

// getServiceIPFamilies returns the IP families required by a Service.
func getServiceIPFamilies(svc *corev1.Service) []corev1.IPFamily {
	_ = "STUB: not implemented"
	return nil

	// getPodIPForFamily returns the Pod IP matching the specified IP family.
	// Returns empty string if no matching IP is found.
}

func getPodIPForFamily(pod *corev1.Pod, ipFamily corev1.IPFamily) string {
	_ = "STUB: not implemented"
	return ""
}

// ipFamilyForAnnotation converts a corev1.IPFamily to the IPFamilyType used in NPL annotations.
func ipFamilyForAnnotation(ipFamily corev1.IPFamily) types.IPFamilyType {
	_ = "STUB: not implemented"
	return *new(types.IPFamilyType)
}
