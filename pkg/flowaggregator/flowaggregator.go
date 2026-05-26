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

package flowaggregator

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/google/uuid"
	"k8s.io/client-go/kubernetes"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
	flowaggregatorconfig "antrea.io/antrea/v2/pkg/config/flowaggregator"
	"antrea.io/antrea/v2/pkg/flowaggregator/certificate"
	"antrea.io/antrea/v2/pkg/flowaggregator/collector"
	"antrea.io/antrea/v2/pkg/flowaggregator/exporter"
	"antrea.io/antrea/v2/pkg/flowaggregator/intermediate"
	"antrea.io/antrea/v2/pkg/flowaggregator/options"
	"antrea.io/antrea/v2/pkg/flowaggregator/querier"
	"antrea.io/antrea/v2/pkg/flowaggregator/ringbuffer"
	"antrea.io/antrea/v2/pkg/ipfix"
	"antrea.io/antrea/v2/pkg/util/objectstore"
)

const aggregationWorkerNum = 2

// these are used for unit testing
var (
	newIPFIXExporter = func(clusterUUID uuid.UUID, clusterID string, opt *options.Options, registry ipfix.IPFIXRegistry) exporter.Runner {
		return exporter.NewIPFIXExporter(clusterUUID, clusterID, opt, registry)
	}
	newClickHouseExporter = func(clusterUUID uuid.UUID, opt *options.Options) (exporter.Runner, error) {
		return exporter.NewClickHouseExporter(clusterUUID, opt)
	}
	newS3Exporter = func(clusterUUID uuid.UUID, opt *options.Options) (exporter.Runner, error) {
		return exporter.NewS3Exporter(clusterUUID, opt)
	}
	newLogExporter = func(opt *options.Options) (exporter.Runner, error) {
		return exporter.NewLogExporter(opt)
	}

	newCertificateProvider = func(k8sClient kubernetes.Interface, addr string) *certificate.Provider {
		return certificate.NewProvider(k8sClient, addr)
	}
)

// exporterHandle tracks the lifecycle of a running exporter goroutine.
type exporterHandle struct {
	exporter exporter.Runner
	// cancel cancels the context passed to the exporter's Run method,
	// signalling it to stop.
	cancel context.CancelFunc
	// doneCh is closed when the exporter's Run goroutine has returned.
	doneCh chan struct{}
}

// stop cancels the exporter's context and waits for its goroutine to exit.
func (h *exporterHandle) stop() { _ = "STUB: not implemented"; return }

type flowAggregator struct {
	aggregatorMode              flowaggregatorconfig.AggregatorMode
	clusterUUID                 uuid.UUID
	clusterID                   string
	aggregatorTransportProtocol flowaggregatorconfig.AggregatorTransportProtocol
	collectorMutex              sync.Mutex
	certificateUpdateCh         chan struct{}
	ipfixCollector              collector.Interface
	grpcCollector               collector.Interface
	aggregationProcess          intermediate.AggregationProcess
	activeFlowRecordTimeout     time.Duration
	inactiveFlowRecordTimeout   time.Duration
	registry                    ipfix.IPFIXRegistry
	flowAggregatorAddress       string
	includePodLabels            bool
	includeK8sUIDs              bool
	k8sClient                   kubernetes.Interface
	podStore                    objectstore.PodStore
	nodeStore                   objectstore.NodeStore
	serviceStore                objectstore.ServiceStore
	numRecordsExported          atomic.Int64
	// numRecordsDropped is always 0 with the current ring-buffer design: records
	// are produced to the buffer (which may discard them when full, counted
	// separately by the buffer itself) and proxyRecord no longer returns an error.
	// Kept for API/metrics compatibility.
	numRecordsDropped atomic.Int64
	updateCh          chan *options.Options
	configFile        string
	configWatcher     *fsnotify.Watcher
	configData        []byte
	APIServer         flowaggregatorconfig.APIServerConfig
	logTickerDuration time.Duration
	recordCh          chan *flowpb.Flow

	recordBuffer        ringbuffer.BroadcastBuffer[*flowpb.Flow]
	ipfixHandle         *exporterHandle
	clickHouseHandle    *exporterHandle
	s3Handle            *exporterHandle
	logHandle           *exporterHandle
	exportersMutex      sync.Mutex
	certificateProvider *certificate.Provider
}

func NewFlowAggregator(
	k8sClient kubernetes.Interface,
	clusterUUID uuid.UUID,
	podStore objectstore.PodStore,
	nodeStore objectstore.NodeStore,
	serviceStore objectstore.ServiceStore,
	configFile string,
) (*flowAggregator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When watching the configuration file directly, we have to add the file back to our watcher whenever the configuration
// file is modified (The watcher cannot track the config file  when the config file is replaced).
// Watching the directory can prevent us from above situation.

// #nosec G703: path is provided by a cluster admin via command-line args to the flow-aggregator; no privilege boundary crossed.

// launchExporter creates a new consumer from the ring buffer, starts the
// exporter's Run method in a goroutine, and returns a handle for lifecycle
// management.
func (fa *flowAggregator) launchExporter(exp exporter.Runner) *exporterHandle {
	_ = "STUB: not implemented"
	return nil
}

func (fa *flowAggregator) initCollectors() error { _ = "STUB: not implemented"; return nil }

func (fa *flowAggregator) InitAggregationProcess() error { _ = "STUB: not implemented"; return nil }

func (fa *flowAggregator) CertificateUpdated() { _ = "STUB: not implemented"; return }

func (fa *flowAggregator) runCollectors(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (fa *flowAggregator) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// We first wait for the object stores to sync to avoid lookup failures when processing records.

// Stores not synced within a reasonable time. We continue with the rest of the
// function but there may be error logs when processing records.

// blocking function, will return when fa.aggregationProcess.Stop() is called

// initExporters reads the initial configuration and launches exporter
// goroutines for all enabled exporters.
func (fa *flowAggregator) initExporters() { _ = "STUB: not implemented"; return }

// stopAllExporters stops all running exporter goroutines.
func (fa *flowAggregator) stopAllExporters() { _ = "STUB: not implemented"; return }

// flowExportLoop reads records from recordCh, enriches them, and produces them
// into the ring buffer. Each exporter independently consumes from the buffer.
func (fa *flowAggregator) flowExportLoop(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (fa *flowAggregator) proxyRecord(record *flowpb.Flow) { _ = "STUB: not implemented"; return }

// !withDestination should be redundant here

// egress

// !withSource should be redundant here

// ingress

// egress

// ingress

// this covers the IntraNode case

// This is the only case where K8s metadata could be missing

func (fa *flowAggregator) flowExportLoopProxy(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// set the channel to nil and essentially disable this select case.
// we could also just return straightaway as this should only happen
// when stopCh is closed, but maybe it's better to keep stopCh as
// the only signal for stopping the event loop.

func (fa *flowAggregator) flowExportLoopAggregate(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// Every 1s, we check for expired records and we export all of them. 1s is small enough to have good accuracy,
	// and long enough that we don't call the function too often. In the future, we can support handling batches of
	// expired records, which we can add to the ring buffer in a single API call.
	// Note that ActiveFlowRecordTimeout / InactiveFlowRecordTimeout values under 1s are not reasonable and not supported.
	return
}

// set the channel to nil and essentially disable this select case.
// we could also just return straightaway as this should only happen
// when stopCh is closed, but maybe it's better to keep stopCh as
// the only signal for stopping the event loop.

// produceRecord publishes a flow record into the ring buffer for all active
// exporter consumers to pick up independently.
func (fa *flowAggregator) produceRecord(record *flowpb.Flow) { _ = "STUB: not implemented"; return }

func (fa *flowAggregator) sendAggregatedRecord(key intermediate.FlowKey, record *intermediate.AggregationFlowRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// Even if fa.includePodLabels is false, we still need to add an empty IE to match the template.

// In Aggregate mode, we need to clone the record before placing it in the ring buffer, as
// it is "owned" by the aggregation process and will be updated as new records are received
// from the FlowExporters.

func (fa *flowAggregator) getNodeUID(nodeName string, startTime time.Time) string {
	_ = "STUB: not implemented"
	return ""
}

// fillK8sMetadata fills Pod name, Pod namespace, Pod UID, Node name and Node UID for inter-Node flows.
// This function is used in Proxy mode, as well as in Aggregate mode when correlation cannot be
// performed because of network policies.
func (fa *flowAggregator) fillK8sMetadata(sourceAddress, destinationAddress string, record *flowpb.Flow, startTime time.Time) {
	_ = "STUB: not implemented"
	// fill source Pod info when sourcePodName is empty
	return
}

// fill destination Pod info when destinationPodName is empty

func (fa *flowAggregator) fetchPodLabels(ip string, startTime time.Time) *flowpb.Labels {
	_ = "STUB: not implemented"
	return nil
}

// Labels field is of type map[string]string.
// Note that Protobuf treats nil and empty maps the same when it comes to
// serialization, and they should be treated the same in our Go code as well.

func (fa *flowAggregator) fillPodLabels(sourceAddress, destinationAddress string, record *flowpb.Flow, startTime time.Time) {
	_ = "STUB: not implemented"
	// If fa.includePodLabels is false, we always use nil.
	// If fa.includePodLabels is true, we use nil in case of error or if the endpoint is not a Pod.
	return
}

func (fa *flowAggregator) fillServiceUID(record *flowpb.Flow, startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func (fa *flowAggregator) fillEgressNodeUID(record *flowpb.Flow, startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func (fa *flowAggregator) GetFlowRecords(flowKey *intermediate.FlowKey) []map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (fa *flowAggregator) getNumFlows() int64 { _ = "STUB: not implemented"; return 0 }

func (fa *flowAggregator) getNumRecordsReceived() int64 { _ = "STUB: not implemented"; return 0 }

func (fa *flowAggregator) getNumConnsToCollector() int64 { _ = "STUB: not implemented"; return 0 }

func (fa *flowAggregator) GetRecordMetrics() querier.Metrics {
	_ = "STUB: not implemented"
	return *new(querier.Metrics)
}

func (fa *flowAggregator) watchConfiguration(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// If configWatcher event channel is closed, we kill the flow-aggregator Pod to restore
// the channel.

// If the watcher cannot add mounted configuration file or the configuration file is not readable,
// we kill the flow-aggregator Pod (serious error)

// If the error happens to watcher, we kill the flow-aggregator Pod.
// watcher might be shut-down or broken in this situation.

func (fa *flowAggregator) handleWatcherEvent() error { _ = "STUB: not implemented"; return nil }

// all updates must be performed within flowExportLoop

func (fa *flowAggregator) updateFlowAggregator(opt *options.Options) {
	_ = "STUB: not implemented"
	// This function potentially modifies the exporter handle fields (e.g.,
	// fa.ipfixHandle). We protect these writes by locking fa.exportersMutex, so
	// that GetRecordMetrics() can safely read the fields (by also locking the mutex).
	return
}

// If user tries to change the mode dynamically, it makes sense to error out
// immediately and ignore other updates, as this is such a major configuration
// parameter. Unsupported "minor" updates are handled at the end of this function.

// IPFIX exporter: stop-and-replace

// ClickHouse exporter: stop-and-replace

// S3 exporter: stop-and-replace

// Log exporter: stop-and-replace
