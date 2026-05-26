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

package multicluster

import (
	"strings"

	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const defaultToken = "default-member-token"

type initOptions struct {
	namespace   string
	clusterSet  string
	clusterID   string
	createToken bool
	output      string
	k8sClient   client.Client
}

var initOpts *initOptions

func (o *initOptions) validate(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

var initExample = strings.Trim(`
# Initialize ClusterSet in the given Namespace of the leader cluster.
  $ antctl mc init --clusterset clusterset1 --clusterid cluster-north -n antrea-multicluster
# Initialize ClusterSet of the leader cluster and save the join config to a file.
  $ antctl mc init --clusterset clusterset1 --clusterid cluster-north -n antrea-multicluster -j join-config.yml
# Initialize ClusterSet with a default member token, and save the join config as well as the token Secret to a file.
  $ antctl mc init --clusterset clusterset1 --clusterid cluster-north --create-token -n antrea-multicluster -j join-config.yml
`, "\n")

func NewInitCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func initRunE(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Declare ClusterSet init succeeded, even if there is a failure later when creating the
// member token or writing the join config file.
