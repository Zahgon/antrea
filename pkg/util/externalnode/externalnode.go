// Copyright 2022 Antrea Authors
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

package externalnode

import (
	// #nosec G505: not used for security purposes

	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha1"
	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha2"
)

const (
	EntityOwnerKind = "ExternalNode"

	interfaceNameLength = 5
)

func GenExternalEntityName(externalNode *v1alpha1.ExternalNode) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// This should not happen since openAPIV3Schema checks it.

// Only one network interface is supported now.
// Other interfaces except interfaces[0] will be ignored if there are more than one interfaces.

// #nosec G401: not used for security purposes

func GenerateEntityNodeKey(externalEntity *v1alpha2.ExternalEntity) string {
	_ = "STUB: not implemented"
	return ""
}
