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

package querier

import (
	v1 "k8s.io/api/core/v1"

	"antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/querier"
)

const (
	serviceName = "antrea"
)

var _ ControllerQuerier = new(controllerQuerier)

type ControllerQuerier interface {
	GetControllerInfo(controllerInfo *v1beta1.AntreaControllerInfo, partial bool)
}

type controllerQuerier struct {
	networkPolicyInfoQuerier querier.ControllerNetworkPolicyInfoQuerier
	apiPort                  int
}

func NewControllerQuerier(networkPolicyInfoQuerier querier.ControllerNetworkPolicyInfoQuerier,
	apiPort int) *controllerQuerier {
	_ = "STUB: not implemented"
	return nil
}

// getNetworkPolicyInfoQuerier gets current network policy info querier.
func (cq controllerQuerier) getNetworkPolicyInfoQuerier() querier.ControllerNetworkPolicyInfoQuerier {
	_ = "STUB: not implemented"
	return *new(querier.ControllerNetworkPolicyInfoQuerier)
}

// getService gets current service.
func (cq controllerQuerier) getService() v1.ObjectReference {
	_ = "STUB: not implemented"
	return *new(v1.ObjectReference)
}

// getNetworkPolicyControllerInfo gets current network policy controller info
// including: number of network policies, address groups and applied to groups.
func (cq controllerQuerier) getNetworkPolicyControllerInfo() v1beta1.NetworkPolicyControllerInfo {
	_ = "STUB: not implemented"
	return *new(v1beta1.NetworkPolicyControllerInfo)
}

func (cq controllerQuerier) getControllerConditions() []v1beta1.ControllerCondition {
	_ = "STUB: not implemented"
	return nil
}

// GetControllerInfo gets current info of controller.
func (cq controllerQuerier) GetControllerInfo(controllerInfo *v1beta1.AntreaControllerInfo, partial bool) {
	_ = "STUB: not implemented"
	return
}
