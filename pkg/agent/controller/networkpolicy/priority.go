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

package networkpolicy

import (
	"antrea.io/antrea/v2/pkg/agent/types"
)

const (
	zoneOffset                   = uint16(5)
	defaultTierPriority          = int32(250)
	baselinePolicyBottomPriority = uint16(10)
	baselinePolicyTopPriority    = uint16(180)
	policyBottomPriority         = uint16(100)
	policyTopPriority            = uint16(65000)
	priorityOffsetBaselineTier   = float64(10)
	tierOffsetBaselineTier       = uint16(0)
	priorityOffsetMultiTier      = float64(20)
	priorityOffsetDefaultTier    = float64(100)
	tierOffsetMultiTier          = uint16(200)
)

// priorityUpdate stores the original and updated ofPriority of a Priority.
type priorityUpdate struct {
	Original uint16
	Updated  uint16
}

// reassignCost stores the cost of reassigning registered Priorities, if all registered
// Priorities in the lowerBound-upperBound range were to be rearranged.
type reassignCost struct {
	lowerBound uint16
	upperBound uint16
	cost       int
}

// priorityUpdatesToOFUpdates converts a map of Priority and its ofPriority update to a map
// of ofPriority updates.
func priorityUpdatesToOFUpdates(allUpdates map[types.Priority]*priorityUpdate) map[uint16]uint16 {
	_ = "STUB: not implemented"
	return nil
}

// priorityAssigner is a struct that maintains the current boundaries of
// all ClusterNetworkPolicy categories/priorities and rule priorities, and knows
// how to re-assign priorities if certain section overflows.
type priorityAssigner struct {
	// priorityMap maintains the current mapping between a known Priority to OpenFlow priority.
	priorityMap map[types.Priority]uint16
	// ofPriorityMap maintains the current mapping of OpenFlow priorities in the table to Priorities.
	ofPriorityMap map[uint16]types.Priority
	// sortedPriorities maintains a list of sorted Priorities currently registered in the table.
	sortedPriorities types.ByPriority
	// isBaselineTier keeps track of if the priorityAssigner is responsible for handling the baseline Tier
	// table (which is shared with K8s NetworkPolicy default tables) or the Antrea Policy tables.
	isBaselineTier bool
	// policyBottomPriority keeps track of the lowest ofPriority allowed for flow creation in the table it manages.
	policyBottomPriority uint16
	// policyTopPriority keeps track of the highest ofPriority allowed for flow creation in the table it manages.
	policyTopPriority uint16
}

func newPriorityAssigner(isBaselineTier bool) *priorityAssigner {
	_ = "STUB: not implemented"
	return nil
}

// initialOFPriority is a heuristic function that will map types.Priority to a specific initial
// OpenFlow priority in a table. It is used to space out the priorities in the OVS table and provide an
// initial guess on the OpenFlow priority that can be assigned to the input Priority. If that OpenFlow
// priority is not available, or if the surrounding priorities are out of place, insertConsecutivePriorities()
// will then search for the appropriate OpenFlow priority to insert the input Priority.
// It computes the initial OpenFlow priority by offsetting the tier priority, policy priority and rule priority
// with pre-determined coefficients.
func (pa *priorityAssigner) initialOFPriority(p types.Priority) uint16 {
	_ = "STUB: not implemented"
	return 0
}

// Use uint32 to prevent arithmetic overflow during calculation

// Check if the calculated offset exceeds the available priority range

// Safe to cast offSet to uint16: the preceding if statement guarantees that
// offSet <= availableRange <= (policyTopPriority - policyBottomPriority) <= math.MaxUint16

// findGap calculates the gap between upperBound and lowerBound with uint16 underflow protection.
// Returns the gap and whether there is any space available.
func findGap(upperBound, lowerBound uint16) (gap uint16, hasSpace bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// updatePriorityAssignment updates all the local maps to correlate input ofPriority and Priority.
func (pa *priorityAssigner) updatePriorityAssignment(ofPriority uint16, p types.Priority) {
	_ = "STUB: not implemented"
	return
}

// idx is the insertion point for the newly registered Priority.

// Move elements starting from idx back one position to make room for the inserting Priority.

// findReassignBoundaries finds the range to reassign Priorities that minimizes the number of
// registered Priorities to be reassigned.
func (pa *priorityAssigner) findReassignBoundaries(lowerBound, upperBound uint16, numNewPriorities, gap int) (uint16, uint16, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil

	// To reach the target number of slots to be added into the gap, Priorities needs to be sifted upwards or
	// downwards (or both), and empty slots from lower and higher ofPriority space will be swapped into the gap.
	// costMap maintains the costs and reassign boundaries for each combination of lower empty slots used and
	// higher empty slots used, with the sum equals the target. For example, if the target is 2, the maps stores
	//  {0: cost of using 0 lower empty slots and 2 higher empty slots,
	//   1: cost of using 1 lower empty slot and 1 higher empty slot each,
	//   2: cost of using 2 lower empty slots and 0 higher empty slots}
}

// Search for empty slots below lowerBound, but don't go below policyBottomPriority

// When upperBound == lowerBound, use upperBound directly to create a valid single-slot range
// When upperBound > lowerBound, use upperBound - 1 to avoid overlapping with the gap

// Special case: when we're at the boundary itself, include it in the range

// Include the boundary slot

// Only create cost map entry if it results in a valid range

// visit costMap in the reverse direction

// only add to the costMap if the counterpart cost is available. i.e. if the target is 4, and cost for
// using 2 empty slots high is computed, it does not make sense to store this cost if there's no entry
// for cost that uses 2 empty slots low (indicating no 2 empty slots can be found starting from lowerBound).

// When upperBound == lowerBound, start from lowerBound itself, not lowerBound + 1

// Include the boundary slot in the reassignment range

// make sure that the reassign range adds up to the number of all Priorities to be registered.

// theoretically this should not happen since Priority overflow is checked earlier.

// reassignBoundaryPriorities reassigns Priorities from lowerBound / upperBound or both, to make room for
// new Priorities to be registered. It also records all the priority updates due to the reassignment in the
// map of updates, which is passed to it as parameter.
func (pa *priorityAssigner) reassignBoundaryPriorities(lowerBound, upperBound uint16, prioritiesToRegister types.ByPriority,
	updates map[types.Priority]*priorityUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

// No available slots when upperBound <= lowerBound

// siftedPrioritiesLow and siftedPrioritiesHigh keep track of Priorities that need to be reassigned,
// below the lowerBound and above the upperBound respectively. Note that we do not include the priority
// currently assigned to policyBottomPriority since it cannot be sifted down further, and vice versa
// for policyTopPriority.

// record the ofPriorities of the reassigned Priorities before the reassignment.

// if exists (the Priority has already been reassigned in a previous step), the original
// ofPriority of that Priority would have been recorded.

// assign ofPriorities by the order of siftedPrioritiesLow, prioritiesToRegister and siftedPrioritiesHigh.

// Protect against overflow when low + i > 65535

// record the ofPriorities of the reassigned Priorities after the reassignment.

// getOFPriority returns if the Priority is registered with the priorityAssigner,
// and retrieves the corresponding ofPriority.
func (pa *priorityAssigner) getOFPriority(p types.Priority) (uint16, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// registerPriorities registers a list of types.Priority with the priorityAssigner. It allocates ofPriorities for
// input priorities that are not yet registered. It also returns the ofPriority updates if there are reassignments,
// as well as a revert function that can undo the registration if any error occurred in data plane.
// Note that this function modifies the priorities slice in the parameter, as it only keeps the Priorities which
// this priorityAssigner has not yet registered. Input priorities are not assumed to be unique or consecutive.
func (pa *priorityAssigner) registerPriorities(priorities []types.Priority) (map[uint16]uint16, func(), error) {
	_ = "STUB: not implemented"
	// create a zero-length slice with the same underlying array to save memory usage.
	return nil, nil, nil
}

// Check for overflow before casting to uint16

// break the list of Priority into lists of consecutive Priority.

// registerConsecutivePriorities registers lists of consecutive Priorities with the priorityAssigner.
func (pa *priorityAssigner) registerConsecutivePriorities(consecutivePriorities [][]types.Priority) (map[uint16]uint16, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// in case of error, all new Priorities need to be unregistered.

// all reassigned Priorities need to be assigned back to the original ofPriorities.

// if failure occurred at any point, revert Priorities registered so far.

// insertConsecutivePriorities inserts a list of consecutive Priorities into the ofPriority space.
// It first identifies the lower and upper bound for insertion, by obtaining the ofPriorities of
// registered Priorities surrounding (immediately lower and higher than) the inserting Priorities.
// It then decides the range to register new Priorities, and reassign existing ones if necessary.
func (pa *priorityAssigner) insertConsecutivePriorities(priorities types.ByPriority, updates map[types.Priority]*priorityUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

// get the index for inserting the lowest Priority into the registered Priorities.

// set lowerBound to the ofPriority of the registered Priority that is immediately lower than the inserting Priorities.

// set upperBound to the ofPriority of the registered Priority that is immediately higher than the inserting Priorities.

// not enough space currently to insert Priorities.

// ofPriorities provided by the heuristic function are good.

// ofPriorities returned by the heuristic function are out of place/overlap with existing Priorities.
// If the Priorities to be registered overlap with lower Priorities/are lower than the lower Priorities,
// and the gap between lowerBound and upperBound for insertion is large, then we insert these Priorities
// above the lowerBound, offsetted by a constant zoneOffset, and vice versa.
// 5 is chosen as the zoneOffset here since it gives some buffer in case Priorities are again created
// in between those zones, while in the meantime keeps priority assignments compact.

// Protect against overflow: lowerBound + zoneOffset + 1

// Protect against underflow: upperBound - zoneOffset - uint16(len(priorities))

// when the window between upper/lowerBound is small, simply put the Priorities in the middle of the window.

// gap is guaranteed >= numPriorities by the condition check above
// So gap - uint16(numPriorities) will not underflow

// Protect against overflow: insertionPointLow + uint16(i)

// release removes the priority that currently corresponds to the input OFPriority from the known priorities.
func (pa *priorityAssigner) release(ofPriority uint16) { _ = "STUB: not implemented"; return }

// deletePriorityMapping removes the Priority <-> ofPriority mapping from the input
func (pa *priorityAssigner) deletePriorityMapping(ofPriority uint16, priority types.Priority) {
	_ = "STUB: not implemented"
	return
}

// unregisterPriority unregisters the Priority from the known Priorities.
func (pa *priorityAssigner) unregisterPriority(priority types.Priority) {
	_ = "STUB: not implemented"
	return
}
