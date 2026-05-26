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

package externalnode

import (
	"time"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/agent/util"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	enlister "antrea.io/antrea/v2/pkg/client/listers/crd/v1alpha1"
	agentConfig "antrea.io/antrea/v2/pkg/config/agent"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
	"antrea.io/antrea/v2/pkg/util/channel"
)

const (
	controllerName = "ExternalNodeController"
	// How long to wait before retrying the processing of an ExternalNode change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Disable resyncing.
	resyncPeriod time.Duration = 0

	ovsExternalIDUplinkName      = "uplink-name"
	ovsExternalIDUplinkPort      = "uplink-port"
	ovsExternalIDEntityName      = "entity-name"
	ovsExternalIDEntityNamespace = "entity-namespace"
	ovsExternalIDIPs             = "ip-address"
	ipsSplitter                  = ","
)

var (
	keyFunc              = cache.MetaNamespaceKeyFunc
	splitKeyFunc         = cache.SplitMetaNamespaceKey
	renameInterface      = util.RenameInterface
	getInterfaceConfig   = util.GetInterfaceConfig
	getIPNetDeviceFromIP = util.GetIPNetDeviceFromIP
	hostInterfaceExists  = util.HostInterfaceExists
)

type ExternalNodeController struct {
	ovsBridgeClient          ovsconfig.OVSBridgeClient
	ovsctlClient             ovsctl.OVSCtlClient
	ofClient                 openflow.Client
	externalNodeInformer     cache.SharedIndexInformer
	externalNodeLister       enlister.ExternalNodeLister
	externalNodeListerSynced cache.InformerSynced
	queue                    workqueue.TypedRateLimitingInterface[string]
	ifaceStore               interfacestore.InterfaceStore
	syncedExternalNode       *v1alpha1.ExternalNode
	// externalEntityUpdateNotifier is used for notifying ExternalEntity updates to NetworkPolicyController.
	externalEntityUpdateNotifier channel.Notifier
	nodeName                     string
	externalNodeNamespace        string
	policyBypassRules            []agentConfig.PolicyBypassRule
}

func NewExternalNodeController(ovsBridgeClient ovsconfig.OVSBridgeClient, ofClient openflow.Client, externalNodeInformer cache.SharedIndexInformer,
	ifaceStore interfacestore.InterfaceStore, externalEntityUpdateNotifier channel.Notifier, externalNodeNamespace string, policyBypassRules []agentConfig.PolicyBypassRule) (*ExternalNodeController, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run will create a worker (goroutine) which will process the ExternalNode events from the work queue.
func (c *ExternalNodeController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *ExternalNodeController) enqueueExternalNodeAdd(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *ExternalNodeController) enqueueExternalNodeUpdate(oldObj interface{}, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *ExternalNodeController) enqueueExternalNodeDelete(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *ExternalNodeController) reconcile() error { _ = "STUB: not implemented"; return nil }

func (c *ExternalNodeController) reconcileHostUplinkFlows() error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExternalNodeController) reconcilePolicyBypassFlows() error {
	_ = "STUB: not implemented"
	return nil
}

// worker is a long-running function that will continuously call the processNextWorkItem function in
// order to read and process a message on the work queue.
func (c *ExternalNodeController) worker() { _ = "STUB: not implemented"; return }

func (c *ExternalNodeController) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// If no error occurs, then forget this item so it does not get queued again until
// another change happens.

// Put the item back on the work queue to handle any transient errors.

func (c *ExternalNodeController) syncExternalNode(key string) error {
	_ = "STUB: not implemented"
	return nil
}

// This err should not occur.

func (c *ExternalNodeController) addExternalNode(en *v1alpha1.ExternalNode) error {
	_ = "STUB: not implemented"
	return nil
}

// Notify the ExternalEntity event to NetworkPolicyController.

func (c *ExternalNodeController) addInterface(ifName string, eeNamespace string, eeName string, ips []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExternalNodeController) updateExternalNode(preEN *v1alpha1.ExternalNode, curEN *v1alpha1.ExternalNode) error {
	_ = "STUB: not implemented"
	return nil
}

// Notify the ExternalEntity event to NetworkPolicyController.

func (c *ExternalNodeController) deleteExternalNode() error { _ = "STUB: not implemented"; return nil }

// Remove any stale configuration that is related to the deleted ExternalNode
// and terminate the process if required.

func (c *ExternalNodeController) deleteInterfaces() error { _ = "STUB: not implemented"; return nil }

func (c *ExternalNodeController) deleteInterface(interfaceConfig *interfacestore.InterfaceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExternalNodeController) createOVSPortsAndFlows(uplinkName, hostIFName, eeNamespace, eeName string, ips []string) (*interfacestore.InterfaceConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create uplink port in OVS.

// Create host port in OVS.

// Move configurations from the uplink to host port

func GetOVSAttachInfo(uplinkName, uplinkUUID, entityName, entityNamespace string, ips []string) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExternalNodeController) updateOVSPortsData(interfaceConfig *interfacestore.InterfaceConfig, portData *ovsconfig.OVSPortData, eeName string, ips []string) (*interfacestore.InterfaceConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ExternalNodeController) removeOVSPortsAndFlows(interfaceConfig *interfacestore.InterfaceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// This is for issue #5111 (https://github.com/antrea-io/antrea/issues/5111), which may happen if an error occurs
// when moving the configuration back from host internal interface to uplink. This logic is run in the second
// try after the error is returned, at this time the host internal interface is already deleted, and the uplink's
// name is recovered. So the ips and routes in "adapterConfig" are actually read from the uplink and no need to
// move the configurations back. The issue was seen on VM with RHEL 8.4 on azure cloud.

// Delete host interface from OVS datapath if it exists.
// This is to resolve an issue that OVS fails to remove the interface from datapath. It might happen because the interface
// is busy when OVS tries to remove it with the OVSDB interface deletion event.

// Wait until the host interface created by OVS is removed.

// Recover the uplink interface's name.

// Move the IP configurations back to the host interface.

func getHostInterfaceName(iface v1alpha1.NetworkInterface) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func ParseHostInterfaceConfig(ovsBridgeClient ovsconfig.OVSBridgeClient, portData *ovsconfig.OVSPortData, portConfig *interfacestore.OVSPortConfig) (*interfacestore.InterfaceConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseProtocol(protocol string) binding.Protocol {
	_ = "STUB: not implemented"
	return *new(binding.Protocol)
}
