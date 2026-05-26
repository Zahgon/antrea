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

package featuregates

import (
	"net/http"

	clientset "k8s.io/client-go/kubernetes"

	"antrea.io/antrea/v2/pkg/apiserver/apis"
)

type Config struct {
	// FeatureGates is a map of feature names to bools that enable or disable experimental features.
	FeatureGates map[string]bool `yaml:"featureGates,omitempty"`
}

const (
	AgentMode            = "agent"
	AgentWindowsMode     = "agent-windows"
	ControllerMode       = "controller"
	agentConfigName      = "antrea-agent.conf"
	controllerConfigName = "antrea-controller.conf"
)

// HandleFunc returns the function which can handle queries issued by 'antctl get featuregates' command.
// The handler function populates Antrea featuregates information to the response.
func HandleFunc(k8sclient clientset.Interface) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func getFeatureGatesResponse(cfg *Config, component string) []apis.FeatureGateResponse {
	_ = "STUB: not implemented"
	return nil
}
