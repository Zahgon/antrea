// Copyright 2025 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exporter

import (
	"crypto/tls"

	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
)

type TLSConfig struct {
	ServerName    string
	CAData        []byte
	CertData      []byte
	KeyData       []byte
	MinTLSVersion string
}

// AsStdConfig converts the TLSConfig to the standard tls.Config.
func (c *TLSConfig) AsStdConfig() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

// Use system roots if c.CAData == nil.

// Don't use a client certificate if c.CertData == nil.

// Implementations of this interface don't provide any guarantees regarding thread-safety.
type Interface interface {
	ConnectToCollector(addr string, tlsConfig *TLSConfig) error
	Export(conn *connection.Connection) error
	CloseConnToCollector()
}
