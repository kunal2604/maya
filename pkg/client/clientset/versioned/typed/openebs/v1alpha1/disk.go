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
package v1alpha1

import (
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	scheme "github.com/openebs/maya/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DisksGetter has a method to return a DiskInterface.
// A group's client should implement this interface.
type DisksGetter interface {
	Disks() DiskInterface
}

// DiskInterface has methods to work with Disk resources.
type DiskInterface interface {
	Create(*v1alpha1.Disk) (*v1alpha1.Disk, error)
	Update(*v1alpha1.Disk) (*v1alpha1.Disk, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Disk, error)
	List(opts v1.ListOptions) (*v1alpha1.DiskList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Disk, err error)
	DiskExpansion
}

// disks implements DiskInterface
type disks struct {
	client rest.Interface
}

// newDisks returns a Disks
func newDisks(c *OpenebsV1alpha1Client) *disks {
	return &disks{
		client: c.RESTClient(),
	}
}

// Get takes name of the disk, and returns the corresponding disk object, and an error if there is any.
func (c *disks) Get(name string, options v1.GetOptions) (result *v1alpha1.Disk, err error) {
	result = &v1alpha1.Disk{}
	err = c.client.Get().
		Resource("disks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Disks that match those selectors.
func (c *disks) List(opts v1.ListOptions) (result *v1alpha1.DiskList, err error) {
	result = &v1alpha1.DiskList{}
	err = c.client.Get().
		Resource("disks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested disks.
func (c *disks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("disks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a disk and creates it.  Returns the server's representation of the disk, and an error, if there is any.
func (c *disks) Create(disk *v1alpha1.Disk) (result *v1alpha1.Disk, err error) {
	result = &v1alpha1.Disk{}
	err = c.client.Post().
		Resource("disks").
		Body(disk).
		Do().
		Into(result)
	return
}

// Update takes the representation of a disk and updates it. Returns the server's representation of the disk, and an error, if there is any.
func (c *disks) Update(disk *v1alpha1.Disk) (result *v1alpha1.Disk, err error) {
	result = &v1alpha1.Disk{}
	err = c.client.Put().
		Resource("disks").
		Name(disk.Name).
		Body(disk).
		Do().
		Into(result)
	return
}

// Delete takes name of the disk and deletes it. Returns an error if one occurs.
func (c *disks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("disks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *disks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("disks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched disk.
func (c *disks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Disk, err error) {
	result = &v1alpha1.Disk{}
	err = c.client.Patch(pt).
		Resource("disks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
