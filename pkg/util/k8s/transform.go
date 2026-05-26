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

package k8s

import (
	"k8s.io/client-go/tools/cache"
)

// NewTrimmer returns a cache.TransformFunc that can be used to trim objects stored in informers.
// The function must be idempotent before client-go v0.31.0 to avoid a race condition happening when objects were
// accessed during Resync operation, see https://github.com/kubernetes/kubernetes/issues/124337.
// But it's generally more efficient to avoid trimming the same object more than once.
//
// Note: be cautious when adding fields to be trimmed, ensuring they do not inadvertently clear the original values when
// objects are updated by Antrea to kube-apiserver.
func NewTrimmer(extraTrimmers ...cache.TransformFunc) cache.TransformFunc {
	_ = "STUB: not implemented"
	return *new(cache.TransformFunc)
}

// It means the objects has been trimmed.

// Trim common fields for all objects.
// According to https://kubernetes.io/docs/reference/using-api/server-side-apply/#clearing-managedfields,
// setting the managedFields to an empty list will not reset the field when the objects are updated, so it's
// safe to trim it for all objects.

// Trim type specific fields for each type.

// TrimPod clears unused fields from a Pod that are not required by Antrea.
// It's safe to do so because Antrea only patches Pod.
func TrimPod(obj interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// TrimNode clears unused fields from a Node that are not required by Antrea.
// It's safe to do so because Antrea only patches Node.
func TrimNode(obj interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }
