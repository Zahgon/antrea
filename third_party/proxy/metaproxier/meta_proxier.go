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

Original file https://raw.githubusercontent.com/kubernetes/kubernetes/refs/tags/v1.34.2/pkg/proxy/metaproxier/meta_proxier.go

Modifies:

- Replace import "k8s.io/kubernetes/pkg/proxy" with "antrea.io/antrea/v2/third_party/proxy".

Adds:

- Add functions:
  - `func (proxier *metaProxier) SyncedOnce() bool`
  - `func (proxier *metaProxier) Run(_ <-chan struct{})`

*/

package metaproxier

import (
	v1 "k8s.io/api/core/v1"
	discovery "k8s.io/api/discovery/v1"

	"antrea.io/antrea/v2/third_party/proxy"
)

type metaProxier struct {
	// actual, wrapped
	ipv4Proxier proxy.Provider
	// actual, wrapped
	ipv6Proxier proxy.Provider
}

// NewMetaProxier returns a dual-stack "meta-proxier". Proxier API
// calls will be dispatched to the ProxyProvider instances depending
// on address family.
func NewMetaProxier(ipv4Proxier, ipv6Proxier proxy.Provider) proxy.Provider {
	_ = "STUB: not implemented"
	return *new(proxy.Provider)
}

// Sync immediately synchronizes the ProxyProvider's current state to
// proxy rules.
func (proxier *metaProxier) Sync() { _ = "STUB: not implemented"; return }

// SyncLoop runs periodic work.  This is expected to run as a
// goroutine or as the main loop of the app.  It does not return.
func (proxier *metaProxier) SyncLoop() { _ = "STUB: not implemented"; return }

// Use go-routine here!
// never returns

// OnServiceAdd is called whenever creation of new service object is observed.
func (proxier *metaProxier) OnServiceAdd(service *v1.Service) { _ = "STUB: not implemented"; return }

// OnServiceUpdate is called whenever modification of an existing
// service object is observed.
func (proxier *metaProxier) OnServiceUpdate(oldService, service *v1.Service) {
	_ = "STUB: not implemented"
	return
}

// OnServiceDelete is called whenever deletion of an existing service
// object is observed.
func (proxier *metaProxier) OnServiceDelete(service *v1.Service) { _ = "STUB: not implemented"; return }

// OnServiceSynced is called once all the initial event handlers were
// called and the state is fully propagated to local cache.
func (proxier *metaProxier) OnServiceSynced() { _ = "STUB: not implemented"; return }

// OnEndpointSliceAdd is called whenever creation of a new endpoint slice object
// is observed.
func (proxier *metaProxier) OnEndpointSliceAdd(endpointSlice *discovery.EndpointSlice) {
	_ = "STUB: not implemented"
	return
}

// OnEndpointSliceUpdate is called whenever modification of an existing endpoint
// slice object is observed.
func (proxier *metaProxier) OnEndpointSliceUpdate(oldEndpointSlice, newEndpointSlice *discovery.EndpointSlice) {
	_ = "STUB: not implemented"
	return
}

// OnEndpointSliceDelete is called whenever deletion of an existing endpoint slice
// object is observed.
func (proxier *metaProxier) OnEndpointSliceDelete(endpointSlice *discovery.EndpointSlice) {
	_ = "STUB: not implemented"
	return
}

// OnEndpointSlicesSynced is called once all the initial event handlers were
// called and the state is fully propagated to local cache.
func (proxier *metaProxier) OnEndpointSlicesSynced() { _ = "STUB: not implemented"; return }

// OnTopologyChange is called whenever change in proxy relevant topology labels is observed.
func (proxier *metaProxier) OnTopologyChange(topologyLabels map[string]string) {
	_ = "STUB: not implemented"
	return
}

// OnServiceCIDRsChanged is called whenever a change is observed
// in any of the ServiceCIDRs, and provides complete list of service cidrs.
func (proxier *metaProxier) OnServiceCIDRsChanged(cidrs []string) {
	_ = "STUB: not implemented"
	return
}

// SyncedOnce returns true if the proxier has synced rules at least once.
func (proxier *metaProxier) SyncedOnce() bool { _ = "STUB: not implemented"; return false }

func (proxier *metaProxier) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }
