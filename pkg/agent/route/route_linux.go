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
	"bytes"
	"context"
	"net"
	"regexp"
	"sync"

	"github.com/vishvananda/netlink"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/knftables"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/servicecidr"
	"antrea.io/antrea/v2/pkg/agent/util/ipset"
	"antrea.io/antrea/v2/pkg/agent/util/iptables"
	utilnetlink "antrea.io/antrea/v2/pkg/agent/util/netlink"
	"antrea.io/antrea/v2/pkg/agent/util/nftables"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

const (
	vxlanPort  = 4789
	genevePort = 6081

	// Antrea managed ipset.
	// antreaPodIPSet contains all Per-Node IPAM Pod CIDRs of this cluster.
	antreaPodIPSet = "ANTREA-POD-IP"
	// antreaPodIP6Set contains all Per-Node IPAM IPv6 Pod CIDRs of this cluster.
	antreaPodIP6Set = "ANTREA-POD-IP6"

	// Antrea managed ipset. Max name length is 31 chars.
	// localAntreaFlexibleIPAMPodIPSet contains all AntreaFlexibleIPAM Pod IPs of this Node.
	localAntreaFlexibleIPAMPodIPSet = "LOCAL-FLEXIBLE-IPAM-POD-IP"
	// localAntreaFlexibleIPAMPodIP6Set contains all AntreaFlexibleIPAM Pod IPv6s of this Node.
	localAntreaFlexibleIPAMPodIP6Set = "LOCAL-FLEXIBLE-IPAM-POD-IP6"
	// clusterNodeIPSet contains all other Node IPs in the cluster.
	clusterNodeIPSet = "CLUSTER-NODE-IP"
	// clusterNodeIP6Set contains all other Node IP6s in the cluster.
	clusterNodeIP6Set = "CLUSTER-NODE-IP6"

	// Antrea managed ipsets for different types of Service IP addresses and ports.
	antreaNodePortIPSet    = "ANTREA-NODEPORT-IP"
	antreaNodePortIP6Set   = "ANTREA-NODEPORT-IP6"
	antreaExternalIPIPSet  = "ANTREA-EXTERNAL-IP"
	antreaExternalIPIP6Set = "ANTREA-EXTERNAL-IP6"

	// Antrea managed iptables chains.
	antreaForwardChain     = "ANTREA-FORWARD"
	antreaPreRoutingChain  = "ANTREA-PREROUTING"
	antreaPostRoutingChain = "ANTREA-POSTROUTING"
	antreaInputChain       = "ANTREA-INPUT"
	antreaOutputChain      = "ANTREA-OUTPUT"

	kubeProxyServiceChain = "KUBE-SERVICES"

	serviceIPv4CIDRKey = "serviceIPv4CIDRKey"
	serviceIPv6CIDRKey = "serviceIPv6CIDRKey"

	preNodeNetworkPolicyIngressRulesChain = "ANTREA-POL-PRE-INGRESS-RULES"
	preNodeNetworkPolicyEgressRulesChain  = "ANTREA-POL-PRE-EGRESS-RULES"

	// All the chains/sets/flowables are created inside our table, so they don't need any "antrea-" prefix of their own.

	// Flowtable and offload chains
	antreaNFTablesFlowtable           = "fastpath"
	antreaNFTablesChainForwardOffload = "forward-offload"
	antreaNFTablesSetPeerPodCIDR      = "peer-pod-cidr"

	// Raw chains for proxyAll
	antreaNFTablesRawChainPreroutingProxyAll = "raw-prerouting-proxy-all"
	antreaNFTablesRawChainOutputProxyAll     = "raw-output-proxy-all"

	// NAT chains for proxyAll
	antreaNFTablesNatChainPreroutingProxyAll  = "nat-prerouting-proxy-all"
	antreaNFTablesNatChainOutputProxyAll      = "nat-output-proxy-all"
	antreaNFTablesNatChainPostroutingProxyAll = "nat-postrouting-proxy-all"

	// NodePort and ExternalIP sets
	antreaNFTablesSetNodePort    = "nodeport"
	antreaNFTablesSetNodePort6   = "nodeport-6"
	antreaNFTablesSetExternalIP  = "externalip"
	antreaNFTablesSetExternalIP6 = "externalip-6"
)

// Client implements Interface.
var _ Interface = &Client{}

var (
	// globalVMAC is used in the IPv6 neighbor configuration to advertise ND solicitation for the IPv6 address of the
	// host gateway interface on other Nodes.
	globalVMAC, _ = net.ParseMAC("aa:bb:cc:dd:ee:ff")

	// The system auto-generated IPv6 link-local route always uses "fe80::/64" as the destination regardless of the
	// interface's global address's mask.
	_, llrCIDR, _ = net.ParseCIDR("fe80::/64")
)

// jumpToAntreaChainPattern matches iptables jump rules that have a comment enclosed in double quotes and target
// Antrea-managed chains.
var jumpToAntreaChainPattern = regexp.MustCompile(`--comment\s+"(Antrea:[^"]+)"\s+-j\s+(ANTREA-[A-Z0-9-]+)`)

// feature identifies a feature that relies on iptables rules. Each component maintains an independent rule
// cache for IPv4 and IPv6.
type feature int

const (
	featureNodeNetworkPolicy feature = iota
	featureWireguard
	featureNodeLatencyMonitor
	featureProxyHealthCheck
)

// iptablesCache stores per-feature iptables state for IPv4 and IPv6. Each feature maintains an independent sync.Map
// for rules/chains.
type iptablesCache struct {
	ipv4 map[feature]*sync.Map
	ipv6 map[feature]*sync.Map
}

// newIPTablesCache initializes an iptablesCache with a sync.Map for each feature and IP family.
func newIPTablesCache() *iptablesCache { _ = "STUB: not implemented"; return nil }

// Client takes care of routing container packets in host network, coordinating ip route, ip rule, iptables and ipset.
type Client struct {
	nodeConfig             *config.NodeConfig
	networkConfig          *config.NetworkConfig
	noSNAT                 bool
	nodeSNATRandomFully    bool
	egressSNATRandomFully  bool
	iptablesHasRandomFully bool
	iptables               iptables.Interface
	nftables               *nftables.Client
	ipset                  ipset.Interface
	netlink                utilnetlink.Interface
	// nodeRoutes caches ip routes to remote Pods. It's a map of podCIDR to routes.
	nodeRoutes sync.Map
	// nodeNeighbors caches IPv6 Neighbors to remote host gateway
	nodeNeighbors sync.Map
	// markToSNATIP caches marks to SNAT IPs. It's used in Egress feature.
	markToSNATIP sync.Map
	// iptablesInitialized is used to notify when iptables initialization is done.
	iptablesInitialized chan struct{}
	proxyAll            bool
	// hostNetworkNFTables is only applicable to proxyAll for now.
	hostNetworkNFTables            bool
	connectUplinkToBridge          bool
	multicastEnabled               bool
	isCloudEKS                     bool
	nodeNetworkPolicyEnabled       bool
	nodeLatencyMonitorEnabled      bool
	egressEnabled                  bool
	hostNetworkAccelerationEnabled bool
	// serviceRoutes caches ip routes about Services.
	serviceRoutes sync.Map
	// serviceExternalIPReferences tracks the references of Service IP. The key is the Service IP and the value is
	// the set of ServiceInfo strings. Because a Service could have multiple ports and each port will generate a
	// ServicePort (which is the unit of the processing), a Service IP route may be required by several ServicePorts.
	// With the references, we install the configurations for a Service IP exactly once as long as it's used by any
	// ServicePorts and uninstall it exactly once when it's no longer used by any ServicePorts.
	// It applies to externalIP and LoadBalancerIP.
	serviceExternalIPReferences map[string]sets.Set[string]
	// serviceNeighbors caches neighbors.
	serviceNeighbors sync.Map
	// serviceIPSets caches ipsets about Services.
	serviceIPSets map[string]*sync.Map
	// serviceIPSets caches nftables sets about Services.
	serviceNFTablesSets map[string]*sync.Map
	// clusterNodeIPs stores the IPv4 of all other Nodes in the cluster
	clusterNodeIPs sync.Map
	// clusterNodeIP6s stores the IPv6 address of all other Nodes in the cluster. It is maintained but not consumed
	// until Multicast supports IPv6.
	clusterNodeIP6s sync.Map
	// egressRoutes caches ip routes about Egresses.
	egressRoutes sync.Map
	// egressRules caches ip rules about Egresses.
	egressRules sync.Map
	// egressNeighbors caches neighbors installed for Egress.
	egressNeighbors sync.Map
	// The latest calculated Service CIDRs can be got from serviceCIDRProvider.
	serviceCIDRProvider servicecidr.Interface
	// nodeNetworkPolicyIPSetsIPv4 caches all existing IPv4 ipsets for NodeNetworkPolicy.
	nodeNetworkPolicyIPSetsIPv4 sync.Map
	// nodeNetworkPolicyIPSetsIPv6 caches all existing IPv6 ipsets for NodeNetworkPolicy.
	nodeNetworkPolicyIPSetsIPv6 sync.Map
	// iptablesCache caches all existing iptables chains and rules for the features relying on it.
	iptablesCache *iptablesCache
	// podCIDRNFTablesSetIPv4 caches all existing IPv4 Pod CIDRs stored in antreaNFTablesSetPeerPodCIDR.
	podCIDRNFTablesSetIPv4 sync.Map
	// podCIDRNFTablesSetIPv6 caches all existing IPv6 Pod CIDRs stored in antreaNFTablesSetPeerPodCIDR.
	podCIDRNFTablesSetIPv6 sync.Map
	// deterministic represents whether to write iptables chains and rules for NodeNetworkPolicy deterministically when
	// syncIPTables is called. Enabling it may carry a performance impact. It's disabled by default and should only be
	// used in testing.
	deterministic bool
	// wireguardPort is the port used for the WireGuard UDP tunnels. When WireGuard is enabled (used as the encryption
	// mode), we add iptables rules to the filter table to accept input and output UDP traffic destined to this port.
	wireguardPort int32
	// proxyHealthCheckPort is the port on which AntreaProxy health check server listens when proxyAll is enabled.
	proxyHealthCheckPort int32
}

// NewClient returns a route client.
func NewClient(networkConfig *config.NetworkConfig,
	noSNAT bool,
	proxyAll bool,
	connectUplinkToBridge bool,
	nodeNetworkPolicyEnabled bool,
	nodeLatencyMonitorEnabled bool,
	multicastEnabled bool,
	egressEnabled bool,
	nodeSNATRandomFully bool,
	egressSNATRandomFully bool,
	serviceCIDRProvider servicecidr.Interface,
	wireguardPort int32,
	proxyHealthCheckPort int32) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize initializes all infrastructures required to route container packets in host network.
// It is idempotent and can be safely called on every startup.
func (c *Client) Initialize(nodeConfig *config.NodeConfig, done func()) error {
	_ = "STUB: not implemented"
	return nil
}

// Sets up the ipset that will be used in iptables.

// Sets up the iptables infrastructure required to route packets in host network.
// It's called in a goroutine because xtables lock may not be acquired immediately.

// ProxyAll depends on nftables; fail if unavailable.

// Host-network acceleration is optional; skip gracefully.

// Sets up the IP routes and IP rule required to route packets in host network.

// Ensure IPv4 forwarding is enabled if it is a dual-stack or IPv4-only cluster.

// Ensure IPv6 forwarding is enabled if it is a dual-stack or IPv6-only cluster.

// In hybrid mode, or encap mode with WireGuard enabled, Egress traffic originating from remote Pod CIDRs is forwarded like this:
//     remote Pods -> tunnel (remote Node OVS) -> tunnel (local Node OVS) -> antrea-gw0 (local Node) -> external network.
//
// To ensure reply packets follow a symmetric path, Antrea uses policy routing on the local Node. However, the
// kernel's strict RPF check only validates source paths against the main routing table. Since the transport
// interface (not antrea‑gw0) is listed as the next-hop for these routes, strict RPF drops the reply packets
// (because policy routing is ignored by rp_filter). As a result, we set its rp_filter to loose mode (2).

// Set up the IP routes and sysctl parameters to support all Services in AntreaProxy.

// Set up the policy routing ip rule to support Egress in hybrid mode or encap mode with WireGuard enabled.

// Build static iptables rules for NodeNetworkPolicy.

// Run waits for iptables initialization, then periodically syncs ipsets, iptables/nftables, routes, neighbors, and
// policy rules. It will not return until ctx is cancelled.
func (c *Client) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

// syncNetworkConfig is idempotent and can be safely called on every sync operation.
func (c *Client) syncNetworkConfig(ctx context.Context) {
	_ = "STUB: not implemented"
	// Sync ipset before syncing iptables rules
	return
}

type routeKey struct {
	linkIndex int
	dst       string
	gw        string
	tableID   int
}

func (c *Client) syncRoute() error { _ = "STUB: not implemented"; return nil }

// These routes are installed automatically by the kernel when the address is configured on
// the interface (with "proto kernel"). If these routes are deleted manually by mistake, we
// restore them as part of this sync (without "proto kernel"). An alternative would be to
// flap the interface, but this seems like a better approach.

// Here we assume the IPv6 link-local address always exists on antrea-gw0
// to avoid unexpected issues in the IPv6 forwarding.

// Restore the IPv6 link-local route.

type neighborKey struct {
	ip  string
	mac string
}

// syncNeighbor ensures that necessary neighbors exist on the Antrea gateway interface, as some routes managed by Antrea
// depend on these neighbors.
func (c *Client) syncNeighbor() error { _ = "STUB: not implemented"; return nil }

type ruleKey struct {
	table  int
	mark   uint32
	mask   uint32
	family int
}

func generateRuleKey(rule *netlink.Rule) ruleKey { _ = "STUB: not implemented"; return *new(ruleKey) }

// syncIPRule ensures that the necessary ip rules required by Antrea exist.
func (c *Client) syncIPRule() error { _ = "STUB: not implemented"; return nil }

// syncIPSet ensures that the required ipset exists, and it has the initial members.
func (c *Client) syncIPSet() error {
	_ = "STUB: not implemented"
	// Create the ipsets to store all Pod CIDRs for constructing full-mesh routing in encap/noEncap/hybrid modes. In
	// networkPolicyOnly mode, Antrea is not responsible for IPAM, so CIDRs are not available and the ipsets should not
	// be created.
	return nil
}

// Loop all valid Pod CIDRs and add them into the corresponding ipset.

// AntreaProxy proxyAll is available in all traffic modes. If proxyAll is enabled, create the ipsets to store the
// pairs of Node IP and NodePort.

// AntreaIPAM is available in noEncap mode. There is a validation in Antrea configuration about this traffic mode
// when AntreaIPAM is enabled.

// Multicast is available in encap/noEncap/hybrid mode, and the ipsets are consumed in encap mode.

// NodeNetworkPolicy is available in all traffic modes.

func getIPSetName(ip net.IP) string { _ = "STUB: not implemented"; return "" }

func getNodePortIPSetName(isIPv6 bool) string { _ = "STUB: not implemented"; return "" }

func getExternalIPIPSetName(isIPv6 bool) string { _ = "STUB: not implemented"; return "" }

func getNodePortNFTablesSet(isIPv6 bool) string { _ = "STUB: not implemented"; return "" }

func getExternalIPNFTablesSet(isIPv6 bool) string { _ = "STUB: not implemented"; return "" }

func getLocalAntreaFlexibleIPAMPodIPSetName(isIPv6 bool) string {
	_ = "STUB: not implemented"
	return ""
}

// writeEKSMangleRules writes an additional iptables mangle rule to the
// iptablesData buffer, to set the traffic mark to the connection mark for
// traffic coming out of the gateway. This rule is needed for 2 cases:
//   - for the reverse path for NodePort Service traffic (see
//     https://github.com/antrea-io/antrea/issues/678).
//   - for Pod-to-external traffic that needs to be SNATed (see
//     https://github.com/antrea-io/antrea/issues/3946).
func (c *Client) writeEKSMangleRules(iptablesData *bytes.Buffer) {
	_ = "STUB: not implemented"
	// TODO: the following should be taking into account:
	//  1. this rule is only needed if AWS_VPC_CNI_NODE_PORT_SUPPORT is set
	//     to true (which is the default) or if AWS_VPC_K8S_CNI_EXTERNALSNAT
	//     is set to false (which is also the default). If both settings are
	//     changed, we do not need to install the rule.
	//  2. this option is not documented but the mark value can be
	//     configured with AWS_VPC_K8S_CNI_CONNMARK.
	//
	// While we do not have access to these environment variables, we could
	// look for existing rules installed by the AWS VPC CNI, and determine
	// whether we need to install this rule.
	return
}

// writeEKSNATRules writes additional iptables nat rules to the iptablesData
// buffer. The first rule sets the connection mark for Pod-to-external traffic
// that needs to be SNATed. Without the mark, this traffic is not routed using
// the correct route table for Pods getting an IP address from a secondary
// network interface (ENI). The second rule restores the packet mark from the
// connection mark. That rule will only apply to the first packet of the
// connection. For subsequent packets, the rule installed buy writeEKSMangleRule
// will take care of restoring the mark. We need that rule because the mangle
// table is traversed before the nat table.
// See https://docs.aws.amazon.com/eks/latest/userguide/external-snat.html and
// https://github.com/antrea-io/antrea/issues/3946 for more details.
func (c *Client) writeEKSNATRules(iptablesData *bytes.Buffer) {
	_ = "STUB: not implemented"
	// TODO: just like for writeEKSMangleRule, these rules should ideally be
	// installed conditionally, when AWS_VPC_K8S_CNI_EXTERNALSNAT is set to
	// false (which is the default value).
	return
}

// The AWS VPC CNI already installs the same rule in the PREROUTING
// chain. However, that rule will typically be visited before the
// ANTREA-PREROUTING chain, hence we need to install our own copy of the
// rule.

func (c *Client) getIPProtocol() iptables.Protocol {
	_ = "STUB: not implemented"
	return *new(iptables.Protocol)
}

// Create the antrea managed chains and link them to built-in chains.
// We cannot use iptables-restore for these jump rules because there
// are non antrea managed rules in built-in chains.
type jumpRule struct {
	table    string
	srcChain string
	dstChain string
	comment  string
	insert   bool
}

type jumpRuleKey struct {
	table    string
	srcChain string
	dstChain string
	comment  string
}

func (j *jumpRule) getKey() jumpRuleKey { _ = "STUB: not implemented"; return *new(jumpRuleKey) }

func (c *Client) removeUnexpectedAntreaJumpRule(protocol iptables.Protocol, jumpRule jumpRule) error {
	_ = "STUB: not implemented"
	// List all the existing rules of the table and the chain where the Antrea jump rule will be added.
	return nil
}

// Construct keywords to identify Antrea and kube-proxy jump rules.

// Check if the current rule is the Antrea jump rule to be added.

// Check if the current rule is the kube-proxy jump rule.

// If the Antrea jump rule is installed after the kube-proxy jump rule, which is not expected, delete the
// existing Antrea jump rule to ensure that a new one will be installed before the kube-proxy one when syncing iptables.

// syncIPTables ensure that the iptables infrastructure we use is set up.
// It's idempotent and can safely be called on every startup.
func (c *Client) syncIPTables(cleanupStaleJumpRules bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Install the static rules (WireGuard + NodeLatencyMonitor) before the dynamic rules (e.g., NodeNetworkPolicy)
// for performance reasons.

// Use iptables-restore to configure IPv4 settings.

// Setting --noflush to keep the previous contents (i.e. non antrea managed chains) of the tables.

// Use ip6tables-restore to configure IPv6 settings.

// Setting --noflush to keep the previous contents (i.e. non antrea managed chains) of the tables.

func (c *Client) getJumpRuleKeys(ipProtocol iptables.Protocol) sets.Set[jumpRuleKey] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) restoreIptablesData(podCIDR *net.IPNet,
	podIPSet,
	localAntreaFlexibleIPAMPodIPSet,
	nodePortIPSet,
	externalIPSet,
	clusterNodeIPSet string,
	nodePortDNATVirtualIP,
	serviceVirtualIP net.IP,
	snatMarkToIP map[uint32]net.IP,
	iptablesFiltersRuleByChain map[string][]string,
	isIPv6 bool) *bytes.Buffer {
	_ = "STUB: not implemented"
	// Create required rules in the antrea chains.
	// Use iptables-restore as it flushes the involved chains and creates the desired rules
	// with a single call, instead of string matching to clean up stale rules.
	return nil
}

// Write head lines anyway so the undesired rules can be deleted when changing encap mode.

// For Geneve and VXLAN encapsulation packets, the request and response packets don't belong to a UDP connection
// so tracking them doesn't give the normal benefits of conntrack. Besides, kube-proxy may install great number
// of iptables rules in nat table. The first encapsulation packets of connections would have to go through all
// of the rules which wastes CPU and increases packet latency.

// If a tunnel port is specified, use it instead of the default port.

// Note: Multicast can only work with IPv4 for now. Remove condition "!isIPv6" in the future after
// IPv6 is supported.

// Drop the multicast packets forwarded from other Nodes in the cluster. This is because
// the packet sent out from the sender Pod is already received via tunnel port with encap mode,
// and the one forwarded via the underlay network is to send to external receivers

// This rule is to bypass conntrack for packets sourced from external and destined to externalIPs, which also
// results in bypassing the chains managed by Antrea Proxy and kube-proxy in nat table.

// This rule is to bypass conntrack for packets sourced from externalIPs, which also results in bypassing the
// chains managed by Antrea Proxy and kube-proxy in nat table.

// This rule is to bypass conntrack for packets sourced from local and destined to externalIPs, which also
// results in bypassing the chains managed by Antrea Proxy and kube-proxy in nat table.

// Write head lines anyway so the undesired rules can be deleted when noEncap -> encap.

// When Antrea is used to enforce NetworkPolicies in EKS, additional iptables
// mangle rules are required. See https://github.com/antrea-io/antrea/issues/678.
// These rules are only needed for IPv4.

// Match packets from established connections.
// Match reply packets.
// Match packets with conntrack mark.

// Restore fwmark from the conntrack mark.

// Don't match packets from local Pods.
// Match the first packet.
// Match packets whose fwmark value (0–7) is non-zero. In the OVS pipeline, this mark is set for Egress connections.

// Load EgressNoEncapReturnToRemoteMark into the conntrack mark so it persists for the lifetime of the Egress connection.

// Match packets from established connections.
// Match reply packets.
// Clear the fwmark EgressNoEncapReturnToRemoteMark from the packet to avoid that the fwmark mark may
// interfere with source IP selection when the packets are encapsulated by OVS flow-based tunnel.

// To make liveness/readiness probe traffic bypass ingress rules of Network Policies, mark locally generated packets
// that will be sent to OVS so we can identify them later in the OVS pipeline.
// It must match source address because kube-proxy ipvs mode will redirect ingress packets to output chain, and they
// will have non local source addresses.

// Match packets from established connections.
// Match reply packets.
// Match packets with conntrack mark.

// Restore fwmark from the conntrack mark.

// Add accept rules for local AntreaFlexibleIPAM
// AntreaFlexibleIPAM Pods -> HostPort Pod
// AntreaFlexibleIPAM Pods -> NodePort Service -> Backend Pod

// The masqueraded multicast traffic will become unicast so we
// stop traversing this antreaPostRoutingChain for multicast traffic.
// Note: Multicast can only work with IPv4 for now. Remove condition "!isIPv6" in the future after
// IPv6 is supported.

// Egress rules must be inserted before the default masquerade rule.

// Cannot reuse snatRuleSpec to generate the rule as it doesn't have "`" in the comment.

// For local traffic going out of the gateway interface, if the source IP does not match any
// of the gateway's IP addresses, the traffic needs to be masqueraded. Otherwise, we observe
// that ARP requests may advertise a different source IP address, in which case they will be
// dropped by the SpoofGuard table in the OVS pipeline. See description for the arp_announce
// sysctl parameter.

// If AntreaProxy full support is enabled, it SNATs the packets whose source IP is VirtualServiceIPv4/VirtualServiceIPv6
// so the packets can be routed back to this Node.

// This generates the rule to masquerade the packets destined for a hostPort whose backend is an AntreaIPAM VLAN Pod.
// For simplicity, in the following descriptions:
//   - per-Node IPAM Pod is referred to as regular Pod.
//   - AntreaIPAM Pod without VLAN is referred to as AntreaIPAM Pod.
//   - AntreaIPAM Pod with VLAN is referred to as AntreaIPAM VLAN Pod.
// The common conditions are:
//   - AntreaIPAM VLAN Pod exposes hostPort.
//   - hostPort traffic is sent to underlay gateway after Node DNATed the traffic.
//   - underlay gateway sends traffic back with a vlan tag.
//   - SNAT is required to guarantee the reply traffic can be sent to Node to de-DNAT.
// Corresponding traffic models are:
//   01. Regular Pod (remote)     -- hostPort [request]              --> AntreaIPAM VLAN Pod
//   02. AntreaIPAM Pod           -- hostPort [request]              --> AntreaIPAM VLAN Pod
//   03. AntreaIPAM VLAN Pod      -- hostPort [request]              --> AntreaIPAM VLAN Pod (different subnet/VLAN)
//   04. External                 -- hostPort [request]              --> AntreaIPAM VLAN Pod
// Below traffic models are already covered by portmap CNI:
//   01. AntreaIPAM VLAN Pod      -- hostPort [request]              --> AntreaIPAM VLAN Pod (same subnet)
//   02. Regular Pod (local)      -- hostPort [request]              --> AntreaIPAM VLAN Pod

// We do not use --random-fully for this rule for consistency with the portmap CNI plugin.
// https://github.com/containernetworking/plugins/blob/c29dc79f96cd50452a247a4591443d2aac033429/plugins/meta/portmap/portmap.go#L321-L345

// When Antrea is used to enforce NetworkPolicies in EKS, additional iptables
// nat rules are required. See https://github.com/antrea-io/antrea/issues/3946.
// These rules are only needed for IPv4.

func (c *Client) initIPRoutes() error { _ = "STUB: not implemented"; return nil }

func (c *Client) initServiceIPRoutes() error { _ = "STUB: not implemented"; return nil }

func (c *Client) initNodeNetworkPolicy() { _ = "STUB: not implemented"; return }

func buildAllowHostIngressPortRule(protocol string, port *intstr.IntOrString, comment string) string {
	_ = "STUB: not implemented"
	return ""
}

func buildAllowHostEgressPortRule(protocol string, port *intstr.IntOrString, comment string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Client) initProxyHealthCheck() { _ = "STUB: not implemented"; return }

// Proxy health check reply packets are sent from the listening port. This rule ensures that the packets are
// allowed to output.

func (c *Client) initWireguard() { _ = "STUB: not implemented"; return }

func (c *Client) initNodeLatencyRules() {
	_ = "STUB: not implemented"
	// the interface on which ICMP probes are sent / received is the Antrea gateway interface, except
	// in networkPolicyOnly mode, for which it is the Node's transport interface.
	return
}

func (c *Client) initEgressIPRules() error { _ = "STUB: not implemented"; return nil }

// Delete all rules whose target table is ReplyEgressRouteTable on startup. These rules will be reinstalled
// each time after restart. This implementation was inspired by the implementation of RestoreEgressRoutesAndRules.

func (c *Client) initEgressIPRoutes() error { _ = "STUB: not implemented"; return nil }

// Reconcile removes orphaned podCIDRs from ipset and removes routes to orphaned podCIDRs
// based on the desired podCIDRs.
func (c *Client) Reconcile(podCIDRs []string) error { _ = "STUB: not implemented"; return nil }

// Get the peer IPv6 gateways from pod CIDRs

// Remove orphaned podCIDRs from ipset.

// Remove any unknown routes on Antrea gateway.

// The route to the IPv6 link-local CIDR is always auto-generated by the system along with
// a link-local address, which is not configured by Antrea and should therefore to be ignored
// in the "deletion" list. Such routes are useful in some cases, e.g., IPv6 NDP.

// IPv6 doesn't support "on-link" route, routes to the peer IPv6 gateways need to
// be added separately. So don't delete such routes.

// Don't delete the routes which are added by AntreaProxy when proxyAll is enabled.

// Return immediately if there is no IPv6 gateway address configured on the Nodes.

// Remove orphaned IPv6 Neighbors from host network.

// Remove any unknown IPv6 neighbors on Antrea gateway.

// Don't delete the virtual Service IP neighbor which is added by AntreaProxy.

func (c *Client) isServiceRoute(route *netlink.Route) bool {
	_ = "STUB: not implemented"
	// If the gateway IP or the destination IP is the virtual Service IP, then it is a route added by AntreaProxy.
	return false
}

// listIPRoutes returns list of routes on Antrea gateway.
func (c *Client) listIPRoutesOnGW() ([]netlink.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RestoreEgressRoutesAndRules simply deletes all IP routes and rules created for Egresses for now.
// It may be better to keep the ones whose Egress IPs are still on this Node, but it's a bit hard to achieve it at the
// moment because the marks are not permanent and could change upon restart.
func (c *Client) RestoreEgressRoutesAndRules(minTableID, maxTableID int) error {
	_ = "STUB: not implemented"
	return nil
}

// Not routes created for Egress.

// Not rules created for Egress.

// getIPv6Gateways returns the IPv6 gateway addresses of the given CIDRs.
func getIPv6Gateways(podCIDRs []string) sets.Set[string] { _ = "STUB: not implemented"; return nil }

func (c *Client) listIPv6NeighborsOnGateway() (map[string]*netlink.Neigh, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddRoutes adds routes to a new podCIDR. It overrides the routes if they already exist.
func (c *Client) AddRoutes(podCIDR *net.IPNet, nodeName string, nodeIP, nodeGwIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// Add this podCIDR to antreaPodIPSet so that packets to them won't be masqueraded when they leave the host.

// Install routes to this Node.

// If WireGuard is enabled, create a route via WireGuard device regardless of the traffic encapsulation modes.

// "on-link" is not identified in IPv6 route entries, so split the configuration into 2 entries.
// TODO: Kernel >= 4.16 supports adding IPv6 route with onlink flag. Delete this route after Kernel version
//       requirement bump in future.

// NoEncap traffic to Node on the same subnet.
// Set the peerNodeIP as next hop.

// Update the nftables set that contains all peer PodCIDRs.

// NetworkPolicyOnly mode or NoEncap traffic to a Node on a different subnet.
// Routing should be handled by a route which is already present on the host.

// Delete stale route and neigh to peer gateway.

// Add IPv6 neighbor if the given podCIDR is using IPv6 address.

// DeleteRoutes deletes routes to a PodCIDR. It does nothing if the routes doesn't exist.
func (c *Client) DeleteRoutes(podCIDR *net.IPNet) error { _ = "STUB: not implemented"; return nil }

// Delete this podCIDR from antreaPodIPSet as the CIDR is no longer for Pods.

// Join all words with spaces, terminate with newline and write to buf.
func writeLine(buf *bytes.Buffer, words ...string) {
	_ = "STUB: not implemented"
	// We avoid strings.Join for performance reasons.
	return
}

// MigrateRoutesToGw moves routes (including assigned IP addresses if any) from link linkName to
// host gateway.
func (c *Client) MigrateRoutesToGw(linkName string) error { _ = "STUB: not implemented"; return nil }

// Swap route first then address, otherwise route gets removed when address is removed.

// Swap address if any.

// UnMigrateRoutesFromGw moves route from gw to link linkName if provided; otherwise route is deleted
func (c *Client) UnMigrateRoutesFromGw(route *net.IPNet, linkName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) snatRuleSpec(snatIP net.IP, snatMark uint32) []string {
	_ = "STUB: not implemented"
	return nil
}

// The condition is needed to prevent the rule from being applied to local out packets destined for Pods, which
// have "0x1/0x1" mark.

func (c *Client) AddSNATRule(snatIP net.IP, mark uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteSNATRule(mark uint32) error { _ = "STUB: not implemented"; return nil }

func (c *Client) AddEgressRoutes(tableID uint32, dev int, gateway net.IP, prefixLength int) error {
	_ = "STUB: not implemented"
	return nil
}

// Install routes for the subnet, for example:
// tableID=101, dev=eth.10, gateway=172.20.10.1, prefixLength=24
// $ ip route show table 101
// 172.20.10.0/24 dev eth0.10 table 101
// default via 172.20.10.1 dev eth0.10 table 101

func (c *Client) DeleteEgressRoutes(tableID uint32) error { _ = "STUB: not implemented"; return nil }

func (c *Client) AddEgressRule(tableID uint32, mark uint32, isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteEgressRule(tableID uint32, mark uint32, isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// addVirtualServiceIPRoute is used to add a route which is used to route the packets whose destination IP is a virtual
// IP to Antrea gateway.
func (c *Client) addVirtualServiceIPRoute(isIPv6 bool) error { _ = "STUB: not implemented"; return nil }

// addEgressReplyDefaultPolicyIPRoute is used to add a default route to policy-routing table ReplyEgressRouteTable when traffic
// mode is hybrid. The route is to route the reply Egress packets originated from remote Nodes to Antrea gateway,
// avoiding the packets to be matched by the routes in main route table. This ensures that the reply Egress packets
// can be forwarded back to the remote Nodes through OVS flow-based tunnel.
func (c *Client) addEgressReplyDefaultPolicyIPRoute(isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// AddNodePortConfigs adds tuples of `IP, protocol, port` for every NodePort IP to the target set. The set is used by
// netfilter rules in the Node's host network as the destination, which directs the traffic to the OVS pipeline via
// Antrea gateway interface for further processing, bypassing kube-proxy processing.
func (c *Client) AddNodePortConfigs(nodePortAddresses []net.IP, port uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteNodePortConfigs deletes corresponding tuples from the target set when a NodePort Service is deleted.
func (c *Client) DeleteNodePortConfigs(nodePortAddresses []net.IP, port uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) addServiceCIDRRoute(serviceCIDR *net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate a route with the new Service CIDR and install it.

// Store the new Service CIDR.

// Collect stale routes.

// If current destination CIDR is not nil, the route with current destination CIDR should be uninstalled.

// If current destination CIDR is nil, which means that Antrea Agent has just started, then all existing routes
// whose destination CIDR contains the first ClusterIP should be uninstalled, except the newly installed route.
// Note that, there may be multiple stale routes prior to this commit. When upgrading, all stale routes will be
// collected. After this commit, there will be only one stale route after Antrea Agent started.

// Not the routes we are interested in.

// It's the latest route we just installed.

// The route covers the desired route. It was installed when the calculated ServiceCIDR was larger than the
// current one, which could happen after some Services are deleted.

// The desired route covers the route. It was installed when the calculated ServiceCIDR was smaller than the
// current one, which could happen after some Services are added.

// Remove stale routes.

func (c *Client) addVirtualNodePortDNATIPRoute(isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// AddExternalIPConfigs adds a route entry to forward traffic destined for the external Service IP to the Antrea
// gateway interface. Additionally, it adds the IP to the target set. The set is used by netfilter rules in the
// Node's host network as the destination, skipping conntrack and bypassing kube-proxy processing.
func (c *Client) AddExternalIPConfigs(svcInfoStr string, externalIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteExternalIPConfigs deletes the route entry to forward traffic destined for the external Service IP to the Antrea
// gateway interface. Additionally, it removes the IP from the target set.
func (c *Client) DeleteExternalIPConfigs(svcInfoStr string, externalIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// AddLocalAntreaFlexibleIPAMPodRule is used to add IP to target ip set when an AntreaFlexibleIPAM Pod is added. An entry is added
// for every Pod IP.
func (c *Client) AddLocalAntreaFlexibleIPAMPodRule(podAddresses []net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip Per-Node IPAM Pod

// DeleteLocalAntreaFlexibleIPAMPodRule is used to delete related IP set entries when an AntreaFlexibleIPAM Pod is deleted.
func (c *Client) DeleteLocalAntreaFlexibleIPAMPodRule(podAddresses []net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// addNodeIP adds nodeIP into the ipset when a new Node joins the cluster.
// The ipset is consumed with encap mode when multicast is enabled.
func (c *Client) addNodeIP(podCIDR *net.IPNet, nodeIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteNodeIP deletes NodeIPs from the ipset when a Node leaves the cluster.
// The ipset is consumed with encap mode when multicast is enabled.
func (c *Client) deleteNodeIP(podCIDR *net.IPNet) error { _ = "STUB: not implemented"; return nil }

func (c *Client) AddRouteForLink(cidr *net.IPNet, linkIndex int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DeleteRouteForLink(cidr *net.IPNet, linkIndex int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) ClearConntrackEntryForService(svcIP net.IP, svcPort uint16, endpointIP net.IP, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func getTransProtocolStr(protocol binding.Protocol) string { _ = "STUB: not implemented"; return "" }

func isIPv6Protocol(protocol binding.Protocol) bool { _ = "STUB: not implemented"; return false }

func generateRule(table int, mark uint32, mask *uint32, family int) *netlink.Rule {
	_ = "STUB: not implemented"
	return nil
}

func generateRoute(ip net.IP,
	mask int,
	gw net.IP,
	linkIndex int,
	scope netlink.Scope,
	table *int,
	flags *int) *netlink.Route {
	_ = "STUB: not implemented"
	return nil
}

func generateNeigh(ip net.IP, linkIndex int) *netlink.Neigh { _ = "STUB: not implemented"; return nil }

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

func (c *Client) nftablesProxyAll(tx *knftables.Transaction, ipProtocol knftables.Family) {
	_ = "STUB: not implemented"
	return
}

// The priority ensures that Antrea takes precedence over kube-proxy when matching Service traffic originating
// from external hosts.

// The priority ensures that Antrea takes precedence over kube-proxy when matching Service traffic originating
// from the local Node.

// Flush the sets if it exists.

// Restore the elements.

func (c *Client) nftablesTrafficAcceleration(tx *knftables.Transaction, ipProtocol knftables.Family) {
	_ = "STUB: not implemented"
	// Add the flowtable for accelerating matched connections.
	return
}

// Add a forward hook chain to contain the rules for matching connections which are eligible for flowtable acceleration.

// Flush the chain if it already exists. This is safe because the entire transaction is applied atomically:
// nftables builds the new config in memory and swaps it in one step, so there is never a moment when rules are
// partially removed or missing. See: https://wiki.nftables.org/wiki-nftables/index.php/Atomic_rule_replacement

// Add the set to contain peer Pod IPv4 CIDRs.

// Flush the set if it exists.

// Restore existing IPv4 peer Pod CIDRs into the set.

// Add rules to accelerate Pod-to-Pod IPv4 connections via flowtable.

// Add the nft set to contain peer Pod IPv6 CIDRs.

// Flush the set if it exists.

// Restore existing IPv6 peer Pod CIDRs into the set.

// Add rules to accelerate Pod-to-Pod IPv6 connections via flowtable.

func (c *Client) syncNFTables(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Add the table whose name is defined when initializing nftables instance.

func (c *Client) addPeerPodCIDRToNFTablesSet(podCIDR *net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) deletePeerPodCIDRFromNFTablesSet(podCIDR *net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: nftables error messages are not stable across versions / distros.
// This is a best-effort check until lib sigs.k8s.io/knftables provides a structured NotFound error for set
// elements.

func (c *Client) getNFT(isIPv6 bool) knftables.Interface {
	_ = "STUB: not implemented"
	return *new(knftables.Interface)
}

func (c *Client) addNodePortConfigsNFTables(nodePortAddresses []net.IP, port uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) addNodePortConfigsIPsets(nodePortAddresses []net.IP, port uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) deleteNodePortConfigsNFTables(nodePortAddresses []net.IP, port uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) deleteNodePortConfigsIPsets(nodePortAddresses []net.IP, port uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

// addExternalIPConfigsRoute is idempotent. If the route has already been added, it returns no error. This ensures
// correctness when addExternalIPConfigsRoute calls it even after a prior partial cleanup or failure.
func (c *Client) addExternalIPConfigsRoute(externalIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) addExternalIPConfigsNFTables(externalIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) addExternalIPConfigsIPsets(externalIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteExternalIPConfigsRoute is idempotent. If the route has already been deleted, it returns no error. This ensures
// correctness when DeleteExternalIPConfigs calls it even after a prior partial cleanup or failure.
func (c *Client) deleteExternalIPConfigsRoute(externalIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) deleteExternalIPConfigsNFTables(externalIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) deleteExternalIPConfigsIPsets(externalIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// shouldEnableEgressPolicyRouting returns true when Egress is enabled and policy based routing
// is needed for Egress traffic, such as when traffic encapsulation mode is hybrid or traffic
// encryption mode is WireGuard. In those cases, Egress uses a tunnel path distinct from the common
// Pod-to-Pod traffic path (e.g. Pod-to-Pod may go via routing in hybrid), so symmetric-path handling
// for Egress reply packets is needed.
func (c *Client) shouldEnableEgressPolicyRouting() bool { _ = "STUB: not implemented"; return false }
