//go:build !windows
// +build !windows

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

package rules

import (
	"bytes"

	"antrea.io/antrea/v2/pkg/agent/util/iptables"
)

// InitRules initializes rules based on the underlying implementation
func InitRules(isIPv6 bool) PodPortRules {
	_ = "STUB: not implemented"
	// This can be extended based on the system capability.
	return *new(PodPortRules)
}

// NodePortLocalChain is the name of the chain in IPTABLES for Node Port Local
const NodePortLocalChain = "ANTREA-NODE-PORT-LOCAL"

// IPTableRules provides a client to perform IPTABLES operations
type iptablesRules struct {
	table    iptables.Interface
	isIPv6   bool
	protocol iptables.Protocol
}

// NewIPTableRules retruns a new instance of IPTableRules
func NewIPTableRules(isIPv6 bool) *iptablesRules { _ = "STUB: not implemented"; return nil }

// Init initializes IPTABLES rules for NPL. Currently it deletes existing rules to ensure that no stale entries are present.
func (ipt *iptablesRules) Init() error { _ = "STUB: not implemented"; return nil }

// initRules creates the NPL chain and links it to the PREROUTING (for incoming
// traffic) and OUTPUT chain (for locally-generated traffic). All NPL DNAT rules
// will be added to this chain.
func (ipt *iptablesRules) initRules() error { _ = "STUB: not implemented"; return nil }

func buildRuleForPod(nodePort int, podIP string, podPort int, protocol string) []string {
	_ = "STUB: not implemented"
	return nil
}

// AddRule appends a DNAT rule in NodePortLocalChain chain of NAT table.
func (ipt *iptablesRules) AddRule(nodePort int, podIP string, podPort int, protocol string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddAllRules constructs a list of iptables rules for the NPL chain and performs a
// iptables-restore on this chain. It uses --no-flush to keep the previous rules intact.
func (ipt *iptablesRules) AddAllRules(nplList []PodNodePort) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteRule deletes a specific NPL rule from NodePortLocalChain chain
func (ipt *iptablesRules) DeleteRule(nodePort int, podIP string, podPort int, protocol string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAllRules deletes all NPL rules programmed in the node
func (ipt *iptablesRules) DeleteAllRules() error { _ = "STUB: not implemented"; return nil }

// Join all words with spaces, terminate with newline and write to buf.
func writeLine(buf *bytes.Buffer, words ...string) {
	_ = "STUB: not implemented"
	// We avoid strings.Join for performance reasons.
	return
}
