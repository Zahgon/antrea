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

package store

import (
	"k8s.io/apimachinery/pkg/watch"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
	"antrea.io/antrea/v2/pkg/apiserver/storage"
	"antrea.io/antrea/v2/pkg/controller/types"
)

// egressGroupEvent implements storage.InternalEvent.
type egressGroupEvent struct {
	// The current version of the stored EgressGroup.
	CurrGroup *types.EgressGroup
	// The previous version of the stored EgressGroup.
	PrevGroup *types.EgressGroup
	// The key of this EgressGroup.
	Key             string
	ResourceVersion uint64
}

// ToWatchEvent converts the egressGroupEvent to *watch.Event based on the provided Selectors. It has the following features:
// 1. Added event will be generated if the Selectors was not interested in the object but is now.
// 2. Modified event will be generated if the Selectors was and is interested in the object.
// 3. Deleted event will be generated if the Selectors was interested in the object but is not now.
// 4. If nodeName is specified, only GroupMembers that hosted by the Node will be in the event.
func (event *egressGroupEvent) ToWatchEvent(selectors *storage.Selectors, isInitEvent bool) *watch.Event {
	_ = "STUB: not implemented"
	return nil
}

// If nodeName is specified in selectors, only GroupMembers that hosted by the Node should be in the event.

// Watcher is not interested in that object.

// Watcher was not interested in that object but is now, an added event will be generated.

// Watcher was and is interested in that object, a modified event will be generated.

// No change for the watcher.

// Watcher was interested in that object but is not interested now, a deleted event will be generated.

func (event *egressGroupEvent) GetResourceVersion() uint64 { _ = "STUB: not implemented"; return 0 }

var _ storage.GenEventFunc = genEgressGroupEvent

// genEgressGroupEvent generates InternalEvent from the given versions of an EgressGroup.
func genEgressGroupEvent(key string, prevObj, currObj interface{}, rv uint64) (storage.InternalEvent, error) {
	_ = "STUB: not implemented"
	return *new(storage.InternalEvent), nil
}

// ToEgressGroupMsg converts the stored EgressGroup to its message form.
// If includeBody is true, GroupMembers will be copied.
// If nodeName is provided, only GroupMembers that hosted by the Node will be copied.
func ToEgressGroupMsg(in *types.EgressGroup, out *controlplane.EgressGroup, includeBody bool, nodeName *string) {
	_ = "STUB: not implemented"
	return
}

// EgressGroupKeyFunc knows how to get the key of an EgressGroup.
func EgressGroupKeyFunc(obj interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

// NewEgressGroupStore creates a store of EgressGroup.
func NewEgressGroupStore() storage.Interface {
	_ = "STUB: not implemented"
	return *new(storage.Interface)
}

// keyAndSpanSelectFunc returns whether the provided selectors match the key and/or the nodeNames.
func keyAndSpanSelectFunc(selectors *storage.Selectors, key string, obj interface{}) bool {
	_ = "STUB: not implemented"
	// If Key is present in selectors, the provided key must match it.
	return false
}

// If nodeName is present in selectors's Field selector, the provided nodeNames must contain it.

// isSelected determines if the previous and the current version of an object should be selected by the given selectors.
func isSelected(key string, prevObj, currObj interface{}, selectors *storage.Selectors, isInitEvent bool) (bool, bool) {
	_ = "STUB: not implemented"
	// We have filtered out init events that we are not interested in, so the current object must be selected.
	return false, false
}
