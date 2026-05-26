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
	"antrea.io/antrea/v2/pkg/apis/controlplane"
	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	antreatypes "antrea.io/antrea/v2/pkg/controller/types"
)

func getANNPReference(annp *crdv1beta1.NetworkPolicy) *controlplane.NetworkPolicyReference {
	_ = "STUB: not implemented"
	return nil
}

// addANNP receives AntreaNetworkPolicy ADD events and enqueues a reference of
// the AntreaNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) addANNP(obj interface{}) { _ = "STUB: not implemented"; return }

// updateANNP receives AntreaNetworkPolicy UPDATE events and enqueues a reference
// of the AntreaNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) updateANNP(old, cur interface{}) {
	_ = "STUB: not implemented"
	return
}

// deleteANNP receives AntreaNetworkPolicy DELETE events and enqueues a reference
// of the AntreaNetworkPolicy to trigger its process.
func (n *NetworkPolicyController) deleteANNP(old interface{}) { _ = "STUB: not implemented"; return }

// processAntreaNetworkPolicy creates an internal NetworkPolicy instance
// corresponding to the crdv1beta1.NetworkPolicy object. This method
// does not commit the internal NetworkPolicy in store, instead returns an
// instance to the caller.
func (n *NetworkPolicyController) processAntreaNetworkPolicy(np *crdv1beta1.NetworkPolicy) (*antreatypes.NetworkPolicy, map[string]*antreatypes.AppliedToGroup, map[string]*antreatypes.AddressGroup) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// appliedToGroups tracks all distinct appliedToGroups referred to by the Antrea NetworkPolicy,
// either in the spec section or in ingress/egress rules.
// The span calculation and stale appliedToGroup cleanup logic would work seamlessly for both cases.

// clusterSetScopeSelectorKeys keeps track of all the ClusterSet-scoped selector keys of the policy.
// During policy peer processing, any ClusterSet-scoped selector will be registered with the
// labelIdentityInterface and added to this set. By the end of the function, this set will
// be used to remove any stale selector from the policy in the labelIdentityInterface.

// Create AppliedToGroup for each AppliedTo present in AntreaNetworkPolicy spec.

// Compute NetworkPolicyRule for Ingress Rule.

// Set default action to ALLOW to allow traffic.

// Create AppliedToGroup for each AppliedTo present in the ingress rule.

// Compute NetworkPolicyRule for Egress Rule.

// Set default action to ALLOW to allow traffic.

// Create AppliedToGroup for each AppliedTo present in the egress rule.

func (n *NetworkPolicyController) processAppliedTo(namespace string, appliedTo []crdv1beta1.AppliedTo) []*antreatypes.AppliedToGroup {
	_ = "STUB: not implemented"
	return nil
}

// ErrNetworkPolicyAppliedToUnsupportedGroup is an error response when
// a Group with Pods in other Namespaces is used as AppliedTo.
type ErrNetworkPolicyAppliedToUnsupportedGroup struct {
	namespace string
	groupName string
}

func (e *ErrNetworkPolicyAppliedToUnsupportedGroup) Error() string {
	_ = "STUB: not implemented"
	return ""
}
