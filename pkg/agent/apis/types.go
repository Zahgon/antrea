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

package apis

import (
	"time"

	corev1 "k8s.io/api/core/v1"

	"antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

// AntreaAgentInfoResponse is the struct for the response of agentinfo command.
// It includes all fields except meta info from v1beta1.AntreaAgentInfo struct.
type AntreaAgentInfoResponse struct {
	Version                     string                              `json:"version,omitempty"`                     // Antrea binary version
	PodRef                      corev1.ObjectReference              `json:"podRef,omitempty"`                      // The Pod that Antrea Agent is running in
	NodeRef                     corev1.ObjectReference              `json:"nodeRef,omitempty"`                     // The Node that Antrea Agent is running in
	NodeSubnets                 []string                            `json:"nodeSubnets,omitempty"`                 // Node subnets
	OVSInfo                     v1beta1.OVSInfo                     `json:"ovsInfo,omitempty"`                     // OVS Information
	NetworkPolicyControllerInfo v1beta1.NetworkPolicyControllerInfo `json:"networkPolicyControllerInfo,omitempty"` // Antrea Agent NetworkPolicy information
	LocalPodNum                 int32                               `json:"localPodNum,omitempty"`                 // The number of Pods which the agent is in charge of
	AgentConditions             []v1beta1.AgentCondition            `json:"agentConditions,omitempty"`             // Agent condition contains types like AgentHealthy
	// The type should have been int32 in the CRD as well, unfortunately it was added as an int.
	APIPort                int32               `json:"apiPort,omitempty"`                // The port of Antrea Agent API Server
	NodePortLocalPortRange string              `json:"nodePortLocalPortRange,omitempty"` // The port range used by NodePortLocal
	NetworkInfo            v1beta1.NetworkInfo `json:"networkInfo,omitempty"`            // Network information
}

func (r AntreaAgentInfoResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r AntreaAgentInfoResponse) getAgentConditionStr() string {
	_ = "STUB: not implemented"
	return ""
}

func (r AntreaAgentInfoResponse) GetTableRow(maxColumnLength int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (r AntreaAgentInfoResponse) SortRows() bool { _ = "STUB: not implemented"; return false }

type FQDNCacheResponse struct {
	FQDNName       string    `json:"fqdnName,omitempty"`
	IPAddress      string    `json:"ipAddress,omitempty"`
	ExpirationTime time.Time `json:"expirationTime,omitempty"`
}

func (r FQDNCacheResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r FQDNCacheResponse) GetTableRow(maxColumn int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (r FQDNCacheResponse) SortRows() bool { _ = "STUB: not implemented"; return false }

type FeatureGateResponse struct {
	Component string `json:"component,omitempty"`
	Name      string `json:"name,omitempty"`
	Status    string `json:"status,omitempty"`
	Version   string `json:"version,omitempty"`
}

// MemberlistResponse describes the response struct of memberlist command.
type MemberlistResponse struct {
	NodeName string `json:"nodeName,omitempty"`
	IP       string `json:"ip,omitempty"`
	Status   string `json:"status,omitempty"`
}

func (r MemberlistResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r MemberlistResponse) GetTableRow(_ int) []string { _ = "STUB: not implemented"; return nil }

func (r MemberlistResponse) SortRows() bool { _ = "STUB: not implemented"; return false }

type MulticastResponse struct {
	PodName      string `json:"name,omitempty" antctl:"name,Name of the Pod"`
	PodNamespace string `json:"podNamespace,omitempty"`
	Inbound      string `json:"inbound,omitempty"`
	Outbound     string `json:"outbound,omitempty"`
}

func (r MulticastResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r MulticastResponse) GetTableRow(_ int) []string { _ = "STUB: not implemented"; return nil }

func (r MulticastResponse) SortRows() bool {
	_ = "STUB: not implemented"

	// OVSFlowResponse is the response struct of ovsflows command.
	return false
}

type OVSFlowResponse struct {
	Flow string `json:"flow,omitempty"`
}

func (r OVSFlowResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r OVSFlowResponse) GetTableRow(maxColumnLength int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (r OVSFlowResponse) SortRows() bool {
	_ = "STUB: not implemented"

	// OVSTracingResponse is the response struct of ovstracing command.
	return false
}

type OVSTracingResponse struct {
	Result string `json:"result,omitempty"`
}

// PodInterfaceResponse describes the response struct of pod-interface command.
type PodInterfaceResponse struct {
	PodName       string   `json:"name,omitempty" antctl:"name,Name of the Pod"`
	PodNamespace  string   `json:"podNamespace,omitempty"`
	InterfaceName string   `json:"interfaceName,omitempty"`
	IPs           []string `json:"ips,omitempty"`
	MAC           string   `json:"mac,omitempty"`
	PortUUID      string   `json:"portUUID,omitempty"`
	OFPort        int32    `json:"ofPort,omitempty"`
	ContainerID   string   `json:"containerID,omitempty"`
}

func (r PodInterfaceResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r PodInterfaceResponse) getContainerIDStr() string { _ = "STUB: not implemented"; return "" }

func (r PodInterfaceResponse) GetTableRow(_ int) []string { _ = "STUB: not implemented"; return nil }

func (r PodInterfaceResponse) SortRows() bool {
	_ = "STUB: not implemented"

	// ServiceExternalIPInfo contains the essential information for Services with type of Loadbalancer managed by Antrea.
	return false
}

type ServiceExternalIPInfo struct {
	ServiceName    string `json:"serviceName,omitempty" antctl:"name,Name of the Service"`
	Namespace      string `json:"namespace,omitempty"`
	ExternalIP     string `json:"externalIP,omitempty"`
	ExternalIPPool string `json:"externalIPPool,omitempty"`
	AssignedNode   string `json:"assignedNode,omitempty"`
}

func (r ServiceExternalIPInfo) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r ServiceExternalIPInfo) GetTableRow(_ int) []string { _ = "STUB: not implemented"; return nil }

func (r ServiceExternalIPInfo) SortRows() bool {
	_ = "STUB: not implemented"

	// BGPPolicyResponse describes the response struct of bgppolicy command.
	return false
}

type BGPPolicyResponse struct {
	BGPPolicyName           string   `json:"name,omitempty"`
	RouterID                string   `json:"routerID,omitempty"`
	LocalASN                int32    `json:"localASN,omitempty"`
	ListenPort              int32    `json:"listenPort,omitempty"`
	ConfederationIdentifier int32    `json:"confederationIdentifier,omitempty"`
	MemberASNs              []uint32 `json:"memberASNs,omitempty"`
}

func (r BGPPolicyResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r BGPPolicyResponse) GetTableRow(maxColumnLength int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (r BGPPolicyResponse) SortRows() bool {
	_ = "STUB: not implemented"

	// BGPPeerResponse describes the response struct of bgppeers command.
	return false
}

type BGPPeerResponse struct {
	Peer  string `json:"peer,omitempty"`
	ASN   int32  `json:"asn,omitempty"`
	State string `json:"state,omitempty"`
}

func (r BGPPeerResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r BGPPeerResponse) GetTableRow(_ int) []string { _ = "STUB: not implemented"; return nil }

func (r BGPPeerResponse) SortRows() bool {
	_ = "STUB: not implemented"

	// BGPRouteResponse describes the response struct of bgproutes command.
	return false
}

type BGPRouteResponse struct {
	Route     string `json:"route,omitempty"`
	Type      string `json:"type,omitempty"`
	K8sObjRef string `json:"k8sObjRef,omitempty"`
}

func (r BGPRouteResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r BGPRouteResponse) GetTableRow(_ int) []string { _ = "STUB: not implemented"; return nil }

func (r BGPRouteResponse) SortRows() bool { _ = "STUB: not implemented"; return false }
