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

	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha2"
	clientset "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	externalnodeinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha1"
	externalentityinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha2"
	externalnodelisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1alpha1"
	externalentitylisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1alpha2"
)

const (
	controllerName = "ExternalNodeController"
	// How long to wait before retrying the processing of an ExternalNode change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing ExternalNode changes.
	defaultWorkers = 4
	// Set resyncPeriod to 0 to disable resyncing.
	resyncPeriod time.Duration = 0
)

var (
	keyFunc      = cache.DeletionHandlingMetaNamespaceKeyFunc
	splitKeyFunc = cache.SplitMetaNamespaceKey
)

type ExternalNodeController struct {
	crdClient clientset.Interface

	externalNodeInformer     externalnodeinformers.ExternalNodeInformer
	externalNodeLister       externalnodelisters.ExternalNodeLister
	externalNodeListerSynced cache.InformerSynced

	externalEntityInformer     externalentityinformers.ExternalEntityInformer
	externalEntityLister       externalentitylisters.ExternalEntityLister
	externalEntityListerSynced cache.InformerSynced

	syncedExternalNode cache.Store
	// queue maintains the ExternalNode objects that need to be synced.
	queue workqueue.TypedRateLimitingInterface[string]
}

func NewExternalNodeController(crdClient clientset.Interface, externalNodeInformer externalnodeinformers.ExternalNodeInformer,
	externalEntityInformer externalentityinformers.ExternalEntityInformer) *ExternalNodeController {
	_ = "STUB: not implemented"
	return nil
}

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

// Run will create defaultWorkers workers (goroutines) which will process the ExternalEntity events from the work queue.
func (c *ExternalNodeController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// reconcileExternalNodes reconciles all the existing ExternalNodes and cleans up the stale ExternalEntities.
func (c *ExternalNodeController) reconcileExternalNodes() error {
	_ = "STUB: not implemented"
	return nil
}

// Clean up stale ExternalEntities when ExternalNode no longer exists or
// when interface[0] name is changed.

// worker is a long-running function that will continually call the processNextWorkItem function in
// order to read and process a message on the work queue.
func (c *ExternalNodeController) worker() { _ = "STUB: not implemented"; return }

func (c *ExternalNodeController) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

func (c *ExternalNodeController) syncExternalNode(key string) error {
	_ = "STUB: not implemented"
	return nil
}

// This err should not occur.

// addExternalNode creates ExternalEntity for each NetworkInterface in the ExternalNode.
// Only one interface is supported for now and there should be one ExternalEntity generated for one ExternalNode.
func (c *ExternalNodeController) addExternalNode(en *v1alpha1.ExternalNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExternalNodeController) createExternalEntity(ee *v1alpha2.ExternalEntity) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExternalNodeController) updateExternalNode(preEn *v1alpha1.ExternalNode, curEn *v1alpha1.ExternalNode) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete the previous ExternalEntity and create a new one if the name of the generated ExternalEntity is changed.
// Otherwise, update the ExternalEntity.

func (c *ExternalNodeController) updateExternalEntity(ee *v1alpha2.ExternalEntity) error {
	_ = "STUB: not implemented"
	// resourceVersion must be specified for update operation,
	// so it gets the existing ExternalEntity and modifies the changed fields.
	return nil
}

func (c *ExternalNodeController) deleteExternalNode(namespace string, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExternalNodeController) deleteExternalEntity(namespace string, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func genExternalEntity(eeName string, en *v1alpha1.ExternalNode) (*v1alpha2.ExternalEntity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This should not happen since openAPIV3Schema checks it.

// Generate one/multiple endpoint(s) if one/multiple IP(s) are specified for interface[0]
