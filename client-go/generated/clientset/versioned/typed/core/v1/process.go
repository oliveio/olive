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
	context "context"

	corev1 "github.com/olive-io/olive/apis/core/v1"
	applyconfigurationcorev1 "github.com/olive-io/olive/client-go/generated/applyconfiguration/core/v1"
	scheme "github.com/olive-io/olive/client-go/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ProcessesGetter has a method to return a ProcessInterface.
// A group's client should implement this interface.
type ProcessesGetter interface {
	Processes(namespace string) ProcessInterface
}

// ProcessInterface has methods to work with Process resources.
type ProcessInterface interface {
	Create(ctx context.Context, process *corev1.Process, opts metav1.CreateOptions) (*corev1.Process, error)
	Update(ctx context.Context, process *corev1.Process, opts metav1.UpdateOptions) (*corev1.Process, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, process *corev1.Process, opts metav1.UpdateOptions) (*corev1.Process, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.Process, error)
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ProcessList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *corev1.Process, err error)
	Apply(ctx context.Context, process *applyconfigurationcorev1.ProcessApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.Process, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, process *applyconfigurationcorev1.ProcessApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.Process, err error)
	ProcessExpansion
}

// processes implements ProcessInterface
type processes struct {
	*gentype.ClientWithListAndApply[*corev1.Process, *corev1.ProcessList, *applyconfigurationcorev1.ProcessApplyConfiguration]
}

// newProcesses returns a Processes
func newProcesses(c *CoreV1Client, namespace string) *processes {
	return &processes{
		gentype.NewClientWithListAndApply[*corev1.Process, *corev1.ProcessList, *applyconfigurationcorev1.ProcessApplyConfiguration](
			"processes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *corev1.Process { return &corev1.Process{} },
			func() *corev1.ProcessList { return &corev1.ProcessList{} },
		),
	}
}
