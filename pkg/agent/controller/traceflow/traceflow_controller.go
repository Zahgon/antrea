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

package traceflow

import (
	"net"
	"net/netip"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/types"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	clientsetversioned "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
	"antrea.io/antrea/v2/pkg/querier"
)

const (
	controllerName = "AntreaAgentTraceflowController"
	// Set resyncPeriod to 0 to disable resyncing.
	resyncPeriod time.Duration = 0
	// How long to wait before retrying the processing of a traceflow.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing traceflow request.
	defaultWorkers = 4
	// Delay in milliseconds before injecting packet into OVS. The time of different nodes may not be completely
	// synchronized, which requires a delay before inject packet.
	injectPacketDelay      = 2000
	injectLocalPacketDelay = 100

	// ICMP Echo Request type and code.
	icmpEchoRequestType   uint8 = 8
	icmpv6EchoRequestType uint8 = 128
	icmpEchoRequestCode   uint8 = 0

	defaultTTL uint8 = 64
)

type traceflowState struct {
	name string
	// Used to uniquely identify Traceflow.
	uid         types.UID
	tag         int8
	liveTraffic bool
	droppedOnly bool
	// Live-traffic Traceflow with only destination Pod specified.
	receiverOnly bool
	isSender     bool
	// Agent received the first Traceflow packet from OVS.
	receivedPacket bool
}

// Controller is responsible for setting up Openflow entries and injecting traceflow packet into
// the switch for traceflow request.
type Controller struct {
	kubeClient             clientset.Interface
	serviceLister          corelisters.ServiceLister
	serviceListerSynced    cache.InformerSynced
	crdClient              clientsetversioned.Interface
	traceflowInformer      crdinformers.TraceflowInformer
	traceflowLister        crdlisters.TraceflowLister
	traceflowListerSynced  cache.InformerSynced
	ofClient               openflow.Client
	networkPolicyQuerier   querier.AgentNetworkPolicyInfoQuerier
	egressQuerier          querier.EgressQuerier
	podSubnetChecker       PodSubnetChecker
	interfaceStore         interfacestore.InterfaceStore
	networkConfig          *config.NetworkConfig
	nodeConfig             *config.NodeConfig
	serviceCIDR            *net.IPNet   // K8s Service ClusterIP CIDR
	podCIDRs               []*net.IPNet // Only used in networkPolicyOnly mode
	queue                  workqueue.TypedRateLimitingInterface[string]
	runningTraceflowsMutex sync.RWMutex
	// runningTraceflows is a map for storing the running Traceflow state
	// with dataplane tag to be the key.
	runningTraceflows map[int8]*traceflowState
	enableAntreaProxy bool
}

// NewTraceflowController instantiates a new Controller object which will process Traceflow
// events.
func NewTraceflowController(
	kubeClient clientset.Interface,
	crdClient clientsetversioned.Interface,
	serviceInformer coreinformers.ServiceInformer,
	traceflowInformer crdinformers.TraceflowInformer,
	client openflow.Client,
	npQuerier querier.AgentNetworkPolicyInfoQuerier,
	egressQuerier querier.EgressQuerier,
	podSubnetChecker PodSubnetChecker,
	interfaceStore interfacestore.InterfaceStore,
	networkConfig *config.NetworkConfig,
	nodeConfig *config.NodeConfig,
	serviceCIDR *net.IPNet,
	podCIDRs []*net.IPNet,
	enableAntreaProxy bool) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Add handlers for Traceflow events.

// Register packetInHandler

// Add serviceLister if AntreaProxy enabled

// enqueueTraceflow adds an object to the controller work queue.
func (c *Controller) enqueueTraceflow(tf *crdv1beta1.Traceflow) { _ = "STUB: not implemented"; return }

// Run will create defaultWorkers workers (go routines) which will process the Traceflow events from the
// workqueue.
func (c *Controller) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *Controller) addTraceflow(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updateTraceflow(_, curObj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) deleteTraceflow(old interface{}) { _ = "STUB: not implemented"; return }

// worker is a long-running function that will continually call the processTraceflowItem function
// in order to read and process a message on the workqueue.
func (c *Controller) worker() { _ = "STUB: not implemented"; return }

// processTraceflowItem processes an item in the "traceflow" work queue, by calling syncTraceflow
// after casting the item to a string (Traceflow name). If syncTraceflow returns an error, this
// function logs error. If syncTraceflow is successful, the Traceflow is removed from the queue
// until we get notified of a new change. This function returns false if and only if the work queue
// was shutdown (no more items will be processed).
func (c *Controller) processTraceflowItem() bool { _ = "STUB: not implemented"; return false }

// We call Done here so the workqueue knows we have finished processing this item. We also
// must remember to call Forget if we do not want this work item being re-queued. For
// example, we do not call Forget if a transient error occurs, instead the item is put back
// on the workqueue and attempted again after a back-off period.

// If no error occurs we Forget this item so it does not get queued again.

// If error occurs we log error.

// TODO: Let controller compute which Node is the sender, and each Node watch the TF CRD with some
// filter to get and process only TF from the Node.
//
// syncTraceflow gets Traceflow CRD by name, update cache and start syncing.
func (c *Controller) syncTraceflow(traceflowName string) error {
	_ = "STUB: not implemented"
	return nil
}

// This may happen if a Traceflow is assigned with a tag that was just released from an old Traceflow but
// the agent hasn't processed the deletion event of the old Traceflow yet.

// startTraceflow deploys OVS flow entries for Traceflow and inject packet if current Node
// is Sender Node.
func (c *Controller) startTraceflow(tf *crdv1beta1.Traceflow) error {
	_ = "STUB: not implemented"
	return nil
}

// Live-traffic Traceflow with only the Destination Pod specified.

// TODO: let controller compute the sender/receiver Node, and the sender
// /receiver Node can just return an error, if fails to find the Pod.

// On the sender or receiver (the receiverOnly case) Node, trace
// the first packet of the first connection that matches the
// Traceflow spec.

// Store Traceflow to cache.

// Install flow entries for traceflow.

// Skip packet injection if the source Pod is not found on the local Node.

// If the destination is Service/IP or the packet will
// be sent to remote Node, wait a small period for other
// Nodes.

// Issue #2116
// Wait a small period after flows installed to avoid unexpected behavior.

func (c *Controller) validateTraceflow(tf *crdv1beta1.Traceflow) error {
	_ = "STUB: not implemented"
	return nil
}

// When AntreaProxy is enabled, serviceCIDR is not required and may be set to a
// default value which does not match the cluster configuration.

func (c *Controller) preparePacket(tf *crdv1beta1.Traceflow, intf *interfacestore.InterfaceConfig, receiverOnly bool) (*binding.Packet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The packet will be matched with the Pod MAC.

// DestinationMAC is nil here, will be set to gateway
// MAC in ofClient.SendTraceflowPacket()

// IP Protocol 0 (IPv6 Hop-by-Hop Option) is not supported by
// Traceflow. If NextHeader is not provided, protocol ICMPv6
// will be used as the default.

// TCP > UDP > ICMP > other IP protocol.

// Defaults to ICMP if not live-traffic Traceflow.

func (c *Controller) errorTraceflowCRD(tf *crdv1beta1.Traceflow, reason string) (*crdv1beta1.Traceflow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete Traceflow state and OVS flows.
func (c *Controller) cleanupTraceflow(tfName string) { _ = "STUB: not implemented"; return }

// This must be executed before deleting the tag from runningTraceflows, otherwise it may uninstall another
// Traceflow's flows if the tag is reassigned.

type PodSubnetChecker interface {
	// LookupIPInPodSubnets returns two boolean values. The first one indicates whether the IP can be
	// found in a PodCIDR for one of the cluster Nodes. The second one indicates whether the IP is used
	// as a gateway IP. The second boolean value can only be true if the first one is true.
	LookupIPInPodSubnets(ip netip.Addr) (isFound bool, isGWIP bool)
}
