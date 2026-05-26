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

package ndp

import (
	"net"
)

const (
	// Option Length, 8-bit unsigned integer. The length of the option (including the type and length fields) in units of 8 octets.
	// The value 0 is invalid. Nodes MUST silently discard a ND packet that contains an option with length zero.
	// https://datatracker.ietf.org/doc/html/rfc4861
	ndpOptionLen = 1

	// ndpOptionType
	// 	Option Name                             Type
	//
	// Source Link-Layer Address                    1
	// Target Link-Layer Address                    2
	// Prefix Information                           3
	// Redirected Header                            4
	// MTU                                          5
	ndpOptionType = 2

	// Minimum byte length values for each type of valid Message.
	naLen = 20

	// Hop limit is always 255, refer RFC 4861.
	hopLimit = 255
)

// NeighborAdvertisement sends an unsolicited Neighbor Advertisement ICMPv6 multicast packet,
// over interface 'iface' from 'srcIP', announcing a given IPv6 address('srcIP') to all IPv6 nodes as per RFC4861.
func NeighborAdvertisement(srcIP net.IP, iface *net.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

func newNDPNeighborAdvertisementMessage(targetAddress net.IP, hwa net.HardwareAddr) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ICMPType = 136, Neighbor Advertisement

// Always zero.

// The ICMP checksum. Calculated by caller or OS.
