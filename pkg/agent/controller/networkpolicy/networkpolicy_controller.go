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

package networkpolicy

import (
	"net"
	"sync"
	"time"

	"antrea.io/ofnet/ofctrl"
	"github.com/spf13/afero"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/client"
	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/controller/networkpolicy/l7engine"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	proxytypes "antrea.io/antrea/v2/pkg/agent/proxy/types"
	"antrea.io/antrea/v2/pkg/agent/route"
	"antrea.io/antrea/v2/pkg/agent/types"
	"antrea.io/antrea/v2/pkg/apis/controlplane/install"
	"antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	"antrea.io/antrea/v2/pkg/querier"
	"antrea.io/antrea/v2/pkg/util/channel"
	utilwait "antrea.io/antrea/v2/pkg/util/wait"
)

const (
	// How long to wait before retrying the processing of a network policy change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing a rule change.
	defaultWorkers = 4
	// Default number of workers for making DNS queries.
	defaultDNSWorkers = 4
	// Reserved OVS rule ID for installing the DNS response intercept rule.
	// It is a special OVS rule which intercepts DNS query responses from DNS
	// services to the workloads that have FQDN policy rules applied.
	dnsInterceptRuleID = uint32(1)
)

const (
	dataPath           = "/var/run/antrea/networkpolicy"
	networkPoliciesDir = "network-policies"
	appliedToGroupsDir = "applied-to-groups"
	addressGroupsDir   = "address-groups"
)

type L7RuleReconciler interface {
	AddRule(ruleID, policyName string, vlanID uint32, l7Protocols []v1beta2.L7Protocol) error
	DeleteRule(ruleID string, vlanID uint32) error
}

var emptyWatch = watch.NewEmptyWatch()

var (
	scheme = runtime.NewScheme()
	codecs = serializer.NewCodecFactory(scheme)
)

func init() {
	install.Install(scheme)
}

type packetInAction func(*ofctrl.PacketIn) error

// Controller is responsible for watching Antrea AddressGroups, AppliedToGroups,
// and NetworkPolicies, feeding them to ruleCache, getting dirty rules from
// ruleCache, invoking reconcilers to reconcile them.
//
//	        a.Feed AddressGroups,AppliedToGroups
//	             and NetworkPolicies
//	|-----------|    <--------    |----------- |  c. Reconcile dirty rules |----------- |
//	| ruleCache |                 | Controller |     ------------>         | reconciler |
//	| ----------|    -------->    |----------- |                           |----------- |
//	            b. Notify dirty rules
type Controller struct {
	// antreaPolicyEnabled indicates whether Antrea NetworkPolicy and
	// ClusterNetworkPolicy are enabled.
	antreaPolicyEnabled      bool
	l7NetworkPolicyEnabled   bool
	nodeNetworkPolicyEnabled bool
	// antreaProxyEnabled indicates whether Antrea proxy is enabled.
	antreaProxyEnabled bool
	// statusManagerEnabled indicates whether a statusManager is configured.
	statusManagerEnabled bool
	// multicastEnabled indicates whether multicast is enabled.
	multicastEnabled bool
	// nodeType indicates type of the Node where Antrea Agent is running on.
	nodeType config.NodeType
	// antreaClientProvider provides interfaces to get antreaClient, which
	// can be used to watch Antrea AddressGroups, AppliedToGroups, and
	// NetworkPolicies. We need to get antreaClient dynamically because we
	// are not relying on the ClusterIP to access the Antrea Service (we
	// resolve the endpoint directly, and the endpoint can change if the
	// antrea-controller Pod is rescheduled), and because the apiserver cert
	// can be rotated and we need a new client with the updated CA cert.
	// Verifying server certificate only takes place for new requests and existing
	// watches won't be interrupted by rotating cert. The new client will be used
	// after the existing watches expire.
	antreaClientProvider client.AntreaClientProvider
	// queue maintains the NetworkPolicy ruleIDs that need to be synced.
	queue workqueue.TypedRateLimitingInterface[string]
	// ruleCache maintains the desired state of NetworkPolicy rules.
	ruleCache *ruleCache
	// podReconciler provides interfaces to reconcile the desired state of
	// NetworkPolicy rules with the actual state of Openflow entries.
	podReconciler Reconciler
	// nodeReconciler provides interfaces to reconcile the desired state of
	// NetworkPolicy rules with the actual state of iptables entries.
	nodeReconciler Reconciler
	// l7RuleReconciler provides interfaces to reconcile the desired state of
	// NetworkPolicy rules which have L7 rules with the actual state of Suricata rules.
	l7RuleReconciler L7RuleReconciler
	// l7VlanIDAllocator allocates a VLAN ID for every L7 rule.
	l7VlanIDAllocator *l7VlanIDAllocator
	// ofClient registers packetin for Antrea Policy logging.
	ofClient    openflow.Client
	auditLogger *AuditLogger
	// statusManager syncs NetworkPolicy statuses with the antrea-controller.
	// It's only for Antrea NetworkPolicies.
	statusManager         StatusManager
	fqdnController        *fqdnController
	networkPolicyWatcher  *watcher
	appliedToGroupWatcher *watcher
	addressGroupWatcher   *watcher
	fullSyncGroup         sync.WaitGroup
	ifaceStore            interfacestore.InterfaceStore
	// denyConnNotifier is used to send denied connection to the store
	denyConnNotifier channel.Notifier
	gwPort           uint32
	tunPort          uint32
	nodeConfig       *config.NodeConfig
	podNetworkWait   *utilwait.Group

	// The fileStores store runtime.Objects in files and use them as the fallback data source when agent can't connect
	// to antrea-controller on startup.
	networkPolicyStore  *fileStore
	appliedToGroupStore *fileStore
	addressGroupStore   *fileStore

	logPacketAction           packetInAction
	rejectRequestAction       packetInAction
	storeDenyConnectionAction packetInAction
}

// NewNetworkPolicyController returns a new *Controller.
func NewNetworkPolicyController(antreaClientGetter client.AntreaClientProvider,
	ofClient openflow.Client,
	routeClient route.Interface,
	ifaceStore interfacestore.InterfaceStore,
	fs afero.Fs,
	nodeName string,
	podUpdateSubscriber channel.Subscriber,
	externalEntityUpdateSubscriber channel.Subscriber,
	groupCounters []proxytypes.GroupCounter,
	groupIDUpdates <-chan string,
	antreaPolicyEnabled bool,
	l7NetworkPolicyEnabled bool,
	nodeNetworkPolicyEnabled bool,
	antreaProxyEnabled bool,
	statusManagerEnabled bool,
	multicastEnabled bool,
	loggerOptions *AuditLoggerOptions, // use nil to disable logging
	asyncRuleDeleteInterval time.Duration,
	dnsServerOverride string,
	nodeType config.NodeType,
	v4Enabled bool,
	v6Enabled bool,
	gwPort, tunPort uint32,
	nodeConfig *config.NodeConfig,
	podNetworkWait *utilwait.Group,
	l7Reconciler *l7engine.Reconciler,
	fqdnCacheMinTTL uint32) (*Controller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a WaitGroup that is used to block network policy workers from asynchronously processing
// NP rules until the events preceding bookmark are synced. It can also be used as part of the
// solution to a deterministic mechanism for when to cleanup flows from previous round.
// Wait until appliedToGroupWatcher, addressGroupWatcher and networkPolicyWatcher to receive bookmark event.

// Register packetInHandler

// Initialize logger for Antrea Policy audit logging

// Use nodeName to filter resources when watching resources.

// Storing the object to file first because its GroupVersionKind can be updated in-place during
// serialization, which may incur data race if we add it to ruleCache first.

// Storing the object to file first because its GroupVersionKind can be updated in-place during
// serialization, which may incur data race if we add it to ruleCache first.

// If any rule or the generation changes, we ensure statusManager will resync the policy's status once, in
// case the changes don't cause any actual rule update but the whole policy's generation is changed.

// When ReplaceFunc is called, either the controller restarted or this was a regular reconnection.
// For the former case, agent must resync the statuses as the controller lost the previous statuses.
// For the latter case, agent doesn't need to do anything. However, we are not able to differentiate the
// two cases. Anyway there's no harm to do a periodical resync.

// Storing the object to file first because its GroupVersionKind can be updated in-place during
// serialization, which may incur data race if we add it to ruleCache first.

// Storing the object to file first because its GroupVersionKind can be updated in-place during
// serialization, which may incur data race if we add it to ruleCache first.

// It's fine to store the object to file after applying the patch to ruleCache because the returned object
// is newly created, and ruleCache itself doesn't use it.

// Storing the object to file first because its GroupVersionKind can be updated in-place during
// serialization, which may incur data race if we add it to ruleCache first.

// Storing the object to file first because its GroupVersionKind can be updated in-place during
// serialization, which may incur data race if we add it to ruleCache first.

// It's fine to store the object to file after applying the patch to ruleCache because the returned object
// is newly created, and ruleCache itself doesn't use it.

// Storing the object to file first because its GroupVersionKind can be updated in-place during
// serialization, which may incur data race if we add it to ruleCache first.

func (c *Controller) GetFQDNCache(fqdnFilter *querier.FQDNCacheFilter) []types.DnsCacheEntry {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) GetNetworkPolicyNum() int { _ = "STUB: not implemented"; return 0 }

func (c *Controller) GetAddressGroupNum() int { _ = "STUB: not implemented"; return 0 }

func (c *Controller) GetAppliedToGroupNum() int { _ = "STUB: not implemented"; return 0 }

// GetNetworkPolicies returns the requested NetworkPolicies.
// This func will return all NetworkPolicies that can match all provided attributes in NetworkPolicyQueryFilter.
// These not provided attributes in NetworkPolicyQueryFilter means match all.
func (c *Controller) GetNetworkPolicies(npFilter *querier.NetworkPolicyQueryFilter) []v1beta2.NetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}

// GetAppliedNetworkPolicies returns the NetworkPolicies applied to the Pod and match the filter.
func (c *Controller) GetAppliedNetworkPolicies(pod, namespace string, npFilter *querier.NetworkPolicyQueryFilter) []v1beta2.NetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) GetAddressGroups() []v1beta2.AddressGroup {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) GetAppliedToGroups() []v1beta2.AppliedToGroup {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) GetNetworkPolicyByRuleFlowID(ruleFlowID uint32) *v1beta2.NetworkPolicyReference {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) GetRuleByFlowID(ruleFlowID uint32) *types.PolicyRule {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) GetControllerConnectionStatus() bool {
	_ = "STUB: not implemented"
	// When the watchers are connected, controller connection status is true. Otherwise, it is false.
	return false
}

func (c *Controller) SetDenyStoreNotifier(notifier channel.Notifier) {
	_ = "STUB: not implemented"
	return
}

// Run begins watching and processing Antrea AddressGroups, AppliedToGroups
// and NetworkPolicies, and spawns workers that reconciles NetworkPolicy rules.
// Run will not return until stopCh is closed.
func (c *Controller) Run(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"

	// If Antrea client is not ready within 5s, we assume that the Antrea Controller is not
	// available. We proceed with our watches, which are likely to fail. In turn, this will
	// trigger the fallback mechanism.
	// 5s should be more than enough if the Antrea Controller is running correctly.
	return
}

// Use NonSlidingUntil so that normal reconnection (disconnected after
// running a while) can reconnect immediately while abnormal reconnection
// won't be too aggressive.

// Batch install all rules in queue after fullSync is finished.

func (c *Controller) matchIGMPType(r *rule, igmpType uint8, groupAddress string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetIGMPNPRuleInfo looks up the IGMP NetworkPolicy rule that matches the given Pod and groupAddress,
// and returns the rule information if found.
func (c *Controller) GetIGMPNPRuleInfo(podName, podNamespace string, groupAddress net.IP, igmpType uint8) (*types.IGMPNPRuleInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Controller) enqueueRule(ruleID string) { _ = "STUB: not implemented"; return }

// worker runs a worker thread that just dequeues items, processes them, and
// marks them done. You may run as many of these in parallel as you wish; the
// workqueue guarantees that they will not end up processing the same rule at
// the same time.
func (c *Controller) worker() { _ = "STUB: not implemented"; return }

func (c *Controller) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// processAllItemsInQueue pops all rule keys queued at the moment and calls syncRules to
// reconcile those rules in batch.
func (c *Controller) processAllItemsInQueue() { _ = "STUB: not implemented"; return }

// set key to done to prevent missing watched updates between here and fullSync finish.

// Reconcile all rule keys at once.

func (c *Controller) syncRule(key string) error { _ = "STUB: not implemented"; return nil }

// Uncertain whether this rule applies to a Node or Pod, but it's safe to delete it redundantly.

// We don't know whether this is a rule owned by Antrea Policy, but
// harmless to delete it.

// If the rule is not realizable, we can simply skip it as it will be marked as dirty
// and queued again when we receive the missing group it missed.

// Allocate VLAN ID for the L7 rule.

// No matter whether the rule reconciliation succeeds or not, fqdnController
// needs to be notified of the status.

// syncRules calls the reconciler to sync all the rules after watchers complete full sync.
// After flows for those init events are installed, subsequent rules will be handled asynchronously
// by the syncRule() function.
func (c *Controller) syncRules(keys []string) error { _ = "STUB: not implemented"; return nil }

// It's normal that a rule is not effective on this Node but abnormal that it is not realizable after watchers
// complete full sync.

// Allocate VLAN ID for the L7 rule.

func (c *Controller) handleErr(err error, key string) { _ = "STUB: not implemented"; return }

// watcher is responsible for watching a given resource with the provided watchFunc
// and calling the eventHandlers when receiving events.
type watcher struct {
	// objectType is the type of objects being watched, used for logging.
	objectType string
	// watchFunc is the function that starts the watch.
	watchFunc func() (watch.Interface, error)
	// AddFunc is the function that handles added event.
	AddFunc func(obj runtime.Object) error
	// UpdateFunc is the function that handles modified event.
	UpdateFunc func(obj runtime.Object) error
	// DeleteFunc is the function that handles deleted event.
	DeleteFunc func(obj runtime.Object) error
	// ReplaceFunc is the function that handles init events.
	ReplaceFunc func(objs []runtime.Object) error
	// FallbackFunc is the function that provides the data when it can't start the watch successfully.
	FallbackFunc func() ([]runtime.Object, error)
	// connected represents whether the watch has connected to apiserver successfully.
	connected bool
	// lock protects connected.
	lock sync.RWMutex
	// group to be notified when each watcher receives bookmark event
	fullSyncWaitGroup *sync.WaitGroup
	// fullSynced indicates if the resource has been synced at least once since agent started.
	fullSynced bool
}

func (w *watcher) isConnected() bool { _ = "STUB: not implemented"; return false }

func (w *watcher) setConnected(connected bool) { _ = "STUB: not implemented"; return }

// fallback gets init events from the FallbackFunc if the watcher hasn't been synced once.
func (w *watcher) fallback() {
	_ = "STUB: not implemented"
	// If the watcher has been synced once, the fallback data source doesn't have newer data, do nothing.
	return
}

func (w *watcher) onFullSync() { _ = "STUB: not implemented"; return }

// Notify fullSyncWaitGroup that all events before bookmark is handled

func (w *watcher) watch() { _ = "STUB: not implemented"; return }

// Watch method doesn't return error but "emptyWatch" in case of some partial data errors,
// e.g. timeout error. Make sure that watcher is not empty and log error otherwise.

// First receive init events from the result channel and buffer them until
// a Bookmark event is received, indicating that all init events have been
// received.
