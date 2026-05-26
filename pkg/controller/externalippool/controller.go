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

package externalippool

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	antreacrds "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	clientset "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	antreainformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
	antrealisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/ipam/ipallocator"
)

const (
	controllerName = "ExternalIPPoolController"
	// Set resyncPeriod to 0 to disable resyncing.
	resyncPeriod time.Duration = 0
	// How long to wait before retrying the processing of an ExternalIPPool change.
	minRetryDelay = 5 * time.Second
	maxRetryDelay = 300 * time.Second
	// Default number of workers processing an ExternalIPPool change.
	defaultWorkers = 4
)

var (
	ErrExternalIPPoolNotFound = errors.New("ExternalIPPool not found")
)

// IPAllocation contains the IP and the IP Pool which allocates it.
type IPAllocation struct {
	// ObjectReference is useful to track the owner of this IP allocation.
	ObjectReference corev1.ObjectReference
	// IPPoolName is the name of the IP pool.
	IPPoolName string
	// IP is the allocated IP.
	IP net.IP
}

// ExternalIPPoolEventHandler defines a consumer to subscribe for external ExternalIPPool events.
type ExternalIPPoolEventHandler func(externalIPPool string)

type ExternalIPAllocator interface {
	// AddEventHandler adds a consumer for ExternalIPPool events. It will block other consumers from allocating new IPs by
	// AllocateIPFromPool() until it calls RestoreIPAllocations().
	AddEventHandler(handler ExternalIPPoolEventHandler)
	// RestoreIPAllocations is used to restore the previous allocated IPs after controller restarts. It will return the
	// succeeded IP Allocations.
	RestoreIPAllocations(allocations []IPAllocation) []IPAllocation
	// AllocateIPFromPool allocates an IP from the given IP pool.
	AllocateIPFromPool(externalIPPool string) (net.IP, error)
	// IPPoolExists checks whether the IP pool exists.
	IPPoolExists(externalIPPool string) bool
	// IPPoolHasIP checks whether the IP pool contains the given IP.
	IPPoolHasIP(externalIPPool string, ip net.IP) bool
	// UpdateIPAllocation marks the IP in the specified ExternalIPPool as occupied.
	UpdateIPAllocation(externalIPPool string, ip net.IP) error
	// ReleaseIP releases the IP to the IP pool.
	// It returns ErrExternalIPPoolNotFound if the externalIPPool does not exist.
	// Any other error indicates that the IP was not allocated, or is not currently allocated.
	// In case of an error, there is no reason to try again with the same arguments, as
	// transient errors are not possible.
	ReleaseIP(externalIPPool string, ip net.IP) error
	// HasSynced indicates ExternalIPAllocator has finished syncing all ExternalIPPool resources.
	HasSynced() bool
}

var _ ExternalIPAllocator = (*ExternalIPPoolController)(nil)

// ExternalIPPoolController is responsible for synchronizing the ExternalIPPool resources.
type ExternalIPPoolController struct {
	crdClient                  clientset.Interface
	externalIPPoolLister       antrealisters.ExternalIPPoolLister
	externalIPPoolListerSynced cache.InformerSynced

	// ipAllocatorMap is a map from ExternalIPPool name to MultiIPAllocator.
	ipAllocatorMap   map[string]ipallocator.MultiIPAllocator
	ipAllocatorMutex sync.RWMutex

	// ipAllocatorInitialized stores a boolean value, which tracks if the ipAllocatorMap has been initialized
	// with the full list of ExternalIPPool.
	ipAllocatorInitialized *atomic.Value

	// handlers is an array of handlers will be notified when ExternalIPPool updates.
	handlers          []ExternalIPPoolEventHandler
	handlersWaitGroup sync.WaitGroup

	// queue maintains the ExternalIPPool objects that need to be synced.
	queue workqueue.TypedRateLimitingInterface[string]
}

// NewExternalIPPoolController returns a new *ExternalIPPoolController.
func NewExternalIPPoolController(crdClient clientset.Interface, externalIPPoolInformer antreainformers.ExternalIPPoolInformer) *ExternalIPPoolController {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExternalIPPoolController) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (c *ExternalIPPoolController) AddEventHandler(handler ExternalIPPoolEventHandler) {
	_ = "STUB: not implemented"
	return
}

func (c *ExternalIPPoolController) RestoreIPAllocations(allocations []IPAllocation) []IPAllocation {
	_ = "STUB: not implemented"
	return nil
}

// Run begins watching and syncing of the ExternalIPPoolController.
func (c *ExternalIPPoolController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Initialize the ipAllocatorMap with the existing ExternalIPPools.

// createOrUpdateIPAllocator creates or updates the IP allocator based on the provided ExternalIPPool.
// Currently it's assumed that only new ranges will be added and existing ranges should not be deleted.
// TODO: Use validation webhook to ensure it.
func (c *ExternalIPPoolController) createOrUpdateIPAllocator(ipPool *antreacrds.ExternalIPPool) bool {
	_ = "STUB: not implemented"
	return false
}

// Must use normalized IPNet string to check if the IP range exists. Otherwise non-strict CIDR like
// 192.168.0.1/24 will be considered new even if it doesn't change.
// Validating or normalizing the input CIDR should be a better solution but the externalIPPools that
// have been created will still have this issue, so we just normalize the CIDR when using it.

// Don't use the IPv4 network's broadcast address.

// The IP range already exists in multiIPAllocator.

// deleteIPAllocator deletes the IP allocator of the given IP pool.
func (c *ExternalIPPoolController) deleteIPAllocator(poolName string) {
	_ = "STUB: not implemented"
	return
}

// getIPAllocator gets the IP allocator of the given IP pool.
func (c *ExternalIPPoolController) getIPAllocator(poolName string) (ipallocator.MultiIPAllocator, bool) {
	_ = "STUB: not implemented"
	return *new(ipallocator.MultiIPAllocator), false
}

// AllocateIPFromPool allocates an IP from the the given IP pool.
func (c *ExternalIPPoolController) AllocateIPFromPool(ipPoolName string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// UpdateIPAllocation sets the IP in the specified ExternalIPPool.
func (c *ExternalIPPoolController) UpdateIPAllocation(poolName string, ip net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExternalIPPoolController) updateExternalIPPoolStatus(poolName string) error {
	_ = "STUB: not implemented"
	return nil
}

// ReleaseIP releases the IP to the pool.
func (c *ExternalIPPoolController) ReleaseIP(poolName string, ip net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ExternalIPPoolController) IPPoolHasIP(poolName string, ip net.IP) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ExternalIPPoolController) IPPoolExists(pool string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ExternalIPPoolController) worker() { _ = "STUB: not implemented"; return }

func (c *ExternalIPPoolController) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// Put the item back in the workqueue to handle any transient errors.

// If no error occurs we Forget this item so it does not get queued again until
// another change happens.

// addExternalIPPool processes ExternalIPPool ADD events. It creates an IPAllocator for the pool and triggers
// reconciliation of consumers that refer to the pool.
func (c *ExternalIPPoolController) addExternalIPPool(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

// updateExternalIPPool processes ExternalIPPool UPDATE events. It updates the IPAllocator for the pool and triggers
// reconciliation of consumers that refer to the pool if the IPAllocator changes.
func (c *ExternalIPPoolController) updateExternalIPPool(_, cur interface{}) {
	_ = "STUB: not implemented"
	return
}

// deleteExternalIPPool processes ExternalIPPool DELETE events. It deletes the IPAllocator for the pool and triggers
// reconciliation of all consumers that refer to the pool.
func (c *ExternalIPPoolController) deleteExternalIPPool(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

// Call consumers to reclaim the IPs allocated from the pool.
