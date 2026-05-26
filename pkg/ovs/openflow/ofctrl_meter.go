// Copyright 2021 Antrea Authors
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
	"antrea.io/libOpenflow/openflow15"
	"antrea.io/ofnet/ofctrl"
)

type ofMeter struct {
	ofctrl *ofctrl.Meter
	bridge *OFBridge
}

func (m *ofMeter) Reset() { _ = "STUB: not implemented"; return }

// Note: use OFSwitch to directly send MeterModification message rather than bundle message is because the
// current ofnet implementation for OpenFlow bundle does not support adding MeterModification.
func (m *ofMeter) Add() error { _ = "STUB: not implemented"; return nil }

func (m *ofMeter) Modify() error { _ = "STUB: not implemented"; return nil }

func (m *ofMeter) Delete() error { _ = "STUB: not implemented"; return nil }

func (m *ofMeter) Type() EntryType { _ = "STUB: not implemented"; return *new(EntryType) }

func (m *ofMeter) GetBundleMessages(entryOper OFOperation) ([]ofctrl.OpenFlowModMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ofMeter) ResetMeterBands() Meter { _ = "STUB: not implemented"; return *new(Meter) }

func (m *ofMeter) MeterBand() MeterBandBuilder {
	_ = "STUB: not implemented"
	return *new(MeterBandBuilder)
}

type meterBandBuilder struct {
	meter           *ofMeter
	meterBandHeader *openflow15.MeterBandHeader
	prevLevel       uint8
	experimenter    uint32
}

func (m *meterBandBuilder) MeterType(meterType ofctrl.MeterType) MeterBandBuilder {
	_ = "STUB: not implemented"
	return *new(MeterBandBuilder)
}

func (m *meterBandBuilder) Rate(rate uint32) MeterBandBuilder {
	_ = "STUB: not implemented"
	return *new(MeterBandBuilder)
}

func (m *meterBandBuilder) Burst(burst uint32) MeterBandBuilder {
	_ = "STUB: not implemented"
	return *new(MeterBandBuilder)
}

func (m *meterBandBuilder) PrecLevel(precLevel uint8) MeterBandBuilder {
	_ = "STUB: not implemented"
	return *new(MeterBandBuilder)
}

func (m *meterBandBuilder) Experimenter(experimenter uint32) MeterBandBuilder {
	_ = "STUB: not implemented"
	return *new(MeterBandBuilder)
}

func (m *meterBandBuilder) Done() Meter { _ = "STUB: not implemented"; return *new(Meter) }
