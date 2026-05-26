//go:build linux
// +build linux

// Copyright 2021 Antrea Authors
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

package multicast

import (
	"time"
)

const (
	mRouteTimeout = time.Minute * 10
)

// parseIGMPMsg parses the kernel version into parsedIGMPMsg. Note we need to consider the change
// after linux 5.9 in the igmpmsg struct when parsing vif. Please check
// https://github.com/torvalds/linux/commit/c8715a8e9f38906e73d6d78764216742db13ba0e.
func (c *MRouteClient) parseIGMPMsg(msg []byte) (*parsedIGMPMsg, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// im_mbz in igmpmsg must be zero, as document by
// https://github.com/torvalds/linux/blob/4634129ad9fdc89d10b597fc6f8f4336fb61e105/include/uapi/linux/mroute.h#L115.

// Kernels >= 5.10 use a 16-bit VIF field (two bytes), whereas older kernels use 8 bits.

// detectVIFMode detects once whether the running kernel uses a 16-bit VIF
// field in igmpmsg (kernels >= 5.10) or the legacy 8-bit field, and caches
// the result in c.vif16bit for use in parseIGMPMsg.
func (c *MRouteClient) detectVIFMode() error { _ = "STUB: not implemented"; return nil }

func (c *MRouteClient) run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// When Antrea FlexibleIPAM is enabled, messages received by the socket
// will be dropped directly because we won't create any route from the upcall igmpmsg messages.
// In addition, by reading the socket, we can avoid potential errors such as memory bloat.

// Check packet count difference every minute for each multicast route and
// remove ones that do not route any packets in past mRouteTimeout.
// The remaining multicast routes' statistics are getting updated by
// this process as well.

func (c *MRouteClient) updateMulticastRouteStatsEntry(entry multicastRouteEntry) (isStale bool, newEntry *multicastRouteEntry) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *MRouteClient) updateInboundMrouteStats() { _ = "STUB: not implemented"; return }

func (c *MRouteClient) updateOutboundMrouteStats() { _ = "STUB: not implemented"; return }

func (c *MRouteClient) updateMrouteStats() { _ = "STUB: not implemented"; return }
