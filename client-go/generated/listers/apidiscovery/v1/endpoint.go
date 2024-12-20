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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	apidiscoveryv1 "github.com/olive-io/olive/apis/apidiscovery/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// EndpointLister helps list Endpoints.
// All objects returned here must be treated as read-only.
type EndpointLister interface {
	// List lists all Endpoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apidiscoveryv1.Endpoint, err error)
	// Endpoints returns an object that can list and get Endpoints.
	Endpoints(namespace string) EndpointNamespaceLister
	EndpointListerExpansion
}

// endpointLister implements the EndpointLister interface.
type endpointLister struct {
	listers.ResourceIndexer[*apidiscoveryv1.Endpoint]
}

// NewEndpointLister returns a new EndpointLister.
func NewEndpointLister(indexer cache.Indexer) EndpointLister {
	return &endpointLister{listers.New[*apidiscoveryv1.Endpoint](indexer, apidiscoveryv1.Resource("endpoint"))}
}

// Endpoints returns an object that can list and get Endpoints.
func (s *endpointLister) Endpoints(namespace string) EndpointNamespaceLister {
	return endpointNamespaceLister{listers.NewNamespaced[*apidiscoveryv1.Endpoint](s.ResourceIndexer, namespace)}
}

// EndpointNamespaceLister helps list and get Endpoints.
// All objects returned here must be treated as read-only.
type EndpointNamespaceLister interface {
	// List lists all Endpoints in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apidiscoveryv1.Endpoint, err error)
	// Get retrieves the Endpoint from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*apidiscoveryv1.Endpoint, error)
	EndpointNamespaceListerExpansion
}

// endpointNamespaceLister implements the EndpointNamespaceLister
// interface.
type endpointNamespaceLister struct {
	listers.ResourceIndexer[*apidiscoveryv1.Endpoint]
}
