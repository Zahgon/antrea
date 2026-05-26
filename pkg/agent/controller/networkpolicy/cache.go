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
	// #nosec G505: not used for security purposes

	"sync"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"

	"antrea.io/antrea/v2/pkg/agent/config"
	v1beta "antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/querier"
	"antrea.io/antrea/v2/pkg/util/channel"
)

const (
	ruleIDLength                  = 16
	appliedToGroupIndex           = "appliedToGroup"
	addressGroupIndex             = "addressGroup"
	policyIndex                   = "policy"
	toServicesIndex               = "toServices"
	toIGMPReportGroupAddressIndex = "toIGMPReportGroupAddress"
)

// rule is the struct stored in ruleCache, it contains necessary information
// to construct a complete rule that can be used by reconciler to enforce.
// The K8s NetworkPolicy object doesn't provide ID for its rule, here we
// calculate an ID based on the rule's fields. That means:
//  1. If a rule's selector/services/direction changes, it becomes "another" rule.
//  2. If inserting rules before a rule or shuffling rules in a NetworkPolicy, we
//     can know the existing rules don't change and skip processing them. Note that
//     if an ACNP/ANNP rule's position (from top down) within a networkpolicy changes,
//     it affects the Priority of the rule.
type rule struct {
	// ID is calculated from the hash value of all other fields.
	ID string
	// Direction of this rule.
	Direction v1beta.Direction
	// Source Address of this rule, can't coexist with To.
	From v1beta.NetworkPolicyPeer
	// Destination Address of this rule, can't coexist with From.
	To v1beta.NetworkPolicyPeer
	// Protocols and Ports of this rule.
	Services []v1beta.Service
	// Layer 7 protocols of this rule.
	L7Protocols []v1beta.L7Protocol
	// Name of this rule. Empty for k8s NetworkPolicy.
	Name string
	// Action of this rule. nil for k8s NetworkPolicy.
	Action *crdv1beta1.RuleAction
	// Priority of this rule within the NetworkPolicy. Defaults to -1 for K8s NetworkPolicy.
	Priority int32
	// The highest rule Priority within the NetworkPolicy. Defaults to -1 for K8s NetworkPolicy.
	MaxPriority int32
	// Priority of the NetworkPolicy to which this rule belong. nil for K8s NetworkPolicy.
	PolicyPriority *float64
	// Priority of the tier that the NetworkPolicy belongs to. nil for K8s NetworkPolicy.
	TierPriority *int32
	// Targets of this rule.
	AppliedToGroups []string
	// The parent Policy ID. Used to identify rules belong to a specified
	// policy for deletion.
	PolicyUID types.UID
	// The metadata of parent Policy. Used to associate the rule with Policy
	// for troubleshooting purpose (logging and CLI).
	PolicyName string
	// Reference to the original NetworkPolicy that the rule belongs to.
	// Note it has a different meaning from PolicyUID, PolicyName, and
	// PolicyNamespace which are the metadata fields of the corresponding
	// controlplane NetworkPolicy. Although they are same for now, it might
	// change in the future, features that need the information of the original
	// NetworkPolicy should use SourceRef.
	SourceRef *v1beta.NetworkPolicyReference
	// EnableLogging is a boolean indicating whether logging is required for Antrea Policies. Always false for K8s NetworkPolicy.
	EnableLogging bool
	// LogLabel is a string associated to the NetworkPolicy rule. Used for logging.
	LogLabel string
}

func (r *rule) Less(r2 *rule) bool {
	_ = "STUB: not implemented"
	// priorities for rule r
	return false
}

// priorities for rule r2

// compare two rules' priorities

// hashRule calculates a string based on the rule's content.
func hashRule(r *rule) string {
	_ = "STUB: not implemented"
	// #nosec G401: not used for security purposes
	return ""
}

// CompletedRule contains IPAddresses and Pods flattened from AddressGroups and AppliedToGroups.
// It's the struct used by reconciler.
type CompletedRule struct {
	*rule
	// Source GroupMembers of this rule, can't coexist with ToAddresses.
	FromAddresses v1beta.GroupMemberSet
	// Destination GroupMembers of this rule, can't coexist with FromAddresses.
	ToAddresses v1beta.GroupMemberSet
	// Target GroupMembers of this rule.
	TargetMembers v1beta.GroupMemberSet
	// Vlan ID allocated for this rule if this rule is for L7 NetworkPolicy.
	L7RuleVlanID *uint32
}

// String returns the string representation of the CompletedRule.
func (r *CompletedRule) String() string { _ = "STUB: not implemented"; return "" }

// isAntreaNetworkPolicyRule returns true if the rule is part of a Antrea policy.
func (r *CompletedRule) isAntreaNetworkPolicyRule() bool { _ = "STUB: not implemented"; return false }

func (r *CompletedRule) isIGMPEgressPolicyRule() bool { _ = "STUB: not implemented"; return false }

func (r *CompletedRule) isNodeNetworkPolicyRule() bool { _ = "STUB: not implemented"; return false }

// ruleCache caches Antrea AddressGroups, AppliedToGroups and NetworkPolicies,
// can construct complete rules that can be used by reconciler to enforce.
type ruleCache struct {
	appliedToSetLock sync.RWMutex
	// appliedToSetByGroup stores the AppliedToGroup members.
	// It is a mapping from group name to a set of GroupMembers.
	appliedToSetByGroup map[string]v1beta.GroupMemberSet

	addressSetLock sync.RWMutex
	// addressSetByGroup stores the AddressGroup members.
	// It is a mapping from group name to a set of GroupMembers.
	addressSetByGroup map[string]v1beta.GroupMemberSet

	policyMapLock sync.RWMutex
	// policyMap is a map using NetworkPolicy UID as the key.
	// TODO: reduce its storage redundancy with rules.
	policyMap map[string]*v1beta.NetworkPolicy

	// rules is a storage that supports listing rules using multiple indexing functions.
	// rules is thread-safe.
	rules cache.Indexer
	// dirtyRuleHandler is a callback that is run upon finding a rule out-of-sync.
	dirtyRuleHandler func(string)

	// groupIDUpdates is a channel for receiving groupID for Service is assigned
	// or released events from groupCounters.
	groupIDUpdates <-chan string
}

func (c *ruleCache) getNetworkPolicies(npFilter *querier.NetworkPolicyQueryFilter) []v1beta.NetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}

// networkPolicyMatchFilter returns true if the provided NetworkPolicy matches the provided NetworkPolicyQueryFilter.
func (c *ruleCache) networkPolicyMatchFilter(npFilter *querier.NetworkPolicyQueryFilter, np *v1beta.NetworkPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ruleCache) getNetworkPolicy(uid string) *v1beta.NetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (c *ruleCache) getAppliedNetworkPolicies(pod, namespace string, npFilter *querier.NetworkPolicyQueryFilter) []v1beta.NetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}

// The Policy might be removed during the query.

func (c *ruleCache) getEffectiveRulesByNetworkPolicy(uid string) []*rule {
	_ = "STUB: not implemented"
	return nil
}

// A rule is considered effective when any of its AppliedToGroups can be populated.

func (c *ruleCache) GetAddressGroups() []v1beta.AddressGroup { _ = "STUB: not implemented"; return nil }

func (c *ruleCache) GetAppliedToGroups() []v1beta.AppliedToGroup {
	_ = "STUB: not implemented"
	return nil
}

// ruleKeyFunc knows how to get key of a *rule.
func ruleKeyFunc(obj interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

// addressGroupIndexFunc knows how to get addressGroups of a *rule.
// It's provided to cache.Indexer to build an index of addressGroups.
func addressGroupIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// appliedToGroupIndexFunc knows how to get appliedToGroups of a *rule.
// It's provided to cache.Indexer to build an index of appliedToGroups.
func appliedToGroupIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// policyIndexFunc knows how to get NetworkPolicy UID of a *rule.
// It's provided to cache.Indexer to build an index of NetworkPolicy.
func policyIndexFunc(obj interface{}) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// toServicesIndexFunc knows how to get NamespacedNames of Services referred in
// ToServices field of a *rule. It's provided to cache.Indexer to build an index of
// NetworkPolicy.
func toServicesIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// toIGMPReportGroupAddressIndexFunc knows how to get IGMP report groupAddresses of a *rule
// It's provided to cache.Indexer to build an index of NetworkPolicy.
func toIGMPReportGroupAddressIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newRuleCache returns a new *ruleCache.
func newRuleCache(dirtyRuleHandler func(string), podUpdateSubscriber channel.Subscriber, externalEntityUpdateSubscriber channel.Subscriber,
	serviceGroupIDUpdate <-chan string, nodeType config.NodeType) *ruleCache {
	_ = "STUB: not implemented"
	return nil
}

// Subscribe Pod update events from CNIServer.

// Subscribe ExternalEntity update events from ExternalNodeController

// processPodUpdate will be called when CNIServer publishes a Pod update event.
// It finds out AppliedToGroups that contain this Pod and triggers reconciliation
// of related rules.
// It can enforce NetworkPolicies to newly added Pods right after CNI ADD is
// done if antrea-controller has computed the Pods' policies and propagated
// them to this Node by their labels and NodeName, instead of waiting for their
// IPs are reported to kube-apiserver and processed by antrea-controller.
func (c *ruleCache) processPodUpdate(e interface{}) { _ = "STUB: not implemented"; return }

// processExternalEntityUpdate will be called when ExternalNodeController publishes an ExternalEntity update event.
// It finds out AppliedToGroups that contain this ExternalNode converted ExternalEntity and triggers reconciliation
// of related rules.
// It can enforce NetworkPolicies to ExternalEntities after ExternalEntityInterface is realised in the interface store.
func (c *ruleCache) processExternalEntityUpdate(e interface{}) { _ = "STUB: not implemented"; return }

// processGroupIDUpdates is an infinite loop that takes Service groupID
// update events from the channel, finds out rules that refer this Service in
// ToServices field and use dirtyRuleHandler to re-queue these rules.
func (c *ruleCache) processGroupIDUpdates() { _ = "STUB: not implemented"; return }

// GetAddressGroupNum gets the number of AddressGroup.
func (c *ruleCache) GetAddressGroupNum() int { _ = "STUB: not implemented"; return 0 }

// ReplaceAddressGroups atomically adds the given groups to the cache and deletes
// the pre-existing groups that are not in the given groups from the cache.
// It makes the cache in sync with the apiserver when restarting a watch.
func (c *ruleCache) ReplaceAddressGroups(groups []*v1beta.AddressGroup) {
	_ = "STUB: not implemented"
	return
}

// AddAddressGroup adds a new *v1beta.AddressGroup to the cache. The rules
// referencing it will be regarded as dirty.
// It's safe to add an AddressGroup multiple times as it only overrides the
// map, this could happen when the watcher reconnects to the Apiserver.
func (c *ruleCache) AddAddressGroup(group *v1beta.AddressGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ruleCache) addAddressGroupLocked(group *v1beta.AddressGroup) error {
	_ = "STUB: not implemented"
	return nil
}

// Must not store address of loop iterator variable as it's the same
// address taking different values in each loop iteration, otherwise
// groupMemberSet would eventually contain only the last value.
// https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable

// PatchAddressGroup updates a cached *v1beta.AddressGroup.
// The rules referencing it will be regarded as dirty.
// It returns a copy of the patched AddressGroup, or an error if the AddressGroup doesn't exist.
func (c *ruleCache) PatchAddressGroup(patch *v1beta.AddressGroupPatch) (*v1beta.AddressGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteAddressGroup deletes a cached *v1beta.AddressGroup.
// It should only happen when a group is no longer referenced by any rule, so
// no need to mark dirty rules.
func (c *ruleCache) DeleteAddressGroup(group *v1beta.AddressGroup) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAppliedToGroupNum gets the number of AppliedToGroup.
func (c *ruleCache) GetAppliedToGroupNum() int { _ = "STUB: not implemented"; return 0 }

// ReplaceAppliedToGroups atomically adds the given groups to the cache and deletes
// the pre-existing groups that are not in the given groups from the cache.
// It makes the cache in sync with the apiserver when restarting a watch.
func (c *ruleCache) ReplaceAppliedToGroups(groups []*v1beta.AppliedToGroup) {
	_ = "STUB: not implemented"
	return
}

// AddAppliedToGroup adds a new *v1beta.AppliedToGroup to the cache. The rules
// referencing it will be regarded as dirty.
// It's safe to add an AppliedToGroup multiple times as it only overrides the
// map, this could happen when the watcher reconnects to the Apiserver.
func (c *ruleCache) AddAppliedToGroup(group *v1beta.AppliedToGroup) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ruleCache) addAppliedToGroupLocked(group *v1beta.AppliedToGroup) error {
	_ = "STUB: not implemented"
	return nil
}

// PatchAppliedToGroup updates a cached *v1beta.AppliedToGroupPatch.
// The rules referencing it will be regarded as dirty.
// It returns a copy of the patched AppliedToGroup, or an error if the AppliedToGroup doesn't exist.
func (c *ruleCache) PatchAppliedToGroup(patch *v1beta.AppliedToGroupPatch) (*v1beta.AppliedToGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteAppliedToGroup deletes a cached *v1beta.AppliedToGroup.
// It may be called when a rule becomes ineffective, so it needs to mark dirty rules.
func (c *ruleCache) DeleteAppliedToGroup(group *v1beta.AppliedToGroup) error {
	_ = "STUB: not implemented"
	return nil
}

// toRule converts v1beta.NetworkPolicyRule to *rule.
func toRule(r *v1beta.NetworkPolicyRule, policy *v1beta.NetworkPolicy, maxPriority int32) *rule {
	_ = "STUB: not implemented"
	return nil
}

// getMaxPriority returns the highest rule priority for v1beta.NetworkPolicy that is created
// by Antrea-native policies. For K8s NetworkPolicies, it always returns -1.
func getMaxPriority(policy *v1beta.NetworkPolicy) int32 { _ = "STUB: not implemented"; return 0 }

// GetNetworkPolicyNum gets the number of NetworkPolicy.
func (c *ruleCache) GetNetworkPolicyNum() int { _ = "STUB: not implemented"; return 0 }

// ReplaceNetworkPolicies atomically adds the given policies to the cache and deletes
// the pre-existing policies that are not in the given policies from the cache.
// It makes the cache in sync with the apiserver when restarting a watch.
func (c *ruleCache) ReplaceNetworkPolicies(policies []*v1beta.NetworkPolicy) {
	_ = "STUB: not implemented"
	return
}

// AddNetworkPolicy adds a new *v1beta.NetworkPolicy to the cache.
// It could happen that an existing NetworkPolicy is "added" again when the
// watcher reconnects to the Apiserver, we use the same processing as
// UpdateNetworkPolicy to ensure orphan rules are removed.
func (c *ruleCache) AddNetworkPolicy(policy *v1beta.NetworkPolicy) {
	_ = "STUB: not implemented"
	return
}

// UpdateNetworkPolicy updates a cached *v1beta.NetworkPolicy and returns whether any rule or the generation changes.
// The added rules and removed rules will be regarded as dirty.
func (c *ruleCache) UpdateNetworkPolicy(policy *v1beta.NetworkPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

// updateNetworkPolicyLocked returns whether any rule or the generation changes.
func (c *ruleCache) updateNetworkPolicyLocked(policy *v1beta.NetworkPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

// If rule already exists, remove it from the map so the ones left are orphaned,
// which means those rules need to be handled by dirtyRuleHandler.

// If rule doesn't exist, add it to cache and mark it as dirty.

// Count up antrea_agent_ingress_networkpolicy_rule_count or antrea_agent_egress_networkpolicy_rule_count

// At this moment, the remaining rules are orphaned, remove them from store and mark them as dirty.

// Count down antrea_agent_ingress_networkpolicy_rule_count or antrea_agent_egress_networkpolicy_rule_count

// DeleteNetworkPolicy deletes a cached *v1beta.NetworkPolicy.
// All its rules will be regarded as dirty.
func (c *ruleCache) DeleteNetworkPolicy(policy *v1beta.NetworkPolicy) {
	_ = "STUB: not implemented"
	return
}

func (c *ruleCache) deleteNetworkPolicyLocked(uid string) { _ = "STUB: not implemented"; return }

// Count down antrea_agent_ingress_networkpolicy_rule_count or antrea_agent_egress_networkpolicy_rule_count

// GetCompletedRule constructs a *CompletedRule for the provided ruleID.
// If the rule is not effective or not realizable due to missing group data, the return value will indicate it.
// A rule is considered effective when any of its AppliedToGroups can be populated.
// A rule is considered realizable when it's effective and all of its AddressGroups can be populated.
// When a rule is not effective, it should be removed from the datapath.
// When a rule is effective but not realizable, the caller should wait for it being realizable before doing anything.
// When a rule is effective and realizable, the caller should realize it.
// This is because some AppliedToGroups in a rule might never be sent to this Node if one of the following is true:
//   - The original policy has multiple AppliedToGroups and some AppliedToGroups' span does not include this Node.
//   - The original policy is appliedTo-per-rule, and some of the rule's AppliedToGroups do not include this Node.
//   - The original policy is appliedTo-per-rule, none of the rule's AppliedToGroups includes this Node, but some other rules' (in the same policy) AppliedToGroups include this Node.
//
// In these cases, it is not guaranteed that all AppliedToGroups in the rule will eventually be present in the cache.
// Only the AppliedToGroups whose span includes this Node will eventually be received.
func (c *ruleCache) GetCompletedRule(ruleID string) (completedRule *CompletedRule, effective bool, realizable bool) {
	_ = "STUB: not implemented"
	return nil, false, false
}

// onAppliedToGroupUpdate gets rules referencing to the provided AppliedToGroup
// and mark them as dirty.
func (c *ruleCache) onAppliedToGroupUpdate(groupName string) { _ = "STUB: not implemented"; return }

// onAddressGroupUpdate gets rules referencing to the provided AddressGroup
// and mark them as dirty.
func (c *ruleCache) onAddressGroupUpdate(groupName string) { _ = "STUB: not implemented"; return }

// unionAddressGroups gets the union of addresses of the provided address groups.
// If any group is not found, nil and false will be returned to indicate the
// set is not complete yet.
func (c *ruleCache) unionAddressGroups(groupNames []string) (v1beta.GroupMemberSet, bool) {
	_ = "STUB: not implemented"
	return *new(v1beta.GroupMemberSet), false
}

// unionAppliedToGroups gets the union of pods of the provided appliedTo groups.
// If any group is found, the union and true will be returned. Otherwise an empty set and false will be returned.
func (c *ruleCache) unionAppliedToGroups(groupNames []string) (v1beta.GroupMemberSet, bool) {
	_ = "STUB: not implemented"
	return *new(v1beta.GroupMemberSet), false
}

// processServiceGroupIDUpdate gets names of AppliedToGroup by Service NamespacedName.
func (c *ruleCache) processServiceGroupIDUpdate(svcStr string) { _ = "STUB: not implemented"; return }

// Reprocess rules if the Service referred by this rule's ToServices has updated.

// Reprocess rules if the Service referred by rule's AppliedToGroup has updated.
