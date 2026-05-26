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

package openflow

import (
	"sync"
	"time"

	"antrea.io/libOpenflow/openflow15"
	"antrea.io/ofnet/ofctrl"
	"golang.org/x/time/rate"
)

// ofTable implements openflow.Table.
type ofTable struct {
	// sync.RWMutex protects ofTable status from concurrent modification and reading.
	sync.RWMutex
	id         uint8
	name       string
	next       uint8
	missAction MissActionType
	flowCount  uint
	updateTime time.Time
	stageID    StageID
	pipelineID PipelineID

	*ofctrl.Table
}

func (t *ofTable) GetID() uint8 { _ = "STUB: not implemented"; return 0 }

func (t *ofTable) GetName() string { _ = "STUB: not implemented"; return "" }

func (t *ofTable) Status() TableStatus { _ = "STUB: not implemented"; return *new(TableStatus) }

func (t *ofTable) GetMissAction() MissActionType {
	_ = "STUB: not implemented"
	return *new(MissActionType)
}

func (t *ofTable) GetNext() uint8 { _ = "STUB: not implemented"; return 0 }

func (t *ofTable) SetNext(next uint8) { _ = "STUB: not implemented"; return }

func (t *ofTable) SetMissAction(action MissActionType) { _ = "STUB: not implemented"; return }

func (t *ofTable) GetStageID() StageID { _ = "STUB: not implemented"; return *new(StageID) }

func (t *ofTable) SetTable() { _ = "STUB: not implemented"; return }

func (t *ofTable) GetPipelineID() PipelineID { _ = "STUB: not implemented"; return *new(PipelineID) }

func (t *ofTable) UpdateStatus(flowCountDelta int) { _ = "STUB: not implemented"; return }

func (t *ofTable) ResetStatus() { _ = "STUB: not implemented"; return }

// BuildFlow returns FlowBuilder object to help construct Openflow entry.
func (t *ofTable) BuildFlow(priority uint16) FlowBuilder {
	_ = "STUB: not implemented"
	return *new(FlowBuilder)
}

// Set ofctl.Table to Flow, otherwise the flow can't find OFSwitch to install.

// DumpFlows dumps all existing Openflow entries from OFSwitch using cookie ID and table ID as filters.
func (t *ofTable) DumpFlows(cookieID, cookieMask uint64) (map[uint64]*FlowStates, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFlowStats(ofStats []*openflow15.FlowDesc) map[uint64]*FlowStates {
	_ = "STUB: not implemented"
	return nil
}

func NewOFTable(id uint8, name string, stageID StageID, pipelineID PipelineID, missAction MissActionType) Table {
	_ = "STUB: not implemented"
	return *new(Table)
}

// OFBridge implements openflow.Bridge.
type OFBridge struct {
	bridgeName string
	// Management address
	mgmtAddr string
	// sync.RWMutex protects tableCache from concurrent modification and iteration.
	sync.RWMutex
	// tableCache is used to cache ofTables.
	tableCache map[uint8]*ofTable

	ofSwitchMutex sync.RWMutex
	// ofSwitch is the target OFSwitch.
	ofSwitch *ofctrl.OFSwitch
	// controller helps maintain connections to remote OFSwitch.
	controller *ofctrl.Controller
	// retryInterval is the interval for retry connection.
	retryInterval time.Duration
	// maxRetrySec is the seconds waiting for connection to the OFSwitch.
	maxRetrySec int

	// channel to notify agent OFSwitch is connected.
	connCh chan struct{}
	// connected is an internal channel to notify if connected to the OFSwitch or not. It is used only in Connect method.
	connected chan bool
	// pktConsumers is a map from PacketIn category to the channel that is used to publish the PacketIn message.
	pktConsumers sync.Map

	// portStatusConsumerCh is a channel to notify agent a PortStatus message is received
	portStatusConsumerCh chan *openflow15.PortStatus
	// portStatusMutex is to guard the access on portStatusConsumerCh.
	portStatusMutex sync.RWMutex

	mpReplyChsMutex sync.RWMutex
	mpReplyChs      map[uint32]chan *openflow15.MultipartReply
	// tunMetadataLengthMap is used to store the tlv-map settings on the OVS bridge. Key is the index of tunnel metadata,
	// and value is the length configured in this tunnel metadata.
	tunMetadataLengthMap map[uint16]uint8
}

func (b *OFBridge) NewGroupTypeAll(id GroupIDType) Group {
	_ = "STUB: not implemented"
	return *new(Group)
}

func (b *OFBridge) NewGroup(id GroupIDType) Group { _ = "STUB: not implemented"; return *new(Group) }

func (b *OFBridge) newGroupWithType(id GroupIDType, groupType ofctrl.GroupType) Group {
	_ = "STUB: not implemented"
	return *new(Group)
}

func (b *OFBridge) NewMeter(id MeterIDType, flags ofctrl.MeterFlag) Meter {
	_ = "STUB: not implemented"
	return *new(Meter)
}

func (b *OFBridge) DeleteMeterAll() error { _ = "STUB: not implemented"; return nil }

func (b *OFBridge) DeleteGroupAll() error { _ = "STUB: not implemented"; return nil }

func (b *OFBridge) GetMeterStats(handleMeterStatsReply func(meterID int, packetCount int64)) error {
	_ = "STUB: not implemented"
	return nil
	// Represents all meters
}

func (b *OFBridge) NewTable(table Table, next uint8, missAction MissActionType) Table {
	_ = "STUB: not implemented"
	return *new(Table)
}

// GetTableByID returns the existing table by the given id. If no table exists, an error is returned.
func (b *OFBridge) GetTableByID(id uint8) (Table, error) {
	_ = "STUB: not implemented"
	return *new(Table), nil
}

// DumpTableStatus dumps table status from local cache.
func (b *OFBridge) DumpTableStatus() []TableStatus { _ = "STUB: not implemented"; return nil }

// PacketRcvd is a callback when a packetIn is received on ofctrl.OFSwitch.
func (b *OFBridge) PacketRcvd(sw *ofctrl.OFSwitch, packet *ofctrl.PacketIn) {
	_ = "STUB: not implemented"
	// Correspond to MessageStream.outbound log level.
	return
}

// SwitchConnected is a callback when the remote OFSwitch is connected.
func (b *OFBridge) SwitchConnected(sw *ofctrl.OFSwitch) { _ = "STUB: not implemented"; return }

// initialize tables.

// b.connected is nil if it is an automatic reconnection but not triggered by OFSwitch.Connect.

func (b *OFBridge) SetOFSwitch(sw *ofctrl.OFSwitch) { _ = "STUB: not implemented"; return }

// MultipartReply is a callback when multipartReply message is received on ofctrl.OFSwitch is connected.
// Client uses this method to handle the reply message if it has customized MultipartRequest message.
func (b *OFBridge) MultipartReply(sw *ofctrl.OFSwitch, rep *openflow15.MultipartReply) {
	_ = "STUB: not implemented"
	return
}

func (b *OFBridge) SwitchDisconnected(sw *ofctrl.OFSwitch) { _ = "STUB: not implemented"; return }

func (b *OFBridge) FlowGraphEnabledOnSwitch() bool { _ = "STUB: not implemented"; return false }

func (b *OFBridge) TLVMapEnabledOnSwitch() bool {
	_ = "STUB: not implemented"

	// Initialize creates ofctrl.Table for each table in the tableCache.
	return false
}

func (b *OFBridge) Initialize() { _ = "STUB: not implemented"; return }

// reset flow counts, which is needed for reconnections

// Connect initiates the connection to the OFSwitch, and initializes ofTables after connected.
func (b *OFBridge) Connect(maxRetrySec int, connectionCh chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Disconnect stops connection to the OFSwitch.
func (b *OFBridge) Disconnect() error { _ = "STUB: not implemented"; return nil }

// DumpFlows queries the Openflow entries from OFSwitch, the filter of the query is Openflow cookieID. The result is
// a map from flow cookieID to FlowStates.
func (b *OFBridge) DumpFlows(cookieID, cookieMask uint64) (map[uint64]*FlowStates, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteFlowsByCookie removes Openflow entries from OFSwitch. The removed Openflow entries use the specific CookieID.
func (b *OFBridge) DeleteFlowsByCookie(cookieID, cookieMask uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *OFBridge) IsConnected() bool { _ = "STUB: not implemented"; return false }

func (b *OFBridge) AddFlowsInBundle(addflows, modFlows, delFlows []*openflow15.FlowMod) error {
	_ = "STUB: not implemented"
	// If no Openflow entries are requested to be added or modified or deleted on the OVS bridge, return immediately.
	return nil
}

// Create a new transaction.

// Open a bundle on the OFSwitch.

// "AddFlow" operation is async, the function only returns error which occur when constructing and sending
// the BundleAdd message. An absence of error does not mean that all Openflow entries are added into the
// bundle by the switch. The number of entries successfully added to the bundle by the switch will be
// returned by function "Complete".

// Close the bundle and cancel it if there is error when adding the FlowMod message.

// Install new Openflow entries with the opened bundle.

// Modify existing Openflow entries with the opened bundle.

// Delete Openflow entries with the opened bundle.

// Close the bundle before committing it to the OFSwitch.

// This case should not be possible if all the calls to "tx.AddFlow" returned nil. This is just a sanity check.

// Commit the bundle to the OFSwitch. The "Commit" operation is sync, and the Openflow entries should be realized if
// there is no error returned.

// Update TableStatus after the transaction is committed successfully.

func (b *OFBridge) AddOFEntriesInBundle(addEntries []OFEntry, modEntries []OFEntry, delEntries []OFEntry) error {
	_ = "STUB: not implemented"
	// If no Openflow entries are requested to be added or modified or deleted on the OVS bridge, return immediately.
	return nil
}

// Classify the entries according to the EntryType, and set a correct operation type.

// Create a new transaction. Use ofctrl.Ordered to ensure the messages are realized on OVS in the order of adding
// messages. This type could ensure Group entry is realized on OVS in advance of Flow entry.

// Open a bundle on the OFSwitch.

// "AddMessage" operation is async, the function only returns error which occur when constructing and sending
// the BundleAdd message. An absence of error does not mean that all OpenFlow entries are added into the
// bundle by the switch. The number of entries successfully added to the bundle by the switch will be
// returned by function "Complete".

// Close the bundle and cancel it if there is error when adding the FlowMod message.

// Add Group modification messages in advance of Flow modification messages, so it can ensure the dependent Group
// exists when adding a new Flow entry. When OVS is deleting the Group, the corresponding Flow entry is removed
// together. It doesn't return an error when OVS is deleting a non-existing Flow entry.

// Close the bundle before committing it to the OFSwitch.

// This case should not be possible if all the calls to "tx.AddMessage" returned nil. This is just a sanity check.

// Commit the bundle to the OFSwitch. The "Commit" operation is sync, and the Openflow entries should be realized if
// there is no error returned.

// Update TableStatus after the transaction is committed successfully.

type PacketInQueue struct {
	// category is used only for logging, to help distinguish the purpose of packets.
	category       uint8
	pktRateLimiter *rate.Limiter
	pktDrops       int64
	logRateLimiter *rate.Limiter
	packetsCh      chan *ofctrl.PacketIn
}

func NewPacketInQueue(category uint8, size int, r rate.Limit) *PacketInQueue {
	_ = "STUB: not implemented"
	return nil
}

// Throttle packet drop logs to once per minute.

func (q *PacketInQueue) AddOrDrop(packet *ofctrl.PacketIn) bool {
	_ = "STUB: not implemented"
	return false
}

// Channel is full.

func (q *PacketInQueue) GetRateLimited(stopCh <-chan struct{}) *ofctrl.PacketIn {
	_ = "STUB: not implemented"
	return nil
}

func (b *OFBridge) SubscribePacketIn(category uint8, pktInQueue *PacketInQueue) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *OFBridge) SendPacketOut(packetOut *ofctrl.PacketOut) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *OFBridge) ResumePacket(packetIn *ofctrl.PacketIn) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *OFBridge) BuildPacketOut() PacketOutBuilder {
	_ = "STUB: not implemented"
	return *new(PacketOutBuilder)
}

// MaxRetry is a callback from OFController. It sets the max retry count that OFController attempts to connect to OFSwitch.
func (b *OFBridge) MaxRetry() int { _ = "STUB: not implemented"; return 0 }

// RetryInterval is a callback from OFController. It sets the interval in that the OFController will initiate next connection
// to OFSwitch if it fails this time.
func (b *OFBridge) RetryInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (b *OFBridge) PortStatusRcvd(status *openflow15.PortStatus) { _ = "STUB: not implemented"; return }

// Correspond to MessageStream.outbound log level.

// We only process add/modified status for now.

func (b *OFBridge) SubscribePortStatusConsumer(statusCh chan *openflow15.PortStatus) {
	_ = "STUB: not implemented"
	return
}

func (b *OFBridge) setPacketInFormatTo2() { _ = "STUB: not implemented"; return }

func (b *OFBridge) queryTableFeatures() { _ = "STUB: not implemented"; return }

// Use a buffer for the channel to avoid blocking the OpenFlow connection inbound channel, since it takes time when
// sending the Multipart Request messages to modify the tables' names. The buffer size "20" is the observed number
// of the Multipart Reply messages sent from OVS.

// Delete the channel which is used to receive the MultipartReply message after all tables' features are received.

func (b *OFBridge) processTableFeatures(ch chan *openflow15.MultipartReply) {
	_ = "STUB: not implemented"
	return
}

// Since the initial MultipartRequest doesn't specify any table ID, OVS will reply all tables' (except the hidden one)
// features in the reply. Here we complete the loop after we receive all the reply messages, while the reply message
// is configured with Flags=0.

// A MultipartReply message may have one or many OFPTableFeatures messages, and MultipartReply.Body is a
// slice of these messages.

// Modify table name if the table is in the pipeline, otherwise use the default table features.
// OVS doesn't allow to skip any table except the hidden table (always the last table) in a table_features
// request. So use the existing table features for the tables that Antrea doesn't define in the pipeline.

// Set table name with the configured value.

// OVS uses "Flags=0" in the last MultipartReply message to indicate all tables' features have been sent.
// Here use this mark to identify all related messages are received and complete the loop.

func (b *OFBridge) registerMpReplyCh(xid uint32, ch chan *openflow15.MultipartReply) {
	_ = "STUB: not implemented"
	return
}

func (b *OFBridge) unregisterMpReplyCh(xid uint32) { _ = "STUB: not implemented"; return }

func NewOFBridge(br string, mgmtAddr string) *OFBridge { _ = "STUB: not implemented"; return nil }

var tableID uint8

func NextTableID() (id uint8) { _ = "STUB: not implemented"; return 0 }

// ResetTableID is used to reset the initial tableID so that the table ID increases from 0.
// This function is only for test.
func ResetTableID() { _ = "STUB: not implemented"; return }
