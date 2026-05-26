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

package networkpolicyevaluation

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"

	"antrea.io/antrea/v2/pkg/controller/networkpolicy"
)

type REST struct {
	querier networkpolicy.PolicyRuleQuerier
}

var (
	_ rest.Storage              = &REST{}
	_ rest.Scoper               = &REST{}
	_ rest.Creater              = &REST{}
	_ rest.SingularNameProvider = &REST{}
)

// NewREST returns a REST object that will work against API services.
func NewREST(querier networkpolicy.PolicyRuleQuerier) *REST { _ = "STUB: not implemented"; return nil }

func (r *REST) New() runtime.Object { _ = "STUB: not implemented"; return *new(runtime.Object) }

func (r *REST) Destroy() { _ = "STUB: not implemented"; return }

func (r *REST) Create(ctx context.Context, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

func (r *REST) NamespaceScoped() bool { _ = "STUB: not implemented"; return false }

func (r *REST) GetSingularName() string { _ = "STUB: not implemented"; return "" }
