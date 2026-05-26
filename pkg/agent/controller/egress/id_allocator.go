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

package egress

import (
	"container/list"
	"sync"
)

type idAllocator struct {
	sync.Mutex
	maxID        uint32
	nextID       uint32
	availableIDs *list.List
}

func (a *idAllocator) allocate() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func (a *idAllocator) release(id uint32) error { _ = "STUB: not implemented"; return nil }

func newIDAllocator(minID, maxID uint32) *idAllocator { _ = "STUB: not implemented"; return nil }
