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
	"sync"

	networkinginformers "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/tools/cache"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	statsv1alpha1 "antrea.io/antrea/v2/pkg/apis/stats/v1alpha1"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
)

const (
	uidIndex           = "uid"
	GroupNameIndexName = "groupName"
)

// Aggregator collects the stats from the antrea-agents, aggregates them, caches the result, and provides interfaces
// for Stats API handlers to query them. It implements the following interfaces:
// - pkg/apiserver/registry/controlplane/nodestatssummary.statsCollector
// - pkg/apiserver/registry/stats/networkpolicystats.statsProvider
// - pkg/apiserver/registry/stats/antreaclusternetworkpolicystats.statsProvider
// - pkg/apiserver/registry/stats/antreanetworkpolicystats.statsProvider
// - pkg/apiserver/registry/stats/multicastgroup.statsProvider
type Aggregator struct {
	// networkPolicyStats caches the statistics of K8s NetworkPolicies collected from the antrea-agents.
	networkPolicyStats cache.Indexer
	// antreaClusterNetworkPolicyStats caches the statistics of Antrea ClusterNetworkPolicies collected from the antrea-agents.
	antreaClusterNetworkPolicyStats cache.Indexer
	// antreaNetworkPolicyStats caches the statistics of Antrea NetworkPolicies collected from the antrea-agents.
	antreaNetworkPolicyStats cache.Indexer
	// groupNodePodsMap caches the information of Pods in a Node that have joined multicast groups collected from the antrea-agents.
	// The map can be interpreted as
	// map[IP of multicast group]map[name of node]list of PodReference.
	groupNodePodsMap      map[string]map[string][]statsv1alpha1.PodReference
	groupNodePodsMapMutex sync.RWMutex
	// dataCh is the channel that buffers the NodeSummaries sent by antrea-agents.
	dataCh chan *controlplane.NodeStatsSummary
	// npListerSynced is a function which returns true if the K8s NetworkPolicy shared informer has been synced at least once.
	npListerSynced cache.InformerSynced
	// acnpListerSynced is a function which returns true if the Antrea ClusterNetworkPolicy shared informer has been synced at least once.
	acnpListerSynced cache.InformerSynced
	// annpListerSynced is a function which returns true if the Antrea NetworkPolicy shared informer has been synced at least once.
	annpListerSynced cache.InformerSynced
}

// uidIndexFunc is an index function that indexes based on an object's UID.
func uidIndexFunc(obj interface{}) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func NewAggregator(networkPolicyInformer networkinginformers.NetworkPolicyInformer, acnpInformer crdinformers.ClusterNetworkPolicyInformer, annpInformer crdinformers.NetworkPolicyInformer) *Aggregator {
	_ = "STUB: not implemented"
	return nil
}

// Add handlers for NetworkPolicy events.
// They are the source of truth of the NetworkPolicyStats, i.e., a NetworkPolicyStats is present only if the
// corresponding NetworkPolicy is present.

// Set resyncPeriod to 0 to disable resyncing.

// Register Informer and add handlers for AntreaPolicy events only if the feature is enabled.
// They are the source of truth of the ClusterNetworkPolicyStats, i.e., a ClusterNetworkPolicyStats is present
// only if the corresponding ClusterNetworkPolicy is present.

// Set resyncPeriod to 0 to disable resyncing.

// Set resyncPeriod to 0 to disable resyncing.

// addNetworkPolicy handles NetworkPolicy ADD events and creates corresponding NetworkPolicyStats objects.
func (a *Aggregator) addNetworkPolicy(obj interface{}) { _ = "STUB: not implemented"; return }

// To indicate the duration that the stats cover, the CreationTimestamp is set to the time that the stats
// start, instead of the CreationTimestamp of the NetworkPolicy.

// deleteNetworkPolicy handles NetworkPolicy DELETE events and deletes corresponding NetworkPolicyStats objects.
func (a *Aggregator) deleteNetworkPolicy(obj interface{}) { _ = "STUB: not implemented"; return }

// addACNP handles ClusterNetworkPolicy ADD events and creates corresponding ClusterNetworkPolicyStats objects.
func (a *Aggregator) addACNP(obj interface{}) { _ = "STUB: not implemented"; return }

// To indicate the duration that the stats covers, the CreationTimestamp is set to the time that the stats
// start, instead of the CreationTimestamp of the ClusterNetworkPolicy.

// deleteACNP handles ClusterNetworkPolicy DELETE events and deletes corresponding ClusterNetworkPolicyStats objects.
func (a *Aggregator) deleteACNP(obj interface{}) { _ = "STUB: not implemented"; return }

// addANNP handles Antrea NetworkPolicy ADD events and creates corresponding AntreaNetworkPolicyStats objects.
func (a *Aggregator) addANNP(obj interface{}) { _ = "STUB: not implemented"; return }

// To indicate the duration that the stats covers, the CreationTimestamp is set to the time that the stats
// start, instead of the CreationTimestamp of the Antrea NetworkPolicy.

// deleteANNP handles Antrea NetworkPolicy DELETE events and deletes corresponding AntreaNetworkPolicyStats objects.
func (a *Aggregator) deleteANNP(obj interface{}) { _ = "STUB: not implemented"; return }

func (a *Aggregator) ListAntreaClusterNetworkPolicyStats() []statsv1alpha1.AntreaClusterNetworkPolicyStats {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aggregator) ListMulticastGroups() []statsv1alpha1.MulticastGroup {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aggregator) GetMulticastGroup(group string) (*statsv1alpha1.MulticastGroup, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a *Aggregator) GetAntreaClusterNetworkPolicyStats(name string) (*statsv1alpha1.AntreaClusterNetworkPolicyStats, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a *Aggregator) ListAntreaNetworkPolicyStats(namespace string) []statsv1alpha1.AntreaNetworkPolicyStats {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aggregator) GetAntreaNetworkPolicyStats(namespace, name string) (*statsv1alpha1.AntreaNetworkPolicyStats, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a *Aggregator) ListNetworkPolicyStats(namespace string) []statsv1alpha1.NetworkPolicyStats {
	_ = "STUB: not implemented"
	return nil
}

func (a *Aggregator) GetNetworkPolicyStats(namespace, name string) (*statsv1alpha1.NetworkPolicyStats, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Collect collects the node summary asynchronously to avoid the competition for the statsLock and to save clients
// from pending on it.
func (a *Aggregator) Collect(summary *controlplane.NodeStatsSummary) {
	_ = "STUB: not implemented"
	return

	// Run runs a loop that keeps taking stats summary from the data channel and actually collecting them until the
	// provided stop channel is closed.
}

func (a *Aggregator) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (a *Aggregator) doCollect(summary *controlplane.NodeStatsSummary) {
	_ = "STUB: not implemented"
	return
}

// The policy might have been removed, skip processing it if missing.

// The object returned by cache is supposed to be read only, create a new object and update it.

// The antrea-agent reports full mcastGroupInfo to the controller, if the group is unreported,
// then this group has no Pod joined in this Node.

// The policy have might been removed, skip processing it if missing.

// The object returned by cache is supposed to be read only, create a new object and update it.

// antrea agents may not be updated and still use TrafficStats to collect overall networkpolicy

// The policy have might been removed, skip processing it if missing.

// The object returned by cache is supposed to be read only, create a new object and update it.

// antrea agents may not be updated and still use TrafficStats to collect overall networkpolicy

func addUp(stats *statsv1alpha1.TrafficStats, inc *statsv1alpha1.TrafficStats) {
	_ = "STUB: not implemented"
	return
}

func addRulesUp(ruleStats *[]statsv1alpha1.RuleTrafficStats, ruleSumStats *statsv1alpha1.TrafficStats, inc []statsv1alpha1.RuleTrafficStats) {
	_ = "STUB: not implemented"
	return
}

// accumulate incMap traffics stats to the current traffic stats

// accumulate the rule traffic stats as the rule has already 'existed' in the ruleStats

// convert remaining incs to RuleTrafficStats and add it to current traffic stats
