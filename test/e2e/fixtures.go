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
	"testing"

	"k8s.io/component-base/featuregate"

	"antrea.io/antrea/v2/pkg/agent/config"
)

func skipIfNotBenchmarkTest(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfNotAntreaIPAMTest(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfAntreaIPAMTest(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfNotFlowVisibilityTest(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfNamespaceIsNotEqual(tb testing.TB, actualNamespace, expectNamespace string) {
	_ = "STUB: not implemented"
	return
}

func skipIfProviderIs(tb testing.TB, name string, reason string) { _ = "STUB: not implemented"; return }

func skipIfExternalFRRNotSet(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfNotRequired(tb testing.TB, keys ...string) { _ = "STUB: not implemented"; return }

func skipIfNumNodesLessThan(tb testing.TB, required int) { _ = "STUB: not implemented"; return }

func skipIfRunCoverage(tb testing.TB, reason string) { _ = "STUB: not implemented"; return }

func skipIfNotIPv4Cluster(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfIPv6Cluster(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfNotIPv6Cluster(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfMissingKernelModule(tb testing.TB, data *TestData, nodeName string, requiredModules []string) {
	_ = "STUB: not implemented"
	return
}

// modprobe with "--dry-run" does not require root privileges

func skipIfEncapModeIsNot(tb testing.TB, data *TestData, encapMode config.TrafficEncapModeType) {
	_ = "STUB: not implemented"
	return
}

func skipIfEncapModeIs(tb testing.TB, data *TestData, encapMode config.TrafficEncapModeType) {
	_ = "STUB: not implemented"
	return
}

func skipIfHasWindowsNodes(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfNoWindowsNodes(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfNoVMs(tb testing.TB) { _ = "STUB: not implemented"; return }

func skipIfMulticastEnabled(tb testing.TB, data *TestData) { _ = "STUB: not implemented"; return }

func skipIfFeatureDisabled(tb testing.TB, feature featuregate.Feature, checkAgent bool, checkController bool) {
	_ = "STUB: not implemented"
	return
}

func skipIfProxyDisabled(t *testing.T, data *TestData) { _ = "STUB: not implemented"; return }

func skipIfEgressShapingDisabled(t *testing.T) { _ = "STUB: not implemented"; return }

/* checkAgent */ /* checkController */

func skipIfProxyAllDisabled(t *testing.T, data *TestData) { _ = "STUB: not implemented"; return }

func skipIfFlowExportProtocolIsNotGRPC(t *testing.T) { _ = "STUB: not implemented"; return }

func ensureAntreaRunning(data *TestData) error { _ = "STUB: not implemented"; return nil }

func createDirectory(path string) error { _ = "STUB: not implemented"; return nil }

func (data *TestData) SetupLogDirectoryForTest(testName string) error {
	_ = "STUB: not implemented"
	// sanitize the testName: it can contain '/' if the test is a subtest
	return nil
}

// remove directory if it already exists. This ensures that we start with an empty
// directory

func setupTest(tb testing.TB) (*TestData, error) { _ = "STUB: not implemented"; return nil, nil }

// sanitize the name: the final name must be a valid lowercase RFC 1123 label

// 50 is a conservative length here (we will append a random suffix)

func setupFlowAggregator(tb testing.TB, testData *TestData, o flowVisibilityTestOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func exportLogsForSubtest(tb testing.TB, data *TestData) func() {
	_ = "STUB: not implemented"
	return nil
}

func exportLogs(tb testing.TB, data *TestData, logsSubDir string, writeNodeLogs bool) {
	_ = "STUB: not implemented"
	return
}

// if test was successful and --logs-export-on-success was not provided, we do not export
// any logs.

// for now we just retrieve the logs for the Antrea Pods, but maybe we can find a good way to
// retrieve the logs for the test Pods in the future (before deleting them) if it is useful
// for debugging.

// getPodWriter creates the file with name nodeName-podName-suffix. It returns nil if the
// file cannot be created. File must be closed by the caller.

// runKubectl runs the provided kubectl command on the control-plane Node and returns the
// output. It returns an empty string in case of error.

// dump the logs for Antrea Pods to disk.

// dump the logs for monitoring Pods to disk.

// dump the logs for all flow-aggregator Pods to disk.

// dump the logs for flow-visibility Pods to disk.

// dump the logs for clickhouse operator Pods to disk.

// dump the output of "kubectl describe" for Antrea pods to disk.

// getNodeWriter creates the file with name nodeName-suffix. It returns nil if the file
// cannot be created. File must be closed by the caller.

// export kubelet logs with journalctl for each Node. If the Nodes do not use journalctl we
// print a log message. If kubelet is not run with systemd, the log file will be empty.

// --no-pager ensures the command does not hang.

// return an error and skip subsequent Nodes

// move on to the next Node

// move on to the next VM

func teardownFlowAggregator(tb testing.TB, data *TestData) { _ = "STUB: not implemented"; return }

func teardownTest(tb testing.TB, data *TestData) { _ = "STUB: not implemented"; return }

func deletePodWrapper(tb testing.TB, data *TestData, namespace, name string) {
	_ = "STUB: not implemented"
	return
}

// createTestToolboxPods creates the desired number of toolbox Pods and wait for their IP address to
// become available. This is a common patter in our tests, so having this helper function makes
// sense. It calls Fatalf in case of error, so it must be called from the goroutine running the test
// or benchmark function. You can create all the Pods on the same Node by setting nodeName. If
// nodeName is the empty string, each Pod will be created on an arbitrary
// Node. createTestToolboxPods returns the cleanupFn function which can be used to delete the
// created Pods. Pods are created in parallel to reduce the time required to run the tests.
func createTestToolboxPods(tb testing.TB, data *TestData, num int, ns string, nodeName string) (
	podNames []string, podIPs []*PodIPs, cleanupFn func(),
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func createTestAgnhostPods(tb testing.TB, data *TestData, num int, ns string, nodeName string) (
	podNames []string, podIPs []*PodIPs, cleanupFn func(),
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func createTestPods(tb testing.TB, data *TestData, num int, ns string, nodeName string, hostNetwork bool, createFunc func(string, string, string, bool) error) (
	podNames []string, podIPs []*PodIPs, cleanupFn func(),
) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// setupLogging creates a temporary directory to export the test logs if necessary. If a directory
// was provided by the user, it checks that the directory exists.
func (tOptions *TestOptions) setupLogging() func() { _ = "STUB: not implemented"; return nil }

// we will delete the temporary directory if no logs are exported

// no-op cleanup function

// setupCoverage checks if the directory provided by the user exists.
func (tOptions *TestOptions) setupCoverage(data *TestData) func() {
	_ = "STUB: not implemented"
	return nil
}

// cpNodeCoverageDir is a directory on the control-plane Node, where tests can deposit test
// coverage data.

// best effort

// testMain is meant to be called by TestMain and enables the use of defer statements.
func testMain(m *testing.M) int { _ = "STUB: not implemented"; return 0 }

// Collect PodCIDRs after Antrea is running as Antrea is responsible for allocating PodCIDRs in some cases.
// Polling is not needed here because antrea-agents won't be up and running if PodCIDRs of their Nodes are not set.

func gracefulExitAntrea(testData *TestData) { _ = "STUB: not implemented"; return }

// The following funcs are used in e2e-secondary-network.

func RunTests(m *testing.M) int { _ = "STUB: not implemented"; return 0 }

func SetupTest(tb testing.TB) (*TestData, error) { _ = "STUB: not implemented"; return nil, nil }

func TeardownTest(tb testing.TB, data *TestData) { _ = "STUB: not implemented"; return }

func NodeName(idx int) string { _ = "STUB: not implemented"; return "" }

func NodeCount() int { _ = "STUB: not implemented"; return 0 }

func SkipIfNotIPv4Cluster(tb testing.TB) { _ = "STUB: not implemented"; return }

func SkipIfNotIPv6Cluster(tb testing.TB) { _ = "STUB: not implemented"; return }

func SkipIfIPv6Cluster(tb testing.TB) { _ = "STUB: not implemented"; return }

func SkipIfNotAntreaIPAMTest(tb testing.TB) { _ = "STUB: not implemented"; return }

func (data *TestData) GetTestNamespace() string { _ = "STUB: not implemented"; return "" }
