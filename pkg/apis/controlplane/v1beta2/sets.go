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

package v1beta2

import (
	"k8s.io/apimachinery/pkg/util/sets"
)

// groupMemberKey is used to uniquely identify GroupMember.
type groupMemberKey string

// GroupMemberSet is a set of GroupMembers.
// +k8s:openapi-gen=false
// +k8s:deepcopy-gen=false
type GroupMemberSet map[groupMemberKey]*GroupMember

// normalizeGroupMember calculates the groupMemberKey of the provided
// GroupMember based on the Pod/ExternalEntity/Service's namespaced name and IPs.
// For GroupMembers in appliedToGroups, the IPs are not set, so the
// generated key does not contain IP information.
func normalizeGroupMember(member *GroupMember) groupMemberKey {
	_ = "STUB: not implemented"
	// "/" is illegal in Namespace and name so is safe as the delimiter.
	return *new(groupMemberKey)
}

// NewGroupMemberSet builds a GroupMemberSet from a list of GroupMember.
func NewGroupMemberSet(items ...*GroupMember) GroupMemberSet {
	_ = "STUB: not implemented"
	return *new(GroupMemberSet)
}

// Insert adds items to the set.
func (s GroupMemberSet) Insert(items ...*GroupMember) { _ = "STUB: not implemented"; return }

// Delete removes all items from the set.
func (s GroupMemberSet) Delete(items ...*GroupMember) { _ = "STUB: not implemented"; return }

// Has returns true if and only if item is contained in the set.
func (s GroupMemberSet) Has(item *GroupMember) bool { _ = "STUB: not implemented"; return false }

// Difference returns a set of GroupMembers that are not in o.
func (s GroupMemberSet) Difference(o GroupMemberSet) GroupMemberSet {
	_ = "STUB: not implemented"
	return *new(GroupMemberSet)
}

// IPDifference returns a String set of GroupMember IPs that are not in o.
func (s GroupMemberSet) IPDifference(o GroupMemberSet) sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// Union returns a new set which includes items in either m or o.
func (s GroupMemberSet) Union(o GroupMemberSet) GroupMemberSet {
	_ = "STUB: not implemented"
	return *new(GroupMemberSet)
}

// Merge merges the other set into the set.
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Merge(s2) = {a1, a2, a3, a4, a5}
// s1 = {a1, a2, a3, a4, a5}
//
// It should be used instead of s1.Union(s2) when constructing a new set is not required.
func (s GroupMemberSet) Merge(o GroupMemberSet) GroupMemberSet {
	_ = "STUB: not implemented"
	return *new(GroupMemberSet)
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s GroupMemberSet) IsSuperset(o GroupMemberSet) bool { _ = "STUB: not implemented"; return false }

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s GroupMemberSet) Equal(o GroupMemberSet) bool { _ = "STUB: not implemented"; return false }

// Items returns the slice with contents in random order.
func (s GroupMemberSet) Items() []*GroupMember { _ = "STUB: not implemented"; return nil }
