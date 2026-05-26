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
	"sync"
	"sync/atomic"

	"antrea.io/libOpenflow/openflow15"
	"antrea.io/ofnet/ofctrl"
	"k8s.io/client-go/tools/cache"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/nodeip"
	"antrea.io/antrea/v2/pkg/agent/openflow/cookie"
	"antrea.io/antrea/v2/pkg/agent/openflow/operations"
	"antrea.io/antrea/v2/pkg/agent/types"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
	"antrea.io/antrea/v2/third_party/proxy"
)

var (
	//      _   _   _             _   _               _
	//     / \ | |_| |_ ___ _ __ | |_(_) ___  _ __   | |
	//    / _ \| __| __/ _ \ '_ \| __| |/ _ \| '_ \  | |
	//   / ___ \ |_| ||  __/ | | | |_| | (_) | | | | |_|
	//  /_/   \_\__|\__\___|_| |_|\__|_|\___/|_| |_| (_)
	//
	// Before adding a new table in FlexiblePipeline, please read the following instructions carefully.
	//
	// - Double confirm the necessity of adding a new table, and consider reusing an existing table to implement the
	//   functionality alternatively.
	// - Choose a name that can help users to understand the function of the table.
	// - Choose a stage. Existing stageIDs are defined in file pkg/agent/openflow/framework.go. If you want to add a new
	//   stage, please discuss with maintainers or OVS pipeline developers of Antrea.
	// - Choose a pipeline. Existing pipelineIDs are defined in file pkg/agent/openflow/framework.go. If you want to add
	//   a new pipeline, please discuss with maintainers or OVS pipeline developers of Antrea.
	// - Decide where to add the new table in the pipeline. The order table declaration decides the order of tables in the
	//   stage. For example:
	//     * If you want to add a table called `FooTable` between `SpoofGuardTable` and `IPv6Table` in pipelineIP, then
	//       the table should be declared after `SpoofGuardTable` and before `IPv6Table`:
	//       ```go
	//          SpoofGuardTable  = newTable("SpoofGuard", stageValidation, pipelineIP)
	//          FooTable         = newTable("Foo", stageValidation, pipelineIP)
	//          IPv6Table        = newTable("IPv6", stageValidation, pipelineIP)
	//       ```
	//      * If you want to add a table called `FooTable` just before `ARPResponderTable` in pipelineARP, then the table
	//        should be declared before `ARPResponderTable`:
	//       ```go
	//          FooTable          = newTable("Foo", stageOutput, binding.PipelineARP)
	//          ARPResponderTable = newTable("ARPResponder", stageOutput, binding.PipelineARP)
	//       ```
	//       * If you want to add a table called `FooTable` just after `ConntrackStateTable` in pipelineARP, then the
	//         table should be declared after `ConntrackStateTable`:
	//       ```go
	//          UnSNATTable         = newTable("UnSNAT", stageConntrackState, pipelineIP)
	//          ConntrackTable      = newTable("ConntrackZone", stageConntrackState, pipelineIP)
	//          ConntrackStateTable = newTable("ConntrackState", stageConntrackState, pipelineIP)
	//          FooTable            = newTable("Foo", stageConntrackState, pipelineIP)
	//       ```
	//  - Reference the new table in a feature in file pkg/agent/openflow/framework.go. The table can be referenced by multiple
	//    features if multiple features need to install flows in the table. Note that, if the newly added table is not
	//    referenced by any feature or the features referencing the table are all inactivated, then the table will not
	//    be realized in OVS; if at least one feature referencing the table is activated, then the table will be realized
	//    at the desired position in OVS pipeline.
	//  - By default, the miss action of the new table is to forward packets to next table. If the miss action needs to
	//    drop packets, add argument defaultDrop when creating the new table.
	//
	// How to forward packet between tables with a proper action in FlexiblePipeline?
	//
	// |   table A   | |   table B   | |   table C   | |   table D   | |   table E   | |   table F   | |   table G   |
	// |   stage S1  | |                           stage S2                          | |          stage S4           |
	//
	//  - NextTable is used to forward packets to the next table. E.g. A -> B, B -> C, C -> D, etc.
	//  - GotoTable is used to forward packets to a specific table, and the target table ID should be greater than the
	//    current table ID. Within a stage, GotoTable should be used to forward packets to a specific table, e.g. B -> D,
	//    C -> E. Today we do not have the case, but if in future there is a case that a packet needs to be forwarded to
	//    a table in another stage directly, e.g. A -> C, B -> G, GotoTable can also be used.
	//  - GotoStage is used to forward packets to a specific stage. Note that, packets are forwarded to the first table of
	//    the target stage, and the first table ID of the target stage should be greater than the current table ID. E.g.
	//    A -> S4 (F), D -> S4 (F) are fine, but D -> S1 (A), F -> S2 (B) are not allowed. It is recommended to use
	//    GotoStage to forward packets across stages.
	//  - ResubmitToTables is used to forward packets to one or multiple tables. It should be used only when the target
	//    table ID is smaller than the current table ID, like E -> B; or when forwarding packets to multiple tables,
	//    like B - > D E; otherwise, in all other cases GotoTable should be used.

	// Tables of PipelineRoot are declared below.

	// PipelineRootClassifierTable is the only table of pipelineRoot at this moment and its table ID should be 0. Packets
	// are forwarded to pipelineIP or pipelineARP in this table.
	PipelineRootClassifierTable = newTable("PipelineRootClassifier", stageStart, pipelineRoot, defaultDrop)

	// Tables of pipelineARP are declared below.

	// Tables in stageValidation:
	ARPSpoofGuardTable = newTable("ARPSpoofGuard", stageValidation, pipelineARP, defaultDrop)

	// Tables in stageOutput:
	ARPResponderTable = newTable("ARPResponder", stageOutput, pipelineARP)

	// Tables of pipelineIP are declared below.

	// Tables in stageClassifier:
	ClassifierTable = newTable("Classifier", stageClassifier, pipelineIP, defaultDrop)

	// Tables in stageValidation:
	SpoofGuardTable           = newTable("SpoofGuard", stageValidation, pipelineIP, defaultDrop)
	IPv6Table                 = newTable("IPv6", stageValidation, pipelineIP)
	PipelineIPClassifierTable = newTable("PipelineIPClassifier", stageValidation, pipelineIP)

	// Tables in stageConntrackState:
	UnSNATTable         = newTable("UnSNAT", stageConntrackState, pipelineIP)
	ConntrackTable      = newTable("ConntrackZone", stageConntrackState, pipelineIP)
	ConntrackStateTable = newTable("ConntrackState", stageConntrackState, pipelineIP)

	// Tables in stagePreRouting:
	// When proxy is enabled.
	PreRoutingClassifierTable = newTable("PreRoutingClassifier", stagePreRouting, pipelineIP)
	NodePortMarkTable         = newTable("NodePortMark", stagePreRouting, pipelineIP)
	SessionAffinityTable      = newTable("SessionAffinity", stagePreRouting, pipelineIP)
	ServiceLBTable            = newTable("ServiceLB", stagePreRouting, pipelineIP)
	DSRServiceMarkTable       = newTable("DSRServiceMark", stagePreRouting, pipelineIP)
	EndpointDNATTable         = newTable("EndpointDNAT", stagePreRouting, pipelineIP)
	// When proxy is disabled.
	DNATTable = newTable("DNAT", stagePreRouting, pipelineIP)

	// Tables in stageEgressSecurity:
	EgressSecurityClassifierTable = newTable("EgressSecurityClassifier", stageEgressSecurity, pipelineIP)
	AntreaPolicyEgressRuleTable   = newTable("AntreaPolicyEgressRule", stageEgressSecurity, pipelineIP)
	EgressRuleTable               = newTable("EgressRule", stageEgressSecurity, pipelineIP)
	EgressDefaultTable            = newTable("EgressDefaultRule", stageEgressSecurity, pipelineIP)
	EgressMetricTable             = newTable("EgressMetric", stageEgressSecurity, pipelineIP)

	// Tables in stageRouting:
	L3ForwardingTable = newTable("L3Forwarding", stageRouting, pipelineIP)
	EgressMarkTable   = newTable("EgressMark", stageRouting, pipelineIP)
	EgressQoSTable    = newTable("EgressQoS", stageRouting, pipelineIP)
	L3DecTTLTable     = newTable("L3DecTTL", stageRouting, pipelineIP)

	// Tables in stagePostRouting:
	SNATMarkTable = newTable("SNATMark", stagePostRouting, pipelineIP)
	SNATTable     = newTable("SNAT", stagePostRouting, pipelineIP)

	// Tables in stageSwitching:
	L2ForwardingCalcTable = newTable("L2ForwardingCalc", stageSwitching, pipelineIP)
	TrafficControlTable   = newTable("TrafficControl", stageSwitching, pipelineIP)

	// Tables in stageIngressSecurity:
	IngressSecurityClassifierTable = newTable("IngressSecurityClassifier", stageIngressSecurity, pipelineIP)
	AntreaPolicyIngressRuleTable   = newTable("AntreaPolicyIngressRule", stageIngressSecurity, pipelineIP)
	IngressRuleTable               = newTable("IngressRule", stageIngressSecurity, pipelineIP)
	IngressDefaultTable            = newTable("IngressDefaultRule", stageIngressSecurity, pipelineIP)
	IngressMetricTable             = newTable("IngressMetric", stageIngressSecurity, pipelineIP)

	// Tables in stageConntrack:
	ConntrackCommitTable = newTable("ConntrackCommit", stageConntrack, pipelineIP)

	// Tables in stageOutput:
	VLANTable   = newTable("VLAN", stageOutput, pipelineIP)
	OutputTable = newTable("Output", stageOutput, pipelineIP)

	// Tables of pipelineMulticast are declared below. Do don't declare any tables of other pipelines here!
	// Tables in stageEgressSecurity:
	// Since IGMP Egress rules only support IGMP report which is handled by packetIn, it is not necessary to add
	// MulticastIGMPEgressMetricTable here.
	MulticastEgressRuleTable   = newTable("MulticastEgressRule", stageEgressSecurity, pipelineMulticast)
	MulticastEgressMetricTable = newTable("MulticastEgressMetric", stageEgressSecurity, pipelineMulticast)

	MulticastEgressPodMetricTable = newTable("MulticastEgressPodMetric", stageEgressSecurity, pipelineMulticast)

	// Tables in stageRouting:
	MulticastRoutingTable = newTable("MulticastRouting", stageRouting, pipelineMulticast)
	// Tables in stageIngressSecurity
	MulticastIngressRuleTable      = newTable("MulticastIngressRule", stageIngressSecurity, pipelineMulticast)
	MulticastIngressMetricTable    = newTable("MulticastIngressMetric", stageIngressSecurity, pipelineMulticast)
	MulticastIngressPodMetricTable = newTable("MulticastIngressPodMetric", stageIngressSecurity, pipelineMulticast)
	// Tables in stageOutput
	MulticastOutputTable = newTable("MulticastOutput", stageOutput, pipelineMulticast)

	// NonIPTable is used when Antrea Agent is running on an external Node. It forwards the non-IP packet
	// between the uplink and its pair port directly.
	NonIPTable = newTable("NonIP", stageClassifier, pipelineNonIP, defaultDrop)

	// Flow priority level
	priorityHigh            = uint16(210)
	priorityNormal          = uint16(200)
	priorityLow             = uint16(190)
	priorityMiss            = uint16(0)
	priorityTopAntreaPolicy = uint16(64990)
	priorityDNSIntercept    = uint16(64991)

	// Index for priority cache
	priorityIndex = "priority"

	// IPv6 multicast prefix
	ipv6MulticastAddr = "FF00::/8"
	// IPv6 link-local prefix
	ipv6LinkLocalAddr = "FE80::/10"

	// Operation field values in ARP packets
	arpOpRequest = uint16(1)
	arpOpReply   = uint16(2)

	tableNameIndex = "tableNameIndex"
)

// tableCache caches the OpenFlow tables used in pipelines, and it supports using the table ID and name as the index to query the OpenFlow table.
var tableCache = cache.NewIndexer(tableIDKeyFunc, cache.Indexers{tableNameIndex: tableNameIndexFunc})

func tableNameIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tableIDKeyFunc(obj interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getTableByID(id uint8) binding.Table { _ = "STUB: not implemented"; return *new(binding.Table) }

// GetFlowTableName returns the flow table name given the table ID. An empty
// string is returned if the table cannot be found.
func GetFlowTableName(tableID uint8) string { _ = "STUB: not implemented"; return "" }

// GetFlowTableID does a case insensitive lookup of the table name, and
// returns the flow table number if the table is found. Otherwise TableIDAll is
// returned if the table cannot be found.
func GetFlowTableID(tableName string) uint8 { _ = "STUB: not implemented"; return 0 }

func GetTableList() []binding.Table { _ = "STUB: not implemented"; return nil }

func GetAntreaPolicyEgressTables() []*Table { _ = "STUB: not implemented"; return nil }

func GetAntreaIGMPIngressTables() []*Table { _ = "STUB: not implemented"; return nil }

func GetAntreaMulticastEgressTables() []*Table { _ = "STUB: not implemented"; return nil }

func GetAntreaPolicyIngressTables() []*Table { _ = "STUB: not implemented"; return nil }

func GetAntreaPolicyBaselineTierTables() []*Table { _ = "STUB: not implemented"; return nil }

func GetAntreaPolicyMultiTierTables() []*Table { _ = "STUB: not implemented"; return nil }

const (
	CtZone       = 0xfff0
	CtZoneV6     = 0xffe6
	SNATCtZone   = 0xfff1
	SNATCtZoneV6 = 0xffe7

	// disposition values used in AP
	DispositionAllow = 0b00
	DispositionDrop  = 0b01
	DispositionRej   = 0b10
	DispositionPass  = 0b11
	// DispositionL7NPRedirect is used when sending packet-in to controller for
	// logging layer 7 NetworkPolicy indicating that this packet is redirected to
	// l7 engine to determine the disposition.
	DispositionL7NPRedirect = 0b1

	// EtherTypeDot1q is used when adding 802.1Q VLAN header in OVS action
	EtherTypeDot1q = 0x8100

	// dsrServiceConnectionIdleTimeout represents the idle timeout of the flows learned for DSR Service.
	// 160 means the learned flows will be deleted if the flow is not used in 160s.
	// It tolerates 1 keep-alive drop (net.ipv4.tcp_keepalive_intvl defaults to 75) and a deviation of 10s for long connections.
	dsrServiceConnectionIdleTimeout = 160
	// dsrServiceConnectionFinIdleTimeout represents the idle timeout of the flows learned for DSR Service after a TCP
	// packet with the FIN or RST flag is received.
	dsrServiceConnectionFinIdleTimeout = 5
)

var DispositionToString = map[uint32]string{
	DispositionAllow: "Allow",
	DispositionDrop:  "Drop",
	DispositionRej:   "Reject",
	DispositionPass:  "Pass",
}

var (
	// snatPktMarkRange takes an 8-bit range of pkt_mark to store the ID of
	// a SNAT IP. The bit range must match SNATIPMarkMask.
	snatPktMarkRange = &binding.Range{0, 7}

	GlobalVirtualMAC, _ = net.ParseMAC("aa:bb:cc:dd:ee:ff")
)

func copyFlowWithNewPriority(flowMod *openflow15.FlowMod, priority uint16) *openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

func flowMessageMatched(oldFlow, newFlow *openflow15.FlowMod) bool {
	_ = "STUB: not implemented"
	return false
}

// isDropFlow returns true if no instructions are defined in the OpenFlow modification message.
// According to the OpenFlow spec, there is no explicit action to represent drops. Instead, the action of dropping
// packets could come from empty instruction sets.
func isDropFlow(f *openflow15.FlowMod) bool { _ = "STUB: not implemented"; return false }

func getFlowModKey(fm *openflow15.FlowMod) string { _ = "STUB: not implemented"; return "" }

func getFlowDumpKey(fm *openflow15.FlowMod) string { _ = "STUB: not implemented"; return "" }

type flowMessageCache map[string]*openflow15.FlowMod

type flowCategoryCache struct {
	sync.Map
}

type client struct {
	enableProxy                bool
	proxyAll                   bool
	enableDSR                  bool
	enableAntreaPolicy         bool
	enableL7NetworkPolicy      bool
	enableDenyTracking         bool
	enableEgress               bool
	enableEgressTrafficShaping bool
	enableMulticast            bool
	enableTrafficControl       bool
	enableMulticluster         bool
	enablePrometheusMetrics    bool
	connectUplinkToBridge      bool
	nodeType                   config.NodeType
	roundInfo                  types.RoundInfo
	cookieAllocator            cookie.Allocator
	bridge                     binding.Bridge
	groupIDAllocator           GroupAllocator

	featurePodConnectivity          *featurePodConnectivity
	featureService                  *featureService
	featureEgress                   *featureEgress
	featureNetworkPolicy            *featureNetworkPolicy
	featureMulticast                *featureMulticast
	featureMulticluster             *featureMulticluster
	featureExternalNodeConnectivity *featureExternalNodeConnectivity
	activatedFeatures               []feature

	featureTraceflow  *featureTraceflow
	traceableFeatures []traceableFeature

	pipelines map[binding.PipelineID]binding.Pipeline

	// ofEntryOperations is a wrapper interface for operating multiple OpenFlow entries with action AddAll / ModifyAll / DeleteAll.
	// It enables convenient mocking in unit tests.
	ofEntryOperations operations.OFEntryOperations
	// replayMutex provides exclusive access to the OFSwitch to the ReplayFlows method.
	replayMutex           sync.RWMutex
	nodeConfig            *config.NodeConfig
	networkConfig         *config.NetworkConfig
	egressConfig          *config.EgressConfig
	serviceConfig         *config.ServiceConfig
	l7NetworkPolicyConfig *config.L7NetworkPolicyConfig
	// ovsMetersAreSupported indicates whether the OVS datapath supports OpenFlow meters.
	ovsMetersAreSupported bool
	// ovsMeterPacketDrops tracks the number of packets dropped by each OVS meter, keyed by meter ID.
	ovsMeterPacketDrops map[int]*atomic.Int64
	// packetInRate defines the OVS controller packet rate limits for different
	// features. All features will apply this rate-limit individually on packet-in
	// messages sent to antrea-agent. The number stands for the rate as packets per
	// second(pps) and the burst size will be automatically set to twice the rate.
	// When the rate and burst size are exceeded, new packets will be dropped.
	packetInRate int
	// packetInHandlers stores handler to process PacketIn event. When a packetIn
	// arrives, openflow send packet to registered handler in this map.
	packetInHandlers map[uint8]PacketInHandler
	// Supported IP Protocols (IP or IPv6) on the current Node.
	ipProtocols []binding.Protocol
	// ovsctlClient is the interface for executing OVS "ovs-ofctl" and "ovs-appctl" commands.
	ovsctlClient ovsctl.OVSCtlClient

	nodeIPChecker nodeip.Checker
}

func (c *client) Run(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// Start PacketIn
	return
}

// Start OVS meter stats collection

func (c *client) GetTunnelVirtualMAC() net.HardwareAddr {
	_ = "STUB: not implemented"
	return *new(net.HardwareAddr)
}

func (c *client) defaultFlows() []*openflow15.FlowMod { _ = "STUB: not implemented"; return nil }

// This generates the default flow for every table in every pipeline.

// This generates the flow to match IPv4 / IPv6 packets and forward them to the first table of pipelineIP in
// PipelineRootClassifierTable.

// This generates the flow to match ARP packets and forward them to the first table of pipelineARP in
// PipelineRootClassifierTable.

// This generates the flow to match multicast packets and forward them to the first table of pipelineMulticast
// in PipelineIPClassifierTable. Note that, PipelineIPClassifierTable is in stageValidation of pipeline for IP. In another word,
// pipelineMulticast is forked from PipelineIPClassifierTable in pipelineIP.

// tunnelClassifierFlow generates the flow to mark the packets from tunnel port.
func (f *featurePodConnectivity) tunnelClassifierFlow(tunnelOFPort uint32) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// gatewayClassifierFlows generates the flow to mark the packets from the Antrea gateway port.
func (f *featurePodConnectivity) gatewayClassifierFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// If the packet is from gateway but its source IP is not the gateway IP, it's considered external sourced traffic.

// podClassifierFlow generates the flow to mark the packets from a local Pod port.
// If multi-cluster is enabled, also load podLabelID into LabelIDField.
func (f *featurePodConnectivity) podClassifierFlow(podOFPort uint32, isAntreaFlexibleIPAM bool, podLabelID *uint32) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// podUplinkClassifierFlows generates the flows to mark the packets with target destination MAC address from uplink/bridge
// port, which are needed when uplink is connected to OVS bridge and Antrea IPAM is configured.
func (f *featurePodConnectivity) podUplinkClassifierFlows(dstMAC net.HardwareAddr, vlanID uint16) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to mark the packets from uplink port.

// This generates the flow to mark the packets from bridge local port.

// conntrackFlows generates the flows about conntrack for feature PodConnectivity.
func (f *featurePodConnectivity) conntrackFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to transform the destination IP of request packets or source IP of reply packets
// from tracked connections in CT zone.

// This generates the flow to match the packets of tracked non-Service connection and forward them to
// stageEgressSecurity directly to bypass stagePreRouting. The first packet of non-Service connection passes
// through stagePreRouting, and the subsequent packets go to stageEgressSecurity directly.

// This generates the flow to drop invalid packets.

// This flow matches the first packet of all non-SNAT connections to commit them to the main /
// DNAT CtZone and to mark the source of the connection by copying PktSourceField to
// ConnSourceCTMarkField. SNAT connections have already been committed to the main / DNAT
// CtZone, prior to being committed to the SNAT CtZone.
// Note that matching on ct_state=-snat with MatchCTStateSNAT(false) does not work because of
// https://github.com/openvswitch/ovs-issues/issues/370. Matching on ct_zone only supports exact
// match, so it is not ideal.

// This generates default flow to match the first packet of a new connection and forward it to stagePreRouting.

// conntrackFlows generates the flows about conntrack for feature Service.
func (f *featureService) conntrackFlows() []binding.Flow { _ = "STUB: not implemented"; return nil }

// This generates the flow to mark tracked DNATed Service connection with RewriteMACRegMark (load-balanced by
// AntreaProxy) and forward the packets to stageEgressSecurity directly to bypass stagePreRouting.

// If DSR is enabled, traffic working in DSR mode will be in invalid state on ingress Node.
// We forward externally originated packets of invalid connections to stagePreRouting to see if it could be
// marked as DSR traffic. If not, they will be dropped in DSRServiceMarkTable.

// snatConntrackFlows generates the flows about conntrack of SNAT connection for feature Service.
func (f *featureService) snatConntrackFlows() []binding.Flow { _ = "STUB: not implemented"; return nil }

// virtualIP is used as SNAT IP when a request's source IP is gateway IP and we need to forward it back to
// gateway interface to avoid asymmetry path.

// SNAT should be performed for the following connections:
// - Hairpin Service connection initiated through a local Pod, and SNAT should be performed with the Antrea
//   gateway IP.
// - Hairpin Service connection initiated through the Antrea gateway, and SNAT should be performed with a
//   virtual IP.
// - Nodeport / LoadBalancer connection initiated through the Antrea gateway and externalTrafficPolicy is
//   Cluster, if the selected Endpoint is not on local Node, then SNAT should be performed with the Antrea
//   gateway IP.
// Note that, for Service connections that require SNAT, ServiceCTMark is loaded in SNAT CT zone when performing
// SNAT since ServiceCTMark loaded in DNAT CT zone cannot be read in SNAT CT zone. For Service connections,
// ServiceCTMark (loaded in DNAT / SNAT CT zone) is used to bypass ConntrackCommitTable which is used to commit
// non-Service connections. For hairpin connections, HairpinCTMark is also loaded in SNAT CT zone when performing
// SNAT since HairpinCTMark loaded in DNAT CT zone also cannot be read in SNAT CT zone. HairpinCTMark is used
// to output packets of hairpin connections in OutputTable.

// This generates the flow to match the first packet of hairpin Service connection initiated through the Antrea
// gateway with ConnSNATCTMark and HairpinCTMark, then perform SNAT in SNAT CT zone with a virtual IP.

// This generates the flow to unSNAT reply packets of connections committed in SNAT CT zone by the above flow.

// This generates the flow to match the first packet of hairpin Service connection initiated through a Pod with
// ConnSNATCTMark and HairpinCTMark, then perform SNAT in SNAT CT zone with the Antrea gateway IP.

// This generates the flow to match the first packet of NodePort / LoadBalancer connection (non-hairpin) initiated
// through the Antrea gateway with ConnSNATCTMark, then perform SNAT in SNAT CT zone with the Antrea gateway IP.

// This generates the flow to unSNAT reply packets of connections committed in SNAT CT zone by the above flows.

// This generates the flow to match the subsequent request packets of connection whose first request packet has
// been committed in SNAT CT zone, then commit the packets in SNAT CT zone again to perform SNAT.
// For example:
/*
	* 192.168.77.1 is the IP address of client.
	* 192.168.77.100 is the IP address of K8s Node.
	* 30001 is the NodePort port.
	* 10.10.0.1 is the IP address of Antrea gateway.
	* 10.10.0.3 is the IP of NodePort Service Endpoint.

	* packet 1 (request)
		* client                     192.168.77.1:12345->192.168.77.100:30001
		* CT zone SNAT 65521         192.168.77.1:12345->192.168.77.100:30001
		* CT zone DNAT 65520         192.168.77.1:12345->192.168.77.100:30001
		* CT commit DNAT zone 65520  192.168.77.1:12345->192.168.77.100:30001  =>  192.168.77.1:12345->10.10.0.3:80
		* CT commit SNAT zone 65521  192.168.77.1:12345->10.10.0.3:80          =>  10.10.0.1:12345->10.10.0.3:80
		* output
	  * packet 2 (reply)
		* Pod                         10.10.0.3:80->10.10.0.1:12345
		* CT zone SNAT 65521          10.10.0.3:80->10.10.0.1:12345            =>  10.10.0.3:80->192.168.77.1:12345
		* CT zone DNAT 65520          10.10.0.3:80->192.168.77.1:12345         =>  192.168.77.1:30001->192.168.77.1:12345
		* output
	  * packet 3 (request)
		* client                     192.168.77.1:12345->192.168.77.100:30001
		* CT zone SNAT 65521         192.168.77.1:12345->192.168.77.100:30001
		* CT zone DNAT 65520         192.168.77.1:12345->10.10.0.3:80
		* CT zone SNAT 65521         192.168.77.1:12345->10.10.0.3:80          =>  10.10.0.1:12345->10.10.0.3:80
		* output
	  * packet ...
*/
// As a result, subsequent request packets like packet 3 will only perform SNAT when they pass through SNAT
// CT zone the second time, after they are DNATed in DNAT CT zone.

// TODO: Use DuplicateToBuilder or integrate this function into original one to avoid unexpected difference.
// flowsToTrace generates Traceflow specific flows in the connectionTrackStateTable or L2ForwardingCalcTable for featurePodConnectivity.
// When packet is not provided, the flows bypass the drop flow in conntrackStateFlow to avoid unexpected drop of the
// injected Traceflow packet, and to drop any Traceflow packet that has ct_state +rpl, which may happen when the Traceflow
// request destination is the Node's IP. When packet is provided, a flow is added to mark - the first packet of the first
// connection that matches the provided packet - as the Traceflow packet. The flow is added in connectionTrackStateTable
// when receiverOnly is false and it also matches in_port to be the provided ofPort (the sender Pod); otherwise when
// receiverOnly is true, the flow is added into L2ForwardingCalcTable and matches the destination MAC (the receiver Pod MAC).
func (f *featurePodConnectivity) flowsToTrace(dataplaneTag uint8,
	ovsMetersAreSupported,
	liveTraffic,
	droppedOnly,
	receiverOnly bool,
	packet *binding.Packet,
	ofPort uint32,
	timeout uint16) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// Match transport header

// Add controller actions unless we are capturing only dropped packets.

// Clear the loaded DSCP bits before output.

// Output the packets if traffic mode uses direct routing (noEncap, hybrid, or WireGuard).
// In direct routing modes, packets are routed directly without overlay encapsulation.
// WireGuard mode utilizes direct routing for all encrypted Pod traffic, similar to the
// behavior seen in noEncap and hybrid modes for direct-reachability traffic.

// This generates Traceflow specific flows that outputs traceflow non-hairpin packets to OVS port and Antrea Agent after
// L2 forwarding calculation.

// SendToController and Output if output port is tunnel port.

// For injected packets, SendToController and Output depending on traffic mode if output port is local gateway.
// - In encap mode, a Traceflow packet going out of the gateway port (i.e. exiting the overlay) essentially means
//   that the Traceflow request is complete. only SendToController if output port is local gateway.
// - In direct routing modes (noEncap, hybrid, or WireGuard), inter-Node Pod-to-Pod traffic is expected to go out of
//   the gateway port on the way to its destination.

// Only SendToController if output port is local gateway and destination IP is gateway.

// Only SendToController if output port is Pod port.

// flowsToTrace is used to generate flows for Traceflow in featureService.
func (f *featureService) flowsToTrace(dataplaneTag uint8,
	ovsMetersAreSupported,
	liveTraffic,
	droppedOnly,
	receiverOnly bool,
	packet *binding.Packet,
	ofPort uint32,
	timeout uint16) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// Add controller actions unless we are capturing only dropped packets.

// Clear the loaded DSCP bits before output.

// This generates Traceflow specific flows that outputs hairpin traceflow packets to OVS port and Antrea Agent after
// L2forwarding calculation.

// Only SendToController for hairpin traffic.
// This flow must have higher priority than the one installed by l2ForwardOutputHairpinServiceFlow.

// flowsToTrace is used to generate flows for Traceflow from globalConjMatchFlowCache and policyCache.
func (f *featureNetworkPolicy) flowsToTrace(dataplaneTag uint8,
	ovsMetersAreSupported,
	liveTraffic,
	droppedOnly,
	receiverOnly bool,
	packet *binding.Packet,
	ofPort uint32,
	timeout uint16) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// Copy Antrea NetworkPolicy drop rules.

// Generate both IPv4 and IPv6 flows if the original drop flow doesn't match IP/IPv6.
// DSCP field is in IP/IPv6 headers so IP/IPv6 match is required in a flow.

// l2ForwardCalcFlow generates the flow to match the destination MAC and load the target ofPort to TargetOFPortField.
func (f *featurePodConnectivity) l2ForwardCalcFlow(dstMAC net.HardwareAddr, ofPort uint32) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// l2ForwardOutputHairpinServiceFlow generates the flow to output the packet of hairpin Service connection with IN_PORT
// action. It matches OutputToOFPortRegMark to ensure only packets explicitly marked for output are processed, excluding
// packets intended for packet-in to the controller.
func (f *featureService) l2ForwardOutputHairpinServiceFlow() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// l2ForwardOutputFlow generates the flow to output the packets to target OVS port according to the value of TargetOFPortField.
func (f *featurePodConnectivity) l2ForwardOutputFlow() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// l3FwdFlowToPod generates the flows to match the packets destined for a local Pod. For a per-Node IPAM Pod, the flow
// rewrites destination MAC to the Pod interface's MAC, and rewrites source MAC to Antrea gateway interface's MAC. For
// an Antrea IPAM Pod, the flow only rewrites the destination MAC to the Pod interface's MAC.
func (f *featurePodConnectivity) l3FwdFlowToPod(localGatewayMAC net.HardwareAddr,
	podInterfaceIPs []net.IP,
	podInterfaceMAC net.HardwareAddr,
	isAntreaFlexibleIPAM bool,
	vlanID uint16) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to match the packets destined for a local Antrea IPAM Pod.

// This generates the flow to match the packets with RewriteMACRegMark and destined for a local per-Node IPAM Pod.

// Only overwrite MAC for untagged traffic which destination is a local per-Node IPAM Pod.

// l3FwdFlowRouteToPod generates the flows to match the packets destined for a Pod based on the destination IPs. It rewrites
// destination MAC to the Pod interface's MAC. The flows are only used in networkPolicyOnly mode.
func (f *featurePodConnectivity) l3FwdFlowRouteToPod(podInterfaceIPs []net.IP, podInterfaceMAC net.HardwareAddr) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// l3FwdFlowRouteToGW generates the flows to match the packets destined for the Antrea gateway. It rewrites destination MAC
// to the Antrea gateway interface's MAC. The flows are used in networkPolicyOnly mode to match the packets sourced from a
// local Pod and destined for remote Pods, Nodes, or external network.
func (f *featurePodConnectivity) l3FwdFlowRouteToGW() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// l3FwdFlowToGateway generates the flows to match the packets destined for the Antrea gateway.
func (f *featurePodConnectivity) l3FwdFlowToGateway() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to match the packets destined for Antrea gateway.

// This generates the flow to match the reply packets of connection with FromGatewayCTMark.

// l3FwdFlowsToRemoteViaTun generates the flows to match the packets destined for remote Pods via tunnel.
func (f *featurePodConnectivity) l3FwdFlowsToRemoteViaTun(localGatewayMAC net.HardwareAddr, peerSubnet net.IPNet, tunnelPeer net.IP) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// Rewrite src MAC to local gateway MAC.
// Rewrite dst MAC to virtual MAC.
// Flow based tunnel. Set tunnel destination.

// The flow handles packets whose destination IP is in the peer subnet.

// If DSR is enabled, packets accessing a DSR Service will not be DNATed on the ingress Node, but EndpointIPField
// holds the selected backend Pod IP, we match it and DSRServiceRegMark to send these packets to corresponding Nodes.

// Like matching destination IP, we only check if the prefix of the EndpointIP stored in EndpointIPField is in
// the subnet. For example, if the peerSubnet is 10.10.1.0/24, we will check reg3=0xa0a0100/0xffffff00.

// TODO: MatchXXReg must support mask to support IPv6.

// l3FwdFlowEgressReturnViaTun generates the flow to match reply packets of Egress connections (whose request packets
// came from remote Pods via tunnel) and forward them back to those Pods via tunnel, ensuring symmetric paths. It is
// used when Egress uses a tunnel path distinct from the common Pod-to-Pod path (hybrid or WireGuard) and the peer is
// reachable via routing.
func (f *featurePodConnectivity) l3FwdFlowEgressReturnViaTun(localGatewayMAC net.HardwareAddr, peerSubnet net.IPNet, tunnelPeer net.IP) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// Match packets received on local gateway only, ensuring they are reply Egress packets.
// Match packets from connections originated from tunnel.

// Rewrite src MAC to local gateway MAC.
// Rewrite dst MAC to virtual MAC.
// Flow based tunnel. Set tunnel destination.

// l3FwdFlowToRemoteViaGW generates the flow to match the packets destined for remote Pods via the Antrea gateway. It is
// used when the cross-Node connections that do not require encapsulation (in noEncap, networkPolicyOnly, or hybrid mode).
func (f *featurePodConnectivity) l3FwdFlowToRemoteViaGW(localGatewayMAC net.HardwareAddr, peerSubnet net.IPNet) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// Exclude the packets from Antrea IPAM Pods.

// This generates the flow to match the packets destined for remote Pods. Note that, this flow is installed in Linux Nodes
// or Windows Nodes whose remote Node's transport interface MAC is unknown.

// Traffic to in-cluster destination should skip EgressMark table.

// l3FwdFlowToRemoteViaUplink generates the flow to match the packets destined for remote Pods via uplink. It is used
// when the cross-Node connections that do not require encapsulation (in noEncap, networkPolicyOnly, hybrid mode).
func (f *featurePodConnectivity) l3FwdFlowToRemoteViaUplink(remoteGatewayMAC net.HardwareAddr,
	peerSubnet net.IPNet,
	isAntreaFlexibleIPAM bool) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// This generates the flow to match the packets destined for remote Pods via uplink directly without passing
// through the Antrea gateway by rewriting destination MAC to remote Node Antrea gateway's MAC. Note that,
// this flow is only installed in Windows Nodes。

// This generates the flow to match the packets sourced Antrea IPAM Pods and destined for remote Pods, and rewrite
// the destination MAC to remote Node Antrea gateway's MAC. Note that, this flow is only used in Linux when AntreaIPAM
// is enabled.

// arpResponderFlow generates the flow to reply to the ARP request with a MAC address for the target IP address.
func (f *featurePodConnectivity) arpResponderFlow(ipAddr net.IP, macAddr net.HardwareAddr) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// arpResponderStaticFlow generates the flow to reply to any ARP request with the same global virtual MAC. It is used
// in policy-only mode, where traffic are routed via IP not MAC.
func (f *featurePodConnectivity) arpResponderStaticFlow() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// podIPSpoofGuardFlow generates the flow to check IP packets from local Pods. Packets from the Antrea gateway will not be
// checked, since it might be Pod to Service connection or host namespace connection.
func (f *featurePodConnectivity) podIPSpoofGuardFlow(ifIPs []net.IP, ifMAC net.HardwareAddr, ifOFPort uint32, vlanID uint16) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// - When IPv4 is enabled only, IPv6Table is not initialized. All packets should be forwarded to the next table of
//   SpoofGuardTable.
// - When IPv6 is enabled only, IPv6Table is initialized, and it is the next table of SpoofGuardTable. All packets
//   should be to IPv6Table.
// - When both IPv4 and IPv6 are enabled, IPv4 packets should skip IPv6Table (which is the next table of SpoofGuardTable)
//   to avoid unnecessary overhead.

func getIPProtocol(ip net.IP) binding.Protocol {
	_ = "STUB: not implemented"
	return *new(binding.Protocol)
}

// arpSpoofGuardFlow generates the flow to check the ARP packets sourced from local Pods or the Antrea gateway.
func (f *featurePodConnectivity) arpSpoofGuardFlow(ifIP net.IP, ifMAC net.HardwareAddr, ifOFPort uint32) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// sessionAffinityReselectFlow generates the flow which resubmits the Service accessing packet back to ServiceLBTable
// if there is no endpointDNAT flow matched. This case will occur if an Endpoint is removed and is the learned Endpoint
// selection of the Service.
func (f *featureService) sessionAffinityReselectFlow() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// gatewayIPSpoofGuardFlows generates the flow to skip spoof guard checking for packets from the Antrea gateway.
func (f *featurePodConnectivity) gatewayIPSpoofGuardFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// - When IPv4 is enabled only, IPv6Table is not initialized. All packets should be forwarded to the next table of
//   SpoofGuardTable.
// - When IPv6 is enabled only, IPv6Table is initialized, and it is the next table of SpoofGuardTable. All packets
//   should be to IPv6Table.
// - When both IPv4 and IPv6 are enabled, IPv4 packets should skip IPv6Table (which is the next table of SpoofGuardTable)
//   to avoid unnecessary overhead.

// Set CtZoneTypeField based on ipProtocol and keep VLANIDField=0

// serviceCIDRDNATFlows generates the flows to match destination IP in Service CIDR and output to the Antrea gateway directly.
func (f *featureService) serviceCIDRDNATFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// serviceNeedLBFlow generates the default flow to mark packets with EpToSelectRegMark.
func (f *featureService) serviceNeedLBFlow() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// arpNormalFlow generates the flow to reply to the ARP request packets in normal way if no flow in ARPResponderTable is matched.
func (f *featurePodConnectivity) arpNormalFlow() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

func (f *featureNetworkPolicy) allowRulesMetricFlows(conjunctionID uint32, ingress bool, tableID uint8) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// We use the 0..31 bits of the ct_label to store the ingress rule ID and use the 32..63 bits to store the
// egress rule ID.

// Unlike rules for unicast traffic, each IGMP and multicast rule uses single metric flow to track stats
// in multicast metric tables.

// These two flows track the number of sessions in addition to the packet and byte counts.
// The flow matching 'ct_state=+new' tracks the number of sessions and byte count of the first packet for each
// session.
// The flow matching 'ct_state=-new' tracks the byte/packet count of an established connection (both directions).

func (f *featureNetworkPolicy) denyRuleMetricFlow(conjunctionID uint32, ingress bool, tableID uint8) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// ipv6Flows generates the flows to allow IPv6 packets from link-local addresses and handle multicast packets, Neighbor
// Solicitation and ND Advertisement packets properly.
func (f *featurePodConnectivity) ipv6Flows() []binding.Flow { _ = "STUB: not implemented"; return nil }

// Allow IPv6 packets (e.g. Multicast Listener Report Message V2) which are sent from link-local addresses in
// SpoofGuardTable, so that these packets will not be dropped.

// Handle IPv6 Neighbor Solicitation and Neighbor Advertisement as a regular L2 learning Switch by using normal.

// Handle IPv6 multicast packets as a regular L2 learning Switch by using normal.
// It is used to ensure that all kinds of IPv6 multicast packets are properly handled (e.g. Multicast Listener
// Report Message V2).

// For normal traffic, conjunctionActionFlow generates the flow to jump to a specific table if policyRuleConjunction ID is matched. Priority of
// conjunctionActionFlow is created at priorityLow for k8s network policies, and *priority assigned by PriorityAssigner for AntreaPolicy.
func (f *featureNetworkPolicy) conjunctionActionFlow(conjunctionID uint32, table binding.Table, nextTable uint8, priority *uint16, enableLogging bool, l7RuleVlanID *uint32) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// Traceflow.
// CT action requires commit flag if actions other than NAT without arguments are specified.

// Mark the packets of the connection should be redirected to an application-aware engine.
// Load the VLAN ID allocated for L7 NetworkPolicy rule to CT mark field L7NPRuleVlanIDCTMarkField.

// AntreaPolicy.

// Traceflow.
// CT action requires commit flag if actions other than NAT without arguments are specified.

// AntreaPolicy.

// Traceflow.
// CT action requires commit flag if actions other than NAT without arguments are specified.

// Mark the packets of the connection should be redirected to an application-aware engine.
// Load the VLAN ID allocated for L7 NetworkPolicy rule to CT mark field L7NPRuleVlanIDCTMarkField.

// Traceflow.
// CT action requires commit flag if actions other than NAT without arguments are specified.

// As IGMP and multicast use a different pipeline 'Multicast', if the rule is
// IGMP ingress or multicast egress，conjunctionActionFlow generates the flow
// to mark the packet to be allowed if policyRuleConjunction ID is matched.
// Any matched flow will be resubmitted to next table in corresponding metric tables.

// conjunctionActionDenyFlow generates the flow to mark the packet to be denied (dropped or rejected) if policyRuleConjunction
// ID is matched. Any matched flow will be dropped in corresponding metric tables.
func (f *featureNetworkPolicy) conjunctionActionDenyFlow(conjunctionID uint32, table binding.Table, priority *uint16,
	disposition uint32, enableLogging bool) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// We do not drop the packet immediately but send the packet to the metric table to update the rule metrics.

func (f *featureNetworkPolicy) conjunctionActionPassFlow(conjunctionID uint32, table binding.Table, priority *uint16, enableLogging bool) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

func (c *client) Disconnect() error { _ = "STUB: not implemented"; return nil }

func newFlowCategoryCache() *flowCategoryCache { _ = "STUB: not implemented"; return nil }

func (f *featureNetworkPolicy) addFlowMatch(fb binding.FlowBuilder, matchKey *types.MatchKey, matchValue interface{}) binding.FlowBuilder {
	_ = "STUB: not implemented"
	return *new(binding.FlowBuilder)
}

// ofport number in NXM_NX_REG1 is used in ingress rule to match packets sent to local Pod.

// conjunctionExceptionFlow generates the flow to jump to a specific table if both policyRuleConjunction ID and except address are matched.
// Keeping this for reference to generic exception flow.
// nolint: unused
func (f *featureNetworkPolicy) conjunctionExceptionFlow(conjunctionID uint32, tableID uint8, nextTable uint8, matchKey *types.MatchKey, matchValue interface{}) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// Traceflow.

// conjunctiveMatchFlow generates the flow to set conjunctive actions if the match condition is matched.
func (f *featureNetworkPolicy) conjunctiveMatchFlow(tableID uint8, matchPairs []matchPair, priority *uint16, actions []*conjunctiveAction) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// defaultDropFlow generates the flow to drop packets if the match condition is matched.
func (f *featureNetworkPolicy) defaultDropFlow(table binding.Table, matchPairs []matchPair, enableLogging bool) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// multiClusterNetworkPolicySecurityDropFlow generates the security drop flows for MultiClusterNetworkPolicy.
func (f *featureNetworkPolicy) multiClusterNetworkPolicySecurityDropFlow(table binding.Table, matchPairs []matchPair) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// dnsPacketInFlow generates the flow to send dns response packets of fqdn policy selected Pods to the fqdnController for
// processing.
func (f *featureNetworkPolicy) dnsPacketInFlow(conjunctionID uint32) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// FQDN should pause DNS response packets and send them to the controller. After
// the controller processes DNS response packets, like creating related flows in
// the OVS or no operations are needed, the controller will resume those packets.

// localProbeFlows generates the flows to forward locally generated request packets to stageConntrack directly, bypassing
// ingress rule of Network Policies. The packets are sent by kubelet to probe the liveness/readiness of local Pods.
// On Linux and when OVS kernel datapath is used, the probe packets are identified by matching the HostLocalSourceMark.
// On Windows or when OVS userspace (netdev) datapath is used, we need a different approach because:
//  1. On Windows, kube-proxy userspace mode is used, and currently there is no way to distinguish kubelet generated traffic
//     from kube-proxy proxied traffic.
//  2. pkt_mark field is not properly supported for OVS userspace (netdev) datapath.
//
// When proxyAll is disabled, the probe packets are identified by matching the source IP is the Antrea gateway IP;
// otherwise, the packets are identified by matching both the Antrea gateway IP and NotServiceCTMark. Note that, when
// proxyAll is disabled, currently there is no way to distinguish kubelet generated traffic from kube-proxy proxied traffic
// only by matching the Antrea gateway IP. There is a defect that NodePort Service access by external clients will be
// masqueraded as the Antrea gateway IP to bypass NetworkPolicies. See https://github.com/antrea-io/antrea/issues/280.
func (f *featurePodConnectivity) localProbeFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// ingressClassifierFlows generates the flows to classify the packets from local Pods or the Antrea gateway to different
// tables within stageIngressSecurity.
func (f *featureNetworkPolicy) ingressClassifierFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to match the packets to the Antrea gateway and forward them to IngressMetricTable.

// This generates the flow to match the packets to tunnel and forward them to IngressMetricTable.

// This generates the flow to match the packets to uplink and forward them to IngressMetricTable.

// This generates the flow to match the hairpin service packets and forward them to stageConntrack.

// This generates the flow to match the NodePort Service packets and forward them to AntreaPolicyIngressRuleTable.
// Policies applied on NodePort Service will be enforced in AntreaPolicyIngressRuleTable.

// snatSkipCIDRFlow generates the flow to skip SNAT for connection destined for the provided CIDR.
func (f *featureEgress) snatSkipCIDRFlow(cidr net.IPNet) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// snatSkipNodeFlow generates the flow to skip SNAT for connection destined for the transport IP of a remote Node.
func (f *featureEgress) snatSkipNodeFlow(nodeIP net.IP) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// snatIPFromTunnelFlow generates the flow that marks SNAT packets tunnelled from remote Nodes. The SNAT IP matches the
// packet's tunnel destination IP.
func (f *featureEgress) snatIPFromTunnelFlow(snatIP net.IP, mark uint32) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// To apply rate-limit on all traffic.

// snatRuleFlow generates the flow that applies the SNAT rule for a local Pod. If the SNAT IP exists on the local Node,
// it sets the packet mark with the ID of the SNAT IP, for the traffic from local Pods to external; if the SNAT IP is
// on a remote Node, it tunnels the packets to the remote Node.
func (f *featureEgress) snatRuleFlow(ofPort uint32, snatIP net.IP, snatMark uint32, localGatewayMAC net.HardwareAddr) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// Local SNAT IP.

// To apply rate-limit on all traffic.

// SNAT IP should be on a remote Node.

// Set tunnel destination to the SNAT IP.

func (f *featureEgress) egressQoSFlow(mark uint32) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

func (f *featureEgress) egressQoSDefaultFlow() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// nodePortMarkFlows generates the flows to mark the first packet of Service NodePort connection with ToNodePortAddressRegMark,
// which indicates the Service type is NodePort.
func (f *featureService) nodePortMarkFlows() []binding.Flow { _ = "STUB: not implemented"; return nil }

// This generates a flow for every NodePort IP. The flows are used to mark the first packet of NodePort connection
// from a local Pod.

// From the perspective of a local Pod, the traffic destined to loopback is not NodePort traffic, so we skip
// the loopback address.

// This generates the flow for the virtual NodePort DNAT IP. The flow is used to mark the first packet of NodePort
// connection sourced from the Antrea gateway (the connection is performed DNAT with the virtual IP in host netns).

// serviceLearnFlow generates the flow with learn action which adds new flows in SessionAffinityTable according to the
// Endpoint selection decision.
func (f *featureService) serviceLearnFlow(config *types.ServiceConfig) binding.Flow {
	_ = "STUB: not implemented"
	// Using unique cookie ID here to avoid learned flow cascade deletion.
	return *new(binding.Flow)
}

// EpToLearnRegMark is required to match the packets that have done Endpoint selection.

// ToNodePortAddressRegMark is required to match the packets if the flow is for NodePort address, otherwise Service IP
// is used.

// affinityTimeout is used as the OpenFlow "hard timeout": learned flow will be removed from
// OVS after that time regarding of whether traffic is still hitting the flow. This is the
// desired behavior based on the K8s spec. Note that existing connections will keep going to
// the same endpoint because of connection tracking; and that is also the desired behavior.

// Loading the EpSelectedRegMark indicates that the Endpoint selection is completed. RewriteMACRegMark must be loaded
// for Service packets.

// If the flow is for external Service IP, which means the Service is accessible externally, the ToExternalAddressRegMark
// should be loaded to determine whether SNAT is required for the connection.

// serviceLBFlows generates the flows which use the specific groups to do Endpoint selection.
func (f *featureService) serviceLBFlows(config *types.ServiceConfig) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// EpToSelectRegMark is required to match the packets that haven't undergone Endpoint selection yet.
// ToNodePortAddressRegMark is required to match the packets if the flow is for NodePort address, otherwise Service IP
// is used.

// RewriteMACRegMark must be loaded for Service packets.

// If the flow is for external Service IP, which means the Service is accessible externally, the ToExternalAddressRegMark
// should be loaded to determine whether SNAT is required for the connection.

// For short-circuiting flow, an extra match condition matching packet from a local Pod or the Node is added.

// For DSR Service, we add a flow to match packets received from tunnel device, which means it has been
// load-balanced once in ingress Node, and we must select a local Endpoint on this Node.

// dsrServiceMarkFlow generates the flow which matches the packets with the following attributes:
//  1. It's accessing the DSR Service's IP and port.
//  2. It's externally originated.
//  3. It's going to be sent to a remote Endpoint.
//  4. It doesn't already have the DSRServiceRegMark.
//
// And it will perform the following actions:
//  1. Load DSRServiceRegMark to indicate this packet uses DSR mode.
//  2. Generate a learned flow which matches the 5-tuple of the connection, to ensure the same Endpoint will be selected
//     for subsequent packets of the connection.
func (f *featureService) dsrServiceMarkFlow(config *types.ServiceConfig) binding.Flow {
	_ = "STUB: not implemented"
	// Using unique cookie ID here to avoid learned flow cascade deletion.
	return *new(binding.Flow)
}

// This learned flow has higher priority than the learned flow generated for ClientIP session affinity because
// we need this connection's traffic to hit this flow to reset the idle duration and its FIN/RST packet to reset
// the idle timeout.

// endpointRedirectFlowForServiceIP generates the flow which uses the specific group for a Service's ClusterIP
// to do final Endpoint selection.
func (f *featureService) endpointRedirectFlowForServiceIP(config *types.ServiceConfig) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// endpointDNATFlow generates the flow which transforms the Service Cluster IP to the Endpoint IP according to the Endpoint
// selection decision which is stored in regs.
func (f *featureService) endpointDNATFlow(endpointIP net.IP, endpointPort uint16, protocol binding.Protocol) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// dsrServiceNoDNATFlows generates the flows which prevent traffic in DSR mode from being DNATed on the ingress Node.
func (f *featureService) dsrServiceNoDNATFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// serviceEndpointGroup creates/modifies the group/buckets of Endpoints. If the withSessionAffinity is true, then buckets
// will resubmit packets back to ServiceLBTable to trigger the learn flow, the learn flow will then send packets to
// EndpointDNATTable. Otherwise, buckets will resubmit packets to EndpointDNATTable directly.
// IMPORTANT: Ensure any changes to this function are tested in TestServiceEndpointGroupMaxBuckets.
func (f *featureService) serviceEndpointGroup(groupID binding.GroupIDType, withSessionAffinity bool, endpoints ...proxy.Endpoint) binding.Group {
	_ = "STUB: not implemented"
	return *new(binding.Group)
}

// It will be EndpointDNATTable if DSR is not enabled, otherwise DSRServiceMarkTable.

// Load RemoteEndpointRegMark for remote non-hostNetwork Endpoints.

// decTTLFlows generates the flow to process TTL. For the packets forwarded across Nodes, TTL should be decremented by one;
// for packets which enter OVS pipeline from the Antrea gateway, as the host IP stack should have decremented the TTL
// already for such packets, TTL should not be decremented again.
func (f *featurePodConnectivity) decTTLFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// Skip packets from the gateway interface.

// externalFlows generates the flows to perform SNAT for the packets of connection to the external network. The flows identify
// the packets to external network, and send them to EgressMarkTable, where SNAT IPs are looked up for the packets.
func (f *featureEgress) externalFlows() []binding.Flow { _ = "STUB: not implemented"; return nil }

// This generates the flow to match the packets sourced from local Pods and destined for external network, then
// forward them to EgressMarkTable.

// This generates the flow to match the packets sourced from tunnel and destined for external network, then
// forward them to EgressMarkTable.

// This generates the default flow to drop the packets from remote Nodes and there is no matched SNAT policy.

// This generates the flow to bypass the packets destined for local Node.

// This generates the flows to bypass the packets sourced from local Pods and destined for the except CIDRs for Egress.

// This generates the flow to match the packets of tracked Egress connection and forward them to stageSwitching.

// policyConjKeyFunc knows how to get key of a *policyRuleConjunction.
func policyConjKeyFunc(obj interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

// priorityIndexFunc knows how to get priority of actionFlows in a *policyRuleConjunction.
// It's provided to cache.Indexer to build an index of policyRuleConjunction.
func priorityIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// genOFMeter generates a meter entry with specific meterID, meterFlag, rate and
// burst. Packets which exceed the rate will be dropped.
func (c *client) genOFMeter(meterID binding.MeterIDType, meterFlags ofctrl.MeterFlag, rate uint32, burst uint32) binding.Meter {
	_ = "STUB: not implemented"
	return *new(binding.Meter)
}

func generatePipeline(pipelineID binding.PipelineID, requiredTables []*Table) binding.Pipeline {
	_ = "STUB: not implemented"
	return *new(binding.Pipeline)
}

// Generate a sequencing ID for the flow table.

// Initialize a flow table.

// realizePipelines sets next ID and missing action for every flow table in every pipeline and realize it on OVS bridge.
func (c *client) realizePipelines() { _ = "STUB: not implemented"; return }

// For the last table in a pipeline, set the miss action to TableMissActionDrop and next ID to LastTableID.

// For a table (not the last one) in a pipeline, set the next ID to the next table ID. If the miss action
// of the table is TableMissActionNone, set the miss action to TableMissActionNext.

func pipelineClassifyFlow(cookieID uint64, protocol binding.Protocol, pipeline binding.Pipeline) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// igmpEgressFlow generates flows to match IGMP report to jump to table MulticastRoutingTable.
// This is because normal multicast egress rule can match IGMP v1 report, when there is egress
// rule to block multicast traffic, IGMP v1 report will also be blocked, which is not expected.
func (f *featureMulticast) igmpEgressFlow() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// igmpPktInFlows generates the flow to load CustomReasonIGMPRegMark to mark the IGMP packet in MulticastRoutingTable
// and sends it to antrea-agent.
func (f *featureMulticast) igmpPktInFlows() []binding.Flow { _ = "STUB: not implemented"; return nil }

// Set a custom category for the IGMP packets, and then send it to antrea-agent. Then antrea-agent can identify
// the local multicast group and its members in the meanwhile.
// Do not set dst IP address because IGMPv1 report message uses target multicast group as IP destination in
// the packet.

// localMulticastForwardFlows generates the flow to forward multicast packets with OVS action "normal", and outputs
// it to Antrea gateway in the meanwhile, so that the packet can be forwarded to local Pods which have joined the Multicast
// group and to the external receivers. For external multicast packets accessing to the given multicast IP also hits the
// flow, and the packet is not sent back to Antrea gateway because OVS datapath will drop it when it finds the output
// port is the same as the input port.
func (f *featureMulticast) localMulticastForwardFlows(multicastIP net.IP, groupID binding.GroupIDType) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// externalMulticastReceiverFlow generates the flow to output multicast packets to Antrea gateway interface (to the host interface
// and the uplink interface when flexibleIPAM is enabled), so that local Pods can send multicast packets to the external receivers.
// For the case that one or more local Pods have joined the target multicast group, it is handled by the flows created by
// function "localMulticastForwardFlows" after local Pods report the IGMP membership.
// Because there are ingress tables between MulticastRoutingTable and MulticastOutputTable, while currently ingress rules only
// support IGMP query, it is not necessary to goto the ingress tables for other multicast traffic.
func (f *featureMulticast) externalMulticastReceiverFlow() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// NewClient is the constructor of the Client interface.
func NewClient(bridgeName string,
	mgmtAddr string,
	nodeIPCheck nodeip.Checker,
	enableProxy bool,
	enableAntreaPolicy bool,
	enableL7NetworkPolicy bool,
	enableEgress bool,
	enableEgressTrafficShaping bool,
	enableDenyTracking bool,
	proxyAll bool,
	enableDSR bool,
	connectUplinkToBridge bool,
	enableMulticast bool,
	enableTrafficControl bool,
	enableMulticluster bool,
	groupIDAllocator GroupAllocator,
	enablePrometheusMetrics bool,
	packetInRate int,
) *client {
	_ = "STUB: not implemented"
	return nil
}

// Pre-initialize the map with all possible keys to avoid concurrent updates and potential race conditions later.

type conjunctiveActionsInOrder []*conjunctiveAction

func (sl conjunctiveActionsInOrder) Len() int           { _ = "STUB: not implemented"; return 0 }
func (sl conjunctiveActionsInOrder) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (sl conjunctiveActionsInOrder) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// l3FwdFlowToLocalPodCIDR generates the flow to match the packets to local per-Node IPAM Pods.
func (f *featurePodConnectivity) l3FwdFlowToLocalPodCIDR() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to match the packets destined for local Pods without RewriteMACRegMark.

// l3FwdFlowToNode generates the flows to match the packets destined for local Node.
func (f *featurePodConnectivity) l3FwdFlowToNode() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to match the packets sourced from local Antrea Pods and destined for local Node
// via bridge local port.

// When Node bridge local port and uplink port connect to OVS, this generates the flow to match the reply
// packets of connection initiated through the bridge local port with FromBridgeCTMark.

// l3FwdFlowToExternal generates the flow to forward packets destined for external network. Corresponding cases are listed
// in the follows:
//   - when Egress is disabled, request packets of connections sourced from local Pods and destined for external network.
//   - when AntreaIPAM is enabled, request packets of connections sourced from local AntreaIPAM Pods and destined for external network.
//
// TODO: load ToUplinkRegMark to packets sourced from AntreaIPAM Pods and destined for external network.
//
// Due to the lack of defined variables of flow priority, there are not enough flow priority to install the flows to
// differentiate the packets sourced from AntreaIPAM Pods and non-AntreaIPAM Pods. For the packets sourced from AntreaIPAM
// Pods and destined for external network, they are forwarded via uplink port, not Antrea gateway. Apparently, loading
// ToGatewayRegMark to such packets is not right. However, packets sourced from AntreaIPAM Pods with ToGatewayRegMark
// don't cause unexpected effects to the consumers of ToGatewayRegMark and ToUplinkRegMark. Consumers of these two
// marks are listed in the follows:
//   - In IngressSecurityClassifierTable, flows are installed to forward the packets with ToGatewayRegMark, ToGatewayRegMark
//     or ToUplinkRegMark to IngressMetricTable directly.
//   - In ServiceMarkTable, ToGatewayRegMark is used with FromGatewayRegMark together.
//   - In ServiceMarkTable, ToUplinkRegMark is only used in noEncap mode + Windows.
func (f *featurePodConnectivity) l3FwdFlowToExternal() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// hostBridgeLocalFlows generates the flows to match the packets forwarded between bridge local port and uplink port.
func (f *featurePodConnectivity) hostBridgeLocalFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to forward the packets from uplink port to bridge local port.

// This generates the flow to forward the packets from bridge local port to uplink port.

// hostBridgeUplinkVLANFlows generates the flows to match VLAN packets from uplink port.
func (f *featurePodConnectivity) hostBridgeUplinkVLANFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// podVLANFlows generates the flows to match the packets from Pod and set VLAN ID.
func (f *featurePodConnectivity) podVLANFlow(podOFPort uint32, vlanID uint16) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// preRoutingClassifierFlows generates the flow to classify packets in stagePreRouting.
func (f *featureService) preRoutingClassifierFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the default flow to match the first packet of a connection.

// l3FwdFlowToExternalEndpoint generates the flow to forward the packets of Service connections sourced from local Antrea
// gateway and destined for external network. Note that, the destination MAC address of the packets should be rewritten to
// local Antrea gateway's so that the packets can be forwarded to external network via local Antrea gateway.
func (f *featureService) l3FwdFlowToExternalEndpoint() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// podHairpinSNATFlow generates the flow to match the first packet of hairpin connection initiated through a local Pod.
// ConnSNATCTMark and HairpinCTMark will be loaded in DNAT CT zone.
func (f *featureService) podHairpinSNATFlow(endpoint net.IP) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// gatewaySNATFlows generate the flows to match the first packet of Service connection initiated through the Antrea gateway,
// and the connection requires SNAT.
func (f *featureService) gatewaySNATFlows() []binding.Flow { _ = "STUB: not implemented"; return nil }

// This generates the flow to match the first packet of hairpin connection initiated through the Antrea gateway.
// ConnSNATCTMark and HairpinCTMark will be loaded in DNAT CT zone.

// This generates the flow to match the first packets of externally-originated connections towards external
// addresses of the Service initiated through the Antrea gateway, and the selected Endpoint is on a remote Node,
// then ConnSNATCTMark will be loaded in DNAT CT zone, indicating that SNAT is required for the connection.

// Do not SNAT DSR traffic.

func getCachedFlowMessages(cache *flowCategoryCache) []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

func getZoneSrcField(connectUplinkToBridge bool) *binding.RegField {
	_ = "STUB: not implemented"
	return nil
}
