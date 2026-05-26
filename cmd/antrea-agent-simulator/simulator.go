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

// The simulator binary is responsible for running simulated antrea agent.
// It watches NetworkPolicies, AddressGroups and AppliedToGroups from antrea controller
// and prints the events of these resources to log.
package main

import (
	"k8s.io/apimachinery/pkg/watch"
)

func run() error { _ = "STUB: not implemented"; return nil }

// Create Antrea Clientset for the given config.

// Create the stop chan with signals

// Generate a context for functions which require one (instead of stopCh).
// We cancel the context when the function returns, which in the normal case will be when
// stopCh is closed.

// Add loop to check whether client is ready

// Wrapper watcher to call watch

// watch NetworkPolicies, AddressGroups, AppliedToGroups

type watchWrapper struct {
	watchFunc func() (watch.Interface, error)
	name      string
}

func (w *watchWrapper) watch() { _ = "STUB: not implemented"; return }

// Call the watch func which is initialized in watchWrapper

// Stop the watcher upon exit

// Watch the init events from chan, and log the events

// Watch the events from chan, and log the events
