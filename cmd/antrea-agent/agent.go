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

package main

import (
	"net"
	"time"
)

// informerDefaultResync is the default resync period if a handler doesn't specify one.
// Use the same default value as kube-controller-manager:
// https://github.com/kubernetes/kubernetes/blob/release-1.17/pkg/controller/apis/config/v1alpha1/defaults.go#L120
const informerDefaultResync = 12 * time.Hour

// resyncPeriodDisabled is 0 to disable resyncing.
// UpdateFunc event handler will be called only when the object is actually updated.
const resyncPeriodDisabled = 0 * time.Minute

// The devices that should be excluded from NodePort.
var (
	excludeNodePortDevices        = []string{"antrea-egress0", "antrea-ingress0", "kube-ipvs0"}
	excludeNodePortDevicePrefixes = []string{"antrea-ext."}
)

var ipv4Localhost = net.ParseIP("127.0.0.1")

// run starts Antrea agent with the given options and waits for termination signal.
func run(o *Options) error { _ = "STUB: not implemented"; return nil }

// Create K8s Clientset, CRD Clientset, Multicluster CRD Clientset and SharedInformerFactory for the given config.

// Create Antrea Clientset for the given config.

// Register Antrea Agent metrics if EnablePrometheusMetrics is set

// Create ovsdb and openflow clients.

// TODO: ovsconfig.NewOVSDBConnectionUDS might return timeout in the future, need to add retry

// Bridging mode will connect the uplink interface to the OVS bridge.

// WithRequiredPortExternalIDs will ensure that whenever we create a port, the required
// external ID (interface type) is provided. This is a sanity check to ensure code
// correctness.

// Create an ifaceStore that caches network interfaces managed by this node.

// podNetworkWait is used to wait and notify that preconditions for Pod network are ready.
// Processes that are supposed to finish before enabling Pod network should increment the wait group and decrement
// it when finished.
// Processes that enable Pod network should wait for it.

// flowRestoreCompleteWait is used to wait until "essential" flows have been installed
// successfully in OVS. These flows include NetworkPolicy flows (guaranteed by
// podNetworkWait), Pod forwarding flows and flows installed by the
// NodeRouteController. Additional requirements may be added in the future.

// staleFlowsDeletedWait starts with one pending unit (Increment); the stale-flow cleanup
// goroutine calls Done when deletion completes.

// set up signal capture: the first SIGTERM / SIGINT signal is handled gracefully and will
// cause the stopCh channel to be closed; if another signal is received before the program
// exits, we will force exit.

// Generate a context for functions which require one (instead of stopCh).

// Must start after registering all event handlers.

// Get all available NodePort addresses.

// Initialize agent and node network.

// podUpdateChannel is a channel for receiving Pod updates from CNIServer and
// notifying NetworkPolicyController, StretchedNetworkPolicyController and
// EgressController to reconcile rules related to the updated Pods.

// externalEntityUpdateChannel is a channel for receiving ExternalEntity updates from ExternalNodeController and
// notifying NetworkPolicyController to reconcile rules related to the updated ExternalEntities.

// Lazily initialize localPodInformer when it's required by any module.

// NamespaceIndex is used in NPLController.

// We pick a time interval for rule deletion in the async rule cache (part of the
// idAllocator) based on the configured flow poll interval for the Flow Exporter. This is to
// preserve the rule info for populating NetworkPolicy fields in the Flow Exporter even
// after rule deletion, and avoid missing or even incorrect information in the flow records.
// In theory, anything slightly longer than the poll interval should work, but to
// accommodate for longer than usual poll cycles we choose to play it safe.
// o.pollInterval will be 0 when the Flow Exporter is not enabled.

// In Antrea agent, status manager will automatically be enabled if
// AntreaPolicy feature is enabled.

// Secondary network controller should be created before CNIServer.Run() to make sure no Pod CNI updates will be missed.

// Initialize the NPL agent.

// Antrea IPAM is needed by bridging mode and secondary network IPAM.

//  Start the localPodInformer

// If AntreaProxy is configured to proxy all Service traffic, we need to wait for it to sync at least once
// before moving forward. Components that rely on Service availability should run after it, otherwise accessing
// Service would fail.

// We ensure that flowRestoreCompleteWait.Wait() cannot return until podNetworkWait.Wait() returns.

// ConnectUplinkToOVSBridge must be run immediately after FlowRestoreComplete

// Restore network config before shutdown. ovsdbConnection must be alive when restore.

// secondaryNetworkController Initialize must be run after FlowRestoreComplete for the case that Node
// IPs are moved to the secondary OVS bridge

// statsCollector collects stats and reports to the antrea-controller periodically. For now it's only used for
// NetworkPolicy stats and Multicast stats.

// staleFlowsDeletedWait: the collector waits inside Run until the initializer's stale-flow
// cleanup calls Done, so we do not report stats for flows from the prior round (see
// Initializer.initOpenFlowPipeline and flowRestoreCompleteWait / podNetworkWait).

// The certificate is static and will not be rotated; it will be re-generated if the Agent restarts.

// The API certificate is passed on directly to the monitor, instead of being provided by
// the agentQuerier. This is to avoid a circular dependency between apiServer and
// agentQuerier. The apiServer already depends on the agentQuerier to implement some API
// handlers. The certificate data is only available after initializing the apiServer.

// Start PacketIn and OVS meter stats collection for Prometheus

// Start the goroutine to periodically export IPFIX flow records.

// Start the node latency monitor if applicable.
