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
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/client"
	"antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
)

const (
	realizedRulePolicyIndex = "policy"
)

// StatusManager keeps track of the realized NetworkPolicy rules. It syncs the status of a NetworkPolicy to the
// antrea-controller once it is realized. A policy is considered realized when all of its desired rules have been
// realized and all of its undesired rules have been removed.
// For each new policy, SetRuleRealization is supposed to be called for each of its desired rules while
// DeleteRuleRealization is supposed to be called for the removed rules.
type StatusManager interface {
	// SetRuleRealization updates the actual status for the given NetworkPolicy rule.
	SetRuleRealization(ruleID string, policyID types.UID)
	// DeleteRuleRealization deletes the actual status for the given NetworkPolicy rule.
	DeleteRuleRealization(ruleID string)
	// Resync triggers syncing status with the antrea-controller for the given NetworkPolicy.
	Resync(policyID types.UID)
	// Start the status sync loop.
	Run(stopCh <-chan struct{})
}

// StatusController implements StatusManager.
type StatusController struct {
	nodeName string
	// statusControlInterface knows how to update control plane NetworkPolicy status.
	statusControlInterface networkPolicyStatusControlInterface
	// ruleCache provides the desired state of NetworkPolicy rules.
	ruleCache *ruleCache
	// realizedRules keeps track of the realized NetworkPolicy rules.
	realizedRules cache.Indexer
	// queue maintains the UIDs of the NetworkPolicy that need to be processed.
	queue workqueue.TypedRateLimitingInterface[types.UID]
}

// realizedRule is the struct kept by StatusController for storing a realized rule.
// It has policyID because "ruleCache" only keeps desired state of policies, so if a rule is no longer in a policy it
// will be deleted immediately from "ruleCache" while we need to know these rules are actually uninstalled from
// dataplane before their policies are considered realized.
type realizedRule struct {
	ruleID   string
	policyID types.UID
}

func realizedRuleKeyFunc(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func realizedRulePolicyIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newStatusController(antreaClientProvider client.AntreaClientProvider, nodeName string, ruleCache *ruleCache) *StatusController {
	_ = "STUB: not implemented"
	return nil
}

func (c *StatusController) SetRuleRealization(ruleID string, policyID types.UID) {
	_ = "STUB: not implemented"
	return
}

// This rule has been realized before. The current call must be triggered by group member updates, which doesn't
// affect the policy's realization status.

func (c *StatusController) DeleteRuleRealization(ruleID string) { _ = "STUB: not implemented"; return }

// This rule hasn't been realized before, so it doesn't affect the policy's realization status.

func (c *StatusController) Resync(policyID types.UID) { _ = "STUB: not implemented"; return }

// worker is a long-running function that will continually call the processNextWorkItem function in
// order to read and process a message on the workqueue.
func (c *StatusController) worker() { _ = "STUB: not implemented"; return }

func (c *StatusController) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// We call Done here so the workqueue knows we have finished processing this item. We also
// must remember to call Forget if we do not want this work item being re-queued. For
// example, we do not call Forget if a transient error occurs, instead the item is put back
// on the workqueue and attempted again after a back-off period.

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

func (c *StatusController) syncHandler(uid types.UID) error { _ = "STUB: not implemented"; return nil }

// The policy must have been deleted, no further processing.

// The policy must have been deleted, no further processing.

// desiredRules should match actualRules exactly.

// At this point, all desired rules have been realized and all undesired rules have been removed, report it to the antrea-controller.

func (c *StatusController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// networkPolicyStatusControlInterface is an interface that knows how to get and update control plane NetworkPolicy status.
// It's created as an interface to allow testing.
type networkPolicyStatusControlInterface interface {
	UpdateNetworkPolicyStatus(name string, status *v1beta2.NetworkPolicyStatus) error
}

type networkPolicyStatusControl struct {
	antreaClientProvider client.AntreaClientProvider
}

func (c *networkPolicyStatusControl) UpdateNetworkPolicyStatus(name string, status *v1beta2.NetworkPolicyStatus) error {
	_ = "STUB: not implemented"
	return nil
}
