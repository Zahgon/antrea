// Copyright 2026 Antrea Authors
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
	"time"

	corev1 "k8s.io/api/core/v1"
)

const (
	tcpdumpPacketCount         = 10
	tcpdumpTimeout             = 10 * time.Second
	pingResponseTimeoutSeconds = 2
	maxDisplayLines            = 10
)

// deployTcpdumpPod deploys a hostNetwork Pod with tcpdump on the specified Node.
// deploymentName controls the name of the Deployment and Pod created.
// The container is granted NET_RAW and NET_ADMIN capabilities, which are required
// by tcpdump to open raw sockets and enable promiscuous mode on interfaces.
func deployTcpdumpPod(ctx context.Context, testContext *testContext, nodeName, deploymentName string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// runTcpdump runs tcpdump in the given Pod on the specified interface, capturing at most
// tcpdumpPacketCount packets that match filter. The capture is bounded by tcpdumpTimeout.
// A context.DeadlineExceeded error from tcpdump (i.e. timeout with no packets) is treated
// as a non-error and the (possibly empty) stdout is returned.
func runTcpdump(ctx context.Context, testContext *testContext, podName, iface string, filter ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// verifyConnectivity sends pings from the client Pod to the target IP to verify connectivity.
func verifyConnectivity(ctx context.Context, testContext *testContext, clientPodName, targetIP string) error {
	_ = "STUB: not implemented"
	return nil
}

// startBackgroundPing starts a continuous ping in the background from the client Pod to the
// target IP. It returns a cleanup function that cancels the ping and waits for it to stop.
func startBackgroundPing(ctx context.Context, testContext *testContext, clientPodName, targetIP string) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// startBackgroundProbes verifies connectivity to each of the target Pod's IPs and then
// starts a background ping for each IP. It returns a cleanup function that stops all
// background pings and should be called (typically via defer) when captures are done.
func startBackgroundProbes(ctx context.Context, testContext *testContext, clientPodName string, targetPod *corev1.Pod) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// countNonEmptyLines counts non-empty lines in the output.
func countNonEmptyLines(output string) int { _ = "STUB: not implemented"; return 0 }

// displayPacketCapture logs the first maxDisplayLines lines of tcpdump output.
func displayPacketCapture(testContext *testContext, output string) {
	_ = "STUB: not implemented"
	return
}
