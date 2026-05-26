// Copyright 2020 Antrea Authors
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

package env

const (
	NodeNameEnvKey        = "NODE_NAME"
	podNameEnvKey         = "POD_NAME"
	PodNamespaceEnvKey    = "POD_NAMESPACE"
	svcAcctNameEnvKey     = "SERVICEACCOUNT_NAME"
	antreaConfigMapEnvKey = "ANTREA_CONFIG_MAP_NAME"

	antreaCloudEKSEnvKey = "ANTREA_CLOUD_EKS"

	defaultAntreaNamespace = "kube-system"

	// #nosec G101 -- not credentials
	allowNoEncapWithoutAntreaProxyEnvKey = "ALLOW_NO_ENCAP_WITHOUT_ANTREA_PROXY"
)

// GetNodeName returns the node's name used in Kubernetes, based on the priority:
// - Environment variable NODE_NAME, which should be set by Downward API
// - OS's hostname
func GetNodeName() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetPodName returns name of the Pod where the code executes.
func GetPodName() string { _ = "STUB: not implemented"; return "" }

// GetAntreaConfigMapName returns the configMap name of Antrea config.
func GetAntreaConfigMapName() string { _ = "STUB: not implemented"; return "" }

// GetPodNamespace returns Namespace of the Pod where the code executes.
func GetPodNamespace() string { _ = "STUB: not implemented"; return "" }

// GetAntreaControllerServiceAccountName returns the ServiceAccount name associated with antrea-controller.
func GetAntreaControllerServiceAccount() string { _ = "STUB: not implemented"; return "" }

// default value set for antrea-controller

func getBoolEnvVar(name string, defaultValue bool) bool { _ = "STUB: not implemented"; return false }

// IsCloudEKS returns true if Antrea is used to enforce NetworkPolicies in an EKS cluster.
func IsCloudEKS() bool { _ = "STUB: not implemented"; return false }

// GetAntreaNamespace tries to determine the Namespace in which Antrea is running by looking at the
// POD_NAMESPACE environment variable. If this environment variable is not set (e.g. because the
// Antrea component is not run as a Pod), "kube-system" is returned.
func GetAntreaNamespace() string { _ = "STUB: not implemented"; return "" }

// GetAllowNoEncapWithoutAntreaProxy returns whether AntreaProxy can be disabled for traffic
// modes which support noEncap.
func GetAllowNoEncapWithoutAntreaProxy() bool { _ = "STUB: not implemented"; return false }
