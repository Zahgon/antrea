// Copyright 2022 Antrea Authors
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

package deploy

import (
	"net/http"

	"github.com/spf13/cobra"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
)

const (
	leaderRole = "leader"
	memberRole = "member"

	latestVersionURL = "https://raw.githubusercontent.com/antrea-io/antrea/main/multicluster/build/yamls"
	downloadURL      = "https://github.com/antrea-io/antrea/releases/download"
	leaderYAML       = "antrea-multicluster-leader.yml"
	memberYAML       = "antrea-multicluster-member.yml"
)

var httpGet = http.Get
var getAPIGroupResources = getAPIGroupResourcesWrapper

func generateManifests(role string, version string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createResources(cmd *cobra.Command, apiGroupResources []*restmapper.APIGroupResources, dynamicClient dynamic.Interface, content []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func deploy(cmd *cobra.Command, role string, version string, namespace string, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec G107

func getAPIGroupResourcesWrapper(k8sClient kubernetes.Interface) ([]*restmapper.APIGroupResources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
