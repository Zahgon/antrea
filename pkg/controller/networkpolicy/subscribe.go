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

package networkpolicy

import "sync"

// notifier notifies multiple subscribers about any events that happen to the objects they have subscribed.
type notifier struct {
	mutex       sync.RWMutex
	subscribers map[string]map[string]func()
}

func newNotifier() *notifier { _ = "STUB: not implemented"; return nil }

// Subscribe the subscriber to the given resourceID with a callback.
// If the subscription already exists, it does nothing.
func (n *notifier) subscribe(resourceID, subscriberID string, callback func()) {
	_ = "STUB: not implemented"
	return
}

// unsubscribe cancels the subscription.
// If the subscription does not exist, it does nothing.
func (n *notifier) unsubscribe(resourceID, subscriberID string) { _ = "STUB: not implemented"; return }

// If the resource is no longer subscribed by any notifier, remove its key.

// Notify the subscribers by calling the callbacks they registered.
func (n *notifier) notify(resourceID string) { _ = "STUB: not implemented"; return }
