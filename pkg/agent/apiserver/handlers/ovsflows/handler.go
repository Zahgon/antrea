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

package ovsflows

import (
	"fmt"
	"net/http"

	"antrea.io/antrea/v2/pkg/agent/apis"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	agentquerier "antrea.io/antrea/v2/pkg/agent/querier"
	cpv1beta "antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

var (
	// Use function variables for tests.
	getFlowTableName = openflow.GetFlowTableName
	getFlowTableID   = openflow.GetFlowTableID
	getFlowTableList = openflow.GetTableList

	errAmbiguousQuery = fmt.Errorf("query is ambiguous and matches more than one policy")
)

func dumpMatchedFlows(aq agentquerier.AgentQuerier, flowKeys []string) ([]apis.OVSFlowResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dumpFlows(aq agentquerier.AgentQuerier, table uint8) ([]apis.OVSFlowResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dumpMatchedGroups(aq agentquerier.AgentQuerier, groupIDs []binding.GroupIDType) ([]apis.OVSFlowResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nil is returned if the flow table can not be found (the passed table name or
// number is invalid).
func getTableFlows(aq agentquerier.AgentQuerier, tables string) ([]apis.OVSFlowResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Table nubmer is a 8-bit unsigned integer.

// nil is returned if the passed group IDs are invalid.
func getGroups(aq agentquerier.AgentQuerier, groups string) ([]apis.OVSFlowResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Group ID is a 32-bit unsigned integer.

func getPodFlows(aq agentquerier.AgentQuerier, podName, namespace string) ([]apis.OVSFlowResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getServiceFlows(aq agentquerier.AgentQuerier, serviceName, namespace string) ([]apis.OVSFlowResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getNetworkPolicyFlows(aq agentquerier.AgentQuerier, npName, namespace string, policyType cpv1beta.NetworkPolicyType) ([]apis.OVSFlowResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NetworkPolicy not found.

func getTableNames() []apis.OVSFlowResponse { _ = "STUB: not implemented"; return nil }

// HandleFunc returns the function which can handle API requests to "/ovsflows".
func HandleFunc(aq agentquerier.AgentQuerier) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Pod Namespace must be provided to dump flows of a Pod.

// policyType string has already been validated above
