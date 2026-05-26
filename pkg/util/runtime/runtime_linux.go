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

package runtime

func parseKernelVersionStr(kernelVersionStr string) (string, error) {
	_ = "STUB: not implemented"
	// "5.4.0-72-generic" is parsed successfully to "v5.4.0-72-generic".
	// "4.13.18-300.el7.x86_64" is reduced to its first three dot-separated
	// components ("4.13.18-300") before being returned as "v4.13.18-300".
	return "", nil
}

// GetKernelVersion returns the Linux kernel version for the current host as a
// canonical semver string (e.g. "v5.4.0").
func GetKernelVersion() (string, error) { _ = "STUB: not implemented"; return "", nil }

// unameBuf.Release is a fixed-size 65-byte array, we need to remove the trailing null
// characters from it first.
