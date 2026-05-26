// Copyright 2021 Antrea Authors
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

package webhook

import (
	"net/http"

	admv1 "k8s.io/api/admission/v1"
)

type jsonPatchOperation string

const (
	jsonPatchReplaceOp jsonPatchOperation = "replace"
	// LabelMetadataName is a well known reserved label key used by Antrea to store the resource's name
	// as a label value.
	LabelMetadataName = "antrea.io/metadata.name"
)

// jsonPatch contains necessary info that MutatingWebhook required
type jsonPatch struct {
	// Op represents the operation of this mutation
	Op jsonPatchOperation `json:"op"`
	// Path is a jsonPath to locate the value that need to be mutated
	Path string `json:"path"`
	// Value represents the value which is used in mutation
	Value interface{} `json:"value,omitempty"`
}

func HandleMutationLabels() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// verify the content type is accurate

// mutateResourceLabels mutates the resource labels and inserts Antrea required labels.
func mutateResourceLabels(ar *admv1.AdmissionReview) *admv1.AdmissionResponse {
	_ = "STUB: not implemented"
	return nil
}

// At the moment we only mutate Namespace labels.

// mutateLabels mutates the resource's labels and forcefully inserts the resource's name as a well
// known label to ensure that the label is never modified or removed by CREATE and UPDATE events.
func mutateLabels(op admv1.Operation, l map[string]string, name string) (string, bool, []byte) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// Forcefully stomp the resource's metadata.name value as a label.

// createLabelsReplacePatch generates a serialized patch from the new list of labels.
func createLabelsReplacePatch(l map[string]string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getAdmissionResponseForErr(err error) *admv1.AdmissionResponse {
	_ = "STUB: not implemented"
	return nil
}
