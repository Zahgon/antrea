//go:build !windows
// +build !windows

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

package iptables

import (
	"time"
)

const (
	XtablesLockFilePath       = "/var/run/xtables.lock"
	xtablesLockFilePermission = 0600
)

// Lock acquires the provided file lock. It's thread-safe.
// It will block until the lock is acquired or the timeout is reached.
func Lock(lockFilePath string, timeout time.Duration) (func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check whether the lock is available every 200ms.
