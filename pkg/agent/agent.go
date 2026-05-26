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

package agent

import (
	"context"
	"net"
	"time"

	"github.com/spf13/afero"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	clockutils "k8s.io/utils/clock"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	"antrea.io/antrea/v2/pkg/agent/route"
	"antrea.io/antrea/v2/pkg/agent/types"
	"antrea.io/antrea/v2/pkg/agent/util"
	"antrea.io/antrea/v2/pkg/agent/wireguard"
	"antrea.io/antrea/v2/pkg/client/clientset/versioned"
	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
	utilip "antrea.io/antrea/v2/pkg/util/ip"
	utilwait "antrea.io/antrea/v2/pkg/util/wait"
)

const (
	// Default name of the default tunnel interface on the OVS bridge.
	defaultTunInterfaceName = "antrea-tun0"
	maxRetryForHostLink     = 5
	// ipsecPSKEnvKey is environment variable.
	ipsecPSKEnvKey          = "ANTREA_IPSEC_PSK"
	roundNumKey             = "roundNum" // round number key in externalIDs.
	initialRoundNum         = 1
	maxRetryForRoundNumSave = 5
	// On Linux, OVS configures the MTU for tunnel interfaces to 65000.
	// See https://github.com/openvswitch/ovs/blame/3e666ba000b5eff58da8abb4e8c694ac3f7b08d6/lib/dpif-netlink-rtnl.c#L348-L360
	// There are some edge cases (e.g., Kind clusters) where the transport Node's MTU may be
	// larger than that (e.g., 65535), and packets may be dropped. To account for this, we use
	// 65000 as an upper bound for the MTU calculated in getInterfaceMTU, when encap is
	// supported. For simplicity's sake, we also use this upper bound for Windows, even if it
	// does not apply.
	ovsTunnelMaxMTU = 65000
	// staleFlowDeleteRetryInterval is the wait between DeleteStaleFlows attempts on error.
	staleFlowDeleteRetryInterval = 2 * time.Second
)

var (
	// getIPNetDeviceFromIP is meant to be overridden for testing.
	getIPNetDeviceFromIP = util.GetIPNetDeviceFromIP

	// getIPNetDeviceByV4CIDR is meant to be overridden for testing.
	getIPNetDeviceByCIDRs = util.GetIPNetDeviceByCIDRs

	// getTransportIPNetDeviceByNameFn is meant to be overridden for testing.
	getTransportIPNetDeviceByNameFn = getTransportIPNetDeviceByName

	// setLinkUp is meant to be overridden for testing
	setLinkUp = util.SetLinkUp

	// configureLinkAddresses is meant to be overridden for testing
	configureLinkAddresses = util.ConfigureLinkAddresses
)

// otherConfigKeysForIPsecCertificates are configurations added to OVS bridge when AuthenticationMode is "cert" and
// need to be deleted when changing to "psk".
var otherConfigKeysForIPsecCertificates = []string{"certificate", "private_key", "ca_cert", "remote_cert", "remote_name"}

var (
	// Declared as variables for testing.
	defaultFs                       = afero.NewOsFs()
	clock     clockutils.WithTicker = &clockutils.RealClock{}

	// Maximum time to wait when retrieving the Node object from the K8s API and checking the
	// presence of PodCIDR(s).
	getNodeTimeout = 60 * time.Second
)

// Initializer knows how to setup host networking, OpenVSwitch, and Openflow.
type Initializer struct {
	client                   clientset.Interface
	crdClient                versioned.Interface
	ovsBridgeClient          ovsconfig.OVSBridgeClient
	ovsCtlClient             ovsctl.OVSCtlClient
	ofClient                 openflow.Client
	routeClient              route.Interface
	wireGuardClient          wireguard.Interface
	ifaceStore               interfacestore.InterfaceStore
	ovsBridge                string
	hostGateway              string // name of gateway port on the OVS bridge
	mtu                      int
	networkConfig            *config.NetworkConfig
	nodeConfig               *config.NodeConfig
	wireGuardConfig          *config.WireGuardConfig
	egressConfig             *config.EgressConfig
	serviceConfig            *config.ServiceConfig
	l7NetworkPolicyConfig    *config.L7NetworkPolicyConfig
	enableL7NetworkPolicy    bool
	connectUplinkToBridge    bool
	enableAntreaProxy        bool
	disableTXChecksumOffload bool
	// podNetworkWait should be decremented once the Node's network is ready.
	// The CNI server will wait for it before handling any CNI Add requests.
	podNetworkWait *utilwait.Group
	// flowRestoreCompleteWait is used to indicate that required flows have
	// been installed. We use it to determine whether flows from previous
	// rounds can be deleted.
	flowRestoreCompleteWait *utilwait.Group
	// staleFlowsDeletedWait is created by the caller with NewGroup().Increment();
	// the stale-flow cleanup goroutine in initOpenFlowPipeline calls Done when
	// deletion has finished (or the goroutine exits after a failed delete).
	// Components that must not observe stale flows from a previous agent round
	// should wait on this group before starting.
	staleFlowsDeletedWait *utilwait.Group
	stopCh                <-chan struct{}
	nodeType              config.NodeType
	externalNodeNamespace string
}

func NewInitializer(
	k8sClient clientset.Interface,
	crdClient versioned.Interface,
	ovsBridgeClient ovsconfig.OVSBridgeClient,
	ovsCtlClient ovsctl.OVSCtlClient,
	ofClient openflow.Client,
	routeClient route.Interface,
	ifaceStore interfacestore.InterfaceStore,
	ovsBridge string,
	hostGateway string,
	mtu int,
	networkConfig *config.NetworkConfig,
	wireGuardConfig *config.WireGuardConfig,
	egressConfig *config.EgressConfig,
	serviceConfig *config.ServiceConfig,
	podNetworkWait *utilwait.Group,
	flowRestoreCompleteWait *utilwait.Group,
	staleFlowsDeletedWait *utilwait.Group,
	stopCh <-chan struct{},
	nodeType config.NodeType,
	externalNodeNamespace string,
	connectUplinkToBridge bool,
	enableAntreaProxy bool,
	enableL7NetworkPolicy bool,
	disableTXChecksumOffload bool,
) *Initializer {
	_ = "STUB: not implemented"
	return nil
}

// GetNodeConfig returns the NodeConfig.
func (i *Initializer) GetNodeConfig() *config.NodeConfig { _ = "STUB: not implemented"; return nil }

// GetWireGuardClient returns the Wireguard client.
func (i *Initializer) GetWireGuardClient() wireguard.Interface {
	_ = "STUB: not implemented"
	return *

	// setupOVSBridge sets up the OVS bridge and create host gateway interface and tunnel port
	new(wireguard.Interface)
}

func (i *Initializer) setupOVSBridge() error { _ = "STUB: not implemented"; return nil }

// Wait for the datapath ID for the bridge to be available, as it indicates that the bridge
// has been configured and that we should be able to query supported datapath features.

// Initialize interface cache

// Set up host gateway interface

func (i *Initializer) validateSupportedDPFeatures() error { _ = "STUB: not implemented"; return nil }

// Basic requirements.

// AntreaProxy requires CTStateNAT feature.

// initInterfaceStore initializes InterfaceStore with all OVS ports retrieved
// from the OVS bridge.
func (i *Initializer) initInterfaceStore() error { _ = "STUB: not implemented"; return nil }

// Set the gateway interface name to the discovered name.

// This function can be called for both the "regular" tunnel port or an IPsec tunnel
// port, but the rest of the function only applies to the "regular" tunnel port.

// Set the default tunnel interface name to the discovered name.

// No need to load the OVS host Interface to the interfaceStore

// The port should be for a container interface.

func (i *Initializer) restorePortConfigs() error { _ = "STUB: not implemented"; return nil }

// Initialize sets up agent initial configurations.
func (i *Initializer) Initialize(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// prepareL7EngineInterfaces must be executed after setupOVSBridge since it requires interfaceStore.

// initializeWireGuard must be executed after setupOVSBridge as it requires gateway addresses on the OVS bridge.

// TODO: clean up WireGuard related configurations.

// Initialize for IPsec PSK mode.

// Initialize for IPsec Certificate mode.

// Clean up certificate and private key files.

// Clean up stale configs in OVS database.

// routeClient.Initialize() should be after i.setupOVSBridge() which
// creates the host gateway interface.

// Install OpenFlow entries on OVS bridge.

// Install OpenFlow entries on OVS bridge.

// persistRoundNum will save the provided round number to OVSDB as an external ID. To account for
// transient failures, this (synchronous) function includes a retry mechanism.
func persistRoundNum(num uint64, bridgeClient ovsconfig.OVSBridgeClient, interval time.Duration, maxRetries int) {
	_ = "STUB: not implemented"
	return
}

// success

// deleteStaleFlowsWithRetry calls DeleteStaleFlows until it succeeds or stopCh is closed
// during the wait before another attempt. It returns false when the new round number must not
// be persisted (agent shutting down before stale flows could be deleted).
func (i *Initializer) deleteStaleFlowsWithRetry() bool { _ = "STUB: not implemented"; return false }

// initOpenFlowPipeline sets up necessary Openflow entries, including pipeline, classifiers, conn_track, and gateway flows
// Every time the agent is (re)started, we go through the following sequence:
//  1. agent determines the new round number (this is done by incrementing the round number
//     persisted in OVSDB, or if it's not available by picking round 1).
//  2. any existing flow for which the round number matches the round number obtained from step 1
//     is deleted.
//  3. all required flows are installed, using the round number obtained from step 1.
//  4. after convergence, all existing flows for which the round number matches the previous round
//     number (i.e. the round number which was persisted in OVSDB, if any) are deleted.
//  5. the new round number obtained from step 1 is persisted to OVSDB.
//
// The rationale for not persisting the new round number until after all previous flows have been
// deleted is to avoid a situation in which some stale flows are never deleted because of successive
// agent restarts (with the agent crashing before step 4 can be completed). With the sequence
// described above, We guarantee that at most two rounds of flows exist in the switch at any given
// time.
func (i *Initializer) initOpenFlowPipeline() error { _ = "STUB: not implemented"; return nil }

// Set up all basic flows.

// Delete stale flows from previous round. We need to wait long enough to ensure
// that all the flow which are still required have received an updated cookie (with
// the new round number), otherwise we would disrupt the dataplane. Unfortunately,
// the time required for convergence may be large and there is no simple way to
// determine when is a right time to perform the cleanup task.
// We took a first step towards introducing a deterministic mechanism through which
// the different entities responsible for installing flows can notify the agent that
// this deletion operation can take place. i.flowRestoreCompleteWait.Wait() will
// block until some key flows (NetworkPolicy flows, Pod flows, Node route flows)
// have been installed. But not all entities responsible for installing flows
// currently use this wait group, so we block for a minimum of 10 seconds.
//
// staleFlowsDeletedWait is satisfied when this goroutine exits so that
// components which must not see stale flows (e.g. the NP stats
// collector) can delay their start until cleanup is guaranteed.

// ofClient and ovsBridgeClient have their own mechanisms to restore connections with OVS, and it could
// happen that ovsBridgeClient's connection is not ready when ofClient completes flow replay. We retry it
// with a timeout that is longer time than ovsBridgeClient's maximum connecting retry interval (8 seconds)
// to ensure the flag can be removed successfully.

// This shouldn't happen unless OVS is disconnected again after replaying flows. If it happens, we will try
// to clean up the config again so an error log should be fine.

func (i *Initializer) FlowRestoreComplete() error {
	_ = "STUB: not implemented"
	// Issue #1600: A rare case has been found that the "flow-restore-wait" config was still true even though the delete
	// call below was considered success. At the moment we don't know if it's a race condition caused by "ovs-vsctl set
	// --no-wait" or a problem with OVSDB golang lib or OVSDB itself. To work around it, we check if the config is true
	// before deleting it and if it is false after deleting it, and we will log warnings and retry a few times if
	// anything unexpected happens.
	// If the issue can still happen, it must be that some other code sets the config back after it's deleted.
	return nil
}

// "flow-restore-wait" is supposed to be true here.

// If the log is seen and the config becomes true later, we should look at why "ovs-vsctl set --no-wait"
// doesn't take effect on ovsdb immediately.

// This could happen if the method is triggered by OVS disconnection event, in which OVS doesn't restart.

// ovs-vswitchd is started with flow-restore-wait set to true for the following reasons:
// 1. It prevents packets from being mishandled by ovs-vswitchd in its default fashion,
//    which could affect existing connections' conntrack state and cause issues like #625.
// 2. It prevents ovs-vswitchd from flushing or expiring previously set datapath flows,
//    so existing connections can achieve 0 downtime during OVS restart.
// As a result, we remove the config here after restoring necessary flows.

// If it is seen, we should look at OVSDB golang lib and OVS.

// setupGatewayInterface creates the host gateway interface which is an internal port on OVS. The ofport for host
// gateway interface is predefined, so invoke CreateInternalPort with a specific ofport_request
func (i *Initializer) setupGatewayInterface() error {
	_ = "STUB: not implemented"
	// Create host Gateway port if it does not exist
	return nil
}

// Idempotent operation to set the gateway's MTU: we perform this operation regardless of
// whether the gateway interface already exists, as the desired MTU may change across
// restarts.

// Set arp_announce to 1 on Linux platform to make the ARP requests sent on the gateway
// interface always use the gateway IP as the source IP, otherwise the ARP requests would be
// dropped by ARP SpoofGuard flow.

func (i *Initializer) configureGatewayInterface(gatewayIface *interfacestore.InterfaceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Host link might not be queried at once after creating OVS internal port; retry max 5 times with 1s
// delay each time to ensure the link is ready.

// Persist the MAC configured in the network interface when the gatewayIface.MAC is not set. This may
// happen in upgrade case.
// Note the "mac" field in Windows OVS internal Interface has no impact on the network adapter's actual MAC,
// set it to the same value just to keep consistency.

// Assign IP to gw as required by SpoofGuard.

// No need to assign local CIDR to gw0 because local CIDR is not managed by Antrea

// Allocate the gateway IP address for each Pod CIDR allocated to the Node. For each CIDR,
// the first address in the subnet is assigned to the host gateway interface.

func (i *Initializer) setupDefaultTunnelInterface() error { _ = "STUB: not implemented"; return nil }

// The correct OVS tunnel type to use GRE with an IPv6 overlay is
// "ip6gre" and not "gre". While it would be possible to support GRE for
// an IPv6-only cluster (by simply setting the tunnel type to "ip6gre"),
// things would be more complicated for a dual-stack cluster. For such a
// cluster, we have both IPv4 and IPv6 tunnels for inter-Node
// traffic. We would therefore need to create 2 default tunnel ports:
// one with type "gre" and one with type "ip6gre". This would introduce
// some complexity as the code currently assumes that we have a single
// default tunnel port. So for now, we just reject configurations that
// request a GRE tunnel when the Node network supports IPv6.
// See https://github.com/antrea-io/antrea/issues/3150

// Enabling UDP checksum can greatly improve the performance for Geneve and
// VXLAN tunnels by triggering GRO on the receiver for old Linux kernel versions.
// It's not necessary for new Linux kernel versions with the following patch:
// https://github.com/torvalds/linux/commit/89e5c58fc1e2857ccdaae506fb8bc5fed57ee063.

// Check the default tunnel port.

// Create the default tunnel port and interface.

// Reset the tunnel interface name to the desired name before
// recreating the tunnel port and interface.

func (i *Initializer) setTunnelCsum(tunnelPortName string, enable bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *Initializer) waitForK8sNode(ctx context.Context, nodeName string) (*v1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use a 10s timeout instead of relying on the default dial timeout of 30s. This way we can avoid long
// TCP retry intervals when using the ClusterIP to access the K8s API and when kube-proxy has not
// installed the rules for the kubernetes Service yet.
// If we exceed the deadline, we will have the opportunity to retry several times before the deadline
// for the parent context is exceeded.
// Note that because 10s is greater than the poll interval (5s), the condition function will be called
// again immediately.

// Immediate is false because we just checked the condition by calling hasPodCIDR on a
// "fresh" Node object, and there is no point in getting the Node again until we wait for
// one interval.

// initK8sNodeLocalConfig retrieves node's subnet CIDR from node.spec.PodCIDR, which is used for IPAM and setup
// host gateway interface.
func (i *Initializer) initK8sNodeLocalConfig(ctx context.Context, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

// nodeInterface is the interface that has K8s Node IP. transportInterface is the interface that is used for
// tunneling or routing the traffic across Nodes. It defaults to nodeInterface and can be overridden by the
// configuration parameters TransportInterface and TransportInterfaceCIDRs.

// nodeIPv4Addr and nodeIPv6Addr are the IP addresses of nodeInterface.
// transportIPv4Addr and transportIPv6Addr are the IP addresses of transportInterface.

// Find the interface configured with Node IP and use it for Pod traffic.

// Find the configured transport interface, and update its IP address in Node's annotation.

// Remove the existing annotation "transport-address" if transportInterface is not set in the configuration.

// Update the Node's MAC address in the annotations of the Node. The MAC address will be used for direct routing by
// OVS in noencap case on Windows Nodes. As a mixture of Linux and Windows nodes is possible, Linux Nodes' MAC
// addresses should be reported too to make them discoverable for Windows Nodes.

// Parse all PodCIDRs first, so that we can support IPv4/IPv6 dual-stack configurations.

// at this stage, node.Spec.PodCIDR is guaranteed to NOT be empty

// waitForIPsecMonitorDaemon checks if preconditions are met for using IPsec.
func (i *Initializer) waitForIPsecMonitorDaemon() error {
	_ = "STUB: not implemented"
	// At the time the agent is initialized and this code is executed, the
	// OVS daemons are already running given that we have successfully
	// connected to OVSDB. Given that the start_ovs script deletes existing
	// PID files before starting the OVS daemons, it is safe to assume that
	// if this file exists, the IPsec monitor is indeed running.
	return nil
}

// initializeWireguard checks if preconditions are met for using WireGuard and initializes WireGuard client or cleans up.
func (i *Initializer) initializeWireGuard() error { _ = "STUB: not implemented"; return nil }

// readIPSecPSK reads the IPsec PSK value from environment variable ANTREA_IPSEC_PSK
func (i *Initializer) readIPSecPSK() error { _ = "STUB: not implemented"; return nil }

// Usually one does not want to log the secret data.

func getLastRoundNum(bridgeClient ovsconfig.OVSBridgeClient) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func saveRoundNum(num uint64, bridgeClient ovsconfig.OVSBridgeClient) error {
	_ = "STUB: not implemented"
	return nil
}

func getRoundInfo(bridgeClient ovsconfig.OVSBridgeClient) types.RoundInfo {
	_ = "STUB: not implemented"
	return *new(types.RoundInfo)
}

// We use a fixed value instead of a randomly-generated value to ensure that stale
// flows can be properly deleted in case of multiple rapid restarts when the agent
// is first deployed to a Node.

func (i *Initializer) getInterfaceMTU(transportInterface *net.Interface) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Make sure MTU is set on the interface.

// See comment for ovsTunnelMaxMTU constant above.

func (i *Initializer) allocateGatewayAddresses(localSubnets []*net.IPNet, gatewayIface *interfacestore.InterfaceConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Check IP address configuration on existing interface first, return if the interface has the desired addresses.
// We perform this check unconditionally, even if the OVS port does not exist when this function is called
// (i.e. portExists is false). Indeed, it may be possible for the interface to exist even if the OVS bridge does
// not exist.
// Configure any missing IP address on the interface. Remove any extra IP address that may exist.

// Periodically check whether IP configuration of the gateway is correct.
// Terminate when stopCh is closed.

func (i *Initializer) patchNodeAnnotations(nodeName, key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// getNodeInterfaceFromIP returns the IPv4/IPv6 configuration, and the associated interface according the give nodeIPs.
// When searching the Node interface, antrea-gw0 is ignored because it is configured with the same address as Node IP
// with NetworkPolicyOnly mode on public cloud setup, e.g., EKS.
func (i *Initializer) getNodeInterfaceFromIP(nodeIPs *utilip.DualStackIPs) (v4IPNet *net.IPNet, v6IPNet *net.IPNet, iface *net.Interface, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (i *Initializer) initNodeLocalConfig(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Only IPv4 is supported on a VM Node.

func (i *Initializer) initVMLocalConfig(nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

// prepareOVSBridge operates OVS bridge.
func (i *Initializer) prepareOVSBridge() error { _ = "STUB: not implemented"; return nil }

// setOVSDatapath generates a static datapath ID for OVS bridge so that the OFSwitch identifier is not
// changed after the physical interface is attached on the switch.
func (i *Initializer) setOVSDatapath() error { _ = "STUB: not implemented"; return nil }

// Check if "datapath-id" exists in "other_config" on OVS bridge or not, and return directly if yes.
// Note: function `ovsBridgeClient.GetDatapathID` is not used here, because OVS always has data in "datapath_id"
// field. If "datapath-id" is not explicitly set in "other_config", the datapath ID in use may change when uplink
// is attached on OVS.
