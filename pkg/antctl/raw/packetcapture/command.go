// Copyright 2025 Antrea Authors.
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

package packetcapture

import (
	"context"
	"io"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"antrea.io/antrea/v2/pkg/antctl/raw"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	antrea "antrea.io/antrea/v2/pkg/client/clientset/versioned"
)

var (
	defaultTimeout          = time.Second * 60
	maxPacketCaptureTimeout = time.Second * 300
	Command                 *cobra.Command
	getCopier               = getPodFileCopier
	defaultFS               = afero.NewOsFs()
)

type packetCaptureOptions struct {
	source       string
	dest         string
	nowait       bool
	timeout      time.Duration
	number       int32
	flow         string
	outputDir    string
	capturePoint string
	direction    string
}

var options = &packetCaptureOptions{}

var packetCaptureExample = `  Start capturing packets from pod1 to pod2, both Pods are in Namespace default
  $ antctl packetcapture -S pod1 -D pod2
  Start capturing packets from pod1 in Namespace ns1 to a destination IP
  $ antctl packetcapture -S ns1/pod1 -D 192.168.123.123
  Start capturing packets from pod1 to pod2, captures at dst pod
  $ antctl packetcapture -S pod1 -D pod2 -p Destination
  Start capturing TCP FIN packets from pod1 to pod2, with destination port 80
  $ antctl packetcapture -S pod1 -D pod2 -f tcp,tcp_dst=80,tcp_flags=+fin
  Start capturing TCP SYNs that are not ACKs from pod1 to pod2, with destination port 80
  $ antctl packetcapture -S pod1 -D pod2 -f tcp,tcp_dst=80,tcp_flags=+syn-ack
  Start capturing IPv6 TCP SYNs packets from pod1 to pod2, with destination port 80
  $ antctl packetcapture -S pod1 -D pod2 -f ipv6,tcp,tcp_dst=80,tcp_flags=+syn
  Start capturing UDP packets from pod1 to pod2, with destination port 1234
  $ antctl packetcapture -S pod1 -D pod2 -f udp,udp_dst=1234
  Start capturing ICMP destination unreachable (host unreachable) packets from pod1 to pod2
  $ antctl packetcapture -S pod1 -D pod2 -f icmp,icmp_type=icmp-unreach,icmp_code=1
  Start capturing ICMP echo packets from pod1 to pod2
  $ antctl packetcapture -S pod1 -D pod2 -f icmp,icmp_type=8
  Start capturing packets in both directions between pod1 and pod2
  $ antctl packetcapture -S pod1 -D pod2 -d Both
  Start capturing ICMPv6 destination unreachable (host unreachable) packets from pod1 to pod2
  $ antctl packetcapture -S pod1 -D pod2 -f icmpv6,icmpv6_type=icmpv6-unreach,icmpv6_code=1
  Start capturing ICMPv6 echo reply packets from pod1 to pod2
  $ antctl packetcapture -S pod1 -D pod2 -f icmpv6,icmpv6_type=129
  Save the packets file to a specified directory
  $ antctl packetcapture -S 192.168.123.123 -D pod2 -f tcp,tcp_dst=80 -o /tmp
`

func init() {
	Command = &cobra.Command{
		Use:     "packetcapture",
		Short:   "Start capture packets",
		Long:    "Start capturing packets on the target flow",
		Aliases: []string{"pc", "packetcaptures"},
		Example: packetCaptureExample,
		RunE:    packetCaptureRunE,
	}

	Command.Flags().StringVarP(&options.source, "source", "S", "", "source of the the PacketCapture: Namespace/Pod, Pod, or IP")
	Command.Flags().StringVarP(&options.dest, "destination", "D", "", "destination of the PacketCapture: Namespace/Pod, Pod, or IP")
	Command.Flags().Int32VarP(&options.number, "number", "n", 1, "target number of packets to capture, the capture will stop when it is reached")
	Command.Flags().StringVarP(&options.flow, "flow", "f", "", "specify the flow (packet headers) of the PacketCapture, including tcp_src, tcp_dst, tcp_flags, udp_src, udp_dst, icmp_type, icmp_code")
	Command.Flags().StringVarP(&options.capturePoint, "capture-point", "p", "", "specify where the packet capture should be performed: Source or Destination")
	Command.Flags().BoolVarP(&options.nowait, "nowait", "", false, "if set, command returns without retrieving results")
	Command.Flags().StringVarP(&options.outputDir, "output-dir", "o", ".", "save the packets file to the target directory")
	Command.Flags().StringVarP(&options.direction, "direction", "d", "SourceToDestination", "direction of the traffic to capture: SourceToDestination, DestinationToSource, or Both")
}

var protocols = map[string]int32{
	"icmp":   1,
	"tcp":    6,
	"udp":    17,
	"icmpv6": 58,
}

var tcpFlags = map[string]int32{
	"fin": 1,
	"syn": 1 << 1,
	"rst": 1 << 2,
	"psh": 1 << 3,
	"ack": 1 << 4,
	"urg": 1 << 5,
	"ece": 1 << 6,
	"cwr": 1 << 7,
}

func getPodFileCopier(config *rest.Config, client kubernetes.Interface) raw.PodFileCopier {
	_ = "STUB: not implemented"
	return *new(raw.PodFileCopier)
}

func getConfigAndClients(cmd *cobra.Command) (*rest.Config, kubernetes.Interface, antrea.Interface, error) {
	_ = "STUB: not implemented"
	return nil, *new(kubernetes.Interface), *new(antrea.Interface), nil
}

func getPCName(options *packetCaptureOptions) string { _ = "STUB: not implemented"; return "" }

func packetCaptureRunE(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func packetCaptureRun(ctx context.Context, out io.Writer, restConfig *rest.Config, k8sClient kubernetes.Interface, antreaClient antrea.Interface, options *packetCaptureOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// add extra timeout to make sure the wait won't be interrupted before PacketCapture timeout.

func parseEndpoint(endpoint string) (*v1alpha1.PodReference, *string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFlowFields(flow string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tokenizeTCPFlags parses tcp_flags value and returns two slices: set (flags that must be set) and unset (flags that must be unset).
func tokenizeTCPFlags(r string) ([]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func parseFlow(options *packetCaptureOptions) (*v1alpha1.Packet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseDirection(direction string) (v1alpha1.CaptureDirection, error) {
	_ = "STUB: not implemented"
	// This case should not occur in practice as the direction flag is defaulted to SourceToDestination
	return *new(v1alpha1.CaptureDirection), nil
}

func parseCapturePoint(captPointStr string) (v1alpha1.CapturePoint, error) {
	_ = "STUB: not implemented"
	return *new(v1alpha1.CapturePoint), nil
}

func newPacketCapture(options *packetCaptureOptions) (*v1alpha1.PacketCapture, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
