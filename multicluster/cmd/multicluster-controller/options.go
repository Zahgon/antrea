// Copyright 2021 Antrea Authors
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
	"github.com/spf13/pflag"
	ctrl "sigs.k8s.io/controller-runtime"

	mcsv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
)

type Options struct {
	// The path of configuration file.
	configFile     string
	SelfSignedCert bool
	// options store some base controller Manager options (initialized from the provided config).
	options ctrl.Options
	// The Service ClusterIP range used in the member cluster.
	ServiceCIDR string
	// PodCIDRs is the Pod IP address CIDRs of the member cluster.
	PodCIDRs []string
	// The precedence about which IP (private or public one) of Node is preferred to
	// be used as tunnel endpoint. If not specified, private IP will be chosen.
	GatewayIPPrecedence mcsv1alpha1.Precedence
	// The type of IP address (ClusterIP or PodIP) to be used as the Multi-cluster
	// Services' Endpoints.
	EndpointIPType string
	// Enable StretchedNetworkPolicy to exchange labelIdentities info among the whole
	// ClusterSet.
	EnableStretchedNetworkPolicy bool
	// Watch EndpointSlice API for exported Service if EndpointSlice API is available.
	EnableEndpointSlice bool
	// ClusterCalimCRDAvailable indicates if the ClusterClaim CRD is available or not
	// in the cluster.
	ClusterCalimCRDAvailable bool
	// WebhookConfig contains the controllers webhook configuration
	WebhookConfig mcsv1alpha1.ControllerWebhook
}

func newOptions() *Options { _ = "STUB: not implemented"; return nil }

func (o *Options) complete(args []string) error { _ = "STUB: not implemented"; return nil }

// addFlags adds flags to fs and binds them to options.
func (o *Options) addFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func (o *Options) setDefaults() { _ = "STUB: not implemented"; return }

func (o *Options) loadConfig(data []byte, multiclusterConfig *mcsv1alpha1.MultiClusterConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *Options) loadConfigFromFile(multiclusterConfig *mcsv1alpha1.MultiClusterConfig) error {
	_ = "STUB: not implemented"
	return nil
}
