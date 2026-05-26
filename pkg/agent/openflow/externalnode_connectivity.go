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

const (
	policyBypassFlowsKey = "policyBypassFlows"
)

type featureExternalNodeConnectivity struct {
	cookieAllocator cookie.Allocator
	ipProtocols     []binding.Protocol
	ctZones         map[binding.Protocol]int
	category        cookie.Category

	uplinkFlowCache *flowCategoryCache
}

func (f *featureExternalNodeConnectivity) getFeatureName() string {
	_ = "STUB: not implemented"
	return ""
}

func newFeatureExternalNodeConnectivity(
	cookieAllocator cookie.Allocator,
	ipProtocols []binding.Protocol) *featureExternalNodeConnectivity {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureExternalNodeConnectivity) vmUplinkFlows(hostOFPort, uplinkOFPort uint32) []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// Set the output port number with the uplink port if the IP packet enters OVS from the
// paired host internal port, and then enforce the packet to go through the IP pipeline.

// Set the output port number with the paired host internal port if the IP packet enters the OVS from
// the uplink port, and then enforce the packet to go through the IP pipeline.

// Output the packet to the uplink port if it is not using IP protocol, and enters OVS from the
// paired host internal port.

// Output the packet to the uplink port if it is not using IP protocol, and enters OVS from the
// paired host internal port.

func (f *featureExternalNodeConnectivity) initFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to maintain tracked connections in CT zone.

func (f *featureExternalNodeConnectivity) replayFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureExternalNodeConnectivity) initGroups() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureExternalNodeConnectivity) replayGroups() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureExternalNodeConnectivity) replayMeters() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureExternalNodeConnectivity) policyBypassFlow(protocol binding.Protocol, ipNet *net.IPNet, port uint16, isIngress bool) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

func (f *featureExternalNodeConnectivity) addPolicyBypassFlows(flow binding.Flow) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) InstallVMUplinkFlows(hostIFName string, hostPort int32, uplinkPort int32) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) UninstallVMUplinkFlows(hostIFName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) InstallPolicyBypassFlows(protocol binding.Protocol, ipNet *net.IPNet, port uint16, isIngress bool) error {
	_ = "STUB: not implemented"
	return nil
}

// nonIPPipelineClassifyFlow generates a flow in PipelineClassifierTable to resubmit packets not using IP protocols to
// pipelineNonIP.
func nonIPPipelineClassifyFlow(cookieID uint64, pipeline binding.Pipeline) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}
