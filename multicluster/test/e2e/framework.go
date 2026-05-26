// Copyright 2022 Antrea Authors
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
	"os"
	"time"

	corev1 "k8s.io/api/core/v1"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	antreae2e "antrea.io/antrea/v2/test/e2e"
	"antrea.io/antrea/v2/test/e2e/providers"
)

var (
	homedir, _ = os.UserHomeDir()
)

const (
	defaultTimeout     = 90 * time.Second
	importServiceDelay = 2 * time.Second

	multiClusterTestNamespace string = "antrea-multicluster-test"
	eastClusterTestService    string = "east-nginx"
	westClusterTestService    string = "west-nginx"
	mcEastClusterTestService  string = "antrea-mc-east-nginx"
	mcWestClusterTestService  string = "antrea-mc-west-nginx"
	eastCluster               string = "east-cluster"
	westCluster               string = "west-cluster"
	leaderCluster             string = "leader-cluster"
	serviceExportYML          string = "serviceexport.yml"

	testServerPod           string = "test-nginx-pod"
	gatewayNodeClientSuffix string = "gateway-client"
	regularNodeClientSuffix string = "regular-client"

	nginxImage   = "antrea/nginx:1.21.6-alpine"
	agnhostImage = "registry.k8s.io/e2e-test-images/agnhost:2.40"
)

var provider providers.ProviderInterface

type TestOptions struct {
	leaderClusterKubeConfigPath string
	westClusterKubeConfigPath   string
	eastClusterKubeConfigPath   string
	enableGateway               bool
	providerName                string
	logsExportDir               string
}

var testOptions TestOptions

type MCTestData struct {
	clusters            []string
	clusterTestDataMap  map[string]*antreae2e.TestData
	controlPlaneNames   map[string]string
	logsDirForTestCase  string
	clusterGateways     map[string]string
	clusterRegularNodes map[string]string
}

var testData *MCTestData

func (data *MCTestData) createClients() error { _ = "STUB: not implemented"; return nil }

func (data *MCTestData) initProviders() error { _ = "STUB: not implemented"; return nil }

func (data *MCTestData) createTestNamespaces() error { _ = "STUB: not implemented"; return nil }

func (data *MCTestData) deleteTestNamespaces() error { _ = "STUB: not implemented"; return nil }

func (data *MCTestData) patchPod(clusterName, namespace, name string, patch []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *MCTestData) deletePod(clusterName, namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *MCTestData) deletePodAndWait(clusterName, namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *MCTestData) deleteService(clusterName, namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *MCTestData) getService(clusterName, namespace, name string) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *MCTestData) createPod(clusterName, name, nodeName, namespace, ctrName, image string, command []string,
	args []string, env []corev1.EnvVar, ports []corev1.ContainerPort, hostNetwork bool, mutateFunc func(pod *corev1.Pod)) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *MCTestData) updatePod(clusterName string, namespace, name string, mutateFunc func(*corev1.Pod)) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *MCTestData) updateNamespace(clusterName string, namespace string, mutateFunc func(*corev1.Namespace)) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *MCTestData) createService(clusterName, serviceName, namespace string, port int32, targetPort int32,
	protocol corev1.Protocol, selector map[string]string, affinity bool, nodeLocalExternal bool, serviceType corev1.ServiceType,
	ipFamily *corev1.IPFamily, annotation map[string]string) (*corev1.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *MCTestData) createOrUpdateANNP(clusterName string, annp *crdv1beta1.NetworkPolicy) (*crdv1beta1.NetworkPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// deleteANNP is a convenience function for deleting ANNP by name and Namespace.
func (data *MCTestData) deleteANNP(clusterName, namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *MCTestData) createOrUpdateACNP(clusterName string, acnp *crdv1beta1.ClusterNetworkPolicy) (*crdv1beta1.ClusterNetworkPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// deleteACNP is a convenience function for deleting ACNP by name.
func (data *MCTestData) deleteACNP(clusterName, name string) error {
	_ = "STUB: not implemented"
	return nil
}

// podWaitFor polls the K8s apiserver until the specified Pod is found (in the test Namespace) and
// the condition predicate is met (or until the provided timeout expires).
func (data *MCTestData) podWaitFor(timeout time.Duration, clusterName, name, namespace string, condition antreae2e.PodCondition) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (data *MCTestData) probeServiceFromPodInCluster(
	cluster string,
	podName string,
	containerName string,
	podNamespace string,
	serviceIP string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (data *MCTestData) probeFromPodInCluster(
	cluster string,
	podNamespace string,
	podName string,
	containerName string,
	dstAddr string,
	dstName string,
	port int32,
	protocol corev1.Protocol,
) antreae2e.PodConnectivityMark {
	_ = "STUB: not implemented"
	return *new(antreae2e.PodConnectivityMark)
}

// Run the provided command in the specified Container for the given Pod and returns the contents of
// stdout and stderr as strings. An error either indicates that the command couldn't be run or that
// the command returned a non-zero error code.
func (data *MCTestData) runCommandFromPod(clusterName, podNamespace, podName, containerName string, cmd []string) (stdout string, stderr string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
