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
	"math/rand/v2"
	"net"

	"antrea.io/libOpenflow/protocol"
	"antrea.io/libOpenflow/util"
	"antrea.io/ofnet/ofctrl"
)

// #nosec G404: random number generator not used for security purposes
var pktRand = rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))

type ofPacketOutBuilder struct {
	pktOut  *ofctrl.PacketOut
	icmpID  *uint16
	icmpSeq *uint16
}

// SetSrcMAC sets the packet's source MAC with the provided value.
func (b *ofPacketOutBuilder) SetSrcMAC(mac net.HardwareAddr) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetDstMAC sets the packet's destination MAC with the provided value.
func (b *ofPacketOutBuilder) SetDstMAC(mac net.HardwareAddr) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetSrcIP sets the packet's source IP with the provided value.
func (b *ofPacketOutBuilder) SetSrcIP(ip net.IP) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetDstIP sets the packet's destination IP with the provided value.
func (b *ofPacketOutBuilder) SetDstIP(ip net.IP) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetIPProtocol sets IP protocol in the packet's IP header.
func (b *ofPacketOutBuilder) SetIPProtocol(proto Protocol) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetIPProtocolValue sets IP protocol in the packet's IP header with the
// intetger protocol value.
func (b *ofPacketOutBuilder) SetIPProtocolValue(isIPv6 bool, protoValue uint8) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetTTL sets TTL in the packet's IP header.
func (b *ofPacketOutBuilder) SetTTL(ttl uint8) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetIPFlags sets flags in the packet's IP header. IPv4 only.
func (b *ofPacketOutBuilder) SetIPFlags(flags uint16) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetIPHeaderID sets identifier field in the packet's IP header. IPv4 only.
func (b *ofPacketOutBuilder) SetIPHeaderID(id uint16) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetTCPSrcPort sets the source port in the packet's TCP header.
func (b *ofPacketOutBuilder) SetTCPSrcPort(port uint16) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetTCPDstPort sets the destination port in the packet's TCP header.
func (b *ofPacketOutBuilder) SetTCPDstPort(port uint16) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetTCPFlags sets the flags in the packet's TCP header.
func (b *ofPacketOutBuilder) SetTCPFlags(flags uint8) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) SetTCPSeqNum(seqNum uint32) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) SetTCPAckNum(ackNum uint32) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) SetTCPHdrLen(hdrLen uint8) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) SetTCPWinSize(winSize uint16) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) SetTCPData(data []byte) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetUDPSrcPort sets the source port in the packet's UDP header.
func (b *ofPacketOutBuilder) SetUDPSrcPort(port uint16) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetUDPDstPort sets the destination port in the packet's UDP header.
func (b *ofPacketOutBuilder) SetUDPDstPort(port uint16) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetICMPType sets the type in the packet's ICMP header.
func (b *ofPacketOutBuilder) SetICMPType(icmpType uint8) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetICMPCode sets the code in the packet's ICMP header.
func (b *ofPacketOutBuilder) SetICMPCode(icmpCode uint8) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetICMPID sets the identifier in the packet's ICMP header.
func (b *ofPacketOutBuilder) SetICMPID(id uint16) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetICMPSequence sets the sequence number in the packet's ICMP header.
func (b *ofPacketOutBuilder) SetICMPSequence(seq uint16) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) SetICMPData(data []byte) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) SetUDPData(data []byte) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetInport sets the in_port field of the packetOut message.
func (b *ofPacketOutBuilder) SetInport(inPort uint32) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetOutport sets the output port of the packetOut message. If the message is expected to go through OVS pipeline
// from table0, use openflow15.P_TABLE, which is also the default value.
func (b *ofPacketOutBuilder) SetOutport(outport uint32) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// SetL4Packet sets the L4 packet of the packetOut message. It provides a generic function to create a packet
// of protocol other than TCP/UDP/ICMP.
func (b *ofPacketOutBuilder) SetL4Packet(packet util.Message) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) SetEthPacket(packet *protocol.Ethernet) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// AddSetIPTOSAction sets the IP_TOS field in the packet-out message. The action clears the two ECN bits as 0,
// and only 2-7 bits of the DSCP field in IP header is set.
func (b *ofPacketOutBuilder) AddSetIPTOSAction(data uint8) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) AddLoadRegMark(mark *RegMark) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) AddResubmitAction(inPort *uint16, table *uint8) PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

func (b *ofPacketOutBuilder) Done() *ofctrl.PacketOut { _ = "STUB: not implemented"; return nil }

// Entire ethernet packet is provided. No need to fill L3/L4 header.

// #nosec G404: random number generator not used for security purposes

// #nosec G404: random number generator not used for security purposes

// #nosec G404: random number generator not used for security purposes

// Set IP version in the IP Header.

// #nosec G404: random number generator not used for security purposes

// #nosec G404: random number generator not used for security purposes

// Set IPv6 version in the IP Header.

func (b *ofPacketOutBuilder) setICMPData() { _ = "STUB: not implemented"; return }

func (b *ofPacketOutBuilder) ipHeaderChecksum() uint16 { _ = "STUB: not implemented"; return 0 }

func (b *ofPacketOutBuilder) icmpHeaderChecksum() uint16 { _ = "STUB: not implemented"; return 0 }

func (b *ofPacketOutBuilder) tcpHeaderChecksum() uint16 { _ = "STUB: not implemented"; return 0 }

func (b *ofPacketOutBuilder) udpHeaderChecksum() uint16 { _ = "STUB: not implemented"; return 0 }

// From RFC 768:
// If the computed checksum is zero, it is transmitted as all ones (the
// equivalent in one's complement arithmetic). An all zero transmitted
// checksum value means that the transmitter generated no checksum (for
// debugging or for higher level protocols that don't care).

func (b *ofPacketOutBuilder) igmpHeaderChecksum() uint16 { _ = "STUB: not implemented"; return 0 }

func (b *ofPacketOutBuilder) generatePseudoHeader(length uint16) []byte {
	_ = "STUB: not implemented"
	return nil
}

func checksum(data []byte) uint16 { _ = "STUB: not implemented"; return 0 }
