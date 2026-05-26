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

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/openflow/cookie"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

type featureEgress struct {
	cookieAllocator cookie.Allocator
	ipProtocols     []binding.Protocol

	cachedFlows *flowCategoryCache
	cachedMeter sync.Map

	exceptCIDRs map[binding.Protocol][]net.IPNet
	nodeIPs     map[binding.Protocol]net.IP
	gatewayMAC  net.HardwareAddr

	category                   cookie.Category
	enableEgressTrafficShaping bool
}

func (f *featureEgress) getFeatureName() string { _ = "STUB: not implemented"; return "" }

func newFeatureEgress(cookieAllocator cookie.Allocator,
	ipProtocols []binding.Protocol,
	nodeConfig *config.NodeConfig,
	egressConfig *config.EgressConfig,
	enableEgressTrafficShaping bool) *featureEgress {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureEgress) initFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	// This installs the flows to enable Pods to communicate to the external IP addresses. The flows identify the packets
	// from local Pods to the external IP address, and mark the packets to be SNAT'd with the configured SNAT IPs.
	return nil
}

func (f *featureEgress) replayFlows() []*openflow15.FlowMod { _ = "STUB: not implemented"; return nil }

func (f *featureEgress) initGroups() []binding.OFEntry { _ = "STUB: not implemented"; return nil }

func (f *featureEgress) replayGroups() []binding.OFEntry { _ = "STUB: not implemented"; return nil }

func (f *featureEgress) replayMeters() []binding.OFEntry { _ = "STUB: not implemented"; return nil }
