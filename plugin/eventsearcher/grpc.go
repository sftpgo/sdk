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

package eventsearcher

import (
	"context"
	"time"

	"github.com/sftpgo/sdk/plugin/eventsearcher/proto"
)

const (
	rpcTimeout = 20 * time.Second
)

// GRPCClient is an implementation of Notifier interface that talks over RPC.
type GRPCClient struct {
	client proto.SearcherClient
}

// SearchFsEvents implements the Searcher interface
func (c *GRPCClient) SearchFsEvents(searchFilters *FsEventSearch) ([]byte, []string, []string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	resp, err := c.client.SearchFsEvents(ctx, &proto.FsEventsFilter{
		StartTimestamp: searchFilters.StartTimestamp,
		EndTimestamp:   searchFilters.EndTimestamp,
		Actions:        searchFilters.Actions,
		Username:       searchFilters.Username,
		Ip:             searchFilters.IP,
		SshCmd:         searchFilters.SSHCmd,
		Protocols:      searchFilters.Protocols,
		InstanceIds:    searchFilters.InstanceIDs,
		Statuses:       searchFilters.Statuses,
		Limit:          int32(searchFilters.Limit),
		ExcludeIds:     searchFilters.ExcludeIDs,
		FsProvider:     int32(searchFilters.FsProvider),
		Bucket:         searchFilters.Bucket,
		Endpoint:       searchFilters.Endpoint,
		Order:          proto.FsEventsFilter_Order(searchFilters.Order),
		Role:           searchFilters.Role,
	})

	if err != nil {
		return nil, nil, nil, err
	}
	return resp.Data, resp.SameTsAtStart, resp.SameTsAtEnd, nil
}

// SearchProviderEvents implements the Searcher interface
func (c *GRPCClient) SearchProviderEvents(searchFilters *ProviderEventSearch) ([]byte, []string, []string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	resp, err := c.client.SearchProviderEvents(ctx, &proto.ProviderEventsFilter{
		StartTimestamp: searchFilters.StartTimestamp,
		EndTimestamp:   searchFilters.EndTimestamp,
		Actions:        searchFilters.Actions,
		Username:       searchFilters.Username,
		Ip:             searchFilters.IP,
		ObjectTypes:    searchFilters.ObjectTypes,
		ObjectName:     searchFilters.ObjectName,
		InstanceIds:    searchFilters.InstanceIDs,
		Limit:          int32(searchFilters.Limit),
		ExcludeIds:     searchFilters.ExcludeIDs,
		Order:          proto.ProviderEventsFilter_Order(searchFilters.Order),
		Role:           searchFilters.Role,
		OmitObjectData: searchFilters.OmitObjectData,
	})

	if err != nil {
		return nil, nil, nil, err
	}
	return resp.Data, resp.SameTsAtStart, resp.SameTsAtEnd, nil
}

// GRPCServer defines the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl Searcher
}

// SearchFsEvents implements the server side fs events search method
func (s *GRPCServer) SearchFsEvents(ctx context.Context, req *proto.FsEventsFilter) (*proto.SearchResponse, error) {
	responseData, sameTsAtStart, sameTsAtEnd, err := s.Impl.SearchFsEvents(&FsEventSearch{
		CommonSearchParams: CommonSearchParams{
			StartTimestamp: req.StartTimestamp,
			EndTimestamp:   req.EndTimestamp,
			Actions:        req.Actions,
			Username:       req.Username,
			IP:             req.Ip,
			InstanceIDs:    req.InstanceIds,
			Limit:          int(req.Limit),
			ExcludeIDs:     req.ExcludeIds,
			Order:          int(req.Order),
			Role:           req.Role,
		},

		SSHCmd:     req.SshCmd,
		Protocols:  req.Protocols,
		Statuses:   req.Statuses,
		FsProvider: int(req.FsProvider),
		Bucket:     req.Bucket,
		Endpoint:   req.Endpoint,
	})

	return &proto.SearchResponse{
		Data:          responseData,
		SameTsAtStart: sameTsAtStart,
		SameTsAtEnd:   sameTsAtEnd,
	}, err
}

// SearchProviderEvents implement the server side provider events search method
func (s *GRPCServer) SearchProviderEvents(ctx context.Context, req *proto.ProviderEventsFilter) (*proto.SearchResponse, error) {
	responseData, sameTsAtStart, sameTsAtEnd, err := s.Impl.SearchProviderEvents(&ProviderEventSearch{
		CommonSearchParams: CommonSearchParams{
			StartTimestamp: req.StartTimestamp,
			EndTimestamp:   req.EndTimestamp,
			Actions:        req.Actions,
			Username:       req.Username,
			IP:             req.Ip,
			InstanceIDs:    req.InstanceIds,
			Limit:          int(req.Limit),
			ExcludeIDs:     req.ExcludeIds,
			Order:          int(req.Order),
			Role:           req.Role,
		},
		ObjectTypes:    req.ObjectTypes,
		ObjectName:     req.ObjectName,
		OmitObjectData: req.OmitObjectData,
	})

	return &proto.SearchResponse{
		Data:          responseData,
		SameTsAtStart: sameTsAtStart,
		SameTsAtEnd:   sameTsAtEnd,
	}, err
}
