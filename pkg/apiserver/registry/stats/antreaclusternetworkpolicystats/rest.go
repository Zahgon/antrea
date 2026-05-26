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

package antreaclusternetworkpolicystats

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"

	statsv1alpha1 "antrea.io/antrea/v2/pkg/apis/stats/v1alpha1"
)

var (
	tableColumnDefinitions = []metav1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: swaggerMetadataDescriptions["name"]},
		{Name: "Sessions", Type: "integer", Description: "The sessions count hit by the Antrea ClusterNetworkPolicy."},
		{Name: "Packets", Type: "integer", Description: "The packets count hit by the Antrea ClusterNetworkPolicy."},
		{Name: "Bytes", Type: "integer", Description: "The bytes count hit by the Antrea ClusterNetworkPolicy."},
		{Name: "Created At", Type: "date", Description: swaggerMetadataDescriptions["creationTimestamp"]},
	}
)

type REST struct {
	statsProvider statsProvider
}

// NewREST returns a REST object that will work against API services.
func NewREST(p statsProvider) *REST { _ = "STUB: not implemented"; return nil }

var (
	_ rest.Storage              = &REST{}
	_ rest.Scoper               = &REST{}
	_ rest.Getter               = &REST{}
	_ rest.Lister               = &REST{}
	_ rest.SingularNameProvider = &REST{}
)

type statsProvider interface {
	ListAntreaClusterNetworkPolicyStats() []statsv1alpha1.AntreaClusterNetworkPolicyStats

	GetAntreaClusterNetworkPolicyStats(name string) (*statsv1alpha1.AntreaClusterNetworkPolicyStats, bool)
}

func (r *REST) New() runtime.Object { _ = "STUB: not implemented"; return *new(runtime.Object) }

func (r *REST) Destroy() { _ = "STUB: not implemented"; return }

func (r *REST) NewList() runtime.Object { _ = "STUB: not implemented"; return *new(runtime.Object) }

func (r *REST) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

func (r *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

var swaggerMetadataDescriptions = metav1.ObjectMeta{}.SwaggerDoc()

func formatTimestamp(t metav1.Time) string { _ = "STUB: not implemented"; return "" }

func (r *REST) ConvertToTable(ctx context.Context, obj runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *REST) NamespaceScoped() bool { _ = "STUB: not implemented"; return false }

func (r *REST) GetSingularName() string { _ = "STUB: not implemented"; return "" }
