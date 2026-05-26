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

package memberlist

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/hashicorp/memberlist"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/consistenthash"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
	crdlister "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
)

const (
	controllerName = "MemberListCluster"
	// Set resyncPeriod to 0 to disable resyncing.
	resyncPeriod time.Duration = 0
	// Set default virtual node replicas num of consistent hash
	// in order to improve the quality of the hash distribution, refs https://github.com/golang/groupcache/issues/29
	defaultVirtualNodeReplicas = 50
	// How long to wait before retrying the processing of an ExternalIPPool change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing an ExternalIPPool change.
	defaultWorkers = 4

	nodeEventTypeJoin   nodeEventType = "Join"
	nodeEventTypeLeave  nodeEventType = "Leave"
	nodeEventTypeUpdate nodeEventType = "Update"

	allNodesConsistentHashMapKey = ""
)

// ErrNoNodeAvailable is the error returned if no Node is chosen in SelectNodeForIP and ShouldSelectIP.
var ErrNoNodeAvailable = errors.New("no Node available")

type nodeEventType string

// Default Hash Fn is crc32.ChecksumIEEE.
var defaultHashFn func(data []byte) uint32

var (
	errDecodingObject          = fmt.Errorf("received unexpected object")
	errDecodingObjectTombstone = fmt.Errorf("deletedFinalStateUnknown contains unexpected object")
)

var mapNodeEventType = map[memberlist.NodeEventType]nodeEventType{
	memberlist.NodeJoin:   nodeEventTypeJoin,
	memberlist.NodeLeave:  nodeEventTypeLeave,
	memberlist.NodeUpdate: nodeEventTypeUpdate,
}

var linuxNodeSelector = labels.SelectorFromSet(labels.Set{corev1.LabelOSStable: "linux"})

type ClusterNodeEventHandler func(objName string)

type Interface interface {
	ShouldSelectIP(ip string, pool string, filters ...func(node string) bool) (bool, error)
	SelectNodeForIP(ip, externalIPPool string, filters ...func(string) bool) (string, error)
	AliveNodes() sets.Set[string]
	AddClusterEventHandler(handler ClusterNodeEventHandler)
}

type Memberlist interface {
	Join(existing []string) (int, error)
	Members() []*memberlist.Node
	Leave(timeout time.Duration) error
	Shutdown() error
}

// Cluster implements ClusterInterface.
type Cluster struct {
	bindPort int
	// Name of local Node. Node name must be unique in the cluster.
	nodeName string

	mList Memberlist
	// consistentHash hold the consistentHashMap, when a Node join cluster, use method Add() to add a key to the hash.
	// when a Node leave the cluster, the consistentHashMap should be update.
	consistentHashMap     map[string]*consistenthash.Map
	consistentHashRWMutex sync.RWMutex
	// nodeEventsCh, the Node join/leave events will be notified via it.
	nodeEventsCh chan memberlist.NodeEvent

	// clusterNodeEventHandlers contains eventHandler which will run when consistentHashMap is updated,
	// which caused by an ExternalIPPool or Node event, such as cluster Node status update(leave of join cluster),
	// ExternalIPPool events(create/update/delete).
	// For example, when a new Node joins the cluster, each Node should compute whether it should still hold all
	// its existing Egresses, and when a Node leaves the cluster,
	// each Node should check whether it is now responsible for some of the Egresses from that Node.
	clusterNodeEventHandlers []ClusterNodeEventHandler

	nodeInformer     coreinformers.NodeInformer
	nodeLister       corelisters.NodeLister
	nodeListerSynced cache.InformerSynced

	externalIPPoolInformer          cache.SharedIndexInformer
	externalIPPoolLister            crdlister.ExternalIPPoolLister
	externalIPPoolInformerHasSynced cache.InformerSynced

	// queue maintains the ExternalIPPool names that need to be synced.
	queue workqueue.TypedRateLimitingInterface[string]
}

// NewCluster returns a new *Cluster.
func NewCluster(
	nodeIP net.IP,
	clusterBindPort int,
	nodeName string,
	nodeInformer coreinformers.NodeInformer,
	externalIPPoolInformer crdinformers.ExternalIPPoolInformer,
	ml Memberlist, // Parameterized for testing, could be left nil for production code.
) (*Cluster, error) {
	_ = "STUB: not implemented"
	// The Node join/leave events will be notified via it.
	return nil, nil
}

// Setting it to a non-zero value to allow reclaiming Nodes with different addresses for Node IP update case.

func shouldJoinCluster(node *corev1.Node) bool {
	_ = "STUB: not implemented"
	// non-Linux Nodes should not join the memberlist cluster as all features relying on it is only supported on Linux.
	return false
}

func (c *Cluster) handleCreateNode(obj interface{}) { _ = "STUB: not implemented"; return }

// Ignore the Node itself.

func (c *Cluster) handleDeleteNode(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Cluster) handleUpdateNode(oldObj, newObj interface{}) { _ = "STUB: not implemented"; return }

func (c *Cluster) enqueueExternalIPPools(eips sets.Set[string]) { _ = "STUB: not implemented"; return }

func (c *Cluster) enqueueExternalIPPool(obj interface{}) { _ = "STUB: not implemented"; return }

// newClusterMember gets the Node's IP and returns it as a cluster member for memberlist cluster to join.
func (c *Cluster) newClusterMember(node *corev1.Node) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Cluster) filterEIPsFromNodeLabels(node *corev1.Node) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// Run will join all the other K8s Nodes in a memberlist cluster
// and will create defaultWorkers workers (go routines) which will process the ExternalIPPool or Node events
// from the work queue.
func (c *Cluster) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// In order to exit the cluster more gracefully, call Leave prior to shutting down.

// Rejoin Nodes periodically in case some Nodes are removed from the member list because of long downtime.

// RejoinNodes rejoins Nodes that were removed from the member list by memberlist because they were unreachable for more
// than 15 seconds (the GossipToTheDeadTime we are using). Without it, once there is a network downtime lasting more
// than 15 seconds, the agent wouldn't try to reach any other Node and would think it's the only alive Node until it's
// restarted.
func (c *Cluster) RejoinNodes() { _ = "STUB: not implemented"; return }

// Every known Node is alive, do nothing.

// The Join method returns an error only when none could be reached.

func (c *Cluster) worker() { _ = "STUB: not implemented"; return }

func (c *Cluster) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the work queue to handle any transient errors.

func (c *Cluster) syncConsistentHash(eipName string) error { _ = "STUB: not implemented"; return nil }

// updateConsistentHash refreshes the consistentHashMap.

// Node alive and Node labels match ExternalIPPool nodeSelector.

func NewNodeConsistentHashMap() *consistenthash.Map { _ = "STUB: not implemented"; return nil }

func (c *Cluster) handleClusterNodeEvents(nodeEvent *memberlist.NodeEvent) {
	_ = "STUB: not implemented"
	return
}

// When a Node joins cluster, all matched ExternalIPPools consistentHash should be updated;
// when a Node leaves cluster, the Node may have failed or have been deleted,
// if the Node has been deleted, affected ExternalIPPool should be enqueued, and deleteNode handler has been executed,
// if the Node has failed, ExternalIPPools consistentHash maybe changed, and affected ExternalIPPool should be enqueued.

// It means the Node has been deleted, no further processing is needed as handleDeleteNode has enqueued
// related ExternalIPPools.

// AliveNodes returns the list of nodeNames in the cluster.
func (c *Cluster) AliveNodes() sets.Set[string] { _ = "STUB: not implemented"; return nil }

// ShouldSelectIP returns true if the local Node is selected as the owner Node of the IP in the specific
// ExternalIPPool. The local Node in the cluster holds the same consistent hash ring for each ExternalIPPool,
// consistentHash.Get gets the closest item (Node name) in the hash to the provided key (IP), if the name of
// the local Node is equal to the name of the selected Node, returns true.
func (c *Cluster) ShouldSelectIP(ip, externalIPPool string, filters ...func(string) bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SelectNodeForIP returns the closest item (Node name) in the hash to the provided key (IP) and ExternalIPPool.
func (c *Cluster) SelectNodeForIP(ip, externalIPPool string, filters ...func(string) bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Cluster) notify(objName string) { _ = "STUB: not implemented"; return }

// AddClusterEventHandler adds a clusterNodeEventHandler, which will run when consistentHashMap is updated,
// due to an ExternalIPPool or Node event.
func (c *Cluster) AddClusterEventHandler(handler ClusterNodeEventHandler) {
	_ = "STUB: not implemented"
	return
}
