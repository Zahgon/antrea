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

package flowlogger

import (
	"bufio"
	"io"
	"sync"
	"time"

	"antrea.io/antrea/v2/pkg/flowaggregator/flowrecord"
)

const MaxLatency = 5 * time.Second

type FlowLogger struct {
	sync.Mutex
	logger     io.Closer
	maxLatency time.Duration
	writer     *bufio.Writer
}

func NewFlowLogger(path string, maxSize int, maxBackups int, maxAge int, compress bool) *FlowLogger {
	_ = "STUB: not implemented"
	return nil
}

func (fl *FlowLogger) FlushLoop(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (fl *FlowLogger) Close() { _ = "STUB: not implemented"; return }

func (fl *FlowLogger) WriteRecord(r *flowrecord.FlowRecord, prettyPrint bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (fl *FlowLogger) Flush() error { _ = "STUB: not implemented"; return nil }
