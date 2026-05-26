//go:build windows
// +build windows

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

package portcache

import (
	"context"

	"antrea.io/antrea/v2/pkg/agent/nodeportlocal/rules"
)

func addRuleForPort(podPortRules rules.PodPortRules, port int, podIP string, podPort int, protocol string) (ProtocolSocketData, error) {
	_ = "STUB: not implemented"
	// Only the protocol used here should be returned if NetNatStaticMapping rule
	// can be inserted to an unused protocol port.
	return *new(ProtocolSocketData), nil
}

func (pt *PortTable) addRuleforFreePort(podIP string, podPort int, protocol string) (int, ProtocolSocketData, error) {
	_ = "STUB: not implemented"
	return 0, *new(ProtocolSocketData), nil
}

// handle wrap around

// protocol port is already taken

func (pt *PortTable) AddRule(podKey string, podPort int, protocol string, podIP string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//success means port, protocol available.

// Only add rules if the entry does not exist.

// RestoreRules should be called at Antrea Agent startup to restore a set of NPL rules.
func (pt *PortTable) RestoreRules(ctx context.Context, allNPLPorts []rules.PodNodePort) {
	_ = "STUB: not implemented"
	return
}

// This will be handled gracefully by the NPL controller: if there is an
// annotation using this port, it will be removed and replaced with a new
// one with a valid port mapping.

func (pt *PortTable) DeleteRule(podKey string, podPort int, protocol string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete not required when the PortTable entry does not exist

// Calling DeleteRule is idempotent.
