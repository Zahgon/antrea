//go:build !windows

// Copyright 2024 Antrea Authors
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

package monitortool

import (
	"net"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"

	statsv1alpha1 "antrea.io/antrea/v2/pkg/apis/stats/v1alpha1"
)

// LatencyStore is a store for latency information of connections between Nodes.
type LatencyStore struct {
	// Lock for the latency store
	mutex sync.RWMutex

	// Whether the agent is running in networkPolicyOnly mode
	isNetworkPolicyOnly bool
	// The map of Node IP to latency entry, it will be changed by latency monitor
	nodeIPLatencyMap map[string]*NodeIPLatencyEntry
	// The map of Node name to Node IP(s), it will be changed by Node watcher
	// If the agent is running in networkPolicyOnly mode, the value will be the transport IP of the Node.
	// Otherwise, the value will be the gateway IP of the Node
	nodeTargetIPsMap map[string][]net.IP
}

// NodeIPLatencyEntry is the entry of the latency map.
type NodeIPLatencyEntry struct {
	// The timestamp of the last sent packet
	LastSendTime time.Time
	// The timestamp of the last received packet
	LastRecvTime time.Time
	// The last valid rtt of the connection
	LastMeasuredRTT time.Duration
}

// NewLatencyStore creates a new LatencyStore.
func NewLatencyStore(isNetworkPolicyOnly bool) *LatencyStore { _ = "STUB: not implemented"; return nil }

// getNodeIPLatencyEntry returns the NodeIPLatencyEntry for the given Node IP
// For now, it is only used for testing purposes.
func (s *LatencyStore) getNodeIPLatencyEntry(nodeIP string) (NodeIPLatencyEntry, bool) {
	_ = "STUB: not implemented"
	return *new(NodeIPLatencyEntry), false
}

// getNodeIPLatencyKeys returns the list of Node IPs for which we currently have
// latency measurements.
// It is only used for testing purposes.
func (s *LatencyStore) getNodeIPLatencyKeys() []string { _ = "STUB: not implemented"; return nil }

// SetNodeIPLatencyEntry sets the NodeIPLatencyEntry for the given Node IP
func (s *LatencyStore) SetNodeIPLatencyEntry(nodeIP string, mutator func(entry *NodeIPLatencyEntry)) {
	_ = "STUB: not implemented"
	return
}

// addNode adds a Node to the latency store
func (s *LatencyStore) addNode(node *corev1.Node) { _ = "STUB: not implemented"; return }

// deleteNode deletes a Node from the latency store
func (s *LatencyStore) deleteNode(node *corev1.Node) { _ = "STUB: not implemented"; return }

// updateNode updates a Node name in the latency store
func (s *LatencyStore) updateNode(new *corev1.Node) { _ = "STUB: not implemented"; return }

// Node name will not be changed in the same Node update operation.

// updateNodeMap updates the nodeTargetIPsMap with the IPs of the given Node.
func (s *LatencyStore) updateNodeMap(node *corev1.Node) { _ = "STUB: not implemented"; return }

// getNodeIPs returns the target IPs of the given Node based on the agent mode.
func (s *LatencyStore) getNodeIPs(node *corev1.Node) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getTransportIPs returns the transport IPs of the given Node.
func getTransportIPs(node *corev1.Node) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getGWIPs returns the gateway IPs of the given Node.
func getGWIPs(node *corev1.Node) ([]net.IP, error) { _ = "STUB: not implemented"; return nil, nil }

// Add first IP in CIDR to the map

// getPodCIDRsOnNode returns the PodCIDRs of the given Node.
func getPodCIDRsOnNode(node *corev1.Node) []string { _ = "STUB: not implemented"; return nil }

// ListNodeIPs returns the list of all Node IPs in the latency store.
func (s *LatencyStore) ListNodeIPs() []net.IP { _ = "STUB: not implemented"; return nil }

// Allocate a slice with a capacity equal to twice the size of the map,
// as we can have up to 2 IP addresses per Node in dual-stack case.

// DeleteStaleNodeIPs deletes the stale Node IPs from the nodeIPLatencyMap.
func (s *LatencyStore) DeleteStaleNodeIPs() { _ = "STUB: not implemented"; return }

// ConvertList converts the latency store to a list of PeerNodeLatencyStats.
func (l *LatencyStore) ConvertList(currentNodeName string) []statsv1alpha1.PeerNodeLatencyStats {
	_ = "STUB: not implemented"
	return nil
}

// PeerNodeLatencyStats should be a list of size N-1, where N is the number of Nodes in the cluster.
// TargetIPLatencyStats will be a list of size 1 (single-stack case) or 2 (dual-stack case).

// Even though the current Node should already be excluded from the map, we add an extra check as an additional guarantee.
