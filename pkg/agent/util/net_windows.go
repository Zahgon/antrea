//go:build windows
// +build windows

// Copyright 2020 Antrea Authors
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

	"github.com/Microsoft/hcsshim"

	"antrea.io/antrea/v2/pkg/agent/util/winnet"
)

const (
	LocalHNSNetwork = "antrea-hnsnetwork"
	HNSNetworkType  = "Transparent"
	namedPipePrefix = `\\.\pipe\`

	AntreaNatName = "antrea-nat"
	LocalVMSwitch = "antrea-switch"
)

var (
	winnetUtil winnet.Interface = &winnet.Handle{}

	getHNSNetworkByName = hcsshim.GetHNSNetworkByName
	hnsNetworkRequest   = hcsshim.HNSNetworkRequest
	hnsNetworkCreate    = (*hcsshim.HNSNetwork).Create
	hnsNetworkDelete    = (*hcsshim.HNSNetwork).Delete
)

func GetNSPath(containerNetNS string) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// CreateHNSNetwork creates a new HNS Network, whose type is "Transparent". The NetworkAdapter is using the host
		// interface which is configured with Node IP. HNS Network properties "ManagementIP" and "SourceMac" are used to record
		// the original IP and MAC addresses on the network adapter.
		nil
}

func CreateHNSNetwork(hnsNetName string, subnetCIDR *net.IPNet, nodeIP *net.IPNet, adapter *net.Interface) (*hcsshim.HNSNetwork, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteHNSNetwork(hnsNetName string) error { _ = "STUB: not implemented"; return nil }

type vSwitchExtensionPolicy struct {
	ExtensionID string `json:"Id,omitempty"`
	IsEnabled   bool
}

type ExtensionsPolicy struct {
	Extensions []vSwitchExtensionPolicy `json:"Extensions"`
}

// EnableHNSNetworkExtension enables the specified vSwitchExtension on the target HNS Network. Antrea calls this function
// to enable OVS Extension on the HNS Network.
func EnableHNSNetworkExtension(hnsNetID string, vSwitchExtension string) error {
	_ = "STUB: not implemented"
	return nil
}

func SetLinkUp(name string) (net.HardwareAddr, int, error) {
	_ = "STUB: not implemented"
	return *new(net.HardwareAddr), 0, nil
}

func addrEqual(addr1, addr2 *net.IPNet) bool { _ = "STUB: not implemented"; return false }

// addrSliceDifference returns elements in s1 but not in s2.
func addrSliceDifference(s1, s2 []*net.IPNet) []*net.IPNet { _ = "STUB: not implemented"; return nil }

// ConfigureLinkAddresses adds the provided addresses to the interface identified by index idx, if
// they are missing from the interface. Any other existing address already configured for the
// interface will be removed, unless it is a link-local address. At the moment, this function only
// supports IPv4 addresses and will ignore any address in ipNets that is not IPv4.
func ConfigureLinkAddresses(idx int, ipNets []*net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

// PrepareHNSNetwork creates HNS Network for containers.
func PrepareHNSNetwork(subnetCIDR *net.IPNet, nodeIPNet *net.IPNet, uplinkAdapter *net.Interface, nodeGateway string, dnsServers string, routes []interface{}, newName string) error {
	_ = "STUB: not implemented"
	return nil
}

// On the current Windows testbed, it takes a maximum of 1.8 seconds to obtain a valid IP.
// Therefore, we set the timeout limit to triple of that value, allowing a maximum wait of 6 seconds here.

// By default, "ipFound" should be true after Windows creates the HNSNetwork. The following check is for some corner
// cases that Windows fails to move the physical adapter's IP address to the virtual network adapter, e.g., DHCP
// Server fails to allocate IP to new virtual network.

// Rename the vnic created by Windows host with the given newName, then it can be used by OVS when creating bridge port.

// Rename NetAdapter in the meanwhile, then the network adapter can be treated as a host network adapter other than
// a vm network adapter.

// Enable OVS Extension on the HNS Network. If an error occurs, delete the HNS Network and return the error.
// While the hnsshim API allows for enabling the OVS extension when creating an HNS network, it can cause the adapter being unable
// to obtain a valid DHCP IP in case of network interruption. Therefore, we have to enable the OVS extension after running adapterIPExists.

// adapterIPExists finds the network adapter configured with the provided IP, MAC and its name has the given prefix.
// If "namePrefix" is empty, it returns the first network adapter with the provided IP and MAC.
// It returns true if the IP is found on the adapter, otherwise it returns false.
func adapterIPExists(ip net.IP, mac net.HardwareAddr, namePrefix string) (*net.Interface, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// GetDefaultGatewayByInterfaceIndex returns the default gateway configured on the specified interface.
func GetDefaultGatewayByInterfaceIndex(ifIndex int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ListenLocalSocket creates a listener on a Unix domain socket or a Windows named pipe.
// - If the specified address starts with "\\.\pipe\",  create a listener on a Windows named pipe path.
// - Else create a listener on a local Unix domain socket.
func ListenLocalSocket(address string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func HostInterfaceExists(ifaceName string) bool { _ = "STUB: not implemented"; return false }

func GetInterfaceConfig(ifName string) (*net.Interface, []*net.IPNet, []interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// Skip the routes automatically generated by Windows host when adding IP address on the network adapter.

func RenameInterface(from, to string) error { _ = "STUB: not implemented"; return nil }

func GenHostInterfaceName(upLinkIfName string) string { _ = "STUB: not implemented"; return "" }
