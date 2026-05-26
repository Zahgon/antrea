// Copyright 2019 Antrea Authors
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

const (
	IsAppliedToServiceIndex = "isAppliedToService"
	SourceGroupIndex        = "sourceGroup"
)

// appliedToGroupEvent implements storage.InternalEvent.
type appliedToGroupEvent struct {
	// The current version of the stored AppliedToGroup.
	CurrGroup *types.AppliedToGroup
	// The previous version of the stored AppliedToGroup.
	PrevGroup *types.AppliedToGroup
	// The key of this AppliedToGroup.
	Key             string
	ResourceVersion uint64
}

// ToWatchEvent converts the appliedToGroupEvent to *watch.Event based on the provided Selectors. It has the following features:
// 1. Added event will be generated if the Selectors was not interested in the object but is now.
// 2. Modified event will be generated if the Selectors was and is interested in the object.
// 3. Deleted event will be generated if the Selectors was interested in the object but is not now.
// 4. If nodeName is specified, only GroupMembers that hosted by the Node will be in the event.
func (event *appliedToGroupEvent) ToWatchEvent(selectors *storage.Selectors, isInitEvent bool) *watch.Event {
	_ = "STUB: not implemented"
	return nil
}

// If nodeName is specified in selectors, only GroupMembers that hosted by the Node should be in the event.

// Watcher is not interested in that object.

// Watcher was not interested in that object but is now, an added event will be generated.

// Watcher was and is interested in that object, a modified event will be generated.

// No change for the watcher.

// Watcher was interested in that object but is not interested now, a deleted event will be generated.

func (event *appliedToGroupEvent) GetResourceVersion() uint64 { _ = "STUB: not implemented"; return 0 }

var _ storage.GenEventFunc = genAppliedToGroupEvent

// genAppliedToGroupEvent generates InternalEvent from the given versions of an AppliedToGroup.
func genAppliedToGroupEvent(key string, prevObj, currObj interface{}, rv uint64) (storage.InternalEvent, error) {
	_ = "STUB: not implemented"
	return *new(storage.InternalEvent), nil
}

// ToAppliedToGroupMsg converts the stored AppliedToGroup to its message form.
// If includeBody is true, GroupMembers will be copied.
// If nodeName is provided, only GroupMembers that hosted by the Node will be copied.
func ToAppliedToGroupMsg(in *types.AppliedToGroup, out *controlplane.AppliedToGroup, includeBody bool, nodeName *string) {
	_ = "STUB: not implemented"
	return
}

// AppliedToGroupKeyFunc knows how to get the key of an AppliedToGroup.
func AppliedToGroupKeyFunc(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewAppliedToGroupStore creates a store of AppliedToGroup.
func NewAppliedToGroupStore() storage.Interface {
	_ = "STUB: not implemented"
	return *new(storage.Interface)
}
