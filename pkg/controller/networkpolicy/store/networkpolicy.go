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
	AppliedToGroupIndex = "appliedToGroup"
	AddressGroupIndex   = "addressGroup"
)

// networkPolicyEvent implements storage.InternalEvent.
type networkPolicyEvent struct {
	// The current version of the stored NetworkPolicy.
	CurrPolicy *types.NetworkPolicy
	// The previous version of the stored NetworkPolicy.
	PrevPolicy *types.NetworkPolicy
	// The current version of the transferred NetworkPolicy, which will be used in Added and Modified events.
	CurrObject *controlplane.NetworkPolicy
	// The previous version of the transferred NetworkPolicy, which will be used in Deleted events.
	// Note that only metadata will be set in Deleted events for efficiency.
	PrevObject *controlplane.NetworkPolicy
	// The key of this NetworkPolicy.
	Key             string
	ResourceVersion uint64
}

// ToWatchEvent converts the networkPolicyEvent to *watch.Event based on the provided Selectors. It has the following features:
// 1. Added event will be generated if the Selectors was not interested in the object but is now.
// 2. Modified event will be generated if the Selectors was and is interested in the object.
// 3. Deleted event will be generated if the Selectors was interested in the object but is not now.
func (event *networkPolicyEvent) ToWatchEvent(selectors *storage.Selectors, isInitEvent bool) *watch.Event {
	_ = "STUB: not implemented"
	return nil
}

func (event *networkPolicyEvent) GetResourceVersion() uint64 { _ = "STUB: not implemented"; return 0 }

var _ storage.GenEventFunc = genNetworkPolicyEvent

// genNetworkPolicyEvent generates InternalEvent from the given versions of a NetworkPolicy.
// It converts the stored NetworkPolicy to its message form.
func genNetworkPolicyEvent(key string, prevObj, currObj interface{}, rv uint64) (storage.InternalEvent, error) {
	_ = "STUB: not implemented"
	return *new(storage.InternalEvent), nil
}

// ToNetworkPolicyMsg converts the stored NetworkPolicy to its message form.
// If includeBody is true, Rules and AppliedToGroups will be copied.
func ToNetworkPolicyMsg(in *types.NetworkPolicy, out *controlplane.NetworkPolicy, includeBody bool) {
	_ = "STUB: not implemented"
	return
}

// Since stored objects are immutable, we just reference the fields here.

// AppliedToGroups at policy level only need to be populated to controlplane msg if
// appliedTo is not set per rule. Otherwise, the agent will read appliedTo from each
// rule, and in that case, the policy level appliedTo only contains span information,
// which the Antrea agent does not care about.

// NetworkPolicyKeyFunc knows how to get the key of a NetworkPolicy.
func NetworkPolicyKeyFunc(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewNetworkPolicyStore creates a store of NetworkPolicy.
func NewNetworkPolicyStore() storage.Interface {
	_ = "STUB: not implemented"
	// Build indices with the appliedToGroups and the addressGroups so that
	// it's efficient to get network policies that have references to specified
	// appliedToGroups or addressGroups.
	return *new(storage.Interface)
}
