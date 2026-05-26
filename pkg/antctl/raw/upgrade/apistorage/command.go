// Copyright 2023 Antrea Authors
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

package apistorage

import (
	"io"
	"strings"

	"github.com/spf13/cobra"
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var example = strings.Trim(`
  Perform a dry-run to upgrade all existing objects of Antrea CRDs to the storage API version
  $ antctl upgrade api-storage --dry-run

  Upgrade all existing objects of Antrea CRDs to the storage version
  $ antctl upgrade api-storage

  Upgrade existing AntreaAgentInfo objects to the storage version
  $ antctl upgrade api-storage --crds=antreaagentinfos.crd.antrea.io

  Upgrade existing Egress and Group objects to the storage version
  $ antctl upgrade api-storage --crds=egresses.crd.antrea.io,groups.crd.antrea.io
`, "\n")

type options struct {
	k8sClient client.Client
	crdNames  []string
	dryRun    bool
}

var opts *options

func NewCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func (o *options) complete(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

func runE(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

// getAntreaCRDNames gets names of all Antrea CRDs.
func getAntreaCRDNames(k8sClient client.Client) (sets.Set[string], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getCRDNamesToUpgrade gets names of Antrea CRDs to upgrade.
func getCRDNamesToUpgrade(writer io.Writer, k8sClient client.Client, crdNamesToUpgrade sets.Set[string]) (sets.Set[string], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the user-provided name list of CRDs to upgrade is empty, upgrade all Antrea CRDs.

// If the user-provided name list of CRDs to upgrade is not empty, and it contains CRDs without suffix
// "crd.antrea.io", skip these CRDs.

// Only upgrade the CRDs with suffix "crd.antrea.io".

// getCRDsToUpgrade gets a list of Antrea CRDs to upgrade.
func getCRDsToUpgrade(writer io.Writer, k8sClient client.Client, crdNamesToUpgrade sets.Set[string]) ([]*apiextv1.CustomResourceDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip the CRD that has only one version.

// Skip the CRD that all stored objects are in the storage version.

// upgradeCRDObjects upgrades the existing objects of a CRD.
func upgradeCRDObjects(writer io.Writer, k8sClient client.Client, crd *apiextv1.CustomResourceDefinition) error {
	_ = "STUB: not implemented"
	return nil
}

func upgradeCRDObject(writer io.Writer, k8sClient client.Client, crd *apiextv1.CustomResourceDefinition, obj unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

// If there is a conflict error, update pointer "objToUpdate" to retry.

// updateCRDStoredVersions updates status.storedVersion of a CRD.
func updateCRDStoredVersions(k8sClient client.Client, crd *apiextv1.CustomResourceDefinition) error {
	_ = "STUB: not implemented"
	return nil
}

func getCRDStorageVersion(crd *apiextv1.CustomResourceDefinition) string {
	_ = "STUB: not implemented"
	return ""
}

func newUnexpectedChangeError(crd *apiextv1.CustomResourceDefinition) error {
	_ = "STUB: not implemented"
	return nil
}
