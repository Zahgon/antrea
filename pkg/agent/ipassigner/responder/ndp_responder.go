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

package responder

import (
	"net"
	"net/netip"
	"sync"

	"antrea.io/ndp"
	"golang.org/x/net/ipv6"
	"k8s.io/apimachinery/pkg/util/sets"
)

var (
	solicitedNodeMulticastAddressPrefix = netip.MustParseAddr("ff02::1:ff00:0")
)

type ndpConn interface {
	WriteTo(message ndp.Message, cm *ipv6.ControlMessage, dstIP netip.Addr) error
	ReadFrom() (ndp.Message, *ipv6.ControlMessage, netip.Addr, error)
	JoinGroup(netip.Addr) error
	LeaveGroup(netip.Addr) error
	Close() error
}

type ndpResponder struct {
	once            sync.Once
	linkName        string
	conn            ndpConn
	linkEventCh     chan struct{}
	assignedIPs     sets.Set[netip.Addr]
	multicastGroups map[netip.Addr]int
	mutex           sync.Mutex
}

var _ Responder = (*ndpResponder)(nil)

func parseIPv6SolicitedNodeMulticastAddress(ip netip.Addr) netip.Addr {
	_ = "STUB: not implemented"
	return *new(netip.Addr)
}

// copy lower 24 bits

func (r *ndpResponder) InterfaceName() string { _ = "STUB: not implemented"; return "" }

func (r *ndpResponder) handleNeighborSolicitation(conn ndpConn, link *net.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ndpResponder) Run(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// The responder instance is created by the factory and can be shared by multiple callers.
	// Using once.Do here ensures it is started only once.
	return
}

func (r *ndpResponder) dialAndHandleRequests(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// It may take time for the interface to be ready for socket binding. For example, IPv6 introduces Duplicate Address Detection,
// which may take time to allow the address to be used for socket binding. EADDRNOTAVAIL (bind: cannot assign requested address)
// may be returned for such cases.

func (r *ndpResponder) AddIP(ip netip.Addr) error { _ = "STUB: not implemented"; return nil }

func (r *ndpResponder) joinMulticastGroup(ip netip.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ndpResponder) leaveMulticastGroup(ip netip.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ndpResponder) RemoveIP(ip netip.Addr) error { _ = "STUB: not implemented"; return nil }

func (r *ndpResponder) isIPAssigned(ip netip.Addr) bool { _ = "STUB: not implemented"; return false }

func (r *ndpResponder) onLinkUpdate(linkName string) { _ = "STUB: not implemented"; return }

// if an event is already present in the channel, we can drop this new one as we only monitor one link
