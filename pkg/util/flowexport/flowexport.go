// Copyright 2021 Antrea Authors
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

package flowexport

import (
	"time"

	flowaggregatorconfig "antrea.io/antrea/v2/pkg/config/flowaggregator"
)

// ParseFlowCollectorAddr parses the flow collector address input for flow exporter and aggregator
func ParseFlowCollectorAddr(addr string, defaultPort string, defaultProtocol string) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

// ParseFlowIntervalString parses the flow poll or export interval input string for flow exporter and aggregator
func ParseFlowIntervalString(intervalString string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

var protocolMap = map[string]flowaggregatorconfig.AggregatorTransportProtocol{
	"tcp":  flowaggregatorconfig.AggregatorTransportProtocolTCP,
	"tls":  flowaggregatorconfig.AggregatorTransportProtocolTLS,
	"udp":  flowaggregatorconfig.AggregatorTransportProtocolUDP,
	"none": flowaggregatorconfig.AggregatorTransportProtocolNone,
}

// ParseTransportProtocol parses the transport protocol input for the flow aggregator
func ParseTransportProtocol(transportProtocolInput flowaggregatorconfig.AggregatorTransportProtocol) (flowaggregatorconfig.AggregatorTransportProtocol, error) {
	_ = "STUB: not implemented"
	return *new(flowaggregatorconfig.AggregatorTransportProtocol), nil
}
