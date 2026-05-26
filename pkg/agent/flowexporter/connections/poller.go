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

package connections

import (
	"time"

	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
	"antrea.io/antrea/v2/pkg/util/channel"
)

type Poller struct {
	connTrackDumper ConnTrackDumper

	pollInterval          time.Duration
	v4Enabled             bool
	v6Enabled             bool
	connectUplinkToBridge bool

	notifier channel.Notifier

	zones []uint16
}

func NewPoller(ctDumper ConnTrackDumper, notifier channel.Notifier, pollInterval time.Duration, v4Enabled, v6Enabled, connectUplinkToBridge bool) *Poller {
	_ = "STUB: not implemented"
	return nil
}

func (p *Poller) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Not failing here as errors can be transient and could be resolved in future poll cycles.
// TODO: Come up with a backoff/retry mechanism by increasing poll interval and adding retry timeout

// Poll calls into conntrackDumper interface to dump conntrack flows. It returns the connections
// filtered by zones and number of connections for each address family, as a slice.
// In dual-stack clusters, the slice will contain 2 values (number of IPv4 connections first, then
// number of IPv6 connections).
// TODO: As optimization, only poll invalid/closed connections during every poll, and poll the established connections right before the export.
func (p *Poller) Poll() ([]*connection.Connection, []int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
