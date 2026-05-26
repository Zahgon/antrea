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

	"antrea.io/arp"
	"k8s.io/apimachinery/pkg/util/sets"
)

type arpResponder struct {
	once        sync.Once
	linkName    string
	assignedIPs sets.Set[netip.Addr]
	mutex       sync.Mutex
	linkEventCh chan struct{}
}

var _ Responder = (*arpResponder)(nil)

func (r *arpResponder) InterfaceName() string { _ = "STUB: not implemented"; return "" }

func (r *arpResponder) AddIP(ip netip.Addr) error { _ = "STUB: not implemented"; return nil }

func (r *arpResponder) RemoveIP(ip netip.Addr) error { _ = "STUB: not implemented"; return nil }

func (r *arpResponder) handleARPRequest(client *arp.Client, iface *net.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *arpResponder) Run(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// The responder instance is created by the factory and can be shared by multiple callers.
	// Using once.Do here ensures it is started only once.
	return
}

func (r *arpResponder) dialAndHandleRequests(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (r *arpResponder) isIPAssigned(ip netip.Addr) bool { _ = "STUB: not implemented"; return false }

func (r *arpResponder) deleteIP(ip netip.Addr) bool { _ = "STUB: not implemented"; return false }

func (r *arpResponder) addIP(ip netip.Addr) bool { _ = "STUB: not implemented"; return false }

func (r *arpResponder) onLinkUpdate(linkName string) { _ = "STUB: not implemented"; return }

// if an event is already present in the channel, we can drop this new one as we only monitor one link
