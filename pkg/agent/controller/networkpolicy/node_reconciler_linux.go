//go:build linux
// +build linux

// Copyright 2024 Antrea Authors
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

	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/agent/route"
	"antrea.io/antrea/v2/pkg/agent/types"
	"antrea.io/antrea/v2/pkg/agent/util/ipset"
	"antrea.io/antrea/v2/pkg/agent/util/iptables"
	"antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	secv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

const (
	ipv4Any = "0.0.0.0/0"
	ipv6Any = "::/0"
)

// The logging of Node NetworkPolicy is implemented by iptables target LOG, which turns on kernel logging of matching
// packets. The label is useful for distinguishing Node NetworkPolicy logs from other kernel logs.
const logLabelPrefix = "Antrea"

var ipsetTypeHashIP = ipset.HashIP

/*
Tips:
In the following, service describes a port to allow traffic on which is defined in pkg/apis/controlplane/v1beta2/types.go

NodeNetworkPolicy data path implementation using iptables/ip6tables involves four components:
1. Core iptables rule:
   - Added to ANTREA-POL-INGRESS-RULES (ingress) or ANTREA-POL-EGRESS-RULES (egress).
   - Matches an ipset created for the NodeNetworkPolicy rule as source (ingress) or destination (egress) when there are
     multiple IP addresses; if there is only one address, matches the address directly.
   - Targets an action (the rule with a single service) or a service chain created for the NodeNetworkPolicy rule (the
     rule with multiple services).
2. Service iptables chain:
   - Created for the NodeNetworkPolicy rule to integrate service iptables rules if a rule has multiple services.
3. Service iptables rules:
   - Added to the service chain created for the NodeNetworkPolicy rule.
   - Constructed from the services of the NodeNetworkPolicy rule.
4. From/To ipset:
   - Created for the NodeNetworkPolicy rule, containing all source IP addresses (ingress) or destination IP addresses (egress).

Assuming four ingress NodeNetworkPolicy rules with IDs RULE1, RULE2, RULE3 and RULE4 prioritized in descending order.
Core iptables rules organized by priorities in ANTREA-POL-INGRESS-RULES like the following.

If the rule has multiple source IP addresses to match, then an ipset will be created for it. The name of the ipset consists
of prefix "ANTREA-POL", rule ID and IP protocol version.

If the rule has multiple services, an iptables chain and related rules will be created for it. The name the chain consists
of prefix "ANTREA-POL" and rule ID.

```
:ANTREA-POL-INGRESS-RULES
-A ANTREA-POL-INGRESS-RULES -m set --match-set ANTREA-POL-RULE1-4 src -j ANTREA-POL-RULE1 -m comment --comment "Antrea: for rule RULE1, policy AntreaClusterNetworkPolicy:name1"
-A ANTREA-POL-INGRESS-RULES -m set --match-set ANTREA-POL-RULE2-4 src -p tcp --dport 8080 -j ACCEPT -m comment --comment "Antrea: for rule RULE2, policy AntreaClusterNetworkPolicy:name2"
-A ANTREA-POL-INGRESS-RULES -s 3.3.3.3/32 src -j ANTREA-POL-RULE3 -m comment --comment "Antrea: for rule RULE3, policy AntreaClusterNetworkPolicy:name3"
-A ANTREA-POL-INGRESS-RULES -s 4.4.4.4/32 -p tcp --dport 80 -j ACCEPT -m comment --comment "Antrea: for rule RULE4, policy AntreaClusterNetworkPolicy:name4"
```

For the first rule, it has multiple services and multiple source IP addresses to match, so there will be service iptables chain
and service iptables rules and ipset created for it.

The iptables chain is like the following:

```
:ANTREA-POL-RULE1
-A ANTREA-POL-RULE1 -j ACCEPT -p tcp --dport 80
-A ANTREA-POL-RULE1 -j ACCEPT -p tcp --dport 443
```

The ipset is like the following:

```
Name: ANTREA-POL-RULE1-4
Type: hash:net
Revision: 6
Header: family inet hashsize 1024 maxelem 65536
Size in memory: 472
References: 1
Number of entries: 2
Members:
1.1.1.1
1.1.1.2
```

For the second rule, it has only one service, so there will be no service iptables chain and service iptables rules created
for it. The core rule will match the service and target the action directly. The rule has multiple source IP addresses to
match, so there will be an ipset `ANTREA-POL-RULE2-4` created for it.

For the third rule, it has multiple services to match, so there will be service iptables chain and service iptables rules
created for it. The rule has only one source IP address to match, so there will be no ipset created for it and just match
the source IP address directly.

For the fourth rule, it has only one service and one source IP address to match, so there will be no service iptables chain
and service iptables rules created for it. The core rule will match the service and source IP address and target the action
directly.
*/

// coreIPTRule is a struct to store the information of a core iptables rule.
type coreIPTRule struct {
	ruleID   string
	priority *types.Priority
	ruleStrs []string
}

type chainKey struct {
	name   string
	isIPv6 bool
}

func newChainKey(name string, isIPv6 bool) chainKey {
	_ = "STUB: not implemented"
	return *new(chainKey)
}

// coreIPTChain caches the sorted iptables rules for a chain where core iptables rules are installed.
type coreIPTChain struct {
	rules []*coreIPTRule
	sync.Mutex
}

func newCoreIPTChain() *coreIPTChain { _ = "STUB: not implemented"; return nil }

// nodePolicyLastRealized is the struct cached by nodeReconciler. It's used to track the actual state of iptables rules
// and chains we have enforced, so that we can know how to reconcile a rule when it's updated/removed.
type nodePolicyLastRealized struct {
	// The desired state of a policy rule.
	*CompletedRule
	// ipsets tracks the last realized ipset names used in core iptables rules. It cannot coexist with ipnets.
	ipsets map[iptables.Protocol]string
	// ipnets tracks the last realized ip nets used in core iptables rules. It cannot coexist with ipsets.
	ipnets map[iptables.Protocol]string
	// serviceIPTChain tracks the last realized service iptables chain if a rule has multiple services.
	serviceIPTChain string
	// coreIPTChain tracks the last realized iptables chain where the core iptables rule is installed.
	coreIPTChain string
}

func newNodePolicyLastRealized(rule *CompletedRule) *nodePolicyLastRealized {
	_ = "STUB: not implemented"
	return nil
}

type nodeReconciler struct {
	ipProtocols   []iptables.Protocol
	routeClient   route.Interface
	coreIPTChains map[chainKey]*coreIPTChain
	// lastRealizeds caches the last realized rules. It's a mapping from ruleID to *nodePolicyLastRealized.
	lastRealizeds sync.Map
}

func newNodeReconciler(routeClient route.Interface, ipv4Enabled, ipv6Enabled bool) *nodeReconciler {
	_ = "STUB: not implemented"
	return nil
}

// Reconcile checks whether the provided rule has been enforced or not, and invoke the add or update method accordingly.
func (r *nodeReconciler) Reconcile(rule *CompletedRule) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *nodeReconciler) RunIDAllocatorWorker(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (r *nodeReconciler) BatchReconcile(rules []*CompletedRule) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *nodeReconciler) batchAdd(rules []*CompletedRule) error {
	_ = "STUB: not implemented"
	return nil
}

// Sync all ipsets.

// Collect all service iptables rules and chains.

// Collect all core iptables rules.

func (r *nodeReconciler) Forget(ruleID string) error { _ = "STUB: not implemented"; return nil }

// No-op if the rule was not realized before.

func (r *nodeReconciler) GetRuleByFlowID(ruleFlowID uint32) (*types.PolicyRule, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (r *nodeReconciler) computeIPTRules(rule *CompletedRule) (map[iptables.Protocol]*types.NodePolicyRule, *nodePolicyLastRealized) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If a rule has multiple services, create a chain to install iptables rules for these services, with the target
// of the services determined by the rule's action. The core iptables rule should target the chain.

// If a rule has no service or a single service, the target is determined by the rule's action, as there is no
// need to create a chain for a single-service iptables rule.

// If a rule has a single service, the core iptables rule directly incorporates the service.

// If a rule matches multiple source or destination ipnets, create an ipset which contains these ipnets and
// use the ipset in core iptables rule.

// If a rule matches single source or destination, use it in core iptables rule directly.

// If the target of a core iptables rule is not a service chain, the iptables rule for logging should be
// generated along with the core iptables rule. Otherwise, the iptables rules for logging should be generated
// along with the service iptables rules.

func (r *nodeReconciler) add(rule *CompletedRule) error { _ = "STUB: not implemented"; return nil }

func (r *nodeReconciler) update(lastRealized *nodePolicyLastRealized, newRule *CompletedRule) error {
	_ = "STUB: not implemented"
	return nil
}

// Core iptables rules should be updated in the following cases:
// - Single IP change: A -> B (prevIPSet = "", ipset = "", prevIPNet = A, ipnet = B).
// - Transition from multiple addresses to a single IP: {A, B} -> A (prevIPSet = "ipset name", ipset = "", prevIPNet = "", ipnet = A).
// - Transition from a single IP to multiple addresses: A -> {A, B} (prevIPNet = A, ipnet = "", prevIPSet = "", ipset = "ipset name").

// The name of ipset for a rule will never change during updates.

// If the current rule uses an ipset, sync the ipset first, then sync the core iptables rule that
// references it.

// If the previous rule used an ipset, sync the new core iptables rule first to remove its reference, then
// delete the unused ipset.

func (r *nodeReconciler) addOrUpdateCoreIPTRules(chain string, isIPv6 bool, isUpdate bool, newRules ...*coreIPTRule) error {
	_ = "STUB: not implemented"
	return nil
}

// Build a map to store the mapping of rule ID to rule for the rules to update.

// Iterate each existing rule. If an existing rule exists in rulesToUpdate, replace it with the new rule.

// If these are new rules, append the new rules then sort all rules.

// Get all iptables rules and synchronize them.

// cache the updated rules.

func (r *nodeReconciler) deleteCoreIPTRule(ruleID string, iptChain string, isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Get all the cached rules, then delete the rule with the given rule ID.

// If the rule is not found, return directly.

// If the rule is found, delete it from the slice.

// Get all the iptables rules and synchronize them.

// cache the updated rules.

func (r *nodeReconciler) getCoreIPTChain(iptChain string, isIPv6 bool) *coreIPTChain {
	_ = "STUB: not implemented"
	// - For IPv4 ingress rules, iptables rules are installed in chain ANTREA-INGRESS-RULES.
	// - For IPv6 ingress rules, ip6tables rules are installed in chain ANTREA-INGRESS-RULES.
	// - For IPv4 egress rules, iptables rules are installed in chain ANTREA-EGRESS-RULES.
	// - For IPv6 egress rules, ip6tables rules are installed in chain ANTREA-EGRESS-RULES.
	return nil
}

func groupMembersToIPNets(groups v1beta2.GroupMemberSet, isIPv6 bool) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

func ipBlocksToIPNets(ipBlocks []v1beta2.IPBlock, isIPv6 bool) []string {
	_ = "STUB: not implemented"
	return nil
}

func getIPNetsFromRule(rule *CompletedRule, isIPv6 bool) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// If the set contains "0.0.0.0/0" or "::/0", it means the rule matches any source or destination IP address, just
// return a new set only containing "0.0.0.0/0" or "::/0".

func buildCoreIPTRules(ipProtocol iptables.Protocol,
	iptChain string,
	ipset string,
	ipnet string,
	iptRuleTarget string,
	iptRuleComment string,
	service *v1beta2.Service,
	isIngress bool,
	enableLogging bool,
	logLabel string) []string {
	_ = "STUB: not implemented"
	return nil
}

// If no source IP address is matched, return an empty slice since the core iptables will never be matched.

// If no destination IP address is matched, return an empty slice since the core iptables will never be matched.

func buildServiceIPTRules(ipProtocol iptables.Protocol,
	services []v1beta2.Service,
	chain string,
	ruleTarget string,
	enableLogging bool,
	logLabel string) []string {
	_ = "STUB: not implemented"
	return nil
}

func ruleActionToIPTTarget(ruleAction *secv1beta1.RuleAction) string {
	_ = "STUB: not implemented"
	return ""
}

func getServiceTransProtocol(protocol *v1beta2.Protocol) string {
	_ = "STUB: not implemented"
	return ""
}

func generateLogLabel(rule *CompletedRule) string {
	_ = "STUB: not implemented"
	// Construct the log label used as iptables log prefix. According to https://ipset.netfilter.org/iptables-extensions.man.html,
	// the log prefix is up to 29 letters long. The log label should include essential information to help filter the
	// generated iptables kernel log. As a result, the user-provided log label is limited to 12 characters.
	// The log label format:
	// |Antrea|:|I|:|Reject|:|user-provided label|:|
	// |6     |1|1|1|4-6   |1|1-12               |1|
	return ""
}

// Truncate the user-provided log label if it exceeds 12 characters.
