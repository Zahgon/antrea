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

func (f *RegField) GetRegID() int { _ = "STUB: not implemented"; return 0 }

func (f *RegField) GetRange() *Range { _ = "STUB: not implemented"; return nil }

func (f *RegField) GetNXFieldName() string { _ = "STUB: not implemented"; return "" }

func (f *RegField) isFullRange() bool { _ = "STUB: not implemented"; return false }

func NewRegField(id int, start, end uint32) *RegField { _ = "STUB: not implemented"; return nil }

func NewOneBitRegMark(id int, bit uint32) *RegMark { _ = "STUB: not implemented"; return nil }

func NewOneBitZeroRegMark(id int, bit uint32) *RegMark { _ = "STUB: not implemented"; return nil }

func NewRegMark(field *RegField, value uint32) *RegMark { _ = "STUB: not implemented"; return nil }

func (m *RegMark) GetValue() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *RegMark) GetField() *RegField { _ = "STUB: not implemented"; return nil }

func (f *XXRegField) GetRegID() int { _ = "STUB: not implemented"; return 0 }

func (f *XXRegField) GetRange() *Range { _ = "STUB: not implemented"; return nil }

func (f *XXRegField) GetNXFieldName() string { _ = "STUB: not implemented"; return "" }

func NewXXRegField(id int, start, end uint32) *XXRegField { _ = "STUB: not implemented"; return nil }

func (m *CtMark) GetRange() *Range {
	_ = "STUB: not implemented"

	// GetValue gets CT mark value with offset since CT mark is used by bit. E.g, CT_MARK_REG[3]==1, the return
	// value of this function is 0b1000.
	return nil
}

func (m *CtMark) GetValue() uint32 { _ = "STUB: not implemented"; return 0 }

func NewCTMarkField(start, end uint32) *CtMarkField { _ = "STUB: not implemented"; return nil }

func NewOneBitCTMark(bit uint32) *CtMark { _ = "STUB: not implemented"; return nil }

func NewOneBitZeroCTMark(bit uint32) *CtMark { _ = "STUB: not implemented"; return nil }

func NewCTMark(field *CtMarkField, value uint32) *CtMark { _ = "STUB: not implemented"; return nil }

func NewCTLabel(start, end uint32) *CtLabel { _ = "STUB: not implemented"; return nil }

func (f *CtLabel) GetNXFieldName() string { _ = "STUB: not implemented"; return "" }

func (f *CtLabel) GetRange() *Range { _ = "STUB: not implemented"; return nil }
