// Copyright 2024 Antrea Authors
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

package e2e

import (
	"context"

	"golang.org/x/crypto/ssh"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

var sftpLabels = map[string]string{"app": "sftp"}

const (
	sftpUser      = "foo"
	sftpPassword  = "pass"
	sftpUploadDir = "upload"
)

func genSFTPService(nodePort int32) *v1.Service { _ = "STUB: not implemented"; return nil }

func genSSHKeysSecret(ed25519Key, rsaKey []byte) *v1.Secret { _ = "STUB: not implemented"; return nil }

func genSFTPDeployment() *appsv1.Deployment { _ = "STUB: not implemented"; return nil }

func (data *TestData) deploySFTPServer(ctx context.Context, nodePort int32) (*appsv1.Deployment, *v1.Service, []ssh.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}
