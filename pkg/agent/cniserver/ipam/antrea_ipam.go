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

package ipam

import (
	"net"
	"sync"

	"github.com/containernetworking/cni/pkg/invoke"
	cnitypes "github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"

	"antrea.io/antrea/v2/pkg/agent/cniserver/types"
	crdv1b1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/ipam/poolallocator"
)

const (
	AntreaIPAMType = "antrea"
)

// Antrea IPAM driver would allocate IP addresses according to object IPAM annotation,
// if present. If annotation is not present, the driver will delegate functionality
// to traditional IPAM driver.
type AntreaIPAM struct {
	controller      *AntreaIPAMController
	controllerMutex sync.RWMutex
}

// Global variable is needed to work around order of initialization
// Controller will be assigned to the driver after it is initialized
// by agent init.
var antreaIPAMDriver *AntreaIPAM

type mineType uint8

const (
	mineUnknown mineType = iota
	mineFalse
	mineTrue
)

// Resource needs to be unique since it is used as identifier in Del.
// Therefore Container ID is used, while Pod/Namespace are shown for visibility.
func getAllocationPodOwner(args *invoke.Args, k8sArgs *types.K8sArgs, reservedOwner *crdv1b1.IPAddressOwner, secondary bool) *crdv1b1.PodOwner {
	_ = "STUB: not implemented"
	return nil
}

// Add interface name for secondary network to uniquely identify
// the secondary network interface.

func getAllocationOwner(args *invoke.Args, k8sArgs *types.K8sArgs, reservedOwner *crdv1b1.IPAddressOwner, secondary bool) *crdv1b1.IPAddressOwner {
	_ = "STUB: not implemented"
	return nil
}

// Helper to generate IP config and default route, taking IP version into account
func generateIPConfig(ip net.IP, prefixLength int, gwIP net.IP) (*current.IPConfig, *cnitypes.Route) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseStaticAddresses(ipamConfig *types.IPAMConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *AntreaIPAM) setController(controller *AntreaIPAMController) {
	_ = "STUB: not implemented"
	return
}

// splitIPsByFamily returns the first IPv4 and IPv6 address found in ips.
// Additional addresses of the same family are silently ignored.
func splitIPsByFamily(ips []net.IP) (v4, v6 net.IP) {
	_ = "STUB: not implemented"
	return *new(net.IP), *new(net.IP)
}

// Add allocates IP addresses from the associated IP Pools. It supports IPv4,
// IPv6, and dual-stack configurations. At most one IP is allocated per address
// family, even when multiple Pools exist for that family. The allocated IPs and
// associated resources will be stored in the IP Pool status.
//
// When multiple Pools of the same IP family are configured and no specific IP
// is requested, Add will try each Pool in order. If a Pool is exhausted (no
// free IPs), the next Pool of the same family is attempted. Other errors
// (e.g. API failures) cause an immediate return.
//
// When a Pod specifies desired IPs via the AntreaIPAMPodIP annotation, at most
// one IPv4 and one IPv6 address are used; additional addresses of the same
// family are silently ignored. The specified IP is always allocated from the
// first Pool of the corresponding IP family. If the allocation fails for any
// reason (IP not in range, already allocated, etc.), the error is returned
// immediately without trying subsequent Pools.
// See https://antrea.io/docs/main/docs/antrea-ipam.md for more details.
func (d *AntreaIPAM) Add(args *invoke.Args, k8sArgs *types.K8sArgs, networkConfig []byte) (bool, *IPAMResult, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// pass this request to next driver

// Release already allocated IPs on error.

// For a requested IP, we only attempt allocation from the first matching pool,
// and do not fall back to other pools if the allocation fails.

// At this point an IP should have been allocated: we already determined that this Pod matches
// at least one Antrea IPPool (otherwise owns / getPoolAllocatorsByPod would have returned an
// error), and any allocation failure should have been caught above.

// All allocations successful, clear the deferred release.

// Del releases the IP associated with the resource from the IP Pool status.
func (d *AntreaIPAM) Del(args *invoke.Args, k8sArgs *types.K8sArgs, networkConfig []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Let the invoker retry at error.

// If no allocation found, pass CNI DEL to the next driver.

// Check verifies the IP associated with the resource is tracked in the IP Pool status.
func (d *AntreaIPAM) Check(args *invoke.Args, k8sArgs *types.K8sArgs, networkConfig []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// pass this request to next driver

// SecondaryNetworkAllocate allocates IP addresses for a Pod secondary network interface, based on
// the IPAM configuration of the passed CNI network configuration.
// It supports IPAM for both Antrea-managed secondary networks and Multus-managed secondary
// networks.
func (d *AntreaIPAM) SecondaryNetworkAllocate(podOwner *crdv1b1.PodOwner, networkConfig *types.NetworkConfig) (*IPAMResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return error to let the invoker retry.

// Try to release the allocated IPs after an error.

// CNI spec 0.2.0 and below support only one v4 and one v6 address. But we
// assume the CNI version >= 0.3.0, and so do not check the number of
// addresses.

// No failed allocation, so do not release allocated IPs.

// Add static addresses.

// Copy routes and DNS from the input IPAM configuration.

// SecondaryNetworkRelease releases the IP addresses allocated for a Pod secondary network interface.
func (d *AntreaIPAM) SecondaryNetworkRelease(owner *crdv1b1.PodOwner) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *AntreaIPAM) secondaryNetworkAdd(args *invoke.Args, k8sArgs *types.K8sArgs, networkConfig *types.NetworkConfig) (*IPAMResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *AntreaIPAM) secondaryNetworkDel(args *invoke.Args, k8sArgs *types.K8sArgs, networkConfig *types.NetworkConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *AntreaIPAM) secondaryNetworkCheck(args *invoke.Args, k8sArgs *types.K8sArgs, networkConfig *types.NetworkConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *AntreaIPAM) del(podOwner *crdv1b1.PodOwner) (foundAllocation bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Return error to let the invoker retry.

// The Pod resource might have been removed; and for a secondary we
// would rely on the passed IPPool for CNI DEL. So, search IPPools with
// the matched PodOwner.

// Multiple allocators can be returned if the network interface has IPs
// allocated from more than one IPPools.

// owns checks whether this driver owns the coming IPAM request. This decision is based on Antrea
// IPAM annotation for the resource (Pod or Namespace). If an annotation is not present, or the
// annotated IP Pool not found, the driver should not own the request and will fall back to the next
// IPAM driver.
// return:
// mineUnknown + PodNotFound error
// mineUnknown + InvalidIPAnnotation error
// mineFalse + nil error
// mineTrue + timeout error
// mineTrue + IPPoolNotFound error
// mineTrue + nil error
func (d *AntreaIPAM) owns(k8sArgs *types.K8sArgs) (mineType, []*poolallocator.IPPoolAllocator, []net.IP, *crdv1b1.IPAddressOwner, error) {
	_ = "STUB: not implemented"
	// Wait controller ready to avoid inappropriate behaviors on the CNI request.
	return *new(mineType), nil, nil, nil, nil
}

// Return mineTrue to make this request fail and kubelet will retry.

func (d *AntreaIPAM) waitForControllerReady() error { _ = "STUB: not implemented"; return nil }

func init() {
	// Antrea driver must come first.
	// NOTE: this is global variable that requires follow-up setup post agent initialization.
	antreaIPAMDriver = &AntreaIPAM{}
	RegisterIPAMDriver(AntreaIPAMType, antreaIPAMDriver)

	// Host local plugin is fallback driver
	RegisterIPAMDriver(AntreaIPAMType, &IPAMDelegator{pluginType: ipamHostLocal})
}
