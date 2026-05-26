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

package traceflow

import (
	"errors"

	"antrea.io/libOpenflow/openflow15"
	"antrea.io/ofnet/ofctrl"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

var errSkipTraceflowUpdate = errors.New("skip Traceflow update")

func (c *Controller) HandlePacketIn(pktIn *ofctrl.PacketIn) error {
	_ = "STUB: not implemented"
	return nil
}

// Retry when update CRD conflict which caused by multiple agents updating one CRD at same time.

func (c *Controller) parsePacketIn(pktIn *ofctrl.PacketIn) (*crdv1beta1.Traceflow, *crdv1beta1.NodeResult, *crdv1beta1.Packet, error) {
	_ = "STUB: not implemented"
	return nil, nil,

		// Get data plane tag.
		// Directly read data plane tag from packet.
		nil, nil
}

// Live Traceflow only considers the first packet of each
// connection. However, it is possible for 2 connections to
// match the Live Traceflow flows in OVS (before the flows can
// be uninstalled below), leading to 2 Packet In messages being
// processed. If we don't ignore all additional Packet Ins, we
// can end up with duplicate Node observations in the Traceflow
// Status. This situation is more likely when the Live TraceFlow
// request does not specify source / destination ports.

// Uninstall the OVS flows after receiving the first packet, to
// avoid capturing too many matched packets.

// Report the captured dropped packet, if the Traceflow is for
// the dropped packet only; report too if only the receiver
// captures packets in the Traceflow (live-traffic Traceflow
// that has only destination Pod set); otherwise only the sender
// should report the first captured packet.

// For SNATed packet(hairpin), ipSrc and ctNwSrc are different.
// We noticed that ctNwSrc is invalid for ICMPv6 packets: it should contain
// the original src Pod IP but it is always empty due to an issue in OVS.
// https://github.com/openvswitch/ovs-issues/issues/327

// In the case of ICMPv6, since ctNwSrc is invalid, we can use ipSrc as
// hairpin is not applicable, so ipSrc always contains src pod IP.

// Collect Service connections.
// - For packet is DNATed only, the final state is that ipDst != ctNwDst (in DNAT CT zone).
// - For packet is both DNATed and SNATed, the first state is also ipDst != ctNwDst (in DNAT CT zone), but the final
//   state is that ipSrc != ctNwSrc (in SNAT CT zone). The state in DNAT CT zone cannot be recognized in SNAT CT zone.

// Collect egress conjunctionID and get NetworkPolicy from cache.

// Collect ingress conjunctionID and get NetworkPolicy from cache.

// Get drop table.

// Get output table.

// decide according to packet.

// an Egress packet, currently on source Node and forwarded to Egress Node.

// encap or hybrid

// Egress packet on Egress Node

// Egress Node is Source Node of this Egress packet

// In hybrid mode or WireGuard mode, packets to Pod IPs in the same subnet are forwarded
// directly without encapsulation. Check if the destination is a Pod IP to determine
// the correct action (Forwarded vs ForwardedOutOfNetwork).

// networkPolicyOnly

// noEncap
// TODO: update this and above case if noEncap mode supports Egress feature

// Output port is Pod port, packet is delivered.

func getMatchPktMarkField(matchers *ofctrl.Matchers) *ofctrl.MatchField {
	_ = "STUB: not implemented"
	return nil
}

func getMatchRegField(matchers *ofctrl.Matchers, field *binding.RegField) *ofctrl.MatchField {
	_ = "STUB: not implemented"
	return nil
}

func getMatchTunnelDstField(matchers *ofctrl.Matchers, isIPv6 bool) *ofctrl.MatchField {
	_ = "STUB: not implemented"
	return nil
}

func getMarkValue(match *ofctrl.MatchField) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getRegValue(regMatch *ofctrl.MatchField, rng *openflow15.NXRange) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getTunnelDstValue(regMatch *ofctrl.MatchField) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getCTDstValue(matchers *ofctrl.Matchers, isIPv6 bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getCTSrcValue(matchers *ofctrl.Matchers, isIPv6 bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getNetworkPolicyObservation(tableID uint8, ingress bool) *crdv1beta1.Observation {
	_ = "STUB: not implemented"
	return nil
}

// Packet dropped by ANP/default drop rule

// Packet dropped by ANP/default drop rule

// Packet dropped by ANP/default drop rule

// Packet dropped by ANP/default drop rule

func isValidCtNw(ipStr string) bool { _ = "STUB: not implemented"; return false }

// Reserved by IETF [RFC3513][RFC4291]

func parseCapturedPacket(pktIn *ofctrl.PacketIn) *crdv1beta1.Packet {
	_ = "STUB: not implemented"
	return nil
}

func getEgressObservation(isEgressNode bool, egressIP, egressName, egressNode string) *crdv1beta1.Observation {
	_ = "STUB: not implemented"
	return nil
}
