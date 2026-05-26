// Copyright 2025 Antrea Authors.
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

package filter

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

var inverseServiceProtocolMap = map[corev1.Protocol]uint8{
	corev1.ProtocolTCP:  6,
	corev1.ProtocolUDP:  17,
	corev1.ProtocolSCTP: 132,
}

// A set of protocols to filter records by
type ProtocolFilter struct {
	protocolNumbers sets.Set[uint8]
}

// For a given protocol, return true if the protocol is allowed
func (p *ProtocolFilter) Allow(protocol uint8) bool { _ = "STUB: not implemented"; return false }

// Returns a new ProtocolFilter with only valid protocols and logs a message
// if invalid or unsupported protocols are found. When protocols is nil, all
// protocols will be allowed. When it is empty, no protocols are allowed.
func NewProtocolFilter(protocols []string) ProtocolFilter {
	_ = "STUB: not implemented"
	return *new(ProtocolFilter)
}
