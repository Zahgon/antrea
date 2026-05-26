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

Original file https://raw.githubusercontent.com/kubernetes/kubernetes/refs/tags/v1.34.2/pkg/proxy/serviceport.go

Modifies:

- Remove import apiservice "k8s.io/kubernetes/pkg/api/v1/service" and replace its usages.
- Replace import proxyutil "k8s.io/kubernetes/pkg/proxy/util" with proxyutil "antrea.io/antrea/v2/third_party/proxy/util".
- Modify "UsesLocalEndpoints()" to support DSR implementation in Antrea.

*/

package proxy

import (
	"net"

	v1 "k8s.io/api/core/v1"
)

// ServicePort is an interface which abstracts information about a service.
type ServicePort interface {
	// String returns service string.  An example format can be: `IP:Port/Protocol`.
	String() string
	// ClusterIP returns service cluster IP in net.IP format.
	ClusterIP() net.IP
	// Port returns service port if present. If return 0 means not present.
	Port() int
	// SessionAffinityType returns service session affinity type
	SessionAffinityType() v1.ServiceAffinity
	// StickyMaxAgeSeconds returns service max connection age
	StickyMaxAgeSeconds() int
	// ExternalIPs returns service ExternalIPs
	ExternalIPs() []net.IP
	// LoadBalancerVIPs returns service LoadBalancerIPs which are VIP mode
	LoadBalancerVIPs() []net.IP
	// Protocol returns service protocol.
	Protocol() v1.Protocol
	// LoadBalancerSourceRanges returns service LoadBalancerSourceRanges if present empty array if not
	LoadBalancerSourceRanges() []*net.IPNet
	// HealthCheckNodePort returns service health check node port if present.  If return 0, it means not present.
	HealthCheckNodePort() int
	// NodePort returns a service Node port if present. If return 0, it means not present.
	NodePort() int
	// ExternalPolicyLocal returns if a service has only node local endpoints for external traffic.
	ExternalPolicyLocal() bool
	// InternalPolicyLocal returns if a service has only node local endpoints for internal traffic.
	InternalPolicyLocal() bool
	// ExternallyAccessible returns true if the service port is reachable via something
	// other than ClusterIP (NodePort/ExternalIP/LoadBalancer)
	ExternallyAccessible() bool
	// UsesClusterEndpoints returns true if the service port ever sends traffic to
	// endpoints based on "Cluster" traffic policy
	UsesClusterEndpoints() bool
	// UsesLocalEndpoints returns true if the service port ever sends traffic to
	// endpoints based on "Local" traffic policy
	UsesLocalEndpoints() bool
}

// BaseServicePortInfo contains base information that defines a service.
// This could be used directly by proxier while processing services,
// or can be used for constructing a more specific ServiceInfo struct
// defined by the proxier if needed.
type BaseServicePortInfo struct {
	clusterIP                net.IP
	port                     int
	protocol                 v1.Protocol
	nodePort                 int
	loadBalancerVIPs         []net.IP
	sessionAffinityType      v1.ServiceAffinity
	stickyMaxAgeSeconds      int
	externalIPs              []net.IP
	loadBalancerSourceRanges []*net.IPNet
	healthCheckNodePort      int
	externalPolicyLocal      bool
	internalPolicyLocal      bool
}

var _ ServicePort = &BaseServicePortInfo{}

// String is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) String() string { _ = "STUB: not implemented"; return "" }

// ClusterIP is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) ClusterIP() net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

// Port is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) Port() int { _ = "STUB: not implemented"; return 0 }

// SessionAffinityType is part of the ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) SessionAffinityType() v1.ServiceAffinity {
	_ = "STUB: not implemented"
	return *new(v1.ServiceAffinity)
}

// StickyMaxAgeSeconds is part of the ServicePort interface
func (bsvcPortInfo *BaseServicePortInfo) StickyMaxAgeSeconds() int {
	_ = "STUB: not implemented"
	return 0
}

// Protocol is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) Protocol() v1.Protocol {
	_ = "STUB: not implemented"
	return *new(v1.Protocol)
}

// LoadBalancerSourceRanges is part of ServicePort interface
func (bsvcPortInfo *BaseServicePortInfo) LoadBalancerSourceRanges() []*net.IPNet {
	_ = "STUB: not implemented"
	return nil
}

// HealthCheckNodePort is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) HealthCheckNodePort() int {
	_ = "STUB: not implemented"
	return 0
}

// NodePort is part of the ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) NodePort() int { _ = "STUB: not implemented"; return 0 }

// ExternalIPs is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) ExternalIPs() []net.IP {
	_ = "STUB: not implemented"
	return nil
}

// LoadBalancerVIPs is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) LoadBalancerVIPs() []net.IP {
	_ = "STUB: not implemented"
	return nil
}

// ExternalPolicyLocal is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) ExternalPolicyLocal() bool {
	_ = "STUB: not implemented"
	return false
}

// InternalPolicyLocal is part of ServicePort interface
func (bsvcPortInfo *BaseServicePortInfo) InternalPolicyLocal() bool {
	_ = "STUB: not implemented"
	return false
}

// ExternallyAccessible is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) ExternallyAccessible() bool {
	_ = "STUB: not implemented"
	return false
}

// UsesClusterEndpoints is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) UsesClusterEndpoints() bool {
	_ = "STUB: not implemented"
	// The service port uses Cluster endpoints if the internal traffic policy is "Cluster",
	// or if it accepts external traffic at all. (Even if the external traffic policy is
	// "Local", we need Cluster endpoints to implement short circuiting.)
	return false
}

// UsesLocalEndpoints is part of ServicePort interface.
func (bsvcPortInfo *BaseServicePortInfo) UsesLocalEndpoints() bool {
	_ = "STUB: not implemented"
	// When DSR is enabled, local group could be used by DSR traffic on backend Node.
	// As the Service's own information is not enough to determine its load balancer mode, we just ensure the local group
	// exist as long as it can potentially work in DSR mode.
	return false
}

func NewBaseServiceInfo(service *v1.Service, ipFamily v1.IPFamily, port *v1.ServicePort) *BaseServicePortInfo {
	_ = "STUB: not implemented"
	return nil
}

// Kube-apiserver side guarantees SessionAffinityConfig won't be nil when session affinity type is ClientIP

// Filter ExternalIPs to correct IP family

// Filter source ranges to correct IP family. Also deal with the fact that
// LoadBalancerSourceRanges validation mistakenly allows whitespace padding

// zero-masked cidr means "allow any", which same as the empty loadBalancerSourceRanges.

// Filter Load Balancer Ingress IPs to correct IP family. While proxying load
// balancers might choose to proxy connections from an LB IP of one family to a
// service IP of another family, that's irrelevant to kube-proxy, which only
// creates rules for VIP-style load balancers.

// (already verified as an IP-address)
