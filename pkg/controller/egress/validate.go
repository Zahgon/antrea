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

package egress

import (
	admv1 "k8s.io/api/admission/v1"
)

func (c *EgressController) ValidateEgress(review *admv1.AdmissionReview) *admv1.AdmissionResponse {
	_ = "STUB: not implemented"
	return nil
}

// Validate Egress trafficShaping

// Allow it if EgressIP and ExternalIPPool don't change.

// Only validate whether the specified Egress IP is in the Pool when they are both set.

// This shouldn't happen with the webhook configuration we include in the Antrea YAML manifests.

// Always allow DELETE request.

func newAdmissionResponseForErr(err error) *admv1.AdmissionResponse {
	_ = "STUB: not implemented"
	return nil
}
