// Copyright 2024 Antrea Authors.
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

package packetcapture

import (
	"context"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/spf13/afero"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/pkg/agent/interfacestore"
	crdv1alpha1 "antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	clientsetversioned "antrea.io/antrea/v2/pkg/client/clientset/versioned"
	crdinformers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha1"
	crdlisters "antrea.io/antrea/v2/pkg/client/listers/crd/v1alpha1"
	"antrea.io/antrea/v2/pkg/util/sftp"
)

type storageProtocolType string

const (
	sftpProtocol storageProtocolType = "sftp"
)

const (
	controllerName               = "PacketCaptureController"
	resyncPeriod   time.Duration = 0

	minRetryDelay = 500 * time.Millisecond
	maxRetryDelay = 30 * time.Second

	defaultWorkers = 4

	// defines how many capture request we can handle concurrently. waiting captures will be
	// marked as Pending until they can be processed.
	maxConcurrentCaptures     = 16
	captureStatusUpdatePeriod = 10 * time.Second
	// PacketCapture uses a dedicated Secret object to store authentication information for a file server.
	// #nosec G101
	fileServerAuthSecretName = "antrea-packetcapture-fileserver-auth"

	// max packet size we can capture.
	snapLen = 65536
)

type packetCapturePhase string

const (
	packetCapturePhasePending  packetCapturePhase = "Pending"
	packetCapturePhaseStarted  packetCapturePhase = "Started"
	packetCapturePhaseComplete packetCapturePhase = "Complete"
)

var (
	packetDirectory = filepath.Join(os.TempDir(), "antrea", "packetcapture", "packets")
	defaultFS       = afero.NewOsFs()
)

type packetCaptureState struct {
	// capturedPacketsNum records how many packets have been captured. Due to the RateLimiter,
	// this may not be the real-time data.
	capturedPacketsNum int32
	// targetCapturedPacketsNum is the target number limit for a PacketCapture. When numCapturedPackets == targetCapturedPacketsNum, it means
	// the PacketCapture is done successfully.
	targetCapturedPacketsNum int32
	// phase is the phase of the PacketCapture.
	phase packetCapturePhase
	// filePath is the final path shown in PacketCapture's status.
	filePath string
	// captureErr is the error observed during the capturing phase.
	captureErr error
	// uploadErr is the error observed during the uploading phase.
	uploadErr error
	// cancel is the cancel function for capture context.
	cancel context.CancelFunc
}

func (pcs *packetCaptureState) isCaptureSuccessful() bool { _ = "STUB: not implemented"; return false }

type Controller struct {
	kubeClient            clientset.Interface
	crdClient             clientsetversioned.Interface
	packetCaptureInformer crdinformers.PacketCaptureInformer
	packetCaptureLister   crdlisters.PacketCaptureLister
	packetCaptureSynced   cache.InformerSynced
	interfaceStore        interfacestore.InterfaceStore
	queue                 workqueue.TypedRateLimitingInterface[string]
	sftpUploader          sftp.Uploader
	captureInterface      PacketCapturer
	mutex                 sync.Mutex
	// A name-state mapping for all PacketCapture CRs.
	captures           map[string]*packetCaptureState
	numRunningCaptures int
}

func NewPacketCaptureController(
	kubeClient clientset.Interface,
	crdClient clientsetversioned.Interface,
	packetCaptureInformer crdinformers.PacketCaptureInformer,
	interfaceStore interfacestore.InterfaceStore,
) (*Controller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Controller) enqueuePacketCapture(pc *crdv1alpha1.PacketCapture) {
	_ = "STUB: not implemented"
	return

	// Run will create defaultWorkers workers (go routines) which will process the PacketCapture events from the
	// workqueue.
}

func (c *Controller) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *Controller) addPacketCapture(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) updatePacketCapture(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) deletePacketCapture(obj interface{}) { _ = "STUB: not implemented"; return }

func nameToPath(name string) string { _ = "STUB: not implemented"; return "" }

func (c *Controller) worker() { _ = "STUB: not implemented"; return }

func (c *Controller) processPacketCaptureItem() bool { _ = "STUB: not implemented"; return false }

func (c *Controller) syncPacketCapture(pcName string) error { _ = "STUB: not implemented"; return nil }

// Lister.Get only returns error when the resource is not found.

// Capture will not occur on this Node if a corresponding Pod interface is not found.

// Do not return the error as it's not a transient error.

// Return the error as it's a transient error.

// The OpenAPI schema for the CRD makes sure Spec.Timeout is not nil.

// Start the capture goroutine in a separate goroutine. The goroutine will decrease numRunningCaptures on exit.

func (c *Controller) validatePacketCapture(spec *crdv1alpha1.PacketCaptureSpec) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) cleanupPacketCapture(pcName string) { _ = "STUB: not implemented"; return }

func getPacketFile(filePath string) (afero.File, error) {
	_ = "STUB: not implemented"
	return *new(afero.File), nil
}

// getTargetCaptureDevice is trying to locate the target device for packet capture. If the target
// Pod does not exist on the current Node, the agent on this Node will not perform the capture.
// In the PacketCapture spec, at least one of `.Spec.Source.Pod` or `.Spec.Destination.Pod`
// should be set.
func (c *Controller) getTargetCaptureDevice(pc *crdv1alpha1.PacketCapture) string {
	_ = "STUB: not implemented"
	// Set CapturePoint to 'Source' if a Source Pod is specified; otherwise, use 'Destination'.
	return ""
}

// getPodDevice returns the network device name for the given PodReference using the interfaceStore.
func (c *Controller) getPodDevice(pod *crdv1alpha1.PodReference) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Controller) startCapture(ctx context.Context, pc *crdv1alpha1.PacketCapture, state *packetCaptureState, device string) {
	_ = "STUB: not implemented"
	return
}

// Resync the PacketCapture on exit of the capture goroutine.

// If nothing is captured, no need to proceed.

// If any is captured, upload it if required and update filePath in the status of the PacketCapture.

// It can't use the same context as performCapture because it might have timed out.

// performCapture blocks until either the target number of packets have been captured, the context is canceled, or the
// context reaches its deadline.
// It returns a boolean indicating whether any packet is captured, and an error if the target number of packets are not
// captured.
func (c *Controller) performCapture(
	ctx context.Context,
	pc *crdv1alpha1.PacketCapture,
	captureState *packetCaptureState,
	file afero.File,
	device string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// set SnapLength here to make tcpdump on Mac OSX works. By default, its value is
// 0 and means unlimited, but tcpdump on Mac OSX will complain:
// 'tcpdump: pcap_loop: invalid packet capture length <len>, bigger than snaplen of 524288'

// Track whether any packet is captured.

// use rate limiter to reduce the times we need to update status.

func (c *Controller) getPodIP(ctx context.Context, podRef *crdv1alpha1.PodReference, ipFamily v1.IPFamily) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

func (c *Controller) parseIPs(ctx context.Context, pc *crdv1alpha1.PacketCapture) (srcIP, dstIP net.IP, err error) {
	_ = "STUB: not implemented"
	return *new(net.IP), *new(net.IP), nil
}

func (c *Controller) getUploaderByProtocol(protocol storageProtocolType) (sftp.Uploader, error) {
	_ = "STUB: not implemented"
	return *new(sftp.Uploader), nil
}

func (c *Controller) generatePacketsPathForServer(name string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Controller) uploadPackets(ctx context.Context, pc *crdv1alpha1.PacketCapture, outputFile afero.File) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) updateStatus(ctx context.Context, pc *crdv1alpha1.PacketCapture, state packetCaptureState) error {
	_ = "STUB: not implemented"
	// Make a deepcopy as the object returned from lister must not be updated directly.
	return nil
}

// Set Uploaded condition if applicable.

// Return the error from UPDATE.

func mergeConditions(oldConditions, newConditions []crdv1alpha1.PacketCaptureCondition) []crdv1alpha1.PacketCaptureCondition {
	_ = "STUB: not implemented"
	return nil
}

// Use the original Condition if the only change is about lastTransition time

// Use the latest Condition.
