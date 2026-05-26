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

package networkpolicy

import (
	"sigs.k8s.io/network-policy-api/apis/v1alpha1"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	antreacrd "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	antreatypes "antrea.io/antrea/v2/pkg/controller/types"
)

var (
	adminNetworkPolicyTierPriority = int32(251)
	banpTierPriority               = int32(254)
	banpPriority                   = float64(1)

	anpActionToAntreaActionMap = map[v1alpha1.AdminNetworkPolicyRuleAction]antreacrd.RuleAction{
		v1alpha1.AdminNetworkPolicyRuleActionAllow: antreacrd.RuleActionAllow,
		v1alpha1.AdminNetworkPolicyRuleActionDeny:  antreacrd.RuleActionDrop,
		v1alpha1.AdminNetworkPolicyRuleActionPass:  antreacrd.RuleActionPass,
	}

	banpActionToAntreaActionMap = map[v1alpha1.BaselineAdminNetworkPolicyRuleAction]antreacrd.RuleAction{
		v1alpha1.BaselineAdminNetworkPolicyRuleActionAllow: antreacrd.RuleActionAllow,
		v1alpha1.BaselineAdminNetworkPolicyRuleActionDeny:  antreacrd.RuleActionDrop,
	}
)

func getAdminNPReference(anp *v1alpha1.AdminNetworkPolicy) *controlplane.NetworkPolicyReference {
	_ = "STUB: not implemented"
	return nil
}

func getBANPReference(banp *v1alpha1.BaselineAdminNetworkPolicy) *controlplane.NetworkPolicyReference {
	_ = "STUB: not implemented"
	return nil
}

// addAdminNP receives AdminNetworkPolicy ADD events and enqueues a reference of
// the AdminNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) addAdminNP(obj interface{}) { _ = "STUB: not implemented"; return }

// updateAdminNP receives AdminNetworkPolicy UPDATE events and enqueues a
// reference of the AdminNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) updateAdminNP(_, cur interface{}) {
	_ = "STUB: not implemented"
	return
}

// deleteAdminNP receives AdminNetworkPolicy DELETE events and enqueues a
// reference of the AdminNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) deleteAdminNP(old interface{}) { _ = "STUB: not implemented"; return }

// addBANP receives BaselineAdminNetworkPolicy ADD events and enqueues a reference of
// the BaselineAdminNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) addBANP(obj interface{}) { _ = "STUB: not implemented"; return }

// updateBANP receives BaselineAdminNetworkPolicy UPDATE events and enqueues a
// reference of the BaselineAdminNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) updateBANP(_, cur interface{}) { _ = "STUB: not implemented"; return }

// deleteBANP receives BaselineAdminNetworkPolicy DELETE events and enqueues a
// reference of the BaselineAdminNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) deleteBANP(old interface{}) { _ = "STUB: not implemented"; return }

// anpHasNamespaceLabelRule returns whether an AdminNetworkPolicy has rules defined by
// advanced Namespace selection (sameLabels and notSameLabels)
func anpHasNamespaceLabelRule(anp *v1alpha1.AdminNetworkPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

// banpHasNamespaceLabelRule returns whether a BaselineAdminNetworkPolicy has rules defined by
// advanced Namespace selection (sameLabels and notSameLabels)
func banpHasNamespaceLabelRule(banp *v1alpha1.BaselineAdminNetworkPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

// toAntreaServicesForPolicyCRD processes ports field for ANPs/BANPs and returns the translated
// Antrea Services.
func toAntreaServicesForPolicyCRD(npPorts []v1alpha1.AdminNetworkPolicyPort) []controlplane.Service {
	_ = "STUB: not implemented"
	return nil
}

// splitPolicyPeersByScope splits the ANP/BANP peers in the rule by whether the peer is cluster scoped
// or per-namespace scoped. Per-namespace peers are those whose defined by sameLabels and
// notSameLabels.
func splitPolicyPeerByScope(peers []v1alpha1.AdminNetworkPolicyPeer) ([]v1alpha1.AdminNetworkPolicyPeer, []v1alpha1.AdminNetworkPolicyPeer) {
	_ = "STUB: not implemented"
	return nil, nil
}

// toAntreaPeerForPolicyCRD processes AdminNetworkPolicyPeers and yield Antrea NetworkPolicyPeers.
func (n *NetworkPolicyController) toAntreaPeerForPolicyCRD(peers []v1alpha1.AdminNetworkPolicyPeer) (*controlplane.NetworkPolicyPeer, []*antreatypes.AddressGroup) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processClusterSubject processes AdminNetworkPolicySubject and yield Antrea AppliedToGroups.
func (n *NetworkPolicyController) processClusterSubject(subject v1alpha1.AdminNetworkPolicySubject) []*antreatypes.AppliedToGroup {
	_ = "STUB: not implemented"
	return nil
}

func anpActionToCRDAction(action v1alpha1.AdminNetworkPolicyRuleAction) *antreacrd.RuleAction {
	_ = "STUB: not implemented"
	return nil
}

func banpActionToCRDAction(action v1alpha1.BaselineAdminNetworkPolicyRuleAction) *antreacrd.RuleAction {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkPolicyController) processAdminNetworkPolicy(anp *v1alpha1.AdminNetworkPolicy) (*antreatypes.NetworkPolicy, map[string]*antreatypes.AppliedToGroup, map[string]*antreatypes.AddressGroup) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//TODO: implement SameLabels and NotSameLabels for per NS label ingress peers

//TODO: implement SameLabels and NotSameLabels for per NS label egress peers

func (n *NetworkPolicyController) processBaselineAdminNetworkPolicy(banp *v1alpha1.BaselineAdminNetworkPolicy) (*antreatypes.NetworkPolicy, map[string]*antreatypes.AppliedToGroup, map[string]*antreatypes.AddressGroup) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//TODO: implement SameLabels and NotSameLabels for per NS label ingress peers

//TODO: implement SameLabels and NotSameLabels for per NS label egress peers
