// Copyright 2022 Antrea Authors
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

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/openflow/cookie"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha2"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

type featurePodConnectivity struct {
	cookieAllocator cookie.Allocator
	ipProtocols     []binding.Protocol

	nodeCachedFlows *flowCategoryCache
	podCachedFlows  *flowCategoryCache
	tcCachedFlows   *flowCategoryCache

	gatewayIPs    map[binding.Protocol]net.IP
	gatewayPort   uint32
	uplinkPort    uint32
	hostIfacePort uint32
	tunnelPort    uint32
	ctZones       map[binding.Protocol]int
	snatCtZones   map[binding.Protocol]int
	localCIDRs    map[binding.Protocol]net.IPNet
	nodeIPs       map[binding.Protocol]net.IP
	nodeConfig    *config.NodeConfig
	networkConfig *config.NetworkConfig

	connectUplinkToBridge bool
	ctZoneSrcField        *binding.RegField
	ipCtZoneTypeRegMarks  map[binding.Protocol]*binding.RegMark
	enableMulticast       bool
	proxyAll              bool
	enableDSR             bool
	enableTrafficControl  bool

	category cookie.Category
}

func (f *featurePodConnectivity) getFeatureName() string { _ = "STUB: not implemented"; return "" }

func newFeaturePodConnectivity(
	cookieAllocator cookie.Allocator,
	ipProtocols []binding.Protocol,
	nodeConfig *config.NodeConfig,
	networkConfig *config.NetworkConfig,
	connectUplinkToBridge bool,
	enableMulticast bool,
	proxyAll bool,
	enableDSR bool,
	enableTrafficControl bool,
) *featurePodConnectivity {
	_ = "STUB: not implemented"
	return nil
}

func (f *featurePodConnectivity) initFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

// This installs the flows between bridge local port and uplink port to support host networking.

// Add flow to ensure the liveliness check packet could be forwarded correctly.

// If IPv6 is enabled, this flow will never get hit. Replies any ARP request with the same global virtual MAC.

// If NetworkPolicyOnly mode is enabled, IPAM is implemented by the primary CNI, which may not use the Pod CIDR
// of the Node. Therefore, it doesn't make sense to install flows for the Pod CIDR. Individual flow for each local
// Pod IP will take care of routing the traffic to destination Pod.

func (f *featurePodConnectivity) replayFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

// Get cached flows.

// trafficControlMarkFlows generates the flows to mark the packets that need to be redirected or mirrored.
func (f *featurePodConnectivity) trafficControlMarkFlows(sourceOFPorts []uint32,
	targetOFPort uint32,
	direction v1alpha2.Direction,
	action v1alpha2.TrafficControlAction,
	priority uint16) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to mark the packets destined for a provided port.

// This generates the flow to mark the packets sourced from a provided port.

// trafficControlReturnClassifierFlow generates the flow to mark the packets from traffic control return port and forward
// the packets to stageRouting directly. Note that, for the packets which are originally to be output to a tunnel port,
// value of NXM_NX_TUN_IPV4_DST for the returned packets needs to be loaded in stageRouting.
func (f *featurePodConnectivity) trafficControlReturnClassifierFlow(returnOFPort uint32) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// trafficControlCommonFlows generates the common flows for traffic control.
func (f *featurePodConnectivity) trafficControlCommonFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to output packets to the original target port as well as mirror the packets to the target
// traffic control port.

// This generates the flow to output the packets to be redirected to the target traffic control port.

// This generates the flow to forward the returned packets (with FromTCReturnRegMark) to stageOutput directly
// after loading output port number to reg1 in L2ForwardingCalcTable.

func (f *featurePodConnectivity) initGroups() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}

func (f *featurePodConnectivity) replayGroups() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}

func (f *featurePodConnectivity) replayMeters() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}
