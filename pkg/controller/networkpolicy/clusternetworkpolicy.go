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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	antreatypes "antrea.io/antrea/v2/pkg/controller/types"
)

const (
	labelValueSeparator = ","
)

func getACNPReference(cnp *crdv1beta1.ClusterNetworkPolicy) *controlplane.NetworkPolicyReference {
	_ = "STUB: not implemented"
	return nil
}

// addCNP receives ClusterNetworkPolicy ADD events and enqueues a reference of
// the ClusterNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) addCNP(obj interface{}) { _ = "STUB: not implemented"; return }

// updateCNP receives ClusterNetworkPolicy UPDATE events and enqueues a
// reference of the ClusterNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) updateCNP(_, cur interface{}) { _ = "STUB: not implemented"; return }

// deleteCNP receives ClusterNetworkPolicy DELETE events and enqueues a
// reference of the ClusterNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) deleteCNP(old interface{}) { _ = "STUB: not implemented"; return }

// filterPerNamespaceRuleACNPsByNSLabels gets all ClusterNetworkPolicy names that will need to be
// re-processed based on the entire label set of an added/updated/deleted Namespace.
func (n *NetworkPolicyController) filterPerNamespaceRuleACNPsByNSLabels(nsLabels labels.Set) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// It's fine to ignore this peer if the ClusterGroup is not found. After the ClusterGroup is created,
// the ClusterNetworkPolicy will be reprocessed anyway.

// An empty nsLabelSelector means select from all Namespaces

// The policy has only spec level AppliedTo.

// The policy has rule level AppliedTo.
// It needs to check each rule's peers. If any peer of the rule has PeerNamespaces selector and its
// AppliedTo selects this Namespace, the ClusterNetworkPolicy will be affected by the Namespace.

// getACNPsWithRulesMatchingAnyLabelKey gets all ACNPs that have relevant rules based on Namespace label keys.
func (n *NetworkPolicyController) getACNPsWithRulesMatchingAnyLabelKey(labelKeys sets.Set[string]) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// getACNPsWithRulesMatchingAnyUpdatedLabels gets all ACNPs that have rules based on Namespace
// label keys, which have changes in value across Namespace update.
func (n *NetworkPolicyController) getACNPsWithRulesMatchingAnyUpdatedLabels(oldNSLabels, newNSLabels map[string]string) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// addNamespace receives Namespace ADD events and triggers all ClusterNetworkPolicies that have a
// per-namespace rule applied to this Namespace to be re-processed.
func (n *NetworkPolicyController) addNamespace(obj interface{}) { _ = "STUB: not implemented"; return }

// Ignore the ClusterNetworkPolicy if it has been removed during the process.

// updateNamespace receives Namespace UPDATE events and triggers all ClusterNetworkPolicies that have a
// per-namespace rule applied to either the original or the new Namespace to be re-processed.
// It also triggers all K8s NetworkPolicies in the new Namespace to be re-processed
// if the logging Annotation changes.
func (n *NetworkPolicyController) updateNamespace(oldObj, curObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// No ClusterNetworkPolicies are affected if the Namespace's labels do not change.

// Any ACNPs that has Namespace label rules that refers to the label key set that has
// changed during the Namespace update will need to be re-processed.

// Ignore the ClusterNetworkPolicy if it has been removed during the process.

// deleteNamespace receives Namespace DELETE events and triggers all ClusterNetworkPolicies that have a
// per-namespace rule applied to this Namespace to be re-processed.
func (n *NetworkPolicyController) deleteNamespace(old interface{}) {
	_ = "STUB: not implemented"
	return
}

// Ignore the ClusterNetworkPolicy if it has been removed during the process.

func (c *NetworkPolicyController) filterAGsFromNodeLabels(node *v1.Node) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

func (c *NetworkPolicyController) getATGsAppliedToService() sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

func (c *NetworkPolicyController) addNode(obj interface{}) { _ = "STUB: not implemented"; return }

// All AppliedToGroups that are applied to Services need re-sync.

func (c *NetworkPolicyController) deleteNode(obj interface{}) { _ = "STUB: not implemented"; return }

// enqueue affected address group

// All AppliedToGroups that are applied to Services need re-sync.

func nodeIPChanged(oldNode, newNode *v1.Node) (changed bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *NetworkPolicyController) updateNode(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// processClusterNetworkPolicy creates an internal NetworkPolicy instance
// corresponding to the crdv1beta1.ClusterNetworkPolicy object. This method
// does not commit the internal NetworkPolicy in store, instead returns an
// instance to the caller wherein, it will be either stored as a new Object
// in case of ADD event or modified and store the updated instance, in case
// of an UPDATE event.
func (n *NetworkPolicyController) processClusterNetworkPolicy(cnp *crdv1beta1.ClusterNetworkPolicy) (*antreatypes.NetworkPolicy, map[string]*antreatypes.AppliedToGroup, map[string]*antreatypes.AddressGroup) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// If one of the ACNP rule is a per-namespace rule (a peer in that rule has namespaces.Match set
// to Self), the policy will need to be converted to appliedTo per rule policy, as the appliedTo
// will be different for rules created for each namespace.

// appliedToGroups tracks all distinct appliedToGroups referred to by the ClusterNetworkPolicy,
// either in the spec section or in ingress/egress rules.
// The span calculation and stale appliedToGroup cleanup logic would work seamlessly for both cases.

// If appliedTo is set at spec level and the ACNP has per-namespace rules, then each appliedTo needs
// to be split into appliedToGroups for each of its affected Namespace.

// When appliedTo is set at spec level and the ACNP has rules that select peer Namespaces by sameLabels,
// this field tracks the labels of all Namespaces selected by the appliedTo.

// clusterSetScopeSelectorKeys keeps track of all the ClusterSet-scoped selector keys of the policy.
// During policy peer processing, any ClusterSet-scoped selector will be registered with the
// labelIdentityInterface and added to this set. By the end of the function, this set will
// be used to remove any stale selector from the policy in the labelIdentityInterface.

// When a rule's NetworkPolicyPeer is empty, a cluster level rule should be created
// with an Antrea peer matching all addresses.

// For ACNPs that have per-namespace rules, cluster-level rules will be created with appliedTo
// set as the spec appliedTo for each rule.

// Create a rule for each affected Namespace of appliedTo at spec level

// Create a rule for each affected Namespace of appliedTo at rule level

// All affected Namespaces and their labels are already stored in labelsPerAffectedNS

// Compute NetworkPolicyRules for Ingress Rules.

// Compute NetworkPolicyRules for Egress Rules.

// Create AppliedToGroup for each AppliedTo present in ClusterNetworkPolicy spec.

// serviceAccountNameToPodSelector returns a PodSelector which could be used to
// select Pods based on their ServiceAccountName.
func serviceAccountNameToPodSelector(saName string) *metav1.LabelSelector {
	_ = "STUB: not implemented"
	return nil
}

// hasPerNamespaceRule returns true if there is at least one per-namespace rule
func hasPerNamespaceRule(cnp *crdv1beta1.ClusterNetworkPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func namespaceRuleLabelKeys(cnp *crdv1beta1.ClusterNetworkPolicy) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkPolicyController) getNamespaceLabels(ns string) labels.Set {
	_ = "STUB: not implemented"
	return *new(labels.Set)
}

// The Namespace referred to (by ServiceAccount etc.) does not exist yet.
// ACNP will be re-queued once that Namespace event is received.

// groupNamespaceByLabelValue groups Namespaces if they have the same label value for all the
// label keys listed. If a Namespace is missing at least one of the label keys, it will not
// be grouped. Example:
//
//	  ns1: app=web, tier=test, tenant=t1
//	  ns2: app=web, tier=test, tenant=t2
//	  ns3: app=web, tier=production, tenant=t1
//	  ns4: app=web, tier=production, tenant=t2
//	  ns5: app=db, tenant=t1
//	labelKeys = [app, tier]
//	Result after grouping:
//	  "web,test,":       [ns1, ns2]
//	  "web,production,": [ns3, ns4]
func groupNamespacesByLabelValue(affectedNSAndLabels map[string]labels.Set, labelKeys []string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func getLabelValues(labels map[string]string, labelKeys []string) string {
	_ = "STUB: not implemented"
	return ""
}

// convertSameLabelsToSelector creates a LabelSelector based on a list of label keys
// and their expected values.
func convertSameLabelsToSelector(labelKeys []string, labelValues string) *metav1.LabelSelector {
	_ = "STUB: not implemented"
	return nil
}

// toAntreaPeerForSameLabelNamespaces computes the appliedToGroups and addressGroups for each
// group of Namespaces who have the same values for the sameLabels keys.
func (n *NetworkPolicyController) toAntreaPeerForSameLabelNamespaces(peer crdv1beta1.NetworkPolicyPeer,
	np metav1.Object, atgPerAffectedNS map[string]*antreatypes.AppliedToGroup,
	labelValues string,
	namespacesByLabelValues []string) (*controlplane.NetworkPolicyPeer, []*antreatypes.AppliedToGroup, []*antreatypes.AddressGroup, sets.Set[string]) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// select Namespaces who, for specific label keys, have the same values as the appliedTo Namespaces.

// In addition to getting the matched Label Identity IDs, AddSelector also registers the selector
// with the labelIdentityInterface.

// processClusterAppliedTo processes appliedTo groups in Antrea ClusterNetworkPolicy set
// at cluster level (appliedTo groups which will not need to be split by Namespaces).
func (n *NetworkPolicyController) processClusterAppliedTo(appliedTo []crdv1beta1.AppliedTo) []*antreatypes.AppliedToGroup {
	_ = "STUB: not implemented"
	return nil
}

// splitPeersByScope splits the ClusterNetworkPolicy peers in the rule by whether the peer
// is cluster-scoped or per-namespace.
func splitPeersByScope(rule *crdv1beta1.Rule, dir controlplane.Direction) ([]crdv1beta1.NetworkPolicyPeer, []crdv1beta1.NetworkPolicyPeer, []crdv1beta1.NetworkPolicyPeer) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// getAffectedNamespacesForAppliedTo computes the Namespaces currently affected by the appliedTo
// Namespace selectors, and returns these Namespaces along with their labels.
func (n *NetworkPolicyController) getAffectedNamespacesForAppliedTo(appliedTo crdv1beta1.AppliedTo) map[string]labels.Set {
	_ = "STUB: not implemented"
	return nil
}

// An empty nsLabelSelector means select from all Namespaces

// processInternalGroupForRule examines the internal group (and its childGroups if applicable)
// to determine whether an addressGroup needs to be created, and returns any ipBlocks contained
// by the internal Group as well.
func (n *NetworkPolicyController) processInternalGroupForRule(group *antreatypes.Group) (bool, []controlplane.IPBlock) {
	_ = "STUB: not implemented"
	return false, nil
}

// processRefGroupOrClusterGroup processes the Group/ClusterGroup reference present in the rule and returns the
// NetworkPolicyPeer with the corresponding AddressGroup or IPBlock.
func (n *NetworkPolicyController) processRefGroupOrClusterGroup(g, namespace string) (*antreatypes.AddressGroup, []controlplane.IPBlock) {
	_ = "STUB: not implemented"
	// Namespaced Group uses NAMESPACE/NAME as the key of the corresponding internal group while ClusterGroup uses Name.
	return nil, nil
}

// Find the internal Group corresponding to this ClusterGroup

// Internal Group was not found. Once the internal Group is created, the sync
// worker for internal group will re-enqueue the ClusterNetworkPolicy processing
// which will trigger the creation of AddressGroup.

// The Group/ClusterGroup referred in the rule might have childGroups defined using selectors
// or ipBlocks (or both). An addressGroup needs to be created as long as there is at least
// one childGroup defined by selectors, or the Group/ClusterGroup itself is defined by selectors.
// In case of updates, the original addressGroup created will be de-referenced and cleaned
// up if the Group/ClusterGroup becomes ipBlocks-only.
