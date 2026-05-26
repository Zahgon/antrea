// Copyright 2019 Antrea Authors
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

package noderoute

import (
	"net"
	"net/netip"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/cache/synctrack"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/controller/ipseccertificate"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/agent/route"
	"antrea.io/antrea/v2/pkg/agent/wireguard"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
	utilip "antrea.io/antrea/v2/pkg/util/ip"
	utilwait "antrea.io/antrea/v2/pkg/util/wait"
)

const (
	controllerName = "AntreaAgentNodeRouteController"
	// Interval of reprocessing every node.
	nodeResyncPeriod = 60 * time.Second
	// How long to wait before retrying the processing of a node change
	minRetryDelay = 2 * time.Second
	maxRetryDelay = 120 * time.Second
	// Default number of workers processing a node change
	defaultWorkers = 4

	ovsExternalIDNodeName = "node-name"

	nodeRouteInfoPodCIDRIndexName = "podCIDR"
)

// Controller is responsible for setting up necessary IP routes and Openflow entries for inter-node traffic.
type Controller struct {
	ovsBridgeClient  ovsconfig.OVSBridgeClient
	ofClient         openflow.Client
	ovsCtlClient     ovsctl.OVSCtlClient
	routeClient      route.Interface
	interfaceStore   interfacestore.InterfaceStore
	networkConfig    *config.NetworkConfig
	nodeConfig       *config.NodeConfig
	nodeInformer     coreinformers.NodeInformer
	nodeLister       corelisters.NodeLister
	nodeListerSynced cache.InformerSynced
	queue            workqueue.TypedRateLimitingInterface[string]
	// installedNodes records routes and flows installation states of Nodes.
	// The key is the host name of the Node, the value is the nodeRouteInfo of the Node.
	// A node will be in the map after its flows and routes are installed successfully.
	installedNodes cache.Indexer
	// podSubnetsMutex protects access to the podSubnets set.
	podSubnetsMutex sync.RWMutex
	// podSubnets is a set which stores all known PodCIDRs in the cluster as masked netip.Prefix objects.
	podSubnets      sets.Set[netip.Prefix]
	maskSizeV4      int
	maskSizeV6      int
	wireGuardClient wireguard.Interface
	// ipsecCertificateManager is useful for determining whether the ipsec certificate has been configured
	// or not when IPsec is enabled with "cert" mode. The NodeRouteController must wait for the certificate
	// to be configured before installing routes/flows to peer Nodes to prevent unencrypted traffic across Nodes.
	ipsecCertificateManager ipseccertificate.Manager
	// flowRestoreCompleteWait is to be decremented after installing flows for initial Nodes.
	flowRestoreCompleteWait *utilwait.Group
	// hasProcessedInitialList keeps track of whether the initial informer list has been
	// processed by workers.
	// See https://github.com/kubernetes/apiserver/blob/v0.36.1/pkg/admission/plugin/policy/internal/generic/controller.go
	hasProcessedInitialList *synctrack.AsyncTracker[string]
	// eventHandlerRegistration.HasSynced will be used to track whether even handlers have been
	// called for the initial list.
	eventHandlerRegistration cache.ResourceEventHandlerRegistration
}

// NewNodeRouteController instantiates a new Controller object which will process Node events
// and ensure connectivity between different Nodes.
func NewNodeRouteController(
	nodeInformer coreinformers.NodeInformer,
	client openflow.Client,
	ovsCtlClient ovsctl.OVSCtlClient,
	ovsBridgeClient ovsconfig.OVSBridgeClient,
	routeClient route.Interface,
	interfaceStore interfacestore.InterfaceStore,
	networkConfig *config.NetworkConfig,
	nodeConfig *config.NodeConfig,
	wireguardClient wireguard.Interface,
	ipsecCertificateManager ipseccertificate.Manager,
	flowRestoreCompleteWait *utilwait.Group,
) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func nodeRouteInfoKeyFunc(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func nodeRouteInfoPodCIDRIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nodeRouteInfo is the route related information extracted from corev1.Node.
type nodeRouteInfo struct {
	nodeName           string
	podCIDRs           []*net.IPNet
	nodeIPs            *utilip.DualStackIPs
	gatewayIPs         *utilip.DualStackIPs
	nodeMAC            net.HardwareAddr
	wireGuardPublicKey string
}

// enqueueNode adds an object to the controller work queue
// obj could be a *corev1.Node, or a DeletionFinalStateUnknown item.
func (c *Controller) enqueueNode(obj interface{}, isInInitialList bool) {
	_ = "STUB: not implemented"
	return
}

// Ignore notifications for this Node, no need to establish connectivity to itself.

// removeStaleGatewayRoutes removes all the gateway routes which no longer correspond to a Node in
// the cluster. If the antrea agent restarts and Nodes have left the cluster, this function will
// take care of removing routes which are no longer valid.
func (c *Controller) removeStaleGatewayRoutes() error { _ = "STUB: not implemented"; return nil }

// We iterate over all current Nodes, including the Node on which this agent is
// running, so the route to local Pods will be desired as well.

// routeClient will remove orphaned routes whose destinations are not in desiredPodCIDRs.

// removeStaleTunnelPorts removes all the tunnel ports which no longer correspond to a Node in the
// cluster. If the antrea agent restarts and Nodes have left the cluster, this function will take
// care of removing tunnel ports which are no longer valid. If the tunnel port configuration has
// changed, the tunnel port will also be deleted (the controller loop will later take care of
// re-creating the port with the correct configuration).
func (c *Controller) removeStaleTunnelPorts() error { _ = "STUB: not implemented"; return nil }

// desiredInterfaces is the set of interfaces we wish to have, based on the current list of
// Nodes. If a tunnel port corresponds to a valid Node but its configuration is wrong, we
// will not include it in the set.

// knownInterfaces is the list of interfaces currently in the local cache.

// Tunnel port not created for this Node, nothing to do.

// remote_name and psk are mutually exclusive.

// remove all ports which are no longer needed or for which the configuration is no longer
// valid.

// this interface matches an existing Node, nothing to do.

// should not happen, nothing should have concurrent access to the interface
// store for tunnel interfaces.

func (c *Controller) compareInterfaceConfig(interfaceConfig *interfacestore.InterfaceConfig,
	peerNodeIP net.IP, psk, remoteName, interfaceName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) reconcile() error { _ = "STUB: not implemented"; return nil }

// reconciliation consists of removing stale routes and stale / invalid tunnel ports:
// missing routes and tunnel ports will be added normally by processNextWorkItem, which will
// also take care of updating incorrect routes.

// removeStaleWireGuardPeers deletes stale WireGuard peers if necessary.
func (c *Controller) removeStaleWireGuardPeers() error { _ = "STUB: not implemented"; return nil }

// Run will create defaultWorkers workers (go routines) which will process the Node events from the
// workqueue.
func (c *Controller) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// If agent is running policy-only mode, it delegates routing to
// underlying network. Therefore it needs not know the routes to
// peer Pod CIDRs.

// After eventHandlerRegistration.HasSynced is true, we need to let hasProcessedInitialList
// known that the source (upstream) is synced.

// When the initial list of Nodes has been processed, we decrement flowRestoreCompleteWait.

// An error here means the context has been cancelled, which means that the stopCh
// has been closed. While it is still possible for c.hasProcessedInitialList.HasSynced
// to become true, as workers may not have returned yet, we should not decrement
// flowRestoreCompleteWait or log the message below.

// HasSynced returns true when the initial list of Nodes has been processed by the controller.
func (c *Controller) HasSynced() bool { _ = "STUB: not implemented"; return false }

// worker is a long-running function that will continually call the processNextWorkItem function in
// order to read and process a message on the workqueue.
func (c *Controller) worker() { _ = "STUB: not implemented"; return }

// processNextWorkItem processes an item in the "node" work queue, by calling syncNodeRoute after
// casting the item to a string (Node name). If syncNodeRoute returns an error, this function
// handles it by requeuing the item so that it can be processed again later. If syncNodeRoute is
// successful, the Node is removed from the queue until we get notified of a new change. This
// function returns false if and only if the work queue was shutdown (no more items will be
// processed).
func (c *Controller) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// We call Done here so the workqueue knows we have finished processing this item. We also
// must remember to call Forget if we do not want this work item being re-queued. For
// example, we do not call Forget if a transient error occurs, instead the item is put back
// on the workqueue and attempted again after a back-off period.

// We call Finished unconditionally even if this only matters for the initial list of
// Nodes. There is no harm in calling Finished without a corresponding call to Start.

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

// syncNode manages connectivity to "peer" Node with name nodeName
// If we have not established connectivity to the Node yet:
//   - we install the appropriate Linux route:
//
// Destination     Gateway         Use Iface
// peerPodCIDR     peerGatewayIP   localGatewayIface (e.g antrea-gw0)
//   - we install the appropriate OpenFlow flows to ensure that all the traffic destined to
//     peerPodCIDR goes through the correct L3 tunnel.
//
// If the Node no longer exists (cannot be retrieved by name from nodeLister) we delete the route
// and OpenFlow flows associated with it.
func (c *Controller) syncNodeRoute(nodeName string) error { _ = "STUB: not implemented"; return nil }

// The work queue guarantees that concurrent goroutines cannot call syncNodeRoute on the
// same Node, which is required by the InstallNodeFlows / UninstallNodeFlows OF Client
// methods.

func (c *Controller) deleteNodeRoute(nodeName string) error { _ = "STUB: not implemented"; return nil }

// Route is not added for this Node.

// Tunnel port not created for this Node.

func (c *Controller) addNodeRoute(nodeName string, node *corev1.Node) error {
	_ = "STUB: not implemented"
	// It is only for Windows Noencap mode to get Node MAC.
	return nil
}

// Route is already added for this Node and Node MAC, transport IP
// and WireGuard public key are not changed.

// If no valid PodCIDR is configured in Node.Spec, return immediately.

// Does not help to return an error and trigger controller retries.

// PodCIDRs can be released from deleted Nodes and allocated to new Nodes. For server side, it won't happen that a
// PodCIDR is allocated to more than one Node at any point. However, for client side, if a resync happens to occur
// when there are Node creation and deletion events, the informer will generate the events in a way that all
// creation events come before deletion ones even they actually happen in the opposite order on the server side.
// See https://github.com/kubernetes/kubernetes/blob/v1.18.2/staging/src/k8s.io/client-go/tools/cache/delta_fifo.go#L503-L512
// Therefore, a PodCIDR may appear in a new Node before the Node that previously owns it is removed. To ensure the
// stale routes, flows, and relevant cache of this podCIDR are removed appropriately, we wait for the Node deletion
// event to be processed before proceeding, or the route installation and uninstallation operations may override or
// conflict with each other.
// For Windows Noencap case, it is possible that nodesHaveSamePodCIDR is the Node itself because the Node
// MAC annotation was not set yet when the Node was initially installed. Then it is processed for the second
// time when its MAC annotation is updated.

// Return an error so that the Node will be put back to the workqueue and will be retried later.

// Create a separate tunnel port for the Node, as OVS IPsec monitor needs to
// read PSK and remote IP from the Node's tunnel interface to create IPsec
// security policies. We use the Node's IPv4 address when present, and the
// Node's IPv6 address otherwise.

func getPodCIDRsOnNode(node *corev1.Node) []string { _ = "STUB: not implemented"; return nil }

// Does not help to return an error and trigger controller retries.

// createIPSecTunnelPort creates an IPsec tunnel port for the remote Node if the
// tunnel does not exist, and returns the ofport number.
func (c *Controller) createIPSecTunnelPort(nodeName string, nodeIP net.IP) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// remote_name and psk are mutually exclusive.

// check if Node IP, PSK, remote name, or tunnel type changes. This can
// happen if removeStaleTunnelPorts fails to remove a "stale"
// tunnel port for which the configuration has changed, return error to requeue the Node.

// ofPortRequest - let OVS allocate OFPort number.

// GetOFPort will wait for up to 1 second for OVSDB to report the OFPort number.

// Could be a temporary OVSDB connection failure or timeout.
// Let NodeRouteController retry at errors.

// Set the port with no-flood to reject ARP flood packets.

// ParseTunnelInterfaceConfig initializes and returns an InterfaceConfig struct
// for a tunnel interface. It reads tunnel type, remote IP, IPsec PSK from the
// OVS interface options, and NodeName from the OVS port external_ids.
// nil is returned, if the OVS port and interface configurations are not valid
// for a tunnel interface.
func ParseTunnelInterfaceConfig(
	portData *ovsconfig.OVSPortData,
	portConfig *interfacestore.OVSPortConfig) *interfacestore.InterfaceConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) findPodSubnetForIP(ip netip.Addr) (netip.Prefix, bool) {
	_ = "STUB: not implemented"
	return *new(netip.Prefix), false
}

// LookupIPInPodSubnets returns two boolean values. The first one indicates whether the IP can be
// found in a PodCIDR for one of the cluster Nodes. The second one indicates whether the IP is used
// as a gateway IP. The second boolean value can only be true if the first one is true.
func (c *Controller) LookupIPInPodSubnets(ip netip.Addr) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// getNodeMAC gets Node's br-int MAC from its annotation. It is only for Windows Noencap mode.
func getNodeMAC(node *corev1.Node) (net.HardwareAddr, error) {
	_ = "STUB: not implemented"
	return *new(net.HardwareAddr), nil
}

func cidrToPrefix(cidr *net.IPNet) (netip.Prefix, error) {
	_ = "STUB: not implemented"
	return *new(netip.Prefix), nil
}

func cidrsToPrefixes(cidrs []*net.IPNet) ([]netip.Prefix, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
