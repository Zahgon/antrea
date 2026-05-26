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

package proxy

import (
	"strings"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/proxy"
)

const (
	defaultPort         = 8001
	defaultStaticPrefix = "/static/"
	defaultAPIPrefix    = "/"
	defaultAddress      = "127.0.0.1"
)

// Command is the proxy command implementation.
var Command *cobra.Command

type proxyOptions struct {
	staticDir     string
	staticPrefix  string
	apiPrefix     string
	acceptPaths   string
	rejectPaths   string
	acceptHosts   string
	rejectMethods string
	port          int
	address       string
	disableFilter bool
	unixSocket    string
	keepalive     time.Duration

	filter *proxy.FilterServer

	controller    bool
	agentNodeName string
	insecure      bool
}

var options *proxyOptions
var defaultFS = afero.NewOsFs()

// validateAndComplete checks the proxyOptions to see if there is sufficient information to run the
// command, and adds default values when needed.
func (o *proxyOptions) validateAndComplete() error { _ = "STUB: not implemented"; return nil }

// default to controller

var proxyCommandExample = strings.Trim(`
  Start a reverse proxy for the Antrea Controller API
  $ antctl proxy --controller
  Start a reverse proxy for the API of an Antrea Agent running on a specific Node
  $ antctl proxy --agent-node <Node Name>
`, "\n")

func init() {
	Command = &cobra.Command{
		Use:     "proxy",
		Short:   "Run a reverse proxy to access Antrea API",
		Long:    "Run a reverse proxy to access Antrea API (Controller or Agent). Command only supports remote mode. HTTPS connections between the proxy and the Antrea API will not be secure (no certificate verification).",
		Example: proxyCommandExample,
		RunE:    runE,
		Args:    cobra.NoArgs,
	}

	o := &proxyOptions{}
	options = o
	// These options are the same as for "kubectl proxy".
	// https://github.com/kubernetes/kubectl/blob/v0.19.0/pkg/cmd/proxy/proxy.go
	Command.Flags().StringVarP(&o.staticDir, "www", "w", "", "Also serve static files from the given directory under the specified prefix.")
	Command.Flags().StringVarP(&o.staticPrefix, "www-prefix", "P", defaultStaticPrefix, "Prefix to serve static files under, if static file directory is specified.")
	Command.Flags().StringVarP(&o.apiPrefix, "api-prefix", "", defaultAPIPrefix, "Prefix to serve the proxied API under.")
	Command.Flags().StringVar(&o.acceptPaths, "accept-paths", proxy.DefaultPathAcceptRE, "Regular expression for paths that the proxy should accept.")
	Command.Flags().StringVar(&o.rejectPaths, "reject-paths", proxy.DefaultPathRejectRE, "Regular expression for paths that the proxy should reject. Paths specified here will be rejected even accepted by --accept-paths.")
	Command.Flags().StringVar(&o.acceptHosts, "accept-hosts", proxy.DefaultHostAcceptRE, "Regular expression for hosts that the proxy should accept.")
	Command.Flags().StringVar(&o.rejectMethods, "reject-methods", proxy.DefaultMethodRejectRE, "Regular expression for HTTP methods that the proxy should reject (example --reject-methods='POST,PUT,PATCH'). ")
	Command.Flags().IntVarP(&o.port, "port", "p", defaultPort, "The port on which to run the proxy. Set to 0 to pick a random port.")
	Command.Flags().StringVarP(&o.address, "address", "", defaultAddress, "The IP address on which to serve on.")
	Command.Flags().BoolVar(&o.disableFilter, "disable-filter", false, "If true, disable request filtering in the proxy. This is dangerous, and can leave you vulnerable to XSRF attacks, when used with an accessible port.")
	Command.Flags().StringVarP(&o.unixSocket, "unix-socket", "u", "", "Unix socket on which to run the proxy.")
	Command.Flags().DurationVar(&o.keepalive, "keepalive", 0, "keepalive specifies the keep-alive period for an active network connection. Set to 0 to disable keepalive.")

	// These options are specific to "antctl proxy".
	Command.Flags().BoolVar(&o.controller, "controller", false, "Run proxy for Antrea Controller API. If both --controller and --agent-node are omitted, the proxy will run for the Controller API.")
	Command.Flags().StringVar(&o.agentNodeName, "agent-node", "", "Run proxy for Antrea Agent API on the provided K8s Node.")
	Command.Flags().BoolVar(&o.insecure, "insecure", false, "Skip TLS verification when connecting to Antrea API.")
}

func runE(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

// The last argument is for "appendLocationPath", which for "kubectl proxy" is used as
// follows: if the Kubeconfig context provides a server URL which includes a Path comppnent
// (e.g., https://example.com/PATH), then this path is automatically added to all incoming
// requests to the proxy.
// See https://github.com/kubernetes/kubernetes/pull/97350
// In our case, we craft the config manually and clientCfg.Host never includes a Path
// component, so we always set "appendLocationPath" to "false", and there is no need to
// expose a flag like --append-server-path for "antctl proxy".

// Separate listening from serving so we can report the bound port when it is chosen by os
// (eg: port == 0).
