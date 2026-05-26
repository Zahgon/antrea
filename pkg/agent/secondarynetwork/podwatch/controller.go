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

package podwatch

import (
	"net"
	"sync"
	"time"

	current "github.com/containernetworking/cni/pkg/types/100"
	netdefv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	netdefclient "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/client/clientset/versioned/typed/k8s.cni.cncf.io/v1"
	netdefutils "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/cniserver/ipam"
	cnitypes "antrea.io/antrea/v2/pkg/agent/cniserver/types"
	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	crdv1b1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
	"antrea.io/antrea/v2/pkg/util/channel"
)

const (
	controllerName = "SecondaryNetworkController"
	minRetryDelay  = 2 * time.Second
	maxRetryDelay  = 120 * time.Second
	numWorkers     = 4
	// Set resyncPeriod to 0 to disable resyncing.
	resyncPeriod = 0 * time.Minute
)

const (
	resourceNameAnnotationKey = "k8s.v1.cni.cncf.io/resourceName"
	startIfaceIndex           = 1
	endIfaceIndex             = 101

	interfaceDefaultMTU = 1500
	vlanIDMax           = 4094
)

type InterfaceConfigurator interface {
	ConfigureSriovSecondaryInterface(podName, podNamespace, containerID, containerNetNS, containerInterfaceName string, mtu int, podSriovVFDeviceID string, result *current.Result) error
	DeleteSriovSecondaryInterface(interfaceConfig *interfacestore.InterfaceConfig) error
	ConfigureVLANSecondaryInterface(podName, podNamespace, containerID, containerNetNS, containerInterfaceName string, mtu int, ipamResult *ipam.IPAMResult, mac net.HardwareAddr) error
	DeleteVLANSecondaryInterface(interfaceConfig *interfacestore.InterfaceConfig) error
}

type IPAMAllocator interface {
	SecondaryNetworkAllocate(podOwner *crdv1b1.PodOwner, networkConfig *cnitypes.NetworkConfig) (*ipam.IPAMResult, error)
	SecondaryNetworkRelease(podOwner *crdv1b1.PodOwner) error
}

type podCNIInfo struct {
	containerID string
	netNS       string
}

type PodController struct {
	kubeClient            clientset.Interface
	netAttachDefClient    netdefclient.K8sCniCncfIoV1Interface
	queue                 workqueue.TypedRateLimitingInterface[string]
	podInformer           cache.SharedIndexInformer
	podLister             corelisters.PodLister
	ipPoolLister          crdlisters.IPPoolLister
	podUpdateSubscriber   channel.Subscriber
	ovsBridgeClient       ovsconfig.OVSBridgeClient
	interfaceStore        interfacestore.InterfaceStore
	primaryInterfaceStore interfacestore.InterfaceStore
	interfaceConfigurator InterfaceConfigurator
	ipamAllocator         IPAMAllocator
	// Map from "namespace/pod" to podCNIInfo.
	cniCache           sync.Map
	vfDeviceIDUsageMap sync.Map
	nodeConfig         *config.NodeConfig
}

func NewPodController(
	kubeClient clientset.Interface,
	netAttachDefClient netdefclient.K8sCniCncfIoV1Interface,
	podInformer cache.SharedIndexInformer,
	podUpdateSubscriber channel.Subscriber,
	primaryInterfaceStore interfacestore.InterfaceStore,
	nodeConfig *config.NodeConfig,
	ovsBridgeClient ovsconfig.OVSBridgeClient,
	ipPoolLister crdlisters.IPPoolLister,
) (*PodController, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// podUpdateSubscriber can be nil with test code.

// Subscribe Pod CNI add/del events.

func podKeyGet(podName, podNamespace string) string { _ = "STUB: not implemented"; return "" }

func allocatePodSecondaryIfaceName(usedIFNames sets.Set[string]) (string, error) {
	_ = "STUB: not implemented"
	// Generate new interface name (eth1,eth2..eth100) and return to caller.
	return "", nil
}

func (pc *PodController) enqueuePod(obj interface{}) { _ = "STUB: not implemented"; return }

// processCNIUpdate will be called when CNIServer publishes a Pod update event.
func (pc *PodController) processCNIUpdate(e interface{}) { _ = "STUB: not implemented"; return }

// handleAddUpdatePod handles Pod Add, Update events and updates annotation if required.
func (pc *PodController) handleAddUpdatePod(pod *corev1.Pod, podCNIInfo *podCNIInfo, storedSecondaryInterfaces []*interfacestore.InterfaceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Primary network configuration is not complete yet. Return nil here to dequeue the
// Pod event. Secondary network configuration will be handled with the following Pod
// update events.

// Parse Pod annotation and proceed with the secondary network configuration.

// Do not return an error as a retry is not appropriate.
// When the annotation is fixed, the Pod will be enqueued again.

// We do not support secondary network update at the moment. Return as long as one
// secondary interface has been created for the Pod.

// Intentionally ignore errors from updating the Pod's network status annotation here.
// Failure to update the annotation does not affect the actual network setup for the Pod.
// The annotation is mainly used for status reporting and restoring SR-IOV interface
// information after agent restarts.

func (pc *PodController) deletePodNetworkStatusAnnotation(pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	setNetworkStatus = netdefutils.SetNetworkStatus
)

// updatePodNetworkStatusAnnotation update the Pod's network status annotation
func (pc *PodController) updatePodNetworkStatusAnnotation(netStatus []netdefv1.NetworkStatus, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the Pod's network status annotation

func (pc *PodController) removeInterfaces(interfaces []*interfacestore.InterfaceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Since only VLAN and SR-IOV interfaces are supported by now, we judge the
// interface type by checking interfaceConfig.OVSPortConfig is set or not.

func (pc *PodController) syncPod(key string) error { _ = "STUB: not implemented"; return nil }

// Pod or its primary interface has been deleted. Remove secondary interfaces too.

// Interfaces created for a previous Pod with the same Namespace/name are
// not deleted yet. First delete them before processing the new Pod's
// secondary networks.

func (pc *PodController) Worker() { _ = "STUB: not implemented"; return }

func (pc *PodController) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// Configure Secondary Network Interface.
func (pc *PodController) configureSecondaryInterface(
	pod *corev1.Pod,
	network *netdefv1.NetworkSelectionElement,
	resourceName string,
	podCNIInfo *podCNIInfo,
	networkConfig *SecondaryNetworkConfig,
	mac net.HardwareAddr,
) (*current.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Interface creation failed. Free allocated IP address

// Let VLAN ID in the CNI network configuration override the IPPool subnet
// VLAN.

func (pc *PodController) configurePodSecondaryNetwork(pod *corev1.Pod, networkList []*netdefv1.NetworkSelectionElement, podCNIInfo *podCNIInfo) ([]netdefv1.NetworkStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NetworkAttachmentDefinition Spec.Config parsing failed. Do not retry.

// Ignore non-Antrea CNI type.

// This annotation is required for SRIOV devices, otherwise there is
// no way to make sure that we allocate the "right" type of device.

// It is probably worth retrying as the NetworkAttachmentDefinition
// may eventually be updated with the missing annotation.

// Generate a new interface name, if the secondary interface name was not provided in the
// Pod annotation.

// Do not return error: no need to requeue.

// Get and validate the MAC from annotation, ignore for SR-IOV

// No MAC requested, skip

// Do not return error: no need to requeue.

// Secondary network information retrieved from API server. Proceed to configure secondary interface now.

// Secondary interfaces are not for the default Pod network

// As we do not support secondary network update, do not return error to
// retry, if at least one secondary network is configured.

func validateNetworkConfig(cniConfig []byte) (*SecondaryNetworkConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: use the physical interface MTU as the default.

func (pc *PodController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Failure of initializeSRIOVSecondaryInterfaceStore() won't stop agent from starting.

func checkForPodSecondaryNetworkAttachment(pod *corev1.Pod) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// When a Kubernetes Node reboots and the OVSDB file is not properly restored,
// the primary and secondary OVS ports in OVSDB are lost. Pods may temporarily
// appear as "unknown" and later get recreated on the same Node with new container IDs.
// The cleanup of secondary interfaces during agent restart depends on the InterfaceStores
// being initialized with existing OVS ports. If the OVSDB file is missing, the
// primary InterfaceStore is empty, preventing the secondary InterfaceStore from
// initializing correctly. Consequently, secondary IPs assigned to old Pods (with
// previous container IDs) are not deleted. To resolve this, all IPPool allocations
// must be checked, and any IPs associated with non-existing container IDs should be
// released. Otherwise, stale secondary IPs remain in the IPPool and cannot be reused.
// We run it periodically to cover the case where a Pod is recreated with
// the same name on another Node and has any secondary IP as well.
func (pc *PodController) cleanUpStaleIPAddresses() { _ = "STUB: not implemented"; return }

// Only consider SecondaryNetwork interfaces

func (pc *PodController) initializeCNICache() { _ = "STUB: not implemented"; return }

// initializeOVSSecondaryInterfaceStore restores secondary interfaceStore for VLAN interfaces when agent restarts.
func (pc *PodController) initializeOVSSecondaryInterfaceStore() error {
	_ = "STUB: not implemented"
	// This is the case when secondary bridge is not configured and no VLAN interface at all.
	return nil
}

// initializeSRIOVSecondaryInterfaceStore restores secondary interfaceStore for SR-IOV interfaces
// when agent restarts. It will get the Pod info from the store of primary interfaces, and check
// the NetworkStatus annotation of a Pod, then restore the SR-IOV interfaces based on the NetworkAttachmentDefinition
// name and device's pci_address in the NetworkStatus.
func (pc *PodController) initializeSRIOVSecondaryInterfaceStore() error {
	_ = "STUB: not implemented"
	return nil
}

// Add the interface to the Secondary interfaceStore.

func parseIPs(ips []string) []net.IP { _ = "STUB: not implemented"; return nil }

// reconcileSecondaryInterfaces deletes stale secondary interfaces after agent restarts.
func (pc *PodController) reconcileSecondaryInterfaces() { _ = "STUB: not implemented"; return }

// secondaryInterfaces is the list of interfaces currently in the secondary local cache.

// Delete an interface when the primary interface has already been deleted,
// and delete the OVS port when a secondary interface is missing
// (OFPort == -1).
// In a normal case, a SR-IOV interface should not be included here, as the
// primary interface cannot be deleted until the Pod's SR-IOV interfaces are
// all deleted.

// If there are any stale interfaces, pass them to removeInterfaces()
