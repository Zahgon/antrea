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

package apiserver

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"

	agentquerier "antrea.io/antrea/v2/pkg/agent/querier"
	systeminstall "antrea.io/antrea/v2/pkg/apis/system/install"
	"antrea.io/antrea/v2/pkg/querier"
)

const CertPairName = "antrea-agent-api"

var (
	scheme = runtime.NewScheme()
	codecs = serializer.NewCodecFactory(scheme)
)

func init() {
	systeminstall.Install(scheme)
	metav1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
}

type agentAPIServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

func (s *agentAPIServer) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *agentAPIServer) GetCertData() []byte { _ = "STUB: not implemented"; return nil }

func installHandlers(aq agentquerier.AgentQuerier, npq querier.AgentNetworkPolicyInfoQuerier, mq querier.AgentMulticastInfoQuerier, seipq querier.ServiceExternalIPStatusQuerier, s *genericapiserver.GenericAPIServer, bgpq querier.AgentBGPPolicyInfoQuerier) {
	_ = "STUB: not implemented"
	return
}

func installAPIGroup(s *genericapiserver.GenericAPIServer, aq agentquerier.AgentQuerier, npq querier.AgentNetworkPolicyInfoQuerier, v4Enabled, v6Enabled bool) error {
	_ = "STUB: not implemented"
	return nil
}

// New creates an APIServer for running in antrea agent.
func New(aq agentquerier.AgentQuerier,
	npq querier.AgentNetworkPolicyInfoQuerier,
	mq querier.AgentMulticastInfoQuerier,
	seipq querier.ServiceExternalIPStatusQuerier,
	bgpq querier.AgentBGPPolicyInfoQuerier,
	secureServing *genericoptions.SecureServingOptionsWithLoopback,
	authentication *genericoptions.DelegatingAuthenticationOptions,
	authorization *genericoptions.DelegatingAuthorizationOptions,
	enableMetrics bool,
	kubeconfig string,
	loopbackClientTokenPath string,
	v4Enabled,
	v6Enabled bool,
) (*agentAPIServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newConfig(aq agentquerier.AgentQuerier,
	npq querier.AgentNetworkPolicyInfoQuerier,
	secureServing *genericoptions.SecureServingOptionsWithLoopback,
	authentication *genericoptions.DelegatingAuthenticationOptions,
	authorization *genericoptions.DelegatingAuthorizationOptions,
	enableMetrics bool,
	kubeconfig string,
	loopbackClientTokenPath string,
) (*genericapiserver.CompletedConfig, error) {
	_ = "STUB: not implemented"
	// kubeconfig file is useful when antrea-agent isn't running as a Pod.
	return nil, nil
}

// Set the PairName but leave certificate directory blank to generate in-memory by default.

// Add readiness probe to check the status of watchers.

// Add liveness probe to check the connection with OFSwitch.
// This helps automatic recovery if some issues cause OFSwitch reconnection to not work properly, e.g. issue #4092.
