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

package supportbundle

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/spf13/afero"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	clockutils "k8s.io/utils/clock"
	"k8s.io/utils/exec"

	agentquerier "antrea.io/antrea/v2/pkg/agent/querier"
	systemv1beta1 "antrea.io/antrea/v2/pkg/apis/system/v1beta1"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
	"antrea.io/antrea/v2/pkg/querier"
	"antrea.io/antrea/v2/pkg/support"
)

const (
	bundleExpireDuration = time.Hour
	modeController       = "controller"
	modeAgent            = "agent"
)

var (
	// Declared as variables for testing.
	defaultFS       = afero.NewOsFs()
	defaultExecutor = exec.New()
	newAgentDumper  = support.NewAgentDumper

	clock clockutils.Clock = clockutils.RealClock{}
)

// NewControllerStorage creates a support bundle storage for working on antrea controller.
func NewControllerStorage() Storage { _ = "STUB: not implemented"; return *new(Storage) }

// NewAgentStorage creates a support bundle storage for working on antrea agent.
func NewAgentStorage(client ovsctl.OVSCtlClient, aq agentquerier.AgentQuerier, npq querier.AgentNetworkPolicyInfoQuerier, v4Enabled, v6Enabled bool) Storage {
	_ = "STUB: not implemented"
	return *new(Storage)
}

// Storage contains REST resources for support bundle, including status query and download.
type Storage struct {
	SupportBundle *supportBundleREST
	Download      *downloadREST
	Mode          string
}

var (
	_ rest.Scoper               = &supportBundleREST{}
	_ rest.Getter               = &supportBundleREST{}
	_ rest.Creater              = &supportBundleREST{}
	_ rest.GracefulDeleter      = &supportBundleREST{}
	_ rest.SingularNameProvider = &supportBundleREST{}
)

// supportBundleREST implements REST interfaces for bundle status querying.
type supportBundleREST struct {
	mode         string
	statusLocker sync.RWMutex
	cancelFunc   context.CancelFunc
	// cache stores the "current" or most recent SupportBundle resource. Because Get returns
	// this value directly, it is important for the supportBundleREST implementation *not* to
	// mutate fields in the SupportBundle object once it has been assigned to cache, in order to
	// ensure thread-safety. Otherwise, we would have a race with Get callers.
	cache *systemv1beta1.SupportBundle

	ovsCtlClient ovsctl.OVSCtlClient
	aq           agentquerier.AgentQuerier
	npq          querier.AgentNetworkPolicyInfoQuerier
	v4Enabled    bool
	v6Enabled    bool
}

// Create triggers a bundle generation. It only allows resource creation when
// the name matches the mode. It returns metav1.Status if there is any error,
// otherwise it returns the SupportBundle.
func (r *supportBundleREST) Create(ctx context.Context, obj runtime.Object, _ rest.ValidateObjectFunc, _ *metav1.CreateOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

func (r *supportBundleREST) New() runtime.Object {
	_ = "STUB: not implemented"
	return *new(runtime.Object)
}

func (r *supportBundleREST) Destroy() {
	_ = "STUB: not implemented"

	// Get returns current status of the bundle. It only allows querying the resource
	// whose name is equal to the mode.
	return
}

func (r *supportBundleREST) Get(_ context.Context, name string, _ *metav1.GetOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

// Delete can remove the current finished bundle or cancel a running bundle
// collecting. It only allows querying the resource whose name is equal to the mode.
func (r *supportBundleREST) Delete(_ context.Context, name string, _ rest.ValidateObjectFunc, _ *metav1.DeleteOptions) (runtime.Object, bool, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), false, nil
}

func (r *supportBundleREST) NamespaceScoped() bool { _ = "STUB: not implemented"; return false }

func (r *supportBundleREST) collect(ctx context.Context, dumpers ...func(string) error) (*systemv1beta1.SupportBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *supportBundleREST) collectAgent(ctx context.Context, since string) (*systemv1beta1.SupportBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *supportBundleREST) collectController(ctx context.Context, since string) (*systemv1beta1.SupportBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *supportBundleREST) clean(ctx context.Context, bundlePath string, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// check the context again in case of cancellation when acquiring the lock.

func (r *supportBundleREST) GetSingularName() string { _ = "STUB: not implemented"; return "" }

var (
	_ rest.Storage         = new(downloadREST)
	_ rest.Getter          = new(downloadREST)
	_ rest.StorageMetadata = new(downloadREST)
)

// downloadREST implements the REST for downloading the bundle.
type downloadREST struct {
	supportBundle *supportBundleREST
}

func (d *downloadREST) New() runtime.Object { _ = "STUB: not implemented"; return *new(runtime.Object) }

func (d *downloadREST) Destroy() { _ = "STUB: not implemented"; return }

func (d *downloadREST) Get(_ context.Context, _ string, _ *metav1.GetOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

func (d *downloadREST) ProducesMIMETypes(_ string) []string { _ = "STUB: not implemented"; return nil }

func (d *downloadREST) ProducesObject(_ string) interface{} { _ = "STUB: not implemented"; return nil }

var (
	_ rest.ResourceStreamer = new(bundleStream)
	_ runtime.Object        = new(bundleStream)
)

type bundleStream struct {
	cache *systemv1beta1.SupportBundle
}

func (b *bundleStream) GetObjectKind() schema.ObjectKind {
	_ = "STUB: not implemented"
	return *new(schema.ObjectKind)
}

func (b *bundleStream) DeepCopyObject() runtime.Object {
	_ = "STUB: not implemented"
	return *new(runtime.Object)
}

func (b *bundleStream) InputStream(_ context.Context, _, _ string) (stream io.ReadCloser, flush bool, mimeType string, err error) {
	_ = "STUB: not implemented"
	// f will be closed by invoker, no need to close in this function.
	return *new(io.ReadCloser), false, "", nil
}
