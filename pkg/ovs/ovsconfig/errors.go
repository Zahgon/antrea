// Copyright 2019 Antrea Authors
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

package ovsconfig

type Error interface {
	error
	Timeout() bool   // Is the error a timeout?
	Temporary() bool // Is the error temporary?
}

type TransactionError struct {
	error
	temporary bool
}

func NewTransactionError(err error, temporary bool) *TransactionError {
	_ = "STUB: not implemented"
	return nil
}

func (e *TransactionError) Temporary() bool { _ = "STUB: not implemented"; return false }

func (e *TransactionError) Timeout() bool { _ = "STUB: not implemented"; return false }

type InvalidArgumentsError string

func newInvalidArgumentsError(err string) InvalidArgumentsError {
	_ = "STUB: not implemented"
	return *new(InvalidArgumentsError)
}

func (e InvalidArgumentsError) Error() string { _ = "STUB: not implemented"; return "" }

func (e InvalidArgumentsError) Temporary() bool { _ = "STUB: not implemented"; return false }

func (e InvalidArgumentsError) Timeout() bool { _ = "STUB: not implemented"; return false }
