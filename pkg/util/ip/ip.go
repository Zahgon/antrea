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

package ip

import (
	"net"
	"net/netip"

	"antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
)

const (
	V4BitLen = 8 * net.IPv4len
	V6BitLen = 8 * net.IPv6len
)

type DualStackIPs struct {
	IPv4 net.IP
	IPv6 net.IP
}

func (ips DualStackIPs) Equal(x DualStackIPs) bool { _ = "STUB: not implemented"; return false }

// This function takes in one allow CIDR and multiple except CIDRs and gives diff CIDRs
// in allowCIDR eliminating except CIDRs. It currently supports only IPv4. except CIDR input
// can be changed.
func DiffFromCIDRs(allowCIDR *net.IPNet, exceptCIDRs []*net.IPNet) ([]*net.IPNet, error) {
	_ = "STUB: not implemented"
	// Remove the redundant CIDRs
	return nil, nil
}

// Consider masked IP from IPNet struct

// Delete the considered CIDR block and add resulting CIDR blocks

// Append the result CIDRs

// This step can be optimized by having iterator over just the index. Went with reinitialization of iterator.

// Just delete the CIDR block

// This function gives diff CIDRs between a superset CIDR (allow CIDR) and subset CIDR
// (except CIDR)
func diffFromCIDR(allowCIDR, exceptCIDR *net.IPNet) []*net.IPNet {
	_ = "STUB: not implemented"
	return nil
}

// Mask the IP to get the start IP of range

// New CIDRs should not contain the IPs in exceptCIDR. Manipulating the bits in start IP of
// exceptCIDR will give remainder IPs in allowCIDR, specifically the masked IPs for remaining
// CIDRs with prefix ranging from [allowPrefix+1, exceptPrefix].

// Flip the (ipBitLen - i)th bit from LSB in exceptCIDR to get the IP which is not in exceptCIDR

func flipSingleBit(ip *net.IP, bitIndex int) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// XOR bit operation to flip

// This function is to check for redundant CIDRs in the list that are
// covered by other CIDRs and remove them. Input array can be modified.
func MergeCIDRs(cidrBlocks []*net.IPNet) []*net.IPNet {
	_ = "STUB: not implemented"
	// Sort the list by netmask in ascending order
	return nil
}

// Check and remove if there are redundant CIDRs that are part of bigger CIDRs
// or repeated CIDRs

// Delete the CIDR block and truncate the slice

// Decrement the tracker to consider next element

// IPNetToNetIPNet converts Antrea IPNet to *net.IPNet.
// Note that K8s allows non-standard CIDRs to be specified (e.g. 10.0.1.1/16, fe80::7015:efff:fe9a:146b/64). However,
// OVS will report OFPBMC_BAD_WILDCARDS error if using them in the OpenFlow messages. The function will normalize the
// CIDR if it's non-standard.
func IPNetToNetIPNet(ipNet *v1beta2.IPNet) *net.IPNet { _ = "STUB: not implemented"; return nil }

const (
	ICMPProtocol   = 1
	IGMPProtocol   = 2
	TCPProtocol    = 6
	UDPProtocol    = 17
	ICMPv6Protocol = 58
	SCTPProtocol   = 132
)

// IPProtocolNumberToString returns the string name of the IP protocol with number protocolNum. If
// the number does not match a "known" protocol, we return the defaultValue string.
func IPProtocolNumberToString(protocolNum uint8, defaultValue string) string {
	_ = "STUB: not implemented"
	return ""
}

// MustParseCIDR turns the given string into IPNet or panics, for tests or other cases where the string must be valid.
func MustParseCIDR(cidr string) *net.IPNet { _ = "STUB: not implemented"; return nil }

func MustParseMAC(mac string) net.HardwareAddr {
	_ = "STUB: not implemented"
	return *new(net.HardwareAddr)
}

// IPNetEqual returns if the provided IPNets are the same subnet.
func IPNetEqual(ipNet1, ipNet2 *net.IPNet) bool { _ = "STUB: not implemented"; return false }

// IPNetContains returns if the first IPNet contains the second IPNet.
// For example:
//
// 10.0.0.0/24 contains 10.0.0.0/24.
// 10.0.0.0/24 contains 10.0.0.0/25.
// 10.0.0.0/24 contains 10.0.0.128/25.
// 10.0.0.0/24 does not contain 10.0.0.0/23.
// 10.0.0.0/24 does not contain 10.0.1.0/25.
func IPNetContains(ipNet1, ipNet2 *net.IPNet) bool { _ = "STUB: not implemented"; return false }

func MustIPv6(s string) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// GetLocalBroadcastIP returns the last IP address in a subnet. This IP is always working as the broadcast address in
// the subnet on Windows, and an active route entry that uses it as the destination is added by default when a new IP is
// configured on the interface.
func GetLocalBroadcastIP(ipNet *net.IPNet) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// AppendPortIfMissing appends the given port to the address if the address doesn't contain any port.
func AppendPortIfMissing(addr, port string) string { _ = "STUB: not implemented"; return "" }

// Return the address directly if it's not a valid address.

// GetStartAndEndOfPrefix retrieves the start and end addresses of a netip.Prefix.
// For example:  10.10.40.0/24 -> 10.10.40.0, 10.10.40.255
func GetStartAndEndOfPrefix(prefix netip.Prefix) (netip.Addr, netip.Addr) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), *new(netip.Addr)
}

// use gateway address, of canonical form of prefix, as start address.

// calculate the end address by performing bitwise OR with the complement of the mask.
