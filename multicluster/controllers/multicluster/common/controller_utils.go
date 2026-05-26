// Copyright 2022 Antrea Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"context"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mcv1alpha2 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha2"
)

var createDiscoveryClientFn = createDiscoveryClient

// discoverServiceCIDRByInvalidServiceCreation creates an invalid Service to get returned error, and analyzes
// the error message to get Service CIDR.
// TODO: add dual-stack support.
func discoverServiceCIDRByInvalidServiceCreation(ctx context.Context, k8sClient client.Client, namespace string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Creating invalid Service didn't fail as expected

// TODO: add dual-stack support.
func parseServiceCIDRFromError(msg string) (string, error) {
	_ = "STUB: not implemented"
	// Expected error message is like below:
	// `The Service "invalid-svc" is invalid: spec.clusterIPs: Invalid value: []string{"0.0.0.0"}:
	// failed to allocate IP 0.0.0.0: provided IP is not in the valid range. The range of valid IPs is 10.19.0.0/18`
	// The CIDR string should be parsed from the error message is:
	//   10.19.0.0/18
	return "", nil
}

func isK8sVersionGreaterThanOrEqualTo(discoveryClient discovery.DiscoveryInterface, expectedVersion string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getClusterServiceCIDR(ctx context.Context, apiReader client.Reader) (string, error) {
	_ = "STUB: not implemented"
	// A ServiceCIDR CR named 'kubernetes' will be created by default since v1.33.0.
	// We will retrieve it to get the Service CIDR.
	return "", nil
}

func createDiscoveryClient(config *rest.Config) (discovery.DiscoveryInterface, error) {
	_ = "STUB: not implemented"
	return *new(discovery.DiscoveryInterface), nil
}

func DiscoverClusterServiceCIDR(ctx context.Context, mgrConfig *rest.Config, apiReader client.Reader, client client.Client, namespace string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func isIPv4CIDR(cidr string) bool { _ = "STUB: not implemented"; return false }

func NewClusterInfoResourceExportName(clusterID string) string {
	_ = "STUB: not implemented"
	return ""
}

func getClusterIDFromClusterClaim(c client.Client, clusterSet *mcv1alpha2.ClusterSet) (ClusterID, error) {
	_ = "STUB: not implemented"
	return *new(ClusterID), nil
}

func GetClusterID(clusterCalimCRDAvailable bool, req ctrl.Request, client client.Client, clusterSet *mcv1alpha2.ClusterSet) (ClusterID, error) {
	_ = "STUB: not implemented"
	return *new(ClusterID), nil
}

// ClusterID is a required field, and the empty value case should only happen
// when Antrea Multi-cluster is upgraded from an old version prior to v1.13.
// Here we try to get the ClusterID from ClusterClaim before returning any error.
