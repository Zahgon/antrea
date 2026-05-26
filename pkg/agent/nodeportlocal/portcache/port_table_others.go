//go:build !windows
// +build !windows

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

func openSocketsForPort(localPortOpener LocalPortOpener, port int, protocol string, isIPv6 bool) (ProtocolSocketData, error) {
	_ = "STUB: not implemented"
	// Port only needs to be available for the protocol used by the NPL rule.
	// We don't need to allocate the same nodePort for all protocols anymore.
	return *new(ProtocolSocketData), nil
}

func (pt *PortTable) getFreePort(podIP string, podPort int, protocol string) (int, ProtocolSocketData, error) {
	_ = "STUB: not implemented"
	return 0, *new(ProtocolSocketData), nil
}

// handle wrap around

// port is already taken

func (pt *PortTable) AddRule(podKey string, podPort int, protocol string, podIP string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Only add rules if the entry does not exist.

func (pt *PortTable) deleteRule(data *NodePortData) error { _ = "STUB: not implemented"; return nil }

// In theory, we should not be modifying a cache item in-place. However, the field we are
// modifying (defunct) does NOT participate in indexing and the modification is thread-safe
// because of pt.tableLock.
// TODO: stop modifying cache items in-place.
// We could set defunct after the call to DeleteRule, because a failed call to DeleteRule
// should mean that the rule is still present and valid, but there is no harm in being more
// conservative.

// Calling DeleteRule is idempotent.

// We don't need to delete cache from different indexes repeatedly because they map to the same entry.
// Deletion errors are not possible because our Index functions cannot return errors.
// See https://github.com/kubernetes/client-go/blob/3aa45779f2e5592d52edf68da66abfbd0805e413/tools/cache/store.go#L189-L196

func (pt *PortTable) DeleteRule(podKey string, podPort int, protocol string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete not required when the PortTable entry does not exist

// syncRules ensures that contents of the port table matches the iptables rules present on the Node.
func (pt *PortTable) syncRules() error { _ = "STUB: not implemented"; return nil }

// RestoreRules should be called at Antrea Agent startup to restore a set of NPL rules. It is
// blocking and no other operations should be performed on the PortTable until the function returns.
func (pt *PortTable) RestoreRules(ctx context.Context, allNPLPorts []rules.PodNodePort) {
	_ = "STUB: not implemented"
	return
}

// This will be handled gracefully by the NPL controller: if there is an
// annotation using this port, it will be removed and replaced with a new
// one with a valid port mapping.

// retry mechanism as iptables-restore can fail if other components (in Antrea or other
// software) are accessing iptables.
