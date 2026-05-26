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

package openflow

import (
	"antrea.io/libOpenflow/protocol"
	"antrea.io/libOpenflow/util"
	"antrea.io/ofnet/ofctrl"
)

const (
	icmpEchoRequestType  uint8 = 8
	icmp6EchoRequestType uint8 = 128
	// tcpStandardHdrLen is the TCP header length without options.
	tcpStandardHdrLen uint8 = 5
)

func GetTCPHeaderData(ipPkt util.Message) (tcpSrcPort, tcpDstPort uint16, tcpSeqNum, tcpAckNum uint32, tcpHdrLen uint8, tcpFlags uint8, tcpWinSize uint16, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0, 0, 0, nil
}

// GetTCPPacketFromIPMessage gets a TCP struct from an IP message.
func GetTCPPacketFromIPMessage(ipPkt util.Message) (tcpPkt *protocol.TCP, err error) {
	_ = "STUB: not implemented"
	return nil,

		// Transfer Buffer to TCP
		nil
}

func GetTCPDNSData(tcpPkt *protocol.TCP) (data []byte, length int, err error) {
	_ = "STUB: not implemented"
	// TCP.HdrLen is 4-octet unit indicating the length of TCP header including options.
	return nil, 0, nil
}

// Move two more octet.
// From RFC 7766:
// DNS clients and servers SHOULD pass the two-octet length field, and
// the message described by that length field, to the TCP layer at the
// same time (e.g., in a single "write" system call) to make it more
// likely that all the data will be transmitted in a single TCP segment.

func GetUDPHeaderData(ipPkt util.Message) (udpSrcPort, udpDstPort uint16, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func getICMPHeaderData(ipPkt util.Message) (icmpType, icmpCode uint8, icmpEchoID, icmpEchoSeq uint16, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func ParsePacketIn(pktIn *ofctrl.PacketIn) (*Packet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IPv6 header includes only playload length. Add 40 to count in
// the IPv6 header length.

// Not an IP packet.
