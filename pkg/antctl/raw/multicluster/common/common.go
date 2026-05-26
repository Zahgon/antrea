// Copyright 2022 Antrea Authors
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

package common

import (
	"io"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mcv1alpha2 "antrea.io/antrea/v2/multicluster/apis/multicluster/v1alpha2"
)

const (
	ClusterSetJoinConfigAPIVersion = "multicluster.antrea.io/v1alpha1"
	ClusterSetJoinConfigKind       = "ClusterSetJoinConfig"

	CreateByAntctlAnnotation = "multicluster.antrea.io/created-by-antctl"

	DefaultMemberNamespace = "kube-system"
	DefaultLeaderNamespace = "antrea-multicluster"
)

// "omitempty" fields (clusterID, namespace, tokenSecretName, tokenSecretFile)
// can be populated by the corresponding command line options if not set in the
// config file.
type ClusterSetJoinConfig struct {
	APIVersion      string `yaml:"apiVersion"`
	Kind            string `yaml:"kind"`
	ClusterSetID    string `yaml:"clusterSetID"`
	ClusterID       string `yaml:"clusterID,omitempty"`
	Namespace       string `yaml:"namespace,omitempty"`
	LeaderClusterID string `yaml:"leaderClusterID"`
	LeaderNamespace string `yaml:"leaderNamespace"`
	LeaderAPIServer string `yaml:"leaderAPIServer"`
	TokenSecretName string `yaml:"tokenSecretName,omitempty"`
	TokenSecretFile string `yaml:"tokenSecretFile,omitempty"`
}

func NewClient(cmd *cobra.Command) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

func CreateClusterSet(cmd *cobra.Command, k8sClient client.Client, namespace string, clusterset string,
	leaderServer string, secret string, memberClusterID string, leaderClusterID string, leaderClusterNamespace string, createdRes *[]map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteClusterSet(cmd *cobra.Command, k8sClient client.Client, namespace string, clusterSet string) {
	_ = "STUB: not implemented"
	return
}

func deleteSecrets(cmd *cobra.Command, k8sClient client.Client, namespace string) {
	_ = "STUB: not implemented"
	return
}

func deleteRoleBindings(cmd *cobra.Command, k8sClient client.Client, namespace string) {
	_ = "STUB: not implemented"
	return
}

func deleteServiceAccounts(cmd *cobra.Command, k8sClient client.Client, namespace string) {
	_ = "STUB: not implemented"
	return
}

// ConvertMemberTokenSecret generates a token Secret manifest for creating the
// input Secret in a member cluster.
func ConvertMemberTokenSecret(secret *corev1.Secret) *corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func CreateMemberToken(cmd *cobra.Command, k8sClient client.Client, name string, namespace string, createdRes *[]map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// It will take one or two seconds to wait for the Data.token to be created.

func DeleteMemberToken(cmd *cobra.Command, k8sClient client.Client, name string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func waitForSecretReady(client client.Client, secretName string, namespace string) error {
	_ = "STUB: not implemented"
	return nil
}

func newClusterSet(name, namespace, leaderServer, secret, memberClusterID, leaderClusterID, leaderNamespace string) *mcv1alpha2.ClusterSet {
	_ = "STUB: not implemented"
	return nil
}

func newRoleBinding(name string, saName string, namespace string) *rbacv1.RoleBinding {
	_ = "STUB: not implemented"
	return nil
}

func newSecret(name string, saName string, namespace string) *corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func newServiceAccount(name string, namespace string) *corev1.ServiceAccount {
	_ = "STUB: not implemented"
	return nil
}

func OutputMemberTokenSecret(tokenSecret *corev1.Secret, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func OutputJoinConfig(cmd *cobra.Command, writer io.Writer, clusterSetID, leaderClusterID, leaderNamespace string, tokenSecret *corev1.Secret) error {
	_ = "STUB: not implemented"
	return nil
}

// We comment out these ClusterSetJoinConfig fields in the generated
// join config file, so they can be populated by command line options of
// the "antctl mc join" command.
