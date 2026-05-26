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

package externalippool

import (
	admv1 "k8s.io/api/admission/v1"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	"antrea.io/antrea/v2/pkg/controller/validation"
)

func (c *ExternalIPPoolController) ValidateExternalIPPool(review *admv1.AdmissionReview) *admv1.AdmissionResponse {
	_ = "STUB: not implemented"
	return nil
}

// Fixed error message to be consistent with IPPool controller

// This shouldn't happen with the webhook configuration we include in the Antrea YAML manifests.

// Always allow DELETE request.

func newAdmissionResponseForErr(err error) *admv1.AdmissionResponse {
	_ = "STUB: not implemented"
	return nil
}

func validateIPRangesAndSubnetInfoForExternalIPPool(externalIPPool *crdv1beta1.ExternalIPPool, existingExternalIPPools []*crdv1beta1.ExternalIPPool) error {
	_ = "STUB: not implemented"
	return nil
}

func collectExistingRanges(pools []*crdv1beta1.ExternalIPPool, skipPool string) ([]validation.NormalizedIPRange, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateNoOverlappingRanges(currentNormalizedIPRanges []validation.NormalizedIPRange, existingExternalIPPools []*crdv1beta1.ExternalIPPool, externalIPPoolName string) error {
	_ = "STUB: not implemented"
	return nil
}
