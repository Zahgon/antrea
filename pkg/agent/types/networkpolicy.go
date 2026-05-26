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

package types

import (
	"net"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	secv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

type DnsCacheEntry struct {
	FQDNName       string
	IPAddress      net.IP
	ExpirationTime time.Time
}

type MatchKey struct {
	ofProtocol    binding.Protocol
	valueCategory AddressCategory
	keyString     string
}

func (m *MatchKey) GetOFProtocol() binding.Protocol {
	_ = "STUB: not implemented"
	return *new(binding.Protocol)
}

func (m *MatchKey) GetValueCategory() AddressCategory {
	_ = "STUB: not implemented"
	return *new(AddressCategory)
}

func (m *MatchKey) GetKeyString() string { _ = "STUB: not implemented"; return "" }

func NewMatchKey(proto binding.Protocol, valueCategory AddressCategory, keyString string) *MatchKey {
	_ = "STUB: not implemented"
	return nil
}

type AddressCategory uint8

const (
	IPAddr AddressCategory = iota
	IPNetAddr
	OFPortAddr
	L4PortAddr
	ICMPAddr
	ServiceGroupIDAddr
	IGMPAddr
	LabelIDAddr
	TCPFlagsAddr
	CTStateAddr
	UnSupported
)

type AddressType int

const (
	SrcAddress AddressType = iota
	DstAddress
)

type Address interface {
	GetMatchValue() string
	GetMatchKey(addrType AddressType) *MatchKey
	GetValue() interface{}
}

type NodePolicyRule struct {
	IPSet           string
	IPSetMembers    sets.Set[string]
	Priority        *Priority
	ServiceIPTChain string
	ServiceIPTRules []string
	CoreIPTChain    string
	CoreIPTRules    []string
	IsIPv6          bool
}

// PolicyRule groups configurations to set up conjunctive match for egress/ingress policy rules.
type PolicyRule struct {
	Direction     v1beta2.Direction
	From          []Address
	To            []Address
	Service       []v1beta2.Service
	L7Protocols   []v1beta2.L7Protocol
	L7RuleVlanID  *uint32
	Action        *secv1beta1.RuleAction
	Priority      *uint16
	Name          string
	FlowID        uint32
	TableID       uint8
	PolicyRef     *v1beta2.NetworkPolicyReference
	EnableLogging bool
	LogLabel      string
}

// IsAntreaNetworkPolicyRule returns if a PolicyRule is created for Antrea NetworkPolicy types.
func (r *PolicyRule) IsAntreaNetworkPolicyRule() bool { _ = "STUB: not implemented"; return false }

// Priority is a struct that is composed of Antrea NetworkPolicy priority, rule priority and Tier priority.
// It is used as the basic unit for priority sorting.
type Priority struct {
	TierPriority   int32
	PolicyPriority float64
	RulePriority   int32
}

func (p *Priority) Less(p2 Priority) bool { _ = "STUB: not implemented"; return false }

func (p *Priority) Equals(p2 Priority) bool { _ = "STUB: not implemented"; return false }

// InSamePriorityZone returns true if two Priorities are of the same Tier and same priority at policy level.
func (p *Priority) InSamePriorityZone(p2 Priority) bool { _ = "STUB: not implemented"; return false }

// IsConsecutive returns true if two Priorties are immediately next to each other.
func (p *Priority) IsConsecutive(p2 Priority) bool { _ = "STUB: not implemented"; return false }

// ByPriority sorts a list of Priority by their relative TierPriority, PolicyPriority and RulePriority, in that order.
// It implements sort.Interface.
type ByPriority []Priority

func (bp ByPriority) Len() int           { _ = "STUB: not implemented"; return 0 }
func (bp ByPriority) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (bp ByPriority) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type RuleMetric struct {
	Bytes, Packets, Sessions uint64
}

func (m *RuleMetric) Merge(m1 *RuleMetric) { _ = "STUB: not implemented"; return }

// A BitRange is a representation of a range of values from base value with a
// bitmask applied.
type BitRange struct {
	Value uint16
	Mask  *uint16
}
