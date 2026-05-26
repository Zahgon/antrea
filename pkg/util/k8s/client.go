// Copyright 2019 Antrea Authors
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

package k8s

import (
	apiextensionclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	componentbaseconfig "k8s.io/component-base/config"
	aggregatorclientset "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
	policyclient "sigs.k8s.io/network-policy-api/pkg/client/clientset/versioned"

	mcclientset "antrea.io/antrea/v2/multicluster/pkg/client/clientset/versioned"
	crdclientset "antrea.io/antrea/v2/pkg/client/clientset/versioned"
)

const (
	kubeServiceHostEnvKey = "KUBERNETES_SERVICE_HOST"
	kubeServicePortEnvKey = "KUBERNETES_SERVICE_PORT"
)

// CreateClients creates kube clients from the given config.
func CreateClients(config componentbaseconfig.ClientConnectionConfiguration, kubeAPIServerOverride string) (
	clientset.Interface, aggregatorclientset.Interface, crdclientset.Interface, apiextensionclientset.Interface, mcclientset.Interface, policyclient.Interface, error) {
	_ = "STUB: not implemented"
	return *new(clientset.Interface), *new(aggregatorclientset.Interface), *new(crdclientset.Interface), *new(apiextensionclientset.Interface), *new(mcclientset.Interface), *new(policyclient.Interface), nil
}

// Create client for CRD operations.

// Create client for CRD manipulations.

// Create client for multicluster CRD operations.

func CreateRestConfig(config componentbaseconfig.ClientConnectionConfiguration, kubeAPIServerOverride string) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OverrideKubeAPIServer overrides the env vars related to the kubernetes service used by InClusterConfig.
// It's required because some K8s libraries like DelegatingAuthenticationOptions and DelegatingAuthorizationOptions
// read the information from env vars and don't support overriding via parameters.
func OverrideKubeAPIServer(kubeAPIServerOverride string) { _ = "STUB: not implemented"; return }

func ParseKubeAPIServerOverride(kubeAPIServerOverride string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// if SplitHostPort returns an error, the entire hostport is considered as host

func EndpointSliceAPIAvailable(k8sClient clientset.Interface) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// The group version doesn't exist.
