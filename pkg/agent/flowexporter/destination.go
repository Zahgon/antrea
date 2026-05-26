// Copyright 2026 Antrea Authors.
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

package flowexporter

import (
	"context"
	"time"

	"k8s.io/client-go/kubernetes"

	"antrea.io/antrea/v2/pkg/agent/controller/noderoute"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/connections"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/exporter"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/priorityqueue"
	"antrea.io/antrea/v2/pkg/agent/proxy"
	api "antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	"antrea.io/antrea/v2/pkg/querier"
	"antrea.io/antrea/v2/pkg/util/channel"
	"antrea.io/antrea/v2/pkg/util/objectstore"
)

type exporterProtocol interface {
	Name() string
	TransportProtocol() api.FlowExporterTransportProtocol
}

type DestinationConfig struct {
	name    string
	address string

	activeFlowTimeout      time.Duration
	idleFlowTimeout        time.Duration
	staleConnectionTimeout time.Duration

	isNetworkPolicyOnly bool
	tlsConfig           *api.FlowExporterTLSConfig

	// allowProtocolFilter specifies whether the incoming connections will be accepted
	allowProtocolFilter []string

	networkPolicyReadyTime time.Time
}

type Destination struct {
	DestinationConfig

	k8sClient          kubernetes.Interface
	ctConnSubscriber   channel.Subscriber
	denyConnSubscriber channel.Subscriber

	conntrackConnStore     *connections.ConntrackConnectionStore
	conntrackPriorityQueue *priorityqueue.ExpirePriorityQueue

	denyConnStore     *connections.DenyConnectionStore
	denyPriorityQueue *priorityqueue.ExpirePriorityQueue

	nodeRouteController *noderoute.Controller
	egressQuerier       querier.EgressQuerier

	exp       exporter.Interface
	connected bool

	exportConns      []connection.Connection
	numConnsExported uint64
}

func NewDestination(
	ctConnSubscriber channel.Subscriber,
	denyConnSubscriber channel.Subscriber,
	exporter exporter.Interface,
	k8sClient kubernetes.Interface,
	nodeRouteController *noderoute.Controller,
	podStore objectstore.PodStore,
	npQuerier querier.AgentNetworkPolicyInfoQuerier,
	proxier proxy.ProxyQuerier,
	egressQuerier querier.EgressQuerier,
	networkPolicyReadyTime time.Time,
	destinationConfig DestinationConfig,
) *Destination {
	_ = "STUB: not implemented"
	return nil
}

func (d *Destination) getExporterTLSConfig(ctx context.Context) (*exporter.TLSConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if CA certificate, client certificate and key do not exist during initialization,
// it will retry to obtain the credentials in next export cycle

func (d *Destination) Connect(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *Destination) resetFlowExporter() { _ = "STUB: not implemented"; return }

func (d *Destination) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Start the goroutine to periodically delete stale deny connections.

// Initializing flow exporter fails, will retry in next cycle.

// If there is an error when sending flow records because of
// intermittent connectivity, we reset the connection to collector
// and retry in the next export cycle to reinitialize the connection
// and send flow records.

func (d *Destination) populateCTStore(e any) { _ = "STUB: not implemented"; return }

func (d *Destination) populateDenyStore(e any) { _ = "STUB: not implemented"; return }

func (d *Destination) sendFlowRecords() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// Select the shorter time out among two connection stores to do the next round of export.

// Clear expiredConns slice after exporting. Allocated memory is kept.

func (d *Destination) exportConn(conn *connection.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip exporting the Pod-to-External connection at the Egress Node if it's different from the Source Node

func (d *Destination) findFlowType(conn connection.Connection) uint8 {
	_ = "STUB: not implemented"
	// TODO: support Pod-To-External flows in network policy only mode.
	return 0
}

// This matches what we do in filterAntreaConns but is more general as we consider
// remote gateways as well.

func (d *Destination) fillEgressInfo(conn *connection.Connection) {
	_ = "STUB: not implemented"
	return
}

// Egress is not enabled or no Egress is applied to this Pod

func getMinTime(t1, t2 time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// resolveCollectorAddress resolves the collector address provided to an IP address if applicable or
// DNS name. The collector address can be a namespaced reference to a K8s Service, and hence needs
// resolution (to the Service's ClusterIP).
func resolveCollectorAddress(ctx context.Context, k8sClient kubernetes.Interface, address string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
