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

package controllerinfo

import (
	"io"

	corev1 "k8s.io/api/core/v1"

	"antrea.io/antrea/v2/pkg/antctl/transform/common"
	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

// Response includes all data fields of clusterinfo.AntreaControllerInfo, but
// removes the resource meta fields.
type Response struct {
	Version                     string                                 `json:"version,omitempty"`                     // Antrea binary version
	PodRef                      corev1.ObjectReference                 `json:"podRef,omitempty"`                      // The Pod that Antrea Controller is running in
	NodeRef                     corev1.ObjectReference                 `json:"nodeRef,omitempty"`                     // The Node that Antrea Controller is running in
	ServiceRef                  corev1.ObjectReference                 `json:"serviceRef,omitempty"`                  // Antrea Controller Service
	NetworkPolicyControllerInfo crdv1beta1.NetworkPolicyControllerInfo `json:"networkPolicyControllerInfo,omitempty"` // Antrea Controller NetworkPolicy information
	ConnectedAgentNum           int32                                  `json:"connectedAgentNum,omitempty"`           // Number of agents which are connected to this controller
	ControllerConditions        []crdv1beta1.ControllerCondition       `json:"controllerConditions,omitempty"`        // Controller condition contains types like ControllerHealthy
}

func Transform(reader io.Reader, _ bool, _ map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ common.TableOutput = new(Response)

func (r Response) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r Response) GetControllerConditionStr() string { _ = "STUB: not implemented"; return "" }

func (r Response) GetTableRow(maxColumnLength int) []string { _ = "STUB: not implemented"; return nil }

func (r Response) SortRows() bool { _ = "STUB: not implemented"; return false }
