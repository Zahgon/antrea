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

package leader

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mcv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
)

const (
	indexKey                       = "spec.clusterID"
	memberClusterAnnounceStaleTime = 24 * time.Hour
)

var getResourceExportsByClusterIDFunc = getResourceExportsByClusterID

// StaleResCleanupController will run periodically (memberClusterAnnounceStaleTime / 2 = 12 Hours)
// to clean up stale MemberClusterAnnounce resources in the leader cluster if the MemberClusterAnnounce
// timestamp annotation has not been updated for memberClusterAnnounceStaleTime (24 Hours).
// It will remove all ResourceExports belong to a member cluster when the corresponding MemberClusterAnnounce
// CR is deleted. It will also try to clean up all stale ResourceExports during start.
type StaleResCleanupController struct {
	client.Client
	Scheme *runtime.Scheme
}

func NewStaleResCleanupController(
	Client client.Client,
	Scheme *runtime.Scheme,
) *StaleResCleanupController {
	_ = "STUB: not implemented"
	return nil
}

// cleanUpExpiredMemberClusterAnnounces will delete any MemberClusterAnnounce if its
// last update timestamp is over 24 hours.
func (c *StaleResCleanupController) cleanUpExpiredMemberClusterAnnounces(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (c *StaleResCleanupController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *StaleResCleanupController) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// Ignore the event if it's not with non-zero DeletionTimestamp

// Clean up all corresponding ResourceExports when the member cluster's
// MemberClusterAnnounce is deleted.

// When cleanup is done, remove the Finalizer of this MemberClusterAnnounce.

// SetupWithManager sets up the controller with the Manager.
func (c *StaleResCleanupController) SetupWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	// Add an Indexer for ResourceExport, so it can be filtered by the ClusterID.
	return nil
}

func deleteResourceExports(ctx context.Context, mgrClient client.Client, resouceExports []mcv1alpha1.ResourceExport) bool {
	_ = "STUB: not implemented"
	return false
}

func getClusterIDFromName(name string) string { _ = "STUB: not implemented"; return "" }

func getResourceExportsByClusterID(c *StaleResCleanupController, ctx context.Context, clusterID string) ([]mcv1alpha1.ResourceExport, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
