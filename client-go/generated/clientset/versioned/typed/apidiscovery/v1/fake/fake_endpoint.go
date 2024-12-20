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
	v1 "github.com/olive-io/olive/apis/apidiscovery/v1"
	apidiscoveryv1 "github.com/olive-io/olive/client-go/generated/applyconfiguration/apidiscovery/v1"
	typedapidiscoveryv1 "github.com/olive-io/olive/client-go/generated/clientset/versioned/typed/apidiscovery/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeEndpoints implements EndpointInterface
type fakeEndpoints struct {
	*gentype.FakeClientWithListAndApply[*v1.Endpoint, *v1.EndpointList, *apidiscoveryv1.EndpointApplyConfiguration]
	Fake *FakeApidiscoveryV1
}

func newFakeEndpoints(fake *FakeApidiscoveryV1, namespace string) typedapidiscoveryv1.EndpointInterface {
	return &fakeEndpoints{
		gentype.NewFakeClientWithListAndApply[*v1.Endpoint, *v1.EndpointList, *apidiscoveryv1.EndpointApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("endpoints"),
			v1.SchemeGroupVersion.WithKind("Endpoint"),
			func() *v1.Endpoint { return &v1.Endpoint{} },
			func() *v1.EndpointList { return &v1.EndpointList{} },
			func(dst, src *v1.EndpointList) { dst.ListMeta = src.ListMeta },
			func(list *v1.EndpointList) []*v1.Endpoint { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.EndpointList, items []*v1.Endpoint) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
