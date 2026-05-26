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

package exporter

import (
	"context"
	"sync"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
	flowaggregatorconfig "antrea.io/antrea/v2/pkg/config/flowaggregator"
	"antrea.io/antrea/v2/pkg/flowaggregator/flowlogger"
	"antrea.io/antrea/v2/pkg/flowaggregator/flowrecord"
	"antrea.io/antrea/v2/pkg/flowaggregator/options"
	"antrea.io/antrea/v2/pkg/flowaggregator/ringbuffer"
)

type flowFilter struct {
	IngressNetworkPolicyRuleActions []uint8
	EgressNetworkPolicyRuleActions  []uint8
}

type LogExporter struct {
	config     flowaggregatorconfig.FlowLoggerConfig
	filters    []flowFilter
	flowLogger *flowlogger.FlowLogger
	stopCh     chan struct{}
	wg         sync.WaitGroup
}

func NewLogExporter(opt *options.Options) (*LogExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *LogExporter) buildFilters() { _ = "STUB: not implemented"; return }

// invalid case

func (e *LogExporter) applyFilters(r *flowrecord.FlowRecord) bool {
	_ = "STUB: not implemented"
	return false
}

// Run consumes flow records from the ring buffer and writes them to a local log file.
// It blocks until ctx is cancelled or the consumer signals shutdown.
func (e *LogExporter) Run(ctx context.Context, buf ringbuffer.BroadcastBuffer[*flowpb.Flow]) {
	_ = "STUB: not implemented"
	return
}
