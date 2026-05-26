/*
Copyright 2022 The Kubernetes Authors.

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
// Copyright 2025 Antrea Authors
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

Original file https://raw.githubusercontent.com/kubernetes/kubernetes/refs/tags/v1.34.2/pkg/proxy/node.go

Modifies:

- Replace import utilnode "k8s.io/kubernetes/pkg/util/node" with utilnode "antrea.io/antrea/v2/third_party/util/node".

*/

package proxy

import (
	"context"
	"net"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	v1informers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
)

// NodeManager handles the life cycle of kube-proxy based on the NodeIPs and PodCIDRs handles
// node watch events and crashes kube-proxy if there are any changes in NodeIPs or PodCIDRs.
// Note: It only crashes on change on PodCIDR when watchPodCIDRs is set to true.
type NodeManager struct {
	nodeInformer  v1informers.NodeInformer
	nodeLister    corelisters.NodeLister
	exitFunc      func(exitCode int)
	watchPodCIDRs bool

	// These are constant after construct time
	nodeIPs  []net.IP
	podCIDRs []string

	mu   sync.Mutex
	node *v1.Node
}

// NewNodeManager initializes node informer that selects for the given node, waits for cache sync
// and returns NodeManager after waiting some amount of time for the node object to exist
// and have NodeIPs (and PodCIDRs if watchPodCIDRs is true). Note: for backward compatibility,
// NewNodeManager doesn't return any error if it failed to retrieve NodeIPs and watchPodCIDRs
// is false.
func NewNodeManager(ctx context.Context, client clientset.Interface,
	resyncInterval time.Duration, nodeName string, watchPodCIDRs bool,
) (*NodeManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newNodeManager implements NewNodeManager with configurable exit function, poll interval and timeouts.
func newNodeManager(ctx context.Context, client clientset.Interface, resyncInterval time.Duration,
	nodeName string, watchPodCIDRs bool, exitFunc func(int),
	pollInterval, nodeIPsTimeout, podCIDRsTimeout time.Duration,
) (*NodeManager, error) {
	_ = "STUB: not implemented"
	// make an informer that selects for the given node
	return nil, nil
}

// initialize the informer and wait for cache sync

// wait for the node object to exist and have NodeIPs.

// wait some additional time for the PodCIDRs.

// For backward-compatibility, we keep going even if we didn't find a node (in
// non-watchPodCIDRs mode) or it didn't have IPs.

func getNodeInfo(nodeLister corelisters.NodeLister, nodeName string) (*v1.Node, []net.IP, []string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// NodeIPs returns the NodeIPs polled in NewNodeManager(). (This may be empty if
// NewNodeManager timed out without getting any IPs.)
func (n *NodeManager) NodeIPs() []net.IP {
	_ = "STUB: not implemented"

	// PodCIDRs returns the PodCIDRs polled in NewNodeManager().
	return nil
}

func (n *NodeManager) PodCIDRs() []string {
	_ = "STUB: not implemented"

	// Node returns a copy of the latest node object, or nil if the Node has not yet been seen.
	return nil
}

func (n *NodeManager) Node() *v1.Node { _ = "STUB: not implemented"; return nil }

// NodeInformer returns the NodeInformer.
func (n *NodeManager) NodeInformer() v1informers.NodeInformer {
	_ = "STUB: not implemented"
	return *

	// OnNodeChange is a handler for Node creation and update.
	new(v1informers.NodeInformer)
}

func (n *NodeManager) OnNodeChange(node *v1.Node) {
	_ = "STUB: not implemented"
	// update the node object
	return
}

// We exit whenever there is a change in PodCIDRs detected initially, and PodCIDRs received
// on node watch event if the node manager is configured with watchPodCIDRs.

// We exit whenever there is a change in NodeIPs detected initially, and NodeIPs received
// on node watch event.

// FIXME: exit
// klog.Flush()
// n.exitFunc(1)

// OnNodeDelete is a handler for Node deletes.
func (n *NodeManager) OnNodeDelete(node *v1.Node) { _ = "STUB: not implemented"; return }

// FIXME: exit
// klog.Flush()
// n.exitFunc(1)

// OnNodeSynced is called after the cache is synced and all pre-existing Nodes have been reported
func (n *NodeManager) OnNodeSynced() { _ = "STUB: not implemented"; return }
