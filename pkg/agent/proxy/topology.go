// Copyright 2022 Antrea Authors
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
	k8sproxy "antrea.io/antrea/v2/third_party/proxy"
)

func (p *proxier) categorizeEndpoints(endpoints map[string]k8sproxy.Endpoint, svcInfo k8sproxy.ServicePort, nodeName string, nodeLabels map[string]string) ([]k8sproxy.Endpoint, []k8sproxy.Endpoint, []k8sproxy.Endpoint) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// If cluster Endpoints is to be used for the Service, generate a list of cluster Endpoints.

// If there is no cluster Endpoint, fallback to any terminating Endpoints that are serving. When falling back to
// terminating Endpoints, and topology aware routing is NOT considered since this is the best effort attempt to
// avoid dropping connections.

// If local Endpoints is not to be used, clusterEndpoints is just allReachableEndpoints, then only return clusterEndpoints
// and allReachableEndpoints.

// If there are any local Endpoints, use local ready Endpoints.

// If there is no local Endpoint, fallback to terminating local Endpoints that are serving. When falling back to
// terminating Endpoints, and topology aware routing is NOT considered since this is the best effort attempt to
// avoid dropping connections.

// If cluster Endpoints is not to be used, localEndpoints is just allReachableEndpoints, then only return localEndpoints
// and allReachableEndpoints.

// !useServingTerminatingEndpoints means that localEndpoints contains only Ready Endpoints. topologyMode == "" means
// that clusterEndpoints contains *every* Ready Endpoint. So clusterEndpoints must be a superset of localEndpoints.

// clusterEndpoints may contain remote Endpoints that aren't in localEndpoints, while localEndpoints may contain
// terminating or topologically-unavailable local endpoints that aren't in clusterEndpoints. So we have to merge
// the two lists.

// topologyModeFromHints returns a topology mode ("", "PreferSameZone", or "PreferSameNode") based on the Endpoint hints:
//   - If the PreferSameTrafficDistribution feature gate is enabled, and every ready endpoint has a node hint, and at
//     least one endpoint is hinted for this node, then it returns "PreferSameNode".
//   - Otherwise, if every ready endpoint has a zone hint, and at least one endpoint is hinted for this node's zone,
//     then it returns "PreferSameZone".
//   - Otherwise it returns "" (meaning, no topology / default traffic distribution).
func (p *proxier) topologyModeFromHints(svcInfo k8sproxy.ServicePort, endpoints map[string]k8sproxy.Endpoint, nodeName, zone string) string {
	_ = "STUB: not implemented"
	return ""
}

// availableForTopology checks if this endpoint is available for use on this node when using the given topologyMode.
// (Note that there's no fallback here; the fallback happens when deciding which mode to use, not when applying that
// decision.)
func availableForTopology(endpoint k8sproxy.Endpoint, topologyMode, nodeName, zone string) bool {
	_ = "STUB: not implemented"
	return false
}

// filterEndpoints filters endpoints according to predicate
func filterEndpoints(endpoints map[string]k8sproxy.Endpoint, predicate func(k8sproxy.Endpoint) bool) []k8sproxy.Endpoint {
	_ = "STUB: not implemented"
	return nil
}
