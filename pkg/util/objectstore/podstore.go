// Copyright 2023 Antrea Authors
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

package objectstore

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
)

const podIPIndex = "podIP"

// PodStore interface provides Pod-specific operations
type PodStore interface {
	GetPodByIPAndTime(ip string, startTime time.Time) (*corev1.Pod, bool)
	Run(stopCh <-chan struct{})
	HasSynced() bool
}

// podStore embeds ObjectStore[*corev1.Pod] to provide Pod-specific methods
type podStore struct {
	*ObjectStore[*corev1.Pod]
}

// Validate that *podStore implements the PodStore interface
var _ PodStore = &podStore{}

func NewPodStore(podInformer cache.SharedIndexInformer) *podStore {
	_ = "STUB: not implemented"
	return nil
}

// GetPodByIPAndTime provides a Pod-specific method for getting Pods by IP and time
func (s *podStore) GetPodByIPAndTime(ip string, startTime time.Time) (*corev1.Pod, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func podIPIndexFunc(obj interface{}) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
