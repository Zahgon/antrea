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

package nodeip

import (
	"sync"

	corev1 "k8s.io/api/core/v1"

	coreinformers "k8s.io/client-go/informers/core/v1"
)

type Checker interface {
	IsNodeIP(ip string) bool

	HasSynced() bool
}

type Tracker struct {
	nodeInformer coreinformers.NodeInformer
	nodeIPs      map[string]string
	mutex        sync.RWMutex
}

func NewTracker(nodeInformer coreinformers.NodeInformer) *Tracker {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tracker) OnNodeAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (t *Tracker) OnNodeUpdate(oldObj, obj interface{}) { _ = "STUB: not implemented"; return }

func (t *Tracker) OnNodeDelete(obj interface{}) { _ = "STUB: not implemented"; return }

// When the informer's watch connection is interrupted and re-established,
// delete events are delivered as cache.DeletedFinalStateUnknown tombstones.

func (t *Tracker) addNodeIPs(node *corev1.Node) { _ = "STUB: not implemented"; return }

func (t *Tracker) deleteNodeIPs(node *corev1.Node) { _ = "STUB: not implemented"; return }

func (t *Tracker) IsNodeIP(ip string) bool { _ = "STUB: not implemented"; return false }

func (t *Tracker) HasSynced() bool { _ = "STUB: not implemented"; return false }
