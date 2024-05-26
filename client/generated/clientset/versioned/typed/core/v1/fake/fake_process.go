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

	v1 "github.com/olive-io/olive/apis/core/v1"
	corev1 "github.com/olive-io/olive/client/generated/applyconfiguration/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProcesses implements ProcessInterface
type FakeProcesses struct {
	Fake *FakeOliveV1
	ns   string
}

var processesResource = v1.SchemeGroupVersion.WithResource("processes")

var processesKind = v1.SchemeGroupVersion.WithKind("Process")

// Get takes name of the process, and returns the corresponding process object, and an error if there is any.
func (c *FakeProcesses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Process, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(processesResource, c.ns, name), &v1.Process{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Process), err
}

// List takes label and field selectors, and returns the list of Processes that match those selectors.
func (c *FakeProcesses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ProcessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(processesResource, processesKind, c.ns, opts), &v1.ProcessList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ProcessList{ListMeta: obj.(*v1.ProcessList).ListMeta}
	for _, item := range obj.(*v1.ProcessList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested processes.
func (c *FakeProcesses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(processesResource, c.ns, opts))

}

// Create takes the representation of a process and creates it.  Returns the server's representation of the process, and an error, if there is any.
func (c *FakeProcesses) Create(ctx context.Context, process *v1.Process, opts metav1.CreateOptions) (result *v1.Process, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(processesResource, c.ns, process), &v1.Process{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Process), err
}

// Update takes the representation of a process and updates it. Returns the server's representation of the process, and an error, if there is any.
func (c *FakeProcesses) Update(ctx context.Context, process *v1.Process, opts metav1.UpdateOptions) (result *v1.Process, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(processesResource, c.ns, process), &v1.Process{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Process), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProcesses) UpdateStatus(ctx context.Context, process *v1.Process, opts metav1.UpdateOptions) (*v1.Process, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(processesResource, "status", c.ns, process), &v1.Process{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Process), err
}

// Delete takes name of the process and deletes it. Returns an error if one occurs.
func (c *FakeProcesses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(processesResource, c.ns, name, opts), &v1.Process{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProcesses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(processesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ProcessList{})
	return err
}

// Patch applies the patch and returns the patched process.
func (c *FakeProcesses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Process, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(processesResource, c.ns, name, pt, data, subresources...), &v1.Process{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Process), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied process.
func (c *FakeProcesses) Apply(ctx context.Context, process *corev1.ProcessApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Process, err error) {
	if process == nil {
		return nil, fmt.Errorf("process provided to Apply must not be nil")
	}
	data, err := json.Marshal(process)
	if err != nil {
		return nil, err
	}
	name := process.Name
	if name == nil {
		return nil, fmt.Errorf("process.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(processesResource, c.ns, *name, types.ApplyPatchType, data), &v1.Process{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Process), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeProcesses) ApplyStatus(ctx context.Context, process *corev1.ProcessApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Process, err error) {
	if process == nil {
		return nil, fmt.Errorf("process provided to Apply must not be nil")
	}
	data, err := json.Marshal(process)
	if err != nil {
		return nil, err
	}
	name := process.Name
	if name == nil {
		return nil, fmt.Errorf("process.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(processesResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.Process{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Process), err
}
