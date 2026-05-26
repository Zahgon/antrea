// Copyright 2024 Antrea Authors
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

package certificate

import (
	"context"
	"crypto/x509"
	"net"
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apiserver/pkg/server/dynamiccertificates"
	"k8s.io/apiserver/pkg/server/options"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	clockutils "k8s.io/utils/clock"
)

var loopbackAddresses = []net.IP{net.ParseIP("127.0.0.1"), net.IPv6loopback}

// generateSelfSignedCertKeyFn represents a function which can create a self-signed certificate and
// key for the given host.
type generateSelfSignedCertKeyFn func(host string, alternateIPs []net.IP, alternateDNS []string) ([]byte, []byte, error)

type selfSignedCertProvider struct {
	client          kubernetes.Interface
	secretInformer  cache.SharedIndexInformer
	secretLister    corelisters.SecretLister
	secretNamespace string
	secureServing   *options.SecureServingOptionsWithLoopback
	caConfig        *CAConfig
	clock           clockutils.Clock

	listeners []dynamiccertificates.Listener
	// queue only ever has one item, but it has nice error handling backoff/retry semantics
	queue workqueue.TypedRateLimitingInterface[string]

	// mutex protects the fields following it.
	mutex sync.RWMutex
	// cert and key represent the contents of the cert file and the key file.
	cert          []byte
	key           []byte
	verifyOptions *x509.VerifyOptions

	// generateSelfSignedCertKey is the function used to generate self-signed certificates and keys.
	// We use a struct member for unit testing.
	generateSelfSignedCertKey generateSelfSignedCertKeyFn
}

var _ dynamiccertificates.CAContentProvider = &selfSignedCertProvider{}
var _ dynamiccertificates.ControllerRunner = &selfSignedCertProvider{}

type providerOption func(p *selfSignedCertProvider)

func withGenerateSelfSignedCertKeyFn(fn generateSelfSignedCertKeyFn) providerOption {
	_ = "STUB: not implemented"
	return *new(providerOption)
}

func withClock(clock clockutils.Clock) providerOption {
	_ = "STUB: not implemented"
	return *new(providerOption)
}

func newSelfSignedCertProvider(client kubernetes.Interface, secureServing *options.SecureServingOptionsWithLoopback, caConfig *CAConfig, options ...providerOption) (*selfSignedCertProvider, error) {
	_ = "STUB: not implemented"
	// Set the CertKey and CertDirectory to generate the certificate files.
	return nil, nil
}

// In clusters where antrea-controller's deployment strategy is set to RollingUpdate, two instances may run
// simultaneously in a short time when the deployment is being updated. The event handlers are for the case that
// the certificate needs rotation during that time window. With it, regardless of which instance updates the
// secret first, the other one will switch to it and stop generating a new one.
// In the future when HA is implemented, we should only let the active instance rotate the certificate, and the
// standby instances should refresh its certificate immediately with the event handlers.

func (p *selfSignedCertProvider) RunOnce(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *selfSignedCertProvider) Run(ctx context.Context, workers int) {
	_ = "STUB: not implemented"
	return
}

// doesn't matter what workers say, only start one.

// check if the certificate should be regenerated periodically.

func (p *selfSignedCertProvider) Name() string { _ = "STUB: not implemented"; return "" }

func (p *selfSignedCertProvider) CurrentCABundleContent() []byte {
	_ = "STUB: not implemented"
	return nil
}

func (p *selfSignedCertProvider) VerifyOptions() (x509.VerifyOptions, bool) {
	_ = "STUB: not implemented"
	return *new(x509.VerifyOptions), false
}

func newVerifyOptions(caBundle []byte) *x509.VerifyOptions {
	_ = "STUB: not implemented"
	// We don't really use the CA bundle to verify clients, this is just to follow DynamicFileCAContent.
	return nil
}

func (p *selfSignedCertProvider) AddListener(listener dynamiccertificates.Listener) {
	_ = "STUB: not implemented"
	return
}

func (p *selfSignedCertProvider) runWorker() { _ = "STUB: not implemented"; return }

func (p *selfSignedCertProvider) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *selfSignedCertProvider) enqueue() {
	_ = "STUB: not implemented"
	// The key can be anything as we only have a single item.
	return
}

func (p *selfSignedCertProvider) shouldRotateCertificate(certBytes []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// rotateSelfSignedCertificate generates a new self-signed certificate if it needs to.
func (p *selfSignedCertProvider) rotateSelfSignedCertificate() error {
	_ = "STUB: not implemented"
	return nil
}

// If Secret is specified, we should prioritize it.

// If Secret is specified, we should save the new certificate and key to it.

// If the certificate and key don't change, do nothing.

func (p *selfSignedCertProvider) getCertKeyFromSecret() (*corev1.Secret, []byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (p *selfSignedCertProvider) saveCertKeyToSecret(secret *corev1.Secret, cert []byte, key []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not update the existing Secret's type. Otherwise, the update would fail if it's not of type
// "kubernetes.io/tls" as the type field is immutable.
