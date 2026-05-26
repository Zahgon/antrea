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

type tokenOptions struct {
	namespace     string
	outputFormat  string
	allNamespaces bool
	k8sClient     client.Client
}

var optionsToken *tokenOptions

var tokenExamples = strings.Trim(`
# Get all member tokens in the specified Namespace
  $ antctl mc get membertoken -n antrea-multicluster
# Get all member tokens in all Namespaces
  $ antctl mc get membertoken -A
# Get the specified member token
  $ antctl mc get membertoken cluster-east-token -n antrea-multicluster
# Get the default member token and print the token Secret in YAML format
  $ antctl mc get membertoken default-member-token -n antrea-multicluster -o yaml
# Save the token Secret manifest to a file (which can be used with "antctl mc join" command)
  $ antctl mc get membertoken cluster-east-token -n antrea-multicluster -o yaml > token.yml
`, "\n")

func (o *tokenOptions) validateAndComplete(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMemberTokenCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runEToken(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Ignore tokens not created by antctl mc command.

// ConvertMemberTokenSecret() does not set Namespace of the Secret.
