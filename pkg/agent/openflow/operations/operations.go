// Copyright 2024 Antrea Authors
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

package operations

import (
	"antrea.io/libOpenflow/openflow15"

	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

type ofAction int32

const (
	add ofAction = iota
	mod
	del
)

func (a ofAction) String() string { _ = "STUB: not implemented"; return "" }

type OFEntryOperations interface {
	AddAll(flows []*openflow15.FlowMod) error
	ModifyAll(flows []*openflow15.FlowMod) error
	BundleOps(adds, mods, dels []*openflow15.FlowMod) error
	DeleteAll(flows []*openflow15.FlowMod) error
	AddOFEntries(ofEntries []binding.OFEntry) error
	ModifyOFEntries(ofEntries []binding.OFEntry) error
	DeleteOFEntries(ofEntries []binding.OFEntry) error
}

type ofEntryOperations struct {
	bridge binding.Bridge
}

func NewOFEntryOperations(b binding.Bridge) OFEntryOperations {
	_ = "STUB: not implemented"
	return *new(OFEntryOperations)
}

func (c *ofEntryOperations) AddAll(flowMessages []*openflow15.FlowMod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ofEntryOperations) ModifyAll(flowMessages []*openflow15.FlowMod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ofEntryOperations) DeleteAll(flowMessages []*openflow15.FlowMod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ofEntryOperations) BundleOps(adds, mods, dels []*openflow15.FlowMod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ofEntryOperations) AddOFEntries(ofEntries []binding.OFEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ofEntryOperations) ModifyOFEntries(ofEntries []binding.OFEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ofEntryOperations) DeleteOFEntries(ofEntries []binding.OFEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ofEntryOperations) changeAll(flowsMap map[ofAction][]*openflow15.FlowMod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ofEntryOperations) changeOFEntries(ofEntries []binding.OFEntry, action ofAction) error {
	_ = "STUB: not implemented"
	return nil
}
