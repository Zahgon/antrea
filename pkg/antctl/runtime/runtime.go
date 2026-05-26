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

package runtime

import (
	"os"
	"strings"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"

	"antrea.io/antrea/v2/pkg/apis"
	"antrea.io/antrea/v2/pkg/util/runtime"
)

const (
	ModeController     string = "controller"
	ModeAgent          string = "agent"
	ModeFlowAggregator string = "flowaggregator"
)

var (
	// Mode tells which mode antctl is running against.
	Mode  string
	InPod bool
)

func ResolveKubeconfig(path string) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G703: Path provided by local user and consumed by CLI; no privilege boundary crossed.

func init() {
	podName, found := os.LookupEnv("POD_NAME")
	InPod = found && (strings.HasPrefix(podName, "antrea-agent") || strings.HasPrefix(podName, "antrea-controller") ||
		strings.HasPrefix(podName, "flow-aggregator"))

	if runtime.IsWindowsPlatform() && !InPod {
		if _, err := os.Stat(apis.APIServerLoopbackTokenPath); err == nil {
			InPod = true
			Mode = ModeAgent
			return
		}
	}

	if strings.HasPrefix(podName, "antrea-agent") {
		Mode = ModeAgent
	} else if strings.HasPrefix(podName, "flow-aggregator") {
		Mode = ModeFlowAggregator
	} else {
		Mode = ModeController
	}
}
