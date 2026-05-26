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
	"crypto"
	"crypto/x509"
	"sync/atomic"
	"time"

	certificatesv1 "k8s.io/api/certificates/v1"
	clientset "k8s.io/client-go/kubernetes"
	csrlister "k8s.io/client-go/listers/certificates/v1"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	ipsecRootCAName               = "antrea-ipsec-ca"
	ipsecCSRSigningControllerName = "IPsecCertificateSigningRequestSigningController"
	workerItemKey                 = "key"
	rootCACertKey                 = "ca.crt"

	duration365d = time.Hour * 24 * 365
	duration10y  = duration365d * 10
)

// IPsecCSRSigningController is responsible for signing CertificateSigningRequests.
type IPsecCSRSigningController struct {
	client          clientset.Interface
	csrInformer     cache.SharedIndexInformer
	csrLister       csrlister.CertificateSigningRequestLister
	csrListerSynced cache.InformerSynced

	configMapInformer     cache.SharedIndexInformer
	configMapLister       corev1listers.ConfigMapLister
	configMapListerSynced cache.InformerSynced

	selfSignedCA bool

	// saved CertificateAuthority
	certificateAuthority atomic.Value

	queue         workqueue.TypedRateLimitingInterface[string]
	fixturesQueue workqueue.TypedRateLimitingInterface[string]
}

// certificateAuthority implements a certificate authority and used by the signing controller.
type certificateAuthority struct {
	// RawCert is an optional field to determine if signing cert/key pairs have changed
	RawCert []byte
	// RawKey is an optional field to determine if signing cert/key pairs have changed
	RawKey []byte

	Certificate *x509.Certificate
	PrivateKey  crypto.Signer
}

func (c *certificateAuthority) signCSR(template *x509.Certificate, requestKey crypto.PublicKey) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newIPsecCSRSigningController supports setting minRetryDelay and maxRetryDelay to non-default
// values for testing.
func newIPsecCSRSigningController(
	client clientset.Interface,
	csrInformer cache.SharedIndexInformer,
	csrLister csrlister.CertificateSigningRequestLister,
	selfSignedCA bool,
	minRetryDelay, maxRetryDelay time.Duration,
) *IPsecCSRSigningController {
	_ = "STUB: not implemented"
	return nil
}

// NewIPsecCSRSigningController returns a new *IPsecCSRSigningController.
func NewIPsecCSRSigningController(
	client clientset.Interface,
	csrInformer cache.SharedIndexInformer,
	csrLister csrlister.CertificateSigningRequestLister,
	selfSignedCA bool,
) *IPsecCSRSigningController {
	_ = "STUB: not implemented"
	return nil
}

// Run begins watching and syncing of the IPsecCSRSigningController.
func (c *IPsecCSRSigningController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *IPsecCSRSigningController) syncRootCertificateAndKey() error {
	_ = "STUB: not implemented"
	return nil
}

func (c *IPsecCSRSigningController) csrWorker() { _ = "STUB: not implemented"; return }

// watchSecretChanges uses watch API directly to watch for Secret changes.
// Antrea Controller should not have List permission for Secrets.
func (c *IPsecCSRSigningController) watchSecretChanges(endCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// re-queue in case of missing events before watcher starts.

// we do not care the actual Event.

func (c *IPsecCSRSigningController) fixturesWorker() { _ = "STUB: not implemented"; return }

func (c *IPsecCSRSigningController) enqueueCertificateSigningRequest(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *IPsecCSRSigningController) syncCSR(key string) error {
	_ = "STUB: not implemented"
	return nil
}

func newCertificateTemplate(certReq *x509.CertificateRequest, usage []certificatesv1.KeyUsage) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// defaults to 1 year

func (c *IPsecCSRSigningController) processNextFixtureWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *IPsecCSRSigningController) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// generateSelfSignedRootCertificate creates self-signed CA certificates and returns the PEM encoded
// certificates and private key.
func generateSelfSignedRootCertificate(commonName string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// valid an hour earlier to avoid flakes due to clock skew
