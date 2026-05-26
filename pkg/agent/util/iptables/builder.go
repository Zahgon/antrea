//go:build !windows
// +build !windows

// Copyright 2024 Antrea Authors
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
	"strings"

	"k8s.io/apimachinery/pkg/util/intstr"

	"antrea.io/antrea/v2/pkg/agent/util/ipset"
)

type iptablesRule struct {
	chain string
	specs *strings.Builder
}

type iptablesRuleBuilder struct {
	iptablesRule
}

func NewRuleBuilder(chain string) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) writeSpec(spec string) { _ = "STUB: not implemented"; return }

func (b *iptablesRuleBuilder) MatchCIDRSrc(cidr string) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) MatchCIDRDst(cidr string) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) SetLogPrefix(prefix string) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) MatchIPSetSrc(ipsetName string, ipsetType ipset.SetType) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) MatchIPSetDst(ipsetName string, ipsetType ipset.SetType) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) MatchTransProtocol(protocol string) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) MatchPortDst(port *intstr.IntOrString, endPort *int32) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) MatchPortSrc(port, endPort *int32) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) MatchICMP(icmpType, icmpCode *int32, ipProtocol Protocol) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) MatchEstablishedOrRelated() IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) MatchInputInterface(interfaceName string) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) MatchOutputInterface(interfaceName string) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) SetTarget(target string) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) SetTargetDNATToDst(dnatIP string, dnatPort *int32) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) SetComment(comment string) IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) CopyBuilder() IPTablesRuleBuilder {
	_ = "STUB: not implemented"
	return *new(IPTablesRuleBuilder)
}

func (b *iptablesRuleBuilder) Done() IPTablesRule {
	_ = "STUB: not implemented"
	return *new(IPTablesRule)
}

func (e *iptablesRule) GetRule() string { _ = "STUB: not implemented"; return "" }
