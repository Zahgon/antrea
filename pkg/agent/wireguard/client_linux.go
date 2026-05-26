//go:build linux
// +build linux

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

package wireguard

import (
	"io"
	"net"
	"sync"

	"github.com/vishvananda/netlink"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/util"
)

const defaultWireGuardInterfaceName = "antrea-wg0"

var zeroKey = wgtypes.Key{}

// wgctrlClient is an interface to mock wgctrl.Client
type wgctrlClient interface {
	io.Closer
	Devices() ([]*wgtypes.Device, error)
	Device(name string) (*wgtypes.Device, error)
	ConfigureDevice(name string, config wgtypes.Config) error
}

var _ Interface = (*client)(nil)

var (
	linkAdd                    = netlink.LinkAdd
	linkSetUp                  = netlink.LinkSetUp
	linkSetMTU                 = netlink.LinkSetMTU
	utilConfigureLinkAddresses = util.ConfigureLinkAddresses
)

type client struct {
	wgClient                wgctrlClient
	nodeName                string
	privateKey              wgtypes.Key
	peerPublicKeyByNodeName *sync.Map
	wireGuardConfig         *config.WireGuardConfig
	gatewayConfig           *config.GatewayConfig
}

func New(nodeConfig *config.NodeConfig, wireGuardConfig *config.WireGuardConfig) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func (client *client) Init(ipv4 net.IP, ipv6 net.IP) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Ignore existing link as it may have already been created or managed by userspace process, just ensure the MTU
// is set correctly.

// Configure the IP addresses same as Antrea gateway so iptables MASQUERADE target will select it as source address.
// It's necessary to make Service traffic requiring SNAT (e.g. host to ClusterIP, external to NodePort) accepted by
// peer Node and to make their response routed back correctly.
// If ipv4 or ipv6 is not provided, the IP address from client's Gateway configuration will be used.
// It uses "/32" mask for IPv4 address and "/128" mask for IPv6 address to avoid impacting routes on Antrea gateway.

// This must be executed after netlink.LinkSetUp as the latter ensures link.Attrs().Index is set.

// WireGuard private key will be persistent across agent restarts. So we only need to
// generate a new private key if it is empty (all zero).

func (client *client) RemoveStalePeers(currentPeerPublickeys map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Save known Node name and public key mappings for tracking of public key changes when calling UpdatePeer.

func (client *client) UpdatePeer(nodeName, publicKeyString string, peerNodeIP net.IP, podCIDRs []*net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

// delete old peer by public key.

func (client *client) deletePeerByPublicKey(pubKey wgtypes.Key) error {
	_ = "STUB: not implemented"
	return nil
}

func (client *client) DeletePeer(nodeName string) error { _ = "STUB: not implemented"; return nil }

func (client *client) CleanUp() error { _ = "STUB: not implemented"; return nil }
