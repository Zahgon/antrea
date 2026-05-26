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

package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"

	npltypes "antrea.io/antrea/v2/pkg/agent/nodeportlocal/types"
)

const NPLEnabledAnnotationIndex = "nplEnabledAnnotation"

func toJSON(serialize interface{}) string { _ = "STUB: not implemented"; return "" }

func (c *NPLController) updatePodNPLAnnotation(ctx context.Context, pod *corev1.Pod, annotations []npltypes.NPLAnnotation) error {
	_ = "STUB: not implemented"
	return nil
}

func patchPod(ctx context.Context, value []npltypes.NPLAnnotation, pod *corev1.Pod, kubeClient clientset.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

// compareNPLAnnotationLists returns true if and only if the two lists contain the same set of
// annotations, irrespective of the order.
func compareNPLAnnotationLists(annotations1, annotations2 []npltypes.NPLAnnotation) bool {
	_ = "STUB: not implemented"
	return false
}

// Use IPFamily for tie-breaking to ensure consistent ordering
