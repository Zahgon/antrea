// Copyright 2021 Antrea Authors
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
	"time"

	corev1 "k8s.io/api/core/v1"

	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/filter"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/priorityqueue"
	"antrea.io/antrea/v2/pkg/agent/proxy"
	"antrea.io/antrea/v2/pkg/querier"
	"antrea.io/antrea/v2/pkg/util/objectstore"
)

var serviceProtocolMap = map[uint8]corev1.Protocol{
	6:   corev1.ProtocolTCP,
	17:  corev1.ProtocolUDP,
	132: corev1.ProtocolSCTP,
}

type ConntrackConnectionStore struct {
	networkPolicyReadyTime time.Time
	protocolFilter         filter.ProtocolFilter
	connectionStore
}

func NewConntrackConnectionStore(
	npQuerier querier.AgentNetworkPolicyInfoQuerier,
	podStore objectstore.PodStore,
	proxier proxy.ProxyQuerier,
	cfg ConnectionStoreConfig,
) *ConntrackConnectionStore {
	_ = "STUB: not implemented"
	return nil
}

func (cs *ConntrackConnectionStore) AddOrUpdateConns(conns []*connection.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset IsPresent flag for all connections in connection map before updating
// the dumped flows information in connection map. If the connection does not
// exist in conntrack table and has been exported, then we will delete it from
// connection map. In addition, if the connection was not exported for a specific
// time period, then we consider it to be stale and delete it.

// Delete the connection if it is ready to delete or it was not exported
// in the time period as specified by the stale connection timeout.

// In case ReadyToDelete is true, item should already have been removed from pq

// Hold the lock until we verify whether the connection exist in conntrack table,
// and finish updating the connection store.

// Update only the Connection store. IPFIX records are generated based on Connection store.

// AddOrUpdateConn updates the connection if it is already present, i.e., update timestamp, counters etc.,
// or adds a new connection with the resolved K8s metadata.
func (cs *ConntrackConnectionStore) AddOrUpdateConn(conn *connection.Connection) {
	_ = "STUB: not implemented"
	return
}

// Update the necessary fields that are used in generating flow records.
// Can same 5-tuple flow get deleted and added to conntrack table? If so use ID.

// If the connKey:pqItem pair does not exist in the map, it shows the
// conn was inactive, and was removed from PQ and map. Since it becomes
// active again now, we create a new pqItem and add it to PQ and map.

// We don't add connections to connection map or expirePriorityQueue if we can't find the pod
// information for both srcPod and dstPod

// This should only happen if we failed to set net.netfilter.nf_conntrack_timestamp

// Add new antrea connection to connection store and PQ.

func (cs *ConntrackConnectionStore) GetExpiredConns(expiredConns []connection.Connection, currTime time.Time, maxSize int) ([]connection.Connection, time.Duration) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration)
}

// If a conntrack connection is in dying state or connection is not
// in the conntrack table, we set the ReadyToDelete flag to true to
// do the deletion later.

// No packets have been received during the idle timeout interval,
// the connection is therefore considered inactive.

// deleteConnWithoutLock deletes the connection from the connection map given
// the connection key without grabbing the lock. Caller is expected to grab lock.
func (cs *ConntrackConnectionStore) deleteConnWithoutLock(connKey connection.ConnectionKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *ConntrackConnectionStore) DeleteAllConnections() int { _ = "STUB: not implemented"; return 0 }

func (cs *ConntrackConnectionStore) GetPriorityQueue() *priorityqueue.ExpirePriorityQueue {
	_ = "STUB: not implemented"
	return nil
}
