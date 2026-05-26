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

	"antrea.io/antrea/v2/pkg/agent/flowexporter/connection"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/filter"
	"antrea.io/antrea/v2/pkg/agent/flowexporter/priorityqueue"
	"antrea.io/antrea/v2/pkg/agent/proxy"
	"antrea.io/antrea/v2/pkg/querier"
	"antrea.io/antrea/v2/pkg/util/objectstore"
)

type DenyConnectionStore struct {
	connectionStore
	protocolFilter filter.ProtocolFilter
}

func NewDenyConnectionStore(
	npQuerier querier.AgentNetworkPolicyInfoQuerier,
	podStore objectstore.PodStore,
	proxier proxy.ProxyQuerier,
	cfg ConnectionStoreConfig,
) *DenyConnectionStore {
	_ = "STUB: not implemented"
	return nil
}

func (ds *DenyConnectionStore) RunPeriodicDeletion(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

// In case ReadyToDelete is true, item should already have been removed from pq

// AddOrUpdateConn updates the connection if it is already present, i.e., update timestamp, counters etc.,
// or adds a new connection with the resolved K8s metadata.
func (ds *DenyConnectionStore) AddOrUpdateConn(conn *connection.Connection) {
	_ = "STUB: not implemented"
	return
}

// We don't add connections to connection map or expirePriorityQueue if we can't find the pod
// information for both srcPod and dstPod

// For intra-Node flows which are denied by an ingress policy rule, we can retrieve
// egress policy information from the CT labels.

func (ds *DenyConnectionStore) GetExpiredConns(expiredConns []connection.Connection, currTime time.Time, maxSize int) ([]connection.Connection, time.Duration) {
	_ = "STUB: not implemented"
	return nil, *new(time.Duration)
}

// If a deny connection item is idle time out, we set the ReadyToDelete
// flag to true to do the deletion later.

// If a deny connection doesn't have increase in packet count,
// we consider the connection to be inactive.

// deleteConnWithoutLock deletes the connection from the connection map given
// the connection key without grabbing the lock. Caller is expected to grab lock.
func (ds *DenyConnectionStore) deleteConnWithoutLock(connKey connection.ConnectionKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (ds *DenyConnectionStore) GetPriorityQueue() *priorityqueue.ExpirePriorityQueue {
	_ = "STUB: not implemented"
	return nil
}
