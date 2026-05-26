// Copyright 2022 Antrea Authors.
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

package openflow

import (
	"net"

	"antrea.io/libOpenflow/openflow15"
	"antrea.io/ofnet/ofctrl"
)

type ofFlowAction struct {
	builder *ofFlowBuilder
}

// Drop is an action to drop packets.
func (a *ofFlowAction) Drop() FlowBuilder { _ = "STUB: not implemented"; return *new(FlowBuilder) }

// Output is an action to output packets to the specified ofport.
func (a *ofFlowAction) Output(port uint32) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// OutputFieldRange is an action to output packets to the port located in the specified NXM field with rng.
func (a *ofFlowAction) OutputFieldRange(name string, rng *Range) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (a *ofFlowAction) OutputToRegField(field *RegField) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// OutputInPort is an action to output packets to the ofport from where the packet enters the OFSwitch.
func (a *ofFlowAction) OutputInPort() FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// CT is an action to set conntrack marks and return CTAction to add actions that is executed with conntrack context.
// zone will be ignored if zoneSrcField is not nil.
func (a *ofFlowAction) CT(commit bool, tableID uint8, zone int, zoneSrcField *RegField) CTAction {
	_ = "STUB: not implemented"
	return *new(CTAction)
}

// ofCTAction is a struct to implement CTAction.
type ofCTAction struct {
	ctBase
	actions []openflow15.Action
	builder *ofFlowBuilder
}

func (a *ofCTAction) LoadToCtMark(marks ...*CtMark) CTAction {
	_ = "STUB: not implemented"
	return *new(CTAction)
}

func (a *ofCTAction) LoadToLabelField(value uint64, labelField *CtLabel) CTAction {
	_ = "STUB: not implemented"
	return *new(CTAction)
}

// MoveToLabel is an action to move data into ct_label.
func (a *ofCTAction) MoveToLabel(fromName string, fromRng, labelRng *Range) CTAction {
	_ = "STUB: not implemented"
	return *new(CTAction)
}

// MoveToCtMarkField is an action to move data into ct_mark.
func (a *ofCTAction) MoveToCtMarkField(fromRegField *RegField, ctMarkField *CtMarkField) CTAction {
	_ = "STUB: not implemented"
	return *new(CTAction)
}

func (a *ofCTAction) move(fromField *openflow15.OxmId, toField *openflow15.OxmId, nBits, fromStart, toStart uint16) {
	_ = "STUB: not implemented"
	return
}

func (a *ofCTAction) natAction(isSNAT bool, ipRange *IPRange, portRange *PortRange) CTAction {
	_ = "STUB: not implemented"
	return *new(CTAction)
}

// ipRange should not be nil. The check here is for code safety.

func (a *ofCTAction) SNAT(ipRange *IPRange, portRange *PortRange) CTAction {
	_ = "STUB: not implemented"
	return *new(CTAction)
}

func (a *ofCTAction) DNAT(ipRange *IPRange, portRange *PortRange) CTAction {
	_ = "STUB: not implemented"
	return *new(CTAction)
}

func (a *ofCTAction) NAT() CTAction { _ = "STUB: not implemented"; return *new(CTAction) }

// CTDone sets the conntrack action in the Openflow rule and it returns FlowBuilder.
func (a *ofCTAction) CTDone() FlowBuilder { _ = "STUB: not implemented"; return *new(FlowBuilder) }

// SetDstMAC is an action to modify packet destination MAC address to the specified address.
func (a *ofFlowAction) SetDstMAC(addr net.HardwareAddr) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SetSrcMAC is an action to modify packet source MAC address to the specified address.
func (a *ofFlowAction) SetSrcMAC(addr net.HardwareAddr) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SetARPSha is an action to modify ARP packet source hardware address to the specified address.
func (a *ofFlowAction) SetARPSha(addr net.HardwareAddr) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SetARPTha is an action to modify ARP packet target hardware address to the specified address.
func (a *ofFlowAction) SetARPTha(addr net.HardwareAddr) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SetARPSpa is an action to modify ARP packet source protocol address to the specified address.
func (a *ofFlowAction) SetARPSpa(addr net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SetARPTpa is an action to modify ARP packet target protocol address to the specified address.
func (a *ofFlowAction) SetARPTpa(addr net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SetSrcIP is an action to modify packet source IP address to the specified address.
func (a *ofFlowAction) SetSrcIP(addr net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SetDstIP is an action to modify packet destination IP address to the specified address.
func (a *ofFlowAction) SetDstIP(addr net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SetTunnelDst is an action to modify packet tunnel destination address to the specified address.
func (a *ofFlowAction) SetTunnelDst(addr net.IP) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SetTunnelID is an action to modify packet tunnel ID to the specified ID.
func (a *ofFlowAction) SetTunnelID(tunnelID uint64) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// PopVLAN is an action to pop VLAN ID.
func (a *ofFlowAction) PopVLAN() FlowBuilder { _ = "STUB: not implemented"; return *new(FlowBuilder) }

// PushVLAN is an action to add VLAN ID.
func (a *ofFlowAction) PushVLAN(etherType uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SetVLAN is an action to set existing VLAN ID.
func (a *ofFlowAction) SetVLAN(vlanID uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// LoadARPOperation is an action to load data to NXM_OF_ARP_OP field.
func (a *ofFlowAction) LoadARPOperation(value uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (a *ofFlowAction) LoadToRegField(field *RegField, value uint32) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (a *ofFlowAction) LoadRegMark(marks ...*RegMark) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// LoadPktMarkRange is an action to load data into pkt_mark at specified range.
func (a *ofFlowAction) LoadPktMarkRange(value uint32, rng *Range) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// LoadIPDSCP is an action to load data to IP DSCP bits.
func (a *ofFlowAction) LoadIPDSCP(value uint8) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (a *ofFlowAction) setField(field *openflow15.MatchField) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// Move is an action to copy all data from "fromField" to "toField". Fields with name "fromField" and "fromField" should
// have the same data length, otherwise there will be error when realizing the flow on OFSwitch.
func (a *ofFlowAction) Move(fromField, toField string) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// MoveRange is an action to move data from "fromField" at "fromRange" to "toField" at "toRange".
func (a *ofFlowAction) MoveRange(fromField, toField string, fromRange, toRange Range) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (a *ofFlowAction) copyField(srcOxmId, dstOxmId *openflow15.OxmId, fromRange, toRange Range) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// Resubmit is an action to resubmit packet to the specified table with the port as new in_port. If port is empty string,
// the in_port field is not changed.
func (a *ofFlowAction) Resubmit(ofPort uint16, tableID uint8) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (a *ofFlowAction) ResubmitToTables(tables ...uint8) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// DecTTL is an action to decrease TTL. It is used in routing functions implemented by Openflow.
func (a *ofFlowAction) DecTTL() FlowBuilder { _ = "STUB: not implemented"; return *new(FlowBuilder) }

// Normal is an action to leverage OVS fwd table to forwarding packets.
func (a *ofFlowAction) Normal() FlowBuilder { _ = "STUB: not implemented"; return *new(FlowBuilder) }

// Conjunction is an action to add new conjunction configuration to conjunctive match flow.
func (a *ofFlowAction) Conjunction(conjID uint32, clauseID uint8, nClause uint8) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// Group is an action to forward packets to groups to do load-balance.
func (a *ofFlowAction) Group(id GroupIDType) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// Note annotates the OpenFlow entry. The notes are presented as hex digits in the OpenFlow entry, and it will be
// padded on the right to make the total number of bytes 6 more than a multiple of 8.
func (a *ofFlowAction) Note(notes string) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// SendToController will send the packet to the OVS controller.
// If pause option is true, the packet will be sent to the controller and meanwhile
// also paused in the pipeline. The controller could use a resume message to resume
// this packet letting it continue its journey in the pipeline from where it was
// paused.
// As for the userdata, the first 2 bytes are used for packetIn. The first byte is
// packetIn category, which indicates the handler of this packetIn. The second
// byte is packetIn operation, which indicates the operation(s) that should be
// executed by the handler.
func (a *ofFlowAction) SendToController(userdata []byte, pause bool) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (a *ofFlowAction) Meter(meterID uint32) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// Learn is an action which adds or modifies a flow in an OpenFlow table.
func (a *ofFlowAction) Learn(id uint8, priority uint16, idleTimeout, hardTimeout, finIdleTimeout, finHardTimeout uint16, cookieID uint64) LearnAction {
	_ = "STUB: not implemented"
	return *new(LearnAction)
}

// ofLearnAction is used to describe actions in the learned flow.
type ofLearnAction struct {
	flowBuilder *ofFlowBuilder
	nxLearn     *ofctrl.FlowLearn
}

// DeleteLearned makes learned flows to be deleted when current flow is being deleted.
func (a *ofLearnAction) DeleteLearned() LearnAction {
	_ = "STUB: not implemented"
	return *new(LearnAction)
}

// MatchEthernetProtocol specifies that the NXM_OF_ETH_TYPE field in the
// learned flow must match IP(0x800) or IPv6(0x86dd).
func (a *ofLearnAction) MatchEthernetProtocol(isIPv6 bool) LearnAction {
	_ = "STUB: not implemented"
	return *new(LearnAction)
}

func (a *ofLearnAction) MatchIPProtocol(protocol Protocol) LearnAction {
	_ = "STUB: not implemented"
	return *new(LearnAction)
}

// Return directly if the protocol is not supported.

// MatchLearnedDstPort specifies that the transport layer destination field
// {tcp|udp|sctp}_dst in the learned flow must match the same field of the packet
// currently being processed. It only accepts ProtocolTCP, ProtocolUDP, or
// ProtocolSCTP, and does nothing for other protocols.
func (a *ofLearnAction) MatchLearnedDstPort(protocol Protocol) LearnAction {
	_ = "STUB: not implemented"
	// OXM_OF fields support TCP, UDP and SCTP, but NXM_OF fields only support TCP and UDP. So here use "OXM_OF_" to
	// generate the field name.
	return *new(LearnAction)
}

// Return directly if the protocol is not supported.

// MatchLearnedSrcPort specifies that the transport layer source field
// {tcp|udp|sctp}_src in the learned flow must match the same field of the packet
// currently being processed. It only accepts ProtocolTCP, ProtocolUDP, or
// ProtocolSCTP, and does nothing for other protocols.
func (a *ofLearnAction) MatchLearnedSrcPort(protocol Protocol) LearnAction {
	_ = "STUB: not implemented"
	// OXM_OF fields support TCP, UDP and SCTP, but NXM_OF fields only support TCP and UDP. So here use "OXM_OF_" to
	// generate the field name.
	return *new(LearnAction)
}

// Return directly if the protocol is not supported.

// MatchLearnedSrcIP makes the learned flow match the nw_src of current IP packet.
func (a *ofLearnAction) MatchLearnedSrcIP(isIPv6 bool) LearnAction {
	_ = "STUB: not implemented"
	return *new(LearnAction)
}

// MatchLearnedDstIP makes the learned flow match the nw_dst of current IP packet.
func (a *ofLearnAction) MatchLearnedDstIP(isIPv6 bool) LearnAction {
	_ = "STUB: not implemented"
	return *new(LearnAction)
}

func (a *ofLearnAction) MatchRegMark(marks ...*RegMark) LearnAction {
	_ = "STUB: not implemented"
	return *new(LearnAction)
}

func (a *ofLearnAction) LoadFieldToField(fromField, toField *RegField) LearnAction {
	_ = "STUB: not implemented"
	return *new(LearnAction)
}

// LoadXXRegToXXReg makes the learned flow to load reg[fromXXField.regID] to reg[toXXField.regID]
// with specific ranges.
func (a *ofLearnAction) LoadXXRegToXXReg(fromXXField, toXXField *XXRegField) LearnAction {
	_ = "STUB: not implemented"
	return *new(LearnAction)
}

func (a *ofLearnAction) LoadRegMark(marks ...*RegMark) LearnAction {
	_ = "STUB: not implemented"
	return *new(LearnAction)
}

func (a *ofLearnAction) Done() FlowBuilder { _ = "STUB: not implemented"; return *new(FlowBuilder) }

func getFieldRange(name string) (*openflow15.MatchField, Range, error) {
	_ = "STUB: not implemented"
	return nil, *new(Range), nil
}

// GotoTable is an action to jump to the specified table.
func (a *ofFlowAction) GotoTable(tableID uint8) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (a *ofFlowAction) NextTable() FlowBuilder { _ = "STUB: not implemented"; return *new(FlowBuilder) }

func (a *ofFlowAction) GotoStage(stage StageID) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}
