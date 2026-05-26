// Copyright 2022 Antrea Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flowaggregator

import (
	"strings"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	flowaggregatorconfig "antrea.io/antrea/v2/pkg/config/flowaggregator"
)

// Command is the support bundle command implementation.
var Command *cobra.Command

type FlowAggregatorConfigMutator func(c *flowaggregatorconfig.FlowAggregatorConfig, value string) error

var mutators map[string]FlowAggregatorConfigMutator

var getClients = getk8sClient

var example = strings.Trim(`
  Enable ClickHouse
  $ antctl set flow-aggregator clickHouse.enable=true
  Update ClickHouse database
  $ antctl set flow-aggregator clickHouse.database=name
  Update ClickHouse databaseURL
  $ antctl set flow-aggregator clickHouse.databaseURL=http://xxxxx
  Update ClickHouse debug
  $ antctl set flow-aggregator clickHouse.debug=true
  Update ClickHouse compress
  $ antctl set flow-aggregator clickHouse.compress=true
  Update ClickHouse commitInterval
  $ antctl set flow-aggregator clickHouse.commitInterval=10s
  Update IPFIX Flow Collector address
  $ antctl set flow-aggregator flowCollector.address=<IP>:<port>[:<proto>]
  Enable IPFIX Flow Collector
  $ antctl set flow-aggregator flowCollector.enable=true
`, "\n")

func NewFlowAggregatorSetCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getk8sClient(cmd *cobra.Command) (kubernetes.Interface, error) {
	_ = "STUB: not implemented"
	return *new(kubernetes.Interface), nil
}

func updateRunE(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// marshal back the changed parameters to configmap

func setBoolOrFail(b *bool, value string) error { _ = "STUB: not implemented"; return nil }

func setStringOrFail(b *string, value string) error { _ = "STUB: not implemented"; return nil }

func setCommitIntervalOrFail(b *string, value string) error { _ = "STUB: not implemented"; return nil }

// GetFAConfigMap is used to get and return the flow-aggregator configmap
func GetFAConfigMap(k8sClient kubernetes.Interface, configMapName string) (*corev1.ConfigMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
