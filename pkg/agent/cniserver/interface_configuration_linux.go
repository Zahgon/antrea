//go:build linux
// +build linux

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

package cniserver

import (
	"net"

	cnitypes "github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/plugins/pkg/ip"
	"github.com/containernetworking/plugins/pkg/ipam"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"

	"antrea.io/antrea/v2/pkg/agent/util"
	"antrea.io/antrea/v2/pkg/agent/util/arping"
	"antrea.io/antrea/v2/pkg/agent/util/ethtool"
	"antrea.io/antrea/v2/pkg/agent/util/ndp"
	netlinkutil "antrea.io/antrea/v2/pkg/agent/util/netlink"
	cnipb "antrea.io/antrea/v2/pkg/apis/cni/v1beta1"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
)

// NetDeviceType type Enum
const (
	netDeviceTypeVeth = "veth"
	netDeviceTypeVF   = "vf"
)

// Declared variables for test
var (
	ipSetupVethWithName            = ip.SetupVethWithName
	ipDelLinkByName                = ip.DelLinkByName
	ipamConfigureIface             = ipam.ConfigureIface
	ethtoolTXHWCsumOff             = ethtool.EthtoolTXHWCsumOff
	renameInterface                = util.RenameInterface
	netInterfaceByName             = net.InterfaceByName
	netInterfaceByIndex            = net.InterfaceByIndex
	arpingGratuitousARPOverIface   = arping.GratuitousARPOverIface
	ndpGratuitousNDPOverIface      = ndp.GratuitousNDPOverIface
	ipValidateExpectedInterfaceIPs = ip.ValidateExpectedInterfaceIPs
	ipValidateExpectedRoute        = ip.ValidateExpectedRoute
	ipGetVethPeerIfindex           = ip.GetVethPeerIfindex
	getNSDevInterface              = util.GetNSDevInterface
	getNSPeerDevBridge             = util.GetNSPeerDevBridge
	nsGetNS                        = ns.GetNS
	nsWithNetNSPath                = ns.WithNetNSPath
	nsIsNSorErr                    = ns.IsNSorErr
	tempNetNS                      = ns.TempNetNS
)

type ifConfigurator struct {
	ovsDatapathType             ovsconfig.OVSDatapathType
	isOvsHardwareOffloadEnabled bool
	disableTXChecksumOffload    bool
	netlink                     netlinkutil.Interface
	sriovnet                    SriovNet
}

func newInterfaceConfigurator(ovsDatapathType ovsconfig.OVSDatapathType, isOvsHardwareOffloadEnabled bool, disableTXChecksumOffload bool) (*ifConfigurator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ic *ifConfigurator) moveIFToNetNS(ifname string, netns ns.NetNS) error {
	_ = "STUB: not implemented"
	return nil
}

// Move VF device to ns

func (ic *ifConfigurator) configureVFLinkAndIPAM(link netlink.Link, containerID string, containerIfaceName string, netnsPath string, mtu int, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// result.Interfaces must be set before this.

func (ic *ifConfigurator) moveOffloadVFToContainerNS(vfNetDevice string, containerID string, containerNetNS string, containerIfaceName string, mtu int, result *current.Result) error {
	_ = "STUB: not implemented"
	// Move VF to Container namespace
	return nil
}

func (ic *ifConfigurator) moveVFtoContainerNS(vfNetDevice string, containerID string, containerNetNS string, containerIfaceName string, mtu int, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// Rename the device in a temp NS to avoid race condition.
// This rename logic is mainly referring to the code of host-device CNI.
// https://github.com/containernetworking/plugins/blob/a5d507e2b884d8bd6a001c9e5a9118113ffef444/plugins/main/host-device/host-device.go#L238-L464

// Move the host VF device into tempNS

// Look up the device in tempNS, as the index might have changed

// Rename the device to the wanted name

// Restore the original device name in case of error

// Remove the alias on error

// Move VF to Container namespace

// Look up the device again on error, as the index might have changed

// Move the interface back to tempNS on error, so that we can undo VF
// rename and alias setting in the previous steps.

// configureContainerSriovLinkOnBridge moves the VF to the container namespace for OVS offload.
func (ic *ifConfigurator) configureContainerSriovLinkOnBridge(
	podName string,
	podNamespace string,
	containerID string,
	containerNetNS string,
	containerIfaceName string,
	mtu int,
	pciAddress string,
	result *current.Result,
) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. get VF netdevice from PCI

// Make sure we have 1 netdevice per PCI address

// 2. get Uplink netdevice

// 3. get VF index from PCI

// 4. lookup representor

// 5. rename VF representor to hostIfaceName

// configureContainerSriovLink moves the VF to the container namespace for Pod link SR-IOV interface;
// intended for multiple interfaces other than the primary interface.
func (ic *ifConfigurator) configureContainerSriovLink(
	podName string,
	podNamespace string,
	containerID string,
	containerNetNS string,
	containerIfaceName string,
	mtu int,
	pciAddress string,
	result *current.Result,
) error {
	_ = "STUB: not implemented"
	return nil
}

// recoverVFInterfaceName rename the interface back to the original VF interface name.
func (ic *ifConfigurator) recoverVFInterfaceName(containerIfaceName string, containerNetNS string) error {
	_ = "STUB: not implemented"
	return nil
}

// Move VF from container namespace to tempNS

// Lookup the device in tempNS (index might have changed)

// Move VF back to container namespace on error

// Rename container device to originalVFName

// Rename the device back to containerIfaceName on error

// Unset device's alias property

// Set back the device alias to originalVFName on error

// Move VF from container namespace back to hostNS

// configureContainerLinkVeth creates a veth pair: one in the container netns and one in the host netns, and configures IP
// address and routes to the container veth.
func (ic *ifConfigurator) configureContainerLinkVeth(
	podName string,
	podNamespace string,
	containerID string,
	containerNetNS string,
	containerIfaceName string,
	mtu int,
	result *current.Result,
	mac net.HardwareAddr,
) error {
	_ = "STUB: not implemented"
	// Include the container veth interface name in the name generation, as one Pod can have more
	// than one interfaces inc. secondary interfaces, while the host interface name must be unique.
	return nil
}

// Disable TX checksum offloading when it's configured explicitly.

// result.Interfaces must be set before this.

// advertiseContainerAddr sends 3 GARP packets in another goroutine with 50ms interval, if the
// container interface is assigned an IPv4 address. It's because Openflow entries are installed
// asynchronously, and the gratuitous ARP could be sent out after the Openflow entries are
// installed. Using another goroutine to ensure the processing of CNI ADD request is not blocked.
func (ic *ifConfigurator) advertiseContainerAddr(containerNetNS string, containerIfaceName string, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// Sending Gratuitous ARP is a best-effort action and is unlikely to fail as we have ensured the netns is valid.

// Send gratuitous ARP/NDP to network in case of stale mappings for this IP address
// (e.g. if a previous - deleted - Pod was using the same IP).

// configureContainerLink creates a veth pair: one in the container netns and one in the host netns, and configures IP
// address and routes to the container veth.
func (ic *ifConfigurator) configureContainerLink(
	podName string,
	podNamespace string,
	containerID string,
	containerNetNS string,
	containerIfaceName string,
	mtu int,
	brSriovVFDeviceID string,
	podSriovVFDeviceID string,
	result *current.Result,
	containerAccess *containerAccessArbitrator,
	mac net.HardwareAddr,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Move SR-IOV VF to network namespace

// For Pod link SR-IOV interface not attached to the OVS bridge

// Create veth pair and link up

func (ic *ifConfigurator) changeContainerMTU(containerNetNS string, containerIFDev string, mtuDeduction int) error {
	_ = "STUB: not implemented"
	return nil
}

func (ic *ifConfigurator) removeContainerLink(containerID, hostInterfaceName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Don't return an error if the device is already removed as CniDel can be called multiple times.

func parseContainerIfaceFromResults(cfgArgs *cnipb.CniCmdArgs, prevResult *current.Result) *current.Interface {
	_ = "STUB: not implemented"
	return nil
}

func (ic *ifConfigurator) checkContainerInterface(
	containerNetns, containerID string,
	containerIface *current.Interface,
	containerIPs []*current.IPConfig,
	containerRoutes []*cnitypes.Route,
	sriovVFDeviceID string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Check netns configuration
}

// Check container interface configuration

// Check container link config

// Check container IP config

// Check container route config

func (ic *ifConfigurator) validateContainerVFInterface(intf *current.Interface, sriovVFDeviceID string) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

// The check makes sure that the SR-IOV VF netdevice is not in the host namespace.
// GetNetDevicesFromPCI is using linux sysfs to find the VF netdevice. The method
// is running in container network namespace, but we are still in the antrea-agent
// filesystem. The antrea-agent container is privileged, which allows access to
// the host sysfs, therefore the validation is to make sure that the VF netdevice
// is not in the host network namespace.

func (ic *ifConfigurator) validateContainerVethInterface(intf *current.Interface) (*vethPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ic *ifConfigurator) validateVFRepInterface(sriovVFDeviceID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ic *ifConfigurator) validateContainerPeerInterface(interfaces []*current.Interface, containerVeth *vethPair) (*vethPair, error) {
	_ = "STUB: not implemented"
	// Iterate all the passed interfaces and look up the host interface by
	// matching the veth peer interface index.
	return nil, nil
}

// Not in the default Namespace. Must be the container interface.

func (ic *ifConfigurator) getInterceptedInterfaces(
	sandbox string,
	containerNetNS string,
	containerIFDev string,
) (*current.Interface, *current.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Setup dev in host ns.

// addPostInterfaceCreateHook is called only on Windows. Adding this function in this file because it is defined in the
// interface `podInterfaceConfigurator`.
func (ic *ifConfigurator) addPostInterfaceCreateHook(containerID, endpointName string, containerAccess *containerAccessArbitrator, hook postInterfaceCreateHook) error {
	_ = "STUB: not implemented"
	return nil
}

func (ic *ifConfigurator) validateInterface(intf *current.Interface, inNetns bool, ifType string) (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}

func isVeth(link netlink.Link) bool { _ = "STUB: not implemented"; return false }
