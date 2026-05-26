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
	"math/rand/v2"
	"net"
	"sync/atomic"
	"time"

	corev1 "k8s.io/api/core/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"

	"antrea.io/antrea/v2/pkg/agent/client"
	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	statsv1alpha1 "antrea.io/antrea/v2/pkg/apis/stats/v1alpha1"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha1"
)

// #nosec G404: random number generator not used for security purposes.
var icmpEchoID = rand.Int32N(1 << 16)

const (
	ipv4ProtocolICMPRaw = "ip4:icmp"
	ipv6ProtocolICMPRaw = "ip6:ipv6-icmp"
	protocolICMP        = 1
	protocolICMPv6      = 58
	minReportInterval   = 10 * time.Second
	reportJitter        = time.Second
)

type PacketListener interface {
	ListenPacket(network, address string) (net.PacketConn, error)
}

type ICMPListener struct{}

func (l *ICMPListener) ListenPacket(network, address string) (net.PacketConn, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), nil
}

// NodeLatencyMonitor is a tool to monitor the latency of the Node.
type NodeLatencyMonitor struct {
	// latencyStore is the cache to store the latency of each Nodes.
	latencyStore *LatencyStore
	// latencyConfigChanged is the channel to notify the latency config changed.
	latencyConfigChanged chan latencyConfig
	// isIPv4Enabled is the flag to indicate whether the IPv4 is enabled.
	isIPv4Enabled bool
	// isIPv6Enabled is the flag to indicate whether the IPv6 is enabled.
	isIPv6Enabled bool

	// antreaClientProvider provides interfaces to get antreaClient, which will be used to report the statistics
	antreaClientProvider client.AntreaClientProvider
	// nodeName is the name of the current Node, used to filter out the current Node from the latency monitor.
	nodeName string

	nodeInformerSynced cache.InformerSynced
	nlmInformerSynced  cache.InformerSynced

	listener PacketListener

	icmpSeqNum atomic.Uint32
}

// latencyConfig is the config for the latency monitor.
type latencyConfig struct {
	// Enable is the flag to enable the latency monitor.
	Enable bool
	// Interval is the interval time to ping all Nodes.
	Interval time.Duration
}

// NewNodeLatencyMonitor creates a new NodeLatencyMonitor.
func NewNodeLatencyMonitor(
	antreaClientProvider client.AntreaClientProvider,
	nodeInformer coreinformers.NodeInformer,
	nlmInformer crdinformers.NodeLatencyMonitorInformer,
	nodeConfig *config.NodeConfig,
	trafficEncapMode config.TrafficEncapModeType,
) *NodeLatencyMonitor {
	_ = "STUB: not implemented"
	return nil
}

// Is current node
func (m *NodeLatencyMonitor) isCurrentNode(node *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

// onNodeAdd is the event handler for adding Node.
func (m *NodeLatencyMonitor) onNodeAdd(obj interface{}) { _ = "STUB: not implemented"; return }

// onNodeUpdate is the event handler for updating Node.
func (m *NodeLatencyMonitor) onNodeUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// onNodeDelete is the event handler for deleting Node.
func (m *NodeLatencyMonitor) onNodeDelete(obj interface{}) { _ = "STUB: not implemented"; return }

// onNodeLatencyMonitorAdd is the event handler for adding NodeLatencyMonitor.
func (m *NodeLatencyMonitor) onNodeLatencyMonitorAdd(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

// onNodeLatencyMonitorUpdate is the event handler for updating NodeLatencyMonitor.
func (m *NodeLatencyMonitor) onNodeLatencyMonitorUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// updateLatencyConfig updates the latency config based on the NodeLatencyMonitor CRD.
func (m *NodeLatencyMonitor) updateLatencyConfig(nlm *v1alpha1.NodeLatencyMonitor) {
	_ = "STUB: not implemented"
	return
}

// onNodeLatencyMonitorDelete is the event handler for deleting NodeLatencyMonitor.
func (m *NodeLatencyMonitor) onNodeLatencyMonitorDelete(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func icmpEchoData(ts time.Time) []byte { _ = "STUB: not implemented"; return nil }

// sendPing sends an ICMP message to the target IP address.
func (m *NodeLatencyMonitor) sendPing(socket net.PacketConn, addr net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// Serialize the ICMP message

// Send the ICMP message

// Create or update the latency store

func (m *NodeLatencyMonitor) handlePing(buffer []byte, peerIP string, isIPv4 bool) {
	_ = "STUB: not implemented"
	// Parse the ICMP message
	return
}

// Ignore ICMP echo messages received from other Nodes (they will be answered by the system)

// Ignore ICMP echo messages received from other Nodes (they will be answered by the system)

// Parse the time from the ICMP data

// Calculate the round-trip time

// Update the latency store

// recvPings receives ICMP messages.
func (m *NodeLatencyMonitor) recvPings(socket net.PacketConn, isIPv4 bool) {
	_ = "STUB: not implemented"
	// We only expect small packets, if we receive a larger packet, we will drop the extra data.
	return
}

// When the socket is closed in the Run method, this error will be logged, which is not ideal.
// In the future, we may try setting a ReadDeadline on the socket before each ReadFrom and using
// a channel to signal that the loop should terminate.

// pingAll sends ICMP messages to all the Nodes.
func (m *NodeLatencyMonitor) pingAll(ipv4Socket, ipv6Socket net.PacketConn) {
	_ = "STUB: not implemented"
	return
}

// getSummary returns the latency summary of the given Node IP.
func (m *NodeLatencyMonitor) getSummary() *statsv1alpha1.NodeLatencyStats {
	_ = "STUB: not implemented"
	return nil
}

func (m *NodeLatencyMonitor) report() { _ = "STUB: not implemented"; return }

// Run starts the NodeLatencyMonitor.
func (m *NodeLatencyMonitor) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// monitorLoop is the main loop to monitor the latency of the Node.
func (m *NodeLatencyMonitor) monitorLoop(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Update ping ticker based on the latencyConfig

// Stop the pingTicker

// Update report ticker with minimum interval and jitter

// Add jitter to avoid lockstep reporting

// Stop the reportTicker

// Start the pingAll goroutine

// Try to send pingAll signal

// We no not delete IPs from nodeIPLatencyMap as part of the Node delete event handler
// to avoid consistency issues and because it would not be sufficient to avoid stale entries completely.
// This means that we have to periodically invoke DeleteStaleNodeIPs to avoid stale entries in the map.

// Start or stop the pingAll goroutine based on the latencyConfig

// latencyConfig changed

// If the recvPing socket is closed,
// recreate it if it is closed (CR is deleted).

// Create a new socket for IPv4 when it is IPv4-only

// Create a new socket for IPv6 when it is IPv6-only

// We close the sockets as a signal to recvPing that it needs to stop.
// Note that at that point, we are guaranteed that there is no ongoing Write
// to the socket, because pingAll runs in the same goroutine as this code.

// After closing the sockets, wait for the recvPing goroutines to return

// getICMPSeqNum returns the sequence number to be used when sending the next
// ICMP echo request. It wraps around to 0 after reaching the maximum value for
// uint16.
func (m *NodeLatencyMonitor) getICMPSeqNum() uint16 { _ = "STUB: not implemented"; return 0 }
