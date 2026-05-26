// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connections

import (
	"net"
	"net/netip"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
)

// Some connections (e.g., Service connections) are committed to conntrack before policy
// enforcement, but are later denied by a policy. We want to ignore these connections to avoid
// duplicates between the conntrack connection store and the deny connection store. These
// connections can be reliably identified by checking if the ConnSourceCTMarkField in the CT mark is
// unset.
var connAllowedCTMarkMask = binding.NewCTMark(openflow.ConnSourceCTMarkField, 0xf).GetValue()

// InitializeConnTrackDumper initializes the ConnTrackDumper interface for different OS and datapath types.
func InitializeConnTrackDumper(nodeConfig *config.NodeConfig, serviceCIDRv4 *net.IPNet, serviceCIDRv6 *net.IPNet, ovsDatapathType ovsconfig.OVSDatapathType, isAntreaProxyEnabled bool) ConnTrackDumper {
	_ = "STUB: not implemented"
	return *new(ConnTrackDumper)
}

func filterAntreaConns(conns []*connection.Connection, nodeConfig *config.NodeConfig, serviceCIDR netip.Prefix, zoneFilter uint16, isAntreaProxyEnabled bool) []*connection.Connection {
	_ = "STUB: not implemented"
	return nil
}

// Consider Pod-to-Pod, Pod-To-Service and Pod-To-External flows.

// Pod-to-Service flows with kube-proxy: There are two conntrack flows
// for every Pod-to-Service flow. One is with ClusterIP as destination
// and the other one is with resolved endpoint PodIP as destination.
// Both conntrack flows have same stats, which makes them duplicates.
// We ignore the connection with ClusterIP and keep the connection with
// the endpoint PodIP, which is essentially Pod-to-Pod flow.
// TODO: Consider the conntrack flows from default zoneID to get iptables
// related flow that has both ClusterIP and resolved endpoint PodIP.
