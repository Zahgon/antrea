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

package ovstracing

import (
	"fmt"
	"net"
	"net/http"
	"regexp"

	corev1 "k8s.io/api/core/v1"

	"antrea.io/antrea/v2/pkg/agent/apiserver/handlers"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/querier"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
)

type tracingPeer struct {
	ovsPort string
	// Name of a Pod or Service
	name string
	// Namespace of Pod or Service.
	namespace string
	ip        net.IP
}

func (p *tracingPeer) getAddressFamily() uint8 { _ = "STUB: not implemented"; return 0 }

type request struct {
	// tracingPeer.ip is invalid for inputPort, as inputPort can only be
	// specified by ovsPort or Pod Namespace/name.
	inputPort   *tracingPeer
	source      *tracingPeer
	destination *tracingPeer
	flow        string
	addrFamily  uint8
}

func getServiceClusterIP(aq querier.AgentQuerier, name, namespace string) (net.IP, *handlers.HandlerError) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func getLocalOVSInterface(aq querier.AgentQuerier, peer *tracingPeer) (*interfacestore.InterfaceConfig, *handlers.HandlerError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Local Pod.

// getPeerAddress looks up a Pod and returns its IP and MAC addresses. It
// first looks up the Pod from the InterfaceStore, and returns the Pod's IP and
// MAC addresses if found. If fails, it then gets the Pod from Kubernetes API,
// and returns the IP address in Pod resource Status if found.
func getPeerAddress(aq querier.AgentQuerier, peer *tracingPeer, addrFamily uint8) (net.IP, *interfacestore.InterfaceConfig, *handlers.HandlerError) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil, nil
}

// Try getting the Pod from K8s API.

// Return IP only assuming it should be a remote Pod.

// Todo: move this function to pkg/agent/util/net.go if it is called by other code
func getPodIPWithAddressFamily(pod *corev1.Pod, addrFamily uint8) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func prepareTracingRequest(aq querier.AgentQuerier, req *request) (*ovsctl.TracingRequest, *handlers.HandlerError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Input port is not specified. Allow "in_port" field in "Flow" to override
// the auto-chosen input port.

// Default source MAC is decided by the input port.

// Input port not specified. Try using the source OVS port.

// The destination might be a Service.

// Must be a local Pod or OVS port. Use interface MAC as the packet
// destination MAC.

// Should be a remote Pod or IP. Use gateway MAC as the destination MAC.

// Source is a remote Pod. Use the default tunnel port as the input port.
// For hybrid TrafficEncapMode, even the remote Node is in the same subnet
// as the source Node, the tunnel port is still used as the input port.

// If the default tunnel port is not found, it might be NoEncap or
// NetworkPolicyOnly mode. Use gateway port as the input port then.

// Use gateway port as the input port when the source is an IP address
// (assuming it is an external IP or a Node IP).

// Use tunnel traffic virtual MAC for both source and destination MAC
// addresses of the trace packet input from the tunnel port.

// Use gateway port as the input port if it could not be figured out from the
// source.

// parseTracingPeer parses Pod/Service name and Namespace or OVS port name or
// IP address from the string. nil is returned if the string is not of a
// valid Pod/Service reference ("Namespace/name") or OVS port name format, and
// not an IP address.
func parseTracingPeer(str string) *tracingPeer { _ = "STUB: not implemented"; return nil }

// Namespace and name must not be empty.

// Probably an OVS port name.

const (
	flowRegexElementPattern = `[a-zA-Z0-9\.\-_]+(=[a-zA-Z0-9\.\-_]+)?`
)

var (
	flowRegexPattern = fmt.Sprintf(`^(%s,\s*)*%s$`, flowRegexElementPattern, flowRegexElementPattern)
	flowRegex        = regexp.MustCompile(flowRegexPattern)
)

func validateRequest(r *http.Request) (*request, *handlers.HandlerError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sanitize user input since it is used to invoke exec.Command

// Input port cannot be specified with an IP.

// HandleFunc returns the function which can handle API requests to "/ovstracing".
func HandleFunc(aq querier.AgentQuerier) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// ovs-appctl has been executed but returned an error (e.g. the provided
// "flow" expression is incorrect). Return the error output to the client in
// this case.
