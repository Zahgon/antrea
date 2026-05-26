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

package cniserver

import (
	"net"
	"time"

	"antrea.io/libOpenflow/openflow15"
	current "github.com/containernetworking/cni/pkg/types/100"
	corev1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/events"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/cniserver/ipam"
	"antrea.io/antrea/v2/pkg/agent/cniserver/types"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/agent/route"
	"antrea.io/antrea/v2/pkg/agent/util"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
	"antrea.io/antrea/v2/pkg/util/channel"
	"antrea.io/antrea/v2/pkg/util/wait"
)

type vethPair struct {
	name      string
	ifIndex   int
	peerIndex int
}

const (
	ovsExternalIDMAC          = "attached-mac"
	ovsExternalIDIP           = "ip-address"
	ovsExternalIDContainerID  = "container-id"
	ovsExternalIDPodName      = "pod-name"
	ovsExternalIDPodNamespace = "pod-namespace"
	ovsExternalIDIFDev        = "if-dev"
	ovsExternalIDNetNS        = "net-ns"
)

const (
	defaultIFDevName = "eth0"
)

var (
	getNSPath = util.GetNSPath
	// retryInterval is the interval to re-install Pod OpenFlow entries if any error happened.
	// Note, using a variable rather than constant for retryInterval because we may use a shorter time in the
	// test code.
	retryInterval = 5 * time.Second
)

type podConfigurator struct {
	ovsBridgeClient ovsconfig.OVSBridgeClient
	ofClient        openflow.Client
	routeClient     route.Interface
	ifaceStore      interfacestore.InterfaceStore
	gatewayMAC      net.HardwareAddr
	ifConfigurator  podInterfaceConfigurator
	// podUpdateNotifier is used for notifying updates of local Pods to other components which may benefit from this
	// information, i.e. NetworkPolicyController, EgressController.
	podUpdateNotifier channel.Notifier
	// isSecondaryNetwork is true if this instance of podConfigurator is used to configure
	// Pod secondary network interfaces.
	isSecondaryNetwork bool

	containerAccess  *containerAccessArbitrator
	eventBroadcaster events.EventBroadcaster
	recorder         events.EventRecorder
	podListerSynced  cache.InformerSynced
	podLister        v1.PodLister
	kubeClient       clientset.Interface
	unreadyPortQueue workqueue.TypedDelayingInterface[string]
	statusCh         chan *openflow15.PortStatus
}

func newPodConfigurator(
	kubeClient clientset.Interface,
	ovsBridgeClient ovsconfig.OVSBridgeClient,
	ofClient openflow.Client,
	routeClient route.Interface,
	ifaceStore interfacestore.InterfaceStore,
	gatewayMAC net.HardwareAddr,
	ovsDatapathType ovsconfig.OVSDatapathType,
	isOvsHardwareOffloadEnabled bool,
	disableTXChecksumOffload bool,
	podUpdateNotifier channel.Notifier,
	podInformer cache.SharedIndexInformer,
	containerAccess *containerAccessArbitrator,
) (*podConfigurator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initiate the PortStatus message listener. This function is a no-op except on Windows.

func parseContainerIPs(ipcs []*current.IPConfig) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildContainerConfig(
	interfaceName, containerID, podName, podNamespace, netNS string,
	containerIface *current.Interface,
	ips []*current.IPConfig,
	vlanID uint16) *interfacestore.InterfaceConfig {
	_ = "STUB: not implemented"
	// A secondary interface can be created without IPs. Ignore the IP parsing error here.
	return nil
}

// containerIface.Mac should be a valid MAC string, otherwise it should throw error before

// BuildOVSPortExternalIDs parses OVS port external_ids from InterfaceConfig.
// external_ids are used to compare and sync container interface configuration.
func BuildOVSPortExternalIDs(containerConfig *interfacestore.InterfaceConfig) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Save interface name for a secondary interface.

// Save NetNS which is needed for secondary interface creation.

func getContainerIPsString(ips []net.IP) string { _ = "STUB: not implemented"; return "" }

// ParseOVSPortInterfaceConfig reads the Pod properties saved in the OVS port
// external_ids, initializes and returns an InterfaceConfig struct.
// nil will be returned, if the OVS port does not have external IDs or it is
// not created for a Pod interface.
func ParseOVSPortInterfaceConfig(portData *ovsconfig.OVSPortData, portConfig *interfacestore.OVSPortConfig) *interfacestore.InterfaceConfig {
	_ = "STUB: not implemented"
	return nil
}

// A secondary interface may not have an IP assigned.

func (pc *podConfigurator) configureInterfacesCommon(
	podName, podNamespace, containerID, containerNetNS string,
	containerIFDev string, mtu int, sriovVFDeviceID string,
	result *ipam.IPAMResult, containerAccess *containerAccessArbitrator, mac net.HardwareAddr) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete veth pair if any failure occurs in later manipulation.

// Not needed for a secondary network interface.

// Note that the IP address should be advertised after Pod OpenFlow entries are installed, otherwise the packet might
// be dropped by OVS.

// Do not return an error and fail the interface creation.

// Mark the manipulation as success to cancel deferred operations.

func (pc *podConfigurator) createOVSPort(ovsPortName string, ovsAttachInfo map[string]interface{}, vlanID uint16) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (pc *podConfigurator) removeInterfaces(containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Deleting veth devices and OVS port must be called after Openflows are uninstalled.
// Otherwise there could be a race condition:
// 1. Pod A's ofport was released
// 2. Pod B got the ofport released above
// 3. Flows for Pod B were installed
// 4. Flows for Pod A were uninstalled
// Because Pod A and Pod B had same ofport, they had overlapping flows, e.g. the
// classifier flow in table 0 which has only in_port as the match condition, then
// step 4 can remove flows owned by Pod B by mistake.
// Note that deleting the interface attached to an OVS port can release the ofport.

func (pc *podConfigurator) checkInterfaces(
	containerID, containerNetNS string,
	containerIface *current.Interface,
	prevResult *current.Result, sriovVFDeviceID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *podConfigurator) checkHostInterface(
	containerID string,
	containerIntf *current.Interface,
	containerIfKind interface{},
	containerIPs []*current.IPConfig,
	interfaces []*current.Interface,
	sriovVFDeviceID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *podConfigurator) validateOVSInterfaceConfig(containerID string, containerMAC string, ips []*current.IPConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func parsePrevResult(conf *types.NetworkConfig) error { _ = "STUB: not implemented"; return nil }

func (pc *podConfigurator) reconcile(pods []corev1.Pod, containerAccess *containerAccessArbitrator, podNetworkWait, flowRestoreCompleteWait *wait.Group) error {
	_ = "STUB: not implemented"
	// desiredPods is the set of Pods that should be present, based on the
	// current list of Pods got from the Kubernetes API.
	return nil
}

// desiredPodIPs is the set of IPs allocated to desiredPods.

// knownInterfaces is the list of interfaces currently in the local cache.

// Find the OVS ports corresponding to Pods which no longer exist. This includes the case that the Pod using the
// Namespaced name constructed from OVSDB does not exist in kube-apiserver, and the case that a Pod with the
// same Namespaced name exists in kube-apiserver but the host interface is disconnected.
// Note: a Pod's host interface is generated by both Pod name and the sandbox container ID, so the new Pod with
// the same name would have a different host interface associated to it.

// This code is executed synchronously as part of cniServer.Initialize and prior to the call to cniServer.Run.
// As a result, concurrent calls to CNI ADD / DEL are not possible.

// This should only happen on Windows Nodes because if this condition (OFPort is -1) is satisfied on Linux,
// the call to isInterfaceInvalid above will return true and the interface will be removed, and this code will
// not be executed. On Windows, an OVS port may be created without connecting to the host interface when agent
// is down, so OVS has no chance to allocate a valid OpenFlow port to the interface.

// Do not install Pod flows until all preconditions are met.

// To avoid race condition with CNIServer CNI event handlers.

// This interface matches an existing Pod.
// We rely on the interface cache / store - which is initialized from the persistent
// OVSDB - to map the Pod to its interface configuration. The interface
// configuration includes the parameters we need to replay the flows.

// clean-up IPs that may still be allocated

// disconnectInterfaceFromOVS disconnects an existing interface from ovs br-int.
func (pc *podConfigurator) disconnectInterfaceFromOVS(containerConfig *interfacestore.InterfaceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// We should not delete OVS port if Pod flows deletion fails, otherwise
// it is possible a new Pod will reuse the reclaimed ofport number, and
// the OVS flows added for the new Pod can conflict with the stale
// flows of the deleted Pod.

// TODO: handle error and introduce garbage collection for failure on deletion

// Remove container configuration from cache.

// connectInterceptedInterface connects intercepted interface to ovs br-int.
func (pc *podConfigurator) connectInterceptedInterface(
	podName string,
	podNamespace string,
	containerID string,
	containerNetNS string,
	containerIFDev string,
	containerIPs []*current.IPConfig,
	containerAccess *containerAccessArbitrator,
) error {
	_ = "STUB: not implemented"
	return nil
}

// disconnectInterceptedInterface disconnects intercepted interface from ovs br-int.
func (pc *podConfigurator) disconnectInterceptedInterface(podName, podNamespace, containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO recover pre-connect state? repatch vethpair to original bridge etc ?? to make first CNI happy??

func (pc *podConfigurator) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// Put the item back on the workqueue to handle any transient errors.

func (pc *podConfigurator) updateUnreadyPod(ovsPort string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the InterfaceConfig again after the lock to avoid race conditions.

// Add Pod not-ready event if the pod flows are not successfully installed, and the OpenFlow port is not allocated.
// Returns error so that we can have a retry after 5s.

// Install OpenFlow entries for the Pod.

// Add Pod not-ready event if the pod flows installation fails.
// Returns error so that we can have a retry after 5s.

// Notify the Pod update event to required components.

func (pc *podConfigurator) recordPodEvent(ifConfig *interfacestore.InterfaceConfig, installed bool) {
	_ = "STUB: not implemented"
	return
}

// Add normal event to record Pod network is ready.

func (pc *podConfigurator) processPortStatusMessage(status *openflow15.PortStatus) {
	_ = "STUB: not implemented"
	return
}

// Update Pod OpenFlow entries only after the OpenFlow port state is live or down.
// Accepting Port state "openflow15.PS_LINK_DOWN" is a workaround for Windows OVS issue https://github.com/openvswitch/ovs-issues/issues/351.
// In which OVS does not correctly implement function netdev_windows_update_flags, so OVS doesn't update ifp_flags
// after a new OpenFlow port is successfully installed. Since this OVS issue doesn't have side impact on datapath
// packets forwarding, antrea-agent will ignore the bad state to ensure the Pod's OpenFlow entries are installed as
// long as the port number is allocated.

// Get the InterfaceConfig again after the lock to avoid race conditions.

// Update interface config with the ofPort.
