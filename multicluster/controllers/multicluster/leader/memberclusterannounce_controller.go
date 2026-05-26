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

// memberclusterannounce_controller is for leader cluster only.
package leader

import (
	"context"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mcv1alpha2 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha2"
	"antrea.io/antrea/v2/multicluster/controllers/multicluster/common"
)

var (
	ReasonConnected    = "Connected"
	ReasonDisconnected = "Disconnected"

	MemberClusterAnnounceFinalizer = "memberclusterannounce.finalizer.antrea.io"

	TimerInterval     = 10 * time.Second
	ConnectionTimeout = 3 * TimerInterval
)

type memberData struct {
	lastUpdateTime time.Time
	status         *mcv1alpha2.ClusterStatus
}

// MemberClusterAnnounceReconciler reconciles a MemberClusterAnnounce object
type MemberClusterAnnounceReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	mapLock         sync.RWMutex
	memberStatusMap map[common.ClusterID]*memberData
}

type MemberClusterStatusManager interface {
	GetMemberClusterStatuses() []mcv1alpha2.ClusterStatus
}

func NewMemberClusterAnnounceReconciler(client client.Client, scheme *runtime.Scheme) *MemberClusterAnnounceReconciler {
	_ = "STUB: not implemented"
	return nil
}

//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=memberclusterannounces,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=memberclusterannounces/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=multicluster.crd.antrea.io,resources=memberclusterannounces/finalizers,verbs=update

// Reconcile implements cluster status management on the leader cluster
func (r *MemberClusterAnnounceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// If MemberClusterAnnounce is deleted, no further processing is needed, as cleanup
// must have been done when the Finalizer was removed.

// SetupWithManager sets up the controller with the Manager.
func (r *MemberClusterAnnounceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"

	// Running background task here.
	return nil
}

func (r *MemberClusterAnnounceReconciler) processMCSStatus() { _ = "STUB: not implemented"; return }

// Check if the member has connected at least once in the last 3 intervals.

// Member has updated MemberClusterStatus at least once in the last 3 intervals.
// If last status is not connected, then update the status.

// Member has not updated MemberClusterStatus in the last 3 intervals, assume it is disconnected

func (r *MemberClusterAnnounceReconciler) addOrUpdateMemberStatus(memberID common.ClusterID) {
	_ = "STUB: not implemented"
	return
}

// Reset lastUpdateTime for this member.

func (r *MemberClusterAnnounceReconciler) removeMemberStatus(memberID common.ClusterID) {
	_ = "STUB: not implemented"
	return
}

/******************************* MemberClusterStatusManager methods *******************************/

func (r *MemberClusterAnnounceReconciler) GetMemberClusterStatuses() []mcv1alpha2.ClusterStatus {
	_ = "STUB: not implemented"
	return nil
}
