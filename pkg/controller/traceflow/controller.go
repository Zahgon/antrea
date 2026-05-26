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
	"sync"
	"time"

	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/client/clientset/versioned"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
)

const (
	controllerName = "TraceflowController"

	// Set resyncPeriod to 0 to disable resyncing.
	resyncPeriod time.Duration = 0

	// How long to wait before retrying the processing of a traceflow.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second

	// Default number of workers processing traceflow request.
	defaultWorkers = 4

	// Min and max data plane tag for traceflow. minTagNum is 7 (0b000111), maxTagNum is 59 (0b111011).
	// As per RFC2474, 16 different DSCP values are we reserved for Experimental or Local Use, which we use as the 16 possible data plane tag values.
	// tagStep is 4 (0b100) to keep last 2 bits at 0b11.
	tagStep   uint8 = 0b100
	minTagNum uint8 = 0b1*tagStep + 0b11
	maxTagNum uint8 = 0b1110*tagStep + 0b11

	// String set to TraceflowStatus.Reason.
	traceflowTimeout = "Traceflow timeout"

	// Traceflow timeout period.
	defaultTimeoutDuration = time.Second * time.Duration(crdv1beta1.DefaultTraceflowTimeout)
)

var (
	timeoutCheckInterval = 10 * time.Second
)

// Controller is for traceflow.
type Controller struct {
	client                 versioned.Interface
	podInformer            coreinformers.PodInformer
	podLister              corelisters.PodLister
	traceflowInformer      crdinformers.TraceflowInformer
	traceflowLister        crdlisters.TraceflowLister
	traceflowListerSynced  cache.InformerSynced
	queue                  workqueue.TypedRateLimitingInterface[string]
	runningTraceflowsMutex sync.Mutex
	runningTraceflows      map[uint8]string // tag->traceflowName if tf.Status.Phase is Running.
}

// NewTraceflowController creates a new traceflow controller and adds podIP indexer to podInformer.
func NewTraceflowController(client versioned.Interface, podInformer coreinformers.PodInformer, traceflowInformer crdinformers.TraceflowInformer) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// Add handlers for ClusterNetworkPolicy events.

// enqueueTraceflow adds an object to the controller work queue.
func (c *Controller) enqueueTraceflow(tf *crdv1beta1.Traceflow) { _ = "STUB: not implemented"; return }

func (c *Controller) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Load all data plane tags from CRD into controller's cache.

func (c *Controller) addTraceflow(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updateTraceflow(_, curObj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) deleteTraceflow(old interface{}) { _ = "STUB: not implemented"; return }

// worker is a long-running function that will continually call the processTraceflowItem function
// in order to read and process a message on the workqueue.
func (c *Controller) worker() { _ = "STUB: not implemented"; return }

func (c *Controller) checkTraceflowTimeout() { _ = "STUB: not implemented"; return }

// Re-post all running Traceflow requests to the work queue to
// be processed and checked for timeout.

// processTraceflowItem processes an item in the "traceflow" work queue, by calling syncTraceflow
// after casting the item to a string (Traceflow name). If syncTraceflow returns an error, this
// function logs the error and adds the Traceflow request back to the queue with a rate limit. If
// no error occurs, the Traceflow request is removed from the queue until we get notified of a new
// change. This function returns false if and only if the work queue was shutdown (no more items
// will be processed).
func (c *Controller) processTraceflowItem() bool { _ = "STUB: not implemented"; return false }

// We call Done here so the workqueue knows we have finished processing this item. We also
// must remember to call Forget if we do not want this work item being re-queued. For
// example, we do not call Forget if a transient error occurs, instead the item is put back
// on the workqueue and attempted again after a back-off period.

func (c *Controller) syncTraceflow(traceflowName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Traceflow CRD has been deleted.

// Deallocate tag when agent set Traceflow status to Failed.

func (c *Controller) startTraceflow(tf *crdv1beta1.Traceflow) error {
	_ = "STUB: not implemented"
	// Allocate data plane tag.
	return nil
}

// checkTraceflowStatus is only called for Traceflows in the Running phase
func (c *Controller) checkTraceflowStatus(tf *crdv1beta1.Traceflow) error {
	_ = "STUB: not implemented"
	return nil
}

// There should be only one reported NodeResult for droppedOnly
// Traceflow.

// Add Pod ns/name to observation if TranslatedDstIP (a.k.a. Service Endpoint address) is Pod IP.

// When the Source Pod is specified, the Traceflow should receive
// results from both the sender and the receiver. When the Source
// Pod is not specified (in live-traffic Traceflow), only the
// receiver Node will report the results.

// a fallback that should not be needed in general since we are in the Running phase
// when upgrading Antrea from a previous version, the field would be empty

func (c *Controller) updateTraceflowStatus(tf *crdv1beta1.Traceflow, phase crdv1beta1.TraceflowPhase, reason string, dataPlaneTag uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) occupyTag(tf *crdv1beta1.Traceflow) error {
	_ = "STUB: not implemented"
	return nil
}

// Allocates a tag. If the Traceflow request has been allocated with a tag
// already, 0 is returned. If number of existing Traceflow requests reaches
// the upper limit, an error is returned.
func (c *Controller) allocateTag(name string) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// The Traceflow request has been processed already.

// Deallocates tag from cache. Ignore DataplaneTag == 0 which is an invalid case.
func (c *Controller) deallocateTagForTF(tf *crdv1beta1.Traceflow) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) deallocateTag(name string, tag uint8) { _ = "STUB: not implemented"; return }
