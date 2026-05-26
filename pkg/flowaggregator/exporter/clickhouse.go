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

package exporter

import (
	"context"
	"time"

	"github.com/google/uuid"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
	"antrea.io/antrea/v2/pkg/flowaggregator/clickhouseclient"
	"antrea.io/antrea/v2/pkg/flowaggregator/options"
	"antrea.io/antrea/v2/pkg/flowaggregator/ringbuffer"
)

type ClickHouseExporter struct {
	chConfig        *clickhouseclient.ClickHouseConfig
	chExportProcess *clickhouseclient.ClickHouseExportProcess
}

const (
	CACertFile      = "ca.crt"
	CertDir         = "/etc/flow-aggregator/certs/clickhouse"
	DefaultInterval = 1 * time.Second
	Timeout         = 1 * time.Minute
)

func buildClickHouseConfig(opt *options.Options) clickhouseclient.ClickHouseConfig {
	_ = "STUB: not implemented"
	return *new(clickhouseclient.ClickHouseConfig)
}

func NewClickHouseExporter(clusterUUID uuid.UUID, opt *options.Options) (*ClickHouseExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run consumes flow records from the ring buffer and writes them to ClickHouse.
// It blocks until ctx is cancelled or the consumer signals shutdown.
func (e *ClickHouseExporter) Run(ctx context.Context, buf ringbuffer.BroadcastBuffer[*flowpb.Flow]) {
	_ = "STUB: not implemented"
	return
}
