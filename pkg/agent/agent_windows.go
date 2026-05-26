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

package agent

import (
	"net"

	"antrea.io/antrea/v2/pkg/agent/util/winnet"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
)

var (
	winnetUtil winnet.Interface = &winnet.Handle{}
	// setInterfaceMTU is meant to be overridden for testing
	setInterfaceMTU = winnetUtil.SetNetAdapterMTU

	// setInterfaceARPAnnounce is meant to be overridden for testing.
	setInterfaceARPAnnounce = func(ifaceName string, value int) error { return nil }
)

func (i *Initializer) prepareHostNetwork() error { _ = "STUB: not implemented"; return nil }

// prepareHNSNetworkAndOVSExtension creates HNS Network for containers, and enables OVS Extension on it.
func (i *Initializer) prepareHNSNetworkAndOVSExtension() error {
	_ = "STUB: not implemented"
	// If the HNS Network already exists, return immediately.
	return nil
}

// Enable OVS Extension on the HNS Network.

// Enable RSC for existing vSwitch.

// Save the uplink adapter name to check if the OVS uplink port has been created in prepareOVSBridge stage.

// Save the uplink adapter MAC to modify Pod traffic source MAC if the packet is directly output to the uplink
// interface in OVS pipeline.

// Get uplink network configuration. The uplink interface is the one used for transporting Pod traffic across Nodes.
// Use the interface specified with "transportInterface" in the configuration if configured, otherwise the interface
// configured with NodeIP is used as uplink.

// To forward container traffic to physical network, Transparent HNSNetwork must have a physical adapter attached,
// otherwise creating it would fail with "The parameter is incorrect" if the provided adapter is virtual or "An
// adapter was not found" if no adapter is provided and no physical adapter is available on the host.
// If the discovered adapter is virtual, it likely means the physical adapter is already attached to another
// HNSNetwork. For example, docker may create HNSNetworks which attach to the physical adapter.

// Save routes which are configured on the uplink interface, and configure them on the management virtual adapter
// if Windows host doesn't move the configuration automatically.

// Create HNS network.

func (i *Initializer) prepareVMNetworkAndOVSExtension() error {
	_ = "STUB: not implemented"
	return nil
}

// Check whether VM Switch is created

// Get the uplink interface configuration

// Rename interfaceName to interfaceName~

// prepareOVSBridgeForK8sNode adds local port and uplink port to OVS bridge after OVS extension is enabled on HNSNetwork.
// This function deletes OVS bridge and HNS network created by Antrea on failure.
func (i *Initializer) prepareOVSBridgeForK8sNode() error { _ = "STUB: not implemented"; return nil }

// prepareOVSBridgeOnHNSNetwork adds local port and uplink to OVS bridge after the OVS Extension is enabled on HNSNetwork.
// This function will delete OVS bridge and HNS network created by Antrea at failures.
func (i *Initializer) prepareOVSBridgeOnHNSNetwork() error { _ = "STUB: not implemented"; return nil }

// prepareOVSBridge only works on Windows platform. The operation has a chance to fail on the first time agent
// starts up when OVS bridge uplink and local interface have not been configured. If the operation fails, the
// host can not communicate with external network. To make sure the agent can connect to API server in
// next retry, this step deletes OVS bridge and HNS network created previously which will restore the
// host network.

// Set datapathID of OVS bridge.
// If no datapathID configured explicitly, the reconfiguration operation will change OVS bridge datapathID
// and break the OpenFlow channel.

// Create local port.

// OVS does not receive "ofport_request" param when creating local port, so here use
// ovsconfig.AutoAssignedOFPort (0).

// If uplink already exists, return early.

// We check if the antrea-type external ID, which is used to store the interface
// type is present. If it is missing, we add it. Prior to Antrea v2.0, this external
// ID was not set for the uplink port, which was a bug. We need this code for
// backwards-compatibility, as other parts of the code may assume this external ID
// always exist. This code can be removed in Antrea v2.3.

// Nothing to do, external ID already exists

// Add missing external ID.
// A copy is required because of the type mismatch.

// Create uplink port.

//nolint: govet

// Enable IP forwarding on the bridge local interface. Traffic from the uplink interface will be output to the bridge
// local interface directly. When an external client connects to a LoadBalancer type Service, and the packets of the
// connection are routed to the selected backend Pod via the bridge interface; if we do not enable IP forwarding on
// the bridge interface, the packet will be discarded on the bridge interface as the destination of the packet
// is not the Node.

// Set the uplink with "no-flood" config, so that the IP of local Pods and "antrea-gw0" will not be leaked to the
// underlay network by the "normal" flow entry.

func (i *Initializer) prepareOVSBridgeForVM() error { _ = "STUB: not implemented"; return nil }

// TODO: Have a separate function for creation of pair ports
// Create uplink port on OVS.

// Manual clean up of OVS configurations is required, when agent exits
// abruptly or when the auto cleanup operation fails.

// Query the uplink port to check if its created

// ExternalEntity is not processed yet, so an empty name is set for entityName in OVSDB,
// which will be updated by ExternalNode controller.

// Create host port on OVS.

// Manual clean up of OVS configurations is required, when agent exits abruptly.

// Query the host port to check if its created

// getTunnelLocalIP returns local_ip of tunnel port
func (i *Initializer) getTunnelPortLocalIP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// saveHostRoutes saves routes configured on the uplink interface before the
// interface is configured as the uplink of Antrea HNS network.
// The routes will be restored on the OVS bridge interface after the IP
// configuration is moved to the OVS bridge.
func (i *Initializer) saveHostRoutes() error {
	_ = "STUB: not implemented"
	// IPv6 is not supported on Windows currently. Please refer to https://github.com/antrea-io/antrea/issues/5162
	// for more information.
	return nil
}

// Skip default route. The default route will be added automatically when
// configuring IP address on OVS bridge interface.

func getTransportIPNetDeviceByName(ifaceName string, ovsBridgeName string) (*net.IPNet, *net.IPNet, *net.Interface, error) {
	_ = "STUB: not implemented"
	// Find transport Interface in the order: ifaceName -> br-int. Return immediately if
	// an interface using the specified name exists. Using br-int is for restart agent case.
	return nil, nil, nil, nil
}

// ConnectUplinkToOVSBridge returns immediately on Windows. The uplink interface
// will be connected to the bridge in prepareOVSBridge().
func (i *Initializer) ConnectUplinkToOVSBridge() error {
	_ = "STUB: not implemented"

	// RestoreOVSBridge returns immediately in Windows.
	// OVS is managed by system in Windows, network config can be retained after Antrea shutdown.
	return nil
}

func (i *Initializer) RestoreOVSBridge() { _ = "STUB: not implemented"; return }

func (i *Initializer) setInterfaceMTU(iface string, mtu int) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *Initializer) setVMNodeConfig(en *v1alpha1.ExternalNode, nodeName string) error {
	_ = "STUB: not implemented"
	// TODO: Handle for multiple interfaces
	return nil
}

// installVMFlows configures default flows between uplink and host port,
// so that antrea-agent can connect to antrea-controller.
func (i *Initializer) installVMInitialFlows() error { _ = "STUB: not implemented"; return nil }

func (i *Initializer) prepareL7EngineInterfaces() error { _ = "STUB: not implemented"; return nil }

func (i *Initializer) setTXChecksumOffloadOnGateway() error { _ = "STUB: not implemented"; return nil }
