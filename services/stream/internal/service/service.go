package service

import (
	"context"
	v1 "github.com/belo4ya/live-streaming-service/api/stream/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	v1.UnimplementedStreamServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ListStreams(_ context.Context, _ *emptypb.Empty) (*v1.ListStreamsResponse, error) {
	results := []*v1.ListStreamsResponse_Stream{
		{
			Id:       1,
			Name:     "Stream 1",
			Username: "Streamer 1",
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
	return &v1.ListStreamsResponse{Results: results}, nil
}
