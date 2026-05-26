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

package stats

import (
	"time"

	"k8s.io/apimachinery/pkg/types"

	"antrea.io/antrea/v2/pkg/agent/client"
	"antrea.io/antrea/v2/pkg/agent/multicast"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	agenttypes "antrea.io/antrea/v2/pkg/agent/types"
	cpv1beta "antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	statsv1alpha1 "antrea.io/antrea/v2/pkg/apis/stats/v1alpha1"
	"antrea.io/antrea/v2/pkg/querier"
	utilwait "antrea.io/antrea/v2/pkg/util/wait"
)

const (
	// Period for performing stats collection and report.
	collectPeriod = 60 * time.Second
)

// statsCollection is a collection of stats.
type statsCollection struct {
	// networkPolicyStats is a mapping from K8s NetworkPolicy UIDs to their traffic stats.
	networkPolicyStats map[types.UID]*statsv1alpha1.TrafficStats
	// antreaClusterNetworkPolicyStats is a mapping from Antrea ClusterNetworkPolicy UIDs to their traffic stats.
	antreaClusterNetworkPolicyStats map[types.UID]map[string]*statsv1alpha1.TrafficStats
	// antreaNetworkPolicyStats is a mapping from Antrea NetworkPolicy UIDs to their traffic stats.
	antreaNetworkPolicyStats map[types.UID]map[string]*statsv1alpha1.TrafficStats
	// multicastGroups is a map that encodes the list of Pods that has joined the multicast group.
	multicastGroups map[string][]cpv1beta.PodReference
}

// Collector is responsible for collecting stats from the Openflow client, calculating the delta compared with the last
// reported stats, and reporting it to the antrea-controller summary API.
type Collector struct {
	nodeName string
	// antreaClientProvider provides interfaces to get antreaClient, which will be used to report the statistics to the
	// antrea-controller.
	antreaClientProvider client.AntreaClientProvider
	// ofClient is the Openflow interface that can fetch the statistic of the Openflow entries.
	ofClient             openflow.Client
	networkPolicyQuerier querier.AgentNetworkPolicyInfoQuerier
	multicastQuerier     querier.AgentMulticastInfoQuerier
	// lastStatsCollection is the last statistics that has been reported to antrea-controller successfully.
	// It is used to calculate the delta of the statistics that will be reported.
	lastStatsCollection *statsCollection
	multicastEnabled    bool
	// staleFlowsDeletedWait, when non-nil, must unblock before the first collect so metric dumps
	// do not include flows from a prior agent round. The initializer's stale-flow goroutine calls
	// Done on this group when deletion finishes or exits.
	staleFlowsDeletedWait *utilwait.Group
}

func NewCollector(antreaClientProvider client.AntreaClientProvider, ofClient openflow.Client, npQuerier querier.AgentNetworkPolicyInfoQuerier, mcQuerier *multicast.Controller, staleFlowsDeletedWait *utilwait.Group) *Collector {
	_ = "STUB: not implemented"
	return nil
}

// Run runs a loop that collects statistics and reports them until the provided channel is closed.
func (m *Collector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Record the initial statistics as the base that will be used to calculate the delta.
// If the counters increase during antrea-agent's downtime, the delta will not be reported to the antrea-controller,
// it's however better than reporting the full statistics twice which could introduce greater deviations.

// Do not update m.lastStatsMap if the report fails so that the next report attempt can add up the
// statistics produced in this duration.

// collect collects the stats of Openflow rules, maps them to the stats of NetworkPolicies.
// It returns a map from NetworkPolicyReferences to their stats.
func (m *Collector) collect() *statsCollection { _ = "STUB: not implemented"; return nil }

// This should not happen because the rule flow ID to rule mapping is
// preserved for at least 5 seconds even after the rule deletion.

func addPolicyStatsUp(statsMap map[types.UID]*statsv1alpha1.TrafficStats, ruleStats *agenttypes.RuleMetric, rule *agenttypes.PolicyRule) {
	_ = "STUB: not implemented"
	return
}

func addRuleStatsUp(ruleStatsMap map[types.UID]map[string]*statsv1alpha1.TrafficStats, ruleStats *agenttypes.RuleMetric, rule *agenttypes.PolicyRule) {
	_ = "STUB: not implemented"
	return
}

func addUp(stats *statsv1alpha1.TrafficStats, inc *agenttypes.RuleMetric) {
	_ = "STUB: not implemented"
	return
}

func isIdenticalMulticastGroupMap(a, b map[string][]cpv1beta.PodReference) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Collector) calculateNPStats(curStatsCollection *statsCollection) (npStats, acnpStats, annpStats []cpv1beta.NetworkPolicyStats) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (m *Collector) calculateNodeStatsSummary(curStatsCollection *statsCollection) *cpv1beta.NodeStatsSummary {
	_ = "STUB: not implemented"
	return nil
}

// Semantically, reporting networkpolicy statistics with zero length is equal to reporting the same multicastGroupInfo.

// mergeStatsWithIGMPReports merges acnpStats or annpStats with IGMP report statistics.
// Unlike other networkpolicystats collection process, IGMP report statistics is not collected from OVS flows. It was collected during IGMP packetIn process by a local cache.
// IGMP report statistics collected for a rule should be merged into already defined networkpolicy statistics before reporting.
func (m *Collector) mergeStatsWithIGMPReports(acnpStats, annpStats []cpv1beta.NetworkPolicyStats) ([]cpv1beta.NetworkPolicyStats, []cpv1beta.NetworkPolicyStats) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertMulticastGroups converts multicastGroupMap into a slice of multicastGroups.
// Calculating diff is not needed because we report full multicast group of the local node.
func (m *Collector) convertMulticastGroups(multicastGroupMap map[string][]cpv1beta.PodReference) []cpv1beta.MulticastGroupInfo {
	_ = "STUB: not implemented"
	return nil
}

// report calculates the delta of the stats and pushes it to the antrea-controller summary API.
// If multicast feature gate is enabled, it also sends the full multicast group and IGMP report stats to the antrea-controller.
func (m *Collector) report(curStatsCollection *statsCollection) error {
	_ = "STUB: not implemented"
	return nil
}

func calculateRuleDiff(curStatsMap, lastStatsMap map[types.UID]map[string]*statsv1alpha1.TrafficStats) []cpv1beta.NetworkPolicyStats {
	_ = "STUB: not implemented"
	return nil
}

// curRuleStats.Bytes < lastRuleStats.Bytes could happen
// as rules with same name can be deleted and recreated later.

func calculateDiff(curStatsMap, lastStatsMap map[types.UID]*statsv1alpha1.TrafficStats) []cpv1beta.NetworkPolicyStats {
	_ = "STUB: not implemented"
	return nil
}

// curStats.Bytes < lastStats.Bytes could happen if one of the following conditions happens:
// 1. OVS is restarted and Openflow entries are reinstalled.
// 2. The NetworkPolicy is removed and recreated in-between two collection.
// In these cases, curStats is the delta it should report.

// If the statistics of the NetworkPolicy remain unchanged, no need to report it.
