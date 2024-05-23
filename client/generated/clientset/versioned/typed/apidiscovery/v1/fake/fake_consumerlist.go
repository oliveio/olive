/*
   Copyright 2024 The olive Authors

   This program is offered under a commercial and under the AGPL license.
   For AGPL licensing, see below.

   AGPL licensing:
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/olive-io/olive/apis/apidiscovery/v1"
	apidiscoveryv1 "github.com/olive-io/olive/client/generated/applyconfiguration/apidiscovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConsumerLists implements ConsumerListInterface
type FakeConsumerLists struct {
	Fake *FakeApidiscoveryV1
	ns   string
}

var consumerlistsResource = v1.SchemeGroupVersion.WithResource("consumerlists")

var consumerlistsKind = v1.SchemeGroupVersion.WithKind("ConsumerList")

// Get takes name of the consumerList, and returns the corresponding consumerList object, and an error if there is any.
func (c *FakeConsumerLists) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ConsumerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(consumerlistsResource, c.ns, name), &v1.ConsumerList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsumerList), err
}

// List takes label and field selectors, and returns the list of ConsumerLists that match those selectors.
func (c *FakeConsumerLists) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ConsumerListList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(consumerlistsResource, consumerlistsKind, c.ns, opts), &v1.ConsumerListList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsumerListList), err
}

// Watch returns a watch.Interface that watches the requested consumerLists.
func (c *FakeConsumerLists) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(consumerlistsResource, c.ns, opts))

}

// Create takes the representation of a consumerList and creates it.  Returns the server's representation of the consumerList, and an error, if there is any.
func (c *FakeConsumerLists) Create(ctx context.Context, consumerList *v1.ConsumerList, opts metav1.CreateOptions) (result *v1.ConsumerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(consumerlistsResource, c.ns, consumerList), &v1.ConsumerList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsumerList), err
}

// Update takes the representation of a consumerList and updates it. Returns the server's representation of the consumerList, and an error, if there is any.
func (c *FakeConsumerLists) Update(ctx context.Context, consumerList *v1.ConsumerList, opts metav1.UpdateOptions) (result *v1.ConsumerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(consumerlistsResource, c.ns, consumerList), &v1.ConsumerList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsumerList), err
}

// Delete takes name of the consumerList and deletes it. Returns an error if one occurs.
func (c *FakeConsumerLists) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(consumerlistsResource, c.ns, name, opts), &v1.ConsumerList{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConsumerLists) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(consumerlistsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ConsumerListList{})
	return err
}

// Patch applies the patch and returns the patched consumerList.
func (c *FakeConsumerLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConsumerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(consumerlistsResource, c.ns, name, pt, data, subresources...), &v1.ConsumerList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsumerList), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied consumerList.
func (c *FakeConsumerLists) Apply(ctx context.Context, consumerList *apidiscoveryv1.ConsumerListApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ConsumerList, err error) {
	if consumerList == nil {
		return nil, fmt.Errorf("consumerList provided to Apply must not be nil")
	}
	data, err := json.Marshal(consumerList)
	if err != nil {
		return nil, err
	}
	name := consumerList.Name
	if name == nil {
		return nil, fmt.Errorf("consumerList.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(consumerlistsResource, c.ns, *name, types.ApplyPatchType, data), &v1.ConsumerList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ConsumerList), err
}
