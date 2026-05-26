//go:build !windows
// +build !windows

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

package iptables

import (
	"github.com/coreos/go-iptables/iptables"
	"k8s.io/apimachinery/pkg/util/intstr"

	"antrea.io/antrea/v2/pkg/agent/util/ipset"
)

const (
	NATTable    = "nat"
	FilterTable = "filter"
	MangleTable = "mangle"
	RawTable    = "raw"

	AcceptTarget     = "ACCEPT"
	DropTarget       = "DROP"
	MasqueradeTarget = "MASQUERADE"
	MarkTarget       = "MARK"
	ReturnTarget     = "RETURN"
	ConnTrackTarget  = "CT"
	NoTrackTarget    = "NOTRACK"
	SNATTarget       = "SNAT"
	DNATTarget       = "DNAT"
	RejectTarget     = "REJECT"
	NotrackTarget    = "NOTRACK"
	LOGTarget        = "LOG"

	PreRoutingChain  = "PREROUTING"
	InputChain       = "INPUT"
	ForwardChain     = "FORWARD"
	PostRoutingChain = "POSTROUTING"
	OutputChain      = "OUTPUT"

	waitSeconds              = 10
	waitIntervalMicroSeconds = 200000
)

type Protocol byte

var protocolStrMap = map[Protocol]string{
	ProtocolIPv4: "IPv4",
	ProtocolIPv6: "IPv6",
}

func (p Protocol) String() string { _ = "STUB: not implemented"; return "" }

const (
	ProtocolDual Protocol = iota
	ProtocolIPv4
	ProtocolIPv6
)

const (
	ProtocolTCP    = "tcp"
	ProtocolUDP    = "udp"
	ProtocolSCTP   = "sctp"
	ProtocolICMP   = "icmp"
	ProtocolICMPv6 = "icmp6"
)

const (
	// https://netfilter.org/projects/iptables/files/changes-iptables-1.6.2.txt:
	// iptables-restore: support acquiring the lock.
	restoreWaitSupportedMinVersion = "v1.6.2"

	// https://netfilter.org/projects/iptables/files/changes-iptables-1.6.0.txt:
	// iptables: snat: add randomize-full support
	// https://netfilter.org/projects/iptables/files/changes-iptables-1.6.2.txt:
	// iptables: masquerade: add randomize-full support
	// In our case, we do not differentiate between SNAT and MASQUERADE support for the option,
	// and we use 1.6.2 as the common minimum version number.
	randomFullySupportedMinVersion = "v1.6.2"
)

type Interface interface {
	EnsureChain(protocol Protocol, table string, chain string) error

	ChainExists(protocol Protocol, table string, chain string) (bool, error)

	AppendRule(protocol Protocol, table string, chain string, ruleSpec []string) error

	InsertRule(protocol Protocol, table string, chain string, ruleSpec []string) error

	DeleteRule(protocol Protocol, table string, chain string, ruleSpec []string) error

	DeleteChain(protocol Protocol, table string, chain string) error

	ListRules(protocol Protocol, table string, chain string) (map[Protocol][]string, error)

	Restore(data string, flush bool, useIPv6 bool) error

	Save() ([]byte, error)

	HasRandomFully() bool
}

type IPTablesRuleBuilder interface {
	MatchCIDRSrc(cidr string) IPTablesRuleBuilder
	MatchCIDRDst(cidr string) IPTablesRuleBuilder
	MatchIPSetSrc(ipset string, ipsetType ipset.SetType) IPTablesRuleBuilder
	MatchIPSetDst(ipset string, ipsetType ipset.SetType) IPTablesRuleBuilder
	MatchTransProtocol(protocol string) IPTablesRuleBuilder
	MatchPortDst(port *intstr.IntOrString, endPort *int32) IPTablesRuleBuilder
	MatchPortSrc(port, endPort *int32) IPTablesRuleBuilder
	MatchICMP(icmpType, icmpCode *int32, ipProtocol Protocol) IPTablesRuleBuilder
	MatchEstablishedOrRelated() IPTablesRuleBuilder
	MatchInputInterface(interfaceName string) IPTablesRuleBuilder
	MatchOutputInterface(interfaceName string) IPTablesRuleBuilder
	SetLogPrefix(prefix string) IPTablesRuleBuilder
	SetTarget(target string) IPTablesRuleBuilder
	SetTargetDNATToDst(dnatIP string, dnatPort *int32) IPTablesRuleBuilder
	SetComment(comment string) IPTablesRuleBuilder
	CopyBuilder() IPTablesRuleBuilder
	Done() IPTablesRule
}

type IPTablesRule interface {
	GetRule() string
}

type Client struct {
	ipts map[Protocol]*iptables.IPTables
	// restoreWaitSupported indicates whether iptables-restore (or ip6tables-restore) supports --wait flag.
	restoreWaitSupported bool
	// randomFullySupported indicates whether --random-fully is supported for SNAT and MASQUERADE rules.
	randomFullySupported bool
}

func New(enableIPV4, enableIPV6 bool) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

func isRestoreWaitSupported(ipt *iptables.IPTables) bool { _ = "STUB: not implemented"; return false }

func isRandomFullySupported(ipt *iptables.IPTables) bool {
	_ = "STUB: not implemented"
	// Note that even if the iptables version supports it, the kernel version may not.
	// For SNAT rules, kernel >= 3.14 is required. For MASQUERADE rules, kernel >= 3.13 is required.
	// Given how old these kernel releases are, we do not check the version here. This is
	// consistent with how K8s checks for --random-fully support:
	// https://github.com/kubernetes/kubernetes/blob/60c4c2b2521fb454ce69dee737e3eb91a25e0535/pkg/util/iptables/iptables.go#L239
	return false
}

// EnsureChain checks if target chain already exists, creates it if not.
func (c *Client) EnsureChain(protocol Protocol, table string, chain string) error {
	_ = "STUB: not implemented"
	return nil
}

// ChainExists checks if target chain already exists in a table
func (c *Client) ChainExists(protocol Protocol, table string, chain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AppendRule checks if target rule already exists with the protocol, appends it if not.
func (c *Client) AppendRule(protocol Protocol, table string, chain string, ruleSpec []string) error {
	_ = "STUB: not implemented"
	return nil
}

// InsertRule checks if target rule already exists, inserts it at the beginning of the chain if not.
func (c *Client) InsertRule(protocol Protocol, table string, chain string, ruleSpec []string) error {
	_ = "STUB: not implemented"
	return nil
}

func matchProtocol(ipt *iptables.IPTables, protocol Protocol) bool {
	_ = "STUB: not implemented"
	return false
}

// DeleteRule checks if target rule already exists, deletes the rule if found.
func (c *Client) DeleteRule(protocol Protocol, table string, chain string, ruleSpec []string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteChain deletes all rules from a chain in a table and then delete the chain.
func (c *Client) DeleteChain(protocol Protocol, table string, chain string) error {
	_ = "STUB: not implemented"
	return nil
}

// ListRules lists all rules from a chain in a table.
func (c *Client) ListRules(protocol Protocol, table string, chain string) (map[Protocol][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Restore calls iptable-restore to restore iptables with the provided content.
// If flush is true, all previous contents of the respective tables will be flushed.
// Otherwise only involved chains will be flushed. Restore supports "ip6tables-restore" for IPv6.
func (c *Client) Restore(data string, flush bool, useIPv6 bool) error {
	_ = "STUB: not implemented"
	return nil
}

// We acquire xtables lock for iptables-restore to prevent it from conflicting
// with iptables/iptables-restore which might being called by kube-proxy.
// iptables supports "--wait" option and go-iptables has enabled it.
// iptables-restore doesn't support the option until 1.6.2. We use "-w" if the
// detected version is greater than or equal to 1.6.2, otherwise we acquire the
// file lock explicitly.
// Note that we cannot just acquire the file lock explicitly for all cases because
// iptables-restore will try acquiring the lock with or without "-w" provided since 1.6.2.

// Save calls iptables-saves to dump chains and tables in iptables.
func (c *Client) Save() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// HasRandomFully returns true if the iptables version supports --random-fully for SNAT and
// MASQUERADE rules.
func (c *Client) HasRandomFully() bool { _ = "STUB: not implemented"; return false }

func MakeChainLine(chain string) string { _ = "STUB: not implemented"; return "" }

func IsIPv6Protocol(protocol Protocol) bool { _ = "STUB: not implemented"; return false }
