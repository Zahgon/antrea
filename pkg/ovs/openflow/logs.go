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

package openflow

import (
	"github.com/sirupsen/logrus"
)

type klogFormatter struct{}

// Format formats logrus log in compliance with k8s log.
func (f *klogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// logrus has seven logging levels: Trace, Debug, Info, Warning, Error, Fatal and Panic.

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&klogFormatter{})
}
