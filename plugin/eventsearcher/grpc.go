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
func (c *GRPCClient) SearchFsEvents(searchFilters *FsEventSearch) ([]byte, error) {
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
		FromId:         searchFilters.FromID,
		FsProvider:     int32(searchFilters.FsProvider),
		Bucket:         searchFilters.Bucket,
		Endpoint:       searchFilters.Endpoint,
		Order:          proto.FsEventsFilter_Order(searchFilters.Order),
		Role:           searchFilters.Role,
	})

	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// SearchProviderEvents implements the Searcher interface
func (c *GRPCClient) SearchProviderEvents(searchFilters *ProviderEventSearch) ([]byte, error) {
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
		FromId:         searchFilters.FromID,
		Order:          proto.ProviderEventsFilter_Order(searchFilters.Order),
		Role:           searchFilters.Role,
		OmitObjectData: searchFilters.OmitObjectData,
	})

	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// SearchLogEvents implements the Searcher interface
func (c *GRPCClient) SearchLogEvents(searchFilters *LogEventSearch) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	resp, err := c.client.SearchLogEvents(ctx, &proto.LogEventsFilter{
		StartTimestamp: searchFilters.StartTimestamp,
		EndTimestamp:   searchFilters.EndTimestamp,
		Events:         searchFilters.Events,
		Username:       searchFilters.Username,
		Ip:             searchFilters.IP,
		Protocols:      searchFilters.Protocols,
		InstanceIds:    searchFilters.InstanceIDs,
		Limit:          int32(searchFilters.Limit),
		FromId:         searchFilters.FromID,
		Order:          proto.LogEventsFilter_Order(searchFilters.Order),
		Role:           searchFilters.Role,
	})

	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// GRPCServer defines the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	Impl Searcher
}

// SearchFsEvents implements the server side fs events search method
func (s *GRPCServer) SearchFsEvents(ctx context.Context, req *proto.FsEventsFilter) (*proto.SearchResponse, error) {
	responseData, err := s.Impl.SearchFsEvents(&FsEventSearch{
		CommonSearchParams: CommonSearchParams{
			StartTimestamp: req.StartTimestamp,
			EndTimestamp:   req.EndTimestamp,
			Username:       req.Username,
			IP:             req.Ip,
			InstanceIDs:    req.InstanceIds,
			Limit:          int(req.Limit),
			FromID:         req.FromId,
			Order:          int(req.Order),
			Role:           req.Role,
		},
		Actions:    req.Actions,
		SSHCmd:     req.SshCmd,
		Protocols:  req.Protocols,
		Statuses:   req.Statuses,
		FsProvider: int(req.FsProvider),
		Bucket:     req.Bucket,
		Endpoint:   req.Endpoint,
	})

	return &proto.SearchResponse{
		Data: responseData,
	}, err
}

// SearchProviderEvents implement the server side provider events search method
func (s *GRPCServer) SearchProviderEvents(ctx context.Context, req *proto.ProviderEventsFilter) (*proto.SearchResponse, error) {
	responseData, err := s.Impl.SearchProviderEvents(&ProviderEventSearch{
		CommonSearchParams: CommonSearchParams{
			StartTimestamp: req.StartTimestamp,
			EndTimestamp:   req.EndTimestamp,
			Username:       req.Username,
			IP:             req.Ip,
			InstanceIDs:    req.InstanceIds,
			Limit:          int(req.Limit),
			FromID:         req.FromId,
			Order:          int(req.Order),
			Role:           req.Role,
		},
		Actions:        req.Actions,
		ObjectTypes:    req.ObjectTypes,
		ObjectName:     req.ObjectName,
		OmitObjectData: req.OmitObjectData,
	})

	return &proto.SearchResponse{
		Data: responseData,
	}, err
}

// SearchLogEvents implement the server side log events search method
func (s *GRPCServer) SearchLogEvents(ctx context.Context, req *proto.LogEventsFilter) (*proto.SearchResponse, error) {
	responseData, err := s.Impl.SearchLogEvents(&LogEventSearch{
		CommonSearchParams: CommonSearchParams{
			StartTimestamp: req.StartTimestamp,
			EndTimestamp:   req.EndTimestamp,
			Username:       req.Username,
			IP:             req.Ip,
			InstanceIDs:    req.InstanceIds,
			Limit:          int(req.Limit),
			FromID:         req.FromId,
			Order:          int(req.Order),
			Role:           req.Role,
		},
		Events:    req.Events,
		Protocols: req.Protocols,
	})

	return &proto.SearchResponse{
		Data: responseData,
	}, err
}
