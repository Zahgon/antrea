// Copyright 2019 Antrea Authors
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

package antctl

import (
	"io"
	"time"

	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
)

// requestOption describes options to issue requests.
type requestOption struct {
	commandDefinition *commandDefinition
	// kubeconfig is the path to the config file for kubectl.
	kubeconfig string
	// args are the parameters of the ongoing resourceRequest.
	args map[string]string
	// timeout specifies a time limit for requests made by the client. The timeout
	// duration includes connection setup, all redirects, and reading of the
	// response body.
	timeout time.Duration
	// server is the address and port of the APIServer specified by user explicitly.
	// If not set, antctl will connect to 127.0.0.1:10350 in agent mode, and will
	// connect to the server set in kubeconfig in controller mode.
	// It set, it takes precedence over the above default endpoints.
	server string
}

type AntctlClient interface {
	request(opt *requestOption) (io.Reader, error)
}

// client issues requests to endpoints.
type client struct {
	// codec is the CodecFactory for this command, it is needed for remote accessing.
	codec serializer.CodecFactory
}

func newClient(codec serializer.CodecFactory) AntctlClient {
	_ = "STUB: not implemented"
	return *new(AntctlClient)
}

// resolveKubeconfig tries to load the kubeconfig specified in the requestOption.
// It will return error if the stating of the file failed or the kubeconfig is malformed.
// If the default kubeconfig not exists, it will try to use an in-cluster config.
func (c *client) resolveKubeconfig(opt *requestOption) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) request(opt *requestOption) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (c *client) nonResourceRequest(e *nonResourceEndpoint, opt *requestOption) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

/* isResourceRequest */

func (c *client) resourceRequest(e *resourceEndpoint, opt *requestOption) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// If timeout is zero, there will be no timeout.

/* isResourceRequest */
