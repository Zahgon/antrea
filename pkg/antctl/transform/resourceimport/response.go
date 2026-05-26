// Copyright 2022 Antrea Authors
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

package resourceimport

import (
	"antrea.io/antrea/v2/pkg/antctl/transform/common"
)

type Response struct {
	Namespace string `json:"namespace" yaml:"namespace"`
	Name      string `json:"name" yaml:"name"`
	Kind      string `json:"kind" yaml:"kind"`
}

func Transform(r interface{}, single bool) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listTransform(l interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func objectTransform(o interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ common.TableOutput = new(Response)

func (r Response) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r Response) GetTableRow(maxColumnLength int) []string { _ = "STUB: not implemented"; return nil }

func (r Response) SortRows() bool { _ = "STUB: not implemented"; return false }
