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

package networkpolicy

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
)

// StatusREST implements the REST endpoint for getting NetworkPolicy's status.
type StatusREST struct {
	collector statusCollector
}

// NewStatusREST returns a REST object that will work against API services.
func NewStatusREST(collector statusCollector) *StatusREST { _ = "STUB: not implemented"; return nil }

// statusCollector is the interface required by the handler.
type statusCollector interface {
	UpdateStatus(status *controlplane.NetworkPolicyStatus) error
}

var _ rest.NamedCreater = &StatusREST{}

func (s StatusREST) New() runtime.Object { _ = "STUB: not implemented"; return *new(runtime.Object) }

func (s StatusREST) Destroy() { _ = "STUB: not implemented"; return }

func (s StatusREST) Create(ctx context.Context, name string, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}
