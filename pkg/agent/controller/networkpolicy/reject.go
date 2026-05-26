// Copyright 2021 Antrea Authors
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

package networkpolicy

import (
	"antrea.io/ofnet/ofctrl"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

type rejectType int

const (
	// rejectPodLocal represents this packetOut is used to reject Pod-to-Pod traffic
	// and for this response, the srcPod and the dstPod are on the same Node.
	rejectPodLocal rejectType = iota
	// rejectPodRemoteToLocal represents this packetOut is used to reject Pod-to-Pod
	// traffic and for this response, the srcPod is on a remote Node and the dstPod is
	// on the local Node.
	rejectPodRemoteToLocal
	// rejectPodLocalToRemote represents this packetOut is used to reject Pod-to-Pod
	// traffic and for this response, the srcPod is on the local Node and the dstPod is
	// on a remote Node.
	rejectPodLocalToRemote
	// rejectServiceLocal represents this packetOut is used to reject Service traffic,
	// when AntreaProxy is enabled. The EndpointPod and the dstPod of the reject
	// response are on the same Node.
	rejectServiceLocal
	// rejectServiceRemoteToLocal represents this packetOut is used to reject Service
	// traffic, when AntreaProxy is enabled. The EndpointPod is on a remote Node and
	// the dstPod of the reject response is on the local Node.
	rejectServiceRemoteToLocal
	// rejectServiceLocalToRemote represents this packetOut is used to reject Service
	// traffic, when AntreaProxy is enabled. The EndpointPod is on the local Node and
	// the dstPod of the reject response is on a remote Node.
	rejectServiceLocalToRemote
	// rejectNoAPServiceLocal represents this packetOut is used to reject Service
	// traffic, when AntreaProxy is disabled. The EndpointPod and the dstPod of the
	// reject response are on the same Node.
	rejectNoAPServiceLocal
	// rejectNoAPServiceRemoteToLocal represents this packetOut is used to reject
	// Service traffic, when AntreaProxy is disabled. The EndpointPod is on a remote
	// Node and the dstPod of the reject response is on the local Node.
	rejectNoAPServiceRemoteToLocal
	// rejectServiceRemoteFromTunToExternal represents this packetOut is used to reject
	// Service traffic, when AntreaProxy is enabled. The EndpointPod is on a remote Node
	// that is reachable from OVS via the tunnel interface, and the destination of the reject
	// response is an external client.
	rejectServiceRemoteFromTunToExternal
	// rejectServiceRemoteFromGwToExternal represents this packetOut is used to reject
	// Service traffic, when AntreaProxy is enabled. The EndpointPod is on a remote Node
	// that is reachable from OVS via the antrea gateway interface, and the destination of
	// the reject response is an external client.
	rejectServiceRemoteFromGwToExternal
	// unsupported indicates that Antrea couldn't generate packetOut for current
	// packetIn.
	unsupported
)

// rejectRequest sends reject response to the requesting client, based on the
// packet-in message.
func (c *Controller) rejectRequest(pktIn *ofctrl.PacketIn) error {
	_ = "STUB: not implemented"
	// All src/dst mean the source/destination of the reject packet, which are destination/source of the incoming packet.
	// Get ethernet data.
	return nil
}

// Get IP data.

// Get IP data.

// dstIsDirect means that the reject packet destination is on the same Node and the reject packet can be forwarded
// without leaving the OVS bridge.

// Check if OVS InPort matches dIface.
// If port doesn't match, set dstIsDirect to false since the reject packet destination should not be sent to
// local Pod directly.

// isServiceTraffic checks if it's a Service traffic when the destination of the
// reject response is on local Node. When the destination of the reject response is
// remote, isServiceTraffic will always return false. Because there is no
// difference between Service traffic and Pod-to-Pod traffic in this case. They all
// belong to RejectLocalToRemote type and use the same logic to handle.
// There are two situations in which it can be determined that this is a service
// traffic:
// 1. When AntreaProxy is enabled, EpSelectedRegMark is set in ServiceEPStateField.
//    AntreaProxy is required for FlexibleIPAM feature.
// 2. When AntreaProxy is disabled, dstIP of reject response is on the local Node
//    and dstMAC of reject response is antrea-gw's MAC. In this case, the reject
//    response is being generated for locally-originated traffic that went through
//    kube-proxy and was re-injected into the bridge through antrea-gw.

// When rejecting external client access to a Service with a remote Endpoint, there are two scenarios:
// 1. If the remote Endpoint is reachable via the OVS tunnel interface (i.e., traffic mode is "encap" or "hybrid"
//    with the remote Node on a different subnet), the packet-out should simulate being sent from the tunnel.
// 2. If the remote Endpoint is reachable via the OVS gateway interface and Node route (i.e., traffic mode is
//    "noEncap" or "hybrid" with the remote Node on the same subnet), the packet-out should simulate being sent
//    from the gateway.
// The PktDestinationField set in L3Forwarding table indicates the intended outPort of the packet-in packet, which
// can help determine the appropriate inPort for the packet-out packet.

// When in AntreaIPAM mode, even though srcPod and dstPod are on the same Node, MAC
// will still be re-written in L3ForwardingTable. During rejection, the reject
// response will be directly sent to the dst OF port without go through
// L3ForwardingTable. So we need to re-write MAC here. There is no need to check
// whether AntreaIPAM mode is enabled. Because if AntreaIPAM mode is disabled,
// this re-write doesn't change anything.

// getRejectType returns rejectType of a rejection.
func getRejectType(isServiceTraffic, antreaProxyEnabled, srcIsLocal, dstIsLocal, srcFromTun bool) rejectType {
	_ = "STUB: not implemented"
	return *new(rejectType)
}

// getRejectOFPorts returns the inPort and outPort of a packetOut based on the rejectType.
func getRejectOFPorts(rejectType rejectType,
	sIface *interfacestore.InterfaceConfig,
	dIface *interfacestore.InterfaceConfig,
	gwOFPort uint32,
	tunOFPort uint32) (uint32, uint32) {
	_ = "STUB: not implemented"
	return 0, 0
}

// For rejectServiceLocal and rejectServiceLocalToRemote, we set inPort as the
// OFPort of the srcPod to simulate its rejection. And we don't set outPort, since
// it's Service traffic load-balanced by AntreaProxy. The reject response packet
// needs to be UnDNATed by the pipeline, instead of directly sending it out
// through outPort.

// If tunnel interface is not found, which means we are in noEncap mode, then use
// gateway port as inPort.

// getRejectPacketOutMutateFunc returns the mutate func of a packetOut based on the rejectType.
func getRejectPacketOutMutateFunc(rejectType rejectType, nodeType config.NodeType, isFlexibleIPAMSrc, isFlexibleIPAMDst bool, ctZone uint32) func(binding.PacketOutBuilder) binding.PacketOutBuilder {
	_ = "STUB: not implemented"
	return nil
}

// L3ForwardingTable is not initialized for ExternalNode case since layer 3 is not needed.

func parseFlexibleIPAMStatus(pktIn *ofctrl.PacketIn, nodeConfig *config.NodeConfig, srcIP string, srcIsLocal bool, dstIP string, dstIsLocal bool) (isFlexibleIPAMSrc bool, isFlexibleIPAMDst bool, ctZone uint32, err error) {
	_ = "STUB: not implemented"
	// isFlexibleIPAMSrc is true if srcIP belongs to a local FlexibleIPAM Pod.
	// isFlexibleIPAMDst is true if dstIP belongs to a local FlexibleIPAM Pod.
	// ctZone is not zero if FlexibleIPAM is enabled.
	return false, false, 0, nil
}

// ctZone is read from the incoming packet.
// The generated reject packet should have same ctZone with the incoming packet, otherwise the conntrack cannot work properly.
