/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
/*
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

Modifies:
- Cleanup unused imports due to code removal and relocation
- Remove the following unused funcs to avoid additional code imports:
	GetNodeHostIPs()
	GetNodeHostIP()
	GetNodeIP()
*/

package node

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	clientset "k8s.io/client-go/kubernetes"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
)

const (
	// NodeUnreachablePodReason is the reason on a pod when its state cannot be confirmed as kubelet is unresponsive
	// on the node it is (was) running.
	NodeUnreachablePodReason = "NodeLost"
	// NodeUnreachablePodMessage is the message on a pod when its state cannot be confirmed as kubelet is unresponsive
	// on the node it is (was) running.
	NodeUnreachablePodMessage = "Node %v which was running pod %v is unresponsive"
)

// GetHostname returns OS's hostname if 'hostnameOverride' is empty; otherwise, return 'hostnameOverride'.
func GetHostname(hostnameOverride string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Trim whitespaces first to avoid getting an empty hostname
// For linux, the hostname is read from file /proc/sys/kernel/hostname directly

// NoMatchError is a typed implementation of the error interface. It indicates a failure to get a matching Node.
type NoMatchError struct {
	addresses []v1.NodeAddress
}

// Error is the implementation of the conventional interface for
// representing an error condition, with the nil value representing no error.
func (e *NoMatchError) Error() string { _ = "STUB: not implemented"; return "" }

// GetPreferredNodeAddress returns the address of the provided node, using the provided preference order.
// If none of the preferred address types are found, an error is returned.
func GetPreferredNodeAddress(node *v1.Node, preferredAddressTypes []v1.NodeAddressType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type nodeForConditionPatch struct {
	Status nodeStatusForPatch `json:"status"`
}

type nodeStatusForPatch struct {
	Conditions []v1.NodeCondition `json:"conditions"`
}

// SetNodeCondition updates specific node condition with patch operation.
func SetNodeCondition(c clientset.Interface, node types.NodeName, condition v1.NodeCondition) error {
	_ = "STUB: not implemented"
	return nil
}

type nodeForCIDRMergePatch struct {
	Spec nodeSpecForMergePatch `json:"spec"`
}

type nodeSpecForMergePatch struct {
	PodCIDR  string   `json:"podCIDR"`
	PodCIDRs []string `json:"podCIDRs,omitempty"`
}

// PatchNodeCIDR patches the specified node's CIDR to the given value.
func PatchNodeCIDR(c clientset.Interface, node types.NodeName, cidr string) error {
	_ = "STUB: not implemented"
	return nil
}

// PatchNodeCIDRs patches the specified node.CIDR=cidrs[0] and node.CIDRs to the given value.
func PatchNodeCIDRs(c clientset.Interface, node types.NodeName, cidrs []string) error {
	_ = "STUB: not implemented"
	// set the pod cidrs list and set the old pod cidr field
	return nil
}

// PatchNodeStatus patches node status.
func PatchNodeStatus(c v1core.CoreV1Interface, nodeName types.NodeName, oldNode *v1.Node, newNode *v1.Node) (*v1.Node, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func preparePatchBytesforNodeStatus(nodeName types.NodeName, oldNode *v1.Node, newNode *v1.Node) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NodeStatus.Addresses is incorrectly annotated as patchStrategy=merge, which
// will cause strategicpatch.CreateTwoWayMergePatch to create an incorrect patch
// if it changed.

// Reset spec to make sure only patch for Status or ObjectMeta is generated.
// Note that we don't reset ObjectMeta here, because:
// 1. This aligns with Nodes().UpdateStatus().
// 2. Some component does use this to update node annotations.

// fixupPatchForNodeStatusAddresses adds a replace-strategy patch for Status.Addresses to
// the existing patch
func fixupPatchForNodeStatusAddresses(patchBytes []byte, addresses []v1.NodeAddress) ([]byte, error) {
	_ = "STUB: not implemented"
	// Given patchBytes='{"status": {"conditions": [ ... ], "phase": ...}}' and
	// addresses=[{"type": "InternalIP", "address": "10.0.0.1"}], we need to generate:
	//
	//   {
	//     "status": {
	//       "conditions": [ ... ],
	//       "phase": ...,
	//       "addresses": [
	//         {
	//           "type": "InternalIP",
	//           "address": "10.0.0.1"
	//         },
	//         {
	//           "$patch": "replace"
	//         }
	//       ]
	//     }
	//   }
	return nil, nil
}
