// Copyright 2025 Antrea Authors
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

package exporter

import (
	ipfixentities "github.com/vmware/go-ipfix/pkg/entities"

	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
	"antrea.io/antrea/v2/pkg/ipfix"
)

var (
	IANAInfoElementsCommon = []string{
		"flowStartSeconds",
		"flowEndSeconds",
		"flowEndReason",
		"sourceTransportPort",
		"destinationTransportPort",
		"protocolIdentifier",
		"packetTotalCount",
		"octetTotalCount",
		"packetDeltaCount",
		"octetDeltaCount",
	}
	IANAInfoElementsIPv4 = append(IANAInfoElementsCommon, []string{"sourceIPv4Address", "destinationIPv4Address"}...)
	IANAInfoElementsIPv6 = append(IANAInfoElementsCommon, []string{"sourceIPv6Address", "destinationIPv6Address"}...)
	// IANAReverseInfoElements contain substring "reverse" which is an indication to get reverse element of go-ipfix library.
	IANAReverseInfoElements = []string{
		"reversePacketTotalCount",
		"reverseOctetTotalCount",
		"reversePacketDeltaCount",
		"reverseOctetDeltaCount",
	}
	antreaInfoElementsCommon = []string{
		"sourcePodName",
		"sourcePodNamespace",
		"sourceNodeName",
		"destinationPodName",
		"destinationPodNamespace",
		"destinationNodeName",
		"destinationServicePort",
		"destinationServicePortName",
		"ingressNetworkPolicyName",
		"ingressNetworkPolicyNamespace",
		"ingressNetworkPolicyType",
		"ingressNetworkPolicyRuleName",
		"ingressNetworkPolicyRuleAction",
		"egressNetworkPolicyName",
		"egressNetworkPolicyNamespace",
		"egressNetworkPolicyType",
		"egressNetworkPolicyRuleName",
		"egressNetworkPolicyRuleAction",
		"tcpState",
		"flowType",
		"egressName",
		"egressIP",
		"egressNodeName",
	}
	AntreaInfoElementsIPv4 = append(antreaInfoElementsCommon, []string{"destinationClusterIPv4"}...)
	AntreaInfoElementsIPv6 = append(antreaInfoElementsCommon, []string{"destinationClusterIPv6"}...)
)

type ipfixExporter struct {
	process        ipfix.IPFIXExportingProcess
	collectorProto string
	v4Enabled      bool
	v6Enabled      bool
	elementsListv4 []ipfixentities.InfoElementWithValue
	elementsListv6 []ipfixentities.InfoElementWithValue
	ipfixSet       ipfixentities.Set
	templateIDv4   uint16
	templateIDv6   uint16
	registry       ipfix.IPFIXRegistry
	nodeName       string
	obsDomainID    uint32
}

func NewIPFIXExporter(collectorProto string, nodeName string, obsDomainID uint32, v4Enabled, v6Enabled bool) *ipfixExporter {
	_ = "STUB: not implemented"
	// Initialize IPFIX registry
	return nil
}

func (e *ipfixExporter) ConnectToCollector(addr string, tlsConfig *TLSConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// TempRefTimeout specifies how often the exporting process should send the template
// again. It is only relevant when using the UDP protocol. We use 0 to tell the go-ipfix
// library to use the default value, which should be 600s as per the IPFIX standards.

func (e *ipfixExporter) Export(conn *connection.Connection) error {
	_ = "STUB: not implemented"
	// TODO: more records per data set will be supported when go-ipfix supports size check when adding records
	return nil
}

func (e *ipfixExporter) CloseConnToCollector() { _ = "STUB: not implemented"; return }

func (e *ipfixExporter) sendTemplateSet(isIPv6 bool) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Get all elements from template record.

func (e *ipfixExporter) addConnToSet(conn *connection.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// Iterate over all infoElements in the list

// Add nodeName only for local Pods whose Pod names are resolved.

// Add nodeName only for local Pods whose Pod names are resolved.

// Sending dummy IP as IPFIX collector expects constant length of data for IP field.
// We should probably think of better approach as this involves customization of IPFIX collector to ignore
// this dummy IP address.

// Same as destinationClusterIPv4.

func (e *ipfixExporter) sendDataSet() (int, error) { _ = "STUB: not implemented"; return 0, nil }
