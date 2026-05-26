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
	"net"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	apimachinerytypes "k8s.io/apimachinery/pkg/types"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/controller/externalippool"
)

const (
	controllerName = "ExternalIPController"
	// Set resyncPeriod to 0 to disable resyncing.
	resyncPeriod time.Duration = 0
	// How long to wait before retrying the processing of an Service change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing an Service change.
	defaultWorkers = 4

	// externalIPPoolIndex is an index of serviceInformer.
	externalIPPoolIndex = "externalIPPool"
	// ipIndex is an index of ipAllocations.
	ipIndex = "ip"
)

// ipAllocation contains the IP and the IP Pool which allocates it.
type ipAllocation struct {
	service  apimachinerytypes.NamespacedName
	ip       net.IP
	ipPool   string
	sharable bool
}

// ServiceExternalIPController is responsible for synchronizing the Services that need external IPs.
type ServiceExternalIPController struct {
	externalIPAllocator externalippool.ExternalIPAllocator
	client              clientset.Interface

	// ipAllocations caches the IP and the IP Pool which allocates it for each Service.
	ipAllocations     cache.Indexer
	ipAllocationMutex sync.RWMutex

	serviceInformer     cache.SharedIndexInformer
	serviceLister       corelisters.ServiceLister
	serviceListerSynced cache.InformerSynced
	// queue maintains the Service objects that need to be synced.
	queue workqueue.TypedRateLimitingInterface[apimachinerytypes.NamespacedName]
}

func NewServiceExternalIPController(
	client clientset.Interface,
	serviceInformer coreinformers.ServiceInformer,
	externalIPAllocator externalippool.ExternalIPAllocator,
) *ServiceExternalIPController {
	_ = "STUB: not implemented"
	return nil
}

func (c *ServiceExternalIPController) enqueueService(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

// enqueueServicesByExternalIPPool enqueues all Services that refer to the provided ExternalIPPool.
// The ExternalIPPool is affected by a Node update/create/delete event or ExternalIPPool changes.
func (c *ServiceExternalIPController) enqueueServicesByExternalIPPool(eipName string) {
	_ = "STUB: not implemented"
	return
}

// enqueueServicesWithoutIPs enqueues all Services that refer to the provided ExternalIPPool and have empty
// LoadBalancerIP in LoadBalancerStatus.
func (c *ServiceExternalIPController) enqueueServicesWithoutIPs(eipName string) {
	_ = "STUB: not implemented"
	return
}

// Run will create defaultWorkers workers (go routines) which will process the Service events from the
// workqueue.
func (c *ServiceExternalIPController) Run(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// restoreIPAllocations restores the existing external IPs of Services and records the successful ones in ipAllocations.
func (c *ServiceExternalIPController) restoreIPAllocations(services []*corev1.Service) {
	_ = "STUB: not implemented"
	// requestedIPAllocations contains IPAllocations we are going to request based on the Service statuses.
	return
}

// knownIPsByPool is used to deduplicate IPAllocations.

// eligibleServices contains Services which had a LoadBalancerIP assigned previously.

// Another Service already tried to restore the IP allocation, and restoring an IP more than once will fail for
// sure. Therefore, we use a set to ensure we restore each IP once. If the IP is restored successfully, all
// Services that had it assigned can continue using it.

// We don't set ObjectReference here as it might be shared between multiple Services.

// Convert the succeeded IPAllocations for ease of querying.

func (c *ServiceExternalIPController) updateIPAllocation(name apimachinerytypes.NamespacedName, ipPool string, ip net.IP, sharable bool) {
	_ = "STUB: not implemented"
	return
}

// Update other Services if the Service changes from unsharable to sharable.

func (c *ServiceExternalIPController) addIPAllocationLocked(name apimachinerytypes.NamespacedName, ipPool string, ip net.IP, sharable bool) {
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

func (c *ServiceExternalIPController) releaseExternalIP(service apimachinerytypes.NamespacedName) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// The IP is exclusively used by the Service.

// Ignore the error since the external IP Pool could be deleted.

// Update other Services as they may be affected by the deletion of this allocation.
// It's necessary even when the IP was not exclusively used by the deleted Service, considering this case:
// 1. Service A requests to share an IP with Service B, and both get the IP assigned.
// 2. Service A changes to not allow shared IP, currently the IP is still shared between A and B, but Service C
//    cannot get the IP assigned even if it allows shared IP.
// 3. Service A is deleted, the only remaining owner of the IP, Service B, allows shared IP, so Service C is
//    eligible for the IP.

func (c *ServiceExternalIPController) allocateExternalIP(service apimachinerytypes.NamespacedName, pool string, requestedIP string, allowSharedIP bool) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// Allocate IP from ExternalIPPool.

// Check whether the requested IP is in the IP pool or not.

// Check whether the requested IP is already used.

// Fail if the Service itself doesn't allow shared IP.

// Fail if any Service already using the IP doesn't allow shared IP.

// The requested IP is not used yet, allocate it.

func (c *ServiceExternalIPController) getExternalIPAllocation(service apimachinerytypes.NamespacedName) (*ipAllocation, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func getServiceExternalIP(service *corev1.Service) string { _ = "STUB: not implemented"; return "" }

func getServiceExternalIPPool(service *corev1.Service) string { _ = "STUB: not implemented"; return "" }

func isServiceExternalIPSharable(service *corev1.Service) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ServiceExternalIPController) syncService(key apimachinerytypes.NamespacedName) error {
	_ = "STUB: not implemented"
	return nil
}

// Service already deleted

// Service does not need external IP or type has changed.

// Ensure the LoadBalancerStatus in Kubernetes API is unset if the IP was allocated by it.

// If user specifies external IP in spec, we should check whether it matches the current external IP.

// Only the sharable annotation changes, update it.

// Ensure the LoadBalancerStatus in Kubernetes API matches the cache.

// The ExternalIPPool does not exist or has been deleted. Reclaim the external IP.

// The external IP or ExternalIPPool changes. Delete the previous allocation.

// updateService updates the Service status in Kubernetes API.
func (c *ServiceExternalIPController) updateServiceLoadBalancerIP(svc *corev1.Service, ip net.IP) error {
	_ = "STUB: not implemented"
	return nil
}
