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

package support

import (
	"time"

	"github.com/spf13/afero"
	"k8s.io/utils/exec"

	agentquerier "antrea.io/antrea/v2/pkg/agent/querier"
	"antrea.io/antrea/v2/pkg/agent/util/ipset"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
	"antrea.io/antrea/v2/pkg/querier"
)

// AgentDumper is the interface for dumping runtime information of the agent. Its
// functions should only work in an agent Pod or a Windows Node which has an agent
// installed.
type AgentDumper interface {
	// DumpFlows should create files that contains flows under the basedir.
	DumpFlows(basedir string) error
	// DumpGroups should create files that contains groups under the basedir
	DumpGroups(basedir string) error
	// DumpHostNetworkInfo should create files that contains host network
	// information under the basedir. Host network information should include
	// links, routes, addresses and etc.
	DumpHostNetworkInfo(basedir string) error
	// DumpLog should create files that contains container logs of the agent
	// Pod under the basedir.
	DumpLog(basedir string) error
	// DumpAgentInfo should create a file that contains AgentInfo of the agent Pod
	// under the basedir.
	DumpAgentInfo(basedir string) error
	// DumpNetworkPolicyResources should create files that contains networkpolicy
	// resources on the agent Pod under the base dir.
	DumpNetworkPolicyResources(basedir string) error
	// DumpHeapPprof should create a pprof file of heap usage of the agent.
	DumpHeapPprof(basedir string) error
	// DumpGoroutinePprof should create a pprof file of goroutine stacks of the agent.
	DumpGoroutinePprof(basedir string) error

	// DumpOVSPorts should create file that contains OF port descriptions under the basedir.
	DumpOVSPorts(basedir string) error
	// DumpMemberlist should create a file that contains state of Memberlist
	// cluster of the agent Pod under the basedir.
	DumpMemberlist(basedir string) error
}

// ControllerDumper is the interface for dumping runtime information of the
// controller. Its functions should only work in the controller Pod.
type ControllerDumper interface {
	// DumpLog should create files that contains container logs of the controller
	// Pod under the basedir.
	DumpLog(basedir string) error
	// DumpControllerInfo should create a file that contains ControllerInfo of
	// the controller Pod under the basedir.
	DumpControllerInfo(basedir string) error
	// DumpNetworkPolicyResources should create files that contains networkpolicy
	// resources on the controller Pod under the base dir.
	DumpNetworkPolicyResources(basedir string) error
	// DumpHeapPprof should create a pprof file of the heap usage of the controller.
	DumpHeapPprof(basedir string) error
	// DumpGoroutinePprof should create a pprof file of goroutine stacks of the controller.
	DumpGoroutinePprof(basedir string) error
}

func DumpHeapPprof(fs afero.Fs, basedir string) error { _ = "STUB: not implemented"; return nil }

func DumpGoroutinePprof(fs afero.Fs, basedir string) error { _ = "STUB: not implemented"; return nil }

func dumpAntctlGet(fs afero.Fs, executor exec.Interface, name, basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func dumpNetworkPolicyResources(fs afero.Fs, executor exec.Interface, basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func timestampFilter(since string) *time.Time { _ = "STUB: not implemented"; return nil }

// parseTimeFromFileName parse time from log file name.
// example log file format: <component>.<hostname>.<user>.log.<level>.<yyyymmdd>-<hhmmss>.1
func parseTimeFromFileName(name string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// parseTimeFromLogLine parse timestamp from the log line.
// example(kubelet/agent/controller): "I0817 06:55:10.804384       1 shared_informer.go:270] caches populated"
// example(ovs): "2021-06-02T16:18:52.285Z|00004|reconnect|INFO|unix:/var/run/openvswitch/db.sock: connecting..."
// the first char indicates the log level.
func parseTimeFromLogLine(log string, year string, prefix string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// directoryCopy copies files under the srcDir to the targetDir. Only files whose name matches
// the prefixFilter will be copied. If prefixFiler is "", no filter is performed. At the same time, if the timeFilter is set,
// only files whose modTime is later than the timeFilter will be copied. If a file contains both older logs and matched logs, only
// the matched logs will be copied. Copied files will be located under the same relative path.
func directoryCopy(fs afero.Fs, targetDir string, srcDir string, prefixFilter string, timeFilter *time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// if name contains timestamp, use it to find the first matched file. If not, such as ovs log file,
// just parse the log file (usually there is only one log file for each component)

// the size limit of single log line is 64k. marked it as known issue and fix it if
// error occurs

// writeFile writes the given data to the specified filePath. Param "resource" is used to identify
// the type of the given data in the error message.
func writeFile(fs afero.Fs, filePath string, resource string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// writeYAMLFile writes the given data to the specified filePath in YAML format. Param "resource" is
// used to identify the type of the given data in the error message.
func writeYAMLFile(fs afero.Fs, filePath string, resource string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type controllerDumper struct {
	fs       afero.Fs
	executor exec.Interface
	since    string
}

func (d *controllerDumper) DumpControllerInfo(basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *controllerDumper) DumpNetworkPolicyResources(basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *controllerDumper) DumpLog(basedir string) error { _ = "STUB: not implemented"; return nil }

func (d *controllerDumper) DumpHeapPprof(basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *controllerDumper) DumpGoroutinePprof(basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func NewControllerDumper(fs afero.Fs, executor exec.Interface, since string) ControllerDumper {
	_ = "STUB: not implemented"
	return *new(ControllerDumper)
}

type agentDumper struct {
	fs           afero.Fs
	executor     exec.Interface
	ovsCtlClient ovsctl.OVSCtlClient
	ipsetClient  ipset.Interface
	aq           agentquerier.AgentQuerier
	npq          querier.AgentNetworkPolicyInfoQuerier
	since        string
	v4Enabled    bool
	v6Enabled    bool
}

func (d *agentDumper) DumpAgentInfo(basedir string) error { _ = "STUB: not implemented"; return nil }

func (d *agentDumper) DumpNetworkPolicyResources(basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *agentDumper) DumpFlows(basedir string) error { _ = "STUB: not implemented"; return nil }

func (d *agentDumper) DumpGroups(basedir string) error { _ = "STUB: not implemented"; return nil }

func (d *agentDumper) DumpHeapPprof(basedir string) error { _ = "STUB: not implemented"; return nil }

func (d *agentDumper) DumpGoroutinePprof(basedir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *agentDumper) DumpOVSPorts(basedir string) error { _ = "STUB: not implemented"; return nil }

func NewAgentDumper(fs afero.Fs, executor exec.Interface, ovsCtlClient ovsctl.OVSCtlClient, aq agentquerier.AgentQuerier, npq querier.AgentNetworkPolicyInfoQuerier, since string, v4Enabled, v6Enabled bool) AgentDumper {
	_ = "STUB: not implemented"
	return *new(AgentDumper)
}
