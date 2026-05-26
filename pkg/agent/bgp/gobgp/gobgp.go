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

package gobgp

import (
	"context"

	gobgpapi "github.com/osrg/gobgp/v3/api"
	"github.com/osrg/gobgp/v3/pkg/server"

	"antrea.io/antrea/v2/pkg/agent/bgp"
)

const (
	ipv4AllZero = "0.0.0.0"
	ipv6AllZero = "::"
)

type Server struct {
	server       *server.BgpServer
	globalConfig *gobgpapi.Global
}

func NewGoBGPServer(globalConfig *bgp.GlobalConfig) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Server) AddPeer(ctx context.Context, peerConf bgp.PeerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) UpdatePeer(ctx context.Context, peerConf bgp.PeerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) RemovePeer(ctx context.Context, peerConf bgp.PeerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) GetPeers(ctx context.Context) ([]bgp.PeerStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) AdvertiseRoutes(ctx context.Context, routes []bgp.Route) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) WithdrawRoutes(ctx context.Context, routes []bgp.Route) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) GetRoutes(ctx context.Context, routeType bgp.RouteType, peerAddress string) ([]bgp.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertGoBGPPeerToPeerStatus(peer *gobgpapi.Peer) *bgp.PeerStatus {
	_ = "STUB: not implemented"
	return nil
}

// According to the gobgp code, all pointer fields in `peer *gobgpapi.Peer` used below should be set and non-nil
// when peers are listed. It's safe and harmless to keep the nil checks.

func convertGoBGPDestinationToRoute(destination *gobgpapi.Destination) *bgp.Route {
	_ = "STUB: not implemented"
	return nil
}

func convertRouteTypeToGoBGPTableType(routeType bgp.RouteType) gobgpapi.TableType {
	_ = "STUB: not implemented"
	return *new(gobgpapi.TableType)
}

func convertRouteToGoBGPPath(route *bgp.Route) *gobgpapi.Path {
	_ = "STUB: not implemented"
	return nil
}

func convertToGoBGPFamilyAfi(isIPv6 bool) gobgpapi.Family_Afi {
	_ = "STUB: not implemented"
	return *new(gobgpapi.Family_Afi)
}

func convertPeerConfigToGoBGPPeer(peerConfig bgp.PeerConfig) (*gobgpapi.Peer, error) {
	_ = "STUB: not implemented"
	// The following fields are required and are validated when the corresponding BGPPolicy is created.
	// Nonetheless, it is both safe and prudent to check them here as an additional safeguard.
	return nil, nil
}

// The following pointer fields are set to default values when the corresponding BGPPolicy is created, so they
// should not be nil. However, it is safe and harmless to include nil checks.

func convertGoBGPSessionStateToSessionState(s gobgpapi.PeerState_SessionState) bgp.SessionState {
	_ = "STUB: not implemented"
	return *new(bgp.SessionState)
}

func isValidIPString(ip string) bool { _ = "STUB: not implemented"; return false }
