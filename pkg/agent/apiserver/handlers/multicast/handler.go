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

package multicast

import (
	"net/http"

	"antrea.io/antrea/v2/pkg/agent/apis"
	"antrea.io/antrea/v2/pkg/agent/multicast"
	"antrea.io/antrea/v2/pkg/querier"
)

func generateResponse(podName string, podNamespace string, trafficStats *multicast.PodTrafficStats) apis.MulticastResponse {
	_ = "STUB: not implemented"
	return *new(apis.MulticastResponse)
}

// HandleFunc returns the function which can handle queries issued by 'antctl get podmulticaststats' command.
// It will return Pod multicast traffic statistics for the local Node.
func HandleFunc(mq querier.AgentMulticastInfoQuerier) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
