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
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	mcinformers "antrea.io/antrea/v2/multicluster/pkg/client/informers/externalversions/multicluster/v1alpha1"
	mclisters "antrea.io/antrea/v2/multicluster/pkg/client/listers/multicluster/v1alpha1"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/util/channel"
)

const (
	stretchedNetworkPolicyWorker         = 4
	stretchedNetworkPolicyControllerName = "AntreaAgentStretchedNetworkPolicyController"

	labelIndex = "Label"
)

type podSet map[types.NamespacedName]struct{}

// StretchedNetworkPolicyController is used to update classifier flows of Pods.
// It will make sure the latest LabelIdentity of the Pod, if available, will be
// loaded into tun_id in the classifier flow of the Pod.
// If the LabelIdentity of the Pod is not available when updating, the
// UnknownLabelIdentity will be loaded. When the actual LabelIdentity is created,
// the classifier flow will be updated accordingly.
type StretchedNetworkPolicyController struct {
	ofClient                  openflow.Client
	interfaceStore            interfacestore.InterfaceStore
	podInformer               cache.SharedIndexInformer
	podLister                 corelisters.PodLister
	podListerSynced           cache.InformerSynced
	namespaceInformer         coreinformers.NamespaceInformer
	namespaceLister           corelisters.NamespaceLister
	namespaceListerSynced     cache.InformerSynced
	labelIdentityInformer     mcinformers.LabelIdentityInformer
	labelIdentityLister       mclisters.LabelIdentityLister
	LabelIdentityListerSynced cache.InformerSynced
	queue                     workqueue.TypedRateLimitingInterface[types.NamespacedName]
	lock                      sync.RWMutex

	labelToPods map[string]podSet
	podToLabel  map[types.NamespacedName]string
}

func NewMCAgentStretchedNetworkPolicyController(
	client openflow.Client,
	interfaceStore interfacestore.InterfaceStore,
	podInformer cache.SharedIndexInformer,
	namespaceInformer coreinformers.NamespaceInformer,
	labelIdentityInformer mcinformers.LabelIdentityInformer,
	podUpdateSubscriber channel.Subscriber,
) *StretchedNetworkPolicyController {
	_ = "STUB: not implemented"
	return nil
}

// Pod add event will be handled by processPodCNIAddEvent.
// We choose to use events from podUpdateSubscriber instead of the informer because
// the controller can only update the Pod classifier flow when the Pod container
// config is available. Events from the Informer may be received way before the Pod
// container config is available, which will cause the work item be continually
// re-queued with an exponential increased delay time. When the Pod container
// config is ready, the work item could wait for a long time to be processed.

func (s *StretchedNetworkPolicyController) Run(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (s *StretchedNetworkPolicyController) enqueueAllPods() { _ = "STUB: not implemented"; return }

// worker is a long-running function that will continually call the processNextWorkItem
// function in order to read and process a message on the workqueue.
func (s *StretchedNetworkPolicyController) worker() { _ = "STUB: not implemented"; return }

func (s *StretchedNetworkPolicyController) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// Put the item back on the workqueue to handle any transient errors.

// syncPodClassifierFlow gets containerConfigs and labelIdentity according to a
// Pod reference and updates this Pod's classifierFlow.
func (s *StretchedNetworkPolicyController) syncPodClassifierFlow(podRef types.NamespacedName) error {
	_ = "STUB: not implemented"
	return nil
}

// getLabelIdentity updates labelToPods and podToLabel and returns the
// LabelIdentity based on the normalizedLabel.
func (s *StretchedNetworkPolicyController) getLabelIdentity(podRef types.NamespacedName, normalizedLabel string) uint32 {
	_ = "STUB: not implemented"
	return 0
}

func (s *StretchedNetworkPolicyController) processPodCNIAddEvent(e interface{}) {
	_ = "STUB: not implemented"
	return
}

// processPodUpdate handles Pod update events. It only enqueues the Pod if the
// Labels of this Pod have been updated.
func (s *StretchedNetworkPolicyController) processPodUpdate(old, cur interface{}) {
	_ = "STUB: not implemented"
	return
}

// processPodDelete handles Pod delete events. It deletes the Pod from the
// labelToPods and podToLabel. After Pod is deleted, its classifier flow will also
// be deleted by podConfigurator. So no need to enqueue this Pod to update its
// classifier flow.
func (s *StretchedNetworkPolicyController) processPodDelete(old interface{}) {
	_ = "STUB: not implemented"
	return
}

// processNamespaceUpdate handles Namespace update events. It only enqueues all
// Pods in this Namespace if the Labels of this Namespace have been updated.
func (s *StretchedNetworkPolicyController) processNamespaceUpdate(old, cur interface{}) {
	_ = "STUB: not implemented"
	return
}

// processLabelIdentityEvent handles labelIdentity add/update/delete event.
// It will enqueue all Pods affected by this labelIdentity
func (s *StretchedNetworkPolicyController) processLabelIdentityEvent(cur interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *StretchedNetworkPolicyController) addLabelToPod(normalizedLabel string, podRef types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

func (s *StretchedNetworkPolicyController) deleteLabelToPod(normalizedLabel string, podRef types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

func getPodReference(pod *v1.Pod) types.NamespacedName {
	_ = "STUB: not implemented"
	return *new(types.NamespacedName)
}
