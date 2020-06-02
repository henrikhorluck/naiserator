// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	naisiov1alpha1 "github.com/nais/naiserator/pkg/apis/nais.io/v1alpha1"
	versioned "github.com/nais/naiserator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nais/naiserator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/nais/naiserator/pkg/client/listers/nais.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AzureAdApplicationInformer provides access to a shared informer and lister for
// AzureAdApplications.
type AzureAdApplicationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AzureAdApplicationLister
}

type azureAdApplicationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAzureAdApplicationInformer constructs a new informer for AzureAdApplication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAzureAdApplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAzureAdApplicationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAzureAdApplicationInformer constructs a new informer for AzureAdApplication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAzureAdApplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NaisV1alpha1().AzureAdApplications(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NaisV1alpha1().AzureAdApplications(namespace).Watch(options)
			},
		},
		&naisiov1alpha1.AzureAdApplication{},
		resyncPeriod,
		indexers,
	)
}

func (f *azureAdApplicationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAzureAdApplicationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *azureAdApplicationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&naisiov1alpha1.AzureAdApplication{}, f.defaultInformer)
}

func (f *azureAdApplicationInformer) Lister() v1alpha1.AzureAdApplicationLister {
	return v1alpha1.NewAzureAdApplicationLister(f.Informer().GetIndexer())
}
