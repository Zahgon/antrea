// Copyright 2022 Antrea Authors
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

package multicluster

import (
	"net"
	"time"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	mcv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
	mcclientset "antrea.io/antrea/v2/multicluster/pkg/client/clientset/versioned"
	mcinformersv1alpha1 "antrea.io/antrea/v2/multicluster/pkg/client/informers/externalversions/multicluster/v1alpha1"
	mclisters "antrea.io/antrea/v2/multicluster/pkg/client/listers/multicluster/v1alpha1"
	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	antrearoute "antrea.io/antrea/v2/pkg/agent/route"
	"antrea.io/antrea/v2/pkg/agent/wireguard"
	"antrea.io/antrea/v2/pkg/config/agent"
)

const (
	controllerName = "MCDefaultRouteController"

	// Set resyncPeriod to 0 to disable resyncing
	resyncPeriod = 0 * time.Second
	// How long to wait before retrying the processing of a resource change
	minRetryDelay = 2 * time.Second
	maxRetryDelay = 120 * time.Second

	workerItemKey = "key"

	multiclusterWireGuardInterface = "antrea-mc-wg0"
	multiclusterWireGuardPublicKey = "publicKey"
)

var (
	wireGuardNewFunc = wireguard.New
)

// MCDefaultRouteController watches Gateway and ClusterInfoImport events.
// It is responsible for setting up necessary Openflow entries for multi-cluster
// traffic on a Gateway or a regular Node.
type MCDefaultRouteController struct {
	mcClient             mcclientset.Interface
	ofClient             openflow.Client
	routeClient          antrearoute.Interface
	wireGuardClient      wireguard.Interface
	nodeConfig           *config.NodeConfig
	networkConfig        *config.NetworkConfig
	wireGuardConfig      *config.WireGuardConfig
	gwInformer           mcinformersv1alpha1.GatewayInformer
	gwLister             mclisters.GatewayLister
	gwListerSynced       cache.InformerSynced
	ciImportInformer     mcinformersv1alpha1.ClusterInfoImportInformer
	ciImportLister       mclisters.ClusterInfoImportLister
	ciImportListerSynced cache.InformerSynced
	queue                workqueue.TypedRateLimitingInterface[string]
	// installedCIImports is for saving ClusterInfos which have been processed
	// in MCDefaultRouteController. Need to use mutex to protect 'installedCIImports' if
	// we change the number of 'defaultWorkers'.
	installedCIImports      map[string]*mcv1alpha1.ClusterInfoImport
	installedWireGuardPeers map[string]*mcv1alpha1.ClusterInfoImport
	// Need to use mutex to protect 'installedActiveGW' if we change to
	// use multiple go routines to handle events
	installedActiveGW *mcv1alpha1.Gateway
	// The Namespace where Antrea Multi-cluster Controller is running.
	namespace                    string
	enableStretchedNetworkPolicy bool
	enablePodToPodConnectivity   bool
	wireGuardInitialized         bool
}

func NewMCDefaultRouteController(
	mcClient mcclientset.Interface,
	gwInformer mcinformersv1alpha1.GatewayInformer,
	ciImportInformer mcinformersv1alpha1.ClusterInfoImportInformer,
	client openflow.Client,
	nodeConfig *config.NodeConfig,
	networkConfig *config.NetworkConfig,
	routeClient antrearoute.Interface,
	multiclusterConfig agent.MulticlusterConfig,
) *MCDefaultRouteController {
	_ = "STUB: not implemented"
	return nil
}

// Regardless of the tunnel type, the WireGuard device must only reduce MTU for encryption because the
// packets it transmits have been encapsulated.

func (c *MCDefaultRouteController) enqueueGateway(obj interface{}, isDelete bool) {
	_ = "STUB: not implemented"
	return
}

func (c *MCDefaultRouteController) enqueueClusterInfoImport(obj interface{}, isDelete bool) {
	_ = "STUB: not implemented"
	return
}

// Run will create a worker (go routines) which will process
// the Gateway events from the workqueue.
func (c *MCDefaultRouteController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// worker is a long-running function that will continually call the processNextWorkItem
// function in order to read and process a message on the workqueue.
func (c *MCDefaultRouteController) worker() { _ = "STUB: not implemented"; return }

func (c *MCDefaultRouteController) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// Put the item back on the workqueue to handle any transient errors.

// syncWireGuard reconciles WireGuard configurations in following way:
//  1. If the current Node is the Multi-cluster Gateway Node, controller will try to initialize corresponding WireGuard
//     configuration and route on the host, then add all existing Gateway Nodes in other member clusters as WireGuard peers.
//  2. If the current Node is not Multi-cluster Gateway Node, controller will try to clean up WireGuard configurations.
//
// Note: MCDefaultRouteController runs only one worker to process Gateway and ClusterInfoImport. So we do not need
// any synchronization mechanism.
func (c *MCDefaultRouteController) syncWireGuard() error { _ = "STUB: not implemented"; return nil }

// Check cache and existing ClusterInfoImports, clean up routes and WireGuard peers of the
// removed ClusterInfoImports.

func (c *MCDefaultRouteController) removeWireGuardRouteAndPeer(ciImport *mcv1alpha1.ClusterInfoImport) error {
	_ = "STUB: not implemented"
	return nil
}

// addWireGuardRouteAndPeer tries to update a WireGuard peer with ClusterInfoImport. If updating successfully,
// it will also create host route to WireGuard peer.
func (c *MCDefaultRouteController) addWireGuardRouteAndPeer(ciImport *mcv1alpha1.ClusterInfoImport) error {
	_ = "STUB: not implemented"
	return nil
}

// The cross-cluster traffic will be both encapsulated and encrypted. To avoid routing loop, we use a tunnel endpoint
// IP different from the WireGuard endpoint IP. Since the ServiceCIDR is guaranteed to be unique across member clusters,
// we choose the ServiceCIDR's network address as the tunnel endpoint IP. For instance, if a cluster's ServiceCIDR is
// 10.96.0.0/16, 10.96.0.0 will be used as the tunnel endpoint IP of the cluster's Gateway Node.

// initializeWireGuard initializes the WireGuard interface and client.
// It will also update Gateway's WireGuard field.
func (c *MCDefaultRouteController) initializeWireGuard(gateway *mcv1alpha1.Gateway) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanUpWireGuard deletes the WireGuard interface on the host.
// The WireGuard route will also be deleted automatically when the interface is deleted.
func (c *MCDefaultRouteController) cleanUpWireGuard() error { _ = "STUB: not implemented"; return nil }

func (c *MCDefaultRouteController) syncMCFlows() error { _ = "STUB: not implemented"; return nil }

// Active Gateway name doesn't change but still do a full flow sync
// for any Gateway Spec or ClusterInfoImport changes.

func (c *MCDefaultRouteController) syncMCFlowsForAllCIImps(activeGW *mcv1alpha1.Gateway) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MCDefaultRouteController) checkGatewayIPChange(activeGW *mcv1alpha1.Gateway) bool {
	_ = "STUB: not implemented"
	return false
}

// On a Gateway Node, the GatewayIP of the active Gateway will impact the Openflow rules.

// On a regular Node, the InternalIP of the active Gateway will impact the Openflow rules.

func (c *MCDefaultRouteController) addMCFlowsForAllCIImps(activeGW *mcv1alpha1.Gateway) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MCDefaultRouteController) addMCFlowsForSingleCIImp(activeGW *mcv1alpha1.Gateway, ciImport *mcv1alpha1.ClusterInfoImport,
	installedCIImp *mcv1alpha1.ClusterInfoImport, activeGWChanged bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MCDefaultRouteController) deleteMCFlowsForSingleCIImp(ciImpName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MCDefaultRouteController) deleteMCFlowsForAllCIImps() error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MCDefaultRouteController) getActiveGateway() (*mcv1alpha1.Gateway, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getActiveGateway(gwLister mclisters.GatewayLister) (*mcv1alpha1.Gateway, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The Gateway webhook guarantees there will be at most one Gateway in a cluster.

func generatePeerConfigs(subnets []string, gatewayIP net.IP) (map[*net.IPNet]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If WireGuard is disabled, getPeerGatewayTunnelIP will return Gateway's GatewayIP.
// If WireGuard is enabled, the WireGuard interfaces use the first IP address of ServiceCIDR
// as its IP address. So getPeerGatewayTunnelIP will return the first IP of the ServiceCIDR
// as the remote Gateway tunnel IP.
func getPeerGatewayTunnelIP(spec mcv1alpha1.ClusterInfo, enableWireGuard bool) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

func getLocalGatewayIP(gateway *mcv1alpha1.Gateway, enableWireGuard bool) net.IP {
	_ = "STUB: not implemented"
	return *new(net.IP)
}

// isWireGuardInfoChanged checks the information in ClusterInfoImport needed by WireGuard change or not.
func isWireGuardInfoChanged(cache, cur *mcv1alpha1.ClusterInfoImport) bool {
	_ = "STUB: not implemented"
	return false
}
