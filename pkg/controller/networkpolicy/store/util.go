// Copyright 2019 Antrea Authors
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

package store

import (
	"antrea.io/antrea/v2/pkg/apiserver/storage"
)

// keyAndSpanSelectFunc returns whether the provided selectors matches the key and/or the nodeNames.
func keyAndSpanSelectFunc(selectors *storage.Selectors, key string, obj interface{}) bool {
	_ = "STUB: not implemented"
	// If Key is present in selectors, the provided key must match it.
	return false
}

// If nodeName is present in selectors's Field selector, the provided nodeNames must contain it.

// isSelected determines if the previous and the current version of an object should be selected by the given selectors.
func isSelected(key string, prevObj, currObj interface{}, selectors *storage.Selectors, isInitEvent bool) (bool, bool) {
	_ = "STUB: not implemented"
	// We have filtered out init events that we are not interested in, so the current object must be selected.
	return false, false
}
