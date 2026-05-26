//go:build windows
// +build windows

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

package winfirewall

import (
	"net"
)

type FWRuleDirection string

const (
	FWRuleIn  FWRuleDirection = "Inbound"
	FWRuleOut FWRuleDirection = "Outbound"
)

type fwRuleAction string

const (
	fwRuleAllow fwRuleAction = "Allow"
	fwRuleDeny  fwRuleAction = "Block"
)

type fwRuleProtocol string

const (
	fwRuleIPProtocol  fwRuleProtocol = "Any"
	fwRuleTCPProtocol fwRuleProtocol = "TCP" //nolint: unused
	fwRuleUDPProtocol fwRuleProtocol = "UDP" //nolint: unused
)

const (
	fwRuleGroup string = "Antrea"
)

type winFirewallRule struct {
	name          string
	action        fwRuleAction
	direction     FWRuleDirection
	protocol      fwRuleProtocol
	localAddress  *net.IPNet
	remoteAddress *net.IPNet
	localPorts    []uint16
	remotePorts   []uint16
}

// add adds Firewall rule on the Windows host. The name and display name of the firewall rule are the same.
func (r *winFirewallRule) add() error { _ = "STUB: not implemented"; return nil }

func (r *winFirewallRule) getCommandString() string { _ = "STUB: not implemented"; return "" }

func getPortsString(ports []uint16) string { _ = "STUB: not implemented"; return "" }

type Client struct {
}

// AddRuleAllowIP adds Windows firewall rule to accept IP packets
func (c *Client) AddRuleAllowIP(name string, direction FWRuleDirection, ipNet *net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

// AddRuleBlockIP adds Windows firewall rule to block IP packets
func (c *Client) AddRuleBlockIP(name string, direction FWRuleDirection, ipNet *net.IPNet) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) FirewallRuleExists(name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func checkDeletionError(err error) error { _ = "STUB: not implemented"; return nil }

func (c *Client) DelFirewallRuleByName(name string) error { _ = "STUB: not implemented"; return nil }

func (c *Client) DelAllFirewallRules() error { _ = "STUB: not implemented"; return nil }

func (c *Client) addIPRule(name string, direction FWRuleDirection, ipNet *net.IPNet, action fwRuleAction) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClient() *Client { _ = "STUB: not implemented"; return nil }
