// Copyright 2025 Antrea Authors
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

package objectstore

import (
	"context"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

const delayTime = time.Minute * 5

// uidGetter is the interface required for objects to be added to / deleted from the indexer
type uidGetter interface {
	GetUID() types.UID
}

// Object defines the minimum interface required for objects stored in ObjectStore
type Object interface {
	uidGetter      // GetUID()
	klog.KMetadata // GetName(), GetNamespace()
	GetCreationTimestamp() metav1.Time
}

// ObjectStore is a generic store for any Kubernetes object type
type ObjectStore[T Object] struct {
	objects         cache.Indexer
	objectsToDelete workqueue.TypedDelayingInterface[types.UID]
	delayTime       time.Duration
	// Mapping object.uuid to objectTimestamps
	timestampMap map[types.UID]*objectTimestamps
	clock        clock.Clock
	mutex        sync.RWMutex
	hasSynced    func() bool
	// Function to get object creation timestamp
	getObjectCreationTimestamp func(T, time.Time) time.Time
}

type objectTimestamps struct {
	CreationTimestamp time.Time
	// DeletionTimestamp is nil if an Object is not deleted.
	DeletionTimestamp *time.Time
}

// StoreConfig holds configuration for creating an ObjectStore
type StoreConfig[T Object] struct {
	// Provide a custom clock for the object store.
	// If omitted (nil), RealClock will be used.
	Clock clock.WithTicker
	// Provide a custom name for the deletion workqueue.
	// If omitted, a generic name will be used.
	DeleteQueueName string
	// Indexers to be added to the store.
	// If omitted, no index will be available.
	Indexers cache.Indexers
	// Filter function to use when receiving events from the informer.
	// If omitted, all objects will always be considered.
	FilterFunc func(T) bool
	// GetObjectCreationTimestamp can be used to customize how the creation timestamp is
	// determined for each object.
	// If omitted, GetCreationTimestamp will be called on the object.
	GetObjectCreationTimestamp func(T, time.Time) time.Time
}

// objectKeyFunc creates a key function that uses the object's UID
func objectKeyFunc(obj interface{}) (string, error) { _ = "STUB: not implemented"; return "", nil }

func NewObjectStore[T Object](informer cache.SharedIndexInformer, config StoreConfig[T]) *ObjectStore[T] {
	_ = "STUB: not implemented"
	return nil
}

// Invalid objects will be rejected by event handlers

// registration.HasSynced returns true when event handlers have been called for the initial list.

func (s *ObjectStore[T]) onObjectUpdate(oldObj interface{}, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// From https://pkg.go.dev/k8s.io/client-go/tools/cache#SharedInformer:
// Because `ObjectMeta.UID` has no role in identifying objects, it is possible that when (1)
// object O1 with ID (e.g. namespace and name) X and `ObjectMeta.UID` U1 in the
// SharedInformer's local cache is deleted and later (2) another object O2 with ID X and
// ObjectMeta.UID U2 is created the informer's clients are not notified of (1) and (2) but
// rather are notified only of an update from O1 to O2. Clients that need to detect such
// cases might do so by comparing the `ObjectMeta.UID` field of the old and the new object
// in the code that handles update notifications (i.e. `OnUpdate` method of
// ResourceEventHandler).

func (s *ObjectStore[T]) onObjectCreate(obj interface{}) { _ = "STUB: not implemented"; return }

func (s *ObjectStore[T]) onObjectDelete(obj interface{}) { _ = "STUB: not implemented"; return }

func (s *ObjectStore[T]) addObject(object T) error { _ = "STUB: not implemented"; return nil }

// Use configurable creation timestamp function

func (s *ObjectStore[T]) updateObject(object T) error { _ = "STUB: not implemented"; return nil }

func (s *ObjectStore[T]) deleteObject(object T) error { _ = "STUB: not implemented"; return nil }

func (s *ObjectStore[T]) checkDeletedObject(obj interface{}) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func (s *ObjectStore[T]) GetObjectByIndexAndTime(indexName, indexedValue string, time time.Time) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// In case the clocks may be skewed between different Nodes in the cluster, we directly return the object if there is only
// one object in the indexer. Otherwise, we check the timestamp for objects in the indexer.

func (s *ObjectStore[T]) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

type objectKey struct {
	uid types.UID
}

func (k *objectKey) GetUID() types.UID {
	_ = "STUB: not implemented"

	// objectKey implements the uidGetter interface
	return *new(types.UID)
}

var _ uidGetter = &objectKey{}

// worker runs a worker thread that just dequeues item from deleteQueue and
// remove the item from prevObject.
func (s *ObjectStore[T]) worker() {
	_ = "STUB: not implemented"
	// Use the same object in each worker to delete from the indexer by key
	// (UID), as there is no reason to allocate a new object for each call
	// to processDeleteQueueItem.
	return
}

func (s *ObjectStore[T]) processDeleteQueueItem(objectDeletionKey *objectKey) bool {
	_ = "STUB: not implemented"
	return false
}

// HasSynced returns true when the event handler has been called for the initial list of Objects.
func (s *ObjectStore[T]) HasSynced() bool { _ = "STUB: not implemented"; return false }

// WaitForStoreSyncs waits for stores to sync. It returns an error if the context is cancelled. You
// need to provide the HasSynced method for each store you want to wait on.
func WaitForStoreSyncs(ctx context.Context, storeSyncs ...func() bool) error {
	_ = "STUB: not implemented"
	return nil
}
