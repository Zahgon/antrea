// Copyright 2025 Antrea Authors
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

package fqdncache

import (
	"net/http"
	"net/url"

	"antrea.io/antrea/v2/pkg/querier"
)

func HandleFunc(npq querier.AgentNetworkPolicyInfoQuerier) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func newFilterFromURLQuery(query url.Values) (*querier.FQDNCacheFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Replace "." as a regex literal, since it's recogized as a separator in FQDN.

// Replace "*" with ".*".

// Anchor the regex match expression.
