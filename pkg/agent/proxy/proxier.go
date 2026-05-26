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

package proxy

import (
	"context"
	"math"
	"net"
	"sync"
	"time"

	"antrea.io/ofnet/ofctrl"
	corev1 "k8s.io/api/core/v1"
	discovery "k8s.io/api/discovery/v1"
	"k8s.io/apimachinery/pkg/labels"
	coreinformers "k8s.io/client-go/informers/core/v1"
	discoveryinformers "k8s.io/client-go/informers/discovery/v1"

	agentconfig "antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/nodeip"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/agent/proxy/types"
	"antrea.io/antrea/v2/pkg/agent/route"
	antreaconfig "antrea.io/antrea/v2/pkg/config/agent"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
	k8sproxy "antrea.io/antrea/v2/third_party/proxy"
	"antrea.io/antrea/v2/third_party/proxy/config"
	"antrea.io/antrea/v2/third_party/proxy/healthcheck"
	"antrea.io/antrea/v2/third_party/proxy/runner"
)

const (
	resyncPeriod  = time.Minute
	componentName = "antrea-agent-proxy"
	// SessionAffinity timeout is implemented using a hard_timeout in OVS. hard_timeout is
	// represented by a uint16 in the OpenFlow protocol.
	maxSupportedAffinityTimeout = math.MaxUint16
	// labelServiceProxyName is the well-known label for service proxy name defined in
	// https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/2447-Make-kube-proxy-service-abstraction-optional
	labelServiceProxyName = "service.kubernetes.io/service-proxy-name"
)

// Proxier extends the standard k8sproxy.Provider interface with additional query capabilities defined in interface
// ProxyQuerier. It serves as an enhanced proxy provider implementation without modifying the original k8sproxy.Provider
// interface.
type Proxier interface {
	k8sproxy.Provider
	ProxyQuerier
}

// ProxyQuerier is the query interface for retrieving information from the Proxy.
type ProxyQuerier interface {
	// GetServiceFlowKeys returns the keys (match strings) of the cached OVS
	// flows and the OVS group IDs for a Service. False is returned if the
	// Service is not found.
	GetServiceFlowKeys(serviceName, namespace string) ([]string, []binding.GroupIDType, bool)
	// GetServiceByIP returns the ServicePortName struct for the given serviceString(ClusterIP:Port/Proto).
	// False is returned if the serviceString is not found in serviceStringMap.
	GetServiceByIP(serviceStr string) (k8sproxy.ServicePortName, bool)
}

type ProxyServer struct {
	endpointSliceConfig *config.EndpointSliceConfig
	serviceConfig       *config.ServiceConfig
	nodeConfig          *config.NodeConfig

	nodeManager   *k8sproxy.NodeManager
	healthzServer *healthcheck.ProxyHealthServer
	ofClient      openflow.Client
	proxier       Proxier
}

type proxier struct {
	// mu protects the fields below, which can be read by GetServiceFlowKeys() called by the "/ovsflows" API handler.
	mu sync.Mutex
	// endpointsChanges and serviceChanges contains all changes to endpoints and
	// services that happened since last syncProxyRules call. For a single object,
	// changes are accumulated. Once both endpointsChanges and serviceChanges
	// have been synced, syncProxyRules will start syncing rules to OVS.
	endpointsChanges *endpointsChangesTracker
	serviceChanges   *serviceChangesTracker
	topologyLabels   map[string]string
	nodeIPChecker    nodeip.Checker
	// serviceMap stores services we expect to be installed.
	serviceMap k8sproxy.ServicePortMap
	// serviceInstalledMap stores services we actually installed.
	serviceInstalledMap k8sproxy.ServicePortMap
	// endpointsMap stores endpoints we expect to be installed.
	endpointsMap k8sproxy.EndpointsMap
	// endpointsInstalledMap stores endpoints we actually installed.
	endpointsInstalledMap k8sproxy.EndpointsMap
	// endpointReferenceCounter stores the number of times an Endpoint is referenced by Services.
	endpointReferenceCounter map[string]int
	// groupCounter is used to allocate groupID.
	groupCounter types.GroupCounter

	ipToServiceMap      *ipToServiceMap
	serviceHealthServer healthcheck.ServiceHealthServer
	healthzServer       *healthcheck.ProxyHealthServer

	// syncedOnce returns true if the proxier has synced rules at least once.
	syncedOnce      bool
	syncedOnceMutex sync.RWMutex

	runner                               *runner.BoundedFrequencyRunner
	stopChan                             <-chan struct{}
	ofClient                             openflow.Client
	routeClient                          route.Interface
	nodePortAddresses                    []net.IP
	hostname                             string
	ipFamily                             corev1.IPFamily
	proxyAll                             bool
	proxyLoadBalancerIPs                 bool
	preferSameTrafficDistributionEnabled bool
	supportNestedService                 bool
	cleanupStaleUDPSvcConntrack          bool

	// When a Service's LoadBalancerMode is DSR, the following changes will be applied to the OpenFlow flows and groups:
	// 1. ClusterGroup will be used by traffic working in DSR mode on ingress Node.
	//   * If a local Endpoint is selected, it will just be handled normally as DSR is not applicable in this case.
	//   * If a remote Endpoint is selected, it will be sent to the backend Node that hosts the Endpoint without being
	//     NAT'd, the eventual Endpoint will be determined on the backend Node and may be different from the one
	//     selected here.
	// 2. LocalGroup will be used by traffic working in DSR mode on backend Node. In this way, each Endpoint has the
	//    same chance to be selected eventually.
	// 3. Traffic working in DSR mode on ingress Node will be marked and treated specially, e.g. bypassing SNAT.
	// 4. Learned flow will be created for each connection to ensure consistent load balance decision for a connection of DSR mode.
	//
	// Learned flow is necessary because connections of DSR mode will remain invalid on ingress Node as it can only see
	// requests and not responses. And OVS doesn't provide ct_state and ct_label for invalid connections. Thus, we can't
	// store the load balance decision of the connection to ct_state or ct_label. To ensure consistent load balancing
	// decision for packets of a connection, we use "learn" action to generate a learned flow when processing the first
	// packet of a connection, and rely on the learned flow to process subsequent packets of the same connection.
	defaultLoadBalancerMode agentconfig.LoadBalancerMode
}

func (p *proxier) SyncedOnce() bool { _ = "STUB: not implemented"; return false }

func endpointKey(endpoint k8sproxy.Endpoint, protocol binding.Protocol) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *proxier) isInitialized() bool {
	_ = "STUB: not implemented"
	// When LoadBalancerModeDSR is enabled, we wait for NodeIPChecker to be initialized before processing any Service, to
	// ensure we get correct result when checking if an Endpoint is running in host network.
	return false
}

// removeStaleServices removes all the configurations of expired Services and their associated Endpoints.
func (p *proxier) removeStaleServices() { _ = "STUB: not implemented"; return }

// Remove Service group which has only local Endpoints.

// Remove Service group which has all Endpoints.

// Remove associated Endpoints flows.

func (p *proxier) removeServiceFlows(svcInfo *types.ServiceInfo) bool {
	_ = "STUB: not implemented"
	return false
}

// Remove ClusterIP flows.

// Remove NodePort flows and configurations.

// Remove ExternalIP flows and configurations.

// Remove LoadBalancer flows and configurations.

func (p *proxier) installServiceGroup(svcPortName k8sproxy.ServicePortName, needUpdate, local, withSessionAffinity bool, endpoints []k8sproxy.Endpoint) (binding.GroupIDType, bool) {
	_ = "STUB: not implemented"
	return *new(binding.GroupIDType), false
}

// If the installation of the group fails, recycle it.

func (p *proxier) removeServiceGroup(svcPortName k8sproxy.ServicePortName, local bool) bool {
	_ = "STUB: not implemented"
	return false
}

// removeStaleEndpoints removes flows for the given Endpoints from the data path if these flows are no longer
// needed by any Service. Endpoints from different Services can have the same characteristics and thus
// can share the same flows. removeStaleEndpoints must be called whenever Endpoints are no longer used by a
// given Service. If the Endpoints are still referenced by any other Services, no flow will be removed.
// The method only returns an error if a data path operation fails. If the flows are successfully
// removed from the data path, the method returns nil.
func (p *proxier) removeStaleEndpoints(svcPortName k8sproxy.ServicePortName, protocol binding.Protocol, staleEndpoints map[string]k8sproxy.Endpoint) bool {
	_ = "STUB: not implemented"
	return false
}

// Get all Endpoints whose reference counter is 1, and these Endpoints should be removed.

// Remove flows for these Endpoints.

// Update the reference counter of Endpoints and remove them from the installed Endpoints of the ServicePortName.

func (p *proxier) removeStaleServiceConntrackEntries(svcPortName k8sproxy.ServicePortName, svcInfo *types.ServiceInfo) bool {
	_ = "STUB: not implemented"
	return false
}

// Clean up the UDP conntrack entries matching the stale Service IPs and ports. For a UDP Service without Endpoint,
// no UDP conntrack entry will have been generated, but there is no harm in calling this function.

func (p *proxier) removeStaleConntrackEntries(svcPortName k8sproxy.ServicePortName, pSvcInfo, svcInfo *types.ServiceInfo, staleEndpoints map[string]k8sproxy.Endpoint) bool {
	_ = "STUB: not implemented"
	return false
}

// If the port of the Service is changed, delete all conntrack entries related to the previous Service IPs and the
// previous Service port. These previous Service IPs includes external IPs, loadBalancer IPs and the ClusterIP.

// If the port of the Service is not changed, delete the conntrack entries related to the stale Service IPs and
// the Service port. These stale Service IPs could be clusterIP, externalIPs or loadBalancerIPs.

// If the NodePort of the Service is changed, delete the conntrack entries related to each of the Node IPs / the
// virtual IP to which NodePort traffic from external will be DNATed and the Service nodePort.

// Clean up the UDP conntrack entries matching the stale Service IPs and ports.

// Get all remaining Service IPs.

// Get all Node IPs.

// Clean up the UDP conntrack entries matching the remaining Service IPs and ports, and the stale Endpoint IPs.

func (p *proxier) addNewEndpoints(svcPortName k8sproxy.ServicePortName, protocol binding.Protocol, newEndpoints map[string]k8sproxy.Endpoint) bool {
	_ = "STUB: not implemented"
	return false
}

// Get all Endpoints whose reference counter is 0, and these Endpoints should be added.

// Add flows for these Endpoints.

// Update the reference counter of Endpoints.

func serviceIdentityChanged(svcInfo, pSvcInfo *types.ServiceInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func serviceExternalAddressesChanged(svcInfo, pSvcInfo *types.ServiceInfo) bool {
	_ = "STUB: not implemented"
	return false
}

// smallSliceDifference builds a slice which includes all the IPs from s1
// which are not in s2.
func smallSliceDifference(s1, s2 []net.IP) []net.IP { _ = "STUB: not implemented"; return nil }

// smallSliceSame builds a slice which includes all the IPs are both in s1 and s2.
func smallSliceSame(s1, s2 []net.IP) []net.IP { _ = "STUB: not implemented"; return nil }

func (p *proxier) installNodePortService(localGroupID, clusterGroupID binding.GroupIDType, svcPort uint16, protocol binding.Protocol, trafficPolicyLocal bool, affinityTimeout uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// Unsupported for NodePort
// Unsupported because external traffic has been DNAT'd in host network before it's forwarded to OVS.

func (p *proxier) uninstallNodePortService(svcPort uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *proxier) installExternalIPService(svcInfoStr string,
	localGroupID,
	clusterGroupID binding.GroupIDType,
	externalIPs []net.IP,
	svcPort uint16,
	protocol binding.Protocol,
	trafficPolicyLocal bool,
	affinityTimeout uint16,
	loadBalancerMode agentconfig.LoadBalancerMode) error {
	_ = "STUB: not implemented"
	return nil
}

// Unsupported for ExternalIP

func (p *proxier) uninstallExternalIPService(svcInfoStr string, externalIPs []net.IP, svcPort uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *proxier) installLoadBalancerService(svcInfoStr string,
	localGroupID,
	clusterGroupID binding.GroupIDType,
	loadBalancerIPs []net.IP,
	svcPort uint16,
	protocol binding.Protocol,
	trafficPolicyLocal bool,
	affinityTimeout uint16,
	loadBalancerMode agentconfig.LoadBalancerMode) error {
	_ = "STUB: not implemented"
	return nil
}

// Unsupported for LoadBalancerIP

func (p *proxier) uninstallLoadBalancerService(svcInfoStr string, loadBalancerIPs []net.IP, svcPort uint16, protocol binding.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *proxier) installServices() { _ = "STUB: not implemented"; return }

// Need to update.

// The changes to serviceIdentity, session affinity config, and traffic policies affect all Service
// flows while the changes to external addresses (NodePort and LoadBalancerIPs) affect external Service
// flows only.

// All Service flows use it.
// All Service flows use it.
// It affects the group ID used by external Service flows.
// It affects the group ID used by internal Service flows.

// We clean the UDP conntrack entries for the following Service update cases:
// - Service port changed, clean the conntrack entries matched by each of the current clusterIP / externalIPs
//   / loadBalancerIPs and the stale Service port.
// - ClusterIP changed, clean the conntrack entries matched by the clusterIP and the Service port.
// - Some externalIPs / loadBalancerIPs are removed, clean the conntrack entries matched by each of the
//   removed Service IPs and the current Service port.
// - Service nodePort changed, clean the conntrack entries matched by each of the Node IPs / the virtual
//   NodePort DNAT IP and the stale Service nodePort.
// However, we DO NOT clean the UDP conntrack entries related to remote Endpoints that are still
// referenced by the Service but are no longer selectable Endpoints for the corresponding Service IPs
// (for externalTrafficPolicy, these IPs are loadBalancerIPs, externalIPs and NodeIPs; for
// internalTrafficPolicy, these IPs clusterIPs) when externalTrafficPolicy or internalTrafficPolicy is
// changed from Cluster to Local. Consequently, the connections, which are supposed to select local
// Endpoints, will continue to send packets to remote Endpoints due to the existing UDP conntrack entries
// until timeout.

// Need to install.

// We need to ensure a group is created for a new Service even if there is no available Endpoints,
// otherwise it would fail to install Service flows because the group doesn't exist.

// Get the stale Endpoints and new Endpoints based on the diff of endpointsInstalled and allReachableEndpoints.

// We also clean the conntrack entries related to the stale Endpoints for a UDP Service. Conntrack entries
// matched by each of stale Endpoint IPs and each of the remaining Service IPs and ports will be deleted.

// categorizeEndpoints has checked if localGroup and clusterGroup should exist. We just create the group if its
// Endpoints is not nil.
// Note that nil represents the group should not exist and empty represents the group should exist but there is
// no available Endpoints.

// Delete previous flows.

// getLoadBalancerMode returns the default load balancer mode if the Service doesn't have the annotation overriding it.
// Otherwise, it returns the mode specified in the annotation.
func (p *proxier) getLoadBalancerMode(svcInfo *types.ServiceInfo) agentconfig.LoadBalancerMode {
	_ = "STUB: not implemented"
	return *new(agentconfig.LoadBalancerMode)
}

func getAffinityTimeout(svcInfo *types.ServiceInfo) uint16 { _ = "STUB: not implemented"; return 0 }

// SessionAffinity timeout is implemented using a hard_timeout in
// OVS. hard_timeout is represented by a uint16 in the OpenFlow protocol,
// hence we cannot support timeouts greater than 65535 seconds. However, the
// K8s Service spec allows timeout values up to 86400 seconds
// (https://godoc.org/k8s.io/api/core/v1#ClientIPConfig). For values greater
// than 65535 seconds, we need to set the hard_timeout to 65535 rather than
// let the timeout value wrap around.

func (p *proxier) installServiceFlows(svcInfo *types.ServiceInfo, localGroupID, clusterGroupID binding.GroupIDType) bool {
	_ = "STUB: not implemented"
	return false
}

// Check the `IsNested` field only when Proxy is enabled with `supportNestedService`.
// It is true only when the Service is an Antrea Multi-cluster Service for now.

// Install ClusterIP flows.

// not applicable for ClusterIP

// Install NodePort flows and configurations.

// Install ExternalIP flows and configurations.

// Install LoadBalancer flows and configurations.

func (p *proxier) updateServiceExternalAddresses(pSvcInfo, svcInfo *types.ServiceInfo, localGroupID, clusterGroupID binding.GroupIDType) bool {
	_ = "STUB: not implemented"
	return false
}

// The deleted ExternalIPs need to be removed from the map explicitly while the added ExternalIPs (they are always included in the current ExternalIPs) will be added at the end of installServices.

// The deleted LoadBalancerIPs need to be removed from the map explicitly while the added LoadBalancerIPs (they are always included in the current LoadBalancerIPs) will be added at the end of installServices.

func compareEndpoints(endpointsCached map[string]k8sproxy.Endpoint, endpointsInstalled []k8sproxy.Endpoint) (map[string]k8sproxy.Endpoint, map[string]k8sproxy.Endpoint) {
	_ = "STUB: not implemented"
	// Map endpointsToRemove is used to store the Endpoints that should be removed.
	return nil, nil
}

// Map endpointsToAdd is used to store the Endpoints that are newly added.

// Copy every Endpoint in endpointsCached to endpointsToRemove. After removing all actually installed Endpoints,
// only stale Endpoints are left.

// If the Endpoint is in the map endpointsCached, then it is not newly installed, remove it from map endpointsToRemove;
// otherwise, add it to map endpointsToAdd.

// syncProxyRules applies current changes in change trackers and then updates
// flows for services and endpoints. It will return immediately if either
// endpoints or services resources are not synced. syncProxyRules is only called
// through the Run method of the runner object, and all calls are serialized.
// This method is the only one that changes internal state, but
// GetServiceFlowKeys(), which is called by the "/ovsflows" API handler,
// also reads service and endpoints maps, so mu is used to protect these two maps.
func (p *proxier) syncProxyRules() error { _ = "STUB: not implemented"; return nil }

// Protect Service and endpoints maps, which can be read by
// GetServiceFlowKeys().

func (p *proxier) SyncLoop() { _ = "STUB: not implemented"; return }

// Sync is called to synchronize the proxier as soon as possible.
func (p *proxier) Sync() { _ = "STUB: not implemented"; return }

func (p *proxier) OnEndpointSliceAdd(endpointSlice *discovery.EndpointSlice) {
	_ = "STUB: not implemented"
	return
}

func (p *proxier) OnEndpointSliceUpdate(oldEndpointSlice, newEndpointSlice *discovery.EndpointSlice) {
	_ = "STUB: not implemented"
	return
}

func (p *proxier) OnEndpointSliceDelete(endpointSlice *discovery.EndpointSlice) {
	_ = "STUB: not implemented"
	return
}

func (p *proxier) OnEndpointSlicesSynced() { _ = "STUB: not implemented"; return }

func (p *proxier) matchAddressFamily(eps *discovery.EndpointSlice) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *proxier) OnServiceAdd(service *corev1.Service) { _ = "STUB: not implemented"; return }

func (p *proxier) serviceSupportsIPFamily(preSvc, curSvc *corev1.Service) bool {
	_ = "STUB: not implemented"
	// Prefer current Service if available.
	return false
}

func (p *proxier) OnServiceUpdate(oldService, service *corev1.Service) {
	_ = "STUB: not implemented"
	return
}

func (p *proxier) OnServiceDelete(service *corev1.Service) { _ = "STUB: not implemented"; return }

func (p *proxier) OnServiceSynced() { _ = "STUB: not implemented"; return }

// OnTopologyChange is called whenever this node's proxy relevant topology-related labels change.
func (p *proxier) OnTopologyChange(topologyLabels map[string]string) {
	_ = "STUB: not implemented"
	return
}

// OnServiceCIDRsChanged is called whenever a change is observed
// in any of the ServiceCIDRs, and provides complete list of service cidrs.
func (p *proxier) OnServiceCIDRsChanged(_ []string) { _ = "STUB: not implemented"; return }

func (p *proxier) GetServiceByIP(serviceStr string) (k8sproxy.ServicePortName, bool) {
	_ = "STUB: not implemented"
	return *new(k8sproxy.ServicePortName), false
}

func (p *proxier) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (p *ProxyServer) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (p *proxier) GetServiceFlowKeys(serviceName, namespace string) ([]string, []binding.GroupIDType, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// Service flows not installed.

func (p *ProxyServer) HandlePacketIn(pktIn *ofctrl.PacketIn) error {
	_ = "STUB: not implemented"
	return nil
}

// Get Ethernet data.

// It cannot use CONTROLLER (the default value when inPort is 0) as the inPort due to a bug in Windows ovsext
// driver, otherwise the Windows OS would crash. See https://github.com/openvswitch/ovs-issues/issues/280.

// newProxier returns a new single-stack proxier.
func newProxier(
	hostname string,
	ofClient openflow.Client,
	ipFamily corev1.IPFamily,
	routeClient route.Interface,
	nodeIPChecker nodeip.Checker,
	nodePortAddresses []net.IP,
	proxyAllEnabled bool,
	skipServices []string,
	proxyLoadBalancerIPs bool,
	defaultLoadBalancerMode agentconfig.LoadBalancerMode,
	groupCounter types.GroupCounter,
	supportNestedService bool,
	serviceHealthServerDisabled bool,
	preferSameTrafficDistributionEnabled bool,
	serviceLabelSelector labels.Selector,
	healthzServer *healthcheck.ProxyHealthServer,
) (*proxier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// metaProxierWrapper wraps metaProxier, and implements the extra methods added
// in interface ProxyQuerier.
type metaProxierWrapper struct {
	k8sproxy.Provider
	ipv4Proxier *proxier
	ipv6Proxier *proxier
}

func (p *metaProxierWrapper) GetServiceFlowKeys(serviceName, namespace string) ([]string, []binding.GroupIDType, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// Return the unions of IPv4 and IPv6 flows and groups.

func (p *metaProxierWrapper) GetServiceByIP(serviceStr string) (k8sproxy.ServicePortName, bool) {
	_ = "STUB: not implemented"
	// Format of serviceStr is <clusterIP>:<svcPort>/<protocol>.
	return *new(k8sproxy.ServicePortName), false
}

func newDualStackProxier(
	hostname string,
	ofClient openflow.Client,
	routeClient route.Interface,
	nodeIPChecker nodeip.Checker,
	nodePortAddressesIPv4 []net.IP,
	nodePortAddressesIPv6 []net.IP,
	proxyAllEnabled bool,
	skipServices []string,
	proxyLoadBalancerIPs bool,
	defaultLoadBalancerMode agentconfig.LoadBalancerMode,
	v4groupCounter types.GroupCounter,
	v6groupCounter types.GroupCounter,
	nestedServiceSupport bool,
	serviceHealthServerDisabled bool,
	preferSameTrafficDistributionEnabled bool,
	serviceLabelSelector labels.Selector,
	healthzServer *healthcheck.ProxyHealthServer,
) (Proxier, error) {
	_ = "STUB: not implemented"
	// Create an IPv4 instance of the single-stack proxier.
	return *new(Proxier), nil
}

// Create an IPv6 instance of the single-stack proxier.

// Create a meta-proxier that dispatch calls between the two
// single-stack proxier instances.

func generateServiceLabelSelector(serviceProxyName string) labels.Selector {
	_ = "STUB: not implemented"
	// TODO: The label selector nonHeadlessServiceSelector was added to pass the Kubernetes e2e test
	//  'Services should implement service.kubernetes.io/headless'. You can find the test case at:
	//  https://github.com/kubernetes/kubernetes/blob/027ac5a426a261ba6b66a40e79e123e75e9baf5b/test/e2e/network/service.go#L2281
	//  However, in AntreaProxy, headless Services are skipped by checking the ClusterIP.
	return *new(labels.Selector)
}

func NewProxyServer(hostname string,
	nodeManager *k8sproxy.NodeManager,
	ofClient openflow.Client,
	routeClient route.Interface,
	nodeIPChecker nodeip.Checker,
	v4Enabled bool,
	v6Enabled bool,
	nodePortAddressesIPv4 []net.IP,
	nodePortAddressesIPv6 []net.IP,
	proxyConfig antreaconfig.AntreaProxyConfig,
	defaultLoadBalancerMode agentconfig.LoadBalancerMode,
	v4GroupCounter types.GroupCounter,
	v6GroupCounter types.GroupCounter,
	nestedServiceSupport bool) (*ProxyServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ProxyServer) Initialize(ctx context.Context,
	serviceInformer coreinformers.ServiceInformer,
	endpointSliceInformer discoveryinformers.EndpointSliceInformer) {
	_ = "STUB: not implemented"
	return
}

func needClearConntrackEntries(protocol binding.Protocol) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *proxier) isIPv6() bool { _ = "STUB: not implemented"; return false }

func serveHealthz(ctx context.Context, hz *healthcheck.ProxyHealthServer) {
	_ = "STUB: not implemented"
	return
}

func (p *ProxyServer) GetProxyQuerier() ProxyQuerier {
	_ = "STUB: not implemented"
	return *new(ProxyQuerier)
}

func (p *ProxyServer) GetProxyProvider() Proxier { _ = "STUB: not implemented"; return *new(Proxier) }

func newIPToServiceMap() *ipToServiceMap { _ = "STUB: not implemented"; return nil }

// ipToServiceMap is a thread-safe store for Service lookup, providing
// access to Service names based IPs such as external IPs, loadBalancer IPs
// and the ClusterIP.
type ipToServiceMap struct {
	// serviceStringMapMutex protects serviceStringMap object.
	serviceStringMapMutex sync.RWMutex
	// serviceStringMap provides map from serviceString(IP:Port/Protocol) to ServicePortName.
	serviceStringMap map[string]k8sproxy.ServicePortName
}

// add registers a new Service to the map.
func (m *ipToServiceMap) add(serviceInfo *types.ServiceInfo, servicePortName k8sproxy.ServicePortName) {
	_ = "STUB: not implemented"
	return
}

// delete removes the Service from the map with thread safety.
func (m *ipToServiceMap) delete(serviceInfo *types.ServiceInfo) { _ = "STUB: not implemented"; return }

// deleteServiceIPs removes the associated keys from the map for the given set
// of Service IPs.
//
// Deleting a key that does not exist is safely ignored.
func (m *ipToServiceMap) deleteServiceIPs(serviceStrings []string) {
	_ = "STUB: not implemented"
	return
}

// get retrieves the associated Service given it's serviceStr of the format
// (IP:Port/Protocol) where IP can be ExternalIPs, LoadBalancerIPs and ClusterIP.
func (m *ipToServiceMap) get(serviceStr string) (k8sproxy.ServicePortName, bool) {
	_ = "STUB: not implemented"
	return *new(k8sproxy.ServicePortName), false
}

// getServiceIPStrings returns a slice of serviceStrings with the format
// "IP:Port/Protocol" for all Service IPs.
func getServiceIPStrings(s *types.ServiceInfo) []string { _ = "STUB: not implemented"; return nil }

// +1 for ClusterIP

// GenerateServiceStrings generates service strings for the given IPs in the format
// "IP:Port/Protocol".
func generateServiceInfoStrings(protocol corev1.Protocol, port int, ips []net.IP) []string {
	_ = "STUB: not implemented"
	return nil
}
