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

package externalnode

import (
	"github.com/vishvananda/netlink"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/util"
)

var (
	linkByName             = netlink.LinkByName
	linkSetMTU             = netlink.LinkSetMTU
	linkSetUp              = netlink.LinkSetUp
	removeLinkIPs          = util.RemoveLinkIPs
	removeLinkRoutes       = util.RemoveLinkRoutes
	configureLinkAddresses = util.ConfigureLinkAddresses
	configureLinkRoutes    = util.ConfigureLinkRoutes
)

func (c *ExternalNodeController) moveIFConfigurations(adapterConfig *config.AdapterNetConfig, src string, dst string) error {
	_ = "STUB: not implemented"
	return nil
}

// Configure the source interface's IPs on the destination interface.

// Configure the source interface's routes on the destination interface.

func (c *ExternalNodeController) removeExternalNodeConfig() error {
	_ = "STUB: not implemented"
	return nil
}
