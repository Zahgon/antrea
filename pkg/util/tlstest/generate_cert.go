// Copyright 2025 Antrea Authors
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

package tlstest

import (
	"time"
)

// Adapted from test certificate generation in src/crypto/tls.

func publicKey(priv any) any { _ = "STUB: not implemented"; return *new(any) }

func GenerateCert(
	hosts []string, // List of hostnames and IPs.
	notBefore time.Time,
	validFor time.Duration,
	isCA bool, // whether this cert should be its own Certificate Authority.
	isClient bool, // whether this cert is to be used as a client certificate.
	rsaBits int, // Size of RSA key to generate. Ignored if ecdsaCurve is not empty.
	ecdsaCurve string, // ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521.
	ed25519Key bool, // Generate an Ed25519 key.
) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ECDSA, ED25519 and RSA subject keys should have the DigitalSignature
// KeyUsage bits set in the x509.Certificate template

// Only RSA subject keys should have the KeyEncipherment KeyUsage bits set. In
// the context of TLS this KeyUsage is particular to RSA key exchange and
// authentication.
