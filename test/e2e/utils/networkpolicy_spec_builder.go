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
	v1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NetworkPolicySpecBuilder struct {
	Spec      networkingv1.NetworkPolicySpec
	Name      string
	Namespace string
}

func (n *NetworkPolicySpecBuilder) Get() *networkingv1.NetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkPolicySpecBuilder) SetPodSelector(labels map[string]string) *NetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkPolicySpecBuilder) SetName(namespace string, name string) *NetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Add tests to match expressions
func (n *NetworkPolicySpecBuilder) AddIngress(protoc v1.Protocol, port *int32, portName *string, cidr *string, exceptCIDRs []string,
	podSelector map[string]string, nsSelector map[string]string,
	podSelectorMatchExp []metav1.LabelSelectorRequirement, nsSelectorMatchExp []metav1.LabelSelectorRequirement) *NetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

// AddEgressDNS mutates the nth policy rule to allow DNS, convenience method
func (n *NetworkPolicySpecBuilder) WithEgressDNS() *NetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkPolicySpecBuilder) AddEgress(protoc v1.Protocol, port *int32, portName *string, cidr *string, exceptCIDRs []string,
	podSelector map[string]string, nsSelector map[string]string,
	podSelectorMatchExp []metav1.LabelSelectorRequirement, nsSelectorMatchExp []metav1.LabelSelectorRequirement) *NetworkPolicySpecBuilder {
	_ = "STUB: not implemented"

	// For simplicity, we just reuse the Ingress code here.  The underlying data model for ingress/egress is identical
	// With the exception of calling the rule `To` vs. `From`.
	return nil
}

func (n *NetworkPolicySpecBuilder) SetTypeIngress() *NetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkPolicySpecBuilder) SetTypeEgress() *NetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkPolicySpecBuilder) SetTypeBoth() *NetworkPolicySpecBuilder {
	_ = "STUB: not implemented"
	return nil
}
