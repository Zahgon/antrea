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

type resourceExportOptions struct {
	namespace     string
	outputFormat  string
	allNamespaces bool
	clusterID     string
	k8sClient     client.Client
}

var optionsResourceExport *resourceExportOptions

var resourceExportExamples = strings.Trim(`
Get all ResourceExports of ClusterSet in default Namesapce
$ antctl mc get resourceexport
Get all ResourceExports of ClusterSet in all Namespaces
$ antctl mc get resourceexport -A
Get all ResourceExports in the specified Namespace
$ antctl mc get resourceexport -n <NAMESPACE>
Get all ResourceExports and print them in JSON format
$ antctl mc get resourceexport -o json
Get the specified ResourceExport
$ antctl mc get resourceexport <RESOURCEEXPORT> -n <NAMESPACE>
`, "\n")

func (o *resourceExportOptions) validateAndComplete(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func NewResourceExportCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runEResourceExport(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
