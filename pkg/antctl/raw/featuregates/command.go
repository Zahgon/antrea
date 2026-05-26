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

package featuregates

import (
	"context"
	"io"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"antrea.io/antrea/v2/pkg/antctl/runtime"
	"antrea.io/antrea/v2/pkg/apiserver/apis"
	antrea "antrea.io/antrea/v2/pkg/client/clientset/versioned"
)

var Command *cobra.Command
var getClients = getConfigAndClients
var getRestClient = getRestClientByMode

var option = &struct {
	insecure bool
}{}

func init() {
	Command = &cobra.Command{
		Use:   "featuregates",
		Short: "Print Antrea feature gates",
	}
	if runtime.Mode == runtime.ModeAgent {
		Command.RunE = agentRunE
		Command.Long = "Print current Antrea agent feature gates info"
	} else if runtime.Mode == runtime.ModeController && runtime.InPod {
		Command.RunE = controllerLocalRunE
		Command.Long = "Print Antrea feature gates info including Controller and Agent"
	} else if runtime.Mode == runtime.ModeController && !runtime.InPod {
		Command.Long = "Print Antrea feature gates info including Controller and Agent"
		Command.Flags().BoolVar(&option.insecure, "insecure", false, "Skip TLS verification when connecting to Antrea API.")
		Command.RunE = controllerRemoteRunE
	}
}

func agentRunE(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func controllerLocalRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func controllerRemoteRunE(cmd *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	return nil
}

func featureGateRequest(cmd *cobra.Command, mode string) error {
	_ = "STUB: not implemented"
	return nil
}

func getConfigAndClients(cmd *cobra.Command) (*rest.Config, kubernetes.Interface, antrea.Interface, error) {
	_ = "STUB: not implemented"
	return nil, *new(kubernetes.Interface), *new(antrea.Interface), nil
}

func getRestClientByMode(ctx context.Context, kubeconfig *rest.Config, k8sClientset kubernetes.Interface, antreaClientset antrea.Interface, mode string) (*rest.RESTClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getControllerClient(ctx context.Context, k8sClientset kubernetes.Interface, antreaClientset antrea.Interface, kubeconfig *rest.Config, insecure bool) (*rest.RESTClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFeatureGatesRequest(client *rest.RESTClient) ([]apis.FeatureGateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func output(resps []apis.FeatureGateResponse, component string, output io.Writer) {
	_ = "STUB: not implemented"
	return
}
