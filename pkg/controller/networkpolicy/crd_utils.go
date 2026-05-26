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
	"net"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	antreatypes "antrea.io/antrea/v2/pkg/controller/types"
)

var (
	// matchAllPodsPeerCrd is a crdv1beta1.NetworkPolicyPeer matching all
	// Pods from all Namespaces.
	matchAllPodsPeerCrd = crdv1beta1.NetworkPolicyPeer{
		NamespaceSelector: &metav1.LabelSelector{},
	}
)

// semanticIgnoreLastTransitionTime does semantic deep equality checks for
// NetworkPolicyCondition but excludes LastTransitionTime. They are used when
// comparing NetworkPolicyCondition in NetworkPolicyStatus objects to avoid
// unnecessary updates caused different status generation time.
var semanticIgnoreLastTransitionTime = conversion.EqualitiesOrDie(
	func(a, b crdv1beta1.NetworkPolicyCondition) bool {
		a.LastTransitionTime = metav1.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
		b.LastTransitionTime = metav1.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
		return a == b
	},
)

// NetworkPolicyStatusEqual compares two NetworkPolicyStatus objects. It disregards
// the LastTransitionTime field in the status Conditions.
func NetworkPolicyStatusEqual(oldStatus, newStatus crdv1beta1.NetworkPolicyStatus) bool {
	_ = "STUB: not implemented"
	return false
}

// groupMembersComputedConditionEqual checks whether the condition status for GroupMembersComputed condition
// is same. Returns true if equal, otherwise returns false. It disregards the lastTransitionTime field.
func groupMembersComputedConditionEqual(conds []crdv1beta1.GroupCondition, condition crdv1beta1.GroupCondition) bool {
	_ = "STUB: not implemented"
	return false
}

// toAntreaServicesForCRD converts a slice of crdv1beta1.NetworkPolicyPort objects
// and a slice of v1beta1.NetworkPolicyProtocol objects to a slice of Antrea
// Service objects. A bool is returned along with the Service objects to indicate
// whether any named port exists.
func toAntreaServicesForCRD(npPorts []crdv1beta1.NetworkPolicyPort, npProtocols []crdv1beta1.NetworkPolicyProtocol) ([]controlplane.Service, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// toAntreaL7ProtocolsForCRD converts a slice of v1beta1.L7Protocol objects to
// a slice of Antrea L7Protocol objects.
func toAntreaL7ProtocolsForCRD(l7Protocols []crdv1beta1.L7Protocol) []controlplane.L7Protocol {
	_ = "STUB: not implemented"
	return nil
}

// toAntreaIPBlockForCRD converts a crdv1beta1.IPBlock to an Antrea IPBlock.
func toAntreaIPBlockForCRD(ipBlock *crdv1beta1.IPBlock) (*controlplane.IPBlock, error) {
	_ = "STUB: not implemented"
	// Convert the allowed IPBlock to networkpolicy.IPNet.
	return nil, nil
}

// Convert the except IPBlock to networkpolicy.IPNet.

// computeEffectiveIPNetForIPBlocks calculates the list of net.IPNet CIDRs after the
// "except" CIDRs are subtracted from each corresponding ipBlock.
func computeEffectiveIPNetForIPBlocks(ipBlocks []crdv1beta1.IPBlock) []*net.IPNet {
	_ = "STUB: not implemented"
	return nil
}

// CIDR format is already validated by the webhook

// This should not happen theoretically since the except CIDRs are all validated
// to be a subnet of the ipBlock.CIDR

// toAntreaPeerForCRD creates an Antrea controlplane NetworkPolicyPeer for crdv1beta1 NetworkPolicyPeer.
// It is used when peer's Namespaces are not matched by NamespaceMatchTypes, for which the controlplane
// NetworkPolicyPeers will need to be created on a per-Namespace basis.
// Any ClusterSet scoped selector in this peer will also be registered with the labelIdentityInterface
// for the policy.
func (n *NetworkPolicyController) toAntreaPeerForCRD(peers []crdv1beta1.NetworkPolicyPeer,
	np metav1.Object, dir controlplane.Direction, namedPortExists bool) (*controlplane.NetworkPolicyPeer, []*antreatypes.AddressGroup, sets.Set[string]) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// NetworkPolicyPeer is supposed to match all addresses when it is empty and no clusterGroup is present.
// It's treated as an IPBlock "0.0.0.0/0".

// For an egress Peer that specifies any named ports, it creates or
// reuses the AddressGroup matching all Pods in all Namespaces and
// appends the AddressGroup UID to the returned Peer such that it can be
// used to resolve the named ports.
// For other cases it uses the IPBlock "0.0.0.0/0" to avoid the overhead
// of handling member updates of the AddressGroup.

// A crdv1beta1.NetworkPolicyPeer will have exactly one of the following fields set:
// - podSelector and/or namespaceSelector (in-cluster scope or ClusterSet scope)
// - reference to a Group/ClusterGroup
// - IPBlocks
// - FQDNs

// In addition to getting the matched Label Identity IDs, AddSelector also registers the selector
// with the labelIdentityInterface.

// toNamespacedPeerForCRD creates an Antrea controlplane NetworkPolicyPeer for crdv1beta1 NetworkPolicyPeer
// for a particular Namespace. It is used when a single crdv1beta1 NetworkPolicyPeer maps to multiple
// controlplane NetworkPolicyPeers because the appliedTo workloads reside in different Namespaces.
func (n *NetworkPolicyController) toNamespacedPeerForCRD(peers []crdv1beta1.NetworkPolicyPeer,
	np metav1.Object, namespace string) (*controlplane.NetworkPolicyPeer, []*antreatypes.AddressGroup, sets.Set[string]) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// In addition to getting the matched Label Identity IDs, AddSelector also registers the selector
// with the labelIdentityInterface.

// svcRefToPeerForCRD creates an Antrea controlplane NetworkPolicyPeer from ServiceReferences in ToServices
// or ToMulticlusterServices field of a crdv1beta1 NetworkPolicyPeer. For ANNP NetworkPolicyPeers, if
// Namespace is not provided in the ServiceReference, the policy's Namespace will be assumed.
func (n *NetworkPolicyController) svcRefToPeerForCRD(svcRefs []crdv1beta1.PeerService, defaultNamespace string) *controlplane.NetworkPolicyPeer {
	_ = "STUB: not implemented"
	return nil
}

// createAppliedToGroupForService creates an AppliedToGroup object corresponding to a Service.
func (n *NetworkPolicyController) createAppliedToGroupForService(service *crdv1beta1.NamespacedName) *antreatypes.AppliedToGroup {
	_ = "STUB: not implemented"
	return nil
}

// Create an AppliedToGroup object for this Service.

// createAppliedToGroupForGroup creates an AppliedToGroup object corresponding to a ClusterGroup or a Group.
// The namespace parameter is only provided when the group is namespace scoped.
func (n *NetworkPolicyController) createAppliedToGroupForGroup(namespace, group string) *antreatypes.AppliedToGroup {
	_ = "STUB: not implemented"
	// Cluster group uses NAME and Namespaced group uses NAMESPACE/NAME as the key of the corresponding internal group.
	return nil
}

// Find the internal Group corresponding to this ClusterGroup/Group.
// There is no need to check if the ClusterGroup/Group exists in clusterGroupLister/groupLister because its
// existence will eventually be reflected in internalGroupStore.

// Internal Group was not found. Once the internal Group is created, the sync worker for internal group will
// re-enqueue the ClusterNetworkPolicy/AntreaNetworkPolicy processing which will call this method again. So it's
// fine to ignore NotFound case.

// A Group may have child Groups, some of which contain regular Pod selectors and some of which contain IPBlocks.
// When the Group is used as AppliedTo, it seems obvious that we should just apply NetworkPolicy to the selected
// Pods and ignore the IPBlocks, instead of reporting errors and asking users to remove IPBlocks from child Groups,
// as the Group could also be used as AddressGroup.
// To keep the behavior consistent regarding IPBlocks, we ignore Groups containing only IPBlocks when it's used as
// AppliedTo.

// getTierPriority retrieves the priority associated with the input Tier name.
// If the Tier name is empty, by default, the lowest priority Application Tier
// is returned.
func (n *NetworkPolicyController) getTierPriority(tier string) int32 {
	_ = "STUB: not implemented"
	return 0
}

// If the tier name is part of the static tier name set, we need to convert
// tier name to lowercase to match the corresponding Tier CRD name. This is
// possible in case of upgrade where in a previously created Antrea Policy
// CRD was referring to an old static tier. Static tiers were introduced in
// release 0.9.0 and deprecated in 0.10.0. So any upgrade from 0.9.0 to a
// later release will undergo this conversion.

// This error should ideally not occur as we perform validation.

// getNormalizedNameForSelector retrieves the normalized name for GroupSelector.
// If the GroupSelector is nil, an empty string is returned.
func getNormalizedNameForSelector(sel *antreatypes.GroupSelector) string {
	_ = "STUB: not implemented"
	return ""
}

func (n *NetworkPolicyController) syncInternalGroup(key string) error {
	_ = "STUB: not implemented"
	return nil
}

// Retrieve the internal Group corresponding to this key.

// Sync the Group as a Namespaced Group.
