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

package e2e

type Pod string

type CustomPod struct {
	Pod    Pod
	Labels map[string]string
}

func NewPod(namespace string, podName string) Pod { _ = "STUB: not implemented"; return *new(Pod) }

func (pod Pod) String() string { _ = "STUB: not implemented"; return "" }

func (pod Pod) split() (string, string) { _ = "STUB: not implemented"; return "", "" }

func (pod Pod) Namespace() string { _ = "STUB: not implemented"; return "" }

func (pod Pod) PodName() string { _ = "STUB: not implemented"; return "" }

type PodConnectivityMark string

const (
	Connected PodConnectivityMark = "Con"
	Unknown   PodConnectivityMark = "Unk"
	Error     PodConnectivityMark = "Err"
	Dropped   PodConnectivityMark = "Drp"
	Rejected  PodConnectivityMark = "Rej"
)

type Connectivity struct {
	From         Pod
	To           Pod
	Connectivity PodConnectivityMark
}

type ConnectivityTable struct {
	Items   []string
	itemSet map[string]bool
	Values  map[string]map[string]PodConnectivityMark
}

type TruthTable struct {
	Items   []string
	itemSet map[string]bool
	Values  map[string]map[string]bool
}

func NewConnectivityTable(items []string, defaultValue *PodConnectivityMark) *ConnectivityTable {
	_ = "STUB: not implemented"
	return nil
}

// IsComplete returns true if there's a value set for every single pair of items, otherwise it returns false.
func (tt *TruthTable) IsComplete() bool { _ = "STUB: not implemented"; return false }

func (ct *ConnectivityTable) Set(from string, to string, value PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (ct *ConnectivityTable) SetAllFrom(from string, value PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (ct *ConnectivityTable) SetAllTo(to string, value PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (ct *ConnectivityTable) Get(from string, to string) PodConnectivityMark {
	_ = "STUB: not implemented"
	return *new(PodConnectivityMark)
}

func (ct *ConnectivityTable) Compare(other *ConnectivityTable) *TruthTable {
	_ = "STUB: not implemented"
	// TODO set equality
	// if tt.itemSet != other.itemSet {
	//	panic()
	// }
	return nil
}

// TODO other.Get(from, to) ?

// TODO check for equality from both sides

func (ct *ConnectivityTable) PrettyPrint(indent string) string {
	_ = "STUB: not implemented"
	return ""
}

func (tt *TruthTable) PrettyPrint(indent string) string { _ = "STUB: not implemented"; return "" }

type Reachability struct {
	Expected        *ConnectivityTable
	Observed        *ConnectivityTable
	Pods            []Pod
	PodsByNamespace map[string][]Pod
}

func NewReachability(pods []Pod, defaultExpectation PodConnectivityMark) *Reachability {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reachability) NewReachabilityWithSameExpectations() *Reachability {
	_ = "STUB: not implemented"
	return nil
}

// ExpectConn is an experimental way to describe connectivity with named fields
func (r *Reachability) ExpectConn(spec *Connectivity) { _ = "STUB: not implemented"; return }

func (r *Reachability) Expect(pod1 Pod, pod2 Pod, connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (r *Reachability) ExpectSelf(allPods []Pod, connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

// ExpectAllIngress defines that any traffic going into the pod will be allowed/dropped/rejected
func (r *Reachability) ExpectAllIngress(pod Pod, connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

// ExpectAllEgress defines that any traffic going out of the pod will be allowed/dropped/rejected
func (r *Reachability) ExpectAllEgress(pod Pod, connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (r *Reachability) ExpectAllSelfNamespace(connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (r *Reachability) ExpectSelfNamespace(namespace string, connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (r *Reachability) ExpectIngressFromNamespace(pod Pod, namespace string, connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (r *Reachability) ExpectEgressToNamespace(pod Pod, namespace string, connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (r *Reachability) ExpectNamespaceIngressFromNamespace(dstNamespace, srcNamespace string, connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (r *Reachability) ExpectNamespaceEgressToNamespace(srcNamespace, dstNamespace string, connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (r *Reachability) Observe(pod1 Pod, pod2 Pod, connectivity PodConnectivityMark) {
	_ = "STUB: not implemented"
	return
}

func (r *Reachability) Summary() (trueObs int, falseObs int, comparison *TruthTable) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (r *Reachability) PrintSummary(printExpected bool, printObserved bool, printComparison bool) {
	_ = "STUB: not implemented"
	return
}

type NPEvalActionType string

const (
	NPEvalNone    NPEvalActionType = "<NONE>"
	NPEvalAllow   NPEvalActionType = "Allow"
	NPEvalDrop    NPEvalActionType = "Drop"
	NPEvalIsolate NPEvalActionType = "Isolate"
	NPEvalReject  NPEvalActionType = "Reject"
)

type NPEvaluationSpec struct {
	Source      Pod
	Destination Pod
	NPName      string
	Action      NPEvalActionType
}

type NPEvaluation struct {
	itemSet    map[string]bool
	Assertions []*NPEvaluationSpec
}

func NewNPEvaluation(pods []Pod) *NPEvaluation { _ = "STUB: not implemented"; return nil }

func (e *NPEvaluation) Expect(from, to Pod, npName string, action NPEvalActionType) *NPEvaluation {
	_ = "STUB: not implemented"
	return nil
}

func (e *NPEvaluation) ExpectNone(from Pod, to Pod) *NPEvaluation {
	_ = "STUB: not implemented"
	return nil
}
