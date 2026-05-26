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

package s3uploader

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	flowpb "antrea.io/antrea/v2/pkg/apis/flow/v1alpha1"
	config "antrea.io/antrea/v2/pkg/config/flowaggregator"
	"antrea.io/antrea/v2/pkg/flowaggregator/flowrecord"
)

const (
	bufferFlushTimeout         = 1 * time.Minute
	maxNumBuffersPendingUpload = 5
)

// GetS3BucketRegion is used for unit testing
var GetS3BucketRegion = getBucketRegion

type stopPayload struct {
	flushQueue bool
}

type S3UploadProcess struct {
	bucketName       string
	bucketPrefix     string
	region           string
	compress         bool
	maxRecordPerFile int32
	// uploadInterval is the interval between batch uploads
	uploadInterval time.Duration
	// uploadTicker is a ticker, containing a channel used to trigger batchUploadAll() for every uploadInterval period
	uploadTicker *time.Ticker
	// stopCh is the channel to receive stop message
	stopCh chan stopPayload
	// exportWg is to ensure that all messages have been flushed from the queue when we stop
	exportWg             sync.WaitGroup
	exportProcessRunning bool
	// mutex protects configuration state from concurrent access
	mutex sync.Mutex
	// queueMutex protects currentBuffer and bufferQueue from concurrent access
	queueMutex sync.Mutex
	// currentBuffer caches flow record
	currentBuffer *bytes.Buffer
	// cachedRecordCount keeps track of the number of flow records written into currentBuffer
	cachedRecordCount int32
	// bufferQueue caches currentBuffer when it is full
	bufferQueue []*bytes.Buffer
	// buffersToUpload stores all the buffers to be uploaded for the current uploadFile() call
	buffersToUpload []*bytes.Buffer
	gzipWriter      *gzip.Writer
	// awsS3Client is used to initialize awsS3Uploader
	awsS3Client *s3.Client
	// awsS3Uploader makes the real call to aws-sdk UploadObject() method to upload an object to S3
	awsS3Uploader *transfermanager.Client
	// s3UploaderAPI wraps the call made by awsS3Uploader
	s3UploaderAPI S3UploaderAPI
	clusterUUID   string
}

type S3Input struct {
	Config         config.S3UploaderConfig
	UploadInterval time.Duration
}

// Define a wrapper interface S3UploaderAPI to assist unit testing.
type S3UploaderAPI interface {
	Upload(ctx context.Context, input *transfermanager.UploadObjectInput, awsS3Uploader *transfermanager.Client, opts ...func(*transfermanager.Options)) (
		*transfermanager.UploadObjectOutput, error,
	)
}

type S3Uploader struct{}

func (u *S3Uploader) Upload(ctx context.Context, input *transfermanager.UploadObjectInput, awsS3Uploader *transfermanager.Client, opts ...func(*transfermanager.Options)) (*transfermanager.UploadObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getBucketRegion determines the exact region in which the bucket is
// located. regionHint can be any region in the same partition as the one in
// which the bucket is located.
func getBucketRegion(ctx context.Context, bucket string, regionHint string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func NewS3UploadProcess(input S3Input, clusterUUID string) (*S3UploadProcess, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *S3UploadProcess) GetBucketName() string { _ = "STUB: not implemented"; return "" }

func (p *S3UploadProcess) GetBucketPrefix() string { _ = "STUB: not implemented"; return "" }

func (p *S3UploadProcess) GetRegion() string { _ = "STUB: not implemented"; return "" }

func (p *S3UploadProcess) GetUploadInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (p *S3UploadProcess) UpdateS3Uploader(bucketName, bucketPrefix, region string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *S3UploadProcess) SetUploadInterval(uploadInterval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (p *S3UploadProcess) CacheRecord(record *flowpb.Flow) error {
	_ = "STUB: not implemented"
	return nil
}

// If the number of pending records in the buffer reaches maxRecordPerFile,
// add the buffer to bufferQueue.

func (p *S3UploadProcess) Start() { _ = "STUB: not implemented"; return }

func (p *S3UploadProcess) Stop() { _ = "STUB: not implemented"; return }

func (p *S3UploadProcess) startExportProcess() { _ = "STUB: not implemented"; return }

func (p *S3UploadProcess) stopExportProcess(flushQueue bool) { _ = "STUB: not implemented"; return }

func (p *S3UploadProcess) flowRecordPeriodicCommit() { _ = "STUB: not implemented"; return }

// batchUploadAll uploads all buffers cached in bufferQueue and previous fail-
// to-upload buffers stored in buffersToUpload. Returns error encountered
// during upload if any.
func (p *S3UploadProcess) batchUploadAll(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// dump cached buffers from bufferQueue to buffersToUpload

func (p *S3UploadProcess) writeRecordToBuffer(record *flowrecord.FlowRecord) {
	_ = "STUB: not implemented"
	return
}

func (p *S3UploadProcess) uploadFile(ctx context.Context, reader *bytes.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// appendBufferToQueue appends currentBuffer to bufferQueue, and reset
// currentBuffer. Caller of this function should acquire queueMutex.
func (p *S3UploadProcess) appendBufferToQueue() { _ = "STUB: not implemented"; return }

// avoid too many memory allocations

func randSeq(n int) string { _ = "STUB: not implemented"; return "" }

// #nosec G404: random number generator not used for security purposes.

func writeRecord(w io.Writer, r *flowrecord.FlowRecord, clusterUUID string) {
	_ = "STUB: not implemented"
	return
}

// Enclose Pod labels with single quote because commas are used to separate
// different columns in CSV and Pod labels json string contains commas.
