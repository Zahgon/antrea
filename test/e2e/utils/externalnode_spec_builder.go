// Copyright 2022 Antrea Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	crdv1alpha1 "antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
)

type ExternalNodeSpecBuilder struct {
	spec      crdv1alpha1.ExternalNodeSpec
	name      string
	namespace string
	labels    map[string]string
}

func (t *ExternalNodeSpecBuilder) SetName(namespace string, name string) *ExternalNodeSpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (t *ExternalNodeSpecBuilder) AddInterface(name string, ips []string) *ExternalNodeSpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (t *ExternalNodeSpecBuilder) AddLabels(labels map[string]string) *ExternalNodeSpecBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (t *ExternalNodeSpecBuilder) Get() *crdv1alpha1.ExternalNode {
	_ = "STUB: not implemented"
	return nil
}
