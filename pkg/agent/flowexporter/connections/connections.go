// Copyright 2020 Antrea Authors
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

package connections

import (
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"

	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/priorityqueue"
	"antrea.io/antrea/v2/pkg/agent/proxy"
	"antrea.io/antrea/v2/pkg/agent/types"
	"antrea.io/antrea/v2/pkg/querier"
	"antrea.io/antrea/v2/pkg/util/objectstore"
)

const (
	periodicDeleteInterval = time.Minute
)

type ConnectionStoreConfig struct {
	ActiveFlowTimeout      time.Duration
	IdleFlowTimeout        time.Duration
	StaleConnectionTimeout time.Duration

	NetworkPolicyReadyTime time.Time
	AllowedProtocols       []string
}

type connectionStore struct {
	connections            map[connection.ConnectionKey]*connection.Connection
	networkPolicyQuerier   querier.AgentNetworkPolicyInfoQuerier
	podStore               objectstore.PodStore
	antreaProxier          proxy.ProxyQuerier
	expirePriorityQueue    *priorityqueue.ExpirePriorityQueue
	staleConnectionTimeout time.Duration
	mutex                  sync.Mutex
}

func NewConnectionStore(
	npQuerier querier.AgentNetworkPolicyInfoQuerier,
	podStore objectstore.PodStore,
	proxier proxy.ProxyQuerier,
	cfg ConnectionStoreConfig) connectionStore {
	_ = "STUB: not implemented"
	return *new(connectionStore)
}

// GetConnByKey gets the connection in connection map given the connection key.
func (cs *connectionStore) GetConnByKey(connKey connection.ConnectionKey) (*connection.Connection, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (cs *connectionStore) NumConnections() int { _ = "STUB: not implemented"; return 0 }

// ForAllConnectionsDo execute the callback for each connection in connection map.
func (cs *connectionStore) ForAllConnectionsDo(callback connection.ConnectionMapCallBack) error {
	_ = "STUB: not implemented"
	return nil
}

// ForAllConnectionsDoWithoutLock execute the callback for each connection in connection
// map, without grabbing the lock. Caller is expected to grab lock.
func (cs *connectionStore) ForAllConnectionsDoWithoutLock(callback connection.ConnectionMapCallBack) error {
	_ = "STUB: not implemented"
	return nil
}

// AddConnToMap adds the connection to connections map given connection key.
// This is used only for unit tests.
func (cs *connectionStore) AddConnToMap(connKey *connection.ConnectionKey, conn *connection.Connection) {
	_ = "STUB: not implemented"
	return
}

func (cs *connectionStore) fillPodInfo(conn *connection.Connection) {
	_ = "STUB: not implemented"
	return
}

// sourceIP/destinationIP are mapped only to local pods and not remote pods.

func (cs *connectionStore) fillServiceInfo(conn *connection.Connection, serviceStr string) {
	_ = "STUB: not implemented"
	// resolve destination Service information
	return
}

// LookupServiceProtocol returns the corresponding Service protocol string for a given protocol identifier
func lookupServiceProtocol(protoID uint8) (corev1.Protocol, error) {
	_ = "STUB: not implemented"
	return *new(corev1.Protocol), nil
}

func (cs *connectionStore) getPolicyRuleMetadata(conn *connection.Connection, ruleID uint32, labelsStart, labelsEnd int) (*types.PolicyRule, uint8, bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}

// deny connections have their ruleIDs set

// allow connections have their flowIDs set in the labels

func (cs *connectionStore) addIngressNetworkPolicyMetadata(conn *connection.Connection) {
	_ = "STUB: not implemented"
	return
}

func (cs *connectionStore) addEgressNetworkPolicyMetadata(conn *connection.Connection) {
	_ = "STUB: not implemented"
	return
}

func (cs *connectionStore) addNetworkPolicyMetadata(conn *connection.Connection) {
	_ = "STUB: not implemented"
	return
}

func (cs *connectionStore) AcquireConnStoreLock() { _ = "STUB: not implemented"; return }

func (cs *connectionStore) ReleaseConnStoreLock() {
	_ = "STUB: not implemented"

	// UpdateConnAndQueue deletes the inactive connection from keyToItem map,
	// without adding it back to the PQ. In this way, we can avoid to reset the
	// item's expire time every time we encounter it in the PQ. The method also
	// updates active connection's stats fields and adds it back to the PQ. Layer 7
	// fields should be set to default to prevent from re-exporting same values.
	return
}

func (cs *connectionStore) UpdateConnAndQueue(pqItem *priorityqueue.ItemToExpire, currTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// For active connections, we update their "prev" stats fields,
// reset active expire time and push back into the PQ.
