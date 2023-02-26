package service

import (
	"context"
	v1 "github.com/belo4ya/live-streaming-service/api/gateway/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type VODService struct {
	v1.UnimplementedVODServiceServer
}

func NewVODService() *VODService {
	return &VODService{}
}

func (s *VODService) ListVODs(_ context.Context, _ *empty.Empty) (*v1.ListVODsResponse, error) {
	results := []*v1.ListVODsResponse_VOD{
		{
			Id:        1,
			Name:      "Stream 1",
			Username:  "Streamer 1",
			CreatedAt: &timestamp.Timestamp{},
		},
		{
			Id:        2,
			Name:      "Stream 2",
			Username:  "Streamer 2",
			CreatedAt: &timestamp.Timestamp{},
		},
		{
			Id:        3,
			Name:      "Stream 3",
			Username:  "Streamer 3",
			CreatedAt: &timestamp.Timestamp{},
		},
	}
	return &v1.ListVODsResponse{Results: results}, nil
}
