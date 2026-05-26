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

import (
	"time"

	"k8s.io/client-go/kubernetes"
)

const informerDefaultResync = 12 * time.Hour

func run(configFile string) error { _ = "STUB: not implemented"; return nil }

// Set up signal capture: the first SIGTERM / SIGINT signal is handled gracefully and will
// cause the stopCh channel to be closed; if another signal is received before the program
// exits, we will force exit.

// Generate a context for functions which require one (instead of stopCh).
// We cancel the context when the function returns, which in the normal case will be when
// stopCh is closed.

func createK8sClient() (kubernetes.Interface, error) {
	_ = "STUB: not implemented"
	return *new(kubernetes.Interface), nil
}
