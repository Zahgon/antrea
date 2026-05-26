// Copyright 2023 Antrea Authors
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

package podwatch

import (
	"time"

	corev1 "k8s.io/api/core/v1"

	// Version v1 of the kubelet API was introduced in K8s v1.20.
	// Using version v1alpha1 instead to support older K8s versions.
	current "github.com/containernetworking/cni/pkg/types/100"
	netdefv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"

	"antrea.io/antrea/v2/pkg/agent/interfacestore"
)

const (
	kubeletPodResourcesPath = "/var/lib/kubelet/pod-resources"
	kubeletSocket           = "kubelet.sock"
	listTimeout             = 10 * time.Second
)

var (
	// getPodContainerDeviceIDsFn is used to retrieve SRIOV device IDs
	// assigned to a specific Pod. It can be overridden by unit tests.
	getPodContainerDeviceIDsFn = getPodContainerDeviceIDs
)

// Structure to associate a unique VF's PCI Address to the Linux ethernet interface.
type podSriovVFDeviceIDInfo struct {
	resourceName string
	vfDeviceID   string
	ifName       string
}

// getPodContainerDeviceIDs returns the device IDs assigned to a Pod's containers.
func getPodContainerDeviceIDs(podName string, podNamespace string) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildVFDeviceIDListPerPod is a helper function to build a cache structure with the
// list of all the PCI addresses allocated per Pod based on their resource requests (in Pod spec).
// When there is a request for a VF resource (to associate it for a secondary network interface),
// assignUnusedSriovVFDeviceID will use this cache information to pick up a unique PCI address
// which is still not associated with a network device name.
// NOTE: buildVFDeviceIDListPerPod is called only if a Pod specific VF to Interface mapping cache
// was not build earlier. Sample initial entry per Pod: "{18:01.1,""},{18:01.2,""},{18:01.3,""}"
func (pc *PodController) buildVFDeviceIDListPerPod(podName, podNamespace string) ([]podSriovVFDeviceIDInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we will set this field when allocating the device

func (pc *PodController) deleteVFDeviceIDListPerPod(podName, podNamespace string) {
	_ = "STUB: not implemented"
	return
}

func (pc *PodController) releaseSriovVFDeviceID(podName, podNamespace, interfaceName string) {
	_ = "STUB: not implemented"
	return
}

func (pc *PodController) assignSriovVFDeviceID(podName, podNamespace, resourceName, interfaceName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// remember the first match of unused PCI address

// Update the cache entry

// Configure SRIOV VF as a Secondary Network Interface.
func (pc *PodController) configureSriovAsSecondaryInterface(
	pod *corev1.Pod,
	network *netdefv1.NetworkSelectionElement,
	resourceName string,
	podCNIInfo *podCNIInfo,
	mtu int,
	result *current.Result,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PodController) deleteSriovSecondaryInterface(interfaceConfig *interfacestore.InterfaceConfig) error {
	_ = "STUB: not implemented"
	// NOTE: SR-IOV VF interface clean-up will be handled by SR-IOV device plugin. The interface
	// is not deleted here.
	return nil
}

// AllowCNIDelete in SecondaryNetwork indicates if a Pod's SR-IOV devices are all detached
// and CNI deletion can be processed to remove the Pod's network namespace.
func (pc *PodController) AllowCNIDelete(podName, podNamespace string) bool {
	_ = "STUB: not implemented"
	return false
}

// SR-IOV VF device found.
