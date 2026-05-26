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
	"crypto/x509"

	certificatesv1 "k8s.io/api/certificates/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	clientset "k8s.io/client-go/kubernetes"
)

const (
	ipsecCSRApproverName = "AntreaIPsecCSRApprover"
)

type ipsecCSRApprover struct {
	client                        clientset.Interface
	antreaAgentServiceAccountName string
}

var ipsecTunnelUsages = sets.New[string](
	string(certificatesv1.UsageIPsecTunnel),
)

var _ approver = (*ipsecCSRApprover)(nil)

func getAntreaAgentServiceAccount() string { _ = "STUB: not implemented"; return "" }

func newIPsecCSRApprover(client clientset.Interface) *ipsecCSRApprover {
	_ = "STUB: not implemented"
	return nil
}

func (ic *ipsecCSRApprover) recognize(csr *certificatesv1.CertificateSigningRequest) bool {
	_ = "STUB: not implemented"
	return false
}

func (ic *ipsecCSRApprover) verify(csr *certificatesv1.CertificateSigningRequest) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ic *ipsecCSRApprover) name() string { _ = "STUB: not implemented"; return "" }

func (ic *ipsecCSRApprover) verifyCertificateRequest(req *x509.CertificateRequest, usages []certificatesv1.KeyUsage) error {
	_ = "STUB: not implemented"
	return nil
}

func (ic *ipsecCSRApprover) verifyIdentity(nodeName string, csr *certificatesv1.CertificateSigningRequest) error {
	_ = "STUB: not implemented"
	return nil
}
