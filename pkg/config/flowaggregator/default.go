// Copyright 2022 Antrea Authors.
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

package flowaggregator

import (
	"time"
)

const (
	DefaultExternalFlowCollectorTransport = "tcp"
	DefaultExternalFlowCollectorPort      = "4739"
	DefaultActiveFlowRecordTimeout        = "60s"
	DefaultInactiveFlowRecordTimeout      = "90s"
	DefaultAggregatorTransportProtocol    = "TLS"
	DefaultRecordFormat                   = "IPFIX"
	DefaultTemplateRefreshTimeout         = "600s"
	MinValidIPFIXMsgSize                  = 512
	MaxValidIPFIXMsgSize                  = 65535

	DefaultClickHouseDatabase       = "default"
	DefaultClickHouseCommitInterval = "8s"
	MinClickHouseCommitInterval     = 1 * time.Second
	DefaultClickHouseDatabaseUrl    = "tcp://clickhouse-clickhouse.flow-visibility.svc:9000"

	DefaultS3Region            = "us-west-2"
	DefaultS3RecordFormat      = "CSV"
	DefaultS3MaxRecordsPerFile = 1000000
	DefaultS3UploadInterval    = "60s"
	MinS3CommitInterval        = 1 * time.Second

	DefaultLoggerMaxSize      = 100
	DefaultLoggerMaxBackups   = 3
	DefaultLoggerRecordFormat = "CSV"

	DefaultRecordBufferSize = 8192
)

func SetConfigDefaults(flowAggregatorConf *FlowAggregatorConfig) { _ = "STUB: not implemented"; return }
