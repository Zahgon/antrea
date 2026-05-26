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
	"sync"

	"antrea.io/libOpenflow/openflow15"

	"antrea.io/antrea/v2/pkg/agent/openflow/cookie"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

type featureMulticast struct {
	cookieAllocator     cookie.Allocator
	ipProtocols         []binding.Protocol
	bridge              binding.Bridge
	gatewayPort         uint32
	encapEnabled        bool
	flexibleIPAMEnabled bool
	tunnelPort          uint32
	uplinkPort          uint32
	hostOFPort          uint32

	cachedFlows        *flowCategoryCache
	groupCache         sync.Map
	enableAntreaPolicy bool

	category cookie.Category
}

func (f *featureMulticast) getFeatureName() string { _ = "STUB: not implemented"; return "" }

func newFeatureMulticast(
	cookieAllocator cookie.Allocator,
	ipProtocols []binding.Protocol,
	bridge binding.Bridge,
	anpEnabled bool,
	gwPort uint32,
	encapEnabled bool,
	tunnelPort uint32,
	uplinkPort uint32,
	hostOFPort uint32,
	flexibleIPAMEnabled bool,
) *featureMulticast {
	_ = "STUB: not implemented"
	return nil
}

func multicastPipelineClassifyFlow(cookieID uint64, pipeline binding.Pipeline) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

func (f *featureMulticast) initFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	// Install flows to send the IGMP report messages to Antrea Agent.
	return nil
}

// Install flow to forward the IGMP query messages to all local Pods.

// Install flows to forward the multicast traffic to antrea-gw0 if no local Pods have joined in the group, and this
// is to ensure local Pods can access the external multicast receivers.

// Install flows to output multicast packets.

func (f *featureMulticast) replayFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	// Get cached flows.
	return nil
}

// IMPORTANT: Ensure any changes to this function are tested in TestMulticastReceiversGroupMaxBuckets.
func (f *featureMulticast) multicastReceiversGroup(groupID binding.GroupIDType, tableID uint8, ports []uint32, remoteIPs []net.IP) binding.Group {
	_ = "STUB: not implemented"
	return *new(binding.Group)
}

func (f *featureMulticast) multicastOutputFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// When running with encap mode, drop the multicast packets if it is received from tunnel port and expected to
// output to antrea-gw0, or received from antrea-gw0 and expected to output to tunnel. These flows are used to
// avoid duplication on packet forwarding. For example, if the packet is received on tunnel port, it means
// the sender is a Pod on other Node, then the packet is already sent to external via antrea-gw0 on the source
// Node. On the reverse, if the packet is received on antrea-gw0, it means the sender is from external, then
// the Pod receivers on other Nodes should also receive the packets from the underlay network.

func (f *featureMulticast) multicastSkipIGMPMetricFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureMulticast) multicastPodMetricFlows(podIP net.IP, podOFPort uint32) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// Generates the flows to forward multicast egress packets before outputting to MulticastOutputTable.
// It matches source IP with IP of the local sender Pod.

// Generates the flows to collect multicast ingress packets metrics before outputting to MulticastOutputTable.
// It matches TargetOFPortField with the OFPort of a multicast receiver Pod.

func (f *featureMulticast) replayGroups() []binding.OFEntry { _ = "STUB: not implemented"; return nil }

func (f *featureMulticast) initGroups() []binding.OFEntry { _ = "STUB: not implemented"; return nil }

func (f *featureMulticast) replayMeters() []binding.OFEntry { _ = "STUB: not implemented"; return nil }

func (f *featureMulticast) multicastForwardFlexibleIPAMFlows(table binding.Table) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureMulticast) multicastRemoteReportFlows(groupID binding.GroupIDType, firstMulticastTable binding.Table) []binding.Flow {
	_ = "STUB: not implemented"
	return nil

	// This flow outputs the IGMP report message sent from Antrea Agent to an OpenFlow group which is expected to
	// broadcast to all the other Nodes in the cluster. The multicast groups in side the IGMP report message
	// include the ones local Pods have joined in.
}

// This flow ensures the IGMP report message sent from Antrea Agent to bypass the check in SpoofGuardTable.

// This flow ensures the multicast packet sent from a different Node via the tunnel port to enter Multicast
// pipeline.
