// Copyright 2020 Antrea Authors
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
	"errors"
	"sync"
	"testing"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	v1net "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/test/e2e/utils"
)

const (
	// standardProbeCount is the number of times we try to probe connectivity.
	standardProbeCount = 3
)

var ErrPodNotFound = errors.New("pod not found")

type KubernetesUtils struct {
	*TestData
	podCache map[string][]v1.Pod
	podLock  sync.Mutex
}

func NewKubernetesUtils(data *TestData) (*KubernetesUtils, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TestCase is a collection of TestSteps to be tested against.
type TestCase struct {
	Name  string
	Steps []*TestStep
}

// TestStep is a single unit of testing spec. It includes the policy specs that need to be
// applied for this test, the port to test traffic on and the expected Reachability matrix.
type TestStep struct {
	Name           string
	Reachability   *Reachability
	NPEvaluation   *NPEvaluation
	TestResources  []metav1.Object
	Ports          []int32
	Protocol       utils.AntreaPolicyProtocol
	Duration       time.Duration
	CustomProbes   []*CustomProbe
	CustomSetup    func()
	CustomTeardown func()
}

// CustomProbe will spin up (or update) SourcePod and DestPod such that Add event of Pods
// can be tested against expected connectivity among those Pods.
type CustomProbe struct {
	// Create or update a source Pod.
	SourcePod CustomPod
	// Create or update a destination Pod.
	DestPod CustomPod
	// Port on which the probe will be made.
	Port int32
	// Set the expected connectivity.
	ExpectConnectivity PodConnectivityMark
}

type probeResult struct {
	podFrom      Pod
	podTo        Pod
	connectivity PodConnectivityMark
	err          error
}

// TestNamespaceMeta holds the relevant metadata of a test Namespace during initialization.
type TestNamespaceMeta struct {
	Name   string
	Labels map[string]string
}

// GetPodByLabel returns a Pod with the matching Namespace and "pod" label if it's found.
// If the pod is not found, GetPodByLabel returns "ErrPodNotFound".
func (k *KubernetesUtils) GetPodByLabel(ns string, name string) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *KubernetesUtils) getPodsUncached(ns string, key, val string) ([]v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPodsByLabel returns an array of all Pods in the given Namespace having a k/v label pair.
func (k *KubernetesUtils) GetPodsByLabel(ns string, key string, val string) ([]v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *KubernetesUtils) LabelPod(ns, name, key, value string) (*v1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *KubernetesUtils) getTCPv4SourcePortRangeFromPod(podNamespace, podNameLabel string) (int32, int32, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// ProbeCommand generates a command to probe the provider url.
// The executor parameter can be used to change where the prober will run. For example, it could be "ip netns exec NAME"
// to run the prober in another namespace.
func ProbeCommand(url, protocol, executor string) []string { _ = "STUB: not implemented"; return nil }

func (data *TestData) RunProbeCommand(
	podNamespace, podName, containerName string,
	srcName string,
	dstName string,
	cmd []string,
	expectedResult *PodConnectivityMark,
) PodConnectivityMark {
	_ = "STUB: not implemented"
	return *new(PodConnectivityMark)
}

// When determining probe outcome, we prioritize stderr over the error returned by RunCommandFromPod.
// There might be an issue in Pod exec API where it sometimes doesn't return error when the probe fails. See #2394.

// For our UDP rejection cases, agnhost will return:
//   For IPv4: 'UNKNOWN: read udp [src]->[dst]: read: no route to host'
//   For IPv6: 'UNKNOWN: read udp [src]->[dst]: read: permission denied'
// To avoid incorrect identification, we use 'no route to host' and
// `permission denied`, instead of 'UNKNOWN' as key string.
// For our other protocols rejection cases, agnhost will return 'REFUSED'.

// unhandled case

// success

func (k *KubernetesUtils) probe(
	pod *v1.Pod,
	podName string,
	containerName string,
	dstAddr string,
	dstName string,
	port int32,
	protocol utils.AntreaPolicyProtocol,
	expectedResult *PodConnectivityMark,
) PodConnectivityMark {
	_ = "STUB: not implemented"
	return *new(PodConnectivityMark)
}

func (k *KubernetesUtils) pingProbe(
	pod *v1.Pod,
	podName string,
	containerName string,
	dstAddr string,
	dstName string,
) PodConnectivityMark {
	_ = "STUB: not implemented"
	return *new(PodConnectivityMark)
}

// decidePingProbeResult uses the pingProbe stdout to decide the connectivity.
func decidePingProbeResult(stdout string, probeNum int) PodConnectivityMark {
	_ = "STUB: not implemented"
	// Provide stdout example for different connectivity:
	// ================== Connected stdout ==================
	// PING 10.10.1.2 (10.10.1.2) 56(84) bytes of data.
	// 64 bytes from 10.10.1.2: icmp_seq=1 ttl=64 time=0.695 ms
	// 64 bytes from 10.10.1.2: icmp_seq=2 ttl=64 time=0.250 ms
	// 64 bytes from 10.10.1.2: icmp_seq=3 ttl=64 time=0.058 ms
	//
	// --- 10.10.1.2 ping statistics ---
	// 3 packets transmitted, 3 received, 0% packet loss, time 2043ms
	// rtt min/avg/max/mdev = 0.058/0.334/0.695/0.266 ms
	// ======================================================
	// =================== Dropped stdout ===================
	// PING 10.10.1.2 (10.10.1.2) 56(84) bytes of data.
	//
	// --- 10.10.1.2 ping statistics ---
	// 3 packets transmitted, 0 received, 100% packet loss, time 2037ms
	// =======================================================
	// =================== Rejected stdout ===================
	// PING 10.10.1.2 (10.10.1.2) 56(84) bytes of data.
	// From 10.10.1.2 icmp_seq=1 Destination Host Prohibited
	// From 10.10.1.2 icmp_seq=2 Destination Host Prohibited
	// From 10.10.1.2 icmp_seq=3 Destination Host Prohibited
	//
	// --- 10.10.1.2 ping statistics ---
	// 3 packets transmitted, 0 received, +3 errors, 100% packet loss, time 2042ms
	// =======================================================
	// =================== Rejected ICMPv6 stdout ===================
	// PING fd02:0:0:f8::11(fd02:0:0:f8::11) 56 data bytes
	// From fd02:0:0:f8::11 icmp_seq=1 Destination unreachable: Administratively prohibited
	// From fd02:0:0:f8::11 icmp_seq=2 Destination unreachable: Administratively prohibited
	// From fd02:0:0:f8::11 icmp_seq=3 Destination unreachable: Administratively prohibited
	//
	// --- fd02:0:0:f8::11 ping statistics ---
	// 3 packets transmitted, 0 received, +3 errors, 100% packet loss, time 2047ms
	// =======================================================
	return *new(PodConnectivityMark)
}

func (k *KubernetesUtils) digDNS(
	podName string,
	podNamespace string,
	dstAddr string,
	useTCP bool,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//========DiG command stdout example========
//; <<>> DiG 9.16.6 <<>> github.com +tcp
//;; global options: +cmd
//;; Got answer:
//;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 21816
//;; flags: qr aa rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1
//
//;; OPT PSEUDOSECTION:
//; EDNS: version: 0, flags:; udp: 4096
//; COOKIE: 2d7fe493ea37c430 (echoed)
//;; QUESTION SECTION:
//;github.com.			IN	A
//
//;; ANSWER SECTION:
//github.com.		6	IN	A	140.82.113.3
//
//;; Query time: 0 msec
//;; SERVER: 10.96.0.10#53(10.96.0.10)
//;; WHEN: Tue Feb 14 22:34:23 UTC 2023
//;; MSG SIZE  rcvd: 77
//==========================================

// Probe execs into a Pod and checks its connectivity to another Pod. It assumes
// that the target Pod is serving on the input port, and also that agnhost is
// installed. The connectivity from source Pod to all IPs of the target Pod
// should be consistent. Otherwise, Error PodConnectivityMark will be returned.
func (k *KubernetesUtils) Probe(ns1, pod1, ns2, pod2 string, port int32, protocol utils.AntreaPolicyProtocol,
	remoteCluster *KubernetesUtils, expectedResult *PodConnectivityMark) (PodConnectivityMark, error) {
	_ = "STUB: not implemented"
	return *new(PodConnectivityMark), nil
}

func (k *KubernetesUtils) probeAndDecideConnectivity(fromPod, toPod v1.Pod,
	fromPodName, toPodName string, port int32, protocol utils.AntreaPolicyProtocol, expectedResult *PodConnectivityMark) (PodConnectivityMark, error) {
	_ = "STUB: not implemented"
	// Both IPv4 and IPv6 address should be tested.
	return *new(PodConnectivityMark), nil
}

// When probing UDP or SCTP from a non-hostNetwork Pod to a hostNetwork Pod within the same Node, the local
// Antrea gateway IPs should be used as the destination IP, rather than the Node external IPs. If using the Node
// external IPs as destination IPs when probing, the UDP or SCTP reply traffic from hostNetwork Pod will choose
// a source IP address based on the routing decision or outgoing interface, which means that the local Antrea
// gateway IPs will be chosen as the source IP address. As a result, the probing will get a failure because the
// source IP address of reply traffic is unexpected. To accommodate with this case, when the target Pod is a
// hostNetwork one, the source Pod is a non-hostNetwork one, and they are on the same Node, the local Antrea gateway
// IPs are used.

// If it's an IPv6 address, add "[]" around it.

// HACK: inferring container name as c80, c81 etc., for simplicity.

// ProbeAddr execs into a Pod and checks its connectivity to an arbitrary destination
// address.
func (k *KubernetesUtils) ProbeAddr(ns, podLabelKey, podLabelValue, dstAddr string, port int32, protocol utils.AntreaPolicyProtocol, expectedResult *PodConnectivityMark) (PodConnectivityMark, error) {
	_ = "STUB: not implemented"
	return *new(PodConnectivityMark), nil
}

// If it's an IPv6 address, add "[]" around it.

// CreateOrUpdateNamespace is a convenience function for idempotent setup of Namespaces
func (data *TestData) CreateOrUpdateNamespace(n string, labels map[string]string) (*v1.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateOrUpdateDeployment is a convenience function for idempotent setup of deployments
func (data *TestData) CreateOrUpdateDeployment(ns string,
	deploymentName string,
	replicas int32,
	labels map[string]string,
	nodeName string,
	hostNetwork bool) (*appsv1.Deployment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BuildService is a convenience function for building a corev1.Service spec.
func (data *TestData) BuildService(svcName, svcNS string, port, targetPort int, selector map[string]string, serviceType *v1.ServiceType) *v1.Service {
	_ = "STUB: not implemented"
	return nil
}

// CreateOrUpdateService is a convenience function for updating/creating Services.
func (data *TestData) CreateOrUpdateService(svc *v1.Service) (*v1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetService is a convenience function for getting Service
func (data *TestData) GetService(namespace, name string) (*v1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) GetConfigMap(namespace, name string) (*v1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *TestData) UpdateConfigMap(configMap *v1.ConfigMap) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *TestData) CreateConfigMap(configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteService is a convenience function for deleting a Service by Namespace and name.
func (data *TestData) DeleteService(ns, name string) error { _ = "STUB: not implemented"; return nil }

// CleanServices is a convenience function for deleting Services in the cluster.
func (data *TestData) CleanServices(namespaces map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// BuildServiceAccount is a convenience function for building a corev1.SerivceAccount spec.
func (data *TestData) BuildServiceAccount(name, ns string, labels map[string]string) *v1.ServiceAccount {
	_ = "STUB: not implemented"
	return nil
}

// CreateOrUpdateServiceAccount is a convenience function for updating/creating ServiceAccount.
func (data *TestData) CreateOrUpdateServiceAccount(sa *v1.ServiceAccount) (*v1.ServiceAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteServiceAccount is a convenience function for deleting a ServiceAccount by Namespace and name.
func (data *TestData) DeleteServiceAccount(ns, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateOrUpdateNetworkPolicy is a convenience function for updating/creating netpols. Updating is important since
// some tests update a network policy to confirm that mutation works with a CNI.
func (data *TestData) CreateOrUpdateNetworkPolicy(netpol *v1net.NetworkPolicy) (*v1net.NetworkPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetNetworkPolicy is a convenience function for getting k8s NetworkPolicies.
func (data *TestData) GetNetworkPolicy(namespace, name string) (*v1net.NetworkPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteNetworkPolicy is a convenience function for deleting NetworkPolicy by name and Namespace.
func (data *TestData) DeleteNetworkPolicy(ns, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// CleanNetworkPolicies is a convenience function for deleting NetworkPolicies in the provided namespaces.
func (data *TestData) CleanNetworkPolicies(namespaces map[string]TestNamespaceMeta) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateTier is a convenience function for creating an Antrea Policy Tier by name and priority.
func (data *TestData) CreateTier(name string, tierPriority int32) (*crdv1beta1.Tier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTier is a convenience function for getting Tier.
func (data *TestData) GetTier(name string) (*crdv1beta1.Tier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateTier is a convenience function for updating an Antrea Policy Tier.
func (data *TestData) UpdateTier(tier *crdv1beta1.Tier) (*crdv1beta1.Tier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isReferencedError(err error) bool { _ = "STUB: not implemented"; return false }

// The message is set by deleteValidate of tierValidator when deleting a Tier that is referenced by any policies.

// DeleteTier is a convenience function for deleting an Antrea Policy Tier with specific name.
// To avoid flakes caused by antrea-controller not in sync with kube-apiserver, it retries a few times if the failure is
// because the Tier is still referenced.
func (data *TestData) DeleteTier(name string) error { _ = "STUB: not implemented"; return nil }

// CreateOrUpdateCG is a convenience function for idempotent setup of crd/v1beta1 ClusterGroups
func (data *TestData) CreateOrUpdateCG(cg *crdv1beta1.ClusterGroup) (*crdv1beta1.ClusterGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateOrUpdateGroup is a convenience function for idempotent setup of crd/v1beta1 Groups
func (k *KubernetesUtils) CreateOrUpdateGroup(g *crdv1beta1.Group) (*crdv1beta1.Group, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetCG is a convenience function for getting ClusterGroups
func (k *KubernetesUtils) GetCG(name string) (*crdv1beta1.ClusterGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGroup is a convenience function for getting Groups
func (k *KubernetesUtils) GetGroup(namespace, name string) (*crdv1beta1.Group, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteCG is a convenience function for deleting core/v1beta1 ClusterGroup by name.
func (data *TestData) DeleteCG(name string) error { _ = "STUB: not implemented"; return nil }

// DeleteGroup is a convenience function for deleting core/v1beta1 Group by namespace and name.
func (k *KubernetesUtils) DeleteGroup(namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// CleanCGs is a convenience function for deleting all ClusterGroups in the cluster.
func (data *TestData) CleanCGs() error { _ = "STUB: not implemented"; return nil }

// CleanGroups is a convenience function for deleting all Groups in the namespace.
func (k *KubernetesUtils) CleanGroups(namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateOrUpdateACNP is a convenience function for updating/creating AntreaClusterNetworkPolicies.
func (data *TestData) CreateOrUpdateACNP(cnp *crdv1beta1.ClusterNetworkPolicy) (*crdv1beta1.ClusterNetworkPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetACNP is a convenience function for getting AntreaClusterNetworkPolicies.
func (data *TestData) GetACNP(name string) (*crdv1beta1.ClusterNetworkPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteACNP is a convenience function for deleting ACNP by name.
func (data *TestData) DeleteACNP(name string) error { _ = "STUB: not implemented"; return nil }

// CleanACNPs is a convenience function for deleting all Antrea ClusterNetworkPolicies in the cluster.
func (data *TestData) CleanACNPs() error { _ = "STUB: not implemented"; return nil }

// CreateOrUpdateANNP is a convenience function for updating/creating Antrea NetworkPolicies.
func (data *TestData) CreateOrUpdateANNP(annp *crdv1beta1.NetworkPolicy) (*crdv1beta1.NetworkPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetANNP is a convenience function for getting AntreaNetworkPolicies.
func (data *TestData) GetANNP(namespace, name string) (*crdv1beta1.NetworkPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteANNP is a convenience function for deleting ANNP by name and Namespace.
func (data *TestData) DeleteANNP(ns, name string) error { _ = "STUB: not implemented"; return nil }

// CleanANNPs is a convenience function for deleting all Antrea NetworkPolicies in provided namespaces.
func (data *TestData) CleanANNPs(namespaces []string) error { _ = "STUB: not implemented"; return nil }

func (data *TestData) WaitForANNPCreationAndRealization(t *testing.T, namespace string, name string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *TestData) WaitForACNPCreationAndRealization(t *testing.T, name string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (k *KubernetesUtils) waitForPodInNamespace(ns string, pod string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type httpServerReadiness struct {
	*KubernetesUtils
	pods              []Pod
	reachability      *Reachability
	remoteCluster     *KubernetesUtils
	protocolPortPairs map[utils.AntreaPolicyProtocol][]int32
}

func (hsr *httpServerReadiness) isReady() bool { _ = "STUB: not implemented"; return false }

func (k *KubernetesUtils) waitForHTTPServers(allPods []Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// Encapsulate the data needed to perform a probe between pods
type probeVector struct {
	fromPod  Pod
	toPod    Pod
	port     int32
	protocol utils.AntreaPolicyProtocol
}

// Populate the channel with all combinations of probes based on the required ports and protocols
func (hsr *httpServerReadiness) buildProbeVectors(probes chan<- probeVector) {
	_ = "STUB: not implemented"
	return
}

// Calculate the number of probes created across all port protocol permutations
func (hsr *httpServerReadiness) numProbes() int { _ = "STUB: not implemented"; return 0 }

// Spawn a fixed set of workers to complete probing of the servers
func (hsr *httpServerReadiness) spawnProberPool(resultsCh chan *probeResult) {
	_ = "STUB: not implemented"
	return
}

// Tested value as the upper limit for running locally with minimal impacts to CI speeds

// Validates two way connectivity between all pods across all protocol and port permutations
func (hsr *httpServerReadiness) validate() { _ = "STUB: not implemented"; return }

// We will receive the connectivity from podFrom to podTo len(ports) times, where
// ports is the parameter to the Validate method.
// If it's the first time we observe the connectivity from podFrom to podTo, just
// store the connectivity we received in reachability matrix.
// If the connectivity from podFrom to podTo has been observed and is different
// from the connectivity we received, store Error connectivity in reachability
// matrix.

// Validate checks the connectivity between all Pods in both directions with a
// list of ports and a protocol. The connectivity from a Pod to another Pod should
// be consistent across all provided ports. Otherwise, this connectivity will be
// treated as Error.
func (k *KubernetesUtils) Validate(allPods []Pod, reachability *Reachability, ports []int32, protocol utils.AntreaPolicyProtocol) {
	_ = "STUB: not implemented"
	return
}

func (k *KubernetesUtils) ValidateRemoteCluster(remoteCluster *KubernetesUtils, allPods []Pod, reachability *Reachability, port int32, protocol utils.AntreaPolicyProtocol) {
	_ = "STUB: not implemented"
	return
}

func (k *KubernetesUtils) Bootstrap(namespaces map[string]TestNamespaceMeta, podsPerNamespace []string, createNamespaces bool, nodeNames map[string]string, hostNetworks map[string]bool) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convenience label for testing

// Ensure that all the HTTP servers have time to start properly.
// See https://github.com/antrea-io/antrea/issues/472.

func (k *KubernetesUtils) Cleanup(namespaces map[string]TestNamespaceMeta) {
	_ = "STUB: not implemented"
	// Cleanup any cluster-scoped resources.
	return
}

func (k *KubernetesUtils) GetClusterSearchDomain() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
