// Copyright 2024 Antrea Authors
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
	"io"

	"antrea.io/antrea/v2/pkg/antctl/transform/common"
	cpv1beta "antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	"antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

const sortByEffectivePriority = "effectivePriority"

type Response struct {
	*cpv1beta.NetworkPolicy
}

// Compute a tierPriority value in between the application tier and the baseline tier,
// which can be used to sort all policies by tier.
var effectiveTierPriorityK8sNP = (v1beta1.DefaultTierPriority + v1beta1.BaselineTierPriority) / 2

type NPSorter struct {
	networkPolicies []cpv1beta.NetworkPolicy
	sortBy          string
}

func objectTransform(o interface{}, _ map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listTransform(l interface{}, opts map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// To check for any special sort cases.

func Transform(reader io.Reader, single bool, opts map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nps *NPSorter) Len() int      { _ = "STUB: not implemented"; return 0 }
func (nps *NPSorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (nps *NPSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Do not need a tie-breaker here since NetworkPolicy names are set as UID
// of the source policy and will be unique.

func priorityToString(p interface{}) string { _ = "STUB: not implemented"; return "" }

var _ common.TableOutput = new(Response)

func (r Response) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r Response) GetTableRow(maxColumnLength int) []string { _ = "STUB: not implemented"; return nil }

func (r Response) SortRows() bool {
	_ = "STUB: not implemented"

	// EvaluationResponse stores the response from NetworkPolicyEvaluation command,
	// and implements TableOutput.
	return false
}

type EvaluationResponse struct {
	*cpv1beta.NetworkPolicyEvaluation
}

func EvaluationTransform(reader io.Reader, _ bool, _ map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ common.TableOutput = new(EvaluationResponse)

func (r EvaluationResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r EvaluationResponse) GetTableRow(_ int) []string { _ = "STUB: not implemented"; return nil }

// Responses from endpoint query with original rules will always have
// valid action fields, except for the synthetic isolation rules,
// identified by a MaxInt32 rule index. "Isolate" corresponds to
// a drop action because of the default isolation model of K8s NPs.

// Should not be possible.

func (r EvaluationResponse) SortRows() bool { _ = "STUB: not implemented"; return false }
