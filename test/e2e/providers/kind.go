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

package providers

import (
	"flag"
	"path"
)

var kindKubeconfigPath = flag.String("kind.kubeconfig", path.Join(homedir, ".kube", "config"), "Path of the kubeconfig of the cluster")

type KindProvider struct {
	controlPlaneNodeName string
}

func (provider *KindProvider) RunCommandOnControlPlaneNode(cmd string) (
	code int, stdout string, stderr string, err error,
) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func (provider *KindProvider) RunCommandOnNode(nodeName string, cmd string) (
	code int, stdout string, stderr string, err error,
) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

func (provider *KindProvider) RunCommandOnNodeExt(nodeName, cmd string, envs map[string]string, stdin string, sudo bool) (
	code int, stdout string, stderr string, err error,
) {
	_ = "STUB: not implemented"
	// sudo is not needed for Docker exec, so ignore the argument.
	return 0, "", "", nil
}

func (provider *KindProvider) GetKubeconfigPath() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// enableKubectlOnControlPlane copies the Kubeconfig file on the Kind control-plane / control-plane Node to the
// default location, in order to make sure that we can run kubectl on the Node.
func (provider *KindProvider) enableKubectlOnControlPlane() error {
	_ = "STUB: not implemented"
	return nil
}

// NewKindProvider returns an implementation of ProviderInterface which is suitable for a
// Kubernetes test cluster created with Kind.
// configPath is unused for the kind provider
func NewKindProvider(configPath string) (ProviderInterface, error) {
	_ = "STUB: not implemented"
	return *new(ProviderInterface), nil
}

// Run docker ps to fetch control-plane Node name
