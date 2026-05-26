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

package portcache

import (
	"io"
	"sync"

	"antrea.io/antrea/v2/pkg/agent/nodeportlocal/rules"

	"k8s.io/client-go/tools/cache"
)

const (
	NodePortIndex    = "nodePortIndex"
	PodEndpointIndex = "podEndpointIndex"
	PodKeyIndex      = "podKeyIndex"
)

type ProtocolSocketData struct {
	Protocol string
	socket   io.Closer
}

type NodePortData struct {
	// PodKey is the namespaced name of the Pod.
	PodKey   string
	NodePort int
	PodPort  int
	PodIP    string
	Protocol ProtocolSocketData
	// defunct is used to indicate that a rule has been partially deleted: it is no longer
	// usable and deletion needs to be re-attempted.
	defunct bool
}

func (d *NodePortData) Defunct() bool { _ = "STUB: not implemented"; return false }

type LocalPortOpener interface {
	OpenLocalPort(port int, protocol string, isIPv6 bool) (io.Closer, error)
}

type localPortOpener struct{}

type PortTable struct {
	PortTableCache  cache.Indexer
	StartPort       int
	EndPort         int
	PortSearchStart int
	PodPortRules    rules.PodPortRules
	LocalPortOpener LocalPortOpener
	IsIPv6          bool
	tableLock       sync.RWMutex
}

func GetPortTableKey(obj interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (pt *PortTable) addPortTableCache(npData *NodePortData) error {
	_ = "STUB: not implemented"
	return nil
}

func (pt *PortTable) deletePortTableCache(npData *NodePortData) error {
	_ = "STUB: not implemented"
	return nil
}

func (pt *PortTable) getPortTableCacheFromNodePortIndex(index string) (*NodePortData, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pt *PortTable) getPortTableCacheFromPodEndpointIndex(index string) (*NodePortData, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pt *PortTable) getPortTableCacheFromPodKeyIndex(index string) ([]*NodePortData, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pt *PortTable) releaseDataFromPortTableCache() error { _ = "STUB: not implemented"; return nil }

func NodePortIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PodEndpointIndexFunc(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PodKeyIndexFunc(obj interface{}) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func NewPortTable(start, end int, isIPv6 bool) (*PortTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pt *PortTable) CleanupAllEntries() { _ = "STUB: not implemented"; return }

func (pt *PortTable) GetDataForPod(podKey string) []*NodePortData {
	_ = "STUB: not implemented"
	return nil
}

func (pt *PortTable) getDataForPod(podKey string) []*NodePortData {
	_ = "STUB: not implemented"
	return nil
}

func (pt *PortTable) GetEntry(podKey string, port int, protocol string) *NodePortData {
	_ = "STUB: not implemented"
	return nil
}

// Return pointer to copy of data from the PodEndpointTable.

// podKeyPortProtoFormat formats the podKey, port and protocol to string key:port:protocol.
func podKeyPortProtoFormat(podKey string, port int, protocol string) string {
	_ = "STUB: not implemented"
	return ""
}

func (pt *PortTable) getEntryByPodKeyPortProto(podKey string, port int, protocol string) *NodePortData {
	_ = "STUB: not implemented"
	return nil
}

func (pt *PortTable) RuleExists(podKey string, podPort int, protocol string) bool {
	_ = "STUB: not implemented"
	return false
}

// nodePortProtoFormat formats the nodeport, protocol to string port:protocol.
func NodePortProtoFormat(nodeport int, protocol string) string {
	_ = "STUB: not implemented"
	return ""
}

// openLocalPort binds to the provided port.
// This is inspired by the openLocalPort function in kube-proxy:
// https://github.com/kubernetes/kubernetes/blob/86f8c3ee91b6faec437f97e3991107747d7fc5e8/pkg/proxy/iptables/proxier.go#L1664
func (lpo *localPortOpener) OpenLocalPort(port int, protocol string, isIPv6 bool) (io.Closer, error) {
	_ = "STUB: not implemented"
	return *new(io.Closer), nil
}
