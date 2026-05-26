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

package create

import (
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type memberTokenOptions struct {
	namespace string
	output    string
	k8sClient client.Client
}

var memberTokenOpts *memberTokenOptions

var memberTokenExamples = strings.Trim(`
# Create a member token in the antrea-multicluster Namespace
  $ antctl mc create membertoken cluster-east-token -n antrea-multicluster
# Create a member token and save the Secret manifest to a file
  $ antctl mc create membertoken cluster-east-token -n antrea-multicluster -o token-secret.yml
`, "\n")

func (o *memberTokenOptions) validateAndComplete(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMemberTokenCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func memberTokenRunE(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
