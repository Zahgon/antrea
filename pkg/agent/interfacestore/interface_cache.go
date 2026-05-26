// Copyright 2019 Antrea Authors
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

package interfacestore

import (
	"k8s.io/client-go/tools/cache"
)

const (
	// interfaceNameIndex is the index built with InterfaceConfig.InterfaceName.
	interfaceNameIndex = "interfaceName"
	// interfaceTypeIndex is the index built with InterfaceConfig.Type.
	interfaceTypeIndex = "interfaceType"
	// containerIDIndex is the index built with InterfaceConfig.ContainerID.
	// Only container interfaces will be indexed.
	// One containerID should get at most one interface in theory.
	containerIDIndex = "containerID"
	// podIndex is the index built with InterfaceConfig.PodNamespace + Podname.
	// Only container interfaces will be indexed.
	// One Pod may get more than one interface.
	podIndex = "pod"
	// interfaceIPIndex is the index built with InterfaceConfig.IP
	// Only the interfaces with IP get indexed.
	interfaceIPIndex = "ip"
	// ofPortIndex is the index built with InterfaceConfig.OFPort
	ofPortIndex = "ofPort"
	// externalEntityIndex is the index built with InterfaceConfig.EntityNamespace + EntityName.
	// Only the interfaces of an ExternalEntity get indexed.
	externalEntityIndex = "externalEntity"
)

// Local cache for interfaces created on node, including container, host gateway, and tunnel
// ports, `Type` field is used to differentiate interface category.
//  1) For container interface, the fields should include: containerID, podName, Namespace,
//     IP, MAC, and OVS Port configurations.
//  2) For host gateway port, the fields should include: name, IP, MAC, and OVS port
//     configurations.
//  3) For tunnel port, the fields include: name and tunnel type; and for an IPsec tunnel,
//     additionally: remoteIP, PSK and remote Node name.
// OVS Port configurations include PortUUID and OFPort.
// Container interface is added into cache after invocation of cniserver.CmdAdd, and removed
// from cache after invocation of cniserver.CmdDel. For cniserver.CmdCheck, the server would
// check previousResult with local cache.
// Host gateway and the default tunnel interfaces are added into cache in node initialization
// phase or retrieved from existing OVS ports.
// An IPsec tunnel interface is added into the cache when IPsec encryption is enabled, and
// NodeRouteController watches a new remote Node from K8s API, and is removed when the remote
// Node is deleted.
// Todo: add periodic task to sync local cache with container veth pair

type interfaceCache struct {
	cache cache.Indexer
}

func (c *interfaceCache) Initialize(interfaces []*InterfaceConfig) {
	_ = "STUB: not implemented"
	return
}

// getInterfaceKey returns the key to access interfaceConfig from the cache.
// It implements cache.KeyFunc.
func getInterfaceKey(obj interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

// IPsec tunnel interface for a Node.

// Use the interface name as the key by default.

// AddInterface adds interfaceConfig into local cache.
func (c *interfaceCache) AddInterface(interfaceConfig *InterfaceConfig) {
	_ = "STUB: not implemented"
	return
}

// UpdateInterface updates interfaceConfig into local cache.
func (c *interfaceCache) UpdateInterface(interfaceConfig *InterfaceConfig) {
	_ = "STUB: not implemented"
	return
}

// DeleteInterface deletes interface from local cache.
func (c *interfaceCache) DeleteInterface(interfaceConfig *InterfaceConfig) {
	_ = "STUB: not implemented"
	return
}

// GetInterface retrieves interface from local cache given the interface key.
func (c *interfaceCache) GetInterface(interfaceKey string) (*InterfaceConfig, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// ListInterfacesByType lists all interfaces from local cache.
func (c *interfaceCache) ListInterfaces() []*InterfaceConfig { _ = "STUB: not implemented"; return nil }

// GetInterfaceByName retrieves interface from local cache given the interface
// name.
func (c *interfaceCache) GetInterfaceByName(interfaceName string) (*InterfaceConfig, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetInterfaceByIP retrieves interface from local cache given the interface IP.
func (c *interfaceCache) GetInterfaceByIP(interfaceIP string) (*InterfaceConfig, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *interfaceCache) GetContainerInterfaceNum() int { _ = "STUB: not implemented"; return 0 }

func (c *interfaceCache) GetInterfacesByType(interfaceType InterfaceType) []*InterfaceConfig {
	_ = "STUB: not implemented"
	return nil
}

func (c *interfaceCache) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *interfaceCache) GetInterfaceKeysByType(interfaceType InterfaceType) []string {
	_ = "STUB: not implemented"
	return nil
}

// GetContainerInterface retrieves InterfaceConfig by the given container ID.
func (c *interfaceCache) GetContainerInterface(containerID string) (*InterfaceConfig, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *interfaceCache) GetInterfacesByEntity(name, namespace string) []*InterfaceConfig {
	_ = "STUB: not implemented"
	return nil
}

// GetContainerInterfacesByPod retrieves InterfaceConfigs for the Pod.
// It's possible that more than one container interface (with different containerIDs) has the same Pod namespace and
// name temporarily when the previous Pod is being deleted and the new Pod is being created almost simultaneously.
// https://github.com/antrea-io/antrea/issues/785#issuecomment-642051884
func (c *interfaceCache) GetContainerInterfacesByPod(podName string, podNamespace string) []*InterfaceConfig {
	_ = "STUB: not implemented"
	return nil
}

// GetNodeTunnelInterface retrieves InterfaceConfig for the tunnel to the Node.
func (c *interfaceCache) GetNodeTunnelInterface(nodeName string) (*InterfaceConfig, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetInterfaceByOFPort retrieves InterfaceConfig by the given ofPort number.
func (c *interfaceCache) GetInterfaceByOFPort(ofPort uint32) (*InterfaceConfig, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func interfaceNameIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func interfaceTypeIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func containerIDIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func podIndexFunc(obj interface{}) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func interfaceIPIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If interfaceConfig IP is not set, we return empty key.

func interfaceOFPortIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OVSPortConfig can be nil for a secondary SR-IOV interface.

// If interfaceConfig OFport is not valid, we return empty key.

func externalEntityIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewInterfaceStore() InterfaceStore { _ = "STUB: not implemented"; return *new(InterfaceStore) }
