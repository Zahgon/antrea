// Copyright 2023 Antrea Authors
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

package networkpolicy

import (
	"github.com/spf13/afero"
	"k8s.io/apimachinery/pkg/runtime"
)

// fileStore encodes and stores runtime.Objects in files. Each object will be stored in a separate file under the given
// directory.
type fileStore struct {
	fs afero.Fs
	// The directory to store the files.
	dir string
	// serializer knows how to encode and decode the objects.
	serializer runtime.Serializer
}

func newFileStore(fs afero.Fs, dir string, serializer runtime.Serializer) (*fileStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// save stores the given object in file with the object's UID as the file name, overwriting any existing content if the
// file already exists. Note the method may update the object's GroupVersionKind in-place during serialization.
func (s fileStore) save(item runtime.Object) error { _ = "STUB: not implemented"; return nil }

// Encode may update the object's GroupVersionKind in-place during serialization.

// delete removes the file with the object's UID as the file name if it exists.
func (s fileStore) delete(item runtime.Object) error { _ = "STUB: not implemented"; return nil }

// replaceAll replaces all files under the directory with the given objects. Existing files not in the given objects
// will be removed. Note the method may update the object's GroupVersionKind in-place during serialization.
func (s fileStore) replaceAll(items []runtime.Object) error { _ = "STUB: not implemented"; return nil }

func (s fileStore) loadAll() ([]runtime.Object, error) { _ = "STUB: not implemented"; return nil, nil }

// If the data is corrupted somehow, we still want to load other data and continue the process.

// Note: we haven't stored a different version so far but version conversion should be performed when the used
// version is upgraded in the future.
