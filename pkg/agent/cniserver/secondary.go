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

package cniserver

import (
	"net"

	current "github.com/containernetworking/cni/pkg/types/100"

	"antrea.io/antrea/v2/pkg/agent/cniserver/ipam"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
)

func NewSecondaryInterfaceConfigurator(ovsBridgeClient ovsconfig.OVSBridgeClient, interfaceStore interfacestore.InterfaceStore) (*podConfigurator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConfigureSriovSecondaryInterface configures a SR-IOV secondary interface for a Pod.
func (pc *podConfigurator) ConfigureSriovSecondaryInterface(
	podName, podNamespace string,
	containerID, containerNetNS, containerInterfaceName string,
	mtu int,
	podSriovVFDeviceID string,
	result *current.Result) error {
	_ = "STUB: not implemented"
	return nil
}

// Use podSriovVFDeviceID as the interface name in the interface store.

// DeleteSriovSecondaryInterface deletes a SRIOV secondary interface.
func (pc *podConfigurator) DeleteSriovSecondaryInterface(interfaceConfig *interfacestore.InterfaceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// ConfigureVLANSecondaryInterface configures a VLAN secondary interface on the secondary network
// OVS bridge, and returns the OVS port UUID.
func (pc *podConfigurator) ConfigureVLANSecondaryInterface(
	podName, podNamespace string,
	containerID, containerNetNS, containerInterfaceName string,
	mtu int, ipamResult *ipam.IPAMResult, mac net.HardwareAddr) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteVLANSecondaryInterface deletes a VLAN secondary interface.
func (pc *podConfigurator) DeleteVLANSecondaryInterface(interfaceConfig *interfacestore.InterfaceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// No retry for interface deletion.
