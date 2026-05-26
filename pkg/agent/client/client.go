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

package client

import (
	"context"
	"sync"

	"k8s.io/apiserver/pkg/server/dynamiccertificates"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/component-base/config"

	"antrea.io/antrea/v2/pkg/client/clientset/versioned"
)

// AntreaClientProvider provides a method to get Antrea client.
type AntreaClientProvider interface {
	GetAntreaClient() (versioned.Interface, error)
}

// antreaClientProvider provides an AntreaClientProvider that can dynamically react to CA bundle
// ConfigMap changes, as well as directly resolve the Antrea Service Endpoint when running inside a K8s cluster.
// The consumers of antreaClientProvider are supposed to always call GetAntreaClient() to get a client and not cache it.
type antreaClientProvider struct {
	config config.ClientConnectionConfiguration
	// mutex protects client.
	mutex sync.RWMutex
	// client is the Antrea client that will be returned. It will be updated when caBundle is updated.
	client versioned.Interface
	// caContentProvider provides the very latest content of the ca bundle.
	caContentProvider *dynamiccertificates.ConfigMapCAController
	// endpointResolver provides a known Endpoint for the Antrea Service. There is usually a
	// single Endpoint at any given time, given that the Antrea Controller runs as a
	// single-replica Deployment. By resolving the Endpoint manually and accessing it directly,
	// instead of depending on the ClusterIP functionality provided by the K8s proxy, we get
	// more flexibility when initializing the Antrea Agent. For example, we can retrieve
	// NetworkPolicies from the Controller even if the proxy is not (yet) available.
	// endpointResolver is only used when no kubeconfig is provided (otherwise we honor the
	// provided config).
	endpointResolver *EndpointResolver
}

// antreaClientProvider must implement the dynamiccertificates.Listener interface to be notified of
// CA bundle updates.
var _ dynamiccertificates.Listener = &antreaClientProvider{}

// antreaClientProvider must implement the Listener interface to be notified of an Endpoint change
// for the Antrea Service.
var _ Listener = &antreaClientProvider{}

func NewAntreaClientProvider(config config.ClientConnectionConfiguration, kubeClient kubernetes.Interface) (*antreaClientProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunOnce runs the task a single time synchronously, ensuring client is initialized if kubeconfig is specified.
func (p *antreaClientProvider) RunOnce() error { _ = "STUB: not implemented"; return nil }

// Run starts the caContentProvider, which watches the ConfigMap and notifies changes
// by calling Enqueue.
func (p *antreaClientProvider) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

// Enqueue implements dynamiccertificates.Listener. It will be called by caContentProvider
// when caBundle is updated.
func (p *antreaClientProvider) Enqueue() { _ = "STUB: not implemented"; return }

// GetAntreaClient implements AntreaClientProvider.
func (p *antreaClientProvider) GetAntreaClient() (versioned.Interface, error) {
	_ = "STUB: not implemented"
	return *new(versioned.Interface), nil
}

func (p *antreaClientProvider) updateAntreaClient() error { _ = "STUB: not implemented"; return nil }

// ContentType will be used to define the Accept header if AcceptContentTypes is not set.

// inClusterConfig returns a config object which uses the service account Kubernetes gives to
// Pods. It's intended for clients that expect to be running inside a Pod running on Kubernetes. It
// will return error if called from a process not running in a Kubernetes environment.
func inClusterConfig(caBundle []byte, endpoint string) (*rest.Config, error) {
	_ = "STUB: not implemented"
	// #nosec G101: false positive triggered by variable name which includes "token"
	return nil, nil
}
