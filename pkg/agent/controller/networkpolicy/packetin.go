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

package networkpolicy

import (
	"net/netip"

	"antrea.io/libOpenflow/openflow15"
	"antrea.io/ofnet/ofctrl"

	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

// HandlePacketIn is the packetIn handler registered to openflow by Antrea network
// policy agent controller. It performs the appropriate operations based on which
// bits are set in the "custom reasons" field of the packet received from OVS.
func (c *Controller) HandlePacketIn(pktIn *ofctrl.PacketIn) error {
	_ = "STUB: not implemented"
	return nil
}

// Choose operations.

// getMatchRegField returns match to the regNum register.
func getMatchRegField(matchers *ofctrl.Matchers, field *binding.RegField) *ofctrl.MatchField {
	_ = "STUB: not implemented"
	return nil
}

// getMatch receives ofctrl matchers and table id, match field.
// Modifies match field to Ingress/Egress register based on tableID.
func getMatch(matchers *ofctrl.Matchers, tableID uint8, disposition uint32) *ofctrl.MatchField {
	_ = "STUB: not implemented"
	// Get match from CNPDenyConjIDReg if disposition is Drop or Reject.
	return nil
}

// Get match from ingress/egress reg if disposition is Allow or Pass.

// getInfoInReg unloads and returns data stored in the match field.
func getInfoInReg(regMatch *ofctrl.MatchField, rng *openflow15.NXRange) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Controller) storeDenyConnection(pktIn *ofctrl.PacketIn) error {
	_ = "STUB: not implemented"
	return nil
}

// storeDenyConnectionParsed takes a parsed packet as input, making it easier to unit test than storeDenyConnection.
func (c *Controller) storeDenyConnectionParsed(pktIn *ofctrl.PacketIn, packet *binding.Packet) error {
	_ = "STUB: not implemented"
	return nil
}

// Get 5-tuple information

// Generate deny connection and add to deny connection store

// OriginalBytes will be added to the total size of the connection
// if the connection is already in the store.

// StartTime identifies when this packet was received. If the connection
// is already in the store, the start time will not be updated. This value
// is also used to approximate the stoptime of the connection.

// Get table ID

// Get disposition Allow, Drop or Reject

// Set match to corresponding ingress/egress reg according to disposition

// For K8s NetworkPolicy implicit drop action, we cannot get Namespace/name.

func isAntreaPolicyIngressTable(tableID uint8) bool { _ = "STUB: not implemented"; return false }

func isAntreaPolicyEgressTable(tableID uint8) bool { _ = "STUB: not implemented"; return false }

// getPacketInTableID returns the OVS table ID in which the packet is sent to antrea-agent. Since L2ForwardOutput is
// the table where all Antrea-native policies logging packets are sent to antrea-agent, "PacketInTableField" is used
// to store the real table requiring "sendToController" action. This function first parses the direct table where
// the packet leaves OVS pipeline, then checks whether "PacketInTableField" is set with a valid value or not. The value
// in the field is returned if yes.
func getPacketInTableID(pktIn *ofctrl.PacketIn) uint8 { _ = "STUB: not implemented"; return 0 }

// This is not expected, so we log an error.

func getCTMarkValue(matchers *ofctrl.Matchers) uint32 { _ = "STUB: not implemented"; return 0 }

// getCTLabelValue returns the conntrack label as a []byte using a big-endian representation.
func getCTLabelValue(matchers *ofctrl.Matchers) []byte { _ = "STUB: not implemented"; return nil }

func getCTNwDstValue(matchers *ofctrl.Matchers) netip.Addr {
	_ = "STUB: not implemented"
	return *new(netip.Addr)
}

func getCTTpDstValue(matchers *ofctrl.Matchers) uint16 { _ = "STUB: not implemented"; return 0 }
