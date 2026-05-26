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

package openflow

import (
	"sync"

	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

type GroupAllocator interface {
	Allocate() binding.GroupIDType
	Next() binding.GroupIDType
	Release(id binding.GroupIDType)
}

type groupAllocator struct {
	// mu is a lock for the groupAllocator.
	mu sync.Mutex

	groupIDCounter binding.GroupIDType
	recycled       []binding.GroupIDType
}

// Allocate allocates a new group ID. It allocates id from the "recycled" slices first, then increases the groupIDCounter if no
// recycled ids exist.
func (a *groupAllocator) Allocate() binding.GroupIDType {
	_ = "STUB: not implemented"
	return *new(binding.GroupIDType)
}

// Next is a readonly method which returns the next available group ID. It's useful in tests to predict the group ID.
func (a *groupAllocator) Next() binding.GroupIDType {
	_ = "STUB: not implemented"
	return *new(binding.GroupIDType)
}

func (a *groupAllocator) Release(id binding.GroupIDType) { _ = "STUB: not implemented"; return }

func NewGroupAllocator() GroupAllocator { _ = "STUB: not implemented"; return *new(GroupAllocator) }
