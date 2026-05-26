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

package openflow

import (
	"net"

	"antrea.io/libOpenflow/openflow15"
	"antrea.io/ofnet/ofctrl"
)

type ofFlowBuilder struct {
	ofFlow
}

// MatchVLAN can be used as follows:
// - to match the packets of a specific VLAN, there are two cases:
//   - to match VLAN 0, nonVLAN should be false, vlanID should be 0, value of vlanMask must be 0x1fff.
//   - VLAN 1-4095, nonVLAN should be false, vlanID should be the VLAN ID, vlanMask can be nil, or its value can be 0x1fff.
//
// - to match the packets of all VLANs, nonVLAN should be false, vlanID must be 0, value of vlanMask must be 0x1000.
// - to match the packets of non-VLAN, nonVLAN should be true, vlanID must be 0, vlanMask can be nil, or its value can be 0x1000.
func (b *ofFlowBuilder) MatchVLAN(nonVLAN bool, vlanID uint16, vlanMask *uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// To match the packets of a VLAN whose VLAN ID is not 0, when vlanMask is nil, set the value of vlanMask to 0x1ffff.

// To match the packets of non-VLAN, when vlanMask is nil, set the value of vlanMask to 0x1000.

func (b *ofFlowBuilder) SetHardTimeout(timout uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) SetIdleTimeout(timeout uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) Done() Flow { _ = "STUB: not implemented"; return *new(Flow) }

// matchReg adds match condition for matching data in the target register.
func (b *ofFlowBuilder) matchReg(regID int, data uint32) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchXXReg adds match condition for matching data in the target xx-register.
func (b *ofFlowBuilder) MatchXXReg(regID int, data []byte) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// matchRegRange adds match condition for matching data in the target register at specified range.
func (b *ofFlowBuilder) matchRegRange(regID int, data uint32, rng *Range) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchRegMark(marks ...*RegMark) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchRegFieldWithValue(field *RegField, data uint32) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTState(ctStates *openflow15.CTStates) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTStateNew(set bool) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTStateRel(set bool) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTStateRpl(set bool) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTStateEst(set bool) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTStateTrk(set bool) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTStateInv(set bool) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTStateDNAT(set bool) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTStateSNAT(set bool) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTMark(marks ...*CtMark) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchPktMark adds match condition for matching pkt_mark. If mask is nil, the mask should be not set in the OpenFlow
// message which is sent to OVS, and OVS should match the value exactly.
func (b *ofFlowBuilder) MatchPktMark(value uint32, mask *uint32) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchTunnelDst adds match condition for matching tun_dst or tun_ipv6_dst.
func (b *ofFlowBuilder) MatchTunnelDst(dstIP net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchTunnelID adds match condition for matching tun_id.
func (b *ofFlowBuilder) MatchTunnelID(tunnelID uint64) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func ctLabelRange(high, low uint64, rng *Range, match *ofctrl.FlowMatch) {
	_ = "STUB: not implemented"
	// [127..64] [63..0]
	//   high     low
	return
}

func (b *ofFlowBuilder) MatchCTLabelField(high, low uint64, field *CtLabel) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchCTZone(zone int) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchInPort adds match condition for matching in_port.
func (b *ofFlowBuilder) MatchInPort(inPort uint32) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchDstIP adds match condition for matching destination IP address.
func (b *ofFlowBuilder) MatchDstIP(ip net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchDstIPNet adds match condition for matching destination IP CIDR.
func (b *ofFlowBuilder) MatchDstIPNet(ipnet net.IPNet) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchICMPType(icmpType byte) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchICMPCode(icmpCode byte) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchICMPv6Type(icmp6Type byte) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchICMPv6Code(icmp6Code byte) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func maskToIP(mask net.IPMask) *net.IP { _ = "STUB: not implemented"; return nil }

// MatchSrcIP adds match condition for matching source IP address.
func (b *ofFlowBuilder) MatchSrcIP(ip net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchSrcIPNet adds match condition for matching source IP CIDR.
func (b *ofFlowBuilder) MatchSrcIPNet(ipnet net.IPNet) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchDstMAC adds match condition for matching destination MAC address.
func (b *ofFlowBuilder) MatchDstMAC(mac net.HardwareAddr) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchSrcMAC adds match condition for matching source MAC address.
func (b *ofFlowBuilder) MatchSrcMAC(mac net.HardwareAddr) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchARPSha adds match condition for matching ARP source host address.
func (b *ofFlowBuilder) MatchARPSha(mac net.HardwareAddr) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchARPTha adds match condition for matching ARP target host address.
func (b *ofFlowBuilder) MatchARPTha(mac net.HardwareAddr) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchARPSpa adds match condition for matching ARP source protocol address.
func (b *ofFlowBuilder) MatchARPSpa(ip net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchARPTpa adds match condition for matching ARP target protocol address.
func (b *ofFlowBuilder) MatchARPTpa(ip net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchARPOp adds match condition for matching ARP operator.
func (b *ofFlowBuilder) MatchARPOp(op uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchIPDSCP adds match condition for matching DSCP field in the IP header. Note, OVS use TOS to present DSCP, and
// the field name is shown as "nw_tos" with OVS command line, and the value is calculated by shifting the given value
// left 2 bits.
func (b *ofFlowBuilder) MatchIPDSCP(dscp uint8) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchConjID adds match condition for matching conj_id.
func (b *ofFlowBuilder) MatchConjID(value uint32) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchPriority(priority uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchProtocol adds match condition for matching protocol type.
func (b *ofFlowBuilder) MatchProtocol(protocol Protocol) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchIPProtocolValue adds match condition for IP protocol with the integer value.
func (b *ofFlowBuilder) MatchIPProtocolValue(isIPv6 bool, protoValue uint8) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchDstPort adds match condition for matching destination port in transport layer. OVS will match the port exactly
// if portMask is nil.
func (b *ofFlowBuilder) MatchDstPort(port uint16, portMask *uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchSrcPort adds match condition for matching source port in transport layer. OVS will match the port exactly
// if portMask is nil.
func (b *ofFlowBuilder) MatchSrcPort(port uint16, portMask *uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) MatchTCPFlags(flag, mask uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchCTSrcIP matches the source IPv4 address of the connection tracker original direction tuple. This match requires
// a match to valid connection tracking state as a prerequisite, and valid connection tracking state matches include
// "+new", "+est", "+rel" and "+trk-inv".
func (b *ofFlowBuilder) MatchCTSrcIP(ip net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchCTSrcIPNet is the same as MatchCTSrcIP but supports IP masking.
func (b *ofFlowBuilder) MatchCTSrcIPNet(ipNet net.IPNet) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchCTDstIP matches the destination IPv4 address of the connection tracker original direction tuple. This match
// requires a match to valid connection tracking state as a prerequisite, and valid connection tracking state matches
// include "+new", "+est", "+rel" and "+trk-inv".
func (b *ofFlowBuilder) MatchCTDstIP(ip net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchCTDstIPNet is the same as MatchCTDstIP but supports IP masking.
func (b *ofFlowBuilder) MatchCTDstIPNet(ipNet net.IPNet) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchCTSrcPort matches the transport source port of the connection tracker original direction tuple. This match requires
// a match to valid connection tracking state as a prerequisite, and valid connection tracking state matches include
// "+new", "+est", "+rel" and "+trk-inv".
func (b *ofFlowBuilder) MatchCTSrcPort(port uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchCTDstPort matches the transport destination port of the connection tracker original direction tuple. This match
// requires a match to valid connection tracking state as a prerequisite, and valid connection tracking state matches
// include "+new", "+est", "+rel" and "+trk-inv".
func (b *ofFlowBuilder) MatchCTDstPort(port uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MatchCTProtocol matches the IP protocol type of the connection tracker original direction tuple. This match requires
// a match to valid connection tracking state as a prerequisite, and a valid connection tracking state matches include
// "+new", "+est", "+rel" and "+trk-inv".
func (b *ofFlowBuilder) MatchCTProtocol(proto Protocol) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// Cookie sets cookie ID for the flow entry.
func (b *ofFlowBuilder) Cookie(cookieID uint64) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (b *ofFlowBuilder) Action() Action { _ = "STUB: not implemented"; return *new(Action) }
