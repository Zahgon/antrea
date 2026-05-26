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

package monitor

import (
	"time"

	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	clientset "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	externalnodeinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha1"
	externalnodelisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1alpha1"
	controllerquerier "antrea.io/antrea/v2/pkg/controller/querier"
)

const (
	controllerName = "AntreaControllerMonitor"
	// How long to wait before retrying the processing of a Node/ExternalNode change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing a Node/ExternalNode change.
	defaultWorkers        = 4
	agentInfoResourceKind = "AntreaAgentInfo"
)

var (
	keyFunc      = cache.DeletionHandlingMetaNamespaceKeyFunc
	splitKeyFunc = cache.SplitMetaNamespaceKey
)

type controllerMonitor struct {
	client       clientset.Interface
	nodeInformer coreinformers.NodeInformer
	nodeLister   corelisters.NodeLister
	// nodeListerSynced is a function which returns true if the node shared informer has been synced at least once.
	nodeListerSynced cache.InformerSynced

	externalNodeInformer     externalnodeinformers.ExternalNodeInformer
	externalNodeLister       externalnodelisters.ExternalNodeLister
	externalNodeListerSynced cache.InformerSynced

	externalNodeEnabled bool

	nodeQueue         workqueue.TypedRateLimitingInterface[string]
	externalNodeQueue workqueue.TypedRateLimitingInterface[string]

	querier controllerquerier.ControllerQuerier
	// controllerCRD is the desired state of controller monitoring CRD which controllerMonitor expects.
	controllerCRD *v1beta1.AntreaControllerInfo
}

// NewControllerMonitor creates a new controller monitor.
func NewControllerMonitor(
	client clientset.Interface,
	nodeInformer coreinformers.NodeInformer,
	externalNodeInformer externalnodeinformers.ExternalNodeInformer,
	querier controllerquerier.ControllerQuerier,
	externalNodeEnabled bool,
) *controllerMonitor {
	_ = "STUB: not implemented"
	return nil
}

// Register Informer and add handlers for ExternalNode events only if the feature is enabled.

// Run creates AntreaControllerInfo CRD first after controller is running.
// Then updates AntreaControllerInfo CRD every 60 seconds if there is any change.
func (monitor *controllerMonitor) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Only wait for externalNodeListerSynced when ExternalNode feature is enabled.

// Sync controller monitoring CRD every minute util stopCh is closed.

func (monitor *controllerMonitor) syncControllerCRD() { _ = "STUB: not implemented"; return }

// getControllerCRD is used to check the existence of controller monitoring CRD.
// So when the Pod restarts, it will update this monitoring CRD instead of creating a new one.
func (monitor *controllerMonitor) getControllerCRD(crdName string) (*v1beta1.AntreaControllerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (monitor *controllerMonitor) createControllerCRD(crdName string) (*v1beta1.AntreaControllerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// updateControllerCRD updates the monitoring CRD.
func (monitor *controllerMonitor) updateControllerCRD(partial bool) (*v1beta1.AntreaControllerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (monitor *controllerMonitor) deleteStaleAgentCRDs() { _ = "STUB: not implemented"; return }

// Delete stale agent monitoring CRD based on existing Nodes and ExternalNodes.

func (monitor *controllerMonitor) enqueueNode(obj interface{}) { _ = "STUB: not implemented"; return }

func (monitor *controllerMonitor) enqueueExternalNode(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (n *controllerMonitor) nodeWorker() { _ = "STUB: not implemented"; return }

func (n *controllerMonitor) externalNodeWorker() { _ = "STUB: not implemented"; return }

func (c *controllerMonitor) processNextNodeWorkItem() bool { _ = "STUB: not implemented"; return false }

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

func (c *controllerMonitor) processNextExternalNodeWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

func (c *controllerMonitor) syncNode(key string) error { _ = "STUB: not implemented"; return nil }

// This err should not occur.

func (c *controllerMonitor) syncExternalNode(key string) error {
	_ = "STUB: not implemented"
	return nil
}

// This err should not occur.

func (monitor *controllerMonitor) createAgentCRD(name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (monitor *controllerMonitor) deleteAgentCRD(name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (monitor *controllerMonitor) antreaAgentInfoAPIAvailable(stopCh <-chan struct{}) bool {
	_ = "STUB: not implemented"
	return false
}
