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

package ipallocator

import (
	"math/big"
	"net"
	"sync"
)

type IPAllocator interface {
	AllocateIP(ip net.IP) error

	AllocateNext() (net.IP, error)

	// Allocate range of continuous IPs in one go. If continuus chunk is not
	// available, error will be returned.
	AllocateRange(size int) ([]net.IP, error)

	Release(ip net.IP) error

	Used() int

	Has(ip net.IP) bool
}

// SingleIPAllocator is responsible for allocating IPs from a contiguous IP range.
type SingleIPAllocator struct {
	// The string format of the IP range. e.g. 10.10.10.0/24, or 10.10.10.10-10.10.10.20.
	ipRangeStr string

	mutex sync.RWMutex
	// base is a cached version of the start IP in the CIDR range as a *big.Int.
	base *big.Int
	// max is the maximum size of the usable addresses in the range.
	max int
	// allocated is a bit array of the allocated items in the range.
	allocated *big.Int
	// count is the number of currently allocated elements in the range.
	count int
	// IPs inside the cidr not available for allocation
	reservedIPs []net.IP
}

// NewCIDRAllocator creates an IPAllocator based on the provided CIDR.
func NewCIDRAllocator(cidr *net.IPNet, reservedIPs []net.IP) (*SingleIPAllocator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start from "x.x.x.1".

// In case a big range occupies too much memory, allow at most 65536 IP for each IP range.

// NewIPRangeAllocator creates an IPAllocator based on the provided start IP and end IP.
// The start IP and end IP are inclusive.
func NewIPRangeAllocator(startIP, endIP net.IP) (*SingleIPAllocator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In case a big range occupies too much memory, allow at most 65536 IP for each ipset.

func (a *SingleIPAllocator) Name() string { _ = "STUB: not implemented"; return "" }

func (a *SingleIPAllocator) checkReserved(ip net.IP) error { _ = "STUB: not implemented"; return nil }

// AllocateIP allocates the specified IP. It returns error if the IP is not in the range or already allocated.
func (a *SingleIPAllocator) AllocateIP(ip net.IP) error { _ = "STUB: not implemented"; return nil }

func (a *SingleIPAllocator) allocateOffset(i int) (net.IP, bool) {
	_ = "STUB: not implemented"
	return *new(net.IP), false
}

// AllocateNext allocates an IP from the IP range. It returns error if no IP is available.
func (a *SingleIPAllocator) AllocateNext() (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// we should never reach here

// AllocateRange allocates continuous range of specified size. If not available, error is returned.
func (a *SingleIPAllocator) AllocateRange(size int) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if this continuous range is available

// perform the actual allocation

func (a *SingleIPAllocator) getOffset(ip net.IP) int { _ = "STUB: not implemented"; return 0 }

// Release releases the provided IP. It returns error if the IP is not in the range or not allocated.
func (a *SingleIPAllocator) Release(ip net.IP) error { _ = "STUB: not implemented"; return nil }

// Used returns the number of the allocated IPs.
func (a *SingleIPAllocator) Used() int { _ = "STUB: not implemented"; return 0 }

// Free returns the number of free IPs.
func (a *SingleIPAllocator) Free() int { _ = "STUB: not implemented"; return 0 }

// Total returns the number total of IPs within the pool.
func (a *SingleIPAllocator) Total() int { _ = "STUB: not implemented"; return 0 }

// Has returns whether the provided IP is in the range or not.
func (a *SingleIPAllocator) Has(ip net.IP) bool { _ = "STUB: not implemented"; return false }

// MultiIPAllocator is responsible for allocating IPs from multiple contiguous IP ranges.
type MultiIPAllocator []*SingleIPAllocator

func (ma MultiIPAllocator) Names() []string { _ = "STUB: not implemented"; return nil }

func (ma MultiIPAllocator) AllocateIP(ip net.IP) error { _ = "STUB: not implemented"; return nil }

func (ma MultiIPAllocator) AllocateNext() (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// AllocateRange allocates continuous range of specified size.
// If not available in any allocator, error is returned.
func (ma MultiIPAllocator) AllocateRange(size int) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ma MultiIPAllocator) Release(ip net.IP) error { _ = "STUB: not implemented"; return nil }

func (ma MultiIPAllocator) Used() int { _ = "STUB: not implemented"; return 0 }

func (ma MultiIPAllocator) Free() int { _ = "STUB: not implemented"; return 0 }

func (ma MultiIPAllocator) Total() int { _ = "STUB: not implemented"; return 0 }

func (ma MultiIPAllocator) Has(ip net.IP) bool { _ = "STUB: not implemented"; return false }
