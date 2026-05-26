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

package hostlocal

import (
	"github.com/spf13/afero"
	"k8s.io/apimachinery/pkg/util/sets"
)

// dataDir is a variable so it can be overridden by tests if needed
var dataDir = "/var/lib/cni/networks"

func networkDir(network string) string { _ = "STUB: not implemented"; return "" }

// This is a hacky approach as we access the internals of the host-local plugin,
// instead of using the CNI interface. However, crafting a CNI DEL request from
// scratch would also be hacky.
func GarbageCollectContainerIPs(network string, desiredIPs sets.Set[string]) error {
	_ = "STUB: not implemented"
	return nil
}

// Internal version of GarbageCollectContainerIPs which does not acquire the
// file lock and can work with an arbitrary afero filesystem.
func gcContainerIPs(fs afero.Fs, dir string, desiredIPs sets.Set[string]) error {
	_ = "STUB: not implemented"
	return nil
}

// not a valid IP, nothing to do

// IP is in-use

// Note that it is perfectly possible for some IPs to be in desiredIPs but not in the
// host-local data directory. This can be the case when another IPAM plugin (e.g.,
// AntreaIPAM) is also used.

func getIPFromPath(path string) string { _ = "STUB: not implemented"; return "" }

// need to unespace IPv6 addresses on Windows
// see https://github.com/containernetworking/plugins/blob/38f18d26ecfef550b8bac02656cc11103fd7cff1/plugins/ipam/host-local/backend/disk/backend.go#L197
