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

package utils

import (
	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
	"antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
)

const (
	connectionDyingFlag = uint32(1 << 9)
)

var (
	Protocols = map[string]uint8{
		"icmp":      1,
		"igmp":      2,
		"tcp":       6,
		"udp":       17,
		"ipv6-icmp": 58,
	}
)

// These constant values are valid for both Protobuf and IPFIX export (same values).

const (
	// FlowTypeUnspecified indicates that we are unable to determine the flow type.
	FlowTypeUnspecified  = uint8(flowpb.FlowType_FLOW_TYPE_UNSPECIFIED)
	FlowTypeInterNode    = uint8(flowpb.FlowType_FLOW_TYPE_INTER_NODE)
	FlowTypeIntraNode    = uint8(flowpb.FlowType_FLOW_TYPE_INTRA_NODE)
	FlowTypeToExternal   = uint8(flowpb.FlowType_FLOW_TYPE_TO_EXTERNAL)
	FlowTypeFromExternal = uint8(flowpb.FlowType_FLOW_TYPE_FROM_EXTERNAL)
	// FlowTypeUnsupported indicates that this type of flow is not supported and that we should
	// skip exporting it.
	FlowTypeUnsupported = uint8(0xff)
)

const (
	PolicyTypeUnspecified                = uint8(flowpb.NetworkPolicyType_NETWORK_POLICY_TYPE_UNSPECIFIED)
	PolicyTypeK8sNetworkPolicy           = uint8(flowpb.NetworkPolicyType_NETWORK_POLICY_TYPE_K8S)
	PolicyTypeAntreaNetworkPolicy        = uint8(flowpb.NetworkPolicyType_NETWORK_POLICY_TYPE_ANP)
	PolicyTypeAntreaClusterNetworkPolicy = uint8(flowpb.NetworkPolicyType_NETWORK_POLICY_TYPE_ACNP)
)

const (
	NetworkPolicyRuleActionNoAction = uint8(flowpb.NetworkPolicyRuleAction_NETWORK_POLICY_RULE_ACTION_NO_ACTION)
	NetworkPolicyRuleActionAllow    = uint8(flowpb.NetworkPolicyRuleAction_NETWORK_POLICY_RULE_ACTION_ALLOW)
	NetworkPolicyRuleActionDrop     = uint8(flowpb.NetworkPolicyRuleAction_NETWORK_POLICY_RULE_ACTION_DROP)
	NetworkPolicyRuleActionReject   = uint8(flowpb.NetworkPolicyRuleAction_NETWORK_POLICY_RULE_ACTION_REJECT)
)

func IsConnectionDying(conn *connection.Connection) bool {
	_ = "STUB: not implemented"
	// "TIME_WAIT" state indicates local endpoint has closed the connection.
	// "CLOSE" state indicates closing RST flag is set and connection is closed.
	return false
}

// connections in other protocol with dying bit set

// Connection no longer exists in conntrack table.

// checkConntrackConnActive returns true if there are changes in connection's stats or
// TCP state, indicating that the connection is active.
func CheckConntrackConnActive(conn *connection.Connection) bool {
	_ = "STUB: not implemented"
	return false
}

// RuleActionToUint8 converts network policy rule action to uint8.
func RuleActionToUint8(action string) uint8 { _ = "STUB: not implemented"; return 0 }

// policyTypeToUint8 converts NetworkPolicy type to uint8
func PolicyTypeToUint8(policyType v1beta2.NetworkPolicyType) uint8 {
	_ = "STUB: not implemented"
	return 0
}

// LookupProtocolMap returns protocol identifier given protocol name
func LookupProtocolMap(name string) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }
