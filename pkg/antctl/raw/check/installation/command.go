// Copyright 2024 Antrea Authors.
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

package installation

import (
	"context"
	"regexp"
	"time"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"antrea.io/antrea/v2/pkg/antctl/raw/check"
	antrea "antrea.io/antrea/v2/pkg/client/clientset/versioned"
)

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Disables positional arguments

type options struct {
	antreaNamespace string
	runFilter       string
	// Container image for the installation checker.
	testImage          string
	networkPolicyDelay string
}

func newOptions() *options { _ = "STUB: not implemented"; return nil }

const (
	testNamespacePrefix         = "antrea-test"
	clientDeploymentName        = "test-client"
	echoSameNodeDeploymentName  = "echo-same-node"
	echoOtherNodeDeploymentName = "echo-other-node"
	kindEchoName                = "echo"
	kindClientName              = "client"
	agentDaemonSetName          = "antrea-agent"
	controllerDeploymentName    = "antrea-controller"
	podReadyTimeout             = 1 * time.Minute
	minNetworkPolicyDelay       = 1 * time.Second
	controllerAvailableTimeout  = 1 * time.Minute
	agentAvailableTimeout       = 5 * time.Minute
)

type notRunnableError struct {
	reason string
}

func (e notRunnableError) Error() string { _ = "STUB: not implemented"; return "" }

func newNotRunnableError(reason string) notRunnableError {
	_ = "STUB: not implemented"
	return *new(notRunnableError)
}

type Test interface {
	// Run executes the test using the provided testContext. It returns a non-nil error when the test doesn't succeed.
	// If a test is not runnable, notRunnableError should be wrapped in the returned error.
	Run(ctx context.Context, testContext *testContext) error
}

var testsRegistry = make(map[string]Test)

func RegisterTest(name string, test Test) { _ = "STUB: not implemented"; return }

type testContext struct {
	check.Logger
	client               kubernetes.Interface
	antreaClient         antrea.Interface
	config               *rest.Config
	clusterName          string
	antreaNamespace      string
	clientPods           []corev1.Pod
	echoSameNodePod      *corev1.Pod
	echoSameNodeService  *corev1.Service
	echoOtherNodePod     *corev1.Pod
	echoOtherNodeService *corev1.Service
	namespace            string
	// A nil regex indicates that all the tests should be run.
	runFilterRegex *regexp.Regexp
	// Container image for the installation checker.
	testImage          string
	networkPolicyDelay time.Duration
}

type testStats struct {
	numSuccess int
	numFailure int
	numSkipped int
}

func (s *testStats) numTotal() int { _ = "STUB: not implemented"; return 0 }

func compileRunFilter(runFilter string) (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Run(o *options) error { _ = "STUB: not implemented"; return nil }

func tcpProbeCommand(ip string, port int) []string { _ = "STUB: not implemented"; return nil }

func tcpServerCommand(port int) []string { _ = "STUB: not implemented"; return nil }

func newService(name string, selector map[string]string, port int32) *corev1.Service {
	_ = "STUB: not implemented"
	return nil
}

func NewTestContext(
	client kubernetes.Interface,
	antreaClient antrea.Interface,
	config *rest.Config,
	clusterName string,
	antreaNamespace string,
	runFilterRegex *regexp.Regexp,
	testImage string,
	networkPolicyDelay time.Duration,
) *testContext {
	_ = "STUB: not implemented"
	return nil
}

func (t *testContext) setup(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (t *testContext) runTests(ctx context.Context) testStats {
	_ = "STUB: not implemented"
	return *new(testStats)
}

func (t *testContext) tcpProbe(ctx context.Context, clientPodName string, container string, target string, targetPort int) error {
	_ = "STUB: not implemented"
	return nil
}

// We log the contents of stderr here for troubleshooting purposes.

func (t *testContext) Header(format string, a ...interface{}) { _ = "STUB: not implemented"; return }
