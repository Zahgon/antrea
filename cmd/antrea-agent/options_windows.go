//go:build windows
// +build windows

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

package main

const (
	defaultNPLPortRange = "40000-41000"
)

func (o *Options) checkUnsupportedFeatures() error { _ = "STUB: not implemented"; return nil }

// First check feature gates.

func (o *Options) validateConfigForPlatform() error {
	_ = "STUB: not implemented"
	// AntreaProxy with proxyAll is required on Windows.
	// The userspace kube-proxy mode (only mode compatible with the Antrea Agent on Windows) was
	// removed in K8s v1.26, hence the requirement for proxyAll.
	// Even prior to that, AntreaProxy was required for correct NetworkPolicy enforcement for
	// Service traffic.
	// While we do not fail initialization at the moment, there should be no valid use case for
	// Antrea on Windows without AntreaProxy + proxyAll.
	return nil
}
