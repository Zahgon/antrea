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

package member

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	k8smcsv1alpha1 "sigs.k8s.io/mcs-api/pkg/apis/v1alpha1"

	multiclusterv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/commonarea"
)

const (
	// cached indexer
	resImportIndexer = "name.kind"
)

func resImportIndexerFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resImportIndexerKeyFunc(obj interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ResourceImportReconciler reconciles a ResourceImport object in the member cluster.
type ResourceImportReconciler struct {
	localClusterClient  client.Client
	localClusterID      string
	namespace           string
	remoteCommonArea    commonarea.RemoteCommonArea
	installedResImports cache.Indexer
	// Saved Manager to indicate SetupWithManager() is done or not.
	manager ctrl.Manager
}

func newResourceImportReconciler(localClusterClient client.Client,
	localClusterID string, namespace string, remoteCommonArea commonarea.RemoteCommonArea) *ResourceImportReconciler {
	_ = "STUB: not implemented"
	return nil
}

// +kubebuilder:rbac:groups=crd.antrea.io,resources=clusternetworkpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=crd.antrea.io,resources=tiers,verbs=get;list;watch
// +kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=resourceimports,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=resourceimports/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=resourceimports/finalizers,verbs=update
// +kubebuilder:rbac:groups=multicluster.x-k8s.io,resources=serviceimports,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=multicluster.x-k8s.io,resources=serviceimports/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=endpoints,verbs=get;list;watch;update;create;patch;delete
// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;update;create;patch;delete
// +kubebuilder:rbac:groups="",resources=events,verbs=create

// Reconcile will attempt to ensure that the imported Resource is installed in local cluster as per the
// ResourceImport object.
func (r *ResourceImportReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// TODO: Must check whether this ResourceImport must be reconciled by this member cluster. Check `spec.clusters` field.

// stale_controller will reconcile and clean up MC Service/ServiceImport, so it's ok to return nil here

func (r *ResourceImportReconciler) handleResImpUpdateForService(ctx context.Context, resImp *multiclusterv1alpha1.ResourceImport) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Here we will skip creating derived MC Service when a Service with the same name
// already exists but it's not previously created by Importer.

// Ignore the error here, and requeue the event again when both Service
// and ServiceImport are created later

// Set multi-cluster Service's ClusterIP as ServiceImport's ClusterSetIP

// Requeue the event to update ServiceImport's ClusterSetIP

// TODO: check label difference ?

func (r *ResourceImportReconciler) handleResImpDeleteForService(ctx context.Context, resImp *multiclusterv1alpha1.ResourceImport) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func (r *ResourceImportReconciler) handleResImpUpdateForEndpoints(ctx context.Context, resImp *multiclusterv1alpha1.ResourceImport) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// TODO: check label difference ?

func (r *ResourceImportReconciler) handleResImpDeleteForEndpoints(ctx context.Context, resImp *multiclusterv1alpha1.ResourceImport) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func getMCService(resImp *multiclusterv1alpha1.ResourceImport) *corev1.Service {
	_ = "STUB: not implemented"
	return nil
}

func getMCServiceImport(resImp *multiclusterv1alpha1.ResourceImport) *k8smcsv1alpha1.ServiceImport {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithManager sets up the controller with the ClusterManager
// which will set up controllers for resources that need to be monitored
// in the remoteCommonArea.
func (r *ResourceImportReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil

	// SetupWithManager was called by a previous RemoteManager.StartWatch() call and already
	// completed with no error.
}

// Ignore status update event via GenerationChangedPredicate

// Register this filter to ignore LabelIdentity kind of ResourceImport
