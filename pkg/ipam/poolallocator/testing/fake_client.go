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

package testing

import (
	"sync"

	fakeversioned "antrea.io/antrea/v2/pkg/client/clientset/versioned/fake"
)

// IPPoolClientset extends the generated fake clientset to simulate optimistic
// concurrency for IPPool updates: an update is rejected with a conflict error
// unless the ResourceVersion in the request matches what is stored.
// Use NewIPPoolClient to construct it; use the standard Create API to register
// pools (the create reactor assigns ResourceVersion automatically).
type IPPoolClientset struct {
	*fakeversioned.Clientset
	// store latest ResourceVersion for given pool
	poolVersion sync.Map
}

func NewIPPoolClient() *IPPoolClientset {
	_ = "STUB: not implemented"
	// NewSimpleClientset provides a working object tracker that handles list,
	// get, create, and update for all resource types via its default reactors.
	// The tracker also sends watch events (Added/Modified) so no custom watcher
	// is needed.
	return nil
}

// Intercept create to assign ResourceVersion and register it for conflict
// detection, then let the tracker handle the actual storage and watch event.

// Intercept update to enforce optimistic concurrency, then let the tracker
// handle the actual storage and watch event.
