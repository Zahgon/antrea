// Copyright 2025 Antrea Authors
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

package collector

import (
	"github.com/vmware/go-ipfix/pkg/entities"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
)

// preprocessor is in charge of converting data records in IPFIX messages received from the IPFIX
// collector to individual Protobuf messages (one per record). If an IPFIX record has extra fields
// (no corresponding field in Protobuf), these will be discarded. If some fields are missing, the
// default Protobuf field value will be used.
type preprocessor struct {
	inCh  <-chan *entities.Message
	outCh chan<- *flowpb.Flow
}

func newPreprocessor(inCh <-chan *entities.Message, outCh chan<- *flowpb.Flow) (*preprocessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *preprocessor) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (p *preprocessor) processMsg(msg *entities.Message) { _ = "STUB: not implemented"; return }

// This is guaranteed to be a slice of length 4, as the Information
// Element has length 4.

// The IE will be a slice of zeros when this IP is not available,
// but for the protobuf message we prefer using the default value (nil).
