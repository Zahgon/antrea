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
	"time"

	"github.com/spf13/pflag"

	"antrea.io/antrea/v2/pkg/agent/config"
	agentconfig "antrea.io/antrea/v2/pkg/config/agent"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
)

const (
	defaultOVSBridge               = "br-int"
	defaultHostGateway             = "antrea-gw0"
	defaultHostProcPathPrefix      = "/host"
	defaultServiceCIDR             = "10.96.0.0/12"
	defaultTunnelType              = ovsconfig.GeneveTunnel
	defaultFlowCollectorAddress    = "flow-aggregator/flow-aggregator:4739:tls"
	defaultFlowCollectorTransport  = "tls"
	defaultFlowCollectorPort       = "4739"
	defaultFlowPollInterval        = "5s"
	defaultActiveFlowExportTimeout = "5s"
	defaultIdleFlowExportTimeout   = "15s"
	defaultIGMPQueryInterval       = 125 * time.Second
	defaultStaleConnectionTimeout  = 5 * time.Minute
	defaultNodeType                = config.K8sNode
	defaultMaxEgressIPsPerNode     = 255
	defaultAuditLogsMaxSize        = 500
	defaultAuditLogsMaxBackups     = 3
	defaultAuditLogsMaxAge         = 28
	defaultAuditLogsCompressed     = true
	defaultPacketInRate            = 5000
)

var defaultIGMPQueryVersions = []int{1, 2, 3}

type Options struct {
	// The path of configuration file.
	configFile string
	// The configuration object
	config *agentconfig.AgentConfig
	// tlsCipherSuites is a slice of TLSCipherSuites mapped to input provided by user.
	tlsCipherSuites []string
	// IPFIX flow collector address
	flowCollectorAddr string
	// IPFIX flow collector protocol
	flowCollectorProto string
	// Flow exporter poll interval
	pollInterval time.Duration
	// Active flow timeout to export records of active flows
	activeFlowTimeout time.Duration
	// Idle flow timeout to export records of inactive flows
	idleFlowTimeout time.Duration
	// Stale connection timeout to delete connections if they are not exported.
	staleConnectionTimeout time.Duration
	igmpQueryInterval      time.Duration
	igmpQueryVersions      []uint8
	nplStartPort           int
	nplEndPort             int
	dnsServerOverride      string
	nodeType               config.NodeType

	// enableEgress represents whether Egress should run or not, calculated from its feature gate configuration and
	// whether the traffic mode supports it.
	enableEgress bool
	// enableAntreaProxy indicates whether AntreaProxy should be enabled, based on feature gate AntreaProxy and options
	// AntreaProxy.Enable. This is used to maintain compatibility with the AntreaProxy feature gate, which was promoted
	// to GA in v1.14.
	enableAntreaProxy bool
	// enableNodePortLocal indicates whether NodePortLocal should be enabled or not, based on feature gate NodePortLocal
	// and options NodePortLocal.Enable. This is used to maintain compatibility with the NodePortLocal feature gate, which
	// was promoted to GA in v1.14
	enableNodePortLocal bool

	defaultLoadBalancerMode config.LoadBalancerMode
}

func newOptions() *Options { _ = "STUB: not implemented"; return nil }

// addFlags adds flags to fs and binds them to options.
func (o *Options) addFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// complete completes all the required options.
func (o *Options) complete(args []string) error { _ = "STUB: not implemented"; return nil }

// validate validates all the required options. It must be called after complete.
func (o *Options) validate(args []string) error { _ = "STUB: not implemented"; return nil }

func (o *Options) loadConfigFromFile() error { _ = "STUB: not implemented"; return nil }

func (o *Options) setDefaults() { _ = "STUB: not implemented"; return }

func (o *Options) validateTLSOptions() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateAntreaProxyConfig(encapMode config.TrafficEncapModeType) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate Service CIDR configuration if AntreaProxy is not enabled.

func (o *Options) validateFlowExporterConfig() error { _ = "STUB: not implemented"; return nil }

// Only validate and parse the static collector address when the legacy static
// destination is enabled; a bad value here should not prevent the agent from starting
// when only FlowExporterDestination CRs are used.

// Parse the given flowPollInterval config

// Parse the given activeFlowExportTimeout config

// Parse the given inactiveFlowExportTimeout config

func (o *Options) validateMulticastConfig(encapMode config.TrafficEncapModeType, encryptionMode config.TrafficEncryptionModeType) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *Options) validateAntreaIPAMConfig() error { _ = "STUB: not implemented"; return nil }

// TODO(gran): support SNAT for Per-Node IPAM Pods
// SNAT needs to be updated to bypass traffic from AntreaIPAM Pod to Per-Node IPAM Pod

func (o *Options) validateMulticlusterConfig(encapMode config.TrafficEncapModeType, encryptionMode config.TrafficEncryptionModeType) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *Options) setK8sNodeDefaultOptions() { _ = "STUB: not implemented"; return }

//It's okay to set the default value of this field even when AntreaProxy is enabled and the field is not used.

// Regardless of whether the egress feature is enabled, o.config.Egress.SNATFullyRandomPorts should not be nil to prevent a crash in NewClient().

func (o *Options) validateEgressConfig(encapMode config.TrafficEncapModeType) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *Options) validateK8sNodeOptions() error { _ = "STUB: not implemented"; return nil }

// Zero for tunnelPort means Antrea will use the assigned IANA port for a given tunnel protocol.

// Check if the enabled features are supported on the OS.

// When using NoEncap traffic mode without AntreaProxy, Pod-to-Service traffic is handled by kube-proxy
// (iptables/ipvs) in the root netns. If the Endpoint is not local the DNATed traffic will be output to
// the physical network directly without going back to OVS for Egress NetworkPolicy enforcement, which
// breaks basic security functionality. Therefore, we usually do not allow the NoEncap traffic mode without
// AntreaProxy. But one can bypass this check and force this feature combination to be allowed, by defining
// the ALLOW_NO_ENCAP_WITHOUT_ANTREA_PROXY environment variable and setting it to true. This may lead to
// better performance when using NoEncap if Egress NetworkPolicy enforcement is not required.

// In the NetworkPolicyOnly mode, Antrea will not perform SNAT
// (but SNAT can be done by the primary CNI).

// Unlike checkUnsupportedFeatures, validateConfigForPlatform runs after all validations and
// after all fields in the Options struct have been initialized (e.g., enableProxy).

// resetVMDefaultFeatures sets the feature's default enablement status as false if it is not supported on a VM or a BM.
func (o *Options) resetVMDefaultFeatures() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateExternalNodeOptions() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validatePolicyBypassRulesConfig() error { _ = "STUB: not implemented"; return nil }

func (o *Options) setExternalNodeDefaultOptions() {
	_ = "STUB: not implemented"
	// Following options are default values for agent running on a Virtual Machine.
	// They are set to avoid unexpected agent crash.
	return
}

// Regardless of whether the egress feature is enabled, o.config.Egress.SNATFullyRandomPorts should not be nil to prevent a crash in NewClient().

func (o *Options) setMulticlusterDefaultOptions() { _ = "STUB: not implemented"; return }

func (o *Options) setAuditLoggingDefaultOptions() { _ = "STUB: not implemented"; return }

func (o *Options) validateSecondaryNetworkConfig() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateNodePortLocalConfig() error { _ = "STUB: not implemented"; return nil }

func (o *Options) validateHostNetworkModeOptions() error { _ = "STUB: not implemented"; return nil }
