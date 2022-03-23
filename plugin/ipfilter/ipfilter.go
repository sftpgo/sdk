// Package ipfilter defines the interface and the GRPC implementation for IP filter plugins.
// IP filter plugins allow to allow/deny access based on the client IP.
package ipfilter

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/sftpgo/sdk/plugin/ipfilter/proto"
)

const (
	// PluginName defines the name for an ipfilter plugin
	PluginName = "ipfilter"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SFTPGO_PLUGIN_IPFILTER",
	MagicCookieValue: "368dbdab-4839-4bec-b75f-fc74a2d50f24",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	PluginName: &Plugin{},
}

// Filter defines the interface for ipfilter plugins
type Filter interface {
	CheckIP(ip string) error
}

// Plugin defines the implementation to serve/connect to a notifier plugin
type Plugin struct {
	plugin.Plugin
	Impl Filter
}

// GRPCServer defines the GRPC server implementation for this plugin
func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterIPFilterServer(s, &GRPCServer{
		Impl: p.Impl,
	})
	return nil
}

// GRPCClient defines the GRPC client implementation for this plugin
func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewIPFilterClient(c),
	}, nil
}
