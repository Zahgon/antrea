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

package exporter

import (
	"context"

	"github.com/google/uuid"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
	"antrea.io/antrea/v2/pkg/flowaggregator/options"
	"antrea.io/antrea/v2/pkg/flowaggregator/ringbuffer"
	"antrea.io/antrea/v2/pkg/flowaggregator/s3uploader"
)

type S3Exporter struct {
	s3Input         *s3uploader.S3Input
	s3UploadProcess *s3uploader.S3UploadProcess
}

func buildS3Input(opt *options.Options) s3uploader.S3Input {
	_ = "STUB: not implemented"
	return *new(s3uploader.S3Input)
}

func NewS3Exporter(clusterUUID uuid.UUID, opt *options.Options) (*S3Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run consumes flow records from the ring buffer and uploads them to S3.
// It blocks until ctx is cancelled or the consumer signals shutdown.
func (e *S3Exporter) Run(ctx context.Context, buf ringbuffer.BroadcastBuffer[*flowpb.Flow]) {
	_ = "STUB: not implemented"
	return
}
