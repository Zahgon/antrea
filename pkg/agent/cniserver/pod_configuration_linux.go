//go:build !windows
// +build !windows

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

	current "github.com/containernetworking/cni/pkg/types/100"
	"k8s.io/client-go/tools/cache"

	"antrea.io/antrea/v2/pkg/agent/cniserver/ipam"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
)

// connectInterfaceToOVS connects an existing interface to the OVS bridge.
func (pc *podConfigurator) connectInterfaceToOVS(
	podName, podNamespace, containerID, netNS string,
	hostIface, containerIface *current.Interface,
	ips []*current.IPConfig,
	vlanID uint16,
	containerAccess *containerAccessArbitrator) (*interfacestore.InterfaceConfig, error) {
	_ = "STUB: not implemented"
	// Use the outer veth interface name as the OVS port name.
	return nil, nil
}

// Create an OVS Port and add container configuration into external_ids.

// Remove OVS port if any failure occurs in later manipulation.

// Not needed for a secondary network interface.

// GetOFPort will wait for up to 1 second for OVSDB to report the OFPort number.

// Add containerConfig into local cache

// Not needed for a secondary network interface.

// Notify the Pod update event to required components.

func (pc *podConfigurator) configureInterfaces(
	podName, podNamespace, containerID, containerNetNS string,
	containerIFDev string, mtu int, sriovVFDeviceID string,
	result *ipam.IPAMResult, createOVSPort bool, containerAccess *containerAccessArbitrator, mac net.HardwareAddr) error {
	_ = "STUB: not implemented"
	return nil
}

// reconcileMissingPods is never called on Linux, see reconcile logic.
func (pc *podConfigurator) reconcileMissingPods(ifConfigs []*interfacestore.InterfaceConfig, containerAccess *containerAccessArbitrator) {
	_ = "STUB: not implemented"

	// isInterfaceInvalid returns true if the OVS interface's ofport is "-1" which means the host interface is disconnected.
	return
}

func (pc *podConfigurator) isInterfaceInvalid(ifaceConfig *interfacestore.InterfaceConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func (pc *podConfigurator) initPortStatusMonitor(_ cache.SharedIndexInformer) {
	_ = "STUB: not implemented"
	return
}

func (pc *podConfigurator) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }
