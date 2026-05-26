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

package testing

import (
	"testing"

	"antrea.io/antrea/v2/pkg/agent/nodeportlocal/types"
)

type ExpectedNPLAnnotations struct {
	nplStartPort int
	nplEndPort   int
	annotations  []types.NPLAnnotation
}

func NewExpectedNPLAnnotations(nplStartPort, nplEndPort int) *ExpectedNPLAnnotations {
	_ = "STUB: not implemented"
	return nil
}

func (a *ExpectedNPLAnnotations) find(podPort int, protocol string, ipFamily types.IPFamilyType) *types.NPLAnnotation {
	_ = "STUB: not implemented"
	return nil
}

func (a *ExpectedNPLAnnotations) Add(ipFamily types.IPFamilyType, nodeIP *string, nodePort *int, podPort int, protocol string) *ExpectedNPLAnnotations {
	_ = "STUB: not implemented"
	return nil
}

func (a *ExpectedNPLAnnotations) Check(t *testing.T, nplValue []types.NPLAnnotation) {
	_ = "STUB: not implemented"
	return
}

// Count returns the number of expected annotations, which corresponds to the number of Add calls
// made so far.
func (a *ExpectedNPLAnnotations) Count() int { _ = "STUB: not implemented"; return 0 }
