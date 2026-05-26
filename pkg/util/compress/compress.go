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

package compress

import (
	"io"

	"github.com/spf13/afero"
)

// sanitizeExtractPath ensures that the target extract path (when joining the destination directory
// and the path from the archive) is within the intended destination directory.
// This is meant to address the "Zip Slip" vulnerability (G305).
// See https://security.snyk.io/research/zip-slip-vulnerability.
func sanitizeExtractPath(filePath string, destination string) (string, error) {
	_ = "STUB: not implemented"
	// If IsLocal(path) returns true, then Join(base, path) will always produce a path contained
	// within base and Clean(path) will always produce an unrooted path with no ".." path
	// elements.
	// IsLocal was introduced in Go 1.20.
	// This will also reject absolute paths, which is not strictly required (e.g., tar can
	// produce such archives when it is run with -P).
	return "", nil
}

// Join also calls Clean on the path.

func UnpackDir(fs afero.Fs, fileName string, targetDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func UnpackReader(fs afero.Fs, file io.Reader, useGzip bool, targetDir string) error {
	_ = "STUB: not implemented"
	return nil
}

// to resolve G110: Potential DoS vulnerability via decompression bomb

// Note in particular that we do not handle symlinks.

func PackDir(fs afero.Fs, dir string, writer io.Writer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
