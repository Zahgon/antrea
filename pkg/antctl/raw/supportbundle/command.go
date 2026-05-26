// Copyright 2020 Antrea Authors
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

package supportbundle

import (
	"context"
	"io"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"antrea.io/antrea/v2/pkg/antctl/runtime"
	antrea "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	systemclientset "antrea.io/antrea/v2/pkg/client/clientset/versioned/typed/system/v1beta1"
)

const (
	barTmpl pb.ProgressBarTemplate = `{{string . "prefix"}}{{bar . }} {{percent . }} {{rtime . "ETA %s"}}` // Example: 'Prefix[-->______] 20%'

	requestRate  = 50
	requestBurst = 100
	timeFormat   = "20060102T150405Z0700"
)

// Command is the support bundle command implementation.
var Command *cobra.Command

var option = &struct {
	dir            string
	labelSelector  string
	controllerOnly bool
	nodeListFile   string
	since          string
	insecure       bool
}{}

var defaultFS = afero.NewOsFs()

var remoteControllerLongDescription = strings.TrimSpace(`
Generate support bundles for the cluster, which include: information about each Antrea agent, information about the Antrea controller and general information about the cluster.
`)

var remoteControllerExample = strings.Trim(`
  Generate support bundles of the controller and agents on all Nodes and save them to current working dir
  $ antctl supportbundle
  Generate support bundle of the controller
  $ antctl supportbundle --controller-only
  Generate support bundle of the controller and agents on all Nodes with only the logs generated during the last 1 hour
  $ antctl supportbundle --since 1h
  Generate support bundles of agents on specific Nodes filtered by name list, no wildcard support
  $ antctl supportbundle node_a node_b node_c
  Generate support bundles of agents on specific Nodes filtered by names in a file (one Node name per line)
  $ antctl supportbundle -f ~/nodelistfile
  Generate support bundles of agents on specific Nodes filtered by name, with support for wildcard expressions
  $ antctl supportbundle '*worker*'
  Generate support bundles of agents on specific Nodes filtered by name and label selectors
  $ antctl supportbundle '*worker*' -l kubernetes.io/os=linux
  Generate support bundles of the controller and agents on all Nodes and save them to specific dir
  $ antctl supportbundle -d ~/Downloads
`, "\n")

func init() {
	Command = &cobra.Command{
		Use:   "supportbundle",
		Short: "Generate support bundle",
	}

	if runtime.Mode == runtime.ModeAgent {
		Command.RunE = agentRunE
		Command.Long = "Generate the support bundle of current Antrea agent."
	} else if runtime.Mode == runtime.ModeController && runtime.InPod {
		Command.RunE = controllerLocalRunE
		Command.Long = "Generate the support bundle of current Antrea controller."
	} else if runtime.Mode == runtime.ModeController && !runtime.InPod {
		Command.Use += " [nodeName...]"
		Command.Long = remoteControllerLongDescription
		Command.Example = remoteControllerExample
		Command.Flags().StringVarP(&option.dir, "dir", "d", "", "support bundles output dir, the path will be created if it doesn't exist")
		Command.Flags().StringVarP(&option.labelSelector, "label-selector", "l", "", "selector (label query) to filter Nodes for agent bundles, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
		Command.Flags().BoolVar(&option.controllerOnly, "controller-only", false, "only collect the support bundle of Antrea controller")
		Command.Flags().StringVarP(&option.nodeListFile, "node-list-file", "f", "", "only collect the support bundle of specific nodes filtered by names in a file (one node name per line)")
		Command.Flags().StringVarP(&option.since, "since", "", "", "only return logs newer than a relative duration like 5s, 2m or 3h. Defaults to all logs")
		Command.Flags().BoolVar(&option.insecure, "insecure", false, "Skip TLS verification when connecting to Antrea API.")
		Command.RunE = controllerRemoteRunE
	}
}

var getSupportBundleClient func() (systemclientset.SupportBundleInterface, error) = setupSupportBundleClient

func setupSupportBundleClient() (systemclientset.SupportBundleInterface, error) {
	_ = "STUB: not implemented"
	return *new(systemclientset.SupportBundleInterface), nil
}

func localSupportBundleRequest(cmd *cobra.Command, mode string, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// will expire after 100ms

// retry again after 500ms

func agentRunE(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func controllerLocalRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func request(ctx context.Context, component string, client systemclientset.SupportBundleInterface) error {
	_ = "STUB: not implemented"
	return nil
}

type result struct {
	nodeName string
	err      error
}

func mapClients(
	ctx context.Context,
	prefix string,
	agentClients map[string]systemclientset.SupportBundleInterface,
	controllerClient systemclientset.SupportBundleInterface,
	bar *pb.ProgressBar,
	af, cf func(ctx context.Context, nodeName string, c systemclientset.SupportBundleInterface) error,
) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

func requestAll(
	ctx context.Context,
	agentClients map[string]systemclientset.SupportBundleInterface,
	controllerClient systemclientset.SupportBundleInterface,
	bar *pb.ProgressBar,
) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

func download(
	ctx context.Context,
	suffix,
	downloadPath string,
	client systemclientset.SupportBundleInterface,
	component string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// will expire after 100ms

// retry again after 500ms

func writeFailedNodes(downloadPath string, nodes []string) error {
	_ = "STUB: not implemented"
	return nil
}

// downloadAll will download all supportBundles. preResults is the request results of node/controller supportBundle.
// if err happens for some nodes or controller, the download step will be skipped for the failed nodes or the controller.
func downloadAll(
	ctx context.Context,
	agentClients map[string]systemclientset.SupportBundleInterface,
	controllerClient systemclientset.SupportBundleInterface,
	downloadPath string,
	bar *pb.ProgressBar,
	preResults map[string]error,
) map[string]error {
	_ = "STUB: not implemented"
	return nil
}

// createAgentClients creates clients for agents on specified nodes. If nameList is set, then nameFilter will be ignored.
func createAgentClients(
	ctx context.Context,
	k8sClientset kubernetes.Interface,
	antreaClientset antrea.Interface,
	kubeconfig *rest.Config,
	nameFilter string,
	nameList []string,
	insecure bool,
) (map[string]systemclientset.SupportBundleInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createControllerClient(
	ctx context.Context,
	k8sClientset kubernetes.Interface,
	antreaClientset antrea.Interface,
	cfgTmpl *rest.Config,
	insecure bool,
) (systemclientset.SupportBundleInterface, error) {
	_ = "STUB: not implemented"
	return *new(systemclientset.SupportBundleInterface), nil
}

func getClusterInfo(w io.Writer, k8sClient kubernetes.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

// These are the ConfigMaps created by Antrea in the the kube-system Namespace.

func controllerRemoteRunE(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect controller bundle when no Node name or label filter is specified, or
// when --controller-only is set.

func genErrorMsg(resultMap map[string]error) string { _ = "STUB: not implemented"; return "" }

// processResults will output the failed nodes and their reasons if any. If no data was collected,
// error is returned, otherwise will return nil. For failed nodes and controller, will also trying to get logs from
// kubernetes api.
func processResults(ctx context.Context, antreaClientset antrea.Interface, k8sClient kubernetes.Interface, resultMap map[string]error, dir string) error {
	_ = "STUB: not implemented"
	return nil
}

// download logs from kubernetes api

func downloadFallbackControllerBundleFromKubernetes(ctx context.Context, antreaClientset antrea.Interface, k8sClient kubernetes.Interface, dir string) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadFallbackAgentBundleFromKubernetes(ctx context.Context, antreaClientset antrea.Interface, k8sClient kubernetes.Interface, failedNodes []string, dir string) error {
	_ = "STUB: not implemented"
	return nil
}

func packPodBundle(pod *corev1.Pod, dir string, bundleDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func downloadPodLogs(ctx context.Context, k8sClient kubernetes.Interface, namespace string, podName string, containers []string, dir string) error {
	_ = "STUB: not implemented"
	return nil
}
