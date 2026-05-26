// Copyright 2024 Antrea Authors
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
	"net/url"
	"sync/atomic"
	"time"

	"k8s.io/apiserver/pkg/util/proxy"
	"k8s.io/client-go/kubernetes"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	// informerDefaultResync is the default resync period if a handler doesn't specify one.
	// Use the same default value as kube-controller-manager:
	// https://github.com/kubernetes/kubernetes/blob/release-1.17/pkg/controller/apis/config/v1alpha1/defaults.go#L120
	informerDefaultResync = 12 * time.Hour

	minRetryDelay = 100 * time.Millisecond
	maxRetryDelay = 30 * time.Second
)

// Listener defines the interface which needs to be implemented by clients which want to subscribe
// to Endpoint updates.
type Listener interface {
	Enqueue()
}

// EndpointResolver is in charge of resolving a specific Service Endpoint, which can then be
// accessed directly instead of depending on the ClusterIP functionality provided by K8s proxies
// (whether it's kube-proxy or AntreaProxy). A new Endpoint is resolved every time the Service's
// Spec or the Endpoints' Subsets are updated, and registered listeners are notified. While this
// EndpointResolver is somewhat generic, at the moment it is only meant to be used for the Antrea
// Service.
type EndpointResolver struct {
	// name is the name of the controller in charge of Endpoint resolution.
	name        string
	namespace   string
	serviceName string
	servicePort int32
	// serviceInformer and endpointSliceInformer are stored here so they can be started in the Run() method.
	serviceInformer       cache.SharedIndexInformer
	endpointSliceInformer cache.SharedIndexInformer
	// serviceLister is used to retrieve the Service when selecting an Endpoint.
	serviceLister       corev1listers.ServiceLister
	serviceListerSynced cache.InformerSynced
	// endpointSliceGetter is used to retrieve the EndpointSlices for the Service during Endpoint selection.
	endpointSliceGetter       proxy.EndpointSliceGetter
	endpointSliceListerSynced cache.InformerSynced
	queue                     workqueue.TypedRateLimitingInterface[string]
	// listeners need to implement the Listerner interface and will get notified when the
	// current Endpoint URL changes.
	listeners   []Listener
	endpointURL atomic.Pointer[url.URL]
}

func NewEndpointResolver(kubeClient kubernetes.Interface, namespace, serviceName string, servicePort int32) *EndpointResolver {
	_ = "STUB: not implemented"
	return nil
}

// We only need a specific Service and corresponding EndpointSlices, so we create
// filtered informers directly without factories for better efficiency.

// Create an EndpointSliceGetter from the lister for use with proxy.ResolveEndpoint

// This should not happen: both objects should be Services in the
// update event handler.

// Ignore changes to metadata or status.

// This should not happen: both objects should be EndpointSlices in the
// update event handler.

// Ignore changes to metadata, only look at changes to endpoints or ports.

func (r *EndpointResolver) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

// We only start one worker for this controller.

func (r *EndpointResolver) runWorker() { _ = "STUB: not implemented"; return }

func (r *EndpointResolver) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

func (r *EndpointResolver) resolveEndpoint() error { _ = "STUB: not implemented"; return nil }

// Typically we will get one of these 2 errors (unavailable or not found).
// In this case, it makes sense to reset the Endpoint URL to nil and notify listeners.
// There is also no need to retry, as we won't find a suitable Endpoint until the Service or
// the Endpoints resource is updated in a way that will cause this function to be called again.

// Unknown error: we err on the side of caution.
// Do not reset the URL or notify listeners, and trigger a retry.

func (r *EndpointResolver) updateEndpointIfNeeded(endpointURL *url.URL) {
	_ = "STUB: not implemented"
	// The separate Load and Store calls are safe because there is a single writer for r.endpointURL.
	return
}

func (r *EndpointResolver) AddListener(listener Listener) { _ = "STUB: not implemented"; return }

func (r *EndpointResolver) CurrentEndpointURL() *url.URL { _ = "STUB: not implemented"; return nil }
