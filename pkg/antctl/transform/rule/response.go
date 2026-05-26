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

package rule

import (
	cpv1beta "antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
)

type service struct {
	Protocol string `json:"protocol,omitempty"`
	Port     string `json:"port,omitempty"`
	EndPort  string `json:"endPort,omitempty"`
}

type ipBlock struct {
	CIDR   string   `json:"cidr" yaml:"cidr"`
	Except []string `json:"except,omitempty"`
}

type peer struct {
	AddressGroups []string  `json:"addressGroups,omitempty"`
	IPBlocks      []ipBlock `json:"ipBlocks,omitempty"`
}

type Response struct {
	Direction string    `json:"direction,omitempty"`
	From      peer      `json:"from,omitempty"`
	To        peer      `json:"to,omitempty"`
	Services  []service `json:"services,omitempty"`
}

func serviceTransform(services ...cpv1beta.Service) []service {
	_ = "STUB: not implemented"
	return nil
}

func ipBlockTransform(block cpv1beta.IPBlock) ipBlock {
	_ = "STUB: not implemented"
	return *new(ipBlock)
}

func peerTransform(p cpv1beta.NetworkPolicyPeer) peer { _ = "STUB: not implemented"; return *new(peer) }

func ObjectTransform(o interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
