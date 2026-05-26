// Copyright 2024 Antrea Authors
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

package linkmonitor

import (
	"sync"

	"github.com/vishvananda/netlink"
	"k8s.io/utils/set"

	utilnetlink "antrea.io/antrea/v2/pkg/agent/util/netlink"
)

const (
	linkAny = ""
)

type linkMonitor struct {
	mutex             sync.RWMutex
	cacheSynced       bool
	linkSubscribeFunc func(ch chan<- netlink.LinkUpdate, done <-chan struct{}, options netlink.LinkSubscribeOptions) error
	eventHandlers     map[string][]LinkEventHandler
	linkNames         set.Set[string]  // known link names
	linkIndexMap      map[int32]string // map from link index to link name
	netlink           utilnetlink.Interface
}

func NewLinkMonitor() *linkMonitor { _ = "STUB: not implemented"; return nil }

func (d *linkMonitor) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (d *linkMonitor) AddEventHandler(handler LinkEventHandler, linkNames ...string) {
	_ = "STUB: not implemented"
	return
}

func (d *linkMonitor) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (d *linkMonitor) listAndWatchLinks(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// For link rename events, notify handlers watching the original name

func (d *linkMonitor) notifyHandlers(linkName string) { _ = "STUB: not implemented"; return }

// LinkExists checks if the provided interface is configured on the Node.
func (d *linkMonitor) LinkExists(name string) bool { _ = "STUB: not implemented"; return false }

func (d *linkMonitor) addLinkName(name string) { _ = "STUB: not implemented"; return }

func (d *linkMonitor) deleteLinkName(name string) { _ = "STUB: not implemented"; return }
