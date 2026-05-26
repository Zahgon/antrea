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

package proxy

import (
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"

	k8sproxy "antrea.io/antrea/v2/third_party/proxy"
)

type serviceChangesTracker struct {
	tracker *k8sproxy.ServiceChangeTracker

	sync.Mutex
	initialized bool
}

func newServiceChangesTracker(ipFamily v1.IPFamily, serviceLabelSelector labels.Selector, skipServices []string) *serviceChangesTracker {
	_ = "STUB: not implemented"
	return nil
}

func (sh *serviceChangesTracker) OnServiceSynced() { _ = "STUB: not implemented"; return }

func (sh *serviceChangesTracker) OnServiceUpdate(previous, current *v1.Service) bool {
	_ = "STUB: not implemented"
	return false
}

func (sh *serviceChangesTracker) Synced() bool { _ = "STUB: not implemented"; return false }

func (sh *serviceChangesTracker) Update(sm k8sproxy.ServicePortMap) k8sproxy.UpdateServiceMapResult {
	_ = "STUB: not implemented"
	return *new(k8sproxy.UpdateServiceMapResult)
}
