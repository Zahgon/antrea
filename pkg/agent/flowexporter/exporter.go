// Copyright 2025 Antrea Authors
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

package flowexporter

import (
	"net"
	"sync"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/controller/noderoute"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/connections"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/exporter"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/options"
	"antrea.io/antrea/v2/pkg/agent/proxy"
	api "antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha1"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1alpha1"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
	"antrea.io/antrea/v2/pkg/querier"
	"antrea.io/antrea/v2/pkg/util/channel"

	"antrea.io/antrea/v2/pkg/util/objectstore"
	utilwait "antrea.io/antrea/v2/pkg/util/wait"
)

// When initializing flowExporter, a slice is allocated with a fixed size to
// store expired connections. The advantage is every time we export, the connection
// store lock will only be held for a bounded time. The disadvantages are: 1. the
// constant is independent of actual number of expired connections 2. when the
// number of expired connections goes over the constant, the export can not be
// finished in a single round. It could be delayed by conntrack connections polling
// routine, which also acquires the connection store lock. The possible solution
// can be taking a fraction of the size of connection store to approximate the
// number of expired connections, while having a min and a max to handle edge cases,
// e.g. min(50 + 0.1 * connectionStore.size(), 200)
const (
	maxConnsToExport = 64
	// How long to wait before retrying the processing of a FlowExporterDestination.
	minRetryDelay  = 5 * time.Second
	maxRetryDelay  = 300 * time.Second
	defaultWorkers = 1

	// We use a buffer of 1 since we batch the connections and send it as a slice.
	ctConnsUpdateChannelBufferSize int = 1
	// We use a buffer of 100 to handle situations where we get a burst of denied
	// denied connections
	denyConnUpdateChannelBufferSize int = 100
)

type destinationObj struct {
	stopCh      chan struct{}
	destination *Destination
}

type FlowExporter struct {
	k8sClient kubernetes.Interface

	destinationInformer crdinformers.FlowExporterDestinationInformer
	destinationSynced   cache.InformerSynced
	destinationLister   crdlisters.FlowExporterDestinationLister

	staleConnectionTimeout time.Duration
	v4Enabled              bool
	v6Enabled              bool
	isNetworkPolicyOnly    bool

	// Destination dependencies
	nodeRouteController *noderoute.Controller
	podStore            objectstore.PodStore
	proxier             proxy.ProxyQuerier
	egressQuerier       querier.EgressQuerier
	npQuerier           querier.AgentNetworkPolicyInfoQuerier

	// networkPolicyWait is used to determine when NetworkPolicy flows have been installed and
	// when the mapping from flow ID to NetworkPolicy rule is available. We will ignore
	// connections which started prior to that time to avoid reporting invalid NetworkPolicy
	// metadata in flow records. This is because the mapping is not "stable" and is expected to
	// change when the Agent restarts.
	networkPolicyWait      *utilwait.Group
	networkPolicyReadyTime time.Time

	poller                *connections.Poller
	ctConnUpdateChannel   *channel.SubscribableChannel
	denyConnUpdateChannel *channel.SubscribableChannel

	// staticDestinationRes is set in NewFlowExporter when static export is enabled. Run() clears
	// it if createDestinationFromResource fails; if still non-nil at shutdown, static destination
	// export was started and the poller must account for it.
	staticDestinationRes *api.FlowExporterDestination
	destinations         map[string]destinationObj
	destinationsMutex    sync.Mutex

	pollerRunning bool
	pollerStopCh  chan struct{}
	pollerDoneCh  chan struct{}

	// Used to create exporter
	nodeName    string
	nodeUID     string
	obsDomainID uint32

	queue workqueue.TypedRateLimitingInterface[string]
}

func NewFlowExporter(
	podStore objectstore.PodStore,
	proxier proxy.ProxyQuerier,
	k8sClient kubernetes.Interface,
	nodeRouteController *noderoute.Controller,
	trafficEncapMode config.TrafficEncapModeType,
	nodeConfig *config.NodeConfig,
	v4Enabled, v6Enabled bool,
	serviceCIDRNet, serviceCIDRNetv6 *net.IPNet,
	ovsDatapathType ovsconfig.OVSDatapathType,
	proxyEnabled bool,
	npQuerier querier.AgentNetworkPolicyInfoQuerier,
	o *options.FlowExporterOptions,
	destinationInformer crdinformers.FlowExporterDestinationInformer,
	egressQuerier querier.EgressQuerier,
	networkPolicyWait *utilwait.Group,
) (*FlowExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fe *FlowExporter) addDestination(obj any) { _ = "STUB: not implemented"; return }

func (fe *FlowExporter) updateDestination(old any, new any) { _ = "STUB: not implemented"; return }

func (fe *FlowExporter) deleteDestination(obj any) { _ = "STUB: not implemented"; return }

func (exp *FlowExporter) GetDenyConnStoreNotifier() channel.Notifier {
	_ = "STUB: not implemented"
	return *new(channel.Notifier)
}

// hasActiveDestinationsLocked reports whether at least one export destination (CR-based or static)
// is currently active. The caller must hold exp.destinationsMutex.
func (exp *FlowExporter) hasActiveDestinationsLocked() bool {
	_ = "STUB: not implemented"
	return false
}

// startPollerIfNeededLocked starts the conntrack poller when there is at least one active destination
// and the poller is not already running. The caller must hold exp.destinationsMutex.
func (exp *FlowExporter) startPollerIfNeededLocked() { _ = "STUB: not implemented"; return }

// If a previous poller goroutine is still exiting, wait for it to finish so we never have
// two concurrent poller.Run goroutines.

// stopPollerIfNeededLocked stops the conntrack poller when there are no active destinations and the
// poller is running. The caller must hold exp.destinationsMutex.
func (exp *FlowExporter) stopPollerIfNeededLocked() { _ = "STUB: not implemented"; return }

func (exp *FlowExporter) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Wait for NodeRouteController to have processed the initial list of Nodes so that
// the list of Pod subnets is up-to-date.

// SubscribableChannels must run whenever the flow exporter is active so Notify does not block
// producers (e.g. denied-connection updates) when there are no export destinations yet.

// Clear staticDestinationRes when createDestinationFromResource fails, so that
// hasActiveDestinationsLocked will function correctly.

// Clear staticDestinationRes so hasActiveDestinationsLocked returns false.

func (exp *FlowExporter) worker() { _ = "STUB: not implemented"; return }

func (exp *FlowExporter) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

func (exp *FlowExporter) syncFlowExporterDestination(key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fe *FlowExporter) createExporter(protocol exporterProtocol) exporter.Interface {
	_ = "STUB: not implemented"
	return *new(exporter.Interface)
}

func ServiceAddressToDNS(address string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (fe *FlowExporter) createDestinationFromResource(res *api.FlowExporterDestination) (*Destination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createStaticDestinationResFromOptions(o *options.FlowExporterOptions) (*api.FlowExporterDestination, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getExporterProtocol(proto api.FlowExporterProtocol) exporterProtocol {
	_ = "STUB: not implemented"
	return *new(exporterProtocol)
}

// This case should never happen on real usage. API server requires at least one to be defined.

func genObservationID(nodeName string) uint32 { _ = "STUB: not implemented"; return 0 }

func validateResource(res *api.FlowExporterDestination) error {
	_ = "STUB: not implemented"
	return nil
}
