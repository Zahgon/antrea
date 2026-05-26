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

package main

import (
	"net"
	"time"

	apiextensionclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	coreinformers "k8s.io/client-go/informers/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	aggregatorclientset "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"

	"antrea.io/antrea/v2/pkg/apiserver"
	"antrea.io/antrea/v2/pkg/apiserver/storage"
	crdv1a2informers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha2"
	"antrea.io/antrea/v2/pkg/controller/egress"
	"antrea.io/antrea/v2/pkg/controller/externalippool"
	antreaipam "antrea.io/antrea/v2/pkg/controller/ipam"
	"antrea.io/antrea/v2/pkg/controller/networkpolicy"
	"antrea.io/antrea/v2/pkg/controller/querier"
	"antrea.io/antrea/v2/pkg/controller/stats"
	"antrea.io/antrea/v2/pkg/controller/supportbundlecollection"
	"antrea.io/antrea/v2/pkg/controller/traceflow"
)

const (
	// informerDefaultResync is the default resync period if a handler doesn't specify one.
	// Use the same default value as kube-controller-manager:
	// https://github.com/kubernetes/kubernetes/blob/release-1.17/pkg/controller/apis/config/v1alpha1/defaults.go#L120
	informerDefaultResync = 12 * time.Hour

	// serverMinWatchTimeout determines the timeout allocated to watches from Antrea
	// clients. Each watch will be allocated a random timeout between this value and twice this
	// value, to help randomly distribute reconnections over time.
	// This parameter corresponds to the MinRequestTimeout server config parameter in
	// https://godoc.org/k8s.io/apiserver/pkg/server#Config.
	// When the Antrea client re-creates a watch, all relevant NetworkPolicy objects need to be
	// sent again by the controller. It may be a good idea to use a value which is larger than
	// the kube-apiserver default (1800s). The K8s documentation states that clients should be
	// able to handle watch timeouts gracefully but recommends using a large value in
	// production.
	serverMinWatchTimeout = 2 * time.Hour
)

var allowedPaths = []string{
	"/healthz",
	"/livez",
	"/readyz",
	"/mutate/acnp",
	"/mutate/annp",
	"/mutate/anp",
	"/mutate/namespace",
	"/validate/tier",
	"/validate/acnp",
	"/validate/annp",
	"/validate/anp",
	"/validate/banp",
	"/validate/clustergroup",
	"/validate/externalippool",
	"/validate/egress",
	"/validate/group",
	"/validate/ippool",
	"/validate/supportbundlecollection",
	"/validate/traceflow",
	"/convert/clustergroup",
	"/convert/ippool",
}

// run starts Antrea Controller with the given options and waits for termination signal.
func run(o *Options) error { _ = "STUB: not implemented"; return nil }

// Create K8s Clientset, Aggregator Clientset, CRD Clientset and SharedInformerFactory for the given config.
// Aggregator Clientset is used to update the CABundle of the APIServices backed by antrea-controller so that
// the aggregator can verify its serving certificate.

// Add IP-Pod index. Each Pod has no more than 2 IPs, the extra overhead is constant and acceptable.
// @tnqn evaluated the performance without/with IP index is 3us vs 4us per pod, i.e. 300ms vs 400ms for 100k Pods.

// Create Antrea object storage.

// statsAggregator takes stats summaries from antrea-agents, aggregates them, and serves the Stats APIs with the
// aggregated data. For now it's only used for NetworkPolicy stats.

// Set up signal capture: the first SIGTERM / SIGINT signal is handled gracefully and will
// cause the stopCh channel to be closed; if another signal is received before the program
// exits, we will force exit.

// Generate a context for functions which require one (instead of stopCh).
// We cancel the context when the function returns, which in the normal case will be when
// stopCh is closed.

// TODO(yang): When the AdminNetworkPolicy graduates to Beta, we need a better mechanism in Antrea controller
//  so that 1. the policyInformerFactory is only started if the AdminNetworkPolicy CRD types are installed in
//  the cluster 2. It retries periodically so that when the CRDs are installed, the policyInformerFactory starts
//  watching and handles policy events.

// It starts dispatching group updates to consumers, should start individually.
// If it's not running, adding Pods/Entities to groupEntityIndex may be blocked because of full channel.

func getNodeCIDRMaskSizes(clusterCIDRs []*net.IPNet, maskSizeIPv4, maskSizeIPv6 int) []int {
	_ = "STUB: not implemented"
	return nil
}

func startNodeIPAM(client clientset.Interface,
	nodeInformer coreinformers.NodeInformer,
	clusterCIDRs []*net.IPNet,
	serviceCIDR *net.IPNet,
	serviceCIDRv6 *net.IPNet,
	nodeCIDRMaskSizeIPv4 int,
	nodeCIDRMaskSizeIPv6 int,
	stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func createAPIServerConfig(kubeconfig string,
	clientCAFile string,
	client clientset.Interface,
	aggregatorClient aggregatorclientset.Interface,
	apiExtensionClient apiextensionclientset.Interface,
	selfSignedCert bool,
	bindPort int,
	addressGroupStore storage.Interface,
	appliedToGroupStore storage.Interface,
	networkPolicyStore storage.Interface,
	egressGroupStore storage.Interface,
	supportBundleCollectionStore storage.Interface,
	podInformer coreinformers.PodInformer,
	nodeInformer coreinformers.NodeInformer,
	eeInformer crdv1a2informers.ExternalEntityInformer,
	controllerQuerier querier.ControllerQuerier,
	endpointQuerier networkpolicy.EndpointQuerier,
	npController *networkpolicy.NetworkPolicyController,
	networkPolicyStatusController *networkpolicy.StatusController,
	egressController *egress.EgressController,
	externalIPPoolController *externalippool.ExternalIPPoolController,
	antreaIPAMController *antreaipam.AntreaIPAMController,
	statsAggregator *stats.Aggregator,
	bundleCollectionStore *supportbundlecollection.Controller,
	traceflowController *traceflow.Controller,
	enableMetrics bool,
	cipherSuites []uint16,
	tlsMinVersion uint16) (*apiserver.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// kubeconfig file is useful when antrea-controller is not running as a pod, like during development.
