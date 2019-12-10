// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha2 "github.com/nais/naiserator/pkg/apis/storage.cnrm.cloud.google.com/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGoogleStorageBuckets implements GoogleStorageBucketInterface
type FakeGoogleStorageBuckets struct {
	Fake *FakeStorageV1alpha2
	ns   string
}

var googlestoragebucketsResource = schema.GroupVersionResource{Group: "storage.cnrm.cloud.google.com", Version: "v1alpha2", Resource: "googlestoragebuckets"}

var googlestoragebucketsKind = schema.GroupVersionKind{Group: "storage.cnrm.cloud.google.com", Version: "v1alpha2", Kind: "GoogleStorageBucket"}

// Get takes name of the googleStorageBucket, and returns the corresponding googleStorageBucket object, and an error if there is any.
func (c *FakeGoogleStorageBuckets) Get(name string, options v1.GetOptions) (result *v1alpha2.GoogleStorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(googlestoragebucketsResource, c.ns, name), &v1alpha2.GoogleStorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.GoogleStorageBucket), err
}

// List takes label and field selectors, and returns the list of GoogleStorageBuckets that match those selectors.
func (c *FakeGoogleStorageBuckets) List(opts v1.ListOptions) (result *v1alpha2.GoogleStorageBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(googlestoragebucketsResource, googlestoragebucketsKind, c.ns, opts), &v1alpha2.GoogleStorageBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.GoogleStorageBucketList{ListMeta: obj.(*v1alpha2.GoogleStorageBucketList).ListMeta}
	for _, item := range obj.(*v1alpha2.GoogleStorageBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested googleStorageBuckets.
func (c *FakeGoogleStorageBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(googlestoragebucketsResource, c.ns, opts))

}

// Create takes the representation of a googleStorageBucket and creates it.  Returns the server's representation of the googleStorageBucket, and an error, if there is any.
func (c *FakeGoogleStorageBuckets) Create(googleStorageBucket *v1alpha2.GoogleStorageBucket) (result *v1alpha2.GoogleStorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(googlestoragebucketsResource, c.ns, googleStorageBucket), &v1alpha2.GoogleStorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.GoogleStorageBucket), err
}

// Update takes the representation of a googleStorageBucket and updates it. Returns the server's representation of the googleStorageBucket, and an error, if there is any.
func (c *FakeGoogleStorageBuckets) Update(googleStorageBucket *v1alpha2.GoogleStorageBucket) (result *v1alpha2.GoogleStorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(googlestoragebucketsResource, c.ns, googleStorageBucket), &v1alpha2.GoogleStorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.GoogleStorageBucket), err
}

// Delete takes name of the googleStorageBucket and deletes it. Returns an error if one occurs.
func (c *FakeGoogleStorageBuckets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(googlestoragebucketsResource, c.ns, name), &v1alpha2.GoogleStorageBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGoogleStorageBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(googlestoragebucketsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.GoogleStorageBucketList{})
	return err
}

// Patch applies the patch and returns the patched googleStorageBucket.
func (c *FakeGoogleStorageBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.GoogleStorageBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(googlestoragebucketsResource, c.ns, name, pt, data, subresources...), &v1alpha2.GoogleStorageBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.GoogleStorageBucket), err
}
