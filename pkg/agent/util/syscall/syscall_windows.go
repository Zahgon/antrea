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

package syscall

import (
	"net"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	AF_UNSPEC uint16 = uint16(windows.AF_UNSPEC)
	AF_INET   uint16 = uint16(windows.AF_INET)
	AF_INET6  uint16 = uint16(windows.AF_INET6)
)

// The following definitions are copied from Nldef header in Win32 API reference documentation.
// RouterDiscoveryBehavior defines the router discovery behavior.
type RouterDiscoveryBehavior int32

const (
	RouterDiscoveryDisabled  RouterDiscoveryBehavior = 0
	RouterDiscoveryEnabled   RouterDiscoveryBehavior = 1
	RouterDiscoveryDHCP      RouterDiscoveryBehavior = 2
	RouterDiscoveryUnchanged RouterDiscoveryBehavior = -1
)

// LinkLocalAddressBehavior defines the link local address behavior.
type LinkLocalAddressBehavior int32

const (
	LinkLocalAlwaysOff LinkLocalAddressBehavior = 0
	LinkLocalDelayed   LinkLocalAddressBehavior = 1
	LinkLocalAlwaysOn  LinkLocalAddressBehavior = 2
	LinkLocalUnchanged LinkLocalAddressBehavior = -1
)

const ScopeLevelCount = 16

// NlInterfaceOffloadRodFlags specifies a set of flags that indicate the offload
// capabilities for an IP interface.
type NlInterfaceOffloadRodFlags uint8

const (
	NlChecksumSupported         NlInterfaceOffloadRodFlags = 0x01
	nlOptionsSupported          NlInterfaceOffloadRodFlags = 0x02
	TlDatagramChecksumSupported NlInterfaceOffloadRodFlags = 0x04
	TlStreamChecksumSupported   NlInterfaceOffloadRodFlags = 0x08
	TlStreamOptionsSupported    NlInterfaceOffloadRodFlags = 0x10
	FastPathCompatible          NlInterfaceOffloadRodFlags = 0x20
	TlLargeSendOffloadSupported NlInterfaceOffloadRodFlags = 0x40
	TlGiantSendOffloadSupported NlInterfaceOffloadRodFlags = 0x80
)

type MibIPInterfaceRow struct {
	Family                               uint16
	Luid                                 uint64
	Index                                uint32
	MaxReassemblySize                    uint32
	Identifier                           uint64
	MinRouterAdvertisementInterval       uint32
	MaxRouterAdvertisementInterval       uint32
	AdvertisingEnabled                   bool
	ForwardingEnabled                    bool
	WeakHostSend                         bool
	WeakHostReceive                      bool
	UseAutomaticMetric                   bool
	UseNeighborUnreachabilityDetection   bool
	ManagedAddressConfigurationSupported bool
	OtherStatefulConfigurationSupported  bool
	AdvertiseDefaultRoute                bool
	RouterDiscoveryBehavior              RouterDiscoveryBehavior
	DadTransmits                         uint32
	BaseReachableTime                    uint32
	RetransmitTime                       uint32
	PathMtuDiscoveryTimeout              uint32
	LinkLocalAddressBehavior             LinkLocalAddressBehavior
	LinkLocalAddressTimeout              uint32
	ZoneIndices                          [ScopeLevelCount]uint32
	SitePrefixLength                     uint32
	Metric                               uint32
	NlMtu                                uint32
	Connected                            bool
	SupportsWakeUpPatterns               bool
	SupportsNeighborDiscovery            bool
	SupportsRouterDiscovery              bool
	ReachableTime                        uint32
	TransmitOffload                      NlInterfaceOffloadRodFlags
	ReceiveOffload                       NlInterfaceOffloadRodFlags
	DisableDefaultRoutes                 bool
}

type RawSockAddrInet struct {
	Family uint16
	data   [26]byte
}

func (a *RawSockAddrInet) IP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func (a *RawSockAddrInet) String() string { _ = "STUB: not implemented"; return "" }

func NewRawSockAddrInetFromIP(ip net.IP) *RawSockAddrInet { _ = "STUB: not implemented"; return nil }

type AddressPrefix struct {
	Prefix       RawSockAddrInet
	prefixLength uint8
	_            [2]byte // Add two bytes to keep alignment.
}

func (p *AddressPrefix) IPNet() *net.IPNet { _ = "STUB: not implemented"; return nil }

func (p *AddressPrefix) EqualsTo(ipNet *net.IPNet) bool { _ = "STUB: not implemented"; return false }

func (p *AddressPrefix) String() string { _ = "STUB: not implemented"; return "" }

func NewAddressPrefixFromIPNet(ipnet *net.IPNet) *AddressPrefix {
	_ = "STUB: not implemented"
	return nil
}

// NlRouteProtocol defines the routing mechanism that an IP route was added with.
type NlRouteProtocol uint32

const (
	RouteProtocolOther   NlRouteProtocol = 1
	RouteProtocolLocal   NlRouteProtocol = 2
	RouteProtocolNetMgmt NlRouteProtocol = 3
	RouteProtocolIcmp    NlRouteProtocol = 4
	RouteProtocolEgp     NlRouteProtocol = 5
	RouteProtocolGgp     NlRouteProtocol = 6
	RouteProtocolHello   NlRouteProtocol = 7
	RouteProtocolRip     NlRouteProtocol = 8
	RouteProtocolIsIs    NlRouteProtocol = 9
	RouteProtocolEsIs    NlRouteProtocol = 10
	RouteProtocolCisco   NlRouteProtocol = 11
	RouteProtocolBbn     NlRouteProtocol = 12
	RouteProtocolOspf    NlRouteProtocol = 13
	RouteProtocolBgp     NlRouteProtocol = 14
	RouteProtocolIdpr    NlRouteProtocol = 15
	RouteProtocolEigrp   NlRouteProtocol = 16
	RouteProtocolDvmrp   NlRouteProtocol = 17
	RouteProtocolRpl     NlRouteProtocol = 18
	RouteProtocolDhcp    NlRouteProtocol = 19

	//
	// Windows-specific definitions.
	//
	NT_AUTOSTATIC     NlRouteProtocol = 10002
	NT_STATIC         NlRouteProtocol = 10006
	NT_STATIC_NON_DOD NlRouteProtocol = 10007
)

// NlRouteOrigin defines the origin of the IP route.
type NlRouteOrigin uint32

const (
	NlroManual              NlRouteOrigin = 0
	NlroWellKnown           NlRouteOrigin = 1
	NlroDHCP                NlRouteOrigin = 2
	NlroRouterAdvertisement NlRouteOrigin = 3
	Nlro6to4                NlRouteOrigin = 4
)

type MibIPForwardRow struct {
	Luid              uint64
	Index             uint32
	DestinationPrefix AddressPrefix
	NextHop           RawSockAddrInet

	SitePrefixLength  uint8
	ValidLifetime     uint32
	PreferredLifetime uint32
	Metric            uint32
	Protocol          NlRouteProtocol

	Loopback             bool
	AutoconfigureAddress bool
	Publish              bool
	Immortal             bool

	Age    uint32
	Origin NlRouteOrigin
}

type MibIPForwardTable struct {
	NumEntries uint32
	Table      [1]MibIPForwardRow
}

var (
	modiphlpapi = syscall.NewLazyDLL("iphlpapi.dll")

	procGetIPInterfaceEntry  = modiphlpapi.NewProc("GetIpInterfaceEntry")
	procSetIPInterfaceEntry  = modiphlpapi.NewProc("SetIpInterfaceEntry")
	procCreateIPForwardEntry = modiphlpapi.NewProc("CreateIpForwardEntry2")
	procDeleteIPForwardEntry = modiphlpapi.NewProc("DeleteIpForwardEntry2")
	procGetIPForwardTable    = modiphlpapi.NewProc("GetIpForwardTable2")
	procFreeMibTable         = modiphlpapi.NewProc("FreeMibTable")
)

type NetIOInterface interface {
	GetIPInterfaceEntry(ipInterfaceRow *MibIPInterfaceRow) (errcode error)

	SetIPInterfaceEntry(ipInterfaceRow *MibIPInterfaceRow) (errcode error)

	CreateIPForwardEntry(ipForwardEntry *MibIPForwardRow) (errcode error)

	DeleteIPForwardEntry(ipForwardEntry *MibIPForwardRow) (errcode error)

	ListIPForwardRows(family uint16) ([]MibIPForwardRow, error)
}

type netIO struct {
	syscallN func(trap uintptr, args ...uintptr) (r1, r2 uintptr, err syscall.Errno)
	// It needs be declared as a variable and replaced during unit tests because the real getIPForwardTable function
	// converts a Pointer to a uintptr as an argument of syscallN, while converting a uintptr back to a Pointer in the
	// fake syscallN is not valid.
	getIPForwardTable func(family uint16, ipForwardTable **MibIPForwardTable) (errcode error)
}

func NewNetIO() NetIOInterface { _ = "STUB: not implemented"; return *new(NetIOInterface) }

func (n *netIO) GetIPInterfaceEntry(ipInterfaceRow *MibIPInterfaceRow) (errcode error) {
	_ = "STUB: not implemented"
	return nil
}

func (n *netIO) SetIPInterfaceEntry(ipInterfaceRow *MibIPInterfaceRow) (errcode error) {
	_ = "STUB: not implemented"
	return nil
}

func (n *netIO) CreateIPForwardEntry(ipForwardEntry *MibIPForwardRow) (errcode error) {
	_ = "STUB: not implemented"
	return nil
}

func (n *netIO) DeleteIPForwardEntry(ipForwardEntry *MibIPForwardRow) (errcode error) {
	_ = "STUB: not implemented"
	return nil
}

func (n *netIO) freeMibTable(table unsafe.Pointer) { _ = "STUB: not implemented"; return }

func getIPForwardTable(family uint16, ipForwardTable **MibIPForwardTable) (errcode error) {
	_ = "STUB: not implemented"
	return nil
}

func (n *netIO) ListIPForwardRows(family uint16) ([]MibIPForwardRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Copy the rows from the table into a new slice as the table's memory will be freed.
// Since MibIPForwardRow contains only value data (no references), the operation performs a deep copy.

func NewIPForwardRow() *MibIPForwardRow { _ = "STUB: not implemented"; return nil }
