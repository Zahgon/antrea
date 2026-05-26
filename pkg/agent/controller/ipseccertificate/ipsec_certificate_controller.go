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

package ipseccertificate

import (
	"crypto"
	"crypto/x509"
	"time"

	certificatesv1 "k8s.io/api/certificates/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/utils/clock"

	"antrea.io/antrea/v2/pkg/ovs/ovsconfig"
)

const (
	controllerName = "AntreaAgentIPsecCertificateController"
	workerItemKey  = "key"
	minRetryDelay  = 5 * time.Second
	maxRetryDelay  = 60 * time.Second

	// the mount path for CA certificate in antrea-ipsec container.
	// StrongSwan will never reads CA certificates from folders other than `/etc/ipsec.d/cacerts`.
	// Though StrongSwan will automatically load CA certificates from the folder, we set the ca_path in other_configs
	// to the correct path for better consistency.
	caCertificatePath = "/etc/ipsec.d/cacerts/ca.crt"

	ovsConfigCACertificateKey = "ca_cert"
	ovsConfigPrivateKeyKey    = "private_key"
	ovsConfigCertificateKey   = "certificate"

	// certificateWaitTimeout controls the amount of time we wait for certificate approval in
	// one iteration.
	certificateWaitTimeout = 15 * time.Minute
)

var defaultCertificatesPath = "/var/run/openvswitch"

// Controller is responsible for requesting certificates by CertificateSigningRequest and configure them to OVS
type Controller struct {
	kubeClient      clientset.Interface
	ovsBridgeClient ovsconfig.OVSBridgeClient
	nodeName        string
	queue           workqueue.TypedRateLimitingInterface[string]

	rotateCertificate  func() (*certificateKeyPair, error)
	certificateKeyPair *certificateKeyPair

	clock clock.WithTicker

	// caPath and is initialized with NewIPSecCertificateController and should not
	// be changed once Controller starts.
	caPath string
	// certificateFolderPath is the folder to store private keys and issued certificates.
	// defaults to defaultCertificatesPath.
	certificateFolderPath string

	syncedOnce uint32
}

// Manager is an interface to track the status of the IPsec certificate controller.
type Manager interface {
	HasSynced() bool
}

var _ Manager = (*Controller)(nil)

func NewIPSecCertificateController(
	kubeClient clientset.Interface,
	ovsBridgeClient ovsconfig.OVSBridgeClient,
	nodeName string,
) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func newIPSecCertificateControllerWithCustomClock(kubeClient clientset.Interface,
	ovsBridgeClient ovsconfig.OVSBridgeClient,
	nodeName string, clock clock.WithTicker) *Controller {
	_ = "STUB: not implemented"
	return nil
}

// worker is a long-running function that will continually call the processNextWorkItem function in
// order to read and process a message on the workqueue.
func (c *Controller) worker() { _ = "STUB: not implemented"; return }

func (c *Controller) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

type certificateKeyPair struct {
	caCertificate    []*x509.Certificate
	certificate      []*x509.Certificate
	privateKey       crypto.Signer
	certificatePath  string
	privateKeyPath   string
	rotationDeadline time.Time
}

func (pair *certificateKeyPair) validate(clock clock.Clock) error {
	_ = "STUB: not implemented"
	return nil
}

//TODO: support key types other than RSA such as *ecdsa.PublicKey.

// cleanup deletes the files of certificate and private key.
func (pair *certificateKeyPair) cleanup() { _ = "STUB: not implemented"; return }

// Delete the old certificate file.

// Delete the old private key file.

// jitteryDuration returns a duration in [totalDuration * 0.7, totalDuration * 0.9].
func jitteryDuration(totalDuration time.Duration) time.Duration {
	_ = "STUB: not implemented"
	// wait.Jitter returns a duration in [totalDuration, totalDuration * 1.2].
	return *new(time.Duration)
}

// nextRotationDeadline returns a value for the threshold at which the
// current certificate should be rotated, 80%+/-10% of the expiration of the
// certificate. The deadline will not change once calculated.
// This function is not thread-safe.
func (pair *certificateKeyPair) nextRotationDeadline() time.Time {
	_ = "STUB: not implemented"
	// Return the previous calculated rotation deadline if applicable.
	return *new(time.Time)
}

func loadCertAndKeyFromFiles(caPath, certPath, keyPath string) (*certificateKeyPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Controller) syncConfigurations() error { _ = "STUB: not implemented"; return nil }

// Validate the existing certificate and key pair.

// Current certificate is about to expire.

// Clean up old certificate and key pair.

// Save the known good certificate and key pair.

// Calculate the rotation deadline of new certificate.

// Re-queue after the interval to renew the certificate.

// Sync OVS bridge configurations.

// HasSynced implements the Manager interface.
func (c *Controller) HasSynced() bool {
	_ = "STUB: not implemented"
	// returns true if the controller has configured certificate successfully
	// at least once.
	return false
}

func loadRootCA(caPath string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newRSAPrivateKey() (crypto.Signer, []byte, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil, nil
}

func loadPrivateKey(privateKeyPath string) (crypto.Signer, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Signer), nil
}

// Load the private key contents from file.

// Try to parse private key from existing file.

func loadCertificate(certPath string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load the certificate from file.

// Try to parse the certificate from the existing file.

func (c *Controller) syncOVSConfigurations(certPath, keyPath, caPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func newCSR(csrNamePrefix, commonName string, privateKey crypto.Signer) (*certificatesv1.CertificateSigningRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Controller) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Load the previous configured certificate path from OVS database.

func (c *Controller) newCertificateKeyPair() (*certificateKeyPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Always create a new CSR for certificate rotation. The old ones will be GCed automatically.

// Use the hash of new certificate and key as the filename suffix.
