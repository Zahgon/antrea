// Copyright 2020 Antrea Authors
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

package arping

import (
	"net"
)

const (
	// 1544 = htons(ETH_P_ARP)
	protoARP = 1544
)

// GratuitousARPOverIface sends an gratuitous arp over interface 'iface' from 'srcIP'.
// It refers to "github.com/j-keck/arping" and is simplified and made thread-safe.
func GratuitousARPOverIface(srcIP net.IP, iface *net.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

func newARPRequest(sha, spa, tha, tpa []byte) []byte { _ = "STUB: not implemented"; return nil }

// Ethernet header.
// Destination MAC address.
// Source MAC address.
// Ethernet protocol type, 0x0806 for ARP.
// ARP message.
// Hardware Type, Ethernet is 1.
// Protocol type, IPv4 is 0x0800.
// Hardware length, Ethernet address length is 6.
// Protocol length, IPv4 address length is 4.
// Operation, request is 1.
// Sender hardware address.
// Sender protocol address.
// Target hardware address.
// Target protocol address.
