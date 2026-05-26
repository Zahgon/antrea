// Copyright 2022 Antrea Authors
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
	"antrea.io/libOpenflow/openflow15"
)

// TableNameCache is for testing.
var TableNameCache map[uint8]string

type fieldMetadata struct {
	name   string
	length uint8
}

func (m *fieldMetadata) getMatchNickname() string { _ = "STUB: not implemented"; return "" }

func (m *fieldMetadata) getActionNickname() string { _ = "STUB: not implemented"; return "" }

var oxxFieldMetadataMap = map[uint16]map[uint8]*fieldMetadata{
	openflow15.OXM_CLASS_NXM_0: {
		openflow15.NXM_OF_IN_PORT:   &fieldMetadata{"NXM_OF_IN_PORT", 2},
		openflow15.NXM_OF_ETH_DST:   &fieldMetadata{"NXM_OF_ETH_DST", 6},
		openflow15.NXM_OF_ETH_SRC:   &fieldMetadata{"NXM_OF_ETH_SRC", 6},
		openflow15.NXM_OF_ETH_TYPE:  &fieldMetadata{"NXM_OF_ETH_TYPE", 2},
		openflow15.NXM_OF_VLAN_TCI:  &fieldMetadata{"NXM_OF_VLAN_TCI", 2},
		openflow15.NXM_OF_IP_TOS:    &fieldMetadata{"NXM_OF_IP_TOS", 1},
		openflow15.NXM_OF_IP_PROTO:  &fieldMetadata{"NXM_OF_IP_PROTO", 1},
		openflow15.NXM_OF_IP_SRC:    &fieldMetadata{"NXM_OF_IP_SRC", 4},
		openflow15.NXM_OF_IP_DST:    &fieldMetadata{"NXM_OF_IP_DST", 4},
		openflow15.NXM_OF_TCP_SRC:   &fieldMetadata{"NXM_OF_TCP_SRC", 2},
		openflow15.NXM_OF_TCP_DST:   &fieldMetadata{"NXM_OF_TCP_DST", 2},
		openflow15.NXM_OF_UDP_SRC:   &fieldMetadata{"NXM_OF_UDP_SRC", 2},
		openflow15.NXM_OF_UDP_DST:   &fieldMetadata{"NXM_OF_UDP_DST", 2},
		openflow15.NXM_OF_ICMP_TYPE: &fieldMetadata{"NXM_OF_ICMP_TYPE", 1},
		openflow15.NXM_OF_ICMP_CODE: &fieldMetadata{"NXM_OF_ICMP_CODE", 1},
		openflow15.NXM_OF_ARP_OP:    &fieldMetadata{"NXM_OF_ARP_OP", 2},
		openflow15.NXM_OF_ARP_SPA:   &fieldMetadata{"NXM_OF_ARP_SPA", 4},
		openflow15.NXM_OF_ARP_TPA:   &fieldMetadata{"NXM_OF_ARP_TPA", 4},
	},
	openflow15.OXM_CLASS_NXM_1: {
		openflow15.NXM_NX_REG0:          &fieldMetadata{"NXM_NX_REG0", 4},
		openflow15.NXM_NX_REG1:          &fieldMetadata{"NXM_NX_REG1", 4},
		openflow15.NXM_NX_REG2:          &fieldMetadata{"NXM_NX_REG2", 4},
		openflow15.NXM_NX_REG3:          &fieldMetadata{"NXM_NX_REG3", 4},
		openflow15.NXM_NX_REG4:          &fieldMetadata{"NXM_NX_REG4", 4},
		openflow15.NXM_NX_REG5:          &fieldMetadata{"NXM_NX_REG5", 4},
		openflow15.NXM_NX_REG6:          &fieldMetadata{"NXM_NX_REG6", 4},
		openflow15.NXM_NX_REG7:          &fieldMetadata{"NXM_NX_REG7", 4},
		openflow15.NXM_NX_REG8:          &fieldMetadata{"NXM_NX_REG8", 4},
		openflow15.NXM_NX_REG9:          &fieldMetadata{"NXM_NX_REG9", 4},
		openflow15.NXM_NX_REG10:         &fieldMetadata{"NXM_NX_REG10", 4},
		openflow15.NXM_NX_REG11:         &fieldMetadata{"NXM_NX_REG11", 4},
		openflow15.NXM_NX_REG12:         &fieldMetadata{"NXM_NX_REG12", 4},
		openflow15.NXM_NX_REG13:         &fieldMetadata{"NXM_NX_REG13", 4},
		openflow15.NXM_NX_REG14:         &fieldMetadata{"NXM_NX_REG14", 4},
		openflow15.NXM_NX_REG15:         &fieldMetadata{"NXM_NX_REG15", 4},
		openflow15.NXM_NX_TUN_ID:        &fieldMetadata{"NXM_NX_TUN_ID", 8},
		openflow15.NXM_NX_ARP_SHA:       &fieldMetadata{"NXM_NX_ARP_SHA", 6},
		openflow15.NXM_NX_ARP_THA:       &fieldMetadata{"NXM_NX_ARP_THA", 6},
		openflow15.NXM_NX_IPV6_SRC:      &fieldMetadata{"NXM_NX_IPV6_SRC", 16},
		openflow15.NXM_NX_IPV6_DST:      &fieldMetadata{"NXM_NX_IPV6_DST", 16},
		openflow15.NXM_NX_ICMPV6_TYPE:   &fieldMetadata{"NXM_NX_ICMPV6_TYPE", 1},
		openflow15.NXM_NX_ICMPV6_CODE:   &fieldMetadata{"NXM_NX_ICMPV6_CODE", 1},
		openflow15.NXM_NX_ND_TARGET:     &fieldMetadata{"NXM_NX_ND_TARGET", 16},
		openflow15.NXM_NX_ND_SLL:        &fieldMetadata{"NXM_NX_ND_SLL", 6},
		openflow15.NXM_NX_ND_TLL:        &fieldMetadata{"NXM_NX_ND_TLL", 6},
		openflow15.NXM_NX_IP_FRAG:       &fieldMetadata{"NXM_NX_IP_FRAG", 1},
		openflow15.NXM_NX_IPV6_LABEL:    &fieldMetadata{"NXM_NX_IPV6_LABEL", 1},
		openflow15.NXM_NX_IP_ECN:        &fieldMetadata{"NXM_NX_IP_ECN", 1},
		openflow15.NXM_NX_IP_TTL:        &fieldMetadata{"NXM_NX_IP_TTL", 1},
		openflow15.NXM_NX_MPLS_TTL:      &fieldMetadata{"NXM_NX_MPLS_TTL", 1},
		openflow15.NXM_NX_TUN_IPV4_SRC:  &fieldMetadata{"NXM_NX_TUN_IPV4_SRC", 4},
		openflow15.NXM_NX_TUN_IPV4_DST:  &fieldMetadata{"NXM_NX_TUN_IPV4_DST", 4},
		openflow15.NXM_NX_PKT_MARK:      &fieldMetadata{"NXM_NX_PKT_MARK", 4},
		openflow15.NXM_NX_TCP_FLAGS:     &fieldMetadata{"NXM_NX_TCP_FLAGS", 2},
		openflow15.NXM_NX_CONJ_ID:       &fieldMetadata{"NXM_NX_CONJ_ID", 4},
		openflow15.NXM_NX_TUN_GBP_ID:    &fieldMetadata{"NXM_NX_TUN_GBP_ID", 2},
		openflow15.NXM_NX_TUN_GBP_FLAGS: &fieldMetadata{"NXM_NX_TUN_GBP_FLAGS", 1},
		openflow15.NXM_NX_TUN_FLAGS:     &fieldMetadata{"NXM_NX_TUN_FLAGS", 2},
		openflow15.NXM_NX_CT_STATE:      &fieldMetadata{"NXM_NX_CT_STATE", 4},
		openflow15.NXM_NX_CT_ZONE:       &fieldMetadata{"NXM_NX_CT_ZONE", 2},
		openflow15.NXM_NX_CT_MARK:       &fieldMetadata{"NXM_NX_CT_MARK", 4},
		openflow15.NXM_NX_CT_LABEL:      &fieldMetadata{"NXM_NX_CT_LABEL", 16},
		openflow15.NXM_NX_TUN_IPV6_SRC:  &fieldMetadata{"NXM_NX_TUN_IPV6_SRC", 16},
		openflow15.NXM_NX_TUN_IPV6_DST:  &fieldMetadata{"NXM_NX_TUN_IPV6_DST", 16},
		openflow15.NXM_NX_CT_NW_PROTO:   &fieldMetadata{"NXM_NX_CT_NW_PROTO", 1},
		openflow15.NXM_NX_CT_NW_SRC:     &fieldMetadata{"NXM_NX_CT_NW_SRC", 4},
		openflow15.NXM_NX_CT_NW_DST:     &fieldMetadata{"NXM_NX_CT_NW_DST", 4},
		openflow15.NXM_NX_CT_IPV6_SRC:   &fieldMetadata{"NXM_NX_CT_IPV6_SRC", 16},
		openflow15.NXM_NX_CT_IPV6_DST:   &fieldMetadata{"NXM_NX_CT_IPV6_DST", 16},
		openflow15.NXM_NX_CT_TP_SRC:     &fieldMetadata{"NXM_NX_CT_TP_SRC", 2},
		openflow15.NXM_NX_CT_TP_DST:     &fieldMetadata{"NXM_NX_CT_TP_DST", 2},
		openflow15.NXM_NX_TUN_METADATA0: &fieldMetadata{"NXM_NX_TUN_METADATA0", 128},
		openflow15.NXM_NX_TUN_METADATA1: &fieldMetadata{"NXM_NX_TUN_METADATA1", 128},
		openflow15.NXM_NX_TUN_METADATA2: &fieldMetadata{"NXM_NX_TUN_METADATA2", 128},
		openflow15.NXM_NX_TUN_METADATA3: &fieldMetadata{"NXM_NX_TUN_METADATA3", 128},
		openflow15.NXM_NX_TUN_METADATA4: &fieldMetadata{"NXM_NX_TUN_METADATA4", 128},
		openflow15.NXM_NX_TUN_METADATA5: &fieldMetadata{"NXM_NX_TUN_METADATA5", 128},
		openflow15.NXM_NX_TUN_METADATA6: &fieldMetadata{"NXM_NX_TUN_METADATA6", 128},
		openflow15.NXM_NX_TUN_METADATA7: &fieldMetadata{"NXM_NX_TUN_METADATA7", 128},
		openflow15.NXM_NX_XXREG0:        &fieldMetadata{"NXM_NX_XXREG0", 16},
		openflow15.NXM_NX_XXREG1:        &fieldMetadata{"NXM_NX_XXREG1", 16},
		openflow15.NXM_NX_XXREG2:        &fieldMetadata{"NXM_NX_XXREG2", 16},
		openflow15.NXM_NX_XXREG3:        &fieldMetadata{"NXM_NX_XXREG3", 16},
	},
	openflow15.OXM_CLASS_OPENFLOW_BASIC: {
		openflow15.OXM_FIELD_IN_PORT:        &fieldMetadata{"OXM_OF_IN_PORT", 4},
		openflow15.OXM_FIELD_IN_PHY_PORT:    &fieldMetadata{"OXM_OF_IN_PHY_PORT", 4},
		openflow15.OXM_FIELD_METADATA:       &fieldMetadata{"OXM_OF_METADATA", 8},
		openflow15.OXM_FIELD_ETH_DST:        &fieldMetadata{"OXM_OF_ETH_DST", 6},
		openflow15.OXM_FIELD_ETH_SRC:        &fieldMetadata{"OXM_OF_ETH_SRC", 6},
		openflow15.OXM_FIELD_ETH_TYPE:       &fieldMetadata{"OXM_OF_ETH_TYPE", 2},
		openflow15.OXM_FIELD_VLAN_VID:       &fieldMetadata{"OXM_OF_VLAN_VID", 2},
		openflow15.OXM_FIELD_VLAN_PCP:       &fieldMetadata{"OXM_OF_VLAN_PCP", 1},
		openflow15.OXM_FIELD_IP_DSCP:        &fieldMetadata{"OXM_OF_IP_DSCP", 1},
		openflow15.OXM_FIELD_IP_ECN:         &fieldMetadata{"OXM_OF_IP_ECN", 1},
		openflow15.OXM_FIELD_IP_PROTO:       &fieldMetadata{"OXM_OF_IP_PROTO", 1},
		openflow15.OXM_FIELD_IPV4_SRC:       &fieldMetadata{"OXM_OF_IPV4_SRC", 4},
		openflow15.OXM_FIELD_IPV4_DST:       &fieldMetadata{"OXM_OF_IPV4_DST", 4},
		openflow15.OXM_FIELD_TCP_SRC:        &fieldMetadata{"OXM_OF_TCP_SRC", 2},
		openflow15.OXM_FIELD_TCP_DST:        &fieldMetadata{"OXM_OF_TCP_DST", 2},
		openflow15.OXM_FIELD_UDP_SRC:        &fieldMetadata{"OXM_OF_UDP_SRC", 2},
		openflow15.OXM_FIELD_UDP_DST:        &fieldMetadata{"OXM_OF_UDP_DST", 2},
		openflow15.OXM_FIELD_SCTP_SRC:       &fieldMetadata{"OXM_OF_SCTP_SRC", 2},
		openflow15.OXM_FIELD_SCTP_DST:       &fieldMetadata{"OXM_OF_SCTP_DST", 2},
		openflow15.OXM_FIELD_ICMPV4_TYPE:    &fieldMetadata{"OXM_OF_ICMPV4_TYPE", 1},
		openflow15.OXM_FIELD_ICMPV4_CODE:    &fieldMetadata{"OXM_OF_ICMPV4_CODE", 1},
		openflow15.OXM_FIELD_ARP_OP:         &fieldMetadata{"OXM_OF_ARP_OP", 2},
		openflow15.OXM_FIELD_ARP_SPA:        &fieldMetadata{"OXM_OF_ARP_SPA", 4},
		openflow15.OXM_FIELD_ARP_TPA:        &fieldMetadata{"OXM_OF_ARP_TPA", 4},
		openflow15.OXM_FIELD_ARP_SHA:        &fieldMetadata{"OXM_OF_ARP_SHA", 6},
		openflow15.OXM_FIELD_ARP_THA:        &fieldMetadata{"OXM_OF_ARP_THA", 6},
		openflow15.OXM_FIELD_IPV6_SRC:       &fieldMetadata{"OXM_OF_IPV6_SRC", 16},
		openflow15.OXM_FIELD_IPV6_DST:       &fieldMetadata{"OXM_OF_IPV6_DST", 16},
		openflow15.OXM_FIELD_IPV6_FLABEL:    &fieldMetadata{"OXM_OF_IPV6_FLABEL", 4},
		openflow15.OXM_FIELD_ICMPV6_TYPE:    &fieldMetadata{"OXM_OF_ICMPV6_TYPE", 1},
		openflow15.OXM_FIELD_ICMPV6_CODE:    &fieldMetadata{"OXM_OF_ICMPV6_CODE", 1},
		openflow15.OXM_FIELD_IPV6_ND_TARGET: &fieldMetadata{"OXM_OF_IPV6_ND_TARGET", 16},
		openflow15.OXM_FIELD_IPV6_ND_SLL:    &fieldMetadata{"OXM_OF_IPV6_ND_SLL", 6},
		openflow15.OXM_FIELD_IPV6_ND_TLL:    &fieldMetadata{"OXM_OF_IPV6_ND_TLL", 6},
		openflow15.OXM_FIELD_MPLS_LABEL:     &fieldMetadata{"OXM_OF_MPLS_LABEL", 4},
		openflow15.OXM_FIELD_MPLS_TC:        &fieldMetadata{"OXM_OF_MPLS_TC", 1},
		openflow15.OXM_FIELD_MPLS_BOS:       &fieldMetadata{"OXM_OF_MPLS_BOS", 1},
		openflow15.OXM_FIELD_PBB_ISID:       &fieldMetadata{"OXM_OF_PBB_ISID", 3},
		openflow15.OXM_FIELD_TUNNEL_ID:      &fieldMetadata{"OXM_OF_TUNNEL_ID", 8},
		openflow15.OXM_FIELD_IPV6_EXTHDR:    &fieldMetadata{"OXM_OF_IPV6_EXTHDR", 2},
		openflow15.OXM_FIELD_TCP_FLAGS:      &fieldMetadata{"OXM_FIELD_TCP_FLAGS", 2},
	},
}

func getFieldNameString(class uint16, field uint8, offset, length uint16, nickName bool, usedForMatching bool) string {
	_ = "STUB: not implemented"
	return ""
}

func matchPktMarkToString(field *openflow15.MatchField) string {
	_ = "STUB: not implemented"
	return ""
}

func matchConjIdToString(field *openflow15.MatchField) string { _ = "STUB: not implemented"; return "" }

func matchCtStateToString(field *openflow15.MatchField) string {
	_ = "STUB: not implemented"
	return ""
}

func matchCtZoneToString(field *openflow15.MatchField) string { _ = "STUB: not implemented"; return "" }

func matchCtMarkToString(field *openflow15.MatchField) string { _ = "STUB: not implemented"; return "" }

func matchCtLabelToString(field *openflow15.MatchField) string {
	_ = "STUB: not implemented"
	return ""
}

func matchCtNwProtoToString(field *openflow15.MatchField) string {
	_ = "STUB: not implemented"
	return ""
}

func matchProtoToString(etherType, ipProto *openflow15.MatchField) string {
	_ = "STUB: not implemented"
	return ""
}

func matchRegToString(idx int, field *openflow15.MatchField) string {
	_ = "STUB: not implemented"
	return ""
}

func matchXXRegToString(idx int, field *openflow15.MatchField) string {
	_ = "STUB: not implemented"
	return ""
}

func matchInPortToString(field *openflow15.MatchField) string { _ = "STUB: not implemented"; return "" }

func matchVlanToString(field *openflow15.MatchField) string { _ = "STUB: not implemented"; return "" }

func matchIpAddrToString(field *openflow15.MatchField, isCt, isSrc, isIPv6 bool) string {
	_ = "STUB: not implemented"
	return ""
}

func matchTunDstToString(field *openflow15.MatchField, isIPv6 bool) string {
	_ = "STUB: not implemented"
	return ""
}

func matchTunIDToString(field *openflow15.MatchField) string { _ = "STUB: not implemented"; return "" }

func matchNwTosToString(field *openflow15.MatchField) string { _ = "STUB: not implemented"; return "" }

func matchIpDscpToString(field *openflow15.MatchField) string { _ = "STUB: not implemented"; return "" }

func matchTpPortToString(field *openflow15.MatchField, isCt, isSrc bool) string {
	_ = "STUB: not implemented"
	return ""
}

func matchEtherAddrToString(field *openflow15.MatchField, isSrc bool) string {
	_ = "STUB: not implemented"
	return ""
}

func matchArpPaAddrToString(field *openflow15.MatchField, isSrc bool) string {
	_ = "STUB: not implemented"
	return ""
}

func matchArpHaAddrToString(field *openflow15.MatchField, isSrc bool) string {
	_ = "STUB: not implemented"
	return ""
}

func matchArpOpToString(field *openflow15.MatchField) string { _ = "STUB: not implemented"; return "" }

func matchIcmpTypeToString(field *openflow15.MatchField) string {
	_ = "STUB: not implemented"
	return ""
}

func matchIcmpCodeToString(field *openflow15.MatchField) string {
	_ = "STUB: not implemented"
	return ""
}

func trimLeadingZero(s string) string { _ = "STUB: not implemented"; return "" }

func getFieldDataString(field *openflow15.MatchField) string { _ = "STUB: not implemented"; return "" }

func actionOutputToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

func actionPopVlanToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

func actionPushToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

func actionCopyFieldToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

func actionSetFieldToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

func actionMeterToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

func nxActionOutputRegToString(action openflow15.Action) string {
	_ = "STUB: not implemented"
	return ""
}

func nxActionConnTrackToString(action openflow15.Action) string {
	_ = "STUB: not implemented"
	return ""
}

func nxActionCTNATToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

func nxActionResubmitTableToString(action openflow15.Action) string {
	_ = "STUB: not implemented"
	return ""
}

func nxActionDecTTLToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

func nxActionConjunctionToString(action openflow15.Action) string {
	_ = "STUB: not implemented"
	return ""
}

func nxActionGroupToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

func nxActionNoteToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

func nxActionControllerToString(action openflow15.Action) string {
	_ = "STUB: not implemented"
	return ""
}

func nxActionController2ToString(action openflow15.Action) string {
	_ = "STUB: not implemented"
	return ""
}

// Add padding

func nxActionLearnToString(action openflow15.Action) string { _ = "STUB: not implemented"; return "" }

//TODO: add isOutput

func getFlowModBaseString(flowMod *openflow15.FlowMod) string {
	_ = "STUB: not implemented"

	// cookie
	return ""
}

// table

// idle_timeout

// hard_timeout

func getFlowModMatch(flowMod *openflow15.FlowMod) string { _ = "STUB: not implemented"; return "" }

// TODO: add support for field "recirc_id"

// TODO: add support for field "dp_hash"

// TODO: add support for field "skb_priority"

// TODO: add support for field "actset_output"

// TODO: add other match conditions about tun

// TODO: add support for field "metadata"

// TODO: add support for field "nw_proto"

// TODO: add support for field "nw_ecn", "nw_ttl", other match conditions about MPLS, and "nw_frag"

func matchTCPFlagsToString(field *openflow15.MatchField) string {
	_ = "STUB: not implemented"
	return ""
}

func getActionString(action openflow15.Action) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getFlowModAction(flowMod *openflow15.FlowMod) string { _ = "STUB: not implemented"; return "" }

func FlowModToString(flowMod *openflow15.FlowMod) string { _ = "STUB: not implemented"; return "" }

func FlowModMatchString(flowMod *openflow15.FlowMod, omitFields ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// Omit specific fields if needed. For example, the priority match field is not supported
// for the ovs-ofctl dump-flows command, and should be removed.

func GroupModToString(groupMod *openflow15.GroupMod) string { _ = "STUB: not implemented"; return "" }
