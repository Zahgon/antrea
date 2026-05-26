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

package networkpolicy

import (
	"sync"
	"time"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/utils/clock"

	"antrea.io/antrea/v2/pkg/agent/types"
)

const (
	MinAllocatorAsyncDeleteInterval = 5 * time.Second

	deleteQueueName = "async_delete_networkpolicyrule"
)

// idAllocator provides interfaces to allocate and release uint32 IDs. It's thread-safe.
// It caches the last allocated ID and the IDs that have been released.
// If no IDs that have been released, the next allocated IP will be lastAllocatedID+1.
// If there are IDs that have been released, they will be reused FIFO.
type idAllocator struct {
	sync.Mutex
	// lastAllocatedID is the last allocated ID.
	// IDs that are greater than it must be available.
	// IDs that are less than or equal to it are available if they are in availableSet,
	// otherwise unavailable.
	lastAllocatedID uint32

	// availableSet maintains the IDs that can be reused for allocation.
	availableSet map[uint32]struct{}
	// availableSlice maintains the order of release.
	availableSlice []uint32
	// asyncRuleCache maintains rules in a cache and deletes the rules asynchronously
	// after a given delete interval.
	asyncRuleCache cache.Store
	// deleteQueue is used to place a rule ID after a given delay for deleting the
	// the rule in the asyncRuleCache.
	deleteQueue workqueue.TypedDelayingInterface[uint32]
	// deleteInterval is the delay interval for deleting the rule in the asyncRuleCache.
	deleteInterval time.Duration
}

// asyncRuleCacheKeyFunc knows how to get key of a *rule.
func asyncRuleCacheKeyFunc(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// newIDAllocatorWithClock creates an ID allocator with a custom clock, which is
// useful when writing unit tests.
func newIDAllocatorWithClock(asyncRuleDeleteInterval time.Duration, clock clock.WithTicker, allocatedIDs ...uint32) *idAllocator {
	_ = "STUB: not implemented"
	return nil
}

// newIDAllocator returns a new *idAllocator.
// It takes a list of allocated IDs, which can be used for the restart case.
func newIDAllocator(asyncRuleDeleteInterval time.Duration, allocatedIDs ...uint32) *idAllocator {
	_ = "STUB: not implemented"
	return nil
}

// allocateForRule allocates an uint32 ID for a given rule if it's available, otherwise
// an error is returned. It will try to reuse the IDs that have been released first,
// then allocate a new ID by incrementing the last allocated one.
func (a *idAllocator) allocateForRule(rule *types.PolicyRule) error {
	_ = "STUB: not implemented"
	return nil
}

// Add ID to the rule and the rule to asyncRuleCache.

// Add ID to the rule and the rule to asyncRuleCache.

// forgetRule adds the rule to the async delete queue with a given delay.
func (a *idAllocator) forgetRule(ruleID uint32) { _ = "STUB: not implemented"; return }

func (a *idAllocator) getRuleFromAsyncCache(ruleID uint32) (*types.PolicyRule, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (a *idAllocator) runWorker(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// worker runs a worker thread that just dequeues item from deleteQueue,
// deletes them from the asyncRuleCache, and releases the associated ID.
func (a *idAllocator) worker() { _ = "STUB: not implemented"; return }

func (a *idAllocator) processDeleteQueueItem() bool { _ = "STUB: not implemented"; return false }

// release releases an uint32 ID if it has been allocated before, otherwise error is returned.
func (a *idAllocator) release(id uint32) error { _ = "STUB: not implemented"; return nil }

// l7VlanIDAllocator provides interfaces to allocate and release VLAN IDs for L7 rules. It also caches the mapping of
// rule IDs to released VLAN IDs and provides an interface for L7 rule to query its allocated VLAN ID.
type l7VlanIDAllocator struct {
	sync.RWMutex

	idCounter      uint32
	recycled       []uint32
	ruleIDToVlanID map[string]uint32
}

func newL7VlanIDAllocator() *l7VlanIDAllocator { _ = "STUB: not implemented"; return nil }

func (l *l7VlanIDAllocator) allocate(ruleID string) uint32 { _ = "STUB: not implemented"; return 0 }

func (l *l7VlanIDAllocator) release(ruleID string) { _ = "STUB: not implemented"; return }

func (l *l7VlanIDAllocator) query(ruleID string) uint32 { _ = "STUB: not implemented"; return 0 }
