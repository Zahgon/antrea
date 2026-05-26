/*
Copyright 2023 Antrea Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package member

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	k8smcv1alpha1 "sigs.k8s.io/mcs-api/pkg/apis/v1alpha1"

	mcv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/commonarea"
	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

// StaleResCleanupController will clean up ServiceImport, MC Service, ACNP, ClusterInfoImport and LabelIdentity
// resources if no corresponding ResourceImports in the leader cluster and remove stale ResourceExports
// in the leader cluster if no corresponding ServiceExport or Gateway in the member cluster when it runs in
// the member cluster. StaleResCleanupController one-time runner will run only once in the member cluster
// during Multi-cluster Controller starts, and it will retry only if there is an error.
// StaleResCleanupController's reconciler will handle ClusterSet deletion event to clean up all
// automatically created resources for the ClusterSet.
type StaleResCleanupController struct {
	client.Client
	Scheme               *runtime.Scheme
	commonAreaCreationCh chan struct{}
	localClusterID       string
	commonAreaGetter     commonarea.RemoteCommonAreaGetter
	namespace            string
}

func NewStaleResCleanupController(
	Client client.Client,
	Scheme *runtime.Scheme,
	commonAreaCreationCh chan struct{},
	namespace string,
	commonAreaGetter commonarea.RemoteCommonAreaGetter,
) *StaleResCleanupController {
	_ = "STUB: not implemented"
	return nil
}

// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;delete
// +kubebuilder:rbac:groups=multicluster.x-k8s.io,resources=serviceimports,verbs=get;list;watch;delete
// +kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=resourceimports,verbs=get;list;watch;
// +kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=resourceexports,verbs=get;list;watch;delete

func (c *StaleResCleanupController) CleanUp(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *StaleResCleanupController) cleanUpStaleResourcesOnMember(ctx context.Context, commonArea commonarea.RemoteCommonArea) error {
	_ = "STUB: not implemented"
	return nil
}

// All previously imported resources need to be listed before ResourceImports are listed.
// This prevents race condition between stale_controller and other reconcilers.
// See https://github.com/antrea-io/antrea/issues/4854

// Clean up any imported Services that do not have corresponding ResourceImport anymore

// Clean up any imported ACNPs that do not have corresponding ResourceImport anymore

// Clean up any imported ClusterInfos that do not have corresponding ResourceImport anymore

// Clean up any imported LabelIdentities that do not have corresponding ResourceImport anymore

// Clean up stale ResourceExports in the leader cluster for a member cluster.
func (c *StaleResCleanupController) cleanUpStaleResourceExportsOnLeader(ctx context.Context, commonArea commonarea.RemoteCommonArea) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *StaleResCleanupController) cleanUpStaleServiceResources(ctx context.Context, svcImpList *k8smcv1alpha1.ServiceImportList,
	svcList *corev1.ServiceList, resImpList *mcv1alpha1.ResourceImportList) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *StaleResCleanupController) cleanUpACNPResources(ctx context.Context, acnpList *crdv1beta1.ClusterNetworkPolicyList,
	resImpList *mcv1alpha1.ResourceImportList) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *StaleResCleanupController) cleanUpClusterInfoImports(ctx context.Context, ciImpList *mcv1alpha1.ClusterInfoImportList,
	resImpList *mcv1alpha1.ResourceImportList) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *StaleResCleanupController) cleanUpLabelIdentities(ctx context.Context, labelIdentityList *mcv1alpha1.LabelIdentityList,
	resImpList *mcv1alpha1.ResourceImportList) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanUpServiceResourceExports removes any Service/Endpoint kind of ResourceExports when there is no
// corresponding ServiceExport in the local cluster.
func (c *StaleResCleanupController) cleanUpServiceResourceExports(ctx context.Context, commonArea commonarea.RemoteCommonArea, resExpList *mcv1alpha1.ResourceExportList) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *StaleResCleanupController) cleanUpLabelIdentityResourceExports(ctx context.Context, commonArea commonarea.RemoteCommonArea, resExpList *mcv1alpha1.ResourceExportList) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanUpClusterInfoResourceExports removes any ClusterInfo kind of ResourceExports when there is no
// Gateway in the local cluster.
func (c *StaleResCleanupController) cleanUpClusterInfoResourceExports(ctx context.Context, commonArea commonarea.RemoteCommonArea) error {
	_ = "STUB: not implemented"
	return nil
}

// Run starts the StaleResCleanupController and blocks until stopCh is closed.
func (c *StaleResCleanupController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *StaleResCleanupController) cleanUpStaleResources(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Clean up stale ResourceExports in the leader cluster for a member cluster.

func cleanUpResourcesCreatedByMC(ctx context.Context, mgrClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanUpMCServicesAndServiceImports(ctx context.Context, mgrClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanUpReplicatedACNPs(ctx context.Context, mgrClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanUpLabelIdentities(ctx context.Context, mgrClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanUpClusterInfoImports(ctx context.Context, mgrClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanUpGateways(ctx context.Context, mgrClient client.Client) error {
	_ = "STUB: not implemented"
	return nil
}
