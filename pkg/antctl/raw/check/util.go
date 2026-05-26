// Copyright 2024 Antrea Authors.
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

package check

import (
	"context"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewClient() (client kubernetes.Interface, config *rest.Config, clusterName string, err error) {
	_ = "STUB: not implemented"
	return *new(kubernetes.Interface), nil, "", nil
}

// getDeploymentCondition returns the condition with the provided type.
func getDeploymentCondition(status appsv1.DeploymentStatus, condType appsv1.DeploymentConditionType) *appsv1.DeploymentCondition {
	_ = "STUB: not implemented"
	return nil
}

// DeploymentIsReady and DaemonSetIsReady are inspired by the implementation of "kubectl rollout status".

// DeploymentIsReady returns a message describing Deployment status, and a bool value indicating if the status is considered ready.
func DeploymentIsReady(deployment *appsv1.Deployment) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// DaemonSetIsReady returns a message describing DaemonSet status, and a bool value indicating if the status is considered ready.
func DaemonSetIsReady(daemonSet *appsv1.DaemonSet) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func NewDeployment(p DeploymentParameters) *appsv1.Deployment {
	_ = "STUB: not implemented"
	return nil
}

type DeploymentParameters struct {
	Name            string
	Role            string
	Image           string
	Replicas        int
	Port            int
	Command         []string
	Args            []string
	Affinity        *corev1.Affinity
	Tolerations     []corev1.Toleration
	Labels          map[string]string
	VolumeMounts    []corev1.VolumeMount
	Volumes         []corev1.Volume
	HostNetwork     bool
	NodeSelector    map[string]string
	SecurityContext *corev1.SecurityContext
}

func WaitForDeploymentsReady(
	ctx context.Context,
	interval, timeout time.Duration,
	immediate bool,
	client kubernetes.Interface,
	clusterName string,
	namespace string,
	deployments ...string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitForDaemonSetReady(
	ctx context.Context,
	interval, timeout time.Duration,
	immediate bool,
	client kubernetes.Interface,
	clusterName string,
	namespace string,
	name string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func GenerateRandomNamespace(baseName string) string { _ = "STUB: not implemented"; return "" }

func Teardown(ctx context.Context, logger Logger, client kubernetes.Interface, namespace string) {
	_ = "STUB: not implemented"
	return
}
