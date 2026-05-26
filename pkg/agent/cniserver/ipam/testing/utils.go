// Copyright 2019 Antrea Authors
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

package testing

import (
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
)

func GenerateIPAMResult(ipConfig []string, routeConfig []string, dnsConfig []string) *current.Result {
	_ = "STUB: not implemented"
	return nil
}

func parseRoute(routeConfig []string) []*types.Route { _ = "STUB: not implemented"; return nil }

func parseIPs(ips []string) []*current.IPConfig { _ = "STUB: not implemented"; return nil }

func parseIPConfig(ipAddress string, gw string, version string) *current.IPConfig {
	_ = "STUB: not implemented"
	return nil
}
