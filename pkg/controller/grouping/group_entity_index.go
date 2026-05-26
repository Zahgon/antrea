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

package grouping

import (
	"sync"
	"sync/atomic"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"

	"antrea.io/antrea/v2/pkg/apis/crd/v1alpha2"
	"antrea.io/antrea/v2/pkg/controller/types"
)

const (
	// Cluster scoped selectors are stored under empty Namespace in the selectorItemIndex.
	emptyNamespace = ""
	// Antrea could add custom labels using CustomLabelKeyPrefix+CustomLabelKeyXXX as
	// the label key to entities for internal process.
	CustomLabelKeyPrefix         = "internal.antrea.io/"
	CustomLabelKeyServiceAccount = "service-account"
)

var (
	// eventChanSize is declared as a variable to allow overriding for testing.
	eventChanSize = 1000
)

type eventHandler func(group string)

// GroupType is a public type used to differentiate Groups.
type GroupType string

// Interface provides methods to query entities that a given group selects and groups that select a given entity.
// It maintains indexes between groups and entities to make the query efficient. It supports callers to register
// callbacks that will be called when a specific type of groups' entities are updated.
type Interface interface {
	// AddGroup adds or updates a group to the index. The caller can then get entities selected by this group.
	AddGroup(groupType GroupType, name string, selector *types.GroupSelector)
	// DeleteGroup deletes a group from the index.
	DeleteGroup(groupType GroupType, name string)
	// AddEventHandler registers an eventHandler for the given type of groups. When any Pod/ExternelEntity/Namespace
	// update affects the given kind of groups, the eventHandler will be called with the affected groups.
	// The eventHandler is supposed to execute quickly and not perform blocking operation. Blocking operation should be
	// deferred to a routine that is triggered by the eventHandler, like the eventHandler + workqueue pattern.
	AddEventHandler(groupType GroupType, handler eventHandler)
	// GetEntities returns the selected Pods or ExternalEntities for the given group.
	GetEntities(groupType GroupType, name string) ([]*v1.Pod, []*v1alpha2.ExternalEntity)
	// GetGroupsForPod returns the groups that select the given Pod.
	GetGroupsForPod(namespace, name string) (map[GroupType][]string, bool)
	// GetGroupsForExternalEntity returns the groups that select the given ExternalEntity.
	GetGroupsForExternalEntity(namespace, name string) (map[GroupType][]string, bool)
	// AddPod adds or updates a Pod to the index. If any existing groups are affected, eventHandlers will be called with
	// the affected groups.
	AddPod(pod *v1.Pod)
	// DeletePod deletes a Pod from the index. If any existing groups are affected, eventHandlers will be called with
	// the affected groups.
	DeletePod(pod *v1.Pod)
	// AddExternalEntity adds or updates an ExternalEntity to the index. If any existing groups are affected,
	// eventHandlers will be called with the affected groups.
	AddExternalEntity(ee *v1alpha2.ExternalEntity)
	// DeleteExternalEntity deletes an ExternalEntity from the index. If any existing groups are affected, eventHandlers
	// will be called with the affected groups.
	DeleteExternalEntity(ee *v1alpha2.ExternalEntity)
	// AddNamespace adds or updates a Namespace to the index. If any existing groups are affected, eventHandlers will be
	// called with the affected groups.
	AddNamespace(namespace *v1.Namespace)
	// DeleteNamespace deletes a Namespace to the index. If any existing groups are affected, eventHandlers will be
	// called with the affected groups.
	DeleteNamespace(namespace *v1.Namespace)
	// Run starts the index.
	Run(stopCh <-chan struct{})
	// HasSynced returns true if the interface has been initialized with the full lists of Pods, Namespaces, and
	// ExternalEntities.
	HasSynced() bool
}

// entityType is an internal type used to differentiate Pod from ExternalEntity.
type entityType int

const (
	podEntityType entityType = iota
	externalEntityType
)

// entityItem contains an entity (either Pod or ExternalEntity) and some relevant information.
type entityItem struct {
	// entity is either a Pod or an ExternalEntity.
	entity metav1.Object
	// labelItemKey is the key of the labelItem that the entityItem is associated with.
	// entityItems will be associated with the same labelItem if they have same Namespace, entityType, and labels.
	labelItemKey string
}

// labelItem represents an individual label set. It's the actual object that will be matched with label selectors.
// Entities of same type in same Namespace having same labels will share a labelItem.
type labelItem struct {
	// The label set that will be used for matching.
	labels labels.Set
	// The Namespace of the entities that share the labelItem.
	namespace string
	// The type of the entities that share the labelItem.
	entityType entityType
	// The keys of the entityItems that share the labelItem.
	entityItemKeys sets.Set[string]
	// The keys of the selectorItems that match the labelItem.
	selectorItemKeys sets.Set[string]
}

// groupItem contains a group's metadata and its selector.
type groupItem struct {
	// The type of the group.
	groupType GroupType
	// The name of the group. It must be unique within its own type.
	name string
	// The selector of the group.
	selector *types.GroupSelector
	// selectorItemKey is the key of the selectorItem that the groupItem is associated with.
	// groupItems will be associated with the same selectorItem if they have same selector.
	selectorItemKey string
}

// selectorItem represents an individual label selector. It's the actual object that will be matched with label sets.
// Groups having same label selector will share a selectorItem.
type selectorItem struct {
	// The label selector that will be used for matching.
	selector *types.GroupSelector
	// The keys of the groupItems that share the selectorItem.
	groupItemKeys sets.Set[string]
	// The keys of the labelItems that match the selectorItem.
	labelItemKeys sets.Set[string]
}

var _ Interface = &GroupEntityIndex{}

// GroupEntityIndex implements Interface.
//
// It abstracts label set and label selector from entities and groups and does actual matching against the formers to
// avoid redundant calculation given that most entities (Pod or ExternalEntity) actually share labels. For example, Pods
// that managed by a deployment controller have same labels.
//
// It maintains indexes between label set and label selector so that querying label sets that match a label selector and
// reversed queries can be performed with a constant time complexity. Indirectly, querying entities that match a group
// and reversed queries can be performed with same complexity.
//
// The relationship of the four items are like below:
// entityItem <===> labelItem <===> selectorItem <===> groupItem
type GroupEntityIndex struct {
	lock sync.RWMutex

	// entityItems stores all entityItems.
	entityItems map[string]*entityItem

	// labelItems stores all labelItems.
	labelItems map[string]*labelItem
	// labelItemIndex is nested map from entityType to Namespace to keys of labelItems.
	// It's used to filter potential labelItems when matching a Namespace scoped selectorItem.
	labelItemIndex map[entityType]map[string]sets.Set[string]

	// groupItems stores all groupItems.
	groupItems map[string]*groupItem

	// selectorItems stores all selectorItems.
	selectorItems map[string]*selectorItem
	// selectorItemIndex is nested map from entityType to Namespace to keys of selectorItems.
	// It's used to filter potential selectorItems when matching an labelItem.
	// Cluster scoped selectorItems are stored under empty Namespace "".
	selectorItemIndex map[entityType]map[string]sets.Set[string]

	// namespaceLabels stores label sets of all Namespaces.
	namespaceLabels map[string]labels.Set

	// eventHandlers is a map from group type to a list of handlers. When a type of group's updated, the corresponding
	// event handlers will be called with the group name provided.
	eventHandlers map[GroupType][]eventHandler

	// eventChan is channel used for calling eventHandlers asynchronously.
	eventChan chan string

	// synced stores a boolean value, which tracks if the GroupEntityIndex has been initialized with the full lists of
	// Pods, Namespaces, and ExternalEntities.
	synced *atomic.Value
}

// NewGroupEntityIndex creates a GroupEntityIndex.
func NewGroupEntityIndex() *GroupEntityIndex { _ = "STUB: not implemented"; return nil }

func (i *GroupEntityIndex) GetEntities(groupType GroupType, name string) ([]*v1.Pod, []*v1alpha2.ExternalEntity) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get the selectorItem the group is associated with.

// Get the keys of the labelItems the selectorItem matches.

// Collect the entityItems that share the labelItem.

func (i *GroupEntityIndex) GetGroupsForPod(namespace, name string) (map[GroupType][]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (i *GroupEntityIndex) GetGroupsForExternalEntity(namespace, name string) (map[GroupType][]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (i *GroupEntityIndex) getGroups(entityType entityType, namespace, name string) (map[GroupType][]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Get the selectorItem the group is associated with.

// Get the keys of the selectorItems the labelItem matches.

// Collect the groupItems that share the selectorItem.

func (i *GroupEntityIndex) AddNamespace(namespace *v1.Namespace) { _ = "STUB: not implemented"; return }

// Do nothing if labels are not updated.

// Resync cluster scoped selectors as they may start or stop matching the Namespace because of the label update.

// Cluster scoped selectors are stored under empty Namespace in the selectorItemIndex.

// If the selector selects all Namespaces, it won't be affected.

// By default, the selector selects Pods. It selects ExternalEntities only if ExternalEntitySelector is set
// explicitly.

// Only labelItems in this Namespace may be affected.

// Notify watchers if the selectorItem is updated.

func (i *GroupEntityIndex) DeleteNamespace(namespace *v1.Namespace) {
	_ = "STUB: not implemented"
	return
}

// deleteEntityFromLabelItem disconnects an entityItem from a labelItem.
// The labelItem will be deleted if it's no longer used by any entityItem.
func (i *GroupEntityIndex) deleteEntityFromLabelItem(label, entity string) *labelItem {
	_ = "STUB: not implemented"
	return nil
}

// If the labelItem is still used by any entities, keep it. Otherwise delete it.

// Delete the labelItem itself.

// Delete it from the labelItemIndex.

// Delete the labelItem from matched selectorItems.

// createLabelItem creates a labelItem based on the provided entityItem.
// It's called when there is no existing labelItem for a label set.
func (i *GroupEntityIndex) createLabelItem(entityType entityType, eItem *entityItem, labels map[string]string) *labelItem {
	_ = "STUB: not implemented"
	return nil
}

// Create the labelItem.

// Add it to the labelItemIndex.

// Scan potential selectorItems and associate the new labelItem with the matched ones.

// SelectorItems in the same Namespace may match the labelItem.

// Cluster scoped selectorItems may match the labelItem.

func (i *GroupEntityIndex) AddPod(pod *v1.Pod) {
	_ = "STUB: not implemented"
	// Create a new map to add custom labels to avoid changing the original labels and
	// introducing data race.
	return
}

func (i *GroupEntityIndex) AddExternalEntity(ee *v1alpha2.ExternalEntity) {
	_ = "STUB: not implemented"
	return
}

func (i *GroupEntityIndex) addEntity(entityType entityType, entity metav1.Object, labels map[string]string) {
	_ = "STUB: not implemented"
	return
}

// If its label doesn't change, its labelItem won't change. We still need to dispatch the updates of the groups
// that select the entity if the entity's attributes that we care about are updated.

// Delete the Pod from the previous labelItem as its label is updated.

// Create a labelItem if it doesn't exist.

// Notify group updates.

// If entity is updated, all previously and currently matched selectors are affected. Otherwise only the
// difference portion are affected.

func (i *GroupEntityIndex) DeletePod(pod *v1.Pod) { _ = "STUB: not implemented"; return }

func (i *GroupEntityIndex) DeleteExternalEntity(ee *v1alpha2.ExternalEntity) {
	_ = "STUB: not implemented"
	return
}

func (i *GroupEntityIndex) deleteEntity(entityType entityType, entity metav1.Object) {
	_ = "STUB: not implemented"
	return
}

// Delete the entity from its associated labelItem and entityItems.

// All selectorItems that match the labelItem are affected.

// deleteGroupFromSelectorItem disconnects a groupItem from a selectorItem.
// The selectorItem will be deleted if it's no longer used by any groupItem.
func (i *GroupEntityIndex) deleteGroupFromSelectorItem(sKey, gKey string) *selectorItem {
	_ = "STUB: not implemented"
	return nil
}

// If the selectorItem is still used by any groups, keep it. Otherwise delete it.

// Delete the selectorItem itself.

// Delete it from the selectorItemIndex.

// Delete the selectorItem from matched labelItems.

// createSelectorItem creates a selectorItem based on the provided groupItem.
// It's called when there is no existing selectorItem for a group selector.
func (i *GroupEntityIndex) createSelectorItem(gItem *groupItem) *selectorItem {
	_ = "STUB: not implemented"
	return nil
}

// Create the selectorItem.

// Add it to the selectorItemIndex.

// Scan potential labelItems and associates the new selectorItem with the matched ones.

// The selector is Namespace scoped, it can only match labelItems in this Namespace.

// The selector is Cluster scoped and has non-empty NamespaceSelector, scan labelItems in a Namespace only if
// the Namespace's labels match.

// The selector is Cluster scoped and match all Namespaces.

// scanLabelItems scans potential labelItems and updates their association.
func (i *GroupEntityIndex) scanLabelItems(labelItemKeys sets.Set[string], sItem *selectorItem) bool {
	_ = "STUB: not implemented"
	return false
}

// Connect the selector and the label if they didn't match before, otherwise do nothing.

// Disconnect the selector and the label if they matched before, otherwise do nothing.

func (i *GroupEntityIndex) AddGroup(groupType GroupType, name string, selector *types.GroupSelector) {
	_ = "STUB: not implemented"
	return
}

// Its selector doesn't change, do nothing.

// Create a selectorItem if it doesn't exist.

func (i *GroupEntityIndex) DeleteGroup(groupType GroupType, name string) {
	_ = "STUB: not implemented"
	return
}

// Delete the group from its associated selectorItem and groupItems.

// notify notifies the affected groups to eventHandlers.
// It's supposed to be called with the lock held as it accesses the selectorItems. Normally the method shouldn't block
// as the event channel is buffered and the consumer Run should execute quickly. If it blocks in practice, we should
// review whether there are unexpected blocking eventHandlers, or consider moving the routine out of locking.
func (i *GroupEntityIndex) notify(selector string) { _ = "STUB: not implemented"; return }

func (i *GroupEntityIndex) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (i *GroupEntityIndex) AddEventHandler(groupType GroupType, handler eventHandler) {
	_ = "STUB: not implemented"
	return
}

func (i *GroupEntityIndex) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (i *GroupEntityIndex) setSynced(synced bool) { _ = "STUB: not implemented"; return }

func (i *GroupEntityIndex) match(entityType entityType, label labels.Set, namespace string, sel *types.GroupSelector) bool {
	_ = "STUB: not implemented"
	return false
}

// Pods or ExternalEntities must be matched within the same Namespace.

// podSelector or externalEntitySelector exists but doesn't match the ExternalEntity or Pod's labels.

// Pod's Namespace does not match namespaceSelector.

// ExternalEntity or Pod's Namespace matches namespaceSelector but
// labels do not match the podSelector or externalEntitySelector.

// Selector only has a PodSelector/ExternalEntitySelector and no sel.Namespace.
// Pods/ExternalEntities must be matched from all Namespaces.

// pod/ee labels do not match PodSelector/ExternalEntitySelector.

// The group selects nothing when all selectors are missing.

func entityAttrsUpdated(oldEntity, newEntity metav1.Object) bool {
	_ = "STUB: not implemented"
	return false
}

// For Pod, we only care about PodIP and NodeName update.
// Also, when a Pod is updated to terminated state, the selectorItems need to be
// notified so that they are excluded from any Network Policy computations in
// appliedTo or address groups.
// Some other attributes we care about are immutable, e.g. the named ContainerPort.

// getEntityItemKey returns the entity key used in entityItems.
func getEntityItemKey(entityType entityType, entity metav1.Object) string {
	_ = "STUB: not implemented"
	return ""
}

// getEntityItemKeyByName returns the entity key used in entityItems.
func getEntityItemKeyByName(entityType entityType, namespace, name string) string {
	_ = "STUB: not implemented"
	return ""
}

// getLabelItemKey returns the label key used in labelItems.
func getLabelItemKey(entityType entityType, obj metav1.Object, allLabels map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// getGroupItemKey returns the group key used in groupItems.
func getGroupItemKey(groupType GroupType, name string) string { _ = "STUB: not implemented"; return "" }

// getSelectorItemKey returns the selector key used in selectorItems.
func getSelectorItemKey(selector *types.GroupSelector) string { _ = "STUB: not implemented"; return "" }
