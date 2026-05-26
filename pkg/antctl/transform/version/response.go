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

package version

import (
	"io"
)

type Response struct {
	AgentVersion          string `json:"agentVersion,omitempty"`
	ControllerVersion     string `json:"controllerVersion,omitempty"`
	FlowAggregatorVersion string `json:"flowAggregatorVersion,omitempty"`
	AntctlVersion         string `json:"antctlVersion,omitempty"`
}

// AgentVersion is the AddonTransform for the version command. This function
// will try to parse the response as a AgentVersionResponse and then populate
// it with the version of antctl to a transformedVersionResponse object.
func AgentTransform(reader io.Reader, _ bool, _ map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ControllerTransform(reader io.Reader, _ bool, _ map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FlowAggregatorTransform is the AddonTransform for the flow aggregator version command.
// This function will try to parse the response as a FlowAggregatorVersionResponse and
// then populate it with the version of antctl to a transformedVersionResponse object.
func FlowAggregatorTransform(reader io.Reader, _ bool, _ map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
