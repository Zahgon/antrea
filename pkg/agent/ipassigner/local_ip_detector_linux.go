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

package ipassigner

import (
	"sync"

	"k8s.io/apimachinery/pkg/util/sets"
)

// The devices that should be excluded from Egress.
var excludeEgressDevices = []string{"kube-ipvs0"}

type localIPDetector struct {
	mutex         sync.RWMutex
	localIPs      sets.Set[string]
	cacheSynced   bool
	eventHandlers []LocalIPEventHandler
}

func NewLocalIPDetector() *localIPDetector { _ = "STUB: not implemented"; return nil }

// IsLocalIP checks if the provided IP is configured on the Node.
func (d *localIPDetector) IsLocalIP(ip string) bool { _ = "STUB: not implemented"; return false }

func (d *localIPDetector) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (d *localIPDetector) AddEventHandler(handler LocalIPEventHandler) {
	_ = "STUB: not implemented"
	return
}

func (d *localIPDetector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (d *localIPDetector) notify(ip string, added bool) { _ = "STUB: not implemented"; return }

func (d *localIPDetector) listAndWatchIPAddresses(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// Subscribe IP address update before listing existing IP addresses to prevent event loss.
	return
}

// List existing IP addresses first.

// List existing excluding devices first.

// Ignore IP Addresses events of excluded devices.

// Find IP addresses removed or added during the period it was not watching and call eventHandlers to process them.

// Ignore IP Addresses events of excluded devices.
