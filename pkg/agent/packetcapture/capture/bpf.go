// Copyright 2024 Antrea Authors.
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

package capture

import (
	"net"

	"golang.org/x/net/bpf"

	crdv1alpha1 "antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
)

const (
	lengthByte    int    = 1
	lengthHalf    int    = 2
	lengthWord    int    = 4
	etherTypeIPv4 uint32 = 0x0800
	etherTypeIPv6 uint32 = 0x86DD

	jumpMask                 uint32 = 0x1fff
	ip4SourceAddrOffset      uint32 = 26
	ip4DestinationAddrOffset uint32 = 30
	ip4SourcePort            uint32 = 14
	ip4DestinationPort       uint32 = 16
	ip4HeaderSize            uint32 = 14
	ip4HeaderFlags           uint32 = 20

	ip6HeaderOffset            uint32 = 14
	ip6HeaderSize              uint32 = 40
	ip6NextHeaderOffset        uint32 = ip6HeaderOffset + 6             // 20
	ip6SourceAddrOffset        uint32 = ip6HeaderOffset + 8             // 22
	ip6DestinationAddrOffset   uint32 = ip6HeaderOffset + 24            // 38
	ip6L4HeaderOffset          uint32 = ip6HeaderOffset + ip6HeaderSize // 54
	ip6SourcePort              uint32 = 0
	ip6DestinationPort         uint32 = 2
	ip6TCPFlags                uint32 = 13
	ip6ICMPv6Type              uint32 = 0
	ip6ICMPv6Code              uint32 = 1
	ip6FragmentNextHeader      uint32 = 44 // IPv6 Fragment Extension Header
	ip6FragExtInstructionCount int    = 3  // number of extra instructions for fragment header handling
)

var (
	returnDrop              = bpf.RetConstant{Val: 0}
	returnKeep              = bpf.RetConstant{Val: 0x40000}
	loadIPv4SourcePort      = bpf.LoadIndirect{Off: ip4SourcePort, Size: lengthHalf}
	loadIPv4DestinationPort = bpf.LoadIndirect{Off: ip4DestinationPort, Size: lengthHalf}
	loadEtherKind           = bpf.LoadAbsolute{Off: 12, Size: lengthHalf}
	loadIPv4Protocol        = bpf.LoadAbsolute{Off: 23, Size: lengthByte}
	loadIPv4TCPFlags        = bpf.LoadIndirect{Off: 27, Size: lengthByte}
	loadIPv4ICMPType        = bpf.LoadIndirect{Off: 14, Size: lengthByte}
	loadIPv4ICMPCode        = bpf.LoadIndirect{Off: 15, Size: lengthByte}

	loadIPv6NextHeader      = bpf.LoadAbsolute{Off: ip6NextHeaderOffset, Size: lengthByte}
	loadIPv6SourcePort      = bpf.LoadAbsolute{Off: ip6L4HeaderOffset + ip6SourcePort, Size: lengthHalf}
	loadIPv6DestinationPort = bpf.LoadAbsolute{Off: ip6L4HeaderOffset + ip6DestinationPort, Size: lengthHalf}
	loadIPv6TCPFlags        = bpf.LoadAbsolute{Off: ip6L4HeaderOffset + ip6TCPFlags, Size: lengthByte}
	loadIPv6ICMPv6Type      = bpf.LoadAbsolute{Off: ip6L4HeaderOffset + ip6ICMPv6Type, Size: lengthByte}
	loadIPv6ICMPv6Code      = bpf.LoadAbsolute{Off: ip6L4HeaderOffset + ip6ICMPv6Code, Size: lengthByte}
)

// Supported protocol strings (must be uppercase, since validation uses strings.ToUpper).
// These values are matched against user input in the controller.
var ProtocolMap = map[string]uint32{
	"UDP":    17,
	"TCP":    6,
	"ICMP":   1,
	"ICMPV6": 58,
}

var ICMPMsgTypeMap = map[crdv1alpha1.ICMPMsgType]uint32{
	crdv1alpha1.ICMPMsgTypeEcho:       8,
	crdv1alpha1.ICMPMsgTypeEchoReply:  0,
	crdv1alpha1.ICMPMsgTypeDstUnreach: 3,
	crdv1alpha1.ICMPMsgTypeTimexceed:  11,
}

var ICMPv6MsgTypeMap = map[crdv1alpha1.ICMPv6MsgType]uint32{
	crdv1alpha1.ICMPv6MsgTypeEcho:         128,
	crdv1alpha1.ICMPv6MsgTypeEchoReply:    129,
	crdv1alpha1.ICMPv6MsgTypeDstUnreach:   1,
	crdv1alpha1.ICMPv6MsgTypeTimexceed:    3,
	crdv1alpha1.ICMPv6MsgTypePacketTooBig: 2,
	crdv1alpha1.ICMPv6MsgTypeParamProblem: 4,
}

// tcpFlagsFilter represents a TCP flag match condition with a value and mask.
type tcpFlagsFilter struct {
	flag uint32
	mask uint32
}

// icmpFilter represents an ICMP or ICMPv6 message filter with type and optional code.
type icmpFilter struct {
	icmpType uint32
	icmpCode *uint32
}

// transportFilters holds the parsed transport-layer filter criteria
// extracted from the PacketCapture CRD spec.
type transportFilters struct {
	srcPort  uint16
	dstPort  uint16
	tcpFlags []tcpFlagsFilter
	icmp     []icmpFilter
}

// hasTransportFilters returns true if any L4-level filters (ports, flags, ICMP messages)
// are configured. This is used to decide whether to add IPv6 extension header handling,
// which is only needed for protocol-only filters.
//
// Example: For 'ip6 proto 58' (ICMPv6 only, no transport header), returns false,
// so Fragment Extension Header checks are added. For 'ip6 proto 6 and dst port 80',
// returns true, so Fragment checks are omitted.
func hasTransportFilters(packet *crdv1alpha1.Packet) bool { _ = "STUB: not implemented"; return false }

// ipFamilyHandler encapsulates protocol-specific constants and filter compilation logic
// to allow for a unified, protocol-agnostic packet filter generation function.
type ipFamilyHandler struct {
	etherType             uint32
	addressChunks         int // IPv4: 1, IPv6: 4
	sourceAddrOffset      uint32
	destinationAddrOffset uint32

	loadProtocol        bpf.Instruction
	loadSourcePort      bpf.Instruction
	loadDestinationPort bpf.Instruction
	loadTCPFlags        bpf.Instruction
	loadICMPType        bpf.Instruction
	loadICMPCode        bpf.Instruction
}

// ipv4Handler provides the IPv4-specific implementations for the ipFamilyHandler.
var ipv4Handler = &ipFamilyHandler{
	etherType:             etherTypeIPv4,
	addressChunks:         1,
	sourceAddrOffset:      ip4SourceAddrOffset,
	destinationAddrOffset: ip4DestinationAddrOffset,

	loadProtocol:        loadIPv4Protocol,
	loadSourcePort:      loadIPv4SourcePort,
	loadDestinationPort: loadIPv4DestinationPort,
	loadTCPFlags:        loadIPv4TCPFlags,
	loadICMPType:        loadIPv4ICMPType,
	loadICMPCode:        loadIPv4ICMPCode,
}

// ipv6Handler provides the IPv6-specific implementations for the ipFamilyHandler.
var ipv6Handler = &ipFamilyHandler{
	etherType:             etherTypeIPv6,
	addressChunks:         4,
	sourceAddrOffset:      ip6SourceAddrOffset,
	destinationAddrOffset: ip6DestinationAddrOffset,

	loadProtocol:        loadIPv6NextHeader,
	loadSourcePort:      loadIPv6SourcePort,
	loadDestinationPort: loadIPv6DestinationPort,
	loadTCPFlags:        loadIPv6TCPFlags,
	loadICMPType:        loadIPv6ICMPv6Type,
	loadICMPCode:        loadIPv6ICMPv6Code,
}

func loadIPv4HeaderOffset(skipTrue uint8) []bpf.Instruction { _ = "STUB: not implemented"; return nil }

// flags+fragment offset, since we need to calc where the src/dst port is
// check if there is a L4 header
// calculate the size of IP header

func compareProtocolIP(etherType uint32, skipTrue, skipFalse uint8) bpf.Instruction {
	_ = "STUB: not implemented"
	return *new(bpf.Instruction)
}

func compareProtocol(protocol uint32, skipTrue, skipFalse uint8) bpf.Instruction {
	_ = "STUB: not implemented"
	return *new(bpf.Instruction)
}

// appendProtocolFilters appends protocol-related checks and computes jump offsets
// based on the current instruction length.
//
// For IPv6 protocol-only filters (no transport-layer filters like ports, flags, or ICMP),
// tcpdump/libpcap adds extra instructions to handle the IPv6 Fragment Extension Header.
// This is because the Next Header field may point to a Fragment header (44) rather
// than the actual transport protocol, so we must check both cases.
//
// Example: 'ip6 proto 58' (ICMPv6 protocol only) generates:
//
//	(000) ldh      [12]                             # Load EtherType
//	(001) jeq      #0x86dd     jt 2    jf 8         # Is IPv6?
//	(002) ldb      [20]                             # Load Next Header
//	(003) jeq      #0x3a       jt 7    jf 4         # Is ICMPv6 (58)? → MATCH, else check Fragment
//	(004) jeq      #0x2c       jt 5    jf 8         # Is Fragment (44)? → check inner, else DROP
//	(005) ldb      [54]                             # Load inner Next Header from Fragment Ext Header
//	(006) jeq      #0x3a       jt 7    jf 8         # Is ICMPv6 (58)? → MATCH, else DROP
//	(007) ret      #262144                          # MATCH
//	(008) ret      #0                               # DROP
//
// In contrast, when transport filters are present (e.g., ports), the Fragment Extension
// Header check is omitted because tcpdump/libpcap does not add it:
//
// Example: 'ip6 proto 6 and dst port 80' generates:
//
//	(000) ldh      [12]                             # Load EtherType
//	(001) jeq      #0x86dd     jt 2    jf 7         # Is IPv6?
//	(002) ldb      [20]                             # Load Next Header
//	(003) jeq      #0x6        jt 4    jf 7         # Is TCP (6)?
//	(004) ldh      [56]                             # Load TCP Dst Port
//	(005) jeq      #0x50       jt 6    jf 7         # Is port 80?
//	(006) ret      #262144                          # MATCH
//	(007) ret      #0                               # DROP
func appendProtocolFilters(inst []bpf.Instruction, handler *ipFamilyHandler, proto uint32, hasTransport bool, size uint8) []bpf.Instruction {
	_ = "STUB: not implemented"
	return nil
}

// For IPv6 without transport filters, include Fragment Extension Header handling
// to match libpcap/tcpdump behavior.

// If protocol matches directly, skip past the extension header check.

// If Next Header is NOT Fragment (44), jump to drop.

// Load the protocol from the byte following the Fragment Extension Header,
// which is at the same offset as the normal L4 header.

// Check if the inner protocol matches. Jump to drop if not.

// getAddressChunk abstracts the process of extracting a 4-byte chunk from an IP address,
// handling the structural differences between IPv4 (one chunk) and IPv6 (four chunks).
func (h *ipFamilyHandler) getAddressChunk(ip net.IP, chunkIndex int) uint32 {
	_ = "STUB: not implemented"
	return 0
}

// calculateSkipOffset determines the correct 'SkipFalse' jump offset for an IP address chunk
// comparison. When checking bidirectional traffic ('Both' direction), a failed check for the first
// direction should not jump to the end (drop), but rather to the start of the check for the other
// direction. jumpToReturnTraffic: If true, calculate the offset to jump to the return traffic block.
// chunkIndex: The current 4-byte chunk index of the IP being checked (0-3 for IPv6).
func (h *ipFamilyHandler) calculateSkipOffset(chunkIndex int, skipFalse, skipToEnd uint8, jumpToReturnTraffic bool) uint8 {
	_ = "STUB: not implemented"
	return 0

	// calculate the relative jump offsets (SkipFalse) that decrease by 2 per chunk
	// for the srcIP and dstIP cases.
}

func (h *ipFamilyHandler) countAddrForSkipFalse(srcIP, dstIP net.IP) uint8 {
	_ = "STUB: not implemented"

	// We keep track of this count so we can correctly calculate the
	// relative jump offsets (SkipFalse) that decrease by 2 per chunk
	// for the srcIP and dstIP cases.
	return 0
}

func calculateSkipFalse(handler *ipFamilyHandler, srcIP, dstIP net.IP, transport *transportFilters) uint8 {
	_ = "STUB: not implemented"
	return 0
}

// load fragment offset

// ret keep

// compileIPFilters generates the BPF instructions for matching source and/or destination
// IP addresses. It is protocol-agnostic, using the handler to abstract the differences
// between IPv4 (1 chunk) and IPv6 (4 chunks). It also manages the complex jump logic
// required for bidirectional traffic matching.
func compileIPFilters(handler *ipFamilyHandler, srcIP, dstIP net.IP, size, curLen, skipFalse uint8, needsOtherTrafficDirectionCheck bool) []bpf.Instruction {
	_ = "STUB: not implemented"
	return nil

	// calculate skip size to jump to the final instruction (NO MATCH)
}

// needsOtherTrafficDirectionCheck indicates if we need to check whether the packet belongs to the
// return traffic flow when source IP from the packet spec and packet header don't match and we are
// capturing packets in both direction. If true, we calculate skipFalse to jump to the instruction
// that compares the destination IP from the packet spec with the loaded source IP from the packet
// header.

// If the dstIP doesn't match, skip to the end (no match), unless a srcIP was not provided and
// we need to check the other direction of traffic (reply). If we don't need to check the other
// direction of traffic, we can already say the packet is not a match. If a srcIP was provided
// and get to that stage in the filter (dstIP check), then it means the srcIP was a match: if
// the srcIP matches but not the dstIP, we don't need to check the other direction of traffic
// (guaranteed no match).

// compileTransportFilters generates BPF instructions for filtering transport-layer
// traffic based on ports, TCP flags, ICMP and ICMPv6 messages.
func compileTransportFilters(handler *ipFamilyHandler, size, curLen uint8, transport *transportFilters) []bpf.Instruction {
	_ = "STUB: not implemented"
	return nil

	// calculate skip size to jump to the final instruction (NO MATCH)
}

// For fragment checks and IP header length calculation to find the L4 header offset,
// as the IP header can have variable options.

// tcp flags

// last flag match condition

// ICMP and ICMPv6 message filters.

// return (accept)

// compilePacketFilter acts as the main entry point for BPF filter generation.
// It inspects the IP family specified in the CRD and dispatches the request
// to the unified compiler with the appropriate protocol-specific handler
// (ipv4Handler for IPv4, ipv6Handler for IPv6).
func compilePacketFilter(packetSpec *crdv1alpha1.Packet, srcIP, dstIP net.IP, direction crdv1alpha1.CaptureDirection) []bpf.Instruction {
	_ = "STUB: not implemented"
	return nil
}

// compileGenericPacketFilter compiles the CRD spec to BPF instructions using a
// protocol-specific handler to manage differences between IPv4 and IPv6.
func compileGenericPacketFilter(handler *ipFamilyHandler, packetSpec *crdv1alpha1.Packet, srcIP, dstIP net.IP, direction crdv1alpha1.CaptureDirection) []bpf.Instruction {
	_ = "STUB: not implemented"
	return nil
}

// Start with checking the EtherType.

// skip means how many instructions we need to skip if the compare fails.
// for example, for now we have 2 instructions, and the total size is 17, if ipv4
// check failed, we need to jump to the end (ret #0), skip 17-3=14 instructions.
// if check succeed, skipTrue means we jump to the next instruction. Here 3 means we
// have 3 instructions so far.

// For IPv6 with transport-level filters and at least one IP filter,
// defer protocol checks until after IP checks for one-way directions.
// This better aligns with tcpdump output ordering for many complex
// expressions where L3 constraints are evaluated before protocol checks.
// IPv4 does not need this reordering: tcpdump/libpcap always emits
// IPv4 protocol checks immediately after the EtherType check,
// regardless of filter complexity.

// ports, TCP flags, ICMP and ICMPv6 messages

// default to flag if not specified

// return (drop)

// We need to figure out how long the instruction list will be first. It will be used in the instructions' jump case.
// For example, If you provide all the filters supported by `PacketCapture`, it will end with the following BPF filter string:
// 'ip proto 6 and src host 127.0.0.1 and dst host 127.0.0.1 and src port 123 and dst port 124'
// And using `tcpdump -i <device> '<filter>' -d` will generate the following BPF instructions:
// (000) ldh      [12]                                     # Load 2B at 12 (Ethertype)
// (001) jeq      #0x800           jt 2	jf 16              # Ethertype: If IPv4, goto #2, else #16
// (002) ldb      [23]                                     # Load 1B at 23 (IPv4 Protocol)
// (003) jeq      #0x6             jt 4	jf 16              # IPv4 Protocol: If TCP, goto #4, #16
// (004) ld       [26]                                     # Load 4B at 26 (source address)
// (005) jeq      #0x7f000001      jt 6	jf 16              # If bytes match(127.0.0.1), goto #6, else #16
// (006) ld       [30]                                     # Load 4B at 30 (dest address)
// (007) jeq      #0x7f000001      jt 8	jf 16              # If bytes match(127.0.0.1), goto #8, else #16
// (008) ldh      [20]                                     # Load 2B at 20 (13b Fragment Offset)
// (009) jset     #0x1fff          jt 16	jf 10          # Use 0x1fff as a mask for fragment offset; If fragment offset != 0, #10, else #16
// (010) ldxb     4*([14]&0xf)                             # x = IP header length
// (011) ldh      [x + 14]                                 # Load 2B at x+14 (TCP Source Port)
// (012) jeq      #0x7b            jt 13	jf 16		   # TCP Source Port: If 123, goto #13, else #16
// (013) ldh      [x + 16]                                 # Load 2B at x+16 (TCP dst port)
// (014) jeq      #0x7c            jt 15	jf 16		   # TCP dst port: If 123, goto #15, else #16
// (015) ret      #262144                                  # MATCH
// (016) ret      #0                                       # NOMATCH

// When capturing return traffic also (i.e., both src -> dst and dst -> src), the filter might look like this:
// 'ip proto 6 and ((src host 10.244.1.2 and dst host 10.244.1.3 and src port 123 and dst port 124) or (src host 10.244.1.3 and dst host 10.244.1.2 and src port 124 and dst port 123))'
// And using `tcpdump -i <device> '<filter>' -d` will generate the following BPF instructions:
// Ethertype, IPv4 protocol...
// (004) ld       [26]									   # Load 4B at 26 (source address)
// (005) jeq      #0xaf40102       jt 6	jf 15			   # If bytes match(10.244.1.2), goto #6, else #15
// (006) ld       [30]									   # Load 4B at 30 (dest address)
// (007) jeq      #0xaf40103       jt 8	jf 26			   # If bytes match(10.244.1.3), goto #8, else #26
// Check fragment offset and calculate IP header length...
// (011) ldh      [x + 14]								   # Load 2B at x+14 (TCP Source Port)
// (012) jeq      #0x7b            jt 13	jf 26		   # TCP Source Port: If 123, goto #13, else #26
// (013) ldh      [x + 16]								   # Load 2B at x+16 (TCP dst port)
// (014) jeq      #0x7c            jt 25	jf 26		   # TCP dst port: If 123, goto #25, else #26
// (015) jeq      #0xaf40103       jt 16	jf 26		   # If bytes match(10.244.1.3), goto #16, else #26
// (016) ld       [30]									   # Load 4B at 30 (return traffic dest address)
// (017) jeq      #0xaf40102       jt 18	jf 26		   # If bytes match(10.244.1.2), goto #18, else #26
// Check fragment offset and calculate IP header length...
// (021) ldh      [x + 14]								   # Load 2B at x+14 (TCP Source Port)
// (022) jeq      #0x7c            jt 23	jf 26		   # TCP Source Port: If 124, goto #23, else #26
// (023) ldh      [x + 16]								   # Load 2B at x+16 (TCP dst port)
// (024) jeq      #0x7b            jt 25	jf 26		   # TCP dst port: If 123, goto #25, else #26
// (025) ret      #262144								   # MATCH
// (026) ret      #0									   # NOMATCH

// For simpler code generation in 'Both' direction, an extra instruction to accept the packet is added after instruction 014.
// The final instruction set looks like this:
// Ethertype, IPv4 protocol...
// Source IP, Destination IP, Source port, Destination port...
// (015) ret      #262144								   # MATCH
// Source IP, Destination IP, Source port, Destination port for return traffic...
// (026) ret      #262144								   # MATCH
// (027) ret      #0									   # NOMATCH

// To capture all TCP packets from 10.0.0.4 to 10.0.0.5 with either SYN or ACK flags set, the filter would be:
// 'ip proto 6 and src host 10.0.0.4 and dst host 10.0.0.5 and ((tcp[tcpflags] & tcp-syn) == tcp-syn) or ((tcp[tcpflags] & tcp-ack) == tcp-ack))'
// And using `tcpdump -i <device> '<filter>' -d` will generate the following BPF instructions:
// Ethertype, IPv4 protocol...
// Source and Destination IP...
// Check fragment offset and calculate IP header length...
// (011) ldh      [x + 27]                                 # Load 1B at x+27 (TCP Flags)
// (012) and	  0x2			            			   # Apply a bitwise AND with 0x2 (SYN flag)
// (013) jeq      #0x2             jt 17    jf 14          # If SYN is set, goto #17, else #14
// (014) ldh      [x + 27]                                 # Again load 1B at x+27 (TCP Flags)
// (015) and	  0x10			            			   # Apply a bitwise AND with 0x10 (ACK flag)
// (016) jeq      #0x10            jt 17    jf 18          # If ACK is set, goto #17, else #18
// (017) ret      #262144                                  # MATCH
// (018) ret      #0                                       # NOMATCH

// To capture ICMP destination unreachable (host unreachable) packets from 10.0.0.1 to 10.0.0.2, the tcpdump filter would be:
// 'ip proto 1 and src host 10.0.0.1 and dst host 10.0.0.2 and icmp[0]=3 and icmp[1]=1'
// And using `tcpdump -i <device> '<filter>' -d` will generate the following BPF instructions:
// Ethertype, IPv4 protocol...
// Source and Destination IP...
// Check fragment offset and calculate IP header length...
// (011) ldb      [x + 14]								   # Load 1B at x+14 (ICMP Type)
// (012) jeq      #0x3             jt 13   jf 16		   # ICMP Type: If 3, goto #13, else #16
// (013) ldb      [x + 15]								   # Load 1B at x+15 (ICMP Code)
// (014) jeq      #0x1             jt 15   jf 16		   # ICMP Code: If 1, goto #15, else #16
// (015) ret      #262144								   # MATCH
// (016) ret      #0									   # NOMATCH

// For IPv6, the filter is similar but accounts for the 16-byte addresses, which are
// loaded and compared in 4-byte chunks. There is also no need for the fragment
// offset calculation to find the L4 header.
// 'ip6 proto 6 and src host fd00::1 and dst host fd00::2 and src port 123 and dst port 124'
// And using `tcpdump -i <device> '<filter>' -d` will generate the following BPF instructions:
// (000) ldh      [12]                                     # Load 2B at 12 (Ethertype)
// (001) jeq      #0x86dd          jt 2	jf 25              # Ethertype: If IPv6, goto #2, else #25
// (002) ldb      [20]                                     # Load 1B at 20 (Next Header)
// (003) jeq      #0x6             jt 4	jf 25              # Next Header: If TCP, goto #4, else #25
// (004) ld       [22]                                     # Load 4B at 22 (Src Addr chunk 1)
// (005) jeq      #0xfd000000      jt 6	jf 25              # If chunk 1 matches, goto #6, else #25
// (006) ld       [26]                                     # Load 4B at 26 (Src Addr chunk 2)
// (007) jeq      #0x0             jt 8	jf 25              # If chunk 2 matches, goto #8, else #25
// (008) ld       [30]                                     # Load 4B at 30 (Src Addr chunk 3)
// (009) jeq      #0x0             jt 10	jf 25          # If chunk 3 matches, goto #10, else #25
// (010) ld       [34]                                     # Load 4B at 34 (Src Addr chunk 4)
// (011) jeq      #0x1             jt 12	jf 25          # If chunk 4 matches (fd00::1), goto #12, else #25
// (012) ld       [38]                                     # Load 4B at 38 (Dst Addr chunk 1)
// (013) jeq      #0xfd000000      jt 14	jf 25          # If chunk 1 matches, goto #14, else #25
// (014) ld       [42]                                     # Load 4B at 42 (Dst Addr chunk 2)
// (015) jeq      #0x0             jt 16	jf 25          # If chunk 2 matches, goto #16, else #25
// (016) ld       [46]                                     # Load 4B at 46 (Dst Addr chunk 3)
// (017) jeq      #0x0             jt 18	jf 25          # If chunk 3 matches, goto #18, else #25
// (018) ld       [50]                                     # Load 4B at 50 (Dst Addr chunk 4)
// (019) jeq      #0x2             jt 20	jf 25          # If chunk 4 matches (fd00::2), goto #20, else #25
// (020) ldh      [54]                                     # Load 2B at 54 (TCP Src Port)
// (021) jeq      #0x7b            jt 22	jf 25		   # TCP Src Port: If 123, goto #22, else #25
// (022) ldh      [56]                                     # Load 2B at 56 (TCP Dst port)
// (023) jeq      #0x7c            jt 24	jf 25		   # TCP Dst port: If 124, goto #24, else #25
// (024) ret      #262144                                # MATCH
// (025) ret      #0                                       # NOMATCH

func calculateInstructionsSize(handler *ipFamilyHandler, packet *crdv1alpha1.Packet, srcIP, dstIP net.IP, direction crdv1alpha1.CaptureDirection) int {
	_ = "STUB: not implemented"

	// load ethertype
	return 0
}

// ip check

// load + compare for each chunk

// load + compare for each chunk

// protocol check

// IPv6 Fragment Extension Header handling adds 3 extra instructions
// when there are no transport-layer filters (ports, flags, ICMP).

// load Fragment Offset

// every TCP Flag match condition will have 3 instructions - load, bitwise AND, compare

// load Fragment Offset

// load Fragment Offset

// load icmp type

// compare icmp type

// load + compare icmp code

// load icmpv6 type

// compare icmpv6 type

// load + compare icmpv6 code

// extra returnKeep

// src and dst ip (return traffic)

// ret command
