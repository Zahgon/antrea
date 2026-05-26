/*
Copyright 2019 The Kubernetes Authors.

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

Original file https://raw.githubusercontent.com/kubernetes/kubernetes/refs/tags/v1.34.2/pkg/proxy/endpointslicecache.go

Modifies:

- Replace import utilfeature "k8s.io/apiserver/pkg/util/feature" with "antrea.io/antrea/v2/pkg/features".
- Remove import "k8s.io/kubernetes/pkg/features".
- Remove import "sort".
- Change type of "EndpointsMap" from "map[ServicePortName][]Endpoint" to "map[ServicePortName]map[string]Endpoint".

*/

package proxy

import (
	"sync"

	discovery "k8s.io/api/discovery/v1"
	"k8s.io/apimachinery/pkg/types"
)

// EndpointSliceCache is used as a cache of EndpointSlice information.
type EndpointSliceCache struct {
	// lock protects trackerByServiceMap.
	lock sync.Mutex

	// trackerByServiceMap is the basis of this cache. It contains endpoint
	// slice trackers grouped by service name and endpoint slice name. The first
	// key represents a namespaced service name while the second key represents
	// an endpoint slice name. Since endpoints can move between slices, we
	// require slice specific caching to prevent endpoints being removed from
	// the cache when they may have just moved to a different slice.
	trackerByServiceMap map[types.NamespacedName]*endpointSliceTracker

	makeEndpointInfo makeEndpointFunc
	nodeName         string
}

// endpointSliceTracker keeps track of EndpointSlices as they have been applied
// by a proxier along with any pending EndpointSlices that have been updated
// in this cache but not yet applied by a proxier.
type endpointSliceTracker struct {
	applied endpointSliceDataByName
	pending endpointSliceDataByName
}

// endpointSliceDataByName groups endpointSliceData by the names of the
// corresponding EndpointSlices.
type endpointSliceDataByName map[string]*endpointSliceData

// endpointSliceData contains information about a single EndpointSlice update or removal.
type endpointSliceData struct {
	endpointSlice *discovery.EndpointSlice
	remove        bool
}

// NewEndpointSliceCache initializes an EndpointSliceCache.
func NewEndpointSliceCache(nodeName string, makeEndpointInfo makeEndpointFunc) *EndpointSliceCache {
	_ = "STUB: not implemented"
	return nil
}

// newEndpointSliceTracker initializes an endpointSliceTracker.
func newEndpointSliceTracker() *endpointSliceTracker { _ = "STUB: not implemented"; return nil }

// standardEndpointInfo is the default makeEndpointFunc.
func standardEndpointInfo(ep *BaseEndpointInfo, _ *ServicePortName) Endpoint {
	_ = "STUB: not implemented"

	// updatePending updates a pending slice in the cache.
	return *new(Endpoint)
}

func (cache *EndpointSliceCache) updatePending(endpointSlice *discovery.EndpointSlice, remove bool) bool {
	_ = "STUB: not implemented"
	return false
}

// checkoutChanges returns a map of all endpointsChanges that are
// pending and then marks them as applied.
func (cache *EndpointSliceCache) checkoutChanges() map[types.NamespacedName]*endpointsChange {
	_ = "STUB: not implemented"
	return nil
}

// spToEndpointMap stores groups Endpoint objects by ServicePortName and
// endpoint string (returned by Endpoint.String()).
type spToEndpointMap map[ServicePortName]map[string]Endpoint

// getEndpointsMap computes an EndpointsMap for a given set of EndpointSlices.
func (cache *EndpointSliceCache) getEndpointsMap(serviceNN types.NamespacedName, sliceDataByName endpointSliceDataByName) EndpointsMap {
	_ = "STUB: not implemented"
	return *new(EndpointsMap)
}

// endpointInfoByServicePort groups endpoint info by service port name and address.
func (cache *EndpointSliceCache) endpointInfoByServicePort(serviceNN types.NamespacedName, sliceDataByName endpointSliceDataByName) spToEndpointMap {
	_ = "STUB: not implemented"
	return *new(spToEndpointMap)
}

// TODO: handle nil ports to mean "all"

// addEndpoints adds an Endpoint for each unique endpoint.
func (cache *EndpointSliceCache) addEndpoints(svcPortName *ServicePortName, portNum int, endpointSet map[string]Endpoint, endpoints []discovery.Endpoint) map[string]Endpoint {
	_ = "STUB: not implemented"
	return nil
}

// iterate through endpoints to add them to endpointSet.

// This logic ensures we're deduplicating potential overlapping endpoints
// isLocal should not vary between matching endpoints, but if it does, we
// favor a true value here if it exists.

func (cache *EndpointSliceCache) isLocal(nodeName string) bool {
	_ = "STUB: not implemented"
	return false
}

// esDataChanged returns true if the esData parameter should be set as a new
// pending value in the cache.
func (cache *EndpointSliceCache) esDataChanged(serviceKey types.NamespacedName, sliceKey string, esData *endpointSliceData) bool {
	_ = "STUB: not implemented"
	return false
}

// If there's already a pending value, return whether or not this would
// change that.

// If there's already an applied value, return whether or not this would
// change that.

// If this is marked for removal and does not exist in the cache, no changes
// are necessary.

// If not in the cache, and not marked for removal, it should be added.

// endpointsMapFromEndpointInfo computes an endpointsMap from endpointInfo that
// has been grouped by service port and IP.
func endpointsMapFromEndpointInfo(endpointInfoBySP map[ServicePortName]map[string]Endpoint) EndpointsMap {
	_ = "STUB: not implemented"
	return *new(EndpointsMap)
}

// transform endpointInfoByServicePort into an endpointsMap with sorted IPs.

// formatEndpointsList returns a string list converted from an endpoints list.
func formatEndpointsList(endpoints map[string]Endpoint) []string {
	_ = "STUB: not implemented"
	return nil
}

// endpointSliceCacheKeys returns cache keys used for a given EndpointSlice.
func endpointSliceCacheKeys(endpointSlice *discovery.EndpointSlice) (types.NamespacedName, string, error) {
	_ = "STUB: not implemented"
	return *new(types.NamespacedName), "", nil
}

// byEndpoint helps sort endpoints by endpoint string.
type byEndpoint []Endpoint

func (e byEndpoint) Len() int { _ = "STUB: not implemented"; return 0 }

func (e byEndpoint) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (e byEndpoint) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
