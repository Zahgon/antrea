// Copyright 2021 Antrea Authors
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

// GroupSpecBuilder builds a Group object.
type GroupSpecBuilder struct {
	Spec      crdv1beta1.GroupSpec
	Name      string
	Namespace string
}

func (b *GroupSpecBuilder) Get() *crdv1beta1.Group { _ = "STUB: not implemented"; return nil }

func (b *GroupSpecBuilder) SetName(name string) *GroupSpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *GroupSpecBuilder) SetNamespace(namespace string) *GroupSpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *GroupSpecBuilder) SetPodSelector(podSelector map[string]string, podSelectorMatchExp []metav1.LabelSelectorRequirement) *GroupSpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *GroupSpecBuilder) SetNamespaceSelector(nsSelector map[string]string, nsSelectorMatchExp []metav1.LabelSelectorRequirement) *GroupSpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *GroupSpecBuilder) SetIPBlocks(ipBlocks []crdv1beta1.IPBlock) *GroupSpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *GroupSpecBuilder) SetServiceReference(svcNS, svcName string) *GroupSpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (b *GroupSpecBuilder) SetChildGroups(cgs []string) *GroupSpecBuilder {
	_ = "STUB: not implemented"
	return nil
}
