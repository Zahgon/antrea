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

package raw

import (
	"context"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	antrea "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	"antrea.io/antrea/v2/pkg/util/ip"
)

func GetNodeAddrs(node *corev1.Node) (*ip.DualStackIPs, error) {
	_ = "STUB: not implemented"
	// We prioritize the external Node IP to support cases where antctl is run outside of the
	// cluster, and the internal Node IP may not be reachable.
	return nil, nil
}

func SetupClients(kubeconfig *rest.Config) (*kubernetes.Clientset, *antrea.Clientset, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ResolveKubeconfig(cmd *cobra.Command) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetupLocalKubeconfig(kubeconfig *rest.Config) {
	_ = "STUB: not implemented"

	// We want to avoid accidental uses of this function
	return
}

// TODO: generate kubeconfig in Antrea agent for antctl in-Pod access.

func GetControllerCACert(ctx context.Context, client kubernetes.Interface, controllerInfo *v1beta1.AntreaControllerInfo) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateAgentClientCfgFromObjects(
	ctx context.Context,
	k8sClientset kubernetes.Interface,
	kubeconfig *rest.Config,
	node *corev1.Node,
	agentInfo *v1beta1.AntreaAgentInfo,
	insecure bool,
) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// v1.13 is when APICABundle was added to the AntreaAgentInfo CRD

// The self-signed Agent certificate is only valid for localhost / 127.0.0.1

func CreateAgentClientCfg(
	ctx context.Context,
	k8sClientset kubernetes.Interface,
	antreaClientset antrea.Interface,
	kubeconfig *rest.Config,
	nodeName string,
	insecure bool,
) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateControllerClientCfg(
	ctx context.Context,
	k8sClientset kubernetes.Interface,
	antreaClientset antrea.Interface,
	kubeconfig *rest.Config,
	insecure bool,
) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExecInPod(ctx context.Context, client kubernetes.Interface, config *rest.Config, namespace, pod, container string, command []string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type PodFileCopier interface {
	CopyFromPod(ctx context.Context, fs afero.Fs, namespace, name, containerName, srcPath, dstDir string) error
}

type podFile struct {
	RestConfig *rest.Config
	Client     kubernetes.Interface
}

func NewPodFileCopier(restConfig *rest.Config, client kubernetes.Interface) *podFile {
	_ = "STUB: not implemented"
	return nil
}

func (p *podFile) CopyFromPod(ctx context.Context, fs afero.Fs, namespace, name, containerName, srcPath, dstDir string) error {
	_ = "STUB: not implemented"
	return nil
}
