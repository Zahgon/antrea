// Copyright 2026 Antrea Authors.
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

	corev1 "k8s.io/api/core/v1"

	agentconfig "antrea.io/antrea/v2/pkg/config/agent"
)

const (
	antreaConfigMapName        = "antrea-config"
	antreaIPsecSecretName      = "antrea-ipsec" // #nosec G101: false positive triggered by variable name which includes "Secret"
	ipsecToolboxDeploymentName = "ipsec-tcpdump"
	defaultIPsecPSK            = "changeme"
)

type IPsecTest struct{}

// agentConfigInfo holds relevant configuration from the Antrea agent
type agentConfigInfo struct {
	ipsecEnabled bool
	authMode     string
	tunnelType   string
	tunnelPort   int32
}

func init() {
	RegisterTest("ipsec", &IPsecTest{})
}

func (t *IPsecTest) Run(ctx context.Context, testContext *testContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if key has been changed in psk mode

// Log warning but don't fail the test

// Deploy hostNetwork Pod with tcpdump on the same Node as client Pod

// Check if we captured any ESP packets

// Get antrea-agent DaemonSet to check DesiredNumberScheduled

// getAgentConfig retrieves and parses the Antrea agent configuration from the ConfigMap
func getAgentConfig(ctx context.Context, testContext *testContext) (*agentconfig.AgentConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getAgentConfigInfo retrieves and parses all necessary configuration from the agent
func getAgentConfigInfo(ctx context.Context, testContext *testContext) (*agentConfigInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GRE doesn't use a UDP port

// hasPSKBeenChanged returns true if the default PSK has been changed
func hasPSKBeenChanged(ctx context.Context, testContext *testContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// captureTunnelTraffic captures unencrypted tunnel packets using tcpdump.
// If any packets are captured, it means IPsec encryption is not working properly.
func captureTunnelTraffic(ctx context.Context, testContext *testContext, podName, tunnelType string, tunnelPort int32) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getAntreaAgentPod gets the antrea-agent Pod running on the specified Node
func getAntreaAgentPod(ctx context.Context, testContext *testContext, nodeName string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getIPsecStatus runs 'ipsec status' in the antrea-ipsec container
func getIPsecStatus(ctx context.Context, testContext *testContext, agentPodName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// parseIPsecStatus parses the output of 'ipsec status' to extract the number of routed connections and security associations
func parseIPsecStatus(output string) (routedConnections int, securityAssociations int, err error) {
	_ = "STUB: not implemented"
	// Parse routed connections - count unique connection names (without -in/-out suffix if present)
	// GRE format: "worker2-a0d026-1{1}:  ROUTED, TRANSPORT, reqid 1"
	// Geneve format: "worker2-a0d026-in-1{3}:  ROUTED, TRANSPORT, reqid 3"
	//                "worker2-a0d026-out-1{4}:  ROUTED, TRANSPORT, reqid 4"
	// Match the connection name prefix before the final "-<number>{<id>}" part, ignoring optional "-in" or "-out"
	return 0, 0, nil
}

// Use a set to count unique connections

// Extract the connection name prefix (e.g., "worker2-a0d026" from both "worker2-a0d026-1" and "worker2-a0d026-in-1")

// Parse security associations from the summary line
// Example: "Security Associations (1 up, 0 connecting):"
