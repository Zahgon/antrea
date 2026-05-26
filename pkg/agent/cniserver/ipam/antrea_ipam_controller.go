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

package ipam

import (
	"net"

	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"

	crdv1b1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	clientsetversioned "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1beta1"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/ipam/poolallocator"
)

const (
	controllerName = "AntreaIPAMController"
	// Pod index name for IPPool cache.
	podIndex = "pod"
)

// Antrea IPAM Controller maintains map of Namespace annotations using
// Namespace informer. In future, which Antrea IPAM support expands,
// this controller can be used to store annotations for other objects,
// such as Statefulsets.
type AntreaIPAMController struct {
	crdClient         clientsetversioned.Interface
	ipPoolInformer    crdinformers.IPPoolInformer
	ipPoolLister      crdlisters.IPPoolLister
	namespaceInformer coreinformers.NamespaceInformer
	namespaceLister   corelisters.NamespaceLister
	podInformer       cache.SharedIndexInformer
	podLister         corelisters.PodLister
}

func podIndexFunc(obj interface{}) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func InitializeAntreaIPAMController(crdClient clientsetversioned.Interface,
	namespaceInformer coreinformers.NamespaceInformer,
	ipPoolInformer crdinformers.IPPoolInformer,
	podInformer cache.SharedIndexInformer, ipamAnnotations bool) (*AntreaIPAMController, error) {
	_ = "STUB: not implemented"
	// Order of init causes antreaIPAMDriver to be initialized first
	// After controller is initialized by agent init, we need to make it
	// know to the driver
	return nil, nil
}

// Create podInformer/Lister and namespaceInformer/Lister if need to read the AntreaIPAM
// annotation on Pods and Namespaces.

// Run starts to watch and process Namespace updates for the Node where Antrea Agent
// is running, and maintain a mapping between Namespace name and IPAM annotations.
func (c *AntreaIPAMController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Look up IPPools from the Pod annotation.
func (c *AntreaIPAMController) getIPPoolsByPod(namespace, name string) ([]string, []net.IP, *crdv1b1.IPAddressOwner, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// Find IPPool by Namespace

// Collect specified IPs from the AntreaIPAMPodIP annotation.
// Multiple IPs (comma-separated) may be provided, but the caller
// (AntreaIPAM.Add) will only consume at most one IPv4 and one IPv6
// address. Each IP must belong to the first Pool of the corresponding
// IP family listed in the AntreaIPAM annotation.

// Parse StatefulSet name/index from Pod name

// This should not occur unless user creates an invalid Pod manually

// getPoolAllocatorsByPod looks up IPPools from the Pod annotation and returns
// allocators for all valid pools. This supports IPv4, IPv6, and dual-stack
// configurations where multiple pools (one per IP family) may be specified.
func (c *AntreaIPAMController) getPoolAllocatorsByPod(namespace, podName string) (mineType, []*poolallocator.IPPoolAllocator, []net.IP, *crdv1b1.IPAddressOwner, error) {
	_ = "STUB: not implemented"
	return *new(mineType), nil, nil, nil, nil
}

// Look up IPPools by matching PodOwner.
func (c *AntreaIPAMController) getPoolAllocatorsByOwner(podOwner *crdv1b1.PodOwner) ([]*poolallocator.IPPoolAllocator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *AntreaIPAMController) getPoolAllocatorByName(poolName string) (*poolallocator.IPPoolAllocator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
