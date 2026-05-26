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

	"antrea.io/antrea/v2/pkg/agent/openflow/cookie"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

// GlobalVirtualMACForMulticluster is a vritual MAC which will be used only
// for cross-cluster traffic to distinguish from in-cluster traffic.
var GlobalVirtualMACForMulticluster, _ = net.ParseMAC("aa:bb:cc:dd:ee:f0")

// UnknownLabelIdentity represents an unknown label identity.
// 24 bits in VNI are used for label identity. The max value is reserved for
// UnknownLabelIdentity.
const UnknownLabelIdentity = uint32(0xffffff)

type featureMulticluster struct {
	cookieAllocator cookie.Allocator
	cachedFlows     *flowCategoryCache
	cachedPodFlows  *flowCategoryCache
	category        cookie.Category
	ipProtocols     []binding.Protocol
	dnatCtZones     map[binding.Protocol]int
	snatCtZones     map[binding.Protocol]int
}

func (f *featureMulticluster) getFeatureName() string { _ = "STUB: not implemented"; return "" }

func newFeatureMulticluster(cookieAllocator cookie.Allocator, ipProtocols []binding.Protocol) *featureMulticluster {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureMulticluster) initFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureMulticluster) replayFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureMulticluster) initGroups() []binding.OFEntry { _ = "STUB: not implemented"; return nil }

func (f *featureMulticluster) replayGroups() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureMulticluster) replayMeters() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureMulticluster) l3FwdFlowToRemoteGateway(
	localGatewayMAC net.HardwareAddr,
	peerServiceCIDR net.IPNet,
	tunnelPeer net.IP,
	remoteGatewayIP net.IP,
	enableStretchedNetworkPolicy bool) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to forward cross-cluster request packets based
// on Service ClusterIP range.

// Rewrite src MAC to local gateway MAC.
// Rewrite dst MAC to virtual MC MAC.
// Flow based tunnel. Set tunnel destination.

// This generates the flow to forward cross-cluster reply traffic based
// on Gateway IP.

// Flow based tunnel. Set tunnel destination.

// This generates the flow to forward cross-cluster reject traffic based
// on Gateway IP and reg.

// Flow based tunnel. Set tunnel destination.

func (f *featureMulticluster) tunnelClassifierFlow(tunnelOFPort uint32) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

func (f *featureMulticluster) outputHairpinTunnelFlow(tunnelOFPort uint32) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// snatConntrackFlows generates flows on a multi-cluster Gateway Node to perform SNAT for cross-cluster connections.
func (f *featureMulticluster) snatConntrackFlows(serviceCIDR net.IPNet, localGatewayIP net.IP) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to match the first packet of multi-cluster Service connection, and commit them into
// DNAT zone to make sure DNAT is performed before SNAT for any remote cluster traffic.

// This generates the flow to perform SNAT for the cross-cluster Service connections.

// This generates the flow to unSNAT reply packets of connections committed in SNAT CT zone by the above flows.

func (f *featureMulticluster) l3FwdFlowToPodViaTun(
	localGatewayMAC net.HardwareAddr,
	podIP net.IP,
	tunnelPeer net.IP) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

// This generates the flow to forward cross-cluster request packets based
// on Pod IP.

// Rewrite src MAC to local gateway MAC.
// Flow based tunnel. Set tunnel destination.
