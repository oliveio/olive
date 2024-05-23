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
	v1 "github.com/olive-io/olive/client/generated/clientset/versioned/typed/apidiscovery/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeApidiscoveryV1 struct {
	*testing.Fake
}

func (c *FakeApidiscoveryV1) Consumers(namespace string) v1.ConsumerInterface {
	return &FakeConsumers{c, namespace}
}

func (c *FakeApidiscoveryV1) ConsumerLists(namespace string) v1.ConsumerListInterface {
	return &FakeConsumerLists{c, namespace}
}

func (c *FakeApidiscoveryV1) Endpoints(namespace string) v1.EndpointInterface {
	return &FakeEndpoints{c, namespace}
}

func (c *FakeApidiscoveryV1) EndpointLists(namespace string) v1.EndpointListInterface {
	return &FakeEndpointLists{c, namespace}
}

func (c *FakeApidiscoveryV1) Nodes(namespace string) v1.NodeInterface {
	return &FakeNodes{c, namespace}
}

func (c *FakeApidiscoveryV1) NodeLists(namespace string) v1.NodeListInterface {
	return &FakeNodeLists{c, namespace}
}

func (c *FakeApidiscoveryV1) Services(namespace string) v1.ServiceInterface {
	return &FakeServices{c, namespace}
}

func (c *FakeApidiscoveryV1) ServiceLists(namespace string) v1.ServiceListInterface {
	return &FakeServiceLists{c, namespace}
}

func (c *FakeApidiscoveryV1) Yards(namespace string) v1.YardInterface {
	return &FakeYards{c, namespace}
}

func (c *FakeApidiscoveryV1) YardLists(namespace string) v1.YardListInterface {
	return &FakeYardLists{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeApidiscoveryV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
