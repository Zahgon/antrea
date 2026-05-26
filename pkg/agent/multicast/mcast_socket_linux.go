//go:build linux && (arm || arm64 || amd64)
// +build linux
// +build arm arm64 amd64

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

package multicast

import (
	"net"

	multicastsyscall "antrea.io/antrea/v2/pkg/agent/util/syscall"
)

const (
	IGMPMsgNocache = multicastsyscall.IGMPMSG_NOCACHE
	MaxVIFs        = multicastsyscall.MAXVIFS
	SizeofIgmpmsg  = multicastsyscall.SizeofIgmpmsg
)

// setVIFToInterface adds a virtual interface to the multicast socket for interface with index ifIndex.
func setVIFToInterface(fd int, vif uint16, ifIndex int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Socket) AddMrouteEntry(src net.IP, group net.IP, iif uint16, oifVIFs []uint16) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// GetMroutePacketCount returns the number of routed packets by the multicast route entry.
// The current implementation only supports IPv4 multicast routes.
func (s *Socket) GetMroutePacketCount(src net.IP, group net.IP) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Socket) DelMrouteEntry(src net.IP, group net.IP, iif uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Socket) FlushMRoute() { _ = "STUB: not implemented"; return }

func CreateMulticastSocket() (*Socket, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Socket) AllocateVIFs(interfaceNames []string, startVIF uint16) ([]uint16, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Socket) MulticastInterfaceJoinMgroup(mgroup net.IP, ifaceIP net.IP, ifaceName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Socket) MulticastInterfaceLeaveMgroup(mgroup net.IP, ifaceIP net.IP, ifaceName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Socket) GetFD() int { _ = "STUB: not implemented"; return 0 }

type Socket struct {
	sockFD int
}
