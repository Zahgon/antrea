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

package commonarea

import (
	"context"
	"sync"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	mcv1alpha2 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha2"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/common"
)

const (
	TimestampAnnotationKey = "touch-ts"
)

var (
	ReasonDisconnected = "Disconnected"
)

// remoteCommonArea implements the CommonArea interface and allows local cluster to read/write into
// the CommonArea of RemoteCommonArea.
type remoteCommonArea struct {
	// mutex to synchronize access to connectivity state since it is updated by
	// a background routine running in remoteCommonArea.
	mutex sync.RWMutex

	// client that provides read/write access into the remoteCommonArea.
	client.Client

	// ClusterManager to set up controllers for resources that need to be monitored in the remoteCommonArea.
	ClusterManager manager.Manager

	// ClusterID of this remoteCommonArea.
	ClusterID common.ClusterID

	// ClusterSetID of this remoteCommonArea.
	ClusterSetID common.ClusterSetID

	// config necessary to access the remoteCommonArea.
	config *rest.Config

	// scheme necessary to access the remoteCommonArea.
	scheme *runtime.Scheme

	// Namespace this ClusterSet is associated with.
	Namespace string

	// connected is a state to know whether the remoteCommonArea is connected or not.
	connected bool

	clusterStatus mcv1alpha2.ClusterCondition
	leaderStatus  mcv1alpha2.ClusterCondition

	// The ID of the local member cluster
	localClusterID common.ClusterID

	// client that provides read/write access into the local cluster
	localClusterClient client.Client

	// localNamespace is the Namespace where the controller is running.
	localNamespace string

	// stopFunc to stop all background operations when the RemoteCommonArea is stopped.
	stopFunc context.CancelFunc

	// managerStopFunc to stop the manager when the RemoteCommonArea is stopped.
	managerStopFunc context.CancelFunc

	// Enable StretchedNetworkPolicy which will export and import labelIdentities in the
	// ClusterSet and allow Antrea-native policies to select peers from other clusters
	// in a ClusterSet.
	enableStretchedNetworkPolicy bool

	// A list of ImportReconcilers to reconcile ResourceImports.
	importReconcilers []ImportReconciler
}

// NewRemoteCommonArea returns a RemoteCommonArea instance which will use access credentials from the Secret to
// connect to the leader cluster's CommonArea.
func NewRemoteCommonArea(clusterID common.ClusterID, clusterSetID common.ClusterSetID, localClusterID common.ClusterID, mgr manager.Manager, remoteClient client.Client,
	scheme *runtime.Scheme, localClusterClient client.Client, clusterSetNamespace string, localNamespace string, config *rest.Config, enableStretchedNetworkPolicy bool) (RemoteCommonArea, error) {
	_ = "STUB: not implemented"
	return *new(RemoteCommonArea), nil
}

func GetRemoteConfigAndClient(secretObj *v1.Secret, url string, clusterID common.ClusterID, clusterSet *mcv1alpha2.ClusterSet, scheme *runtime.Scheme) (*rest.Config,
	manager.Manager, client.Client, error) {
	_ = "STUB: not implemented"
	return nil, *new(manager.Manager), *new(client.Client), nil
}

/**
 * GetSecretCACrtAndToken returns the access credentials from Secret.
 */
func getSecretCACrtAndToken(secretObj *v1.Secret) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (r *remoteCommonArea) SendMemberAnnounce() error { _ = "STUB: not implemented"; return nil }

// Add timestamp to force update on MemberClusterAnnounce. Leader cluster requires
// periodic updates to detect connectivity. Without this, no-op updates will be ignored.

// Create happens first before the leader validation passes. When the creation is successful,
// it marks the connectivity status and then the validation on the leader can happen.

func (r *remoteCommonArea) updateRemoteCommonAreaStatus(connected bool, err error) {
	_ = "STUB: not implemented"
	return
}

// TODO: Tolerate transient failures so we dont oscillate between connected and disconnected.

func (r *remoteCommonArea) updateLeaderStatus() { _ = "STUB: not implemented"; return }

/**
 * ---------------------------
 * CommonArea Implementation
 * ---------------------------
 */

func (r *remoteCommonArea) GetClusterID() common.ClusterID {
	_ = "STUB: not implemented"
	return *new(common.ClusterID)
}

func (r *remoteCommonArea) GetNamespace() string {
	_ = "STUB: not implemented"

	/**
	 * ---------------------------
	 * RemoteCommonArea Implementation
	 * ---------------------------
	 */return ""
}

// Start starts a background routine.
// Once connected to the RemoteCommonArea, the Start method runs a timer
// on a go routine to periodically write MemberClusterAnnounce into the
// RemoteCommonArea's CommonArea and also maintain its connectivity status.
func (r *remoteCommonArea) Start() context.CancelFunc {
	_ = "STUB: not implemented"
	return *new(context.CancelFunc)
}

// Will retry in next tick.

func (r *remoteCommonArea) doMemberAnnounce() { _ = "STUB: not implemented"; return }

func (r *remoteCommonArea) Stop() { _ = "STUB: not implemented"; return }

func (r *remoteCommonArea) IsConnected() bool { _ = "STUB: not implemented"; return false }

func (r *remoteCommonArea) AddImportReconciler(reconciler ImportReconciler) {
	_ = "STUB: not implemented"
	return
}

func (r *remoteCommonArea) StartWatching() error { _ = "STUB: not implemented"; return nil }

// This starts the Manager and blocks; Manager performs reconciliation of resources from the RemoteCommonArea.
// When this RemoteCommonArea is not a leader anymore, stopCtx will be closed in StopWatching,
// so this blocking routine can return and finish. And the next time this RemoteCommonArea is connected as
// the leader again, it starts the Manager again.

func (r *remoteCommonArea) StopWatching() { _ = "STUB: not implemented"; return }

func (r *remoteCommonArea) GetStatus() []mcv1alpha2.ClusterCondition {
	_ = "STUB: not implemented"
	return nil
}

// This will be a copy
// This will be a copy

func (r *remoteCommonArea) GetLocalClusterID() string { _ = "STUB: not implemented"; return "" }
