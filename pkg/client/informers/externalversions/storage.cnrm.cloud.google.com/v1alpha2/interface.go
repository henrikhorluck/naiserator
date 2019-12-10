// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	internalinterfaces "github.com/nais/naiserator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// GoogleStorageBuckets returns a GoogleStorageBucketInformer.
	GoogleStorageBuckets() GoogleStorageBucketInformer
	// GoogleStorageBucketAccessControls returns a GoogleStorageBucketAccessControlInformer.
	GoogleStorageBucketAccessControls() GoogleStorageBucketAccessControlInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// GoogleStorageBuckets returns a GoogleStorageBucketInformer.
func (v *version) GoogleStorageBuckets() GoogleStorageBucketInformer {
	return &googleStorageBucketInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GoogleStorageBucketAccessControls returns a GoogleStorageBucketAccessControlInformer.
func (v *version) GoogleStorageBucketAccessControls() GoogleStorageBucketAccessControlInformer {
	return &googleStorageBucketAccessControlInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
