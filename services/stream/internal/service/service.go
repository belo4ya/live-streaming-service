package service

import (
	"context"
	v1 "github.com/belo4ya/live-streaming-service/api/stream/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewService)

type Service struct {
	v1.UnimplementedStreamServiceServer
	log *log.Helper
}

func NewService(l log.Logger) *Service {
	return &Service{log: log.NewHelper(log.With(l, "module", "service/service"))}
}

func (s *Service) ListStreams(_ context.Context, _ *empty.Empty) (*v1.ListStreamsResponse, error) {
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
