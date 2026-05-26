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
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/pflag"
)

const (
	logToStdErrFlag = "logtostderr"
	logDirFlag      = "log_dir"
	logFileFlag     = "log_file"
	maxSizeFlag     = "log_file_max_size"
	maxNumFlag      = "log_file_max_num"

	// Check log file number every 10 mins.
	logFileCheckInterval = time.Minute * 10
	// Allowed maximum value for the maximum file size limit.
	maxMaxSizeMB = 1024 * 100
)

var (
	maxNumArg     = uint16(0)
	logFileMaxNum = uint16(0)
	logDir        = ""

	executableName = filepath.Base(os.Args[0])
)

// initLogFileLimits initializes log file maximum size and maximum number limits based on the
// command line flags.
func initLogFileLimits(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Should not happen. Return for safety.

// Logging to files is not enabled.

// Log to a single file. klog will take care of the max size limit.

// Max log file size in MB.

// klog does not respect the max file size specified by --log_file_max_size
// when --log_file is not used. Here as a workaround, we directly set the
// specified max size to klog.MaxSize.

// Log to the tmp dir.

// StartLogFileNumberMonitor starts monitoring the log files to make sure the
// number of log files does not exceed the maximum limit, when the log file
// number limit is configured.
func StartLogFileNumberMonitor(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// The maximum log file number limit is not configured.

func checkLogFiles() { _ = "STUB: not implemented"; return }

// Skip dir, symbol link, etc.

// Sort files by modification time.

// Remove the oldest files.

// #nosec G703: Path provided via config by admin; no privilege boundary crossed.
