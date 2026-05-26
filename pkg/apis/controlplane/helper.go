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

package controlplane

func (r *NetworkPolicyReference) ToString() string { _ = "STUB: not implemented"; return "" }

func (r *GroupReference) ToGroupName() string { _ = "STUB: not implemented"; return "" }

// ToTypedString returns the Group or ClusterGroup namespaced name as a string along with its type.
// Typed strings are typically used in log messages.
func (r *GroupReference) ToTypedString() string { _ = "STUB: not implemented"; return "" }

func IsSourceAntreaNativePolicy(npRef *NetworkPolicyReference) bool {
	_ = "STUB: not implemented"
	return false
}
