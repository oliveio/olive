// Copyright 2023 The olive Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-olive. DO NOT EDIT.
// versions:
// - protoc-gen-go-olive v0.1.0
// - protoc             v4.25.3
// source: github.com/olive-io/olive/api/gatewaypb/rpc.proto

package gatewaypb

import (
	server "github.com/olive-io/olive/pkg/proxy/server"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

func RegisterGatewayServerHandler(s server.IServer, srv GatewayServer) {
	handler := s.NewHandler(srv, server.WithServerDesc(&Gateway_ServiceDesc))
	s.Handle(handler)
}

func RegisterTestServiceServerHandler(s server.IServer, srv TestServiceServer) {
	handler := s.NewHandler(srv, server.WithServerDesc(&TestService_ServiceDesc))
	s.Handle(handler)
}
