// Copyright 2021 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the “License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an “AS IS” BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package priorityqueue

import (
	"time"

	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
)

// minExpiryTime provides two usages: 1. We want to avoid passing a non positive
// value to ticker 2. We want to avoid processing a single expired item per call.
// If multiple items have very close expiry time(<100ms), by adding a small constant,
// we can make these items expired and process in one call.
const minExpiryTime = 100 * time.Millisecond

type ExpirePriorityQueue struct {
	items             []*ItemToExpire
	ActiveFlowTimeout time.Duration
	IdleFlowTimeout   time.Duration
	KeyToItem         map[connection.ConnectionKey]*ItemToExpire
}

func NewExpirePriorityQueue(activeFlowTimeout time.Duration, idleFlowTimeout time.Duration) *ExpirePriorityQueue {
	_ = "STUB: not implemented"
	return nil
}

func (pq *ExpirePriorityQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (pq *ExpirePriorityQueue) minExpireTime(i int) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (pq *ExpirePriorityQueue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pq *ExpirePriorityQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pq *ExpirePriorityQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (pq *ExpirePriorityQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// Peek returns the item at the beginning of the queue, without removing the
// item or otherwise mutating the queue. It is safe to call directly.
func (pq *ExpirePriorityQueue) Peek() *ItemToExpire { _ = "STUB: not implemented"; return nil }

// Update modifies the priority of an Item in the queue.
func (pq *ExpirePriorityQueue) Update(item *ItemToExpire, activeExpireTime time.Time, idleExpireTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// Remove removes and returns an Item by key from priority queue if it exists.
func (pq *ExpirePriorityQueue) Remove(connKey connection.ConnectionKey) *ItemToExpire {
	_ = "STUB: not implemented"
	return nil
}

// Clear removes all items from the queue and key index map.
func (pq *ExpirePriorityQueue) Clear() { _ = "STUB: not implemented"; return }

// GetExpiryFromExpirePriorityQueue returns the shortest expire time duration
// from expire priority queue.
func (pq *ExpirePriorityQueue) GetExpiryFromExpirePriorityQueue() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Get the minExpireTime of the top item in ExpirePriorityQueue.

// WriteItemToQueue adds conn with connKey into the queue. If an existing item
// has the same connKey, it will be overwritten by the new item.
func (pq *ExpirePriorityQueue) WriteItemToQueue(connKey connection.ConnectionKey, conn *connection.Connection) {
	_ = "STUB: not implemented"
	return
}

// If connKey exists in pq, it is removed first to avoid having multiple pqItems with same key
// in the queue, which can cause memory leak as the previous one can't be updated or removed.

func (pq *ExpirePriorityQueue) ResetActiveExpireTimeAndPush(pqItem *ItemToExpire, currTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func (pq *ExpirePriorityQueue) RemoveItemFromMap(conn *connection.Connection) {
	_ = "STUB: not implemented"
	return
}

func (pq *ExpirePriorityQueue) GetTopExpiredItem(currTime time.Time) *ItemToExpire {
	_ = "STUB: not implemented"
	// If the queue is empty or top item is not timeout, then we do not have to
	// check the following items.
	return nil
}
