// Copyright 2021 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	networkingv1 "k8s.io/api/networking/v1"
	apiextensionclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientset "k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	aggregatorclientset "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	k8smcsv1alpha1 "sigs.k8s.io/mcs-api/pkg/apis/v1alpha1"

	mcv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
	mcv1alpha2 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha2"
	antreacrdv1alpha1 "antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	antreacrdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/apiserver/certificate"
	// +kubebuilder:scaffold:imports
)

var (
	// The unit test code will change the function to set up a mock manager.
	setupManagerAndCertControllerFunc = setupManagerAndCertController
)

const (
	selfSignedCertDir = "/var/run/antrea/multicluster-controller-self-signed"
	certDir           = "/var/run/antrea/multicluster-controller-tls"
	serviceName       = "antrea-mc-webhook-service"
	configMapName     = "antrea-mc-ca"
	leaderRole        = "leader"
	memberRole        = "member"
)

var (
	// mcDefaultServedLabels contains the labels added on the Webhooks which are needed by both leader and member controllers.
	mcDefaultServedLabels = map[string]string{
		"app":       "antrea",
		"served-by": "antrea-mc-controller",
	}
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(k8smcsv1alpha1.AddToScheme(scheme))
	utilruntime.Must(mcv1alpha1.AddToScheme(scheme))
	utilruntime.Must(mcv1alpha2.AddToScheme(scheme))
	utilruntime.Must(antreacrdv1alpha1.AddToScheme(scheme))
	utilruntime.Must(antreacrdv1beta1.AddToScheme(scheme))
	utilruntime.Must(networkingv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

// createClients creates kube clients from the given config.
func createClients(kubeConfig *rest.Config) (
	clientset.Interface, aggregatorclientset.Interface, apiextensionclientset.Interface, error) {
	_ = "STUB: not implemented"
	return *new(clientset.Interface), *new(aggregatorclientset.Interface), *new(apiextensionclientset.Interface), nil
}

// Create client for crd manipulations

func getCaConfig(isLeader bool, controllerNs string) *certificate.CAConfig {
	_ = "STUB: not implemented"
	return nil
}

// the key pair name has to be "tls" https://github.com/kubernetes-sigs/controller-runtime/blob/master/pkg/manager/manager.go#L221

// Rotate the certificate 90 days in advance.

func getWebhookLabel(isLeader bool, controllerNs string) *metav1.LabelSelector {
	_ = "STUB: not implemented"
	return nil
}

// It is allowed that multiple leader controllers running in different Namespace in the same cluster.
// "served-in: $controllerNS" is useful to select the Webhooks managed by the current mc-controller.

func setupManagerAndCertController(isLeader bool, o *Options) (manager.Manager, error) {
	_ = "STUB: not implemented"
	return *new(manager.Manager), nil
}

// build up cert controller to manage certificate for MC Controller

// For the leader, restrict the cache to the controller's Namespace.

// For a member, restict the cache to the controller's Namespace for the following objects.

// EndpointSlice is enabled in AntreaProxy by default since v1.11, so Antrea MC
// will use EndpointSlice API by default to keep consistent with AntreaProxy.

// ClusterClaim CRD is removed since v1.13. Check the existence of
// ClusterClaim API before using ClusterClaim API.

//+kubebuilder:scaffold:builder

func clusterClaimCRDAvailable(k8sClient clientset.Interface) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// The group version doesn't exist.
