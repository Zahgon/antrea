// Copyright 2020 Antrea Authors
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

package traceflow

import (
	"io"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"

	"antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	antrea "antrea.io/antrea/v2/pkg/client/clientset/versioned"
)

const defaultTimeout time.Duration = time.Second * 10

var (
	Command *cobra.Command
	option  = &struct {
		source      string
		destination string
		outputType  string
		flow        string
		liveTraffic bool
		droppedOnly bool
		timeout     time.Duration
		nowait      bool
	}{}
	getClients = getK8sClient
)

var protocols = map[string]int32{
	"icmp": 1,
	"tcp":  6,
	"udp":  17,
}

type CapturedPacket struct {
	SrcIP           string                   `json:"srcIP" yaml:"srcIP"`
	DstIP           string                   `json:"dstIP" yaml:"dstIP"`
	Length          int32                    `json:"length" yaml:"length"`
	IPHeader        *v1beta1.IPHeader        `json:"ipHeader,omitempty" yaml:"ipHeader,omitempty"`
	IPv6Header      *v1beta1.IPv6Header      `json:"ipv6Header,omitempty" yaml:"ipv6Header,omitempty"`
	TransportHeader *v1beta1.TransportHeader `json:"transportHeader,omitempty" yaml:"tranportHeader,omitempty"`
}

// Response is the response of antctl Traceflow.
type Response struct {
	Name           string                 `json:"name" yaml:"name"`                                         // Traceflow name
	Phase          v1beta1.TraceflowPhase `json:"phase,omitempty" yaml:"phase,omitempty"`                   // Traceflow phase
	Reason         string                 `json:"reason,omitempty" yaml:"reason,omitempty"`                 // Traceflow phase reason
	Source         string                 `json:"source,omitempty" yaml:"source,omitempty"`                 // Traceflow source, e.g. "default/pod0"
	Destination    string                 `json:"destination,omitempty" yaml:"destination,omitempty"`       // Traceflow destination, e.g. "default/pod1"
	NodeResults    []v1beta1.NodeResult   `json:"results,omitempty" yaml:"results,omitempty"`               // Traceflow node results
	CapturedPacket *CapturedPacket        `json:"capturedPacket,omitempty" yaml:"capturedPacket,omitempty"` // Captured packet in live-traffic Traceflow
}

func init() {
	Command = &cobra.Command{
		Use:     "traceflow",
		Short:   "Start a Traceflows",
		Long:    "Start a Traceflows from one Pod to another Pod/Service/IP.",
		Aliases: []string{"tf", "traceflows"},
		Example: `  Start a Traceflow from pod1 to pod2, both Pods are in Namespace default
  $antctl traceflow -S pod1 -D pod2
  Start a Traceflow from pod1 in Namepace ns1 to a destination IP
  $antctl traceflow -S ns1/pod1 -D 123.123.123.123
  Start a Traceflow from pod1 to Service svc1 in Namespace ns1
  $antctl traceflow -S pod1 -D ns1/svc1 -f tcp,tcp_dst=80
  Start a Traceflow from pod1 to pod2, with a UDP packet to destination port 1234
  $antctl traceflow -S pod1 -D pod2 -f udp,udp_dst=1234
  Start a Traceflow for live TCP traffic from pod1 to svc1, with 1 minute timeout
  $antctl traceflow -S pod1 -D svc1 -f tcp --live-traffic -t 1m
  Start a Traceflow to capture the first dropped TCP packet to pod1 on port 80, within 10 minutes
  $antctl traceflow -D pod1 -f tcp,tcp_dst=80 --live-traffic --dropped-only -t 10m
`,
		RunE: runE,
		Args: cobra.NoArgs,
	}

	Command.Flags().StringVarP(&option.source, "source", "S", "", "source of the Traceflow: Namespace/Pod, Pod, or IP")
	Command.Flags().StringVarP(&option.destination, "destination", "D", "", "destination of the Traceflow: Namespace/Pod, Pod, Namespace/Service, Service or IP")
	Command.Flags().StringVarP(&option.outputType, "output", "o", "yaml", "output type: yaml (default), json")
	Command.Flags().StringVarP(&option.flow, "flow", "f", "", "specify the flow (packet headers) of the Traceflow packet, including tcp_src, tcp_dst, tcp_flags, udp_src, udp_dst, ipv6")
	Command.Flags().BoolVarP(&option.liveTraffic, "live-traffic", "L", false, "if set, the Traceflow will trace the first packet of the matched live traffic flow")
	Command.Flags().BoolVarP(&option.droppedOnly, "dropped-only", "", false, "if set, capture only the dropped packet in a live-traffic Traceflow")
	Command.Flags().BoolVarP(&option.nowait, "nowait", "", false, "if set, command returns without retrieving results")
}

func getK8sClient(cmd *cobra.Command) (kubernetes.Interface, antrea.Interface, error) {
	_ = "STUB: not implemented"
	return *new(kubernetes.Interface), *new(antrea.Interface), nil
}

func runE(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

// Still output the Traceflow results if any.

func newTraceflow(client kubernetes.Interface) (*v1beta1.Traceflow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dstIsPod(client kubernetes.Interface, ns string, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func parseFlow() (*v1beta1.Packet, error) { _ = "STUB: not implemented"; return nil, nil }

func getPortFields(cleanFlow string) (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func output(tf *v1beta1.Traceflow, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

func yamlOutput(r *Response, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

func jsonOutput(r *Response, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

func getTFName(prefix string) string {
	_ = "STUB: not implemented"
	// prefix may contain IPv6 address. Replace "::"  and ":" to make it a valid RFC 1123 subdomain.
	return ""
}
