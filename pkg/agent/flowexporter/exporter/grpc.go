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
	"google.golang.org/grpc"

	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
)

type grpcExporter struct {
	nodeName    string
	nodeUID     string
	obsDomainID uint32
	grpcClient  *grpc.ClientConn
	client      flowpb.FlowExportServiceClient
	stream      flowpb.FlowExportService_ExportClient
}

func NewGRPCExporter(nodeName string, nodeUID string, obsDomainID uint32) *grpcExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *grpcExporter) ConnectToCollector(addr string, tlsConfig *TLSConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *grpcExporter) Export(conn *connection.Connection) error {
	_ = "STUB: not implemented"
	// It is not safe to modify the message after calling SendMsg, so we need to allocate a
	// brand new message every time.
	// We could investigate using grpc.PreparedMsg to see if it helps reduce the number of
	// allocations, but that is an experimental API:
	// https://github.com/grpc/grpc-go/issues/8186
	return nil
}

// At the moment, we send a single flow per stream message because it's simpler.
// Note that there should still be some batching at the transport layer.
// In the future, we should try to do explicit batching.

func (e *grpcExporter) CloseConnToCollector() { _ = "STUB: not implemented"; return }

func (e *grpcExporter) createMessage(conn *connection.Connection) *flowpb.Flow {
	_ = "STUB: not implemented"
	return nil
}

// not used currently

// Add nodeName / nodeUID only for local Pods whose Pod names are resolved.
