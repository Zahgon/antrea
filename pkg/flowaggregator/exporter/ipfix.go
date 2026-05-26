// Copyright 2022 Antrea Authors
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

package exporter

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/afero"
	ipfixentities "github.com/vmware/go-ipfix/pkg/entities"
	"github.com/vmware/go-ipfix/pkg/exporter"
	"k8s.io/apimachinery/pkg/util/wait"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
	flowaggregatorconfig "antrea.io/antrea/v2/pkg/config/flowaggregator"
	"antrea.io/antrea/v2/pkg/flowaggregator/options"
	"antrea.io/antrea/v2/pkg/flowaggregator/ringbuffer"
	"antrea.io/antrea/v2/pkg/ipfix"
)

var (
	// this is used for unit testing
	initIPFIXExportingProcess = func(exporter *IPFIXExporter) error {
		return exporter.initExportingProcessImpl()
	}

	defaultFS = afero.NewOsFs()
)

const (
	flowCollectorCertDir = "/etc/flow-aggregator/certs/flow-collector"

	flushInterval = 1 * time.Second
)

type IPFIXExporter struct {
	config                     flowaggregatorconfig.FlowCollectorConfig
	externalFlowCollectorAddr  string
	externalFlowCollectorProto string
	exportingProcess           ipfix.IPFIXExportingProcess
	bufferedExporter           ipfix.IPFIXBufferedExporter
	sendJSONRecord             bool
	includeK8sNames            bool
	includeK8sUIDs             bool
	aggregatorMode             flowaggregatorconfig.AggregatorMode
	observationDomainID        uint32
	templateRefreshTimeout     time.Duration
	templateIDv4               uint16
	templateIDv6               uint16
	elementsV4                 []ipfixentities.InfoElementWithValue
	elementsV6                 []ipfixentities.InfoElementWithValue
	registry                   ipfix.IPFIXRegistry
	clusterUUID                uuid.UUID
	clusterID                  string
	maxIPFIXMsgSize            int
	tls                        ipfixExporterTLSConfig
}

type ipfixExporterTLSConfig struct {
	enable                          bool
	minVersion                      uint16
	externalFlowCollectorCAPath     string
	externalFlowCollectorServerName string
	exporterCertPath                string
	exporterKeyPath                 string
}

func newIPFIXExporterTLSConfig(config flowaggregatorconfig.FlowCollectorTLSConfig) ipfixExporterTLSConfig {
	_ = "STUB: not implemented"
	return *new(ipfixExporterTLSConfig)
}

// config.MinVersion has already been validated during FA config validation.

// genObservationDomainID generates an IPFIX Observation Domain ID when one is not provided by the
// user through the flow aggregator configuration. It is generated as a hash of the cluster UUID.
func genObservationDomainID(clusterUUID uuid.UUID) uint32 { _ = "STUB: not implemented"; return 0 }

func newInitBackoff() wait.Backoff { _ = "STUB: not implemented"; return *new(wait.Backoff) }

func NewIPFIXExporter(
	clusterUUID uuid.UUID,
	clusterID string,
	opt *options.Options,
	registry ipfix.IPFIXRegistry,
) *IPFIXExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *IPFIXExporter) reset() { _ = "STUB: not implemented"; return }

// Run consumes flow records from the ring buffer and exports them via IPFIX.
// It blocks until ctx is cancelled or the consumer signals shutdown.
func (e *IPFIXExporter) Run(ctx context.Context, buf ringbuffer.BroadcastBuffer[*flowpb.Flow]) {
	_ = "STUB: not implemented"
	return
}

// consumeDeadline must be <= flushInterval so that Consume returns often
// enough for the flush ticker to be checked promptly.

// Using AfterFunc makes more sense than After: we want to use a custom
// channel and close it when we need the <- waitCh case to act as a
// "default" case in the select statement below.

// initBackoff is used to enforce some minimum delay between initialization attempts.

// initNextAttempt is the time after which the next initialization can be attempted.

// Note that e.flush() will be a no-op and return nil if the exporting
// process is not initialized.

// if waitCh is closed, this case acts as a "default" case

// Safety net: in the normal flow waitFor() is always called with the
// exact remaining duration before scheduling the next attempt, so
// this branch should not be reached. It protects against any future
// code path that reaches the init block without having waited.

// In case of error, we drop the current record and reset the
// exporting process. The next iteration of the loop will be
// responsible for re-initializing the exporting process.

func (e *IPFIXExporter) flush() error { _ = "STUB: not implemented"; return nil }

func (e *IPFIXExporter) makeIPFIXRecord(flow *flowpb.Flow, isIPv6 bool) ipfixentities.Record {
	_ = "STUB: not implemented"
	return *new(ipfixentities.Record)
}

// next is a convenience function to access and write information elements sequentially.
// All elements in the slice should be set. In other words, the number of calls to next() in
// this function should exactly match the length of the elements slice (created by
// prepareElements). We rely on unit testing to ensure this.

// IANA IEs

// IANAReverse IEs

// Antrea IEs

// Use Getter functions in case transport is not TCP

// Add Antrea source stats fields

// Add Antrea destination stats fields

// Add Antrea flow end seconds fields

// Add common throughput fields

// Add Pod label fields

// flow.K8S.SourcePodLabels.Labels can be nil or an empty map
// both cases should be treated the same

// Proxy-mode specific IEs

func (e *IPFIXExporter) sendRecord(flow *flowpb.Flow, isRecordIPv6 bool) error {
	_ = "STUB: not implemented"
	// Run() always initializes the exporting process before calling sendRecord,
	// so this guard should never be triggered in practice.
	return nil
}

func inPod() bool { _ = "STUB: not implemented"; return false }

func getMTU(ifaceName string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (e *IPFIXExporter) prepareExportingProcessTLSClientConfig() (*exporter.ExporterTLSClientConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *IPFIXExporter) initExportingProcess() error { _ = "STUB: not implemented"; return nil }

func (e *IPFIXExporter) initExportingProcessImpl() error {
	_ = "STUB: not implemented"
	// We reload the certificate data every time, in case the files have been updated.
	return nil
}

// TCP transport does not need any tempRefTimeout, so sending 0.

// In a Pod, the primary network interface is always "eth0", and we assume
// this is the interface used to connect to the IPFIX collector.
// The FlowAggregator is not meant to be run in the host network.

// In practice the only guarantee we have is that PMTU <=
// MTU. However, this is a reasonable approximation for most
// scenarios. Note that MaxMessageSize is an available override in
// the config.

// Currently, we send two templates for IPv4 and IPv6 regardless of the IP families supported by cluster

func (e *IPFIXExporter) createAndSendTemplate(isRecordIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// These elements will be used for data records as well, to avoid extra memory allocations.

// No need to flush first, as no data records should have been sent yet.

func (e *IPFIXExporter) prepareElements(isIPv6 bool) ([]ipfixentities.InfoElementWithValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add Antrea source stats fields

// Add Antrea destination stats fields

// Add common throughput fields

// Add source node specific throughput fields

// Add destination node specific throughput fields

func (e *IPFIXExporter) sendTemplateSet(isIPv6 bool) error { _ = "STUB: not implemented"; return nil }

// Ideally we would not have to do it explicitly, it would be taken care of by the go-ipfix library.

func (e *IPFIXExporter) createInfoElement(ieName string, enterpriseID uint32) (ipfixentities.InfoElementWithValue, error) {
	_ = "STUB: not implemented"
	return *new(ipfixentities.InfoElementWithValue), nil
}
