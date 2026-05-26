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

package flowrecord

import (
	"time"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
)

type FlowRecord struct {
	FlowStartSeconds                     time.Time
	FlowEndSeconds                       time.Time
	FlowEndSecondsFromSourceNode         time.Time
	FlowEndSecondsFromDestinationNode    time.Time
	FlowEndReason                        uint8
	SourceIP                             string
	DestinationIP                        string
	SourceTransportPort                  uint16
	DestinationTransportPort             uint16
	ProtocolIdentifier                   uint8
	PacketTotalCount                     uint64
	OctetTotalCount                      uint64
	PacketDeltaCount                     uint64
	OctetDeltaCount                      uint64
	ReversePacketTotalCount              uint64
	ReverseOctetTotalCount               uint64
	ReversePacketDeltaCount              uint64
	ReverseOctetDeltaCount               uint64
	SourcePodName                        string
	SourcePodNamespace                   string
	SourceNodeName                       string
	DestinationPodName                   string
	DestinationPodNamespace              string
	DestinationNodeName                  string
	DestinationClusterIP                 string
	DestinationServicePort               uint16
	DestinationServicePortName           string
	IngressNetworkPolicyName             string
	IngressNetworkPolicyNamespace        string
	IngressNetworkPolicyRuleName         string
	IngressNetworkPolicyRuleAction       uint8
	IngressNetworkPolicyType             uint8
	EgressNetworkPolicyName              string
	EgressNetworkPolicyNamespace         string
	EgressNetworkPolicyRuleName          string
	EgressNetworkPolicyRuleAction        uint8
	EgressNetworkPolicyType              uint8
	TcpState                             string
	FlowType                             uint8
	SourcePodLabels                      string
	DestinationPodLabels                 string
	Throughput                           uint64
	ReverseThroughput                    uint64
	ThroughputFromSourceNode             uint64
	ThroughputFromDestinationNode        uint64
	ReverseThroughputFromSourceNode      uint64
	ReverseThroughputFromDestinationNode uint64
	EgressName                           string
	EgressIP                             string
	EgressNodeName                       string
}

// GetFlowRecord converts flowpb.Flow to FlowRecord.
// It assumes that record.Aggregation is set, so it should only be used in Aggregate mode.
func GetFlowRecord(record *flowpb.Flow) (*FlowRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// flow.K8S.SourcePodLabels.Labels can be nil or an empty map
// both cases should be treated the same

// handles the case where the protocol is not TCP
