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

package podinterface

import (
	"net"
	"net/http"

	"antrea.io/antrea/v2/pkg/agent/apis"
	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	"antrea.io/antrea/v2/pkg/agent/querier"
)

func generateResponse(i *interfacestore.InterfaceConfig) apis.PodInterfaceResponse {
	_ = "STUB: not implemented"
	return *new(apis.PodInterfaceResponse)
}

func getPodIPs(ips []net.IP) []string { _ = "STUB: not implemented"; return nil }

// HandleFunc returns the function which can handle queries issued by the pod-interface command.
func HandleFunc(aq querier.AgentQuerier) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
