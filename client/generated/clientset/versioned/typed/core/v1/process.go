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

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/olive-io/olive/apis/core/v1"
	corev1 "github.com/olive-io/olive/client/generated/applyconfiguration/core/v1"
	scheme "github.com/olive-io/olive/client/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProcessesGetter has a method to return a ProcessInterface.
// A group's client should implement this interface.
type ProcessesGetter interface {
	Processes(namespace string) ProcessInterface
}

// ProcessInterface has methods to work with Process resources.
type ProcessInterface interface {
	Create(ctx context.Context, process *v1.Process, opts metav1.CreateOptions) (*v1.Process, error)
	Update(ctx context.Context, process *v1.Process, opts metav1.UpdateOptions) (*v1.Process, error)
	UpdateStatus(ctx context.Context, process *v1.Process, opts metav1.UpdateOptions) (*v1.Process, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Process, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ProcessList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Process, err error)
	Apply(ctx context.Context, process *corev1.ProcessApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Process, err error)
	ApplyStatus(ctx context.Context, process *corev1.ProcessApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Process, err error)
	ProcessExpansion
}

// processes implements ProcessInterface
type processes struct {
	client rest.Interface
	ns     string
}

// newProcesses returns a Processes
func newProcesses(c *OliveV1Client, namespace string) *processes {
	return &processes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the process, and returns the corresponding process object, and an error if there is any.
func (c *processes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Process, err error) {
	result = &v1.Process{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("processes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Processes that match those selectors.
func (c *processes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ProcessList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ProcessList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("processes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested processes.
func (c *processes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("processes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a process and creates it.  Returns the server's representation of the process, and an error, if there is any.
func (c *processes) Create(ctx context.Context, process *v1.Process, opts metav1.CreateOptions) (result *v1.Process, err error) {
	result = &v1.Process{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("processes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(process).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a process and updates it. Returns the server's representation of the process, and an error, if there is any.
func (c *processes) Update(ctx context.Context, process *v1.Process, opts metav1.UpdateOptions) (result *v1.Process, err error) {
	result = &v1.Process{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("processes").
		Name(process.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(process).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *processes) UpdateStatus(ctx context.Context, process *v1.Process, opts metav1.UpdateOptions) (result *v1.Process, err error) {
	result = &v1.Process{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("processes").
		Name(process.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(process).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the process and deletes it. Returns an error if one occurs.
func (c *processes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("processes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *processes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("processes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched process.
func (c *processes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Process, err error) {
	result = &v1.Process{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("processes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied process.
func (c *processes) Apply(ctx context.Context, process *corev1.ProcessApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Process, err error) {
	if process == nil {
		return nil, fmt.Errorf("process provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(process)
	if err != nil {
		return nil, err
	}
	name := process.Name
	if name == nil {
		return nil, fmt.Errorf("process.Name must be provided to Apply")
	}
	result = &v1.Process{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("processes").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *processes) ApplyStatus(ctx context.Context, process *corev1.ProcessApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Process, err error) {
	if process == nil {
		return nil, fmt.Errorf("process provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(process)
	if err != nil {
		return nil, err
	}

	name := process.Name
	if name == nil {
		return nil, fmt.Errorf("process.Name must be provided to Apply")
	}

	result = &v1.Process{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("processes").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
