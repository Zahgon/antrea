// Copyright 2020 Antrea Authors
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

package networkpolicy

import (
	// #nosec G505: not used for security purposes

	admv1 "k8s.io/api/admission/v1"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

type NetworkPolicyMutator struct {
	networkPolicyController *NetworkPolicyController
}

// NewNetworkPolicyMutator returns a new *NetworkPolicyMutator.
func NewNetworkPolicyMutator(networkPolicyController *NetworkPolicyController) *NetworkPolicyMutator {
	_ = "STUB: not implemented"
	return nil
}

// Mutate function mutates an Antrea-native policy object
func (m *NetworkPolicyMutator) Mutate(ar *admv1.AdmissionReview) *admv1.AdmissionResponse {
	_ = "STUB: not implemented"
	return nil
}

// For a mutating webhook, patch and patchType must be provided together only when patch exists

// mutateAntreaPolicy mutates names of rules of an Antrea NetworkPolicy CRD.
// If users didn't specify the name of an ingress or egress rule,
// mutateAntreaPolicy will auto-generate a name for this rule. In
// addition to the rule names, it also mutates the Tier field to the default
// tier name if it is unset.
func (m *NetworkPolicyMutator) mutateAntreaPolicy(op admv1.Operation, ingress, egress []crdv1beta1.Rule, tier string) (string, bool, []byte) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// Mutate Antrea-native policy rule names.

// Mutate empty tier name to the name of the Default application Tier.

// Delete of Antrea Policies have no mutation

// generateRuleNames generates unique rule names and returns a list of json paths and the corresponding list of generated names
func generateRuleNames(prefix string, rules []crdv1beta1.Rule) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

type jsonPatchOperation string

const (
	jsonPatchReplaceOp jsonPatchOperation = "replace"
)

// jsonPatch contains necessary info that MutatingWebhook required
type jsonPatch struct {
	// Op represent the operation of this mutation
	Op jsonPatchOperation `json:"op"`
	// Path is a jsonPath to locate the value that need to be mutated
	Path string `json:"path"`
	// Value represent the value which is used in mutation
	Value interface{} `json:"value,omitempty"`
}

// createReplacePatch use paths and values that need to be replace to generate a serialized patch
func createReplacePatch(paths []string, values []string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const ruleNameSuffixLen = 7

// hashRule calculates a string based on the rule's content.
func hashRule(r crdv1beta1.Rule) string {
	_ = "STUB: not implemented"
	// #nosec G401: not used for security purposes
	return ""
}
