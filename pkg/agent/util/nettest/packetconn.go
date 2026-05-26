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

package nettest

import (
	"net"
	"time"
)

type Packet struct {
	Bytes []byte
	Addr  net.Addr
}

// PacketConn implements the net.PacketConn interface, and can be used to test networking code that
// requires a packet-oriented network connection. The connection peer is represented by 2 channels
// to send and receive packets.
type PacketConn struct {
	addr    net.Addr
	inCh    chan *Packet
	outCh   chan *Packet
	closeCh chan struct{}
}

var _ net.PacketConn = (*PacketConn)(nil)

func NewPacketConn(localAddr net.Addr, inCh chan *Packet, outCh chan *Packet) *PacketConn {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PacketConn) ReadFrom(p []byte) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	// Check if connection is closed once before the select statement. Otherwise we may end up
	// reading a packet even if the connection was already closed before calling this
	// function. It is still possible for the connection to be closed between this check and the
	// select, but it doesn't matter in this case because that would mean the 2 function calls
	// (Close and ReadFrom) are concurrent.
	return 0, *new(net.Addr), nil
}

func (pc *PacketConn) WriteTo(p []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"
	// See the comment in ReadFrom.
	return 0, nil
}

func (pc *PacketConn) Close() error {
	_ = "STUB: not implemented"
	// panic if the connection has already been closed
	return nil
}

func (pc *PacketConn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (pc *PacketConn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (pc *PacketConn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (pc *PacketConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// Send is a convenience function that will send a packet without blocking. If this is not possible,
// it will return an error. Send does not check whether the connection is closed.
func (pc *PacketConn) Send(p []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Receive is a convenience function that will receive a packet without blocking. If no packet is
// available, it will return an error. Receive does not check whether the connection is closed.
func (pc *PacketConn) Receive() ([]byte, net.Addr, error) {
	_ = "STUB: not implemented"
	return nil, *new(net.Addr), nil
}

func (pc *PacketConn) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (pc *PacketConn) closedConnectionError(op string) error { _ = "STUB: not implemented"; return nil }

// PacketConnPipe creates 2 instances of PacketConn which represent the 2 endpoints of a
// packet-oriented connection. Every packet sent on one side will be received on the other side,
// and we never check whether the address actually matches when sending a packet.
func PacketConnPipe(addr1, addr2 net.Addr, capacity int) (*PacketConn, *PacketConn) {
	_ = "STUB: not implemented"
	return nil, nil
}
