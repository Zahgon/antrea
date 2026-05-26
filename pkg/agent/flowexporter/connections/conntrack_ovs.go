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
	"net/netip"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
)

// Following map is for converting protocol name (string) to protocol identifier
var (
	// Mapping is defined at https://github.com/torvalds/linux/blob/v5.9/include/uapi/linux/netfilter/nf_conntrack_common.h#L42
	conntrackStatusMap = map[string]uint32{
		"EXPECTED":      uint32(1),
		"SEEN_REPLY":    uint32(1 << 1),
		"ASSURED":       uint32(1 << 2),
		"CONFIRMED":     uint32(1 << 3),
		"SRC_NAT":       uint32(1 << 4),
		"DST_NAT":       uint32(1 << 5),
		"NAT_MASK":      uint32(1<<5 | 1<<4),
		"SEQ_ADJUST":    uint32(1 << 6),
		"SRC_NAT_DONE":  uint32(1 << 7),
		"DST_NAT_DONE":  uint32(1 << 8),
		"NAT_DONE_MASK": uint32(1<<8 | 1<<7),
		"DYING":         uint32(1 << 9),
		"FIXED_TIMEOUT": uint32(1 << 10),
		"TEMPLATE":      uint32(1 << 11),
		"UNTRACKED":     uint32(1 << 12),
		"HELPER":        uint32(1 << 13),
		"OFFLOAD":       uint32(1 << 14),
	}
)

// connTrackOvsCtl implements ConnTrackDumper. This supports OVS userspace datapath scenarios.
var _ ConnTrackDumper = new(connTrackOvsCtl)

type connTrackOvsCtl struct {
	nodeConfig           *config.NodeConfig
	serviceCIDRv4        netip.Prefix
	serviceCIDRv6        netip.Prefix
	ovsctlClient         ovsctl.OVSCtlClient
	isAntreaProxyEnabled bool
}

func NewConnTrackOvsAppCtl(nodeConfig *config.NodeConfig, serviceCIDRv4 netip.Prefix, serviceCIDRv6 netip.Prefix, isAntreaProxyEnabled bool) *connTrackOvsCtl {
	_ = "STUB: not implemented"
	return nil
}

// DumpFlows uses "ovs-appctl dpctl/dump-conntrack" to dump conntrack flows in the Antrea ZoneID.
func (ct *connTrackOvsCtl) DumpFlows(zoneFilter uint16) ([]*connection.Connection, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (ct *connTrackOvsCtl) ovsAppctlDumpConnections(zoneFilter uint16) ([]*connection.Connection, int, error) {
	_ = "STUB: not implemented"
	// Dump conntrack using ovs-appctl dpctl/dump-conntrack
	return nil, 0, nil
}

// Parse the output to get the flow strings and convert them to Antrea connections.

// flowStringToAntreaConnection parses the flow string and converts to Antrea connection.
// Example of flow string:
// "tcp,orig=(src=127.0.0.1,dst=127.0.0.1,sport=45218,dport=2379,packets=320108,bytes=24615344),reply=(src=127.0.0.1,dst=127.0.0.1,sport=2379,dport=45218,packets=239595,bytes=24347883),start=2020-07-24T05:07:03.998,id=3750535678,status=SEEN_REPLY|ASSURED|CONFIRMED|SRC_NAT_DONE|DST_NAT_DONE,timeout=86399,labels=0x200000001,protoinfo=(state_orig=ESTABLISHED,state_reply=ESTABLISHED,wscale_orig=7,wscale_reply=7,flags_orig=WINDOW_SCALE|SACK_PERM|MAXACK_SET,flags_reply=WINDOW_SCALE|SACK_PERM|MAXACK_SET)"
func flowStringToAntreaConnection(flow string, zoneFilter uint16) (*connection.Connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Indicator to populate reply or reverse fields

// Proto identifier

// dport field could be the last tuple field in ovs-dpctl output format.

// Append "Z" to meet RFC3339 standard because flow string doesn't have timezone information

// TODO: We didn't find stoptime related field in flow string right now, need to investigate how stoptime is recorded and dumped.

// Add leading zeros since DecodeString() expects the input string to have
// even length and we expect conn.Labels to be a []byte of length 16.

// retrieve tcpState from state or state_orig

// Add current time as stop time.

func hasAnyProto(text string) bool { _ = "STUB: not implemented"; return false }

func (ct *connTrackOvsCtl) GetMaxConnections() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func statusStringToStateFlag(status string) uint32 { _ = "STUB: not implemented"; return 0 }
