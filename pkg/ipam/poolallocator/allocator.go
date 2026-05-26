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

package poolallocator

import (
	"errors"
	"net"
	"time"

	"antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	crdclientset "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	informers "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/ipam/ipallocator"

	"k8s.io/apimachinery/pkg/util/wait"
	utilnet "k8s.io/utils/net"
)

// ErrPoolExhausted is returned when an IPPool has no available IPs left.
var ErrPoolExhausted = errors.New("pool exhausted")

// ipPoolStatusRetry backs off longer than retry.DefaultRetry (5 fixed ~10ms
// sleeps) but avoids an overly long tail. client-go's retry.DefaultBackoff uses
// fewer steps with a larger factor for a sub-second total; here we use a mild
// exponential (1.5) with more attempts than DefaultRetry so many Nodes updating
// the same IPPool status (shared pools, multi-NIC) can clear conflicts without
// multi-second CNI stalls.
var ipPoolStatusRetry = wait.Backoff{
	Steps:    8,
	Duration: 10 * time.Millisecond,
	Factor:   1.5,
	Jitter:   0.1,
}

// IPPoolAllocator is responsible for allocating IPs from IP set defined in IPPool CRD.
// The will update CRD usage accordingly.
// Pool Allocator assumes that pool with allocated IPs can not be deleted. Pool ranges can
// only be extended.
type IPPoolAllocator struct {
	// IP version of the IPPool
	IPVersion utilnet.IPFamily

	// Name of IPPool custom resource
	ipPoolName string

	// crd client to update the pool
	crdClient crdclientset.Interface

	// pool lister for reading the pool
	ipPoolLister informers.IPPoolLister
}

// NewIPPoolAllocator creates an IPPoolAllocator based on the provided IP pool.
func NewIPPoolAllocator(poolName string, client crdclientset.Interface, poolLister informers.IPPoolLister) (*IPPoolAllocator, error) {
	_ = "STUB: not implemented"
	// Validate the pool exists.
	return nil, nil
}

func (a *IPPoolAllocator) getPool() (*v1beta1.IPPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initAllocatorList reads IP Pool status and initializes a list of allocators based on
// IP Pool spec and state of allocation recorded in the status
func (a *IPPoolAllocator) initIPAllocators(ipPool *v1beta1.IPPool) (ipallocator.MultiIPAllocator, error) {
	_ = "STUB: not implemented"
	return *new(ipallocator.MultiIPAllocator), nil
}

// Initialize a list of IP allocators based on pool spec

// Reserve gateway address and broadcast address

// Allocation CIDR covers entire subnet, thus we need
// to reserve broadcast IP as well for IPv4

// Mark allocated IPs from pool status as unavailable

// TODO - fix state if possible

func (a *IPPoolAllocator) getPoolAndInitIPAllocators() (*v1beta1.IPPool, ipallocator.MultiIPAllocator, error) {
	_ = "STUB: not implemented"
	return nil, *new(ipallocator.MultiIPAllocator), nil
}

func (a *IPPoolAllocator) appendPoolUsage(ipPool *v1beta1.IPPool, ip net.IP, state v1beta1.IPAddressPhase, owner v1beta1.IPAddressOwner) error {
	_ = "STUB: not implemented"
	return nil
}

// updateIPAddressState updates the status of the specified IP in the provided IPPool. It requires the IP is already in the IPAddresses list of the IPPool's status.
func (a *IPPoolAllocator) updateIPAddressState(ipPool *v1beta1.IPPool, ip net.IP, state v1beta1.IPAddressPhase, owner v1beta1.IPAddressOwner) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *IPPoolAllocator) appendPoolUsageForStatefulSet(ipPool *v1beta1.IPPool, ips []net.IP, namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// removeIPAddressState updates ipPool status to delete released IP allocation, and keeps preallocation information
func (a *IPPoolAllocator) removeIPAddressState(ipPool *v1beta1.IPPool, ip net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// getExistingAllocation looks up the existing IP allocation for a Pod network interface, and
// returns the IP address and SubnetInfo if found.
func (a *IPPoolAllocator) getExistingAllocation(podOwner *v1beta1.PodOwner) (net.IP, *v1beta1.SubnetInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil, nil
}

// AllocateIP allocates the specified IP. It returns error if the IP is not in the range or already
// allocated, or in case CRD failed to update its state.
// In case of success, IP pool CRD status is updated with allocated IP/state/resource/container.
// AllocateIP returns subnet details for the requested IP, as defined in IP pool spec.
func (a *IPPoolAllocator) AllocateIP(ip net.IP, state v1beta1.IPAddressPhase, owner v1beta1.IPAddressOwner) (*v1beta1.SubnetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retry on CRD update conflict which is caused by multiple agents updating a pool at same time.

// Failed to find matching range

// AllocateNext allocates the next available IP. It returns error if pool is exhausted,
// or in case CRD failed to update its state.
// In case of success, IPPool CRD status is updated with allocated IP/state/resource/container.
// AllocateIP returns subnet details for the requested IP, as defined in IP pool spec.
func (a *IPPoolAllocator) AllocateNext(state v1beta1.IPAddressPhase, owner v1beta1.IPAddressOwner) (net.IP, *v1beta1.SubnetInfo, error) {
	_ = "STUB: not implemented"
	return *

	// Same resource can not ask for allocation twice without release.
	// This needs to be verified even at the expense of another API call.
	new(net.IP), nil, nil
}

// This can happen when the container requests IPs from multiple pools, and after an
// allocation failure, not all allocated IPs were successfully released, and then
// CNI ADD is retried.

// Retry on CRD update conflict which is caused by multiple agents updating a pool at same time.

// successful allocation

// AllocateReservedOrNext allocates the reserved IP if it exists, else allocates next available IP.
// It returns error if pool is exhausted, or in case it fails to update IPPool's state. In case of
// success, IP pool status is updated with allocated IP/state/resource/container.
// AllocateReservedOrNext returns subnet details for the requested IP, as defined in IP pool spec.
func (a *IPPoolAllocator) AllocateReservedOrNext(state v1beta1.IPAddressPhase, owner v1beta1.IPAddressOwner) (net.IP, *v1beta1.SubnetInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil, nil
}

// IP is not reserved, allocate next available IP.

// Retry on CRD update conflict which is caused by multiple agents updating a pool at same time.

// Failed to find matching range

// AllocateStatefulSet pre-allocates continuous range of IPs for StatefulSet.
// This functionality is useful when StatefulSet does not have a dedicated IP Pool assigned.
// It returns error if such range is not available. In this case IPs for the StatefulSet will
// be allocated on the fly, and there is no guarantee for continuous IPs.
func (a *IPPoolAllocator) AllocateStatefulSet(namespace, name string, size int, ip net.IP) error {
	_ = "STUB: not implemented"
	// Retry on CRD update conflict which is caused by multiple agents updating a pool at same time.
	return nil
}

// Make sure there is no double allocation for this StatefulSet

// Release releases the provided IP. It returns error if the IP is not in the range or not allocated,
// or in case CRD failed to update its state.
// In case of success, IP pool CRD status is updated with released IP/state/resource.
func (a *IPPoolAllocator) Release(ip net.IP) error {
	_ = "STUB: not implemented"

	// Retry on CRD update conflict which is caused by multiple agents updating a pool at same time.
	return nil
}

// Failed to find matching range

// ReleaseStatefulSet releases all IPs associated with specified StatefulSet. It returns error
// in case CRD failed to update its state.
// In case of success, IP pool CRD status is updated with released entries.
func (a *IPPoolAllocator) ReleaseStatefulSet(namespace, name string) error {
	_ = "STUB: not implemented"

	// Retry on CRD update conflict which is caused by multiple agents updating a pool at same time.
	return nil
}

// no change

// ReleaseContainer releases the IP associated with the specified container ID and interface name,
// and updates the IPPool CR status.
// If no IP is allocated to the Pod according to the IPPool CR status, the func just returns with no
// change.
func (a *IPPoolAllocator) ReleaseContainer(containerID, ifName string) error {
	_ = "STUB: not implemented"
	// Retry on CRD update conflict which is caused by multiple agents updating a pool at same time.
	return nil
}

// Mark the released IPs as available in the IPPool status.

// hasPod checks whether an IP was associated with specified pod. It returns the error if fails to
// retrieve the IPPool CR.
func (a *IPPoolAllocator) hasPod(namespace, podName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetContainerIP returns the IP allocated for the container interface if found.
func (a *IPPoolAllocator) GetContainerIP(containerID, ifName string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// getReservedIP checks whether an IP was reserved with specified owner. It returns error if the resource crd fails to be retrieved.
func (a *IPPoolAllocator) getReservedIP(reservedOwner v1beta1.IPAddressOwner) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// Name returns the name of the IPPool managed by this allocator.
func (a *IPPoolAllocator) Name() string { _ = "STUB: not implemented"; return "" }

func (a *IPPoolAllocator) Total() int { _ = "STUB: not implemented"; return 0 }

func (a *IPPoolAllocator) updateUsage(ipPool *v1beta1.IPPool) { _ = "STUB: not implemented"; return }
