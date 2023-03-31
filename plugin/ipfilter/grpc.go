// Copyright (C) 2022-2023 Nicola Murino
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

package ipfilter

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sftpgo/sdk/plugin/ipfilter/proto"
)

const (
	rpcTimeout = 20 * time.Second
)

// GRPCClient is an implementation of IPFilter interface that talks over RPC.
type GRPCClient struct {
	client proto.IPFilterClient
}

// CheckIP returns an error if the specified IP is not allowed
func (c *GRPCClient) CheckIP(ip, protocol string) error {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	_, err := c.client.CheckIP(ctx, &proto.CheckIPRequest{
		Ip:       ip,
		Protocol: protocol,
	})
	return err
}

// Reload instructs the filter to reload its configuration
func (c *GRPCClient) Reload() error {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	_, err := c.client.Reload(ctx, &emptypb.Empty{})
	return err
}

// GRPCServer defines the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl Filter
}

// CheckIP implements the server side check IP method
func (s *GRPCServer) CheckIP(ctx context.Context, req *proto.CheckIPRequest) (*emptypb.Empty, error) {
	err := s.Impl.CheckIP(req.Ip, req.Protocol)

	return &emptypb.Empty{}, err
}

// Reload implements the server side reload method
func (s *GRPCServer) Reload(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	err := s.Impl.Reload()

	return &emptypb.Empty{}, err
}
