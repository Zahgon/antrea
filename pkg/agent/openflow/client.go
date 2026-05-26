// Copyright 2019 Antrea Authors
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

package openflow

import (
	"net"

	"antrea.io/libOpenflow/openflow15"
	"antrea.io/libOpenflow/protocol"
	ofutil "antrea.io/libOpenflow/util"
	"antrea.io/ofnet/ofctrl"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/types"
	"antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	crdv1alpha2 "antrea.io/antrea/v2/pkg/apis/crd/v1alpha2"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
	utilip "antrea.io/antrea/v2/pkg/util/ip"
	"antrea.io/antrea/v2/third_party/proxy"
)

const maxRetryForOFSwitch = 5

func tcPriorityToOFPriority(p types.TrafficControlFlowPriority) uint16 {
	_ = "STUB: not implemented"
	return 0
}

// Client is the interface to program OVS flows for entity connectivity of Antrea.
type Client interface {
	// Initialize sets up all basic flows on the specific OVS bridge. It returns a channel which
	// is used to notify the caller in case of a reconnection, in which case ReplayFlows should
	// be called to ensure that the set of OVS flows is correct. All flows programmed in the
	// switch which match the current round number will be deleted before any new flow is
	// installed.
	Initialize(roundInfo types.RoundInfo,
		config *config.NodeConfig,
		networkConfig *config.NetworkConfig,
		egressConfig *config.EgressConfig,
		serviceConfig *config.ServiceConfig,
		l7NetworkPolicyConfig *config.L7NetworkPolicyConfig) (<-chan struct{}, error)

	// InstallNodeFlows should be invoked when a connection to a remote Node is going to be set
	// up. The hostname is used to identify the added flows. When IPsec tunnel is enabled,
	// ipsecTunOFPort must be set to the OFPort number of the IPsec tunnel port to the remote Node;
	// otherwise ipsecTunOFPort must be 0. In dual-stack, IPv6 traffic will be encapsulated in the
	// IPv4 tunnel so it goes through the IPsec tunnel and gets encrypted.
	// InstallNodeFlows has all-or-nothing semantics(call succeeds if all the flows are installed
	// successfully, otherwise no flows will be installed). Calls to InstallNodeFlows are idempotent.
	// Concurrent calls to InstallNodeFlows and / or UninstallNodeFlows are supported as long as they
	// are all for different hostnames.
	InstallNodeFlows(
		hostname string,
		peerConfigs map[*net.IPNet]net.IP,
		peerNodeIPs *utilip.DualStackIPs,
		ipsecTunOFPort uint32,
		peerNodeMAC net.HardwareAddr) error

	// UninstallNodeFlows removes the connection to the remote Node specified with the
	// hostname. UninstallNodeFlows will do nothing if no connection to the host was established.
	UninstallNodeFlows(hostname string) error

	// InstallPodFlows should be invoked when a connection to a Pod on current Node. The
	// interfaceName is used to identify the added flows. InstallPodFlows has all-or-nothing
	// semantics(call succeeds if all the flows are installed successfully, otherwise no
	// flows will be installed). Calls to InstallPodFlows are idempotent. Concurrent calls
	// to InstallPodFlows and / or UninstallPodFlows are supported as long as they are all
	// for different interfaceNames.
	InstallPodFlows(interfaceName string, podInterfaceIPs []net.IP, podInterfaceMAC net.HardwareAddr, ofPort uint32, vlanID uint16, labelID *uint32) error

	// UninstallPodFlows removes the connection to the local Pod specified with the
	// interfaceName. UninstallPodFlows will do nothing if no connection to the Pod was established.
	UninstallPodFlows(interfaceName string) error

	// InstallServiceGroup installs a group for Service LB. Each endpoint
	// is a bucket of the group. For now, each bucket has the same weight.
	InstallServiceGroup(groupID binding.GroupIDType, withSessionAffinity bool, endpoints []proxy.Endpoint) error
	// UninstallServiceGroup removes the group and its buckets that are
	// installed by InstallServiceGroup.
	UninstallServiceGroup(groupID binding.GroupIDType) error

	// InstallEndpointFlows installs flows for accessing Endpoints.
	// If an Endpoint is on the current Node, then flows for hairpin and endpoint
	// L2 forwarding should also be installed.
	InstallEndpointFlows(protocol binding.Protocol, endpoints []proxy.Endpoint) error
	// UninstallEndpointFlows removes flows of the Endpoint installed by
	// InstallEndpointFlows.
	UninstallEndpointFlows(protocol binding.Protocol, endpoints []proxy.Endpoint) error

	// InstallServiceFlows installs flows for accessing Service NodePort, LoadBalancer, ExternalIP and ClusterIP. It
	// installs the flow that uses the group/bucket to do Service LB. If the affinityTimeout is not zero, it also
	// installs the flow which has a learn action to maintain the LB decision. The group with the groupID must be
	// installed before, otherwise the installation will fail.
	// For an external IP with Local traffic policy (IsExternal == true and TrafficPolicyLocal == true), it also
	// installs the flow to implement short-circuiting for internally originated traffic towards external IPs.
	InstallServiceFlows(config *types.ServiceConfig) error
	// UninstallServiceFlows removes flows installed by InstallServiceFlows.
	UninstallServiceFlows(svcIP net.IP, svcPort uint16, protocol binding.Protocol) error

	// GetFlowTableStatus should return an array of flow table status, all existing flow tables should be included in the list.
	GetFlowTableStatus() []binding.TableStatus

	// InstallPolicyRuleFlows installs flows for a new NetworkPolicy rule. Rule should include all fields in the
	// NetworkPolicy rule. Each ingress/egress policy rule installs Openflow entries on two tables, one for
	// ruleTable and the other for dropTable. If a packet does not pass the ruleTable, it will be dropped by the
	// dropTable.
	InstallPolicyRuleFlows(ofPolicyRule *types.PolicyRule) error

	// BatchInstallPolicyRuleFlows installs multiple flows for NetworkPolicy rules in batch.
	BatchInstallPolicyRuleFlows(ofPolicyRules []*types.PolicyRule) error

	// UninstallPolicyRuleFlows removes the Openflow entry relevant to the specified NetworkPolicy rule.
	// It also returns a slice of stale ofPriorities used by ClusterNetworkPolicies.
	// UninstallPolicyRuleFlows will do nothing if no Openflow entry for the rule is installed.
	UninstallPolicyRuleFlows(ruleID uint32) ([]string, error)

	// AddPolicyRuleAddress adds one or multiple addresses to the specified NetworkPolicy rule. If addrType is true, the
	// addresses are added to PolicyRule.From, else to PolicyRule.To.
	AddPolicyRuleAddress(ruleID uint32, addrType types.AddressType, addresses []types.Address, priority *uint16, enableLogging, isMCNPRule bool) error

	// DeletePolicyRuleAddress removes addresses from the specified NetworkPolicy rule. If addrType is srcAddress, the addresses
	// are removed from PolicyRule.From, else from PolicyRule.To.
	DeletePolicyRuleAddress(ruleID uint32, addrType types.AddressType, addresses []types.Address, priority *uint16) error

	// InstallSNATBypassServiceFlows installs flows to prevent traffic destined for the specified Service CIDRs from
	// being SNAT'd. Otherwise, such Pod-to-Service traffic would be forwarded to Egress Node and be load-balanced
	// remotely, as opposed to locally, when AntreaProxy is asked to skip some Services or is not running at all.
	// Calling the method with new CIDRs will override the flows installed for previous CIDRs.
	InstallSNATBypassServiceFlows(serviceCIDRs []*net.IPNet) error

	// InstallSNATMarkFlows installs flows for a local SNAT IP. On Linux, a
	// single flow is added to mark the packets tunnelled from remote Nodes
	// that should be SNAT'd with the SNAT IP.
	InstallSNATMarkFlows(snatIP net.IP, mark uint32) error

	// UninstallSNATMarkFlows removes the flows installed to set the packet
	// mark for a SNAT IP.
	UninstallSNATMarkFlows(mark uint32) error

	// InstallPodSNATFlows installs the SNAT flows for a local Pod. If the
	// SNAT IP for the Pod is on the local Node, a non-zero SNAT ID should
	// be allocated for the SNAT IP, and the installed flow sets the SNAT IP
	// mark on the egress packets from the ofPort; if the SNAT IP is on a
	// remote Node, snatMark should be set to 0, and the installed flow
	// tunnels egress packets to the remote Node using the SNAT IP as the
	// tunnel destination, and the packets should be SNAT'd on the remote
	// Node. As of now, a Pod can be configured to use only a single SNAT
	// IP in a single address family (IPv4 or IPv6).
	InstallPodSNATFlows(ofPort uint32, snatIP net.IP, snatMark uint32) error

	// UninstallPodSNATFlows removes the SNAT flows for the local Pod.
	UninstallPodSNATFlows(ofPort uint32) error

	// InstallEgressQoS installs an OF meter with specific meterID, rate
	// and burst used for QoS of Egress and a QoS flow that direct packets
	// into the meter.
	InstallEgressQoS(meterID, rate, burst uint32) error

	// UninstallEgressQoS removes the flow and OF meter used by QoS of Egress.
	UninstallEgressQoS(meterID uint32) error

	// Disconnect disconnects the connection between client and OFSwitch.
	Disconnect() error

	// IsConnected returns the connection status between client and OFSwitch. The return value is true if the OFSwitch is connected.
	IsConnected() bool

	// ReplayFlows should be called when a spurious disconnection occurs. After we reconnect to
	// the OFSwitch, we need to replay all the flows cached by the client. ReplayFlows will try
	// to replay as many flows as possible, and will log an error when a flow cannot be
	// installed.
	ReplayFlows()

	// DeleteStaleFlows deletes all flows from the previous round which are no longer needed. It
	// should be called by the agent after all required flows have been installed / updated with
	// the new round number.
	DeleteStaleFlows() error

	// GetTunnelVirtualMAC() returns GlobalVirtualMAC used for tunnel traffic.
	GetTunnelVirtualMAC() net.HardwareAddr

	// GetPodFlowKeys returns the keys (match strings) of the cached flows for a
	// Pod.
	GetPodFlowKeys(interfaceName string) []string

	// GetServiceFlowKeys returns the keys (match strings) of the cached
	// flows for a Service (port) and its endpoints.
	GetServiceFlowKeys(svcIP net.IP, svcPort uint16, protocol binding.Protocol, endpoints []proxy.Endpoint) []string

	// GetNetworkPolicyFlowKeys returns the keys (match strings) of the cached
	// flows for a NetworkPolicy. Flows are grouped by policy rules, and duplicated
	// entries can be added due to conjunctive match flows shared by multiple
	// rules.
	GetNetworkPolicyFlowKeys(npName, npNamespace string, npType v1beta2.NetworkPolicyType) []string

	// ReassignFlowPriorities takes a list of priority updates, and update the actionFlows to replace
	// the old priority with the desired one, for each priority update on that table.
	ReassignFlowPriorities(updates map[uint16]uint16, table uint8) error

	// SubscribePacketIn subscribes to packet in messages for the given category. Packets
	// will be placed in the queue and if the queue is full, the packet in messages
	// will be dropped. pktInQueue supports rate-limiting for the consumer, in order to
	// constrain the compute resources that may be used by the consumer.
	SubscribePacketIn(reason uint8, pktInQueue *binding.PacketInQueue) error

	// SendTraceflowPacket injects packet to specified OVS port for Openflow.
	SendTraceflowPacket(dataplaneTag uint8, packet *binding.Packet, inPort uint32, outPort int32) error

	// InstallTraceflowFlows installs flows for a Traceflow request.
	InstallTraceflowFlows(dataplaneTag uint8, liveTraffic, droppedOnly, receiverOnly bool, packet *binding.Packet, ofPort uint32, timeoutSeconds uint16) error

	// UninstallTraceflowFlows uninstalls flows for a Traceflow request.
	UninstallTraceflowFlows(dataplaneTag uint8) error

	// GetPolicyInfoFromConjunction returns the following policy information for the provided conjunction ID:
	// NetworkPolicy reference, OF priority, rule name, label
	// The boolean return value indicates whether the policy information was found.
	GetPolicyInfoFromConjunction(ruleID uint32) (bool, *v1beta2.NetworkPolicyReference, string, string, string)

	// RegisterPacketInHandler uses SubscribePacketIn to get PacketIn message and process received
	// packets through registered handler.
	RegisterPacketInHandler(packetHandlerReason uint8, packetInHandler PacketInHandler)

	StartPacketInHandler(stopCh <-chan struct{})
	// Get traffic metrics of each NetworkPolicy rule.
	NetworkPolicyMetrics() map[uint32]*types.RuleMetric

	// Get multicast ingress metrics of each Pod in MulticastIngressPodMetricTable.
	MulticastIngressPodMetrics() map[uint32]*types.RuleMetric
	// Get multicast Pod ingress statistics from MulticastIngressPodMetricTable with specified ofPort.
	MulticastIngressPodMetricsByOFPort(ofPort int32) *types.RuleMetric
	// Get multicast egress metrics of each Pod in MulticastEgressPodMetricTable.
	MulticastEgressPodMetrics() map[string]*types.RuleMetric
	// Get multicast Pod ingress statistics from MulticastEgressPodMetricTable with specified src IP.
	MulticastEgressPodMetricsByIP(ip net.IP) *types.RuleMetric

	// SendTCPPacketOut sends TCP packet as a packet-out to OVS.
	SendTCPPacketOut(
		srcMAC string,
		dstMAC string,
		srcIP string,
		dstIP string,
		inPort uint32,
		outPort uint32,
		isIPv6 bool,
		tcpSrcPort uint16,
		tcpDstPort uint16,
		tcpSeqNum uint32,
		tcpAckNum uint32,
		tcpHdrLen uint8,
		tcpFlag uint8,
		tcpWinSize uint16,
		tcpData []byte,
		mutatePacketOut func(builder binding.PacketOutBuilder) binding.PacketOutBuilder) error
	// SendICMPPacketOut sends ICMP packet as a packet-out to OVS.
	SendICMPPacketOut(
		srcMAC string,
		dstMAC string,
		srcIP string,
		dstIP string,
		inPort uint32,
		outPort uint32,
		isIPv6 bool,
		icmpType uint8,
		icmpCode uint8,
		icmpData []byte,
		mutatePacketOut func(builder binding.PacketOutBuilder) binding.PacketOutBuilder) error
	// SendUDPPacketOut sends UDP packet as a packet-out to OVS.
	SendUDPPacketOut(
		srcMAC string,
		dstMAC string,
		srcIP string,
		dstIP string,
		inPort uint32,
		outPort uint32,
		isIPv6 bool,
		udpSrcPort uint16,
		udpDstPort uint16,
		udpData []byte,
		mutatePacketOut func(builder binding.PacketOutBuilder) binding.PacketOutBuilder) error
	// SendEthPacketOut sends ethernet packet as a packet-out to OVS.
	SendEthPacketOut(inPort, outPort uint32, ethPkt *protocol.Ethernet, mutatePacketOut func(builder binding.PacketOutBuilder) binding.PacketOutBuilder) error
	// ResumePausePacket resumes a paused packetIn.
	ResumePausePacket(packetIn *ofctrl.PacketIn) error
	// NewDNSPacketInConjunction creates a policyRuleConjunction for the dns response interception flows.
	NewDNSPacketInConjunction(id uint32) error
	// AddAddressToDNSConjunction adds addresses to the toAddresses of the dns packetIn conjunction,
	// so that dns response packets sent towards these addresses will be intercepted and parsed by
	// the fqdnController.
	AddAddressToDNSConjunction(id uint32, addrs []types.Address) error
	// DeleteAddressFromDNSConjunction removes addresses from the toAddresses of the dns packetIn conjunction.
	DeleteAddressFromDNSConjunction(id uint32, addrs []types.Address) error

	// InstallMulticastFlows installs the flow to forward Multicast traffic normally, and output it to antrea-gw0
	// to ensure it can be forwarded to the external addresses.
	InstallMulticastFlows(multicastIP net.IP, groupID binding.GroupIDType) error

	// UninstallMulticastFlows removes the flow matching the given multicastIP.
	UninstallMulticastFlows(multicastIP net.IP) error
	// InstallMulticastFlexibleIPAMFlows installs two flows and forwards them to the first table of Multicast Pipeline
	// when flexibleIPAM is enabled, with one flow matching inbound multicast traffic from the uplink and the other from
	// the host interface, making multicast packets coming from the host and other Nodes be forward to the OVS multicast pipeline.
	InstallMulticastFlexibleIPAMFlows() error
	// InstallMulticastRemoteReportFlows installs flows to forward the IGMP report messages to the other Nodes,
	// and packetIn the report messages to Antrea Agent which is received via tunnel port.
	// The OpenFlow group identified by groupID is used to forward packet to all other Nodes in the cluster
	// over tunnel.
	InstallMulticastRemoteReportFlows(groupID binding.GroupIDType) error
	// SendIGMPQueryPacketOut sends the IGMPQuery packet as a packet-out to OVS from the gateway port.
	SendIGMPQueryPacketOut(
		dstMAC net.HardwareAddr,
		dstIP net.IP,
		outPort uint32,
		igmp ofutil.Message) error

	// InstallTrafficControlMarkFlows installs the flows to mark the packets for a traffic control rule.
	InstallTrafficControlMarkFlows(name string,
		sourceOFPorts []uint32,
		targetOFPort uint32,
		direction crdv1alpha2.Direction,
		action crdv1alpha2.TrafficControlAction,
		priority types.TrafficControlFlowPriority) error

	// UninstallTrafficControlMarkFlows removes the flows for a traffic control rule.
	UninstallTrafficControlMarkFlows(name string) error

	// InstallTrafficControlReturnPortFlow installs the flow to classify the packets from a return port.
	InstallTrafficControlReturnPortFlow(returnOFPort uint32) error

	// UninstallTrafficControlReturnPortFlow removes the flow to classify the packets from a return port.
	UninstallTrafficControlReturnPortFlow(returnOFPort uint32) error

	InstallMulticastGroup(ofGroupID binding.GroupIDType, localReceivers []uint32, remoteNodeReceivers []net.IP) error
	// UninstallMulticastGroup removes the group and its buckets that are
	// installed by InstallMulticastGroup.
	UninstallMulticastGroup(groupID binding.GroupIDType) error

	// SendIGMPRemoteReportPacketOut sends the IGMP report packet as a packet-out to remote Nodes via the tunnel port.
	SendIGMPRemoteReportPacketOut(
		dstMAC net.HardwareAddr,
		dstIP net.IP,
		igmp ofutil.Message) error

	// InstallMulticlusterNodeFlows installs flows to handle cross-cluster packets between a regular
	// Node and a local Gateway.
	InstallMulticlusterNodeFlows(
		clusterID string,
		peerConfigs map[*net.IPNet]net.IP,
		tunnelPeerIP net.IP,
		enableStretchedNetworkPolicy bool) error

	// InstallMulticlusterGatewayFlows installs flows to handle cross-cluster packets between Gateways.
	InstallMulticlusterGatewayFlows(
		clusterID string,
		peerConfigs map[*net.IPNet]net.IP,
		tunnelPeerIP net.IP,
		localGatewayIP net.IP,
		enableStretchedNetworkPolicy bool) error

	// InstallMulticlusterClassifierFlows installs flows to classify cross-cluster packets.
	InstallMulticlusterClassifierFlows(tunnelOFPort uint32, isGateway bool) error

	// InstallMulticlusterPodFlows installs flows to handle cross-cluster packets from Multi-cluster Gateway to
	// regular Nodes.
	InstallMulticlusterPodFlows(podIP net.IP, tunnelPeerIP net.IP) error

	// UninstallMulticlusterFlows removes cross-cluster flows matching the given cache key on
	// a regular Node or a Gateway.
	UninstallMulticlusterFlows(clusterID string) error

	// UninstallMulticlusterPodFlows removes Pod flows matching the given cache key on
	// a Gateway. When the podIP is empty, all Pod flows will be removed.
	UninstallMulticlusterPodFlows(podIP string) error

	// InstallVMUplinkFlows installs flows to forward packet between uplinkPort and hostPort. On a VM, the
	// uplink and host internal port are paired directly, and no layer 2/3 forwarding flow is installed.
	InstallVMUplinkFlows(hostInterfaceName string, hostPort int32, uplinkPort int32) error

	// UninstallVMUplinkFlows removes the flows installed to forward packet between uplinkPort and hostPort.
	UninstallVMUplinkFlows(hostInterfaceName string) error

	// InstallPolicyBypassFlows installs flows to bypass the NetworkPolicy rules on the traffic with the given ipnet
	// or ip, port, protocol and direction. It is used to bypass NetworkPolicy enforcement on a VM for the particular
	// traffic.
	InstallPolicyBypassFlows(protocol binding.Protocol, ipNet *net.IPNet, port uint16, isIngress bool) error

	// SubscribeOFPortStatusMessage registers a channel to listen the OpenFlow PortStatus message.
	SubscribeOFPortStatusMessage(statusCh chan *openflow15.PortStatus)

	// InstallL7NetworkPolicyFlows will be called only when at least one L7 NetworkPolicy is applied locally.
	InstallL7NetworkPolicyFlows() error
}

// GetFlowTableStatus returns an array of flow table status.
func (c *client) GetFlowTableStatus() []binding.TableStatus { _ = "STUB: not implemented"; return nil }

// IsConnected returns the connection status between client and OFSwitch.
func (c *client) IsConnected() bool { _ = "STUB: not implemented"; return false }

// addFlows installs the flows on the OVS bridge and then add them into the flow cache. If the flow cache exists,
// it will return immediately, otherwise it will use Bundle to add all flows, and then add them into the flow cache.
// If it fails to add the flows with Bundle, it will return the error and no flow cache is created.
func (c *client) addFlows(cache *flowCategoryCache, flowCacheKey string, flows []binding.Flow) error {
	_ = "STUB: not implemented"
	return nil
}

// addFlowsWithMultipleKeys installs the flows with different flowMessageCache keys and adds them into the cache on success.
// It will skip flows whose cache already exists. All flows will be installed via a bundle.
func (c *client) addFlowsWithMultipleKeys(cache *flowCategoryCache, keyToFlows map[string][]binding.Flow) error {
	_ = "STUB: not implemented"
	// allMessages keeps the OpenFlow modification messages we will install via a bundle.
	return nil
}

// flowCacheMap keeps the flowMessageCache items we will add to the cache on bundle success.

// If a flow cache entry already exists for the key, skip it.

// Add the installed flows into the flow cache.

// modifyFlows sets the flows of flowCategoryCache be exactly same as the provided slice for the given flowCacheKey.
func (c *client) modifyFlows(cache *flowCategoryCache, flowCacheKey string, flows []binding.Flow) error {
	_ = "STUB: not implemented"
	return nil
}

// Modify the flows in the flow cache.

// deleteFlows deletes all the flows in the flow cache indexed by the provided flowCacheKey.
func (c *client) deleteFlows(cache *flowCategoryCache, flowMessageCacheKey string) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteFlowsWithMultipleKeys uninstalls the flows with different flowMessageCache keys and remove them from the cache on success.
// It will skip the keys which are not in the cache. All flows will be uninstalled via a bundle.
func (c *client) deleteFlowsWithMultipleKeys(cache *flowCategoryCache, keys []string) error {
	_ = "STUB: not implemented"
	// allFlows keeps the flows we will delete via a bundle.
	return nil
}

// If a flow cache entry of the key does not exist, skip it.

// Delete the keys and corresponding flows from the flow cache.

func (c *client) deleteAllFlows(cache *flowCategoryCache) error {
	_ = "STUB: not implemented"
	return nil
}

// InstallNodeFlows installs flows for peer Nodes. Parameter remoteGatewayMAC is only for Windows.
func (c *client) InstallNodeFlows(hostname string,
	peerConfigs map[*net.IPNet]net.IP,
	peerNodeIPs *utilip.DualStackIPs,
	ipsecTunOFPort uint32,
	remoteGatewayMAC net.HardwareAddr,
) error {
	_ = "STUB: not implemented"
	return nil
}

// When IPsec is enabled, prioritize using the Node's IPv4 address for the tunnel endpoint.
// In dual-stack clusters, IPv6 traffic is encapsulated in IPv4 and transmitted through
// the IPsec tunnel.

// Since broadcast is not supported in IPv6, ARP should happen only with IPv4 address, and ARP responder flows
// only work for IPv4 addresses.
// arpResponderFlow() adds a flow to resolve peer gateway IPs to GlobalVirtualMAC.
// This flow replies to ARP requests sent from the local gateway asking for the MAC address of a remote peer gateway. It ensures that the local Node can reach any remote Pod.

// peerNodeIP is the peer Node's transport address. In a dual-stack setup without
// IPsec enabled, each Node has 2 transport addresses (IPv4 and IPv6). With IPsec
// enabled, we always use the IPv4 address for tunneling encrypted traffic between Nodes.

// Flow to forward the reply packets of Egress connections, whose request packets came from remote Pods
// via tunnel, back to those Pods via tunnel, ensuring symmetric paths of the connections. This flow is
// needed when Egress uses a tunnel path distinct from the common Pod-to-Pod path (hybrid or WireGuard)
// and the peer is reachable via routing.

// flow to catch traffic from AntreaFlexibleIPAM Pod to remote Per-Node IPAM Pod

// When IPsec tunnel is enabled, packets received from the remote Node are
// input from the Node's IPsec tunnel port, not the default tunnel port. So,
// add a separate tunnelClassifierFlow for the IPsec tunnel port.

// For Windows Noencap Mode, the OVS flows for Node need to be exactly same as the provided 'flows' slice because
// the Node flows may be processed more than once if the MAC annotation is updated.

func (c *client) UninstallNodeFlows(hostname string) error { _ = "STUB: not implemented"; return nil }

func (c *client) InstallPodFlows(interfaceName string, podInterfaceIPs []net.IP, podInterfaceMAC net.HardwareAddr, ofPort uint32, vlanID uint16, labelID *uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(gran): support IPv6

// Add support for IPv4 ARP responder.

// Add IP SpoofGuard flows for all validate IPs.

// Add L3 Routing flows to rewrite Pod's dst MAC for all validate IPs.

// In policy-only mode, traffic to local Pod is routed based on destination IP.

// Add Pod uplink classifier flows for AntreaFlexibleIPAM Pods.

// Multicast pod statistics is currently only supported for pods running IPv4 address.

func (c *client) installMulticastPodMetricFlows(interfaceName string, podIP net.IP, ofPort uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallPodFlows(interfaceName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) getFlowKeysFromCache(cache *flowCategoryCache, cacheKey string) []string {
	_ = "STUB: not implemented"
	return nil
}

// ReplayFlows() could change Flow internal state. Although its current
// implementation does not impact Flow match string generation, we still
// acquire read lock of replayMutex here for logic cleanliness.

func (c *client) GetPodFlowKeys(interfaceName string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) InstallServiceGroup(groupID binding.GroupIDType, withSessionAffinity bool, endpoints []proxy.Endpoint) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallServiceGroup(groupID binding.GroupIDType) error {
	_ = "STUB: not implemented"
	return nil
}

func generateEndpointFlowCacheKey(endpointIP string, endpointPort int, protocol binding.Protocol) string {
	_ = "STUB: not implemented"
	return ""
}

func generateServicePortFlowCacheKey(svcIP net.IP, svcPort uint16, protocol binding.Protocol) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *client) InstallEndpointFlows(protocol binding.Protocol, endpoints []proxy.Endpoint) error {
	_ = "STUB: not implemented"
	return nil
}

// keyToFlows is a map from the flows' cache key to the flows.

func (c *client) UninstallEndpointFlows(protocol binding.Protocol, endpoints []proxy.Endpoint) error {
	_ = "STUB: not implemented"
	return nil
}

// keyToFlows is a map from the flows' cache key to the flows.

func (c *client) InstallServiceFlows(config *types.ServiceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Currently, this flow is only used in multi-cluster.

func (c *client) UninstallServiceFlows(svcIP net.IP, svcPort uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) GetServiceFlowKeys(svcIP net.IP, svcPort uint16, protocol binding.Protocol, endpoints []proxy.Endpoint) []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) initialize() error {
	_ = "STUB: not implemented"
	// After a connection or re-connection, delete all existing group and meter entries, to
	// avoid "already exist" errors. This will typically happen if the antrea-agent container is
	// restarted (but not the antrea-ovs one). We do this in initialize(), and not directly in
	// Initialize(), to ensure that the deletion happen on every re-connection (when
	// ReplayFlows() is called), even though we typically only see reconnections when the OVS
	// daemons are restarted. When ovs-vswitchd restarts, group and meter entries are empty by
	// default and these calls are not required.
	// This is specific to groups and meters. Flows are replayed with a different cookie number
	// and conflicts are not possible.
	return nil
}

func (c *client) Initialize(roundInfo types.RoundInfo,
	nodeConfig *config.NodeConfig,
	networkConfig *config.NetworkConfig,
	egressConfig *config.EgressConfig,
	serviceConfig *config.ServiceConfig,
	l7NetworkPolicyConfig *config.L7NetworkPolicyConfig) (<-chan struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiate connections to target OFswitch, and create tables on the switch.

// Ignore first notification, it is not a "reconnection".

// In the normal case, there should be no existing flows with the current round number. This
// is needed in case the agent was restarted before we had a chance to increment the round
// number (incrementing the round number happens once we are satisfied that stale flows from
// the previous round have been deleted).

// generatePipelines generates table list for every pipeline from all activated features. Note that, tables are not realized
// in OVS bridge in this function.
func (c *client) generatePipelines() { _ = "STUB: not implemented"; return }

// TODO: add support for IPv6 protocol

// Pipelines to generate.

// For every pipeline, get required tables from every active feature and store the required tables in a map to avoid
// duplication.

// Iterate the table order cache to generate a sorted table list with required tables.

// generate a pipeline from the required table list.

func (c *client) InstallSNATBypassServiceFlows(serviceCIDRs []*net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) InstallSNATMarkFlows(snatIP net.IP, mark uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallSNATMarkFlows(mark uint32) error { _ = "STUB: not implemented"; return nil }

func (c *client) InstallPodSNATFlows(ofPort uint32, snatIP net.IP, snatMark uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallPodSNATFlows(ofPort uint32) error { _ = "STUB: not implemented"; return nil }

func (c *client) InstallEgressQoS(meterID, rate, burst uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// Install Egress QoS meter.

// Install Egress QoS flow.

func (c *client) UninstallEgressQoS(meterID uint32) error { _ = "STUB: not implemented"; return nil }

// Uninstall Egress QoS flow.

// Uninstall Egress QoS meter.

func (c *client) ReplayFlows() { _ = "STUB: not implemented"; return }

// Openflow bundle message doesn't support meter. Add meter individually instead of
// calling AddOFEntries function.

func (c *client) deleteFlowsByRoundNum(roundNum uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) DeleteStaleFlows() error { _ = "STUB: not implemented"; return nil }

func (c *client) SubscribePacketIn(category uint8, pktInQueue *binding.PacketInQueue) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) SendTraceflowPacket(dataplaneTag uint8, packet *binding.Packet, inPort uint32, outPort int32) error {
	_ = "STUB: not implemented"
	return nil
}

// Set ethernet header

// Set IP header

// Set transport header

// #nosec G404: random number generator not used for security purposes.

// #nosec G404: random number generator not used for security purposes.

func (c *client) InstallTraceflowFlows(dataplaneTag uint8, liveTraffic, droppedOnly, receiverOnly bool, packet *binding.Packet, ofPort uint32, timeoutSeconds uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallTraceflowFlows(dataplaneTag uint8) error {
	_ = "STUB: not implemented"
	return nil
}

// setBasePacketOutBuilder sets base IP properties of a packetOutBuilder which can have more packet data added.
func setBasePacketOutBuilder(packetOutBuilder binding.PacketOutBuilder, srcMAC string, dstMAC string, srcIP string, dstIP string, inPort uint32, outPort uint32) (binding.PacketOutBuilder, error) {
	_ = "STUB: not implemented"
	// Set ethernet header.
	return *new(binding.PacketOutBuilder), nil
}

// Set IP header.

// SendTCPPacketOut generates TCP packet as a packet-out and sends it to OVS.
func (c *client) SendTCPPacketOut(
	srcMAC string,
	dstMAC string,
	srcIP string,
	dstIP string,
	inPort uint32,
	outPort uint32,
	isIPv6 bool,
	tcpSrcPort uint16,
	tcpDstPort uint16,
	tcpSeqNum uint32,
	tcpAckNum uint32,
	tcpHdrLen uint8,
	tcpFlag uint8,
	tcpWinSize uint16,
	tcpData []byte,
	mutatePacketOut func(builder binding.PacketOutBuilder) binding.PacketOutBuilder) error {
	_ = "STUB: not implemented"
	// Generate a base IP PacketOutBuilder.
	return nil
}

// Set protocol.

// Set TCP header data.

// SendICMPPacketOut generates ICMP packet as a packet-out and send it to OVS.
func (c *client) SendICMPPacketOut(
	srcMAC string,
	dstMAC string,
	srcIP string,
	dstIP string,
	inPort uint32,
	outPort uint32,
	isIPv6 bool,
	icmpType uint8,
	icmpCode uint8,
	icmpData []byte,
	mutatePacketOut func(builder binding.PacketOutBuilder) binding.PacketOutBuilder) error {
	_ = "STUB: not implemented"
	// Generate a base IP PacketOutBuilder.
	return nil
}

// Set protocol.

// Set ICMP header data.

// SendUDPPacketOut generates UDP packet as a packet-out and sends it to OVS.
func (c *client) SendUDPPacketOut(
	srcMAC string,
	dstMAC string,
	srcIP string,
	dstIP string,
	inPort uint32,
	outPort uint32,
	isIPv6 bool,
	udpSrcPort uint16,
	udpDstPort uint16,
	udpData []byte,
	mutatePacketOut func(builder binding.PacketOutBuilder) binding.PacketOutBuilder) error {
	_ = "STUB: not implemented"
	// Generate a base IP PacketOutBuilder.
	return nil
}

// Set protocol.

// Set UDP header data.

func (c *client) ResumePausePacket(packetIn *ofctrl.PacketIn) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) SendEthPacketOut(inPort, outPort uint32, ethPkt *protocol.Ethernet, mutatePacketOut func(builder binding.PacketOutBuilder) binding.PacketOutBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) InstallMulticastFlows(multicastIP net.IP, groupID binding.GroupIDType) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallMulticastFlows(multicastIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) InstallMulticastFlexibleIPAMFlows() error { _ = "STUB: not implemented"; return nil }

func (c *client) InstallMulticastRemoteReportFlows(groupID binding.GroupIDType) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) SendIGMPQueryPacketOut(
	dstMAC net.HardwareAddr,
	dstIP net.IP,
	outPort uint32,
	igmp ofutil.Message) error {
	_ = "STUB: not implemented"
	// Generate a base IP PacketOutBuilder.
	return nil
}

// Set protocol and L4 message.

func (c *client) InstallTrafficControlMarkFlows(name string,
	sourceOFPorts []uint32,
	targetOFPort uint32,
	direction crdv1alpha2.Direction,
	action crdv1alpha2.TrafficControlAction,
	priority types.TrafficControlFlowPriority) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallTrafficControlMarkFlows(name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) InstallTrafficControlReturnPortFlow(returnOFPort uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallTrafficControlReturnPortFlow(returnOFPort uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) SendIGMPRemoteReportPacketOut(
	dstMAC net.HardwareAddr,
	dstIP net.IP,
	igmp ofutil.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// Set protocol, L4 message, and target OF Group ID.

func (c *client) InstallMulticastGroup(groupID binding.GroupIDType, localReceivers []uint32, remoteNodeReceivers []net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallMulticastGroup(groupID binding.GroupIDType) error {
	_ = "STUB: not implemented"
	return nil
}

// InstallMulticlusterNodeFlows installs flows to handle cross-cluster packets between a regular
// Node and a local Gateway.
func (c *client) InstallMulticlusterNodeFlows(clusterID string,
	peerConfigs map[*net.IPNet]net.IP,
	tunnelPeerIP net.IP,
	enableStretchedNetworkPolicy bool) error {
	_ = "STUB: not implemented"
	return nil
}

// InstallMulticlusterGatewayFlows installs flows to handle cross-cluster packets between Gateways.
func (c *client) InstallMulticlusterGatewayFlows(clusterID string,
	peerConfigs map[*net.IPNet]net.IP,
	tunnelPeerIP net.IP,
	localGatewayIP net.IP,
	enableStretchedNetworkPolicy bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Add SNAT flows to change cross-cluster packets' source IP to local Gateway IP.

// InstallMulticlusterClassifierFlows adds the following flows:
//   - One flow in L2ForwardingCalcTable for the global virtual multicluster MAC 'aa:bb:cc:dd:ee:f0'
//     to set its target output port as 'antrea-tun0'. This flow will be on both Gateway and regular Node.
//   - One flow in ClassifierTable for the tunnel traffic if it's not Encap mode.
//   - One flow to match MC virtual MAC 'aa:bb:cc:dd:ee:f0' in ClassifierTable for Gateway only.
//   - One flow in OutputTable to allow multicluster hairpin traffic for Gateway only.
func (c *client) InstallMulticlusterClassifierFlows(tunnelOFPort uint32, isGateway bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) InstallMulticlusterPodFlows(podIP net.IP, tunnelPeerIP net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallMulticlusterFlows(clusterID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallMulticlusterPodFlows(podIP string) error {
	_ = "STUB: not implemented"
	return nil
}

// Clean up all flows.

func GetFlowModMessages(flows []binding.Flow, op binding.OFOperation) []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

func getFlowModMessage(flow binding.Flow, op binding.OFOperation) *openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

// getMeterStats sends a multipart request to get all the meter statistics and
// sets values for antrea_agent_ovs_meter_packet_dropped_count.
func (c *client) getMeterStats() { _ = "STUB: not implemented"; return }

// Log an error if dropped packets increased in the last round.

func (c *client) SubscribeOFPortStatusMessage(statusCh chan *openflow15.PortStatus) {
	_ = "STUB: not implemented"
	return
}

// InstallL7NetworkPolicyFlows will be called only when at least one L7 NetworkPolicy is applied locally.
func (c *client) InstallL7NetworkPolicyFlows() error { _ = "STUB: not implemented"; return nil }
