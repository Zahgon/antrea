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

package intermediate

import (
	"sync"
	"time"

	"k8s.io/utils/clock"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
)

var MaxRetries = 2

type aggregationProcess struct {
	// flowKeyRecordMap maps each connection (5-tuple) with its records
	flowKeyRecordMap map[FlowKey]*AggregationFlowRecord
	// expirePriorityQueue helps to maintain a priority queue for the records given
	// active expiry and inactive expiry timeouts.
	expirePriorityQueue TimeToExpirePriorityQueue
	// mutex allows multiple readers or one writer at the same time
	mutex sync.RWMutex
	// recordChan is the channel to receive the flow records to process
	recordChan <-chan *flowpb.Flow
	// workerNum is the number of workers to process the messages
	workerNum int
	// workerList is the list of workers
	workerList []aggregationWorker
	// activeExpiryTimeout helps in identifying records that elapsed active expiry
	// timeout. Active expiry timeout is a periodic expiry interval for every flow
	// record in the aggregation record map.
	activeExpiryTimeout time.Duration
	// inactiveExpiryTimeout helps in identifying records that elapsed inactive expiry
	// timeout. Inactive expiry timeout is an expiry interval that gets reset every
	// time a new record is received for the existing record in the aggregation
	// record map.
	inactiveExpiryTimeout time.Duration
	// stopChan is the channel to receive stop message
	stopChan chan bool
	clock    clock.Clock
}

type AggregationInput struct {
	RecordChan            <-chan *flowpb.Flow
	WorkerNum             int
	ActiveExpiryTimeout   time.Duration
	InactiveExpiryTimeout time.Duration
}

func initAggregationProcessWithClock(input AggregationInput, clock clock.Clock) (*aggregationProcess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InitAggregationProcess(input AggregationInput) (*aggregationProcess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *aggregationProcess) Start() { _ = "STUB: not implemented"; return }

func (a *aggregationProcess) Stop() { _ = "STUB: not implemented"; return }

// GetNumFlows returns total number of connections/flows stored in map
func (a *aggregationProcess) GetNumFlows() int64 { _ = "STUB: not implemented"; return 0 }

func (a *aggregationProcess) aggregateRecordByFlowKey(record *flowpb.Flow) error {
	_ = "STUB: not implemented"
	return nil
}

// ForAllRecordsDo takes in callback function to process the operations to flowkey->records pairs in the map
func (a *aggregationProcess) ForAllRecordsDo(callback FlowKeyRecordMapCallBack) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *aggregationProcess) deleteFlowKeyFromMap(flowKey FlowKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *aggregationProcess) deleteFlowKeyFromMapWithoutLock(flowKey FlowKey) error {
	_ = "STUB: not implemented"
	return nil
}

// GetRecords returns map format flow records given a flow key.
// In order to preserve backwards-compatibility (after migrating to Protobuf to represent flow
// records), map keys are the names of the corresponding information elements, and values are typed
// based on the IE type. Not all "elements" are included.
// Returns partially matched flow records if the flow key is not complete.
// Returns all the flow records if the flow key is not provided.
func (a *aggregationProcess) GetRecords(flowKey *FlowKey) []map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Complete filter

// Partial filter

func (a *aggregationProcess) ForAllExpiredFlowRecordsDo(callback FlowKeyRecordMapCallBack) error {
	_ = "STUB: not implemented"
	return nil
}

// We do not have to check other items anymore.

// Pop the record item from the priority queue

// Reset the timeouts and add the record to priority queue.
// Delete the record after max retries.

// We need to use !expireTime.After(currTime) and not expireTime.Before(currTime) to account for the
// equality case (expireTime == currTime) and match the if statement at the beginning of the for loop
// (which determines whether the item is expired or not).

// Delete the flow record if it is expired because of inactive expiry timeout.

// Reset the expireTime for the popped item and push it to the priority queue.

// Reset the active expire timeout and push the record into priority
// queue.

func (a *aggregationProcess) SetCorrelatedFieldsFilled(record *AggregationFlowRecord, isFilled bool) {
	_ = "STUB: not implemented"
	return
}

func (a *aggregationProcess) AreCorrelatedFieldsFilled(record AggregationFlowRecord) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *aggregationProcess) SetExternalFieldsFilled(record *AggregationFlowRecord, isFilled bool) {
	_ = "STUB: not implemented"
	return
}

func (a *aggregationProcess) AreExternalFieldsFilled(record AggregationFlowRecord) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *aggregationProcess) IsAggregatedRecordIPv4(record AggregationFlowRecord) bool {
	_ = "STUB: not implemented"
	return false

	// addOrUpdateRecordInMap either adds the record to flowKeyMap or updates the record in
	// flowKeyMap by doing correlation or updating the stats.
}

func (a *aggregationProcess) addOrUpdateRecordInMap(flowKey *FlowKey, record *flowpb.Flow, isIPv4 bool) {
	_ = "STUB: not implemented"
	return
}

// Do correlation of records if record belongs to inter-node flow and
// records from source and destination node are not received.

// Aggregation of incoming flow record with existing by updating stats
// and flow timestamps.

// For flows that do not need correlation, just do aggregation of the
// flow record with existing record by updating the stats and flow timestamps.

// Reset the inactive expiry time in the queue item with updated aggregate record.
// If this is the first time we have a "complete" record (ReadyToSend becomes true),
// export it right away by setting activeExpireTime to now.

// Add all the new stat fields and initialize them.

// If no correlation is required for an Inter-Node record, K8s metadata is
// expected to be not completely filled. For Intra-Node flows and ToExternal
// flows, areCorrelatedFieldsFilled is set to true by default.

// Push the record to the priority queue.

// If no correlation is required and the record is ReadyToSend, export it right away.

// correlateRecords correlate the incomingRecord with existingRecord using correlation
// fields. This is called for records whose flowType is InterNode.
func (a *aggregationProcess) correlateRecords(incomingRecord, existingRecord *flowpb.Flow) {
	_ = "STUB: not implemented"
	return
}

// aggregateRecords aggregate the incomingRecord with existingRecord by updating
// stats and flow timestamps.
func (a *aggregationProcess) aggregateRecords(incomingRecord, existingRecord *flowpb.Flow, fillSrcStats, fillDstStats bool) {
	_ = "STUB: not implemented"
	return
}

// Update the flowEndSecondsFromSource/DestinationNode fields, and compute
// the time difference between the incoming record and the last record.

// Skip the aggregation process if the incoming record is not the latest
// from its coming node; for intra-node flows. Also to avoid to assign
// zero value to flowEndSecondsDiff.

// If the aggregated flow is set with flowEndReason as "EndOfFlowReason", then we do not have to set again.

// Update tcpState when flow end timestamp is the latest.

// This code will need to change if more fields are added to Transport.Protocol.

// Update the throughput & reverseThroughput fields:
// throughput = (octetTotalCount - prevOctetTotalCount) / (flowEndSeconds - prevFlowEndSeconds)
// reverseThroughput = (reverseOctetTotalCount - prevReverseOctetTotalCount) / (flowEndSeconds - prevFlowEndSeconds)

// ResetStatAndThroughputElementsInRecord is called by the user after the aggregation
// record is sent after its expiry either by active or inactive expiry interval. This
// should be called by user after acquiring the mutex in the Aggregation process.
func (a *aggregationProcess) ResetStatAndThroughputElementsInRecord(record *flowpb.Flow) error {
	_ = "STUB: not implemented"
	// TotalCount statistic elements should not be reset to zeroes as they are used in the
	// throughput calculation.
	return nil
}

func (a *aggregationProcess) addFieldsForStatsAggregation(record *flowpb.Flow, fillSrcStats, fillDstStats bool) {
	_ = "STUB: not implemented"
	return
}

func (a *aggregationProcess) addFieldsForThroughputCalculation(record *flowpb.Flow, fillSrcStats, fillDstStats bool) {
	_ = "STUB: not implemented"
	return
}

// Initialize the throughput elements.

// For the edge case when the record has the same timeEnd and timeStart values,
// we will initialize the throughput fields with zero values.

// updateFlowEndSecondsFromNodes updates the value of flowEndSecondsFromSourceNode
// or flowEndSecondsFromDestinationNode, returning the previous value before update.
func (a *aggregationProcess) updateFlowEndSecondsFromNodes(incomingRecord, existingRecord *flowpb.Flow, isSrc bool, incomingVal int64) int64 {
	_ = "STUB: not implemented"
	return 0
}

// When the incoming record is the first record from its node, the existingVal of the field
// is zero, we set it by flowStartSeconds. time_diff = flowEndSeconds - flowStartSeconds

// isRecordFromSrc returns true if record belongs to inter-node flow and from source node.
func isRecordFromSrc(record *flowpb.Flow) bool { _ = "STUB: not implemented"; return false }

// isRecordFromDst returns true if record belongs to inter-node flow and from destination node.
func isRecordFromDst(record *flowpb.Flow) bool { _ = "STUB: not implemented"; return false }

func areRecordsFromSameNode(record1, record2 *flowpb.Flow) bool {
	_ = "STUB: not implemented"
	// If both records of inter-node flow are from source node, then send true.
	return false
}

// If both records of inter-node flow are from destination node, then send true.

// getFlowKeyFromRecord returns 5-tuple from data record
func getFlowKeyFromRecord(record *flowpb.Flow) (*FlowKey, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// isCorrelationRequired returns true for InterNode flowType when
// either the egressNetworkPolicyRuleAction is not deny (drop/reject) or
// the ingressNetworkPolicyRuleAction is not reject.
func isCorrelationRequired(flowType flowpb.FlowType, record *flowpb.Flow) bool {
	_ = "STUB: not implemented"
	return false
}
