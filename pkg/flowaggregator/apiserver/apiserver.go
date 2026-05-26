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

	systeminstall "antrea.io/antrea/v2/pkg/apis/system/install"
	"antrea.io/antrea/v2/pkg/flowaggregator/querier"
)

const (
	Name = "flow-aggregator-api"
	// authenticationTimeout specifies a time limit for requests made by the authorization webhook client
	// The default value (10 seconds) is not long enough as defined in
	// https://pkg.go.dev/k8s.io/apiserver@v0.21.0/pkg/server/options#NewDelegatingAuthenticationOptions
	// A value of zero means no timeout.
	authenticationTimeout = 0
)

var (
	// Scheme defines methods for serializing and deserializing API objects.
	scheme = runtime.NewScheme()
	// Codecs provides methods for retrieving codecs and serializers for specific
	// versions and content types.
	codecs = serializer.NewCodecFactory(scheme)
)

func init() {
	systeminstall.Install(scheme)
	metav1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
}

type flowAggregatorAPIServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

func (s *flowAggregatorAPIServer) Run(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func installHandlers(s *genericapiserver.GenericAPIServer, faq querier.FlowAggregatorQuerier) {
	_ = "STUB: not implemented"
	return
}

// New creates an APIServer for running in flow aggregator.
func New(faq querier.FlowAggregatorQuerier, bindPort int, cipherSuites []uint16, tlsMinVersion uint16) (*flowAggregatorAPIServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newConfig(bindPort int) (*genericapiserver.CompletedConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set the PairName but leave certificate directory blank to generate in-memory by default.
