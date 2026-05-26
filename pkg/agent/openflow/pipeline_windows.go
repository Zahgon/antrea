//go:build windows
// +build windows

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

package openflow

import (
	"net"

	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

// matchUplinkInPortInClassifierTable matches dstIP field to prevent unintended forwarding when promiscuous mode is enabled on Windows.
func (f *featurePodConnectivity) matchUplinkInPortInClassifierTable(flowBuilder binding.FlowBuilder) binding.FlowBuilder {
	_ = "STUB: not implemented"
	return *new(binding.FlowBuilder)
}

// hostBridgeUplinkFlows generates the flows that forward traffic between the bridge local port and the uplink port to
// support the host traffic with outside.
func (f *featurePodConnectivity) hostBridgeUplinkFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to forward ARP packets from uplink port to bridge local port since uplink port is set
// to disable flood.

// This generates the flow to forward ARP packets from bridge local port to uplink port since uplink port is set
// to disable flood.

// TODO: support IPv6

// If NoEncap is enabled, the reply packets from remote Pod can be forwarded to local Pod directly.
// by explicitly resubmitting them to ConntrackState stage and marking "macRewriteMark" at same time.

func (f *featurePodConnectivity) l3FwdFlowToRemoteViaRouting(localGatewayMAC net.HardwareAddr,
	remoteGatewayMAC net.HardwareAddr,
	peerIP net.IP,
	peerPodCIDR *net.IPNet) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// It enhances Windows Noencap mode performance by bypassing host network.

// Output the reply packet to the uplink interface if the destination is another Node's IP.
// This is for the scenario that another Node directly accesses Pods on this Node. Since the request
// packet enters OVS from the uplink interface, the reply should go back in the same path. Otherwise,
// Windows host will perform stateless SNAT on the reply, and the packets are possibly dropped on peer
// Node because of the wrong source address.

// This generates the flow to match the packets destined for remote Node by matching destination MAC, then
// load the ofPort number of uplink to TargetOFPortField.
