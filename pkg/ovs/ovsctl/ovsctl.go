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

package ovsctl

import (
	"net"
)

// Shell exits with 127 if the command to execute is not found.
const exitCodeCommandNotFound = 127

var (
	IPAndNWProtos = []string{"ip", "icmp", "tcp", "udp", "sctp", "ipv6", "icmp6", "tcp6", "udp6", "sctp6"}
	// Some typical non-IP packet types.
	// "dl_type=0x0800" can be used to indicate an IP packet too, but as it is not
	// a common way, here we simply assume "dl_type=" is used for non-IP types
	// only.
	nonIPDLTypes = []string{"arp", "rarp", "dl_type="}
)

type DPFeature string

const (
	CTStateFeature    DPFeature = "CT state"
	CTZoneFeature     DPFeature = "CT zone"
	CTMarkFeature     DPFeature = "CT mark"
	CTLabelFeature    DPFeature = "CT label"
	CTStateNATFeature DPFeature = "CT state NAT"
)

var knownFeatures = map[DPFeature]struct{}{
	CTStateFeature:    {},
	CTZoneFeature:     {},
	CTMarkFeature:     {},
	CTLabelFeature:    {},
	CTStateNATFeature: {},
}

// TracingRequest defines tracing request parameters.
type TracingRequest struct {
	InPort string // Input port.
	SrcIP  net.IP
	DstIP  net.IP
	SrcMAC net.HardwareAddr
	DstMAC net.HardwareAddr
	Flow   string
	// Whether in_port field in Flow can override InPort.
	AllowOverrideInPort bool
}

type ovsCtlClient struct {
	bridge          string
	ovsOfctlRunner  OVSOfctlRunner
	ovsAppctlRunner OVSAppctlRunner
}

func NewClient(bridge string) *ovsCtlClient { _ = "STUB: not implemented"; return nil }

func (c *ovsCtlClient) Trace(req *TracingRequest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Do not allow overriding destination IP.

// Always allow overriding source and destination MACs.

// IP or IP protocol is already specified in flow. No need to add "ip"/"ipv6" in
// flow.

// Add default IP TTL.

// "ip" or IP protocol must be set before "nw_ttl", "nw_src", "nw_dst", and
// "tp_port". For IPv6 packet, "ipv6" is required as a precondition.

// getNwSrcKey returns keys of IP address family and IP source which are supported in ovs-appctl command according
// to the given IP.
func getNwSrcKey(ip net.IP) (string, string) { _ = "STUB: not implemented"; return "", "" }

func getNwDstKey(ip net.IP) (string, string) { _ = "STUB: not implemented"; return "", "" }

func (c *ovsCtlClient) runTracing(flow string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Remove "\r" to avoid format issue on Windows.

func (c *ovsCtlClient) RunAppctlCmd(cmd string, needsBridge bool, args ...string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ovsCtlClient) GetDPFeatures() (map[DPFeature]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteDPInterface deletes OVS datapath interface, and it returns with no error if the interface does not exist.
func (c *ovsCtlClient) DeleteDPInterface(name string) error { _ = "STUB: not implemented"; return nil }

func newBadRequestError(msg string) BadRequestError {
	_ = "STUB: not implemented"
	return *new(BadRequestError)
}

func NewExecError(err error, errorOutput string) *ExecError { _ = "STUB: not implemented"; return nil }

func (c *ovsCtlClient) DumpFlows(args ...string) ([]string, error) {
	_ = "STUB: not implemented"
	// Print table and port names.
	return nil, nil
}

func (c *ovsCtlClient) DumpFlowsWithoutTableNames(args ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ovsCtlClient) parseFlowEntries(flowDump []byte) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip the non-flow line, which is printed when using parameter "--no-names" in tests.

func (c *ovsCtlClient) DumpMatchedFlow(matchStr string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ovs-ofctl dump-flows can return multiple flows that match matchStr, here we
// check and return only the one that exactly matches matchStr (no extra match
// conditions).

// No exactly matched flow found.

func (c *ovsCtlClient) DumpTableFlows(table uint8, filters ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ovsCtlClient) DumpGroup(groupID uint32) (string, error) {
	_ = "STUB: not implemented"
	// Only OpenFlow 1.5 and later support dumping a specific group. Earlier
	// versions of OpenFlow always dump all groups. But when OpenFlow
	// version is not specified, ovs-ofctl defaults to use OpenFlow10 but
	// with the Nicira extensions enabled, which can support dumping a
	// single group too. So here, we do not specify Openflow15 to run the
	// command.
	return "", nil
}

// Skip the first line.

// No group found.

// Should have at most one line (group) returned.

func (c *ovsCtlClient) DumpGroups() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Skip the first line.

func (c *ovsCtlClient) DumpPortsDesc() ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip the first line.

// If the line starts with a port number, it should be the first line of an OF port. There should be some
// subsequent lines to describe the status of the current port, which start with multiple while-spaces.

func (c *ovsCtlClient) SetPortNoFlood(ofport int) error {
	_ = "STUB: not implemented"
	// This command does not have standard output, and only has standard err when running with error.
	// NOTE THAT, THIS CONFIGURATION MUST WORK WITH OpenFlow10.
	return nil
}

func (c *ovsCtlClient) RunOfctlCmd(cmd string, args ...string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func flowExactMatch(matchStr, flowStr string) bool {
	_ = "STUB: not implemented"
	// Get the match string which starts with "priority=".
	return false
}

// Skip "priority=".

// The match condition is not included in matchStr.
