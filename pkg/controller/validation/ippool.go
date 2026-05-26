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

package validation

import (
	"net/netip"

	"k8s.io/apimachinery/pkg/util/sets"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

// GetIPRangeSet returns a set of string representations of IP ranges
func GetIPRangeSet(ipRanges []crdv1beta1.IPRange) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// parseIPRangeCIDR parses a CIDR string into a netip.Prefix
func parseIPRangeCIDR(cidrStr string) (netip.Prefix, error) {
	_ = "STUB: not implemented"
	return *new(netip.Prefix), nil
}

// parseIPRangeStartEnd parses start and end IP addresses
func parseIPRangeStartEnd(startStr, endStr string) (netip.Addr, netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), *new(netip.Addr), nil
}

// validateIPRange validates an IP range specification
func validateIPRange(ipRange crdv1beta1.IPRange) error { _ = "STUB: not implemented"; return nil }

// ValidateIPRangesAndSubnetInfo validates IP ranges and SubnetInfo
func ValidateIPRangesAndSubnetInfo(subnetInfo *crdv1beta1.SubnetInfo, ipRanges []crdv1beta1.IPRange) ([]NormalizedIPRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate range is within subnet

// Check for overlaps with other ranges in the same pool

// NormalizedIPRange represents a normalized IP range
type NormalizedIPRange struct {
	Start  netip.Addr
	End    netip.Addr
	Origin string // describes the origin of the range
}

// NormalizeRanges normalizes all IP ranges
func NormalizeRanges(ipRanges []crdv1beta1.IPRange, ctx string) ([]NormalizedIPRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// normalizeRange normalizes an IP range specification
func normalizeRange(ipRange crdv1beta1.IPRange, context string) (NormalizedIPRange, error) {
	_ = "STUB: not implemented"
	return *new(NormalizedIPRange), nil
}

// RangesOverlap checks if two IP ranges overlap
func RangesOverlap(start1, end1, start2, end2 netip.Addr) bool {
	_ = "STUB: not implemented"
	return false
}
