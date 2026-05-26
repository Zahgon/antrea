// Copyright 2022 Antrea Authors
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

package options

import (
	"time"

	flowaggregatorconfig "antrea.io/antrea/v2/pkg/config/flowaggregator"
)

type Options struct {
	// The configuration object
	Config *flowaggregatorconfig.FlowAggregatorConfig
	// Mode is the mode in which to run the flow aggregator (with aggregation or just as a proxy)
	AggregatorMode flowaggregatorconfig.AggregatorMode
	// Expiration timeout for active flow records in the flow aggregator
	ActiveFlowRecordTimeout time.Duration
	// Expiration timeout for inactive flow records in the flow aggregator
	InactiveFlowRecordTimeout time.Duration
	// Transport protocol over which the aggregator collects IPFIX records from all Agents
	AggregatorTransportProtocol flowaggregatorconfig.AggregatorTransportProtocol
	// IPFIX flow collector address
	ExternalFlowCollectorAddr string
	// IPFIX flow collector transport protocol
	ExternalFlowCollectorProto string
	//  Template retransmission interval when using the UDP protocol to export records.
	TemplateRefreshTimeout time.Duration
	// clickHouseCommitInterval flow records batch commit interval to clickhouse in the flow aggregator
	ClickHouseCommitInterval time.Duration
	// Flow records batch upload interval from flow aggregator to S3 bucket
	S3UploadInterval time.Duration
}

func LoadConfig(configBytes []byte) (*Options, error) { _ = "STUB: not implemented"; return nil, nil }

// Validate all the required options.

// Validate common parameters

// Validate flow collector specific parameters

// Validate clickhouse specific parameters

// Validate S3Uploader specific parameters

// Validate FlowLogger specific parameters
