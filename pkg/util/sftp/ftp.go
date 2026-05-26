// Copyright 2024 Antrea Authors.
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

package sftp

import (
	"io"
	"net/url"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	uploadToFileServerMaxRetries = 5
	uploadToFileServerRetryDelay = 5 * time.Second
)

func ParseSFTPUploadUrl(uploadUrl string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Uploader interface {
	// Upload uploads a file to the target sftp address using ssh config.
	Upload(url string, fileName string, config *ssh.ClientConfig, outputFile io.Reader) error
}

type sftpUploader struct {
}

func NewUploader() Uploader { _ = "STUB: not implemented"; return *new(Uploader) }

func (uploader *sftpUploader) Upload(url string, fileName string, config *ssh.ClientConfig, outputFile io.Reader) error {
	_ = "STUB: not implemented"
	// url should be like: 10.92.23.154:22/path or sftp://10.92.23.154:22/path
	return nil
}

func upload(address string, path string, config *ssh.ClientConfig, file io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}
