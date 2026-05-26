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

const IsNodeAddressGroupIndex = "isNodeAddressGroup"

// addressGroupEvent implements storage.InternalEvent.
type addressGroupEvent struct {
	// The current version of the stored AddressGroup.
	CurrGroup *types.AddressGroup
	// The previous version of the stored AddressGroup.
	PrevGroup *types.AddressGroup
	// The current version of the transferred AddressGroup, which will be used in Added events.
	CurrObject *controlplane.AddressGroup
	// The previous version of the transferred AddressGroup, which will be used in Deleted events.
	// Note that only metadata will be set in Deleted events for efficiency.
	PrevObject *controlplane.AddressGroup
	// The patch object of the message for transferring, which will be used in Modified events.
	PatchObject *controlplane.AddressGroupPatch
	// The key of this AddressGroup.
	Key             string
	ResourceVersion uint64
}

// ToWatchEvent converts the addressGroupEvent to *watch.Event based on the provided Selectors. It has the following features:
// 1. Added event will be generated if the Selectors was not interested in the object but is now.
// 2. Modified event will be generated if the Selectors was and is interested in the object.
// 3. Deleted event will be generated if the Selectors was interested in the object but is not now.
func (event *addressGroupEvent) ToWatchEvent(selectors *storage.Selectors, isInitEvent bool) *watch.Event {
	_ = "STUB: not implemented"
	return nil
}

// Watcher is not interested in that object.

// Watcher was not interested in that object but is now, an added event will be generated.

// Watcher was and is interested in that object, a modified event will be generated, unless there's no address change.

// Watcher was interested in that object but is not interested now, a deleted event will be generated.

func (event *addressGroupEvent) GetResourceVersion() uint64 { _ = "STUB: not implemented"; return 0 }

// ToAddressGroupMsg converts the stored AddressGroup to its message form.
// If includeBody is true, IPAddresses will be copied.
func ToAddressGroupMsg(in *types.AddressGroup, out *controlplane.AddressGroup, includeBody bool) {
	_ = "STUB: not implemented"
	return
}

var _ storage.GenEventFunc = genAddressGroupEvent

// genAddressGroupEvent generates InternalEvent from the given versions of an AddressGroup.
// It converts the stored AddressGroup to its message form, and calculates the incremental
// message - an AddressGroupPatch object.
func genAddressGroupEvent(key string, prevObj, currObj interface{}, rv uint64) (storage.InternalEvent, error) {
	_ = "STUB: not implemented"
	return *new(storage.InternalEvent), nil
}

// Calculate PatchObject in advance so that we don't need to do it for
// each watcher when generating *event.Event.

// PatchObject will not be generated when only span changes.

// AddressGroupKeyFunc knows how to get the key of an AddressGroup.
func AddressGroupKeyFunc(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewAddressGroupStore creates a store of AddressGroup.
func NewAddressGroupStore() storage.Interface {
	_ = "STUB: not implemented"
	return *new(storage.Interface)
}
