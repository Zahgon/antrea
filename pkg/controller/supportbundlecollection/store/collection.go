// Copyright 2022 Antrea Authors
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

// supportBundleCollectionEvent implements storage.InternalEvent.
type supportBundleCollectionEvent struct {
	// The current version of the stored SupportBundleCollection.
	currBundleCollection *types.SupportBundleCollection
	// The previous version of the stored SupportBundleCollection.
	prevBundleCollection *types.SupportBundleCollection
	// The key of this SupportBundleCollection.
	Key             string
	ResourceVersion uint64
}

// ToWatchEvent converts the supportBundleCollectionEvent to *watch.Event based on the provided Selectors. It has the following features:
// 1. Added event will be generated if the Selectors was not interested in the object but is now.
// 2. Deleted event will be generated if the Selectors was interested in the object but is not now.
func (event *supportBundleCollectionEvent) ToWatchEvent(selectors *storage.Selectors, isInitEvent bool) *watch.Event {
	_ = "STUB: not implemented"
	return nil
}

// Watcher is not interested in that object.

// Watcher was not interested in that object but is now, an added event will be generated.

// Watcher was interested in that object but is not interested now, a deleted event will be generated.

func (event *supportBundleCollectionEvent) GetResourceVersion() uint64 {
	_ = "STUB: not implemented"
	return 0
}

var _ storage.GenEventFunc = genSupportBundleEvent

// genSupportBundleEvent generates InternalEvent from the given versions of an SupportBundleCollection.
func genSupportBundleEvent(key string, prevObj, currObj interface{}, rv uint64) (storage.InternalEvent, error) {
	_ = "STUB: not implemented"
	return *new(storage.InternalEvent), nil
}

// ToSupportBundleCollectionMsg converts the stored SupportBundleCollection to its message form.
// If includeBody is true, the detailed configurations are copied.
func ToSupportBundleCollectionMsg(in *types.SupportBundleCollection, out *controlplane.SupportBundleCollection, includeBody bool) {
	_ = "STUB: not implemented"
	return
}

// SupportBundleCollectionKeyFunc knows how to get the key of a SupportBundleCollection.
func SupportBundleCollectionKeyFunc(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewSupportBundleCollectionStore creates a store of SupportBundleCollection.
func NewSupportBundleCollectionStore() storage.Interface {
	_ = "STUB: not implemented"
	return *new(storage.Interface)
}

// keyAndSpanSelectFunc returns whether the provided selectors match the key and/or the nodeNames.
func keyAndSpanSelectFunc(selectors *storage.Selectors, key string, obj interface{}) bool {
	_ = "STUB: not implemented"
	// If Key is present in selectors, the provided key must match it.
	return false
}

// If nodeName is present in selectors' Field selector, the provided nodeNames must contain it.

// isSelected determines if the previous and the current version of an object should be selected by the given selectors.
func isSelected(key string, prevObj, currObj interface{}, selectors *storage.Selectors, isInitEvent bool) (bool, bool) {
	_ = "STUB: not implemented"
	// We have filtered out init events that we are not interested in, so the current object must be selected.
	return false, false
}
