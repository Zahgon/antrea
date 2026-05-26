//go:build windows
// +build windows

// Copyright 2024 Antrea Authors
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

package winnet

import (
	"errors"
	"net"

	"golang.org/x/sys/windows"

	ps "antrea.io/antrea/v2/pkg/agent/util/powershell"
	antreasyscall "antrea.io/antrea/v2/pkg/agent/util/syscall"
)

const (
	ContainerVNICPrefix = "vEthernet"
	OVSExtensionID      = "583CC151-73EC-4A6A-8B47-578297AD7623"
	ovsExtensionName    = "Open vSwitch Extension"

	MetricDefault = 256
	MetricHigh    = 50

	// Filter masks are used to indicate the attributes used for route filtering.
	RT_FILTER_IF uint64 = 1 << (1 + iota)
	RT_FILTER_METRIC
	RT_FILTER_DST
	RT_FILTER_GW

	// IP_ADAPTER_DHCP_ENABLED is defined in the Win32 API document.
	// https://learn.microsoft.com/en-us/windows/win32/api/iptypes/ns-iptypes-ip_adapter_addresses_lh
	IP_ADAPTER_DHCP_ENABLED = 0x00000004

	// GAA_FLAG_INCLUDE_ALL_COMPARTMENTS is used in windows.GetAdapterAddresses parameter
	// flags to return addresses in all routing compartments.
	GAA_FLAG_INCLUDE_ALL_COMPARTMENTS = 0x00000200

	// GAA_FLAG_INCLUDE_ALL_INTERFACES is used in windows.GetAdapterAddresses parameter
	// flags to return addresses for all NDIS interfaces.
	GAA_FLAG_INCLUDE_ALL_INTERFACES = 0x00000100
)

type Handle struct{}

var (
	// Declared variables which are meant to be overridden for testing.
	antreaNetIO          = antreasyscall.NewNetIO()
	getAdaptersAddresses = windows.GetAdaptersAddresses
	runCommand           = ps.RunCommand
)

func routeFromIPForwardRow(row *antreasyscall.MibIPForwardRow) *Route {
	_ = "STUB: not implemented"
	return nil
}

// IsVirtualNetAdapter checks if the provided network adapter is virtual.
func (h *Handle) IsVirtualNetAdapter(adapterName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsNetAdapterStatusUp checks if the status of the provided network adapter is UP.
func (h *Handle) IsNetAdapterStatusUp(adapterName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// EnableNetAdapter sets the specified network adapter status as UP.
func (h *Handle) EnableNetAdapter(adapterName string) error { _ = "STUB: not implemented"; return nil }

// AddNetAdapterIPAddress adds the specified IP address on the specified network adapter.
func (h *Handle) AddNetAdapterIPAddress(adapterName string, ipConfig *net.IPNet, gateway string) error {
	_ = "STUB: not implemented"
	return nil
}

// If the address already exists, ignore the error.

// RemoveNetAdapterIPAddress removes the specified IP address from the specified network adapter.
func (h *Handle) RemoveNetAdapterIPAddress(adapterName string, ipAddr net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// If the address does not exist, ignore the error.

// EnableIPForwarding enables the network adapter to forward IP packets that arrive at this network adapter to other ones.
func (h *Handle) EnableIPForwarding(adapterName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handle) RenameVMNetworkAdapter(networkName, macStr, newName string, renameNetAdapter bool) error {
	_ = "STUB: not implemented"
	return nil
}

// EnableRSCOnVSwitch enables RSC in the vSwitch to reduce host CPU utilization and increase throughput for virtual
// workloads by coalescing multiple TCP segments into fewer, but larger segments.
func (h *Handle) EnableRSCOnVSwitch(vSwitch string) error { _ = "STUB: not implemented"; return nil }

// RSC doc says it applies to Windows Server 2019, which is the only Windows operating system supported so far, so
// this should not happen. However, this is only an optimization, no need to crash the process even if it's not
// supported.
// https://docs.microsoft.com/en-us/windows-server/networking/technologies/hpn/rsc-in-the-vswitch

// GetDefaultGatewayByNetAdapterIndex returns the default gateway configured on the specified network adapter.
func (h *Handle) GetDefaultGatewayByNetAdapterIndex(adapterIndex int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetDNServersByNetAdapterIndex returns the DNS servers configured on the specified network adapter.
func (h *Handle) GetDNServersByNetAdapterIndex(adapterIndex int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetNetAdapterDNSServers configures DNS servers on network adapter.
func (h *Handle) SetNetAdapterDNSServers(adapterName, dnsServers string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handle) NetAdapterExists(adapterName string) bool { _ = "STUB: not implemented"; return false }

// IsNetAdapterIPv4DHCPEnabled returns the IPv4 DHCP status on the specified network adapter.
func (h *Handle) IsNetAdapterIPv4DHCPEnabled(adapterName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// SetNetAdapterMTU configures network adapter MTU on host for Pods. MTU change cannot be realized with HNSEndpoint because
// there's no MTU field in HNSEndpoint:
// https://github.com/Microsoft/hcsshim/blob/4a468a6f7ae547974bc32911395c51fb1862b7df/internal/hns/hnsendpoint.go#L12
func (h *Handle) SetNetAdapterMTU(adapterName string, mtu int) error {
	_ = "STUB: not implemented"
	return nil
}

func AddressFamilyByIP(ip net.IP) uint16 { _ = "STUB: not implemented"; return 0 }

func VirtualAdapterName(name string) string { _ = "STUB: not implemented"; return "" }

func toMibIPForwardRow(r *Route) *antreasyscall.MibIPForwardRow {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handle) AddNetRoute(route *Route) error { _ = "STUB: not implemented"; return nil }

func (h *Handle) RemoveNetRoute(route *Route) error { _ = "STUB: not implemented"; return nil }

func (h *Handle) ReplaceNetRoute(route *Route) error { _ = "STUB: not implemented"; return nil }

func (h *Handle) RouteListFiltered(family uint16, filter *Route, filterMask uint64) ([]Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseCmdResult(result string, columns int) [][]string { _ = "STUB: not implemented"; return nil }

// Skip if an empty line or something similar

func (h *Handle) AddNetNat(netNatName string, subnetCIDR *net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handle) ReplaceNetNatStaticMapping(mapping *NetNatStaticMapping) error {
	_ = "STUB: not implemented"
	return nil
}

// getNetNatStaticMapping checks if a NetNatStaticMapping exists.
func getNetNatStaticMapping(mapping *NetNatStaticMapping) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AddNetNatStaticMapping adds a static mapping to a NAT instance.
func (h *Handle) AddNetNatStaticMapping(mapping *NetNatStaticMapping) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveNetNatStaticMapping removes a static mapping from a NetNat instance.
func (h *Handle) RemoveNetNatStaticMapping(mapping *NetNatStaticMapping) error {
	_ = "STUB: not implemented"
	return nil
}

func removeNetNatStaticMappingByID(netNatName string, id int) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveNetNatStaticMappingsByNetNat removes all static mappings from a NetNat instance.
func (h *Handle) RemoveNetNatStaticMappingsByNetNat(netNatName string) error {
	_ = "STUB: not implemented"
	return nil
}

// getNetNeighbor gets neighbor cache entries with Get-NetNeighbor.
func getNetNeighbor(neighbor *Neighbor) ([]Neighbor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get-NetNeighbor returns LinkLayerAddress like "AA-BB-CC-DD-EE-FF".

// newNetNeighbor creates a new neighbor cache entry with New-NetNeighbor.
func newNetNeighbor(neighbor *Neighbor) error { _ = "STUB: not implemented"; return nil }

func removeNetNeighbor(neighbor *Neighbor) error { _ = "STUB: not implemented"; return nil }

func (h *Handle) ReplaceNetNeighbor(neighbor *Neighbor) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handle) GetVMSwitchNetAdapterName(vmSwitch string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Remove the leading and trailing {} brackets

func (h *Handle) VMSwitchExists(vmSwitch string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AddVMSwitch creates a VMSwitch and enables OVS extension. Connection to VMSwitch is lost for few seconds.
// TODO: Handle for multiple interfaces
func (h *Handle) AddVMSwitch(adapterName, vmSwitch string) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handle) RemoveVMSwitch(vmSwitch string) error { _ = "STUB: not implemented"; return nil }

type updateIPInterfaceFunc func(entry *antreasyscall.MibIPInterfaceRow) *antreasyscall.MibIPInterfaceRow

type adapter struct {
	net.Interface
	compartmentID uint32
	flags         uint32
}

func (a *adapter) setMTU(mtu int, family uint16) error { _ = "STUB: not implemented"; return nil }

func (a *adapter) setForwarding(enabledForwarding bool, family uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *adapter) setIPInterfaceEntry(family uint16, updateFunc updateIPInterfaceFunc) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	errInvalidInterfaceName = errors.New("invalid network interface name")
	errNoSuchInterface      = errors.New("no such network interface")
)

func getAdapterInAllCompartmentsByName(name string) (*adapter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *Handle) EnableVMSwitchOVSExtension(vmSwitch string) error {
	_ = "STUB: not implemented"
	return nil
}

// parseOVSExtensionOutput parses the VM extension output and returns the value of Enabled field.
func parseOVSExtensionOutput(s string) bool { _ = "STUB: not implemented"; return false }

func (h *Handle) IsVMSwitchOVSExtensionEnabled(vmSwitch string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (h *Handle) RenameNetAdapter(oriName string, newName string) error {
	_ = "STUB: not implemented"
	return nil
}

func getAdaptersByName(name string) ([]adapter, error) { _ = "STUB: not implemented"; return nil, nil }

// ipv6IfIndex is a substitute for ifIndex

// For now we need to infer link-layer service capabilities from media types.
// TODO: use MIB_IF_ROW2.AccessType now that we no longer support Windows XP.

// assume all services available; LANE, point-to-point and point-to-multipoint

// adapterAddresses returns a list of IpAdapterAddresses structures. The structure
// contains an IP adapter and flattened multiple IP addresses including unicast, anycast
// and multicast addresses.
// This function is copied from go/src/net/interface_windows.go, with a change that flag
// GAA_FLAG_INCLUDE_ALL_COMPARTMENTS is introduced to query interfaces in all compartments,
// and GAA_FLAG_INCLUDE_ALL_INTERFACES is introduced to query all NDIS interfaces even they
// are not configured with any IP addresses, e.g., uplink.
func adapterAddresses() ([]*windows.IpAdapterAddresses, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// recommended initial size
