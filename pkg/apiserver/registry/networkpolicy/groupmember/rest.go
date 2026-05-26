// Copyright 2023 Antrea Authors
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

package groupmember

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"

	"antrea.io/antrea/v2/pkg/apiserver/registry/networkpolicy/clustergroupmember"
)

type REST struct {
	querier clustergroupmember.GroupMembershipQuerier
}

var (
	_ rest.Storage              = &REST{}
	_ rest.Scoper               = &REST{}
	_ rest.GetterWithOptions    = &REST{}
	_ rest.SingularNameProvider = &REST{}
)

// NewREST returns a REST object that will work against API services.
func NewREST(querier clustergroupmember.GroupMembershipQuerier) *REST {
	_ = "STUB: not implemented"
	return nil
}

func (r *REST) New() runtime.Object { _ = "STUB: not implemented"; return *new(runtime.Object) }

func (r *REST) Destroy() { _ = "STUB: not implemented"; return }

func (r *REST) Get(ctx context.Context, name string, options runtime.Object) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

// NewGetOptions returns the default options for Get, so options object is never nil.
func (r *REST) NewGetOptions() (runtime.Object, bool, string) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), false, ""
}

func (r *REST) NamespaceScoped() bool { _ = "STUB: not implemented"; return false }

func (r *REST) GetSingularName() string { _ = "STUB: not implemented"; return "" }
