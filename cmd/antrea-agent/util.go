// Copyright 2021 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net"

	clientset "k8s.io/client-go/kubernetes"

	"antrea.io/antrea/v2/pkg/agent/util"
	k8sutil "antrea.io/antrea/v2/pkg/util/k8s"
)

var (
	// Declared variables which are meant to be overridden for testing.
	getAllNodeAddresses = util.GetAllNodeAddresses

	getPodCIDRsFromKubeProxy = k8sutil.GetPodCIDRsFromKubeProxy
	getPodCIDRsFromKubeadm   = k8sutil.GetPodCIDRsFromKubeadm
)

func getAvailableNodePortAddresses(nodePortAddressesFromConfig []string, excludeDevices []string, excludeDevicePrefixes []string) ([]net.IP, []net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get all IP addresses of Node

// If option `NodePortAddresses` is not set, then all Node IP addresses will be used as NodePort IP address.

// parsePortRange parses a port range ("<start>-<end>") and checks that it is valid.
func parsePortRange(portRangeStr string) (start, end int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// getPodCIDRs gets the cluster-wide Pod CIDRs (IPv4 and IPv6) by attempting the following sources in order:
// 1. Agent configuration if field `podCIDRs` is not empty.
// 2. kube-proxy ConfigMap.
// 3. kubeadm-config ConfigMap.
func getPodCIDRs(o *Options, k8sClient clientset.Interface) ([]*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseCIDRs(s string) ([]*net.IPNet, error) { _ = "STUB: not implemented"; return nil, nil }
