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

type AntreaNetworkPolicySpecBuilder struct {
	Spec      crdv1beta1.NetworkPolicySpec
	Name      string
	Namespace string
}

type ANNPAppliedToSpec struct {
	ExternalEntitySelector         map[string]string
	ExternalEntitySelectorMatchExp []metav1.LabelSelectorRequirement
	PodSelector                    map[string]string
	PodSelectorMatchExp            []metav1.LabelSelectorRequirement
	Group                          string
}

func (b *AntreaNetworkPolicySpecBuilder) Get() *crdv1beta1.NetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (b *AntreaNetworkPolicySpecBuilder) SetName(namespace string, name string) *AntreaNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AntreaNetworkPolicySpecBuilder) SetPriority(p float64) *AntreaNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AntreaNetworkPolicySpecBuilder) SetTier(tier string) *AntreaNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AntreaNetworkPolicySpecBuilder) SetAppliedToGroup(specs []ANNPAppliedToSpec) *AntreaNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func ANNPGetAppliedToPeer(podSelector map[string]string,
	podSelectorMatchExp []metav1.LabelSelectorRequirement,
	entitySelector map[string]string,
	entitySelectorMatchExp []metav1.LabelSelectorRequirement,
	appliedToGrp string) crdv1beta1.AppliedTo {
	_ = "STUB: not implemented"
	return *new(crdv1beta1.AppliedTo)
}

func (b *AntreaNetworkPolicySpecBuilder) AddIngress(rb RuleBuilder) *AntreaNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AntreaNetworkPolicySpecBuilder) AddEgress(rb RuleBuilder) *AntreaNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AntreaNetworkPolicySpecBuilder) AddToServicesRule(svcRefs []crdv1beta1.PeerService,
	name string, ruleAppliedToSpecs []ANNPAppliedToSpec, action crdv1beta1.RuleAction) *AntreaNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AntreaNetworkPolicySpecBuilder) AddEgressLogging(logLabel string) *AntreaNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *AntreaNetworkPolicySpecBuilder) AddFQDNRule(fqdn string,
	protoc AntreaPolicyProtocol, port *int32, portName *string, endPort *int32, name string,
	specs []ANNPAppliedToSpec, action crdv1beta1.RuleAction) *AntreaNetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}
