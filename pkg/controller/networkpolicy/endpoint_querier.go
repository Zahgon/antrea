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

// Package networkpolicy provides NetworkPolicyController implementation to manage
// and synchronize the Pods and Namespaces affected by Network Policies and enforce
// their rules.
package networkpolicy

import (
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	antreatypes "antrea.io/antrea/v2/pkg/controller/types"
)

// EndpointQuerier handles requests for querying NetworkPolicies of the endpoint.
type EndpointQuerier interface {
	// QueryNetworkPolicyRules returns the list of NetworkPolicies which apply to the provided Pod,
	// along with the list of NetworkPolicy ingress/egress rules which select the provided Pod.
	QueryNetworkPolicyRules(namespace, podName string) (*antreatypes.EndpointNetworkPolicyRules, error)
}

// EndpointQuerierImpl implements the EndpointQuerier interface
type EndpointQuerierImpl struct {
	networkPolicyController *NetworkPolicyController
}

// NewEndpointQuerier returns a new *EndpointQuerierImpl.
func NewEndpointQuerier(networkPolicyController *NetworkPolicyController) *EndpointQuerierImpl {
	_ = "STUB: not implemented"
	return nil
}

// PolicyRuleQuerier handles requests for querying effective policy rule on entities.
type PolicyRuleQuerier interface {
	QueryNetworkPolicyEvaluation(entities *controlplane.NetworkPolicyEvaluationRequest) (*controlplane.NetworkPolicyEvaluationResponse, error)
}

// policyRuleQuerier implements the PolicyRuleQuerier interface
type policyRuleQuerier struct {
	endpointQuerier EndpointQuerier
}

// NewPolicyRuleQuerier returns a new *policyRuleQuerier
func NewPolicyRuleQuerier(endpointQuerier EndpointQuerier) *policyRuleQuerier {
	_ = "STUB: not implemented"
	return nil
}

type lessFunc func(p1, p2 *antreatypes.RuleInfo) int

// ByRulePriority implements the Sort interface, sorting the rules within.
// Comparators should be ordered by their importance in terms of determining rule priority.
type ByRulePriority struct {
	rules       []*antreatypes.RuleInfo
	comparators []lessFunc
}

func (s ByRulePriority) Len() int { _ = "STUB: not implemented"; return 0 }

func (s ByRulePriority) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s ByRulePriority) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// p < q

// p > q

// p == q; try the next comparison.

// QueryNetworkPolicyRules returns network policies and rules relevant to the selected
// network endpoint. Relevant network policies fall into three categories: applied policies
// are policies which directly apply to an endpoint, egress/ingress rules are rules which
// reference the endpoint respectively.
func (eq *EndpointQuerierImpl) QueryNetworkPolicyRules(namespace, podName string) (*antreatypes.EndpointNetworkPolicyRules, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create network policies categories

// get all appliedToGroups using filter, then get applied policies using appliedToGroup

// We iterate over all AppliedToGroups (same for AddressGroups below). This is acceptable
// since this implementation only supports user queries (in particular through antctl) and
// should return within a reasonable amount of time. We experimented with adding Pod
// Indexers to the AppliedToGroup and AddressGroup stores, but we felt that this use case
// did not justify the memory overhead. If we can find another use for the Indexers as part
// of the NetworkPolicy Controller implementation, we may consider adding them back.

// get all addressGroups using filter, then get ingress and egress policies using addressGroup

// an AddressGroup can only be referenced in a rule once

// an AddressGroup can only be referenced in a rule once

// IngressIndex/egressIndex indicates the current rule's index among this policy's original ingress/egress
// rules. The calculation accounts for policy rules not referencing this pod, and guarantees that
// users can reference the rules from configuration without accessing the internal policies.

// processEndpointAppliedRules processes NetworkPolicy rules applied to an endpoint,
// returns a set of the corresponding policy UIDs, and manually generates Kubernetes
// NetworkPolicy default isolation rules if they exist. The default isolation rule's
// direction depends on isSourceEndpoint, and has the lowest precedence.
func processEndpointAppliedRules(appliedPolicies []*antreatypes.NetworkPolicy, isSourceEndpoint bool) (sets.Set[types.UID], []*antreatypes.RuleInfo) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if the Kubernetes NetworkPolicy creates ingress or egress isolationRules

// predictEndpointsRules returns the predicted rules effective from srcEndpoints to dstEndpoints.
// Rules returned satisfy a. in source applied policies and destination egress rules,
// or b. in source ingress rules and destination applied policies or c. applied to KNP default isolation.
func predictEndpointsRules(srcEndpointRules, dstEndpointRules *antreatypes.EndpointNetworkPolicyRules) (commonRule *antreatypes.RuleInfo) {
	_ = "STUB: not implemented"
	return nil
}

// sort the common rules based on multiple closures, the top rule has the highest precedence

// filter Antrea-native policy rules with Pass action
// if pass rule currently has the highest precedence, skip the remaining rules
// until the next K8s rule or Baseline rule, or return the pass rule otherwise

// QueryNetworkPolicyEvaluation returns the effective NetworkPolicy rule on given
// source and destination entities.
func (eq *policyRuleQuerier) QueryNetworkPolicyEvaluation(entities *controlplane.NetworkPolicyEvaluationRequest) (*controlplane.NetworkPolicyEvaluationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// query endpoints and handle response errors
