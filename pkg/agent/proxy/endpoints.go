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

package proxy

import (
	"sync"

	v1 "k8s.io/api/core/v1"
	discovery "k8s.io/api/discovery/v1"

	k8sproxy "antrea.io/antrea/v2/third_party/proxy"
)

// endpointsChangesTracker tracks Endpoints changes.
type endpointsChangesTracker struct {
	sync.RWMutex
	// initialized tells whether Endpoints have been synced.
	initialized bool

	tracker *k8sproxy.EndpointsChangeTracker
}

func newEndpointsChangesTracker(hostname string, ipFamily v1.IPFamily) *endpointsChangesTracker {
	_ = "STUB: not implemented"
	return nil
}

func (t *endpointsChangesTracker) OnEndpointsSynced() { _ = "STUB: not implemented"; return }

func (t *endpointsChangesTracker) OnEndpointSliceUpdate(endpointSlice *discovery.EndpointSlice, removeSlice bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *endpointsChangesTracker) Synced() bool { _ = "STUB: not implemented"; return false }

// Update updates an EndpointsMap and numLocalEndpoints based on current changes.
func (t *endpointsChangesTracker) Update(em k8sproxy.EndpointsMap) k8sproxy.UpdateEndpointsMapResult {
	_ = "STUB: not implemented"
	return *new(k8sproxy.UpdateEndpointsMapResult)
}
