/*
Copyright 2017 The Kubernetes Authors.

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
// Copyright 2020 Antrea Authors
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

Original file https://raw.githubusercontent.com/kubernetes/kubernetes/refs/tags/v1.34.2/pkg/proxy/util/utils.go

Modifies:

- Remove imports:
  - "fmt"
  - "time"
  - utilfeature "k8s.io/apiserver/pkg/util/feature" and its usages.
  - utilsysctl "k8s.io/component-helpers/node/util/sysctl"
  - "k8s.io/kubernetes/pkg/apis/core/v1/helper"
  - "k8s.io/kubernetes/pkg/features"

- Remove consts:
  - IPv4ZeroCIDR
  - IPv6ZeroCIDR
  - FullSyncPeriod

- Remove functions
  - func AddressSet(isValid func(ip net.IP) bool, addrs []net.Addr) sets.Set[string]
  - func OtherIPFamily(ipFamily v1.IPFamily) v1.IPFamily
  - func AppendPortIfNeeded(addr string, port int32) string
  - func EnsureSysctl(sysctl utilsysctl.Interface, name string, newVal int) error
  - func GetClusterIPByFamily(ipFamily v1.IPFamily, service *v1.Service) string

- Modify functions
  - Update `func ShouldSkipService(service *v1.Service) bool` to `func ShouldSkipService(service *v1.Service, skipServices sets.Set[string], serviceLabelSelector labels.Selector) bool`.
  - func IsVIPMode(ing v1.LoadBalancerIngress) bool

*/

package util

import (
	"net"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	netutils "k8s.io/utils/net"
)

// IsZeroCIDR checks whether the input CIDR string is either
// the IPv4 or IPv6 zero CIDR
func IsZeroCIDR(cidr *net.IPNet) bool { _ = "STUB: not implemented"; return false }

// ShouldSkipService checks if a given service should skip proxying
func ShouldSkipService(service *v1.Service, skipServices sets.Set[string], serviceLabelSelector labels.Selector) bool {
	_ = "STUB: not implemented"
	// Skip proxying if the Service label doesn't match the serviceLabelSelector.
	return false
}

// if ClusterIP is "None" or empty, skip proxying

// Even if ClusterIP is set, ServiceTypeExternalName services don't get proxied

// GetClusterIPByFamily returns a service clusterip by family
func GetClusterIPByFamily(ipFamily v1.IPFamily, service *v1.Service) string {
	_ = "STUB: not implemented"
	// allowing skew
	return ""
}

// MapIPsByIPFamily maps a slice of IPs to their respective IP families (v4 or v6)
func MapIPsByIPFamily(ipStrings []string) map[v1.IPFamily][]net.IP {
	_ = "STUB: not implemented"
	return nil
}

// Since ip is parsed ok, GetIPFamilyFromIP will never return v1.IPFamilyUnknown

// ExternalIPs may not be validated by the api-server.
// Specifically empty strings validation, which yields into a lot
// of bad error logs.

// MapCIDRsByIPFamily maps a slice of CIDRs to their respective IP families (v4 or v6)
func MapCIDRsByIPFamily(cidrsStrings []string) map[v1.IPFamily][]*net.IPNet {
	_ = "STUB: not implemented"
	return nil
}

// Ignore empty strings. Same as in MapIPsByIPFamily

// since we just succefully parsed the CIDR, IPFamilyOfCIDR will never return "IPFamilyUnknown"

// GetIPFamilyFromIP Returns the IP family of ipStr, or IPFamilyUnknown if ipStr can't be parsed as an IP
func GetIPFamilyFromIP(ip net.IP) v1.IPFamily { _ = "STUB: not implemented"; return *new(v1.IPFamily) }

// Convert netutils.IPFamily to v1.IPFamily
func convertToV1IPFamily(ipFamily netutils.IPFamily) v1.IPFamily {
	_ = "STUB: not implemented"
	return *new(v1.IPFamily)
}

func IsVIPMode(ing v1.LoadBalancerIngress) bool { _ = "STUB: not implemented"; return false }
