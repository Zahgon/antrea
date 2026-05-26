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
	"net"
	"sync"

	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	proxytypes "antrea.io/antrea/v2/pkg/agent/proxy/types"
	"antrea.io/antrea/v2/pkg/agent/types"
	"antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
)

var (
	baselineTierPriority int32 = 253
	banpTierPriority     int32 = 254
)

type ruleType int

const (
	unicast   ruleType = 0
	igmp      ruleType = 1
	multicast ruleType = 2

	igmpServicesKey = "igmp-services-key"
)

// Reconciler is an interface that knows how to reconcile the desired state of
// CompletedRule with the actual state of Openflow entries.
type Reconciler interface {
	// Reconcile reconciles the desired state of the provided CompletedRule
	// with the actual state of Openflow entries.
	Reconcile(rule *CompletedRule) error

	// BatchReconcile reconciles the desired state of the provided CompletedRules
	// with the actual state of Openflow entries in batch. It should only be invoked
	// if all rules are newly added without last realized status.
	BatchReconcile(rules []*CompletedRule) error

	// Forget cleanups the actual state of Openflow entries of the specified ruleID.
	Forget(ruleID string) error

	// GetRuleByFlowID returns the rule from the async rule cache in idAllocator cache.
	GetRuleByFlowID(ruleID uint32) (*types.PolicyRule, bool, error)

	// RunIDAllocatorWorker runs the worker that deletes the rules from the cache
	// in idAllocator.
	RunIDAllocatorWorker(stopCh <-chan struct{})
}

// servicesKey is used to identify Services based on their numbered ports.
type servicesKey string

// normalizeServices calculates the servicesKey of the provided services based
// on their numbered ports.
// It constructs the string with strings.Builder instead of using the spew
// library for efficiency consideration as this function is called quite often.
// It ignores the difference of protocol and non resolved ports because the
// servicesKey is only used to distinguish the results obtained by resolving
// named ports for the same services.
func normalizeServices(services []v1beta2.Service) servicesKey {
	_ = "STUB: not implemented"
	return *new(servicesKey)
}

// podPolicyLastRealized is the struct cached by podReconciler. It's used to track the
// actual state of rules we have enforced, so that we can know how to reconcile
// a rule when it's updated/removed.
// It includes the last version of CompletedRule the podReconciler has realized
// and the related runtime information including the ofIDs, the Openflow ports
// or the IP addresses of the target Pods got from the InterfaceStore.
//
// Note that a policy rule can be split into multiple Openflow rules based on
// the named port resolving result. Pods that have same port numbers for the
// port names defined in the rule will share an Openflow rule. For example,
// if an ingress rule applies to 3 Pods like below:
//
// NetworkPolicy rule:
//
//	spec:
//	 ingress:
//	 - from:
//	   - namespaceSelector: {}
//	   ports:
//	   - port: http
//	     protocol: TCP
//
// Pod A and Pod B:
//
//	spec:
//	  containers:
//	  - ports:
//	    - containerPort: 80
//	      name: http
//	      protocol: TCP
//
// Pod C:
//
//	spec:
//	  containers:
//	  - ports:
//	    - containerPort: 8080
//	      name: http
//	      protocol: TCP
//
// Then Pod A and B will share an Openflow rule as both of them resolve "http" to 80,
// while Pod C will have another Openflow rule as it resolves "http" to 8080.
// In the implementation, we group Pods by their resolved services value so Pod A and B
// can be mapped to same group.
type podPolicyLastRealized struct {
	// ofIDs identifies Openflow rules in Openflow implementation.
	// It's a map of servicesKey to Openflow rule ID.
	ofIDs map[servicesKey]uint32
	// The desired state of a policy rule.
	*CompletedRule
	// The OFPort set we have realized for target Pods. We need to record them
	// because this info will be removed from InterfaceStore after CNI DEL, we
	// can't know which OFPort to delete when deleting a Pod from the rule. So
	// we compare the last realized OFPorts and the new desired one to identify
	// difference, which could also cover the stale OFPorts produced by the case
	// that Kubelet calls CNI ADD for a Pod more than once due to non CNI issues.
	// It's only used for ingress rule as its "to" addresses.
	// It's grouped by servicesKey, mapping to multiple Openflow rules.
	podOFPorts map[servicesKey]sets.Set[int32]
	// The IP set we have realized for target Pods. Same as podOFPorts.
	// It's only used for egress rule as its "from" addresses.
	// It's same in all Openflow rules, because named port is only for
	// destination Pods.
	podIPs sets.Set[string]
	// fqdnIPaddresses tracks the last realized set of IP addresses resolved for
	// the fqdn selector of this policy rule. It must be empty for policy rule
	// that is not egress and does not have toFQDN field.
	fqdnIPAddresses sets.Set[string]
	// serviceGroupIDs tracks the last realized set of groupIDs resolved for the
	// toServices of this policy rule or services of TargetMember of this policy rule.
	// It must be empty for policy rule that is neither an egress rule with toServices
	// field nor an ingress rule that is applied to Services.
	serviceGroupIDs sets.Set[int64]
	// groupAddresses track the latest realized set of multicast groups for the multicast traffic
	groupAddresses sets.Set[string]
}

func newPodPolicyLastRealized(rule *CompletedRule) *podPolicyLastRealized {
	_ = "STUB: not implemented"
	return nil
}

// tablePriorityAssigner groups the priorityAssigner and mutex for a single OVS table
// that is reserved for installing Antrea policy rules.
type tablePriorityAssigner struct {
	assigner *priorityAssigner
	mutex    sync.RWMutex
}

// podReconciler implements Reconciler.
// Note that although its Reconcile and Forget methods are thread-safe, it's
// assumed each rule can only be processed by a single client at any given
// time. Different rules can be processed in parallel.
type podReconciler struct {
	// ofClient is the Openflow interface.
	ofClient openflow.Client

	// ifaceStore provides container interface OFPort and IP information.
	ifaceStore interfacestore.InterfaceStore

	// lastRealizeds caches the last realized rules.
	// It's a mapping from ruleID to *podPolicyLastRealized.
	lastRealizeds sync.Map

	// idAllocator provides interfaces to allocateForRule and release uint32 id.
	idAllocator *idAllocator

	// priorityAssigners provides interfaces to manage OF priorities for each OVS table.
	priorityAssigners map[uint8]*tablePriorityAssigner
	// ipv4Enabled tells if IPv4 is supported on this Node or not.
	ipv4Enabled bool
	// ipv6Enabled tells is IPv6 is supported on this Node or not.
	ipv6Enabled bool

	// fqdnController manages dns cache of FQDN rules. It provides interfaces for the
	// podReconciler to register FQDN policy rules and query the IP addresses corresponded
	// to a FQDN.
	fqdnController *fqdnController

	// groupCounters is a list of GroupCounter for v4 and v6 env. podReconciler uses these
	// GroupCounters to get the groupIDs of a specific Service.
	groupCounters []proxytypes.GroupCounter

	// multicastEnabled indicates whether multicast is enabled
	multicastEnabled bool
}

// newPodReconciler returns a new *podReconciler.
func newPodReconciler(ofClient openflow.Client,
	ifaceStore interfacestore.InterfaceStore,
	idAllocator *idAllocator,
	fqdnController *fqdnController,
	groupCounters []proxytypes.GroupCounter,
	v4Enabled bool,
	v6Enabled bool,
	antreaPolicyEnabled bool,
	multicastEnabled bool,
) *podReconciler {
	_ = "STUB: not implemented"
	return nil
}

// Check if ofClient is nil or not to be compatible with unit tests.

// RunIDAllocatorWorker runs the worker that deletes the rules from the cache in
// idAllocator.
func (r *podReconciler) RunIDAllocatorWorker(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// Reconcile checks whether the provided rule has been enforced or not, and
// invoke the add or update method accordingly.
func (r *podReconciler) Reconcile(rule *CompletedRule) error { _ = "STUB: not implemented"; return nil }

// IGMP Egress policy is enforced in userspace via packet-in message, there won't be OpenFlow
// rules created for such rules. Therefore, assigning priority is not required.

// For CNP, only release priorityMutex after rule is installed on OVS. Otherwise,
// priority re-assignments for flows that have already been assigned priorities but
// not yet installed on OVS will be missed.

func (r *podReconciler) getRuleType(rule *CompletedRule) ruleType {
	_ = "STUB: not implemented"
	return *new(ruleType)
}

// getOFRuleTable retrieves the OpenFlow table to install the CompletedRule.
// The decision is made based on whether the rule is created for an ACNP/ANNP, and
// the Tier of that NetworkPolicy.
func (r *podReconciler) getOFRuleTable(rule *CompletedRule) uint8 {
	_ = "STUB: not implemented"
	return 0
}

// Multicast NetworkPolicy only supports egress so far, we leave tableID as 0
// for ingress rules for multicast, later we will return empty flows for it.

// getOFPriority retrieves the OFPriority for the input CompletedRule to be installed,
// and re-arranges installed priorities on OVS if necessary.
func (r *podReconciler) getOFPriority(rule *CompletedRule, tableID uint8, pa *tablePriorityAssigner) (*uint16, bool, error) {
	_ = "STUB: not implemented"
	// IGMP Egress policy is enforced in userspace via packet-in message, there won't be OpenFlow
	// rules created for such rules. Therefore, assigning priority is not required.
	return nil, false, nil
}

// Re-assign installed priorities on OVS

// BatchReconcile reconciles the desired state of the provided CompletedRules
// with the actual state of Openflow entries in batch. It should only be invoked
// if all rules are newly added without last realized status.
func (r *podReconciler) BatchReconcile(rules []*CompletedRule) error {
	_ = "STUB: not implemented"
	return nil
}

// If batch reconcile fails, all priorities should be released and the
// priorityAssigners should return to the initial state.

// registerOFPriorities constructs a Priority type for each CompletedRule in the input list,
// and registers those Priorities with appropriate tablePriorityAssigner based on Tier.
func (r *podReconciler) registerOFPriorities(rules []*CompletedRule) error {
	_ = "STUB: not implemented"
	return nil
}

// IGMP Egress policy is enforced in userspace via packet-in message, there won't be OpenFlow
// rules created for such rules. Therefore, assigning priority is not required.

// add converts CompletedRule to PolicyRule(s) and invokes installOFRule to install them.
func (r *podReconciler) add(rule *CompletedRule, ofPriority *uint16, table uint8) error {
	_ = "STUB: not implemented"
	return nil
}

// Each pod group gets an Openflow ID.

// Record ofID only if its Openflow is installed successfully.

func (r *podReconciler) computeOFRulesForAdd(rule *CompletedRule, ofPriority *uint16, table uint8) (
	map[servicesKey]*types.PolicyRule, *podPolicyLastRealized) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Handle the case that the following processing fails or partially succeeds.

// IGMP query

// IGMP report

// Addresses got from source GroupMembers' IPs.

// Get addresses that in From IPBlock but not in Except IPBlocks.

// If rule is applied to Services, there will be only one svcKey, which is "", in
// membersByServicesMap. So podPolicyLastRealized.serviceGroupIDs won't be overwritten in
// this for-loop.

// TODO: addFQDNRule installs new conjunctive flows, so maybe it doesn't
// belong in computeOFRulesForAdd. The error handling needs to be corrected
// as well: if the flows failed to install, there should be a retry
// mechanism.

// If there are no "ToAddresses", the above process doesn't create any PolicyRule.
// We must ensure there is at least one PolicyRule, otherwise the Pods won't be
// isolated, so we create a PolicyRule with the original services if it doesn't exist.
// If there are IPBlocks or Pods that cannot resolve any named port, they will share
// this PolicyRule. Antrea policies do not need this default isolation.

// Create a new Openflow rule if the group doesn't exist.

// Diff Addresses between To and Except of IPBlocks

// If the rule installation fails, this will be reset

// If the rule installation fails, this will be reset.

// batchAdd converts CompletedRules to PolicyRules and invokes BatchInstallPolicyRuleFlows to install them.
func (r *podReconciler) batchAdd(rules []*CompletedRule, ofPriorities []*uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// update calculates the difference of Addresses between oldRule and newRule,
// and invokes Openflow client's methods to reconcile them.
func (r *podReconciler) update(lastRealized *podPolicyLastRealized, newRule *CompletedRule, ofPriority *uint16, table uint8) error {
	_ = "STUB: not implemented"
	return nil
}

// staleOFIDs tracks servicesKey that are no long needed.
// Firstly fill it with the last realized ofIDs.

// IGMP query

// Install a new Openflow rule if this group doesn't exist, otherwise do incremental update.

// IGMP report

// As rule identifier is calculated from the rule's content, the update can
// only happen to Group members.

// Install a new Openflow rule if this group doesn't exist, otherwise do incremental update.

// Delete valid servicesKey from staleOFIDs.

// Same as the process in `add`, we must ensure the group for the original services is present
// in memberByServicesMap, so that this group won't be removed and its "From" will be updated.

// If the PolicyRule for the original services doesn't exist and IPBlocks is present, it means the
// podReconciler hasn't installed flows for IPBlocks, then it must be added to the new PolicyRule.

// Update the FQDN address set if rule installation succeeds.

// Update the groupID address set if rule installation succeeds.

// Delete valid servicesKey from staleOFIDs.

// Remove stale Openflow rules.

func (r *podReconciler) installOFRule(ofRule *types.PolicyRule) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *podReconciler) updateOFRule(ofID uint32, addedFrom []types.Address, addedTo []types.Address, deletedFrom []types.Address, deletedTo []types.Address, priority *uint16, enableLogging, isMCNPRule bool) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: This might be unnecessarily complex and hard for error handling, consider revising the Openflow interfaces.

func (r *podReconciler) uninstallOFRule(ofID uint32, table uint8) error {
	_ = "STUB: not implemented"
	return nil
}

// Cannot parse the priority str. Theoretically this should never happen.

// If there are stalePriorities, priorityAssigners[table] must not be nil.

// Forget invokes UninstallPolicyRuleFlows to uninstall Openflow entries
// associated with the provided ruleID if it was enforced before.
func (r *podReconciler) Forget(ruleID string) error { _ = "STUB: not implemented"; return nil }

// No-op if the rule was not realized before.

func (r *podReconciler) isIGMPRule(rule *CompletedRule) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *podReconciler) GetRuleByFlowID(ruleFlowID uint32) (*types.PolicyRule, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (r *podReconciler) getOFPorts(members v1beta2.GroupMemberSet) sets.Set[int32] {
	_ = "STUB: not implemented"
	return nil
}

// This might be because the container has been deleted during realization or hasn't been set up yet.

func (r *podReconciler) getIPs(members v1beta2.GroupMemberSet) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// This might be because the container has been deleted during realization or hasn't been set up yet.

func (r *podReconciler) getSvcGroupIDs(members v1beta2.GroupMemberSet) sets.Set[int64] {
	_ = "STUB: not implemented"
	return nil
}

// groupMembersByServices groups the provided groupMembers based on their services resolving result.
// A map of servicesHash to the grouped members and a map of servicesHash to the services resolving result will be returned.
func groupMembersByServices(services []v1beta2.Service, memberSet v1beta2.GroupMemberSet) (map[servicesKey]v1beta2.GroupMemberSet, map[servicesKey][]v1beta2.Service) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If there is no named port in services, all members are in same group.

// Reuse the slice to avoid memory reallocations in the following loop. The
// optimization makes difference as the number of group members might get up to tens
// of thousands.

// Copy resolvedServices as it may be updated in next iteration.

func ofPortsToOFAddresses(ofPorts sets.Set[int32]) []types.Address {
	_ = "STUB: not implemented"
	// Must not return nil as it means not restricted by addresses in Openflow implementation.
	return nil
}

func (r *podReconciler) svcRefsToGroupIDs(svcRefs []v1beta2.ServiceReference) sets.Set[int64] {
	_ = "STUB: not implemented"
	return nil
}

func svcGroupIDsToOFAddresses(groupIDs sets.Set[int64]) []types.Address {
	_ = "STUB: not implemented"
	// Must not return nil as it means not restricted by addresses in Openflow implementation.
	return nil
}

func groupMembersToOFAddresses(groupMemberSet v1beta2.GroupMemberSet) []types.Address {
	_ = "STUB: not implemented"
	// Must not return nil as it means not restricted by addresses in Openflow implementation.
	return nil
}

func ipBlocksToOFAddresses(ipBlocks []v1beta2.IPBlock, ipv4Enabled, ipv6Enabled, ctMatch bool) []types.Address {
	_ = "STUB: not implemented"
	// Must not return nil as it means not restricted by addresses in Openflow implementation.
	return nil
}

// This is part of normal operations: "allow all" in a policy is represented
// by the combination of 2 CIDRs. One is "0.0.0.0/0" (any v4) and one is
// "::/0" (any v6). In single-stack clusters, one of these CIDRs is
// irrelevant and should be ignored.

func labelIDToOFAddresses(labelIDs []uint32) []types.Address {
	_ = "STUB: not implemented"
	// Must not return nil as it means not restricted by addresses in Openflow implementation.
	return nil
}

func isIPNetSupportedByAF(ipnet *net.IPNet, ipv4Enabled, ipv6Enabled bool) bool {
	_ = "STUB: not implemented"
	return false
}

func ipsToOFAddresses(ips sets.Set[string]) []types.Address {
	_ = "STUB: not implemented"
	// Must not return nil as it means not restricted by addresses in Openflow implementation.
	return nil
}

func filterUnresolvablePort(in []v1beta2.Service) []v1beta2.Service {
	_ = "STUB: not implemented"
	// Empty or nil slice means allowing all ports in Kubernetes.
	// nil must be returned to meet ofClient's expectation for this behavior.
	return nil
}

// It makes sure `out` won't be nil, so that even if only named ports are
// specified and none of them are resolvable, the rule just falls back to
// allowing no port, instead of all ports.

// All resolvable named port have been converted to intstr.Int,
// ignore unresolvable ones.

// resolveService resolves the port name of the provided service to a port number for the provided groupMember.
// This function should eventually supersede resolveServiceForPod.
func resolveService(service *v1beta2.Service, member *v1beta2.GroupMember) *v1beta2.Service {
	_ = "STUB: not implemented"
	// If port is not specified or is already a number, return it as is.
	return nil
}

// For K8s NetworkPolicy and Antrea-native policies, the Service.Protocol field will never be nil since TCP
// will be filled as default. However, for AdminNetworkPolicy and BaselineAdminNetworkPolicy, the Protocol
// field will be nil for ports written as named ports. In that case, the member port will match as long
// as the port name is matched.

// Derive named port protocol from the container spec

// If not resolvable, return it as is.
// The group members that cannot resolve it will be grouped together.

func isRuleAppliedToService(memberSet v1beta2.GroupMemberSet) bool {
	_ = "STUB: not implemented"
	return false
}
