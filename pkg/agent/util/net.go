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

	// #nosec G505: not used for security purposes

	"net"
	"net/netip"

	"k8s.io/apimachinery/pkg/util/sets"

	utilip "antrea.io/antrea/v2/pkg/util/ip"
)

const (
	interfaceNameLength   = 15
	interfacePrefixLength = 8
	interfaceKeyLength    = interfaceNameLength - (interfacePrefixLength + 1)

	FamilyIPv4 uint8 = 4
	FamilyIPv6 uint8 = 6

	bridgedUplinkSuffix = "~"
)

var (
	// Declared variables which are meant to be overridden for testing.
	netInterfaceByName  = net.InterfaceByName
	netInterfaceByIndex = net.InterfaceByIndex
	netInterfaces       = net.Interfaces
	netInterfaceAddrs   = (*net.Interface).Addrs
)

func generateInterfaceName(key string, name string, useHead bool) string {
	_ = "STUB: not implemented"
	// #nosec G401: not used for security purposes
	return ""
}

// We use Node/Pod name to generate the interface name,
// valid chars for Node/Pod name are ASCII letters from a to z,
// the digits from 0 to 9, and the hyphen (-).
// Hyphen (-) is the only char which will impact command-line interpretation
// if the interface name starts with one, so we remove it here.

// GenerateContainerInterfaceKey generates a unique string for a Pod's
// interface as: "c/<Container-ID>/<IFDev-Name>".
// We must use ContainerID instead of PodNamespace + PodName because there could
// be more than one container associated with the same Pod at some point.
// For example, when deleting a StatefulSet Pod with 0 second grace period, the
// Pod will be removed from the Kubernetes API very quickly and a new Pod will
// be created immediately, and kubelet may process the deletion of the previous
// Pod and the addition of the new Pod simultaneously.
func GenerateContainerInterfaceKey(containerID, ifDev string) string {
	_ = "STUB: not implemented"
	return ""
}

// GenerateNodeTunnelInterfaceKey generates a unique string for a Node's
// tunnel interface as: node/<Node-name>.
func GenerateNodeTunnelInterfaceKey(nodeName string) string { _ = "STUB: not implemented"; return "" }

// GenerateContainerInterfaceName generates a unique interface name using the
// Pod's Namespace, name and container ID. The output should be deterministic
// (so that multiple calls to GenerateContainerInterfaceName with the same
// parameters return the same value). The output has the length of
// interfaceNameLength(15).
// The probability of collision should be neglectable.
func GenerateContainerInterfaceName(podName, podNamespace, containerID string) string {
	_ = "STUB: not implemented"
	// Use the podName as the prefix and the containerID as the hashing key.
	// podNamespace is not used currently.
	return ""
}

// GenerateContainerHostVethName generates a unique interface name using the
// Pod's Name, container ID, and the container veth interface name. The output
// should be deterministic.
func GenerateContainerHostVethName(podName, podNamespace, containerID, containerVeth string) string {
	_ = "STUB: not implemented"
	return ""
}

// Secondary interface.

// GenerateNodeTunnelInterfaceName generates a unique interface name for the
// tunnel to the Node, using the Node's name.
func GenerateNodeTunnelInterfaceName(nodeName string) string { _ = "STUB: not implemented"; return "" }

type LinkNotFound struct {
	error
}

func newLinkNotFoundError(name string) LinkNotFound {
	_ = "STUB: not implemented"
	return *new(LinkNotFound)
}

func listenUnix(address string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// GetIPNetDeviceFromIP returns local IPs/masks and associated device from IP, and ignores the interfaces which have
// names in the ignoredInterfaces.
func GetIPNetDeviceFromIP(localIPs *utilip.DualStackIPs, ignoredInterfaces sets.Set[string]) (v4IPNet *net.IPNet, v6IPNet *net.IPNet, iface *net.Interface, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// localIPs includes at most one IPv4 address and one IPv6 address. For each device in linkList, all its addresses
// are compared with IPs in localIPs. If found, the iface is set to the device and v4IPNet, v6IPNet are set to
// the matching addresses.

func GetIPNetDeviceByName(ifaceName string) (v4IPNet *net.IPNet, v6IPNet *net.IPNet, link *net.Interface, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func GetIPNetDeviceByCIDRs(cidrsList []string) (v4IPNet, v6IPNet *net.IPNet, link *net.Interface, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func GetIPv4Addr(ips []net.IP) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func GetIPWithFamily(ips []net.IP, addrFamily uint8) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// ExtendCIDRWithIP is used for extending an IPNet with an IP.
func ExtendCIDRWithIP(ipNet *net.IPNet, ip net.IP) (*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is copied from func commonPrefixLen in net/addrselect.go and modified:
// - Replace argument type IP with argument type net.IP.
// - Remove the prefix limit (64 bits) for IPv6.
func longestCommonPrefixLen(a, b net.IP) (cpl int) { _ = "STUB: not implemented"; return 0 }

// GetAllNodeAddresses gets all Node IP addresses (not including IPv6 link local address).
func GetAllNodeAddresses(excludeDeviceMatchers []func(string) bool) ([]net.IP, []net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get all interfaces.

// Get all IPs of every interface

// Skip IPv6 link local address

// Copied from github.com/vishvananda/netlink/netlink.go
// NewIPNet generates an IPNet from an ip address using a netmask of 32 or 128.
func NewIPNet(ip net.IP) *net.IPNet { _ = "STUB: not implemented"; return nil }

func PortToUint16(port int) uint16 { _ = "STUB: not implemented"; return 0 }

// GenerateUplinkInterfaceName generates the uplink interface name after bridged to OVS
func GenerateUplinkInterfaceName(name string) string { _ = "STUB: not implemented"; return "" }

func GenerateRandomMAC() net.HardwareAddr { _ = "STUB: not implemented"; return *new(net.HardwareAddr) }

// Unset the multicast bit.

// Set the local bit.

func getIPNetsByLink(link *net.Interface) ([]*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateOVSDatapathID generates an OVS datapath ID string.
func GenerateOVSDatapathID(macString string) string {
	_ = "STUB: not implemented"
	// The length of datapathID is 64 bits, the lower 48-bits are for a MAC address, while the
	// upper 16-bits are implementer-defined. Antrea uses "0x0000" for the upper 16-bits.
	return ""
}

// GetGatewayIPForPodCIDR returns the gateway IP for a given Pod CIDR.
func GetGatewayIPForPodCIDR(cidr *net.IPNet) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// GetGatewayIPForPodPrefix acts like GetGatewayIPForPodCIDR but takes a netip.Prefix as a parameter
// and returns a netip.Addr value.
func GetGatewayIPForPodPrefix(prefix netip.Prefix) netip.Addr {
	_ = "STUB: not implemented"
	return *new(netip.Addr)
}
