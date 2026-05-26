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

// Package networkpolicy provides NetworkPolicyController implementation to manage
// and synchronize the GroupMembers and Namespaces affected by Network Policies and enforce
// their rules.
package networkpolicy

import (
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	v1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	coreinformers "k8s.io/client-go/informers/core/v1"
	networkinginformers "k8s.io/client-go/informers/networking/v1"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	networkinglisters "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	policyinformers "sigs.k8s.io/network-policy-api/pkg/client/informers/externalversions/apis/v1alpha1"
	policylisters "sigs.k8s.io/network-policy-api/pkg/client/listers/apis/v1alpha1"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha2"
	secv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/apiserver/storage"
	"antrea.io/antrea/v2/pkg/client/clientset/versioned"
	crdv1b1informers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
	crdv1b1listers "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/controller/grouping"
	"antrea.io/antrea/v2/pkg/controller/labelidentity"
	antreatypes "antrea.io/antrea/v2/pkg/controller/types"
)

const (
	controllerName = "NetworkPolicyController"
	// NetworkPolicyController is the only writer of the antrea network policy
	// storages and will keep re-enqueuing failed items until they succeed.
	// Set resyncPeriod to 0 to disable resyncing.
	resyncPeriod time.Duration = 0
	// How long to wait before retrying the processing of a NetworkPolicy change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing a NetworkPolicy change.
	defaultWorkers = 4
	// Default rule priority for K8s NetworkPolicy rules.
	defaultRulePriority = -1
	// TierIndex is used to index ClusterNetworkPolicies by Tier names.
	TierIndex = "tier"
	// PriorityIndex is used to index Tiers by their priorities.
	PriorityIndex = "priority"
	// ClusterGroupIndex is used to index ClusterNetworkPolicies by ClusterGroup names.
	ClusterGroupIndex = "clustergroup"
	// GroupIndex is used to index Antrea NetworkPolicies by Group names.
	GroupIndex = "group"

	// EnableNPLoggingAnnotationKey can be added to Namespace to enable logging K8s NP.
	EnableNPLoggingAnnotationKey = "networkpolicy.antrea.io/enable-logging"

	appliedToGroupType grouping.GroupType = "appliedToGroup"
	addressGroupType   grouping.GroupType = "addressGroup"
	internalGroupType  grouping.GroupType = "internalGroup"

	perNamespaceRuleIndex      = "hasPerNamespaceRule"
	namespaceRuleLabelKeyIndex = "namespaceRuleLabelKeys"
	indexValueTrue             = "true"
)

var (
	// uuidNamespace is a uuid.UUID type generated from a string to be
	// used to generate uuid.UUID for internal Antrea objects like
	// AppliedToGroup, AddressGroup etc.
	// e4f24a48-ca1f-4d5b-819c-ea7632b22115 was generated using
	// uuid.NewRandom() function.
	uuidNamespace = uuid.Must(uuid.Parse("e4f24a48-ca1f-4d5b-819c-ea7632b22115"))

	// matchAllPeer is a NetworkPolicyPeer matching all source/destination IP addresses. Both IPv4 Any (0.0.0.0/0) and
	// IPv6 Any (::/0) are added into the IPBlocks, and Antrea Agent should decide if both two are used according the
	// supported IP protocols configured in the cluster.
	matchAllPeer = controlplane.NetworkPolicyPeer{
		IPBlocks: []controlplane.IPBlock{
			{CIDR: controlplane.IPNet{IP: controlplane.IPAddress(net.IPv4zero), PrefixLength: 0}},
			{CIDR: controlplane.IPNet{IP: controlplane.IPAddress(net.IPv6zero), PrefixLength: 0}},
		},
	}
	// matchAllPodsPeer is a networkingv1.NetworkPolicyPeer matching all Pods from all Namespaces.
	matchAllPodsPeer = networkingv1.NetworkPolicyPeer{
		NamespaceSelector: &metav1.LabelSelector{},
	}
	// defaultAction is a RuleAction which sets the default Action for the NetworkPolicy rule.
	defaultAction = secv1beta1.RuleActionAllow
)

func getKNPReference(knp *networkingv1.NetworkPolicy) *controlplane.NetworkPolicyReference {
	_ = "STUB: not implemented"
	return nil
}

// NetworkPolicyController is responsible for synchronizing the Namespaces and Pods
// affected by a Network Policy.
type NetworkPolicyController struct {
	// kubeClient is a standard Kubernetes clientset.
	kubeClient clientset.Interface
	// crdClient is the clientset for CRD API group.
	crdClient versioned.Interface

	namespaceInformer coreinformers.NamespaceInformer
	// namespaceLister is able to list/get Namespaces and is populated by the shared informer passed to
	// NewNetworkPolicyController.
	namespaceLister corelisters.NamespaceLister
	// namespaceListerSynced is a function which returns true if the Namespace shared informer has been synced at least once.
	namespaceListerSynced cache.InformerSynced

	serviceInformer coreinformers.ServiceInformer
	// serviceLister is able to list/get Services and is populated by the shared informer passed to
	// NewNetworkPolicyController.
	serviceLister corelisters.ServiceLister
	// serviceListerSynced is a function which returns true if the Service shared informer has been synced at least once.
	serviceListerSynced cache.InformerSynced

	networkPolicyInformer networkinginformers.NetworkPolicyInformer
	// networkPolicyLister is able to list/get Network Policies and is populated by the shared informer passed to
	// NewNetworkPolicyController.
	networkPolicyLister networkinglisters.NetworkPolicyLister
	// networkPolicyListerSynced is a function which returns true if the Network Policy shared informer has been synced at least once.
	networkPolicyListerSynced cache.InformerSynced

	acnpInformer crdv1b1informers.ClusterNetworkPolicyInformer
	// acnpLister is able to list/get AntreaClusterNetworkPolicies and is populated by the shared informer passed to
	// NewClusterNetworkPolicyController.
	acnpLister crdv1b1listers.ClusterNetworkPolicyLister
	// acnpListerSynced is a function which returns true if the AntreaClusterNetworkPolicies shared informer has been synced at least once.
	acnpListerSynced cache.InformerSynced

	annpInformer crdv1b1informers.NetworkPolicyInformer
	// annpLister is able to list/get AntreaNetworkPolicies and is populated by the shared informer passed to
	// NewNetworkPolicyController.
	annpLister crdv1b1listers.NetworkPolicyLister
	// annpListerSynced is a function which returns true if the AntreaNetworkPolicies shared informer has been synced at least once.
	annpListerSynced cache.InformerSynced

	tierInformer crdv1b1informers.TierInformer
	// tierLister is able to list/get Tiers and is populated by the shared informer passed to
	// NewNetworkPolicyController.
	tierLister crdv1b1listers.TierLister
	// tierListerSynced is a function which returns true if the Tiers shared informer has been synced at least once.
	tierListerSynced cache.InformerSynced

	cgInformer crdv1b1informers.ClusterGroupInformer
	// cgLister is able to list/get ClusterGroups and is populated by the shared informer passed to
	// NewClusterGroupController.
	cgLister crdv1b1listers.ClusterGroupLister
	// cgListerSynced is a function which returns true if the ClusterGroup shared informer has been synced at least
	// once.
	cgListerSynced cache.InformerSynced

	nodeInformer coreinformers.NodeInformer
	// nodeLister is able to list/get Nodes and is populated by the shared informer passed to
	// NewNetworkPolicyController.
	nodeLister corelisters.NodeLister
	// nodeListerSynced is a function which returns true if the Node shared informer has been synced at least once.
	nodeListerSynced cache.InformerSynced

	grpInformer crdv1b1informers.GroupInformer
	// grpLister is able to list/get Groups and is populated by the shared informer passed to
	// NewGroupController.
	grpLister crdv1b1listers.GroupLister
	// grpListerSynced is a function which returns true if the Group shared informer has been synced at least
	// once.
	grpListerSynced cache.InformerSynced

	adminNetworkPolicyInformer policyinformers.AdminNetworkPolicyInformer
	// adminNetworkPolicyLister is able to list/get AdminNetworkPolicy objects.
	adminNetworkPolicyLister policylisters.AdminNetworkPolicyLister
	// AdminNetworkPolicySynced is a function which returns true if the AdminNetworkPolicy shared informer has
	// been synced at least once.
	adminNetworkPolicyListerSynced cache.InformerSynced

	banpInformer policyinformers.BaselineAdminNetworkPolicyInformer
	// banpLister is able to list/get BaselineAdminNetworkPolicy objects.
	banpLister policylisters.BaselineAdminNetworkPolicyLister
	// banpListerSynced is a function which returns true if the BaselineAdminNetworkPolicy shared informer has
	// been synced at least once.
	banpListerSynced cache.InformerSynced

	// addressGroupStore is the storage where the populated Address Groups are stored.
	addressGroupStore storage.Interface
	// appliedToGroupStore is the storage where the populated AppliedTo Groups are stored.
	appliedToGroupStore storage.Interface
	// internalNetworkPolicyStore is the storage where the populated internal Network Policy are stored.
	internalNetworkPolicyStore storage.Interface
	// internalGroupStore is a simple store which maintains the internal Group types which can be later
	// converted to AppliedToGroup or AddressGroup based on usage.
	internalGroupStore storage.Interface

	// appliedToGroupQueue maintains the networkpolicy.AppliedToGroup objects that
	// need to be synced.
	appliedToGroupQueue workqueue.TypedRateLimitingInterface[string]
	// addressGroupQueue maintains the networkpolicy.AddressGroup objects that
	// need to be synced.
	addressGroupQueue workqueue.TypedRateLimitingInterface[string]
	// internalNetworkPolicyQueue maintains the networkpolicy.NetworkPolicy objects that
	// need to be synced.
	internalNetworkPolicyQueue workqueue.TypedRateLimitingInterface[controlplane.NetworkPolicyReference]
	// internalGroupQueue maintains the networkpolicy.Group objects that needs to be
	// synced.
	internalGroupQueue workqueue.TypedRateLimitingInterface[string]

	// internalNetworkPolicyMutex prevents concurrent processing of internal networkpolicies who refer
	// to the same addressgroups/appliedtogroups.
	internalNetworkPolicyMutex sync.RWMutex

	// appliedToGroupNotifier is responsible for notifying subscribers of an AppliedToGroup about its update.
	// The typical subscribers of AppliedToGroup are NetworkPolicies.
	appliedToGroupNotifier *notifier

	groupingInterface grouping.Interface
	// Added as a member to the struct to allow injection for testing.
	groupingInterfaceSynced func() bool

	labelIdentityInterface labelidentity.Interface
	// Enable Stretched Networkpolicy feature which allows Antrea-native policies to select peer
	// from other clusters in a ClusterSet.
	stretchNPEnabled bool
	// heartbeatCh is an internal channel for testing. It's used to know whether all tasks have been
	// processed, and to count executions of each function.
	heartbeatCh chan heartbeat
}

type heartbeat struct {
	name      string
	timestamp time.Time
}

var tierIndexers = cache.Indexers{
	PriorityIndex: func(obj interface{}) ([]string, error) {
		tr, ok := obj.(*secv1beta1.Tier)
		if !ok {
			return []string{}, nil
		}
		return []string{strconv.FormatInt(int64(tr.Spec.Priority), 10)}, nil
	},
}

var acnpIndexers = cache.Indexers{
	TierIndex: func(obj interface{}) ([]string, error) {
		acnp, ok := obj.(*secv1beta1.ClusterNetworkPolicy)
		if !ok {
			return []string{}, nil
		}
		return []string{acnp.Spec.Tier}, nil
	},
	ClusterGroupIndex: func(obj interface{}) ([]string, error) {
		acnp, ok := obj.(*secv1beta1.ClusterNetworkPolicy)
		if !ok {
			return []string{}, nil
		}
		groupNames := sets.Set[string]{}
		for _, appTo := range acnp.Spec.AppliedTo {
			if appTo.Group != "" {
				groupNames.Insert(appTo.Group)
			}
		}
		if len(acnp.Spec.Ingress) == 0 && len(acnp.Spec.Egress) == 0 {
			return sets.List(groupNames), nil
		}
		appendGroups := func(rule secv1beta1.Rule) {
			for _, peer := range rule.To {
				if peer.Group != "" {
					groupNames.Insert(peer.Group)
				}
			}
			for _, peer := range rule.From {
				if peer.Group != "" {
					groupNames.Insert(peer.Group)
				}
			}
			for _, appTo := range rule.AppliedTo {
				if appTo.Group != "" {
					groupNames.Insert(appTo.Group)
				}
			}
		}
		for _, rule := range acnp.Spec.Egress {
			appendGroups(rule)
		}
		for _, rule := range acnp.Spec.Ingress {
			appendGroups(rule)
		}
		return sets.List(groupNames), nil
	},
	perNamespaceRuleIndex: func(obj interface{}) ([]string, error) {
		acnp, ok := obj.(*secv1beta1.ClusterNetworkPolicy)
		if !ok {
			return []string{}, nil
		}
		if hasPerNamespaceRule(acnp) {
			return []string{indexValueTrue}, nil
		}
		return []string{}, nil
	},
	namespaceRuleLabelKeyIndex: func(obj interface{}) ([]string, error) {
		cnp, ok := obj.(*secv1beta1.ClusterNetworkPolicy)
		if !ok {
			return []string{}, nil
		}
		return namespaceRuleLabelKeys(cnp).UnsortedList(), nil
	},
}

var annpIndexers = cache.Indexers{
	TierIndex: func(obj interface{}) ([]string, error) {
		annp, ok := obj.(*secv1beta1.NetworkPolicy)
		if !ok {
			return []string{}, nil
		}
		return []string{annp.Spec.Tier}, nil
	},
	GroupIndex: func(obj interface{}) ([]string, error) {
		annp, ok := obj.(*secv1beta1.NetworkPolicy)
		if !ok {
			return []string{}, nil
		}
		ns := annp.Namespace + "/"
		groupNames := sets.Set[string]{}
		for _, appTo := range annp.Spec.AppliedTo {
			if appTo.Group != "" {
				groupNames.Insert(ns + appTo.Group)
			}
		}
		if len(annp.Spec.Ingress) == 0 && len(annp.Spec.Egress) == 0 {
			return sets.List(groupNames), nil
		}
		appendGroups := func(rule secv1beta1.Rule) {
			for _, peer := range rule.To {
				if peer.Group != "" {
					groupNames.Insert(ns + peer.Group)
				}
			}
			for _, peer := range rule.From {
				if peer.Group != "" {
					groupNames.Insert(ns + peer.Group)
				}
			}
			for _, appTo := range rule.AppliedTo {
				if appTo.Group != "" {
					groupNames.Insert(ns + appTo.Group)
				}
			}
		}
		for _, rule := range annp.Spec.Egress {
			appendGroups(rule)
		}
		for _, rule := range annp.Spec.Ingress {
			appendGroups(rule)
		}
		return sets.List(groupNames), nil
	},
}

// NewNetworkPolicyController returns a new *NetworkPolicyController.
func NewNetworkPolicyController(kubeClient clientset.Interface,
	crdClient versioned.Interface,
	groupingInterface grouping.Interface,
	labelIdentityInterface labelidentity.Interface,
	namespaceInformer coreinformers.NamespaceInformer,
	serviceInformer coreinformers.ServiceInformer,
	networkPolicyInformer networkinginformers.NetworkPolicyInformer,
	nodeInformer coreinformers.NodeInformer,
	acnpInformer crdv1b1informers.ClusterNetworkPolicyInformer,
	annpInformer crdv1b1informers.NetworkPolicyInformer,
	adminNPInformer policyinformers.AdminNetworkPolicyInformer,
	banpInformer policyinformers.BaselineAdminNetworkPolicyInformer,
	tierInformer crdv1b1informers.TierInformer,
	cgInformer crdv1b1informers.ClusterGroupInformer,
	grpInformer crdv1b1informers.GroupInformer,
	addressGroupStore storage.Interface,
	appliedToGroupStore storage.Interface,
	internalNetworkPolicyStore storage.Interface,
	internalGroupStore storage.Interface,
	stretchedNPEnabled bool) *NetworkPolicyController {
	_ = "STUB: not implemented"
	return nil
}

// Add handlers for NetworkPolicy events.

// Register Informer and add handlers for AntreaPolicy events only if the feature is enabled.

// Add handlers for Namespace events.

// Add handlers for Node events.

// Add event handlers for ClusterGroup notification.

// Add event handlers for Group notification.

func (n *NetworkPolicyController) heartbeat(name string) { _ = "STUB: not implemented"; return }

func (n *NetworkPolicyController) GetNetworkPolicyNum() int { _ = "STUB: not implemented"; return 0 }

func (n *NetworkPolicyController) GetAddressGroupNum() int { _ = "STUB: not implemented"; return 0 }

func (n *NetworkPolicyController) GetAppliedToGroupNum() int { _ = "STUB: not implemented"; return 0 }

// GetConnectedAgentNum gets the number of Agents which are connected to this Controller.
// Since Agent will watch all the three stores (internalNetworkPolicyStore, appliedToGroupStore, addressGroupStore),
// the number of watchers of one of these three stores is equal to the number of connected Agents.
// Here, we uses the number of watchers of appliedToGroupStore to represent the number of connected Agents as
// internalNetworkPolicyStore is also watched by the StatusController of the process itself.
func (n *NetworkPolicyController) GetConnectedAgentNum() int { _ = "STUB: not implemented"; return 0 }

// getNormalizedUID generates a unique UUID based on a given string.
// For example, it can be used to generate keys using normalized selectors
// unique within the Namespace by adding the constant UID.
func getNormalizedUID(name string) string { _ = "STUB: not implemented"; return "" }

// createAppliedToGroup creates an AppliedToGroup object corresponding to the provided selectors.
func (n *NetworkPolicyController) createAppliedToGroup(npNsName string, pSel, nSel, eSel, nodeSel *metav1.LabelSelector) *antreatypes.AppliedToGroup {
	_ = "STUB: not implemented"
	return nil
}

// Construct a new AppliedToGroup.

// createAddressGroup creates an AddressGroup object corresponding to a
// NetworkPolicyPeer object in NetworkPolicyRule. This function simply
// creates the object without actually populating the PodAddresses as the
// affected GroupMembers are calculated during sync process.
func (n *NetworkPolicyController) createAddressGroup(namespace string, podSelector, nsSelector, eeSelector, nodeSelector *metav1.LabelSelector) *antreatypes.AddressGroup {
	_ = "STUB: not implemented"
	return nil
}

// Create an AddressGroup object per Peer object.

// toAntreaProtocol converts a v1.Protocol object to an Antrea Protocol object.
func toAntreaProtocol(npProtocol *v1.Protocol) *controlplane.Protocol {
	_ = "STUB: not implemented"
	// If Protocol is unset, it must default to TCP protocol.
	return nil
}

// toAntreaServices converts a slice of networkingv1.NetworkPolicyPort objects
// to a slice of Antrea Service objects. A bool is returned along with the
// Service objects to indicate whether any named port exists.
func toAntreaServices(npPorts []networkingv1.NetworkPolicyPort) ([]controlplane.Service, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// toAntreaIPBlock converts a networkingv1.IPBlock to an Antrea IPBlock.
func toAntreaIPBlock(ipBlock *networkingv1.IPBlock) (*controlplane.IPBlock, error) {
	_ = "STUB: not implemented"
	// Convert the allowed IPBlock to networkpolicy.IPNet.
	return nil, nil
}

// Convert the except IPBlock to networkpolicy.IPNet.

// processNetworkPolicy creates an internal NetworkPolicy instance corresponding
// to the networkingv1.NetworkPolicy object. This method does not commit the
// internal NetworkPolicy in store, instead returns an instance to the caller
// wherein, it will be either stored as a new Object in case of ADD event or
// modified and store the updated instance, in case of an UPDATE event.
func (n *NetworkPolicyController) processNetworkPolicy(np *networkingv1.NetworkPolicy) (*antreatypes.NetworkPolicy, map[string]*antreatypes.AppliedToGroup, map[string]*antreatypes.AddressGroup) {
	_ = "STUB: not implemented"
	// appliedToGroups tracks all distinct AppliedToGroups referred to by the K8s NetworkPolicy.
	return nil, nil, nil
}

// addressGroups tracks all distinct AddressGroups referred to by the K8s NetworkPolicy.

// Retrieve Namespace logging annotation.

// Compute NetworkPolicyRule for Ingress Rule.

// Compute NetworkPolicyRule for Egress Rule.

// Traffic in a direction must be isolated if Spec.PolicyTypes specify it explicitly.

// If ingress isolation is specified explicitly and there's no ingress rule, append a deny-all ingress rule.
// See https://kubernetes.io/docs/concepts/services-networking/network-policies/#default-deny-all-ingress-traffic

// If egress isolation is specified explicitly and there's no egress rule, append a deny-all egress rule.
// See https://kubernetes.io/docs/concepts/services-networking/network-policies/#default-deny-all-egress-traffic

func (n *NetworkPolicyController) toAntreaPeer(peers []networkingv1.NetworkPolicyPeer, np *networkingv1.NetworkPolicy, dir controlplane.Direction, namedPortExists bool) (*controlplane.NetworkPolicyPeer, []*antreatypes.AddressGroup) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Empty NetworkPolicyPeer is supposed to match all addresses.
// See https://kubernetes.io/docs/concepts/services-networking/network-policies/#default-allow-all-ingress-traffic.
// It's treated as an IPBlock "0.0.0.0/0".

// For an egress Peer that specifies any named ports, it creates or
// reuses the AddressGroup matching all Pods in all Namespaces and
// appends the AddressGroup UID to the returned Peer such that it can be
// used to resolve the named ports.
// For other cases it uses the IPBlock "0.0.0.0/0" to avoid the overhead
// of handling member updates of the AddressGroup.

// A controlplane.NetworkPolicyPeer will either have an IPBlock or a
// podSelector and/or namespaceSelector set.

// addNetworkPolicy receives NetworkPolicy ADD events and creates resources
// which can be consumed by agents to configure corresponding rules on the Nodes.
func (n *NetworkPolicyController) addNetworkPolicy(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

// updateNetworkPolicy receives NetworkPolicy UPDATE events and updates resources
// which can be consumed by agents to configure corresponding rules on the Nodes.
func (n *NetworkPolicyController) updateNetworkPolicy(old, cur interface{}) {
	_ = "STUB: not implemented"
	return
}

// deleteNetworkPolicy receives NetworkPolicy DELETED events and deletes resources
// which can be consumed by agents to delete corresponding rules on the Nodes.
func (n *NetworkPolicyController) deleteNetworkPolicy(old interface{}) {
	_ = "STUB: not implemented"
	return
}

// addService retrieves all internal Groups which refers to this Service
// and enqueues the group keys for further processing.
func (n *NetworkPolicyController) addService(obj interface{}) { _ = "STUB: not implemented"; return }

// Find all internal Group keys which refers to this Service.

// Enqueue internal groups to its queue for group processing.

// updatePod retrieves all internal Groups which refers to this Service
// and enqueues the group keys for further processing.
func (n *NetworkPolicyController) updateService(oldObj, curObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// No need to trigger processing of groups if there is no change in the Service selectors.

// Find all internal Group keys which refers to this Service.

// Enqueue internal groups to its queue for group processing.

// deleteService retrieves all internal Groups which refers to this Service
// and enqueues the group keys for further processing.
func (n *NetworkPolicyController) deleteService(old interface{}) { _ = "STUB: not implemented"; return }

// Find all internal Group keys which refers to this Service.

// Enqueue internal groups to its queue for group processing.

func (n *NetworkPolicyController) enqueueAppliedToGroup(key string) {
	_ = "STUB: not implemented"
	return
}

func (n *NetworkPolicyController) enqueueAddressGroup(key string) {
	_ = "STUB: not implemented"
	return
}

func (n *NetworkPolicyController) enqueueInternalNetworkPolicy(key *controlplane.NetworkPolicyReference) {
	_ = "STUB: not implemented"
	return
}

// It must use value instead of pointer as the key, otherwise the same NetworkPolicies will not be treated as same
// item because the pointers may be different.

// Run begins watching and syncing of a NetworkPolicyController.
func (n *NetworkPolicyController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Only wait for acnpListerSynced and annpListerSynced when AntreaPolicy feature gate is enabled.

func (n *NetworkPolicyController) appliedToGroupWorker() { _ = "STUB: not implemented"; return }

func (n *NetworkPolicyController) addressGroupWorker() { _ = "STUB: not implemented"; return }

func (n *NetworkPolicyController) internalNetworkPolicyWorker() { _ = "STUB: not implemented"; return }

// Processes an item in the "internalNetworkPolicy" work queue, by calling
// syncInternalNetworkPolicy after casting the item to a string
// (NetworkPolicy key). If syncInternalNetworkPolicy returns an error, this
// function handles it by requeuing the item so that it can be processed again
// later. If syncInternalNetworkPolicy is successful, the NetworkPolicy is
// removed from the queue until we get notify of a new change. This function
// return false if and only if the work queue was shutdown (no more items will
// be processed).
func (n *NetworkPolicyController) processNextInternalNetworkPolicyWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// We call Done here so the workqueue knows we have finished processing this item. We also
// must remember to call Forget if we do not want this work item being re-queued. For
// example, we do not call Forget if a transient error occurs, instead the item is put back
// on the workqueue and attempted again after a back-off period.

// Put the item back on the workqueue to handle any transient errors.

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Processes an item in the "addressGroup" work queue, by calling
// syncAddressGroup after casting the item to a string (addressGroup key).
// If syncAddressGroup returns an error, this function handles it by requeuing
// the item so that it can be processed again later. If syncAddressGroup is
// successful, the AddressGroup is removed from the queue until we get notify
// of a new change. This function return false if and only if the work queue
// was shutdown (no more items will be processed).
func (n *NetworkPolicyController) processNextAddressGroupWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// Put the item back on the workqueue to handle any transient errors.

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Processes an item in the "appliedToGroup" work queue, by calling
// syncAppliedToGroup after casting the item to a string (appliedToGroup key).
// If syncAppliedToGroup returns an error, this function handles it by
// requeuing the item so that it can be processed again later. If
// syncAppliedToGroup is successful, the AppliedToGroup is removed from the
// queue until we get notify of a new change. This function return false if
// and only if the work queue was shutdown (no more items will be processed).
func (n *NetworkPolicyController) processNextAppliedToGroupWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// Put the item back on the workqueue to handle any transient errors.

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// syncAddressGroup retrieves all the internal NetworkPolicies which have a
// reference to this AddressGroup and updates its Pod IPAddresses set to
// reflect the current state of affected GroupMembers based on the GroupSelector.
func (n *NetworkPolicyController) syncAddressGroup(key string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get all internal NetworkPolicy objects that refers this AddressGroup.

// AddressGroup was already deleted. No need to process further.

// NodeNames set must be considered immutable once generated and updated
// in the store. If any change is needed, the set must be regenerated with
// the new NodeNames and the store must be updated.

func (c *NetworkPolicyController) getNodeMemberSet(selector labels.Selector) controlplane.GroupMemberSet {
	_ = "STUB: not implemented"
	return *new(controlplane.GroupMemberSet)
}

// getAddressGroupMemberSet knows how to construct a GroupMemberSet that contains
// all the entities selected by an AddressGroup.
func (n *NetworkPolicyController) getAddressGroupMemberSet(g *antreatypes.AddressGroup) controlplane.GroupMemberSet {
	_ = "STUB: not implemented"
	// This AddressGroup is derived from a ClusterGroup/Group.
	return *new(controlplane.GroupMemberSet)
}

// Check if an internal Group object exists corresponding to this AddressGroup.

// In case the ClusterGroup/Group is defined by a mix of childGroup with selectors and
// childGroup with ipBlocks, this function only returns the aggregated GroupMemberSet
// computed from childGroup with selectors, as ipBlocks will be processed differently.

// The internal Group doesn't exist yet or has been deleted. The AddressGroup selects nothing at the moment.
// Once the internalGroup is created, the AddressGroup will be resynced.

// Selector can't be nil when it reaches here.

// getInternalGroupMembers knows how to construct a GroupMemberSet and ipBlocks that contains
// all the entities selected by an internal Group. For internal Groups that has childGroups,
// the members are computed as the union of all its childGroup's members.
func (n *NetworkPolicyController) getInternalGroupMembers(group *antreatypes.Group) (controlplane.GroupMemberSet, []controlplane.IPBlock) {
	_ = "STUB: not implemented"
	return *new(controlplane.GroupMemberSet), nil
}

// getMemberSetForGroupType knows how to construct a GroupMemberSet for the given
// groupType and group name.
func (n *NetworkPolicyController) getMemberSetForGroupType(groupType grouping.GroupType, name string) controlplane.GroupMemberSet {
	_ = "STUB: not implemented"
	return *new(controlplane.GroupMemberSet)
}

// HostNetwork Pods should be excluded from group members: https://github.com/antrea-io/antrea/issues/3078.
// Terminated Pods should be excluded as their IPs can be recycled and used by other Pods.

// podToGroupMember is util function to convert a Pod to a GroupMember type.
// A controlplane.NamedPort item will be set in the GroupMember, only if the
// Pod contains a Port with the name field set. PodReference will also be set
// for converting GroupMember to GroupMemberPod for clients using older version
// of the controlplane API.
func podToGroupMember(pod *v1.Pod, includeIP bool) *controlplane.GroupMember {
	_ = "STUB: not implemented"
	return nil
}

// Only include container ports with name set.

func nodeToGroupMember(node *v1.Node, includeIP bool) (member *controlplane.GroupMember) {
	_ = "STUB: not implemented"
	return nil
}

// Sort the IPs to ensure the GroupMemberKey is deterministic.

func serviceToGroupMember(serviceReference *controlplane.ServiceReference) (member *controlplane.GroupMember) {
	_ = "STUB: not implemented"
	return nil
}

func externalEntityToGroupMember(ee *v1alpha2.ExternalEntity, includeIP bool) *controlplane.GroupMember {
	_ = "STUB: not implemented"
	return nil
}

// syncAppliedToGroup enqueues all the internal NetworkPolicy keys that
// refer this AppliedToGroup and update the AppliedToGroup Pod
// references by Node to reflect the latest set of affected GroupMembers based
// on it's GroupSelector.
func (n *NetworkPolicyController) syncAppliedToGroup(key string) error {
	_ = "STUB: not implemented"
	return nil
}

// AppliedToGroup for NodePort Service span to all Nodes.

// No need to process Pod when it's not scheduled or is already terminated.
// HostNetwork Pods will not be applied to by policies.

// Update the Pod references by Node.

// Update the NodeNames in order to set the SpanMeta for AppliedToGroup.

// Note that this must be executed after storing the result, to ensure that
// the notified subscribers get the latest state.

// getAppliedToWorkloads returns a list of workloads (Pods, ExternalEntities or Nodes) selected by an AppliedToGroup
// for standalone selectors or Pods and ExternalEntities corresponding to a ClusterGroup.
func (n *NetworkPolicyController) getAppliedToWorkloads(g *antreatypes.AppliedToGroup) ([]*v1.Pod, []*v1alpha2.ExternalEntity, []*v1.Node, error) {
	_ = "STUB: not implemented"
	// This AppliedToGroup is derived from a ClusterGroup/Group.
	return nil, nil, nil, nil
}

// Check if an internal Group object exists corresponding to this AppliedToGroup

// The internal Group doesn't exist yet or has been deleted. The AppliedToGroup selects nothing at the moment.
// Once the internalGroup is created, the AppliedToGroup will be resynced.

// Selector can't be nil when it reaches here.

// getInternalGroupWorkloads returns a list of workloads (Pods and ExternalEntities) selected by a ClusterGroup.
// For ClusterGroup that has childGroups, the workloads are computed as the union of all its childGroup's workloads.
func (n *NetworkPolicyController) getInternalGroupWorkloads(group *antreatypes.Group) ([]*v1.Pod, []*v1alpha2.ExternalEntity, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ClusterGroup can select entities in all Namespaces when used as AppliedTo.

// Namespaced Group can only select entities in the same Namespace as the Group when used as AppliedTo.

// childNameString will either be name of the child ClusterGroup or Namespaced name of the child Group.

func (n *NetworkPolicyController) triggerPolicyResyncForLabelIdentityUpdates(key string) {
	_ = "STUB: not implemented"
	return
}

// syncInternalNetworkPolicy retrieves all the AppliedToGroups associated with
// itself in order to calculate the Node span for this policy.
func (n *NetworkPolicyController) syncInternalNetworkPolicy(key *controlplane.NetworkPolicyReference) error {
	_ = "STUB: not implemented"
	return nil
}

// We need to check if the UID matches because it's possible another policy is created with the same name after
// the policy is deleted. It's safe to just delete the internal NetworkPolicy associated with the old policy as
// the two policies are different items in the workqueue and internalNetworkPolicyStore due to different UIDs.

// The NetworkPolicy must subscribe to the updates of AppliedToGroups before calculating span based on them,
// otherwise the calculated span may be outdated as AppliedToGroups can be updated concurrently and the
// NetworkPolicy wouldn't be notified.

// Calculate the set of Node names based on the span of the
// AppliedToGroups referenced by this NetworkPolicy.

// appliedToGroupsToSync tracks new AppliedToGroups created by this NetworkPolicy.

// Create the internal NetworkPolicy, AppliedToGroups and AddressGroups if they don't exist. They need to be updated
// atomically to avoid race conditions between workers that process multiple NetworkPolicies.

// AppliedToGroup is named based on its selector, so only its members can change.
// We don't need to update its selector if it already exists.

// AddressGroup is named based on its selector, so only its members can change.
// We don't need to update its selector if it already exists.

// For an AddressGroup that selects Nodes via nodeSelector, we calculate its members via NodeLister
// directly, instead of groupingInterface which handles Pod and ExternalEntity currently.

// Clean up orphan AddressGroups and AppliedToGroups that are no longer referenced by any NetworkPolicy.

// Enqueue AppliedToGroups that are newly created for this NetworkPolicy.

// Enqueue AddressGroups that are affected by this NetworkPolicy.

// Unsubscribe to the updates of the stale AppliedToGroups.

// deleteInternalNetworkPolicy deletes the internal NetworkPolicy and the referenced AppliedToGroups and AddressGroups
// if they are no longer referenced by any NetworkPolicy. They need to be updated atomically to avoid race conditions
// between workers that process multiple NetworkPolicies.
func (n *NetworkPolicyController) deleteInternalNetworkPolicy(name string) {
	_ = "STUB: not implemented"
	return
}

// Unsubscribe to the updates of the AppliedToGroups.

// Enqueue AddressGroups previously used by this NetworkPolicy as their span may change due to the removal.

// cleanupOrphanGroups deletes AddressGroups and AppliedToGroups that are no longer referenced by any NetworkPolicy.
func (n *NetworkPolicyController) cleanupOrphanGroups(internalNetworkPolicy *antreatypes.NetworkPolicy) {
	_ = "STUB: not implemented"
	return
}

// denyAllRule returns a NetworkPolicyRule which denies all traffic in the given direction.
func denyAllRule(direction controlplane.Direction, enableLogging bool) controlplane.NetworkPolicyRule {
	_ = "STUB: not implemented"
	return *new(controlplane.NetworkPolicyRule)
}

// ipStrToIPAddress converts an IP string to a controlplane.IPAddress.
// nil will returned if the IP string is not valid.
func ipStrToIPAddress(ip string) controlplane.IPAddress {
	_ = "STUB: not implemented"
	return *new(controlplane.IPAddress)
}

// cidrStrToIPNet converts a CIDR (eg. 10.0.0.0/16) to a *controlplane.IPNet.
func cidrStrToIPNet(cidr string) (*controlplane.IPNet, error) {
	_ = "STUB: not implemented"
	// Split the cidr to retrieve the IP and prefix.
	return nil, nil
}

// Convert prefix length to int32

// internalNetworkPolicyKeyFunc knows how to generate the key for an internal NetworkPolicy based on the object metadata
// of the corresponding original NetworkPolicy resource (also referred to as the "source").
// The key must be unique across K8s NetworkPolicies, Antrea NetworkPolicies, and Antrea ClusterNetworkPolicies.
// Currently the UID of the original NetworkPolicy is used to ensure uniqueness.
func internalNetworkPolicyKeyFunc(obj metav1.Object) string { _ = "STUB: not implemented"; return "" }

// internalGroupKeyFunc knows how to generate the key for an internal Group based on the object metadata
// of the corresponding Group and ClusterGroup resource. Currently the Name of the ClusterGroup is used to ensure
// uniqueness. Similarly, the Namespaced Name of the Group is used to ensure uniqueness for the Group resource.
func internalGroupKeyFunc(obj metav1.Object) string { _ = "STUB: not implemented"; return "" }

func getAppliedToGroupNames(groups []*antreatypes.AppliedToGroup) []string {
	_ = "STUB: not implemented"
	return nil
}

func getAddressGroupNames(groups []*antreatypes.AddressGroup) []string {
	_ = "STUB: not implemented"
	return nil
}

func mergeAppliedToGroups(dst map[string]*antreatypes.AppliedToGroup, src ...*antreatypes.AppliedToGroup) map[string]*antreatypes.AppliedToGroup {
	_ = "STUB: not implemented"
	return nil
}

func mergeAddressGroups(dst map[string]*antreatypes.AddressGroup, src ...*antreatypes.AddressGroup) map[string]*antreatypes.AddressGroup {
	_ = "STUB: not implemented"
	return nil
}
