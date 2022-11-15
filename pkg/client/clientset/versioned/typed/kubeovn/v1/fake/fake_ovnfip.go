/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kubeovnv1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOvnFips implements OvnFipInterface
type FakeOvnFips struct {
	Fake *FakeKubeovnV1
}

var ovnfipsResource = schema.GroupVersionResource{Group: "kubeovn.io", Version: "v1", Resource: "ovn-fips"}

var ovnfipsKind = schema.GroupVersionKind{Group: "kubeovn.io", Version: "v1", Kind: "OvnFip"}

// Get takes name of the ovnFip, and returns the corresponding ovnFip object, and an error if there is any.
func (c *FakeOvnFips) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubeovnv1.OvnFip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ovnfipsResource, name), &kubeovnv1.OvnFip{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubeovnv1.OvnFip), err
}

// List takes label and field selectors, and returns the list of OvnFips that match those selectors.
func (c *FakeOvnFips) List(ctx context.Context, opts v1.ListOptions) (result *kubeovnv1.OvnFipList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ovnfipsResource, ovnfipsKind, opts), &kubeovnv1.OvnFipList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubeovnv1.OvnFipList{ListMeta: obj.(*kubeovnv1.OvnFipList).ListMeta}
	for _, item := range obj.(*kubeovnv1.OvnFipList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ovnFips.
func (c *FakeOvnFips) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ovnfipsResource, opts))
}

// Create takes the representation of a ovnFip and creates it.  Returns the server's representation of the ovnFip, and an error, if there is any.
func (c *FakeOvnFips) Create(ctx context.Context, ovnFip *kubeovnv1.OvnFip, opts v1.CreateOptions) (result *kubeovnv1.OvnFip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ovnfipsResource, ovnFip), &kubeovnv1.OvnFip{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubeovnv1.OvnFip), err
}

// Update takes the representation of a ovnFip and updates it. Returns the server's representation of the ovnFip, and an error, if there is any.
func (c *FakeOvnFips) Update(ctx context.Context, ovnFip *kubeovnv1.OvnFip, opts v1.UpdateOptions) (result *kubeovnv1.OvnFip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ovnfipsResource, ovnFip), &kubeovnv1.OvnFip{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubeovnv1.OvnFip), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOvnFips) UpdateStatus(ctx context.Context, ovnFip *kubeovnv1.OvnFip, opts v1.UpdateOptions) (*kubeovnv1.OvnFip, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(ovnfipsResource, "status", ovnFip), &kubeovnv1.OvnFip{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubeovnv1.OvnFip), err
}

// Delete takes name of the ovnFip and deletes it. Returns an error if one occurs.
func (c *FakeOvnFips) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ovnfipsResource, name, opts), &kubeovnv1.OvnFip{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOvnFips) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ovnfipsResource, listOpts)

	_, err := c.Fake.Invokes(action, &kubeovnv1.OvnFipList{})
	return err
}

// Patch applies the patch and returns the patched ovnFip.
func (c *FakeOvnFips) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubeovnv1.OvnFip, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ovnfipsResource, name, pt, data, subresources...), &kubeovnv1.OvnFip{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubeovnv1.OvnFip), err
}