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

package egress

import (
	"net"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/watch"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/events"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/client"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/ipassigner"
	"antrea.io/antrea/v2/pkg/agent/ipassigner/linkmonitor"
	"antrea.io/antrea/v2/pkg/agent/memberlist"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/agent/route"
	"antrea.io/antrea/v2/pkg/agent/servicecidr"
	"antrea.io/antrea/v2/pkg/agent/types"
	cpv1b2 "antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	crdv1b1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	clientsetversioned "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/util/channel"
)

const (
	controllerName = "AntreaAgentEgressController"
	// How long to wait before retrying the processing of an Egress change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing an Egress change.
	defaultWorkers = 4
	// Disable resyncing.
	resyncPeriod time.Duration = 0
	// minEgressMark is the minimum mark of Egress IPs can be configured on a Node.
	minEgressMark = 1
	// maxEgressMark is the maximum mark of Egress IPs can be configured on a Node.
	maxEgressMark = 255

	egressIPIndex       = "egressIP"
	externalIPPoolIndex = "externalIPPool"

	// egressDummyDevice is the dummy device that holds the Egress IPs configured to the system by antrea-agent.
	egressDummyDevice = "antrea-egress0"
)

var maxSubnetsPerNodes = types.MaxRequestEgressRouteTable - types.MinRequestEgressRouteTable + 1

var emptyWatch = watch.NewEmptyWatch()

var newIPAssigner = ipassigner.NewIPAssigner

// egressState keeps the actual state of an Egress that has been realized.
type egressState struct {
	// The actual egress IP of the Egress. If it's different from the desired IP, there is an update to EgressIP, and we
	// need to remove previously installed flows.
	egressIP string
	// The actual datapath mark of this Egress. Used to check if the mark changes since last process.
	mark uint32
	// The actual openflow ports for which we have installed SNAT rules. Used to identify stale openflow ports when
	// updating or deleting an Egress.
	ofPorts sets.Set[int32]
	// The actual Pods of the Egress. Used to identify stale Pods when updating or deleting an Egress.
	pods sets.Set[string]
	// Rate-limit of this Egress.
	rateLimitMeter *rateLimitMeter
}

type rateLimitMeter struct {
	MeterID uint32
	Rate    uint32
	Burst   uint32
}

func (r *rateLimitMeter) Equals(rateLimit *rateLimitMeter) bool {
	_ = "STUB: not implemented"
	return false
}

// egressIPState keeps the actual state of an Egress IP. It's maintained separately from egressState because
// multiple Egresses can share an EgressIP.
type egressIPState struct {
	egressIP net.IP
	// The names of the Egresses that are currently referring to it.
	egressNames sets.Set[string]
	// The datapath mark of this Egress IP. 0 if this is not a local IP.
	mark uint32
	// Whether its flows have been installed.
	flowsInstalled bool
	// Whether its iptables rule has been installed.
	ruleInstalled bool
	// The subnet the Egress IP is associated with.
	subnetInfo *crdv1b1.SubnetInfo
}

// egressRouteTable stores the route table ID created for a subnet and the marks that are referencing it.
type egressRouteTable struct {
	// The route table ID.
	tableID uint32
	// The marks referencing the table. Once it's empty, the route table should be deleted.
	marks sets.Set[uint32]
}

// egressBinding keeps the Egresses applying to a Pod.
// There is one effective Egress for a Pod at any given time.
type egressBinding struct {
	effectiveEgress     string
	alternativeEgresses sets.Set[string]
}

type EgressController struct {
	ofClient             openflow.Client
	routeClient          route.Interface
	k8sClient            kubernetes.Interface
	crdClient            clientsetversioned.Interface
	antreaClientProvider client.AntreaClientProvider

	egressInformer     cache.SharedIndexInformer
	egressLister       crdlisters.EgressLister
	egressListerSynced cache.InformerSynced
	queue              workqueue.TypedRateLimitingInterface[string]

	externalIPPoolLister       crdlisters.ExternalIPPoolLister
	externalIPPoolListerSynced cache.InformerSynced

	// Use an interface for IP detector to enable testing.
	localIPDetector ipassigner.LocalIPDetector
	ifaceStore      interfacestore.InterfaceStore
	nodeName        string
	markAllocator   *idAllocator

	egressGroups      map[string]sets.Set[string]
	egressGroupsMutex sync.RWMutex

	egressBindings      map[string]*egressBinding
	egressBindingsMutex sync.RWMutex

	egressStates map[string]*egressState
	// The mutex is to protect the map, not the egressState items. The workqueue guarantees an Egress will only be
	// processed by a single worker at any time. So the returned EgressState has no race condition.
	egressStatesMutex sync.RWMutex

	egressIPStates      map[string]*egressIPState
	egressIPStatesMutex sync.Mutex

	cluster    memberlist.Interface
	ipAssigner ipassigner.IPAssigner

	egressIPScheduler *egressIPScheduler

	serviceCIDRInterface servicecidr.Interface
	serviceCIDRUpdateCh  chan struct{}
	// Declared for testing.
	serviceCIDRUpdateRetryDelay time.Duration

	trafficShapingEnabled bool

	eventBroadcaster events.EventBroadcaster
	record           events.EventRecorder
	// Whether to support non-default subnets.
	supportSeparateSubnet bool
	// Used to allocate route table ID.
	tableAllocator *idAllocator
	// Each subnet has its own route table.
	egressRouteTables map[crdv1b1.SubnetInfo]*egressRouteTable

	linkMonitor linkmonitor.Interface
}

func NewEgressController(
	ofClient openflow.Client,
	k8sClient kubernetes.Interface,
	antreaClientGetter client.AntreaClientProvider,
	crdClient clientsetversioned.Interface,
	ifaceStore interfacestore.InterfaceStore,
	routeClient route.Interface,
	nodeName string,
	nodeTransportInterface string,
	cluster memberlist.Interface,
	egressInformer crdinformers.EgressInformer,
	externalIPPoolInformer crdinformers.ExternalIPPoolInformer,
	nodeInformers coreinformers.NodeInformer,
	podUpdateSubscriber channel.Subscriber,
	serviceCIDRInterface servicecidr.Interface,
	maxEgressIPsPerNode int,
	trafficShapingEnabled bool,
	supportSeparateSubnet bool,
	linkMonitor linkmonitor.Interface,
	uniqueMACForSubInterfaces bool,
) (*EgressController, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// One buffer is enough as we just use it to ensure the target handler is executed once.

// egressIPIndex will be used to get all Egresses sharing the same Egress IP.

// Subscribe Pod update events from CNIServer to enforce Egress earlier, instead of waiting for their IPs are
// reported to kube-apiserver and processed by antrea-controller.

// onEgressIPSchedule will be called when EgressIPScheduler reschedules an Egress's IP.
func (c *EgressController) onEgressIPSchedule(egress string) { _ = "STUB: not implemented"; return }

// onServiceCIDRUpdate will be called when ServiceCIDRs change.
// It ensures updateServiceCIDRs will be executed once after this call.
func (c *EgressController) onServiceCIDRUpdate(_ []*net.IPNet) { _ = "STUB: not implemented"; return }

// The previous event is not processed yet, discard the new event.

func (c *EgressController) updateServiceCIDRs(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// Consume the first tick.

// No need to retry in this case as the Service CIDRs won't be available until it receives a service CIDRs update.

// Schedule a retry as it should be transient error.

// processPodUpdate will be called when CNIServer publishes a Pod update event.
// It triggers reconciling the effective Egress of the Pod.
func (c *EgressController) processPodUpdate(e interface{}) { _ = "STUB: not implemented"; return }

// addEgress processes Egress ADD events.
func (c *EgressController) addEgress(obj interface{}) { _ = "STUB: not implemented"; return }

// updateEgress processes Egress UPDATE events.
func (c *EgressController) updateEgress(old, cur interface{}) { _ = "STUB: not implemented"; return }

// Ignore handling the Egress Status change if Egress IP already has been assigned on current node.

// deleteEgress processes Egress DELETE events.
func (c *EgressController) deleteEgress(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *EgressController) addExternalIPPool(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *EgressController) updateExternalIPPool(old, cur interface{}) {
	_ = "STUB: not implemented"
	return
}

// We only care about SubnetInfo here.

func (c *EgressController) onExternalIPPoolUpdated(pool string) { _ = "STUB: not implemented"; return }

func (c *EgressController) onLocalIPUpdate(ip string, added bool) {
	_ = "STUB: not implemented"
	return
}

// Run will create defaultWorkers workers (go routines) which will process the Egress events from the
// workqueue.
func (c *EgressController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// replaceEgressIPs unassigns stale Egress IPs that shouldn't be present on this Node and assigns the missing IPs
// on this node. The unassigned IPs are from Egresses that were either deleted from the Kubernetes API or migrated
// to other Nodes when the agent on this Node was not running.
func (c *EgressController) replaceEgressIPs() error { _ = "STUB: not implemented"; return nil }

// Ignore the Egress if the ExternalIPPool doesn't exist.

// Record the Egress's state as we assign their IPs to this Node in the following call. It makes sure these
// Egress IPs will be unassigned when the Egresses are deleted.

// worker is a long-running function that will continually call the processNextWorkItem function in
// order to read and process a message on the workqueue.
func (c *EgressController) worker() { _ = "STUB: not implemented"; return }

func (c *EgressController) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

// installPolicyRoute ensures Egress traffic with the given mark access external network via the subnet's gateway, and
// tagged with the subnet's VLAN ID if present.
func (c *EgressController) installPolicyRoute(ipState *egressIPState, subnetInfo *crdv1b1.SubnetInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// Deletes stale policy route first.

// If the subnetInfo is nil, policy routing is not needed. The Egress IP should just use the main route table.

// Get or create a route table for this subnet.

// Get the index of the network interface to which IPs in the subnet are assigned.
// The network interface will be used as the device via which the Egress traffic leaves.

// This should never happen.

// Add an IP rule to make the marked Egress traffic look up the table.

// Track the route table's usage.

// Track the current subnet of the Egress IP.

// uninstallPolicyRoute deletes the policy route of the Egress IP.
func (c *EgressController) uninstallPolicyRoute(ipState *egressIPState) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete the route table if it is not used by any Egress.

// realizeEgressIP realizes an Egress IP. Multiple Egresses can share the same Egress IP.
// If it's called the first time for a local Egress IP, it allocates a locally-unique mark for the IP and installs flows
// and iptables rule for this IP and the mark.
// If the Egress IP is changed from local to non local, it uninstalls flows and iptables rule and releases the mark.
// The method returns the mark on success. Non local Egresses use 0 as the mark.
func (c *EgressController) realizeEgressIP(egressName, egressIP string, subnetInfo *crdv1b1.SubnetInfo) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Create an egressIPState if this is the first Egress using the IP.

// Ensure the Egress IP has a mark allocated when it's a local IP.

// Ensure datapath is installed properly.

// Ensure datapath is uninstalled properly.

func bandwidthToRateLimitMeter(bandwidth *crdv1b1.Bandwidth, meterID uint32) *rateLimitMeter {
	_ = "STUB: not implemented"
	return nil
}

func (c *EgressController) realizeEgressQoS(egressName string, eState *egressState, mark uint32, bandwidth *crdv1b1.Bandwidth) error {
	_ = "STUB: not implemented"
	return nil
}

// QoS is desired only if the Egress is configured on this Node.

// Nothing changes.

// It's desired to have QoS on this Node, install/override it.

// It's undesired to have QoS on this Node, uninstall it.

// unrealizeEgressIP unrealizes an Egress IP, reverts what realizeEgressIP does.
// For a local Egress IP, only when the last Egress unrealizes the Egress IP, it will releases the IP's mark and
// uninstalls corresponding flows and iptables rule.
func (c *EgressController) unrealizeEgressIP(egressName, egressIP string) error {
	_ = "STUB: not implemented"
	return nil
}

// The Egress IP was not configured before, do nothing.

// Unlink the Egress from the EgressIP. If it's the last Egress referring to it, uninstall its datapath rules and
// release the mark if installed.

func (c *EgressController) getEgressState(egressName string) (*egressState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *EgressController) deleteEgressState(egressName string) { _ = "STUB: not implemented"; return }

func (c *EgressController) newEgressState(egressName string, egressIP string) *egressState {
	_ = "STUB: not implemented"
	return nil
}

// bindPodEgress binds the Pod with the Egress and returns whether this Egress is the effective one for the Pod.
func (c *EgressController) bindPodEgress(pod, egress string) bool {
	_ = "STUB: not implemented"
	return false
}

// Promote itself as the effective Egress if there was not one.

// unbindPodEgress unbinds the Pod with the Egress.
// If the unbound Egress was the effective one for the Pod and there are any alternative ones, it will return the new
// effective Egress and true. Otherwise it return empty string and false.
func (c *EgressController) unbindPodEgress(pod, egress string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// The binding must exist.

// Remove the Pod's binding if there is no alternative.

func (c *EgressController) updateEgressStatus(egress *crdv1b1.Egress, egressIP string, scheduleErr error) error {
	_ = "STUB: not implemented"
	return nil
}

// Select one Node to update false status among all Nodes.
// We don't care about the value of egress.Spec.EgressIP, just use it to reach a consensus among all agents
// about which one should do the update.

// Skip if the Node is not the selected one.

// If the error is nil, it means the Egress hasn't been processed yet.
// The scheduler will get a result for the Egress very soon regardless of success or failure and trigger the
// controller to process it another time, so we avoid generating a transient state here, which may lead to some
// back-off retries due to updating conflict.

// The Egress IP is assigned to a Node (egressIP != "") but it's not this Node (isLocal == false), do nothing.

// Must make a copy here as we will append more conditions. If it's appended to desiredStatus directly, there
// would be duplicate conditions when the function retries.

// Copy conditions other than crdv1b1.IPAssigned to statusToUpdate.

// Return the error from UPDATE.

func (c *EgressController) syncEgress(egressName string) error {
	_ = "STUB: not implemented"
	return nil
}

// The Egress has been removed, clean it up.

// The Egress hasn't been installed, do nothing.

// Only check whether the Egress IP should be assigned to this Node when the Egress is schedulable.
// Otherwise, users are responsible for assigning the Egress IP to Nodes.

// If the EgressIP changes, uninstalls this Egress first.

// Do not proceed if EgressIP is empty.

// Ensure the Egress IP is assigned to the system. Force advertising the IP if it was previously assigned to
// another Node in the Egress API. This could force refreshing other peers' neighbor cache when the Egress IP is
// obtained by this Node and another Node at the same time in some situations, e.g. split brain.

// Unassign the Egress IP from the local Node if it was assigned by the agent.

// Realize the latest EgressIP and get the desired mark.

// If the mark changes, uninstall all of the Egress's Pod flows first, then installs them with new mark.
// It could happen when the Egress IP is added to or removed from the Node.

// Uninstall all of its Pod flows.

// Copy the previous ofPorts and Pods. They will be used to identify stale ofPorts and Pods.

// Get a copy of the desired Pods.

// Install SNAT flows for desired Pods.

// If the Egress is not the effective one for the Pod, do nothing.

// Get the Pod's openflow port.

// Uninstall SNAT flows for stale Pods.

func (c *EgressController) uninstallEgress(egressName string, eState *egressState, egress *crdv1b1.Egress) error {
	_ = "STUB: not implemented"
	// Uninstall all of its Pod flows.
	return nil
}

// Release the EgressIP's mark if the Egress is the last one referring to it.

// Uninstall its meter.

// Unassign the Egress IP from the local Node if it was assigned by the agent.

// Remove the Egress's state.

func (c *EgressController) uninstallPodFlows(egressName string, egressState *egressState, ofPorts sets.Set[int32], pods sets.Set[string]) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove Pods from the Egress state after uninstalling Pod's flows to avoid overlapping. Otherwise another Egress
// may install new flows for the Pod before this Egress uninstalls its previous flows, causing conflicts.
// For each Pod, if the Egress was the Pod's effective Egress and there are other Egresses applying to it, it will
// pick one and trigger its resync.

// Trigger resyncing of the new effective Egresses of the removed Pods.

func (c *EgressController) watchEgressGroup() { _ = "STUB: not implemented"; return }

// Watch method doesn't return error but "emptyWatch" in case of some partial data errors,
// e.g. timeout error. Make sure that watcher is not empty and log error otherwise.

// First receive init events from the result channel and buffer them until
// a Bookmark event is received, indicating that all init events have been
// received.

func (c *EgressController) replaceEgressGroups(groups []*cpv1b2.EgressGroup) {
	_ = "STUB: not implemented"
	return
}

func (c *EgressController) addEgressGroup(group *cpv1b2.EgressGroup) {
	_ = "STUB: not implemented"
	return
}

func (c *EgressController) patchEgressGroup(patch *cpv1b2.EgressGroupPatch) {
	_ = "STUB: not implemented"
	return
}

func (c *EgressController) deleteEgressGroup(group *cpv1b2.EgressGroup) {
	_ = "STUB: not implemented"
	return
}

// GetEgressIPByMark returns the Egress IP associated with the snatMark.
func (c *EgressController) GetEgressIPByMark(mark uint32) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetEgress returns the Egress configuration applied to this Pod.
// If no Egress is applied to the Pod, an error will be returned.
func (c *EgressController) GetEgress(ns, podName string) (types.EgressConfig, error) {
	_ = "STUB: not implemented"
	return *new(types.EgressConfig), nil
}

// An Egress is schedulable if its Egress IP is allocated from ExternalIPPool.
func isEgressSchedulable(egress *crdv1b1.Egress) bool { _ = "STUB: not implemented"; return false }

// compareEgressStatus compares two Egress Statuses, ignoring LastTransitionTime and conditions other than IPAssigned, returns true if they are equal.
func compareEgressStatus(currentStatus, desiredStatus *crdv1b1.EgressStatus) bool {
	_ = "STUB: not implemented"
	return false
}
