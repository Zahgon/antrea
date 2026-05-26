// Copyright 2021 Antrea Authors
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

package clustergroupmember

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"

	"antrea.io/antrea/v2/pkg/apis/controlplane"
)

type REST struct {
	querier GroupMembershipQuerier
}

var (
	_ rest.Storage              = &REST{}
	_ rest.Scoper               = &REST{}
	_ rest.GetterWithOptions    = &REST{}
	_ rest.SingularNameProvider = &REST{}
)

// NewREST returns a REST object that will work against API services.
func NewREST(querier GroupMembershipQuerier) *REST { _ = "STUB: not implemented"; return nil }

type GroupMembershipQuerier interface {
	GetGroupMembers(name string) (controlplane.GroupMemberSet, []controlplane.IPBlock, error)
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

func GetPaginatedMembers(querier GroupMembershipQuerier, name string, options runtime.Object) (members []controlplane.GroupMember, ipNets []controlplane.IPNet, totalMembers, totalPages, currentPage int64, err error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, 0, 0, nil
}

// Retrieve options used for pagination.

// ClusterGroup ipBlock does not support Except slices, so no need to generate an effective
// list of IPs by removing Except slices from allowed CIDR.

// PaginateMemberList returns paginated results if meaningful options are provided. Options should never be nil.
// Paginated results are continuous only when there is no member change across multiple calls.
// Pagination is not enabled if either page number or limit = 0, in which the full member list is returned.
// An error is returned for invalid options, and an empty list is returned for a page number out of the pages range.
func PaginateMemberList(effectiveMembers *[]controlplane.GroupMember, pageInfo *controlplane.PaginationGetOptions) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// Sort members based on EE/Pod names to realize consistent pagination results.

// Returns an empty member list if the page number exceeds total pages, to indicate end of list.
