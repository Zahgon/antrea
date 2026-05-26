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
	"k8s.io/apimachinery/pkg/runtime"
)

// parsePeer parses Namespace/Pod name, empty string is returned if the argument is not of a
// valid Namespace/Pod reference (missing pod name or invalid format). Namespace will be set
// as default if missing, string without separator will be considered as pod name.
func parsePeer(str string) (string, string) { _ = "STUB: not implemented"; return "", "" }

// NewNetworkPolicyEvaluation creates a new NetworkPolicyEvaluation resource
// request from the command-line arguments provided to antctl.
func NewNetworkPolicyEvaluation(args map[string]string) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}
