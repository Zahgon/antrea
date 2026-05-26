//go:build windows
// +build windows

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

package route

import (
	"context"
	"net"
	"sync"

	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/servicecidr"
	"antrea.io/antrea/v2/pkg/agent/util"
	"antrea.io/antrea/v2/pkg/agent/util/winfirewall"
	"antrea.io/antrea/v2/pkg/agent/util/winnet"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

const (
	inboundFirewallRuleName  = "Antrea: accept packets from local Pods"
	outboundFirewallRuleName = "Antrea: accept packets to local Pods"

	antreaNatNodePort = "antrea-nat-nodeport"

	serviceIPv4CIDRKey = "serviceIPv4CIDRKey"
)

var (
	antreaNat                  = util.AntreaNatName
	virtualServiceIPv4Net      = util.NewIPNet(config.VirtualServiceIPv4)
	virtualNodePortDNATIPv4Net = util.NewIPNet(config.VirtualNodePortDNATIPv4)
	PodCIDRIPv4                *net.IPNet
)

type Client struct {
	nodeConfig    *config.NodeConfig
	networkConfig *config.NetworkConfig
	winnet        winnet.Interface
	// nodeRoutes caches ip routes to remote Pods. It's a map of podCIDR to routes.
	nodeRoutes sync.Map
	// serviceRoutes caches ip routes about Services.
	serviceRoutes sync.Map
	// serviceExternalIPReferences tracks the references of Service IP. The key is the Service IP and the value is
	// the set of ServiceInfo strings. Because a Service could have multiple ports and each port will generate a
	// ServicePort (which is the unit of the processing), a Service IP route may be required by several ServicePorts.
	// With the references, we install the configurations for a Service IP exactly once as long as it's used by any
	// ServicePorts and uninstall it exactly once when it's no longer used by any ServicePorts.
	// It applies to externalIP and LoadBalancerIP.
	serviceExternalIPReferences map[string]sets.Set[string]
	// netNatStaticMappings caches Windows NetNat for NodePort.
	netNatStaticMappings sync.Map
	fwClient             *winfirewall.Client
	bridgeInfIndex       int
	noSNAT               bool
	proxyAll             bool
	// The latest calculated Service CIDRs can be got from serviceCIDRProvider.
	serviceCIDRProvider servicecidr.Interface
}

// NewClient returns a route client.
// nodeSNATRandomFully and egressSNATRandomFully are ignored on Windows.
func NewClient(networkConfig *config.NetworkConfig,
	noSNAT bool,
	proxyAll bool,
	connectUplinkToBridge bool,
	nodeNetworkPolicyEnabled bool,
	nodeLatencyMonitorEnabled bool,
	multicastEnabled bool,
	egressEnabled bool, // ignored
	nodeSNATRandomFully bool, // ignored
	egressSNATRandomFully bool, // ignored
	serviceCIDRProvider servicecidr.Interface,
	wireguardPort int32,
	proxyHealthCheckPort int32) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize sets nodeConfig on Window.
// Service LoadBalancing is provided by OpenFlow.
func (c *Client) Initialize(nodeConfig *config.NodeConfig, done func()) error {
	_ = "STUB: not implemented"
	return nil
}

// Enable IP-Forwarding on the host gateway interface, thus the host networking stack can be used to forward the
// SNAT packet from local Pods. The SNAT packet is leaving the OVS pipeline with the Node's IP as the source IP,
// the external address as the destination IP, and the antrea-gw0's MAC as the dst MAC. Then it will be forwarded
// to the host network stack from the host gateway interface, and its dst MAC could be resolved to the right one.
// At last, the packet is sent back to OVS from the bridge Interface, and the OpenFlow entries will output it to
// the uplink interface directly.

// For NodePort Service, a NetNatStaticMapping is needed.

func (c *Client) initServiceIPRoutes() error { _ = "STUB: not implemented"; return nil }

// Reconcile removes the orphaned routes and related configuration based on the desired podCIDRs and Service IPs. Only
// the route entries on the host gateway interface are stored in the cache.
func (c *Client) Reconcile(podCIDRs []string) error { _ = "STUB: not implemented"; return nil }

// Don't remove the route entry that does not use global unicast IP address as destination, like multicast, IPv6
// link local or loopback.

// When configuring an IP address to an interface on Windows, three route entries will be automatically added.
// For example, if the IP address is 10.10.0.1/24, the following three routes will be created:
// Network Destination   Netmask          Gateway  Interface  Metric
// 10.10.0.1             255.255.255.255  On-link  10.10.0.1  281
// 10.10.0.0             255.255.255.0    On-link  10.10.0.1  281
// 10.10.0.255           255.255.255.255  On-link  10.10.0.1  281
// The host (10.10.0.1) and broadcast (10.10.0.255) routes should be ignored since they are not supposed to be
// managed by Antrea. We can ignore them by comparing them to the calculated broadcast IP. Don't remove them since
// removing those route entries might introduce host networking issues.

// Don't remove the route entry that uses local Pod CIDR as destination.

// Don't remove the route entry whose destination is included in the desired Pod CIDRs.

// Don't remove the route entries which are added by AntreaProxy when proxyAll is enabled.

// AddRoutes adds routes to the provided podCIDR.
// It overrides the routes if they already exist, without error.
func (c *Client) AddRoutes(podCIDR *net.IPNet, nodeName string, peerNodeIP, peerGwIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// NoEncap traffic to Node on the same subnet.
// Set the peerNodeIP as next hop.

// NoEncap traffic to Node on the different subnet needs underlying routing support.
// Use host default route inside the Node.

// Remove the existing route entry if the gateway address is not as expected.

// DeleteRoutes deletes routes to the provided podCIDR.
// It does nothing if the routes don't exist, without error.
func (c *Client) DeleteRoutes(podCIDR *net.IPNet) error { _ = "STUB: not implemented"; return nil }

// addVirtualServiceIPRoute is used to add a route for a virtual IP. The virtual IP is used as the next hop IP for ClusterIP,
// NodePort and LoadBalancer routes. Without this route, routes for Service cannot be installed on Windows host.
func (c *Client) addVirtualServiceIPRoute(isIPv6 bool) error { _ = "STUB: not implemented"; return nil }

func (c *Client) addServiceCIDRRoute(serviceCIDR *net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate a route with the new ClusterIP CIDR and install it.

// Store the new ClusterIP CIDR and the new generated route to serviceRoutes. Then the calculated route can be restored
// when it was deleted since members of serviceRoutes are synchronized periodically.

// Collect stale routes.

// If current destination CIDR is not nil, the route with current destination CIDR should be uninstalled since
// a new route with a newly calculated destination CIDR has been installed.

// It's the latest route we just installed.

// The route covers the desired route. It was installed when the calculated ServiceCIDR is larger than the current one, which could happen after some Services are deleted.

// The desired route covers the route. It was either installed when the calculated ServiceCIDR is smaller than the current one, or a per-IP route generated before v1.12.0.

// Remove stale routes.

// addVirtualNodePortDNATIPRoute is used to add a route which is used to route DNATed NodePort traffic to Antrea gateway.
func (c *Client) addVirtualNodePortDNATIPRoute(isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// MigrateRoutesToGw is not supported on Windows.
func (c *Client) MigrateRoutesToGw(linkName string) error { _ = "STUB: not implemented"; return nil }

// UnMigrateRoutesFromGw is not supported on Windows.
func (c *Client) UnMigrateRoutesFromGw(route *net.IPNet, linkName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Run periodically syncs netNatStaticMapping and route. It will not return until ctx is cancelled.
func (c *Client) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

// syncNetworkConfig is idempotent and can be safely called on every sync operation.
func (c *Client) syncNetworkConfig(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *Client) syncRoute() error { _ = "STUB: not implemented"; return nil }

// The route is installed automatically by the kernel when the address is configured on the interface. If the route
// is deleted manually by mistake, we restore it.

func (c *Client) syncNetNatStaticMapping() error { _ = "STUB: not implemented"; return nil }

func (c *Client) isServiceRoute(route *winnet.Route) bool {
	_ = "STUB: not implemented"
	// If the gateway IP or the destination IP is the virtual Service IP, then it is a route added by AntreaProxy.
	return false
}

func (c *Client) listIPRoutesOnGW() ([]winnet.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initFwRules adds Windows Firewall rules to accept the traffic that is sent to or from local Pods.
func (c *Client) initFwRules() error { _ = "STUB: not implemented"; return nil }

func (c *Client) AddSNATRule(snatIP net.IP, mark uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteSNATRule(mark uint32) error {
	_ = "STUB: not implemented"

	// TODO: nodePortAddresses is not supported currently.
	return nil
}

func (c *Client) AddNodePortConfigs(nodePortAddresses []net.IP, port uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteNodePortConfigs(nodePortAddresses []net.IP, port uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

// AddExternalIPConfigs adds a route entry to forward traffic destined for the external Service IP to the Antrea
// gateway interface.
func (c *Client) AddExternalIPConfigs(svcInfoStr string, externalIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteExternalIPConfigs deletes the route entry to forward traffic destined for the external Service IP to the Antrea
// gateway interface.
func (c *Client) DeleteExternalIPConfigs(svcInfoStr string, externalIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) AddLocalAntreaFlexibleIPAMPodRule(podAddresses []net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteLocalAntreaFlexibleIPAMPodRule(podAddresses []net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func generateRoute(ipNet *net.IPNet, gw net.IP, linkIndex int, metric int) *winnet.Route {
	_ = "STUB: not implemented"
	return nil
}

func generateNeigh(ip net.IP, linkIndex int) *winnet.Neighbor {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) AddRouteForLink(dstCIDR *net.IPNet, linkIndex int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteRouteForLink(dstCIDR *net.IPNet, linkIndex int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) ClearConntrackEntryForService(svcIP net.IP, svcPort uint16, endpointIP net.IP, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) RestoreEgressRoutesAndRules(minTableID, maxTableID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) AddEgressRoutes(tableID uint32, dev int, gateway net.IP, prefixLength int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteEgressRoutes(tableID uint32) error { _ = "STUB: not implemented"; return nil }

func (c *Client) AddEgressRule(tableID uint32, mark uint32, isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteEgressRule(tableID uint32, mark uint32, isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) AddOrUpdateNodeNetworkPolicyIPSet(ipsetName string, ipsetEntries sets.Set[string], isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteNodeNetworkPolicyIPSet(ipsetName string, isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) AddOrUpdateNodeNetworkPolicyIPTables(iptablesChains []string, iptablesRules [][]string, isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteNodeNetworkPolicyIPTables(iptablesChains []string, isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}
