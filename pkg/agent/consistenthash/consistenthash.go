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

// The idea is borrowed from https://github.com/golang/groupcache/blob/master/consistenthash/consistenthash.go with the following modifications:
// - Store the replicas in a 2-3-4 btree so that we can add/remove keys in O(log N) without rebuilding the whole cache.
// - Add a function GetWithFilters() to allow filtering keys with desired filters.

// Package consistenthash provides an implementation of a ring hash.
package consistenthash

import (
	"github.com/google/btree"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     map[string]struct{}
	tree     *btree.BTree
}

type replica struct {
	key  string
	hash uint32
}

func (v *replica) Less(than btree.Item) bool { _ = "STUB: not implemented"; return false }

var _ btree.Item = (*replica)(nil)

func New(replicas int, fn Hash) *Map { _ = "STUB: not implemented"; return nil }

// IsEmpty returns true if there are no items available.
func (m *Map) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Add adds some keys to the hash.
func (m *Map) Add(keys ...string) { _ = "STUB: not implemented"; return }

// Remove removes keys from existing hash ring.
func (m *Map) Remove(keys ...string) { _ = "STUB: not implemented"; return }

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string { _ = "STUB: not implemented"; return "" }

// GetWithFilters gets the closest item in the hash to which passes all filters.
func (m *Map) GetWithFilters(key string, filters ...func(string) bool) string {
	_ = "STUB: not implemented"
	return ""
}

// all keys visited

// stop iterating

// search in [pivot, last]

// search in [first, pivot)

// no key passes all filters
