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

package k8s

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/util/ip"
)

// GetNodeAddrsWithType gets the available IP addresses of a Node. It will consider address types
// specified in the types slice, in the provided order.
// If no error is returned, the returned DualStackIPs includes at least one IPv4 or IPv6 address.
func GetNodeAddrsWithType(node *v1.Node, types []v1.NodeAddressType) (*ip.DualStackIPs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNodeAddrs gets the available IP addresses of a Node. GetNodeAddrs will first try to get the
// NodeInternalIP, then try to get the NodeExternalIP.
// If no error is returned, the returned DualStackIPs includes at least one IPv4 or IPv6 address.
func GetNodeAddrs(node *v1.Node) (*ip.DualStackIPs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNodeAddrsFromAnnotations gets available IPs from the Node Annotation. The annotations are set by Antrea.
func GetNodeAddrsFromAnnotations(node *v1.Node, annotationKey string) (*ip.DualStackIPs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNodeGatewayAddrs gets Node Antrea gateway IPs from the Node Spec.
func GetNodeGatewayAddrs(node *v1.Node) (*ip.DualStackIPs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNodeAllAddrs gets all Node IPs from the Node.
func GetNodeAllAddrs(node *v1.Node) (ips sets.Set[string], err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetNodeTransportAddrs(node *v1.Node) (*ip.DualStackIPs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use NodeIP if the transport IP address is not set or not found.
