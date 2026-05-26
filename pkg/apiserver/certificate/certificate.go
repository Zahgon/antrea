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

package certificate

import (
	apiextensionclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apiserver/pkg/server/options"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
)

const (
	// The names of the files that should contain the CA certificate and the TLS key pair.
	CACertFile  = "ca.crt"
	TLSCertFile = "tls.crt"
	TLSKeyFile  = "tls.key"
)

func ApplyServerCert(selfSignedCert bool,
	client kubernetes.Interface,
	aggregatorClient clientset.Interface,
	apiExtensionClient apiextensionclientset.Interface,
	secureServing *options.SecureServingOptionsWithLoopback,
	caConfig *CAConfig) (*CACertController, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The secret may be created after the Pod is created, for example, when cert-manager is used the secret
// is created asynchronously. It waits for a while before it's considered to be failed.

// Since 1.17.0 (https://github.com/kubernetes/kubernetes/commit/3f5fbfbfac281f40c11de2f57d58cc332affc37b),
// apiserver reloads certificate cert and key file from disk every minute, allowing serving tls config to be updated.
