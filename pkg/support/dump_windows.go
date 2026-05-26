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

//go:build windows
// +build windows

package support

const (
	antreaWindowsOVSLogDir     = `C:\openvswitch\var\log\openvswitch`
	antreaWindowsKubeletLogDir = `C:\var\log\kubelet`
)

// Todo: Logs for OVS and kubelet are collected from the fixed path currently, more enhancements are needed to support
// collecting them from a configurable path in the future.
func (d *agentDumper) DumpLog(basedir string) error { _ = "STUB: not implemented"; return nil }

// Dump OVS logs.

// Dump kubelet logs.

func (d *agentDumper) DumpHostNetworkInfo(basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *agentDumper) dumpNetworkConfig(basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *agentDumper) dumpHNSResources(basedir string) error { _ = "STUB: not implemented"; return nil }

func (d *agentDumper) DumpMemberlist(basedir string) error {
	_ = "STUB: not implemented"
	// memberlist never runs on Windows.
	return nil
}
