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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1 "github.com/olive-io/olive/apis/apidiscovery/v1"
	corev1 "github.com/olive-io/olive/apis/core/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=apidiscovery.olive.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("edges"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apidiscovery().V1().Edges().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("endpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apidiscovery().V1().Endpoints().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apidiscovery().V1().Services().Informer()}, nil

		// Group=core.olive.io, Version=v1
	case corev1.SchemeGroupVersion.WithResource("definitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Definitions().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("namespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Namespaces().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("processes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Processes().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("regions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Regions().Informer()}, nil
	case corev1.SchemeGroupVersion.WithResource("runners"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1().Runners().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
