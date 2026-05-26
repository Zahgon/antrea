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

package types

import (
	"sync"

	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/agent/openflow"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
	k8sproxy "antrea.io/antrea/v2/third_party/proxy"
)

// GroupCounter generates and manages global unique group ID.
type GroupCounter interface {
	// AllocateIfNotExist generates a global unique group ID for a Service if the group ID has not been generated, then
	// return the group ID (newly allocated or already allocated).
	AllocateIfNotExist(svcPortName k8sproxy.ServicePortName, isEndpointsLocal bool) binding.GroupIDType
	// Get gets the group ID for the Service.
	Get(svcPortName k8sproxy.ServicePortName, isEndpointsLocal bool) (binding.GroupIDType, bool)
	// Recycle removes the Service group ID mapping. The recycled group ID can be reused.
	Recycle(svcPortName k8sproxy.ServicePortName, isEndpointsLocal bool) bool
	// GetAllGroupIDs gets all group IDs related to the Service.
	GetAllGroupIDs(svcNamespacedName string) []binding.GroupIDType
}

type groupCounter struct {
	mu             sync.Mutex
	groupAllocator openflow.GroupAllocator
	groupIDUpdates chan<- string

	servicePortNamesMap map[string]sets.Set[string]
	groupMap            map[string]binding.GroupIDType
}

func NewGroupCounter(groupAllocator openflow.GroupAllocator, groupIDUpdates chan<- string) *groupCounter {
	_ = "STUB: not implemented"
	return nil
}

func keyString(svcPortName k8sproxy.ServicePortName, isEndpointsLocal bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *groupCounter) updateServicePortNameMap(svcNamespacedName string, svcKeyString string) {
	_ = "STUB: not implemented"
	return
}

func (c *groupCounter) deleteServicePortNameMap(svcNamespacedName string, svcKeyString string) {
	_ = "STUB: not implemented"
	return
}

func (c *groupCounter) AllocateIfNotExist(svcPortName k8sproxy.ServicePortName, isEndpointsLocal bool) binding.GroupIDType {
	_ = "STUB: not implemented"
	return *new(binding.GroupIDType)
}

func (c *groupCounter) Get(svcPortName k8sproxy.ServicePortName, isEndpointsLocal bool) (binding.GroupIDType, bool) {
	_ = "STUB: not implemented"
	return *new(binding.GroupIDType), false
}

func (c *groupCounter) Recycle(svcPortName k8sproxy.ServicePortName, isEndpointsLocal bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *groupCounter) GetAllGroupIDs(svcNamespacedName string) []binding.GroupIDType {
	_ = "STUB: not implemented"
	return nil
}
