// Copyright 2021 Antrea Authors
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

package serviceexternalip

import (
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	apimachinerytypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	coreinformers "k8s.io/client-go/informers/core/v1"
	discoveryinformers "k8s.io/client-go/informers/discovery/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	discoverylisters "k8s.io/client-go/listers/discovery/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/apis"
	"antrea.io/antrea/v2/pkg/agent/ipassigner"
	"antrea.io/antrea/v2/pkg/agent/ipassigner/linkmonitor"
	"antrea.io/antrea/v2/pkg/agent/memberlist"
	"antrea.io/antrea/v2/pkg/querier"
)

const (
	controllerName = "ServiceExternalIPController"
	// How long to wait before retrying the processing of an Service change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing an Service change.
	defaultWorkers = 1
	// Disable resyncing.
	resyncPeriod time.Duration = 0

	externalIPIndex     = "externalIP"
	externalIPPoolIndex = "externalIPPool"
)

type externalIPState struct {
	ip           string
	ipPool       string
	assignedNode string
}

type ServiceExternalIPController struct {
	nodeName            string
	serviceInformer     cache.SharedIndexInformer
	serviceLister       corelisters.ServiceLister
	serviceListerSynced cache.InformerSynced

	endpointSliceInformer     cache.SharedIndexInformer
	endpointSliceLister       discoverylisters.EndpointSliceLister
	endpointSliceListerSynced cache.InformerSynced

	queue workqueue.TypedRateLimitingInterface[apimachinerytypes.NamespacedName]

	externalIPStates      map[apimachinerytypes.NamespacedName]externalIPState
	externalIPStatesMutex sync.RWMutex

	cluster    memberlist.Interface
	ipAssigner ipassigner.IPAssigner

	assignedIPs      map[string]sets.Set[string]
	assignedIPsMutex sync.Mutex

	linkMonitor linkmonitor.Interface
}

var _ querier.ServiceExternalIPStatusQuerier = (*ServiceExternalIPController)(nil)

func NewServiceExternalIPController(
	nodeName string,
	nodeTransportInterface string,
	cluster memberlist.Interface,
	serviceInformer coreinformers.ServiceInformer,
	endpointSliceInformer discoveryinformers.EndpointSliceInformer,
	linkMonitor linkmonitor.Interface,
) (*ServiceExternalIPController, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// On Windows, ipassigner.NewIPAssigner always returns a non-nil error (see ip_assigner_windows.go);
// on Linux, err is nil when initialization succeeds. golangci runs staticcheck with GOOS=windows too.
//nolint:staticcheck // SA4023: err is always non-nil on Windows only.

func (c *ServiceExternalIPController) enqueueService(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *ServiceExternalIPController) enqueueServiceForEndpointSlice(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

// Get the service name from the EndpointSlice label

// EndpointSlice doesn't have the service name label, skip it

// The only possible error Lister.Get can return is NotFound.
// It's fine to ignore the error as the Service's add event will enqueue it when the Service is synced.

// we only care services with ServiceExternalTrafficPolicy setting to local.

// enqueueServicesByExternalIPPool enqueues all services that refer to the provided ExternalIPPool,
// the ExternalIPPool is affected by a Node update/create/delete event or Node leaves/join cluster
// event or ExternalIPPool changed event.
func (c *ServiceExternalIPController) enqueueServicesByExternalIPPool(eipName string) {
	_ = "STUB: not implemented"
	return
}

// Run will create defaultWorkers workers (go routines) which will process the Service events from the
// workqueue.
func (c *ServiceExternalIPController) Run(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// worker is a long-running function that will continually call the processNextWorkItem function in
// order to read and process a message on the workqueue.
func (c *ServiceExternalIPController) worker() { _ = "STUB: not implemented"; return }

func (c *ServiceExternalIPController) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// Put the item back on the workqueue to handle any transient errors.

func (c *ServiceExternalIPController) deleteService(service apimachinerytypes.NamespacedName) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ServiceExternalIPController) getServiceState(service *corev1.Service) (externalIPState, bool) {
	_ = "STUB: not implemented"
	return *new(externalIPState), false
}

func (c *ServiceExternalIPController) saveServiceState(service *corev1.Service, state *externalIPState) {
	_ = "STUB: not implemented"
	return
}

func (c *ServiceExternalIPController) getServiceExternalIP(service *corev1.Service) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ServiceExternalIPController) syncService(key apimachinerytypes.NamespacedName) error {
	_ = "STUB: not implemented"
	return nil
}

// External IP of the Service has changed. Delete the previous assigned IP if exists.

// No Node is available at the moment. The Service will be requeued by EndpointSlice, Node, or Memberlist update events.

func (c *ServiceExternalIPController) assignIP(ip string, service apimachinerytypes.NamespacedName) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ServiceExternalIPController) unassignIP(ip string, service apimachinerytypes.NamespacedName) error {
	_ = "STUB: not implemented"
	return nil
}

// nodesHasHealthyServiceEndpoint returns the set of Nodes which has at least one healthy endpoint
// for the address family matching the service's external IP.
func (c *ServiceExternalIPController) nodesHasHealthyServiceEndpoint(service *corev1.Service) (sets.Set[string], error) {
	_ = "STUB: not implemented"
	return nil,

		// List all EndpointSlices for this service using the label selector
		nil
}

// Determine the address type matching the service's external IP, so that on dual-stack clusters
// we only consider EndpointSlices of the correct address family.

// Check the ready condition first to respect the Service's publishNotReadyAddresses setting.
// The ready condition is true when:
// - publishNotReadyAddresses is true (all endpoints are considered ready), OR
// - the endpoint is serving AND not terminating
// If ready is true (or nil, which means true), we can use this endpoint.

// If ready is false, fall back to checking the serving condition directly.
// This handles cases where the endpoint might still be serving but is marked not ready
// (e.g., during termination but still draining connections).

func (c *ServiceExternalIPController) GetServiceExternalIPStatus() []apis.ServiceExternalIPInfo {
	_ = "STUB: not implemented"
	return nil
}
