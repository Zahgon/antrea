//go:build linux
// +build linux

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

package sysctl

const (
	sysctlNet = "/proc/sys/net"
)

// GetSysctlNet returns the value for sysctl net.* settings
func GetSysctlNet(sysctl string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// SetSysctlNet sets the specified sysctl net.* parameter to the new value.
func SetSysctlNet(sysctl string, newVal int) error {
	_ = "STUB: not implemented"
	// #nosec G306: provided permissions match /proc/sys file permissions
	return nil
}

// EnsureSysctlNetValue checks if the specified sysctl net.* parameter is already set to the
// provided value, and if not, it makes it so.
func EnsureSysctlNetValue(sysctl string, value int) error { _ = "STUB: not implemented"; return nil }

// If permission error, please provide access to sysctl setting

// If permission error, please provide access to sysctl setting
