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

package cluster

import (
	"context"
	"time"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"antrea.io/antrea/v2/pkg/antctl/raw/check"
)

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

// Disables positional arguments

const (
	testNamespacePrefix = "antrea-test"
	deploymentName      = "cluster-checker"
	podReadyTimeout     = 1 * time.Minute
)

type options struct {
	// Container image for the cluster checker.
	testImage string
}

func newOptions() *options { _ = "STUB: not implemented"; return nil }

type uncertainError struct {
	reason string
}

func (e uncertainError) Error() string { _ = "STUB: not implemented"; return "" }

func newUncertainError(reason string, a ...interface{}) uncertainError {
	_ = "STUB: not implemented"
	return *new(uncertainError)
}

type Test interface {
	Run(ctx context.Context, testContext *testContext) error
}

var testsRegistry = make(map[string]Test)

func RegisterTest(name string, test Test) { _ = "STUB: not implemented"; return }

type testContext struct {
	check.Logger
	client      kubernetes.Interface
	config      *rest.Config
	clusterName string
	namespace   string
	testPod     *corev1.Pod
	// Container image for the cluster checker.
	testImage string
}

func Run(o *options) error { _ = "STUB: not implemented"; return nil }

func (t *testContext) setup(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func NewTestContext(client kubernetes.Interface, config *rest.Config, clusterName, testImage string) *testContext {
	_ = "STUB: not implemented"
	return nil
}

func (t *testContext) Header(format string, a ...interface{}) { _ = "STUB: not implemented"; return }
