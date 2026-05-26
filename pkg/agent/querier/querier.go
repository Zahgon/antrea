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

package querier

import (
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/memberlist"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/agent/proxy"
	"antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
	"antrea.io/antrea/v2/pkg/querier"
)

var _ AgentQuerier = new(agentQuerier)

type AgentQuerier interface {
	GetNodeConfig() *config.NodeConfig
	GetNetworkConfig() *config.NetworkConfig
	GetInterfaceStore() interfacestore.InterfaceStore
	GetK8sClient() clientset.Interface
	GetAgentInfo(agentInfo *v1beta1.AntreaAgentInfo, partial bool)
	GetOpenflowClient() openflow.Client
	GetOVSCtlClient() ovsctl.OVSCtlClient
	GetProxier() proxy.ProxyQuerier
	GetNetworkPolicyInfoQuerier() querier.AgentNetworkPolicyInfoQuerier
	GetMemberlistCluster() memberlist.Interface
	GetNodeLister() corelisters.NodeLister
	GetBGPPolicyInfoQuerier() querier.AgentBGPPolicyInfoQuerier
}

type agentQuerier struct {
	nodeConfig               *config.NodeConfig
	networkConfig            *config.NetworkConfig
	interfaceStore           interfacestore.InterfaceStore
	k8sClient                clientset.Interface
	ofClient                 openflow.Client
	ovsBridgeClient          ovsconfig.OVSBridgeClient
	proxier                  proxy.ProxyQuerier
	networkPolicyInfoQuerier querier.AgentNetworkPolicyInfoQuerier
	apiPort                  int
	nplRange                 string
	memberlistCluster        memberlist.Interface
	nodeLister               corelisters.NodeLister
	bgpPolicyInfoQuerier     querier.AgentBGPPolicyInfoQuerier
}

func NewAgentQuerier(
	nodeConfig *config.NodeConfig,
	networkConfig *config.NetworkConfig,
	interfaceStore interfacestore.InterfaceStore,
	k8sClient clientset.Interface,
	ofClient openflow.Client,
	ovsBridgeClient ovsconfig.OVSBridgeClient,
	proxier proxy.ProxyQuerier,
	networkPolicyInfoQuerier querier.AgentNetworkPolicyInfoQuerier,
	apiPort int,
	nplRange string,
	memberlistCluster memberlist.Interface,
	nodeLister corelisters.NodeLister,
	bgpPolicyInfoQuerier querier.AgentBGPPolicyInfoQuerier,
) *agentQuerier {
	_ = "STUB: not implemented"
	return nil
}

// GetNodeLister returns NodeLister.
func (aq agentQuerier) GetNodeLister() corelisters.NodeLister {
	_ = "STUB: not implemented"
	return *

	// GetMemberlistCluster returns MemberlistCluster Interface.
	new(corelisters.NodeLister)
}

func (aq agentQuerier) GetMemberlistCluster() memberlist.Interface {
	_ = "STUB: not implemented"
	return *new(memberlist.Interface)
}

// GetNodeConfig returns NodeConfig.
func (aq agentQuerier) GetNodeConfig() *config.NodeConfig { _ = "STUB: not implemented"; return nil }

// GetNetworkConfig returns NetworkConfig.
func (aq agentQuerier) GetNetworkConfig() *config.NetworkConfig {
	_ = "STUB: not implemented"
	return nil

	// GetInterfaceStore returns InterfaceStore.
}

func (aq agentQuerier) GetInterfaceStore() interfacestore.InterfaceStore {
	_ = "STUB: not implemented"
	return *

	// GetK8sClient returns Kubernetes client.
	new(interfacestore.InterfaceStore)
}

func (aq agentQuerier) GetK8sClient() clientset.Interface {
	_ = "STUB: not implemented"
	return *

	// GetOpenflowClient returns openflow.Client.
	new(clientset.Interface)
}

func (aq *agentQuerier) GetOpenflowClient() openflow.Client {
	_ = "STUB: not implemented"
	return *

	// GetOVSCtlClient returns a new OVSCtlClient.
	new(openflow.Client)
}

func (aq *agentQuerier) GetOVSCtlClient() ovsctl.OVSCtlClient {
	_ = "STUB: not implemented"
	return *new(ovsctl.OVSCtlClient)
}

// GetProxier returns proxy.ProxyQuerier.
func (aq *agentQuerier) GetProxier() proxy.ProxyQuerier {
	_ = "STUB: not implemented"

	// GetNetworkPolicyInfoQuerier returns AgentNetworkPolicyInfoQuerier.
	return *new(proxy.ProxyQuerier)
}

func (aq agentQuerier) GetNetworkPolicyInfoQuerier() querier.AgentNetworkPolicyInfoQuerier {
	_ = "STUB: not implemented"
	return *new(querier.AgentNetworkPolicyInfoQuerier)
}

// getOVSVersion gets current OVS version.
func (aq agentQuerier) getOVSVersion() string { _ = "STUB: not implemented"; return "" }

// getOVSFlowTable gets current OVS flow tables.
func (aq agentQuerier) getOVSFlowTable() map[string]int32 { _ = "STUB: not implemented"; return nil }

// getAgentConditions gets current conditions of agent pod.
func (aq agentQuerier) getAgentConditions(ovsConnected bool) []v1beta1.AgentCondition {
	_ = "STUB: not implemented"
	return nil
}

// getNetworkPolicyControllerInfo gets current network policy controller info
// including: number of network policies, address groups and applied to groups.
func (aq agentQuerier) getNetworkPolicyControllerInfo() v1beta1.NetworkPolicyControllerInfo {
	_ = "STUB: not implemented"
	return *new(v1beta1.NetworkPolicyControllerInfo)
}

// GetAgentInfo gets current agent pod info.
func (aq agentQuerier) GetAgentInfo(agentInfo *v1beta1.AntreaAgentInfo, partial bool) {
	_ = "STUB: not implemented"
	// LocalPodNum, FlowTable, NetworkPolicyControllerInfo, OVSVersion and AgentConditions can be changed, so reset these fields.
	// Only these fields are updated when partial is true.
	return
}

// OVS version query will fail and return empty string when OVSDB connection is down.
// Only change OVS version when the query gets a valid version.

// Some other fields are needed when partial is false.

// Make a new string slice instead of appending agentInfo.NodeSubnets directly to avoid duplicate CIDRs.

// getNetworkInfo gets network information including transport interface details and Pod MTU.
func (aq agentQuerier) getNetworkInfo() v1beta1.NetworkInfo {
	_ = "STUB: not implemented"
	return *new(v1beta1.NetworkInfo)
}

// Collect transport interface IPs with CIDR notation

// GetBGPPolicyInfoQuerier returns AgentBGPPolicyInfoQuerier.
func (aq agentQuerier) GetBGPPolicyInfoQuerier() querier.AgentBGPPolicyInfoQuerier {
	_ = "STUB: not implemented"
	return *new(querier.AgentBGPPolicyInfoQuerier)
}
