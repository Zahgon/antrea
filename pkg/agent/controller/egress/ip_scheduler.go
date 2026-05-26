// Copyright 2023 Antrea Authors
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

package egress

import (
	"sync"
	"sync/atomic"

	corev1 "k8s.io/api/core/v1"
	corev1informers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/memberlist"
	crdv1b1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
)

const (
	// workItem is the only item that will be enqueued, used to trigger Egress IP scheduling.
	workItem = "key"
)

// scheduleEventHandler is a callback when an Egress is rescheduled.
type scheduleEventHandler func(egress string)

// scheduleResult is the schedule result of an Egress, including the effective Egress IP and Node.
type scheduleResult struct {
	ip   string
	node string
	err  error
}

// egressIPScheduler is responsible for scheduling Egress IPs to appropriate Nodes according to the Node selector of the
// IP pool, taking Node's capacity into consideration.
type egressIPScheduler struct {
	// cluster is responsible for selecting a Node for a given IP and pool.
	cluster memberlist.Interface

	egressLister       crdlisters.EgressLister
	egressListerSynced cache.InformerSynced

	// queue is used to trigger scheduling. Triggering multiple times before the item is consumed will only cause one
	// execution of scheduling.
	queue workqueue.TypedInterface[string]

	// mutex is used to protect scheduleResults.
	mutex           sync.RWMutex
	scheduleResults map[string]*scheduleResult
	// scheduledOnce indicates whether scheduling has been executed at lease once.
	scheduledOnce *atomic.Bool

	// eventHandlers is the registered callbacks.
	eventHandlers []scheduleEventHandler

	// The default maximum number of Egress IPs a Node can accommodate.
	maxEgressIPsPerNode int
	// nodeToMaxEgressIPs caches the maximum number of Egress IPs of each Node gotten from Node annotation.
	// It takes precedence over the default value.
	nodeToMaxEgressIPs      map[string]int
	nodeToMaxEgressIPsMutex sync.RWMutex
}

func NewEgressIPScheduler(cluster memberlist.Interface, egressInformer crdinformers.EgressInformer, nodeInformer corev1informers.NodeInformer, maxEgressIPsPerNode int) *egressIPScheduler {
	_ = "STUB: not implemented"
	return nil
}

// Trigger scheduling regardless of which pool is changed.

func getMaxEgressIPsFromAnnotation(node *corev1.Node) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// updateNode processes Node ADD and UPDATE events.
func (s *egressIPScheduler) updateNode(obj interface{}) { _ = "STUB: not implemented"; return }

// deleteNode processes Node DELETE events.
func (s *egressIPScheduler) deleteNode(obj interface{}) { _ = "STUB: not implemented"; return }

// addEgress processes Egress ADD events.
func (s *egressIPScheduler) addEgress(obj interface{}) { _ = "STUB: not implemented"; return }

// updateEgress processes Egress UPDATE events.
func (s *egressIPScheduler) updateEgress(old, cur interface{}) { _ = "STUB: not implemented"; return }

// deleteEgress processes Egress DELETE events.
func (s *egressIPScheduler) deleteEgress(obj interface{}) { _ = "STUB: not implemented"; return }

func (s *egressIPScheduler) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Schedule at least once even if there is no Egress to unblock clients waiting for HasScheduled to return true.

func (s *egressIPScheduler) HasScheduled() bool { _ = "STUB: not implemented"; return false }

func (s *egressIPScheduler) AddEventHandler(handler scheduleEventHandler) {
	_ = "STUB: not implemented"
	return
}

func (s *egressIPScheduler) GetEgressIPAndNode(egress string) (string, string, error, bool) {
	_ = "STUB: not implemented"
	return "", "", nil, false
}

// EgressesByCreationTimestamp sorts a list of Egresses by creation timestamp.
type EgressesByCreationTimestamp []*crdv1b1.Egress

func (o EgressesByCreationTimestamp) Len() int           { _ = "STUB: not implemented"; return 0 }
func (o EgressesByCreationTimestamp) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (o EgressesByCreationTimestamp) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// updateMaxEgressIPsByNode updates the maxEgressIPs for a given Node in the cache.
// It returns whether there is a real change, which indicates if rescheduling is required.
func (s *egressIPScheduler) updateMaxEgressIPsByNode(nodeName string, maxEgressIPs int) bool {
	_ = "STUB: not implemented"
	return false
}

// If the value equals to the default value, no need to cache it and trigger rescheduling.

// deleteMaxEgressIPsByNode deletes the maxEgressIPs for a given Node in the cache.
// It returns whether there is a real change, which indicates if rescheduling is required.
func (s *egressIPScheduler) deleteMaxEgressIPsByNode(nodeName string) bool {
	_ = "STUB: not implemented"
	return false
}

// getMaxEgressIPsByNode gets the maxEgressIPs for a given Node.
// If there isn't a value for the Node, the default value will be returned.
func (s *egressIPScheduler) getMaxEgressIPsByNode(nodeName string) int {
	_ = "STUB: not implemented"
	return 0
}

// schedule takes the spec of Egress and ExternalIPPool and the state of memberlist cluster as inputs, generates
// scheduling results deterministically. When every Node's capacity is sufficient, each Egress's schedule is independent
// and is only determined by the consistent hash map. When any Node's capacity is insufficient, one Egress's schedule
// may be affected by Egresses created before it. It will be triggerred when any schedulable Egress changes or the state
// of memberlist cluster changes, and will notify Egress schedule event subscribers of Egresses that are rescheduled.
//
// Note that it's possible that different agents decide different IP - Node assignment because their caches of Egress or
// the states of memberlist cluster are inconsistent at a moment. But all agents should get the same schedule results
// and correct IP assignment when their caches converge.
func (s *egressIPScheduler) schedule() { _ = "STUB: not implemented"; return }

// Sort Egresses by creation timestamp to make the result deterministic and prioritize objected created earlier
// when the total capacity is insufficient.

// Ignore Egresses that shouldn't be scheduled.

// Count the Egress IPs that are already assigned to this Node.

// Check if this Node can accommodate the new Egress IP.

// Store error in its result to differentiate scheduling error from unprocessed case.

// Identify Egresses whose schedule results are updated.

// Record the new results.
