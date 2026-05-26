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

package cni

import (
	"github.com/containernetworking/cni/pkg/skel"

	cnipb "antrea.io/antrea/v2/pkg/apis/cni/v1beta1"
)

type Action int

const (
	ActionAdd Action = iota
	ActionCheck
	ActionDel
)

// AntreaCNIVersion is the full semantic version (https://semver.org/) of our CNI Protobuf / gRPC
// service.
//
// We follow these best practices (https://cloud.google.com/apis/design/versioning) for the
// versioning of the CNI Protobuf / gRPC service. The major version number is encoded as the last
// component of the proto package name. For pre-GA releases, the last component also includes the
// pre-release version name (e.g. beta) and the pre-release version number. As the API evolves, the
// major version number (and therefore the proto package name) will change if and only if API
// backwards-compatibility is broken.
//
// Here are some potential scenarios we need to accommodate:
//
//  1. major API refactor that breaks backwards-compatibility: in this case we would increase the
//     major version number.
//
//  2. support for a new CNI version:
//
//     - introduction of a new RPC (e.g. when the CHECK command was added in version 0.4.0 of the
//     CNI spec). In such a case we would increment the minor version number (backwards-
//     compatibility is not broken). If antrea-cni does not support this new version, it will
//     not list the new CNI spec version as supported and there will be no issue. If antrea-cni
//     supports it but not the antrea-agent, the gRPC server will return an UNIMPLEMENTED error
//     which we can propagate to the runtime. There is no way to handle this last case better
//     with the current design unless we introduce a different RPC (e.g. Capabilities) early to
//     query the server for the supported API version. This would also require an additional RPC
//     for each CNI binary invocation.
//
//     - introduction of a new field to a proto message: highly unlikely because we just send the
//     CNI input / output as bytes.
//
//     -  no changes are needed if only the CNI parameters or CNI result format changed. In this
//     case either antrea-cni or antrea-agent will reject the CNI request by validating the
//     cniVersion against the list of supported versions. This is independent of which version
//     of the gRPC service is used by either antrea-cni or antrea-agent.
//
// The gRPC server will return UNIMPLEMENTED if the service is unknown (mismatch in package name,
// i.e. mismatch in major version number) or if the method is unknown. In both cases we return an
// INCOMPATIBLE_API_VERSION error to the container runtime.
//
// To limit incompatibility cases, we can strive to support multiple releases (and in particular all
// pre-GA releases of a major version, along with that major version release itself) in the
// server. This is harder to do on the client side (need to fallback to a previous version when
// getting an UNIMPLEMENTED error).
const AntreaCNIVersion = "1.0.0-beta.1"

// To allow for testing with a fake client.
var withClient = rpcClient

func rpcClient(f func(client cnipb.CniClient) error) error {
	_ = "STUB: not implemented"
	// When using a custom dialer, it makes more sense to use the deprecated grpc.Dial function
	// instead of grpc.NewClient, as grpc.Dial will default to the "dns" resolver (instead of
	// the legacy "passthrough" resolver). Using grpc.NewClient would require us to explicitly
	// add the "passthrough" scheme to the target address. Note that grpc.Dial will stay
	// supported throughout gRPC 1.x.
	// A custom dialer is required for Windows named-pipe support.
	// nolint: staticcheck
	return nil
}

// Request requests the antrea-agent to execute the specified action with the provided arguments via RPC.
// If successful, it outputs the result to stdout and returns nil. Otherwise types.Error is returned.
func (a Action) Request(arg *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Handle gRPC errors.

// network errors, could be transient.

// all other RPC errors.

// Handle errors during CNI execution.
