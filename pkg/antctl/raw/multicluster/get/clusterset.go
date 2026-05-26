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

package get

import (
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type clusterSetOptions struct {
	namespace     string
	outputFormat  string
	allNamespaces bool
	k8sClient     client.Client
}

var optionsClusterSet *clusterSetOptions

var clusterSetExamples = strings.Trim(`
Gel all ClusterSets in the default Namesapce
$ antctl mc get clusterset
Get all ClusterSets in all Namespaces
$ antctl mc get clusterset -A
Get all ClusterSets in the specified Namespace
$ antctl mc get clusterset -n <NAMESPACE>
Get all ClusterSets and print them in JSON format
$ antctl mc get clusterset -o json
Get the specified ClusterSet
$ antctl mc get clusterset <CLUSTERSET_ID>
`, "\n")

func (o *clusterSetOptions) validateAndComplete(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterSetCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runEClusterSet(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }
