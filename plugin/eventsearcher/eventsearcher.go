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

// Package eventsearcher defines the interface and the GRPC implementation for events search plugins.
// Events search plugins allow to search for filesystem and provider events.
package eventsearcher

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/sftpgo/sdk/plugin/eventsearcher/proto"
)

const (
	// PluginName defines the name for an events search plugin
	PluginName = "eventsearcher"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SFTPGO_PLUGIN_EVENTSEARCHER",
	MagicCookieValue: "2b523805-0279-471c-895e-6c0d39002ca4",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	PluginName: &Plugin{},
}

// CommonSearchParams defines common parameters for both filesystem and provider search
type CommonSearchParams struct {
	StartTimestamp int64
	EndTimestamp   int64
	Actions        []string
	Username       string
	IP             string
	InstanceIDs    []string
	ExcludeIDs     []string
	Limit          int
	Order          int
	Role           string
}

// FsEventSearch defines the fields for a filesystem event search
type FsEventSearch struct {
	CommonSearchParams
	SSHCmd     string
	Protocols  []string
	Statuses   []int32
	FsProvider int
	Bucket     string
	Endpoint   string
}

// ProviderEventSearch defines the fields for a provider event search
type ProviderEventSearch struct {
	CommonSearchParams
	ObjectName     string
	ObjectTypes    []string
	OmitObjectData bool
}

// Searcher defines the interface for events search plugins
type Searcher interface {
	SearchFsEvents(searchFilters *FsEventSearch) ([]byte, []string, []string, error)
	SearchProviderEvents(searchFilters *ProviderEventSearch) ([]byte, []string, []string, error)
}

// Plugin defines the implementation to serve/connect to a event search plugin
type Plugin struct {
	plugin.Plugin
	Impl Searcher
}

// GRPCServer defines the GRPC server implementation for this plugin
func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterSearcherServer(s, &GRPCServer{
		Impl: p.Impl,
	})
	return nil
}

// GRPCClient defines the GRPC client implementation for this plugin
func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewSearcherClient(c),
	}, nil
}
