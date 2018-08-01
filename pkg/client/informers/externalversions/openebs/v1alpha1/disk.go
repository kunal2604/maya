/*
Copyright 2018 The OpenEBS Authors

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	openebs_io_v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	versioned "github.com/openebs/maya/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openebs/maya/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openebs/maya/pkg/client/listers/openebs/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// DiskInformer provides access to a shared informer and lister for
// Disks.
type DiskInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DiskLister
}

type diskInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewDiskInformer constructs a new informer for Disk type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDiskInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.OpenebsV1alpha1().Disks().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.OpenebsV1alpha1().Disks().Watch(options)
			},
		},
		&openebs_io_v1alpha1.Disk{},
		resyncPeriod,
		indexers,
	)
}

func defaultDiskInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewDiskInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *diskInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&openebs_io_v1alpha1.Disk{}, defaultDiskInformer)
}

func (f *diskInformer) Lister() v1alpha1.DiskLister {
	return v1alpha1.NewDiskLister(f.Informer().GetIndexer())
}
