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

package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

type ClusterNetworkPolicySpecBuilder struct {
	Spec crdv1beta1.ClusterNetworkPolicySpec
	Name string
}

type ACNPAppliedToSpec struct {
	PodSelector          map[string]string
	NodeSelector         map[string]string
	NSSelector           map[string]string
	PodSelectorMatchExp  []metav1.LabelSelectorRequirement
	NodeSelectorMatchExp []metav1.LabelSelectorRequirement
	NSSelectorMatchExp   []metav1.LabelSelectorRequirement
	Group                string
	Service              *crdv1beta1.NamespacedName
}

func (b *ClusterNetworkPolicySpecBuilder) Get() *crdv1beta1.ClusterNetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterNetworkPolicySpecBuilder) SetName(name string) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterNetworkPolicySpecBuilder) SetPriority(p float64) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterNetworkPolicySpecBuilder) SetTier(tier string) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterNetworkPolicySpecBuilder) SetAppliedToGroup(specs []ACNPAppliedToSpec) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func ACNPGetAppliedToPeer(podSelector map[string]string,
	nodeSelector map[string]string,
	nsSelector map[string]string,
	podSelectorMatchExp []metav1.LabelSelectorRequirement,
	nodeSelectorMatchExp []metav1.LabelSelectorRequirement,
	nsSelectorMatchExp []metav1.LabelSelectorRequirement,
	appliedToCG string,
	service *crdv1beta1.NamespacedName) crdv1beta1.AppliedTo {
	_ = "STUB: not implemented"
	return *new(crdv1beta1.AppliedTo)
}

func (b *ClusterNetworkPolicySpecBuilder) AddIngress(rb RuleBuilder) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterNetworkPolicySpecBuilder) AddEgress(rb RuleBuilder) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterNetworkPolicySpecBuilder) AddNodeSelectorRule(nodeSelector *metav1.LabelSelector, protoc AntreaPolicyProtocol, port *int32, name string,
	ruleAppliedToSpecs []ACNPAppliedToSpec, action crdv1beta1.RuleAction, isEgress bool) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterNetworkPolicySpecBuilder) AddFQDNRule(fqdn string,
	protoc AntreaPolicyProtocol, port *int32, portName *string, endPort *int32, name string,
	ruleAppliedToSpecs []ACNPAppliedToSpec, action crdv1beta1.RuleAction) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterNetworkPolicySpecBuilder) AddToServicesRule(svcRefs []crdv1beta1.PeerService,
	name string, ruleAppliedToSpecs []ACNPAppliedToSpec, action crdv1beta1.RuleAction) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterNetworkPolicySpecBuilder) AddStretchedIngressRule(pSel, nsSel map[string]string,
	name string, ruleAppliedToSpecs []ACNPAppliedToSpec, action crdv1beta1.RuleAction) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

// AddEgressDNS mutates the nth policy rule to allow DNS, convenience method
func (b *ClusterNetworkPolicySpecBuilder) WithEgressDNS() *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterNetworkPolicySpecBuilder) AddEgressLogging(logLabel string) *ClusterNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}
