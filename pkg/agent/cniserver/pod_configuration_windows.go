//go:build windows
// +build windows

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
	"time"

	current "github.com/containernetworking/cni/pkg/types/100"
	"k8s.io/client-go/tools/cache"

	"antrea.io/antrea/v2/pkg/agent/cniserver/ipam"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
)

var (
	workerName = "podConfigurator"
)

const (
	podNotReadyTime        = 30 * time.Second
	ovsInterfaceTypeForPod = "internal"
)

// connectInterfaceToOVSAsync waits for an interface to be created and connects it to OVS br-int asynchronously
// in another goroutine. The function is for containerd runtime. The host interface is created after
// CNI call completes.
func (pc *podConfigurator) connectInterfaceToOVSAsync(ifConfig *interfacestore.InterfaceConfig, containerAccess *containerAccessArbitrator) error {
	_ = "STUB: not implemented"
	return nil
}

// Add the OVS port into the queue after 30s in case the OFPort is still not ready. This
// operation is performed before we update OVSDB, otherwise we
// need to think about the race condition between the current goroutine with the listener.
// It may generate a duplicated PodIsReady event if the Pod's OpenFlow entries are installed
// before the time, then the library shall merge the event.

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

// The container interface is created after the CNI returns the network setup result.
// Because of this, we need to wait asynchronously for the interface to be created: we create the OVS port
// and set the OVS Interface type "" first, and change the OVS Interface type to "internal" to connect to the
// container interface after it is created. After OVS connects to the container interface, an OFPort is allocated.

// Add containerConfig into local cache

func (pc *podConfigurator) configureInterfaces(
	podName, podNamespace, containerID, containerNetNS string,
	containerIFDev string, mtu int, sriovVFDeviceID string,
	result *ipam.IPAMResult, createOVSPort bool, containerAccess *containerAccessArbitrator, mac net.HardwareAddr) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if the OVS configurations for the container exists or not. If yes, return
// immediately. This check is used on Windows, as kubelet on Windows will call CNI ADD
// multiple times for the infrastructure container to query IP of the Pod. But there should
// be only one OVS port created for the same Pod (identified by its sandbox container ID),
// and if the OVS port is added more than once, OVS will return an error.
// See: https://github.com/kubernetes/kubernetes/issues/57253#issuecomment-358897721.

// isInterfaceInvalid returns false because we now don't support detecting the disconnected host interface on Windows
// due to the OVS issue (https://github.com/openvswitch/ovs-issues/issues/353), by which we can't differentiate from
// the case that a Pod's host interface is created during agent downtime and is expected to re-connect after agent
// is restarted.
func (pc *podConfigurator) isInterfaceInvalid(ifaceConfig *interfacestore.InterfaceConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func (pc *podConfigurator) reconcileMissingPods(ifConfigs []*interfacestore.InterfaceConfig, containerAccess *containerAccessArbitrator) {
	_ = "STUB: not implemented"
	return
}

// initPortStatusMonitor has subscribed a channel to listen for the OpenFlow PortStatus message, and it also
// initiates the Pod recorder.
func (pc *podConfigurator) initPortStatusMonitor(podInformer cache.SharedIndexInformer) {
	_ = "STUB: not implemented"
	return
}

func (pc *podConfigurator) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Update Pod OpenFlow entries only after the OpenFlow port state is live.

// worker is a long-running function that will continually call the processNextWorkItem function in
// order to read and process a message on the workqueue.
func (pc *podConfigurator) worker() { _ = "STUB: not implemented"; return }
