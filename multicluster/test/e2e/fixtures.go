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

package e2e

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func createDirectory(path string) error { _ = "STUB: not implemented"; return nil }

func (data *MCTestData) setupLogDirectoryForTest(testName string) error {
	_ = "STUB: not implemented"
	return nil
}

// remove directory if it already exists. This ensures that we start with an empty
// directory

func setupTest(tb testing.TB) (*MCTestData, error) { _ = "STUB: not implemented"; return nil, nil }

func teardownTest(tb testing.TB, data *MCTestData) { _ = "STUB: not implemented"; return }

func createPodWrapper(tb testing.TB, data *MCTestData, cluster string, namespace string, name string, nodeName string, image string, ctr string, command []string,
	args []string, env []corev1.EnvVar, ports []corev1.ContainerPort, hostNetwork bool, mutateFunc func(pod *corev1.Pod)) error {
	_ = "STUB: not implemented"
	return nil
}

func deletePodWrapper(tb testing.TB, data *MCTestData, clusterName string, namespace string, name string) {
	_ = "STUB: not implemented"
	return
}

func deletePodAndWaitWrapper(tb testing.TB, data *MCTestData, clusterName string, namespace string, name string) {
	_ = "STUB: not implemented"
	return
}

func deleteServiceWrapper(tb testing.TB, data *MCTestData, clusterName string, namespace string, name string) {
	_ = "STUB: not implemented"
	return
}
