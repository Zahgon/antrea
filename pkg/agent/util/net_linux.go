//go:build linux
// +build linux

// Copyright 2019 Antrea Authors
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

package util

import (
	"net"

	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"

	utilnetlink "antrea.io/antrea/v2/pkg/agent/util/netlink"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
)

var (
	netlinkUtil utilnetlink.Interface = &netlink.Handle{}

	// Declared variables which are meant to be overridden for testing.
	getNS              = ns.GetNS
	netNSDo            = ns.NetNS.Do
	netNSPath          = ns.NetNS.Path
	netNSClose         = ns.NetNS.Close
	getVethPeerIfindex = ip.GetVethPeerIfindex
	netlinkAttrs       = netlink.Link.Attrs
)

// GetNSPeerDevBridge returns peer device and its attached bridge (if applicable)
// for device dev in network space indicated by nsPath
func GetNSPeerDevBridge(nsPath, dev string) (*net.Interface, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// not attached to a bridge.

// master link is not bridge

// GetNSDevInterface returns interface of dev in namespace nsPath.
func GetNSDevInterface(nsPath, dev string) (*net.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNSPath returns the path of the specified netns.
func GetNSPath(netnsName string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func SetLinkUp(name string) (net.HardwareAddr, int, error) {
	_ = "STUB: not implemented"
	return *new(net.HardwareAddr), 0, nil
}

// Set host gateway interface up.

// addrSliceDifference returns elements in s1 but not in s2.
func addrSliceDifference(s1, s2 []netlink.Addr) []*netlink.Addr {
	_ = "STUB: not implemented"
	return nil
}

// ConfigureLinkAddresses adds the provided addresses to the interface identified by index idx, if
// they are missing from the interface. Any other existing address already configured for the
// interface will be removed, unless it is a link-local address.
func ConfigureLinkAddresses(idx int, ipNets []*net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove link-local address from list

// ListenLocalSocket creates a listener on a Unix domain socket.
func ListenLocalSocket(address string) (net.Listener, error) {
	_ = "STUB: not implemented"
	// remove before bind to avoid "address already in use" errors
	return *new(net.Listener), nil
}

// SetAdapterMACAddress set specified MAC address on interface.
func SetAdapterMACAddress(adapterName string, macConfig *net.HardwareAddr) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteOVSPort deletes specific OVS port. This function calls ovs-vsctl command to bypass
// OVS bridge client to work when agent exiting.
func deleteOVSPort(brName, portName string) error { _ = "STUB: not implemented"; return nil }

func HostInterfaceExists(ifName string) bool { _ = "STUB: not implemented"; return false }

func GetInterfaceConfig(ifName string) (*net.Interface, []*net.IPNet, []interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func RenameInterface(from, to string) error { _ = "STUB: not implemented"; return nil }

// Fix for the issue https://github.com/antrea-io/antrea/issues/6301.
// In some new Linux versions which support AltName, if the only valid altname of the interface is the same as the
// interface name, it would be left empty when the name is occupied by the interface name; after we rename the
// interface name to another value, the altname of the interface would be set to the original interface name by the
// system.
// This altname must be removed as we need to reserve the name for an OVS internal port.

func RemoveLinkIPs(link netlink.Link) error { _ = "STUB: not implemented"; return nil }

func RemoveLinkRoutes(link netlink.Link) error { _ = "STUB: not implemented"; return nil }

func ConfigureLinkRoutes(link netlink.Link, routes []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Route already exists

func EnsureIPv6EnabledOnInterface(ifaceName string) error { _ = "STUB: not implemented"; return nil }

func EnsureARPAnnounceOnInterface(ifaceName string, value int) error {
	_ = "STUB: not implemented"
	return nil
}

func EnsureRPFilterOnInterface(ifaceName string, value int) error {
	_ = "STUB: not implemented"
	return nil
}

func EnsurePromoteSecondariesOnInterface(ifaceName string) error {
	_ = "STUB: not implemented"
	return nil
}

func getRoutesOnInterface(linkIndex int) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func renameHostInterface(oriName string, newName string) error {
	_ = "STUB: not implemented"
	return nil
}

// removeInterfaceAltName removes altName on interface with provided name. altName not found will return nil.
func removeInterfaceAltName(name string, altName string) error {
	_ = "STUB: not implemented"
	return nil
}

// PrepareHostInterfaceConnection prepares host interface connection to the OVS bridge client by:
// 1. Renaming the host interface (a bridged suffix will be added to it).
// 2. Creating an internal port (original name of the host interface will be used here).
// 3. Set the MTU of this new link/internal-port to the provided mtu parameter value, unless mtu is zero.
// 4. Moving IPs of host interface to this new link/internal-port.
// 5. Moving routes of host interface to the new link/internal-port.
// and returns the bridged name, true if it already exists, and error.
func PrepareHostInterfaceConnection(
	bridge ovsconfig.OVSBridgeClient,
	ifaceName string,
	ifaceOFPort int32,
	externalIDs map[string]interface{},
	mtu int,
) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// If the port already exists, just return.

// Wait a few seconds for OVS bridge local port.

// Check if interface is configured with an IPv6 address: if it is, we need to ensure that IPv6
// is enabled on the OVS internal port as we need to move all IP addresses over.

// Restore the host routes which are lost when moving the network configuration of the
// host interface to OVS bridge interface.

// RestoreHostInterfaceConfiguration restore the configuration from bridge back to host interface, reverting the
// actions taken in PrepareHostInterfaceConnection.
func RestoreHostInterfaceConfiguration(brName string, interfaceName string) {
	_ = "STUB: not implemented"
	return
}

// restore only when interface eth0~ exists

// get interface config

// delete internal port (eth0)

// remove host interface (eth0~) from bridge

// rename host interface(eth0~ -> eth0)

// restore IPs to eth0

// restore routes to eth0
