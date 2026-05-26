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

package openflow

import (
	"net"

	"antrea.io/libOpenflow/openflow15"
	"antrea.io/ofnet/ofctrl"
)

var (
	MaxBucketsPerMessage = 700
)

type ofGroup struct {
	ofctrl *ofctrl.Group
	bridge *OFBridge
}

// Reset updates ofctrl.Group.Switch with the updated ofSwitch.
func (g *ofGroup) Reset() { _ = "STUB: not implemented"; return }

func (g *ofGroup) Add() error { _ = "STUB: not implemented"; return nil }

func (g *ofGroup) Modify() error { _ = "STUB: not implemented"; return nil }

func (g *ofGroup) Delete() error { _ = "STUB: not implemented"; return nil }

func (g *ofGroup) Type() EntryType { _ = "STUB: not implemented"; return *new(EntryType) }

func (g *ofGroup) Bucket() BucketBuilder { _ = "STUB: not implemented"; return *new(BucketBuilder) }

func (g *ofGroup) GetBundleMessages(entryOper OFOperation) ([]ofctrl.OpenFlowModMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the operation is to delete the group, empty the slice storing buckets since the number of buckets could
// be greater than MaxBucketsPerMessage.

// Get the range of buckets to generate a temp group.

// For the message which is not the first, insert_buckets is used to add buckets to the group on OVS.

// There is no need to generate an insert_buckets message without bucket.

// Generate a temp group to get an OVS message. Note that, the original group should not be modified since it is
// also stored in group cache, and the group cache is used when replaying groups.

func (g *ofGroup) ResetBuckets() Group { _ = "STUB: not implemented"; return *new(Group) }

func (g *ofGroup) GetID() GroupIDType { _ = "STUB: not implemented"; return *new(GroupIDType) }

type bucketBuilder struct {
	group  *ofGroup
	bucket *openflow15.Bucket
}

// LoadXXReg makes the learned flow to load data to xxreg[regID] with specific range.
func (b *bucketBuilder) LoadXXReg(regID int, data []byte) BucketBuilder {
	_ = "STUB: not implemented"
	return *new(BucketBuilder)
}

func (b *bucketBuilder) LoadToRegField(field *RegField, data uint32) BucketBuilder {
	_ = "STUB: not implemented"
	return *new(BucketBuilder)
}

func (b *bucketBuilder) LoadRegMark(mark *RegMark) BucketBuilder {
	_ = "STUB: not implemented"
	return *new(BucketBuilder)
}

// ResubmitToTable is an action to resubmit packet to the specified table when the bucket is selected.
func (b *bucketBuilder) ResubmitToTable(tableID uint8) BucketBuilder {
	_ = "STUB: not implemented"
	return *new(BucketBuilder)
}

// SetTunnelDst is an action to set tunnel destination address when the bucket is selected.
func (b *bucketBuilder) SetTunnelDst(addr net.IP) BucketBuilder {
	_ = "STUB: not implemented"
	return *new(BucketBuilder)
}

// Weight sets the weight of a bucket.
func (b *bucketBuilder) Weight(val uint16) BucketBuilder {
	_ = "STUB: not implemented"
	return *new(BucketBuilder)
}

func (b *bucketBuilder) Done() Group { _ = "STUB: not implemented"; return *new(Group) }
