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

package supportbundlecollection

import (
	"sync"

	"github.com/spf13/afero"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/utils/exec"

	"antrea.io/antrea/v2/pkg/agent/client"
	agentquerier "antrea.io/antrea/v2/pkg/agent/querier"
	"antrea.io/antrea/v2/pkg/apis/controlplane"
	cpv1b2 "antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	"antrea.io/antrea/v2/pkg/ovs/ovsctl"
	"antrea.io/antrea/v2/pkg/querier"
	"antrea.io/antrea/v2/pkg/support"
	"antrea.io/antrea/v2/pkg/util/sftp"
)

type ProtocolType string

const (
	sftpProtocol ProtocolType = "sftp"

	controllerName = "SupportBundleCollectionController"
)

var (
	emptyWatch      = watch.NewEmptyWatch()
	defaultFS       = afero.NewOsFs()
	defaultExecutor = exec.New()
	// Declared as variable for testing.
	newAgentDumper = support.NewAgentDumper
)

type SupportBundleController struct {
	nodeName                     string
	supportBundleNodeType        controlplane.SupportBundleCollectionNodeType
	namespace                    string
	antreaClientGetter           client.AntreaClientProvider
	queue                        workqueue.TypedInterface[string]
	supportBundleCollection      *cpv1b2.SupportBundleCollection
	supportBundleCollectionMutex sync.RWMutex
	ovsCtlClient                 ovsctl.OVSCtlClient
	aq                           agentquerier.AgentQuerier
	npq                          querier.AgentNetworkPolicyInfoQuerier
	v4Enabled                    bool
	v6Enabled                    bool
	sftpUploader                 sftp.Uploader
}

func NewSupportBundleController(nodeName string,
	supportBundleNodeType controlplane.SupportBundleCollectionNodeType,
	namespace string,
	antreaClientGetter client.AntreaClientProvider,
	ovsCtlClient ovsctl.OVSCtlClient,
	aq agentquerier.AgentQuerier,
	npq querier.AgentNetworkPolicyInfoQuerier,
	v4Enabled,
	v6Enabled bool) *SupportBundleController {
	_ = "STUB: not implemented"
	return nil
}

func (c *SupportBundleController) watchSupportBundleCollections() {
	_ = "STUB: not implemented"
	return
}

// Watch method doesn't return error but "emptyWatch" in case of some partial data errors,
// e.g. timeout error. Make sure that watcher is not empty and log warning otherwise.

func (c *SupportBundleController) addSupportBundleCollection(supportBundle *cpv1b2.SupportBundleCollection) {
	_ = "STUB: not implemented"
	return
}

func (c *SupportBundleController) deleteSupportBundleCollection(supportBundle *cpv1b2.SupportBundleCollection) {
	_ = "STUB: not implemented"
	return
}

func (c *SupportBundleController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *SupportBundleController) worker() { _ = "STUB: not implemented"; return }

func (c *SupportBundleController) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// Skip retrying as the time may not meet the requirements for SupportBundle.

func (c *SupportBundleController) syncSupportBundleCollection(key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SupportBundleController) generateSupportBundle(supportBundle *cpv1b2.SupportBundleCollection) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SupportBundleController) uploadSupportBundle(supportBundle *cpv1b2.SupportBundleCollection, outputFile afero.File) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SupportBundleController) getUploaderByProtocol(protocol ProtocolType) (sftp.Uploader, error) {
	_ = "STUB: not implemented"
	return *new(sftp.Uploader), nil
}

func (c *SupportBundleController) updateSupportBundleCollectionStatus(key string, complete bool, genErr error) error {
	_ = "STUB: not implemented"
	return nil
}
