// Copyright 2021 Antrea Authors
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

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/controller/grouping"
	antreatypes "antrea.io/antrea/v2/pkg/controller/types"
)

// addClusterGroup is responsible for processing the ADD event of a ClusterGroup resource.
func (c *NetworkPolicyController) addClusterGroup(curObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// updateClusterGroup is responsible for processing the UPDATE event of a ClusterGroup resource.
func (c *NetworkPolicyController) updateClusterGroup(oldObj, curObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// No change in the contents of the ClusterGroup. No need to enqueue for further sync.

// deleteClusterGroup is responsible for processing the DELETE event of a ClusterGroup resource.
func (c *NetworkPolicyController) deleteClusterGroup(oldObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *NetworkPolicyController) processClusterGroup(cg *crdv1beta1.ClusterGroup) *antreatypes.Group {
	_ = "STUB: not implemented"
	return nil
}

// ServiceReference will be converted to groupSelector once the internalGroup is synced.

// filterInternalGroupsForService computes a list of internal Group keys which references the Service.
func (c *NetworkPolicyController) filterInternalGroupsForService(obj metav1.Object) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

func (c *NetworkPolicyController) enqueueInternalGroup(key string) {
	_ = "STUB: not implemented"
	return
}

func (c *NetworkPolicyController) internalGroupWorker() { _ = "STUB: not implemented"; return }

// Processes an item in the "internalGroup" work queue, by calling
// syncInternalGroup after casting the item to a string (Group key).
// If syncInternalGroup returns an error, this function handles it by re-queueing
// the item so that it can be processed again later. If syncInternalGroup is
// successful, the ClusterGroup is removed from the queue until we get notify
// of a new change. This function return false if and only if the work queue
// was shutdown (no more items will be processed).
func (c *NetworkPolicyController) processNextInternalGroupWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// Put the item back in the workqueue to handle any transient errors.

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

func (c *NetworkPolicyController) syncInternalClusterGroup(grp *antreatypes.Group) error {
	_ = "STUB: not implemented"
	return nil
}

// Retrieve the ClusterGroup corresponding to this key.

// Update the ClusterGroup status to Realized as Antrea has recognized the Group and
// processed its group members. The ClusterGroup is considered realized if:
//   1. It does not have child groups. The group members are immediately considered
//      computed during syncInternalGroup, as the group selector is finalized.
//   2. All its child groups are created and realized.

// Update the internal Group object in the store with the new selector and status.

func getClusterGroupSourceRef(cg *crdv1beta1.ClusterGroup) *controlplane.GroupReference {
	_ = "STUB: not implemented"
	return nil
}

func (c *NetworkPolicyController) triggerParentGroupUpdates(grp string) {
	_ = "STUB: not implemented"
	// TODO: if the max supported group nesting level increases, a Group having children
	//  will no longer be a valid indication that it cannot have parents.
	return
}

// triggerDerivedGroupUpdates triggers processing of AppliedToGroup and AddressGroup derived from the provided group.
func (c *NetworkPolicyController) triggerDerivedGroupUpdates(grp string) {
	_ = "STUB: not implemented"
	return
}

// It's fine if the group is deleted after checking its existence as syncAppliedToGroup will do nothing when it
// doesn't find the group.

// It's fine if the group is deleted after checking its existence as syncAddressGroup will do nothing when it
// doesn't find the group.

// triggerCNPUpdates triggers processing of ClusterNetworkPolicies associated with the input ClusterGroup.
func (c *NetworkPolicyController) triggerCNPUpdates(cg string) {
	_ = "STUB: not implemented"
	// If a ClusterGroup is added/updated, it might have a reference in ClusterNetworkPolicy.
	return
}

// updateClusterGroupStatus updates the Status subresource for a ClusterGroup.
func (c *NetworkPolicyController) updateClusterGroupStatus(cg *crdv1beta1.ClusterGroup, cStatus v1.ConditionStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// There is no change in conditions.

// processServiceReference knows how to process the serviceReference in the group, and set the group
// selector based on the Service referenced. It returns true if the group's selector needs to be
// updated after serviceReference processing, and false otherwise.
func (c *NetworkPolicyController) processServiceReference(group *antreatypes.Group) bool {
	_ = "STUB: not implemented"
	return false
}

// serviceToGroupSelector knows how to generate GroupSelector for a Service.
func (c *NetworkPolicyController) serviceToGroupSelector(service *v1.Service) *antreatypes.GroupSelector {
	_ = "STUB: not implemented"
	return nil
}

// Convert Service.spec.selector to GroupSelector by setting the Namespace to the Service's Namespace
// and podSelector to Service's selector.

// getGroupsForNode returns the set of groups associated with the given node or false
// if there aren't any
func (c *NetworkPolicyController) getGroupsForNode(name string) (map[grouping.GroupType][]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Error ignored safely as groups will remain empty and return false

// GetAssociatedGroups retrieves the internal Groups associated with the entity being
// queried (Pod, Node or ExternalEntity identified by name and namespace).
func (c *NetworkPolicyController) GetAssociatedGroups(name, namespace string) []antreatypes.Group {
	_ = "STUB: not implemented"
	return nil
}

// Try Pod first, then ExternalEntity.

// Remove duplicates in the groupObj slice.

// getAssociatedGroupsByName retrieves the internal Group and all it's parent Group objects
// (if any) by Group name.
func (c *NetworkPolicyController) getAssociatedGroupsByName(grpName string) []antreatypes.Group {
	_ = "STUB: not implemented"
	return nil
}

func (c *NetworkPolicyController) getParentGroups(grpName string) []antreatypes.Group {
	_ = "STUB: not implemented"
	return nil
}

// GetGroupMembers returns the current members of a ClusterGroup/Group.
// If the ClusterGroup/Group is defined with IPBlocks, the returned members will be []controlplane.IPBlock.
// Otherwise, the returned members will be of type controlplane.GroupMemberSet.
func (c *NetworkPolicyController) GetGroupMembers(name string) (controlplane.GroupMemberSet, []controlplane.IPBlock, error) {
	_ = "STUB: not implemented"
	return *new(controlplane.GroupMemberSet), nil, nil
}

func (c *NetworkPolicyController) GetAssociatedIPBlockGroups(ip net.IP) []antreatypes.Group {
	_ = "STUB: not implemented"
	return nil
}

// Append all parent groups to matchedGroups
