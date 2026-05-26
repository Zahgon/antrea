// Copyright 2025 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package intermediate

type aggregationWorker interface {
	start()
	stop()
}

type worker[T any] struct {
	id        int
	inputChan <-chan T
	errChan   chan bool
	job       func(T) error
}

func createWorker[T any](id int, inputChan <-chan T, job func(T) error) *worker[T] {
	_ = "STUB: not implemented"
	return nil
}

func (w *worker[T]) start() { _ = "STUB: not implemented"; return }

// inputChan is closed and empty

func (w *worker[T]) stop() { _ = "STUB: not implemented"; return }
