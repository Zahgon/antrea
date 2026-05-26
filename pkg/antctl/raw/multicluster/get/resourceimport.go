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

type resourceImportOptions struct {
	namespace     string
	outputFormat  string
	allNamespaces bool
	k8sClient     client.Client
}

var options *resourceImportOptions

var resourceImportExamples = strings.Trim(`
Gel all ResourceImports of a ClusterSet in default Namespace
$ antctl mc get resourceimport
Get all ResourceImports of a ClusterSet in all Namespaces
$ antctl mc get resourceimport -A
Get all ResourceImports in the specified Namespace
$ antctl mc get resourceimport -n <NAMESPACE>
Get all ResourceImports and print them in JSON format
$ antctl mc get resourceimport -o json
Get the specified ResourceImport
$ antctl mc get resourceimport <RESOURCEIMPORT> -n <NAMESPACE>
`, "\n")

func (o *resourceImportOptions) validateAndComplete(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func NewResourceImportCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runE(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }
