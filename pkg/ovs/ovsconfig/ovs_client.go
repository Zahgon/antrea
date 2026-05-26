// Copyright 2019 Antrea Authors
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

package ovsconfig

import (
	"net"
	"time"

	"github.com/TomCodeLV/OVSDB-golang-lib/pkg/ovsdb"
)

const defaultOVSDBFile = "db.sock"

type OVSBridge struct {
	ovsdb                    *ovsdb.OVSDB
	name                     string
	datapathType             OVSDatapathType
	mcastSnoopingEnable      bool
	uuid                     string
	isHardwareOffloadEnabled bool
	requiredPortExternalIDs  []string
}

type OVSPortData struct {
	UUID   string
	Name   string
	VLANID uint16
	// Interface type.
	IFType      string
	IFName      string
	OFPort      int32
	ExternalIDs map[string]string
	Options     map[string]string
	MAC         net.HardwareAddr
}

const (
	openvSwitchSchema = "Open_vSwitch"
	// Openflow protocol version 1.0.
	openflowProtoVersion10 = "OpenFlow10"
	// Openflow protocol version 1.5.
	openflowProtoVersion15 = "OpenFlow15"
	// Maximum allowed value of ofPortRequest.
	ofPortRequestMax = 65279
	hardwareOffload  = "hw-offload"
)

// NewOVSDBConnectionUDS connects to the OVSDB server on the UNIX domain socket
// or named pipe (on Windows) specified by address, never using any SSL connection option.
// If address is set to "", the default UNIX domain socket path
// "/run/openvswitch/db.sock" will be used.
// Returns the OVSDB struct on success.
func NewOVSDBConnectionUDS(address string) (*ovsdb.OVSDB, Error) {
	_ = "STUB: not implemented"
	return nil, *new(Error)
}

// For the sake of debugging, we keep logging messages until the
// connection is successful. We use exponential backoff to determine the
// sleep duration between two successive log messages (up to
// maxBackoffTime).

type OVSBridgeOption func(*OVSBridge)

func WithRequiredPortExternalIDs(keys ...string) OVSBridgeOption {
	_ = "STUB: not implemented"
	return *new(OVSBridgeOption)
}

func WithMcastSnooping() OVSBridgeOption { _ = "STUB: not implemented"; return *new(OVSBridgeOption) }

// NewOVSBridge creates and returns a new OVSBridge struct.
func NewOVSBridge(bridgeName string, ovsDatapathType OVSDatapathType, ovsdb *ovsdb.OVSDB, options ...OVSBridgeOption) OVSBridgeClient {
	_ = "STUB: not implemented"
	return *new(OVSBridgeClient)
}

// Create looks up or creates the bridge. If the bridge with name bridgeName
// does not exist, it will be created. Openflow protocol version 1.0 and 1.5
// will be enabled for the bridge.
func (br *OVSBridge) Create() Error { _ = "STUB: not implemented"; return *new(Error) }

// Update OpenFlow protocol versions and datapath type on existent bridge.

func (br *OVSBridge) lookupByName() (bool, Error) {
	_ = "STUB: not implemented"
	return false, *new(Error)
}

func (br *OVSBridge) updateBridgeConfiguration() Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

// Use Openflow protocol version 1.0 and 1.5.

func (br *OVSBridge) create() Error { _ = "STUB: not implemented"; return *new(Error) }

// Use Openflow protocol version 1.0 and 1.5.

func (br *OVSBridge) Delete() Error { _ = "STUB: not implemented"; return *new(Error) }

// GetExternalIDs returns the external IDs of the bridge.
func (br *OVSBridge) GetExternalIDs() (map[string]string, Error) {
	_ = "STUB: not implemented"
	return nil, *new(Error)
}

// SetExternalIDs sets the provided external IDs to the bridge.
func (br *OVSBridge) SetExternalIDs(externalIDs map[string]interface{}) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

// SetDatapathID sets the provided datapath ID to the bridge.
// If datapath ID is not configured, reconfigure bridge(add/delete port or set different Mac address for local port)
// will change its datapath ID. And the change of datapath ID and interrupt OpenFlow connection.
// See question "My bridge disconnects from my controller on add-port/del-port" in：
// http://openvswitch.org/support/dist-docs-2.5/FAQ.md.html
func (br *OVSBridge) SetDatapathID(datapathID string) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

func (br *OVSBridge) GetDatapathID() (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

func (br *OVSBridge) WaitForDatapathID(timeout time.Duration) (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

// GetPortUUIDList returns UUIDs of all ports on the bridge.
func (br *OVSBridge) GetPortUUIDList() ([]string, Error) {
	_ = "STUB: not implemented"
	return nil, *new(Error)
}

// DeletePorts deletes ports in portUUIDList on the bridge
func (br *OVSBridge) DeletePorts(portUUIDList []string) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

// DeletePort deletes the port with the provided portUUID.
// If the port does not exist no change will be done.
func (br *OVSBridge) DeletePort(portUUID string) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

// CreateInternalPort creates an internal port with the specified name on the
// bridge.
// If externalIDs is not empty, the map key/value pairs will be set to the
// port's external_ids.
// If ofPortRequest is not zero, it will be passed to the OVS port creation.
func (br *OVSBridge) CreateInternalPort(name string, ofPortRequest int32, mac string, externalIDs map[string]interface{}) (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

// CreateTunnelPort creates a tunnel port with the specified name and type on
// the bridge.
// If ofPortRequest is not zero, it will be passed to the OVS port creation.
func (br *OVSBridge) CreateTunnelPort(name string, tunnelType TunnelType, ofPortRequest int32) (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

// CreateTunnelPortExt creates a tunnel port with the specified name and type
// on the bridge.
// If ofPortRequest is not zero, it will be passed to the OVS port creation.
// If remoteIP is not empty, it will be set to the tunnel port interface
// options; otherwise flow based tunneling will be configured.
// psk is for the pre-shared key of IPsec ESP tunnel. If it is not empty, it
// will be set to the tunnel port interface options. Flow based IPsec tunnel is
// not supported, so remoteIP must be provided too when psk is not empty.
// If externalIDs is not nil, the IDs in it will be added to the port's
// external_ids.
func (br *OVSBridge) CreateTunnelPortExt(
	name string,
	tunnelType TunnelType,
	ofPortRequest int32,
	csum bool,
	localIP string,
	remoteIP string,
	remoteName string,
	psk string,
	extraOptions map[string]interface{},
	externalIDs map[string]interface{}) (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

func (br *OVSBridge) createTunnelPort(
	name string,
	tunnelType TunnelType,
	ofPortRequest int32,
	csum bool,
	localIP string,
	remoteIP string,
	remoteName string,
	psk string,
	extraOptions map[string]interface{},
	externalIDs map[string]interface{}) (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

// Flow based tunnel.

// GetInterfaceOptions returns the options of the provided interface.
func (br *OVSBridge) GetInterfaceOptions(name string) (map[string]string, Error) {
	_ = "STUB: not implemented"
	return nil, *new(Error)
}

// SetInterfaceOptions sets the specified options of the provided interface.
func (br *OVSBridge) SetInterfaceOptions(name string, options map[string]interface{}) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

// ParseTunnelInterfaceOptions reads remote IP, local IP, IPsec PSK, and csum
// from the tunnel interface options and returns them.
func ParseTunnelInterfaceOptions(portData *OVSPortData) (net.IP, net.IP, int32, string, string, bool) {
	_ = "STUB: not implemented"
	return *new(net.IP), *new(net.IP), 0, "", "", false
}

// CreateUplinkPort creates uplink port.
func (br *OVSBridge) CreateUplinkPort(name string, ofPortRequest int32, externalIDs map[string]interface{}) (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

// CreatePort creates a port with the specified name on the bridge, and connects
// the interface specified by ifDev to the port.
// If externalIDs is not empty, the map key/value pairs will be set to the
// port's external_ids.
func (br *OVSBridge) CreatePort(name, ifDev string, externalIDs map[string]interface{}) (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

// CreateAccessPort creates a port with the specified name and VLAN ID on the bridge, and connects
// the interface specified by ifDev to the port.
// If externalIDs is not empty, the map key/value pairs will be set to the
// port's external_ids.
// vlanID=0 will perform same behavior as CreatePort.
func (br *OVSBridge) CreateAccessPort(name, ifDev string, externalIDs map[string]interface{}, vlanID uint16) (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

func (br *OVSBridge) createPort(name, ifName, ifType string, ofPortRequest int32, vlanID uint16, mac string, externalIDs, options map[string]interface{}) (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

// GetOFPort retrieves the ofport value of an interface given the interface name.
// The function will invoke OVSDB "wait" operation with 5 seconds timeout to
// wait the ofport is set on the interface, and so could be blocked for 5
// seconds. If the "wait" operation times out or the interface is not found, or
// the ofport is invalid, value 0 and an error will be returned.
// If waitUntilValid is true, the function will wait the ofport is not -1 with
// 5 seconds timeout. This parameter is used after the interface type is changed
// by the client.
func (br *OVSBridge) GetOFPort(ifName string, waitUntilValid bool) (int32, Error) {
	_ = "STUB: not implemented"
	return 0, *new(Error)
}

// If an OVS port is newly created, the ofport field is expected to change from empty to a int value.

// If an OVS port is updated from invalid status to valid, the ofport field is expected to change from "-1" to a
// value that is larger than 0.

// ofport value -1 means that the interface could not be created due to an error.

func makeOVSDBSetFromList(list []string) []interface{} { _ = "STUB: not implemented"; return nil }

func buildMapFromOVSDBMap(data []interface{}) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Should not be possible

func buildPortDataCommon(port, intf map[string]interface{}, portData *OVSPortData) {
	_ = "STUB: not implemented"
	return
}

// ofport not assigned by OVS yet

// GetPortData retrieves port data given the OVS port UUID and interface name.
// nil is returned, if the port or interface could not be found, or the
// interface is not attached to the port.
// The port's OFPort will be set to 0, if its ofport is not assigned by OVS yet.
func (br *OVSBridge) GetPortData(portUUID, ifName string) (*OVSPortData, Error) {
	_ = "STUB: not implemented"
	return nil, *new(Error)
}

// GetPortList returns all ports on the bridge.
// A port's OFPort will be set to 0, if its ofport is not assigned by OVS yet.
func (br *OVSBridge) GetPortList() ([]OVSPortData, Error) {
	_ = "STUB: not implemented"
	return nil, *new(Error)
}

// Port should have one interface

// GetOVSVersion either returns the version of OVS, or an error.
func (br *OVSBridge) GetOVSVersion() (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

// parseOvsVersion parses the version from an interface type, which can be a map of string[interface] or string[string], and returns it as a string, we have special logic here so that a panic doesn't happen.
func parseOvsVersion(ovsReturnRow interface{}) (string, Error) {
	_ = "STUB: not implemented"
	return "", *new(Error)
}

// AddOVSOtherConfig adds the given configs to the "other_config" column of
// the single record in the "Open_vSwitch" table.
// For each config, it will only be added if its key doesn't already exist.
// No error is returned if configs already exist.
func (br *OVSBridge) AddOVSOtherConfig(configs map[string]interface{}) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

func (br *OVSBridge) GetOVSOtherConfig() (map[string]string, Error) {
	_ = "STUB: not implemented"
	return nil, *new(Error)
}

// UpdateOVSOtherConfig updates the given configs to the "other_config" column of
// the single record in the "Open_vSwitch" table.
// For each config, it will be updated if the existing value does not match the given one,
// and it will be added if its key does not exist.
// It the configs are already up to date, this function will be a no-op.
func (br *OVSBridge) UpdateOVSOtherConfig(configs map[string]interface{}) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

// DeleteOVSOtherConfig deletes the given configs from the "other_config" column of
// the single record in the "Open_vSwitch" table.
// For each config, it will be deleted if its key exists and the given value is empty string or
// its value matches the given one. No error is returned if configs don't exist or don't match.
func (br *OVSBridge) DeleteOVSOtherConfig(configs map[string]interface{}) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

// AddBridgeOtherConfig adds the given configs to the "other_config" column of
// the single record in the "Bridge" table.
// For each config, it will only be added if its key doesn't already exist.
// No error is returned if configs already exist.
func (br *OVSBridge) AddBridgeOtherConfig(configs map[string]interface{}) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

func (br *OVSBridge) GetBridgeName() string { _ = "STUB: not implemented"; return "" }

func (br *OVSBridge) IsHardwareOffloadEnabled() bool { _ = "STUB: not implemented"; return false }

func (br *OVSBridge) getHardwareOffload() (bool, Error) {
	_ = "STUB: not implemented"
	return false, *new(Error)
}

func (br *OVSBridge) GetOVSDatapathType() OVSDatapathType {
	_ = "STUB: not implemented"
	return *

	// SetInterfaceType modifies the OVS Interface type to the given ifType.
	// This function is used on Windows when the Pod interface is created after the OVS port creation.
	new(OVSDatapathType)
}

func (br *OVSBridge) SetInterfaceType(name, ifType string) Error {
	_ = "STUB: not implemented"
	// Update Interface type, and the caller ensures the host Interface exists.
	return *new(Error)
}

func (br *OVSBridge) SetPortExternalIDs(portName string, externalIDs map[string]interface{}) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

func (br *OVSBridge) GetPortExternalIDs(portName string) (map[string]string, Error) {
	_ = "STUB: not implemented"
	return nil, *new(Error)
}

func (br *OVSBridge) SetInterfaceMTU(name string, MTU int) error {
	_ = "STUB: not implemented"
	return nil
}

func (br *OVSBridge) SetInterfaceMAC(name string, mac net.HardwareAddr) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

func (br *OVSBridge) GetBridgeMcastSnoopingEnable() (bool, Error) {
	_ = "STUB: not implemented"
	return false, *new(Error)
}
