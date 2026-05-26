// Copyright 2020 Antrea Authors
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

//go:build !windows
// +build !windows

package support

import (
	"sync"

	"k8s.io/klog/v2"

	"antrea.io/antrea/v2/pkg/agent/util/nftables"
)

// nftablesIPv4Supported and nftablesIPv6Supported check if the kernel supports nftables.
// They initialize the client once to verify support, but the returned clients are not used.
var nftablesIPv4Supported = sync.OnceValue(func() bool {
	if _, err := nftables.New(true, false); err != nil {
		klog.InfoS("NFTables IPv4 not supported on this Node", "err", err)
		return false
	}
	return true
})

var nftablesIPv6Supported = sync.OnceValue(func() bool {
	if _, err := nftables.New(false, true); err != nil {
		klog.InfoS("NFTables IPv6 not supported on this Node", "err", err)
		return false
	}
	return true
})

func (d *agentDumper) DumpLog(basedir string) error { _ = "STUB: not implemented"; return nil }

func (d *agentDumper) DumpHostNetworkInfo(basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *agentDumper) dumpIPTables(basedir string) error { _ = "STUB: not implemented"; return nil }

func (d *agentDumper) dumpIPSet(basedir string) error { _ = "STUB: not implemented"; return nil }

func (d *agentDumper) dumpNFTables(basedir string) error { _ = "STUB: not implemented"; return nil }

func (d *agentDumper) dumpIPToolInfo(basedir string) error { _ = "STUB: not implemented"; return nil }

// Dump routes from all routing tables (Antrea installs per-Egress custom tables).

func (d *agentDumper) DumpMemberlist(basedir string) error { _ = "STUB: not implemented"; return nil }
