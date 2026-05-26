// Copyright 2023 Antrea Authors
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

package memberlist

import (
	"net/http"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/agent/apis"
	"antrea.io/antrea/v2/pkg/agent/querier"
)

func generateResponse(node *v1.Node, aliveNodes sets.Set[string]) apis.MemberlistResponse {
	_ = "STUB: not implemented"
	return *new(apis.MemberlistResponse)
}

// HandleFunc returns the function which can handle queries issued by the memberlist command.
func HandleFunc(aq querier.AgentQuerier) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// The error message must match the "FOO is not enabled" pattern to pass antctl e2e tests.
