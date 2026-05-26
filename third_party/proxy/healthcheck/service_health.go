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

Original file https://raw.githubusercontent.com/kubernetes/kubernetes/refs/tags/v1.34.2/pkg/proxy/healthcheck/service_health.go

Modifies:
- Replace api.EventTypeWarning with string "Warning".
- Remove import proxyutil "k8s.io/kubernetes/pkg/proxy/util" and its usages.
- Modify "newServiceHealthServer" to remove calls to "GetNodeIPs" to get Node IPs. This is not needed as Antrea passes
  actual IP addresses of interfaces.

*/

package healthcheck

import (
	"net/http"
	"sync"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/events"
)

// ServiceHealthServer serves HTTP endpoints for each service name, with results
// based on the endpoints.  If there are 0 endpoints for a service, it returns a
// 503 "Service Unavailable" error (telling LBs not to use this node).  If there
// are 1 or more endpoints, it returns a 200 "OK".
type ServiceHealthServer interface {
	// Make the new set of services be active.  Services that were open before
	// will be closed.  Services that are new will be opened.  Service that
	// existed and are in the new set will be left alone.  The value of the map
	// is the healthcheck-port to listen on.
	SyncServices(newServices map[types.NamespacedName]uint16) error
	// Make the new set of endpoints be active.  Endpoints for services that do
	// not exist will be dropped.  The value of the map is the number of
	// endpoints the service has on this node.
	SyncEndpoints(newEndpoints map[types.NamespacedName]int) error
}

type proxyHealthChecker interface {
	// Health returns the proxy's health state and last updated time.
	Health() ProxyHealth
}

func newServiceHealthServer(nodeName string, recorder events.EventRecorder, listener listener, factory httpServerFactory, nodeIPs []string, healthzServer proxyHealthChecker) ServiceHealthServer {
	_ = "STUB: not implemented"
	return *new(ServiceHealthServer)
}

// NewServiceHealthServer allocates a new service healthcheck server manager
func NewServiceHealthServer(nodeName string, recorder events.EventRecorder, nodePortAddresses []string, healthzServer proxyHealthChecker) ServiceHealthServer {
	_ = "STUB: not implemented"
	return *new(ServiceHealthServer)
}

type server struct {
	nodeName string
	// node addresses where health check port will listen on
	nodeIPs     []string
	recorder    events.EventRecorder // can be nil
	listener    listener
	httpFactory httpServerFactory

	healthzServer proxyHealthChecker

	lock     sync.RWMutex
	services map[types.NamespacedName]*hcInstance
}

func (hcs *server) SyncServices(newServices map[types.NamespacedName]uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove any that are not needed any more.

// errors are loged in closeAll()

// Add any that are needed.

type hcInstance struct {
	nsn  types.NamespacedName
	port uint16

	httpServers []httpServer

	endpoints int // number of local endpoints for a service
}

// listenAll opens health check port on all the addresses provided
func (hcI *hcInstance) listenAndServeAll(hcs *server) error { _ = "STUB: not implemented"; return nil }

// for each of the node addresses start listening and serving

// create http server

// start listener

// must close whatever have been previously opened
// to allow a retry/or port ownership change as needed

// start serving

// Serve() will exit and return ErrServerClosed when the http server is closed.

func (hcI *hcInstance) closeAll() error { _ = "STUB: not implemented"; return nil }

type hcHandler struct {
	name types.NamespacedName
	hcs  *server
}

var _ http.Handler = hcHandler{}

func (h hcHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (hcs *server) SyncEndpoints(newEndpoints map[types.NamespacedName]int) error {
	_ = "STUB: not implemented"
	return nil
}

// FakeServiceHealthServer is a fake ServiceHealthServer for test programs
type FakeServiceHealthServer struct{}

// NewFakeServiceHealthServer allocates a new fake service healthcheck server manager
func NewFakeServiceHealthServer() ServiceHealthServer {
	_ = "STUB: not implemented"
	return *new(ServiceHealthServer)
}

// SyncServices is part of ServiceHealthServer
func (fake FakeServiceHealthServer) SyncServices(_ map[types.NamespacedName]uint16) error {
	_ = "STUB: not implemented"

	// SyncEndpoints is part of ServiceHealthServer
	return nil
}

func (fake FakeServiceHealthServer) SyncEndpoints(_ map[types.NamespacedName]int) error {
	_ = "STUB: not implemented"
	return nil
}
