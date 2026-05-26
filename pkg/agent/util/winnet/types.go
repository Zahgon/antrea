// Copyright 2024 Antrea Authors
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

package winnet

import (
	"net"

	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

type Route struct {
	LinkIndex         int
	DestinationSubnet *net.IPNet
	GatewayAddress    net.IP
	RouteMetric       int
}

type Neighbor struct {
	LinkIndex        int
	IPAddress        net.IP
	LinkLayerAddress net.HardwareAddr
	State            string
}

type NetNatStaticMapping struct {
	Name         string
	ExternalIP   net.IP
	ExternalPort uint16
	InternalIP   net.IP
	InternalPort uint16
	Protocol     binding.Protocol
}

func (r *Route) String() string { _ = "STUB: not implemented"; return "" }

func (r *Route) Equal(x Route) bool { _ = "STUB: not implemented"; return false }

func (n *Neighbor) String() string { _ = "STUB: not implemented"; return "" }

func (n *NetNatStaticMapping) String() string { _ = "STUB: not implemented"; return "" }
