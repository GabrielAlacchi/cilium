// Copyright 2017-2020 Authors of Cilium
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

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	"context"
	time "time"

	ciliumiov2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	versioned "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/cilium/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v2 "github.com/cilium/cilium/pkg/k8s/client/listers/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CiliumClusterwideLocalRedirectPolicyInformer provides access to a shared informer and lister for
// CiliumClusterwideLocalRedirectPolicies.
type CiliumClusterwideLocalRedirectPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.CiliumClusterwideLocalRedirectPolicyLister
}

type ciliumClusterwideLocalRedirectPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCiliumClusterwideLocalRedirectPolicyInformer constructs a new informer for CiliumClusterwideLocalRedirectPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCiliumClusterwideLocalRedirectPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCiliumClusterwideLocalRedirectPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCiliumClusterwideLocalRedirectPolicyInformer constructs a new informer for CiliumClusterwideLocalRedirectPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCiliumClusterwideLocalRedirectPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumClusterwideLocalRedirectPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumClusterwideLocalRedirectPolicies().Watch(context.TODO(), options)
			},
		},
		&ciliumiov2.CiliumClusterwideLocalRedirectPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *ciliumClusterwideLocalRedirectPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCiliumClusterwideLocalRedirectPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ciliumClusterwideLocalRedirectPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ciliumiov2.CiliumClusterwideLocalRedirectPolicy{}, f.defaultInformer)
}

func (f *ciliumClusterwideLocalRedirectPolicyInformer) Lister() v2.CiliumClusterwideLocalRedirectPolicyLister {
	return v2.NewCiliumClusterwideLocalRedirectPolicyLister(f.Informer().GetIndexer())
}