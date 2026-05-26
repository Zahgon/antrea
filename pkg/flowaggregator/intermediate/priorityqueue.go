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
	"time"
)

type ItemToExpire struct {
	// Flow related info
	flowKey            *FlowKey
	flowRecord         *AggregationFlowRecord
	activeExpireTime   time.Time
	inactiveExpireTime time.Time
	// Index in the priority queue (heap)
	index int
}

type TimeToExpirePriorityQueue []*ItemToExpire

func (pq TimeToExpirePriorityQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (pq TimeToExpirePriorityQueue) minExpireTime(i int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (pq TimeToExpirePriorityQueue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pq TimeToExpirePriorityQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pq *TimeToExpirePriorityQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (pq *TimeToExpirePriorityQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// Peek returns the item at the beginning of the queue, without removing the
// item or otherwise mutating the queue. It is safe to call directly.
func (pq TimeToExpirePriorityQueue) Peek() *ItemToExpire {
	_ = "STUB: not implemented"

	// update modifies the priority and flow record of an Item in the queue.
	return nil
}

func (pq *TimeToExpirePriorityQueue) Update(item *ItemToExpire, flowKey *FlowKey, flowRecord *AggregationFlowRecord, activeExpireTime time.Time, inactiveExpireTime time.Time) {
	_ = "STUB: not implemented"
	return
}
