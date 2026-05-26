// Copyright 2023 Antrea Authors
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

package servicecidr

import (
	"net"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const (
	// Disable resyncing.
	resyncPeriod time.Duration = 0
)

type EventHandler func(serviceCIDRs []*net.IPNet)

type Interface interface {
	GetServiceCIDRs() ([]*net.IPNet, error)
	// The added handlers will be called when Service CIDR changes.
	AddEventHandler(handler EventHandler)
}

type Discoverer struct {
	serviceInformer cache.SharedIndexInformer
	serviceLister   corelisters.ServiceLister
	sync.RWMutex
	serviceIPv4CIDR *net.IPNet
	serviceIPv6CIDR *net.IPNet
	eventHandlers   []EventHandler
	// queue maintains the Service objects that need to be synced.
	queue workqueue.TypedInterface[types.NamespacedName]
	// initialized indicates whether the Discoverer has been initialized.
	initialized bool
}

func NewServiceCIDRDiscoverer(serviceInformer coreinformers.ServiceInformer) *Discoverer {
	_ = "STUB: not implemented"
	return nil
}

func (d *Discoverer) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Ignore it if not found.

func (d *Discoverer) GetServiceCIDRs() ([]*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Discoverer) AddEventHandler(handler EventHandler) { _ = "STUB: not implemented"; return }

func (d *Discoverer) addService(obj interface{}) { _ = "STUB: not implemented"; return }

func (d *Discoverer) updateService(old, obj interface{}) { _ = "STUB: not implemented"; return }

func (d *Discoverer) updateServiceCIDR(svcs ...*corev1.Service) { _ = "STUB: not implemented"; return }

// If the calculated Service CIDR exists but doesn't contain the ClusterIP, calculate a new Service CIDR by
// enlarging the current Service CIDR with the ClusterIP.

// If the calculated Service CIDR doesn't exist, generate a new Service CIDR with the ClusterIP.
