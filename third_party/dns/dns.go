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

// This package is copied and modified from https://github.com/miekg/dns because the original function Unpack cannot
// unpack fragmented DNS message.

package dns

import (
	godns "github.com/miekg/dns"
)

const (
	// Header.Bits
	_QR = 1 << 15 // query/response (response=1)
	_AA = 1 << 10 // authoritative
	_TC = 1 << 9  // truncated
	_RD = 1 << 8  // recursion desired
	_RA = 1 << 7  // recursion available
	_Z  = 1 << 6  // Z
	_AD = 1 << 5  // authenticated data
	_CD = 1 << 4  // checking disabled
)

// UnpackDNSMsgPartially is modified from https://github.com/miekg/dns/blob/6ad6301ae27dca6d7822baf1b05ff9c9e4ba56f4/msg.go#L883.
// It can unpack a DNS response with partial data only. More specifically, it unpacks the message header, the question
// section, and the answer section, while ignores the authority section and the additional section.
// It's used to get the question and answer sections when a DNS response is carried by a TCP packet but is fragmented.
func UnpackDNSMsgPartially(msg []byte, dns *godns.Msg) error { _ = "STUB: not implemented"; return nil }

func unpackMsgHdr(msg []byte, off int) (godns.Header, int, error) {
	_ = "STUB: not implemented"
	return *new(godns.Header), 0, nil
}

func unpackUint16(msg []byte, off int) (i uint16, off1 int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func setHdr(dns *godns.Msg, dh godns.Header) { _ = "STUB: not implemented"; return }

// _Z covers the zero bit, which should be zero; not sure why we set it to the opposite.

// unpackPartially is modified from https://github.com/miekg/dns/blob/6ad6301ae27dca6d7822baf1b05ff9c9e4ba56f4/msg.go#L826.
// It unpacks the message header, the question section, and the answer section, while ignores the authority section and
// the additional section.
func unpackPartially(dns *godns.Msg, dh godns.Header, msg []byte, off int) (err error) {
	_ = "STUB: not implemented"
	// If we are at the end of the message we should return *just* the
	// header. This can still be useful to the caller. 9.9.9.9 sends these
	// when responding with REFUSED for instance.
	return nil
}

// reset sections before returning

// Qdcount, Ancount, Nscount, Arcount can't be trusted, as they are
// attacker controlled. This means we can't use them to pre-allocate
// slices.

// Offset does not increase anymore, dh.Qdcount is a lie!

// The header counts might have been wrong so we need to update it

// Skip unpacking the authority section and the additional section.

func unpackQuestion(msg []byte, off int) (godns.Question, int, error) {
	_ = "STUB: not implemented"
	return *new(godns.Question), 0, nil
}

// unpackRRslice unpacks msg[off:] into an []RR.
// If we cannot unpack the whole array, then it will return nil
func unpackRRslice(l int, msg []byte, off int) (dst1 []godns.RR, off1 int, err error) {
	_ = "STUB: not implemented"

	// Don't pre-allocate, l may be under attacker control
	return nil, 0, nil
}

// If offset does not increase anymore, l is a lie
