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
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	mcv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
	mcv1alpha2 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha2"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/common"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/commonarea"
)

var (
	ServiceCIDRDiscoverFn = common.DiscoverClusterServiceCIDR

	statusReadyPredicateFunc = func(e event.UpdateEvent) bool {
		if e.ObjectOld == nil || e.ObjectNew == nil {
			return false
		}
		oldClusterSet := e.ObjectOld.(*mcv1alpha2.ClusterSet)
		newClusterSet := e.ObjectNew.(*mcv1alpha2.ClusterSet)
		oldConditionSize := len(oldClusterSet.Status.Conditions)
		newConditionSize := len(newClusterSet.Status.Conditions)
		if oldConditionSize == 0 && newConditionSize > 0 && newClusterSet.Status.Conditions[0].Status == corev1.ConditionTrue {
			return true
		}
		if oldConditionSize > 0 && newConditionSize > 0 &&
			(oldClusterSet.Status.Conditions[0].Status == corev1.ConditionFalse || oldClusterSet.Status.Conditions[0].Status == corev1.ConditionUnknown) &&
			newClusterSet.Status.Conditions[0].Status == corev1.ConditionTrue {
			return true
		}
		return false
	}

	statusReadyPredicate = predicate.Funcs{
		UpdateFunc: statusReadyPredicateFunc,
	}
)

type (
	// NodeReconciler is for member cluster only.
	NodeReconciler struct {
		client.Client
		Scheme             *runtime.Scheme
		namespace          string
		precedence         mcv1alpha1.Precedence
		gatewayCandidates  map[string]bool
		activeGatewayMutex sync.Mutex
		commonAreaGetter   commonarea.RemoteCommonAreaGetter
		activeGateway      string
		serviceCIDR        string
		initialized        bool
	}
)

// NewNodeReconciler creates a NodeReconciler to watch Node resource changes.
// It's responsible for creating a Gateway for the first ready Node with
// annotation `multicluster.antrea.io/gateway:true` if there is no existing Gateway.
// It guarantees there is always only one Gateway CR when there are multiple Nodes
// with annotation `multicluster.antrea.io/gateway:true`.
func NewNodeReconciler(
	client client.Client,
	scheme *runtime.Scheme,
	namespace string,
	serviceCIDR string,
	precedence mcv1alpha1.Precedence,
	commonAreaGetter commonarea.RemoteCommonAreaGetter) *NodeReconciler {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=gateways,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=nodes,verbs=get;list;watch;
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=gateways/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=gateways/finalizers,verbs=update

func (r *NodeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// initialize initializes 'activeGateway' and 'gatewayCandidates' and removes
// stale Gateway during controller startup.
func (r *NodeReconciler) initialize() error { _ = "STUB: not implemented"; return nil }

// Gateway webhook guarantees that there is at most one Gateway in the member cluster.

func (r *NodeReconciler) updateActiveGateway(ctx context.Context, newGateway *mcv1alpha1.Gateway) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: cache might be stale. Need to revisit here and other reconcilers to
// check if we can improve this with 'Owns' or other methods.

// If the Gateway version in the client cache is stale, the update operation will fail,
// then the reconciler will retry with latest state again.

// recreateActiveGateway will delete the existing Gateway CR and create a new Gateway
// from the pool of Gateway candidates.
func (r *NodeReconciler) recreateActiveGateway(ctx context.Context, gateway *mcv1alpha1.Gateway) error {
	_ = "STUB: not implemented"
	return nil
}

// Check remaining Gateway candidates and create a new Gateway.

// getValidGatewayFromCandidates picks a valid Node from Gateway candidates and
// creates a Gateway. It returns no error if no good Gateway candidate.
func (r *NodeReconciler) getValidGatewayFromCandidates() (*mcv1alpha1.Gateway, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *NodeReconciler) createGateway(gateway *mcv1alpha1.Gateway) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *NodeReconciler) getGatawayNodeIP(node *corev1.Node) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *NodeReconciler) clusterSetMapFunc(ctx context.Context, a client.Object) []reconcile.Request {
	_ = "STUB: not implemented"
	return nil
}

// All auto-generated resources will be deleted by the ClusterSet controller when a ClusterSet is
// deleted, so here we can set the activeGateway to empty directly.

func isReadyNode(node *corev1.Node) bool { _ = "STUB: not implemented"; return false }
