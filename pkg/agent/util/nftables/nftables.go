// Copyright 2025 Antrea Authors
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

package nftables

import (
	"iter"

	"sigs.k8s.io/knftables"
)

const table = "antrea"

type Client struct {
	IPv4 knftables.Interface
	IPv6 knftables.Interface
}

func (c *Client) All() iter.Seq2[knftables.Family, knftables.Interface] {
	_ = "STUB: not implemented"
	return nil
}

func New(enableIPv4, enableIPv6 bool) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

func newNFTables(ipFamily knftables.Family) (knftables.Interface, error) {
	_ = "STUB: not implemented"
	// knftables.New validates:
	//  - nft binary is available
	//  - sufficient permissions
	//  - kernel version compatibility
	//  - "nft destroy" support
	return *new(knftables.Interface), nil
}

// Verify nft_flow_offload support with a dry-run.
