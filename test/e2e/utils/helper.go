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

package utils

import (
	v1 "k8s.io/api/core/v1"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

type AntreaPolicyProtocol string

const (
	ProtocolTCP  AntreaPolicyProtocol = "TCP"
	ProtocolUDP  AntreaPolicyProtocol = "UDP"
	ProtocolSCTP AntreaPolicyProtocol = "SCTP"
	ProtocolICMP AntreaPolicyProtocol = "ICMP"
	ProtocolIGMP AntreaPolicyProtocol = "IGMP"
)

func AntreaPolicyProtocolToK8sProtocol(antreaProtocol AntreaPolicyProtocol) (v1.Protocol, error) {
	_ = "STUB: not implemented"
	return *new(v1.Protocol), nil
}

func GenPortsOrProtocols(ruleBuilder BaseRuleBuilder) ([]crdv1beta1.NetworkPolicyPort, []crdv1beta1.NetworkPolicyProtocol) {
	_ = "STUB: not implemented"
	return nil, nil
}
