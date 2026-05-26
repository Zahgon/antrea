// Copyright 2025 Antrea Authors
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

package collector

import (
	ipfixcollector "github.com/vmware/go-ipfix/pkg/collector"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
	flowaggregatorconfig "antrea.io/antrea/v2/pkg/config/flowaggregator"
)

const (
	udpTransport          = "udp"
	tcpTransport          = "tcp"
	ipfixCollectorAddress = "0.0.0.0:4739"
)

type ipfixCollector struct {
	collectingProcess *ipfixcollector.CollectingProcess
	preprocessor      *preprocessor
}

func NewIPFIXCollector(
	recordCh chan *flowpb.Flow,
	aggregatorTransportProtocol flowaggregatorconfig.AggregatorTransportProtocol,
	caCert, serverKey, serverCert []byte,
) (*ipfixCollector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use default value from go-ipfix library

// use default value from go-ipfix library

// use default value from go-ipfix library

// Tell the collector to accept IEs which are not part of the IPFIX registry (hardcoded in
// the go-ipfix library). The preprocessor will take care of removing these elements.

func (c *ipfixCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// blocking function, will return when c.collectingProcess.Stop() is called

func (c *ipfixCollector) GetNumRecordsReceived() int64 { _ = "STUB: not implemented"; return 0 }

func (c *ipfixCollector) GetNumConnsToCollector() int64 { _ = "STUB: not implemented"; return 0 }
