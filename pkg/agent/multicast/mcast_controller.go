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

package multicast

import (
	"net"
	"sync"
	"time"

	apitypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/agent/types"
	"antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
	"antrea.io/antrea/v2/pkg/util/channel"
)

type eventType uint8

const (
	groupJoin eventType = iota
	groupLeave

	podInterfaceIndex = "podInterface"

	// How long to wait before retrying the processing of a multicast group change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second

	// Interval of reprocessing every node.
	nodeResyncPeriod = 60 * time.Second

	// nodeUpdateKey is a key to trigger the Node list operation and update the OpenFlow group buckets to report
	// the local multicast groups to other Nodes.
	nodeUpdateKey = "nodeUpdate"
)

var workerCount uint8 = 2

type mcastGroupEvent struct {
	group net.IP
	eType eventType
	time  time.Time
	iface *interfacestore.InterfaceConfig
	// srcNode is the Node IP where the IGMP report message is sent from. It is set only with encap mode.
	srcNode net.IP
}

type GroupMemberStatus struct {
	group net.IP
	// localMembers is a map for the local Pod member and its last update time, key is the Pod's interface name,
	// and value is its last update time.
	localMembers map[string]time.Time
	// remoteMembers is a set for Nodes which have joined the multicast group in the cluster. The Node's IP is
	// added in the set.
	remoteMembers sets.Set[string]
	ofGroupID     binding.GroupIDType
}

// eventHandler process the multicast Group membership report or leave messages.
func (c *Controller) eventHandler(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// addGroupMemberStatus adds the new group into groupCache.
func (c *Controller) addGroupMemberStatus(e *mcastGroupEvent) { _ = "STUB: not implemented"; return }

// updateGroupMemberStatus updates the group status in groupCache. If a "join" message is sent from an existing member,
// only updates the lastIGMPReport time. If a "join" message is sent from an "unknown" member, updates the lastIGMPReport time and
// adds the new member into the group's local member set. If a "leave" message is sent from an existing member, removes
// it from the group's local member set, and if the member is the last one in local cache, a query message on the group
// is sent out to check if there are still local members in the group.
func (c *Controller) updateGroupMemberStatus(obj interface{}, e *mcastGroupEvent) {
	_ = "STUB: not implemented"
	return
}

// Notify worker immediately about the member leave event if the member doesn't exist on the Node, or there are
// other local members in the multicast group.

// Check if all local members have left the multicast group.

// checkLastMember sends out a query message on the group to check if there are still members in the group. If no new
// membership report is received in the max response time, the group is removed from groupCache.
func (c *Controller) checkLastMember(group net.IP) { _ = "STUB: not implemented"; return }

// clearStaleGroups checks the stale group members which have not been updated for c.mcastGroupTimeout, and then notifies worker
// to remove them from groupCache.
func (c *Controller) clearStaleGroups() { _ = "STUB: not implemented"; return }

// Create a "leave" event for a local member if it is not updated before mcastGroupTimeout.

// removeLocalInterface searches the GroupMemberStatus which the deleted interface has joined, and then triggers a member
// leave event so that Antrea can remove the corresponding interface from local multicast receivers on OVS. This function
// should be called if the removed Pod receiver fails to send IGMP leave message before deletion.
func (c *Controller) removeLocalInterface(podEvent types.PodUpdate) {
	_ = "STUB: not implemented"
	// Ignore Pod creation event.
	return
}

type Controller struct {
	ofClient         openflow.Client
	v4GroupAllocator openflow.GroupAllocator
	ifaceStore       interfacestore.InterfaceStore
	nodeConfig       *config.NodeConfig
	igmpSnooper      *IGMPSnooper
	groupEventCh     chan *mcastGroupEvent
	groupCache       cache.Indexer
	queue            workqueue.TypedRateLimitingInterface[string]
	nodeInformer     coreinformers.NodeInformer
	nodeLister       corelisters.NodeLister
	nodeListerSynced cache.InformerSynced
	nodeUpdateQueue  workqueue.TypedRateLimitingInterface[string]
	// installedGroups saves the groups which are configured on OVS.
	// With encap mode, the entries in installedGroups include all multicast groups identified in the cluster.
	installedGroups      sets.Set[string]
	installedGroupsMutex sync.RWMutex
	// installedLocalGroups saves the groups which are configured on OVS and host. The entries in installedLocalGroups
	// include the multicast groups that local Pod members join.
	installedLocalGroups      sets.Set[string]
	installedLocalGroupsMutex sync.RWMutex
	mRouteClient              *MRouteClient
	// queryInterval is the interval to send IGMP query messages.
	queryInterval time.Duration
	// mcastGroupTimeout is the timeout to detect a group as stale if no IGMP report is received within the time.
	mcastGroupTimeout time.Duration
	// the group ID in OVS for group which IGMP queries are sent to
	queryGroupId binding.GroupIDType
	// nodeGroupID is the OpenFlow group ID in OVS which is used to send IGMP report messages to other Nodes.
	nodeGroupID binding.GroupIDType
	// installedNodes is the installed Node set that the IGMP report message is sent to.
	installedNodes      sets.Set[string]
	encapEnabled        bool
	flexibleIPAMEnabled bool
	// ipv4Enabled is the flag that if it is running on IPv4 cluster. An error is returned if IPv4Enabled is false
	// in Initialize as Multicast does not support IPv6 for now.
	// TODO: remove this flag after IPv6 is supported in Multicast.
	ipv4Enabled bool
	// ipv6Enabled is the flag that if it is running on IPv6 cluster.
	// TODO: remove this flag after IPv6 is supported in Multicast.
	ipv6Enabled bool
}

func NewMulticastController(ofClient openflow.Client,
	v4GroupAllocator openflow.GroupAllocator,
	nodeConfig *config.NodeConfig,
	ifaceStore interfacestore.InterfaceStore,
	multicastSocket RouteInterface,
	multicastInterfaces sets.Set[string],
	podUpdateSubscriber channel.Subscriber,
	igmpQueryInterval time.Duration,
	igmpQueryVersions []uint8,
	validator types.McastNetworkPolicyController,
	isEncap bool,
	nodeInformer coreinformers.NodeInformer,
	enableFlexibleIPAM bool,
	ipv4Enabled bool,
	ipv6Enabled bool) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Initialize() error { _ = "STUB: not implemented"; return nil }

// Install OpenFlow group to send the multicast groups that local Pods joined to all other Nodes in the cluster.

func (c *Controller) Run(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// Periodically query Multicast Groups on OVS.
	return
}

// Periodically check the group member status, and remove the groups in which no members exist

// Process multicast Group membership report or leave messages.

func (c *Controller) worker() { _ = "STUB: not implemented"; return }

// getGroupMemberStatusesByPod returns all GroupMemberStatus that the given podInterface is included in its localMembers.
func (c *Controller) getGroupMemberStatusesByPod(podInterface string) []*GroupMemberStatus {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

func (c *Controller) syncGroup(groupKey string) error { _ = "STUB: not implemented"; return nil }

// Send IGMP leave message to other Nodes to notify the current Node leaves the given multicast group.

// TODO: add check on the stale multicast group that is joined by the Pods on a different Node.
// remoteMembers is always empty with noEncap mode.

// Remove the multicast OpenFlow flow and group entries if none Pod member on local or remote Node is in the group.

// Remove the multicast flow entry if no local Pod is in the group.

// Install multicast flows and routing entries for the multicast group that local Pods join.

// Reinstall OpenFlow group because either the remote node receivers or local Pod receivers have changed.

// Install OpenFlow group for a new multicast group which has local Pod receivers joined.

// Install OpenFlow flow to forward packets to local Pod receivers which are included in the group.

func (c *Controller) groupHasInstalled(groupKey string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) addInstalledGroup(groupKey string) { _ = "STUB: not implemented"; return }

func (c *Controller) delInstalledGroup(groupKey string) { _ = "STUB: not implemented"; return }

func (c *Controller) localGroupHasInstalled(groupKey string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) addInstalledLocalGroup(groupKey string) { _ = "STUB: not implemented"; return }

func (c *Controller) delInstalledLocalGroup(groupKey string) { _ = "STUB: not implemented"; return }

func (c *Controller) addOrUpdateGroupEvent(e *mcastGroupEvent) { _ = "STUB: not implemented"; return }

func (c *Controller) memberChanged(e interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) initQueryGroup() error { _ = "STUB: not implemented"; return nil }

// updateQueryGroup gets all containers' interfaces, and add all ofports into IGMP query group.
func (c *Controller) updateQueryGroup() error { _ = "STUB: not implemented"; return nil }

// Install OpenFlow group for a new multicast group which has local Pod receivers joined.

// syncLocalGroupsToOtherNodes sends IGMP join message to other Nodes in the same cluster to notify what multicast groups
// are joined by this Node. This function is used only with encap mode.
func (c *Controller) syncLocalGroupsToOtherNodes() { _ = "STUB: not implemented"; return }

func (c *Controller) syncNodes() error { _ = "STUB: not implemented"; return nil }

// Notify local installed multicast groups to other Nodes in the cluster.

func podInterfaceIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getGroupEventKey(obj interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *Controller) CollectIGMPReportNPStats() (igmpANNPStats, igmpACNPStats map[apitypes.UID]map[string]*types.RuleMetric) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Controller) GetGroupPods() map[string][]v1beta2.PodReference {
	_ = "STUB: not implemented"
	return nil
}

// PodTrafficStats encodes the inbound and outbound multicast statistics of each Pod.
type PodTrafficStats struct {
	Inbound, Outbound uint64
}

func (c *Controller) GetPodStats(podName string, podNamespace string) *PodTrafficStats {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) GetAllPodsStats() map[*interfacestore.InterfaceConfig]*PodTrafficStats {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) checkNodeUpdate(old interface{}, cur interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) nodeWorker() { _ = "STUB: not implemented"; return }

func (c *Controller) processNextNodeItem() bool { _ = "STUB: not implemented"; return false }

// We call Done here so the workqueue knows we have finished processing this item. We also
// must remember to call Forget if we do not want this work item being re-queued. For
// example, we do not call Forget if a transient error occurs, instead the item is put back
// on the workqueue and attempted again after a back-off period.

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

func memberExists(status *GroupMemberStatus, e *mcastGroupEvent) bool {
	_ = "STUB: not implemented"
	return false
}

func addGroupMember(status *GroupMemberStatus, e *mcastGroupEvent) *GroupMemberStatus {
	_ = "STUB: not implemented"
	return nil
}

func deleteGroupMember(status *GroupMemberStatus, e *mcastGroupEvent) *GroupMemberStatus {
	_ = "STUB: not implemented"
	return nil
}
