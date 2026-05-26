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
	"antrea.io/antrea/v2/pkg/apiserver/storage"
)

const (
	ServiceIndex      = "service"
	ChildGroupIndex   = "childGroup"
	IPBlockGroupIndex = "hasIPBlocks"
	HasIPBlocks       = "true"
	NodeSelectorIndex = "nodeSelectorIndex"
	HasNodeSelector   = "true"
)

// GroupKeyFunc knows how to get the key of a Group.
func GroupKeyFunc(obj interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

// NewGroupStore creates a store of Group.
func NewGroupStore() storage.Interface { _ = "STUB: not implemented"; return *new(storage.Interface) }

// genEventFunc is set to nil, thus watchers of this store will not be created.
