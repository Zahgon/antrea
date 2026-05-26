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

package certificatesigningrequest

import (
	certificatesv1 "k8s.io/api/certificates/v1"
	clientset "k8s.io/client-go/kubernetes"
	csrlisters "k8s.io/client-go/listers/certificates/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	approvingControllerName = "CertificateSigningRequestApprovingController"
)

type approver interface {
	recognize(csr *certificatesv1.CertificateSigningRequest) bool
	verify(csr *certificatesv1.CertificateSigningRequest) (bool, error)
	name() string
}

// CSRApprovingController is responsible for approving CertificateSigningRequests.
type CSRApprovingController struct {
	client          clientset.Interface
	csrInformer     cache.SharedIndexInformer
	csrLister       csrlisters.CertificateSigningRequestLister
	csrListerSynced cache.InformerSynced
	queue           workqueue.TypedRateLimitingInterface[string]
	approvers       []approver
}

// NewCSRApprovingController returns a new *CSRApprovingController.
func NewCSRApprovingController(client clientset.Interface, csrInformer cache.SharedIndexInformer, csrLister csrlisters.CertificateSigningRequestLister) *CSRApprovingController {
	_ = "STUB: not implemented"
	return nil
}

// Run begins watching and syncing of the CSRApprovingController.
func (c *CSRApprovingController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *CSRApprovingController) worker() { _ = "STUB: not implemented"; return }

func (c *CSRApprovingController) enqueueCertificateSigningRequest(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *CSRApprovingController) syncCSR(key string) error { _ = "STUB: not implemented"; return nil }

// The spec will not be updated by antrea-agent once it is approved or denied.

func appendApprovalCondition(csr *certificatesv1.CertificateSigningRequest, message string) {
	_ = "STUB: not implemented"
	return
}

func (c *CSRApprovingController) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}
