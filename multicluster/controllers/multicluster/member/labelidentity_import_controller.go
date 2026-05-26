/*
Copyright 2022 Antrea Authors.

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

	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	multiclusterv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/commonarea"
)

// LabelIdentityResourceImportReconciler reconciles a LabelIdentity kind of ResourceImport object in the member cluster.
type LabelIdentityResourceImportReconciler struct {
	localClusterClient client.Client
	localClusterID     string
	namespace          string
	remoteCommonArea   commonarea.RemoteCommonArea
	// Saved Manager to indicate SetupWithManager() is done or not.
	manager ctrl.Manager
}

func newLabelIdentityResourceImportReconciler(localClusterClient client.Client,
	localClusterID string, namespace string, remoteCommonArea commonarea.RemoteCommonArea) *LabelIdentityResourceImportReconciler {
	_ = "STUB: not implemented"
	return nil
}

func (r *LabelIdentityResourceImportReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func (r *LabelIdentityResourceImportReconciler) handleLabelIdentityImpCreateOrUpdate(ctx context.Context,
	labelResImp *multiclusterv1alpha1.ResourceImport) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func (r *LabelIdentityResourceImportReconciler) handleLabelIdentityImpDelete(ctx context.Context,
	resImpName types.NamespacedName) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func (r *LabelIdentityResourceImportReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil

	// SetupWithManager was called by a previous RemoteManager.StartWatch() call and already
	// completed with no error.
}

// Ignore status update event via GenerationChangedPredicate

// Only register this controller to reconcile LabelIdentity kind of ResourceImport
