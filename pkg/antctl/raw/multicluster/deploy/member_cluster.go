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

package deploy

import (
	"strings"

	"github.com/spf13/cobra"
)

type memberClusterOptions struct {
	namespace     string
	antreaVersion string
	filename      string
}

var memberClusterOpts *memberClusterOptions

var memberClusterExamples = strings.Trim(`
# Deploy Antrea Multi-cluster of the specified version into a Namespace
  $ antctl mc deploy membercluster --antrea-version <ANTREA_VERSION> -n <NAMESPACE>
# Deploy Antrea Multi-cluster using a pre-downloaded manifest
  $ antctl mc deploy membercluster -f <PATH_TO_MANIFEST>

The following CRDs will be defined:
- CRDs: ClusterSet, MemberClusterAnnounce, ResourceExport, ResourceImport, ServiceExport, ServiceImport
`, "\n")

func (o *memberClusterOptions) validateAndComplete() error { _ = "STUB: not implemented"; return nil }

func NewMemberClusterCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func memberClusterRunE(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }
