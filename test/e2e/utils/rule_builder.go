// Copyright 2025 Antrea Authors
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

package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

type RuleBuilder interface {
	GetIngress() crdv1beta1.Rule
	GetEgress() crdv1beta1.Rule
}

type BaseRuleBuilder struct {
	Protoc               AntreaPolicyProtocol
	Port                 *int32
	PortName             *string
	EndPort              *int32
	ICMPType             *int32
	ICMPCode             *int32
	IGMPType             *int32
	GroupAddress         *string
	PodSelector          map[string]string
	NSSelector           map[string]string
	PodSelectorMatchExp  []metav1.LabelSelectorRequirement
	NodeSelectorMatchExp []metav1.LabelSelectorRequirement
	NSSelectorMatchExp   []metav1.LabelSelectorRequirement
	Action               crdv1beta1.RuleAction
	Name                 string
	SelfNS               bool
	SrcPort              *int32
	SrcEndPort           *int32
	IPBlock              *crdv1beta1.IPBlock
}

type ACNPRuleBuilder struct {
	BaseRuleBuilder
	NodeSelector     map[string]string
	Namespaces       *crdv1beta1.PeerNamespaces
	AppliedToSpecs   []ACNPAppliedToSpec
	ServiceAccount   *crdv1beta1.NamespacedName
	RuleClusterGroup string
}

type ANNPRuleBuilder struct {
	BaseRuleBuilder
	L7Protocols        []crdv1beta1.L7Protocol
	RuleGroup          string
	EESelector         map[string]string
	EESelectorMatchExp []metav1.LabelSelectorRequirement
	AppliedToSpecs     []ANNPAppliedToSpec
}

func toEgress(ingressRule crdv1beta1.Rule) crdv1beta1.Rule {
	_ = "STUB: not implemented"
	return *new(crdv1beta1.Rule)
}

func (rb ANNPRuleBuilder) GetEgress() crdv1beta1.Rule {
	_ = "STUB: not implemented"
	return *new(crdv1beta1.Rule)
}

func (rb ANNPRuleBuilder) GetIngress() crdv1beta1.Rule {
	_ = "STUB: not implemented"
	return *new(crdv1beta1.Rule)
}

// An empty From/To in ANNP rules evaluates to match all addresses.

func (rb ACNPRuleBuilder) GetIngress() crdv1beta1.Rule {
	_ = "STUB: not implemented"
	return *new(crdv1beta1.Rule)
}

// An empty From/To in ACNP rules evaluates to match all addresses.

func (rb ACNPRuleBuilder) GetEgress() crdv1beta1.Rule {
	_ = "STUB: not implemented"
	return *new(crdv1beta1.Rule)
}

func (rb BaseRuleBuilder) generatePodSelector() (podSel *metav1.LabelSelector) {
	_ = "STUB: not implemented"
	return nil
}

func (rb BaseRuleBuilder) generateNSSelector() (nsSel *metav1.LabelSelector) {
	_ = "STUB: not implemented"
	return nil
}
