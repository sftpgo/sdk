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
func (c *GRPCClient) CheckIP(ip string) error {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	_, err := c.client.CheckIP(ctx, &proto.CheckIPRequest{
		Ip: ip,
	})
	return err
}

// GRPCServer defines the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl Filter
}

// CheckIP implements the server side check IP method
func (s *GRPCServer) CheckIP(ctx context.Context, req *proto.CheckIPRequest) (*emptypb.Empty, error) {
	err := s.Impl.CheckIP(req.Ip)

	return &emptypb.Empty{}, err
}
