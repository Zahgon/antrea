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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	mcv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/commonarea"
)

type (
	// GatewayReconciler is for member cluster only.
	GatewayReconciler struct {
		client.Client
		Scheme           *runtime.Scheme
		commonAreaGetter commonarea.RemoteCommonAreaGetter
		namespace        string
		localClusterID   string
		podCIDRs         []string
		leaderNamespace  string
	}
)

// NewGatewayReconciler creates a GatewayReconciler which will watch Gateway events
// and create a ClusterInfo kind of ResourceExport in the leader cluster.
func NewGatewayReconciler(
	client client.Client,
	scheme *runtime.Scheme,
	namespace string,
	podCIDRs []string,
	commonAreaGetter commonarea.RemoteCommonAreaGetter) *GatewayReconciler {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=gateways,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=gateways/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=gateways/finalizers,verbs=update
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=clusterinfoimports,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=clusterinfoimports/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=clusterinfoimports/finalizers,verbs=update

func (r *GatewayReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// TODO: there is a possibility that the ClusterSet is to be deleted after getting the commonArea,
// then there might be new ResourceExports created in the leader after the member cluster
// is removed from the ClusterSet. We need to handle such corner cases in the future release.

// updateResourceExport will update latest Gateway information with the existing ResourceExport's resourceVersion.
// It will return an error and retry when there is a version conflict.

func (r *GatewayReconciler) updateResourceExport(ctx context.Context, req ctrl.Request,
	commonArea commonarea.RemoteCommonArea, existingResExport *mcv1alpha1.ResourceExport, gw *mcv1alpha1.Gateway) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *GatewayReconciler) createResourceExport(ctx context.Context, req ctrl.Request,
	commonArea commonarea.RemoteCommonArea, gateway *mcv1alpha1.Gateway) error {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GatewayReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: add a lock for r.serviceCIDR and r.localClusterID if
//  there is any plan to increase this concurrent number.

func (r *GatewayReconciler) clusterSetMapFunc(ctx context.Context, a client.Object) []reconcile.Request {
	_ = "STUB: not implemented"
	return nil
}

func (r *GatewayReconciler) getClusterInfo(gateway *mcv1alpha1.Gateway) *mcv1alpha1.ClusterInfo {
	_ = "STUB: not implemented"
	return nil
}
