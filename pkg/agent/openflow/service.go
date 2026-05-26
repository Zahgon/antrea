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
	"antrea.io/antrea/v2/pkg/agent/nodeip"
	"antrea.io/antrea/v2/pkg/agent/openflow/cookie"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

type featureService struct {
	cookieAllocator cookie.Allocator
	nodeIPChecker   nodeip.Checker
	ipProtocols     []binding.Protocol
	bridge          binding.Bridge

	cachedFlows *flowCategoryCache
	groupCache  sync.Map

	gatewayIPs             map[binding.Protocol]net.IP
	virtualIPs             map[binding.Protocol]net.IP
	virtualNodePortDNATIPs map[binding.Protocol]net.IP
	dnatCtZones            map[binding.Protocol]int
	snatCtZones            map[binding.Protocol]int
	gatewayMAC             net.HardwareAddr
	nodePortAddresses      map[binding.Protocol][]net.IP
	serviceCIDRs           map[binding.Protocol]net.IPNet
	networkConfig          *config.NetworkConfig
	gatewayPort            uint32

	enableAntreaPolicy    bool
	enableProxy           bool
	proxyAll              bool
	enableDSR             bool
	connectUplinkToBridge bool
	ctZoneSrcField        *binding.RegField

	category cookie.Category
}

func (f *featureService) getFeatureName() string { _ = "STUB: not implemented"; return "" }

func newFeatureService(
	cookieAllocator cookie.Allocator,
	nodeIPChecker nodeip.Checker,
	ipProtocols []binding.Protocol,
	nodeConfig *config.NodeConfig,
	networkConfig *config.NetworkConfig,
	serviceConfig *config.ServiceConfig,
	bridge binding.Bridge,
	enableAntreaPolicy,
	enableProxy,
	proxyAll,
	enableDSR,
	connectUplinkToBridge bool) *featureService {
	_ = "STUB: not implemented"
	return nil
}

// serviceNoEndpointFlow generates the flow to match the packets to Service without Endpoint and send them to controller.
func (f *featureService) serviceNoEndpointFlow() binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

func (f *featureService) initFlows() []*openflow15.FlowMod { _ = "STUB: not implemented"; return nil }

// This installs the flows to match the first packet of NodePort connection. The flows set a bit of a register
// to mark the Service type of the packet as NodePort, and the mark is consumed in table serviceLBTable.

// This installs the flows to enable Service connectivity. Upstream kube-proxy is leveraged to provide load-balancing,
// and the flows installed by this method ensure that traffic sent from local Pods to any Service address can be
// forwarded to the host gateway interface correctly, otherwise packets might be dropped by egress rules before
// they are DNATed to backend Pods.

func (f *featureService) replayFlows() []*openflow15.FlowMod { _ = "STUB: not implemented"; return nil }

func (f *featureService) replayGroups() []binding.OFEntry { _ = "STUB: not implemented"; return nil }

func (f *featureService) initGroups() []binding.OFEntry { _ = "STUB: not implemented"; return nil }

func (f *featureService) replayMeters() []binding.OFEntry { _ = "STUB: not implemented"; return nil }
