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

package trafficcontrol

import (

	// #nosec G505: not used for security purposes

	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha2"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha2"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1alpha2"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
	"antrea.io/antrea/v2/pkg/util/channel"
)

const (
	controllerName = "TrafficControlController"
	// How long to wait before retrying the processing of a TrafficControl change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing a TrafficControl change.
	defaultWorkers = 4
	// Disable resyncing.
	resyncPeriod time.Duration = 0

	// Default VXLAN tunnel destination port.
	defaultVXLANTunnelDestinationPort = int32(4789)
	// Default GENEVE tunnel destination port.
	defaultGENEVETunnelDestinationPort = int32(6081)

	portNamePrefixVXLAN  = "vxlan"
	portNamePrefixGENEVE = "geneve"
	portNamePrefixGRE    = "gre"
	portNamePrefixERSPAN = "erspan"
)

var (
	trafficControlPortExternalIDs = map[string]interface{}{
		interfacestore.AntreaInterfaceTypeKey: interfacestore.AntreaTrafficControl,
	}
)

// trafficControlState keeps the actual state of a TrafficControl that has been realized.
type trafficControlState struct {
	// The actual name of target port used by a TrafficControl.
	targetPortName string
	// The actual openflow port for which we have installed for a TrafficControl.
	targetOFPort uint32
	// The actual name of return port used by a TrafficControl.
	returnPortName string
	// The actual action of a TrafficControl.
	action v1alpha2.TrafficControlAction
	// The actual direction of a TrafficControl.
	direction v1alpha2.Direction
	// The actual openflow ports for which we have installed flows for a TrafficControl. Note that, flows are only installed
	// for the Pods whose effective TrafficControl is the current TrafficControl, and the ports are these Pods'.
	ofPorts sets.Set[int32]
	// The actual Pods applied with the TrafficControl. Note that, a TrafficControl can be either effective TrafficControl
	// or alternative TrafficControl for these Pods.
	pods sets.Set[string]
}

// podToTCBinding keeps the TrafficControls applied to a Pod. There is only one effective TrafficControl for a Pod at any
// given time.
type podToTCBinding struct {
	effectiveTC    string
	alternativeTCs sets.Set[string]
}

// portToTCBinding keeps the TrafficControls using an OVS port.
type portToTCBinding struct {
	interfaceConfig *interfacestore.InterfaceConfig
	trafficControls sets.Set[string]
}

type Controller struct {
	ofClient openflow.Client

	portToTCBindings   map[string]*portToTCBinding
	ovsBridgeClient    ovsconfig.OVSBridgeClient
	ovsCtlClient       ovsctl.OVSCtlClient
	ovsPortUpdateMutex sync.Mutex

	interfaceStore interfacestore.InterfaceStore

	podInformer     cache.SharedIndexInformer
	podLister       corelisters.PodLister
	podListerSynced cache.InformerSynced

	namespaceInformer     cache.SharedIndexInformer
	namespaceLister       corelisters.NamespaceLister
	namespaceListerSynced cache.InformerSynced

	podToTCBindings      map[string]*podToTCBinding
	podToTCBindingsMutex sync.RWMutex

	tcStates      map[string]*trafficControlState
	tcStatesMutex sync.RWMutex

	trafficControlInformer     cache.SharedIndexInformer
	trafficControlLister       crdlisters.TrafficControlLister
	trafficControlListerSynced cache.InformerSynced
	queue                      workqueue.TypedRateLimitingInterface[string]
}

func NewTrafficControlController(ofClient openflow.Client,
	interfaceStore interfacestore.InterfaceStore,
	ovsBridgeClient ovsconfig.OVSBridgeClient,
	ovsCtlClient ovsctl.OVSCtlClient,
	tcInformer crdinformers.TrafficControlInformer,
	podInformer cache.SharedIndexInformer,
	namespaceInformer coreinformers.NamespaceInformer,
	podUpdateSubscriber channel.Subscriber) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// processPodUpdate will be called when CNIServer publishes a Pod update event, and the event of TrafficControl which is
// the effective one of the Pod is triggered.
func (c *Controller) processPodUpdate(e interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) matchedPod(pod *v1.Pod, to *v1alpha2.AppliedTo) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) filterAffectedTCsByPod(pod *v1.Pod) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) addPod(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updatePod(oldObj interface{}, obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) deletePod(obj interface{}) { _ = "STUB: not implemented"; return }

func matchedNamespace(namespace *v1.Namespace, to *v1alpha2.AppliedTo) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) filterAffectedTCsByNS(namespace *v1.Namespace) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) addNamespace(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updateNamespace(oldObj, obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) addTC(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updateTC(oldObj interface{}, obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) deleteTC(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *Controller) worker() { _ = "STUB: not implemented"; return }

func (c *Controller) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// If no error occurs we Forget this item, so it does not get queued again until
// another change happens.

// Put the item back on the work queue to handle any transient errors.

func (c *Controller) newTrafficControlState(tcName string, action v1alpha2.TrafficControlAction, direction v1alpha2.Direction) *trafficControlState {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) getTrafficControlState(tcName string) (*trafficControlState, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Controller) deleteTrafficControlState(tcName string) { _ = "STUB: not implemented"; return }

func (c *Controller) filterPods(appliedTo *v1alpha2.AppliedTo) ([]*v1.Pod, error) {
	_ = "STUB: not implemented"
	// If both selectors are nil, no Pod should be selected.
	return nil, nil
}

// If Pod selector is not nil, use it to select Pods.

// If Pod selector is nil, then Namespace selector will not be nil, select all Pods from the selected Namespaces.

// If Namespace selector is not nil, use it to select Namespaces.

// Select Pods with Pod selector from the selected Namespaces.

// If Namespace selector is nil, use Pod selector to select Pods from all Namespaces.

// TrafficControl does not support host network Pods.

func genVXLANPortName(tunnel *v1alpha2.UDPTunnel) string {
	_ = "STUB: not implemented"
	// #nosec G401: not used for security purposes
	return ""
}

func genGENEVEPortName(tunnel *v1alpha2.UDPTunnel) string {
	_ = "STUB: not implemented"
	// #nosec G401: not used for security purposes
	return ""
}

func genGREPortName(tunnel *v1alpha2.GRETunnel) string {
	_ = "STUB: not implemented"
	// #nosec G401: not used for security purposes
	return ""
}

// genERSPANPortName generates a port name for the given ERSPAN tunnel.
// Note that ERSPAN tunnel's uniqueness is based on the remote IP and the session ID only, which means if there are two
// tunnels having same remote IP and session ID but different other attributes, creating the second port would fail in
// OVS.
func genERSPANPortName(tunnel *v1alpha2.ERSPANTunnel) string {
	_ = "STUB: not implemented"
	// #nosec G401: not used for security purposes
	return ""
}

func ParseTrafficControlInterfaceConfig(portData *ovsconfig.OVSPortData, portConfig *interfacestore.OVSPortConfig) *interfacestore.InterfaceConfig {
	_ = "STUB: not implemented"
	return nil
}

// createOVSInternalPort creates an OVS internal port on OVS and corresponding interface on host. Note that, host interface
// might not be available immediately after creating OVS internal port.
func (c *Controller) createOVSInternalPort(portName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Controller) createUDPTunnelPort(portName string, tunnelType ovsconfig.TunnelType, tunnelConfig *v1alpha2.UDPTunnel) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Controller) createGREPort(portName string, tunnelConfig *v1alpha2.GRETunnel) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Controller) createERSPANPort(portName string, tunnelConfig *v1alpha2.ERSPANTunnel) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Controller) getPortName(port *v1alpha2.TrafficControlPort) string {
	_ = "STUB: not implemented"
	return ""
}

// getOrCreateTrafficControlPort ensures that there is an OVS port for the given TrafficControlPort and binds the port
// to the TrafficControl. The OVS port will be created if the port doesn't exist. It returns the ofPort of the OVS port
// on success, an error if there is.
func (c *Controller) getOrCreateTrafficControlPort(port *v1alpha2.TrafficControlPort, portName, tcName string, isReturnPort bool) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Query the port binding information from portToTCBindings. If the corresponding binding information exists, indicating
// that the port has been created, then insert the TrafficControl to the set of TrafficControls using the port.

// If there is no binding information of the port in portToTCBindings, query the interface store. If corresponding
// config is found, create binding information for the port. Note that, this is used to rebuild portToTCBindings
// after restarting Antrea Agent.

// If the port is a return port, although the port is not newly created here, return flow should be installed for
// the port when it is used by a TrafficControl for the first time.

// Set the port with no-flood to reject ARP flood packets.

// If the port is a return port and is newly created, install a return flow for the port.

// Create binding for the newly created port.

// releaseTrafficControlPort releases the port from the TrafficControl and deletes the port if it is no longer used by
// any TrafficControl.
func (c *Controller) releaseTrafficControlPort(portName, tcName string, isReturnPort bool) error {
	_ = "STUB: not implemented"
	return nil
}

// If the port is no longer used by any TrafficControl, delete the port.

// Uninstall corresponding return flow if the port is a return port.

func (c *Controller) syncTrafficControl(tcName string) error { _ = "STUB: not implemented"; return nil }

// If the TrafficControl is deleted and the corresponding state doesn't exist, just return.

// If a TrafficControl is deleted but the corresponding state exists, do some cleanup for the deleted
// TrafficControl.

// Delete the state of the deleted TrafficControl.

// Get the TrafficControl state.

// If the TrafficControl exists and corresponding state doesn't exist, create state for the TrafficControl.

// Get name of the return port.

// If the name is different from the cached name in the TrafficControl state, it could be caused by the return
// port update of the TrafficControl or the creation of the TrafficControl.

// If the stale return port name cached in TrafficControl state is not empty, release the stale return port
// from the TrafficControl.

// Get or create the return port.

// Update return port name in state.

// Get name of the target port.

// If the name is different from the cached name in the TrafficControl state, it could be caused by the target port
// update of the TrafficControl or the creation of the TrafficControl.

// If the stale target port name cached in TrafficControl state is not empty, release the stale target port
// from the TrafficControl.

// Update target port name in state.

// Get or create the target port.

// Check if the mark flows should be updated.

// Get the list of Pods applying to the TrafficControl.

// If the TrafficControl is not the effective TrafficControl for the Pod, do nothing.

// If the TrafficControl is the effective TrafficControl for the Pod, insert the port to the new set in
// TrafficControl state.

// If target ofPort / direction / action in TrafficControl is updated, the mark flows should be reinstalled; if the
// new ofPort set is different from the old ofPort set, the mark flows should be also reinstalled.

// Update TrafficControl state.

// Resync the Pods applying to the TrafficControl to be deleted.

func (c *Controller) uninstallTrafficControl(tcName string, tcState *trafficControlState) error {
	_ = "STUB: not implemented"
	// Uninstall the mark flows of the TrafficControl.
	return nil
}

// Release the target port from the deleted TrafficControl.

// Release the return port from the deleted TrafficControl.

// Resync the Pods applying to the deleted TrafficControl.

func (c *Controller) podsResync(pods sets.Set[string], tcName string) {
	_ = "STUB: not implemented"
	// Resync the Pods that have new effective TrafficControl.
	return
}

// Trigger resyncing of the new effective TrafficControls of the Pods.

// bindPodToTrafficControl binds the Pod with the TrafficControl and returns whether this TrafficControl is the effective
// one for the Pod.
func (c *Controller) bindPodToTrafficControl(pod, tc string) bool {
	_ = "STUB: not implemented"
	return false
}

// Promote itself as the effective TrafficControl for the Pod if there is no binding information for the Pod.

// unbindPodFromTrafficControl unbinds the Pod with the TrafficControl. If the unbound TrafficControl was the effective
// one for the Pod and there are alternative ones, it will return the new effective TrafficControl, otherwise return empty
// string.
func (c *Controller) unbindPodFromTrafficControl(pod, tcName string) string {
	_ = "STUB: not implemented"
	return ""
}

// The binding must exist.

// Select a new effective TrafficControl.

// Remove the binding information for the Pod if there is no alternative TrafficControls.
