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

package ipassigner

import (
	"net"
	"sync"

	"github.com/vishvananda/netlink"
	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/agent/ipassigner/linkmonitor"
	"antrea.io/antrea/v2/pkg/agent/ipassigner/responder"
	crdv1b1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

// VLAN interfaces created by antrea-agent will be named with the prefix.
// For example, when VLAN ID is 10, the name will be antrea-ext.10.
// It can be used to determine whether it's safe to delete an interface when it's no longer used.
const vlanInterfacePrefix = "antrea-ext."

// assignee is the unit that IPs are assigned to. All IPs from the same VLAN share an assignee.
type assignee struct {
	// logicalInterface is the interface IPs should be logically assigned to. It's also used for IP advertisement.
	// The field must not be nil.
	logicalInterface *net.Interface
	// link is used for IP link management and IP address add/del operation. The field can be nil if IPs don't need to
	// be assigned to an interface physically.
	link netlink.Link
	// arpResponder is used for ARP responder for IPv4 address. The field should be nil if the interface can respond to
	// ARP queries itself.
	arpResponder responder.Responder
	// ndpResponder is used for NDP responder for IPv6 address. The field should be nil if the interface can respond to
	// NDP queries itself.
	ndpResponder responder.Responder
	// ips tracks IPs that have been assigned to this assignee.
	ips sets.Set[string]
}

// deletable returns whether this assignee can be safely deleted.
func (as *assignee) deletable() bool { _ = "STUB: not implemented"; return false }

// It never has a real link.

// Do not delete non VLAN interfaces.

// Do not delete VLAN interfaces not created by antrea-agent.

func (as *assignee) destroy() error { _ = "STUB: not implemented"; return nil }

func (as *assignee) assign(ip net.IP, subnetInfo *crdv1b1.SubnetInfo) error {
	_ = "STUB: not implemented"
	// If there is a real link, add the IP to its address list.
	return nil
}

// Always advertise the IP when the IP is newly assigned to this Node.

func (as *assignee) advertise(ip net.IP) { _ = "STUB: not implemented"; return }

func (as *assignee) unassign(ip net.IP, subnetInfo *crdv1b1.SubnetInfo) error {
	_ = "STUB: not implemented"
	// If there is a real link, delete the IP from its address list.
	return nil
}

func (as *assignee) getVLANID() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (as *assignee) loadIPAddresses() (map[string]*crdv1b1.SubnetInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only include global unicast addresses, otherwise addresses like link local ones may be mistakenly deleted.

// subnetInfo should be nil for the dummy interface.

// ipAssigner creates dummy/vlan devices and assigns IPs to them.
// It's supposed to be used in the cases that external IPs should be configured on the system so that they can be used
// for SNAT (egress scenario) or DNAT (ingress scenario).
// By default, a dummy device is used because the IPs just need to be present in any device to be functional, and using
// dummy device avoids touching system managed devices and is easy to know IPs that are assigned by antrea-agent.
// If an IP is associated with a VLAN ID, it will be assigned to a vlan device which is a sub-interface of the external
// device for proper VLAN tagging and untagging.
type ipAssigner struct {
	// externalInterface is the device that GARP (IPv4) and Unsolicited NA (IPv6) will eventually be sent from.
	externalInterface *net.Interface
	// defaultAssignee is the assignee that IPs without VLAN tag will be assigned to.
	defaultAssignee *assignee
	// vlanAssignees contains the vlan-based assignees that IPs with VLAN tag will be assigned to, keyed by VLAN ID.
	vlanAssignees map[int32]*assignee
	// assignIPs caches the IPs that have been assigned.
	// TODO: Add a goroutine to ensure that the cache is in sync with the IPs assigned to the dummy device in case the
	// IPs are removed by users accidentally.
	assignedIPs map[string]*crdv1b1.SubnetInfo
	mutex       sync.RWMutex
	// uniqueMACForSubInterfaces indicates whether to assign a unique MAC address to VLAN sub-interfaces.
	uniqueMACForSubInterfaces bool
}

// NewIPAssigner returns an *ipAssigner.
func NewIPAssigner(nodeTransportInterface string, dummyDeviceName string, linkMonitor linkmonitor.Interface, uniqueMACForSubInterfaces bool) (IPAssigner, error) {
	_ = "STUB: not implemented"
	return *new(IPAssigner), nil
}

// For the Egress scenario, the external IPs should always be present on the dummy
// interface as they are used as tunnel endpoints. If arp_ignore is set to a value
// other than 0, the host will not reply to ARP requests received on the transport
// interface when the target IPs are assigned on the dummy interface. So a userspace
// ARP responder is needed to handle ARP requests for the Egress IPs.

// getVLANInterfaces returns VLAN sub-interfaces of the given parent interface
// that were created by antrea-agent (i.e., whose names start with vlanInterfacePrefix).
// User-managed VLAN sub-interfaces are ignored to avoid accidentally loading
// their IPs and deleting them as stale during reconciliation.
func getVLANInterfaces(parentIndex int) ([]*netlink.Vlan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getARPIgnoreForInterface gets the max value of conf/{all,interface}/arp_ignore form sysctl.
func getARPIgnoreForInterface(iface string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ensureDummyDevice creates the dummy device if it doesn't exist.
func ensureDummyDevice(deviceName string) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

// When the primary IP address is removed from the interface, promote a corresponding secondary IP address
// instead of removing all the corresponding secondary IP addresses. Otherwise, the deletion of one IP address
// can trigger the automatic removal of all other IP addresses in the same subnet, if the deleted IP happens to
// be the primary (first one assigned chronologically).

// loadIPAddresses gets the IP addresses on the default device and the vlan devices.
func (a *ipAssigner) loadIPAddresses() error {
	_ = "STUB: not implemented"
	// Load IPs assigned to the default interface.
	return nil
}

// Load IPs assigned to the vlan interfaces.

// AssignIP ensures the provided IP is assigned to the system and advertised to its neighbors.
//   - If subnetInfo is nil or the vlan is 0, the IP will be assigned to the default interface, and its advertisement
//     will be sent through the external interface.
//   - Otherwise, the IP will be assigned to a corresponding vlan sub-interface of the external interface, and its
//     advertisement will be sent through the vlan sub-interface (though via the external interface eventually).
func (a *ipAssigner) AssignIP(ip string, subnetInfo *crdv1b1.SubnetInfo, forceAdvertise bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ipAssigner doesn't care about the gateway.

// UnassignIP ensures the provided IP is not assigned to the dummy/vlan device.
func (a *ipAssigner) UnassignIP(ip string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (a *ipAssigner) unassign(ip net.IP, subnetInfo *crdv1b1.SubnetInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// The assignee doesn't exist, meaning the IP has been unassigned previously.

// AssignedIPs return the IPs that are assigned to the dummy device.
func (a *ipAssigner) AssignedIPs() map[string]*crdv1b1.SubnetInfo {
	_ = "STUB: not implemented"
	return nil
}

// Return a copy.

// InitIPs loads the IPs from the dummy/vlan devices and replaces the IPs that are assigned to it
// with the given ones. This function also adds the given IPs to the ARP/NDP responder if
// applicable. It can be used to recover the IP assigner to the desired state after Agent restarts.
// It's not thread-safe and should only be called once for initialization before calling other methods.
func (a *ipAssigner) InitIPs(desired map[string]*crdv1b1.SubnetInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *ipAssigner) GetInterfaceID(subnetInfo *crdv1b1.SubnetInfo) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// The assignee doesn't exist, meaning the IP has been unassigned previously.

// Run starts the ARP responder and NDP responder.
func (a *ipAssigner) Run(ch <-chan struct{}) { _ = "STUB: not implemented"; return }

// getAssignee gets or creates the vlan device for the subnet if it doesn't exist.
func (a *ipAssigner) getAssignee(subnetInfo *crdv1b1.SubnetInfo, createIfNotExist bool) (*assignee, error) {
	_ = "STUB: not implemented"
	// Use the default assignee if subnet info is nil or the vlan is not set.
	return nil, nil
}

func (a *ipAssigner) addVLANAssignee(link netlink.Link, vlan int32) (*assignee, error) {
	_ = "STUB: not implemented"
	return nil,

		// Loose mode is needed because incoming traffic received on the interface is expected to be received on the parent
		// external interface when looking up the main table. To make it look up the custom table, we will need to restore
		// the mark on the reply traffic and turn on src_valid_mark on this interface, which is more complicated.
		nil
}

// When the primary IP address is removed from the interface, promote a corresponding secondary IP address
// instead of removing all the corresponding secondary IP addresses. Otherwise, the deletion of one IP address
// can trigger the automatic removal of all other IP addresses in the same subnet, if the deleted IP happens to
// be the primary (first one assigned chronologically).

// VLAN interface can answer ARP/NDP directly, no need to create userspace responders.

func getIPNet(ip net.IP, subnetInfo *crdv1b1.SubnetInfo) *net.IPNet {
	_ = "STUB: not implemented"
	return nil
}
