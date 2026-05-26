// Copyright 2019 Antrea Authors
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

package ipam

import (
	"os"

	"github.com/containernetworking/cni/pkg/invoke"
	"github.com/containernetworking/cni/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"

	argtypes "antrea.io/antrea/v2/pkg/agent/cniserver/types"
)

const (
	ipamHostLocal  = "host-local"
	defaultCNIPath = "/opt/cni/bin"
)

type IPAMDelegator struct {
	pluginType string
}

var (
	// Declare these two functions as variable for test
	execPluginWithResultFunc = invoke.ExecPluginWithResult
	execPluginNoResultFunc   = invoke.ExecPluginWithoutResult
)

func (d *IPAMDelegator) Add(args *invoke.Args, k8sArgs *argtypes.K8sArgs, networkConfig []byte) (bool, *IPAMResult, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Rollback to delete assigned network configuration for failed to execute Add operation

// IPAM Delegator always owns the request

func (d *IPAMDelegator) Del(args *invoke.Args, k8sArgs *argtypes.K8sArgs, networkConfig []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IPAM Delegator always owns the request

func (d *IPAMDelegator) Check(args *invoke.Args, k8sArgs *argtypes.K8sArgs, networkConfig []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GarbageCollectContainerIPs will release IPs allocated by the delegated IPAM
// plugin that are no longer in-use (if there is any). It should be called on an
// agent restart to provide garbage collection for IPs, and to avoid IP leakage
// in case of missed CNI DEL events. Normally, it is not Antrea's responsibility
// to implement this, as the above layers should ensure that there is always one
// successful CNI DEL for every corresponding CNI ADD. However, we include this
// support to increase robustness in case of a container runtime bug.
// Only the host-local plugin is supported.
func GarbageCollectContainerIPs(network string, desiredIPs sets.Set[string]) error {
	_ = "STUB: not implemented"
	return nil
}

var defaultExec invoke.Exec = &invoke.DefaultExec{
	RawExec: &invoke.RawExec{Stderr: os.Stderr},
}

func delegateCommon(delegatePlugin string, exec invoke.Exec, cniPath string) (string, invoke.Exec, error) {
	_ = "STUB: not implemented"
	// The CNI searching paths passed from kubelet.
	return "", *new(invoke.Exec), nil
}

// When Antrea agent runs as a Pod, the IPAM plugin is always installed in
// defaultCNIPath, but kubelet can be configured to use different paths to
// search for CNI plugins. So here we always add defaultCNIPath to the CNI
// plugin searching paths to make sure the IPAM plugin installed in the agent
// Pod can be found.

func delegateWithResult(delegatePlugin string, networkConfig []byte, args *invoke.Args) (types.Result, error) {
	_ = "STUB: not implemented"
	return *new(types.Result), nil
}

func delegateNoResult(delegatePlugin string, networkConfig []byte, args *invoke.Args) error {
	_ = "STUB: not implemented"
	return nil
}

func init() {
	RegisterIPAMDriver(ipamHostLocal, &IPAMDelegator{pluginType: ipamHostLocal})
}
