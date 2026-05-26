/*
HTTP Content-Type Autonegotiation.

The functions in this package implement the behaviour specified in
http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html

Copyright (c) 2011, Open Knowledge Foundation Ltd.
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

    Redistributions of source code must retain the above copyright
    notice, this list of conditions and the following disclaimer.

    Redistributions in binary form must reproduce the above copyright
    notice, this list of conditions and the following disclaimer in
    the documentation and/or other materials provided with the
    distribution.

    Neither the name of the Open Knowledge Foundation Ltd. nor the
    names of its contributors may be used to endorse or promote
    products derived from this software without specific prior written
    permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Original source: https://github.com/munnerz/goautoneg

package goautoneg

// Accept represents a clause in an HTTP Accept Header.
type Accept struct {
	Type, SubType string
	Q             float64
	Params        map[string]string
}

// acceptSlice is defined to implement sort interface.
type acceptSlice []Accept

func (slice acceptSlice) Len() int { _ = "STUB: not implemented"; return 0 }

func (slice acceptSlice) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (slice acceptSlice) Swap(i, j int) { _ = "STUB: not implemented"; return }

func stringTrimSpaceCutset(r rune) bool { _ = "STUB: not implemented"; return false }

func nextSplitElement(s, sep string) (item string, remaining string) {
	_ = "STUB: not implemented"
	return "", ""
}

// ParseAccept parses an Accept Header string returning a sorted list of clauses.
func ParseAccept(header string) acceptSlice { _ = "STUB: not implemented"; return *new(acceptSlice) }

// Negotiate the most appropriate content_type given the accept header
// and a list of alternatives.
func Negotiate(header string, alternatives []string) (content_type string) {
	_ = "STUB: not implemented"
	return ""
}
