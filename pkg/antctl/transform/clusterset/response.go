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

package clusterset

import (
	mcv1alpha2 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha2"
	"antrea.io/antrea/v2/pkg/antctl/transform/common"
)

type Response struct {
	ClusterID    string `json:"clusterID" yaml:"clusterID"`
	Namespace    string `json:"namespace" yaml:"namespace"`
	ClusterSetID string `json:"clusterSetID" yaml:"clusterSetID"`
	Type         string `json:"type" yaml:"type"`
	Status       string `json:"status" yaml:"status"`
	Reason       string `json:"reason" yaml:"reason"`
}

func Transform(r interface{}, single bool) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listTransform(l interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// When the ClusterSet has no status, we should print it with empty status.

func objectTransform(clusterSet mcv1alpha2.ClusterSet, status mcv1alpha2.ClusterStatus,
	condition mcv1alpha2.ClusterCondition) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ common.TableOutput = new(Response)

func (r Response) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r Response) GetTableRow(maxColumnLength int) []string { _ = "STUB: not implemented"; return nil }

func (r Response) SortRows() bool { _ = "STUB: not implemented"; return false }
