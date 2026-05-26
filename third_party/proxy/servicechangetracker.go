/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
/*
// Copyright 2025 Antrea Authors
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

Original file https://raw.githubusercontent.com/kubernetes/kubernetes/refs/tags/v1.34.2/pkg/proxy/servicechangetracker.go

Modifies:

- Remove import "k8s.io/kubernetes/pkg/proxy/metrics" and its usages.
- Replace import from proxyutil "k8s.io/kubernetes/pkg/proxy/util" to proxyutil "antrea.io/antrea/v2/third_party/proxy/util".

Adds:

- Add "serviceLabelSelector labels.Selector" and "skipServices sets.Set[string]" to "ServiceChangeTracker".
- Add parameters to `NewServiceChangeTracker` for initializing "serviceLabelSelector labels.Selector" and
  "skipServices sets.Set[string]".

*/

package proxy

import (
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
)

// ServiceChangeTracker carries state about uncommitted changes to an arbitrary number of
// Services, keyed by their namespace and name.
type ServiceChangeTracker struct {
	// lock protects items.
	lock sync.Mutex
	// items maps a service to its serviceChange.
	items map[types.NamespacedName]*serviceChange

	// makeServiceInfo allows the proxier to inject customized information when
	// processing services.
	makeServiceInfo makeServicePortFunc
	// processServiceMapChange is invoked by the apply function on every change. This
	// function should not modify the ServicePortMaps, but just use the changes for
	// any Proxier-specific cleanup.
	processServiceMapChange processServiceMapChangeFunc

	ipFamily v1.IPFamily

	serviceLabelSelector labels.Selector
	// skipServices indicates the service list for which we should skip proxying
	// it will be initialized from antrea-agent.conf
	skipServices sets.Set[string]
}

type makeServicePortFunc func(*v1.ServicePort, *v1.Service, *BaseServicePortInfo) ServicePort
type processServiceMapChangeFunc func(previous, current ServicePortMap)

// serviceChange contains all changes to services that happened since proxy rules were synced.  For a single object,
// changes are accumulated, i.e. previous is state from before applying the changes,
// current is state after applying all of the changes.
type serviceChange struct {
	previous ServicePortMap
	current  ServicePortMap
}

// NewServiceChangeTracker initializes a ServiceChangeTracker
func NewServiceChangeTracker(ipFamily v1.IPFamily, makeServiceInfo makeServicePortFunc, processServiceMapChange processServiceMapChangeFunc, serviceLabelSelector labels.Selector, skipServices []string) *ServiceChangeTracker {
	_ = "STUB: not implemented"
	return nil
}

// Update updates the ServiceChangeTracker based on the <previous, current> service pair
// (where either previous or current, but not both, can be nil). It returns true if sct
// contains changes that need to be synced (whether or not those changes were caused by
// this update); note that this is different from the return value of
// EndpointChangeTracker.EndpointSliceUpdate().
func (sct *ServiceChangeTracker) Update(previous, current *v1.Service) bool {
	_ = "STUB: not implemented"
	// This is unexpected, we should return false directly.
	return false
}

// if change.previous equal to change.current, it means no change

// ServicePortMap maps a service to its ServicePort.
type ServicePortMap map[ServicePortName]ServicePort

// UpdateServiceMapResult is the updated results after applying service changes.
type UpdateServiceMapResult struct {
	// UpdatedServices lists the names of all services added/updated/deleted since the
	// last Update.
	UpdatedServices sets.Set[types.NamespacedName]
}

// HealthCheckNodePorts returns a map of Service names to HealthCheckNodePort values
// for all Services in sm with non-zero HealthCheckNodePort.
func (sm ServicePortMap) HealthCheckNodePorts() map[types.NamespacedName]uint16 {
	_ = "STUB: not implemented"
	// TODO: If this will appear to be computationally expensive, consider
	// computing this incrementally similarly to svcPortMap.
	return nil
}

// serviceToServiceMap translates a single Service object to a ServicePortMap.
//
// NOTE: service object should NOT be modified.
func (sct *ServiceChangeTracker) serviceToServiceMap(service *v1.Service) ServicePortMap {
	_ = "STUB: not implemented"
	return *new(ServicePortMap)
}

// Update updates ServicePortMap base on the given changes, returns information about the
// diff since the last Update, triggers processServiceMapChange on every change, and
// clears the changes map.
func (sm ServicePortMap) Update(sct *ServiceChangeTracker) UpdateServiceMapResult {
	_ = "STUB: not implemented"
	return *new(UpdateServiceMapResult)
}

// filter out the Update event of current changes from previous changes
// before calling unmerge() so that can skip deleting the Update events.

// clear changes after applying them to ServicePortMap.

// merge adds other ServicePortMap's elements to current ServicePortMap.
// If collision, other ALWAYS win. Otherwise add the other to current.
// In other words, if some elements in current collisions with other, update the current by other.
func (sm *ServicePortMap) merge(other ServicePortMap) { _ = "STUB: not implemented"; return }

// filter filters out elements from ServicePortMap base on given ports string sets.
func (sm *ServicePortMap) filter(other ServicePortMap) { _ = "STUB: not implemented"; return }

// skip the delete for Update event.

// unmerge deletes all other ServicePortMap's elements from current ServicePortMap.
func (sm *ServicePortMap) unmerge(other ServicePortMap) { _ = "STUB: not implemented"; return }
