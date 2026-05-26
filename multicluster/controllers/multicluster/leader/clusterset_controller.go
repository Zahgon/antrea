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
	"sync"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mcv1alpha2 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha2"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/common"
)

var (
	NoReadyCluster = "NoReadyCluster"
)

// LeaderClusterSetReconciler reconciles a ClusterSet object in the leader cluster deployment.
// Each ClusterSet should have one Multi-cluster Controller running in the ClusterSet' leader
// Namespace, so a MC Controller will be handling only a single ClusterSet in the given Namespace.
type LeaderClusterSetReconciler struct {
	client.Client
	namespace                string
	clusterCalimCRDAvailable bool
	statusManager            MemberClusterStatusManager

	clusterSetID common.ClusterSetID
	clusterID    common.ClusterID
	mutex        sync.Mutex
}

func NewLeaderClusterSetReconciler(client client.Client, namespace string,
	clusterCalimCRDAvailable bool,
	statusManager MemberClusterStatusManager) *LeaderClusterSetReconciler {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=clustersets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=clustersets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=clustersets/finalizers,verbs=update

// Reconcile ClusterSet changes
func (r *LeaderClusterSetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Not the current ClusterSet.

// Handle create or update

// ClusterID is a required field, and the empty value case should only happen
// when Antrea Multi-cluster is upgraded from an old version prior to v1.13.
// Here we try to update the ClusterSet's ClusterID when it's configured in an
// existing ClusterClaim.

// SetupWithManager sets up the controller with the Manager.
func (r *LeaderClusterSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil

	// Ignore status update event via GenerationChangedPredicate
}

func (r *LeaderClusterSetReconciler) runBackgroundTasks() {
	_ = "STUB: not implemented"
	// Update status periodically
	return
}

// updateStatus updates ClusterSet Status as follows:
//  1. TotalClusters is the number of member clusters in the
//     ClusterSet resource last processed.
//  2. ObservedGeneration is the Generation from the last processed
//     ClusterSet resource.
//  3. Individual cluster status is obtained from MemberClusterAnnounce
//     controller.
//  4. ReadyClusters is the number of member clusters with "Ready" = "True"
//  5. Overall condition of the ClusterSet is also computed as follows:
//     a. "Ready" = "True" if all clusters have "Ready" = "True".
//     Message & Reason will be absent.
//     b. "Ready" = "Unknown" if all clusters have "Ready" = "Unknown".
//     Message will be "All clusters have an unknown status"
//     and Reason will be "NoReadyCluster"
//     c. "Ready" = "False" for any other combination of cluster
//     statues across all clusters. Message will be empty and Reason
//     will be "NoReadyCluster"
func (r *LeaderClusterSetReconciler) updateStatus() { _ = "STUB: not implemented"; return }

// Nothing to do.

func validateMemberClusterExists(clusterID common.ClusterID, clusters []mcv1alpha2.LeaderClusterInfo) (err error) {
	_ = "STUB: not implemented"
	return nil
}
