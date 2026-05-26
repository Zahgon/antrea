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

package apiserver

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"

	cpinstall "antrea.io/antrea/v2/pkg/apis/controlplane/install"
	statsinstall "antrea.io/antrea/v2/pkg/apis/stats/install"
	systeminstall "antrea.io/antrea/v2/pkg/apis/system/install"
	"antrea.io/antrea/v2/pkg/apiserver/certificate"
	"antrea.io/antrea/v2/pkg/apiserver/storage"
	crdv1a2informers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha2"
	"antrea.io/antrea/v2/pkg/controller/egress"
	"antrea.io/antrea/v2/pkg/controller/externalippool"
	"antrea.io/antrea/v2/pkg/controller/ipam"
	controllernetworkpolicy "antrea.io/antrea/v2/pkg/controller/networkpolicy"
	"antrea.io/antrea/v2/pkg/controller/querier"
	"antrea.io/antrea/v2/pkg/controller/stats"
	controllerbundlecollection "antrea.io/antrea/v2/pkg/controller/supportbundlecollection"
	"antrea.io/antrea/v2/pkg/controller/traceflow"
)

var (
	// Scheme defines methods for serializing and deserializing API objects.
	Scheme = runtime.NewScheme()
	// Codecs provides methods for retrieving codecs and serializers for specific
	// versions and content types.
	Codecs = serializer.NewCodecFactory(Scheme)
	// ParameterCodec defines methods for serializing and deserializing url values
	// to versioned API objects and back.
	parameterCodec = runtime.NewParameterCodec(Scheme)

	// antreaServedLabelSelector selects resources served by antrea-controller.
	antreaServedLabelSelector = &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app":       "antrea",
			"served-by": "antrea-controller",
		},
	}
)

func init() {
	cpinstall.Install(Scheme)
	systeminstall.Install(Scheme)
	statsinstall.Install(Scheme)

	// We need to add the options to empty v1, see sample-apiserver/pkg/apiserver/apiserver.go.
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
}

// ExtraConfig holds custom apiserver config.
type ExtraConfig struct {
	k8sClient                     kubernetes.Interface
	addressGroupStore             storage.Interface
	appliedToGroupStore           storage.Interface
	networkPolicyStore            storage.Interface
	egressGroupStore              storage.Interface
	bundleCollectionStore         storage.Interface
	podInformer                   coreinformers.PodInformer
	nodeInformer                  coreinformers.NodeInformer
	eeInformer                    crdv1a2informers.ExternalEntityInformer
	controllerQuerier             querier.ControllerQuerier
	endpointQuerier               controllernetworkpolicy.EndpointQuerier
	networkPolicyController       *controllernetworkpolicy.NetworkPolicyController
	egressController              *egress.EgressController
	externalIPPoolController      *externalippool.ExternalIPPoolController
	ipamController                *ipam.AntreaIPAMController
	caCertController              *certificate.CACertController
	statsAggregator               *stats.Aggregator
	networkPolicyStatusController *controllernetworkpolicy.StatusController
	bundleCollectionController    *controllerbundlecollection.Controller
	traceflowController           *traceflow.Controller
}

// Config defines the config for Antrea apiserver.
type Config struct {
	genericConfig *genericapiserver.Config
	extraConfig   ExtraConfig
}

// APIServer contains state for a Kubernetes cluster apiserver.
type APIServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
	caCertController *certificate.CACertController
}

func (s *APIServer) Run(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Make sure CACertController runs once to publish the CA cert before starting APIServer.
	return nil
}

type completedConfig struct {
	genericConfig genericapiserver.CompletedConfig
	extraConfig   *ExtraConfig
}

func NewConfig(
	genericConfig *genericapiserver.Config,
	k8sClient kubernetes.Interface,
	addressGroupStore, appliedToGroupStore, networkPolicyStore, egressGroupStore, supportBundleCollectionStore storage.Interface,
	podInformer coreinformers.PodInformer,
	nodeInformer coreinformers.NodeInformer,
	eeInformer crdv1a2informers.ExternalEntityInformer,
	caCertController *certificate.CACertController,
	statsAggregator *stats.Aggregator,
	controllerQuerier querier.ControllerQuerier,
	networkPolicyStatusController *controllernetworkpolicy.StatusController,
	endpointQuerier controllernetworkpolicy.EndpointQuerier,
	npController *controllernetworkpolicy.NetworkPolicyController,
	egressController *egress.EgressController,
	externalIPPoolController *externalippool.ExternalIPPoolController,
	ipamController *ipam.AntreaIPAMController,
	bundleCollectionController *controllerbundlecollection.Controller,
	traceflowController *traceflow.Controller) *Config {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) Complete(informers informers.SharedInformerFactory) completedConfig {
	_ = "STUB: not implemented"
	return *new(completedConfig)
}

func installAPIGroup(s *APIServer, c completedConfig) error { _ = "STUB: not implemented"; return nil }

func (c completedConfig) New() (*APIServer, error) { _ = "STUB: not implemented"; return nil, nil }

// CleanupDeprecatedAPIServices deletes the registered APIService resources for
// the deprecated Antrea API groups.
func CleanupDeprecatedAPIServices(aggregatorClient clientset.Interface) error {
	_ = "STUB: not implemented"
	// The APIService of a deprecated API group should be added to the slice.
	// After Antrea upgrades from an old version to a new version that
	// deprecates a registered APIService, the APIService should be deleted,
	// otherwise K8s will fail to delete an existing Namespace.
	// Also check: https://github.com/antrea-io/antrea/issues/494
	return nil
}

func installHandlers(c *ExtraConfig, s *genericapiserver.GenericAPIServer) {
	_ = "STUB: not implemented"
	return
}

// Webhook to mutate Namespace labels and add its metadata.name as a label

// Get new NetworkPolicyMutator

// Install handlers for NetworkPolicy related mutation

// Get new NetworkPolicyValidator

// Install handlers for NetworkPolicy related validation

// Install a post start hook to initialize Tiers on start-up

// context gets cancelled when the server stops.

func DefaultCAConfig() *certificate.CAConfig { _ = "STUB: not implemented"; return nil }

// Rotate the certificate 90 days in advance.
