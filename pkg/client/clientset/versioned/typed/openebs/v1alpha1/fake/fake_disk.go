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
package fake

import (
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDisks implements DiskInterface
type FakeDisks struct {
	Fake *FakeOpenebsV1alpha1
}

var disksResource = schema.GroupVersionResource{Group: "openebs.io", Version: "v1alpha1", Resource: "disks"}

var disksKind = schema.GroupVersionKind{Group: "openebs.io", Version: "v1alpha1", Kind: "Disk"}

// Get takes name of the disk, and returns the corresponding disk object, and an error if there is any.
func (c *FakeDisks) Get(name string, options v1.GetOptions) (result *v1alpha1.Disk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(disksResource, name), &v1alpha1.Disk{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disk), err
}

// List takes label and field selectors, and returns the list of Disks that match those selectors.
func (c *FakeDisks) List(opts v1.ListOptions) (result *v1alpha1.DiskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(disksResource, disksKind, opts), &v1alpha1.DiskList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DiskList{}
	for _, item := range obj.(*v1alpha1.DiskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested disks.
func (c *FakeDisks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(disksResource, opts))
}

// Create takes the representation of a disk and creates it.  Returns the server's representation of the disk, and an error, if there is any.
func (c *FakeDisks) Create(disk *v1alpha1.Disk) (result *v1alpha1.Disk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(disksResource, disk), &v1alpha1.Disk{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disk), err
}

// Update takes the representation of a disk and updates it. Returns the server's representation of the disk, and an error, if there is any.
func (c *FakeDisks) Update(disk *v1alpha1.Disk) (result *v1alpha1.Disk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(disksResource, disk), &v1alpha1.Disk{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disk), err
}

// Delete takes name of the disk and deletes it. Returns an error if one occurs.
func (c *FakeDisks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(disksResource, name), &v1alpha1.Disk{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDisks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(disksResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DiskList{})
	return err
}

// Patch applies the patch and returns the patched disk.
func (c *FakeDisks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Disk, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(disksResource, name, data, subresources...), &v1alpha1.Disk{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Disk), err
}
