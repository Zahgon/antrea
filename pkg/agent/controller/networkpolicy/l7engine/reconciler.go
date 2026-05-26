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

package l7engine

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/spf13/afero"
	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/openflow"
	v1beta "antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	utilsync "antrea.io/antrea/v2/pkg/util/sync"
)

const (
	defaultSuricataConfigPath = "/etc/suricata/suricata.yaml"
	antreaSuricataConfigPath  = "/etc/suricata/antrea.yaml"
	antreaSuricataLogSubdir   = "networkpolicy/l7engine"

	tenantConfigsDir = "/etc/suricata"
	tenantRulesDir   = "/etc/suricata/rules"

	suricataCommandSocket = "/var/run/suricata/suricata-command.socket"

	protocolHTTP = "http"
	protocolTLS  = "tls"

	scCmdOK = "OK"
)

type scCmdRet struct {
	Message string `json:"message"`
	Return  string `json:"return"`
}

var (
	// Declared as a variable for testing.
	defaultFS = afero.NewOsFs()

	// Create the config file /etc/suricata/antrea.yaml for Antrea which will be included in the default Suricata config file
	// /etc/suricata/suricata.yaml. Two event logs in the config serve alert gilogging and http event logging purposes respectively.
	suricataAntreaConfigData = fmt.Sprintf(`%%YAML 1.1
---
outputs:
  - eve-log:
      enabled: yes
      filetype: regular
      filename: eve-%%Y-%%m-%%d.json
      rotate-interval: day
      pcap-file: false
      community-id: false
      community-id-seed: 0
      xff:
        enabled: no
      types:
        - alert:
            packet: yes
        - http:
            extended: yes
        - tls:
            extended: yes
af-packet:
  - interface: %[1]s
    threads: auto
    cluster-id: 80
    cluster-type: cluster_flow
    defrag: no
    use-mmap: yes
    tpacket-v2: yes
    checksum-checks: no
    copy-mode: ips
    copy-iface: %[2]s
  - interface:  %[2]s
    threads: auto
    cluster-id: 81
    cluster-type: cluster_flow
    defrag: no
    use-mmap: yes
    tpacket-v2: yes
    checksum-checks: no
    copy-mode: ips
    copy-iface: %[1]s
multi-detect:
  enabled: yes
  selector: vlan
`, config.L7RedirectTargetPortName, config.L7RedirectReturnPortName)
)

type threadSafeSet[T comparable] struct {
	sync.RWMutex
	cached sets.Set[T]
}

func (g *threadSafeSet[T]) has(key T) bool { _ = "STUB: not implemented"; return false }

func (g *threadSafeSet[T]) insert(key T) { _ = "STUB: not implemented"; return }

func (g *threadSafeSet[T]) delete(key T) { _ = "STUB: not implemented"; return }

type Reconciler struct {
	// Declared as member variables for testing.
	startSuricataFn func()
	suricataScFn    func(scCmd string) (*scCmdRet, error)

	suricataTenantCache        *threadSafeSet[uint32]
	suricataTenantHandlerCache *threadSafeSet[uint32]

	ofClient openflow.Client

	startSuricataOnce     utilsync.OnceWithNoError
	initializeL7FlowsOnce utilsync.OnceWithNoError
}

func NewReconciler(ofClient openflow.Client) *Reconciler { _ = "STUB: not implemented"; return nil }

func generateTenantRulesData(policyName string, protoKeywords map[string]sets.Set[string]) *bytes.Buffer {
	_ = "STUB: not implemented"
	return nil
}

// Generate default reject rule.

// Generate rules.

// It is a convention that the sid is provided as the last keyword (or second-to-last if there is a rev)
// of a rule.

func generateTenantRulesPath(vlanID uint32) string { _ = "STUB: not implemented"; return "" }

func generateTenantConfigPath(vlanID uint32) string { _ = "STUB: not implemented"; return "" }

func writeConfigFile(path string, data *bytes.Buffer) error { _ = "STUB: not implemented"; return nil }

// By default, Suricata performs pattern-matching for provided content. To support exact match, prefix match, and suffix
// match, we use wildcards to indicate whether an exact match is expected.
// - A string starting with * means suffix match. For example, "*.foo.com" matches "www.foo.com".
// - A string ending with * means prefix match. For example, "/public/*" matches "/public/index.html".
// - A string starting with and ending with * means pattern-matching. For example, "*/v2/*" matches "/api/v2/pods".
// - A string having no * means exact match. For example, "/index.html" can only match "/index.html".
func convertContent(content string) string { _ = "STUB: not implemented"; return "" }

func convertProtocolHTTP(http *v1beta.HTTPProtocol) string { _ = "STUB: not implemented"; return "" }

func convertProtocolTLS(tls *v1beta.TLSProtocol) string { _ = "STUB: not implemented"; return "" }

func (r *Reconciler) StartSuricataOnce() error { _ = "STUB: not implemented"; return nil }

func (r *Reconciler) initializeL7Flows() error { _ = "STUB: not implemented"; return nil }

func (r *Reconciler) AddRule(ruleID, policyName string, vlanID uint32, l7Protocols []v1beta.L7Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate the keyword part used in Suricata rules.

// Write the Suricata rules to file.

// Add a Suricata tenant.

func (r *Reconciler) DeleteRule(ruleID string, vlanID uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete the Suricata tenant.

// Delete the Suricata rules file.

func (r *Reconciler) addBindingSuricataTenant(vlanID uint32, rulesPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// If the tenant config file exists, it means that this tenant has been added, just reload the tenant to load the
// updated rules.

// If the tenant config file doesn't exist, create a config file for the tenant.

// Delete the config file regardless if it is created.

// Register the tenant with the config file. Note that, to be simple, use the VLAN id as the tenant ID.

// Register the tenant handler by mapping the tenant to the allocated VLAN ID.

func (r *Reconciler) deleteBindingSuricataTenant(vlanID uint32) error {
	_ = "STUB: not implemented"
	// Unregister the tenant handler.
	return nil
}

// Unregister the tenant.

// Delete the tenant config file.

func (r *Reconciler) reloadSuricataTenant(tenantID uint32, tenantConfigPath string) (*scCmdRet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reconciler) registerSuricataTenant(tenantID uint32, tenantConfigPath string) (*scCmdRet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reconciler) unregisterSuricataTenant(tenantID uint32) (*scCmdRet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reconciler) registerSuricataTenantHandler(tenantID, vlanID uint32) (*scCmdRet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reconciler) unregisterSuricataTenantHandler(tenantID, vlanID uint32) (*scCmdRet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reconciler) startSuricata() error { _ = "STUB: not implemented"; return nil }

// Open the default Suricata config file /etc/suricata/suricata.yaml.

// Include the config file /etc/suricata/antrea.yaml for Antrea in the default Suricata config file /etc/suricata/suricata.yaml.

// Wait Suricata command socket file to be ready.

func startSuricata() {
	_ = "STUB: not implemented"
	// Ensure that rules directory exists.
	return
}

// Create log directory for Suricata.

// Start Suricata with default Suricata config file /etc/suricata/suricata.yaml.

func suricataSc(scCmd string) (*scCmdRet, error) { _ = "STUB: not implemented"; return nil, nil }
