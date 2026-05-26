//go:build windows
// +build windows

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

package rules

import (
	"antrea.io/antrea/v2/pkg/agent/util"
	"antrea.io/antrea/v2/pkg/agent/util/winnet"
)

// Use antrea-nat netnatstaticmapping rules as NPL implementation
var (
	antreaNatNPL = util.AntreaNatName
)

// InitRules initializes rules based on the netnatstaticmapping implementation on windows
func InitRules(_ bool) PodPortRules { _ = "STUB: not implemented"; return *new(PodPortRules) }

type netnatRules struct {
	name   string
	winnet winnet.Interface
}

// NewNetNatRules returns a new instance of netnatRules.
func NewNetNatRules() *netnatRules { _ = "STUB: not implemented"; return nil }

// Init initializes NetNat rules for NPL.
func (nn *netnatRules) Init() error { _ = "STUB: not implemented"; return nil }

// initRules creates or reuses NetNat table as NPL rule instance on Windows.
func (nn *netnatRules) initRules() error { _ = "STUB: not implemented"; return nil }

// AddRule appends a NetNatStaticMapping rule.
func (nn *netnatRules) AddRule(nodePort int, podIP string, podPort int, protocol string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddAllRules constructs a list of NPL rules and performs NetNatStaticMapping replacement.
func (nn *netnatRules) AddAllRules(nplList []PodNodePort) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteRule deletes a specific NPL rule from NetNatStaticMapping table
func (nn *netnatRules) DeleteRule(nodePort int, podIP string, podPort int, protocol string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAllRules deletes the NetNatStaticMapping table in the node
func (nn *netnatRules) DeleteAllRules() error { _ = "STUB: not implemented"; return nil }
