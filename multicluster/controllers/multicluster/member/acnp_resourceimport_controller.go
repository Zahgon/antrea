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

	ctrl "sigs.k8s.io/controller-runtime"

	multiclusterv1alpha1 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha1"
	"antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

const acnpImportFailed string = "ACNPImportFailed"

var (
	resourceImportAPIVersion     = "multicluster.crd.antrea.io/v1alpha1"
	resourceImportKind           = "ResourceImport"
	acnpEventReportingController = "resourceimport-controller"
	// TODO(yang): add run-time pod suffix
	acnpEventReportingInstance = "antrea-mc-controller"
)

func (r *ResourceImportReconciler) handleResImpUpdateForClusterNetworkPolicy(ctx context.Context, resImp *multiclusterv1alpha1.ResourceImport) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

// If the ACNP Tier exists in the importing member cluster, then the policy is realizable.
// Create or update the ACNP if necessary.

// The ACNP Tier does not exist, and the policy cannot be realized in this particular importing member cluster.
// If there is an ACNP previously created via import (which has a valid Tier by then), it should be cleaned up.

func (r *ResourceImportReconciler) handleResImpDeleteForClusterNetworkPolicy(ctx context.Context, resImp *multiclusterv1alpha1.ResourceImport) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}

func getMCAntreaClusterPolicy(resImp *multiclusterv1alpha1.ResourceImport) *v1beta1.ClusterNetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResourceImportReconciler) reportStatusEvent(errMsg string, ctx context.Context, resImp *multiclusterv1alpha1.ResourceImport) error {
	_ = "STUB: not implemented"
	return nil
}
