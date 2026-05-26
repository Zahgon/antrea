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

package addressgroup

import (
	"io"

	"antrea.io/antrea/v2/pkg/antctl/transform/common"
)

type Response struct {
	Name  string               `json:"name" yaml:"name"`
	Pods  []common.GroupMember `json:"pods,omitempty"`
	Nodes []common.GroupMember `json:"nodes,omitempty"`
}

func listTransform(l interface{}, opts map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func objectTransform(o interface{}, _ map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Transform(reader io.Reader, single bool, opts map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ common.TableOutput = new(Response)

func (r Response) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r Response) GetPodIPs(maxColumnLength int) string { _ = "STUB: not implemented"; return "" }

func (r Response) GetNodeIPs(maxColumnLength int) string { _ = "STUB: not implemented"; return "" }

func (r Response) GetTableRow(maxColumnLength int) []string { _ = "STUB: not implemented"; return nil }

func (r Response) SortRows() bool { _ = "STUB: not implemented"; return false }
