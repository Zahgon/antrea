// Copyright 2019 Antrea Authors
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

package ovs

import (
	"testing"
	"time"

	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
)

const (
	openFlowCheckTimeout  = 500 * time.Millisecond
	openFlowCheckInterval = 100 * time.Millisecond
)

func PrepareOVSBridge(brName string) error {
	_ = "STUB: not implemented"
	// using the netdev datapath type does not impact test coverage but
	// ensures that the integration tests can be run with Docker Desktop on
	// macOS.
	return nil
}

func DeleteOVSBridge(brName string) error { _ = "STUB: not implemented"; return nil }

type ExpectFlow struct {
	MatchStr string
	ActStr   string
}

func (f ExpectFlow) flowStr(name string) string { _ = "STUB: not implemented"; return "" }

func CheckFlowExists(t *testing.T, ovsCtlClient ovsctl.OVSCtlClient, tableName string, tableID uint8, expectFound bool, flows []*ExpectFlow) []string {
	_ = "STUB: not implemented"
	return nil
}

func CheckGroupExists(t *testing.T, ovsCtlClient ovsctl.OVSCtlClient, groupID binding.GroupIDType, groupType string, buckets []string, expectFound bool) {
	_ = "STUB: not implemented"
	return
}

func OfctlFlowMatch(flowList []string, tableName string, flow *ExpectFlow) bool {
	_ = "STUB: not implemented"
	return false
}

// trimOVSFlowLine matches historical pkg/ovs/ovsctl.trimFlowStr: find " table" and
// return the substring from the "t" of "table=..." onward so each dump line is shaped
// like the output ovsctl used to return. pkg/ovs/ovsctl no longer applies this trim, so
// the integration test helpers do it before formatFlowDump.
func trimOVSFlowLine(line string) string { _ = "STUB: not implemented"; return "" }

func formatFlowDump(rawFlows []string) []string { _ = "STUB: not implemented"; return nil }

func OfctlDumpFlows(ovsCtlClient ovsctl.OVSCtlClient, args ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OfctlDumpTableFlows(ovsCtlClient ovsctl.OVSCtlClient, table string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OfctlDumpTableFlowsWithoutName(ovsCtlClient ovsctl.OVSCtlClient, table uint8) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OfctlDeleteFlows(ovsCtlClient ovsctl.OVSCtlClient) error {
	_ = "STUB: not implemented"
	return nil
}

func OfCtlDumpGroups(ovsCtlClient ovsctl.OVSCtlClient) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OfctlDeleteGroups(ovsCtlClient ovsctl.OVSCtlClient) error {
	_ = "STUB: not implemented"
	return nil
}
