//go:build linux
// +build linux

// Copyright 2023 Antrea Authors
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

package secondarynetwork

import (
	"net"

	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
)

var (
	// Funcs which will be overridden with mock funcs in tests.
	interfaceByNameFn = net.InterfaceByName
)

// Initialize sets up OVS bridges.
func (c *Controller) Initialize() error {
	_ = "STUB: not implemented"
	// We only support moving and restoring of interface configuration to OVS Bridge for the single physical interface case.
	return nil
}

// do not request a specific MTU

// Restore restores interface configuration from secondary-bridge back to host-interface.
func (c *Controller) Restore() { _ = "STUB: not implemented"; return }

func connectPhyInterfacesToOVSBridge(ovsBridgeClient ovsconfig.OVSBridgeClient, phyInterfaces []string) error {
	_ = "STUB: not implemented"
	return nil
}
