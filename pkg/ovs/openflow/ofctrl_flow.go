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
	"antrea.io/libOpenflow/openflow15"
	"antrea.io/ofnet/ofctrl"
)

type FlowStates struct {
	TableID         uint8
	PacketCount     uint64
	DurationNSecond uint32
}

type ofFlow struct {
	table *ofTable
	// The Flow.Table field can be updated by Reset(), which can be called by
	// ReplayFlows() when replaying the Flow to OVS. For thread safety, any access
	// to Flow.Table should hold the replayMutex read lock.
	*ofctrl.Flow

	// protocol adds a readable protocol type in the match string of ofFlow.
	protocol Protocol
	// ctStates is a temporary variable to maintain openflow15.CTStates. When FlowBuilder.Done is called, it is used to
	// set the CtStates field in ofctrl.Flow.Match.
	ctStates *openflow15.CTStates
}

func (f *ofFlow) String() string { _ = "STUB: not implemented"; return "" }

func (f *ofFlow) getFlowMod() (*openflow15.FlowMod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset updates the ofFlow.Flow.Table field with ofFlow.table.Table.
// In the case of reconnecting to OVS, the ofnet library creates new OFTable
// objects. Reset() can be called to reset ofFlow.Flow.Table to the right value,
// before replaying the Flow to OVS.
func (f *ofFlow) Reset() { _ = "STUB: not implemented"; return }

func (f *ofFlow) Add() error { _ = "STUB: not implemented"; return nil }

func (f *ofFlow) Modify() error { _ = "STUB: not implemented"; return nil }

func (f *ofFlow) Delete() error { _ = "STUB: not implemented"; return nil }

func (f *ofFlow) Type() EntryType { _ = "STUB: not implemented"; return *new(EntryType) }

func (f *ofFlow) MatchString() string { _ = "STUB: not implemented"; return "" }

func (f *ofFlow) FlowPriority() uint16 { _ = "STUB: not implemented"; return 0 }

func (f *ofFlow) FlowProtocol() Protocol { _ = "STUB: not implemented"; return *new(Protocol) }

func (f *ofFlow) GetBundleMessages(entryOper OFOperation) ([]ofctrl.OpenFlowModMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CopyToBuilder returns a new FlowBuilder that copies the table, protocols,
// matches, and CookieID of the Flow, but does not copy private status fields
// of the ofctrl.Flow, e.g. "realized" and "isInstalled". It copies the
// original actions of the Flow only if copyActions is set to true, and
// resets the priority in the new FlowBuilder if it is provided.
func (f *ofFlow) CopyToBuilder(priority uint16, copyActions bool) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

func (r *Range) ToNXRange() *openflow15.NXRange { _ = "STUB: not implemented"; return nil }

func (r *Range) Length() uint32 { _ = "STUB: not implemented"; return 0 }

func (r *Range) Offset() uint32 { _ = "STUB: not implemented"; return 0 }
