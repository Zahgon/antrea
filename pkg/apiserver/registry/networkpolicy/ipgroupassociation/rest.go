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

package ipgroupassociation

import (
	"context"
	"net"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
	coreinformers "k8s.io/client-go/informers/core/v1"

	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha2"
	ga "antrea.io/antrea/v2/pkg/apiserver/registry/networkpolicy/groupassociation"
	crdv1a2informers "antrea.io/antrea/v2/pkg/client/informers/externalversions/crd/v1alpha2"
	"antrea.io/antrea/v2/pkg/controller/types"
)

type REST struct {
	podInformer  coreinformers.PodInformer
	eeInformer   crdv1a2informers.ExternalEntityInformer
	ipbQuerier   ipBlockGroupAssociationQuerier
	groupQuerier ga.GroupAssociationQuerier
	nodeInformer coreinformers.NodeInformer
}

var (
	_ rest.Storage              = &REST{}
	_ rest.Scoper               = &REST{}
	_ rest.Getter               = &REST{}
	_ rest.SingularNameProvider = &REST{}
)

// NewREST returns a REST object that will work against API services.
func NewREST(podQuerier coreinformers.PodInformer,
	nodeInformer coreinformers.NodeInformer,
	eeQuerier crdv1a2informers.ExternalEntityInformer,
	ipbQuerier ipBlockGroupAssociationQuerier,
	groupQuerier ga.GroupAssociationQuerier) *REST {
	_ = "STUB: not implemented"
	return nil
}

type ipBlockGroupAssociationQuerier interface {
	GetAssociatedIPBlockGroups(ip net.IP) []types.Group
}

func (r *REST) New() runtime.Object { _ = "STUB: not implemented"; return *new(runtime.Object) }

func (r *REST) Destroy() { _ = "STUB: not implemented"; return }

func (r *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

// The namespace/name might correspond to an ExternalEntity rather than a Pod

// Check for ipBlock group association no matter the Pod/EE query result is

func (r *REST) getAssociatedPod(ip string) *v1.Pod { _ = "STUB: not implemented"; return nil }

func (r *REST) getAssociatedNode(ip string) *v1.Node { _ = "STUB: not implemented"; return nil }

func (r *REST) getAssociatedExternalEntities(ip string) []*v1alpha2.ExternalEntity {
	_ = "STUB: not implemented"
	return nil
}

func (r *REST) NamespaceScoped() bool { _ = "STUB: not implemented"; return false }

func (r *REST) GetSingularName() string { _ = "STUB: not implemented"; return "" }
