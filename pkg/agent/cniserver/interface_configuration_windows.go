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

package cniserver

import (
	"net"
	"sync"

	"github.com/Microsoft/hcsshim"
	"github.com/Microsoft/hcsshim/hcn"
	cnitypes "github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"

	"antrea.io/antrea/v2/pkg/agent/util"
	"antrea.io/antrea/v2/pkg/agent/util/winnet"
	cnipb "antrea.io/antrea/v2/pkg/apis/cni/v1beta1"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
)

const (
	notFoundHNSEndpoint = "The endpoint was not found"
)

var (
	getHnsNetworkByNameFunc         = hcsshim.GetHNSNetworkByName
	listHnsEndpointFunc             = hcsshim.HNSListEndpointRequest
	hostInterfaceExistsFunc         = util.HostInterfaceExists
	getNetInterfaceAddrsFunc        = getNetInterfaceAddrs
	createHnsEndpointFunc           = createHnsEndpoint
	attachEndpointInNamespaceFunc   = attachEndpointInNamespace
	getHcnEndpointByIDFunc          = hcn.GetEndpointByID
	deleteHnsEndpointFunc           = deleteHnsEndpoint
	removeEndpointFromNamespaceFunc = hcn.RemoveNamespaceEndpoint
	getHnsEndpointByNameFunc        = hcsshim.GetHNSEndpointByName
	getNetInterfaceByNameFunc       = net.InterfaceByName
)

type ifConfigurator struct {
	hnsNetwork *hcsshim.HNSNetwork
	epCache    *sync.Map
	winnet     winnet.Interface
}

// disableTXChecksumOffload is ignored on Windows.
func newInterfaceConfigurator(ovsDatapathType ovsconfig.OVSDatapathType, isOvsHardwareOffloadEnabled bool, disableTXChecksumOffload bool) (*ifConfigurator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ic *ifConfigurator) addEndpoint(ep *hcsshim.HNSEndpoint) { _ = "STUB: not implemented"; return }

func (ic *ifConfigurator) getEndpoint(name string) (*hcsshim.HNSEndpoint, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (ic *ifConfigurator) delEndpoint(name string) { _ = "STUB: not implemented"; return }

// findContainerIPConfig finds a valid IPv4 address since IPv6 is not supported for Windows at this stage.
func findContainerIPConfig(ips []*current.IPConfig) (*current.IPConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SR-IOV is not supported on Windows.
func (ic *ifConfigurator) recoverVFInterfaceName(containerNetNS string, containerIfaceName string) error {
	_ = "STUB: not implemented"
	return nil
}

// configureContainerLink creates a HNSEndpoint for the container using the IPAM result, and then attach it on the container interface.
func (ic *ifConfigurator) configureContainerLink(
	podName string,
	podNamespace string,
	containerID string,
	containerNetNS string,
	containerIFDev string,
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

// We must use the infra container to generate the endpoint name to ensure infra and workload containers use the
// same HNSEndpoint.

// Search endpoint from local cache.

// Only create HNS Endpoint for infra container.

// Attach HNSEndpoint to the container. Note that HNSEndpoint must be attached to the container before adding OVS port,
// otherwise an error will be returned when creating OVS port.

// Update IPConfig with the index of target interface in the result. The index is used in CNI CmdCheck.

// MTU is configured only when the infrastructure container is created.

// Configure MTU in another separate goroutine to ensure it is executed after the host interface is created.
// The reasons include, 1) for containerd runtime, the interface is created by containerd after the CNI
// CmdAdd request is returned; 2) for Docker runtime, the interface is created after hcsshim.HotAttachEndpoint,
// and the hcsshim call is not synchronized from the observation.

// changeContainerMTU is only used for Antrea Multi-cluster with networkPolicyOnly
// mode, and this mode doesn't support Windows platform yet.
func (ic *ifConfigurator) changeContainerMTU(containerNetNS string, containerIFDev string, mtuDeduction int) error {
	_ = "STUB: not implemented"
	return nil
}

// createContainerLink creates HNSEndpoint using the IP configuration in the IPAM result.
func (ic *ifConfigurator) createContainerLink(endpointName string, result *current.Result, containerID, podName, podNamespace string) (hostLink *hcsshim.HNSEndpoint, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add the new created Endpoint into local cache.

// attachContainerLink takes the result of the IPAM plugin, and adds the appropriate IP
// addresses and routes to the interface.
// For different CRI runtimes we need to use the appropriate Windows container API:
//   - containerd runtime: HCS API
func attachContainerLink(ep *hcsshim.HNSEndpoint, containerID, sandbox, containerIFDev string) (*current.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func attachEndpointInNamespace(hcnEp *hcn.HostComputeEndpoint, sandbox string) error {
	_ = "STUB: not implemented"
	return nil
}

// advertiseContainerAddr returns immediately as the address is advertised automatically after it is configured on an
// network interface on Windows.
func (ic *ifConfigurator) advertiseContainerAddr(containerNetNS string, containerIfaceName string, result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// removeContainerLink removes the HNSEndpoint attached on the Pod.
func (ic *ifConfigurator) removeContainerLink(containerID, epName string) error {
	_ = "STUB: not implemented"
	return nil
}

// removeHNSEndpoint removes the HNSEndpoint from HNS and local cache.
func (ic *ifConfigurator) removeHNSEndpoint(endpoint *hcsshim.HNSEndpoint, containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove HNSEndpoint.

// Deleting HNS Endpoint is blocking in some corner cases. It might be a bug in Windows HNS service. To avoid
// hanging in cniserver, add timeout control in HNSEndpoint deletion.

// Delete HNSEndpoint from local cache.

func deleteHnsEndpoint(endpoint *hcsshim.HNSEndpoint) (*hcsshim.HNSEndpoint, error) {
	_ = "STUB: not implemented"
	return nil,

		// isValidHostNamespace checks if the hostNamespace is valid or not. When using Docker runtime, the hostNamespace
		// is not set, and Windows HCN should use a default value "00000000-0000-0000-0000-000000000000". An error returns
		// when removing HostComputeEndpoint in this namespace. This field is set with a valid value when containerd is used.
		nil
}

func isValidHostNamespace(hostNamespace string) bool { _ = "STUB: not implemented"; return false }

func parseContainerIfaceFromResults(cfgArgs *cnipb.CniCmdArgs, prevResult *current.Result) *current.Interface {
	_ = "STUB: not implemented"
	return nil
}

// checkContainerInterface finds the virtual interface of the container, and compares the network configurations with
// the previous result.
func (ic *ifConfigurator) checkContainerInterface(
	sandboxID, containerID string,
	containerIface *current.Interface,
	containerIPs []*current.IPConfig,
	containerRoutes []*cnitypes.Route,
	sriovVFDeviceID string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check container sandbox configuration.

// Check container MAC configuration.

// Parse container IP configuration from previous result.

// Check container IP configuration.

// Todo: add check for container route configuration.

func getNetInterfaceAddrs(intf *net.Interface) ([]net.Addr, error) {
	_ = "STUB: not implemented"
	return nil,

		// validateExpectedInterfaceIPs checks if the vNIC for the container has configured with correct IP address.
		nil
}

func validateExpectedInterfaceIPs(containerIPConfig *current.IPConfig, intf *net.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

func (ic *ifConfigurator) validateVFRepInterface(sriovVFDeviceID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// validateContainerPeerInterface checks HNSEndpoint configuration.
func (ic *ifConfigurator) validateContainerPeerInterface(interfaces []*current.Interface, containerVeth *vethPair) (*vethPair, error) {
	_ = "STUB: not implemented"
	// Iterate all the passed interfaces and look up the host interface by
	// matching the veth peer interface index.
	return nil, nil
}

// Not in the default Namespace. Must be the container interface.

// getInterceptedInterfaces is not supported on Windows.
func (ic *ifConfigurator) getInterceptedInterfaces(
	sandbox string,
	containerNetNS string,
	containerIFDev string,
) (*current.Interface, *current.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (ic *ifConfigurator) addPostInterfaceCreateHook(containerID, endpointName string, containerAccess *containerAccessArbitrator, hook postInterfaceCreateHook) error {
	_ = "STUB: not implemented"
	return nil
}

func createHnsEndpoint(epRequest *hcsshim.HNSEndpoint) (*hcsshim.HNSEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
