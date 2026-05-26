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

	apis "antrea.io/antrea/v2/pkg/apis"
)

const (
	wireGuardToolboxDeploymentName = "wireguard-tcpdump"
	defaultWireGuardPort           = apis.WireGuardListenPort
	wireGuardInterfaceName         = "antrea-wg0"
)

type WireGuardTest struct{}

func init() {
	RegisterTest("wireguard", &WireGuardTest{})
}

func (t *WireGuardTest) Run(ctx context.Context, testContext *testContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Check 1: plaintext Pod traffic is visible on the WireGuard interface (antrea-wg0).
// This interface carries decrypted traffic on its way in and out of the WireGuard
// tunnel, so Pod IPs should be visible here in plaintext.

// Check 2: WireGuard UDP packets are visible on the transport interface.
// This confirms that traffic leaving the node is encrypted and carried as WireGuard
// UDP on the configured port.

// getNodeTransportInterface returns the name of the transport interface for the given Node
// by reading the AntreaAgentInfo CR. If the field is not set (e.g. older Antrea versions),
// it logs a warning and returns "any" so tcpdump still captures on all interfaces.
func getNodeTransportInterface(ctx context.Context, testContext *testContext, nodeName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
