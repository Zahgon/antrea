/*
Copyright 2021 Antrea Authors.

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

package leader

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	mcs "sigs.k8s.io/mcs-api/pkg/apis/v1alpha1"

	mcsv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
)

type (
	// ResourceExportReconciler reconciles a ResourceExport object in the leader cluster.
	ResourceExportReconciler struct {
		client.Client
		Scheme *runtime.Scheme
	}
)

type resReason int

const (
	succeed resReason = iota
	failed
)

func NewResourceExportReconciler(
	client client.Client,
	scheme *runtime.Scheme) *ResourceExportReconciler {
	_ = "STUB: not implemented"
	return nil
}

// +kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=resourceexports,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=resourceexports/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=resourceexports/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// Reconcile will process all kinds of ResourceExport. Service and Endpoint kinds of ResourceExport
// will be handled in this file, and all other kinds will have their own handler files, eg: newkind_handler.go
func (r *ResourceExportReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// We are using Finalizers to implement asynchronous pre-delete hooks.
// When a ResourceExport is deleted, it will have non-zero DeletionTimestamp
// but controller can still get the deleted ResourceExport object, so it can
// clean up any replicated resources like ResourceImport.
// For more details about using Finalizers, please refer to https://book.kubebuilder.io/reference/using-finalizers.html.

// There might be some changes from ResourceExport triggered reconciling but actually no need to update
// ResourceImport, we still need to update the ResourceExport status to reflect event have been handled
// successfully.

func (r *ResourceExportReconciler) handleUpdateEvent(ctx context.Context,
	resImport *mcsv1alpha1.ResourceImport, resExport *mcsv1alpha1.ResourceExport) error {
	_ = "STUB: not implemented"
	return nil
}

// handleDeleteEvent will either delete the corrsponding ResourceImport if no more ResourceExport exists
// or regenerate ResourceImport's Subsets from latest ResourceExports without Endpoints from
// the deleted ResourceExport.
func (r *ResourceExportReconciler) handleDeleteEvent(ctx context.Context, resExport *mcsv1alpha1.ResourceExport) error {
	_ = "STUB: not implemented"
	return nil
}

// should update ResourceImport status when one of ResourceExports is removed?

func (r *ResourceExportReconciler) cleanUpResourceImport(ctx context.Context,
	resImp types.NamespacedName, re interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResourceExportReconciler) updateEndpointResourceImport(ctx context.Context,
	existRe *mcsv1alpha1.ResourceExport, resImpName types.NamespacedName) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResourceExportReconciler) getExistingResImport(ctx context.Context,
	resExport mcsv1alpha1.ResourceExport) (bool, *mcsv1alpha1.ResourceImport, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// refreshServiceResourceImport returns a new Service kind of ResourceImport or
// updates existing one to reflect any change from member cluster's ResourceExport.
func (r *ResourceExportReconciler) refreshServiceResourceImport(
	resExport *mcsv1alpha1.ResourceExport,
	resImport *mcsv1alpha1.ResourceImport,
	createResImport bool) (*mcsv1alpha1.ResourceImport, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// TODO: check ClusterIPs difference if it is being used in ResrouceImport later

// When there is only one Service ResourceExport, ResourceImport should reflect the change
// otherwise, it should always return error so controller can retry later assuming users can fix the conflicts

// refreshEndpointsResourceImport returns a new Endpoints kind of ResourceImport or
// updates existing one to reflect any change from member cluster's ResourceExport
func (r *ResourceExportReconciler) refreshEndpointsResourceImport(
	resExport *mcsv1alpha1.ResourceExport,
	resImport *mcsv1alpha1.ResourceImport,
	createResImport bool) (*mcsv1alpha1.ResourceImport, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// check corresponding Service type of ResourceExport, if there is any failure,
// skip adding Endpoints of this ResourceExport and update Endpoint type of
// ResourceExport's status.

// check all matched Endpoints ResourceExport and generate a new EndpointSubset

func (r *ResourceExportReconciler) refreshACNPResourceImport(
	resExport *mcsv1alpha1.ResourceExport,
	resImport *mcsv1alpha1.ResourceImport,
	createResImport bool) (*mcsv1alpha1.ResourceImport, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (r *ResourceExportReconciler) getNotDeletedResourceExports(resExport *mcsv1alpha1.ResourceExport) ([]mcsv1alpha1.ResourceExport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ResourceExportReconciler) updateResourceExportStatus(resExport *mcsv1alpha1.ResourceExport, res resReason) {
	_ = "STUB: not implemented"
	return
}

// deleteResourceExport removes ResourceExport finalizer string and updates it, so Kubernetes can complete deletion.
func (r *ResourceExportReconciler) deleteResourceExport(resExport *mcsv1alpha1.ResourceExport) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ResourceExportReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	// Ignore status update event via GenerationChangedPredicate
	return nil
}

// Register this controller to ignore LabelIdentity kind of ResourceExport

func getLabelSelector(resExport *mcsv1alpha1.ResourceExport) labels.Selector {
	_ = "STUB: not implemented"
	return *new(labels.Selector)
}

func SvcPortsConverter(svcPort []corev1.ServicePort) []mcs.ServicePort {
	_ = "STUB: not implemented"
	return nil
}

func GetResourceImportName(resExport *mcsv1alpha1.ResourceExport) types.NamespacedName {
	_ = "STUB: not implemented"
	return *new(types.NamespacedName)
}

// We use finalizers as ResourceExport pre-delete hooks, which means when
// we list the ResourceExports, it will also return deleted items.
// RemoveDeletedResourceExports remove any ResourceExports with non-zero DeletionTimestamp
// which is actually deleted object.
func RemoveDeletedResourceExports(items []mcsv1alpha1.ResourceExport) []mcsv1alpha1.ResourceExport {
	_ = "STUB: not implemented"
	return nil
}
