// Copyright 2017 DigitalOcean.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This whole file is from
// https://github.com/digitalocean/go-openvswitch/blob/master/ovs/portrange.go
package networkpolicy

import (
	"errors"
)

var (
	// ErrInvalidPortRange is returned when there's a port range that invalid.
	ErrInvalidPortRange = errors.New("invalid port range")
)

// An PortRange represents a range of ports expressed in 16 bit integers.  The start and
// end values of this range are inclusive.
type PortRange struct {
	Start uint16
	End   uint16
}

// A BitRange is a representation of a range of values from base value with a bitmask
// applied.
type BitRange struct {
	Value uint16
	Mask  uint16
}

// BitwiseMatch returns an array of BitRanges that represent the range of integers
// in the PortRange.
func (r *PortRange) BitwiseMatch() ([]BitRange, error) { _ = "STUB: not implemented"; return nil, nil }

// Find the largest window we can get on a binary boundary

// Decrement our mask until we fit inside the range we want from a binary boundary.

// The range we picked out was from the middle of our set, so we'll need to recurse on
// the remaining values for anything less than or greater than the current
// range.

// We append our current range here, so we're ordered properly.

func getMask(bitLength uint) uint16 {
	_ = "STUB: not implemented"
	// All 1s for everything that doesn't change in the range
	return 0
}

func getRange(end uint16, bitLength uint) (rangeStart uint16, rangeEnd uint16) {
	_ = "STUB: not implemented"
	// Represents the upper bound of our range window (all 1s to binary boundary)
	return 0, 0
}

// Zero out our mask, so we start at a binary boundary.

// Simply add the mask so we end at a binary boundary.
