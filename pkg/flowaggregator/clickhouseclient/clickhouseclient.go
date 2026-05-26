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

package clickhouseclient

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gammazero/deque"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
	"antrea.io/antrea/v2/pkg/flowaggregator/flowrecord"
)

const (
	ProtocolUnknown   = -1
	maxQueueSize      = 1 << 19 // 524288. ~500MB assuming 1KB per record
	queueFlushTimeout = 10 * time.Second
	insertQuery       = `INSERT INTO flows (
                   flowStartSeconds,
                   flowEndSeconds,
                   flowEndSecondsFromSourceNode,
                   flowEndSecondsFromDestinationNode,
                   flowEndReason,
                   sourceIP,
                   destinationIP,
                   sourceTransportPort,
                   destinationTransportPort,
                   protocolIdentifier,
                   packetTotalCount,
                   octetTotalCount,
                   packetDeltaCount,
                   octetDeltaCount,
                   reversePacketTotalCount,
                   reverseOctetTotalCount,
                   reversePacketDeltaCount,
                   reverseOctetDeltaCount,
                   sourcePodName,
                   sourcePodNamespace,
                   sourceNodeName,
                   destinationPodName,
                   destinationPodNamespace,
                   destinationNodeName,
                   destinationClusterIP,
                   destinationServicePort,
                   destinationServicePortName,
                   ingressNetworkPolicyName,
                   ingressNetworkPolicyNamespace,
                   ingressNetworkPolicyRuleName,
                   ingressNetworkPolicyRuleAction,
                   ingressNetworkPolicyType,
                   egressNetworkPolicyName,
                   egressNetworkPolicyNamespace,
                   egressNetworkPolicyRuleName,
                   egressNetworkPolicyRuleAction,
                   egressNetworkPolicyType,
                   tcpState,
                   flowType,
                   sourcePodLabels,
                   destinationPodLabels,
                   throughput,
                   reverseThroughput,
                   throughputFromSourceNode,
                   throughputFromDestinationNode,
                   reverseThroughputFromSourceNode,
                   reverseThroughputFromDestinationNode,
                   clusterUUID,
                   egressName,
                   egressIP,
                   egressNodeName)
                   VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 
                           ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
                           ?, ?, ?, ?, ?)`
)

// PrepareClickHouseConnection is used for unit testing
var PrepareClickHouseConnection = prepareConnection

type stopPayload struct {
	flushQueue bool
}

type ClickHouseExportProcess struct {
	// db holds sql connection struct to clickhouse db.
	db     *sql.DB
	config ClickHouseConfig
	// deque buffers flows records between batch commits.
	deque deque.Deque[*flowrecord.FlowRecord]
	// dequeMutex is for concurrency between adding and removing records from deque.
	dequeMutex sync.Mutex
	// queueSize is the max size of deque
	queueSize int
	// stopCh is the channel to receive stop message
	stopCh chan stopPayload
	// exportWg is to ensure that all messages have been flushed from the queue when we stop
	exportWg sync.WaitGroup
	// commitTicker is a ticker, containing a channel used to trigger batchCommitAll() for every commitInterval period
	commitTicker         *time.Ticker
	exportProcessRunning bool
	// mutex protects configuration state from concurrent access
	mutex       sync.Mutex
	clusterUUID string
}

type ClickHouseConfig struct {
	Username           string
	Password           string
	Database           string
	DatabaseURL        string
	Debug              bool
	Compress           *bool
	CommitInterval     time.Duration
	CACert             bool
	InsecureSkipVerify bool
	Certificate        []byte
}

func NewClickHouseClient(config ClickHouseConfig, clusterUUID string) (*ClickHouseExportProcess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ch *ClickHouseExportProcess) CacheRecord(record *flowpb.Flow) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *ClickHouseExportProcess) Start() { _ = "STUB: not implemented"; return }

func (ch *ClickHouseExportProcess) Stop() { _ = "STUB: not implemented"; return }

func (ch *ClickHouseExportProcess) startExportProcess() { _ = "STUB: not implemented"; return }

func (ch *ClickHouseExportProcess) stopExportProcess(flushQueue bool) {
	_ = "STUB: not implemented"
	return
}

func (ch *ClickHouseExportProcess) flowRecordPeriodicCommit() { _ = "STUB: not implemented"; return }

// batchCommitAll commits all flow records cached in local deque in one INSERT query.
// Returns the number of records successfully committed, and error if encountered.
// Cached records will be removed only after successful commit.
func (ch *ClickHouseExportProcess) batchCommitAll(ctx context.Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// start new connection

// populate items from deque

// currSize could have increased due to CacheRecord being called in between.

// pushRecordsToFrontOfQueue pushes records to the front of deque without exceeding its capacity.
// Items with lower index (older records) will be dropped first if deque is to be filled.
func (ch *ClickHouseExportProcess) pushRecordsToFrontOfQueue(records []*flowrecord.FlowRecord) {
	_ = "STUB: not implemented"
	return
}

func prepareConnection(config ClickHouseConfig) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Test open Transaction

func (ch *ClickHouseExportProcess) UpdateCH(config ClickHouseConfig, connect *sql.DB) {
	_ = "STUB: not implemented"
	return
}

// do not flush the queue

func (ch *ClickHouseExportProcess) GetCommitInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (ch *ClickHouseExportProcess) SetCommitInterval(commitInterval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (ch *ClickHouseExportProcess) GetClickHouseConfig() ClickHouseConfig {
	_ = "STUB: not implemented"
	return *new(ClickHouseConfig)
}

func ConnectClickHouse(config *ClickHouseConfig) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Connect to ClickHouse in a loop

// Open the database and ping it

// #nosec G402: ignore insecure options

func parseDatabaseURL(dbUrl string) (clickhouse.Protocol, string, bool, error) {
	_ = "STUB: not implemented"
	return *new(clickhouse.Protocol), "", false, nil
}
