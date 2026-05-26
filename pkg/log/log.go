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

// Package log processes the klog flags, and enforces the maximum log file
// size and maximum log file number limits.
package log

import (
	"flag"
	"time"

	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

const (
	logFlushFreqFlag = "log-flush-frequency"
)

var (
	klogFlags = flag.NewFlagSet("logging", flag.ContinueOnError)

	logFlushFreq time.Duration
)

func init() {
	klog.InitFlags(klogFlags)
}

func addKlogFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func AddFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

type Options struct {
	withFlushDaemon bool
}

func WithoutFlushDaemon(options *Options) { _ = "STUB: not implemented"; return }

func initKlog(options *Options) { _ = "STUB: not implemented"; return }

func InitLogs(fs *pflag.FlagSet, opts ...func(options *Options)) { _ = "STUB: not implemented"; return }

func FlushLogs() { _ = "STUB: not implemented"; return }
