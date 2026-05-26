/*
Copyright 2016 The Kubernetes Authors.

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

Original file https://raw.githubusercontent.com/kubernetes/kubernetes/refs/tags/v1.34.2/pkg/proxy/healthcheck/proxy_health.go

Modifies:

- Remove import "k8s.io/kubernetes/pkg/proxy/metrics" and its usages.
- Replace import "k8s.io/kubernetes/pkg/proxy" with "antrea.io/antrea/v2/third_party/proxy".

*/

package healthcheck

import (
	"context"
	"net/http"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/utils/clock"

	"antrea.io/antrea/v2/third_party/proxy"
)

const (
	// ToBeDeletedTaint is a taint used by the CLuster Autoscaler before marking a node for deletion. Defined in
	// https://github.com/kubernetes/autoscaler/blob/e80ab518340f88f364fe3ef063f8303755125971/cluster-autoscaler/utils/deletetaint/delete.go#L36
	ToBeDeletedTaint = "ToBeDeletedByClusterAutoscaler"
)

// ProxierHealth represents the health of a proxier which operates on a single IP family.
type ProxierHealth struct {
	LastUpdated time.Time `json:"lastUpdated"`
	Healthy     bool      `json:"healthy"`
}

// ProxyHealth represents the health of kube-proxy, embeds health of individual proxiers.
type ProxyHealth struct {
	// LastUpdated is the last updated time of the proxier
	// which was updated most recently.
	// This is kept for backward-compatibility.
	LastUpdated  time.Time `json:"lastUpdated"`
	CurrentTime  time.Time `json:"currentTime"`
	NodeEligible *bool     `json:"nodeEligible,omitempty"`
	// Healthy is true when all the proxiers are healthy,
	// false otherwise.
	Healthy bool `json:"healthy"`
	// status of the health check per IP family
	Status map[v1.IPFamily]ProxierHealth `json:"status,omitempty"`
}

// ProxyHealthServer allows callers to:
//  1. run a http server with /healthz and /livez endpoint handlers.
//  2. update healthz timestamps before and after synchronizing dataplane.
//  3. sync node status, for reporting unhealthy /healthz response
//     if the node is marked for deletion by autoscaler.
//  4. get proxy health by verifying that the delay between QueuedUpdate()
//     calls and Updated() calls exceeded healthTimeout or not.
type ProxyHealthServer struct {
	listener    listener
	httpFactory httpServerFactory
	clock       clock.Clock

	nodeManager *proxy.NodeManager

	addr          string
	healthTimeout time.Duration

	lock                   sync.RWMutex
	lastUpdatedMap         map[v1.IPFamily]time.Time
	oldestPendingQueuedMap map[v1.IPFamily]time.Time
}

// NewProxyHealthServer returns a proxy health http server.
func NewProxyHealthServer(addr string, healthTimeout time.Duration, nodeManager *proxy.NodeManager) *ProxyHealthServer {
	_ = "STUB: not implemented"
	return nil
}

func newProxyHealthServer(listener listener, httpServerFactory httpServerFactory, c clock.Clock, addr string, healthTimeout time.Duration, nodeManager *proxy.NodeManager) *ProxyHealthServer {
	_ = "STUB: not implemented"
	return nil
}

// Updated should be called when the proxier of the given IP family has successfully updated
// the service rules to reflect the current state and should be considered healthy now.
func (hs *ProxyHealthServer) Updated(ipFamily v1.IPFamily) { _ = "STUB: not implemented"; return }

// QueuedUpdate should be called when the proxier receives a Service or Endpoints event
// from API Server containing information that requires updating service rules. It
// indicates that the proxier for the given IP family has received changes but has not
// yet pushed them to its backend. If the proxier does not call Updated within the
// healthTimeout time then it will be considered unhealthy.
func (hs *ProxyHealthServer) QueuedUpdate(ipFamily v1.IPFamily) { _ = "STUB: not implemented"; return }

// Set oldestPendingQueuedMap[ipFamily] only if it's currently unset

// Health returns proxy health status.
func (hs *ProxyHealthServer) Health() ProxyHealth {
	_ = "STUB: not implemented"
	return *new(ProxyHealth)
}

// initialize the health status of each proxier
// with healthy=true and the last updated time
// of the proxier.

// the proxier is healthy while it's starting up
// or the proxier is fully synced.

// there's an unprocessed update queued for this proxier, but it's not late yet.

// mark the status unhealthy.

// NodeEligible returns if node is eligible or not. Eligible is defined
// as being: not tainted by ToBeDeletedTaint and not deleted.
func (hs *ProxyHealthServer) NodeEligible() bool { _ = "STUB: not implemented"; return false }

// Run starts the healthz HTTP server and blocks until it exits.
func (hs *ProxyHealthServer) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type healthzHandler struct {
	hs *ProxyHealthServer
}

func (h healthzHandler) ServeHTTP(resp http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

// updating the node eligibility here (outside of Health() call) as we only want responses
// of /healthz calls (not /livez) to have that.

// In older releases, the returned "lastUpdated" time indicated the last
// time the proxier sync loop ran, even if nothing had changed. To
// preserve compatibility, we use the same semantics: the returned
// lastUpdated value is "recent" if the server is healthy. The kube-proxy
// metrics provide more detailed information.

type livezHandler struct {
	hs *ProxyHealthServer
}

func (h livezHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// In older releases, the returned "lastUpdated" time indicated the last
// time the proxier sync loop ran, even if nothing had changed. To
// preserve compatibility, we use the same semantics: the returned
// lastUpdated value is "recent" if the server is healthy. The kube-proxy
// metrics provide more detailed information.
