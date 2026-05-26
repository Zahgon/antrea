// Copyright 2024 Antrea Authors.
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

package capture

import (
	"context"
	"net"

	"github.com/gopacket/gopacket"
	"golang.org/x/net/bpf"

	crdv1alpha1 "antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
)

type pcapCapture struct {
}

func NewPcapCapture() (*pcapCapture, error) {
	_ = "STUB: not implemented"
	return nil,

		// zeroFilter is a filter that will drop all packets.
		// see: https://github.com/antrea-io/antrea/issues/6815 for the user case.
		nil
}

func zeroFilter() []bpf.Instruction { _ = "STUB: not implemented"; return nil }

func (p *pcapCapture) Capture(ctx context.Context, device string, snapLen int, srcIP, dstIP net.IP, packet *crdv1alpha1.Packet, direction crdv1alpha1.CaptureDirection) (chan gopacket.Packet, error) {
	_ = "STUB: not implemented"
	// Compile the BPF filter in advance to reduce the time window between starting the capture and applying the filter.
	return nil, nil
}

// Install a BPF filter that won't match any packets
// see: https://natanyellin.com/posts/ebpf-filtering-done-right/.
// Packets which don’t match the target BPF can be received after the socket
// is created and before setsockopt is called. Those packets will remain
// in the socket’s buffer even after the BPF is applied and will later
// be transferred to the application via recv. Here we use a zero
// bpf filter(match no packet), then empty out any packets that arrived
// before the “zero-BPF” filter was applied. At this point the socket is
// definitely empty and it can’t fill up with junk because the zero-BPF
// is in place. Then we replace the zero-BPF with the real BPF we want.

// Drain the channel

// timeout: channel is drained so socket is drained
// install the correct BPF filter
