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

package agent

import (
	"net"

	"antrea.io/antrea/v2/pkg/agent/util"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
)

var (
	// getInterfaceByName is meant to be overridden for testing.
	getInterfaceByName = net.InterfaceByName

	// setInterfaceARPAnnounce is meant to be overridden for testing.
	setInterfaceARPAnnounce = util.EnsureARPAnnounceOnInterface
)

// prepareHostNetwork returns immediately on Linux.
func (i *Initializer) prepareHostNetwork() error {
	_ = "STUB: not implemented"

	// Assuming a page cache of 4096, based on Suricata source code from L1752-L1798
	// at https://github.com/OISF/suricata/blob/49713ebaa0b8edb057d60f1cfe9126946645a848/src/source-af-packet.c#L1757C2-L1777C129.
	// The maximum supported MTU by Suricata is 32678 after calculation.
	return nil
}

const maxMTUSupportedBySuricata = 32678

// prepareOVSBridgeForK8sNode returns immediately on Linux if connectUplinkToBridge is false.
func (i *Initializer) prepareOVSBridgeForK8sNode() error { _ = "STUB: not implemented"; return nil }

// Get uplink network configuration.
// TODO(gran): support IPv6

// Gateway and DNSServers are not configured at adapter in Linux
// Limitation: dynamic DNS servers will be lost after DHCP lease expired

// Set datapathID of OVS bridge.
// If no datapathID configured explicitly, the reconfiguration operation will change OVS bridge datapathID
// and break the OpenFlow channel.

// If local port exists, get the real uplink interface.
// This branch is used when antrea-agent had a hard restart (e.g. SIGKILL)

// getTunnelLocalIP returns local_ip of tunnel port.
// On linux platform, local_ip option is not needed.
func (i *Initializer) getTunnelPortLocalIP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func getTransportIPNetDeviceByName(ifaceName string, ovsBridgeName string) (*net.IPNet, *net.IPNet, *net.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (i *Initializer) ConnectUplinkToOVSBridge() error {
	_ = "STUB: not implemented"
	// Return immediately on Linux if connectUplinkToBridge is false.
	return nil
}

// We request the same MTU for the bridge interface as for the uplink adapter. If we don't,
// OVS will default to the lowest MTU among all existing bridge ports, including container
// ports. There may be some existing workloads with a lower MTU, and using that lower value
// may impact host connectivity.

// Create uplink port.

// Add newly created uplinkInterface to interface cache.

//nolint: govet

// RestoreOVSBridge returns immediately on Linux if connectUplinkToBridge is false.
// OVS is managed by Antrea in Linux, network config must be restored to uplink before Antrea Agent shutdown.
func (i *Initializer) RestoreOVSBridge() { _ = "STUB: not implemented"; return }

func (i *Initializer) setInterfaceMTU(iface string, mtu int) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *Initializer) setVMNodeConfig(en *v1alpha1.ExternalNode, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *Initializer) prepareOVSBridgeForVM() error { _ = "STUB: not implemented"; return nil }

func (i *Initializer) installVMInitialFlows() error {
	_ = "STUB: not implemented"

	// prepareL7EngineInterfaces creates two OVS internal ports. An application-aware engine will connect to OVS
	// through these two ports.
	return nil
}

func (i *Initializer) prepareL7EngineInterfaces() error { _ = "STUB: not implemented"; return nil }

// Set the ports with no-flood to reject ARP flood packets at every startup.

// Set MTU of the ports to the calculated MTU value at every startup.

// Currently, the maximum of MTU supported by L7 NetworkPolicy engine Suricata is 32678 (assuming that the page size
// is 4096). If the calculated MTU value is greater than 32678, Suricata may fail to start.

func (i *Initializer) setTXChecksumOffloadOnGateway() error { _ = "STUB: not implemented"; return nil }
