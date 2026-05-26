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

package ipset

import (
	"regexp"

	"k8s.io/utils/exec"
)

type SetType string

const (
	// The hash:net set type uses a hash to store different sized IP network addresses.
	// The lookup time grows linearly with the number of the different prefix values added to the set.
	HashNet    SetType = "hash:net"
	HashIP     SetType = "hash:ip"
	HashIPPort SetType = "hash:ip,port"
)

// memberPattern is used to match the members part of ipset list result.
var memberPattern = regexp.MustCompile("(?m)^(.*\n)*Members:\n")

type Interface interface {
	CreateIPSet(name string, setType SetType, isIPv6 bool) error

	DestroyIPSet(name string) error

	AddEntry(name string, entry string) error

	DelEntry(name string, entry string) error

	ListEntries(name string) ([]string, error)

	Save() ([]byte, error)
}

type Client struct {
	exec exec.Interface
}

var _ Interface = &Client{}

func NewClient() *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) DestroyIPSet(name string) error { _ = "STUB: not implemented"; return nil }

// CreateIPSet creates a new set, it will ignore error when the set already exists.
func (c *Client) CreateIPSet(name string, setType SetType, isIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G204 -- inputs are not controlled by users

// #nosec G204 -- inputs are not controlled by users

// AddEntry adds a new entry to the set, it will ignore error when the entry already exists.
func (c *Client) AddEntry(name string, entry string) error { _ = "STUB: not implemented"; return nil }

// DelEntry deletes the entry from the set, it will ignore error when the entry doesn't exist.
func (c *Client) DelEntry(name string, entry string) error { _ = "STUB: not implemented"; return nil }

// ListEntries lists all the entries of the set.
func (c *Client) ListEntries(name string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Save() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
