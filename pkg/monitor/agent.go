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

package monitor

import (
	agentquerier "antrea.io/antrea/v2/pkg/agent/querier"
	"antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	clientset "antrea.io/antrea/v2/pkg/client/clientset/versioned"
)

type agentMonitor struct {
	client  clientset.Interface
	querier agentquerier.AgentQuerier
	// apiCertData is not provided by the querier to avoid a circular dependency between
	// apiServer and querier.
	apiCertData []byte
	// agentCRD is the desired state of agent monitoring CRD which agentMonitor expects.
	agentCRD *v1beta1.AntreaAgentInfo
}

// NewAgentMonitor creates a new agent monitor.
func NewAgentMonitor(client clientset.Interface, querier agentquerier.AgentQuerier, apiCertData []byte) *agentMonitor {
	_ = "STUB: not implemented"
	return nil
}

// Run creates AntreaAgentInfo CRD first after controller is running.
// Then updates AntreaAgentInfo CRD every 60 seconds.
func (monitor *agentMonitor) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Sync agent monitoring CRD every minute util stopCh is closed.

func (monitor *agentMonitor) syncAgentCRD() { _ = "STUB: not implemented"; return }

// getAgentCRD is used to check the existence of agent monitoring CRD.
// So when the pod restarts, it will update this monitoring CRD instead of creating a new one.
func (monitor *agentMonitor) getAgentCRD() (*v1beta1.AntreaAgentInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// updateAgentCRD updates the monitoring CRD.
func (monitor *agentMonitor) updateAgentCRD(partial bool) (*v1beta1.AntreaAgentInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
