// Copyright 2024 Antrea Authors.
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

package installation

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DenyAllConnectivityTest struct {
	networkPolicy *networkingv1.NetworkPolicy
}

func init() {
	RegisterTest("egress-deny-all-connectivity", &DenyAllConnectivityTest{networkPolicy: &networkingv1.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name: "egress-deny-all",
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: metav1.LabelSelector{
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{
						Key:      "name",
						Operator: metav1.LabelSelectorOpIn,
						Values:   []string{clientDeploymentName},
					},
				},
			},
			PolicyTypes: []networkingv1.PolicyType{
				networkingv1.PolicyTypeEgress,
			},
		},
	}})
	RegisterTest("ingress-deny-all-connectivity", &DenyAllConnectivityTest{networkPolicy: &networkingv1.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ingress-deny-all",
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: metav1.LabelSelector{
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{
						Key:      "name",
						Operator: metav1.LabelSelectorOpIn,
						Values:   []string{echoSameNodeDeploymentName, echoOtherNodeDeploymentName},
					},
				},
			},
			PolicyTypes: []networkingv1.PolicyType{
				networkingv1.PolicyTypeIngress,
			},
		},
	}})
}

func (t *DenyAllConnectivityTest) Run(ctx context.Context, testContext *testContext) error {
	_ = "STUB: not implemented"
	return nil
}

// Instead of an arbitrary sleep, we could poll and invoke tcpProbe at reasonable intervals?
// Unlike Antrea NetworkPolicies, K8s NetworkPolicies don't have a Status that can be used
// to determined whether the policy has been realized.
