//go:build linux
// +build linux

// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connections

import (
	"net/netip"

	"github.com/ti-mo/conntrack"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
)

// connTrackSystem implements ConnTrackDumper. This is for linux kernel datapath.
var _ ConnTrackDumper = new(connTrackSystem)

type connTrackSystem struct {
	nodeConfig           *config.NodeConfig
	serviceCIDRv4        netip.Prefix
	serviceCIDRv6        netip.Prefix
	isAntreaProxyEnabled bool
	connTrack            NetFilterConnTrack
}

// TODO: detect the endianness of the system when initializing conntrack dumper to handle situations on big-endian platforms.
// All connection labels are required to store in little endian format in conntrack dumper.
func NewConnTrackSystem(nodeConfig *config.NodeConfig, serviceCIDRv4 netip.Prefix, serviceCIDRv6 netip.Prefix, isAntreaProxyEnabled bool) *connTrackSystem {
	_ = "STUB: not implemented"
	return nil
}

// Do not fail, but continue after logging an error as we can still dump flows with missing information.

// DumpFlows opens netlink connection and dumps all the flows in Antrea ZoneID of conntrack table.
func (ct *connTrackSystem) DumpFlows(zoneFilter uint16) ([]*connection.Connection, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Get connection to netlink socket

// ZoneID filter is not supported currently in tl-mo/conntrack library.
// Link to issue: https://github.com/ti-mo/conntrack/issues/23
// Dump all flows in the conntrack table for now.

// NetFilterConnTrack interface helps for testing the code that contains the third party library functions ("github.com/ti-mo/conntrack")
type NetFilterConnTrack interface {
	Dial() error
	Close() error
	DumpFlowsInCtZone(zoneFilter uint16) ([]*connection.Connection, error)
}

type netFilterConnTrack struct {
	netlinkConn *conntrack.Conn
}

func (nfct *netFilterConnTrack) Dial() error {
	_ = "STUB: not implemented"
	// Get netlink client in current namespace
	return nil
}

func (nfct *netFilterConnTrack) Close() error { _ = "STUB: not implemented"; return nil }

func (nfct *netFilterConnTrack) DumpFlowsInCtZone(zoneFilter uint16) ([]*connection.Connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NetlinkFlowToAntreaConnection(conn *conntrack.Flow) *connection.Connection {
	_ = "STUB: not implemented"
	return nil
}

// github.com/ti-mo/conntrack uses native endianness (binary.NativeEndian), but we require a
// big-endian representation for the Labels / LabelsMask fields in connection.Connection.

// Get the stop time from dumped connection if the connection is terminated(dying state).

func SetupConntrackParameters() error { _ = "STUB: not implemented"; return nil }

func (ct *connTrackSystem) GetMaxConnections() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// reference: https://github.com/torvalds/linux/blob/master/net/netfilter/nf_conntrack_proto_tcp.c#L51-L62
func stateToString(state uint8) string { _ = "STUB: not implemented"; return "" }

// invalid state number
