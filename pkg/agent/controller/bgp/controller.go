// Copyright 2024 Antrea Authors
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

package bgp

import (
	"context"
	"errors"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	coreinformers "k8s.io/client-go/informers/core/v1"
	discoveryinformers "k8s.io/client-go/informers/discovery/v1"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	discoverylisters "k8s.io/client-go/listers/discovery/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/bgp"
	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	crdinformersv1a1 "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha1"
	crdinformersv1b1 "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
	crdlistersv1a1 "antrea.io/antrea/v2/pkg/client/listers/crd/v1alpha1"
	crdlistersv1b1 "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
)

const (
	controllerName = "BGPPolicyController"
	// How long to wait before retrying the processing of a BGPPolicy change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Disable resyncing.
	resyncPeriod time.Duration = 0
)

const (
	ipv4Suffix = "/32"
	ipv6Suffix = "/128"
)

const dummyKey = "dummyKey"

var (
	ErrBGPPolicyNotFound = errors.New("BGPPolicy not found")
)

type AdvertisedRouteType string

const (
	EgressIP              AdvertisedRouteType = "EgressIP"
	ServiceLoadBalancerIP AdvertisedRouteType = "ServiceLoadBalancerIP"
	ServiceExternalIP     AdvertisedRouteType = "ServiceExternalIP"
	ServiceClusterIP      AdvertisedRouteType = "ServiceClusterIP"
	NodeIPAMPodCIDR       AdvertisedRouteType = "NodeIPAMPodCIDR"
)

type RouteMetadata struct {
	Type      AdvertisedRouteType
	K8sObjRef string
}

type confederationConfig struct {
	identifier int32
	memberASNs sets.Set[uint32]
}

type bgpPolicyState struct {
	// The local BGP server.
	bgpServer bgp.Interface
	// name of the BGP policy.
	bgpPolicyName string
	// The port on which the local BGP server listens.
	listenPort int32
	// The AS number used by the local BGP server.
	localASN int32
	// The router ID used by the local BGP server.
	routerID string
	// The confederation config used by the local BGP server.
	confederationConfig *confederationConfig
	// routes stores all BGP routes advertised to BGP peers.
	routes map[bgp.Route]RouteMetadata
	// peerConfigs is a map that stores configurations of BGP peers. The map keys are the concatenated strings of BGP
	// peer IP address and ASN (e.g., "192.168.77.100-65000", "2001::1-65000").
	peerConfigs map[string]bgp.PeerConfig
}

type BGPPolicyInfo struct {
	BGPPolicyName           string
	RouterID                string
	LocalASN                int32
	ListenPort              int32
	ConfederationIdentifier int32
	MemberASNs              []uint32
}

type Controller struct {
	nodeInformer     cache.SharedIndexInformer
	nodeLister       corelisters.NodeLister
	nodeListerSynced cache.InformerSynced

	serviceInformer     cache.SharedIndexInformer
	serviceLister       corelisters.ServiceLister
	serviceListerSynced cache.InformerSynced

	egressInformer     cache.SharedIndexInformer
	egressLister       crdlistersv1b1.EgressLister
	egressListerSynced cache.InformerSynced

	bgpPolicyInformer     cache.SharedIndexInformer
	bgpPolicyLister       crdlistersv1a1.BGPPolicyLister
	bgpPolicyListerSynced cache.InformerSynced

	endpointSliceInformer     cache.SharedIndexInformer
	endpointSliceLister       discoverylisters.EndpointSliceLister
	endpointSliceListerSynced cache.InformerSynced

	secretInformer cache.SharedIndexInformer

	bgpPolicyState      *bgpPolicyState
	bgpPolicyStateMutex sync.RWMutex

	k8sClient             kubernetes.Interface
	bgpPeerPasswords      map[string]string
	bgpPeerPasswordsMutex sync.RWMutex

	nodeName     string
	enabledIPv4  bool
	enabledIPv6  bool
	podIPv4CIDR  string
	podIPv6CIDR  string
	nodeIPv4Addr string

	egressEnabled bool

	newBGPServerFn func(globalConfig *bgp.GlobalConfig) bgp.Interface

	queue workqueue.TypedRateLimitingInterface[string]
}

func NewBGPPolicyController(nodeInformer coreinformers.NodeInformer,
	serviceInformer coreinformers.ServiceInformer,
	egressInformer crdinformersv1b1.EgressInformer,
	bgpPolicyInformer crdinformersv1a1.BGPPolicyInformer,
	endpointSliceInformer discoveryinformers.EndpointSliceInformer,
	egressEnabled bool,
	k8sClient kubernetes.Interface,
	nodeConfig *config.NodeConfig,
	networkConfig *config.NetworkConfig) (*Controller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Controller) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *Controller) worker() { _ = "STUB: not implemented"; return }

func (c *Controller) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// If no error occurs we Forget this item, so it does not get queued again until another change happens.

// Put the item back on the work queue to handle any transient errors.

func (c *Controller) getEffectiveBGPPolicy() *v1alpha1.BGPPolicy {
	_ = "STUB: not implemented"
	return nil
}

func getConfederationConfig(conf *v1alpha1.Confederation) *confederationConfig {
	_ = "STUB: not implemented"
	return nil
}

func confederationConfigEqual(a, b *confederationConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) syncBGPPolicy(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the oldest BGPPolicy applied to the current Node as the effective BGPPolicy.

// When the effective BGPPolicy is nil, it means that there is no available BGPPolicy.

// If the BGPPolicy state is nil, just return.

// If the BGPPolicy state is not nil, stop the BGP server and reset the state to nil, then return.

// Retrieve the BGP policy name, listen port, local AS number and router ID from the effective BGPPolicy, and update them to the
// current state.

// If the BGPPolicy state is nil, a new BGP server should be started, initialize the BGPPolicy state to store the
// new BGP server, BGP policy name, listen port, local ASN, and router ID.
// If the BGPPolicy is not nil, any of the listen port, local AS number, router ID or confederation configuration
// has changed, stop the current BGP server first and reset the BGPPolicy state to nil; then start a new BGP server
// and initialize the BGPPolicy state to store the new BGP server, listen port, local ASN, and router ID.

// Stop the current BGP server.

// Reset the BGPPolicy state to nil.

// Create a new BGP server.

// Start the new BGP server.

// Initialize the BGPPolicy state to store the new BGP server, BGP policy name, listen port, local ASN, and router ID.

// It may happen that only BGP policy name has changed in effective BGP policy.

// Reconcile BGP peers.

// Reconcile BGP advertisements.

func (c *Controller) reconcileBGPPeers(ctx context.Context, bgpPeers []v1alpha1.BGPPeer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) reconcileBGPAdvertisements(ctx context.Context, bgpAdvertisements v1alpha1.Advertisements) error {
	_ = "STUB: not implemented"
	return nil
}

func hashNodeNameToIP(s string) string {
	_ = "STUB: not implemented"
	// Create a new FNV hash
	return ""
}

// Get the 32-bit hash

// Convert the hash to a 4-byte slice

func (c *Controller) getRouterID() (string, error) {
	_ = "STUB: not implemented"
	// According to RFC 4271:
	// BGP Identifier:
	//
	//	This 4-octet unsigned integer indicates the BGP Identifier of
	//	the sender.  A given BGP speaker sets the value of its BGP
	//	Identifier to an IP address that is assigned to that BGP
	//	speaker.  The value of the BGP Identifier is determined upon
	//	startup and is the same for every local interface and BGP peer.
	//
	// In goBGP, only an IPv4 address can be used as the BGP Identifier (BGP router ID).
	// The router ID could be specified in the Node annotation `node.antrea.io/bgp-router-id`.
	// For IPv4-only or dual-stack Kubernetes clusters, if the annotation is not present,
	// the Node's IPv4 address is used as the BGP router ID, ensuring uniqueness, and updated
	// to the Node annotation `node.antrea.io/bgp-router-id`.
	// For IPv6-only Kubernetes clusters without a Node IPv4 address, if the annotation is
	// not present, an IPv4 address will be generated by hashing the Node name and updated
	// to the Node annotation `node.antrea.io/bgp-router-id`.
	return "", nil
}

func (c *Controller) getRoutes(advertisements v1alpha1.Advertisements) map[bgp.Route]RouteMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) addServiceRoutes(advertisement *v1alpha1.ServiceAdvertisement, allRoutes map[bgp.Route]RouteMetadata) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) addEgressRoutes(allRoutes map[bgp.Route]RouteMetadata) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) addPodRoutes(allRoutes map[bgp.Route]RouteMetadata) {
	_ = "STUB: not implemented"
	return
}

func addRoutes(allRoutes map[bgp.Route]RouteMetadata, prefix, k8sObjRef string, routeType AdvertisedRouteType) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) hasLocalEndpoints(svc *corev1.Service) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) getPeerConfigs(peers []v1alpha1.BGPPeer) map[string]bgp.PeerConfig {
	_ = "STUB: not implemented"
	return nil
}

func generateBGPPeerKey(address string, asn int32) string { _ = "STUB: not implemented"; return "" }

func (c *Controller) addBGPPolicy(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updateBGPPolicy(oldObj, obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) deleteBGPPolicy(obj interface{}) { _ = "STUB: not implemented"; return }

func getIngressIPs(svc *corev1.Service) []string { _ = "STUB: not implemented"; return nil }

func (c *Controller) matchesCurrentNode(bgpPolicy *v1alpha1.BGPPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func matchesNode(node *corev1.Node, bgpPolicy *v1alpha1.BGPPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func matchesService(svc *corev1.Service, bgpPolicy *v1alpha1.BGPPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) hasAffectedPolicyByService(svc *corev1.Service) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) addService(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updateService(oldObj, obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) deleteService(obj interface{}) { _ = "STUB: not implemented"; return }

func noLocalTrafficPolicy(svc *corev1.Service) bool { _ = "STUB: not implemented"; return false }

func (c *Controller) addEndpointSlice(obj interface{}) { _ = "STUB: not implemented"; return }

// Events of EndpointSlices for Services without a `Local` traffic policy are ignored, as the Service IPs will
// always be advertised.

func (c *Controller) updateEndpointSlice(_, obj interface{}) { _ = "STUB: not implemented"; return }

// Events of EndpointSlices for Services without a `Local` traffic policy are ignored, as the Service IPs will
// always be advertised.

func (c *Controller) deleteEndpointSlice(obj interface{}) { _ = "STUB: not implemented"; return }

// Events of EndpointSlices for Services without a `Local` traffic policy are ignored, as the Service IPs will
// always be advertised.

func (c *Controller) hasAffectedPolicyByEgress() bool { _ = "STUB: not implemented"; return false }

func (c *Controller) addEgress(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updateEgress(oldObj, obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) deleteEgress(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) hasAffectedPolicyByNode(node *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) addNode(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updateNode(oldObj, obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) addSecret(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updateSecret(_, obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) deleteSecret(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updateBGPPeerPasswords(secret *corev1.Secret) {
	_ = "STUB: not implemented"
	return
}

// GetBGPPolicyInfo returns BGPPolicyInfo which includes
// BGPPolicyName, RouterID, LocalASN, ListenPort, ConfederationIdentifier
// and MemberASNs of effective BGP Policy applied on the Node.
func (c *Controller) GetBGPPolicyInfo() *BGPPolicyInfo { _ = "STUB: not implemented"; return nil }

// GetBGPPeerStatus returns current status of BGP Peers of effective BGP Policy applied on the Node.
func (c *Controller) GetBGPPeerStatus(ctx context.Context) ([]bgp.PeerStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBGPRoutes returns the advertised BGP routes.
func (c *Controller) GetBGPRoutes(ctx context.Context) (map[bgp.Route]RouteMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
