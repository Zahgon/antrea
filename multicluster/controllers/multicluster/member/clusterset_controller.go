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
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mcv1alpha2 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha2"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/common"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/commonarea"
)

type leaderClusterInfo struct {
	clusterID  string
	serverUrl  string
	secretName string
}

var getRemoteConfigAndClient = commonarea.GetRemoteConfigAndClient

// MemberClusterSetReconciler reconciles a ClusterSet object in the member cluster deployment.
type MemberClusterSetReconciler struct {
	client.Client
	scheme                   *runtime.Scheme
	namespace                string
	clusterCalimCRDAvailable bool

	// commonAreaLock protects the access to RemoteCommonArea.
	commonAreaLock       sync.RWMutex
	commonAreaCreationCh chan struct{}

	clusterSetID    common.ClusterSetID
	clusterID       common.ClusterID
	installedLeader leaderClusterInfo

	remoteCommonArea             commonarea.RemoteCommonArea
	enableStretchedNetworkPolicy bool
}

func NewMemberClusterSetReconciler(client client.Client,
	scheme *runtime.Scheme,
	namespace string,
	enableStretchedNetworkPolicy bool,
	clusterCalimCRDAvailable bool,
	commonAreaCreationCh chan struct{},
) *MemberClusterSetReconciler {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=clustersets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=clustersets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=clustersets/finalizers,verbs=update

// Reconcile ClusterSet changes
func (r *MemberClusterSetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Not the current ClusterSet.

// Handle create or update

// ClusterSet deletion may fail and retry, but a ClusterSet may have been created just right before the retry.
// In that case, ClusterSet deletion retry will be like an update action, so try to delete stale resources
// here before initilizing a new ClusterSet.

// ClusterID is a required field, and the empty value case should only happen
// when Antrea Multi-cluster is upgraded from an old version prior to v1.13.
// Here we try to update the ClusterSet's ClusterID when it's configured in an
// existing ClusterClaim.

// The CommonArea creation succeeded here and so notify StaleController to
// clean up stale imported resources and ResourceExports.

// The notification has been sent and hasn't been consumed yet,
// no need to send another one.

func (r *MemberClusterSetReconciler) cleanUpResources(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Any ResourceExports belong to this member cluster will be cleaned up by the leader cluster
// when the MemberClusterAnnounce is deleted.

// MemberClusterAnnounce could be kept in the leader cluster, if antrea-mc-controller crashes after the failure.
// Leader cluster will delete the stale MemberClusterAnnounce with a garbage collection mechanism in this case.

func (r *MemberClusterSetReconciler) deleteMemberClusterAnnounce(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MemberClusterSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	// Update status periodically
	return nil
}

// Ignore status update event via GenerationChangedPredicate

func (r *MemberClusterSetReconciler) createRemoteCommonArea(clusterSet *mcv1alpha2.ClusterSet) error {
	_ = "STUB: not implemented"
	return nil
}

// Read Secret to access the leader cluster. Assume Secret is present in the same Namespace as the ClusterSet.

// Create import reconcilers and add them to RemoteCommonArea (to be started with
// RemoteCommonArea.StartWatching).

// getSecretForLeader returns the Secret associated with this local cluster(which is a member)
// for the given leader.
// When a member is added to a ClusterSet, a specific ServiceAccount is created on the
// leader cluster which allows the member access into the CommonArea. This ServiceAccount
// has an associated Secret which must be copied into the member cluster as an opaque secret.
// Name of this secret is part of the ClusterSet spec for this leader. This method reads
// the Secret given by that name.
func (r *MemberClusterSetReconciler) getSecretForLeader(secretName string, secretNs string) (secretObj *v1.Secret, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *MemberClusterSetReconciler) updateStatus() { _ = "STUB: not implemented"; return }

// Nothing to do.

// The total cluster should always be 1 to include the member cluster itself.

// SetRemoteCommonArea is for testing only
func (r *MemberClusterSetReconciler) SetRemoteCommonArea(commanArea commonarea.RemoteCommonArea) commonarea.RemoteCommonArea {
	_ = "STUB: not implemented"
	return *new(commonarea.RemoteCommonArea)
}

func (r *MemberClusterSetReconciler) GetRemoteCommonAreaAndLocalID() (commonarea.RemoteCommonArea, string, error) {
	_ = "STUB: not implemented"
	return *new(commonarea.RemoteCommonArea), "", nil
}
