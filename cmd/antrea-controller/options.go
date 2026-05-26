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
	"github.com/spf13/pflag"

	controllerconfig "antrea.io/antrea/v2/pkg/config/controller"
)

const (
	// Use higher QPS and Burst rather than the default settings (QPS: 5, Burst: 10), otherwise synchronizing resources
	// like ClusterNetworkPolicies and Egresses would be rather slow when they are created in bulk.
	// The following values are recommended by RecommendedDefaultClientConnectionConfiguration.
	// https://github.com/kubernetes/kubernetes/blob/b722d017a34b300a2284b890448e5a605f21d01e/staging/src/k8s.io/component-base/config/v1alpha1/defaults.go#L65-L75
	defaultClientQPS   = 50
	defaultClientBurst = 100

	ipamIPv4MaskLo      = 16
	ipamIPv4MaskHi      = 30
	ipamIPv4MaskDefault = 24
	ipamIPv6MaskLo      = 64
	ipamIPv6MaskHi      = 126
	ipamIPv6MaskDefault = 64
)

type Options struct {
	// The path of configuration file.
	configFile string
	// The configuration object
	config *controllerconfig.ControllerConfig
}

func newOptions() *Options { _ = "STUB: not implemented"; return nil }

// addFlags adds flags to fs and binds them to options.
func (o *Options) addFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// complete completes all the required options.
func (o *Options) complete() error { _ = "STUB: not implemented"; return nil }

// validate validates all the required options.
func (o *Options) validate(args []string) error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateNodeIPAMControllerOptions() error {
	_ = "STUB: not implemented"
	// Validate ClusterCIDRs
	return nil
}

// The subnet mask size cannot be greater than 16 more than the cluster mask size.
// See https://github.com/kubernetes/kubernetes/issues/44918 for more information.

// Validate ServiceCIDR and ServiceCIDRv6. Service CIDRs can be empty when there is no overlap with ClusterCIDR

func (o *Options) loadConfigFromFile() error { _ = "STUB: not implemented"; return nil }

func (o *Options) setDefaults() { _ = "STUB: not implemented"; return }

func ptrBool(value bool) *bool { _ = "STUB: not implemented"; return nil }
