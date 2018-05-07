/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	time "time"

	warmimage_v2 "github.com/mattmoor/warm-image/pkg/apis/warmimage/v2"
	versioned "github.com/mattmoor/warm-image/pkg/client/clientset/versioned"
	internalinterfaces "github.com/mattmoor/warm-image/pkg/client/informers/externalversions/internalinterfaces"
	v2 "github.com/mattmoor/warm-image/pkg/client/listers/warmimage/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WarmImageInformer provides access to a shared informer and lister for
// WarmImages.
type WarmImageInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.WarmImageLister
}

type warmImageInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWarmImageInformer constructs a new informer for WarmImage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWarmImageInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWarmImageInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWarmImageInformer constructs a new informer for WarmImage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWarmImageInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MattmoorV2().WarmImages(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MattmoorV2().WarmImages(namespace).Watch(options)
			},
		},
		&warmimage_v2.WarmImage{},
		resyncPeriod,
		indexers,
	)
}

func (f *warmImageInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWarmImageInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *warmImageInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&warmimage_v2.WarmImage{}, f.defaultInformer)
}

func (f *warmImageInformer) Lister() v2.WarmImageLister {
	return v2.NewWarmImageLister(f.Informer().GetIndexer())
}
