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

package collector

import (
	"sync/atomic"

	"google.golang.org/grpc"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
)

const grpcCollectorAddress = "0.0.0.0:14739"

type grpcCollector struct {
	service *grpcService
	server  *grpc.Server
}

func NewGRPCCollector(recordCh chan *flowpb.Flow, caCert, serverKey, serverCert []byte) (*grpcCollector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *grpcCollector) Run(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// #nosec G102: binding to all network interfaces is intentional
	return
}

// c.server.Stop() will close the listener

func (c *grpcCollector) GetNumRecordsReceived() int64 { _ = "STUB: not implemented"; return 0 }

func (c *grpcCollector) GetNumConnsToCollector() int64 { _ = "STUB: not implemented"; return 0 }

type grpcService struct {
	flowpb.UnimplementedFlowExportServiceServer
	recordCh           chan *flowpb.Flow
	numRecordsReceived atomic.Int64
	numConns           atomic.Int64
}

func (s *grpcService) Export(stream flowpb.FlowExportService_ExportServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Natches the go-ipfix code:
// https://github.com/vmware/go-ipfix/blob/961f78e9fa2d7a417ee4dd1b95f29b08fa2a794d/pkg/collector/process.go#L274-L279
// handle IPv6 address which may involve []
