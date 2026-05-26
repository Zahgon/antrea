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

package supportbundlecollection

import (
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/util/sets"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	"antrea.io/antrea/v2/pkg/apiserver/storage"
	clientset "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha1"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1alpha1"
	"antrea.io/antrea/v2/pkg/controller/types"
)

const (
	controllerName = "SupportBundleCollectionController"
	// How long to wait before retrying the processing of an ExternalNode change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Set resyncPeriod to 0 to disable resyncing.
	resyncPeriod time.Duration = 0
	// supportBundleCollectionRetryPeriod is the duration after which to retry a SupportBundleCollection
	// request if it conflicts with a processing request.
	supportBundleCollectionRetryPeriod = time.Second * 10
)

const (
	processingNodesIndex         = "processingNodes"
	processingExternalNodesIndex = "processingExternalNodes"
	processingNodesIndexValue    = "processingNodes"
)

// supportBundleCollectionAppliedTo is defined to maintain a SupportBundleCollection's required Nodes and ExternalNodes.
type supportBundleCollectionAppliedTo struct {
	// The name of a SupportBundleCollection
	name         string
	processNodes bool
	enNamespace  string
}

func getSupportBundleCollectionKey(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func processingNodesIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processingExternalNodesIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Controller struct {
	kubeClient kubernetes.Interface
	crdClient  clientset.Interface

	supportBundleCollectionInformer     crdinformers.SupportBundleCollectionInformer
	supportBundleCollectionLister       crdlisters.SupportBundleCollectionLister
	supportBundleCollectionListerSynced cache.InformerSynced
	nodeLister                          corelisters.NodeLister
	nodeListerSynced                    cache.InformerSynced
	externalNodeLister                  crdlisters.ExternalNodeLister
	externalNodeListerSynced            cache.InformerSynced

	// queue maintains the ExternalNode objects that need to be synced.
	queue workqueue.TypedRateLimitingInterface[string]

	// supportBundleCollectionStore is the storage where the populated internal support bundle collections are stored.
	supportBundleCollectionStore storage.Interface
	// supportBundleCollectionAppliedToStore is the storage where the required Nodes or ExternalNodes of a
	// SupportBundleCollection are stored.
	supportBundleCollectionAppliedToStore cache.Indexer

	// statuses is a nested map that keeps the realization statuses reported by antrea-agents.
	// The outer map's keys are the SupportBundleCollection names. The inner map's keys are the Node names. The inner
	// map's values are statuses reported by each Node for a SupportBundleCollection.
	statuses     map[string]map[string]*controlplane.SupportBundleCollectionNodeStatus
	statusesLock sync.RWMutex
}

func NewSupportBundleCollectionController(
	kubeClient kubernetes.Interface,
	crdClient clientset.Interface,
	supportBundleInformer crdinformers.SupportBundleCollectionInformer,
	nodeInformer coreinformers.NodeInformer,
	externalNodeInformer crdinformers.ExternalNodeInformer,
	supportBundleCollectionStore storage.Interface) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Run will create defaultWorkers workers (goroutines) which will process the SupportBundle events from the work queue.
func (c *Controller) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// UpdateStatus is called when Agent reports status to the internal SupportBundleCollection resource.
func (c *Controller) UpdateStatus(status *controlplane.SupportBundleCollectionStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) addSupportBundleCollection(obj interface{}) { _ = "STUB: not implemented"; return }

// updateSupportBundleCollection adds the SupportBundleCollection name into queue if the conditions are updated. The
// changes in SupportBundleCollection.Spec is ignored, as we do not support Spec changes after the collection is started.
func (c *Controller) updateSupportBundleCollection(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) deleteSupportBundleCollection(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

// reconcileExternalNodes reconciles all the existing support bundles which are in BundleProcessing phase.
func (c *Controller) reconcileSupportBundleCollections() error {
	_ = "STUB: not implemented"
	return nil
}

// Continue processing the started SupportBundleCollection request if it is not completed before restart.

// worker is a long-running function that will continually call the processNextWorkItem function in
// order to read and process a message on the work queue.
func (c *Controller) worker() { _ = "STUB: not implemented"; return }

func (c *Controller) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// If no error occurs we Forget this item, so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

func (c *Controller) syncSupportBundleCollection(key string) error {
	_ = "STUB: not implemented"
	return nil
}

// If the SupportBundleCollection is completed, remove the internal SupportBundleCollection.

// createInternalSupportBundleCollection creates internal SupportBundle object and saves it into the storage.
func (c *Controller) createInternalSupportBundleCollection(bundle *v1alpha1.SupportBundleCollection) (*types.SupportBundleCollection, error) {
	_ = "STUB: not implemented"
	// Calculate the expiration time with ExpirationMinutes and the created time of the resource.
	// It is to ensure the expiration time of the internal support bundle collection resource is constant, and to
	// ensure the processing is atomic even if SupportBundleCollectionController is restarted.
	return nil, nil
}

// Create a CollectionStarted failure condition on the CRD if time is expired. Return nil to avoid the event
// to be re-enqueued.

// Get expected Kubernetes Nodes defined in the CRD.

// Get expected external Nodes defined in the CRD.

// Get authentication from the Secret provided in authentication field in the CRD

// Process the support bundle collection when time is up, this will create a CollectionFailure condition if the
// bundle collection is not completed in time because any Agent fails to upload the files and does not report
// the failure.

// getBundleNodes returns the names of the Nodes configured in the SupportBundleCollection.
func (c *Controller) getBundleNodes(nodes *v1alpha1.BundleNodes) (sets.Set[string], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return all Kubernetes Nodes if NodeNames is empty and NodeSelector is not specified.

// Add the Nodes which are configured with the names defined in the resource.

// Add the Nodes which are configured with the given label.

// getBundleExternalNodes returns the names of the ExternalNodes configured in the SupportBundleCollection.
func (c *Controller) getBundleExternalNodes(en *v1alpha1.BundleExternalNodes) (sets.Set[string], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return all ExternalNodes in the en.Namespace if both NodeNames is empty and NodeSelector is not specified.

func (c *Controller) deleteInternalSupportBundleCollection(key string) error {
	_ = "STUB: not implemented"
	return nil
}

// addInternalSupportBundleCollection adds internalBundle into supportBundleCollectionStore, and creates a
// supportBundleCollectionAppliedTo resource to maintain the SupportBundleCollection's required Nodes or ExternalNodes.
func (c *Controller) addInternalSupportBundleCollection(
	bundleCollection *v1alpha1.SupportBundleCollection,
	nodeSpan sets.Set[string],
	authentication *controlplane.BundleServerAuthConfiguration,
	expiredAt metav1.Time) *types.SupportBundleCollection {
	_ = "STUB: not implemented"
	return nil
}

// Create internal SupportBundleCollection resource.

// processConflictedCollection adds a Started failure condition on the conflicted the SupportBundleCollection request,
// and re-enqueue the request after 10s.
func (c *Controller) processConflictedCollection(bundle *v1alpha1.SupportBundleCollection) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) addConditions(bundleCollectionName string, conditions []v1alpha1.SupportBundleCollectionCondition) error {
	_ = "STUB: not implemented"
	return nil
}

// Return the error from UPDATE.

// isCollectionAvailable checks if the bundleCollection can be processed at once. It returns true with these conditions:
//  1. the bundleCollection is started processing;
//  2. there are no processing SupportBundleCollections requiring to collect bundle files on any Nodes, if this one requires
//     to collection files on Nodes;
//  3. there are no processing SupportBundleCollections requiring to collect bundle files on the ExternalNodes in the same
//     Namespace as this one.
func (c *Controller) isCollectionAvailable(bundleCollection *v1alpha1.SupportBundleCollection) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) updateStatus(internalBundleCollection *types.SupportBundleCollection) error {
	_ = "STUB: not implemented"
	return nil
}

// The node is no longer in the span of this Support Bundle Collection, delete its status.

// Mark the support bundle collection as started since the internal resource successfully created.
// It will not be added as a duplication if it already exists.

func (c *Controller) getNodeStatuses(key string) []*controlplane.SupportBundleCollectionNodeStatus {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) deleteNodeStatus(key string, nodeName string) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) clearStatuses(key string) { _ = "STUB: not implemented"; return }

func (c *Controller) updateSupportBundleCollectionStatus(name string, updatedStatus *v1alpha1.SupportBundleCollectionStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// If the current status equals to the desired status, no need to update.

// Return the error from UPDATE.

// isCollectionCompleted check if CollectionCompleted condition with status ConditionTrue is added in the bundleCollection or not.
func isCollectionCompleted(bundleCollection *v1alpha1.SupportBundleCollection) bool {
	_ = "STUB: not implemented"
	return false
}

// isCollectionProcessing check if CollectionStarted condition with status ConditionTrue is added in the bundleCollection or not.
func isCollectionProcessing(bundleCollection *v1alpha1.SupportBundleCollection) bool {
	_ = "STUB: not implemented"
	return false
}

func conditionEqualsIgnoreLastTransitionTime(a, b v1alpha1.SupportBundleCollectionCondition) bool {
	_ = "STUB: not implemented"
	return false
}

func conditionExistsIgnoreLastTransitionTime(conditions []v1alpha1.SupportBundleCollectionCondition, condition v1alpha1.SupportBundleCollectionCondition) bool {
	_ = "STUB: not implemented"
	return false
}

func mergeConditions(oldConditions, newConditions []v1alpha1.SupportBundleCollectionCondition) []v1alpha1.SupportBundleCollectionCondition {
	_ = "STUB: not implemented"
	return nil
}

// Use the original Condition if the only change is about lastTransition time

// Use the latest Condition.

func getNodeKey(status *controlplane.SupportBundleCollectionNodeStatus) string {
	_ = "STUB: not implemented"
	return ""
}

var semanticIgnoreLastTransitionTime = conversion.EqualitiesOrDie(
	conditionSliceEqualsIgnoreLastTransitionTime,
)

// supportBundleCollectionStatusEqual compares two SupportBundleCollectionStatus objects. It disregards
// the LastTransitionTime field in the status Conditions.
func supportBundleCollectionStatusEqual(oldStatus, newStatus v1alpha1.SupportBundleCollectionStatus) bool {
	_ = "STUB: not implemented"
	return false
}

func conditionSliceEqualsIgnoreLastTransitionTime(as, bs []v1alpha1.SupportBundleCollectionCondition) bool {
	_ = "STUB: not implemented"
	return false
}
