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
	v1 "k8s.io/api/core/v1"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	antreatypes "antrea.io/antrea/v2/pkg/controller/types"
)

// addGroup is responsible for processing the ADD event of a Group resource.
func (n *NetworkPolicyController) addGroup(curObj interface{}) { _ = "STUB: not implemented"; return }

// updateGroup is responsible for processing the UPDATE event of a Group resource.
func (n *NetworkPolicyController) updateGroup(oldObj, curObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// No change in the contents of the Group. No need to enqueue for further sync.

// deleteGroup is responsible for processing the DELETE event of a Group resource.
func (n *NetworkPolicyController) deleteGroup(oldObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (n *NetworkPolicyController) processGroup(g *crdv1beta1.Group) *antreatypes.Group {
	_ = "STUB: not implemented"
	return nil
}

// ServiceReference will be converted to groupSelector once the internalGroup is synced.

func getGroupSourceRef(g *crdv1beta1.Group) *controlplane.GroupReference {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkPolicyController) syncInternalNamespacedGroup(grp *antreatypes.Group) error {
	_ = "STUB: not implemented"
	return nil
}

// Retrieve the Group corresponding to this key.

// Update the Group status to Realized as Antrea has recognized the Group and
// processed its group members. The Group is considered realized if:
//   1. It does not have child groups. The group members are immediately considered
//      computed during syncInternalGroup, as the group selector is finalized.
//   2. All its child groups are created and realized.

// Update the internal Group object in the store with the new selector and status.

// triggerANNPUpdates triggers processing of Antrea NetworkPolicies associated with the input Group.
func (n *NetworkPolicyController) triggerANNPUpdates(g string) {
	_ = "STUB: not implemented"
	// If a Group is added/updated, it might have a reference in Antrea NetworkPolicy.
	return
}

// updateGroupStatus updates the Status subresource for a Group.
func (n *NetworkPolicyController) updateGroupStatus(g *crdv1beta1.Group, cStatus v1.ConditionStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// There is no change in conditions.
