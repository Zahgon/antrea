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

package networkpolicy

import (
	"sync"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/apiserver/storage"
	antreaclientset "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
)

const (
	statusControllerName = "NetworkPolicyStatusController"
)

var (
	// maxConditionMessageLength defines the max length of the message field in one Condition. If the actual message
	// length is over size, truncate the string and use "..." in the end.
	// Use a variable for test.
	maxConditionMessageLength = 256
)

// StatusController is responsible for synchronizing the status of Antrea ClusterNetworkPolicy and Antrea NetworkPolicy.
type StatusController struct {
	// npControlInterface knows how to update Antrea NetworkPolicy status.
	npControlInterface networkPolicyControlInterface

	// queue maintains the keys of the NetworkPolicy objects that need to be synced.
	queue workqueue.TypedRateLimitingInterface[string]

	// internalNetworkPolicyStore is the storage where the populated internal Network Policy are stored.
	internalNetworkPolicyStore storage.Interface

	// statuses is a nested map that keeps the realization statuses reported by antrea-agents.
	// The outer map's keys are the NetworkPolicy keys. The inner map's keys are the Node names. The inner map's values
	// are statuses reported by each Node for a NetworkPolicy.
	statuses     map[string]map[string]*controlplane.NetworkPolicyNodeStatus
	statusesLock sync.RWMutex

	// acnpListerSynced is a function which returns true if the ClusterNetworkPolicies shared informer has been synced at least once.
	acnpListerSynced cache.InformerSynced
	// annpListerSynced is a function which returns true if the AntreaNetworkPolicies shared informer has been synced at least once.
	annpListerSynced cache.InformerSynced
}

func NewStatusController(antreaClient antreaclientset.Interface, internalNetworkPolicyStore storage.Interface, acnpInformer crdinformers.ClusterNetworkPolicyInformer, annpInformer crdinformers.NetworkPolicyInformer) *StatusController {
	_ = "STUB: not implemented"
	return nil
}

// To save a "GET" query before each update, UpdateAntreaClusterNetworkPolicyStatus treats the cache of Lister as
// the state of kube-apiserver. In some cases the cache may not be in sync, then we might skip updating a policy's
// status by mistake. To resolve it, add update event handlers which trigger resync of a policy if its status is
// updated. This could also ensure we can reconcile a policy's status if it's updated by other clients by accident.
// However, a normal update made by the controller itself will trigger resync as well, which could lead to duplicate
// computation.
// TODO: Evaluate if we can avoid the duplicate computation by comparing the updated status with some internal state.

func (c *StatusController) updateACNP(old, cur interface{}) { _ = "STUB: not implemented"; return }

func (c *StatusController) updateANNP(old, cur interface{}) { _ = "STUB: not implemented"; return }

func (c *StatusController) UpdateStatus(status *controlplane.NetworkPolicyStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *StatusController) getNodeStatuses(key string) []*controlplane.NetworkPolicyNodeStatus {
	_ = "STUB: not implemented"
	return nil
}

func (c *StatusController) clearStatuses(key string) { _ = "STUB: not implemented"; return }

func (c *StatusController) deleteNodeStatus(key string, nodeName string) {
	_ = "STUB: not implemented"
	return
}

// Run begins watching and syncing of a StatusController.
func (c *StatusController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *StatusController) watchInternalNetworkPolicy() { _ = "STUB: not implemented"; return }

// Skip handling Bookmark events.

func (c *StatusController) runWorker() { _ = "STUB: not implemented"; return }

// processNextWorkItem deals with one key off the queue.  It returns false when it's time to quit.
func (c *StatusController) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// syncHandler calculates the NetworkPolicy status based on the desired state from the internalNetworkPolicyStore and
// the actual state from the statuses map, and syncs it with the Kubernetes API.
// Each status update from agents can trigger syncHandler, however, the status updates' arrival time should not differ
// too much so some of them can be merged in workqueue. Besides, there are a limited number of workers. If there are
// many policies need to sync, some policies will have to wait in workqueue, during which their status updates can be
// merged. Therefore, it shouldn't happen that each status update leads to one CR update.
func (c *StatusController) syncHandler(key string) error { _ = "STUB: not implemented"; return nil }

// It has been deleted, cleaning its statuses.

// It means the NetworkPolicy has been processed, and marked as unrealizable. It will enter unrealizable phase
// instead of being further realized. Antrea-agents will not process further.

// It means the NetworkPolicy hasn't been processed once. Set it to Pending to differentiate from NetworkPolicies
// that spans 0 Node.

// The node is no longer in the span of this policy, delete its status.

// networkPolicyControlInterface is an interface that knows how to update Antrea NetworkPolicy status.
// It's created as an interface to allow testing.
type networkPolicyControlInterface interface {
	UpdateAntreaNetworkPolicyStatus(namespace, name string, status *crdv1beta1.NetworkPolicyStatus) error
	UpdateAntreaClusterNetworkPolicyStatus(name string, status *crdv1beta1.NetworkPolicyStatus) error
}

type networkPolicyControl struct {
	antreaClient antreaclientset.Interface
	acnpLister   crdlisters.ClusterNetworkPolicyLister
	annpLister   crdlisters.NetworkPolicyLister
}

func (c *networkPolicyControl) UpdateAntreaNetworkPolicyStatus(namespace, name string, status *crdv1beta1.NetworkPolicyStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// Return the error from UPDATE.

func (c *networkPolicyControl) UpdateAntreaClusterNetworkPolicyStatus(name string, status *crdv1beta1.NetworkPolicyStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// If the current status equals to the desired status, no need to update.

// Return the error from UPDATE.

// GenerateNetworkPolicyCondition generates conditions based on the given error type.
// Error of nil type means the NetworkPolicyCondition status is True.
// Supports ErrNetworkPolicyAppliedToUnsupportedGroup error.
func GenerateNetworkPolicyCondition(err error) []crdv1beta1.NetworkPolicyCondition {
	_ = "STUB: not implemented"
	return nil
}
