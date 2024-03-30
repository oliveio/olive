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

package server

import (
	"context"

	"go.uber.org/zap"

	pb "github.com/olive-io/olive/api/discoverypb"
)

type gatewayRpc struct {
	pb.UnsafeGatewayServer

	gw *Gateway
}

func (gr *gatewayRpc) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Reply: req.Message}, nil
}

func (gr *gatewayRpc) Transmit(ctx context.Context, req *pb.TransmitRequest) (*pb.TransmitResponse, error) {
	lg := gr.gw.Logger()
	lg.Info("transmit executed", zap.String("activity", req.Activity.String()))
	for key, value := range req.Properties {
		lg.Sugar().Infof("%s=%+v", key, value.Value())
	}
	resp := &pb.TransmitResponse{
		Properties:  map[string]*pb.Box{},
		DataObjects: map[string]*pb.Box{},
	}
	return resp, nil
}

func (gr *gatewayRpc) ClientStream(server pb.Gateway_ClientStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (gr *gatewayRpc) ServerStream(request *pb.PingRequest, server pb.Gateway_ServerStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (gr *gatewayRpc) BinaryStream(server pb.Gateway_BinaryStreamServer) error {
	//TODO implement me
	panic("implement me")
}
