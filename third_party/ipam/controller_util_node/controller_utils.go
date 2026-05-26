/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
/*
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

Modifies:
- Relocate imports to antrea.io/third_party
- Cleanup unused imports due to code removal and relocation
- Remove the following unused funcs to avoid additional code imports:
	RecordNodeEvent()
	DeletePods()
	SetPodTerminationReason()
	MarkPodsNotReady()
	SwapNodeControllerTaint()
	AddOrUpdateLabelsOnNode()
    GetNodeCondition()
 - In RecordNodeStatusChange(): Remove recorder parameter, remove ref initialization and comment out recorder.Eventf
   call
*/

package node

// RecordNodeStatusChange records a event related to a node status change. (Common to lifecycle and ipam)
func RecordNodeStatusChange( /*recorder record.EventRecorder,*/ node *v1.Node, newStatus string) {
	_ = "STUB: not implemented"
	//ref := &v1.ObjectReference{
	//	APIVersion: "v1",
	//	Kind:       "Node",
	//	Name:       node.Name,
	//	UID:        node.UID,
	//	Namespace:  "",
	//}
	return
}

// TODO: This requires a transaction, either both node status is updated
// and event is recorded or neither should happen, see issue #6055.
//recorder.Eventf(ref, v1.EventTypeNormal, newStatus, "Node %s status is now: %s", node.Name, newStatus)

// CreateAddNodeHandler creates an add node handler.
func CreateAddNodeHandler(f func(node *v1.Node) error) func(obj interface{}) {
	_ = "STUB: not implemented"
	return nil
}

// CreateUpdateNodeHandler creates a node update handler. (Common to lifecycle and ipam)
func CreateUpdateNodeHandler(f func(oldNode, newNode *v1.Node) error) func(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return nil
}

// CreateDeleteNodeHandler creates a delete node handler. (Common to lifecycle and ipam)
func CreateDeleteNodeHandler(f func(node *v1.Node) error) func(obj interface{}) {
	_ = "STUB: not implemented"
	return nil
}

// We can get DeletedFinalStateUnknown instead of *v1.Node here and
// we need to handle that correctly. #34692
