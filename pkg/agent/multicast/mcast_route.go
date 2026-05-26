// Copyright 2021 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint: unused // a lot of this code is unused for Windows since the multicast feature is not implemented yet
package multicast

import (
	"net"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/cache"

	"antrea.io/antrea/v2/pkg/agent/config"
)

const (
	GroupNameIndexName      = "groupName"
	MulticastFlag           = "multicast"
	MulticastRecvBufferSize = 128
)

func newRouteClient(nodeconfig *config.NodeConfig, groupCache cache.Indexer, multicastSocket RouteInterface, multicastInterfaces sets.Set[string], flexibleIPAMEnabled bool) *MRouteClient {
	_ = "STUB: not implemented"
	return nil
}

func (c *MRouteClient) Initialize() error { _ = "STUB: not implemented"; return nil }

// Allocate VIF for each interface in multicastInterfaceNames and gatewayInterface.
// The VIFs will be later used for multicast route configuration.

// MRouteClient configures static multicast route.
type MRouteClient struct {
	// igmpMsgChan is used for processing IGMPMsg reading from sockFD in parallel
	igmpMsgChan               chan []byte
	nodeConfig                *config.NodeConfig
	multicastInterfaces       []string
	inboundRouteCache         cache.Indexer
	outboundRouteCache        cache.Indexer
	groupCache                cache.Indexer
	socket                    RouteInterface
	multicastInterfaceConfigs []multicastInterfaceConfig
	internalInterfaceVIF      uint16
	externalInterfaceVIFs     []uint16
	flexibleIPAMEnabled       bool
	// vif16bit is true on kernels >= 5.10 where the VIF field in igmpmsg
	// is 16 bits wide; false on older kernels where it is 8 bits wide.
	vif16bit bool
}

// multicastInterfacesJoinMgroup allows multicast interfaces to join multicast group,
// by making these interfaces accept multicast traffic with multicast ip:mgroup.
// https://tldp.org/HOWTO/Multicast-HOWTO-6.html#ss6.4
func (c *MRouteClient) multicastInterfacesJoinMgroup(mgroup net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MRouteClient) multicastInterfacesLeaveMgroup(mgroup net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// processIGMPNocacheMsg reads igmpMsg from the multicast socket and configures
// multicast route based on VIF value in the message.
func (c *MRouteClient) processIGMPNocacheMsg(igmpMsg []byte) { _ = "STUB: not implemented"; return }

// Skip inbound multicast traffic when there is no multicast receiver Pod
// listening the msg.Dst group.

// Prevent adding route entries for unrecognized VIF.

func (c *MRouteClient) deleteInboundMrouteEntryByGroup(group net.IP) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *MRouteClient) deleteInboundMRoute(mRoute *inboundMulticastRouteEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MRouteClient) deleteOutboundMRoute(mRoute *outboundMulticastRouteEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// addOutboundMrouteEntry configures multicast route from Antrea gateway to all the multicast interfaces,
// allowing multicast srcNode Pods to send multicast traffic to external.
func (c *MRouteClient) addOutboundMrouteEntry(src net.IP, group net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// addInboundMrouteEntry configures multicast route from multicast interface to Antrea gateway
// to allow multicast receiver Pods to receive multicast traffic from external.
func (c *MRouteClient) addInboundMrouteEntry(src net.IP, group net.IP, inboundVIF uint16) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Field pktCount and updatedTime are used for removing stale multicast routes.
type multicastRouteEntry struct {
	group       string
	src         string
	pktCount    uint32
	updatedTime time.Time
}

// outboundMulticastRouteEntry encodes the outbound multicast routing entry.
// For example,
//
//	type outboundMulticastRouteEntry struct {
//		group "226.94.9.9"
//		src   "10.0.0.55"
//	} encodes the multicast route entry from Antrea gateway to multicast interfaces
//
// (10.0.0.55,226.94.9.9)           Iif: antrea-gw0      Oifs: list of multicastInterfaces.
//
// The iif is always Antrea gateway and oifs are always outbound interfaces
// so we do not put them in the struct.
type outboundMulticastRouteEntry struct {
	multicastRouteEntry
}

// inboundMulticastRouteEntry encodes the inbound multicast routing entry.
// It has extra field vif to represent inbound interface VIF.
// For example,
//
//	type inboundMulticastRouteEntry struct {
//		group "226.94.9.9"
//		src   "10.0.0.55"
//		vif   vif of wlan0
//	} encodes the multicast route entry from wlan0 to Antrea gateway
//
// (10.0.0.55,226.94.9.9)           Iif: wlan0      Oifs: antrea-gw0.
// The oif is always Antrea gateway so we do not put it in the struct.
type inboundMulticastRouteEntry struct {
	multicastRouteEntry
	vif uint16
}

func getMulticastInboundEntryKey(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getMulticastOutboundEntryKey(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func inboundGroupIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// setMulticastInterfaces tries to compute all the multicast interfaces used to
// accept and send multicast traffic based on the provided multicastInterfaces.
func (c *MRouteClient) setMulticastInterfaces() { _ = "STUB: not implemented"; return }

func (c *MRouteClient) worker(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// This struct is result of parsing igmpmsg from the kernel
// with fields we interest.
type parsedIGMPMsg struct {
	VIF uint16
	Src net.IP
	Dst net.IP
}

type multicastInterfaceConfig struct {
	Name     string
	IPv4Addr *net.IPNet
	IPv6Addr *net.IPNet
}

type RouteInterface interface {
	// MulticastInterfaceJoinMgroup enables interface with name ifaceName and IP ifaceIP
	// joins multicast group IP mgroup.
	MulticastInterfaceJoinMgroup(mgroup net.IP, ifaceIP net.IP, ifaceName string) error
	// MulticastInterfaceLeaveMgroup enables interface with name ifaceName and IP ifaceIP
	// leaves multicast group IP mgroup.
	MulticastInterfaceLeaveMgroup(mgroup net.IP, ifaceIP net.IP, ifaceName string) error
	// AddMrouteEntry adds multicast route with specified source(src), multicast group IP(group),
	// inbound multicast interface(iif) and outbound multicast interfaces(oifs).
	AddMrouteEntry(src net.IP, group net.IP, iif uint16, oifs []uint16) error
	// GetMroutePacketCount returns the number of routed packets by the multicast route entry.
	GetMroutePacketCount(src net.IP, group net.IP) (uint32, error)
	// DelMrouteEntry deletes multicast route with specified source(src), multicast group IP(group),
	// inbound multicast interface(iif).
	DelMrouteEntry(src net.IP, group net.IP, iif uint16) error
	// FlushMRoute flushes static multicast routing entries.
	FlushMRoute()
	// GetFD returns socket file descriptor.
	GetFD() int
	// AllocateVIFs allocate VIFs to interfaces, starting from startVIF.
	AllocateVIFs(interfaceNames []string, startVIF uint16) ([]uint16, error)
}
