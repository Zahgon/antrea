// Copyright 2024 Antrea Authors
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

package nodelatencystats

import (
	"context"

	"k8s.io/utils/clock"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/client-go/tools/cache"
)

type REST struct {
	indexer cache.Indexer
	clock   clock.Clock
}

var (
	_ rest.Storage              = &REST{}
	_ rest.Scoper               = &REST{}
	_ rest.Getter               = &REST{}
	_ rest.Lister               = &REST{}
	_ rest.GracefulDeleter      = &REST{}
	_ rest.SingularNameProvider = &REST{}
)

// NewREST returns a REST object that will work against API services.
func NewREST() *REST { _ = "STUB: not implemented"; return nil }

func newRESTWithClock(clock clock.Clock) *REST { _ = "STUB: not implemented"; return nil }

func (r *REST) New() runtime.Object { _ = "STUB: not implemented"; return *new(runtime.Object) }

func (r *REST) Destroy() { _ = "STUB: not implemented"; return }

func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	// Update will add the object if the key does not exist.
	return *new(runtime.Object), nil
}

func (r *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

func (r *REST) NewList() runtime.Object { _ = "STUB: not implemented"; return *new(runtime.Object) }

func (r *REST) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *

	// Due to the unordered nature of map iteration and the complexity of controlling 'continue',
	// we will ignore paging here and plan to implement it in the future.
	new(runtime.Object), nil
}

func (r *REST) ConvertToTable(ctx context.Context, obj runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate the max and average latency values.

// Due to int64 max value is enough for the sum of all latencies,
// we don't need to check overflow in this case.

func (r *REST) Delete(ctx context.Context, name string, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	_ = "STUB: not implemented"
	// Ignore the deleteValidation and options for now.
	return *new(runtime.Object), false, nil
}

func (r *REST) NamespaceScoped() bool { _ = "STUB: not implemented"; return false }

func (r *REST) GetSingularName() string { _ = "STUB: not implemented"; return "" }
