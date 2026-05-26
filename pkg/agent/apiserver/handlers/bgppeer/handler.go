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

package bgppeer

import (
	"net/http"

	"antrea.io/antrea/v2/pkg/querier"
)

// HandleFunc returns the function which can handle queries issued by the bgppeers command.
func HandleFunc(bq querier.AgentBGPPolicyInfoQuerier) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// The error message must match the "FOO is not enabled" pattern to pass antctl e2e tests.

// make sure that we provide a stable order for the API response

// IPv4 routes first, then IPv6 routes
