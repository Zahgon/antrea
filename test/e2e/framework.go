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

package e2e

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/component-base/featuregate"
	aggregatorclientset "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"

	"antrea.io/antrea/v2/pkg/agent/config"
	crdclientset "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	agentconfig "antrea.io/antrea/v2/pkg/config/agent"
	controllerconfig "antrea.io/antrea/v2/pkg/config/controller"
	flowaggregatorconfig "antrea.io/antrea/v2/pkg/config/flowaggregator"
	"antrea.io/antrea/v2/test/e2e/providers"
)

var AntreaConfigMap *corev1.ConfigMap

var (
	errConnectionLost = fmt.Errorf("http2: client connection lost")
	errNoAggregators  = fmt.Errorf("no flow aggregators found")
)

const (
	defaultTimeout  = 90 * time.Second
	defaultInterval = 1 * time.Second

	// antreaNamespace is the K8s Namespace in which all Antrea resources are running.
	antreaNamespace             = "kube-system"
	kubeNamespace               = "kube-system"
	flowAggregatorNamespace     = "flow-aggregator"
	flowAggregatorNamespace1    = "flow-aggregator-1"
	flowAggregatorNamespace2    = "flow-aggregator-2"
	antreaConfigVolume          = "antrea-config"
	antreaWindowsConfigVolume   = "antrea-windows-config"
	flowAggregatorConfigVolume  = "flow-aggregator-config"
	antreaDaemonSet             = "antrea-agent"
	antreaWindowsDaemonSet      = "antrea-agent-windows"
	antreaDeployment            = "antrea-controller"
	flowAggregatorDeployment    = "flow-aggregator"
	flowAggregatorCHSecret      = "clickhouse-ca"
	antreaDefaultGW             = "antrea-gw0"
	testAntreaIPAMNamespace     = "antrea-ipam-test"
	testAntreaIPAMNamespace11   = "antrea-ipam-test-11"
	testAntreaIPAMNamespace12   = "antrea-ipam-test-12"
	mcjoinContainerName         = "mcjoin"
	agnhostContainerName        = "agnhost"
	toolboxContainerName        = "toolbox"
	nginxContainerName          = "nginx"
	controllerContainerName     = "antrea-controller"
	ovsContainerName            = "antrea-ovs"
	agentContainerName          = "antrea-agent"
	flowAggregatorContainerName = "flow-aggregator"

	antreaYML               = "antrea.yml"
	antreaIPSecYML          = "antrea-ipsec.yml"
	antreaCovYML            = "antrea-coverage.yml"
	antreaIPSecCovYML       = "antrea-ipsec-coverage.yml"
	flowAggregatorYML       = "flow-aggregator.yml"
	flowAggregator1YML      = "flow-aggregator-1.yml"
	flowAggregator2YML      = "flow-aggregator-2.yml"
	flowVisibilityYML       = "flow-visibility.yml"
	flowVisibilityTLSYML    = "flow-visibility-tls.yml"
	chOperatorYML           = "clickhouse-operator-install-bundle.yml"
	flowVisibilityCHPodName = "chi-clickhouse-clickhouse-0-0-0"
	flowVisibilityNamespace = "flow-visibility"
	defaultBridgeName       = "br-int"
	monitoringNamespace     = "monitoring"
	// #nosec G101: not credentials
	flowAggregatorIPFIXClientTLSSecretName = "ipfix-client-cert"
	// #nosec G101: not credentials
	flowAggregatorIPFIXCASecretName = "ipfix-server-ca"

	cpNodeCoverageDir = "/tmp/antrea-e2e-coverage"

	antreaAgentConfName      = "antrea-agent.conf"
	antreaControllerConfName = "antrea-controller.conf"
	flowAggregatorConfName   = "flow-aggregator.conf"

	agnhostImage        = "registry.k8s.io/e2e-test-images/agnhost:2.40"
	ToolboxImage        = "antrea/toolbox:1.5-1"
	mcjoinImage         = "antrea/mcjoin:v2.9"
	nginxImage          = "antrea/nginx:1.21.6-alpine"
	iisImage            = "mcr.microsoft.com/windows/servercore/iis"
	ipfixCollectorImage = "antrea/ipfix-collector:v0.16.0"

	nginxLBService = "nginx-loadbalancer"

	// Port the IPFIX collector test pod listens on for receiving flow records.
	ipfixCollectorPort                  = "4739"
	exporterFlowPollInterval            = 1 * time.Second
	exporterActiveFlowExportTimeout     = 2 * time.Second
	exporterIdleFlowExportTimeout       = 1 * time.Second
	aggregatorActiveFlowRecordTimeout   = 3500 * time.Millisecond
	aggregatorInactiveFlowRecordTimeout = 6 * time.Second
	aggregatorClickHouseCommitInterval  = 1 * time.Second
	clickHouseHTTPPort                  = "8123"
	defaultCHDatabaseURL                = "tcp://clickhouse-clickhouse.flow-visibility.svc:9000"

	statefulSetRestartAnnotationKey = "antrea-e2e/restartedAt"

	iperfPort    = 5201
	iperfSvcPort = 9999
)

type ClusterNode struct {
	idx              int // 0 for control-plane Node
	name             string
	ipv4Addr         string
	ipv6Addr         string
	podV4NetworkCIDR string
	podV6NetworkCIDR string
	gwV4Addr         string
	gwV6Addr         string
	os               string
}

func (n ClusterNode) ip() string { _ = "STUB: not implemented"; return "" }

type ClusterInfo struct {
	numNodes             int
	podV4NetworkCIDR     string
	podV6NetworkCIDR     string
	svcV4NetworkCIDR     string
	svcV6NetworkCIDR     string
	controlPlaneNodeName string
	controlPlaneNodeIPv4 string
	controlPlaneNodeIPv6 string
	nodes                map[int]*ClusterNode
	nodesOS              map[string]string
	windowsNodes         []int
	k8sServerVersion     string
	k8sServiceHost       string
	k8sServicePort       int32
}

type ExternalInfo struct {
	externalServerIPv4 string
	externalServerIPv6 string

	vlanSubnetIPv4  string
	vlanGatewayIPv4 string
	vlanSubnetIPv6  string
	vlanGatewayIPv6 string
	vlanID          int

	externalFRRIPv4 string
	externalFRRIPv6 string
	externalFRRCID  string
}

var clusterInfo ClusterInfo
var externalInfo ExternalInfo

type TestOptions struct {
	providerName        string
	providerConfigPath  string
	logsExportDir       string
	logsExportOnSuccess bool
	withBench           bool
	enableCoverage      bool
	enableAntreaIPAM    bool
	flowVisibility      bool
	npEvaluation        bool
	coverageDir         string
	skipCases           string
	linuxVMs            string
	windowsVMs          string
	// deployAntrea determines whether to deploy Antrea before running tests. It requires antrea.yml to be present in
	// the home directory of the control-plane Node. Note it doesn't affect the tests that redeploy Antrea themselves.
	deployAntrea bool

	externalAgnhostIPs string
	vlanSubnets        string

	externalFRRIPs string
	// FRR cannot currently be configured remotely over networking. As a result, the e2e tests for BGPPolicy can only
	// be run in a Kind cluster, where the FRR container can be configured using Docker exec with the container ID.
	// TODO: Introduce a BGP router implementation that can be configured remotely over networking to replace FRR.
	// This would allow the e2e tests for BGPPolicy to be run in environments other than just a Kind cluster.
	externalFRRCID string

	flowVisibilityProtocol string
}

type flowVisibilityIPFIXTestOptions struct {
	name            string
	tls             bool
	clientAuth      bool
	includeK8sNames *bool
	includeK8sUIDs  *bool
}

type flowAggregatorTestOptions struct {
	disableTLS         bool
	selectedAggregator int
	numReplicas        int
}

type flowVisibilityTestOptions struct {
	mode                     flowaggregatorconfig.AggregatorMode
	databaseURL              string
	databaseSecureConnection bool
	clusterID                string
	ipfixCollector           flowVisibilityIPFIXTestOptions
	flowAggregator           flowAggregatorTestOptions
}

var testOptions TestOptions

// PodInfo combines OS info with a Pod name. It is useful when choosing commands and options on Pods of different OS (Windows, Linux).
type PodInfo struct {
	Name      string
	OS        string
	NodeName  string
	Namespace string
}

// TestData stores the state required for each test case.
type TestData struct {
	ClusterName        string
	provider           providers.ProviderInterface
	KubeConfig         *restclient.Config
	clientset          kubernetes.Interface
	aggregatorClient   aggregatorclientset.Interface
	CRDClient          crdclientset.Interface
	logsDirForTestCase string
	testNamespace      string
}

var testData *TestData

type PodIPs struct {
	IPv4      *net.IP
	IPv6      *net.IP
	IPStrings []string
}

type deployAntreaOptions int

const (
	deployAntreaDefault deployAntreaOptions = iota
	deployAntreaIPsec
	deployAntreaCoverageOffset
)

func (o deployAntreaOptions) WithCoverage() deployAntreaOptions {
	_ = "STUB: not implemented"
	return *new(deployAntreaOptions)
}

func (o deployAntreaOptions) DeployYML() string { _ = "STUB: not implemented"; return "" }

func (o deployAntreaOptions) String() string { _ = "STUB: not implemented"; return "" }

var (
	deployAntreaOptionsString = [...]string{
		"AntreaDefault",
		"AntreaWithIPSec",
	}
	deployAntreaOptionsYML = [...]string{
		antreaYML,
		antreaIPSecYML,
		antreaCovYML,
		antreaIPSecCovYML,
	}
	flowAggYamls = [...]string{
		flowAggregatorYML,
		flowAggregator1YML,
		flowAggregator2YML,
	}
	flowAggregatorNamespaces = [...]string{
		flowAggregatorNamespace,
		flowAggregatorNamespace1,
		flowAggregatorNamespace2,
	}
)

func (p PodIPs) String() string { _ = "STUB: not implemented"; return "" }

func (p *PodIPs) hasSameIP(p1 *PodIPs) bool { _ = "STUB: not implemented"; return false }

func (p *PodIPs) AsSlice() []*net.IP { _ = "STUB: not implemented"; return nil }

func (p *PodIPs) AsStrings() (ipv4, ipv6 string) { _ = "STUB: not implemented"; return "", "" }

// workerNodeName returns an empty string if there is no worker Node with the provided idx
// (including if idx is 0, which is reserved for the control-plane Node)
func workerNodeName(idx int) string {
	_ = "STUB: not implemented"
	// control-plane Node
	return ""
}

func workerNodeIPv4(idx int) string {
	_ = "STUB: not implemented"
	// control-plane Node
	return ""
}

func workerNodeIPv6(idx int) string {
	_ = "STUB: not implemented"
	// control-plane Node
	return ""
}

// workerNodeIP returns an empty string if there is no worker Node with the provided idx
// (including if idx is 0, which is reserved for the control-plane Node)
func workerNodeIP(idx int) string {
	_ = "STUB: not implemented"
	// control-plane Node
	return ""
}

// nodeGatewayIPs returns the Antrea gateway's IPv4 address and IPv6 address for the provided Node
// (if applicable), in that order.
func nodeGatewayIPs(idx int) (string, string) { _ = "STUB: not implemented"; return "", "" }

func controlPlaneNodeName() string { _ = "STUB: not implemented"; return "" }

func controlPlaneNodeIPv4() string { _ = "STUB: not implemented"; return "" }

func controlPlaneNodeIPv6() string { _ = "STUB: not implemented"; return "" }

// nodeName returns an empty string if there is no Node with the provided idx. If idx is 0, the name
// of the control-plane Node will be returned.
func nodeName(idx int) string { _ = "STUB: not implemented"; return "" }

// nodeIPv4 returns an empty string if there is no Node with the provided idx. If idx is 0, the IPv4
// Address of the control-plane Node will be returned.
func nodeIPv4(idx int) string { _ = "STUB: not implemented"; return "" }

// nodeIPv6 returns an empty string if there is no Node with the provided idx. If idx is 0, the IPv6
// Address of the control-plane Node will be returned.
func nodeIPv6(idx int) string { _ = "STUB: not implemented"; return "" }

// nodeIP returns an empty string if there is no Node with the provided idx. If idx is 0, the IP
// of the control-plane Node will be returned.
func nodeIP(idx int) string { _ = "STUB: not implemented"; return "" }

// isIPv4Enabled returns true if and only if IPv4 is enabled in the cluster.
func isIPv4Enabled() bool { _ = "STUB: not implemented"; return false }

// isIPv6Enabled returns true if and only if IPv6 is enabled in the cluster.
func isIPv6Enabled() bool { _ = "STUB: not implemented"; return false }

func labelNodeRoleControlPlane() string { _ = "STUB: not implemented"; return "" }

func controlPlaneNoScheduleTolerations() []corev1.Toleration {
	_ = "STUB: not implemented"
	// "node-role.kubernetes.io/control-plane" was added in K8s 1.20
	// "node-role.kubernetes.io/master" was removed in K8s 1.24
	return nil
}

func (data *TestData) getDefaultLoadBalancerMode() (config.LoadBalancerMode, error) {
	_ = "STUB: not implemented"
	return *new(config.LoadBalancerMode), nil
}

func (data *TestData) InitProvider(providerName, providerConfigPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// RunCommandOnNode is a convenience wrapper around the Provider interface RunCommandOnNode method.
func (data *TestData) RunCommandOnNode(nodeName string, cmd string) (code int, stdout string, stderr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func (data *TestData) RunCommandOnNodeExt(nodeName, cmd string, envs map[string]string, stdin string, sudo bool) (
	code int, stdout, stderr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func (data *TestData) collectExternalInfo() error { _ = "STUB: not implemented"; return nil }

func (data *TestData) collectClusterInfo() error {
	_ = "STUB: not implemented"
	// retrieve K8s server version
	// this needs to be done first, as there may be dependencies on the
	// version later in this function (e.g., for labelNodeRoleControlPlane()).
	return nil
}

// retrieve Node information

// If multiple control-plane Nodes (HA), we will select the last one in the list

// Retrieve cluster CIDRs

// Retrieve cluster CIDRs for Rancher clusters.

// Retrieve service CIDRs

// Retrieve service CIDRs for Rancher clusters.

// Retrieve kubernetes Service host and Port

func getNodeByName(name string) *ClusterNode { _ = "STUB: not implemented"; return nil }

func (data *TestData) collectPodCIDRs() error { _ = "STUB: not implemented"; return nil }

// CreateNamespace creates the provided namespace.
func (data *TestData) CreateNamespace(namespace string, mutateFunc func(*corev1.Namespace)) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore error if the Namespace already exists

// When Namespace already exists, check phase

func (data *TestData) UpdateNamespace(namespace string, mutateFunc func(*corev1.Namespace)) error {
	_ = "STUB: not implemented"
	return nil
}

// Check Namespace phase

// createNamespaceWithAnnotations creates the Namespace with Annotations.
func (data *TestData) createNamespaceWithAnnotations(namespace string, annotations map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// updateNamespaceWithAnnotations updates the given Namespace with Annotations.
func (data *TestData) updateNamespaceWithAnnotations(namespace string, annotations map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// generateAnnotationsMutateFunc generates a mutate function to add given Annotations to a Namespace.
func (data *TestData) generateNamespaceAnnotationsMutateFunc(annotations map[string]string) func(*corev1.Namespace) {
	_ = "STUB: not implemented"
	return nil
}

// DeleteNamespace deletes the provided Namespace, and waits for deletion to actually complete if timeout>=0
func (data *TestData) DeleteNamespace(namespace string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// To log time statistics

// namespace does not exist, we return right away

// Success

// Keep trying

// deployAntreaCommon deploys Antrea using kubectl on the control-plane Node.
func (data *TestData) deployAntreaCommon(yamlFile string, extraOptions string, waitForAgentRollout bool) error {
	_ = "STUB: not implemented"
	// TODO: use the K8s apiserver when server side apply is available?
	// See https://kubernetes.io/docs/reference/using-api/api-concepts/#server-side-apply
	return nil
}

// deployAntrea deploys Antrea with deploy options.
func (data *TestData) deployAntrea(option deployAntreaOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// deployFlowVisibilityClickHouse deploys ClickHouse operator and DB.
func (data *TestData) deployFlowVisibilityClickHouse(o flowVisibilityTestOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ClickHouseInstallation CRD from ClickHouse Operator install bundle applied soon before
// applying CR. Sometimes apiserver validation fails to recognize resource of
// kind: ClickHouseInstallation. Retry in such scenario.

// check for clickhouse pod Ready. Wait for 2x timeout as ch operator needs to be running first to handle chi

// check clickhouse service http port for service connectivity

func (data *TestData) deleteFlowVisibility() error { _ = "STUB: not implemented"; return nil }

func (data *TestData) deleteClickHouseOperator() error { _ = "STUB: not implemented"; return nil }

func (data *TestData) deployIPFIXCollector(serverCert []byte, serverKey []byte, clientCA []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (data *TestData) deployIPFIXCollectorWithName(name string, serverCert []byte, serverKey []byte, clientCA []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// deployFlowAggregator deploys the Flow Aggregator.
func (data *TestData) deployFlowAggregator(
	ipfixCollectorAddr string,
	ipfixClientCert, ipfixClientKey, ipfixServerCA []byte,
	o flowVisibilityTestOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Create flow-aggregator Namespace first, so that we can create the necessary Secrets prior
// to applying the Flow Aggregator manifest.

// clickhouse-ca Secret is created in the flow-visibility Namespace. In order to make it accessible to the Flow Aggregator,
// we copy it from Namespace flow-visibility to Namespace flow-aggregator when databaseSecureConnection is true.

func (data *TestData) mutateFlowAggregatorConfigMap(ipfixCollectorAddr string, o flowVisibilityTestOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// By default, the YAML manifest used for testing already has CASecretName and
// ClientSecretName set (which is a no-op unless TLS is enabled). However, when
// client auth is disabled by the test, we have to make sure that ClientSecretName
// is set to the empty string.

func (data *TestData) GetFlowAggregatorConfigMap(o flowVisibilityTestOptions) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getAgentContainersRestartCount reads the restart count for every container across all Antrea
// Agent Pods and returns the sum of all the read values.
func (data *TestData) getAgentContainersRestartCount() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// waitForAntreaDaemonSetPods waits for the K8s apiserver to report that all the Antrea Pods are
// available, i.e. all the Nodes have one or more of the Antrea daemon Pod running and available.
func (data *TestData) waitForAntreaDaemonSetPods(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure that all antrea-agent Pods are not terminating. This is required because NumberAvailable of
// DaemonSet counts Pods even if they are terminating. Deleting antrea-agent Pods directly does not cause the
// number to decrease if the process doesn't quit immediately, e.g. when the signal is caught by bincover
// program and triggers coverage calculation.

// waitForCoreDNSPods waits for the K8s apiserver to report that all the CoreDNS Pods are available.
func (data *TestData) waitForCoreDNSPods(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep trying

// restartCoreDNSPods deletes all the CoreDNS Pods to force them to be re-scheduled. It then waits
// for all the Pods to become available, by calling waitForCoreDNSPods.
func (data *TestData) restartCoreDNSPods(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// checkCoreDNSPods checks that all the Pods for the CoreDNS deployment are ready. If not, it
// deletes all the Pods to force them to restart and waits up to timeout for the Pods to become
// ready.
func (data *TestData) checkCoreDNSPods(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// deployment ready, nothing to do

// CreateClient initializes the K8s clientset in the TestData structure.
func (data *TestData) CreateClient(kubeconfigPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteAntrea deletes the Antrea DaemonSet; we use cascading deletion, which means all the Pods created
// by Antrea will be deleted. After issuing the deletion request, we poll the K8s apiserver to ensure
// that the DaemonSet does not exist any more. This function is a no-op if the Antrea DaemonSet does
// not exist at the time the function is called.
func (data *TestData) deleteAntrea(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Foreground deletion policy ensures that by the time the DaemonSet is deleted, there are
// no Antrea Pods left.

// no Antrea DaemonSet running, we return right away

// Antrea DaemonSet does not exist any more, success

// Keep trying

// getImageName gets the image name from the fully qualified URI.
// For example: "gcr.io/kubernetes-e2e-test-images/agnhost:2.8" gets "agnhost".
func getImageName(uri string) string { _ = "STUB: not implemented"; return "" }

type PodBuilder struct {
	Name               string
	Namespace          string
	VolumeMounts       []corev1.VolumeMount
	Volumes            []corev1.Volume
	Image              string
	ContainerName      string
	Command            []string
	Args               []string
	Env                []corev1.EnvVar
	Ports              []corev1.ContainerPort
	HostNetwork        bool
	IsPrivileged       bool
	ServiceAccountName string
	Annotations        map[string]string
	Labels             map[string]string
	NodeName           string
	MutateFunc         func(*corev1.Pod)
	ResourceRequests   corev1.ResourceList
	ResourceLimits     corev1.ResourceList
	ReadinessProbe     *corev1.Probe
	DnsConfig          *corev1.PodDNSConfig
}

func NewPodBuilder(name, ns, image string) *PodBuilder { _ = "STUB: not implemented"; return nil }

func (b *PodBuilder) WithContainerName(ctrName string) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) WithCommand(command []string) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) WithArgs(args []string) *PodBuilder { _ = "STUB: not implemented"; return nil }

func (b *PodBuilder) WithEnv(env []corev1.EnvVar) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) WithPorts(ports []corev1.ContainerPort) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) WithHostNetwork(v bool) *PodBuilder { _ = "STUB: not implemented"; return nil }

func (b *PodBuilder) InHostNetwork() *PodBuilder { _ = "STUB: not implemented"; return nil }

func (b *PodBuilder) Privileged() *PodBuilder { _ = "STUB: not implemented"; return nil }

func (b *PodBuilder) WithServiceAccountName(name string) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) WithAnnotations(annotations map[string]string) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) WithLabels(labels map[string]string) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) OnNode(nodeName string) *PodBuilder { _ = "STUB: not implemented"; return nil }

func (b *PodBuilder) WithMutateFunc(f func(*corev1.Pod)) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) WithResources(ResourceRequests, ResourceLimits corev1.ResourceList) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) AddVolume(volume corev1.Volume) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) AddVolumeMount(volumeMount corev1.VolumeMount) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) MountConfigMap(configMapName string, mountPath string, volumeName string) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) MountSecret(secretName string, mountPath string, volumeName string) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) MountHostPath(hostPath string, hostPathType corev1.HostPathType, mountPath string, volumeName string) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) WithReadinessProbe(probe *corev1.Probe) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithCustomDNSConfig adds a custom DNS Configuration to the Pod spec.
// It ensures that the DNSPolicy is set to 'None' and assigns the provided DNSConfig.
func (b *PodBuilder) WithCustomDNSConfig(dnsConfig *corev1.PodDNSConfig) *PodBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *PodBuilder) Create(data *TestData) error { _ = "STUB: not implemented"; return nil }

// Set it to 1s for immediate shutdown to reduce test run time and to avoid affecting subsequent tests.

// tolerate NoSchedule taint if we want Pod to run on control-plane Node

// Set DNSPolicy to None to allow custom DNSConfig

// Assign the provided DNSConfig to the Pod's DNSConfig field

func (data *TestData) UpdatePod(namespace, name string, mutateFunc func(*corev1.Pod)) error {
	_ = "STUB: not implemented"
	return nil
}

// createMcJoinPodOnNode creates a Pod in the test namespace with a single mcjoin container. The
// Pod will be scheduled on the specified Node (if nodeName is not empty).
func (data *TestData) createMcJoinPodOnNode(name string, ns string, nodeName string, hostNetwork bool) error {
	_ = "STUB: not implemented"
	return nil
}

// createToolboxPodOnNode creates a Pod in the test namespace with a single toolbox container. The
// Pod will be scheduled on the specified Node (if nodeName is not empty).
func (data *TestData) createToolboxPodOnNode(name string, ns string, nodeName string, hostNetwork bool) error {
	_ = "STUB: not implemented"
	return nil
}

// createNginxPodOnNode creates a Pod in the test namespace with a single nginx container. The
// Pod will be scheduled on the specified Node (if nodeName is not empty).
func (data *TestData) createNginxPodOnNode(name string, ns string, nodeName string, hostNetwork bool) error {
	_ = "STUB: not implemented"
	return nil
}

// createServerPod creates a Pod that can listen to specified port and have named port set.
func (data *TestData) createServerPod(name string, ns string, portName string, portNum int32, setHostPort bool, hostNetwork bool) error {
	_ = "STUB: not implemented"
	// See https://github.com/kubernetes/kubernetes/blob/master/test/images/agnhost/porter/porter.go#L17 for the image's detail.
	return nil
}

// If hostPort is to be set, it must match the container port number.

// createCustomPod creates a Pod in given Namespace with custom labels.
func (data *TestData) createServerPodWithLabels(name, ns string, portNum int32, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *TestData) PatchPod(namespace, name string, patch []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// DeletePod deletes a Pod in the test namespace.
func (data *TestData) DeletePod(namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// Deletes a Pod in the test namespace then waits us to timeout for the Pod not to be visible to the
// client anymore.
func (data *TestData) DeletePodAndWait(timeout time.Duration, name string, ns string) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep trying

type PodCondition func(*corev1.Pod) (bool, error)

// PodWaitFor polls the K8s apiserver until the specified Pod is found (in the test Namespace) and
// the condition predicate is met (or until the provided timeout expires).
func (data *TestData) PodWaitFor(timeout time.Duration, name, namespace string, condition PodCondition) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// podWaitForRunning polls the k8s apiserver until the specified Pod is in the "running" state (or
// until the provided timeout expires).
func (data *TestData) podWaitForRunning(timeout time.Duration, name, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// podWaitForReady polls the k8s apiserver until the specified Pod is in the "Ready" status (or
// until the provided timeout expires).
func (data *TestData) podWaitForReady(timeout time.Duration, name, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// podWaitForIPs polls the K8s apiserver until the specified Pod is in the "running" state (or until
// the provided timeout expires). The function then returns the IP addresses assigned to the Pod. If the
// Pod is not using "hostNetwork", the function also checks that an IP address exists in each required
// Address Family in the cluster.
func (data *TestData) podWaitForIPs(timeout time.Duration, name, namespace string) (*PodIPs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// According to the K8s API documentation (https://godoc.org/k8s.io/api/core/v1#PodStatus),
// the PodIP field should only be empty if the Pod has not yet been scheduled, and "running"
// implies scheduled.

func parsePodIPs(pod *corev1.Pod) (*PodIPs, error) { _ = "STUB: not implemented"; return nil, nil }

// deleteAntreaAgentOnNode deletes the antrea-agent Pod on a specific Node and measure how long it
// takes for the Pod not to be visible to the client any more. It also waits for a new antrea-agent
// Pod to be running on the Node.
func (data *TestData) deleteAntreaAgentOnNode(nodeName string, gracePeriodSeconds int64, timeout time.Duration) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// we do not use DeleteCollection directly because we want to ensure the resources no longer
// exist by the time we return

// in the normal case, there should be a single Pod in the list

// Keep trying, at least one Pod left

// wait for new antrea-agent Pod

// keep trying

// getAntreaPodOnNode retrieves the name of the Antrea Pod (antrea-agent-*) running on a specific Node.
func (data *TestData) getAntreaPodOnNode(nodeName string) (podName string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (data *TestData) RunCommandFromAntreaPodOnNode(nodeName string, cmd []string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// getFlowAggregators retrieves all Flow-Aggregator Pods (flow-aggregator-*) with a specific label from the specified namespace.
func (data *TestData) getFlowAggregators(namespace string) ([]corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getAntreaController retrieves the name of the Antrea Controller (antrea-controller-*) running in the k8s cluster.
func (data *TestData) getAntreaController() (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// restartAntreaControllerPod deletes the antrea-controller Pod to force it to be re-scheduled. It then waits
// for the new Pod to become available, and returns it.
func (data *TestData) restartAntreaControllerPod(timeout time.Duration) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// wait for new antrea-controller Pod

// Even though the strategy is "Recreate", the old Pod might still be in terminating state when the new Pod is
// running as this is deleting a Pod manually, not upgrade.
// See https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#recreate-deployment.
// So we should ensure there's only 1 Pod and it's running.

// RestartAntreaAgentPods deletes all the antrea-agent Pods to force them to be re-scheduled. It
// then waits for the new Pods to become available.
func (data *TestData) RestartAntreaAgentPods(timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// validatePodIP checks that the provided IP address is in the Pod Network CIDR for the cluster.
func validatePodIP(podNetworkCIDR string, ip net.IP) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// CreateService creates a service with port and targetPort.
func (data *TestData) CreateService(serviceName, namespace string, port, targetPort int32, selector map[string]string, affinity, nodeLocalExternal bool,
	serviceType corev1.ServiceType, ipFamily *corev1.IPFamily) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateServiceWithAnnotations creates a service with Annotation
func (data *TestData) CreateServiceWithAnnotations(serviceName, namespace string, port, targetPort int32, protocol corev1.Protocol, selector map[string]string, affinity, nodeLocalExternal bool,
	serviceType corev1.ServiceType, ipFamily *corev1.IPFamily, annotations map[string]string, mutators ...func(service *corev1.Service)) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createNginxClusterIPService creates a nginx service with the given name.
func (data *TestData) createNginxClusterIPService(name, namespace string, affinity bool, ipFamily *corev1.IPFamily) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createAgnhostClusterIPService creates a ClusterIP agnhost service with the given name.
func (data *TestData) createAgnhostClusterIPService(serviceName string, affinity bool, ipFamily *corev1.IPFamily) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createAgnhostNodePortService creates a NodePort agnhost service with the given name.
func (data *TestData) createAgnhostNodePortService(serviceName string, affinity, nodeLocalExternal bool, ipFamily *corev1.IPFamily) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createNginxNodePortService creates a NodePort nginx service with the given name.
func (data *TestData) createNginxNodePortService(serviceName, namespace string, affinity, nodeLocalExternal bool, ipFamily *corev1.IPFamily) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) updateServiceExternalTrafficPolicy(serviceName string, nodeLocalExternal bool) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) updateServiceInternalTrafficPolicy(serviceName string, nodeLocalInternal bool) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) updateService(serviceName string, mutateFunc func(service *corev1.Service)) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createAgnhostLoadBalancerService creates a LoadBalancer agnhost service with the given name.
func (data *TestData) createAgnhostLoadBalancerService(serviceName string, affinity, nodeLocalExternal bool, ingressIPs []string, ipFamily *corev1.IPFamily, annotations map[string]string) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) createNginxLoadBalancerService(affinity bool, ingressIPs []string, ipFamily *corev1.IPFamily) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// deleteService deletes the service.
func (data *TestData) deleteService(namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// Deletes a Service in the test namespace then waits us to timeout for the Service not to be visible to the
// client anymore.
func (data *TestData) deleteServiceAndWait(timeout time.Duration, name, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep trying

// createNetworkPolicy creates a network policy with spec.
func (data *TestData) createNetworkPolicy(name string, spec *networkingv1.NetworkPolicySpec) (*networkingv1.NetworkPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// deleteNetworkpolicy deletes the network policy.
func (data *TestData) deleteNetworkpolicy(policy *networkingv1.NetworkPolicy) error {
	_ = "STUB: not implemented"
	return nil
}

// A DNS-1123 subdomain must consist of lower case alphanumeric characters
var lettersAndDigits = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func randSeq(n int) string { _ = "STUB: not implemented"; return "" }

// #nosec G404: random number generator not used for security purposes

// randName generates a DNS-1123 subdomain name
func randName(prefix string) string { _ = "STUB: not implemented"; return "" }

// Run the provided command in the specified Container for the give Pod and returns the contents of
// stdout and stderr as strings. An error either indicates that the command couldn't be run or that
// the command returned a non-zero error code.
func (data *TestData) RunCommandFromPod(podNamespace string, podName string, containerName string, cmd []string) (stdout string, stderr string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func forAllNodes(fn func(nodeName string) error) error { _ = "STUB: not implemented"; return nil }

// forAllMatchingPodsInNamespace invokes the provided function for every Pod currently running on every Node in a given
// namespace and which matches labelSelector criteria.
func (data *TestData) forAllMatchingPodsInNamespace(
	labelSelector, nsName string, fn func(nodeName string, podName string, nsName string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func parseArpingStdout(out string) (sent uint32, received uint32, loss float32, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

// RunPingCommandFromTestPod uses ping to check connectivity between the Pod and the given target Pod IPs.
// If dontFragment is true and size is 0, it will set the size to the maximum value allowed by the Pod's MTU.
func (data *TestData) RunPingCommandFromTestPod(podInfo PodInfo, ns string, targetPodIPs *PodIPs, ctrName string, count int, size int, dontFragment bool) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: GetPodInterfaceMTU should work for Windows.

// 8 ICMP header, 20 IPv4 header, 40 IPv6 header

func (data *TestData) runNetcatCommandFromTestPod(podName string, ns string, server string, port int32) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *TestData) runNetcatCommandFromTestPodWithProtocol(podName string, ns string, containerName string, server string, port int32, protocol string) error {
	_ = "STUB: not implemented"
	// No parameter required for TCP connections.
	return nil
}

// Retrying several times to avoid flakes as the test may involve DNS (coredns) and Service/Endpoints (kube-proxy).

func (data *TestData) runWgetCommandOnToolboxWithRetry(podName string, ns string, url string, maxAttempts int) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (data *TestData) runWgetCommandFromTestPodWithRetry(podName string, ns string, containerName string, url string, maxAttempts int) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (data *TestData) doesOVSPortExist(antreaPodName string, portName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (data *TestData) doesOVSPortExistOnWindows(nodeName, portName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (data *TestData) GetAntreaAgentConf() (*agentconfig.AgentConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) GetEncapMode() (config.TrafficEncapModeType, error) {
	_ = "STUB: not implemented"
	return *new(config.TrafficEncapModeType), nil
}

// default encap mode

func (data *TestData) GetEncryptionnMode() (config.TrafficEncryptionModeType, error) {
	_ = "STUB: not implemented"
	return *new(config.TrafficEncryptionModeType), nil
}

// default encryption mode

func (data *TestData) isProxyAll() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func GetAgentFeatures() (featuregate.FeatureGate, error) {
	_ = "STUB: not implemented"
	return *new(featuregate.FeatureGate), nil
}

func GetControllerFeatures() (featuregate.FeatureGate, error) {
	_ = "STUB: not implemented"
	return *new(featuregate.FeatureGate), nil
}

func (data *TestData) GetAntreaWindowsConfigMap(antreaNamespace string) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) GetAntreaConfigMap(antreaNamespace string) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) getAgentConf(antreaNamespace string) (*agentconfig.AgentConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) GetGatewayInterfaceName(antreaNamespace string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (data *TestData) GetMulticastInterfaces(antreaNamespace string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTransportInterfaceName returns the transport interface name for cluster Nodes, assuming all Nodes have the same one.
func (data *TestData) GetTransportInterfaceName() (string, error) {
	_ = "STUB: not implemented"
	// It assumes all Nodes have the same transport interface name.
	return "", nil
}

func (data *TestData) GetTransportInterfaceForNode(nodeIdx int) (string, *net.IPNet, *net.IPNet, error) {
	_ = "STUB: not implemented"
	return "", nil, nil, nil
}

// Example stdout:
// eth0@if461       UP             172.18.0.2/16 fc00:f853:ccd:e793::2/64 fe80::42:acff:fe12:2/64
// eno1             UP             10.176.3.138/22 fe80::e643:4bff:fe43:a30e/64

// findInterfaceForIP finds the interface name and IPNet for a given node IP.

// no IP set, not an error

// Extract interface name.

// Find matching CIDR.

func (data *TestData) GetPodInterfaceMTU(namespace string, podName string, containerName string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (data *TestData) GetNodeMACAddress(node, device string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// mutateAntreaConfigMap will perform the specified updates on the antrea-agent config and the
// antrea-controller config by updating the antrea-config ConfigMap. It will then restart Agents and
// Controller if needed. Note that if the specified updates do not result in any actual change to
// the ConfigMap, this function is a complete no-op.
func (data *TestData) mutateAntreaConfigMap(
	controllerChanges func(config *controllerconfig.ControllerConfig),
	agentChanges func(config *agentconfig.AgentConfig),
	restartController bool,
	restartAgent bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// for each config (Agent and Controller), we unmarshal twice and apply changes on one of
// the copy. We use the unchanged copy to detect whether any actual change was made to the
// config by the client-provided functions (controllerConfOut and agentChanges).

// as a convenience, we initialize the FeatureGates map if it is nil

// getAgentConf should be able to process both windows and linux configmap.

// as a convenience, we initialize the FeatureGates map if it is nil

// no config was changed, no need to call Update or restart anything

// we currently restart the controller after the agents

func (data *TestData) killProcessAndCollectCovFiles(namespace, podName, containerName, processName, covDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// gracefulExitAntreaController copies the Antrea controller binary coverage data file out before terminating the Pod
func (data *TestData) gracefulExitAntreaController(covDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// gracefulExitAntreaAgent copies the Antrea agent binary coverage data file out before terminating the Pod
func (data *TestData) gracefulExitAntreaAgent(covDir string, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

// gracefulExitFlowAggregators copies the Flow Aggregator binary coverage data file out before terminating the Pod.
func (data *TestData) gracefulExitFlowAggregators(covDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// collectCovFiles collects coverage files from the Pod and saves them to the coverage directory
func (data *TestData) collectCovFiles(podName string, containerName string, nsName string, covDir string) error {
	_ = "STUB: not implemented"
	// copy antctl coverage files from Pod to the coverage directory
	return nil
}

// collectAntctlCovFilesFromControlPlaneNode collects coverage files for the antctl binary from the control-plane Node and saves them to the coverage directory
func (data *TestData) collectAntctlCovFilesFromControlPlaneNode(covDir string) error {
	_ = "STUB: not implemented"
	// copy antctl coverage files from node to the coverage directory
	return nil
}

// readPodFile reads a file from a Pod and returns the file contents as a string.
func (data *TestData) readPodFile(podName string, containerName string, nsName string, fileName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// copyPodFile copies a file from a Pod and save it to specified directory.
func (data *TestData) copyPodFile(podName string, containerName string, nsName string, fileName string, destDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// dump the file from Antrea Pods to disk.

// copyNodeFiles copies a file from a Node and save it to specified directory
func (data *TestData) copyNodeFiles(fileName string, destDir string) error {
	_ = "STUB: not implemented"
	// getNodeWriter creates the file with name nodeName-suffix. It returns nil if the file
	// cannot be created. File must be closed by the caller.
	return nil
}

// dump the file from Nodes to disk.

// createAgnhostPodOnNode creates a Pod in the test namespace with a single agnhost container. The
// Pod will be scheduled on the specified Node (if nodeName is not empty).
func (data *TestData) createAgnhostPodOnNode(name string, ns string, nodeName string, hostNetwork bool) error {
	_ = "STUB: not implemented"
	return nil
}

// createAgnhostPodWithSAOnNode creates a Pod in the test namespace with a single
// agnhost container and a specific ServiceAccount. The Pod will be scheduled on
// the specified Node (if nodeName is not empty).
func (data *TestData) createAgnhostPodWithSAOnNode(name string, ns string, nodeName string, hostNetwork bool, serviceAccountName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *TestData) createDaemonSet(name string, ns string, ctrName string, image string, cmd []string, args []string) (*appsv1.DaemonSet, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Set it to 1s for immediate shutdown to reduce test run time and to avoid affecting subsequent tests.

func (data *TestData) waitForDaemonSetPods(timeout time.Duration, dsName string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *TestData) createStatefulSet(name string, ns string, size int32, ctrName string, image string, cmd []string, args []string, mutateFunc func(*appsv1.StatefulSet)) (*appsv1.StatefulSet, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Set it to 1s for immediate shutdown to reduce test run time and to avoid affecting subsequent tests.

func (data *TestData) updateStatefulSetSize(name string, ns string, size int32) (*appsv1.StatefulSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) restartStatefulSet(name string, ns string) (*appsv1.StatefulSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Modify StatefulSet PodTemplate annotation to trigger a restart for StatefulSet Pods.

func (data *TestData) waitForStatefulSetPods(timeout time.Duration, stsName string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func isConnectionLostError(err error) bool { _ = "STUB: not implemented"; return false }

// retryOnConnectionLostError allows the caller to retry fn in case the error is ConnectionLost.
// e2e script might get ConnectionLost error when accessing k8s apiserver if AntreaIPAM is enabled and antrea-agent is restarted.
func retryOnConnectionLostError(backoff wait.Backoff, fn func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *TestData) checkAntreaAgentInfo(interval time.Duration, timeout time.Duration, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// keep trying

// Validate that the podRef in AntreaAgentInfo matches the name of the current Pod for the Node

// If err is NotFound, we should keep trying

func getPingCommand(count int, size int, os string, ip *net.IP, dontFragment bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// getCommandInFakeExternalNetwork fakes executing a command from external network by creating a netns and link the netns
// with the host network.
func getCommandInFakeExternalNetwork(cmd string, prefixLength int, externalIP string, localIP string, otherLocalIPs ...string) (string, string) {
	_ = "STUB: not implemented"
	// Create another netns to fake an external network on the host network Pod.
	return "", ""
}

// GetPodLogs returns the current logs for the specified Pod container. If container is empty, it
// defaults to only container when there is one container in the Pod.
func (data *TestData) GetPodLogs(ctx context.Context, namespace, name, container string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (data *TestData) runDNSQuery(
	podName string,
	containerName string,
	podNamespace string,
	dstAddr string,
	useTCP bool,
	dnsServiceIP string) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// setPodAnnotation Patches a pod by adding an annotation with a specified key and value.
func (data *TestData) setPodAnnotation(namespace, podName, annotationKey string, annotationValue string) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *TestData) waitForDeploymentReady(t *testing.T, namespace string, name string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *TestData) getAntreaClusterUUID(timeout time.Duration) (uuid.UUID, error) {
	_ = "STUB: not implemented"
	return *new(uuid.UUID), nil
}

func getHttpURL(ip string, port string) string { _ = "STUB: not implemented"; return "" }
