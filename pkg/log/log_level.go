// Copyright 2020 Antrea Authors
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

package log

import (
	"github.com/spf13/pflag"
)

const logVerbosityFlag = "v"

type logLevelManager struct {
	// we use this to access the verbosity value at runtime
	flag *pflag.Flag
}

var logLevelMgr = &logLevelManager{}

func (m *logLevelManager) getCurrentLogLevel() string { _ = "STUB: not implemented"; return "" }

func (m *logLevelManager) setLogLevel(level string) error { _ = "STUB: not implemented"; return nil }

func initLogLevelManager(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// GetCurrentLogLevel returns the current log verbosity level.
func GetCurrentLogLevel() string { _ = "STUB: not implemented"; return "" }

// SetLogLevel sets the log verbosity level. level must be a string
// representation of a decimal integer.
func SetLogLevel(level string) error { _ = "STUB: not implemented"; return nil }
