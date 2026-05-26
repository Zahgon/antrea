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

package multicluster

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"antrea.io/antrea/v2/multicluster/pkg/client/informers/externalversions/multicluster/v1alpha1"
	mclisters "antrea.io/antrea/v2/multicluster/pkg/client/listers/multicluster/v1alpha1"
	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/openflow"
)

const (
	// The number of workers processing a Pod change
	podWorkerNum = 5

	dummyKey               = "key"
	podIndexKey            = "podIP"
	podRouteControllerName = "MCPodRouteController"
)

// MCPodRouteController generates L3 forwarding flows to forward cross-cluster
// traffic from MC Gateway to Pods on other Nodes inside a member cluster. It is
// required when networkPolicyOnly, noEncap or hybrid mode are configured, to forward
// the traffic through tunnels between Gateway and other Nodes, as otherwise the
// traffic will not go through tunnels in those modes.
type MCPodRouteController struct {
	k8sClient   kubernetes.Interface
	ofClient    openflow.Client
	nodeConfig  *config.NodeConfig
	podQueue    workqueue.TypedRateLimitingInterface[string]
	gwQueue     workqueue.TypedRateLimitingInterface[string]
	podInformer cache.SharedIndexInformer
	podLister   corelisters.PodLister
	gwInformer  cache.SharedIndexInformer
	gwLister    mclisters.GatewayLister
	// podWorkersStarted is a boolean which tracks if the Pod flow controller has been started.
	podWorkersStarted      bool
	podWorkersStartedMutex sync.RWMutex
	podWorkerStopCh        chan struct{}
}

func NewMCPodRouteController(
	k8sClient kubernetes.Interface,
	gwInformer v1alpha1.GatewayInformer,
	client openflow.Client,
	nodeConfig *config.NodeConfig,
) *MCPodRouteController {
	_ = "STUB: not implemented"
	return nil
}

// Gateway UPDATE event doesn't impact Pod flows, so ignore it.

func podIPIndexFunc(obj interface{}) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *MCPodRouteController) createPodInformer() { _ = "STUB: not implemented"; return }

func (c *MCPodRouteController) enqueueGateway(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *MCPodRouteController) createPod(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *MCPodRouteController) updatePod(old, cur interface{}) { _ = "STUB: not implemented"; return }

func (c *MCPodRouteController) deletePod(obj interface{}) { _ = "STUB: not implemented"; return }

func isValidPod(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func (c *MCPodRouteController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Run a single routine to handle Gateway events.

func (c *MCPodRouteController) gatewayWorker() { _ = "STUB: not implemented"; return }

func (c *MCPodRouteController) processGatewayNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *MCPodRouteController) syncGateway() error { _ = "STUB: not implemented"; return nil }

// Stop Pod flow controller and clean up all installed Multi-cluster Pod flows,
// if the Node was a Gateway before.

// Do nothing when the Pod flow controller is already started since
// Pod flow controller will be responsible for handling Pod events to install flows.

func (c *MCPodRouteController) podWorker() { _ = "STUB: not implemented"; return }

func (c *MCPodRouteController) processPodNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *MCPodRouteController) syncPod(podIP string) error { _ = "STUB: not implemented"; return nil }

func (c *MCPodRouteController) getLatestPod(pods []interface{}) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}
