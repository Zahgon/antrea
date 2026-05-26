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

package installation

import (
	"context"

	corev1 "k8s.io/api/core/v1"
)

type PodToServiceConnectivityTest struct {
	getService func(testContext *testContext) (*corev1.Service, error)
}

func init() {
	RegisterTest("pod-to-service-intranode-connectivity", &PodToServiceConnectivityTest{
		getService: func(testContext *testContext) (*corev1.Service, error) {
			return testContext.echoSameNodeService, nil
		},
	})
	RegisterTest("pod-to-service-internode-connectivity", &PodToServiceConnectivityTest{
		getService: func(testContext *testContext) (*corev1.Service, error) {
			if testContext.echoOtherNodeService == nil {
				return nil, newNotRunnableError("Inter-Node test requires multiple Nodes")
			}
			return testContext.echoOtherNodeService, nil
		},
	})
}

func (t *PodToServiceConnectivityTest) Run(ctx context.Context, testContext *testContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Service is realized asynchronously, retry a few times.
