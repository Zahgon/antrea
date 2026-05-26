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

type joinConfigOptions struct {
	namespace   string
	memberToken string
	k8sClient   client.Client
}

var joinConfigOpts *joinConfigOptions

var joinConfigExamples = strings.Trim(`
Print member join parameters of the ClusterSet in the antrea-multicluster Namespace
$ antctl mc get joinconfig -n antrea-multicluster
Print member join parameters and the Secret manifest of the default member token
$ antctl mc get joinconfig --member-token default-member-token -n antrea-multicluster
Print member join parameters and the Secret manifest of the specified member token
$ antctl mc get joinconfig --member-token cluster-east-token -n antrea-multicluster
`, "\n")

func (o *joinConfigOptions) validateAndComplete(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

// For unit test.

func NewJoinConfigCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runEJoinConfig(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }
