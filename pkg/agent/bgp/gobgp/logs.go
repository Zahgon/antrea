// Copyright 2024 Antrea Authors
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

package gobgp

import (
	gobgplog "github.com/osrg/gobgp/v3/pkg/log"
)

// goBGPLogger implements https://github.com/osrg/gobgp/blob/master/pkg/log/logger.go interface.
type goBGPLogger struct {
	routerID string
}

func newGoBGPLogger(routerID string) *goBGPLogger { _ = "STUB: not implemented"; return nil }

// We use a depth of 1 for all log messages to get more useful source information.
// Otherwise, the reported source location would be the line where the klog function is invoked.

func (g *goBGPLogger) Panic(msg string, fields gobgplog.Fields) { _ = "STUB: not implemented"; return }

func (g *goBGPLogger) Fatal(msg string, fields gobgplog.Fields) { _ = "STUB: not implemented"; return }

func (g *goBGPLogger) Error(msg string, fields gobgplog.Fields) { _ = "STUB: not implemented"; return }

func (g *goBGPLogger) Warn(msg string, fields gobgplog.Fields) { _ = "STUB: not implemented"; return }

func (g *goBGPLogger) Info(msg string, fields gobgplog.Fields) { _ = "STUB: not implemented"; return }

func (g *goBGPLogger) Debug(msg string, fields gobgplog.Fields) { _ = "STUB: not implemented"; return }

// This should never be used in Antrea.
func (g *goBGPLogger) SetLevel(level gobgplog.LogLevel) {
	_ = "STUB: not implemented"

	// This should never be used in Antrea.
	return
}

func (g *goBGPLogger) GetLevel() gobgplog.LogLevel {
	_ = "STUB: not implemented"
	return *new(gobgplog.LogLevel)
}

func (g *goBGPLogger) logFieldsToKeysAndValues(fields gobgplog.Fields) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Add routerID to all log messages for more context.
