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

package multicast

import (
	"net"
	"sync"
	"time"

	"antrea.io/libOpenflow/protocol"
	"antrea.io/libOpenflow/util"
	"antrea.io/ofnet/ofctrl"
	apitypes "k8s.io/apimachinery/pkg/types"

	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/agent/types"
)

const (
	IGMPProtocolNumber = 2
)

var (
	// igmpMaxResponseTime is the maximum time allowed before sending a responding report which is used for the
	// "Max Resp Code" field in the IGMP query message. It is also the maximum time to wait for the IGMP report message
	// when checking the last group member.
	igmpMaxResponseTime = time.Second * 10
	// igmpQueryDstMac is the MAC address used in the dst MAC field in the IGMP query message
	igmpQueryDstMac, _ = net.ParseMAC("01:00:5e:00:00:01")
	// igmpReportDstMac is the MAC address used in the dst MAC field in the IGMP report message
	igmpReportDstMac, _ = net.ParseMAC("01:00:5e:00:00:16")
)

type IGMPSnooper struct {
	ofClient      openflow.Client
	ifaceStore    interfacestore.InterfaceStore
	eventCh       chan *mcastGroupEvent
	validator     types.McastNetworkPolicyController
	queryInterval time.Duration
	queryVersions []uint8
	// igmpReportANNPStats is a map that saves AntreaNetworkPolicyStats of IGMP report packets.
	// The map can be interpreted as
	// map[UID of the AntreaNetworkPolicy]map[name of AntreaNetworkPolicy rule]statistics of rule.
	igmpReportANNPStats      map[apitypes.UID]map[string]*types.RuleMetric
	igmpReportANNPStatsMutex sync.Mutex
	// Similar to igmpReportANNPStats, it stores ACNP stats for IGMP reports.
	igmpReportACNPStats      map[apitypes.UID]map[string]*types.RuleMetric
	igmpReportACNPStatsMutex sync.Mutex
	encapEnabled             bool
}

func (s *IGMPSnooper) parseSrcInterface(pktIn *ofctrl.PacketIn) (*interfacestore.InterfaceConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *IGMPSnooper) queryIGMP(group net.IP) error { _ = "STUB: not implemented"; return nil }

// outPort sets the output port of the packetOut message. We expect the message to go through OVS pipeline
// from table0. The final OpenFlow message will use a standard OpenFlow port number OFPP_TABLE = 0xfffffff9 corrected
// by ofnet.

func (s *IGMPSnooper) validate(event *mcastGroupEvent, igmpType uint8, packetInData protocol.Ethernet) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// Return true directly if there is no validator.
		nil
}

// MulticastValidator only validates the IGMP report message sent from Pods. The report message received from tunnel
// port is sent from Antrea Agent on a different Node, and returns true directly.

// It shall drop the packet if function Validate returns error

func (s *IGMPSnooper) validatePacketAndNotify(event *mcastGroupEvent, igmpType uint8, packetInData protocol.Ethernet) {
	_ = "STUB: not implemented"
	return
}

// Antrea Agent does not remove the Pod from the OpenFlow group bucket immediately when an error is returned,
// but it will be removed when after timeout (Controller.mcastGroupTimeout)

// If any rule is desired to drop the traffic, Antrea Agent removes the Pod from
// the OpenFlow group bucket directly

func (s *IGMPSnooper) addToIGMPReportNPStatsMap(item types.IGMPNPRuleInfo, packetLen uint64) {
	_ = "STUB: not implemented"
	return
}

// WARNING: This func will reset the saved stats.
func (s *IGMPSnooper) collectStats() (igmpANNPStats, igmpACNPStats map[apitypes.UID]map[string]*types.RuleMetric) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *IGMPSnooper) sendIGMPReport(groupRecordType uint8, groups []net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *IGMPSnooper) generateIGMPReportPacket(groupRecordType uint8, groups []net.IP) (util.Message, error) {
	_ = "STUB: not implemented"
	return *new(util.Message), nil
}

func (s *IGMPSnooper) sendIGMPJoinReport(groups []net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *IGMPSnooper) sendIGMPLeaveReport(groups []net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *IGMPSnooper) HandlePacketIn(pktIn *ofctrl.PacketIn) error {
	_ = "STUB: not implemented"
	return nil
}

// If an IGMP report arrives via a tunnel, extract the Node IP from its source IP.
// This works because for remote IGMP reports (sent via packet-out), the source IP
// is set to the Node's transport IP (see pkg/agent/openflow/client.go SendIGMPRemoteReportPacketOut).

func generateIGMPQueryPacket(group net.IP, version uint8, queryInterval time.Duration) (util.Message, error) {
	_ = "STUB: not implemented"
	// The max response time field in IGMP protocol uses a value in units of 1/10 second.
	// See https://datatracker.ietf.org/doc/html/rfc2236 and https://datatracker.ietf.org/doc/html/rfc3376
	return *new(util.Message), nil
}

func parseIPv4Packet(pkt *protocol.Ethernet) (*protocol.IPv4, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseIGMPPacket(ipPacket *protocol.IPv4) (protocol.IGMPMessage, error) {
	_ = "STUB: not implemented"
	return *new(protocol.IGMPMessage), nil
}

func newSnooper(ofClient openflow.Client, ifaceStore interfacestore.InterfaceStore, eventCh chan *mcastGroupEvent, queryInterval time.Duration, igmpQueryVersions []uint8, multicastValidator types.McastNetworkPolicyController, encapEnabled bool) *IGMPSnooper {
	_ = "STUB: not implemented"
	return nil
}
