package service

import (
	"context"
	v1 "github.com/belo4ya/live-streaming-service/api/vod/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Service struct {
	v1.UnimplementedVODServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ListVODs(_ context.Context, _ *empty.Empty) (*v1.ListVODsResponse, error) {
	results := []*v1.ListVODsResponse_VOD{
		{
			Id:        1,
			Name:      "Stream 1",
			Username:  "Streamer 1",
			CreatedAt: &timestamp.Timestamp{},
		},
		{
			Id:       2,
			Name:     "Stream 2",
			Username: "Streamer 2",
		},
		{
			Id:       3,
			Name:     "Stream 3",
			Username: "Streamer 3",
		},
	}
	return &v1.ListVODsResponse{Results: results}, nil
}
